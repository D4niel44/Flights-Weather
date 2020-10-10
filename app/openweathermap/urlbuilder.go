package openweathermap

import "fmt"

type getRequestURLBuilder struct {
	base   string
	params map[string]string
}

func newGetRequestURLBuilder(baseURL string) *getRequestURLBuilder {
	return &getRequestURLBuilder{base: baseURL, params: make(map[string]string)}
}

func (r *getRequestURLBuilder) addParameter(key, value string) {
	r.params[key] = value
}

func (r *getRequestURLBuilder) makeURL() string {
	url := r.base
	i := 1
	length := len(r.params)
	for key, value := range r.params {
		url += fmt.Sprintf("%s=%s", key, value)
		if i != length {
			url += "&"
		}
		i++
	}
	return url
}
