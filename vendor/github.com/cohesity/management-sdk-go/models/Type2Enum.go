package models

import(
    "encoding/json"
)

/**
 * Type definition for Type2Enum enum
 */
type Type2Enum int

/**
 * Value collection for Type2Enum enum
 */
const (
    Type2_KSUBSCRIPTION Type2Enum = 1 + iota
    Type2_KRESOURCEGROUP
    Type2_KVIRTUALMACHINE
    Type2_KSTORAGEACCOUNT
    Type2_KSTORAGEKEY
    Type2_KSTORAGECONTAINER
    Type2_KSTORAGEBLOB
    Type2_KSTORAGERESOURCEGROUP
    Type2_KNETWORKSECURITYGROUP
    Type2_KVIRTUALNETWORK
    Type2_KNETWORKRESOURCEGROUP
    Type2_KSUBNET
    Type2_KCOMPUTEOPTIONS
)

func (r Type2Enum) MarshalJSON() ([]byte, error) { 
    s := Type2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type2Enum to its string representation
 */
func Type2EnumToValue(type2Enum Type2Enum) string {
    switch type2Enum {
        case Type2_KSUBSCRIPTION:
    		return "kSubscription"		
        case Type2_KRESOURCEGROUP:
    		return "kResourceGroup"		
        case Type2_KVIRTUALMACHINE:
    		return "kVirtualMachine"		
        case Type2_KSTORAGEACCOUNT:
    		return "kStorageAccount"		
        case Type2_KSTORAGEKEY:
    		return "kStorageKey"		
        case Type2_KSTORAGECONTAINER:
    		return "kStorageContainer"		
        case Type2_KSTORAGEBLOB:
    		return "kStorageBlob"		
        case Type2_KSTORAGERESOURCEGROUP:
    		return "kStorageResourceGroup"		
        case Type2_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"		
        case Type2_KVIRTUALNETWORK:
    		return "kVirtualNetwork"		
        case Type2_KNETWORKRESOURCEGROUP:
    		return "kNetworkResourceGroup"		
        case Type2_KSUBNET:
    		return "kSubnet"		
        case Type2_KCOMPUTEOPTIONS:
    		return "kComputeOptions"		
        default:
        	return "kSubscription"
    }
}

/**
 * Converts Type2Enum Array to its string Array representation
*/
func Type2EnumArrayToValue(type2Enum []Type2Enum) []string {
    convArray := make([]string,len( type2Enum))
    for i:=0; i<len(type2Enum);i++ {
        convArray[i] = Type2EnumToValue(type2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type2EnumFromValue(value string) Type2Enum {
    switch value {
        case "kSubscription":
            return Type2_KSUBSCRIPTION
        case "kResourceGroup":
            return Type2_KRESOURCEGROUP
        case "kVirtualMachine":
            return Type2_KVIRTUALMACHINE
        case "kStorageAccount":
            return Type2_KSTORAGEACCOUNT
        case "kStorageKey":
            return Type2_KSTORAGEKEY
        case "kStorageContainer":
            return Type2_KSTORAGECONTAINER
        case "kStorageBlob":
            return Type2_KSTORAGEBLOB
        case "kStorageResourceGroup":
            return Type2_KSTORAGERESOURCEGROUP
        case "kNetworkSecurityGroup":
            return Type2_KNETWORKSECURITYGROUP
        case "kVirtualNetwork":
            return Type2_KVIRTUALNETWORK
        case "kNetworkResourceGroup":
            return Type2_KNETWORKRESOURCEGROUP
        case "kSubnet":
            return Type2_KSUBNET
        case "kComputeOptions":
            return Type2_KCOMPUTEOPTIONS
        default:
            return Type2_KSUBSCRIPTION
    }
}
