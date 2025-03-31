// Package jobs provides the functionality to parse CSV files and insert
// the transactions into the database in the background.
package jobs

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/13excite/c24-expences/pkg/c24parser"
	"github.com/13excite/c24-expences/pkg/config"
	"github.com/13excite/c24-expences/pkg/driver"
	"github.com/13excite/c24-expences/pkg/filemanager"
	"github.com/13excite/c24-expences/pkg/models"
)

type parser interface {
	ParseFile(string) error
}

// Job struct that holds the logger, parser and configuration of the job
type Job struct {
	logger *zap.SugaredLogger
	parser parser
	config *config.Config
}

// New returns a new Job struct
func New(conf *config.Config, parser parser) *Job {
	return &Job{
		config: conf,
		logger: zap.S().With("package", "job"),
		parser: parser,
	}
}

func (j *Job) parserRunner() {
	// job runs not so often, so we can afford to create a new connection every time
	conn, err := driver.OpenDB(j.config.Clickhouse.Username,
		j.config.Clickhouse.Password, j.config.Clickhouse.Address, j.config.Clickhouse.Database)
	if err != nil {
		j.logger.Error("Error opening database connection", zap.Error(err))
		return
	}
	defer conn.Close()

	model := models.NewModel(conn)

	fileMgr := filemanager.NewFileManager(j.config.InputDir, &model.DB)
	files, err := fileMgr.GetFilesToUpload()
	if err != nil {
		j.logger.Error("Error getting files to upload", zap.Error(err))
		return
	}

	csvParser := c24parser.NewParser()
	for _, file := range files {
		if err := csvParser.ParseFile(file.Path); err != nil {
			j.logger.Error("Error parsing file", zap.Error(err))
			continue
		}
		j.logger.Info("Starts to create transaction")
		for _, t := range csvParser.GetTransactions() {
			err := model.DB.InsertTransaction(t)
			if err != nil {
				j.logger.Error("Error inserting transaction", zap.Error(err))
				continue
			}
		}
	}
}

// RunBackgroundParseJob runs the background job that parses the CSV files
func (j *Job) RunBackgroundParseJob(ctx context.Context) error {
	j.logger.Info("Background ParseFileJob is starting with run every", j.config.RunEvery)

	ticker := time.NewTicker(time.Duration(j.config.RunEvery) * time.Minute)

	for {
		select {
		case <-ticker.C:

			j.parserRunner()

		case <-ctx.Done():
			return nil
		}
	}
}
