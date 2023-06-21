package entities

import (
	"context"
)

// Host is a struct that represents a host
type Host struct {
	// host fields
	ID       int64  `json:"id"`
	HostName string `json:"host_name"`
	HostIP   string `json:"host_ip"`
}

type HostUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Host, string, error)
}

// HostRepository is an interface that represents a host repository
type HostRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Host, string, error)
}
