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
	}
}
