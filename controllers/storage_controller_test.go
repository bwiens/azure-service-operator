// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all storage

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	config "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageControllerHappyPathWithoutNetworkRule(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	StorageAccountName := GenerateAlphaNumTestResourceName("sadev")
	// Create the ResourceGroup object and expect the Reconcile to be created
	saInstance := &azurev1alpha1.Storage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      StorageAccountName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.StorageSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Sku: azurev1alpha1.StorageSku{
				Name: "Standard_RAGRS",
			},
			Kind:                   "StorageV2",
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}
	// create rg
	EnsureInstance(ctx, t, tc, saInstance)
	// delete rg
	EnsureDelete(ctx, t, tc, saInstance)
}
func TestStorageControllerHappyPathWithNetworkRule(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error

	StorageAccountName := GenerateTestResourceNameWithRandom("storage", 10)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	VNetName := GenerateTestResourceNameWithRandom("svnet", 10)
	subnetName := "subnet-storage-test"

	// Create a VNET as prereq for the test

	VNetSubNetInstance := azurev1alpha1.VNetSubnets{
		SubnetName:          subnetName,
		SubnetAddressPrefix: "110.1.0.0/16",
	}

	// Create a VNET
	VNetInstance := &azurev1alpha1.VirtualNetwork{
		ObjectMeta: metav1.ObjectMeta{
			Name:      VNetName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.VirtualNetworkSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			AddressSpace:  "110.0.0.0/8",
			Subnets:       []azurev1alpha1.VNetSubnets{VNetSubNetInstance},
		},
	}

	EnsureInstance(ctx, t, tc, VNetInstance)

	subnetID := "/subscriptions/" + config.SubscriptionID() + "/resourceGroups/" + rgName + "/providers/Microsoft.Network/virtualNetworks/" + VNetName + "/subnets/" + subnetName
	vnetRules := []azurev1alpha1.VirtualNetworkRule{
		{
			SubnetId: &subnetID,
		},
	}
	ipAddress := "1.1.1.1"
	ipRange := "2.2.2.2/24"
	ipRules := []azurev1alpha1.IPRule{
		{
			IPAddressOrRange: &ipAddress,
		},
		{
			IPAddressOrRange: &ipRange,
		},
	}

	// Create the storage account object and expect the Reconcile to be created
	cnInstance := &azurev1alpha1.Storage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      StorageAccountName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.StorageSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.StorageSku{
				Name: "Standard_RAGRS",
			},
			Kind:                   "StorageV2",
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
			NetworkRule: &azurev1alpha1.StorageNetworkRuleSet{
				Bypass:              "AzureServices",
				VirtualNetworkRules: &vnetRules,
				IPRules:             &ipRules,
				DefaultAction:       "Deny",
			},
		},
	}

	err = tc.k8sClient.Create(ctx, cnInstance)
	assert.Equal(nil, err, "create StorageAccount  in k8s")

	storageAccountNamespacedName := types.NamespacedName{Name: StorageAccountName, Namespace: "default"}

	// Wait for the APIMgmtAPI instance to be written to k8s
	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, storageAccountNamespacedName, cnInstance)
		return strings.Contains(cnInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "awaiting storageAccount instance creation")

}
