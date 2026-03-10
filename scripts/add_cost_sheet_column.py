#!/usr/bin/env python3
"""
Script to add a cost_sheet column to price_comparison_all_03062026.csv
based on model/provider discount rules.
Uses updated_price_per_unit when not null/empty, else price_per_unit.

- Anthropic models (anthropic, bedrock, azure) -> 8.5% less
- Gemini 2 and 2.5 (all except 3s and 3.5s) -> 10% less
- Gemini 3s and 3.5 -> 5% less
- OpenAI models via OpenAI -> 6% less
- OpenAI models via Azure -> 10% less
- No match -> empty
"""

import csv
from decimal import Decimal, InvalidOperation, ROUND_HALF_UP

# Discount multipliers: (1 - discount_percent)
ANTHROPIC_MULTIPLIER = Decimal("0.915")   # 8.5% less
GEMINI_2_2_5_MULTIPLIER = Decimal("0.90")   # 10% less
GEMINI_3_3_5_MULTIPLIER = Decimal("0.95")   # 5% less
OPENAI_MULTIPLIER = Decimal("0.94")       # 6% less
AZURE_OPENAI_MULTIPLIER = Decimal("0.90")   # 10% less


def parse_price(val):
    """Parse price string to Decimal. Returns None if invalid."""
    try:
        if not val or not val.strip():
            return None
        return Decimal(val.strip())
    except (InvalidOperation, ValueError):
        return None


def get_cost_sheet_value(feature_name: str, event_name: str, price_per_unit: str) -> str:
    """
    Determine cost sheet value based on feature/event naming and discount rules.
    Returns empty string if no match or invalid price.
    """
    price = parse_price(price_per_unit)
    if price is None:
        return ""

    # Use both feature_name and event_name for matching (event_name often has cleaner model info)
    combined = f"{feature_name} {event_name}".lower()

    # Check in order of specificity (more specific first)

    # 1. Azure OpenAI models -> 10% less
    if "azure-openai" in combined or combined.startswith("azure-openai-"):
        return format_decimal(price * AZURE_OPENAI_MULTIPLIER)

    # 2. OpenAI via OpenAI (not Azure) -> 6% less
    # Must be openai- prefixed but NOT azure-openai
    if (combined.startswith("openai-") or "openai-" in combined) and "azure-openai" not in combined:
        return format_decimal(price * OPENAI_MULTIPLIER)

    # 3. Anthropic models (any provider: anthropic, bedrock, azure) -> 8.5% less
    # anthropic-, anthropic-bedrock-, anthropic-azure-, or bedrock/azure with claude
    if (
        "anthropic-" in combined
        or ("bedrock" in combined and "claude" in combined)
        or ("azure" in combined and "claude" in combined)
    ):
        return format_decimal(price * ANTHROPIC_MULTIPLIER)

    # 4. Gemini 3 and 3.5 -> 5% less (check before Gemini 2/2.5)
    if "gemini-3" in combined or "gemini-3.5" in combined:
        return format_decimal(price * GEMINI_3_3_5_MULTIPLIER)

    # 5. Gemini 2 and 2.5 (all models except 3s and 3.5s) -> 10% less
    if "gemini-2" in combined or "gemini-2.5" in combined:
        return format_decimal(price * GEMINI_2_2_5_MULTIPLIER)

    return ""


def format_decimal(d):
    """Format Decimal preserving enough precision for pricing. Avoids scientific notation."""
    quantize = Decimal("0.000000000000001")
    q = d.quantize(quantize, rounding=ROUND_HALF_UP)
    # 15 decimal places to match original CSV style, no scientific notation
    return f"{float(q):.15f}"


def process_file(input_file: str, output_file: str, price_col: str = "price_per_unit", updated_col: str | None = "updated_price_per_unit", delimiter: str = ","):
    """Process a CSV file and add cost_sheet column."""
    rows = []
    with open(input_file, "r", encoding="utf-8") as f:
        reader = csv.DictReader(f, delimiter=delimiter)
        fieldnames = list(reader.fieldnames) + ["cost_sheet"]
        for row in reader:
            price_to_use = row.get(price_col, "")
            if updated_col and row.get(updated_col, "").strip():
                price_to_use = row.get(updated_col, "")
            cost_sheet = get_cost_sheet_value(
                row.get("feature_name", ""),
                row.get("event_name", ""),
                price_to_use,
            )
            row["cost_sheet"] = cost_sheet
            rows.append(row)

    with open(output_file, "w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames, delimiter=delimiter)
        writer.writeheader()
        writer.writerows(rows)

    matched = sum(1 for r in rows if r["cost_sheet"])
    print(f"Processed {len(rows)} rows")
    print(f"Matched and applied discount: {matched}")
    print(f"Output written to: {output_file}")


def main():
    import argparse
    parser = argparse.ArgumentParser(description="Add cost_sheet column to CSV")
    parser.add_argument("--missing-features", action="store_true", help="Process missing features file")
    args = parser.parse_args()

    base_dir = "/Users/prathamkhodwe/Work/flexprice"

    if args.missing_features:
        input_file = f"{base_dir}/missing_features_vapi_format_03062026.csv"
        output_file = f"{base_dir}/missing_features_vapi_format_with_cost_sheet_03062026.csv"
        process_file(input_file, output_file, updated_col=None, delimiter="\t")
    else:
        input_file = f"{base_dir}/price_comparison_all_03062026.csv"
        output_file = f"{base_dir}/price_comparison_all_with_cost_sheet_03062026.csv"
        process_file(input_file, output_file)


if __name__ == "__main__":
    main()
