package handlers

import (
	"encoding/json"
	"net/http"

	"go-meetup-payment/internal/domain"
	"go-meetup-payment/internal/queue"
)

// NewPublishHandler registra o endpoint HTTP responsável por enfileirar mensagens.
// Método e rota: POST /publish
// Entrada esperada: JSON no formato {"body":"..."}
// Saída: JSON confirmando o enfileiramento {"status":"enqueued"}
func NewPublishHandler(q *queue.Queue) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /publish", func(w http.ResponseWriter, r *http.Request) {
		var m domain.Message
		// Decodifica o corpo JSON da requisição para o struct Message.
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Publica a mensagem na fila para processamento assíncrono pelos workers.
		q.Publish(m)
		// Resposta simples informando que a mensagem foi enfileirada com sucesso.
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "enqueued"})
	})
	return mux
}


