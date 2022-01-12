package bootstrap

import (
	"api-http-hex-golang/internal/platform/server"
	"api-http-hex-golang/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "localhost"
	port   = 8080
	dbUser = "root"
	dbPass = "secret-pass"
	dbHost = "127.0.0.1"
	dbPort = "3306"
	dbName = "go_ddd"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
