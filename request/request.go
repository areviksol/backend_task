package request

import (
	"net/http"
)

func GetIdentifier(r *http.Request) string {
	return r.URL.Query().Get("id")
}
