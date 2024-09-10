package transactionlogger

type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
	ReadEvents()
	Run()
}

type EventType uint

const (
	_ = iota
	EventPut
	EventDelete
)
