package main

import "time"

// Ticket  -- карточка, сообщающая нам о том, какой клиент сидит за столом и в какое время он за него сел
type Ticket struct {
	name      string    // Имя клиента
	startTime time.Time // Время, когда клиент сел за стол
}
