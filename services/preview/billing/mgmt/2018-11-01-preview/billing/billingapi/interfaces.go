package billingapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/tcz001/azure-sdk-for-go/services/preview/billing/mgmt/2018-11-01-preview/billing"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, expand string) (result billing.Account, err error)
	List(ctx context.Context, expand string) (result billing.AccountListResult, err error)
	Update(ctx context.Context, billingAccountName string, parameters billing.AccountUpdateProperties) (result billing.AccountsUpdateFuture, err error)
}

var _ AccountsClientAPI = (*billing.AccountsClient)(nil)

// PaymentMethodsClientAPI contains the set of methods on the PaymentMethodsClient type.
type PaymentMethodsClientAPI interface {
	ListByBillingAccountName(ctx context.Context, billingAccountName string) (result billing.PaymentMethodsListResultPage, err error)
	ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.PaymentMethodsListResultPage, err error)
}

var _ PaymentMethodsClientAPI = (*billing.PaymentMethodsClient)(nil)

// AddressesClientAPI contains the set of methods on the AddressesClient type.
type AddressesClientAPI interface {
	Validate(ctx context.Context, address billing.Address) (result billing.ValidateAddressResponse, err error)
}

var _ AddressesClientAPI = (*billing.AddressesClient)(nil)

// AvailableBalancesClientAPI contains the set of methods on the AvailableBalancesClient type.
type AvailableBalancesClientAPI interface {
	GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.AvailableBalance, err error)
}

var _ AvailableBalancesClientAPI = (*billing.AvailableBalancesClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	Create(ctx context.Context, billingAccountName string, parameters billing.ProfileCreationParameters) (result billing.ProfilesCreateFuture, err error)
	Get(ctx context.Context, billingAccountName string, billingProfileName string, expand string) (result billing.Profile, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string) (result billing.ProfileListResult, err error)
	Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters billing.Profile) (result billing.ProfilesUpdateFuture, err error)
}

var _ ProfilesClientAPI = (*billing.ProfilesClient)(nil)

// InvoiceSectionsClientAPI contains the set of methods on the InvoiceSectionsClient type.
type InvoiceSectionsClientAPI interface {
	Create(ctx context.Context, billingAccountName string, parameters billing.InvoiceSectionCreationRequest) (result billing.InvoiceSectionsCreateFuture, err error)
	ElevateToBillingProfile(ctx context.Context, billingAccountName string, invoiceSectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, billingAccountName string, invoiceSectionName string, expand string) (result billing.InvoiceSection, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string) (result billing.InvoiceSectionListResult, err error)
	ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.InvoiceSectionListResult, err error)
	ListByCreateSubscriptionPermission(ctx context.Context, billingAccountName string, expand string) (result billing.InvoiceSectionListResult, err error)
	Update(ctx context.Context, billingAccountName string, invoiceSectionName string, parameters billing.InvoiceSection) (result billing.InvoiceSectionsUpdateFuture, err error)
}

var _ InvoiceSectionsClientAPI = (*billing.InvoiceSectionsClient)(nil)

// DepartmentsClientAPI contains the set of methods on the DepartmentsClient type.
type DepartmentsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, departmentName string, expand string, filter string) (result billing.Department, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string, filter string) (result billing.DepartmentListResult, err error)
}

var _ DepartmentsClientAPI = (*billing.DepartmentsClient)(nil)

// EnrollmentAccountsClientAPI contains the set of methods on the EnrollmentAccountsClient type.
type EnrollmentAccountsClientAPI interface {
	GetByEnrollmentAccountID(ctx context.Context, billingAccountName string, enrollmentAccountName string, expand string, filter string) (result billing.EnrollmentAccount, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string, filter string) (result billing.EnrollmentAccountListResult, err error)
}

var _ EnrollmentAccountsClientAPI = (*billing.EnrollmentAccountsClient)(nil)

// InvoicesClientAPI contains the set of methods on the InvoicesClient type.
type InvoicesClientAPI interface {
	Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceName string) (result billing.InvoiceSummary, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (result billing.InvoiceListResult, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string) (result billing.InvoiceListResult, err error)
}

var _ InvoicesClientAPI = (*billing.InvoicesClient)(nil)

// PriceSheetClientAPI contains the set of methods on the PriceSheetClient type.
type PriceSheetClientAPI interface {
	Download(ctx context.Context, billingAccountName string, invoiceName string) (result billing.PriceSheetDownloadFuture, err error)
}

var _ PriceSheetClientAPI = (*billing.PriceSheetClient)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string) (result billing.SubscriptionSummary, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string) (result billing.SubscriptionsListResultPage, err error)
	ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.SubscriptionsListResult, err error)
	ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string) (result billing.SubscriptionsListResult, err error)
	Transfer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string, parameters billing.TransferBillingSubscriptionRequestProperties) (result billing.SubscriptionsTransferFuture, err error)
	ValidateTransfer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingSubscriptionName string, parameters billing.TransferBillingSubscriptionRequestProperties) (result billing.ValidateSubscriptionTransferEligibilityResult, err error)
}

var _ SubscriptionsClientAPI = (*billing.SubscriptionsClient)(nil)

// ProductsClientAPI contains the set of methods on the ProductsClient type.
type ProductsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string) (result billing.ProductSummary, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, filter string) (result billing.ProductsListResultPage, err error)
	ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, filter string) (result billing.ProductsListResult, err error)
	Transfer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters billing.TransferProductRequestProperties) (result billing.ProductSummary, err error)
	UpdateAutoRenewByBillingAccountName(ctx context.Context, billingAccountName string, productName string, body billing.UpdateAutoRenewRequest) (result billing.UpdateAutoRenewOperationSummary, err error)
	UpdateAutoRenewByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, body billing.UpdateAutoRenewRequest) (result billing.UpdateAutoRenewOperationSummary, err error)
	ValidateTransfer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters billing.TransferProductRequestProperties) (result billing.ValidateProductTransferEligibilityResult, err error)
}

