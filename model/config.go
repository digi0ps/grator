package model

type Config struct {
	BaseURL   string
	MaxActors int
	Actions   []Action
}

func NewConfig(baseURL string, maxActors int) Config {
	return Config{
		BaseURL:   baseURL,
		MaxActors: maxActors,
		Actions:   GetSampleActions(),
	}
}
