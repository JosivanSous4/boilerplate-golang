package config

import (
	"os"
)

type Config struct {
    MySQLDSN   string
    MongoDBURI string
    MessagingURI string
    JWTSecret  string
    AwsEndpoint string
    AwsRegion string
}

func LoadConfig() Config {
    return Config{
        MySQLDSN:   os.Getenv("MYSQL_DSN"),
        MongoDBURI: os.Getenv("MONGODB_URI"),
        MessagingURI: os.Getenv("MESSAGING_URI"),
        JWTSecret:  os.Getenv("JWT_SECRET"),
        AwsEndpoint: os.Getenv("AWS_ENDPOINT"),
        AwsRegion: os.Getenv("AWS_DEFAULT_REGION"),
    }
}
