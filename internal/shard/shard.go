package shard

import (
	"db-sharding/internal/store"
	"fmt"
)

type Shard struct {
	ID       int
	Name     string
	Database string
	Host     string
	Port     int
	Store    *store.Store
}

func (s Shard) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s Shard) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%d%s",
		s.Host,
		s.Port, s.Database)
}
