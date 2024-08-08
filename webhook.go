package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handleWebhook)

	panic(http.ListenAndServe(":60000", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "webssh resp.")
}
