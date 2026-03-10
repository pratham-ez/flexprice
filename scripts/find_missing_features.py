#!/usr/bin/env python3
"""
Script to find features that exist in latest_prices_03062026.csv (slug column)
but are missing from vapi_current_features_in_flexprice.csv (feature_name column).

- Only considers rows where component_type == 'feature-term' (excludes feature-multiplier).
- Only considers rows with component_type=feature-term (excludes feature-multiplier).
Only considers rows with component_type=feature-term (excludes feature-multiplier).
Only considers rows with component_type=feature-term (excludes feature-multiplier).
Filters out false positives: features where the base (without metric suffix like
-characters, -channels, -durationMS, etc.) already exists in vapi.
"""

import csv
import os

# Metric suffixes that indicate a variant of a base feature.
# If base (feature minus suffix) exists in vapi, we don't count as truly missing.
KNOWN_METRIC_SUFFIXES = {
    'characters', 'channels', 'durationMS', 'completionTokens', 'promptTokens',
    'numCharacters', 'billable_value', 'cacheCreationInputTokens', 'cacheReadInputTokens',
    'cachedPromptTokens',
}


def read_latest_prices_features(filepath, component_type_filter='feature-term'):
    """
    Read latest_prices_03062026.csv and extract unique slugs (feature names).
    Only includes rows where component_type matches the filter (default: feature-term).
    Excludes feature-multiplier and other component types.
    """
    features = set()
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            comp_type = row.get('component_type', '').strip()
            if comp_type != component_type_filter:
                continue
            slug = row['slug'].strip()
            if slug:
                features.add(slug)
    
    return features


def read_vapi_features(filepath):
    """
    Read vapi_current_features_in_flexprice.csv and extract all unique feature names
    """
    features = set()
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            feature_name = row['feature_name'].strip()
            if feature_name:  # Only add non-empty feature names
                features.add(feature_name)
    
    return features


def get_base_feature(feature: str) -> str | None:
    """
    If the last segment is a known metric suffix, return the base (without it).
    Otherwise return None.
    """
    if '-' not in feature:
        return None
    parts = feature.rsplit('-', 1)
    if len(parts) != 2:
        return None
    base, suffix = parts
    if suffix in KNOWN_METRIC_SUFFIXES:
        return base
    return None


def is_base_in_vapi(feature: str, vapi_features: set) -> bool:
    """
    Check if the base feature (after stripping known metric suffix) exists in vapi.
    """
    base = get_base_feature(feature)
    if base is None:
        return False
    return base in vapi_features


def find_missing_features(latest_prices_features, vapi_features):
    """
    Find features that are in latest_prices but missing from vapi_features
    """
    missing = latest_prices_features - vapi_features
    return sorted(missing)


def filter_truly_missing(missing_features, vapi_features):
    """
    Exclude features where the base (without metric suffix) already exists in vapi.
    """
    truly_missing = []
    excluded = []
    for f in missing_features:
        if is_base_in_vapi(f, vapi_features):
            excluded.append(f)
        else:
            truly_missing.append(f)
    return truly_missing, excluded


def write_missing_features(missing_features, output_file, headers=None):
    """
    Write missing features to a CSV file
    """
    with open(output_file, 'w', newline='', encoding='utf-8') as f:
        writer = csv.writer(f)
        writer.writerow(headers or ['feature_name'])
        for feature in missing_features:
            row = [feature] if isinstance(feature, str) else feature
            writer.writerow(row)


def main():
    # File paths
    base_dir = '/Users/prathamkhodwe/Work/flexprice'
    latest_prices_file = os.path.join(base_dir, 'latest_prices_03062026.csv')
    vapi_features_file = os.path.join(base_dir, 'vapi_current_features_in_flexprice.csv')
    output_file_raw = os.path.join(base_dir, 'missing_features_03062026.csv')
    output_file_actual = os.path.join(base_dir, 'missing_features_actual_03062026.csv')
    
    print("Finding missing features (feature-term only, excluding feature-multiplier)...")
    print(f"Reading feature names from latest prices: {latest_prices_file}")
    latest_prices_features = read_latest_prices_features(latest_prices_file)
    print(f"Found {len(latest_prices_features)} unique features in latest_prices")
    
    print(f"\nReading feature names from VAPI features: {vapi_features_file}")
    vapi_features = read_vapi_features(vapi_features_file)
    print(f"Found {len(vapi_features)} unique features in vapi_current_features")
    
    print("\nFinding features that are in latest_prices but missing from vapi_current_features...")
    missing_features = find_missing_features(latest_prices_features, vapi_features)
    
    # Filter out false positives (base exists in vapi)
    truly_missing, excluded = filter_truly_missing(missing_features, vapi_features)
    
    # Write raw missing (original behavior)
    if missing_features:
        write_missing_features(missing_features, output_file_raw)
        print(f"\n📄 Raw output (all missing): {output_file_raw}")
        print(f"   - Count: {len(missing_features)}")
    
    # Write actual missing (filtered)
    if truly_missing:
        write_missing_features(truly_missing, output_file_actual)
        print(f"\n✅ Actual missing (base not in vapi): {output_file_actual}")
        print(f"   - Count: {len(truly_missing)}")
        print(f"\nFirst 15 truly missing features:")
        for i, feature in enumerate(truly_missing[:15], 1):
            print(f"   {i}. {feature}")
        if len(truly_missing) > 15:
            print(f"   ... and {len(truly_missing) - 15} more")
    else:
        print("\n✅ No truly missing features!")
    
    if excluded:
        print(f"\n📋 Excluded {len(excluded)} false positives (base exists in vapi):")
        for i, feature in enumerate(excluded[:10], 1):
            base = get_base_feature(feature)
            print(f"   {i}. {feature} → base '{base}' exists")
        if len(excluded) > 10:
            print(f"   ... and {len(excluded) - 10} more")
    
    print("\n" + "="*70)
    print("Summary:")
    print(f"  - Features in latest_prices: {len(latest_prices_features)}")
    print(f"  - Features in vapi_current_features: {len(vapi_features)}")
    print(f"  - Raw missing (before filter): {len(missing_features)}")
    print(f"  - Excluded (base exists): {len(excluded)}")
    print(f"  - Truly missing: {len(truly_missing)}")
    print("="*70)


if __name__ == '__main__':
    main()
