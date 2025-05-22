package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, struct{}{})
}