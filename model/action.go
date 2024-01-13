package model

type Action struct {
	Method      string
	URL         string
	Body        string
	Headers     map[string]string
	StoreValues map[string]string
}

func GetSampleActions() []Action {
	return []Action{
		{Method: "GET", URL: "/api/ping/"},
		{Method: "POST", URL: "/api/otp/generate", Body: "{\"phone_number\": \"9840543050\"}", StoreValues: map[string]string{"otp": "data.otp"}},
	}
}
