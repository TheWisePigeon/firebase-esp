package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Received call")
  w.WriteHeader(http.StatusOK)
  return
}
