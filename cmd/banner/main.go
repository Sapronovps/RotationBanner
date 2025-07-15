package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/app"
	"github.com/Sapronovps/RotationBanner/internal/broker/kafka"
	"github.com/Sapronovps/RotationBanner/internal/logger"
	"github.com/Sapronovps/RotationBanner/internal/server/grpc"
	"github.com/Sapronovps/RotationBanner/internal/server/http"
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"github.com/Sapronovps/RotationBanner/internal/storage/memory"
	"github.com/Sapronovps/RotationBanner/internal/storage/sql"
	_ "github.com/lib/pq" // for postgres
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Username, config.DB.Password, config.DB.DBName)
	var storageApp storage.Storage
	if config.DB.InMemory {
		storageApp = memory.New()
	} else {
		newDB := sql.NewDB(dsn)
		defer newDB.Close()
		storageApp = sql.New(newDB)
	}

	logg := logger.New(config.Logger.Level, config.Logger.File)
	kafkaProducer, err := kafka.NewKafkaProducer(config.Kafka.Brokers, config.Kafka.EventsTopic, config.Kafka.RetryMax)
	if err != nil {
		logg.Fatal("failed create kafka producer: " + err.Error())
	}
	application := app.NewApp(logg, storageApp, kafkaProducer)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	if config.Server.IsHTTP {
		serverAddress := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
		server := http.NewServer(serverAddress, application, logg)

		go func() {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err := server.Stop(ctx); err != nil {
				logg.Error("failed to stop http server: " + err.Error())
			}
		}()

		logg.Info("Microservice Rotation Banner is running...")

		if err := server.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	} else {
		server := grpc.NewBannerGrpcServer(config.Server.AddressGrpc, logg, application)

		go func() {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err := server.Stop(ctx); err != nil {
				logg.Error("failed to stop grpc server: " + err.Error())
			}
		}()

		logg.Info("Microservice Rotation Banner is running...")

		if err := server.Start(ctx); err != nil {
			logg.Error("failed to start grpc server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	}
}
