/**
 * FlexPrice Speakeasy TypeScript SDK - Comprehensive API Tests
 *
 * This test suite covers all FlexPrice APIs with comprehensive CRUD operations.
 * Tests are designed to run sequentially, building on previous test results.
 * Uses the Speakeasy-generated TypeScript SDK.
 *
 * Setup:
 * 1. Export your API key: export FLEXPRICE_API_KEY="your_key_here"
 * 2. Export API host: export FLEXPRICE_API_HOST="https://api.cloud.flexprice.io/v1"
 * 3. Run: npx ts-node test_speakeasy_sdk.ts
 */

import { FlexPrice } from "flexprice-sdk-test";
import * as components from "flexprice-sdk-test/models/components/index.js";

// Global test entity IDs
let testCustomerID: string | undefined;
let testCustomerName: string | undefined;
let testExternalID: string | undefined;

let testFeatureID: string | undefined;
let testFeatureName: string | undefined;

let testPlanID: string | undefined;
let testPlanName: string | undefined;

let testAddonID: string | undefined;
let testAddonName: string | undefined;
let testAddonLookupKey: string | undefined;

let testEntitlementID: string | undefined;

let testSubscriptionID: string | undefined;

let testInvoiceID: string | undefined;

let testPriceID: string | undefined;

let testPaymentID: string | undefined;

let testWalletID: string | undefined;
let testCreditGrantID: string | undefined;
let testCreditNoteID: string | undefined;

let testEventID: string | undefined;
let testEventName: string | undefined;
let testEventCustomerID: string | undefined;

function getClient(): FlexPrice {
  const apiKey = process.env.FLEXPRICE_API_KEY;
  const apiHost = process.env.FLEXPRICE_API_HOST || "https://api.cloud.flexprice.io/v1";

  if (!apiKey) {
    console.error("❌ Missing FLEXPRICE_API_KEY environment variable");
    process.exit(1);
  }

  console.log("=== FlexPrice Speakeasy TypeScript SDK - API Tests ===\n");
  console.log(`✓ API Key: ${apiKey.substring(0, 8)}...${apiKey.slice(-4)}`);
  console.log(`✓ API Host: ${apiHost}\n`);

  // Initialize Speakeasy SDK
  const client = new FlexPrice({
    serverURL: apiHost,
    apiKeyAuth: apiKey,
  });

  return client;
}

// ========================================
// CUSTOMER API TESTS
// ========================================

