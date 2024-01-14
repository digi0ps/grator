package main

import (
	"fmt"
	"grator/httpclient"
	"grator/model"
	"grator/stagemanager"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting executing...")

	config := model.NewConfigFromYamlFile("/Users/sravichandran/Code/digi0ps/grator/data/config.yml")
	client := httpclient.NewDefaultClient(&http.Client{
		Timeout: time.Duration(10 * time.Second),
	})

	sm := stagemanager.NewStageManager(client, config)
	sm.Start()
}
