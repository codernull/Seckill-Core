//go:build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"Seckill-Core/internal/biz"
	"Seckill-Core/internal/conf"
	"Seckill-Core/internal/data"
	"Seckill-Core/internal/interfaces"
	"Seckill-Core/internal/server"
	"Seckill-Core/internal/service"
	"Seckill-Core/internal/task"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: wireApp init kratos application.
//	@param *conf.Server
//	@param *conf.Data
//	@return *kratos.App
//	@return func()
//	@return error
func wireApp(*conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		interfaces.ProviderSet,
		task.ProviderSet,
		newApp))
}
