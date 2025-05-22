package main

import "net/http"

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusBadRequest, "Something went wrong")
}