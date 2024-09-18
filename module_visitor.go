package main

import "fmt"

type ModuleVisitor interface {
	VisitCompressor(*Compressor)
	VisitDispenser(*Dispenser)
}

type ReportGeneratorVisitor struct{}

func (v *ReportGeneratorVisitor) VisitCompressor(c *Compressor) {
	fmt.Println("Generating report for Compressor")
	data := c.ControlData.Data["compressor"].(CompressorData)
	data.ValueY = 20
	c.ControlData.Data["compressor"] = data

}

func (v *ReportGeneratorVisitor) VisitDispenser(d *Dispenser) {
	fmt.Println("Generating report for dispenser")
	data := d.ControlData.Data["dispenser"].(DispenserData)
	data.ValueB = 241.123
	d.ControlData.Data["dispenser"] = data
}
