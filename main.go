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

	config := model.NewConfig("http://157.245.97.150:80", 1)
	client := httpclient.NewDefaultClient(&http.Client{
		Timeout: time.Duration(10 * time.Second),
	})

	sm := stagemanager.NewStageManager(client, config)
	sm.Start()
}
