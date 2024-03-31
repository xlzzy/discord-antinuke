package main

import (
	"antinuke-bot/analysis"
	"antinuke-bot/config"
	"antinuke-bot/data"
	"antinuke-bot/decision"
  "antinuke-bot/utils"
	"fmt"
)

func main() {
	// Load configurations
	cfg := config.LoadConfig()

	// Fetch data
	newData := data.FetchData(cfg)

	// Analyze data
	analysisResult := analysis.AnalyzeData(newData)

	// Make decisions based on analysis
	decision.MakeDecision(analysisResult)

	fmt.Println("Anti-nuke bot execution complete.")
}
