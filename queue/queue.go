package queue

import "go-meetup-payment/models"

// Queue é uma fila simples baseada em channel.
// O channel com buffer permite enfileirar mensagens sem bloquear imediatamente.
type Queue struct {
	Messages chan models.Message
}

// NewQueue cria a fila com capacidade configurada (buffer).
// Se o buffer encher, novas publicações irão bloquear até um worker consumir.
func NewQueue() *Queue {
	return &Queue{
		Messages: make(chan models.Message, 50),
	}
}

// Publish coloca a mensagem na fila para processamento assíncrono pelos workers.
func (q *Queue) Publish(msg models.Message) {
	q.Messages <- msg
}


