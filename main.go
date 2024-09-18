package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create modules that automatically reference the global data singleton
	compressor := NewCompressor()
	dispenser := NewDispenser()
	globalData := GetControlData()

	prettyJSON, err := json.MarshalIndent(globalData.Data, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate JSON:", err)
		return
	}

	fmt.Println("Global Data Pool:", string(prettyJSON))
	// Create a visitor (ReportGenerator)
	visitor := &ReportGeneratorVisitor{}

	// Use the visitor with different modules
	compressor.Accept(visitor)
	dispenser.Accept(visitor)

	// Print the global data to see updates

	prettyJSON, err = json.MarshalIndent(globalData.Data, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate JSON:", err)
		return
	}

	fmt.Println("Global Data Pool:", string(prettyJSON))
}
