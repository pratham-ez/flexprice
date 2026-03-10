#!/usr/bin/env python3
"""
Script to compare prices between latest_prices_03062026.csv and vapi_current_features_in_flexprice.csv
and generate a new CSV with updated prices where they differ.
"""

import csv
import os
from decimal import Decimal, InvalidOperation

def normalize_for_comparison(value_str):
    """
    Normalize a numeric string for comparison by converting to Decimal.
    This handles scientific notation and different representations of the same number.
    """
    try:
        if not value_str or value_str.strip() == '':
            return Decimal('0')
        return Decimal(value_str.strip())
    except (InvalidOperation, ValueError):
        return Decimal('0')

def read_latest_prices(filepath):
    """
    Read latest_prices_03062026.csv and create a mapping of slug -> (numerical_value_string, normalized_decimal)
    Stores both the original string representation and normalized value for comparison
    """
    price_map = {}
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            slug = row['slug'].strip()
            numerical_value = row['numerical_value'].strip()
            
            # Store both original string and normalized value
            normalized = normalize_for_comparison(numerical_value)
            price_map[slug] = (numerical_value, normalized)
    
    return price_map

def read_vapi_features(filepath):
    """
    Read vapi_current_features_in_flexprice.csv
    """
    features = []
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            features.append(row)
    
    return features

def compare_and_generate_output(vapi_features, latest_prices_map, output_files):
    """
    Compare the features and generate 3 output CSV files:
    1. All data combined
    2. Only rows with updates
    3. Only rows without updates
    Preserves original number format exactly as it appears in source files
    """
    rows_with_updates = []
    rows_without_updates = []
    
    for feature in vapi_features:
        feature_name = feature['feature_name'].strip()
        event_name = feature['event_name'].strip()
        aggregation_type = feature['aggregation_type'].strip()
        aggregation_field = feature['aggregation_field'].strip()
        price_per_unit = feature['price_per_unit'].strip()
        
        # Normalize current price for comparison
        current_price_normalized = normalize_for_comparison(price_per_unit)
        
        # Check if feature exists in latest prices
        if feature_name in latest_prices_map:
            new_price_str, new_price_normalized = latest_prices_map[feature_name]
            
            # Compare normalized values (but output original strings)
            if current_price_normalized != new_price_normalized:
                # Prices differ - add to updated list with original format preserved
                output_row = {
                    'feature_name': feature_name,
                    'event_name': event_name,
                    'aggregation_type': aggregation_type,
                    'aggregation_field': aggregation_field,
                    'price_per_unit': price_per_unit if price_per_unit else '0',
                    'updated_price_per_unit': new_price_str
                }
                rows_with_updates.append(output_row)
            else:
                # Prices are the same - no update needed
                output_row = {
                    'feature_name': feature_name,
                    'event_name': event_name,
                    'aggregation_type': aggregation_type,
                    'aggregation_field': aggregation_field,
                    'price_per_unit': price_per_unit if price_per_unit else '0',
                    'updated_price_per_unit': ''  # Empty since no update
                }
                rows_without_updates.append(output_row)
        else:
            # Feature not found in latest prices - no update
            output_row = {
                'feature_name': feature_name,
                'event_name': event_name,
                'aggregation_type': aggregation_type,
                'aggregation_field': aggregation_field,
                'price_per_unit': price_per_unit if price_per_unit else '0',
                'updated_price_per_unit': ''  # Empty since not found
            }
            rows_without_updates.append(output_row)
    
    # Define fieldnames
    fieldnames = ['feature_name', 'event_name', 'aggregation_type', 'aggregation_field', 'price_per_unit', 'updated_price_per_unit']
    
    # Write Sheet 1: All data (rows with updates first, then rows without)
    all_rows = rows_with_updates + rows_without_updates
    with open(output_files['all'], 'w', newline='', encoding='utf-8') as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(all_rows)
    
    # Write Sheet 2: Only rows with updates
    with open(output_files['with_updates'], 'w', newline='', encoding='utf-8') as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(rows_with_updates)
    
    # Write Sheet 3: Only rows without updates
    with open(output_files['without_updates'], 'w', newline='', encoding='utf-8') as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(rows_without_updates)
    
    return len(rows_with_updates), len(rows_without_updates)

def main():
    # File paths
    base_dir = '/Users/prathamkhodwe/Work/flexprice'
    latest_prices_file = os.path.join(base_dir, 'latest_prices_03062026.csv')
    vapi_features_file = os.path.join(base_dir, 'vapi_current_features_in_flexprice.csv')
    
    # Output files
    output_files = {
        'all': os.path.join(base_dir, 'price_comparison_all_03062026.csv'),
        'with_updates': os.path.join(base_dir, 'price_comparison_with_updates_03062026.csv'),
        'without_updates': os.path.join(base_dir, 'price_comparison_without_updates_03062026.csv')
    }
    
    print("Starting price comparison...")
    print(f"Reading latest prices from: {latest_prices_file}")
    latest_prices_map = read_latest_prices(latest_prices_file)
    print(f"Found {len(latest_prices_map)} entries in latest prices")
    
    print(f"\nReading VAPI features from: {vapi_features_file}")
    vapi_features = read_vapi_features(vapi_features_file)
    print(f"Found {len(vapi_features)} features in VAPI current features")
    
    print("\nComparing prices and generating output files...")
    rows_with_updates, rows_without_updates = compare_and_generate_output(
        vapi_features, 
        latest_prices_map, 
        output_files
    )
    
    print(f"\n✅ Output files generated:")
    print(f"\n   📊 Sheet 1 (All Data): {output_files['all']}")
    print(f"      - Total rows: {rows_with_updates + rows_without_updates}")
    print(f"      - Rows with updates listed first, then rows without updates")
    
    print(f"\n   🔄 Sheet 2 (With Updates): {output_files['with_updates']}")
    print(f"      - Rows with price updates: {rows_with_updates}")
    print(f"      - Contains only features where prices differ")
    
    print(f"\n   ✓ Sheet 3 (Without Updates): {output_files['without_updates']}")
    print(f"      - Rows without updates: {rows_without_updates}")
    print(f"      - Contains features where prices match or not found in latest_prices")
    
    print("\n" + "="*70)
    print("Summary:")
    print(f"  - Total features processed: {rows_with_updates + rows_without_updates}")
    print(f"  - Features with price updates: {rows_with_updates}")
    print(f"  - Features without price updates: {rows_without_updates}")
    print("="*70)

if __name__ == '__main__':
    main()
