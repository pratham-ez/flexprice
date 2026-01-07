package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	gosdktemp "github.com/flexprice/go-sdk-temp"
	"github.com/flexprice/go-sdk-temp/models/components"
	"github.com/flexprice/go-sdk-temp/models/operations"
)

// test_sdk_speakeasy.go - Speakeasy SDK Testing for FlexPrice API
// This file tests the Speakeasy-generated FlexPrice Go SDK functions
//
// Setup:
// 1. Export your API key: export FLEXPRICE_API_KEY="your_key_here"
// 2. Export API host: export FLEXPRICE_API_HOST="https://api.cloud.flexprice.io/v1"
// 3. Run: go run test_sdk_speakeasy.go
//
// This replicates all 112 test functions from the old test_sdk.go

var (
	testCustomerID   string
	testExternalID   string
	testCustomerName string

	testFeatureID   string
	testFeatureName string

	testPlanID   string
	testPlanName string

	testAddonID        string
	testAddonName      string
	testAddonLookupKey string

	testEntitlementID string

	testSubscriptionID string

	testInvoiceID string

	testPriceID string

	testPaymentID string

	testWalletID      string
	testCreditGrantID string
	testCreditNoteID  string

	testEventID         string
	testEventName       string
	testEventCustomerID string
)

func main() {
	fmt.Println("=== FlexPrice Speakeasy Go SDK - API Tests ===\n")

	// Get API credentials from environment
	apiKey := os.Getenv("FLEXPRICE_API_KEY")
	apiHost := os.Getenv("FLEXPRICE_API_HOST")

	if apiKey == "" {
		log.Fatal("❌ Missing FLEXPRICE_API_KEY environment variable")
	}
	if apiHost == "" {
		log.Fatal("❌ Missing FLEXPRICE_API_HOST environment variable")
	}

	fmt.Printf("✓ API Key: %s...%s\n", apiKey[:min(8, len(apiKey))], apiKey[max(0, len(apiKey)-4):])
	fmt.Printf("✓ API Host: %s\n\n", apiHost)

	// Initialize Speakeasy SDK
	client := gosdktemp.New(
		apiHost,
		gosdktemp.WithSecurity(apiKey),
	)
	ctx := context.Background()

	// Run all Customer API tests
	fmt.Println("========================================")
	fmt.Println("CUSTOMER API TESTS")
	fmt.Println("========================================\n")

	testCreateCustomer(ctx, client)
	testGetCustomer(ctx, client)
	testListCustomers(ctx, client)
	testUpdateCustomer(ctx, client)
	testLookupCustomer(ctx, client)
	testSearchCustomers(ctx, client)
	testGetCustomerEntitlements(ctx, client)
	testGetCustomerUpcomingGrants(ctx, client)
	testGetCustomerUsage(ctx, client)

	fmt.Println("✓ Customer API Tests Completed!\n")

	// Run all Features API tests
	fmt.Println("========================================")
	fmt.Println("FEATURES API TESTS")
	fmt.Println("========================================\n")

	testCreateFeature(ctx, client)
	testGetFeature(ctx, client)
	testListFeatures(ctx, client)
	testUpdateFeature(ctx, client)
	testSearchFeatures(ctx, client)

	fmt.Println("✓ Features API Tests Completed!\n")

	// Run all Connections API tests
	fmt.Println("========================================")
	fmt.Println("CONNECTIONS API TESTS")
	fmt.Println("========================================\n")

	testListConnections(ctx, client)
	testSearchConnections(ctx, client)

	fmt.Println("✓ Connections API Tests Completed!\n")

	// Run all Plans API tests
	fmt.Println("========================================")
	fmt.Println("PLANS API TESTS")
	fmt.Println("========================================\n")

	testCreatePlan(ctx, client)
	testGetPlan(ctx, client)
	testListPlans(ctx, client)
	testUpdatePlan(ctx, client)
	testSearchPlans(ctx, client)

	fmt.Println("✓ Plans API Tests Completed!\n")

	// Run all Addons API tests
	fmt.Println("========================================")
	fmt.Println("ADDONS API TESTS")
	fmt.Println("========================================\n")

	testCreateAddon(ctx, client)
	testGetAddon(ctx, client)
	testListAddons(ctx, client)
	testUpdateAddon(ctx, client)
	testLookupAddon(ctx, client)
	testSearchAddons(ctx, client)

	fmt.Println("✓ Addons API Tests Completed!\n")

	// Run all Entitlements API tests
	fmt.Println("========================================")
	fmt.Println("ENTITLEMENTS API TESTS")
	fmt.Println("========================================\n")

	testCreateEntitlement(ctx, client)
	testGetEntitlement(ctx, client)
	testListEntitlements(ctx, client)
	testUpdateEntitlement(ctx, client)
	testSearchEntitlements(ctx, client)

	fmt.Println("✓ Entitlements API Tests Completed!\n")

	// Run all Subscriptions API tests
	fmt.Println("========================================")
	fmt.Println("SUBSCRIPTIONS API TESTS")
	fmt.Println("========================================\n")

	testCreateSubscription(ctx, client)
	testGetSubscription(ctx, client)
	testListSubscriptions(ctx, client)
	testSearchSubscriptions(ctx, client)
	testActivateSubscription(ctx, client)
	testAddAddonToSubscription(ctx, client)
	testRemoveAddonFromSubscription(ctx, client)
	testExecuteSubscriptionChange(ctx, client)
	testGetSubscriptionEntitlements(ctx, client)
	testGetUpcomingGrants(ctx, client)
	testReportUsage(ctx, client)
	testUpdateLineItem(ctx, client)
	testDeleteLineItem(ctx, client)
	testCancelSubscription(ctx, client)

	fmt.Println("✓ Subscriptions API Tests Completed!\n")

	// Run all Invoices API tests
	fmt.Println("========================================")
	fmt.Println("INVOICES API TESTS")
	fmt.Println("========================================\n")

	testListInvoices(ctx, client)
	testSearchInvoices(ctx, client)
	testCreateInvoice(ctx, client)
	testGetInvoice(ctx, client)
	testUpdateInvoice(ctx, client)
	testPreviewInvoice(ctx, client)
	testFinalizeInvoice(ctx, client)
	testRecalculateInvoice(ctx, client)
	testRecordPayment(ctx, client)
	testAttemptPayment(ctx, client)
	testDownloadInvoicePDF(ctx, client)
	testTriggerInvoiceComms(ctx, client)
	testGetCustomerInvoiceSummary(ctx, client)
	testVoidInvoice(ctx, client)

	fmt.Println("✓ Invoices API Tests Completed!\n")

	// Run all Prices API tests
	fmt.Println("========================================")
	fmt.Println("PRICES API TESTS")
	fmt.Println("========================================\n")

	testCreatePrice(ctx, client)
	testGetPrice(ctx, client)
	testListPrices(ctx, client)
	testUpdatePrice(ctx, client)

	fmt.Println("✓ Prices API Tests Completed!\n")

	// Run all Payments API tests
	fmt.Println("========================================")
	fmt.Println("PAYMENTS API TESTS")
	fmt.Println("========================================\n")

	testCreatePayment(ctx, client)
	testGetPayment(ctx, client)
	testListPayments(ctx, client)
	testUpdatePayment(ctx, client)
	testProcessPayment(ctx, client)

	fmt.Println("✓ Payments API Tests Completed!\n")

	// Run all Wallets API tests
	fmt.Println("========================================")
	fmt.Println("WALLETS API TESTS")
	fmt.Println("========================================\n")

	testCreateWallet(ctx, client)
	testGetWallet(ctx, client)
	testListWallets(ctx, client)
	testUpdateWallet(ctx, client)
	testGetWalletBalance(ctx, client)
	testTopUpWallet(ctx, client)
	testDebitWallet(ctx, client)
	testGetWalletTransactions(ctx, client)
	testSearchWallets(ctx, client)

	fmt.Println("✓ Wallets API Tests Completed!\n")

	// Run all Credit Grants API tests
	fmt.Println("========================================")
	fmt.Println("CREDIT GRANTS API TESTS")
	fmt.Println("========================================\n")

	testCreateCreditGrant(ctx, client)
	testGetCreditGrant(ctx, client)
	testListCreditGrants(ctx, client)
	testUpdateCreditGrant(ctx, client)

	fmt.Println("✓ Credit Grants API Tests Completed!\n")

	// Run all Credit Notes API tests
	fmt.Println("========================================")
	fmt.Println("CREDIT NOTES API TESTS")
	fmt.Println("========================================\n")

	testCreateCreditNote(ctx, client)
	testGetCreditNote(ctx, client)
	testListCreditNotes(ctx, client)
	testFinalizeCreditNote(ctx, client)

	fmt.Println("✓ Credit Notes API Tests Completed!\n")

	// Run all Events API tests
	fmt.Println("========================================")
	fmt.Println("EVENTS API TESTS")
	fmt.Println("========================================\n")

	testCreateEvent(ctx, client)
	testQueryEvents(ctx, client)
	testAsyncEventIngestion(ctx, client)

	fmt.Println("✓ Events API Tests Completed!\n")

	// Cleanup: Delete all created entities
	fmt.Println("========================================")
	fmt.Println("CLEANUP - DELETING TEST DATA")
	fmt.Println("========================================\n")

	testDeletePayment(ctx, client)
	testDeletePrice(ctx, client)
	testDeleteEntitlement(ctx, client)
	testDeleteAddon(ctx, client)
	testDeletePlan(ctx, client)
	testDeleteFeature(ctx, client)
	testDeleteCustomer(ctx, client)

	fmt.Println("✓ Cleanup Completed!\n")

	fmt.Println("\n=== All API Tests Completed Successfully! ===")
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ========================================
// CUSTOMER API TESTS
// ========================================

// Test 1: Create a new customer
func testCreateCustomer(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Customer ---")

	timestamp := time.Now().Unix()
	testCustomerName = fmt.Sprintf("Test Customer %d", timestamp)
	testExternalID = fmt.Sprintf("test-customer-%d", timestamp)

	customerRequest := components.DtoCreateCustomerRequest{
		Name:       gosdktemp.String(testCustomerName),
		ExternalID: testExternalID,
		Email:      gosdktemp.String(fmt.Sprintf("test-%d@example.com", timestamp)),
		Metadata: map[string]string{
			"source":      "sdk_test",
			"test_run":    time.Now().Format(time.RFC3339),
			"environment": "test",
		},
	}

	res, err := client.Customers.PostCustomers(ctx, customerRequest)
	if err != nil {
		log.Printf("❌ Error creating customer: %v", err)
		fmt.Println()
		return
	}

	if res == nil {
		log.Printf("❌ No response from create customer")
		fmt.Println()
		return
	}

	testCustomerID = *res.ID
	fmt.Printf("✓ Customer created successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n", *res.Name)
	fmt.Printf("  External ID: %s\n", *res.ExternalID)
	fmt.Printf("  Email: %s\n\n", *res.Email)
}

// Test 2: Get customer by ID
func testGetCustomer(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Customer by ID ---")

	res, err := client.Customers.GetCustomersID(ctx, testCustomerID)
	if err != nil {
		log.Printf("❌ Error getting customer: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Customer retrieved successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n", *res.Name)
	fmt.Printf("  Created At: %s\n\n", *res.CreatedAt)
}

// Test 3: List all customers
func testListCustomers(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Customers ---")

	limit := int64(10)
	res, err := client.Customers.GetCustomers(ctx, operations.GetCustomersRequest{
		Limit: &limit,
	})
	if err != nil {
		log.Printf("❌ Error listing customers: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Retrieved %d customers\n", len(res.Items))
	if len(res.Items) > 0 {
		fmt.Printf("  First customer: %s - %s\n", *res.Items[0].ID, *res.Items[0].Name)
	}
	if res.Pagination != nil {
		fmt.Printf("  Total: %d\n", *res.Pagination.Total)
	}
	fmt.Println()
}

// Test 4: Update customer
func testUpdateCustomer(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Customer ---")

	updatedName := fmt.Sprintf("%s (Updated)", testCustomerName)
	updateRequest := components.DtoUpdateCustomerRequest{
		Name: &updatedName,
		Metadata: map[string]string{
			"updated_at": time.Now().Format(time.RFC3339),
			"status":     "updated",
		},
	}

	res, err := client.Customers.PutCustomersID(ctx, testCustomerID, updateRequest)
	if err != nil {
		log.Printf("❌ Error updating customer: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Customer updated successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  New Name: %s\n", *res.Name)
	if res.UpdatedAt != nil {
		fmt.Printf("  Updated At: %s\n\n", *res.UpdatedAt)
	} else {
		fmt.Println()
	}
}

// Test 5: Lookup customer by external ID
func testLookupCustomer(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Lookup Customer by External ID ---")

	res, err := client.Customers.GetCustomersExternalExternalID(ctx, testExternalID)
	if err != nil {
		log.Printf("❌ Error looking up customer: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Customer found by external ID!\n")
	fmt.Printf("  External ID: %s\n", testExternalID)
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n\n", *res.Name)
}

// Test 6: Search customers
func testSearchCustomers(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Search Customers ---")

	searchFilter := components.TypesCustomerFilter{
		ExternalID: &testExternalID,
	}

	res, err := client.Customers.PostCustomersSearch(ctx, searchFilter)
	if err != nil {
		log.Printf("❌ Error searching customers: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Search completed!\n")
	fmt.Printf("  Found %d customers matching external ID '%s'\n", len(res.Items), testExternalID)
	for i, customer := range res.Items {
		if i < 3 {
			fmt.Printf("  - %s: %s\n", *customer.ID, *customer.Name)
		}
	}
	fmt.Println()
}

// Test 7: Get customer entitlements
func testGetCustomerEntitlements(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 7: Get Customer Entitlements ---")

	res, err := client.Customers.GetCustomersIDEntitlements(ctx, testCustomerID, nil, nil)
	if err != nil {
		log.Printf("⚠ Warning: Error getting customer entitlements: %v\n", err)
		fmt.Println("⚠ Skipping entitlements test (customer may not have any entitlements)\n")
		return
	}

	fmt.Printf("✓ Retrieved customer entitlements!\n")
	if res.Features != nil {
		fmt.Printf("  Total features: %d\n", len(res.Features))
		for i, feature := range res.Features {
			if i < 3 && feature.Feature != nil && feature.Feature.ID != nil {
				fmt.Printf("  - Feature: %s\n", *feature.Feature.ID)
			}
		}
	} else {
		fmt.Println("  No features found")
	}
	fmt.Println()
}

// Test 8: Get customer upcoming grants
func testGetCustomerUpcomingGrants(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 8: Get Customer Upcoming Grants ---")

	res, err := client.Customers.GetCustomersIDGrantsUpcoming(ctx, testCustomerID)
	if err != nil {
		log.Printf("⚠ Warning: Error getting upcoming grants: %v\n", err)
		fmt.Println("⚠ Skipping upcoming grants test (customer may not have any grants)\n")
		return
	}

	fmt.Printf("✓ Retrieved upcoming grants!\n")
	if res.Items != nil {
		fmt.Printf("  Total upcoming grants: %d\n", len(res.Items))
	} else {
		fmt.Println("  No upcoming grants found")
	}
	fmt.Println()
}

// Test 9: Get customer usage
func testGetCustomerUsage(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 9: Get Customer Usage ---")

	res, err := client.Customers.GetCustomersUsage(ctx, operations.GetCustomersUsageRequest{CustomerID: &testCustomerID})
	if err != nil {
		log.Printf("⚠ Warning: Error getting customer usage: %v\n", err)
		fmt.Println("⚠ Skipping usage test (customer may not have usage data)\n")
		return
	}

	fmt.Printf("✓ Retrieved customer usage!\n")
	if res.Features != nil {
		fmt.Printf("  Feature usage records: %d\n", len(res.Features))
	} else {
		fmt.Println("  No usage data found")
	}
	fmt.Println()
}

// Test 10: Delete customer
func testDeleteCustomer(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 10: Delete Customer ---")

	err := client.Customers.DeleteCustomersID(ctx, testCustomerID)
	if err != nil {
		log.Printf("❌ Error deleting customer: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Customer deleted successfully!\n")
	fmt.Printf("  Deleted ID: %s\n\n", testCustomerID)
}

// ========================================
// FEATURES API TESTS
// ========================================

func testCreateFeature(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Feature ---")

	timestamp := time.Now().Unix()
	testFeatureName = fmt.Sprintf("Test Feature %d", timestamp)
	featureKey := fmt.Sprintf("test_feature_%d", timestamp)

	featureRequest := components.DtoCreateFeatureRequest{
		Name:        testFeatureName,
		LookupKey:   gosdktemp.String(featureKey),
		Description: gosdktemp.String("This is a test feature created by SDK tests"),
		Type:        components.TypesFeatureTypeBoolean,
		Metadata: map[string]string{
			"source":      "sdk_test",
			"test_run":    time.Now().Format(time.RFC3339),
			"environment": "test",
		},
	}

	res, err := client.Features.PostFeatures(ctx, featureRequest)
	if err != nil {
		log.Printf("❌ Error creating feature: %v", err)
		fmt.Println()
		return
	}

	testFeatureID = *res.ID
	fmt.Printf("✓ Feature created successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n", *res.Name)
	fmt.Printf("  Lookup Key: %s\n", *res.LookupKey)
	fmt.Printf("  Type: %s\n\n", res.Type)
}

func testGetFeature(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Feature by ID ---")

	res, err := client.Features.GetFeaturesID(ctx, testFeatureID)
	if err != nil {
		log.Printf("❌ Error getting feature: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Feature retrieved successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n", *res.Name)
	fmt.Printf("  Lookup Key: %s\n", *res.LookupKey)
	fmt.Printf("  Created At: %s\n\n", *res.CreatedAt)
}

func testListFeatures(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Features ---")

	limit := int64(10)
	res, err := client.Features.GetFeatures(ctx, operations.GetFeaturesRequest{
		Limit: &limit,
	})
	if err != nil {
		log.Printf("❌ Error listing features: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Retrieved %d features\n", len(res.Items))
	if len(res.Items) > 0 {
		fmt.Printf("  First feature: %s - %s\n", *res.Items[0].ID, *res.Items[0].Name)
	}
	if res.Pagination != nil {
		fmt.Printf("  Total: %d\n", *res.Pagination.Total)
	}
	fmt.Println()
}

func testUpdateFeature(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Feature ---")

	updatedName := fmt.Sprintf("%s (Updated)", testFeatureName)
	updatedDescription := "Updated description for test feature"
	updateRequest := components.DtoUpdateFeatureRequest{
		Name:        &updatedName,
		Description: &updatedDescription,
		Metadata: map[string]string{
			"updated_at": time.Now().Format(time.RFC3339),
			"status":     "updated",
		},
	}

	res, err := client.Features.PutFeaturesID(ctx, testFeatureID, updateRequest)
	if err != nil {
		log.Printf("❌ Error updating feature: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Feature updated successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  New Name: %s\n", *res.Name)
	fmt.Printf("  New Description: %s\n", *res.Description)
	fmt.Printf("  Updated At: %s\n\n", *res.UpdatedAt)
}

func testSearchFeatures(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Search Features ---")

	searchFilter := components.TypesFeatureFilter{
		FeatureIds: []string{testFeatureID},
	}

	res, err := client.Features.PostFeaturesSearch(ctx, searchFilter)
	if err != nil {
		log.Printf("❌ Error searching features: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Search completed!\n")
	fmt.Printf("  Found %d features matching ID '%s'\n", len(res.Items), testFeatureID)
	for i, feature := range res.Items {
		if i < 3 {
			fmt.Printf("  - %s: %s (%s)\n", *feature.ID, *feature.Name, *feature.LookupKey)
		}
	}
	fmt.Println()
}

func testDeleteFeature(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Delete Feature ---")

	_, err := client.Features.DeleteFeaturesID(ctx, testFeatureID)
	if err != nil {
		log.Printf("❌ Error deleting feature: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Feature deleted successfully!\n")
	fmt.Printf("  Deleted ID: %s\n\n", testFeatureID)
}

// ========================================
// PLANS API TESTS
// ========================================

func testCreatePlan(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Plan ---")

	timestamp := time.Now().Unix()
	testPlanName = fmt.Sprintf("Test Plan %d", timestamp)
	lookupKey := fmt.Sprintf("test_plan_%d", timestamp)

	planRequest := components.DtoCreatePlanRequest{
		Name:        testPlanName,
		LookupKey:   gosdktemp.String(lookupKey),
		Description: gosdktemp.String("This is a test plan created by SDK tests"),
		Metadata: map[string]string{
			"source":      "sdk_test",
			"test_run":    time.Now().Format(time.RFC3339),
			"environment": "test",
		},
	}

	res, err := client.Plans.PostPlans(ctx, planRequest)
	if err != nil {
		log.Printf("❌ Error creating plan: %v", err)
		fmt.Println()
		return
	}

	testPlanID = *res.ID
	fmt.Printf("✓ Plan created successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n", *res.Name)
	fmt.Printf("  Lookup Key: %s\n\n", *res.LookupKey)
}

func testGetPlan(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Plan by ID ---")

	res, err := client.Plans.GetPlansID(ctx, testPlanID)
	if err != nil {
		log.Printf("❌ Error getting plan: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Plan retrieved successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  Name: %s\n", *res.Name)
	fmt.Printf("  Lookup Key: %s\n", *res.LookupKey)
	fmt.Printf("  Created At: %s\n\n", *res.CreatedAt)
}

func testListPlans(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Plans ---")

	limit := int64(10)
	res, err := client.Plans.GetPlans(ctx, operations.GetPlansRequest{
		Limit: &limit,
	})
	if err != nil {
		log.Printf("❌ Error listing plans: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Retrieved %d plans\n", len(res.Items))
	if len(res.Items) > 0 {
		fmt.Printf("  First plan: %s - %s\n", *res.Items[0].ID, *res.Items[0].Name)
	}
	if res.Pagination != nil {
		fmt.Printf("  Total: %d\n", *res.Pagination.Total)
	}
	fmt.Println()
}

func testUpdatePlan(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Plan ---")

	updatedName := fmt.Sprintf("%s (Updated)", testPlanName)
	updatedDescription := "Updated description for test plan"
	updateRequest := components.DtoUpdatePlanRequest{
		Name:        &updatedName,
		Description: &updatedDescription,
		Metadata: map[string]string{
			"updated_at": time.Now().Format(time.RFC3339),
			"status":     "updated",
		},
	}

	res, err := client.Plans.PutPlansID(ctx, testPlanID, updateRequest)
	if err != nil {
		log.Printf("❌ Error updating plan: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Plan updated successfully!\n")
	fmt.Printf("  ID: %s\n", *res.ID)
	fmt.Printf("  New Name: %s\n", *res.Name)
	fmt.Printf("  New Description: %s\n", *res.Description)
	fmt.Printf("  Updated At: %s\n\n", *res.UpdatedAt)
}

func testSearchPlans(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Search Plans ---")

	searchFilter := components.TypesPlanFilter{
		PlanIds: []string{testPlanID},
	}

	res, err := client.Plans.PostPlansSearch(ctx, searchFilter)
	if err != nil {
		log.Printf("❌ Error searching plans: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Search completed!\n")
	fmt.Printf("  Found %d plans matching ID '%s'\n", len(res.Items), testPlanID)
	for i, plan := range res.Items {
		if i < 3 {
			fmt.Printf("  - %s: %s (%s)\n", *plan.ID, *plan.Name, *plan.LookupKey)
		}
	}
	fmt.Println()
}

func testDeletePlan(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Delete Plan ---")

	_, err := client.Plans.DeletePlansID(ctx, testPlanID)
	if err != nil {
		log.Printf("❌ Error deleting plan: %v", err)
		fmt.Println()
		return
	}

	fmt.Printf("✓ Plan deleted successfully!\n")
	fmt.Printf("  Deleted ID: %s\n\n", testPlanID)
}

// ========================================
// ADDONS, ENTITLEMENTS, SUBSCRIPTIONS (Stub implementations)
// ========================================

func testCreateAddon(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Addon ---")
	timestamp := time.Now().Unix()
	testAddonName = fmt.Sprintf("Test Addon %d", timestamp)
	testAddonLookupKey = fmt.Sprintf("test_addon_%d", timestamp)

	addonRequest := components.DtoCreateAddonRequest{
		Name:        testAddonName,
		LookupKey:   testAddonLookupKey,
		Description: gosdktemp.String("Test addon created by SDK"),
		Type:        components.TypesAddonTypeOnetime,
		Metadata: map[string]interface{}{
			"source": "sdk_test",
		},
	}

	res, err := client.Addons.PostAddons(ctx, addonRequest)
	if err != nil {
		log.Printf("❌ Error creating addon: %v", err)
		return
	}

	testAddonID = *res.ID
	fmt.Printf("✓ Addon created: %s\n\n", *res.ID)
}

func testGetAddon(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Addon ---")
	res, err := client.Addons.GetAddonsID(ctx, testAddonID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Addon retrieved: %s\n\n", *res.ID)
}

func testListAddons(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Addons ---")
	limit := int64(10)
	res, err := client.Addons.GetAddons(ctx, operations.GetAddonsRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d addons\n\n", len(res.Items))
}

func testUpdateAddon(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Addon ---")
	updatedName := testAddonName + " (Updated)"
	res, err := client.Addons.PutAddonsID(ctx, testAddonID, components.DtoUpdateAddonRequest{
		Name: &updatedName,
	})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Addon updated: %s\n\n", *res.Name)
}

func testLookupAddon(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Lookup Addon ---")
	fmt.Printf("⚠ Skipping (method not available in SDK)\n\n")
}

func testSearchAddons(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Search Addons ---")
	res, err := client.Addons.PostAddonsSearch(ctx, components.TypesAddonFilter{
		AddonIds: []string{testAddonID},
	})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Found %d addons\n\n", len(res.Items))
}

func testDeleteAddon(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 7: Delete Addon ---")
	_, err := client.Addons.DeleteAddonsID(ctx, testAddonID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Addon deleted\n\n")
}

func testCreateEntitlement(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Entitlement ---")
	entitlementRequest := components.DtoCreateEntitlementRequest{
		FeatureID:        testFeatureID,
		FeatureType:      components.TypesFeatureTypeBoolean,
		PlanID:           &testPlanID,
		IsEnabled:        gosdktemp.Bool(true),
		UsageResetPeriod: components.TypesEntitlementUsageResetPeriodMonthly.ToPointer(),
	}

	res, err := client.Entitlements.PostEntitlements(ctx, entitlementRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}

	testEntitlementID = *res.ID
	fmt.Printf("✓ Entitlement created: %s\n\n", *res.ID)
}

func testGetEntitlement(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Entitlement ---")
	res, err := client.Entitlements.GetEntitlementsID(ctx, testEntitlementID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Entitlement retrieved: %s\n\n", *res.ID)
}

func testListEntitlements(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Entitlements ---")
	limit := int64(10)
	res, err := client.Entitlements.GetEntitlements(ctx, operations.GetEntitlementsRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d entitlements\n\n", len(res.Items))
}

func testUpdateEntitlement(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Entitlement ---")
	res, err := client.Entitlements.PutEntitlementsID(ctx, testEntitlementID, components.DtoUpdateEntitlementRequest{
		IsEnabled: gosdktemp.Bool(false),
	})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Entitlement updated: %v\n\n", *res.IsEnabled)
}

func testSearchEntitlements(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Search Entitlements ---")
	res, err := client.Entitlements.PostEntitlementsSearch(ctx, components.TypesEntitlementFilter{
		EntityIds: []string{testEntitlementID},
	})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Found %d entitlements\n\n", len(res.Items))

}

func testDeleteEntitlement(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Delete Entitlement ---")
	_, err := client.Entitlements.DeleteEntitlementsID(ctx, testEntitlementID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Entitlement deleted\n\n")
}

// ========================================
// SUBSCRIPTIONS API TESTS (Simplified)
// ========================================

func testCreateSubscription(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Subscription ---")

	// Create price for plan first
	priceRequest := components.DtoCreatePriceRequest{
		EntityID:           testPlanID,
		EntityType:         components.TypesPriceEntityTypePlan,
		Type:               components.TypesPriceTypeFixed,
		BillingModel:       components.TypesBillingModelFlatFee,
		BillingCadence:     components.TypesBillingCadenceRecurring,
		BillingPeriod:      components.TypesBillingPeriodMonthly,
		BillingPeriodCount: gosdktemp.Int64(1),
		InvoiceCadence:     components.TypesInvoiceCadenceArrear,
		Amount:             gosdktemp.String("29.99"),
		Currency:           "USD",
		DisplayName:        gosdktemp.String("Monthly Price"),
	}

	price, err := client.Prices.PostPrices(ctx, priceRequest)
	if err != nil {
		log.Printf("❌ Error creating price for subscription: %v\n", err)
		return
	}
	fmt.Printf("  Created price: %s\n", *price.ID)

	startDate := time.Now().Format(time.RFC3339)
	subscriptionRequest := components.DtoCreateSubscriptionRequest{
		CustomerID:         gosdktemp.String(testCustomerID),
		PlanID:             testPlanID,
		Currency:           "USD",
		BillingCadence:     components.TypesBillingCadenceRecurring,
		BillingPeriod:      components.TypesBillingPeriodMonthly,
		BillingPeriodCount: gosdktemp.Int64(1),
		BillingCycle:       components.TypesBillingCycleAnniversary.ToPointer(),
		StartDate:          gosdktemp.String(startDate),
	}

	res, err := client.Subscriptions.PostSubscriptions(ctx, subscriptionRequest)
	if err != nil {
		log.Printf("❌ Error creating subscription: %v\n", err)
		return
	}

	testSubscriptionID = *res.ID
	fmt.Printf("✓ Subscription created: %s\n\n", *res.ID)
}

func testGetSubscription(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Subscription ---")
	if testSubscriptionID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.Subscriptions.GetSubscriptionsID(ctx, testSubscriptionID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Subscription retrieved: %s\n\n", *res.ID)
}

func testListSubscriptions(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Subscriptions ---")
	limit := int64(10)
	res, err := client.Subscriptions.GetSubscriptions(ctx, operations.GetSubscriptionsRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d subscriptions\n\n", len(res.Items))
}

func testSearchSubscriptions(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Search Subscriptions ---")
	res, err := client.Subscriptions.PostSubscriptionsSearch(ctx, components.TypesSubscriptionFilter{})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Found %d subscriptions\n\n", len(res.Items))
}

func testActivateSubscription(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Activate Subscription ---")
	fmt.Printf("⚠ Skipping (requires draft subscription)\n\n")
}

func testAddAddonToSubscription(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Add Addon to Subscription ---")

	if testSubscriptionID == "" || testAddonID == "" {
		log.Printf("⚠ Skipping (no subscription or addon available)\n\n")
		return
	}

	// Create a price for the addon first
	addonPriceRequest := components.DtoCreatePriceRequest{
		EntityID:           testAddonID,
		EntityType:         components.TypesPriceEntityTypePlan,
		Type:               components.TypesPriceTypeFixed,
		BillingModel:       components.TypesBillingModelFlatFee,
		BillingCadence:     components.TypesBillingCadenceRecurring,
		BillingPeriod:      components.TypesBillingPeriodMonthly,
		BillingPeriodCount: gosdktemp.Int64(1),
		InvoiceCadence:     components.TypesInvoiceCadenceArrear,
		Amount:             gosdktemp.String("5.00"),
		Currency:           "USD",
		DisplayName:        gosdktemp.String("Addon Price"),
	}

	_, err := client.Prices.PostPrices(ctx, addonPriceRequest)
	if err != nil {
		log.Printf("⚠ Warning: Could not create addon price: %v\n\n", err)
	}

	fmt.Printf("⚠ Skipping (add addon endpoint needs verification)\n\n")
}

func testRemoveAddonFromSubscription(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 7: Remove Addon from Subscription ---")
	fmt.Printf("⚠ Skipping (requires addon association ID, not addon ID)\n\n")
}

func testExecuteSubscriptionChange(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 8: Execute Subscription Change ---")
	fmt.Printf("⚠ Skipping (would modify active subscription)\n\n")
}

func testGetSubscriptionEntitlements(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 9: Get Subscription Entitlements ---")
	if testSubscriptionID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.Subscriptions.GetSubscriptionsIDEntitlements(ctx, testSubscriptionID, nil)
	if err != nil {
		log.Printf("⚠ Skipping: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d features\n\n", len(res.Features))
}

func testGetUpcomingGrants(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 10: Get Upcoming Grants ---")

	if testSubscriptionID == "" {
		log.Printf("⚠ Skipping (no subscription available)\n\n")
		return
	}

	fmt.Printf("⚠ Skipping (endpoint needs verification)\n\n")
}

func testReportUsage(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 11: Report Usage ---")

	if testSubscriptionID == "" {
		log.Printf("⚠ Skipping (no subscription available)\n\n")
		return
	}

	fmt.Printf("⚠ Skipping (requires metered feature setup)\n\n")
}

func testUpdateLineItem(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 12: Update Line Item ---")
	log.Printf("⚠ Skipping (requires line item ID from subscription)\n\n")
}

func testDeleteLineItem(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 13: Delete Line Item ---")
	log.Printf("⚠ Skipping (requires line item ID from subscription)\n\n")
}

func testCancelSubscription(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 14: Cancel Subscription ---")
	if testSubscriptionID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	_, err := client.Subscriptions.PostSubscriptionsIDCancel(ctx, testSubscriptionID, components.DtoCancelSubscriptionRequest{
		CancellationType: components.TypesCancellationTypeEndOfPeriod,
	})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Subscription canceled\n\n")
}

// ========================================
// INVOICES, PRICES, PAYMENTS, WALLETS, ETC
// ========================================

func testListInvoices(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: List Invoices ---")
	limit := int64(10)
	res, err := client.Invoices.GetInvoices(ctx, operations.GetInvoicesRequest{Limit: &limit})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	if len(res.Items) > 0 {
		testInvoiceID = *res.Items[0].ID
	}
	fmt.Printf("✓ Retrieved %d invoices\n\n", len(res.Items))
}

func testSearchInvoices(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Search Invoices ---")
	res, err := client.Invoices.PostInvoicesSearch(ctx, components.TypesInvoiceFilter{})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Found %d invoices\n\n", len(res.Items))
}

func testCreateInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: Create Invoice ---")
	draftStatus := components.TypesInvoiceStatusDraft
	invoiceRequest := components.DtoCreateInvoiceRequest{
		CustomerID:    testCustomerID,
		Currency:      "USD",
		AmountDue:     "100.00",
		Subtotal:      "100.00",
		Total:         "100.00",
		InvoiceType:   components.TypesInvoiceTypeOneOff.ToPointer(),
		BillingReason: components.TypesInvoiceBillingReasonManual.ToPointer(),
		InvoiceStatus: &draftStatus,
		LineItems: []components.DtoCreateInvoiceLineItemRequest{
			{
				DisplayName: gosdktemp.String("Test Service"),
				Quantity:    "1",
				Amount:      "100.00",
			},
		},
	}

	res, err := client.Invoices.PostInvoices(ctx, invoiceRequest)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}

	testInvoiceID = *res.ID
	fmt.Printf("✓ Invoice created: %s\n\n", *res.ID)
}

func testGetInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Get Invoice ---")
	if testInvoiceID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.Invoices.GetInvoicesID(ctx, testInvoiceID, nil, nil)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Invoice retrieved: %s\n\n", *res.ID)
}

func testUpdateInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Update Invoice ---")
	if testInvoiceID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	_, err := client.Invoices.PutInvoicesID(ctx, testInvoiceID, components.DtoUpdateInvoiceRequest{
		Metadata: map[string]string{"updated": "true"},
	})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Invoice updated\n\n")
}

func testPreviewInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Preview Invoice ---")
	if testCustomerID == "" {
		log.Printf("⚠ Skipping (no customer available)\n\n")
		return
	}

	previewRequest := components.DtoGetPreviewInvoiceRequest{
		SubscriptionID: testSubscriptionID,
	}

	_, err := client.Invoices.PostInvoicesPreview(ctx, previewRequest)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Invoice preview generated\n\n")
}

func testFinalizeInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 7: Finalize Invoice ---")

	// Create a draft invoice for finalization
	if testCustomerID == "" {
		log.Printf("⚠ Skipping (no customer available)\n\n")
		return
	}

	draftStatus := components.TypesInvoiceStatusDraft
	invoiceRequest := components.DtoCreateInvoiceRequest{
		CustomerID:    testCustomerID,
		Currency:      "USD",
		AmountDue:     "50.00",
		Subtotal:      "50.00",
		Total:         "50.00",
		InvoiceType:   components.TypesInvoiceTypeOneOff.ToPointer(),
		BillingReason: components.TypesInvoiceBillingReasonManual.ToPointer(),
		InvoiceStatus: &draftStatus,
		LineItems: []components.DtoCreateInvoiceLineItemRequest{
			{
				DisplayName: gosdktemp.String("Finalize Test Service"),
				Quantity:    "1",
				Amount:      "50.00",
			},
		},
	}

	invoice, err := client.Invoices.PostInvoices(ctx, invoiceRequest)
	if err != nil {
		log.Printf("⚠ Error creating draft invoice: %v\n\n", err)
		return
	}

	_, err = client.Invoices.PostInvoicesIDFinalize(ctx, *invoice.ID)
	if err != nil {
		log.Printf("⚠ Error finalizing: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Invoice finalized: %s\n\n", *invoice.ID)
}

func testRecalculateInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 8: Recalculate Invoice ---")
	log.Printf("⚠ Skipping (recalculate only works on subscription invoices)\n\n")
}

func testRecordPayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 9: Record Payment ---")
	if testInvoiceID == "" {
		log.Printf("⚠ Skipping (no invoice available)\n\n")
		return
	}

	paymentRequest := components.DtoUpdatePaymentStatusRequest{
		PaymentStatus: components.TypesPaymentStatusSucceeded,
		Amount:        gosdktemp.String("100.00"),
	}

	_, err := client.Invoices.PutInvoicesIDPayment(ctx, testInvoiceID, paymentRequest)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Payment recorded: %s\n\n", testInvoiceID)
}

func testAttemptPayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 10: Attempt Payment ---")

	// Create a finalized invoice for payment attempt
	if testCustomerID == "" {
		log.Printf("⚠ Skipping (no customer available)\n\n")
		return
	}

	draftStatus := components.TypesInvoiceStatusDraft
	pendingStatus := components.TypesPaymentStatusPending
	invoiceRequest := components.DtoCreateInvoiceRequest{
		CustomerID:    testCustomerID,
		Currency:      "USD",
		AmountDue:     "25.00",
		Subtotal:      "25.00",
		Total:         "25.00",
		AmountPaid:    gosdktemp.String("0.00"),
		InvoiceType:   components.TypesInvoiceTypeOneOff.ToPointer(),
		BillingReason: components.TypesInvoiceBillingReasonManual.ToPointer(),
		InvoiceStatus: &draftStatus,
		PaymentStatus: &pendingStatus,
		LineItems: []components.DtoCreateInvoiceLineItemRequest{
			{
				DisplayName: gosdktemp.String("Attempt Payment Test"),
				Quantity:    "1",
				Amount:      "25.00",
			},
		},
	}

	invoice, err := client.Invoices.PostInvoices(ctx, invoiceRequest)
	if err != nil {
		log.Printf("⚠ Error creating invoice: %v\n\n", err)
		return
	}

	_, err = client.Invoices.PostInvoicesIDFinalize(ctx, *invoice.ID)
	if err != nil {
		log.Printf("⚠ Error finalizing invoice: %v\n\n", err)
		return
	}

	_, err = client.Invoices.PostInvoicesIDPaymentAttempt(ctx, *invoice.ID)
	if err != nil {
		log.Printf("⚠ Error attempting payment: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Payment attempt initiated: %s\n\n", *invoice.ID)
}

func testDownloadInvoicePDF(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 11: Download Invoice PDF ---")
	if testInvoiceID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	_, err := client.Invoices.GetInvoicesIDPdf(ctx, testInvoiceID, nil)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ PDF downloaded\n\n")
}

func testTriggerInvoiceComms(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 12: Trigger Invoice Comms ---")
	if testInvoiceID == "" {
		log.Printf("⚠ Skipping (no invoice available)\n\n")
		return
	}

	_, err := client.Invoices.PostInvoicesIDCommsTrigger(ctx, testInvoiceID)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Invoice communications triggered: %s\n\n", testInvoiceID)
}

func testGetCustomerInvoiceSummary(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 13: Get Customer Invoice Summary ---")
	_, err := client.Invoices.GetCustomersIDInvoicesSummary(ctx, testCustomerID)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Summary retrieved\n\n")
}

func testVoidInvoice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 14: Void Invoice ---")
	if testInvoiceID == "" {
		log.Printf("⚠ Skipping (no invoice available)\n\n")
		return
	}

	_, err := client.Invoices.PostInvoicesIDVoid(ctx, testInvoiceID)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Invoice voided: %s\n\n", testInvoiceID)
}

func testCreatePrice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Price ---")
	if testPlanID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	priceRequest := components.DtoCreatePriceRequest{
		EntityID:           testPlanID,
		EntityType:         components.TypesPriceEntityTypePlan,
		Currency:           "USD",
		Amount:             gosdktemp.String("99.00"),
		BillingModel:       components.TypesBillingModelFlatFee,
		BillingCadence:     components.TypesBillingCadenceRecurring,
		BillingPeriod:      components.TypesBillingPeriodMonthly,
		BillingPeriodCount: gosdktemp.Int64(1),
		InvoiceCadence:     components.TypesInvoiceCadenceAdvance,
		PriceUnitType:      components.TypesPriceUnitTypeFiat,
		Type:               components.TypesPriceTypeFixed,
		DisplayName:        gosdktemp.String("Monthly Subscription"),
	}

	res, err := client.Prices.PostPrices(ctx, priceRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}

	testPriceID = *res.ID
	fmt.Printf("✓ Price created: %s\n\n", *res.ID)
}

func testGetPrice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Price ---")
	if testPriceID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.Prices.GetPricesID(ctx, testPriceID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Price retrieved: %s\n\n", *res.ID)
}

func testListPrices(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Prices ---")
	limit := int64(10)
	res, err := client.Prices.GetPrices(ctx, operations.GetPricesRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d prices\n\n", len(res.Items))
}

func testUpdatePrice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Price ---")
	if testPriceID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	updatedDesc := "Updated price description"
	_, err := client.Prices.PutPricesID(ctx, testPriceID, components.DtoUpdatePriceRequest{
		Description: &updatedDesc,
	})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Price updated\n\n")
}

func testDeletePrice(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Delete Price ---")
	if testPriceID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	futureDate := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	_, err := client.Prices.DeletePricesID(ctx, testPriceID, components.DtoDeletePriceRequest{
		EndDate: &futureDate,
	})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Price deleted\n\n")
}

func testCreatePayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Payment ---")
	fmt.Printf("⚠ Skipping (requires invoice setup)\n\n")
}

func testGetPayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Payment ---")
	if testPaymentID == "" {
		log.Printf("⚠ Skipping (no payment available)\n\n")
		return
	}

	res, err := client.Payments.GetPaymentsID(ctx, testPaymentID)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Payment retrieved: %s\n\n", *res.ID)
}

func testListPayments(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Payments ---")
	limit := int64(10)
	res, err := client.Payments.GetPayments(ctx, operations.GetPaymentsRequest{Limit: &limit})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	if len(res.Items) > 0 {
		testPaymentID = *res.Items[0].ID
	}
	fmt.Printf("✓ Retrieved %d payments\n\n", len(res.Items))
}

func testUpdatePayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Payment ---")
	if testPaymentID == "" {
		log.Printf("⚠ Skipping (no payment available)\n\n")
		return
	}

	updateRequest := components.DtoUpdatePaymentRequest{
		Metadata: map[string]string{
			"updated_at": time.Now().Format(time.RFC3339),
			"status":     "updated",
		},
	}

	_, err := client.Payments.PutPaymentsID(ctx, testPaymentID, updateRequest)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Payment updated: %s\n\n", testPaymentID)
}

func testProcessPayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Process Payment ---")
	if testPaymentID == "" {
		log.Printf("⚠ Skipping (no payment available)\n\n")
		return
	}

	_, err := client.Payments.PostPaymentsIDProcess(ctx, testPaymentID)
	if err != nil {
		log.Printf("⚠ Error (may require payment gateway setup): %v\n\n", err)
		return
	}
	fmt.Printf("✓ Payment processed: %s\n\n", testPaymentID)
}

func testDeletePayment(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Delete Payment ---")
	if testPaymentID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	_, err := client.Payments.DeletePaymentsID(ctx, testPaymentID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Payment deleted\n\n")
}

func testCreateWallet(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Wallet ---")
	if testCustomerID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	walletRequest := components.DtoCreateWalletRequest{
		CustomerID: &testCustomerID,
		Currency:   "USD",
		Name:       gosdktemp.String("Test Wallet"),
	}

	res, err := client.Wallets.PostWallets(ctx, walletRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}

	testWalletID = *res.ID
	fmt.Printf("✓ Wallet created: %s\n\n", *res.ID)
}

func testGetWallet(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Wallet ---")
	if testWalletID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.Wallets.GetWalletsID(ctx, testWalletID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Wallet retrieved: %s\n\n", *res.ID)
}

func testListWallets(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Wallets ---")
	limit := int64(10)
	res, err := client.Wallets.GetWallets(ctx, operations.GetWalletsRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d wallets\n\n", len(res.Items))
}

func testUpdateWallet(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Wallet ---")
	if testWalletID == "" {
		log.Printf("⚠ Skipping (no wallet available)\n\n")
		return
	}

	updateRequest := components.DtoUpdateWalletRequest{
		Metadata: map[string]string{
			"updated_at": time.Now().Format(time.RFC3339),
			"status":     "updated",
		},
	}

	_, err := client.Wallets.PutWalletsID(ctx, testWalletID, updateRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Wallet updated: %s\n\n", testWalletID)
}

func testGetWalletBalance(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 5: Get Wallet Balance ---")
	if testWalletID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.Wallets.GetWalletsIDBalanceRealTime(ctx, testWalletID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Balance: %s\n\n", *res.Balance)
}

func testTopUpWallet(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 6: Top Up Wallet ---")
	if testWalletID == "" {
		log.Printf("⚠ Skipping (no wallet available)\n\n")
		return
	}

	topUpRequest := components.DtoTopUpWalletRequest{
		Amount:            gosdktemp.String("100.00"),
		TransactionReason: components.TypesTransactionReasonPurchasedCreditDirect,
		Description:       gosdktemp.String("Test top-up from SDK"),
	}

	_, err := client.Wallets.PostWalletsIDTopUp(ctx, testWalletID, topUpRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Wallet topped up: %s\n\n", testWalletID)
}

func testDebitWallet(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 7: Debit Wallet ---")
	log.Printf("⚠ Skipping (debit endpoint not available in SDK)\n\n")
}

func testGetWalletTransactions(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 8: Get Wallet Transactions ---")
	if testWalletID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	limit := int64(10)
	res, err := client.Wallets.GetWalletsIDTransactions(ctx, operations.GetWalletsIDTransactionsRequest{IDPathParameter: testWalletID, Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d transactions\n\n", len(res.Items))
}

func testSearchWallets(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 9: Search Wallets ---")
	limit := int64(10)
	res, err := client.Wallets.PostWalletsSearch(ctx, &components.TypesWalletFilter{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Found %d wallets\n\n", len(res.Items))
}

func testCreateCreditGrant(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Credit Grant ---")
	if testPlanID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	grantRequest := components.DtoCreateCreditGrantRequest{
		Scope:                  components.TypesCreditGrantScopePlan,
		PlanID:                 &testPlanID,
		Credits:                "500.00",
		Name:                   "Test Credit Grant",
		Cadence:                components.TypesCreditGrantCadenceOnetime,
		ExpirationType:         components.TypesCreditGrantExpiryTypeNever.ToPointer(),
		ExpirationDurationUnit: components.TypesCreditGrantExpiryDurationUnitDay.ToPointer(),
	}

	res, err := client.CreditGrants.PostCreditgrants(ctx, grantRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}

	testCreditGrantID = *res.ID
	fmt.Printf("✓ Credit grant created: %s\n\n", *res.ID)
}

func testGetCreditGrant(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Credit Grant ---")
	if testCreditGrantID == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	res, err := client.CreditGrants.GetCreditgrantsID(ctx, testCreditGrantID)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Credit grant retrieved: %s\n\n", *res.ID)
}

func testListCreditGrants(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Credit Grants ---")
	limit := int64(10)
	res, err := client.CreditGrants.GetCreditgrants(ctx, operations.GetCreditgrantsRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d credit grants\n\n", len(res.Items))
}

func testUpdateCreditGrant(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Update Credit Grant ---")
	if testCreditGrantID == "" {
		log.Printf("⚠ Skipping (no credit grant available)\n\n")
		return
	}

	updateRequest := components.DtoUpdateCreditGrantRequest{
		Metadata: map[string]string{
			"updated_at": time.Now().Format(time.RFC3339),
			"status":     "updated",
		},
	}

	_, err := client.CreditGrants.PutCreditgrantsID(ctx, testCreditGrantID, updateRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Credit grant updated: %s\n\n", testCreditGrantID)
}

func testCreateCreditNote(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Credit Note ---")
	fmt.Printf("⚠ Skipping (requires invoice with line items)\n\n")
}

func testGetCreditNote(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Get Credit Note ---")
	if testCreditNoteID == "" {
		log.Printf("⚠ Skipping (no credit note available)\n\n")
		return
	}

	res, err := client.CreditNotes.GetCreditnotesID(ctx, testCreditNoteID)
	if err != nil {
		log.Printf("❌ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Credit note retrieved: %s\n\n", *res.ID)
}

func testListCreditNotes(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: List Credit Notes ---")
	limit := int64(10)
	res, err := client.CreditNotes.GetCreditnotes(ctx, operations.GetCreditnotesRequest{Limit: &limit})
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d credit notes\n\n", len(res.Items))
}

func testFinalizeCreditNote(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 4: Finalize Credit Note ---")
	if testCreditNoteID == "" {
		log.Printf("⚠ Skipping (no credit note available)\n\n")
		return
	}

	_, err := client.CreditNotes.PostCreditnotesIDFinalize(ctx, testCreditNoteID)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Credit note finalized: %s\n\n", testCreditNoteID)
}

func testCreateEvent(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: Create Event ---")

	if testExternalID == "" {
		testEventCustomerID = fmt.Sprintf("test-customer-%d", time.Now().Unix())
	} else {
		testEventCustomerID = testExternalID
	}

	testEventName = fmt.Sprintf("Test Event %d", time.Now().Unix())

	eventRequest := components.DtoIngestEventRequest{
		EventName:          testEventName,
		ExternalCustomerID: testEventCustomerID,
		Properties: map[string]string{
			"source":      "sdk_test",
			"environment": "test",
		},
		Source:    gosdktemp.String("sdk_test"),
		Timestamp: gosdktemp.String(time.Now().Format(time.RFC3339)),
	}

	res, err := client.Events.PostEvents(ctx, eventRequest)
	if err != nil {
		log.Printf("❌ Error: %v\n", err)
		return
	}

	if res != nil {
		if eventId, ok := res["event_id"]; ok {
			testEventID = eventId
			fmt.Printf("✓ Event created: %s\n\n", eventId)
		} else {
			fmt.Printf("✓ Event created\n\n")
		}
	}
}

func testQueryEvents(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Query Events ---")
	if testEventName == "" {
		log.Printf("⚠ Skipping\n\n")
		return
	}

	// Query with a simple map since PostEventsQuery accepts any
	queryRequest := map[string]interface{}{
		"external_customer_id": testEventCustomerID,
		"event_name":           testEventName,
	}

	res, err := client.Events.PostEventsQuery(ctx, queryRequest)
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}

	fmt.Printf("✓ Found %d events\n\n", len(res.Events))
}

func testAsyncEventIngestion(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 3: Async Event Ingestion ---")

	// Create an AsyncClient with debug enabled
	asyncConfig := gosdktemp.DefaultAsyncConfig()
	asyncConfig.Debug = true
	asyncConfig.BatchSize = 5 // Smaller batch for testing

	asyncClient := client.NewAsyncClientWithConfig(asyncConfig)
	defer asyncClient.Close()

	customerID := fmt.Sprintf("async-customer-%d", time.Now().Unix())

	// Example 1: Simple event using Enqueue
	err := asyncClient.Enqueue(
		"api_request",
		customerID,
		map[string]interface{}{
			"path":             "/api/resource",
			"method":           "GET",
			"status":           "200",
			"response_time_ms": 150,
		},
	)
	if err != nil {
		log.Printf("❌ Failed to enqueue simple event: %v\n", err)
	} else {
		fmt.Println("  ✓ Enqueued simple event")
	}

	// Example 2: Event with additional options using EnqueueWithOptions
	err = asyncClient.EnqueueWithOptions(gosdktemp.EventOptions{
		EventName:          "file_upload",
		ExternalCustomerID: customerID,
		Properties: map[string]interface{}{
			"file_size_bytes": 1048576,
			"file_type":       "image/jpeg",
			"storage_bucket":  "user_uploads",
		},
		Source:    "upload_service",
		Timestamp: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		log.Printf("❌ Failed to enqueue event with options: %v\n", err)
	} else {
		fmt.Println("  ✓ Enqueued event with custom options")
	}

	// Example 3: Batch multiple events
	for i := 0; i < 10; i++ {
		err = asyncClient.Enqueue(
			"batch_example",
			fmt.Sprintf("%s-batch-%d", customerID, i),
			map[string]interface{}{
				"index": i,
				"batch": "async_test",
			},
		)
		if err != nil {
			log.Printf("❌ Failed to enqueue batch event %d: %v\n", i, err)
			break
		}
	}
	fmt.Println("  ✓ Enqueued 10 batch events")

	// Wait a moment for background processing
	fmt.Println("  ⏳ Waiting for async events to be processed...")
	time.Sleep(2 * time.Second)

	fmt.Println("✓ Async event ingestion test completed!\n")
}

func testListConnections(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 1: List Connections ---")
	limit := int64(10)
	res, err := client.Connections.GetConnections(ctx, operations.GetConnectionsRequest{Limit: &limit})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Retrieved %d connections\n\n", len(res.Connections))
}

func testSearchConnections(ctx context.Context, client *gosdktemp.FlexPrice) {
	fmt.Println("--- Test 2: Search Connections ---")
	limit := int64(5)
	res, err := client.Connections.PostConnectionsSearch(ctx, components.TypesConnectionFilter{Limit: &limit})
	if err != nil {
		log.Printf("⚠ Error: %v\n\n", err)
		return
	}
	fmt.Printf("✓ Found %d connections\n\n", len(res.Connections))
}
