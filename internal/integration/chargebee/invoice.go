package chargebee

import (
	"context"
	"fmt"
	"time"

	"github.com/chargebee/chargebee-go/v3/enum"
	chargebeeInvoice "github.com/chargebee/chargebee-go/v3/models/invoice"
	"github.com/flexprice/flexprice/internal/api/dto"
	"github.com/flexprice/flexprice/internal/domain/entityintegrationmapping"
	"github.com/flexprice/flexprice/internal/domain/invoice"
	"github.com/flexprice/flexprice/internal/domain/payment"
	"github.com/flexprice/flexprice/internal/domain/price"
	ierr "github.com/flexprice/flexprice/internal/errors"
	"github.com/flexprice/flexprice/internal/logger"
	"github.com/flexprice/flexprice/internal/types"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// ChargebeeInvoiceService defines the interface for Chargebee invoice operations
type ChargebeeInvoiceService interface {
	CreateInvoice(ctx context.Context, req *InvoiceCreateRequest) (*InvoiceResponse, error)
	RetrieveInvoice(ctx context.Context, invoiceID string) (*InvoiceResponse, error)
	SyncInvoiceToChargebee(ctx context.Context, req ChargebeeInvoiceSyncRequest) (*ChargebeeInvoiceSyncResponse, error)
}

// InvoiceServiceParams holds dependencies for InvoiceService
type InvoiceServiceParams struct {
	Client                       ChargebeeClient
	CustomerSvc                  ChargebeeCustomerService
	InvoiceRepo                  invoice.Repository
	PaymentRepo                  payment.Repository
	EntityIntegrationMappingRepo entityintegrationmapping.Repository
	PriceRepo                    price.Repository
	Logger                       *logger.Logger
}

// InvoiceService handles Chargebee invoice operations
type InvoiceService struct {
	InvoiceServiceParams
	itemFamilyService ChargebeeItemFamilyService
	itemService       ChargebeeItemService
	itemPriceService  ChargebeeItemPriceService
}

// NewInvoiceService creates a new Chargebee invoice service
func NewInvoiceService(params InvoiceServiceParams) ChargebeeInvoiceService {
	// Initialize dependent services for item and item price creation
	itemFamilyService := NewItemFamilyService(ItemFamilyServiceParams{
		Client: params.Client,
		Logger: params.Logger,
	})
	itemService := NewItemService(ItemServiceParams{
		Client: params.Client,
		Logger: params.Logger,
	})
	itemPriceService := NewItemPriceService(ItemPriceServiceParams{
		Client: params.Client,
		Logger: params.Logger,
	})

	return &InvoiceService{
		InvoiceServiceParams: params,
		itemFamilyService:    itemFamilyService,
		itemService:          itemService,
		itemPriceService:     itemPriceService,
	}
}

// ChargebeeInvoiceSyncRequest represents a request to sync an invoice to Chargebee
type ChargebeeInvoiceSyncRequest struct {
	InvoiceID string `json:"invoice_id" validate:"required"`
}

// ChargebeeInvoiceSyncResponse represents the response from syncing an invoice to Chargebee
type ChargebeeInvoiceSyncResponse struct {
	ChargebeeInvoiceID string `json:"chargebee_invoice_id"`
	Status             string `json:"status"`
	Total              int64  `json:"total"`
	AmountDue          int64  `json:"amount_due"`
	CurrencyCode       string `json:"currency_code"`
}

// CreateInvoice creates a new invoice in Chargebee using charge items
func (s *InvoiceService) CreateInvoice(ctx context.Context, req *InvoiceCreateRequest) (*InvoiceResponse, error) {
	s.Logger.Infow("creating invoice in Chargebee",
		"customer_id", req.CustomerID,
		"auto_collection", req.AutoCollection,
		"line_items_count", len(req.LineItems))

	// Prepare request params
	createParams := &chargebeeInvoice.CreateForChargeItemsAndChargesRequestParams{
		CustomerId:     req.CustomerID,
		AutoCollection: enum.AutoCollection(req.AutoCollection),
	}

	// Add item prices (line items)
	if len(req.LineItems) > 0 {
		itemPrices := make([]*chargebeeInvoice.CreateForChargeItemsAndChargesItemPriceParams, 0, len(req.LineItems))
		for _, item := range req.LineItems {
			itemPrice := &chargebeeInvoice.CreateForChargeItemsAndChargesItemPriceParams{
				ItemPriceId: item.ItemPriceID,
				Quantity:    lo.ToPtr(int32(item.Quantity)),
			}

			if item.UnitAmount > 0 {
				itemPrice.UnitPrice = lo.ToPtr(item.UnitAmount)
			}

			if item.DateFrom != nil {
				itemPrice.DateFrom = lo.ToPtr(item.DateFrom.Unix())
			}

			if item.DateTo != nil {
				itemPrice.DateTo = lo.ToPtr(item.DateTo.Unix())
			}

			itemPrices = append(itemPrices, itemPrice)
		}
		createParams.ItemPrices = itemPrices
	}

	// Add optional dates
	if req.Date != nil {
		createParams.InvoiceDate = lo.ToPtr(req.Date.Unix())
	}

	// Create invoice using client wrapper
	result, err := s.Client.CreateInvoice(ctx, createParams)
	if err != nil {
		s.Logger.Errorw("failed to create invoice in Chargebee",
			"customer_id", req.CustomerID,
			"error", err)
		return nil, ierr.NewError("failed to create invoice in Chargebee").
			WithReportableDetails(map[string]interface{}{
				"error":       err.Error(),
				"customer_id": req.CustomerID,
			}).
			WithHint("Check Chargebee API credentials and invoice data").
			Mark(ierr.ErrValidation)
	}

	invoiceData := result.Invoice

	s.Logger.Infow("successfully created invoice in Chargebee",
		"invoice_id", invoiceData.Id,
		"customer_id", invoiceData.CustomerId,
		"status", invoiceData.Status,
		"total", invoiceData.Total)

	// Convert to our DTO format
	invoiceResponse := &InvoiceResponse{
		ID:              invoiceData.Id,
		CustomerID:      invoiceData.CustomerId,
		Status:          string(invoiceData.Status),
		AutoCollection:  req.AutoCollection, // Use the request value since API doesn't return it
		Total:           invoiceData.Total,
		AmountDue:       invoiceData.AmountDue,
		AmountPaid:      invoiceData.AmountPaid,
		CurrencyCode:    invoiceData.CurrencyCode,
		Date:            time.Unix(invoiceData.Date, 0),
		CreatedAt:       time.Unix(invoiceData.GeneratedAt, 0), // Use GeneratedAt as CreatedAt
		UpdatedAt:       time.Unix(invoiceData.UpdatedAt, 0),
		ResourceVersion: invoiceData.ResourceVersion,
	}

	if invoiceData.DueDate > 0 {
		dueDate := time.Unix(invoiceData.DueDate, 0)
		invoiceResponse.DueDate = &dueDate
	}

	// Convert line items if present
	if len(invoiceData.LineItems) > 0 {
		lineItems := make([]InvoiceLineItemResponse, 0, len(invoiceData.LineItems))
		for _, item := range invoiceData.LineItems {
			lineItem := InvoiceLineItemResponse{
				ID:          item.Id,
				ItemPriceID: item.EntityId, // EntityId contains the item price ID
				EntityType:  string(item.EntityType),
				Quantity:    int(item.Quantity),
				UnitAmount:  item.UnitAmount,
				Amount:      item.Amount,
				Description: item.Description,
				DateFrom:    time.Unix(item.DateFrom, 0),
				DateTo:      time.Unix(item.DateTo, 0),
			}
			lineItems = append(lineItems, lineItem)
		}
		invoiceResponse.LineItems = lineItems
	}

	return invoiceResponse, nil
}

// SyncInvoiceToChargebee syncs a FlexPrice invoice to Chargebee
func (s *InvoiceService) SyncInvoiceToChargebee(
	ctx context.Context,
	req ChargebeeInvoiceSyncRequest,
) (*ChargebeeInvoiceSyncResponse, error) {
	s.Logger.Infow("starting Chargebee invoice sync",
		"invoice_id", req.InvoiceID)

	// Step 1: Check if Chargebee connection exists
	if !s.Client.HasChargebeeConnection(ctx) {
		return nil, ierr.NewError("Chargebee connection not available").
			WithHint("Chargebee integration must be configured for invoice sync").
			Mark(ierr.ErrNotFound)
	}

	// Step 2: Get FlexPrice invoice
	flexInvoice, err := s.InvoiceRepo.Get(ctx, req.InvoiceID)
	if err != nil {
		return nil, ierr.WithError(err).
			WithHint("Failed to get FlexPrice invoice").
			Mark(ierr.ErrDatabase)
	}

	// Step 3: Check if invoice is already synced to avoid duplicates
	existingMapping, err := s.getExistingChargebeeMapping(ctx, req.InvoiceID)
	if err != nil && !ierr.IsNotFound(err) {
		return nil, err
	}

	var chargebeeInvoiceID string
	if existingMapping != nil {
		chargebeeInvoiceID = existingMapping.ProviderEntityID
		s.Logger.Infow("invoice already synced to Chargebee",
			"invoice_id", req.InvoiceID,
			"chargebee_invoice_id", chargebeeInvoiceID)
		// Fetch existing invoice details and return
		invoiceResp, err := s.RetrieveInvoice(ctx, chargebeeInvoiceID)
		if err != nil {
			return nil, err
		}
		return &ChargebeeInvoiceSyncResponse{
			ChargebeeInvoiceID: invoiceResp.ID,
			Status:             invoiceResp.Status,
			Total:              invoiceResp.Total,
			AmountDue:          invoiceResp.AmountDue,
			CurrencyCode:       invoiceResp.CurrencyCode,
		}, nil
	}

	// Step 4: Ensure customer is synced to Chargebee
	flexpriceCustomer, err := s.CustomerSvc.EnsureCustomerSyncedToChargebee(ctx, flexInvoice.CustomerID)
	if err != nil {
		return nil, ierr.WithError(err).
			WithHint("Failed to sync customer to Chargebee").
			Mark(ierr.ErrInternal)
	}

	chargebeeCustomerID := flexpriceCustomer.Metadata["chargebee_customer_id"]
	if chargebeeCustomerID == "" {
		// Try to get from entity mapping if not in metadata
		chargebeeCustomerID, err = s.CustomerSvc.GetChargebeeCustomerID(ctx, flexInvoice.CustomerID)
		if err != nil {
			return nil, ierr.WithError(err).
				WithHint("Failed to get Chargebee customer ID").
				Mark(ierr.ErrInternal)
		}
	}

	s.Logger.Infow("customer synced to Chargebee",
		"customer_id", flexInvoice.CustomerID,
		"chargebee_customer_id", chargebeeCustomerID)

	// Step 5: Check if customer has payment method and set auto_collection accordingly
	hasPaymentMethod, err := s.customerHasPaymentMethod(ctx, chargebeeCustomerID)
	if err != nil {
		s.Logger.Warnw("failed to check customer payment method, defaulting to auto_collection off",
			"error", err,
			"chargebee_customer_id", chargebeeCustomerID)
		hasPaymentMethod = false
	}

	autoCollection := "off"
	if hasPaymentMethod {
		autoCollection = "on"
	}

	s.Logger.Infow("determined auto_collection setting",
		"chargebee_customer_id", chargebeeCustomerID,
		"has_payment_method", hasPaymentMethod,
		"auto_collection", autoCollection)

	// Step 6: Build line items with Chargebee item price IDs
	lineItems, err := s.buildLineItems(ctx, flexInvoice)
	if err != nil {
		return nil, err
	}

	if len(lineItems) == 0 {
		return nil, ierr.NewError("invoice has no line items").
			WithHint("Cannot create Chargebee invoice without line items").
			Mark(ierr.ErrValidation)
	}

	// Step 7: Create invoice in Chargebee
	invoiceReq := &InvoiceCreateRequest{
		CustomerID:     chargebeeCustomerID,
		AutoCollection: autoCollection,
		LineItems:      lineItems,
	}

	// Set invoice date if available (use FinalizedAt or CreatedAt)
	if flexInvoice.FinalizedAt != nil {
		invoiceReq.Date = flexInvoice.FinalizedAt
	} else {
		// Use CreatedAt as fallback
		createdAt := flexInvoice.CreatedAt
		invoiceReq.Date = &createdAt
	}

	chargebeeInvoice, err := s.CreateInvoice(ctx, invoiceReq)
	if err != nil {
		return nil, ierr.WithError(err).
			WithHint("Failed to create invoice in Chargebee").
			Mark(ierr.ErrInternal)
	}

	chargebeeInvoiceID = chargebeeInvoice.ID
	s.Logger.Infow("successfully created invoice in Chargebee",
		"invoice_id", req.InvoiceID,
		"chargebee_invoice_id", chargebeeInvoiceID)

	// Step 8: Create entity integration mapping
	if err := s.createInvoiceMapping(ctx, req.InvoiceID, chargebeeInvoiceID, flexInvoice.EnvironmentID); err != nil {
		s.Logger.Errorw("failed to create invoice mapping",
			"error", err,
			"invoice_id", req.InvoiceID,
			"chargebee_invoice_id", chargebeeInvoiceID)
		// Don't fail the sync, just log the error
	}

	// Step 9: Build and return response
	return &ChargebeeInvoiceSyncResponse{
		ChargebeeInvoiceID: chargebeeInvoice.ID,
		Status:             chargebeeInvoice.Status,
		Total:              chargebeeInvoice.Total,
		AmountDue:          chargebeeInvoice.AmountDue,
		CurrencyCode:       chargebeeInvoice.CurrencyCode,
	}, nil
}

// buildLineItems converts FlexPrice line items to Chargebee format
// Maps FlexPrice price IDs to Chargebee item price IDs using entity mapping
func (s *InvoiceService) buildLineItems(ctx context.Context, flexInvoice *invoice.Invoice) ([]InvoiceLineItem, error) {
	lineItems := make([]InvoiceLineItem, 0)

	for _, item := range flexInvoice.LineItems {
		// Skip zero-amount items
		if item.Amount.IsZero() {
			s.Logger.Debugw("skipping zero-amount line item",
				"invoice_id", flexInvoice.ID,
				"line_item_id", item.ID)
			continue
		}

		// Skip if price ID is not available
		if item.PriceID == nil || *item.PriceID == "" {
			s.Logger.Debugw("line item has no price ID, skipping",
				"invoice_id", flexInvoice.ID,
				"line_item_id", item.ID)
			continue
		}

		// Get Chargebee item price ID AND check if it's tiered
		chargebeeItemPriceID, isTiered, err := s.getChargebeeItemPriceIDAndCheckTiered(ctx, *item.PriceID)
		if err != nil {
			s.Logger.Errorw("failed to get Chargebee item price ID for line item",
				"invoice_id", flexInvoice.ID,
				"line_item_id", item.ID,
				"price_id", *item.PriceID,
				"error", err)
			// Skip this line item if we can't find the mapping
			continue
		}

		lineItem := InvoiceLineItem{
			ItemPriceID: chargebeeItemPriceID,
		}

		// CRITICAL: Different handling based on pricing model
		if isTiered {
			// For tiered/volume/stairstep pricing: Send ONLY quantity, let Chargebee calculate
			// NOTE: Chargebee's calculation may differ slightly due to tier precision rounding
			lineItem.Quantity = int(item.Quantity.IntPart())
			// DO NOT set UnitAmount - Chargebee will reject it
			s.Logger.Debugw("tiered price line item - using quantity only",
				"item_price_id", chargebeeItemPriceID,
				"quantity", lineItem.Quantity,
				"flexprice_amount", item.Amount.String())
		} else {
			// For flat_fee/per_unit/package pricing: Use Quantity=1 + exact amount for precision
			// This ensures exact amount matching with FlexPrice's calculation
			lineItem.Quantity = 1
			lineItem.UnitAmount = convertAmountToSmallestUnit(item.Amount.InexactFloat64(), flexInvoice.Currency)
			s.Logger.Debugw("non-tiered price line item - using exact amount",
				"item_price_id", chargebeeItemPriceID,
				"quantity", 1,
				"unit_amount", lineItem.UnitAmount,
				"flexprice_amount", item.Amount.String())
		}

		// Add description if available (use DisplayName or PlanDisplayName)
		if item.DisplayName != nil && *item.DisplayName != "" {
			lineItem.Description = *item.DisplayName
		} else if item.PlanDisplayName != nil && *item.PlanDisplayName != "" {
			lineItem.Description = *item.PlanDisplayName
		}

		// Add date range if available
		if item.PeriodStart != nil {
			lineItem.DateFrom = item.PeriodStart
		}
		if item.PeriodEnd != nil {
			lineItem.DateTo = item.PeriodEnd
		}

		lineItems = append(lineItems, lineItem)
	}

	return lineItems, nil
}

// getChargebeeItemPriceIDAndCheckTiered retrieves or creates the Chargebee item price ID and checks if it uses tiered pricing
// If the price is not found in Chargebee but exists in FlexPrice, it creates the item and item price in Chargebee
func (s *InvoiceService) getChargebeeItemPriceIDAndCheckTiered(ctx context.Context, flexPriceID string) (string, bool, error) {
	if flexPriceID == "" {
		return "", false, ierr.NewError("price ID is required").
			WithHint("Line item must have a price ID").
			Mark(ierr.ErrValidation)
	}

	// Query entity mapping table for price ID
	filter := types.NewEntityIntegrationMappingFilter()
	filter.EntityID = flexPriceID
	filter.EntityType = types.IntegrationEntityTypeItemPrice
	filter.ProviderTypes = []string{string(types.SecretProviderChargebee)}

	mappings, err := s.EntityIntegrationMappingRepo.List(ctx, filter)
	if err != nil {
		return "", false, ierr.WithError(err).
			WithHint("Failed to get Chargebee item price mapping").
			Mark(ierr.ErrDatabase)
	}

	var chargebeeItemPriceID string
	var pricingModel string

	// If mapping exists, use it
	if len(mappings) > 0 {
		chargebeeItemPriceID = mappings[0].ProviderEntityID

		// Fetch the item price from Chargebee to check its pricing model
		result, err := s.Client.RetrieveItemPrice(ctx, chargebeeItemPriceID)
		if err != nil {
			s.Logger.Warnw("failed to retrieve item price from Chargebee, assuming flat_fee",
				"item_price_id", chargebeeItemPriceID,
				"error", err)
			// Fallback: assume flat_fee if we can't fetch
			return chargebeeItemPriceID, false, nil
		}

		pricingModel = string(result.ItemPrice.PricingModel)
	} else {
		// Mapping doesn't exist - create item and item price in Chargebee
		s.Logger.Infow("Chargebee item price not found, creating in Chargebee",
			"flexprice_price_id", flexPriceID)

		// Get the FlexPrice price details
		flexPrice, err := s.PriceRepo.Get(ctx, flexPriceID)
		if err != nil {
			return "", false, ierr.WithError(err).
				WithHint("Failed to get FlexPrice price").
				WithReportableDetails(map[string]interface{}{
					"flexprice_price_id": flexPriceID,
				}).
				Mark(ierr.ErrDatabase)
		}

		// Create item and item price in Chargebee
		createdItemPriceID, createdPricingModel, err := s.createItemAndItemPriceInChargebee(ctx, flexPrice)
		if err != nil {
			return "", false, ierr.WithError(err).
				WithHint("Failed to create item price in Chargebee").
				Mark(ierr.ErrInternal)
		}

		chargebeeItemPriceID = createdItemPriceID
		pricingModel = createdPricingModel

		s.Logger.Infow("successfully created item price in Chargebee",
			"flexprice_price_id", flexPriceID,
			"chargebee_item_price_id", chargebeeItemPriceID,
			"pricing_model", pricingModel)
	}

	isTiered := pricingModel == "tiered" || pricingModel == "volume" || pricingModel == "stairstep"

	s.Logger.Debugw("retrieved pricing model for item price",
		"item_price_id", chargebeeItemPriceID,
		"pricing_model", pricingModel,
		"is_tiered", isTiered)

	return chargebeeItemPriceID, isTiered, nil
}

// createItemAndItemPriceInChargebee creates both an item and item price in Chargebee for a FlexPrice price
// Returns the created item price ID and pricing model
func (s *InvoiceService) createItemAndItemPriceInChargebee(ctx context.Context, flexPrice *price.Price) (string, string, error) {
	// Step 1: Get or select the latest item family
	itemFamily, err := s.itemFamilyService.GetLatestItemFamily(ctx)
	if err != nil {
		s.Logger.Errorw("failed to get item family from Chargebee",
			"price_id", flexPrice.ID,
			"error", err)
		return "", "", ierr.WithError(err).
			WithHint("Failed to get item family from Chargebee. Please create an item family first").
			Mark(ierr.ErrNotFound)
	}

	s.Logger.Infow("using item family for price",
		"price_id", flexPrice.ID,
		"item_family_id", itemFamily.ID,
		"item_family_name", itemFamily.Name)

	// Step 2: Create unique item for this price
	// Format: charge_{uuid}
	uniqueUUID := types.GenerateUUID()
	itemID := fmt.Sprintf("charge_%s", uniqueUUID)

	// Use price ID as display name (or description if available)
	displayName := flexPrice.ID
	if flexPrice.Description != "" {
		displayName = flexPrice.Description
	}

	// Use display name as external name for better invoice display
	// Format: {display_name} - {currency} (e.g., "API Calls - USD" or "Pro Plan - USD")
	externalName := fmt.Sprintf("%s - %s", displayName, flexPrice.Currency)

	itemReq := &ItemCreateRequest{
		ID:              itemID,
		Name:            itemID,   // Name matches ID format (must be unique)
		Type:            "charge", // Charge type for one-time/recurring charges
		ItemFamilyID:    itemFamily.ID,
		Description:     flexPrice.Description,
		ExternalName:    externalName, // Customer-facing name on invoices
		EnabledInPortal: true,
	}

	item, err := s.itemService.CreateItem(ctx, itemReq)
	if err != nil {
		s.Logger.Errorw("failed to create item in Chargebee",
			"price_id", flexPrice.ID,
			"item_id", itemID,
			"error", err)
		return "", "", ierr.WithError(err).
			WithHint("Failed to create item in Chargebee").
			Mark(ierr.ErrInternal)
	}

	s.Logger.Infow("successfully created item in Chargebee",
		"price_id", flexPrice.ID,
		"item_id", item.ID)

	// Step 3: Create item price for this price
	// Item price ID should just be the FlexPrice price ID
	itemPriceID := flexPrice.ID

	// Map FlexPrice pricing model to Chargebee pricing model
	pricingModel := mapPricingModel(flexPrice)

	// Use display name as external name for better invoice display
	itemPriceExternalName := fmt.Sprintf("%s - %s", displayName, flexPrice.Currency)

	// Build item price request
	itemPriceReq := &ItemPriceCreateRequest{
		ID:           itemPriceID,
		ItemID:       item.ID,
		Name:         itemPriceID,
		ExternalName: itemPriceExternalName, // Customer-facing name on invoices
		PricingModel: pricingModel,
		CurrencyCode: flexPrice.Currency,
		Description:  flexPrice.Description,
	}

	// Set price or tiers based on billing model
	switch flexPrice.BillingModel {
	case types.BILLING_MODEL_TIERED:
		// For tiered/volume pricing, send tiers array
		// Convert domain tiers to types.PriceTier slice
		tiers := make([]*types.PriceTier, len(flexPrice.Tiers))
		for i := range flexPrice.Tiers {
			tiers[i] = &types.PriceTier{
				UpTo:       flexPrice.Tiers[i].UpTo,
				UnitAmount: flexPrice.Tiers[i].UnitAmount,
				FlatAmount: flexPrice.Tiers[i].FlatAmount,
			}
		}
		itemPriceReq.Tiers = convertTiersForChargebee(tiers, flexPrice.Currency)
		s.Logger.Infow("creating tiered pricing in Chargebee",
			"price_id", flexPrice.ID,
			"tier_mode", flexPrice.TierMode,
			"tier_count", len(flexPrice.Tiers))

	case types.BILLING_MODEL_PACKAGE:
		// For package pricing, set price and optionally period
		itemPriceReq.Price = convertAmountToSmallestUnit(flexPrice.Amount.InexactFloat64(), flexPrice.Currency)
		if flexPrice.TransformQuantity.DivideBy > 0 {
			divideBy := flexPrice.TransformQuantity.DivideBy
			itemPriceReq.Period = &divideBy
		}

	case types.BILLING_MODEL_FLAT_FEE:
		// For flat fee, just set the price
		itemPriceReq.Price = convertAmountToSmallestUnit(flexPrice.Amount.InexactFloat64(), flexPrice.Currency)

	default:
		// Default to flat fee
		itemPriceReq.Price = convertAmountToSmallestUnit(flexPrice.Amount.InexactFloat64(), flexPrice.Currency)
	}

	itemPrice, err := s.itemPriceService.CreateItemPrice(ctx, itemPriceReq)
	if err != nil {
		s.Logger.Errorw("failed to create item price in Chargebee",
			"price_id", flexPrice.ID,
			"item_price_id", itemPriceID,
			"item_id", item.ID,
			"error", err)
		return "", "", ierr.WithError(err).
			WithHint("Failed to create item price in Chargebee").
			Mark(ierr.ErrInternal)
	}

	s.Logger.Infow("successfully created item price in Chargebee",
		"price_id", flexPrice.ID,
		"item_price_id", itemPrice.ID,
		"item_id", item.ID,
		"amount", itemPrice.Price,
		"currency", itemPrice.CurrencyCode)

	// Step 4: Save entity mapping for price -> item_price
	mapping := &entityintegrationmapping.EntityIntegrationMapping{
		ID:               types.GenerateUUIDWithPrefix(types.UUID_PREFIX_ENTITY_INTEGRATION_MAPPING),
		EntityType:       types.IntegrationEntityTypeItemPrice,
		EntityID:         flexPrice.ID,
		ProviderType:     string(types.SecretProviderChargebee),
		ProviderEntityID: itemPrice.ID,
		EnvironmentID:    flexPrice.EnvironmentID,
		BaseModel:        types.GetDefaultBaseModel(ctx),
		Metadata: map[string]interface{}{
			"chargebee_charge_item_id": item.ID,
		},
	}
	// Override tenant_id from price
	mapping.TenantID = flexPrice.TenantID

	err = s.EntityIntegrationMappingRepo.Create(ctx, mapping)
	if err != nil {
		s.Logger.Errorw("failed to save price entity mapping",
			"price_id", flexPrice.ID,
			"chargebee_item_price_id", itemPrice.ID,
			"chargebee_item_id", item.ID,
			"error", err)
		// Don't fail the entire operation, just log the error
	}

	return itemPrice.ID, itemPrice.PricingModel, nil
}

// getChargebeeItemPriceIDSimple retrieves the Chargebee item price ID from entity mapping
func (s *InvoiceService) getChargebeeItemPriceIDSimple(ctx context.Context, flexPriceID string) (string, error) {
	if flexPriceID == "" {
		return "", ierr.NewError("price ID is required").
			WithHint("Line item must have a price ID").
			Mark(ierr.ErrValidation)
	}

	// Query entity mapping table for price ID
	filter := types.NewEntityIntegrationMappingFilter()
	filter.EntityID = flexPriceID
	filter.EntityType = types.IntegrationEntityTypeItemPrice
	filter.ProviderTypes = []string{string(types.SecretProviderChargebee)}

	mappings, err := s.EntityIntegrationMappingRepo.List(ctx, filter)
	if err != nil {
		return "", ierr.WithError(err).
			WithHint("Failed to get Chargebee item price mapping").
			Mark(ierr.ErrDatabase)
	}

	if len(mappings) == 0 {
		return "", ierr.NewError("Chargebee item price not found for FlexPrice price").
			WithHint("Price must be synced to Chargebee before creating invoice").
			WithReportableDetails(map[string]interface{}{
				"flexprice_price_id": flexPriceID,
			}).
			Mark(ierr.ErrNotFound)
	}

	return mappings[0].ProviderEntityID, nil
}

// getExistingChargebeeMapping checks if invoice is already synced to Chargebee
func (s *InvoiceService) getExistingChargebeeMapping(ctx context.Context, invoiceID string) (*entityintegrationmapping.EntityIntegrationMapping, error) {
	filter := types.NewEntityIntegrationMappingFilter()
	filter.EntityID = invoiceID
	filter.EntityType = types.IntegrationEntityTypeInvoice
	filter.ProviderTypes = []string{string(types.SecretProviderChargebee)}

	mappings, err := s.EntityIntegrationMappingRepo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return nil, ierr.NewError("invoice mapping not found").
			Mark(ierr.ErrNotFound)
	}

	return mappings[0], nil
}

