package response

import "net/http"

type Raw struct {
	Data   interface{}
	Cookie *http.Cookie
}
