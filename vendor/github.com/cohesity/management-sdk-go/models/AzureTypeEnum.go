// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AzureTypeEnum enum
 */
type AzureTypeEnum int

/**
 * Value collection for AzureTypeEnum enum
 */
const (
    AzureType_KSUBSCRIPTION AzureTypeEnum = 1 + iota
    AzureType_KRESOURCEGROUP
    AzureType_KVIRTUALMACHINE
    AzureType_KSTORAGEACCOUNT
    AzureType_KSTORAGEKEY
    AzureType_KSTORAGECONTAINER
    AzureType_KSTORAGEBLOB
    AzureType_KSTORAGERESOURCEGROUP
    AzureType_KNETWORKSECURITYGROUP
    AzureType_KVIRTUALNETWORK
    AzureType_KNETWORKRESOURCEGROUP
    AzureType_KSUBNET
    AzureType_KCOMPUTEOPTIONS
)

func (r AzureTypeEnum) MarshalJSON() ([]byte, error) {
    s := AzureTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AzureTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AzureTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AzureTypeEnum to its string representation
 */
func AzureTypeEnumToValue(azureTypeEnum AzureTypeEnum) string {
    switch azureTypeEnum {
        case AzureType_KSUBSCRIPTION:
    		return "kSubscription"
        case AzureType_KRESOURCEGROUP:
    		return "kResourceGroup"
        case AzureType_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case AzureType_KSTORAGEACCOUNT:
    		return "kStorageAccount"
        case AzureType_KSTORAGEKEY:
    		return "kStorageKey"
        case AzureType_KSTORAGECONTAINER:
    		return "kStorageContainer"
        case AzureType_KSTORAGEBLOB:
    		return "kStorageBlob"
        case AzureType_KSTORAGERESOURCEGROUP:
    		return "kStorageResourceGroup"
        case AzureType_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"
        case AzureType_KVIRTUALNETWORK:
    		return "kVirtualNetwork"
        case AzureType_KNETWORKRESOURCEGROUP:
    		return "kNetworkResourceGroup"
        case AzureType_KSUBNET:
    		return "kSubnet"
        case AzureType_KCOMPUTEOPTIONS:
    		return "kComputeOptions"
        default:
        	return "kSubscription"
    }
}

/**
 * Converts AzureTypeEnum Array to its string Array representation
*/
func AzureTypeEnumArrayToValue(azureTypeEnum []AzureTypeEnum) []string {
    convArray := make([]string,len( azureTypeEnum))
    for i:=0; i<len(azureTypeEnum);i++ {
        convArray[i] = AzureTypeEnumToValue(azureTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AzureTypeEnumFromValue(value string) AzureTypeEnum {
    switch value {
        case "kSubscription":
            return AzureType_KSUBSCRIPTION
        case "kResourceGroup":
            return AzureType_KRESOURCEGROUP
        case "kVirtualMachine":
            return AzureType_KVIRTUALMACHINE
        case "kStorageAccount":
            return AzureType_KSTORAGEACCOUNT
        case "kStorageKey":
            return AzureType_KSTORAGEKEY
        case "kStorageContainer":
            return AzureType_KSTORAGECONTAINER
        case "kStorageBlob":
            return AzureType_KSTORAGEBLOB
        case "kStorageResourceGroup":
            return AzureType_KSTORAGERESOURCEGROUP
        case "kNetworkSecurityGroup":
            return AzureType_KNETWORKSECURITYGROUP
        case "kVirtualNetwork":
            return AzureType_KVIRTUALNETWORK
        case "kNetworkResourceGroup":
            return AzureType_KNETWORKRESOURCEGROUP
        case "kSubnet":
            return AzureType_KSUBNET
        case "kComputeOptions":
            return AzureType_KCOMPUTEOPTIONS
        default:
            return AzureType_KSUBSCRIPTION
    }
}
