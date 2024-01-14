package model

type Action struct {
	Method      string            `yaml:"method"`
	URL         string            `yaml:"url"`
	Body        string            `yaml:"body"`
	Repeat      int               `yaml:"repeat"`
	WaitFor     int               `yaml:"waitFor"`
	Headers     map[string]string `yaml:"headers"`
	StoreValues map[string]string `yaml:"storeValues"`
}

func (a *Action) Times() int {
	if a.Repeat == 0 {
		return 1
	}
	return a.Repeat
}

func (a *Action) ShouldWait() bool {
	return a.WaitFor > 0
}

func GetSampleActions() []Action {
	return []Action{
		{Method: "GET", URL: "/api/ping/"},
		{
			Method:      "POST",
			URL:         "/api/otp/generate",
			Body:        "{\"phone_number\": \"9840543050\"}",
			StoreValues: map[string]string{"phone_otp": "data.otp"},
		},
		{
			Method:      "POST",
			URL:         "/api/otp/customer/authenticate",
			Body:        "{\"phone_number\": \"9840543050\", \"otp\": \"%phone_otp%\"}",
			StoreValues: map[string]string{"token": "data.token"},
		},
		{
			Method:  "GET",
			URL:     "/api/geo/revgeo?latlng=12.835787029432241,80.13895099577203",
			Headers: map[string]string{"Authorization": "Bearer %token%"},
		},
	}
}
