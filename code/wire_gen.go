// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"lifelog-grpc/code/grpc"
	"lifelog-grpc/code/ioc"
	"lifelog-grpc/code/repository"
	"lifelog-grpc/code/repository/cache"
	"lifelog-grpc/code/service"
)

// Injectors from wire.go:

// InitCodeServiceGRPCService 初始化CodeServiceGRPCService
func InitCodeServiceGRPCService() *grpc.CodeServiceGRPCService {
	cmdable := ioc.InitRedis()
	logger := ioc.InitLogger()
	bloomFilter := ioc.InitBloomFilter(cmdable)
	codeCache := cache.NewCodeCache(cmdable, logger, bloomFilter)
	codeRepository := repository.NewCodeRepository(codeCache)
	sendSmsService := ioc.InitSms(logger, cmdable)
	codeService := service.NewCodeService(codeRepository, sendSmsService)
	codeServiceGRPCService := grpc.NewCodeServiceGRPCService(codeService)
	return codeServiceGRPCService
}

// wire.go:

// codeSet 注入
var codeSet = wire.NewSet(service.NewCodeService, repository.NewCodeRepository, cache.NewCodeCache)

var third = wire.NewSet(ioc.InitRedis, ioc.InitSms, ioc.InitLogger, ioc.InitBloomFilter)