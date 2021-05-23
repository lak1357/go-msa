package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/go-msa/database"
	"example.com/go-msa/user"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var httpAddr = flag.String("http", ":8080", "http listner address")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "user", "time:", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	ctx := context.Background()

	initDatabase(logger)
	defer database.DBcon.Close()

	var srv user.Service
	{
		repository := user.NewRepo(logger)
		srv = user.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := user.MakeEndpoints(srv)

	go func() {
		fmt.Println("listning on port", *httpAddr)
		handler := user.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}

func initDatabase(logger log.Logger) {
	var err error
	database.DBcon, err = gorm.Open("sqlite3", "user.db")
	if err != nil {
		level.Error(logger).Log("Failed to connect to database", err)
		panic("Failed to connect to database")
	}
	logger.Log("msg", "Database connection successfully opened")
	database.DBcon.AutoMigrate(&user.User{})
	logger.Log("msg", "Database migrated")
}
