package server


type sendMessageRequest struct {
  Name    string `json:"name"`
  Content string `json:"content"`
}