// createInvoiceMapping creates an entity integration mapping for the invoice
func (s *InvoiceService) createInvoiceMapping(ctx context.Context, invoiceID, chargebeeInvoiceID, environmentID string) error {
	mapping := &entityintegrationmapping.EntityIntegrationMapping{
		ID:               types.GenerateUUIDWithPrefix(types.UUID_PREFIX_ENTITY_INTEGRATION_MAPPING),
		EntityType:       types.IntegrationEntityTypeInvoice,
		EntityID:         invoiceID,
		ProviderType:     string(types.SecretProviderChargebee),
		ProviderEntityID: chargebeeInvoiceID,
		EnvironmentID:    environmentID,
		BaseModel:        types.GetDefaultBaseModel(ctx),
	}

	// Get tenant_id from context or invoice
	// We'll need to get it from the invoice entity
	inv, err := s.InvoiceRepo.Get(ctx, invoiceID)
	if err == nil {
		mapping.TenantID = inv.TenantID
	}

	err = s.EntityIntegrationMappingRepo.Create(ctx, mapping)
	if err != nil {
		return ierr.WithError(err).
			WithHint("Failed to create invoice mapping").
			Mark(ierr.ErrDatabase)
	}

	return nil
}

// customerHasPaymentMethod checks if a Chargebee customer has a payment method
func (s *InvoiceService) customerHasPaymentMethod(ctx context.Context, chargebeeCustomerID string) (bool, error) {
	// Retrieve customer from Chargebee using client wrapper
	result, err := s.Client.RetrieveCustomer(ctx, chargebeeCustomerID)
	if err != nil {
		return false, ierr.WithError(err).
			WithHint("Failed to retrieve customer from Chargebee").
			Mark(ierr.ErrInternal)
	}

	customer := result.Customer

	// Check if customer has a primary payment source or payment method
	hasPaymentMethod := customer.PrimaryPaymentSourceId != "" || customer.PaymentMethod != nil

	return hasPaymentMethod, nil
}

