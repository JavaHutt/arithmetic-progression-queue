package main

import (
	"fmt"
	"os"

	"github.com/JavaHutt/arithmetic-progression-queue/config"
	"github.com/JavaHutt/arithmetic-progression-queue/internal/action"
	"github.com/JavaHutt/arithmetic-progression-queue/internal/http"
	"github.com/JavaHutt/arithmetic-progression-queue/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger := logrus.New()

	arithmeticProcessor := action.NewArithmeticProcessor(*logger, cfg.ConcurrencyLimit())
	arithmeticProcessor.StartWorkers()

	taskService := service.NewTaskService(*logger, arithmeticProcessor)
	server := http.NewServer(*logger, cfg.HTTPServerPort(), taskService)

	doneChannel := make(chan bool)
	_ = doneChannel
	errorChannel := make(chan error)
	go func() {
		errorChannel <- server.Open()
	}()

	action.GracefulShutdown(cfg, errorChannel, server, doneChannel)
}
