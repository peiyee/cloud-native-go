package filetransactionlogger

import (
	"fmt"
	"os"
	"time"

	"circuitbreaker/chapter5/transactionlogger"
)

type Event struct {
	EventType transactionlogger.EventType
	Key       string
	Value     string
}

type Logger struct {
	LastSequence uint64
	EventsCh     chan *Event
	file         *os.File
}

func NewTransactionLogger() transactionlogger.TransactionLogger {
	fileLogger := NewLogger("transaction.log")
	fileLogger.Run()

	return fileLogger
}

func RunMain() {
	logger := NewTransactionLogger()
	logger.WriteDelete("key")
	logger.WritePut("key", "value")
	time.Sleep(time.Second * 2)
}

func NewLogger(filename string) *Logger {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	return &Logger{file: file}
}
func (l *Logger) ReadEvents() {

}

func (l *Logger) Run() {
	events := make(chan *Event, 16)
	l.EventsCh = events

	go func() {
		for e := range l.EventsCh {
			fmt.Println("received", e.EventType)
			l.LastSequence++
			fmt.Fprintf(l.file,
				"%d\t%d\t%s\t%s\n",
				l.LastSequence, e.EventType, e.Key, e.Value,
			)
		}
	}()
}

func (l *Logger) WritePut(key, value string) {
	l.EventsCh <- &Event{EventType: transactionlogger.EventPut, Key: key, Value: value}
}

func (l *Logger) WriteDelete(key string) {
	l.EventsCh <- &Event{EventType: transactionlogger.EventDelete, Key: key}
}
