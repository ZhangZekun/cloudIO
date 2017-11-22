package service

import (
	"net/http"
)


func myError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	var a []byte = []byte("501 not implemented!")
	w.Write(a)
}