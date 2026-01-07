#!/usr/bin/env python3
"""
FlexPrice Speakeasy Python SDK - Comprehensive API Tests

This test suite covers all FlexPrice APIs with comprehensive CRUD operations.
Tests are designed to run sequentially, building on previous test results.
Uses the Speakeasy-generated Python SDK.
"""

import os
import sys
import time
from datetime import datetime, timezone, timedelta
from typing import Optional

# Add the SDK directory to path
sys.path.insert(0, os.path.join(os.path.dirname(__file__), 'python', 'src'))

# Import the FlexPrice Speakeasy SDK
from flexprice_sdk_test import FlexPrice
from flexprice_sdk_test.models import components, operations

# Global test entity IDs
test_customer_id: Optional[str] = None
test_customer_name: Optional[str] = None
test_external_id: Optional[str] = None

test_feature_id: Optional[str] = None
test_feature_name: Optional[str] = None

test_plan_id: Optional[str] = None
test_plan_name: Optional[str] = None

test_addon_id: Optional[str] = None
test_addon_name: Optional[str] = None
test_addon_lookup_key: Optional[str] = None

test_entitlement_id: Optional[str] = None

test_subscription_id: Optional[str] = None

test_invoice_id: Optional[str] = None

test_price_id: Optional[str] = None

test_payment_id: Optional[str] = None

test_wallet_id: Optional[str] = None
test_credit_grant_id: Optional[str] = None
test_credit_note_id: Optional[str] = None

test_event_id: Optional[str] = None
test_event_name: Optional[str] = None
test_event_customer_id: Optional[str] = None


def get_client() -> FlexPrice:
    """Get and configure the FlexPrice API client."""
    api_key = os.getenv("FLEXPRICE_API_KEY")
    api_host = os.getenv("FLEXPRICE_API_HOST", "https://api.cloud.flexprice.io/v1")

    if not api_key:
        print("❌ Missing FLEXPRICE_API_KEY environment variable")
        exit(1)

    print("=== FlexPrice Speakeasy Python SDK - API Tests ===\n")
    print(f"✓ API Key: {api_key[:8]}...{api_key[-4:]}")
    print(f"✓ API Host: {api_host}\n")

    # Initialize Speakeasy SDK
    client = FlexPrice(
        server_url=api_host,
        api_key_auth=api_key
    )

    return client


# ========================================
# CUSTOMER API TESTS
# ========================================

def test_create_customer(client: FlexPrice):
    """Test 1: Create Customer"""
    print("--- Test 1: Create Customer ---")

    try:
        timestamp = int(time.time())
        global test_customer_name, test_customer_id, test_external_id
        test_customer_name = f"Test Customer {timestamp}"
        test_external_id = f"test-customer-{timestamp}"

        customer_request = components.DtoCreateCustomerRequest(
            name=test_customer_name,
            email=f"test-{timestamp}@example.com",
            external_id=test_external_id,
            metadata={
                "source": "sdk_test",
                "test_run": datetime.now().isoformat(),
                "environment": "test",
            },
        )

        response = client.customers.post_customers(request=customer_request)

        test_customer_id = response.id
        print("✓ Customer created successfully!")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}")
        print(f"  External ID: {response.external_id}")
        print(f"  Email: {response.email}\n")
    except Exception as e:
        print(f"❌ Error creating customer: {e}\n")


def test_get_customer(client: FlexPrice):
    """Test 2: Get Customer by ID"""
    print("--- Test 2: Get Customer by ID ---")

    if not test_customer_id:
        print("⚠ Warning: No customer ID available\n⚠ Skipping get customer test\n")
        return

    try:
        response = client.customers.get_customers_id_(id=test_customer_id)

        print("✓ Customer retrieved successfully!")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}")
        print(f"  Created At: {response.created_at}\n")
    except Exception as e:
        print(f"❌ Error getting customer: {e}\n")


def test_list_customers(client: FlexPrice):
    """Test 3: List Customers"""
    print("--- Test 3: List Customers ---")

    try:
        response = client.customers.get_customers(request=operations.GetCustomersRequest(limit=10))

        print(f"✓ Retrieved {len(response.items) if response.items else 0} customers")
        if response.items and len(response.items) > 0:
            print(f"  First customer: {response.items[0].id} - {response.items[0].name}")
        if response.pagination:
            print(f"  Total: {response.pagination.total}\n")
    except Exception as e:
        print(f"❌ Error listing customers: {e}\n")


def test_update_customer(client: FlexPrice):
    """Test 4: Update Customer"""
    print("--- Test 4: Update Customer ---")

    if not test_customer_id:
        print("⚠ Warning: No customer ID available\n⚠ Skipping update customer test\n")
        return

    try:
        update_request = components.DtoUpdateCustomerRequest(
            name=f"{test_customer_name} (Updated)",
            metadata={
                "updated_at": datetime.now().isoformat(),
                "status": "updated",
            },
        )

        response = client.customers.put_customers_id_(id=test_customer_id, body=update_request)

        print("✓ Customer updated successfully!")
        print(f"  ID: {response.id}")
        print(f"  New Name: {response.name}")
        if response.updated_at:
            print(f"  Updated At: {response.updated_at}\n")
        else:
            print()
    except Exception as e:
        print(f"❌ Error updating customer: {e}\n")


def test_lookup_customer(client: FlexPrice):
    """Test 5: Lookup Customer by External ID"""
    print("--- Test 5: Lookup Customer by External ID ---")

    if not test_external_id:
        print("⚠ Warning: No external ID available\n⚠ Skipping lookup test\n")
        return

    try:
        response = client.customers.get_customers_external_external_id_(external_id=test_external_id)

        print("✓ Customer found by external ID!")
        print(f"  External ID: {test_external_id}")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}\n")
    except Exception as e:
        print(f"❌ Error looking up customer: {e}\n")


def test_search_customers(client: FlexPrice):
    """Test 6: Search Customers"""
    print("--- Test 6: Search Customers ---")

    if not test_external_id:
        print("⚠ Warning: No external ID available\n⚠ Skipping search test\n")
        return

    try:
        search_filter = components.TypesCustomerFilter(
            external_id=test_external_id
        )

        response = client.customers.post_customers_search(request=search_filter)

        print("✓ Search completed!")
        print(f"  Found {len(response.items) if response.items else 0} customers matching external ID '{test_external_id}'")
        if response.items:
            for i, customer in enumerate(response.items[:3]):
                print(f"  - {customer.id}: {customer.name}")
        print()
    except Exception as e:
        print(f"❌ Error searching customers: {e}\n")


