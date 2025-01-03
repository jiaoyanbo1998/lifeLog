// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	repository6 "lifelog-grpc/code/repository"
	cache4 "lifelog-grpc/code/repository/cache"
	service6 "lifelog-grpc/code/service"
	repository4 "lifelog-grpc/collectClip/repository"
	dao3 "lifelog-grpc/collectClip/repository/dao"
	service4 "lifelog-grpc/collectClip/service"
	repository5 "lifelog-grpc/comment/repository"
	dao4 "lifelog-grpc/comment/repository/dao"
	service5 "lifelog-grpc/comment/service"
	"lifelog-grpc/event/commentEvent"
	"lifelog-grpc/event/lifeLogEvent"
	repository2 "lifelog-grpc/interactive/repository"
	cache2 "lifelog-grpc/interactive/repository/cache"
	dao2 "lifelog-grpc/interactive/repository/dao"
	service2 "lifelog-grpc/interactive/service"
	"lifelog-grpc/ioc"
	"lifelog-grpc/lifeLog/repository"
	"lifelog-grpc/lifeLog/repository/cache"
	"lifelog-grpc/lifeLog/repository/dao"
	"lifelog-grpc/lifeLog/service"
	repository3 "lifelog-grpc/ranking/repository"
	cache3 "lifelog-grpc/ranking/repository/cache"
	service3 "lifelog-grpc/ranking/service"
	repository7 "lifelog-grpc/user/repository"
	cache5 "lifelog-grpc/user/repository/cache"
	dao5 "lifelog-grpc/user/repository/dao"
	service7 "lifelog-grpc/user/service"
	"lifelog-grpc/web"
)

// Injectors from wire.go:

func InitApp() *App {
	userServiceClient := ioc.InitUserServiceGRPCClient()
	logger := ioc.InitLogger()
	jwtHandler := web.NewJWTHandler(logger)
	userHandler := web.NewUserHandler(userServiceClient, logger, jwtHandler)
	cmdable := ioc.InitRedis()
	v := ioc.InitMiddlewares(logger, cmdable)
	db := ioc.GetMysql(logger)
	lifeLogDao := dao.NewLifeLogDao(db, logger)
	lifeLogCache := cache.NewLifeLogRedisCache(cmdable, logger)
	lifeLogRepository := repository.NewLifeLogRepository(lifeLogDao, logger, lifeLogCache)
	lifeLogService := service.NewLifeLogService(lifeLogRepository)
	interactiveDao := dao2.NewInteractiveDao(db, logger)
	interactiveCache := cache2.NewInteractiveCache(cmdable, logger)
	interactiveRepository := repository2.NewInteractiveRepository(interactiveDao, interactiveCache)
	interactiveService := service2.NewInteractiveService(interactiveRepository)
	rankingCache := cache3.NewRankingCacheRedis(cmdable)
	rankingRepository := repository3.NewRankingRepository(rankingCache)
	rankingService := service3.NewRankingService(rankingRepository)
	job := ioc.InitRankingJob(rankingService, logger, cmdable)
	client := ioc.InitKafka()
	syncProducer := ioc.InitSyncProducer(client)
	producer := lifeLogEvent.NewSaramaSyncProducer(syncProducer)
	lifeLogHandler := web.NewLifeLogHandler(logger, lifeLogService, interactiveService, job, producer)
	collectClipDao := dao3.NewCollectClipDao(db, logger)
	collectClipRepository := repository4.NewCollectClipRepository(collectClipDao)
	collectClipService := service4.NewCollectClipService(collectClipRepository)
	collectClipHandler := web.NewCollectClipHandler(collectClipService)
	commentDao := dao4.NewCommentDaoGorm(db, logger)
	commentRepository := repository5.NewCommentRepository(commentDao)
	commentService := service5.NewCommentService(commentRepository)
	commentHandler := web.NewCommentHandler(logger, commentService)
	bloomFilter := ioc.InitBloomFilter(cmdable)
	codeCache := cache4.NewCodeCache(cmdable, logger, bloomFilter)
	codeRepository := repository6.NewCodeRepository(codeCache)
	sendSmsService := ioc.InitSms(logger, cmdable)
	codeService := service6.NewCodeService(codeRepository, sendSmsService)
	codeHandler := web.NewCodeHandler(logger, codeService)
	engine := ioc.InitGin(userHandler, v, lifeLogHandler, collectClipHandler, commentHandler, codeHandler)
	readEventConsumer := lifeLogEvent.NewReadEventConsumer(client, logger, interactiveService)
	v2 := ioc.InitConsumers(readEventConsumer)
	cron := ioc.InitCronRankingJob(logger, job)
	app := &App{
		server:    engine,
		consumers: v2,
		cron:      cron,
	}
	return app
}

// wire.go:

// userSet user模块的依赖注入
var userSet = wire.NewSet(web.NewUserHandler, service7.NewUserService, repository7.NewUserRepository, dao5.NewUserDao, cache5.NewUserCache, ioc.InitUserServiceGRPCClient)

// codeSet code模块的依赖注入
var codeSet = wire.NewSet(service6.NewCodeService, repository6.NewCodeRepository, cache4.NewCodeCache, web.NewCodeHandler, ioc.InitCodeServiceGRPCClient)

// JwtSet 初始化jwt模块
var JwtSet = wire.NewSet(web.NewJWTHandler)

// 短信模块
var smsSet = wire.NewSet(ioc.InitSms)

// interactiveSet interactive模块的依赖注入
var interactiveSet = wire.NewSet(service2.NewInteractiveService, repository2.NewInteractiveRepository, dao2.NewInteractiveDao, cache2.NewInteractiveCache)

// LifeLog模块
var lifeLogSet = wire.NewSet(web.NewLifeLogHandler, service.NewLifeLogService, repository.NewLifeLogRepository, dao.NewLifeLogDao, cache.NewLifeLogRedisCache)

// collectClipSet collectClip模块的依赖注入
var collectClipSet = wire.NewSet(web.NewCollectClipHandler, service4.NewCollectClipService, repository4.NewCollectClipRepository, dao3.NewCollectClipDao)

// kafkaSet kafka模块的依赖注入
var kafkaSet = wire.NewSet(ioc.InitKafka, lifeLogEvent.NewReadEventBatchConsumer, lifeLogEvent.NewReadEventConsumer, lifeLogEvent.NewSaramaSyncProducer, commentEvent.NewCommentEventBatchConsumer, ioc.InitConsumers, ioc.InitSyncProducer)

// rankingSet ranking模块的依赖注入
var rankingSet = wire.NewSet(service3.NewRankingService, repository3.NewRankingRepository, cache3.NewRankingCacheRedis)

// rankingJobCronSet 热榜定时任务的依赖注入
var rankingJobCronSet = wire.NewSet(ioc.InitRankingJob, ioc.InitCronRankingJob)

// commentSet 评论
var commentSet = wire.NewSet(web.NewCommentHandler, service5.NewCommentService, repository5.NewCommentRepository, dao4.NewCommentDaoGorm)
