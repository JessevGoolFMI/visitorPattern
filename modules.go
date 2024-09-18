package main

type ControlCore struct {
	BaseModule
}

type Compressor struct {
	BaseModule
}

func (c *Compressor) Accept(visitor ModuleVisitor) {
	visitor.VisitCompressor(c)
}

type CompressorData struct {
	ValueX        int
	ValueY        float64
	SomethingElse string
}

func NewCompressor() *Compressor {
	compressor := Compressor{NewBaseModule()}
	compressor.ControlData.Data["compressor"] = CompressorData{
		ValueX:        20,
		ValueY:        5.53,
		SomethingElse: "test",
	}
	return &compressor
}

type Dispenser struct {
	BaseModule
}

func (d *Dispenser) Accept(visitor ModuleVisitor) {
	visitor.VisitDispenser(d)
}

type DispenserData struct {
	ValueA        int
	ValueB        float64
	SomethingElse EmbeddedData
}

type EmbeddedData struct {
	Id      string
	Payload []byte
}

func NewDispenser() *Dispenser {
	dispenser := Dispenser{NewBaseModule()}
	dispenser.ControlData.Data["dispenser"] = DispenserData{
		ValueA: 41,
		ValueB: 352.2352,
		SomethingElse: EmbeddedData{
			Id:      "testId",
			Payload: []byte("testPayload"),
		},
	}
	return &dispenser
}
