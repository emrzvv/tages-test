package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/emrzvv/tages-test/proto"
)

func main() {
	serverAddr := flag.String("addr", "localhost:50051", "gRPC server address")
	command := flag.String("cmd", "", "Command to run: upload | download | list")
	filePath := flag.String("file", "", "Local file path for upload or filename on server for download")
	dirPath := flag.String("out", "./", "Local directory path for file download")

	flag.Parse()
	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewImageServiceClient(conn)

	switch *command {
	case "upload":
		if *filePath == "" {
			fmt.Println("Please specify --file for upload")
			os.Exit(1)
		}

		if err := doUpload(context.Background(), client, *filePath); err != nil {
			fmt.Printf("Upload error: %v\n", err)
			os.Exit(1)
		}

	case "download":
		if *filePath == "" {
			fmt.Println("Please specify --file for download")
			os.Exit(1)
		}
		if err := doDownload(context.Background(), client, *filePath, *dirPath); err != nil {
			fmt.Printf("Download error: %v\n", err)
			os.Exit(1)
		}

	case "list":
		if err := doList(context.Background(), client); err != nil {
			fmt.Printf("List error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Usage:")
		fmt.Println("  --cmd upload   --file /path/to/localfile")
		fmt.Println("  --cmd download --file filename_on_server --out dir_path")
		fmt.Println("  --cmd list")
		os.Exit(1)
	}
}

func doUpload(ctx context.Context, client pb.ImageServiceClient, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error when opening local file: %w", err)
	}
	defer f.Close()

	stream, err := client.UploadImage(ctx)
	if err != nil {
		return fmt.Errorf("error when starting UploadImage stream: %w", err)
	}

	metaReq := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Meta{
			Meta: &pb.UploadMeta{
				Name: filePath[strings.LastIndex(filePath, "/")+1:],
			},
		},
	}
	if sendErr := stream.Send(metaReq); sendErr != nil {
		return fmt.Errorf("error when sending metadata: %w", sendErr)
	}
	chunkSize := 4096
	buffer := make([]byte, chunkSize)
	for {
		n, readErr := f.Read(buffer)
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return fmt.Errorf("error when reading local file chunk: %w", readErr)
		}

		chunkReq := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_Chunk{
				Chunk: buffer[:n],
			},
		}
		if sendErr := stream.Send(chunkReq); sendErr != nil {
			return fmt.Errorf("error when sending chunk: %w", sendErr)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("error when recieving UploadImage response: %w", err)
	}

	fmt.Println(resp.GetMessage())
	if resp.GetInfo() != nil {
		fmt.Printf("image info: Name=%s, CreatedAt=%s, LastModifiedAt=%s\n",
			resp.Info.GetName(), resp.Info.GetCreatedAt(), resp.Info.GetLastModifiedAt())
	}
	return nil
}

func doDownload(ctx context.Context, client pb.ImageServiceClient, remoteName string, dirPath string) error {
	stream, err := client.DownloadImage(ctx, &pb.DownloadImageRequest{Name: remoteName})
	if err != nil {
		return fmt.Errorf("error when calling DownloadImage: %w", err)
	}

	outFile, err := os.Create(path.Join(dirPath, remoteName))
	if err != nil {
		return fmt.Errorf("error when creating local file: %w", err)
	}
	defer outFile.Close()

	for {
		chunkResp, recvErr := stream.Recv()
		if recvErr == io.EOF {
			break
		}
		if recvErr != nil {
			return fmt.Errorf("error when receiving chunk: %w", recvErr)
		}

		if _, writeErr := outFile.Write(chunkResp.Chunk); writeErr != nil {
			return fmt.Errorf("error when writing chunk to local file: %w", writeErr)
		}
	}

	fmt.Printf("File %s downloaded successfully\n", remoteName)
	return nil
}

func doList(ctx context.Context, client pb.ImageServiceClient) error {
	resp, err := client.GetImagesList(ctx, &pb.GetImagesListRequest{})
	if err != nil {
		return fmt.Errorf("error when calling GetImagesList: %w", err)
	}

	for _, info := range resp.ImageInfo {
		fmt.Println(info)
	}
	return nil
}
