package service

import (
	"io"
	"strings"
	"time"

	"github.com/emrzvv/tages-test/cfg"
	"github.com/emrzvv/tages-test/internal/app/model"
	"github.com/emrzvv/tages-test/internal/app/storage"
)

type ImgService struct {
	config      *cfg.Config
	fileStorage storage.FileStorage
	metaStorage storage.MetaStorage
}

func NewImgService(config *cfg.Config, fs storage.FileStorage, ms storage.MetaStorage) *ImgService {
	return &ImgService{
		config:      config,
		fileStorage: fs,
		metaStorage: ms,
	}
}

func (s *ImgService) FormatName(name string) string {
	// name validation & formatting?
	return strings.TrimSpace(name)
}

func (s *ImgService) SaveMeta(name string, currentTime time.Time) error {
	// currentTimeStr := currentTime.Format(time.RFC3339)
	meta, exists := s.metaStorage.GetMetaByName(name)
	if !exists {
		err := s.metaStorage.InsertMeta(model.MetaData{
			Name:       name,
			CreatedAt:  currentTime,
			ModifiedAt: currentTime,
		})
		return err
	}
	updatedMeta := model.MetaData{
		Name:       meta.Name,
		CreatedAt:  meta.CreatedAt,
		ModifiedAt: currentTime,
	}
	err := s.metaStorage.UpdateMeta(updatedMeta)
	return err
}

func (s *ImgService) FormatTime(toFormat time.Time) string {
	return toFormat.Format(time.RFC3339)
}

func (s *ImgService) GetMetaByName(name string) (model.MetaData, bool) {
	return s.metaStorage.GetMetaByName(name)
}

func (s *ImgService) CreateFile(name string) (io.WriteCloser, error) {
	return s.fileStorage.Save(name)
}

func (s *ImgService) GetFile(name string) (io.ReadCloser, error) {
	return s.fileStorage.Get(name)
}

func (s *ImgService) GetImagesMetaInfoList() []string {
	metaList := s.metaStorage.GetMetaList()
	result := make([]string, 0, len(metaList))
	for _, meta := range metaList {
		result = append(result, meta.String())
	}
	return result
}
