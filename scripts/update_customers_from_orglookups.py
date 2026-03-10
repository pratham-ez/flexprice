#!/usr/bin/env python3
"""
Script to update/create customers from org_lookups.csv
- Updates existing customers with domain name and email
- Creates new customers with automatic onboarding if not found
"""

import csv
import requests
import json
from typing import Optional, Dict, Any

# Configuration
TENANT_ID = "tenant_01KF5GXB4S7YKWH2Y3YQ1TEMQ3"
ENVIRONMENT_ID = "env_01KG4E6FR5YCNW0742N6CA1YD1"
API_KEY = "sk_01KK27QPXS5JF2JPXRJRAFRSCX"
BASE_URL = "https://us.api.flexprice.io/v1"

# Test mode - set to number of customers to process (None for all)
TEST_LIMIT = None  # Change to None to process all customers

# File paths
INPUT_FILE = "/Users/prathamkhodwe/Work/flexprice/org_lookups.csv"
OUTPUT_FILE = "/Users/prathamkhodwe/Work/flexprice/customer_update_results.csv"


def get_headers() -> Dict[str, str]:
    """Return headers for API requests"""
    return {
        "x-api-key": API_KEY,
        "x-tenant-id": TENANT_ID,
        "x-environment-id": ENVIRONMENT_ID,
        "Content-Type": "application/json"
    }


def extract_domain_from_email(email: str) -> str:
    """
    Extract domain from email address
    Example: pranay@notablehealth.com -> notablehealth.com
    """
    if not email or "@" not in email:
        return ""
    return email.split("@")[1].strip()


def get_customer_by_lookup_key(external_id: str) -> Optional[Dict[str, Any]]:
    """
    Get customer by external ID (lookup key)
    Returns customer data if found, None if not found
    """
    url = f"{BASE_URL}/customers/lookup/{external_id}"
    
    try:
        response = requests.get(url, headers=get_headers())
        
        if response.status_code == 200:
            return response.json()
        elif response.status_code == 404:
            return None
        else:
            print(f"  ⚠️  Error getting customer: {response.status_code} - {response.text}")
            return None
    except Exception as e:
        print(f"  ❌ Exception getting customer: {str(e)}")
        return None


def update_customer(customer_id: str, name: str, email: str) -> Optional[Dict[str, Any]]:
    """
    Update existing customer with new name and email
    """
    url = f"{BASE_URL}/customers/{customer_id}"
    
    payload = {
        "name": name,
        "email": email
    }
    
    try:
        response = requests.put(url, headers=get_headers(), json=payload)
        
        if response.status_code == 200:
            return response.json()
        else:
            print(f"  ❌ Error updating customer: {response.status_code} - {response.text}")
            return None
    except Exception as e:
        print(f"  ❌ Exception updating customer: {str(e)}")
        return None


def create_customer(external_id: str, name: str, email: str) -> Optional[Dict[str, Any]]:
    """
    Create new customer with automatic onboarding
    """
    url = f"{BASE_URL}/customers"
    
    payload = {
        "external_id": external_id,
        "name": name,
        "email": email
    }
    
    try:
        response = requests.post(url, headers=get_headers(), json=payload)
        
        if response.status_code == 201 or response.status_code == 200:
            return response.json()
        else:
            print(f"  ❌ Error creating customer: {response.status_code} - {response.text}")
            return None
    except Exception as e:
        print(f"  ❌ Exception creating customer: {str(e)}")
        return None


