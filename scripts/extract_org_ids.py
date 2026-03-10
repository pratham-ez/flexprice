#!/usr/bin/env python3
"""
Script to extract only the IDs from org_lookups.csv
Outputs a single-column CSV with just the organization IDs
"""

import csv
import os

def extract_ids(input_file, output_file):
    """
    Read org_lookups.csv and extract only the id column
    """
    ids = []
    
    with open(input_file, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            org_id = row['id'].strip()
            if org_id:  # Only add non-empty IDs
                ids.append(org_id)
    
    # Write IDs to output file
    with open(output_file, 'w', newline='', encoding='utf-8') as f:
        writer = csv.writer(f)
        writer.writerow(['id'])  # Header
        for org_id in ids:
            writer.writerow([org_id])
    
    return len(ids)

def main():
    # File paths
    input_file = '/Users/prathamkhodwe/Work/flexprice-front/org_lookups.csv'
    output_file = '/Users/prathamkhodwe/Work/flexprice/org_ids_output.csv'
    
    print("Extracting organization IDs...")
    print(f"Reading from: {input_file}")
    
    total_ids = extract_ids(input_file, output_file)
    
    print(f"\n✅ Output generated: {output_file}")
    print(f"   - Total IDs extracted: {total_ids}")
    
    print(f"\nFirst 10 IDs:")
    with open(output_file, 'r', encoding='utf-8') as f:
        reader = csv.reader(f)
        next(reader)  # Skip header
        for i, row in enumerate(reader, 1):
            if i <= 10:
                print(f"   {i}. {row[0]}")
            else:
                break
    
    if total_ids > 10:
        print(f"   ... and {total_ids - 10} more")
    
    print("\n" + "="*70)
    print("Summary:")
    print(f"  - Input file: {input_file}")
    print(f"  - Output file: {output_file}")
    print(f"  - Total IDs: {total_ids}")
    print("="*70)

if __name__ == '__main__':
    main()
