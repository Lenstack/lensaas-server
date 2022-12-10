package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Postgres struct {
	host         string
	port         string
	databaseName string
	user         string
	password     string
	Database     squirrel.StatementBuilderType
}

func NewPostgres(host string, port string, databaseName string, user string, password string, loggerManager *zap.Logger) *Postgres {
	datasource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		host, user, password, databaseName, port)
	loggerManager.Sugar().Infof("Connecting to database: %s", databaseName)
	database, err := sql.Open("postgres", datasource)
	if err != nil {
		loggerManager.Sugar().Errorf("Error connecting to database: %s", err)
		return nil
	}
	return &Postgres{Database: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(database)}
}
