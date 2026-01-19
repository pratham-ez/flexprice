package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const million = 1000000.0

// GenerateTtsStFeatures generates TTS/STT features from synthesizers/transcribers pricing CSV
func GenerateTtsStFeatures() error {
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

	colIndex := map[string]int{}
	for i, col := range header {
		colIndex[strings.TrimSpace(col)] = i
	}

	requiredCols := []string{
		"Provider",
		"Model/Service",
		"ServiceType",
		"DurationCostPerSecond",
		"CharCost/M",
	}

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

		provider := strings.TrimSpace(getCellValue(row, colIndex, "Provider"))
		model := strings.TrimSpace(getCellValue(row, colIndex, "Model/Service"))
		serviceType := strings.TrimSpace(getCellValue(row, colIndex, "ServiceType"))

		if provider == "" || model == "" || serviceType == "" {
			continue
		}

		// method mapping
		method := ""
		if strings.EqualFold(serviceType, "TTS") || strings.EqualFold(serviceType, "synthesizers") {
			method = "synthesizers"
		} else if strings.EqualFold(serviceType, "STT") || strings.EqualFold(serviceType, "transcribers") {
			method = "transcribers"
		} else {
			// Only synthesizers/transcribers allowed
			continue
		}

		// pricing detection (prefer Duration if both exist)
		durationRaw := strings.TrimSpace(getCellValue(row, colIndex, "DurationCostPerSecond"))
		charRaw := strings.TrimSpace(getCellValue(row, colIndex, "CharCost/M"))

		aggregationField := ""
		pricePerUnit := ""

		if durationRaw != "" {
			// DurationCostPerSecond -> durationMS
			val, ok := parseFloatValue(durationRaw)
			if !ok {
				continue
			}
			aggregationField = "durationMS"
			pricePerUnit = formatFloatValue(val) // exact value
		} else if charRaw != "" {
			// CharCost/M -> numCharacters
			val, ok := parseFloatValue(charRaw)
			if !ok {
				continue
			}
			aggregationField = "numCharacters"
			pricePerUnit = formatFloatValue(val / million)
		} else {
			continue
		}

		// feature/event name => provider-method-model
		name := strings.ToLower(fmt.Sprintf("%s-%s-%s", provider, method, model))

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

	fmt.Printf("✅ Successfully generated TTS/STT features\n")
	fmt.Printf("   Input:  %s\n", inputPath)
	fmt.Printf("   Output: %s\n", outputPath)
	fmt.Printf("   Processed %d output rows\n", rowCount)

	return nil
}

func getCellValue(row []string, colIndex map[string]int, colName string) string {
	idx, ok := colIndex[colName]
	if !ok {
		return ""
	}
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return row[idx]
}

func parseFloatValue(s string) (float64, bool) {
	clean := strings.ReplaceAll(strings.TrimSpace(s), ",", "")
	f, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func formatFloatValue(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