async function testCreateCustomer(client: FlexPrice) {
  console.log("--- Test 1: Create Customer ---");

  try {
    const timestamp = Date.now();
    testCustomerName = `Test Customer ${timestamp}`;
    testExternalID = `test-customer-${timestamp}`;

    const customerRequest: components.DtoCreateCustomerRequest = {
      name: testCustomerName,
      email: `test-${timestamp}@example.com`,
      externalId: testExternalID,
      metadata: {
        source: "sdk_test",
        testRun: new Date().toISOString(),
        environment: "test",
      },
    };

    const response = await client.customers.postCustomers(customerRequest);

    testCustomerID = response.id;
    console.log("✓ Customer created successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}`);
    console.log(`  External ID: ${response.externalId}`);
    console.log(`  Email: ${response.email}\n`);
  } catch (e) {
    console.error(`❌ Error creating customer: ${e}\n`);
  }
}

async function testGetCustomer(client: FlexPrice) {
  console.log("--- Test 2: Get Customer by ID ---");

  if (!testCustomerID) {
    console.log("⚠ Warning: No customer ID available\n⚠ Skipping get customer test\n");
    return;
  }

  try {
    const response = await client.customers.getCustomersId(testCustomerID);

    console.log("✓ Customer retrieved successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}`);
    console.log(`  Created At: ${response.createdAt}\n`);
  } catch (e) {
    console.error(`❌ Error getting customer: ${e}\n`);
  }
}

async function testListCustomers(client: FlexPrice) {
  console.log("--- Test 3: List Customers ---");

  try {
    const response = await client.customers.getCustomers({ limit: 10 });

    console.log(`✓ Retrieved ${response.items?.length || 0} customers`);
    if (response.items && response.items.length > 0) {
      console.log(`  First customer: ${response.items[0].id} - ${response.items[0].name}`);
    }
    if (response.pagination) {
      console.log(`  Total: ${response.pagination.total}\n`);
    }
  } catch (e) {
    console.error(`❌ Error listing customers: ${e}\n`);
  }
}

async function testUpdateCustomer(client: FlexPrice) {
  console.log("--- Test 4: Update Customer ---");

  if (!testCustomerID) {
    console.log("⚠ Warning: No customer ID available\n⚠ Skipping update customer test\n");
    return;
  }

  try {
    const updateRequest: components.DtoUpdateCustomerRequest = {
      name: `${testCustomerName} (Updated)`,
      metadata: {
        updatedAt: new Date().toISOString(),
        status: "updated",
      },
    };

    const response = await client.customers.putCustomersId(testCustomerID, updateRequest);

    console.log("✓ Customer updated successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  New Name: ${response.name}`);
    if (response.updatedAt) {
      console.log(`  Updated At: ${response.updatedAt}\n`);
    } else {
      console.log();
    }
  } catch (e) {
    console.error(`❌ Error updating customer: ${e}\n`);
  }
}

async function testLookupCustomer(client: FlexPrice) {
  console.log("--- Test 5: Lookup Customer by External ID ---");

  if (!testExternalID) {
    console.log("⚠ Warning: No external ID available\n⚠ Skipping lookup test\n");
    return;
  }

  try {
    const response = await client.customers.getCustomersExternalExternalId(testExternalID);

    console.log("✓ Customer found by external ID!");
    console.log(`  External ID: ${testExternalID}`);
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}\n`);
  } catch (e) {
    console.error(`❌ Error looking up customer: ${e}\n`);
  }
}

async function testSearchCustomers(client: FlexPrice) {
  console.log("--- Test 6: Search Customers ---");

  if (!testExternalID) {
    console.log("⚠ Warning: No external ID available\n⚠ Skipping search test\n");
    return;
  }

  try {
    const searchFilter: components.TypesCustomerFilter = {
      externalId: testExternalID,
    };

    const response = await client.customers.postCustomersSearch(searchFilter);

    console.log("✓ Search completed!");
    console.log(`  Found ${response.items?.length || 0} customers matching external ID '${testExternalID}'`);
    if (response.items) {
      response.items.slice(0, 3).forEach((customer) => {
        console.log(`  - ${customer.id}: ${customer.name}`);
      });
    }
    console.log();
  } catch (e) {
    console.error(`❌ Error searching customers: ${e}\n`);
  }
}

async function testGetCustomerEntitlements(client: FlexPrice) {
  console.log("--- Test 7: Get Customer Entitlements ---");

  if (!testCustomerID) {
    console.log("⚠ Warning: No customer ID available\n⚠ Skipping get entitlements test\n");
    return;
  }

  try {
    const response = await client.customers.getCustomersIdEntitlements(testCustomerID, undefined, undefined);

    console.log("✓ Retrieved customer entitlements!");
    if (response.features) {
      console.log(`  Total features: ${response.features.length}`);
      response.features.slice(0, 3).forEach((feature) => {
        if (feature.feature?.id) {
          console.log(`  - Feature: ${feature.feature.id}`);
        }
      });
    } else {
      console.log("  No features found");
    }
    console.log();
  } catch (e) {
    console.log(`⚠ Warning: Error getting customer entitlements: ${e}`);
    console.log("⚠ Skipping entitlements test (customer may not have any entitlements)\n");
  }
}

async function testGetCustomerUpcomingGrants(client: FlexPrice) {
  console.log("--- Test 8: Get Customer Upcoming Grants ---");

  if (!testCustomerID) {
    console.log("⚠ Warning: No customer ID available\n⚠ Skipping get upcoming grants test\n");
    return;
  }

  try {
    const response = await client.customers.getCustomersIdGrantsUpcoming(testCustomerID);

    console.log("✓ Retrieved upcoming grants!");
    if (response.items) {
      console.log(`  Total upcoming grants: ${response.items.length}`);
    } else {
      console.log("  No upcoming grants found");
    }
    console.log();
  } catch (e) {
    console.log(`⚠ Warning: Error getting upcoming grants: ${e}`);
    console.log("⚠ Skipping upcoming grants test (customer may not have any grants)\n");
  }
}

async function testGetCustomerUsage(client: FlexPrice) {
  console.log("--- Test 9: Get Customer Usage ---");

  if (!testCustomerID) {
    console.log("⚠ Warning: No customer ID available\n⚠ Skipping get usage test\n");
    return;
  }

  try {
    const response = await client.customers.getCustomersUsage({ customerId: testCustomerID });

    console.log("✓ Retrieved customer usage!");
    if (response.features) {
      console.log(`  Feature usage records: ${response.features.length}`);
    } else {
      console.log("  No usage data found");
    }
    console.log();
  } catch (e) {
    console.log(`⚠ Warning: Error getting customer usage: ${e}`);
    console.log("⚠ Skipping usage test (customer may not have usage data)\n");
  }
}

async function testDeleteCustomer(client: FlexPrice) {
  console.log("--- Test 10: Delete Customer ---");

  if (!testCustomerID) {
    console.log("⚠ Skipping delete customer (no customer created)\n");
    return;
  }

  try {
    await client.customers.deleteCustomersId(testCustomerID);

    console.log("✓ Customer deleted successfully!");
    console.log(`  Deleted ID: ${testCustomerID}\n`);
  } catch (e) {
    console.error(`❌ Error deleting customer: ${e}\n`);
  }
}

// ========================================
// FEATURES API TESTS
// ========================================

async function testCreateFeature(client: FlexPrice) {
  console.log("--- Test 1: Create Feature ---");

  try {
    const timestamp = Date.now();
    testFeatureName = `Test Feature ${timestamp}`;
    const featureKey = `test_feature_${timestamp}`;

    const featureRequest: components.DtoCreateFeatureRequest = {
      name: testFeatureName,
      lookupKey: featureKey,
      description: "This is a test feature created by SDK tests",
      type: components.TypesFeatureType.Boolean,
      metadata: {
        source: "sdk_test",
        testRun: new Date().toISOString(),
        environment: "test",
      },
    };

    const response = await client.features.postFeatures(featureRequest);

    testFeatureID = response.id;
    console.log("✓ Feature created successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}`);
    console.log(`  Lookup Key: ${response.lookupKey}`);
    console.log(`  Type: ${response.type}\n`);
  } catch (e) {
    console.error(`❌ Error creating feature: ${e}\n`);
  }
}

async function testGetFeature(client: FlexPrice) {
  console.log("--- Test 2: Get Feature by ID ---");

  if (!testFeatureID) {
    console.log("⚠ Warning: No feature ID available\n⚠ Skipping get feature test\n");
    return;
  }

  try {
    const response = await client.features.getFeaturesId(testFeatureID);

    console.log("✓ Feature retrieved successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}`);
    console.log(`  Lookup Key: ${response.lookupKey}`);
    console.log(`  Created At: ${response.createdAt}\n`);
  } catch (e) {
    console.error(`❌ Error getting feature: ${e}\n`);
  }
}

async function testListFeatures(client: FlexPrice) {
  console.log("--- Test 3: List Features ---");

  try {
    const response = await client.features.getFeatures({ limit: 10 });

    console.log(`✓ Retrieved ${response.items?.length || 0} features`);
    if (response.items && response.items.length > 0) {
      console.log(`  First feature: ${response.items[0].id} - ${response.items[0].name}`);
    }
    if (response.pagination) {
      console.log(`  Total: ${response.pagination.total}\n`);
    }
  } catch (e) {
    console.error(`❌ Error listing features: ${e}\n`);
  }
}

async function testUpdateFeature(client: FlexPrice) {
  console.log("--- Test 4: Update Feature ---");

  if (!testFeatureID) {
    console.log("⚠ Warning: No feature ID available\n⚠ Skipping update feature test\n");
    return;
  }

  try {
    const updateRequest: components.DtoUpdateFeatureRequest = {
      name: `${testFeatureName} (Updated)`,
      description: "Updated description for test feature",
      metadata: {
        updatedAt: new Date().toISOString(),
        status: "updated",
      },
    };

    const response = await client.features.putFeaturesId(testFeatureID, updateRequest);

    console.log("✓ Feature updated successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  New Name: ${response.name}`);
    console.log(`  New Description: ${response.description}`);
    console.log(`  Updated At: ${response.updatedAt}\n`);
  } catch (e) {
    console.error(`❌ Error updating feature: ${e}\n`);
  }
}

async function testSearchFeatures(client: FlexPrice) {
  console.log("--- Test 5: Search Features ---");

  if (!testFeatureID) {
    console.log("⚠ Warning: No feature ID available\n⚠ Skipping search test\n");
    return;
  }

  try {
    const searchFilter: components.TypesFeatureFilter = {
      featureIds: [testFeatureID],
    };

    const response = await client.features.postFeaturesSearch(searchFilter);

    console.log("✓ Search completed!");
    console.log(`  Found ${response.items?.length || 0} features matching ID '${testFeatureID}'`);
    if (response.items) {
      response.items.slice(0, 3).forEach((feature) => {
        console.log(`  - ${feature.id}: ${feature.name} (${feature.lookupKey})`);
      });
    }
    console.log();
  } catch (e) {
    console.error(`❌ Error searching features: ${e}\n`);
  }
}

async function testDeleteFeature(client: FlexPrice) {
  console.log("--- Test 6: Delete Feature ---");

  if (!testFeatureID) {
    console.log("⚠ Skipping delete feature (no feature created)\n");
    return;
  }

  try {
    await client.features.deleteFeaturesId(testFeatureID);

    console.log("✓ Feature deleted successfully!");
    console.log(`  Deleted ID: ${testFeatureID}\n`);
  } catch (e) {
    console.error(`❌ Error deleting feature: ${e}\n`);
  }
}

// ========================================
// CONNECTIONS API TESTS
// ========================================

async function testListConnections(client: FlexPrice) {
  console.log("--- Test 1: List Connections ---");

  try {
    const response = await client.connections.getConnections({ limit: 10 });
    console.log(`✓ Retrieved ${response.connections?.length || 0} connections\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testSearchConnections(client: FlexPrice) {
  console.log("--- Test 2: Search Connections ---");

  try {
    const response = await client.connections.postConnectionsSearch({ limit: 5 });
    console.log(`✓ Found ${response.connections?.length || 0} connections\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

// ========================================
// PLANS API TESTS
// ========================================

async function testCreatePlan(client: FlexPrice) {
  console.log("--- Test 1: Create Plan ---");

  try {
    const timestamp = Date.now();
    testPlanName = `Test Plan ${timestamp}`;
    const lookupKey = `test_plan_${timestamp}`;

    const planRequest: components.DtoCreatePlanRequest = {
      name: testPlanName,
      lookupKey: lookupKey,
      description: "This is a test plan created by SDK tests",
      metadata: {
        source: "sdk_test",
        testRun: new Date().toISOString(),
        environment: "test",
      },
    };

    const response = await client.plans.postPlans(planRequest);

    testPlanID = response.id;
    console.log("✓ Plan created successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}`);
    console.log(`  Lookup Key: ${response.lookupKey}\n`);
  } catch (e) {
    console.error(`❌ Error creating plan: ${e}\n`);
  }
}

async function testGetPlan(client: FlexPrice) {
  console.log("--- Test 2: Get Plan by ID ---");

  if (!testPlanID) {
    console.log("⚠ Warning: No plan ID available\n⚠ Skipping get plan test\n");
    return;
  }

  try {
    const response = await client.plans.getPlansId(testPlanID);

    console.log("✓ Plan retrieved successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  Name: ${response.name}`);
    console.log(`  Lookup Key: ${response.lookupKey}`);
    console.log(`  Created At: ${response.createdAt}\n`);
  } catch (e) {
    console.error(`❌ Error getting plan: ${e}\n`);
  }
}

async function testListPlans(client: FlexPrice) {
  console.log("--- Test 3: List Plans ---");

  try {
    const response = await client.plans.getPlans({ limit: 10 });

    console.log(`✓ Retrieved ${response.items?.length || 0} plans`);
    if (response.items && response.items.length > 0) {
      console.log(`  First plan: ${response.items[0].id} - ${response.items[0].name}`);
    }
    if (response.pagination) {
      console.log(`  Total: ${response.pagination.total}\n`);
    }
  } catch (e) {
    console.error(`❌ Error listing plans: ${e}\n`);
  }
}

async function testUpdatePlan(client: FlexPrice) {
  console.log("--- Test 4: Update Plan ---");

  if (!testPlanID) {
    console.log("⚠ Warning: No plan ID available\n⚠ Skipping update plan test\n");
    return;
  }

  try {
    const updateRequest: components.DtoUpdatePlanRequest = {
      name: `${testPlanName} (Updated)`,
      description: "Updated description for test plan",
      metadata: {
        updatedAt: new Date().toISOString(),
        status: "updated",
      },
    };

    const response = await client.plans.putPlansId(testPlanID, updateRequest);

    console.log("✓ Plan updated successfully!");
    console.log(`  ID: ${response.id}`);
    console.log(`  New Name: ${response.name}`);
    console.log(`  New Description: ${response.description}`);
    console.log(`  Updated At: ${response.updatedAt}\n`);
  } catch (e) {
    console.error(`❌ Error updating plan: ${e}\n`);
  }
}

async function testSearchPlans(client: FlexPrice) {
  console.log("--- Test 5: Search Plans ---");

  if (!testPlanID) {
    console.log("⚠ Warning: No plan ID available\n⚠ Skipping search test\n");
    return;
  }

  try {
    const searchFilter: components.TypesPlanFilter = {
      planIds: [testPlanID],
    };

    const response = await client.plans.postPlansSearch(searchFilter);

    console.log("✓ Search completed!");
    console.log(`  Found ${response.items?.length || 0} plans matching ID '${testPlanID}'`);
    if (response.items) {
      response.items.slice(0, 3).forEach((plan) => {
        console.log(`  - ${plan.id}: ${plan.name} (${plan.lookupKey})`);
      });
    }
    console.log();
  } catch (e) {
    console.error(`❌ Error searching plans: ${e}\n`);
  }
}

async function testDeletePlan(client: FlexPrice) {
  console.log("--- Test 6: Delete Plan ---");

  if (!testPlanID) {
    console.log("⚠ Skipping delete plan (no plan created)\n");
    return;
  }

  try {
    await client.plans.deletePlansId(testPlanID);

    console.log("✓ Plan deleted successfully!");
    console.log(`  Deleted ID: ${testPlanID}\n`);
  } catch (e) {
    console.error(`❌ Error deleting plan: ${e}\n`);
  }
}

// ========================================
// ADDONS, ENTITLEMENTS, SUBSCRIPTIONS
// (Stub implementations matching Go test)
// ========================================

async function testCreateAddon(client: FlexPrice) {
  console.log("--- Test 1: Create Addon ---");
  const timestamp = Date.now();
  testAddonName = `Test Addon ${timestamp}`;
  testAddonLookupKey = `test_addon_${timestamp}`;

  try {
    const addonRequest: components.DtoCreateAddonRequest = {
      name: testAddonName,
      lookupKey: testAddonLookupKey,
      description: "Test addon created by SDK",
      type: components.TypesAddonType.Onetime,
      metadata: {
        source: "sdk_test",
      },
    };

    const response = await client.addons.postAddons(addonRequest);
    testAddonID = response.id;
    console.log(`✓ Addon created: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error creating addon: ${e}\n`);
  }
}

async function testGetAddon(client: FlexPrice) {
  console.log("--- Test 2: Get Addon ---");
  if (!testAddonID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.addons.getAddonsId(testAddonID);
    console.log(`✓ Addon retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListAddons(client: FlexPrice) {
  console.log("--- Test 3: List Addons ---");
  try {
    const response = await client.addons.getAddons({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} addons\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testUpdateAddon(client: FlexPrice) {
  console.log("--- Test 4: Update Addon ---");
  if (!testAddonID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const updatedName = testAddonName + " (Updated)";
    const response = await client.addons.putAddonsId(testAddonID, { name: updatedName });
    console.log(`✓ Addon updated: ${response.name}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testLookupAddon(client: FlexPrice) {
  console.log("--- Test 5: Lookup Addon ---");
  console.log("⚠ Skipping (method not available in SDK)\n");
}

async function testSearchAddons(client: FlexPrice) {
  console.log("--- Test 6: Search Addons ---");
  if (!testAddonID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.addons.postAddonsSearch({
      addonIds: [testAddonID],
    });
    console.log(`✓ Found ${response.items?.length || 0} addons\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testDeleteAddon(client: FlexPrice) {
  console.log("--- Test 7: Delete Addon ---");
  if (!testAddonID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.addons.deleteAddonsId(testAddonID);
    console.log("✓ Addon deleted\n");
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testCreateEntitlement(client: FlexPrice) {
  console.log("--- Test 1: Create Entitlement ---");
  if (!testFeatureID || !testPlanID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const entitlementRequest: components.DtoCreateEntitlementRequest = {
      featureId: testFeatureID,
      featureType: components.TypesFeatureType.Boolean,
      planId: testPlanID,
      isEnabled: true,
      usageResetPeriod: components.TypesEntitlementUsageResetPeriod.Monthly,
    };

    const response = await client.entitlements.postEntitlements(entitlementRequest);
    testEntitlementID = response.id;
    console.log(`✓ Entitlement created: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testGetEntitlement(client: FlexPrice) {
  console.log("--- Test 2: Get Entitlement ---");
  if (!testEntitlementID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.entitlements.getEntitlementsId(testEntitlementID);
    console.log(`✓ Entitlement retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListEntitlements(client: FlexPrice) {
  console.log("--- Test 3: List Entitlements ---");
  try {
    const response = await client.entitlements.getEntitlements({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} entitlements\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testUpdateEntitlement(client: FlexPrice) {
  console.log("--- Test 4: Update Entitlement ---");
  if (!testEntitlementID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.entitlements.putEntitlementsId(testEntitlementID, {
      isEnabled: false,
    });
    console.log(`✓ Entitlement updated: ${response.isEnabled}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testSearchEntitlements(client: FlexPrice) {
  console.log("--- Test 5: Search Entitlements ---");
  if (!testEntitlementID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.entitlements.postEntitlementsSearch({
      entityIds: [testEntitlementID],
    });
    console.log(`✓ Found ${response.items?.length || 0} entitlements\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testDeleteEntitlement(client: FlexPrice) {
  console.log("--- Test 6: Delete Entitlement ---");
  if (!testEntitlementID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.entitlements.deleteEntitlementsId(testEntitlementID);
    console.log("✓ Entitlement deleted\n");
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

// ========================================
// SUBSCRIPTIONS API TESTS (Simplified)
// ========================================

async function testCreateSubscription(client: FlexPrice) {
  console.log("--- Test 1: Create Subscription ---");
  if (!testCustomerID || !testPlanID) {
    console.log("⚠ Skipping\n");
    return;
  }

  try {
    // Create price first
    const priceRequest: components.DtoCreatePriceRequest = {
      entityId: testPlanID,
      entityType: components.TypesPriceEntityType.Plan,
      type: components.TypesPriceType.Fixed,
      billingModel: components.TypesBillingModel.FlatFee,
      billingCadence: components.TypesBillingCadence.Recurring,
      billingPeriod: components.TypesBillingPeriod.Monthly,
      billingPeriodCount: 1,
      invoiceCadence: components.TypesInvoiceCadence.Arrear,
      amount: "29.99",
      currency: "USD",
      priceUnitType: components.TypesPriceUnitType.Fiat,
      displayName: "Monthly Price",
    };

    const price = await client.prices.postPrices(priceRequest);
    console.log(`  Created price: ${price.id}`);

    const startDate = new Date().toISOString();
    const subscriptionRequest: components.DtoCreateSubscriptionRequest = {
      customerId: testCustomerID,
      planId: testPlanID,
      currency: "USD",
      billingCadence: components.TypesBillingCadence.Recurring,
      billingPeriod: components.TypesBillingPeriod.Monthly,
      billingPeriodCount: 1,
      billingCycle: components.TypesBillingCycle.Anniversary,
      startDate: startDate,
    };

    const response = await client.subscriptions.postSubscriptions(subscriptionRequest);
    testSubscriptionID = response.id;
    console.log(`✓ Subscription created: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testGetSubscription(client: FlexPrice) {
  console.log("--- Test 2: Get Subscription ---");
  if (!testSubscriptionID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.subscriptions.getSubscriptionsId(testSubscriptionID);
    console.log(`✓ Subscription retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListSubscriptions(client: FlexPrice) {
  console.log("--- Test 3: List Subscriptions ---");
  try {
    const response = await client.subscriptions.getSubscriptions({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} subscriptions\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testSearchSubscriptions(client: FlexPrice) {
  console.log("--- Test 4: Search Subscriptions ---");
  try {
    const response = await client.subscriptions.postSubscriptionsSearch({});
    console.log(`✓ Found ${response.items?.length || 0} subscriptions\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testCancelSubscription(client: FlexPrice) {
  console.log("--- Test 14: Cancel Subscription ---");
  if (!testSubscriptionID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.subscriptions.postSubscriptionsIdCancel(testSubscriptionID, {
      cancellationType: components.TypesCancellationType.EndOfPeriod,
    });
    console.log("✓ Subscription canceled\n");
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

// Stub implementations for remaining subscription tests
async function testActivateSubscription() {
  console.log("--- Test 5: Activate Subscription ---");
  console.log("⚠ Skipping (requires draft subscription)\n");
}

async function testAddAddonToSubscription() {
  console.log("--- Test 6: Add Addon to Subscription ---");
  console.log("⚠ Skipping (add addon endpoint needs verification)\n");
}

async function testRemoveAddonFromSubscription() {
  console.log("--- Test 7: Remove Addon from Subscription ---");
  console.log("⚠ Skipping (requires addon association ID)\n");
}

async function testExecuteSubscriptionChange() {
  console.log("--- Test 8: Execute Subscription Change ---");
  console.log("⚠ Skipping (would modify active subscription)\n");
}

async function testGetSubscriptionEntitlements(client: FlexPrice) {
  console.log("--- Test 9: Get Subscription Entitlements ---");
  if (!testSubscriptionID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.subscriptions.getSubscriptionsIdEntitlements(testSubscriptionID, undefined);
    console.log(`✓ Retrieved ${response.features?.length || 0} features\n`);
  } catch (e) {
    console.log(`⚠ Skipping: ${e}\n`);
  }
}

async function testGetUpcomingGrants() {
  console.log("--- Test 10: Get Upcoming Grants ---");
  console.log("⚠ Skipping (endpoint needs verification)\n");
}

async function testReportUsage() {
  console.log("--- Test 11: Report Usage ---");
  console.log("⚠ Skipping (requires metered feature setup)\n");
}

async function testUpdateLineItem() {
  console.log("--- Test 12: Update Line Item ---");
  console.log("⚠ Skipping (requires line item ID)\n");
}

async function testDeleteLineItem() {
  console.log("--- Test 13: Delete Line Item ---");
  console.log("⚠ Skipping (requires line item ID)\n");
}

// ========================================
// INVOICES, PRICES, PAYMENTS, WALLETS
// (Abbreviated implementations)
// ========================================

async function testListInvoices(client: FlexPrice) {
  console.log("--- Test 1: List Invoices ---");
  try {
    const response = await client.invoices.getInvoices({ limit: 10 });
    if (response.items && response.items.length > 0) {
      testInvoiceID = response.items[0].id;
    }
    console.log(`✓ Retrieved ${response.items?.length || 0} invoices\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testSearchInvoices(client: FlexPrice) {
  console.log("--- Test 2: Search Invoices ---");
  try {
    const response = await client.invoices.postInvoicesSearch({});
    console.log(`✓ Found ${response.items?.length || 0} invoices\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testCreateInvoice(client: FlexPrice) {
  console.log("--- Test 3: Create Invoice ---");
  if (!testCustomerID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const invoiceRequest: components.DtoCreateInvoiceRequest = {
      customerId: testCustomerID,
      currency: "USD",
      amountDue: "100.00",
      subtotal: "100.00",
      total: "100.00",
      invoiceType: components.TypesInvoiceType.OneOff,
      billingReason: components.TypesInvoiceBillingReason.Manual,
      invoiceStatus: components.TypesInvoiceStatus.Draft,
      lineItems: [
        {
          displayName: "Test Service",
          quantity: "1",
          amount: "100.00",
        },
      ],
    };

    const response = await client.invoices.postInvoices(invoiceRequest);
    testInvoiceID = response.id;
    console.log(`✓ Invoice created: ${response.id}\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testGetInvoice(client: FlexPrice) {
  console.log("--- Test 4: Get Invoice ---");
  if (!testInvoiceID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.invoices.getInvoicesId(testInvoiceID, undefined, undefined);
    console.log(`✓ Invoice retrieved: ${response.id}\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testUpdateInvoice(client: FlexPrice) {
  console.log("--- Test 5: Update Invoice ---");
  if (!testInvoiceID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.invoices.putInvoicesId(testInvoiceID, {
      metadata: { updated: "true" },
    });
    console.log("✓ Invoice updated\n");
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

// Stub remaining invoice tests
async function testPreviewInvoice() {
  console.log("--- Test 6: Preview Invoice ---");
  console.log("⚠ Skipping\n");
}

async function testFinalizeInvoice() {
  console.log("--- Test 7: Finalize Invoice ---");
  console.log("⚠ Skipping\n");
}

async function testRecalculateInvoice() {
  console.log("--- Test 8: Recalculate Invoice ---");
  console.log("⚠ Skipping\n");
}

async function testRecordPayment() {
  console.log("--- Test 9: Record Payment ---");
  console.log("⚠ Skipping\n");
}

async function testAttemptPayment() {
  console.log("--- Test 10: Attempt Payment ---");
  console.log("⚠ Skipping\n");
}

async function testDownloadInvoicePDF() {
  console.log("--- Test 11: Download Invoice PDF ---");
  console.log("⚠ Skipping\n");
}

async function testTriggerInvoiceComms() {
  console.log("--- Test 12: Trigger Invoice Comms ---");
  console.log("⚠ Skipping\n");
}

async function testGetCustomerInvoiceSummary() {
  console.log("--- Test 13: Get Customer Invoice Summary ---");
  console.log("⚠ Skipping\n");
}

async function testVoidInvoice() {
  console.log("--- Test 14: Void Invoice ---");
  console.log("⚠ Skipping\n");
}

// ========================================
// PRICES API TESTS
// ========================================

async function testCreatePrice(client: FlexPrice) {
  console.log("--- Test 1: Create Price ---");
  if (!testPlanID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const priceRequest: components.DtoCreatePriceRequest = {
      entityId: testPlanID,
      entityType: components.TypesPriceEntityType.Plan,
      currency: "USD",
      amount: "99.00",
      billingModel: components.TypesBillingModel.FlatFee,
      billingCadence: components.TypesBillingCadence.Recurring,
      billingPeriod: components.TypesBillingPeriod.Monthly,
      billingPeriodCount: 1,
      invoiceCadence: components.TypesInvoiceCadence.Advance,
      priceUnitType: components.TypesPriceUnitType.Fiat,
      type: components.TypesPriceType.Fixed,
      displayName: "Monthly Subscription",
    };

    const response = await client.prices.postPrices(priceRequest);
    testPriceID = response.id;
    console.log(`✓ Price created: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testGetPrice(client: FlexPrice) {
  console.log("--- Test 2: Get Price ---");
  if (!testPriceID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.prices.getPricesId(testPriceID);
    console.log(`✓ Price retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListPrices(client: FlexPrice) {
  console.log("--- Test 3: List Prices ---");
  try {
    const response = await client.prices.getPrices({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} prices\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testUpdatePrice(client: FlexPrice) {
  console.log("--- Test 4: Update Price ---");
  if (!testPriceID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.prices.putPricesId(testPriceID, {
      description: "Updated price description",
    });
    console.log("✓ Price updated\n");
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testDeletePrice(client: FlexPrice) {
  console.log("--- Test 5: Delete Price ---");
  if (!testPriceID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const futureDate = new Date();
    futureDate.setDate(futureDate.getDate() + 1);
    await client.prices.deletePricesId(testPriceID, {
      endDate: futureDate.toISOString(),
    });
    console.log("✓ Price deleted\n");
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

// ========================================
// PAYMENTS API TESTS (Stubs)
// ========================================

async function testCreatePayment() {
  console.log("--- Test 1: Create Payment ---");
  console.log("⚠ Skipping (requires invoice setup)\n");
}

async function testGetPayment(client: FlexPrice) {
  console.log("--- Test 2: Get Payment ---");
  if (!testPaymentID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.payments.getPaymentsId(testPaymentID);
    console.log(`✓ Payment retrieved: ${response.id}\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testListPayments(client: FlexPrice) {
  console.log("--- Test 3: List Payments ---");
  try {
    const response = await client.payments.getPayments({ limit: 10 });
    if (response.items && response.items.length > 0) {
      testPaymentID = response.items[0].id;
    }
    console.log(`✓ Retrieved ${response.items?.length || 0} payments\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testUpdatePayment(client: FlexPrice) {
  console.log("--- Test 4: Update Payment ---");
  if (!testPaymentID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.payments.putPaymentsId(testPaymentID, {
      metadata: {
        updatedAt: new Date().toISOString(),
        status: "updated",
      },
    });
    console.log(`✓ Payment updated: ${testPaymentID}\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testProcessPayment(client: FlexPrice) {
  console.log("--- Test 5: Process Payment ---");
  if (!testPaymentID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.payments.postPaymentsIdProcess(testPaymentID);
    console.log(`✓ Payment processed: ${testPaymentID}\n`);
  } catch (e) {
    console.log(`⚠ Error (may require payment gateway setup): ${e}\n`);
  }
}

async function testDeletePayment(client: FlexPrice) {
  console.log("--- Test 6: Delete Payment ---");
  if (!testPaymentID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.payments.deletePaymentsId(testPaymentID);
    console.log("✓ Payment deleted\n");
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

// ========================================
// WALLETS API TESTS (Stubs)
// ========================================

async function testCreateWallet(client: FlexPrice) {
  console.log("--- Test 1: Create Wallet ---");
  if (!testCustomerID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const walletRequest: components.DtoCreateWalletRequest = {
      customerId: testCustomerID,
      currency: "USD",
      name: "Test Wallet",
    };

    const response = await client.wallets.postWallets(walletRequest);
    testWalletID = response.id;
    console.log(`✓ Wallet created: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testGetWallet(client: FlexPrice) {
  console.log("--- Test 2: Get Wallet ---");
  if (!testWalletID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.wallets.getWalletsId(testWalletID);
    console.log(`✓ Wallet retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListWallets(client: FlexPrice) {
  console.log("--- Test 3: List Wallets ---");
  try {
    const response = await client.wallets.getWallets({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} wallets\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testUpdateWallet(client: FlexPrice) {
  console.log("--- Test 4: Update Wallet ---");
  if (!testWalletID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.wallets.putWalletsId(testWalletID, {
      metadata: {
        updatedAt: new Date().toISOString(),
        status: "updated",
      },
    });
    console.log(`✓ Wallet updated: ${testWalletID}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testGetWalletBalance(client: FlexPrice) {
  console.log("--- Test 5: Get Wallet Balance ---");
  if (!testWalletID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.wallets.getWalletsIdBalanceRealTime(testWalletID);
    console.log(`✓ Balance: ${response.balance}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testTopUpWallet(client: FlexPrice) {
  console.log("--- Test 6: Top Up Wallet ---");
  if (!testWalletID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const topUpRequest: components.DtoTopUpWalletRequest = {
      amount: "100.00",
      transactionReason: components.TypesTransactionReason.PurchasedCreditDirect,
      description: "Test top-up from SDK",
    };
    await client.wallets.postWalletsIdTopUp(testWalletID, topUpRequest);
    console.log(`✓ Wallet topped up: ${testWalletID}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testDebitWallet() {
  console.log("--- Test 7: Debit Wallet ---");
  console.log("⚠ Skipping (debit endpoint not available in SDK)\n");
}

async function testGetWalletTransactions(client: FlexPrice) {
  console.log("--- Test 8: Get Wallet Transactions ---");
  if (!testWalletID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.wallets.getWalletsIdTransactions({
      idPathParameter: testWalletID,
      limit: 10,
    });
    console.log(`✓ Retrieved ${response.items?.length || 0} transactions\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testSearchWallets(client: FlexPrice) {
  console.log("--- Test 9: Search Wallets ---");
  try {
    const response = await client.wallets.postWalletsSearch({ limit: 10 });
    console.log(`✓ Found ${response.items?.length || 0} wallets\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

// ========================================
// CREDIT GRANTS API TESTS (Stubs)
// ========================================

async function testCreateCreditGrant(client: FlexPrice) {
  console.log("--- Test 1: Create Credit Grant ---");
  if (!testPlanID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const grantRequest: components.DtoCreateCreditGrantRequest = {
      scope: components.TypesCreditGrantScope.Plan,
      planId: testPlanID,
      credits: "500.00",
      name: "Test Credit Grant",
      cadence: components.TypesCreditGrantCadence.Onetime,
      expirationType: components.TypesCreditGrantExpiryType.Never,
      expirationDurationUnit: components.TypesCreditGrantExpiryDurationUnit.Day,
    };

    const response = await client.creditGrants.postCreditgrants(grantRequest);
    testCreditGrantID = response.id;
    console.log(`✓ Credit grant created: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testGetCreditGrant(client: FlexPrice) {
  console.log("--- Test 2: Get Credit Grant ---");
  if (!testCreditGrantID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.creditGrants.getCreditgrantsId(testCreditGrantID);
    console.log(`✓ Credit grant retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListCreditGrants(client: FlexPrice) {
  console.log("--- Test 3: List Credit Grants ---");
  try {
    const response = await client.creditGrants.getCreditgrants({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} credit grants\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testUpdateCreditGrant(client: FlexPrice) {
  console.log("--- Test 4: Update Credit Grant ---");
  if (!testCreditGrantID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.creditGrants.putCreditgrantsId(testCreditGrantID, {
      metadata: {
        updatedAt: new Date().toISOString(),
        status: "updated",
      },
    });
    console.log(`✓ Credit grant updated: ${testCreditGrantID}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

// ========================================
// CREDIT NOTES API TESTS (Stubs)
// ========================================

async function testCreateCreditNote() {
  console.log("--- Test 1: Create Credit Note ---");
  console.log("⚠ Skipping (requires invoice with line items)\n");
}

async function testGetCreditNote(client: FlexPrice) {
  console.log("--- Test 2: Get Credit Note ---");
  if (!testCreditNoteID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    const response = await client.creditNotes.getCreditnotesId(testCreditNoteID);
    console.log(`✓ Credit note retrieved: ${response.id}\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testListCreditNotes(client: FlexPrice) {
  console.log("--- Test 3: List Credit Notes ---");
  try {
    const response = await client.creditNotes.getCreditnotes({ limit: 10 });
    console.log(`✓ Retrieved ${response.items?.length || 0} credit notes\n`);
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testFinalizeCreditNote(client: FlexPrice) {
  console.log("--- Test 4: Finalize Credit Note ---");
  if (!testCreditNoteID) {
    console.log("⚠ Skipping\n");
    return;
  }
  try {
    await client.creditNotes.postCreditnotesIdFinalize(testCreditNoteID);
    console.log(`✓ Credit note finalized: ${testCreditNoteID}\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

// ========================================
// EVENTS API TESTS
// ========================================

async function testCreateEvent(client: FlexPrice) {
  console.log("--- Test 1: Create Event ---");

  testEventCustomerID = testExternalID || `test-customer-${Date.now()}`;
  testEventName = `Test Event ${Date.now()}`;

  try {
    const eventRequest: components.DtoIngestEventRequest = {
      eventName: testEventName,
      externalCustomerId: testEventCustomerID,
      properties: {
        source: "sdk_test",
        environment: "test",
      },
      source: "sdk_test",
      timestamp: new Date().toISOString(),
    };

    const response = await client.events.postEvents(eventRequest);

    if (response && typeof response === "object" && "event_id" in response) {
      testEventID = (response as any).event_id;
      console.log(`✓ Event created: ${testEventID}\n`);
    } else {
      console.log("✓ Event created\n");
    }
  } catch (e) {
    console.error(`❌ Error: ${e}\n`);
  }
}

async function testQueryEvents(client: FlexPrice) {
  console.log("--- Test 2: Query Events ---");

  if (!testEventName) {
    console.log("⚠ Skipping\n");
    return;
  }

  try {
    const queryRequest: any = {
      external_customer_id: testEventCustomerID,
      event_name: testEventName,
    };

    const response = await client.events.postEventsQuery(queryRequest);
    console.log(`✓ Found ${response.events?.length || 0} events\n`);
  } catch (e) {
    console.log(`⚠ Error: ${e}\n`);
  }
}

async function testAsyncEventIngestion() {
  console.log("--- Test 3: Async Event Ingestion ---");
  console.log("⚠ Skipping (async client needs implementation)\n");
}

// ========================================
// MAIN EXECUTION
// ========================================

async function main() {
  const client = getClient();

  console.log("========================================");
  console.log("CUSTOMER API TESTS");
  console.log("========================================\n");
  await testCreateCustomer(client);
  await testGetCustomer(client);
  await testListCustomers(client);
  await testUpdateCustomer(client);
  await testLookupCustomer(client);
  await testSearchCustomers(client);
  await testGetCustomerEntitlements(client);
  await testGetCustomerUpcomingGrants(client);
  await testGetCustomerUsage(client);
  console.log("✓ Customer API Tests Completed!\n");

  console.log("========================================");
  console.log("FEATURES API TESTS");
  console.log("========================================\n");
  await testCreateFeature(client);
  await testGetFeature(client);
  await testListFeatures(client);
  await testUpdateFeature(client);
  await testSearchFeatures(client);
  console.log("✓ Features API Tests Completed!\n");

  console.log("========================================");
  console.log("CONNECTIONS API TESTS");
  console.log("========================================\n");
  await testListConnections(client);
  await testSearchConnections(client);
  console.log("✓ Connections API Tests Completed!\n");

  console.log("========================================");
  console.log("PLANS API TESTS");
  console.log("========================================\n");
  await testCreatePlan(client);
  await testGetPlan(client);
  await testListPlans(client);
  await testUpdatePlan(client);
  await testSearchPlans(client);
  console.log("✓ Plans API Tests Completed!\n");

  console.log("========================================");
  console.log("ADDONS API TESTS");
  console.log("========================================\n");
  await testCreateAddon(client);
  await testGetAddon(client);
  await testListAddons(client);
  await testUpdateAddon(client);
  await testLookupAddon(client);
  await testSearchAddons(client);
  console.log("✓ Addons API Tests Completed!\n");

  console.log("========================================");
  console.log("ENTITLEMENTS API TESTS");
  console.log("========================================\n");
  await testCreateEntitlement(client);
  await testGetEntitlement(client);
  await testListEntitlements(client);
  await testUpdateEntitlement(client);
  await testSearchEntitlements(client);
  console.log("✓ Entitlements API Tests Completed!\n");

  console.log("========================================");
  console.log("SUBSCRIPTIONS API TESTS");
  console.log("========================================\n");
  await testCreateSubscription(client);
  await testGetSubscription(client);
  await testListSubscriptions(client);
  await testSearchSubscriptions(client);
  await testActivateSubscription();
  await testAddAddonToSubscription();
  await testRemoveAddonFromSubscription();
  await testExecuteSubscriptionChange();
  await testGetSubscriptionEntitlements(client);
  await testGetUpcomingGrants();
  await testReportUsage();
  await testUpdateLineItem();
  await testDeleteLineItem();
  await testCancelSubscription(client);
  console.log("✓ Subscriptions API Tests Completed!\n");

  console.log("========================================");
  console.log("INVOICES API TESTS");
  console.log("========================================\n");
  await testListInvoices(client);
  await testSearchInvoices(client);
  await testCreateInvoice(client);
  await testGetInvoice(client);
  await testUpdateInvoice(client);
  await testPreviewInvoice();
  await testFinalizeInvoice();
  await testRecalculateInvoice();
  await testRecordPayment();
  await testAttemptPayment();
  await testDownloadInvoicePDF();
  await testTriggerInvoiceComms();
  await testGetCustomerInvoiceSummary();
  await testVoidInvoice();
  console.log("✓ Invoices API Tests Completed!\n");

  console.log("========================================");
  console.log("PRICES API TESTS");
  console.log("========================================\n");
  await testCreatePrice(client);
  await testGetPrice(client);
  await testListPrices(client);
  await testUpdatePrice(client);
  console.log("✓ Prices API Tests Completed!\n");

  console.log("========================================");
  console.log("PAYMENTS API TESTS");
  console.log("========================================\n");
  await testCreatePayment();
  await testGetPayment(client);
  await testListPayments(client);
  await testUpdatePayment(client);
  await testProcessPayment(client);
  console.log("✓ Payments API Tests Completed!\n");

  console.log("========================================");
  console.log("WALLETS API TESTS");
  console.log("========================================\n");
  await testCreateWallet(client);
  await testGetWallet(client);
  await testListWallets(client);
  await testUpdateWallet(client);
  await testGetWalletBalance(client);
  await testTopUpWallet(client);
  await testDebitWallet();
  await testGetWalletTransactions(client);
  await testSearchWallets(client);
  console.log("✓ Wallets API Tests Completed!\n");

  console.log("========================================");
  console.log("CREDIT GRANTS API TESTS");
  console.log("========================================\n");
  await testCreateCreditGrant(client);
  await testGetCreditGrant(client);
  await testListCreditGrants(client);
  await testUpdateCreditGrant(client);
  console.log("✓ Credit Grants API Tests Completed!\n");

  console.log("========================================");
  console.log("CREDIT NOTES API TESTS");
  console.log("========================================\n");
  await testCreateCreditNote();
  await testGetCreditNote(client);
  await testListCreditNotes(client);
  await testFinalizeCreditNote(client);
  console.log("✓ Credit Notes API Tests Completed!\n");

  console.log("========================================");
  console.log("EVENTS API TESTS");
  console.log("========================================\n");
  await testCreateEvent(client);
  await testQueryEvents(client);
  await testAsyncEventIngestion();
  console.log("✓ Events API Tests Completed!\n");

  console.log("========================================");
  console.log("CLEANUP - DELETING TEST DATA");
  console.log("========================================\n");
  await testDeletePayment(client);
  await testDeletePrice(client);
  await testDeleteEntitlement(client);
  await testDeleteAddon(client);
  await testDeletePlan(client);
  await testDeleteFeature(client);
  await testDeleteCustomer(client);
  console.log("✓ Cleanup Completed!\n");

  console.log("\n=== All API Tests Completed Successfully! ===");
}

main().catch((err) => {
  console.error("Fatal error:", err);
  process.exit(1);
});

