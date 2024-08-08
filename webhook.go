package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var Lark_Webhook_Uri string

func main() {
	Lark_Webhook_Uri = os.Getenv("LARK_WEBHOOK_URI")

	http.HandleFunc("/webhook", handleWebhook)

	panic(http.ListenAndServe(":60000", nil))
}

type Alert struct {
	Status      string            `json:"status"`
	StartsAt    string            `json:"startsAt"`
	EndsAt      string            `json:"endsAt"`
	Values      []interface{}     `json:"values"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}
type Payload struct {
	Receiver string  `json:"receiver"`
	Status   string  `json:"status"`
	Alerts   []Alert `json:"alerts"`
	Title    string  `json:"title"`
	State    string  `json:"state"`
	Message  string  `json:"message"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析 JSON 负载
	var payload *Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Received: %+v\n", payload)
	// 输出解析结果
	// fmt.Fprintf(w, "Received: %+v\n", payload)

	fmt.Fprintln(w, "webssh resp.")
}
