package networkapi

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
	"github.com/tcz001/azure-sdk-for-go/services/network/mgmt/2015-06-15/network"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string) (result network.DNSNameAvailabilityResult, err error)
}

var _ BaseClientAPI = (*network.BaseClient)(nil)

// ApplicationGatewaysClientAPI contains the set of methods on the ApplicationGatewaysClient type.
type ApplicationGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters network.ApplicationGateway) (result network.ApplicationGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGateway, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ApplicationGatewayListResultPage, err error)
	ListAll(ctx context.Context) (result network.ApplicationGatewayListResultPage, err error)
	Start(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysStopFuture, err error)
}

var _ ApplicationGatewaysClientAPI = (*network.ApplicationGatewaysClient)(nil)

// ExpressRouteCircuitAuthorizationsClientAPI contains the set of methods on the ExpressRouteCircuitAuthorizationsClient type.
type ExpressRouteCircuitAuthorizationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters network.ExpressRouteCircuitAuthorization) (result network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result network.ExpressRouteCircuitAuthorizationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result network.ExpressRouteCircuitAuthorization, err error)
	List(ctx context.Context, resourceGroupName string, circuitName string) (result network.AuthorizationListResultPage, err error)
}

var _ ExpressRouteCircuitAuthorizationsClientAPI = (*network.ExpressRouteCircuitAuthorizationsClient)(nil)

// ExpressRouteCircuitPeeringsClientAPI contains the set of methods on the ExpressRouteCircuitPeeringsClient type.
type ExpressRouteCircuitPeeringsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters network.ExpressRouteCircuitPeering) (result network.ExpressRouteCircuitPeeringsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitPeeringsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitPeering, err error)
	List(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitPeeringListResultPage, err error)
}

var _ ExpressRouteCircuitPeeringsClientAPI = (*network.ExpressRouteCircuitPeeringsClient)(nil)

// ExpressRouteCircuitsClientAPI contains the set of methods on the ExpressRouteCircuitsClient type.
type ExpressRouteCircuitsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters network.ExpressRouteCircuit) (result network.ExpressRouteCircuitsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuit, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ExpressRouteCircuitListResultPage, err error)
	ListAll(ctx context.Context) (result network.ExpressRouteCircuitListResultPage, err error)
	ListArpTable(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitsArpTableListResultPage, err error)
	ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitsRoutesTableListResultPage, err error)
	ListStats(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitsStatsListResultPage, err error)
}

var _ ExpressRouteCircuitsClientAPI = (*network.ExpressRouteCircuitsClient)(nil)

// ExpressRouteServiceProvidersClientAPI contains the set of methods on the ExpressRouteServiceProvidersClient type.
type ExpressRouteServiceProvidersClientAPI interface {
	List(ctx context.Context) (result network.ExpressRouteServiceProviderListResultPage, err error)
}

var _ ExpressRouteServiceProvidersClientAPI = (*network.ExpressRouteServiceProvidersClient)(nil)

// LoadBalancersClientAPI contains the set of methods on the LoadBalancersClient type.
type LoadBalancersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters network.LoadBalancer) (result network.LoadBalancersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, expand string) (result network.LoadBalancer, err error)
	List(ctx context.Context, resourceGroupName string) (result network.LoadBalancerListResultPage, err error)
	ListAll(ctx context.Context) (result network.LoadBalancerListResultPage, err error)
}

var _ LoadBalancersClientAPI = (*network.LoadBalancersClient)(nil)

// InterfacesClientAPI contains the set of methods on the InterfacesClient type.
type InterfacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters network.Interface) (result network.InterfacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand string) (result network.Interface, err error)
	GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.Interface, err error)
	List(ctx context.Context, resourceGroupName string) (result network.InterfaceListResultPage, err error)
	ListAll(ctx context.Context) (result network.InterfaceListResultPage, err error)
	ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.InterfaceListResultPage, err error)
	ListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result network.InterfaceListResultPage, err error)
}

var _ InterfacesClientAPI = (*network.InterfacesClient)(nil)

// SecurityGroupsClientAPI contains the set of methods on the SecurityGroupsClient type.
type SecurityGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters network.SecurityGroup) (result network.SecurityGroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityGroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (result network.SecurityGroup, err error)
	List(ctx context.Context, resourceGroupName string) (result network.SecurityGroupListResultPage, err error)
	ListAll(ctx context.Context) (result network.SecurityGroupListResultPage, err error)
}

