package storage

import (
	"database/sql"
	"sync"
	"time"

	"github.com/emrzvv/tages-test/cfg"
	"github.com/emrzvv/tages-test/internal/app/model"
	_ "github.com/mattn/go-sqlite3"
)

type MetaStorage interface {
	InsertMeta(meta model.MetaData) error
	UpdateMeta(meta model.MetaData) error
	GetMetaByName(name string) (model.MetaData, bool)
	GetMetaList() []model.MetaData
}

type InMemoryMetaStorage struct {
	config *cfg.Config
	mutex  sync.RWMutex
	db     map[string]model.MetaData
}

func NewInMemoryMetaStorage(config *cfg.Config) *InMemoryMetaStorage {
	return &InMemoryMetaStorage{
		config: config,
		mutex:  sync.RWMutex{},
		db:     make(map[string]model.MetaData),
	}
}

func (imms *InMemoryMetaStorage) InsertMeta(meta model.MetaData) error {
	imms.mutex.Lock()
	defer imms.mutex.Unlock()
	imms.db[meta.Name] = meta
	return nil
}

func (imms *InMemoryMetaStorage) UpdateMeta(meta model.MetaData) error {
	imms.mutex.Lock()
	defer imms.mutex.Unlock()
	imms.db[meta.Name] = meta
	return nil
}

func (imms *InMemoryMetaStorage) GetMetaByName(name string) (model.MetaData, bool) {
	imms.mutex.RLock()
	defer imms.mutex.RUnlock()
	value, ok := imms.db[name]
	return value, ok
}

func (imms *InMemoryMetaStorage) GetMetaList() []model.MetaData {
	imms.mutex.Lock()
	defer imms.mutex.Unlock()
	result := make([]model.MetaData, 0, len(imms.db))
	for _, value := range imms.db {
		result = append(result, value)
	}
	return result
}

type SQLiteMetaStorage struct {
	config *cfg.Config
	db     *sql.DB
}

// комментарий к решению: sqlite не хранит даты,
// поэтому использую костыль с преобразованием в строку и обратно
func NewSQLiteMetaStorage(config *cfg.Config) (MetaStorage, error) {
	db, err := sql.Open("sqlite3", config.SQLiteDBPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`create table if not exists meta(name text pk, created_at text, modified_at text)`)
	if err != nil {
		return nil, err
	}
	return &SQLiteMetaStorage{
		config: config,
		db:     db,
	}, nil
}

func (s *SQLiteMetaStorage) InsertMeta(meta model.MetaData) error {
	_, err := s.db.Exec(`
        insert into meta (name, created_at, modified_at)
        values (?, ?, ?)
    `, meta.Name, meta.CreatedAt.Format(time.RFC3339), meta.ModifiedAt.Format(time.RFC3339))
	return err
}

func (s *SQLiteMetaStorage) UpdateMeta(meta model.MetaData) error {
	_, err := s.db.Exec(`
        update meta set created_at = ?, modified_at = ?
        where name = ?
    `, meta.CreatedAt.Format(time.RFC3339), meta.ModifiedAt.Format(time.RFC3339), meta.Name)
	return err
}

func (s *SQLiteMetaStorage) GetMetaByName(name string) (model.MetaData, bool) {
	row := s.db.QueryRow(`
        select name, created_at, modified_at from meta
        where name = ?
        limit 1
    `, name)

	var result model.MetaData
	var createdAtStr, modifiedAtStr string
	err := row.Scan(&result.Name, &createdAtStr, &modifiedAtStr)
	if err == sql.ErrNoRows {
		return model.MetaData{}, false
	} else if err != nil {
		return model.MetaData{}, false
	}

	result.CreatedAt, err = time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return model.MetaData{}, false
	}
	result.ModifiedAt, err = time.Parse(time.RFC3339, modifiedAtStr)
	if err != nil {
		return model.MetaData{}, false
	}
	return result, true
}

func (s *SQLiteMetaStorage) GetMetaList() []model.MetaData {
	rows, err := s.db.Query(`
        select name, created_at, modified_at
        from meta
    `)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var metas []model.MetaData
	for rows.Next() {
		var m model.MetaData
		var createdAtStr, modifiedAtStr string
		if err := rows.Scan(&m.Name, &createdAtStr, &modifiedAtStr); err != nil {
			continue
		}
		m.CreatedAt, err = time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			continue
		}
		m.ModifiedAt, err = time.Parse(time.RFC3339, modifiedAtStr)
		if err != nil {
			continue
		}
		metas = append(metas, m)
	}
	return metas
}