// RetrieveInvoice retrieves an invoice from Chargebee
func (s *InvoiceService) RetrieveInvoice(ctx context.Context, invoiceID string) (*InvoiceResponse, error) {
	s.Logger.Infow("retrieving invoice from Chargebee",
		"invoice_id", invoiceID)

	// Retrieve invoice using client wrapper
	result, err := s.Client.RetrieveInvoice(ctx, invoiceID, &chargebeeInvoice.RetrieveRequestParams{})
	if err != nil {
		s.Logger.Errorw("failed to retrieve invoice from Chargebee",
			"invoice_id", invoiceID,
			"error", err)
		return nil, ierr.NewError("failed to retrieve invoice from Chargebee").
			WithReportableDetails(map[string]interface{}{
				"error":      err.Error(),
				"invoice_id": invoiceID,
			}).
			WithHint("Check if invoice exists in Chargebee").
			Mark(ierr.ErrNotFound)
	}

	invoiceData := result.Invoice

	s.Logger.Infow("successfully retrieved invoice from Chargebee",
		"invoice_id", invoiceData.Id,
		"customer_id", invoiceData.CustomerId)

	// Convert to our DTO format
	invoiceResponse := &InvoiceResponse{
		ID:              invoiceData.Id,
		CustomerID:      invoiceData.CustomerId,
		Status:          string(invoiceData.Status),
		AutoCollection:  "", // Not available in invoice object
		Total:           invoiceData.Total,
		AmountDue:       invoiceData.AmountDue,
		AmountPaid:      invoiceData.AmountPaid,
		CurrencyCode:    invoiceData.CurrencyCode,
		Date:            time.Unix(invoiceData.Date, 0),
		CreatedAt:       time.Unix(invoiceData.GeneratedAt, 0), // Use GeneratedAt as CreatedAt
		UpdatedAt:       time.Unix(invoiceData.UpdatedAt, 0),
		ResourceVersion: invoiceData.ResourceVersion,
	}

	if invoiceData.DueDate > 0 {
		dueDate := time.Unix(invoiceData.DueDate, 0)
		invoiceResponse.DueDate = &dueDate
	}

	// Convert line items if present
	if len(invoiceData.LineItems) > 0 {
		lineItems := make([]InvoiceLineItemResponse, 0, len(invoiceData.LineItems))
		for _, item := range invoiceData.LineItems {
			lineItem := InvoiceLineItemResponse{
				ID:          item.Id,
				ItemPriceID: item.EntityId, // EntityId contains the item price ID
				EntityType:  string(item.EntityType),
				Quantity:    int(item.Quantity),
				UnitAmount:  item.UnitAmount,
				Amount:      item.Amount,
				Description: item.Description,
				DateFrom:    time.Unix(item.DateFrom, 0),
				DateTo:      time.Unix(item.DateTo, 0),
			}
			lineItems = append(lineItems, lineItem)
		}
		invoiceResponse.LineItems = lineItems
	}

	return invoiceResponse, nil
}

