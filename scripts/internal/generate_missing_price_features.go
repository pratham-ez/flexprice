package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type MissingPriceCol struct {
	AggregationField string
}

var missingPriceColumns = []MissingPriceCol{
	{AggregationField: "promptTokens"},
	{AggregationField: "completionTokens"},
	{AggregationField: "cachedPromptTokens"},
	{AggregationField: "cacheCreationInputTokens"},
	{AggregationField: "cacheReadInputTokens"},
}

var allowedEventNames = map[string]bool{
	"openai-synthesizer":                     true,
	"google-transcriber-gemini-2.0-flash":    true,
	"google-synthesizer-gemini-2.0-flash":    true,
	"google-transcriber-gemini-2.0-flash-lite": true,
	"google-transcriber-gemini-2.5-flash":    true,
	"google-transcriber-gemini-1.5-flash":    true,
	"google-synthesizer-gemini-2.0-flash-lite": true,
}

// GenerateMissingPriceFeatures generates features with 0 price from missing prices CSV
func GenerateMissingPriceFeatures() error {
	inputPath := os.Getenv("INPUT_FILE")
	outputPath := os.Getenv("OUTPUT_FILE")

	if inputPath == "" {
		return fmt.Errorf("INPUT_FILE environment variable is required")
	}

	if outputPath == "" {
		return fmt.Errorf("OUTPUT_FILE environment variable is required")
	}

	inFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inFile.Close()

	reader := csv.NewReader(inFile)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	// Map header -> index
	colIndex := map[string]int{}
	for i, col := range header {
		colIndex[strings.TrimSpace(col)] = i
	}

	requiredCols := []string{"event_name", "methodName"}
	for _, c := range requiredCols {
		if _, ok := colIndex[c]; !ok {
			return fmt.Errorf("missing required column: %s", c)
		}
	}

	// Output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Output header
	err = writer.Write([]string{
		"feature_name",
		"event_name",
		"aggregation_type",
		"aggregation_field",
		"price_per_unit",
	})
	if err != nil {
		return fmt.Errorf("failed to write output header: %w", err)
	}

	// Process rows
	rowCount := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read row: %w", err)
		}

		eventName := strings.TrimSpace(getMissingPriceCell(row, colIndex, "event_name"))
		methodName := strings.TrimSpace(getMissingPriceCell(row, colIndex, "methodName"))

		// skip empty rows
		if eventName == "" || methodName == "" {
			continue
		}

		// Only process rows with allowed event_name values
		if !allowedEventNames[eventName] {
			continue
		}

		// For each of 5 aggregation fields, generate a feature with 0 price
		for _, pc := range missingPriceColumns {
			featureName := fmt.Sprintf("%s-%s", eventName, pc.AggregationField)

			out := []string{
				featureName,
				eventName,
				"SUM",
				pc.AggregationField,
				"0",
			}

			if err := writer.Write(out); err != nil {
				return fmt.Errorf("failed to write output row: %w", err)
			}
			rowCount++
		}
	}

	fmt.Printf("✅ Successfully generated missing price features\n")
	fmt.Printf("   Input:  %s\n", inputPath)
	fmt.Printf("   Output: %s\n", outputPath)
	fmt.Printf("   Processed %d output rows\n", rowCount)

	return nil
}

func getMissingPriceCell(row []string, colIndex map[string]int, colName string) string {
	idx, ok := colIndex[colName]
	if !ok {
		return ""
	}
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return row[idx]
}
