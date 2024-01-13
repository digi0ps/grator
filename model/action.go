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
		{
			Method:  "GET",
			URL:     "/api/geo/revgeo?latlng=12.835787029432241,80.13895099577203",
			Headers: map[string]string{"Authorization": "Bearer %token%"},
		},
	}
}