// ReconcileInvoicePayment reconciles an invoice payment when a payment succeeds in Chargebee
// This is called from the webhook handler after a payment_succeeded event
// This method replicates the logic from invoiceService.ReconcilePaymentStatus but uses repositories directly
func (s *InvoiceService) ReconcileInvoicePayment(ctx context.Context, invoiceID string, paymentAmount decimal.Decimal) error {
	s.Logger.Infow("starting payment reconciliation with invoice",
		"invoice_id", invoiceID,
		"payment_amount", paymentAmount.String())

	// Get the invoice using repository
	inv, err := s.InvoiceRepo.Get(ctx, invoiceID)
	if err != nil {
		s.Logger.Errorw("failed to get invoice for reconciliation",
			"error", err,
			"invoice_id", invoiceID)
		return err
	}

	// Calculate new amounts (matching ReconcilePaymentStatus logic)
	newAmountPaid := inv.AmountPaid.Add(paymentAmount)

	// Determine payment status based on amounts
	var newPaymentStatus types.PaymentStatus
	var newAmountRemaining decimal.Decimal

	// Check if invoice is overpaid
	if newAmountPaid.GreaterThan(inv.AmountDue) {
		newPaymentStatus = types.PaymentStatusOverpaid
		// For overpaid invoices, amount_remaining is always 0
		newAmountRemaining = decimal.Zero
	} else if newAmountPaid.Equal(inv.AmountDue) {
		newPaymentStatus = types.PaymentStatusSucceeded
		newAmountRemaining = decimal.Zero
	} else {
		newPaymentStatus = types.PaymentStatusPending
		newAmountRemaining = inv.AmountDue.Sub(newAmountPaid)
	}

	s.Logger.Infow("calculated new amounts for reconciliation",
		"invoice_id", invoiceID,
		"payment_amount", paymentAmount.String(),
		"new_amount_paid", newAmountPaid.String(),
		"new_amount_remaining", newAmountRemaining.String(),
		"new_payment_status", newPaymentStatus)

	// Update invoice using repository
	inv.AmountPaid = newAmountPaid
	inv.AmountRemaining = newAmountRemaining
	inv.PaymentStatus = newPaymentStatus

	// Set PaidAt timestamp if payment succeeded or overpaid (and not already set)
	if (newPaymentStatus == types.PaymentStatusSucceeded || newPaymentStatus == types.PaymentStatusOverpaid) && inv.PaidAt == nil {
		now := time.Now().UTC()
		inv.PaidAt = &now
	}

	err = s.InvoiceRepo.Update(ctx, inv)
	if err != nil {
		s.Logger.Errorw("failed to update invoice payment status",
			"error", err,
			"invoice_id", invoiceID)
		return err
	}

	s.Logger.Infow("successfully reconciled invoice",
		"invoice_id", invoiceID,
		"payment_amount", paymentAmount.String(),
		"new_payment_status", newPaymentStatus)

	return nil
}

