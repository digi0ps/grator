package stagemanager

import (
	"fmt"
	"grator/httpclient"
	"grator/model"
	"sync"
)

type StageManager struct {
	httpClient httpclient.HttpClient
	config     model.Config
	actors     []*Actor
	wg         sync.WaitGroup
}

func (sm *StageManager) Start() {
	for i := 1; i <= sm.config.MaxActors; i++ {
		uuid := fmt.Sprintf("actor-%d", i)
		actor := NewActor(uuid, sm.httpClient, sm.config.BaseURL, sm.config.Actions)
		sm.actors = append(sm.actors, actor)

		sm.wg.Add(1)
		go func() {
			defer sm.wg.Done()
			actor.Play()
		}()
	}

	sm.wg.Wait()
	fmt.Println("Execution complete")
}

func NewStageManager(
	client httpclient.HttpClient,
	config model.Config,
) *StageManager {
	return &StageManager{
		httpClient: client,
		config:     config,
	}
}
