package dto

import (
	"context"

	"github.com/flexprice/flexprice/internal/domain/entityintegrationmapping"
	"github.com/flexprice/flexprice/internal/types"
	"github.com/flexprice/flexprice/internal/validator"
)

type CreateEntityIntegrationMappingRequest struct {
	EntityID         string                      `json:"entity_id" validate:"required,max=255"`
	EntityType       types.IntegrationEntityType `json:"entity_type" validate:"required"`
	ProviderType     string                      `json:"provider_type" validate:"required,max=50"`
	ProviderEntityID string                      `json:"provider_entity_id" validate:"required,max=255"`
	Metadata         map[string]interface{}      `json:"metadata,omitempty"`
}

type EntityIntegrationMappingResponse struct {
	ID               string                      `json:"id"`
	EntityID         string                      `json:"entity_id"`
	EntityType       types.IntegrationEntityType `json:"entity_type"`
	ProviderType     string                      `json:"provider_type"`
	ProviderEntityID string                      `json:"provider_entity_id"`
	EnvironmentID    string                      `json:"environment_id"`
	TenantID         string                      `json:"tenant_id"`
	Status           types.Status                `json:"status"`
	CreatedAt        string                      `json:"created_at"`
	UpdatedAt        string                      `json:"updated_at"`
	CreatedBy        string                      `json:"created_by"`
	UpdatedBy        string                      `json:"updated_by"`
}

// ListEntityIntegrationMappingsResponse represents the response for listing entity integration mappings
type ListEntityIntegrationMappingsResponse = types.ListResponse[*EntityIntegrationMappingResponse]

// ConnectCustomerToProviderRequest represents the request to connect a customer to an existing provider customer
type ConnectCustomerToProviderRequest struct {
	CustomerID         string `json:"customer_id" validate:"required,max=255"`
	ProviderCustomerID string `json:"provider_customer_id" validate:"required,max=255"`
	ProviderType       string `json:"provider_type" validate:"required,max=50"`
}

// ConnectCustomerToProviderResponse represents the response after connecting a customer to a provider
type ConnectCustomerToProviderResponse struct {
	Message              string                            `json:"message"`
	CustomerID           string                            `json:"customer_id"`
	ProviderCustomerID   string                            `json:"provider_customer_id"`
	ProviderType         string                            `json:"provider_type"`
	IntegrationMappingID string                            `json:"integration_mapping_id"`
	Mapping              *EntityIntegrationMappingResponse `json:"mapping,omitempty"`
}

func (r *CreateEntityIntegrationMappingRequest) Validate() error {
	return validator.ValidateRequest(r)
}

func (r *ConnectCustomerToProviderRequest) Validate() error {
	return validator.ValidateRequest(r)
}

func (r *CreateEntityIntegrationMappingRequest) ToEntityIntegrationMapping(ctx context.Context) *entityintegrationmapping.EntityIntegrationMapping {
	return &entityintegrationmapping.EntityIntegrationMapping{
		ID:               types.GenerateUUIDWithPrefix(types.UUID_PREFIX_ENTITY_INTEGRATION_MAPPING),
		EntityID:         r.EntityID,
		EntityType:       r.EntityType,
		ProviderType:     r.ProviderType,
		ProviderEntityID: r.ProviderEntityID,
		Metadata:         r.Metadata,
		EnvironmentID:    types.GetEnvironmentID(ctx),
		BaseModel:        types.GetDefaultBaseModel(ctx),
	}
}
