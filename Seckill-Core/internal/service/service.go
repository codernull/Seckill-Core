package service

import (
	pb "api/Seckill-Core/v1"

	"Seckill-Core/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSeckill-CoreService)

type Seckill-CoreService struct {
	pb.UnimplementedSeckill-CoreServer
	uc *biz.Seckill-CoreUseCase
}

// NewSeckill-CoreService
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@param uc
//	@return *Seckill-CoreService
func NewSeckill-CoreService(uc *biz.Seckill-CoreUseCase) *Seckill-CoreService {
	return &Seckill-CoreService{uc: uc}
}
