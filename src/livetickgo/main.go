package main

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/livetickgo/grpc/livetickgo"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/livetickgo/service"
	"google.golang.org/grpc"
)

func init() {
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)
}

var grpcPort = 9111

func main() {
	log.Info("Starting Livetickgo server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on grpc port %d: %v", grpcPort, err)
	}
	grpcServer := grpc.NewServer()
	transferService := service.NewBankTransferService()
	transferService.Start()
	defer transferService.Stop()
	banktransfer.RegisterBankTransferServer(grpcServer, service.NewBankTransferService())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
