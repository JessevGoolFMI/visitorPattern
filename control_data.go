package main

import "sync"

type ControlData struct {
	Data map[string]interface{}
	once sync.Once
}

var (
	globalControlDataInstance *ControlData
	once                      sync.Once
)

func GetControlData() *ControlData {
	once.Do(func() {
		globalControlDataInstance = &ControlData{
			Data: make(map[string]interface{}),
		}
	})
	return globalControlDataInstance
}
