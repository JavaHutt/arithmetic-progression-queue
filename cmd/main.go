package main

import (
	"fmt"
	"os"
	"strconv"

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

	concurrencyLimit := cfg.ConcurrencyLimit()
	if len(os.Args) > 1 {
		concurrencyLimit, err = strconv.Atoi(os.Args[1])
		if err != nil {
			concurrencyLimit = cfg.ConcurrencyLimit()
		}
	}

	logger := logrus.New()

	arithmeticProcessor := action.NewArithmeticProcessor(*logger, concurrencyLimit)
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
