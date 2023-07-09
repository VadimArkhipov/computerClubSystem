package main

import "time"

type Ticket struct {
	name      string    // Имя клиента
	startTime time.Time // Время, когда клиент сел за стол
}
