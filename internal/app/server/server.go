package server

import (
	"context"
	"io"
	"time"

	"github.com/emrzvv/tages-test/cfg"
	"github.com/emrzvv/tages-test/internal/app/service"
	pb "github.com/emrzvv/tages-test/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedImageServiceServer
	config     *cfg.Config
	imgService *service.ImgService
}

func NewServer(config *cfg.Config, imgService *service.ImgService) *Server {
	return &Server{
		config:     config,
		imgService: imgService,
	}
}

func (server *Server) UploadImage(stream pb.ImageService_UploadImageServer) error {
	var name string
	var writer io.WriteCloser
	var currentTime time.Time = time.Now()

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "failed to receive stream: %v", err)
		}

		switch data := request.Data.(type) {

		case *pb.UploadImageRequest_Meta:
			if writer != nil {
				return status.Errorf(codes.InvalidArgument, "meta has been already received")
			}
			name = server.imgService.FormatName(data.Meta.GetName())
			w, openErr := server.imgService.CreateFile(name)
			if openErr != nil {
				return status.Errorf(codes.Internal, "error when creating img file: %v", openErr)
			}
			writer = w
			currentTime = time.Now()

		case *pb.UploadImageRequest_Chunk:
			if writer == nil {
				return status.Errorf(codes.InvalidArgument, "no metadata received before streaming")
			}
			if _, writeErr := writer.Write(data.Chunk); writeErr != nil {
				return status.Errorf(codes.Internal, "error when writing file: %v", writeErr)
			}

		default:
			return status.Errorf(codes.InvalidArgument, "invalid request")
		}
	}

	if err := writer.Close(); err != nil {
		return status.Errorf(codes.Internal, "error when closing file writer: %v", err)
	}

	metaErr := server.imgService.SaveMeta(name, currentTime)
	if metaErr != nil {
		return status.Errorf(codes.Internal, "error when saving metadata: %v", metaErr)
	}
	meta, _ := server.imgService.GetMetaByName(name)
	return stream.SendAndClose(&pb.UploadImageResponse{
		Message: "Image uploaded successfully",
		Info: &pb.ImageInfo{
			Name:           meta.Name,
			CreatedAt:      meta.CreatedAt,
			LastModifiedAt: meta.ModifiedAt,
		},
	})
}

func (server *Server) DownloadImage(in *pb.DownloadImageRequest, stream pb.ImageService_DownloadImageServer) error {
	reader, err := server.imgService.GetFile(in.GetName())
	if err != nil {
		return status.Errorf(codes.Internal, "error when reading file: %v", err)
	}
	defer reader.Close()

	buffer := make([]byte, server.config.ChunkSize)
	for {
		n, readErr := reader.Read(buffer)
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return status.Errorf(codes.Internal, "error when reading file to buffer: %v", readErr)
		}

		sendErr := stream.Send(&pb.DownloadImageResponse{
			Chunk: buffer[:n],
		})
		if sendErr != nil {
			return status.Errorf(codes.Internal, "error when sending chunk to stream: %v", sendErr)
		}
	}

	return nil
}

func (server *Server) GetImagesList(ctx context.Context, in *pb.GetImagesListRequest) (*pb.GetImagesListStrResponse, error) {
	list := server.imgService.GetImagesMetaInfoList()
	return &pb.GetImagesListStrResponse{
		ImageInfo: list,
	}, nil
}
