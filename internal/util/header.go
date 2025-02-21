package util

import "net/http"

func HeaderValues(key string, h http.Header) string {
	return h.Get(key)
}