var _ SecurityGroupsClientAPI = (*network.SecurityGroupsClient)(nil)

// SecurityRulesClientAPI contains the set of methods on the SecurityRulesClient type.
type SecurityRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters network.SecurityRule) (result network.SecurityRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result network.SecurityRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result network.SecurityRule, err error)
	List(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityRuleListResultPage, err error)
}

var _ SecurityRulesClientAPI = (*network.SecurityRulesClient)(nil)

// PublicIPAddressesClientAPI contains the set of methods on the PublicIPAddressesClient type.
type PublicIPAddressesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (result network.PublicIPAddressesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error)
	List(ctx context.Context, resourceGroupName string) (result network.PublicIPAddressListResultPage, err error)
	ListAll(ctx context.Context) (result network.PublicIPAddressListResultPage, err error)
}

var _ PublicIPAddressesClientAPI = (*network.PublicIPAddressesClient)(nil)

// RouteTablesClientAPI contains the set of methods on the RouteTablesClient type.
type RouteTablesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters network.RouteTable) (result network.RouteTablesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteTablesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeTableName string, expand string) (result network.RouteTable, err error)
	List(ctx context.Context, resourceGroupName string) (result network.RouteTableListResultPage, err error)
	ListAll(ctx context.Context) (result network.RouteTableListResultPage, err error)
}

var _ RouteTablesClientAPI = (*network.RouteTablesClient)(nil)

// RoutesClientAPI contains the set of methods on the RoutesClient type.
type RoutesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route) (result network.RoutesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (result network.RoutesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (result network.Route, err error)
	List(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteListResultPage, err error)
}

var _ RoutesClientAPI = (*network.RoutesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result network.UsagesListResultPage, err error)
}

var _ UsagesClientAPI = (*network.UsagesClient)(nil)

// VirtualNetworksClientAPI contains the set of methods on the VirtualNetworksClient type.
type VirtualNetworksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters network.VirtualNetwork) (result network.VirtualNetworksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, expand string) (result network.VirtualNetwork, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkListResultPage, err error)
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
}

var _ VirtualNetworksClientAPI = (*network.VirtualNetworksClient)(nil)

// SubnetsClientAPI contains the set of methods on the SubnetsClient type.
type SubnetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters network.Subnet) (result network.SubnetsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result network.SubnetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result network.Subnet, err error)
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.SubnetListResultPage, err error)
}

var _ SubnetsClientAPI = (*network.SubnetsClient)(nil)

// VirtualNetworkGatewaysClientAPI contains the set of methods on the VirtualNetworkGatewaysClient type.
type VirtualNetworkGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VirtualNetworkGateway) (result network.VirtualNetworkGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewaysDeleteFuture, err error)
	Generatevpnclientpackage(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VpnClientParameters) (result network.VirtualNetworkGatewaysGeneratevpnclientpackageFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGateway, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayListResultPage, err error)
	Reset(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VirtualNetworkGateway) (result network.VirtualNetworkGatewaysResetFuture, err error)
}

var _ VirtualNetworkGatewaysClientAPI = (*network.VirtualNetworkGatewaysClient)(nil)

// VirtualNetworkGatewayConnectionsClientAPI contains the set of methods on the VirtualNetworkGatewayConnectionsClient type.
type VirtualNetworkGatewayConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.VirtualNetworkGatewayConnection) (result network.VirtualNetworkGatewayConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.VirtualNetworkGatewayConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.VirtualNetworkGatewayConnection, err error)
	GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.ConnectionSharedKeyResult, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayConnectionListResultPage, err error)
	ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.ConnectionResetSharedKey) (result network.VirtualNetworkGatewayConnectionsResetSharedKeyFuture, err error)
	SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.ConnectionSharedKey) (result network.VirtualNetworkGatewayConnectionsSetSharedKeyFuture, err error)
}

var _ VirtualNetworkGatewayConnectionsClientAPI = (*network.VirtualNetworkGatewayConnectionsClient)(nil)

// LocalNetworkGatewaysClientAPI contains the set of methods on the LocalNetworkGatewaysClient type.
type LocalNetworkGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters network.LocalNetworkGateway) (result network.LocalNetworkGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result network.LocalNetworkGatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result network.LocalNetworkGateway, err error)
	List(ctx context.Context, resourceGroupName string) (result network.LocalNetworkGatewayListResultPage, err error)
}

var _ LocalNetworkGatewaysClientAPI = (*network.LocalNetworkGatewaysClient)(nil)
