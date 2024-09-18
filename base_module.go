package main

type BaseModule struct {
	ControlData *ControlData
}

func NewBaseModule() BaseModule {
	return BaseModule{
		ControlData: GetControlData(),
	}
}
