// Package main is the entry point for the c24-expense application.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/13excite/c24-expense/pkg/config"
	"github.com/13excite/c24-expense/pkg/driver"
	"github.com/13excite/c24-expense/pkg/helper"
	"github.com/13excite/c24-expense/pkg/jobs"
	"github.com/13excite/c24-expense/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	configPath := flag.String("config", "", "path to the configuration file")
	flag.Parse()

	conf := config.Config{}
	conf.Defaults()
	if *configPath != "" {
		conf.ReadConfigFile(*configPath)
	}

	// initialize the logger
	err := logger.InitLogger(&conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger := zap.S().With("package", "cmd")

	// Open a connection to the database
	conn, err := driver.OpenDB(conf.Clickhouse.Username,
		conf.Clickhouse.Password, conf.Clickhouse.Address, conf.Clickhouse.Database)
	if err != nil {
		logger.Error("Error opening database connection", zap.Error(err))
		os.Exit(1)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Wait for shutdown signal
	go func(ctx context.Context, cancel context.CancelFunc) {
		defer cancel()
		helper.WaitForShutdown(ctx)
	}(ctx, cancel)

	// Start the background job to parse CSV files
	parseJob := jobs.New(&conf)
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return parseJob.RunBackgroundParseJob(ctx)
	})

	// Wait for the background job to finish
	err = group.Wait()
	if err != nil {
		logger.Error("Error in background job", zap.Error(err))
	}
}