// GetFlexPriceInvoiceIDByChargebeeInvoiceID retrieves the FlexPrice invoice ID from entity mapping
func (s *InvoiceService) GetFlexPriceInvoiceIDByChargebeeInvoiceID(ctx context.Context, chargebeeInvoiceID string) (string, error) {
	filter := types.NewEntityIntegrationMappingFilter()
	filter.ProviderEntityIDs = []string{chargebeeInvoiceID}
	filter.ProviderTypes = []string{string(types.SecretProviderChargebee)}
	filter.EntityType = types.IntegrationEntityTypeInvoice

	mappings, err := s.EntityIntegrationMappingRepo.List(ctx, filter)
	if err != nil {
		return "", ierr.WithError(err).
			WithHint("Failed to get FlexPrice invoice mapping").
			Mark(ierr.ErrDatabase)
	}

	if len(mappings) == 0 {
		return "", ierr.NewError("FlexPrice invoice not found for Chargebee invoice").
			WithHint("Chargebee invoice not synced to FlexPrice").
			WithReportableDetails(map[string]interface{}{
				"chargebee_invoice_id": chargebeeInvoiceID,
			}).
			Mark(ierr.ErrNotFound)
	}

	return mappings[0].EntityID, nil
}

// ProcessChargebeePaymentFromWebhook processes a Chargebee payment and creates a FlexPrice payment record
func (s *InvoiceService) ProcessChargebeePaymentFromWebhook(
	ctx context.Context,
	flexpriceInvoiceID string,
	chargebeeTransactionID string,
	chargebeeInvoiceID string,
	amount decimal.Decimal,
	currency string,
	paymentMethod string,
) error {
	s.Logger.Infow("processing Chargebee payment from webhook",
		"flexprice_invoice_id", flexpriceInvoiceID,
		"chargebee_invoice_id", chargebeeInvoiceID,
		"chargebee_transaction_id", chargebeeTransactionID,
		"amount", amount.String(),
		"currency", currency)

	// Step 1: Check if payment already exists (idempotency check)
	// This prevents duplicate payment records when Chargebee retries webhooks
	exists, err := s.PaymentExistsByGatewayPaymentID(ctx, chargebeeTransactionID)
	if err != nil {
		s.Logger.Errorw("failed to check if payment exists by gateway payment ID",
			"error", err,
			"chargebee_transaction_id", chargebeeTransactionID)
		// Continue processing on error - fail-safe behavior
	} else if exists {
		s.Logger.Infow("payment already exists for this Chargebee transaction, skipping",
			"chargebee_transaction_id", chargebeeTransactionID,
			"chargebee_invoice_id", chargebeeInvoiceID,
			"flexprice_invoice_id", flexpriceInvoiceID)
		return nil
	}

	// Step 2: Get FlexPrice invoice using repository
	inv, err := s.InvoiceRepo.Get(ctx, flexpriceInvoiceID)
	if err != nil {
		s.Logger.Errorw("failed to get FlexPrice invoice",
			"error", err,
			"flexprice_invoice_id", flexpriceInvoiceID)
		return ierr.WithError(err).
			WithHint("Failed to get FlexPrice invoice").
			Mark(ierr.ErrDatabase)
	}

	// Step 3: Check if invoice is already succeeded (secondary check)
	if inv.PaymentStatus == types.PaymentStatusSucceeded {
		s.Logger.Infow("invoice already succeeded, skipping duplicate payment",
			"flexprice_invoice_id", flexpriceInvoiceID,
			"chargebee_invoice_id", chargebeeInvoiceID)
		return nil
	}

	// Step 3: Create payment record in FlexPrice using repository
	now := time.Now()
	createPaymentReq := dto.CreatePaymentRequest{
		IdempotencyKey:    chargebeeTransactionID, // Use transaction ID as idempotency key to prevent duplicates
		Amount:            amount,
		Currency:          currency,
		PaymentMethodType: types.PaymentMethodTypeCard, // Default to card
		DestinationType:   types.PaymentDestinationTypeInvoice,
		DestinationID:     flexpriceInvoiceID,
		ProcessPayment:    false, // Already processed by Chargebee
		Metadata: types.Metadata{
			"chargebee_transaction_id": chargebeeTransactionID,
			"chargebee_invoice_id":     chargebeeInvoiceID,
			"payment_method":           paymentMethod,
			"source":                   "chargebee_webhook",
		},
	}

	paymentModel, err := createPaymentReq.ToPayment(ctx)
	if err != nil {
		s.Logger.Errorw("failed to convert payment request to domain model",
			"error", err,
			"flexprice_invoice_id", flexpriceInvoiceID)
		return ierr.WithError(err).
			WithHint("Failed to create payment record").
			Mark(ierr.ErrValidation)
	}

	// Set gateway payment ID and succeeded status
	paymentModel.GatewayPaymentID = &chargebeeTransactionID
	paymentModel.PaymentStatus = types.PaymentStatusSucceeded
	paymentModel.SucceededAt = &now

	err = s.PaymentRepo.Create(ctx, paymentModel)
	if err != nil {
		s.Logger.Errorw("failed to create payment record",
			"error", err,
			"flexprice_invoice_id", flexpriceInvoiceID,
			"chargebee_transaction_id", chargebeeTransactionID)
		return ierr.WithError(err).
			WithHint("Failed to create payment record").
			Mark(ierr.ErrDatabase)
	}

	s.Logger.Infow("created payment record",
		"payment_id", paymentModel.ID,
		"flexprice_invoice_id", flexpriceInvoiceID,
		"chargebee_transaction_id", chargebeeTransactionID,
		"amount", amount.String())

	// Step 4: Reconcile invoice
	err = s.ReconcileInvoicePayment(ctx, flexpriceInvoiceID, amount)
	if err != nil {
		s.Logger.Errorw("failed to reconcile payment with invoice",
			"error", err,
			"flexprice_invoice_id", flexpriceInvoiceID,
			"amount", amount.String())
		return ierr.WithError(err).
			WithHint("Failed to reconcile invoice payment").
			Mark(ierr.ErrDatabase)
	}

	s.Logger.Infow("successfully processed Chargebee payment",
		"payment_id", paymentModel.ID,
		"flexprice_invoice_id", flexpriceInvoiceID,
		"chargebee_transaction_id", chargebeeTransactionID)

	return nil
}

// PaymentExistsByGatewayPaymentID checks if a payment already exists with the given gateway payment ID
// This is used for idempotency checks to prevent duplicate payment records from webhook retries
func (s *InvoiceService) PaymentExistsByGatewayPaymentID(
	ctx context.Context,
	gatewayPaymentID string,
) (bool, error) {
	if gatewayPaymentID == "" {
		return false, nil
	}

	// Create filter to query payments by gateway_payment_id
	filter := types.NewNoLimitPaymentFilter()
	limit := 1
	filter.QueryFilter.Limit = &limit
	filter.GatewayPaymentID = &gatewayPaymentID

	// Query payments using repository
	payments, err := s.PaymentRepo.List(ctx, filter)
	if err != nil {
		return false, err
	}

	// Return true if any payment exists with this gateway payment ID
	return len(payments) > 0, nil
}
