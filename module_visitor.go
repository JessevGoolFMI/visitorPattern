package main

import (
	"fmt"
	"sync"
)

type ModuleVisitor interface {
	VisitCompressor(*Compressor)
	VisitDispenser(*Dispenser)
}

type ReportGeneratorVisitor struct {
	mu sync.Mutex
}

func (v *ReportGeneratorVisitor) VisitCompressor(c *Compressor) {
	v.mu.Lock()
	fmt.Println("Generating report for Compressor")
	data := c.ControlData.Data["compressor"].(CompressorData)
	data.ValueY = 20
	c.ControlData.Data["compressor"] = data
	v.mu.Unlock()

}

func (v *ReportGeneratorVisitor) VisitDispenser(d *Dispenser) {
	v.mu.Lock()
	fmt.Println("Generating report for dispenser")
	data := d.ControlData.Data["dispenser"].(DispenserData)
	data.ValueB = 241.123
	d.ControlData.Data["dispenser"] = data
	v.mu.Unlock()
}
