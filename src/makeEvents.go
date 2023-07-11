package main

import (
	"fmt"
	"time"
)

// Создание сообщений с нужным содержанием
func makeEvent(eventTime time.Time, code int, msg string) string {
	eventTimeString := eventTime.Format("15:04")
	return fmt.Sprintf("%s %d %s", eventTimeString, code, msg)
}
