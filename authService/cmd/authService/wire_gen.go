// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/biz"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/casbin"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/conf"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/redis"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/server"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/service"
	"github.com/chencheng8888/tiktok_e-commence/authService/internal/token"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, confToken *conf.Token, logger log.Logger) (*kratos.App, func(), error) {
	authCase := casbin.NewAuthCase(data)
	client := redis.NewRedisDB(data)
	cache := redis.NewCache(client)
	jwTer := token.NewJWTer()
	signer := biz.NewSigner(cache, jwTer, confToken)
	verifyer := biz.NewVerifier(cache, jwTer, confToken)
	authUserCase := biz.NewAuthUserCase(authCase, signer, verifyer)
	authServiceService := service.NewAuthServiceService(authUserCase)
	grpcServer := server.NewGRPCServer(confServer, authServiceService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
	}, nil
}
