package stagemanager

import (
	"encoding/json"
	"fmt"
	"grator/httpclient"
	"grator/model"
	"grator/utils"
	"time"
)

type Actor struct {
	httpClient httpclient.HttpClient
	baseURL    string
	actions    []model.Action
	uuid       string
	session    map[string]interface{}
}

func NewActor(uuid string, httpClient httpclient.HttpClient, baseURL string, actions []model.Action) *Actor {
	return &Actor{
		uuid:       uuid,
		httpClient: httpClient,
		baseURL:    baseURL,
		actions:    actions,
		session:    map[string]interface{}{},
	}
}

func (a *Actor) Play() {
	for _, action := range a.actions {
		a.execute(action)
	}
	return
}

func (a *Actor) formURL(action model.Action) string {
	if a.baseURL == "" {
		return action.URL
	}

	return fmt.Sprintf("%s%s", a.baseURL, action.URL)
}

func (a *Actor) parseJson(body string, storeValues map[string]string) error {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return err
	}

	for key, value := range storeValues {
		foundVal, err := utils.ExtractValue(data, value)
		if err != nil {
			return err
		}
		a.session[key] = foundVal
		fmt.Printf("[%s] Storing %s = %v\n", a.uuid, key, foundVal)
	}

	return nil
}

func (a *Actor) execute(action model.Action) {
	fmt.Printf("[%s] Executing %s %s Body = %s\n", a.uuid, action.Method, action.URL, action.Body)

	url := a.formURL(action)
	startTs := time.Now()
	body, statusCode, err := a.httpClient.MakeRequest(action.Method, url, action.Body, action.Headers)
	timeTaken := time.Since(startTs).Milliseconds()
	if err != nil {
		fmt.Printf("[%s] ERROR | url: %s | status: %d | %dms | err: %d", a.uuid, action, statusCode, timeTaken, err)
		return
	}

	if action.StoreValues != nil {
		err := a.parseJson(body, action.StoreValues)
		if err != nil {
			fmt.Printf("[%s] ERROR | url: %s | status: %d | json.err: %s", a.uuid, action, statusCode, err)
			return
		}
	}

	fmt.Printf("[%s] SUCCESS | url: %s | status: %d | %dms | body: %s\n", a.uuid, action, statusCode, timeTaken, body)
}
