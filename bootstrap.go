package main

import (
	"account_service/http"
	"account_service/library"
	"account_service/library/env"
	"account_service/logger"
	"account_service/providers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

//bootstrap providers,以及routines
func bootStrap() (err error) {
	//加载环境变量
	filePath := ".env"
	if err = godotenv.Load(filePath); err != nil {
		return
	}
	log.Println("env loadded from file ", filePath)

	err, shutdownLogger := logger.Start()
	if err != nil {
		return
	}
	log.Println("Logger Started ")

	//加载Redis连接池
	port, err := env.GetIntVal("REDIS_PORT_ACCOUNT")
	if err != nil {
		return
	}
	poolSize, err := env.GetIntVal("REDIS_POOL_SIZE")
	if err != nil {
		return
	}
	redisConf := library.RedisConfig{
		ConnectionName: os.Getenv("SERVICE_NAME"),
		Addr:           os.Getenv("REDIS_ADDR_ACCOUNT"),
		Port:           port,
		Password:       env.GetStringVal("REDIS_PSWD_ACCOUNT"),
		DB:             0,
		PoolSize:       poolSize,
	}
	providers.RedisClient, err = library.NewRedisClient(&redisConf)
	if err != nil {
		return
	}

	//http server
	err, shutdownHttpServer := http.Start()
	if err != nil {
		return
	}
	log.Println("Httpserver started ")

	//wait for sys signals
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sig := <-exitChan:
		log.Println("Doing cleaning works before shutdown...")
		shutdownLogger()
		shutdownHttpServer()
		log.Println("You abandoned me, bye bye", sig)
	}
	return
}
