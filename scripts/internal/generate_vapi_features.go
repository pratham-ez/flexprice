package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// GenerateVapiFeatures generates VAPI features from provider/method/model CSV
func GenerateVapiFeatures() error {
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

	// Build column index map
	colIndex := map[string]int{}
	for i, col := range header {
		colIndex[strings.TrimSpace(col)] = i
	}

	// Required columns in input CSV
	requiredCols := []string{"provider_name", "method_name", "model_name"}
	for _, c := range requiredCols {
		if _, ok := colIndex[c]; !ok {
			return fmt.Errorf("missing required column in input CSV: %s", c)
		}
	}

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

	rowCount := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read row: %w", err)
		}

		provider := strings.TrimSpace(row[colIndex["provider_name"]])
		method := strings.TrimSpace(row[colIndex["method_name"]])
		model := strings.TrimSpace(row[colIndex["model_name"]])

		// Skip empty rows
		if provider == "" && method == "" && model == "" {
			continue
		}

		// Method name transformations
		displayMethod := method
		if method == "vapi-call-service" {
			displayMethod = "vapi-hosting"
		} else if method == "BEDROCK_LLM" {
			displayMethod = "modeling"
		}

		// Build feature/event name: provider-method-model (model optional)
		name := fmt.Sprintf("%s-%s", provider, displayMethod)
		if model != "" {
			name = fmt.Sprintf("%s-%s", name, model)
		}

		// Determine aggregation field based on method type
		aggregationField := ""
		pricePerUnit := "0"

		switch method {
		case "synthesizer":
			aggregationField = "durationMS"
		case "transcriber":
			aggregationField = "numCharacters"
			// Special pricing for cartesia-transcriber
			if provider == "cartesia" {
				pricePerUnit = "0.000000027"
			}
		case "smallest-call-service":
			aggregationField = "durationMS"
			pricePerUnit = "0.0000003333333333"
		case "vapi-call-service", "video-recording", "vapi-transport", "audio-recording":
			aggregationField = "durationMS"
		}

		out := []string{
			name,             // feature_name
			name,             // event_name
			"SUM",            // aggregation_type
			aggregationField, // aggregation_field
			pricePerUnit,     // price_per_unit
		}

		if err := writer.Write(out); err != nil {
			return fmt.Errorf("failed to write output row: %w", err)
		}
		rowCount++
	}

	fmt.Printf("✅ Successfully generated VAPI features\n")
	fmt.Printf("   Input:  %s\n", inputPath)
	fmt.Printf("   Output: %s\n", outputPath)
	fmt.Printf("   Processed %d rows\n", rowCount)

	return nil
}
