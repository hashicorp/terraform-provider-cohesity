// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeAzureProtectionSourceEnum enum
 */
type TypeAzureProtectionSourceEnum int

/**
 * Value collection for TypeAzureProtectionSourceEnum enum
 */
const (
    TypeAzureProtectionSource_KSUBSCRIPTION TypeAzureProtectionSourceEnum = 1 + iota
    TypeAzureProtectionSource_KRESOURCEGROUP
    TypeAzureProtectionSource_KVIRTUALMACHINE
    TypeAzureProtectionSource_KSTORAGEACCOUNT
    TypeAzureProtectionSource_KSTORAGEKEY
    TypeAzureProtectionSource_KSTORAGECONTAINER
    TypeAzureProtectionSource_KSTORAGEBLOB
    TypeAzureProtectionSource_KSTORAGERESOURCEGROUP
    TypeAzureProtectionSource_KNETWORKSECURITYGROUP
    TypeAzureProtectionSource_KVIRTUALNETWORK
    TypeAzureProtectionSource_KNETWORKRESOURCEGROUP
    TypeAzureProtectionSource_KSUBNET
    TypeAzureProtectionSource_KCOMPUTEOPTIONS
)

func (r TypeAzureProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeAzureProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeAzureProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeAzureProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeAzureProtectionSourceEnum to its string representation
 */
func TypeAzureProtectionSourceEnumToValue(typeAzureProtectionSourceEnum TypeAzureProtectionSourceEnum) string {
    switch typeAzureProtectionSourceEnum {
        case TypeAzureProtectionSource_KSUBSCRIPTION:
    		return "kSubscription"
        case TypeAzureProtectionSource_KRESOURCEGROUP:
    		return "kResourceGroup"
        case TypeAzureProtectionSource_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case TypeAzureProtectionSource_KSTORAGEACCOUNT:
    		return "kStorageAccount"
        case TypeAzureProtectionSource_KSTORAGEKEY:
    		return "kStorageKey"
        case TypeAzureProtectionSource_KSTORAGECONTAINER:
    		return "kStorageContainer"
        case TypeAzureProtectionSource_KSTORAGEBLOB:
    		return "kStorageBlob"
        case TypeAzureProtectionSource_KSTORAGERESOURCEGROUP:
    		return "kStorageResourceGroup"
        case TypeAzureProtectionSource_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"
        case TypeAzureProtectionSource_KVIRTUALNETWORK:
    		return "kVirtualNetwork"
        case TypeAzureProtectionSource_KNETWORKRESOURCEGROUP:
    		return "kNetworkResourceGroup"
        case TypeAzureProtectionSource_KSUBNET:
    		return "kSubnet"
        case TypeAzureProtectionSource_KCOMPUTEOPTIONS:
    		return "kComputeOptions"
        default:
        	return "kSubscription"
    }
}

/**
 * Converts TypeAzureProtectionSourceEnum Array to its string Array representation
*/
func TypeAzureProtectionSourceEnumArrayToValue(typeAzureProtectionSourceEnum []TypeAzureProtectionSourceEnum) []string {
    convArray := make([]string,len( typeAzureProtectionSourceEnum))
    for i:=0; i<len(typeAzureProtectionSourceEnum);i++ {
        convArray[i] = TypeAzureProtectionSourceEnumToValue(typeAzureProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeAzureProtectionSourceEnumFromValue(value string) TypeAzureProtectionSourceEnum {
    switch value {
        case "kSubscription":
            return TypeAzureProtectionSource_KSUBSCRIPTION
        case "kResourceGroup":
            return TypeAzureProtectionSource_KRESOURCEGROUP
        case "kVirtualMachine":
            return TypeAzureProtectionSource_KVIRTUALMACHINE
        case "kStorageAccount":
            return TypeAzureProtectionSource_KSTORAGEACCOUNT
        case "kStorageKey":
            return TypeAzureProtectionSource_KSTORAGEKEY
        case "kStorageContainer":
            return TypeAzureProtectionSource_KSTORAGECONTAINER
        case "kStorageBlob":
            return TypeAzureProtectionSource_KSTORAGEBLOB
        case "kStorageResourceGroup":
            return TypeAzureProtectionSource_KSTORAGERESOURCEGROUP
        case "kNetworkSecurityGroup":
            return TypeAzureProtectionSource_KNETWORKSECURITYGROUP
        case "kVirtualNetwork":
            return TypeAzureProtectionSource_KVIRTUALNETWORK
        case "kNetworkResourceGroup":
            return TypeAzureProtectionSource_KNETWORKRESOURCEGROUP
        case "kSubnet":
            return TypeAzureProtectionSource_KSUBNET
        case "kComputeOptions":
            return TypeAzureProtectionSource_KCOMPUTEOPTIONS
        default:
            return TypeAzureProtectionSource_KSUBSCRIPTION
    }
}
