package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/13excite/c24-expences/pkg/c24parser"
	"github.com/13excite/c24-expences/pkg/config"
	"github.com/13excite/c24-expences/pkg/driver"
	"github.com/13excite/c24-expences/pkg/filemanager"
	"github.com/13excite/c24-expences/pkg/models"
)

func main() {
	configPath := flag.String("config", "", "path to the configuration file")
	flag.Parse()

	conf := config.Config{}
	conf.Defaults()
	if *configPath != "" {
		conf.ReadConfigFile(*configPath)
	}

	conn, err := driver.OpenDB(conf.Clickhouse.Username,
		conf.Clickhouse.Password, conf.Clickhouse.Address, conf.Clickhouse.Database)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	model := models.NewModel(conn)

	fileMgr := filemanager.NewFileManager("./input/", &model.DB)
	files, err := fileMgr.GetFilesToUpload()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csvParser := c24parser.NewParser()
	for _, file := range files {
		if err := csvParser.ParseFile(file.Path); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Starts to create transaction:")
		for _, t := range csvParser.GetTransactions() {
			err := model.DB.InsertTransaction(t)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

}
