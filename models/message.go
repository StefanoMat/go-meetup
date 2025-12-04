package models

// Message representa o payload recebido via HTTP.
// O campo Body é populado a partir do JSON {"body":"..."}.
type Message struct {
	Body string `json:"body"`
}


