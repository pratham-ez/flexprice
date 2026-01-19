# Generate VAPI Features Script

This script generates VAPI features from a CSV file containing provider, method, and model information.

## Purpose

Converts a simple CSV with provider/method/model data into a features CSV that can be imported into the system.

## Input CSV Format

The input CSV must have the following columns:
- `provider_name` - The provider name (e.g., "OpenAI", "Azure", "Deepgram")
- `method_name` - The method type (e.g., "synthesizer", "transcriber")
- `model_name` - The model name (optional, can be empty)

### Example Input CSV

```csv
provider_name,method_name,model_name
OpenAI,synthesizer,tts-1
Azure,synthesizer,
Deepgram,transcriber,nova-2
AssemblyAI,transcriber,
```

## Output CSV Format

The script generates a CSV with the following columns:
- `feature_name` - Generated from provider-method-model (model optional)
- `event_name` - Same as feature_name
- `aggregation_type` - Always "SUM"
- `aggregation_field` - "durationMS" for synthesizer, "numCharacters" for transcriber
- `price_per_unit` - Always "0" (placeholder)

### Example Output CSV

```csv
feature_name,event_name,aggregation_type,aggregation_field,price_per_unit
OpenAI-synthesizer-tts-1,OpenAI-synthesizer-tts-1,SUM,durationMS,0
Azure-synthesizer,Azure-synthesizer,SUM,durationMS,0
Deepgram-transcriber-nova-2,Deepgram-transcriber-nova-2,SUM,numCharacters,0
AssemblyAI-transcriber,AssemblyAI-transcriber,SUM,numCharacters,0
```

## How to Run

### Basic Usage

```bash
cd /Users/prathamkhodwe/Work/flexprice

go run scripts/main.go -cmd generate-vapi-features \
  -input-file ./vapi_input.csv \
  -output-file ./vapi_features.csv
```

### Command Options

- `-cmd generate-vapi-features` - Specifies the command to run
- `-input-file <path>` - Path to the input CSV file
- `-output-file <path>` - Path where the output CSV will be saved

## Input File Location

You can place your input CSV anywhere, but a common location would be:
- In the project root: `/Users/prathamkhodwe/Work/flexprice/vapi_input.csv`
- Or in a data directory: `/Users/prathamkhodwe/Work/flexprice/data/vapi_input.csv`

## Rules

1. **Model Name is Optional**: If `model_name` is empty, the feature/event name will be `provider-method`
2. **Aggregation Field Logic**:
   - `method_name = "synthesizer"` → `aggregation_field = "durationMS"`
   - `method_name = "transcriber"` → `aggregation_field = "numCharacters"`
   - Other methods → `aggregation_field` is empty
3. **Empty Rows**: Rows where all three columns (provider, method, model) are empty will be skipped

## Example Workflow

1. Create your input CSV with provider/method/model data:
```bash
cat > vapi_input.csv << EOF
provider_name,method_name,model_name
OpenAI,synthesizer,tts-1
Azure,synthesizer,
Deepgram,transcriber,nova-2
EOF
```

2. Run the script:
```bash
go run scripts/main.go -cmd generate-vapi-features \
  -input-file ./vapi_input.csv \
  -output-file ./vapi_features.csv
```

3. Check the output:
```bash
cat vapi_features.csv
```

## Notes

- The script uses the same `-input-file` and `-output-file` flags as the `convert-pricing` command
- All feature names and event names are identical in the output
- The `price_per_unit` is set to "0" as a placeholder - you can update prices later if needed