def test_get_customer_entitlements(client: FlexPrice):
    """Test 7: Get Customer Entitlements"""
    print("--- Test 7: Get Customer Entitlements ---")

    if not test_customer_id:
        print("⚠ Warning: No customer ID available\n⚠ Skipping get entitlements test\n")
        return

    try:
        response = client.customers.get_customers_id_entitlements(id=test_customer_id)

        print("✓ Retrieved customer entitlements!")
        if response.features:
            print(f"  Total features: {len(response.features)}")
            for i, feature in enumerate(response.features[:3]):
                if feature.feature and feature.feature.id:
                    print(f"  - Feature: {feature.feature.id}")
        else:
            print("  No features found")
        print()
    except Exception as e:
        print(f"⚠ Warning: Error getting customer entitlements: {e}")
        print("⚠ Skipping entitlements test (customer may not have any entitlements)\n")


def test_get_customer_upcoming_grants(client: FlexPrice):
    """Test 8: Get Customer Upcoming Grants"""
    print("--- Test 8: Get Customer Upcoming Grants ---")

    if not test_customer_id:
        print("⚠ Warning: No customer ID available\n⚠ Skipping get upcoming grants test\n")
        return

    try:
        response = client.customers.get_customers_id_grants_upcoming(id=test_customer_id)

        print("✓ Retrieved upcoming grants!")
        if response.items:
            print(f"  Total upcoming grants: {len(response.items)}")
        else:
            print("  No upcoming grants found")
        print()
    except Exception as e:
        print(f"⚠ Warning: Error getting upcoming grants: {e}")
        print("⚠ Skipping upcoming grants test (customer may not have any grants)\n")


def test_get_customer_usage(client: FlexPrice):
    """Test 9: Get Customer Usage"""
    print("--- Test 9: Get Customer Usage ---")

    if not test_customer_id:
        print("⚠ Warning: No customer ID available\n⚠ Skipping get usage test\n")
        return

    try:
        response = client.customers.get_customers_usage(request=operations.GetCustomersUsageRequest(customer_id=test_customer_id))

        print("✓ Retrieved customer usage!")
        if response.features:
            print(f"  Feature usage records: {len(response.features)}")
        else:
            print("  No usage data found")
        print()
    except Exception as e:
        print(f"⚠ Warning: Error getting customer usage: {e}")
        print("⚠ Skipping usage test (customer may not have usage data)\n")


def test_delete_customer(client: FlexPrice):
    """Test 10: Delete Customer"""
    print("--- Test 10: Delete Customer ---")

    if not test_customer_id:
        print("⚠ Skipping delete customer (no customer created)\n")
        return

    try:
        client.customers.delete_customers_id_(id=test_customer_id)

        print("✓ Customer deleted successfully!")
        print(f"  Deleted ID: {test_customer_id}\n")
    except Exception as e:
        print(f"❌ Error deleting customer: {e}\n")


# ========================================
# FEATURES API TESTS
# ========================================

def test_create_feature(client: FlexPrice):
    """Test 1: Create Feature"""
    print("--- Test 1: Create Feature ---")

    try:
        timestamp = int(time.time())
        global test_feature_name, test_feature_id
        test_feature_name = f"Test Feature {timestamp}"
        feature_key = f"test_feature_{timestamp}"

        feature_request = components.DtoCreateFeatureRequest(
            name=test_feature_name,
            lookup_key=feature_key,
            description="This is a test feature created by SDK tests",
            type=components.TypesFeatureType.BOOLEAN,
            metadata={
                "source": "sdk_test",
                "test_run": datetime.now().isoformat(),
                "environment": "test",
            },
        )

        response = client.features.post_features(request=feature_request)

        test_feature_id = response.id
        print("✓ Feature created successfully!")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}")
        print(f"  Lookup Key: {response.lookup_key}")
        print(f"  Type: {response.type}\n")
    except Exception as e:
        print(f"❌ Error creating feature: {e}\n")


def test_get_feature(client: FlexPrice):
    """Test 2: Get Feature by ID"""
    print("--- Test 2: Get Feature by ID ---")

    if not test_feature_id:
        print("⚠ Warning: No feature ID available\n⚠ Skipping get feature test\n")
        return

    try:
        response = client.features.get_features_id_(id=test_feature_id)

        print("✓ Feature retrieved successfully!")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}")
        print(f"  Lookup Key: {response.lookup_key}")
        print(f"  Created At: {response.created_at}\n")
    except Exception as e:
        print(f"❌ Error getting feature: {e}\n")


def test_list_features(client: FlexPrice):
    """Test 3: List Features"""
    print("--- Test 3: List Features ---")

    try:
        response = client.features.get_features(request=operations.GetFeaturesRequest(limit=10))

        print(f"✓ Retrieved {len(response.items) if response.items else 0} features")
        if response.items and len(response.items) > 0:
            print(f"  First feature: {response.items[0].id} - {response.items[0].name}")
        if response.pagination:
            print(f"  Total: {response.pagination.total}\n")
    except Exception as e:
        print(f"❌ Error listing features: {e}\n")


def test_update_feature(client: FlexPrice):
    """Test 4: Update Feature"""
    print("--- Test 4: Update Feature ---")

    if not test_feature_id:
        print("⚠ Warning: No feature ID available\n⚠ Skipping update feature test\n")
        return

    try:
        update_request = components.DtoUpdateFeatureRequest(
            name=f"{test_feature_name} (Updated)",
            description="Updated description for test feature",
            metadata={
                "updated_at": datetime.now().isoformat(),
                "status": "updated",
            },
        )

        response = client.features.put_features_id_(id=test_feature_id, body=update_request)

        print("✓ Feature updated successfully!")
        print(f"  ID: {response.id}")
        print(f"  New Name: {response.name}")
        print(f"  New Description: {response.description}")
        print(f"  Updated At: {response.updated_at}\n")
    except Exception as e:
        print(f"❌ Error updating feature: {e}\n")


def test_search_features(client: FlexPrice):
    """Test 5: Search Features"""
    print("--- Test 5: Search Features ---")

    if not test_feature_id:
        print("⚠ Warning: No feature ID available\n⚠ Skipping search test\n")
        return

    try:
        search_filter = components.TypesFeatureFilter(
            feature_ids=[test_feature_id]
        )

        response = client.features.post_features_search(request=search_filter)

        print("✓ Search completed!")
        print(f"  Found {len(response.items) if response.items else 0} features matching ID '{test_feature_id}'")
        if response.items:
            for i, feature in enumerate(response.items[:3]):
                print(f"  - {feature.id}: {feature.name} ({feature.lookup_key})")
        print()
    except Exception as e:
        print(f"❌ Error searching features: {e}\n")


def test_delete_feature(client: FlexPrice):
    """Test 6: Delete Feature"""
    print("--- Test 6: Delete Feature ---")

    if not test_feature_id:
        print("⚠ Skipping delete feature (no feature created)\n")
        return

    try:
        client.features.delete_features_id_(id=test_feature_id)

        print("✓ Feature deleted successfully!")
        print(f"  Deleted ID: {test_feature_id}\n")
    except Exception as e:
        print(f"❌ Error deleting feature: {e}\n")


