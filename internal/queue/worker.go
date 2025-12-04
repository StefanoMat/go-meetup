package queue

import (
	"fmt"
	"time"
)

// StartWorker inicia um worker concorrente que consome mensagens da fila.
// Cada worker:
// - aguarda 2 minutos antes de começar (tempo para você enfileirar várias mensagens)
// - lê mensagens do channel q.Messages
// - "processa" cada mensagem (simulado com Sleep) e registra logs
func StartWorker(id int, q *Queue) {
	go func() {
		// Espera inicial para demonstrar acúmulo de mensagens na fila.
		fmt.Printf("[worker-%d] Aguardando 2m para iniciar o processamento...\n", id)
		time.Sleep(2 * time.Minute)
		// Loop de consumo: lê do channel; a leitura bloqueia até haver mensagens disponíveis.
		for msg := range q.Messages {
			fmt.Printf("[worker-%d] Processando: %s\n", id, msg.Body)
			// Simula trabalho de 1 segundo por mensagem.
			time.Sleep(1 * time.Second)
			// Log para observar a concorrência entre múltiplos workers.
			fmt.Printf("[worker-%d] Concluído: %s\n", id, msg.Body)
		}
	}()
}


