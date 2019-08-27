package computeapi

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
	"github.com/tcz001/azure-sdk-for-go/services/compute/mgmt/2015-06-15/compute"
)

// AvailabilitySetsClientAPI contains the set of methods on the AvailabilitySetsClient type.
type AvailabilitySetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters compute.AvailabilitySet) (result compute.AvailabilitySet, err error)
	Delete(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.OperationStatusResponse, err error)
	Get(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.AvailabilitySet, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.AvailabilitySetListResultPage, err error)
	ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.VirtualMachineSizeListResult, err error)
}

var _ AvailabilitySetsClientAPI = (*compute.AvailabilitySetsClient)(nil)

// VirtualMachineExtensionImagesClientAPI contains the set of methods on the VirtualMachineExtensionImagesClient type.
type VirtualMachineExtensionImagesClientAPI interface {
	Get(ctx context.Context, location string, publisherName string, typeParameter string, version string) (result compute.VirtualMachineExtensionImage, err error)
	ListTypes(ctx context.Context, location string, publisherName string) (result compute.ListVirtualMachineExtensionImage, err error)
	ListVersions(ctx context.Context, location string, publisherName string, typeParameter string, filter string, top *int32, orderby string) (result compute.ListVirtualMachineExtensionImage, err error)
}

var _ VirtualMachineExtensionImagesClientAPI = (*compute.VirtualMachineExtensionImagesClient)(nil)

// VirtualMachineExtensionsClientAPI contains the set of methods on the VirtualMachineExtensionsClient type.
type VirtualMachineExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, extensionParameters compute.VirtualMachineExtension) (result compute.VirtualMachineExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string) (result compute.VirtualMachineExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, expand string) (result compute.VirtualMachineExtension, err error)
	Update(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, extensionParameters compute.VirtualMachineExtensionUpdate) (result compute.VirtualMachineExtensionsUpdateFuture, err error)
}

var _ VirtualMachineExtensionsClientAPI = (*compute.VirtualMachineExtensionsClient)(nil)

// VirtualMachineImagesClientAPI contains the set of methods on the VirtualMachineImagesClient type.
type VirtualMachineImagesClientAPI interface {
	Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (result compute.VirtualMachineImage, err error)
	List(ctx context.Context, location string, publisherName string, offer string, skus string, filter string, top *int32, orderby string) (result compute.ListVirtualMachineImageResource, err error)
	ListOffers(ctx context.Context, location string, publisherName string) (result compute.ListVirtualMachineImageResource, err error)
	ListPublishers(ctx context.Context, location string) (result compute.ListVirtualMachineImageResource, err error)
	ListSkus(ctx context.Context, location string, publisherName string, offer string) (result compute.ListVirtualMachineImageResource, err error)
}

var _ VirtualMachineImagesClientAPI = (*compute.VirtualMachineImagesClient)(nil)

// UsageClientAPI contains the set of methods on the UsageClient type.
type UsageClientAPI interface {
	List(ctx context.Context, location string) (result compute.ListUsagesResultPage, err error)
}

var _ UsageClientAPI = (*compute.UsageClient)(nil)

// VirtualMachineSizesClientAPI contains the set of methods on the VirtualMachineSizesClient type.
type VirtualMachineSizesClientAPI interface {
	List(ctx context.Context, location string) (result compute.VirtualMachineSizeListResult, err error)
}

var _ VirtualMachineSizesClientAPI = (*compute.VirtualMachineSizesClient)(nil)

// VirtualMachinesClientAPI contains the set of methods on the VirtualMachinesClient type.
type VirtualMachinesClientAPI interface {
	Capture(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineCaptureParameters) (result compute.VirtualMachinesCaptureFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachine) (result compute.VirtualMachinesCreateOrUpdateFuture, err error)
	Deallocate(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesDeleteFuture, err error)
	Generalize(ctx context.Context, resourceGroupName string, VMName string) (result compute.OperationStatusResponse, err error)
	Get(ctx context.Context, resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineListResultPage, err error)
	ListAll(ctx context.Context) (result compute.VirtualMachineListResultPage, err error)
	ListAvailableSizes(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachineSizeListResult, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesPowerOffFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesRedeployFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesStartFuture, err error)
}

var _ VirtualMachinesClientAPI = (*compute.VirtualMachinesClient)(nil)

// VirtualMachineScaleSetsClientAPI contains the set of methods on the VirtualMachineScaleSetsClient type.
type VirtualMachineScaleSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters compute.VirtualMachineScaleSet) (result compute.VirtualMachineScaleSetsCreateOrUpdateFuture, err error)
	Deallocate(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetsDeleteFuture, err error)
	DeleteInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (result compute.VirtualMachineScaleSetsDeleteInstancesFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSet, err error)
	GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetInstanceView, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineScaleSetListResultPage, err error)
	ListAll(ctx context.Context) (result compute.VirtualMachineScaleSetListWithLinkResultPage, err error)
	ListSkus(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListSkusResultPage, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsPowerOffFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetsReimageFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsStartFuture, err error)
	UpdateInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (result compute.VirtualMachineScaleSetsUpdateInstancesFuture, err error)
}

var _ VirtualMachineScaleSetsClientAPI = (*compute.VirtualMachineScaleSetsClient)(nil)

// VirtualMachineScaleSetVMsClientAPI contains the set of methods on the VirtualMachineScaleSetVMsClient type.
type VirtualMachineScaleSetVMsClientAPI interface {
	Deallocate(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVM, err error)
	GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMInstanceView, err error)
	List(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result compute.VirtualMachineScaleSetVMListResultPage, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsPowerOffFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsReimageFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsStartFuture, err error)
}

var _ VirtualMachineScaleSetVMsClientAPI = (*compute.VirtualMachineScaleSetVMsClient)(nil)
