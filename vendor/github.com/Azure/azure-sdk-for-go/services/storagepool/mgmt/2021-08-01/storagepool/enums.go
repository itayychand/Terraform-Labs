package storagepool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// DiskPoolTier enumerates the values for disk pool tier.
type DiskPoolTier string

const (
	// DiskPoolTierBasic ...
	DiskPoolTierBasic DiskPoolTier = "Basic"
	// DiskPoolTierPremium ...
	DiskPoolTierPremium DiskPoolTier = "Premium"
	// DiskPoolTierStandard ...
	DiskPoolTierStandard DiskPoolTier = "Standard"
)

// PossibleDiskPoolTierValues returns an array of possible values for the DiskPoolTier const type.
func PossibleDiskPoolTierValues() []DiskPoolTier {
	return []DiskPoolTier{DiskPoolTierBasic, DiskPoolTierPremium, DiskPoolTierStandard}
}

// IscsiTargetACLMode enumerates the values for iscsi target acl mode.
type IscsiTargetACLMode string

const (
	// IscsiTargetACLModeDynamic ...
	IscsiTargetACLModeDynamic IscsiTargetACLMode = "Dynamic"
	// IscsiTargetACLModeStatic ...
	IscsiTargetACLModeStatic IscsiTargetACLMode = "Static"
)

// PossibleIscsiTargetACLModeValues returns an array of possible values for the IscsiTargetACLMode const type.
func PossibleIscsiTargetACLModeValues() []IscsiTargetACLMode {
	return []IscsiTargetACLMode{IscsiTargetACLModeDynamic, IscsiTargetACLModeStatic}
}

// OperationalStatus enumerates the values for operational status.
type OperationalStatus string

const (
	// OperationalStatusHealthy ...
	OperationalStatusHealthy OperationalStatus = "Healthy"
	// OperationalStatusInvalid ...
	OperationalStatusInvalid OperationalStatus = "Invalid"
	// OperationalStatusRunning ...
	OperationalStatusRunning OperationalStatus = "Running"
	// OperationalStatusStopped ...
	OperationalStatusStopped OperationalStatus = "Stopped"
	// OperationalStatusStoppeddeallocated ...
	OperationalStatusStoppeddeallocated OperationalStatus = "Stopped (deallocated)"
	// OperationalStatusUnhealthy ...
	OperationalStatusUnhealthy OperationalStatus = "Unhealthy"
	// OperationalStatusUnknown ...
	OperationalStatusUnknown OperationalStatus = "Unknown"
	// OperationalStatusUpdating ...
	OperationalStatusUpdating OperationalStatus = "Updating"
)

// PossibleOperationalStatusValues returns an array of possible values for the OperationalStatus const type.
func PossibleOperationalStatusValues() []OperationalStatus {
	return []OperationalStatus{OperationalStatusHealthy, OperationalStatusInvalid, OperationalStatusRunning, OperationalStatusStopped, OperationalStatusStoppeddeallocated, OperationalStatusUnhealthy, OperationalStatusUnknown, OperationalStatusUpdating}
}

// ProvisioningStates enumerates the values for provisioning states.
type ProvisioningStates string

const (
	// ProvisioningStatesCanceled ...
	ProvisioningStatesCanceled ProvisioningStates = "Canceled"
	// ProvisioningStatesCreating ...
	ProvisioningStatesCreating ProvisioningStates = "Creating"
	// ProvisioningStatesDeleting ...
	ProvisioningStatesDeleting ProvisioningStates = "Deleting"
	// ProvisioningStatesFailed ...
	ProvisioningStatesFailed ProvisioningStates = "Failed"
	// ProvisioningStatesInvalid ...
	ProvisioningStatesInvalid ProvisioningStates = "Invalid"
	// ProvisioningStatesPending ...
	ProvisioningStatesPending ProvisioningStates = "Pending"
	// ProvisioningStatesSucceeded ...
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	// ProvisioningStatesUpdating ...
	ProvisioningStatesUpdating ProvisioningStates = "Updating"
)

// PossibleProvisioningStatesValues returns an array of possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{ProvisioningStatesCanceled, ProvisioningStatesCreating, ProvisioningStatesDeleting, ProvisioningStatesFailed, ProvisioningStatesInvalid, ProvisioningStatesPending, ProvisioningStatesSucceeded, ProvisioningStatesUpdating}
}

// ResourceSkuRestrictionsReasonCode enumerates the values for resource sku restrictions reason code.
type ResourceSkuRestrictionsReasonCode string

const (
	// ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ...
	ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	// ResourceSkuRestrictionsReasonCodeQuotaID ...
	ResourceSkuRestrictionsReasonCodeQuotaID ResourceSkuRestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSkuRestrictionsReasonCodeValues returns an array of possible values for the ResourceSkuRestrictionsReasonCode const type.
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return []ResourceSkuRestrictionsReasonCode{ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription, ResourceSkuRestrictionsReasonCodeQuotaID}
}

// ResourceSkuRestrictionsType enumerates the values for resource sku restrictions type.
type ResourceSkuRestrictionsType string

const (
	// ResourceSkuRestrictionsTypeLocation ...
	ResourceSkuRestrictionsTypeLocation ResourceSkuRestrictionsType = "Location"
	// ResourceSkuRestrictionsTypeZone ...
	ResourceSkuRestrictionsTypeZone ResourceSkuRestrictionsType = "Zone"
)

// PossibleResourceSkuRestrictionsTypeValues returns an array of possible values for the ResourceSkuRestrictionsType const type.
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return []ResourceSkuRestrictionsType{ResourceSkuRestrictionsTypeLocation, ResourceSkuRestrictionsTypeZone}
}
