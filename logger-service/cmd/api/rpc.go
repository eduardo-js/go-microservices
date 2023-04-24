package main

import (
	"log"
	"logger-service/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	d := data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	}
	err := d.Insert(d)
	if err != nil {
		log.Println("Error writing to mongo", err)
	}
	*resp = "Processed log payload " + payload.Name + " successfully"
	return nil
}
