//go:build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/codernull/seckill/internal/biz"
	"github.com/codernull/seckill/internal/conf"
	"github.com/codernull/seckill/internal/data"
	"github.com/codernull/seckill/internal/interfaces"
	"github.com/codernull/seckill/internal/server"
	"github.com/codernull/seckill/internal/service"
	"github.com/codernull/seckill/internal/task"
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
