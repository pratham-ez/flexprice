#!/usr/bin/env python3
"""
Takes the 66 missing features from missing_features_actual_03062026.csv and produces
a CSV in vapi feature format using data from latest_prices_03062026.csv.

Output columns: feature_name, event_name, aggregation_type, aggregation_field, price_per_unit
"""

import csv
import os

BASE_DIR = '/Users/prathamkhodwe/Work/flexprice'
MISSING_FEATURES_FILE = os.path.join(BASE_DIR, 'missing_features_actual_03062026.csv')
LATEST_PRICES_FILE = os.path.join(BASE_DIR, 'latest_prices_03062026.csv')
OUTPUT_FILE = os.path.join(BASE_DIR, 'missing_features_vapi_format_03062026.csv')


def load_latest_prices_lookup(filepath):
    """
    Build slug -> {event_name, aggregation_field, price_per_unit} from latest_prices.
    Only feature-term rows.
    """
    lookup = {}
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            if row.get('component_type', '').strip() != 'feature-term':
                continue
            slug = row['slug'].strip()
            if not slug:
                continue
            lookup[slug] = {
                'event_name': row['sku_logical_name'].strip(),
                'aggregation_field': row['component_feature'].strip(),
                'price_per_unit': row['numerical_value'].strip(),
            }
    return lookup


def load_missing_features(filepath):
    """Load feature names from missing_features_actual CSV."""
    features = []
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name = row.get('feature_name', '').strip()
            if name:
                features.append(name)
    return features


def main():
    print("Loading missing features...")
    missing_features = load_missing_features(MISSING_FEATURES_FILE)
    print(f"  Found {len(missing_features)} features")

    print("Loading latest prices (feature-term only)...")
    prices_lookup = load_latest_prices_lookup(LATEST_PRICES_FILE)
    print(f"  Found {len(prices_lookup)} price entries")

    results = []
    not_found = []
    for feature_name in missing_features:
        if feature_name not in prices_lookup:
            not_found.append(feature_name)
            continue
        data = prices_lookup[feature_name]
        results.append({
            'feature_name': feature_name,
            'event_name': data['event_name'],
            'aggregation_type': 'SUM',
            'aggregation_field': data['aggregation_field'],
            'price_per_unit': data['price_per_unit'],
        })

    # Write output
    fieldnames = ['feature_name', 'event_name', 'aggregation_type', 'aggregation_field', 'price_per_unit']
    with open(OUTPUT_FILE, 'w', newline='', encoding='utf-8') as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames, delimiter='\t')
        writer.writeheader()
        writer.writerows(results)

    print(f"\n✅ Output: {OUTPUT_FILE}")
    print(f"   Rows written: {len(results)}")
    if not_found:
        print(f"\n⚠️  Not found in latest_prices: {len(not_found)}")
        for f in not_found[:5]:
            print(f"     - {f}")
        if len(not_found) > 5:
            print(f"     ... and {len(not_found) - 5} more")


if __name__ == '__main__':
    main()