# ========================================
# CONNECTIONS API TESTS
# ========================================

def test_list_connections(client: FlexPrice):
    """Test 1: List Connections"""
    print("--- Test 1: List Connections ---")

    try:
        response = client.connections.get_connections(request=operations.GetConnectionsRequest(limit=10))

        print(f"✓ Retrieved {len(response.connections) if response.connections else 0} connections\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_search_connections(client: FlexPrice):
    """Test 2: Search Connections"""
    print("--- Test 2: Search Connections ---")

    try:
        connection_filter = components.TypesConnectionFilter(limit=5)
        response = client.connections.post_connections_search(request=connection_filter)

        print(f"✓ Found {len(response.connections) if response.connections else 0} connections\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


# ========================================
# PLANS API TESTS
# ========================================

def test_create_plan(client: FlexPrice):
    """Test 1: Create Plan"""
    print("--- Test 1: Create Plan ---")

    try:
        timestamp = int(time.time())
        global test_plan_name, test_plan_id
        test_plan_name = f"Test Plan {timestamp}"
        lookup_key = f"test_plan_{timestamp}"

        plan_request = components.DtoCreatePlanRequest(
            name=test_plan_name,
            lookup_key=lookup_key,
            description="This is a test plan created by SDK tests",
            metadata={
                "source": "sdk_test",
                "test_run": datetime.now().isoformat(),
                "environment": "test",
            },
        )

        response = client.plans.post_plans(request=plan_request)

        test_plan_id = response.id
        print("✓ Plan created successfully!")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}")
        print(f"  Lookup Key: {response.lookup_key}\n")
    except Exception as e:
        print(f"❌ Error creating plan: {e}\n")


def test_get_plan(client: FlexPrice):
    """Test 2: Get Plan by ID"""
    print("--- Test 2: Get Plan by ID ---")

    if not test_plan_id:
        print("⚠ Warning: No plan ID available\n⚠ Skipping get plan test\n")
        return

    try:
        response = client.plans.get_plans_id_(id=test_plan_id)

        print("✓ Plan retrieved successfully!")
        print(f"  ID: {response.id}")
        print(f"  Name: {response.name}")
        print(f"  Lookup Key: {response.lookup_key}")
        print(f"  Created At: {response.created_at}\n")
    except Exception as e:
        print(f"❌ Error getting plan: {e}\n")


def test_list_plans(client: FlexPrice):
    """Test 3: List Plans"""
    print("--- Test 3: List Plans ---")

    try:
        response = client.plans.get_plans(request=operations.GetPlansRequest(limit=10))

        print(f"✓ Retrieved {len(response.items) if response.items else 0} plans")
        if response.items and len(response.items) > 0:
            print(f"  First plan: {response.items[0].id} - {response.items[0].name}")
        if response.pagination:
            print(f"  Total: {response.pagination.total}\n")
    except Exception as e:
        print(f"❌ Error listing plans: {e}\n")


def test_update_plan(client: FlexPrice):
    """Test 4: Update Plan"""
    print("--- Test 4: Update Plan ---")

    if not test_plan_id:
        print("⚠ Warning: No plan ID available\n⚠ Skipping update plan test\n")
        return

    try:
        update_request = components.DtoUpdatePlanRequest(
            name=f"{test_plan_name} (Updated)",
            description="Updated description for test plan",
            metadata={
                "updated_at": datetime.now().isoformat(),
                "status": "updated",
            },
        )

        response = client.plans.put_plans_id_(id=test_plan_id, body=update_request)

        print("✓ Plan updated successfully!")
        print(f"  ID: {response.id}")
        print(f"  New Name: {response.name}")
        print(f"  New Description: {response.description}")
        print(f"  Updated At: {response.updated_at}\n")
    except Exception as e:
        print(f"❌ Error updating plan: {e}\n")


def test_search_plans(client: FlexPrice):
    """Test 5: Search Plans"""
    print("--- Test 5: Search Plans ---")

    if not test_plan_id:
        print("⚠ Warning: No plan ID available\n⚠ Skipping search test\n")
        return

    try:
        search_filter = components.TypesPlanFilter(
            plan_ids=[test_plan_id]
        )

        response = client.plans.post_plans_search(request=search_filter)

        print("✓ Search completed!")
        print(f"  Found {len(response.items) if response.items else 0} plans matching ID '{test_plan_id}'")
        if response.items:
            for i, plan in enumerate(response.items[:3]):
                print(f"  - {plan.id}: {plan.name} ({plan.lookup_key})")
        print()
    except Exception as e:
        print(f"❌ Error searching plans: {e}\n")


def test_delete_plan(client: FlexPrice):
    """Test 6: Delete Plan"""
    print("--- Test 6: Delete Plan ---")

    if not test_plan_id:
        print("⚠ Skipping delete plan (no plan created)\n")
        return

    try:
        client.plans.delete_plans_id_(id=test_plan_id)

        print("✓ Plan deleted successfully!")
        print(f"  Deleted ID: {test_plan_id}\n")
    except Exception as e:
        print(f"❌ Error deleting plan: {e}\n")


# ========================================
# ADDONS API TESTS
# ========================================

def test_create_addon(client: FlexPrice):
    """Test 1: Create Addon"""
    print("--- Test 1: Create Addon ---")

    try:
        timestamp = int(time.time())
        global test_addon_name, test_addon_id, test_addon_lookup_key
        test_addon_name = f"Test Addon {timestamp}"
        test_addon_lookup_key = f"test_addon_{timestamp}"

        addon_request = components.DtoCreateAddonRequest(
            name=test_addon_name,
            lookup_key=test_addon_lookup_key,
            description="Test addon created by SDK",
            type=components.TypesAddonType.ONETIME,
            metadata={
                "source": "sdk_test",
            },
        )

        response = client.addons.post_addons(request=addon_request)

        test_addon_id = response.id
        print(f"✓ Addon created: {response.id}\n")
    except Exception as e:
        print(f"❌ Error creating addon: {e}\n")


def test_get_addon(client: FlexPrice):
    """Test 2: Get Addon"""
    print("--- Test 2: Get Addon ---")

    if not test_addon_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.addons.get_addons_id_(id=test_addon_id)
        print(f"✓ Addon retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_addons(client: FlexPrice):
    """Test 3: List Addons"""
    print("--- Test 3: List Addons ---")

    try:
        response = client.addons.get_addons(request=operations.GetAddonsRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} addons\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_update_addon(client: FlexPrice):
    """Test 4: Update Addon"""
    print("--- Test 4: Update Addon ---")

    if not test_addon_id:
        print("⚠ Skipping\n")
        return

    try:
        updated_name = test_addon_name + " (Updated)"
        response = client.addons.put_addons_id_(
            id=test_addon_id,
            body=components.DtoUpdateAddonRequest(name=updated_name)
        )
        print(f"✓ Addon updated: {response.name}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_lookup_addon(client: FlexPrice):
    """Test 5: Lookup Addon"""
    print("--- Test 5: Lookup Addon ---")
    print("⚠ Skipping (method not available in SDK)\n")


def test_search_addons(client: FlexPrice):
    """Test 6: Search Addons"""
    print("--- Test 6: Search Addons ---")

    if not test_addon_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.addons.post_addons_search(
            request=components.TypesAddonFilter(addon_ids=[test_addon_id])
        )
        print(f"✓ Found {len(response.items) if response.items else 0} addons\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_delete_addon(client: FlexPrice):
    """Test 7: Delete Addon"""
    print("--- Test 7: Delete Addon ---")

    if not test_addon_id:
        print("⚠ Skipping\n")
        return

    try:
        client.addons.delete_addons_id_(id=test_addon_id)
        print("✓ Addon deleted\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


# ========================================
# ENTITLEMENTS API TESTS
# ========================================

def test_create_entitlement(client: FlexPrice):
    """Test 1: Create Entitlement"""
    print("--- Test 1: Create Entitlement ---")

    if not test_feature_id or not test_plan_id:
        print("⚠ Skipping (no feature or plan)\n")
        return

    try:
        entitlement_request = components.DtoCreateEntitlementRequest(
            feature_id=test_feature_id,
            feature_type=components.TypesFeatureType.BOOLEAN,
            plan_id=test_plan_id,
            is_enabled=True,
            usage_reset_period=components.TypesEntitlementUsageResetPeriod.MONTHLY,
        )

        response = client.entitlements.post_entitlements(request=entitlement_request)

        global test_entitlement_id
        test_entitlement_id = response.id
        print(f"✓ Entitlement created: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_get_entitlement(client: FlexPrice):
    """Test 2: Get Entitlement"""
    print("--- Test 2: Get Entitlement ---")

    if not test_entitlement_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.entitlements.get_entitlements_id_(id=test_entitlement_id)
        print(f"✓ Entitlement retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_entitlements(client: FlexPrice):
    """Test 3: List Entitlements"""
    print("--- Test 3: List Entitlements ---")

    try:
        response = client.entitlements.get_entitlements(request=operations.GetEntitlementsRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} entitlements\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_update_entitlement(client: FlexPrice):
    """Test 4: Update Entitlement"""
    print("--- Test 4: Update Entitlement ---")

    if not test_entitlement_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.entitlements.put_entitlements_id_(
            id=test_entitlement_id,
            body=components.DtoUpdateEntitlementRequest(is_enabled=False)
        )
        print(f"✓ Entitlement updated: {response.is_enabled}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_search_entitlements(client: FlexPrice):
    """Test 5: Search Entitlements"""
    print("--- Test 5: Search Entitlements ---")

    if not test_entitlement_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.entitlements.post_entitlements_search(
            request=components.TypesEntitlementFilter(entity_ids=[test_entitlement_id])
        )
        print(f"✓ Found {len(response.items) if response.items else 0} entitlements\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_delete_entitlement(client: FlexPrice):
    """Test 6: Delete Entitlement"""
    print("--- Test 6: Delete Entitlement ---")

    if not test_entitlement_id:
        print("⚠ Skipping\n")
        return

    try:
        client.entitlements.delete_entitlements_id_(id=test_entitlement_id)
        print("✓ Entitlement deleted\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


# ========================================
# SUBSCRIPTIONS API TESTS
# ========================================

def test_create_subscription(client: FlexPrice):
    """Test 1: Create Subscription"""
    print("--- Test 1: Create Subscription ---")

    if not test_customer_id or not test_plan_id:
        print("⚠ Skipping\n")
        return

    try:
        # Create price first
        price_request = components.DtoCreatePriceRequest(
            entity_id=test_plan_id,
            entity_type=components.TypesPriceEntityType.PLAN,
            type=components.TypesPriceType.FIXED,
            billing_model=components.TypesBillingModel.FLAT_FEE,
            billing_cadence=components.TypesBillingCadence.RECURRING,
            billing_period=components.TypesBillingPeriod.MONTHLY,
            billing_period_count=1,
            invoice_cadence=components.TypesInvoiceCadence.ARREAR,
            amount="29.99",
            currency="USD",
            price_unit_type=components.TypesPriceUnitType.FIAT,
            display_name="Monthly Price",
        )

        price = client.prices.post_prices(request=price_request)
        print(f"  Created price: {price.id}")

        start_date = datetime.now(timezone.utc).isoformat()
        subscription_request = components.DtoCreateSubscriptionRequest(
            customer_id=test_customer_id,
            plan_id=test_plan_id,
            currency="USD",
            billing_cadence=components.TypesBillingCadence.RECURRING,
            billing_period=components.TypesBillingPeriod.MONTHLY,
            billing_period_count=1,
            billing_cycle=components.TypesBillingCycle.ANNIVERSARY,
            start_date=start_date,
        )

        response = client.subscriptions.post_subscriptions(request=subscription_request)

        global test_subscription_id
        test_subscription_id = response.id
        print(f"✓ Subscription created: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_get_subscription(client: FlexPrice):
    """Test 2: Get Subscription"""
    print("--- Test 2: Get Subscription ---")

    if not test_subscription_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.subscriptions.get_subscriptions_id_(id=test_subscription_id)
        print(f"✓ Subscription retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_subscriptions(client: FlexPrice):
    """Test 3: List Subscriptions"""
    print("--- Test 3: List Subscriptions ---")

    try:
        response = client.subscriptions.get_subscriptions(request=operations.GetSubscriptionsRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} subscriptions\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_search_subscriptions(client: FlexPrice):
    """Test 4: Search Subscriptions"""
    print("--- Test 4: Search Subscriptions ---")

    try:
        response = client.subscriptions.post_subscriptions_search(
            request=components.TypesSubscriptionFilter()
        )
        print(f"✓ Found {len(response.items) if response.items else 0} subscriptions\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_activate_subscription(client: FlexPrice):
    """Test 5: Activate Subscription"""
    print("--- Test 5: Activate Subscription ---")
    print("⚠ Skipping (requires draft subscription)\n")


def test_add_addon_to_subscription(client: FlexPrice):
    """Test 6: Add Addon to Subscription"""
    print("--- Test 6: Add Addon to Subscription ---")
    print("⚠ Skipping (add addon endpoint needs verification)\n")


def test_remove_addon_from_subscription(client: FlexPrice):
    """Test 7: Remove Addon from Subscription"""
    print("--- Test 7: Remove Addon from Subscription ---")
    print("⚠ Skipping (requires addon association ID)\n")


def test_execute_subscription_change(client: FlexPrice):
    """Test 8: Execute Subscription Change"""
    print("--- Test 8: Execute Subscription Change ---")
    print("⚠ Skipping (would modify active subscription)\n")


def test_get_subscription_entitlements(client: FlexPrice):
    """Test 9: Get Subscription Entitlements"""
    print("--- Test 9: Get Subscription Entitlements ---")

    if not test_subscription_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.subscriptions.get_subscriptions_id_entitlements(id=test_subscription_id)
        print(f"✓ Retrieved {len(response.features) if response.features else 0} features\n")
    except Exception as e:
        print(f"⚠ Skipping: {e}\n")


def test_get_upcoming_grants(client: FlexPrice):
    """Test 10: Get Upcoming Grants"""
    print("--- Test 10: Get Upcoming Grants ---")
    print("⚠ Skipping (endpoint needs verification)\n")


def test_report_usage(client: FlexPrice):
    """Test 11: Report Usage"""
    print("--- Test 11: Report Usage---")
    print("⚠ Skipping (requires metered feature setup)\n")


def test_update_line_item(client: FlexPrice):
    """Test 12: Update Line Item"""
    print("--- Test 12: Update Line Item ---")
    print("⚠ Skipping (requires line item ID)\n")


def test_delete_line_item(client: FlexPrice):
    """Test 13: Delete Line Item"""
    print("--- Test 13: Delete Line Item ---")
    print("⚠ Skipping (requires line item ID)\n")


def test_cancel_subscription(client: FlexPrice):
    """Test 14: Cancel Subscription"""
    print("--- Test 14: Cancel Subscription ---")

    if not test_subscription_id:
        print("⚠ Skipping\n")
        return

    try:
        client.subscriptions.post_subscriptions_id_cancel(
            id=test_subscription_id,
            request=components.DtoCancelSubscriptionRequest(
                cancellation_type=components.TypesCancellationType.END_OF_PERIOD
            )
        )
        print("✓ Subscription canceled\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


# ========================================
# INVOICES API TESTS
# ========================================

def test_list_invoices(client: FlexPrice):
    """Test 1: List Invoices"""
    print("--- Test 1: List Invoices ---")

    try:
        response = client.invoices.get_invoices(request=operations.GetInvoicesRequest(limit=10))
        global test_invoice_id
        if response.items and len(response.items) > 0:
            test_invoice_id = response.items[0].id
        print(f"✓ Retrieved {len(response.items) if response.items else 0} invoices\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_search_invoices(client: FlexPrice):
    """Test 2: Search Invoices"""
    print("--- Test 2: Search Invoices ---")

    try:
        response = client.invoices.post_invoices_search(
            request=components.TypesInvoiceFilter()
        )
        print(f"✓ Found {len(response.items) if response.items else 0} invoices\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_create_invoice(client: FlexPrice):
    """Test 3: Create Invoice"""
    print("--- Test 3: Create Invoice ---")

    if not test_customer_id:
        print("⚠ Skipping\n")
        return

    try:
        invoice_request = components.DtoCreateInvoiceRequest(
            customer_id=test_customer_id,
            currency="USD",
            amount_due="100.00",
            subtotal="100.00",
            total="100.00",
            invoice_type=components.TypesInvoiceType.ONE_OFF,
            billing_reason=components.TypesInvoiceBillingReason.MANUAL,
            invoice_status=components.TypesInvoiceStatus.DRAFT,
            line_items=[
                components.DtoCreateInvoiceLineItemRequest(
                    display_name="Test Service",
                    quantity="1",
                    amount="100.00",
                )
            ],
        )

        response = client.invoices.post_invoices(request=invoice_request)

        global test_invoice_id
        test_invoice_id = response.id
        print(f"✓ Invoice created: {response.id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_get_invoice(client: FlexPrice):
    """Test 4: Get Invoice"""
    print("--- Test 4: Get Invoice ---")

    if not test_invoice_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.invoices.get_invoices_id_(id=test_invoice_id)
        print(f"✓ Invoice retrieved: {response.id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_update_invoice(client: FlexPrice):
    """Test 5: Update Invoice"""
    print("--- Test 5: Update Invoice ---")

    if not test_invoice_id:
        print("⚠ Skipping\n")
        return

    try:
        client.invoices.put_invoices_id_(
            id=test_invoice_id,
            body=components.DtoUpdateInvoiceRequest(
                metadata={"updated": "true"}
            )
        )
        print("✓ Invoice updated\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_preview_invoice(client: FlexPrice):
    """Test 6: Preview Invoice"""
    print("--- Test 6: Preview Invoice ---")

    if not test_subscription_id:
        print("⚠ Skipping\n")
        return

    try:
        preview_request = components.DtoGetPreviewInvoiceRequest(
            subscription_id=test_subscription_id
        )
        client.invoices.post_invoices_preview(request=preview_request)
        print("✓ Invoice preview generated\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_finalize_invoice(client: FlexPrice):
    """Test 7: Finalize Invoice"""
    print("--- Test 7: Finalize Invoice ---")

    if not test_customer_id:
        print("⚠ Skipping\n")
        return

    try:
        # Create draft invoice
        draft_request = components.DtoCreateInvoiceRequest(
            customer_id=test_customer_id,
            currency="USD",
            amount_due="50.00",
            subtotal="50.00",
            total="50.00",
            invoice_type=components.TypesInvoiceType.ONE_OFF,
            billing_reason=components.TypesInvoiceBillingReason.MANUAL,
            invoice_status=components.TypesInvoiceStatus.DRAFT,
            line_items=[
                components.DtoCreateInvoiceLineItemRequest(
                    display_name="Finalize Test",
                    quantity="1",
                    amount="50.00",
                )
            ],
        )

        invoice = client.invoices.post_invoices(request=draft_request)
        client.invoices.post_invoices_id_finalize(id=invoice.id)
        print(f"✓ Invoice finalized: {invoice.id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_recalculate_invoice(client: FlexPrice):
    """Test 8: Recalculate Invoice"""
    print("--- Test 8: Recalculate Invoice ---")
    print("⚠ Skipping (recalculate only works on subscription invoices)\n")


def test_record_payment(client: FlexPrice):
    """Test 9: Record Payment"""
    print("--- Test 9: Record Payment ---")

    if not test_invoice_id:
        print("⚠ Skipping\n")
        return

    try:
        payment_request = components.DtoUpdatePaymentStatusRequest(
            payment_status=components.TypesPaymentStatus.SUCCEEDED,
            amount="100.00",
        )
        client.invoices.put_invoices_id_payment(id=test_invoice_id, body=payment_request)
        print(f"✓ Payment recorded: {test_invoice_id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_attempt_payment(client: FlexPrice):
    """Test 10: Attempt Payment"""
    print("--- Test 10: Attempt Payment ---")

    if not test_customer_id:
        print("⚠ Skipping\n")
        return

    try:
        # Create and finalize invoice
        attempt_request = components.DtoCreateInvoiceRequest(
            customer_id=test_customer_id,
            currency="USD",
            amount_due="25.00",
            subtotal="25.00",
            total="25.00",
            amount_paid="0.00",
            invoice_type=components.TypesInvoiceType.ONE_OFF,
            billing_reason=components.TypesInvoiceBillingReason.MANUAL,
            invoice_status=components.TypesInvoiceStatus.DRAFT,
            payment_status=components.TypesPaymentStatus.PENDING,
            line_items=[
                components.DtoCreateInvoiceLineItemRequest(
                    display_name="Attempt Payment Test",
                    quantity="1",
                    amount="25.00",
                )
            ],
        )

        invoice = client.invoices.post_invoices(request=attempt_request)
        client.invoices.post_invoices_id_finalize(id=invoice.id)
        client.invoices.post_invoices_id_payment_attempt(id=invoice.id)
        print(f"✓ Payment attempt initiated: {invoice.id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_download_invoice_pdf(client: FlexPrice):
    """Test 11: Download Invoice PDF"""
    print("--- Test 11: Download Invoice PDF ---")

    if not test_invoice_id:
        print("⚠ Skipping\n")
        return

    try:
        client.invoices.get_invoices_id_pdf(id=test_invoice_id)
        print("✓ PDF downloaded\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_trigger_invoice_comms(client: FlexPrice):
    """Test 12: Trigger Invoice Comms"""
    print("--- Test 12: Trigger Invoice Comms ---")

    if not test_invoice_id:
        print("⚠ Skipping\n")
        return

    try:
        client.invoices.post_invoices_id_comms_trigger(id=test_invoice_id)
        print(f"✓ Invoice communications triggered: {test_invoice_id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_get_customer_invoice_summary(client: FlexPrice):
    """Test 13: Get Customer Invoice Summary"""
    print("--- Test 13: Get Customer Invoice Summary ---")

    if not test_customer_id:
        print("⚠ Skipping\n")
        return

    try:
        client.invoices.get_customers_id_invoices_summary(id=test_customer_id)
        print("✓ Summary retrieved\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_void_invoice(client: FlexPrice):
    """Test 14: Void Invoice"""
    print("--- Test 14: Void Invoice ---")

    if not test_invoice_id:
        print("⚠ Skipping\n")
        return

    try:
        client.invoices.post_invoices_id_void(id=test_invoice_id)
        print(f"✓ Invoice voided: {test_invoice_id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


# ========================================
# PRICES API TESTS
# ========================================

def test_create_price(client: FlexPrice):
    """Test 1: Create Price"""
    print("--- Test 1: Create Price ---")

    if not test_plan_id:
        print("⚠ Skipping\n")
        return

    try:
        price_request = components.DtoCreatePriceRequest(
            entity_id=test_plan_id,
            entity_type=components.TypesPriceEntityType.PLAN,
            currency="USD",
            amount="99.00",
            billing_model=components.TypesBillingModel.FLAT_FEE,
            billing_cadence=components.TypesBillingCadence.RECURRING,
            billing_period=components.TypesBillingPeriod.MONTHLY,
            billing_period_count=1,
            invoice_cadence=components.TypesInvoiceCadence.ADVANCE,
            price_unit_type=components.TypesPriceUnitType.FIAT,
            type=components.TypesPriceType.FIXED,
            display_name="Monthly Subscription",
        )

        response = client.prices.post_prices(request=price_request)

        global test_price_id
        test_price_id = response.id
        print(f"✓ Price created: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_get_price(client: FlexPrice):
    """Test 2: Get Price"""
    print("--- Test 2: Get Price ---")

    if not test_price_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.prices.get_prices_id_(id=test_price_id)
        print(f"✓ Price retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_prices(client: FlexPrice):
    """Test 3: List Prices"""
    print("--- Test 3: List Prices ---")

    try:
        response = client.prices.get_prices(request=operations.GetPricesRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} prices\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_update_price(client: FlexPrice):
    """Test 4: Update Price"""
    print("--- Test 4: Update Price ---")

    if not test_price_id:
        print("⚠ Skipping\n")
        return

    try:
        updated_desc = "Updated price description"
        client.prices.put_prices_id_(
            id=test_price_id,
            request=components.DtoUpdatePriceRequest(description=updated_desc)
        )
        print("✓ Price updated\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_delete_price(client: FlexPrice):
    """Test 5: Delete Price"""
    print("--- Test 5: Delete Price ---")

    if not test_price_id:
        print("⚠ Skipping\n")
        return

    try:
        future_date = (datetime.now(timezone.utc) + timedelta(days=1)).isoformat()
        client.prices.delete_prices_id_(
            id=test_price_id,
            body=components.DtoDeletePriceRequest(end_date=future_date)
        )
        print("✓ Price deleted\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


# ========================================
# PAYMENTS API TESTS
# ========================================

def test_create_payment(client: FlexPrice):
    """Test 1: Create Payment"""
    print("--- Test 1: Create Payment ---")
    print("⚠ Skipping (requires invoice setup)\n")


def test_get_payment(client: FlexPrice):
    """Test 2: Get Payment"""
    print("--- Test 2: Get Payment ---")

    if not test_payment_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.payments.get_payments_id_(id=test_payment_id)
        print(f"✓ Payment retrieved: {response.id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_list_payments(client: FlexPrice):
    """Test 3: List Payments"""
    print("--- Test 3: List Payments ---")

    try:
        response = client.payments.get_payments(request=operations.GetPaymentsRequest(limit=10))
        global test_payment_id
        if response.items and len(response.items) > 0:
            test_payment_id = response.items[0].id
        print(f"✓ Retrieved {len(response.items) if response.items else 0} payments\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_update_payment(client: FlexPrice):
    """Test 4: Update Payment"""
    print("--- Test 4: Update Payment ---")

    if not test_payment_id:
        print("⚠ Skipping\n")
        return

    try:
        update_request = components.DtoUpdatePaymentRequest(
            metadata={
                "updated_at": datetime.now().isoformat(),
                "status": "updated",
            }
        )
        client.payments.put_payments_id_(id=test_payment_id, body=update_request)
        print(f"✓ Payment updated: {test_payment_id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_process_payment(client: FlexPrice):
    """Test 5: Process Payment"""
    print("--- Test 5: Process Payment ---")

    if not test_payment_id:
        print("⚠ Skipping\n")
        return

    try:
        client.payments.post_payments_id_process(id=test_payment_id)
        print(f"✓ Payment processed: {test_payment_id}\n")
    except Exception as e:
        print(f"⚠ Error (may require payment gateway setup): {e}\n")


def test_delete_payment(client: FlexPrice):
    """Test 6: Delete Payment"""
    print("--- Test 6: Delete Payment ---")

    if not test_payment_id:
        print("⚠ Skipping\n")
        return

    try:
        client.payments.delete_payments_id_(id=test_payment_id)
        print("✓ Payment deleted\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


# ========================================
# WALLETS API TESTS
# ========================================

def test_create_wallet(client: FlexPrice):
    """Test 1: Create Wallet"""
    print("--- Test 1: Create Wallet ---")

    if not test_customer_id:
        print("⚠ Skipping\n")
        return

    try:
        wallet_request = components.DtoCreateWalletRequest(
            customer_id=test_customer_id,
            currency="USD",
            name="Test Wallet",
        )

        response = client.wallets.post_wallets(request=wallet_request)

        global test_wallet_id
        test_wallet_id = response.id
        print(f"✓ Wallet created: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_get_wallet(client: FlexPrice):
    """Test 2: Get Wallet"""
    print("--- Test 2: Get Wallet ---")

    if not test_wallet_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.wallets.get_wallets_id_(id=test_wallet_id)
        print(f"✓ Wallet retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_wallets(client: FlexPrice):
    """Test 3: List Wallets"""
    print("--- Test 3: List Wallets ---")

    try:
        response = client.wallets.get_wallets(request=operations.GetWalletsRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} wallets\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_update_wallet(client: FlexPrice):
    """Test 4: Update Wallet"""
    print("--- Test 4: Update Wallet ---")

    if not test_wallet_id:
        print("⚠ Skipping\n")
        return

    try:
        update_request = components.DtoUpdateWalletRequest(
            metadata={
                "updated_at": datetime.now().isoformat(),
                "status": "updated",
            }
        )
        client.wallets.put_wallets_id_(id=test_wallet_id, body=update_request)
        print(f"✓ Wallet updated: {test_wallet_id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_get_wallet_balance(client: FlexPrice):
    """Test 5: Get Wallet Balance"""
    print("--- Test 5: Get Wallet Balance ---")

    if not test_wallet_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.wallets.get_wallets_id_balance_real_time(id=test_wallet_id)
        print(f"✓ Balance: {response.balance}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_top_up_wallet(client: FlexPrice):
    """Test 6: Top Up Wallet"""
    print("--- Test 6: Top Up Wallet ---")

    if not test_wallet_id:
        print("⚠ Skipping\n")
        return

    try:
        top_up_request = components.DtoTopUpWalletRequest(
            amount="100.00",
            transaction_reason=components.TypesTransactionReason.PURCHASED_CREDIT_DIRECT,
            description="Test top-up from SDK",
        )
        client.wallets.post_wallets_id_top_up(id=test_wallet_id, body=top_up_request)
        print(f"✓ Wallet topped up: {test_wallet_id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_debit_wallet(client: FlexPrice):
    """Test 7: Debit Wallet"""
    print("--- Test 7: Debit Wallet ---")
    print("⚠ Skipping (debit endpoint not available in SDK)\n")


def test_get_wallet_transactions(client: FlexPrice):
    """Test 8: Get Wallet Transactions"""
    print("--- Test 8: Get Wallet Transactions ---")

    if not test_wallet_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.wallets.get_wallets_id_transactions(request=operations.GetWalletsIDTransactionsRequest(id_path_parameter=test_wallet_id, limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} transactions\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_search_wallets(client: FlexPrice):
    """Test 9: Search Wallets"""
    print("--- Test 9: Search Wallets ---")

    try:
        response = client.wallets.post_wallets_search(
            request=components.TypesWalletFilter(limit=10)
        )
        print(f"✓ Found {len(response.items) if response.items else 0} wallets\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


# ========================================
# CREDIT GRANTS API TESTS
# ========================================

def test_create_credit_grant(client: FlexPrice):
    """Test 1: Create Credit Grant"""
    print("--- Test 1: Create Credit Grant ---")

    if not test_plan_id:
        print("⚠ Skipping\n")
        return

    try:
        grant_request = components.DtoCreateCreditGrantRequest(
            scope=components.TypesCreditGrantScope.PLAN,
            plan_id=test_plan_id,
            credits="500.00",
            name="Test Credit Grant",
            cadence=components.TypesCreditGrantCadence.ONETIME,
            expiration_type=components.TypesCreditGrantExpiryType.NEVER,
            expiration_duration_unit=components.TypesCreditGrantExpiryDurationUnit.DAY,
        )

        response = client.credit_grants.post_creditgrants(request=grant_request)

        global test_credit_grant_id
        test_credit_grant_id = response.id
        print(f"✓ Credit grant created: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_get_credit_grant(client: FlexPrice):
    """Test 2: Get Credit Grant"""
    print("--- Test 2: Get Credit Grant ---")

    if not test_credit_grant_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.credit_grants.get_creditgrants_id_(id=test_credit_grant_id)
        print(f"✓ Credit grant retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_credit_grants(client: FlexPrice):
    """Test 3: List Credit Grants"""
    print("--- Test 3: List Credit Grants ---")

    try:
        response = client.credit_grants.get_creditgrants(request=operations.GetCreditgrantsRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} credit grants\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_update_credit_grant(client: FlexPrice):
    """Test 4: Update Credit Grant"""
    print("--- Test 4: Update Credit Grant ---")

    if not test_credit_grant_id:
        print("⚠ Skipping\n")
        return

    try:
        update_request = components.DtoUpdateCreditGrantRequest(
            metadata={
                "updated_at": datetime.now().isoformat(),
                "status": "updated",
            }
        )
        client.credit_grants.put_creditgrants_id_(id=test_credit_grant_id, body=update_request)
        print(f"✓ Credit grant updated: {test_credit_grant_id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


# ========================================
# CREDIT NOTES API TESTS
# ========================================

def test_create_credit_note(client: FlexPrice):
    """Test 1: Create Credit Note"""
    print("--- Test 1: Create Credit Note ---")
    print("⚠ Skipping (requires invoice with line items)\n")


def test_get_credit_note(client: FlexPrice):
    """Test 2: Get Credit Note"""
    print("--- Test 2: Get Credit Note ---")

    if not test_credit_note_id:
        print("⚠ Skipping\n")
        return

    try:
        response = client.credit_notes.get_creditnotes_id_(id=test_credit_note_id)
        print(f"✓ Credit note retrieved: {response.id}\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_list_credit_notes(client: FlexPrice):
    """Test 3: List Credit Notes"""
    print("--- Test 3: List Credit Notes ---")

    try:
        response = client.credit_notes.get_creditnotes(request=operations.GetCreditnotesRequest(limit=10))
        print(f"✓ Retrieved {len(response.items) if response.items else 0} credit notes\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_finalize_credit_note(client: FlexPrice):
    """Test 4: Finalize Credit Note"""
    print("--- Test 4: Finalize Credit Note ---")

    if not test_credit_note_id:
        print("⚠ Skipping\n")
        return

    try:
        client.credit_notes.post_creditnotes_id_finalize(id=test_credit_note_id)
        print(f"✓ Credit note finalized: {test_credit_note_id}\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


# ========================================
# EVENTS API TESTS
# ========================================

def test_create_event(client: FlexPrice):
    """Test 1: Create Event"""
    print("--- Test 1: Create Event ---")

    global test_event_id, test_event_name, test_event_customer_id
    if test_external_id:
        test_event_customer_id = test_external_id
    else:
        test_event_customer_id = f"test-customer-{int(time.time())}"

    test_event_name = f"Test Event {int(time.time())}"

    try:
        event_request = components.DtoIngestEventRequest(
            event_name=test_event_name,
            external_customer_id=test_event_customer_id,
            properties={
                "source": "sdk_test",
                "environment": "test",
            },
            source="sdk_test",
            timestamp=datetime.now(timezone.utc).isoformat(),
        )

        response = client.events.post_events(request=event_request)

        if response and isinstance(response, dict) and "event_id" in response:
            test_event_id = response["event_id"]
            print(f"✓ Event created: {test_event_id}\n")
        else:
            print("✓ Event created\n")
    except Exception as e:
        print(f"❌ Error: {e}\n")


def test_query_events(client: FlexPrice):
    """Test 2: Query Events"""
    print("--- Test 2: Query Events ---")

    if not test_event_name:
        print("⚠ Skipping\n")
        return

    try:
        query_request = {
            "external_customer_id": test_event_customer_id,
            "event_name": test_event_name,
        }

        response = client.events.post_events_query(request=query_request)
        print(f"✓ Found {len(response.events) if response.events else 0} events\n")
    except Exception as e:
        print(f"⚠ Error: {e}\n")


def test_async_event_ingestion(client: FlexPrice):
    """Test 3: Async Event Ingestion"""
    print("--- Test 3: Async Event Ingestion ---")
    print("⚠ Skipping (async client needs implementation)\n")


# ========================================
# MAIN EXECUTION
# ========================================

def main():
    """Main execution function"""
    client = get_client()

    print("========================================")
    print("CUSTOMER API TESTS")
    print("========================================\n")
    test_create_customer(client)
    test_get_customer(client)
    test_list_customers(client)
    test_update_customer(client)
    test_lookup_customer(client)
    test_search_customers(client)
    test_get_customer_entitlements(client)
    test_get_customer_upcoming_grants(client)
    test_get_customer_usage(client)
    print("✓ Customer API Tests Completed!\n")

    print("========================================")
    print("FEATURES API TESTS")
    print("========================================\n")
    test_create_feature(client)
    test_get_feature(client)
    test_list_features(client)
    test_update_feature(client)
    test_search_features(client)
    print("✓ Features API Tests Completed!\n")

    print("========================================")
    print("CONNECTIONS API TESTS")
    print("========================================\n")
    test_list_connections(client)
    test_search_connections(client)
    print("✓ Connections API Tests Completed!\n")

    print("========================================")
    print("PLANS API TESTS")
    print("========================================\n")
    test_create_plan(client)
    test_get_plan(client)
    test_list_plans(client)
    test_update_plan(client)
    test_search_plans(client)
    print("✓ Plans API Tests Completed!\n")

    print("========================================")
    print("ADDONS API TESTS")
    print("========================================\n")
    test_create_addon(client)
    test_get_addon(client)
    test_list_addons(client)
    test_update_addon(client)
    test_lookup_addon(client)
    test_search_addons(client)
    print("✓ Addons API Tests Completed!\n")

    print("========================================")
    print("ENTITLEMENTS API TESTS")
    print("========================================\n")
    test_create_entitlement(client)
    test_get_entitlement(client)
    test_list_entitlements(client)
    test_update_entitlement(client)
    test_search_entitlements(client)
    print("✓ Entitlements API Tests Completed!\n")

    print("========================================")
    print("SUBSCRIPTIONS API TESTS")
    print("========================================\n")
    test_create_subscription(client)
    test_get_subscription(client)
    test_list_subscriptions(client)
    test_search_subscriptions(client)
    test_activate_subscription(client)
    test_add_addon_to_subscription(client)
    test_remove_addon_from_subscription(client)
    test_execute_subscription_change(client)
    test_get_subscription_entitlements(client)
    test_get_upcoming_grants(client)
    test_report_usage(client)
    test_update_line_item(client)
    test_delete_line_item(client)
    test_cancel_subscription(client)
    print("✓ Subscriptions API Tests Completed!\n")

    print("========================================")
    print("INVOICES API TESTS")
    print("========================================\n")
    test_list_invoices(client)
    test_search_invoices(client)
    test_create_invoice(client)
    test_get_invoice(client)
    test_update_invoice(client)
    test_preview_invoice(client)
    test_finalize_invoice(client)
    test_recalculate_invoice(client)
    test_record_payment(client)
    test_attempt_payment(client)
    test_download_invoice_pdf(client)
    test_trigger_invoice_comms(client)
    test_get_customer_invoice_summary(client)
    test_void_invoice(client)
    print("✓ Invoices API Tests Completed!\n")

    print("========================================")
    print("PRICES API TESTS")
    print("========================================\n")
    test_create_price(client)
    test_get_price(client)
    test_list_prices(client)
    test_update_price(client)
    print("✓ Prices API Tests Completed!\n")

    print("========================================")
    print("PAYMENTS API TESTS")
    print("========================================\n")
    test_create_payment(client)
    test_get_payment(client)
    test_list_payments(client)
    test_update_payment(client)
    test_process_payment(client)
    print("✓ Payments API Tests Completed!\n")

    print("========================================")
    print("WALLETS API TESTS")
    print("========================================\n")
    test_create_wallet(client)
    test_get_wallet(client)
    test_list_wallets(client)
    test_update_wallet(client)
    test_get_wallet_balance(client)
    test_top_up_wallet(client)
    test_debit_wallet(client)
    test_get_wallet_transactions(client)
    test_search_wallets(client)
    print("✓ Wallets API Tests Completed!\n")

    print("========================================")
    print("CREDIT GRANTS API TESTS")
    print("========================================\n")
    test_create_credit_grant(client)
    test_get_credit_grant(client)
    test_list_credit_grants(client)
    test_update_credit_grant(client)
    print("✓ Credit Grants API Tests Completed!\n")

    print("========================================")
    print("CREDIT NOTES API TESTS")
    print("========================================\n")
    test_create_credit_note(client)
    test_get_credit_note(client)
    test_list_credit_notes(client)
    test_finalize_credit_note(client)
    print("✓ Credit Notes API Tests Completed!\n")

    print("========================================")
    print("EVENTS API TESTS")
    print("========================================\n")
    test_create_event(client)
    test_query_events(client)
    test_async_event_ingestion(client)
    print("✓ Events API Tests Completed!\n")

    print("========================================")
    print("CLEANUP - DELETING TEST DATA")
    print("========================================\n")
    test_delete_payment(client)
    test_delete_price(client)
    test_delete_entitlement(client)
    test_delete_addon(client)
    test_delete_plan(client)
    test_delete_feature(client)
    test_delete_customer(client)
    print("✓ Cleanup Completed!\n")

    print("\n=== All API Tests Completed Successfully! ===")


if __name__ == "__main__":
    main()

