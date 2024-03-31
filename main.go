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

	// Initialize logger
	logger := utils.NewLogger()

	// Fetch data
	newData, err := data.FetchData(cfg)
	if err != nil {
		logger.Error("Failed to fetch data:", err)
		return
	}

	// Analyze data
	analysisResult := analysis.AnalyzeData(newData)

	// Make decisions based on analysis
	decisionResult := decision.MakeDecision(analysisResult)

	// Log actions taken
	logger.Info("Decision made:", decisionResult)

	fmt.Println("Anti-nuke bot execution complete.")
}
