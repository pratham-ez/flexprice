package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const millionDivisor = 1000000.0

type MethodPriceCol struct {
	ColName          string
	AggregationField string
}

var methodNames = []string{
	"modeling",
	"voicemail-detection",
	"modeling-completion",
	"knowledge-base",
}

var methodPriceColumns = []MethodPriceCol{
	{ColName: "PromptTokenCostPerMillion", AggregationField: "promptTokens"},
	{ColName: "CompletionTokenCostPerMillion", AggregationField: "completionTokens"},
	{ColName: "CachedPromptTokenCostPerMillion", AggregationField: "cachedPromptTokens"},
	{ColName: "CacheWriteTokenCostPerMillion", AggregationField: "cacheCreationInputTokens"},
	{ColName: "CacheReadTokenCostPerMillion", AggregationField: "cacheReadInputTokens"},
}

// GenerateMethodFeatures generates features for multiple method types from LLM pricing CSV
func GenerateMethodFeatures() error {
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

	requiredCols := []string{"Provider", "Model/Service"}
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

		provider := strings.TrimSpace(getCell(row, colIndex, "Provider"))
		model := strings.TrimSpace(getCell(row, colIndex, "Model/Service"))

		// skip empty identity rows
		if provider == "" && model == "" {
			continue
		}

		// For each of 5 price columns
		for _, pc := range methodPriceColumns {
			val := strings.TrimSpace(getCell(row, colIndex, pc.ColName))

			// If empty => no feature for this price type
			if val == "" {
				continue
			}

			// Parse the price value and divide by 1,000,000
			priceValue, ok := parseMethodFloat(val)
			if !ok {
				continue
			}
			pricePerUnit := formatMethodFloat(priceValue / millionDivisor)

			// For each of 4 method names => generate row
			for _, method := range methodNames {
				eventName := strings.ToLower(fmt.Sprintf("%s-%s-%s", provider, method, model))
				featureName := fmt.Sprintf("%s-%s", eventName, pc.AggregationField)

				out := []string{
					featureName,
					eventName,
					"SUM",
					pc.AggregationField,
					pricePerUnit,
				}

				if err := writer.Write(out); err != nil {
					return fmt.Errorf("failed to write output row: %w", err)
				}
				rowCount++
			}
		}
	}

	fmt.Printf("✅ Successfully generated method features\n")
	fmt.Printf("   Input:  %s\n", inputPath)
	fmt.Printf("   Output: %s\n", outputPath)
	fmt.Printf("   Processed %d output rows\n", rowCount)

	return nil
}

func getCell(row []string, colIndex map[string]int, colName string) string {
	idx, ok := colIndex[colName]
	if !ok {
		return ""
	}
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return row[idx]
}

func parseMethodFloat(s string) (float64, bool) {
	clean := strings.ReplaceAll(strings.TrimSpace(s), ",", "")
	f, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func formatMethodFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