def process_customers(input_file: str, output_file: str, limit: Optional[int] = None):
    """
    Process customers from org_lookups.csv
    """
    results = []
    
    with open(input_file, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        rows = list(reader)
        
        # Apply limit if in test mode
        if limit:
            rows = rows[:limit]
            print(f"🧪 TEST MODE: Processing first {limit} customers\n")
        else:
            print(f"Processing all {len(rows)} customers\n")
        
        for i, row in enumerate(rows, 1):
            external_id = row['id'].strip()
            email = row['stripe_email'].strip()
            domain = extract_domain_from_email(email)
            
            print(f"[{i}/{len(rows)}] Processing: {email}")
            print(f"  External ID: {external_id}")
            print(f"  Domain: {domain}")
            
            if not domain:
                print(f"  ⚠️  Skipping - invalid email format")
                results.append({
                    'external_id': external_id,
                    'email': email,
                    'domain': domain,
                    'action': 'skipped',
                    'status': 'invalid_email',
                    'customer_id': '',
                    'error': 'Invalid email format'
                })
                print()
                continue
            
            # Try to get existing customer
            print(f"  🔍 Looking up customer...")
            existing_customer = get_customer_by_lookup_key(external_id)
            
            if existing_customer:
                # Customer exists - update
                customer_id = existing_customer.get('id')
                current_name = existing_customer.get('name', '')
                current_email = existing_customer.get('email', '')
                
                print(f"  ✓ Customer found (ID: {customer_id})")
                print(f"    Current name: {current_name}")
                print(f"    Current email: {current_email}")
                print(f"  🔄 Updating customer...")
                
                result = update_customer(customer_id, domain, email)
                
                if result:
                    print(f"  ✅ Customer updated successfully")
                    results.append({
                        'external_id': external_id,
                        'email': email,
                        'domain': domain,
                        'action': 'updated',
                        'status': 'success',
                        'customer_id': customer_id,
                        'error': ''
                    })
                else:
                    print(f"  ❌ Failed to update customer")
                    results.append({
                        'external_id': external_id,
                        'email': email,
                        'domain': domain,
                        'action': 'update_failed',
                        'status': 'error',
                        'customer_id': customer_id,
                        'error': 'Update API call failed'
                    })
            else:
                # Customer not found - create
                print(f"  ➕ Customer not found - creating new customer...")
                
                result = create_customer(external_id, domain, email)
                
                if result:
                    customer_id = result.get('id')
                    print(f"  ✅ Customer created successfully (ID: {customer_id})")
                    print(f"     Automatic onboarding triggered")
                    results.append({
                        'external_id': external_id,
                        'email': email,
                        'domain': domain,
                        'action': 'created',
                        'status': 'success',
                        'customer_id': customer_id,
                        'error': ''
                    })
                else:
                    print(f"  ❌ Failed to create customer")
                    results.append({
                        'external_id': external_id,
                        'email': email,
                        'domain': domain,
                        'action': 'create_failed',
                        'status': 'error',
                        'customer_id': '',
                        'error': 'Create API call failed'
                    })
            
            print()
    
    # Write results to CSV
    with open(output_file, 'w', newline='', encoding='utf-8') as f:
        fieldnames = ['external_id', 'email', 'domain', 'action', 'status', 'customer_id', 'error']
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(results)
    
    return results


def print_summary(results):
    """Print summary of operations"""
    total = len(results)
    updated = len([r for r in results if r['action'] == 'updated'])
    created = len([r for r in results if r['action'] == 'created'])
    skipped = len([r for r in results if r['action'] == 'skipped'])
    update_failed = len([r for r in results if r['action'] == 'update_failed'])
    create_failed = len([r for r in results if r['action'] == 'create_failed'])
    
    print("="*70)
    print("SUMMARY")
    print("="*70)
    print(f"Total customers processed: {total}")
    print(f"  ✅ Updated: {updated}")
    print(f"  ✅ Created: {created}")
    print(f"  ⚠️  Skipped: {skipped}")
    print(f"  ❌ Update failed: {update_failed}")
    print(f"  ❌ Create failed: {create_failed}")
    print("="*70)
    print(f"\nResults saved to: {OUTPUT_FILE}")


def main():
    print("="*70)
    print("CUSTOMER UPDATE/CREATE SCRIPT")
    print("="*70)
    print(f"Tenant ID: {TENANT_ID}")
    print(f"Environment ID: {ENVIRONMENT_ID}")
    print(f"Base URL: {BASE_URL}")
    print(f"Input file: {INPUT_FILE}")
    print(f"Output file: {OUTPUT_FILE}")
    print("="*70)
    print()
    
    results = process_customers(INPUT_FILE, OUTPUT_FILE, limit=TEST_LIMIT)
    print_summary(results)


if __name__ == '__main__':
    main()
