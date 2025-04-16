package model

import "fmt"

type MetaData struct {
	Name       string
	CreatedAt  string
	ModifiedAt string
}

func (md *MetaData) String() string {
	return fmt.Sprintf("%s | %s | %s", md.Name, md.CreatedAt, md.ModifiedAt)
}