var _ ProductsClientAPI = (*billing.ProductsClient)(nil)

// TransactionsClientAPI contains the set of methods on the TransactionsClient type.
type TransactionsClientAPI interface {
	ListByBillingAccountName(ctx context.Context, billingAccountName string, startDate string, endDate string, filter string) (result billing.TransactionsListResultPage, err error)
	ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, startDate string, endDate string, filter string) (result billing.TransactionsListResult, err error)
	ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, startDate string, endDate string, filter string) (result billing.TransactionsListResult, err error)
}

var _ TransactionsClientAPI = (*billing.TransactionsClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	GetByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.Policy, err error)
	Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters billing.Policy) (result billing.Policy, err error)
}

var _ PoliciesClientAPI = (*billing.PoliciesClient)(nil)

// PropertyClientAPI contains the set of methods on the PropertyClient type.
type PropertyClientAPI interface {
	Get(ctx context.Context) (result billing.Property, err error)
}

var _ PropertyClientAPI = (*billing.PropertyClient)(nil)

// TransfersClientAPI contains the set of methods on the TransfersClient type.
type TransfersClientAPI interface {
	Cancel(ctx context.Context, billingAccountName string, invoiceSectionName string, transferName string) (result billing.TransferDetails, err error)
	Get(ctx context.Context, billingAccountName string, invoiceSectionName string, transferName string) (result billing.TransferDetails, err error)
	Initiate(ctx context.Context, billingAccountName string, invoiceSectionName string, body billing.InitiateTransferRequest) (result billing.TransferDetails, err error)
	List(ctx context.Context, billingAccountName string, invoiceSectionName string) (result billing.TransferDetailsListResultPage, err error)
}

var _ TransfersClientAPI = (*billing.TransfersClient)(nil)

// RecipientTransfersClientAPI contains the set of methods on the RecipientTransfersClient type.
type RecipientTransfersClientAPI interface {
	Accept(ctx context.Context, transferName string, body billing.AcceptTransferRequest) (result billing.RecipientTransferDetails, err error)
	Decline(ctx context.Context, transferName string) (result billing.RecipientTransferDetails, err error)
	Get(ctx context.Context, transferName string) (result billing.RecipientTransferDetails, err error)
	List(ctx context.Context) (result billing.RecipientTransferDetailsListResultPage, err error)
}

var _ RecipientTransfersClientAPI = (*billing.RecipientTransfersClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result billing.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*billing.OperationsClient)(nil)

// PermissionsClientAPI contains the set of methods on the PermissionsClient type.
type PermissionsClientAPI interface {
	ListByBillingAccount(ctx context.Context, billingAccountName string) (result billing.PermissionsListResult, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.PermissionsListResult, err error)
	ListByInvoiceSections(ctx context.Context, billingAccountName string, invoiceSectionName string) (result billing.PermissionsListResult, err error)
}

var _ PermissionsClientAPI = (*billing.PermissionsClient)(nil)

// RoleDefinitionsClientAPI contains the set of methods on the RoleDefinitionsClient type.
type RoleDefinitionsClientAPI interface {
	GetByBillingAccountName(ctx context.Context, billingAccountName string, billingRoleDefinitionName string) (result billing.RoleDefinition, err error)
	GetByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string) (result billing.RoleDefinition, err error)
	GetByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, billingRoleDefinitionName string) (result billing.RoleDefinition, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string) (result billing.RoleDefinitionListResult, err error)
	ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.RoleDefinitionListResult, err error)
	ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string) (result billing.RoleDefinitionListResult, err error)
}

var _ RoleDefinitionsClientAPI = (*billing.RoleDefinitionsClient)(nil)

// RoleAssignmentsClientAPI contains the set of methods on the RoleAssignmentsClient type.
type RoleAssignmentsClientAPI interface {
	AddByBillingAccountName(ctx context.Context, billingAccountName string, parameters billing.RoleAssignmentPayload) (result billing.RoleAssignmentListResult, err error)
	AddByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, parameters billing.RoleAssignmentPayload) (result billing.RoleAssignmentListResult, err error)
	AddByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, parameters billing.RoleAssignmentPayload) (result billing.RoleAssignmentListResult, err error)
	DeleteByBillingAccountName(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	DeleteByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	DeleteByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	GetByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	GetByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string) (result billing.RoleAssignmentListResult, err error)
	ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.RoleAssignmentListResult, err error)
	ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string) (result billing.RoleAssignmentListResult, err error)
}

var _ RoleAssignmentsClientAPI = (*billing.RoleAssignmentsClient)(nil)

// AgreementsClientAPI contains the set of methods on the AgreementsClient type.
type AgreementsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, agreementName string, expand string) (result billing.Agreement, err error)
	ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string) (result billing.AgreementListResult, err error)
}

var _ AgreementsClientAPI = (*billing.AgreementsClient)(nil)

// LineOfCreditsClientAPI contains the set of methods on the LineOfCreditsClient type.
type LineOfCreditsClientAPI interface {
	Get(ctx context.Context) (result billing.LineOfCredit, err error)
	Update(ctx context.Context, parameters billing.LineOfCredit) (result billing.LineOfCreditsUpdateFuture, err error)
}

var _ LineOfCreditsClientAPI = (*billing.LineOfCreditsClient)(nil)
