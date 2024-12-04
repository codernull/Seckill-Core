package service

import (
	pb "github.com/codernull/seckill/api/seckill/v1"
	"github.com/codernull/seckill/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserXService)

type UserXService struct {
	pb.UnimplementedUserXServer
	uc *biz.UserXUseCase
}

// NewUserXService
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:
//	@param uc
//	@return *UserXService
func NewUserXService(uc *biz.UserXUseCase) *UserXService {
	return &UserXService{uc: uc}
}
