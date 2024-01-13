package httpclient

type HttpClient interface {
	MakeRequest(method string, url string, body string, headers map[string]string) (string, int, error)
}
