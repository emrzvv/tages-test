package model

import (
	"fmt"
	"time"
)

type MetaData struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (md *MetaData) String() string {
	return fmt.Sprintf("%s | %s | %s", md.Name, md.CreatedAt.Format(time.RFC3339), md.ModifiedAt.Format(time.RFC3339))
}
