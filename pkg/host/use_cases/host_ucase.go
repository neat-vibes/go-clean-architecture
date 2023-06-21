package usecases

import (
	"context"
	"time"

	"github.com/neat-vibes/go-clean-architecture/entities"
)

type hostUsecase struct {
	hostRepo       entities.HostRepository
	contextTimeout time.Duration
}

func NewHostUsecase(hostRepo entities.HostRepository, timeout time.Duration) entities.HostUsecase {
	return &hostUsecase{
		hostRepo:       hostRepo,
		contextTimeout: timeout,
	}
}

func (h *hostUsecase) Fetch(ctx context.Context, cursor string, num int64) ([]entities.Host, string, error) {

	return nil, "", nil
}
