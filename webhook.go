package main

import (
	"bytes"
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
	if Lark_Webhook_Uri == "" {
		panic("no uri found, set env: LARK_WEBHOOK_URI before start.")
	}
	log.Printf("Lark_Webhook_Uri: %s", Lark_Webhook_Uri)
	http.HandleFunc("/webhook", handleWebhook)

	panic(http.ListenAndServe(":60000", nil))
}

type Alert struct {
	Status      string            `json:"status"`
	StartsAt    string            `json:"startsAt"`
	EndsAt      string            `json:"endsAt"`
	Values      interface{}       `json:"values"`
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
		log.Printf("read request body error, %+v", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析 JSON 负载
	var payload *Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Printf("unmarshal payload error, %+v", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Received: %+v\n", payload)
	// 输出解析结果
	// fmt.Fprintf(w, "Received: %+v\n", payload)

	// 转发告警
	var larkPayLoad = NewLarkPayLoad(payload.Message)
	larkPLJson, err := json.Marshal(larkPayLoad)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("get lark request json: %s", larkPLJson)
	httpc, err := http.NewRequest("POST", Lark_Webhook_Uri, bytes.NewBuffer(larkPLJson))
	if err != nil {
		log.Println("Error creating request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 设置请求头
	httpc.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(httpc)
	if err != nil {
		log.Println("Error sending request:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	// 读取响应
	if resp.StatusCode != http.StatusOK {
		log.Printf("lark webhook response error, status code: %v", resp.StatusCode)
		http.Error(w, "", http.StatusBadRequest)
	}
	log.Println("send")

	fmt.Fprintln(w, "webssh resp.")
}

type LarkPayLoadTextContent struct {
	Text string `json:"text"`
}

type LarkPayLoad struct {
	MsgType string                 `json:"msg_type"`
	Content LarkPayLoadTextContent `json:"content"`
}

func NewLarkPayLoad(text string) *LarkPayLoad {
	content := LarkPayLoadTextContent{
		Text: text,
	}
	return &LarkPayLoad{
		MsgType: "text",
		Content: content,
	}
}
