package rest

import "net/http"

func PingHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Server Active"))
}
