package infrastructure

import (
	"github.com/spf13/viper"
	"github.com/surrealdb/surrealdb.go"
	"go.uber.org/zap"
)

type SurrealDB struct {
	Database *surrealdb.DB
	log      *zap.Logger
}

func NewSurrealDB(log *zap.Logger) *SurrealDB {
	log.Info("Connecting to surrealdb")
	db, err := surrealdb.New(viper.Get("DB_HOST").(string))
	if err != nil {
		log.Fatal("failed to connect to surrealdb", zap.Error(err))
	}

	log.Info("Connected to surrealdb at")
	log.Info("Signing in to surrealdb as root")
	_, err = db.Signin(map[string]interface{}{
		"user": viper.Get("DB_USER").(string),
		"pass": viper.Get("DB_PASSWORD").(string),
	})
	if err != nil {
		log.Fatal("failed to sign in to surrealdb", zap.Error(err))
	}

	log.Info("Using namespace and database")
	_, err = db.Use(viper.Get("DB_NAMESPACE").(string), viper.Get("DB_NAME").(string))
	if err != nil {
		log.Fatal("failed to use namespace and database", zap.Error(err))
	}

	return &SurrealDB{Database: db}
}
