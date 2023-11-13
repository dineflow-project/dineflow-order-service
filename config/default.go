package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	env := os.Getenv("USE_DOT_ENV")
	if env == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func EnvAMQPURL() string {
	loadDotEnv()
	return os.Getenv("AMQP_URL")
}

func EnvNotiQueueName() string {
	loadDotEnv()
	return os.Getenv("NOTI_QUEUE_NAME")
}

func EnvMongoInitDBRootUsername() string {
	loadDotEnv()
	return os.Getenv("MONGO_INITDB_ROOT_USERNAME")
}

func EnvMongoInitDBRootPassword() string {
	loadDotEnv()
	return os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
}

func EnvMongoDBURI() string {
	loadDotEnv()
	return os.Getenv("MONGODB_LOCAL_URI")
}

func EnvGRPCServerAddress() string {
	loadDotEnv()
	return os.Getenv("GRPC_SERVER_ADDRESS")
}
