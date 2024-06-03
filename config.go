package goconfig

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/julhan07/infrastructure-be/pgx"
	"github.com/julhan07/infrastructure-be/redis"
)

type Config struct {
	AppVersion string
	AppPort    string
	AppHost    string
	Pgx        pgx.Connection
	Redis      redis.Connection
	JWTSecret  string
	StaticPath string
}

func InitConfig() (conf Config) {

	envConf, err := godotenv.Read("./.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	conf.AppVersion = envConf["APP_VERSION"]
	conf.AppPort = envConf["APP_PORT"]
	conf.AppHost = envConf["APP_HOST"]

	portPgx, _ := strconv.Atoi(envConf["DB_PORT"])

	conf.Pgx = pgx.Connection{
		Host:         envConf["DB_HOST"],
		Port:         portPgx,
		User:         envConf["DB_USER"],
		Password:     envConf["DB_PASSWORD"],
		DatabaseName: envConf["DB_NAME"],
		SslMode:      envConf["DB_SSL_MODE"],
	}

	redisPort, _ := strconv.Atoi(envConf["REDIS_PORT"])
	conf.Redis = redis.Connection{
		Host:     envConf["REDIS_HOST"],
		Port:     redisPort,
		Password: envConf["REDIS_PASSWORD"],
	}

	conf.JWTSecret = envConf["JWT_SECRET"]
	conf.StaticPath = envConf["STATIC_PATH"]

	return conf
}
