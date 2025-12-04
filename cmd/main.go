package main

import (
	"fmt"
	"net/http"

	"go-meetup-payment/internal/handlers"
	"go-meetup-payment/internal/queue"
)

// main inicializa a fila, cria workers concorrentes e sobe o servidor HTTP.
func main() {

	// Cria a fila compartilhada entre produtores (HTTP) e consumidores (workers).
	q := queue.NewQueue()
	// Inicia 3 workers; todos leem do mesmo channel (concorrência).
	for i := 1; i <= 3; i++ {
		queue.StartWorker(i, q)
	}
	// Registra o endpoint POST /publish que enfileira mensagens.
	handler := handlers.NewPublishHandler(q)
	// Sobe o servidor em :8080. Envie requisições enquanto os workers aguardam.
	fmt.Println("Rodando em :8080")
	_ = http.ListenAndServe(":8080", handler)
}


