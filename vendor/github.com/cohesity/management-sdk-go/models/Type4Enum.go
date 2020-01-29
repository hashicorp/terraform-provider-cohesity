package models

import(
    "encoding/json"
)

/**
 * Type definition for Type4Enum enum
 */
type Type4Enum int

/**
 * Value collection for Type4Enum enum
 */
const (
    Type4_KIAMUSER Type4Enum = 1 + iota
    Type4_KPROJECT
    Type4_KREGION
    Type4_KAVAILABILITYZONE
    Type4_KVIRTUALMACHINE
    Type4_KVPC
    Type4_KSUBNET
    Type4_KNETWORKSECURITYGROUP
    Type4_KINSTANCETYPE
)

func (r Type4Enum) MarshalJSON() ([]byte, error) { 
    s := Type4EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type4Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type4EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type4Enum to its string representation
 */
func Type4EnumToValue(type4Enum Type4Enum) string {
    switch type4Enum {
        case Type4_KIAMUSER:
    		return "kIAMUser"		
        case Type4_KPROJECT:
    		return "kProject"		
        case Type4_KREGION:
    		return "kRegion"		
        case Type4_KAVAILABILITYZONE:
    		return "kAvailabilityZone"		
        case Type4_KVIRTUALMACHINE:
    		return "kVirtualMachine"		
        case Type4_KVPC:
    		return "kVPC"		
        case Type4_KSUBNET:
    		return "kSubnet"		
        case Type4_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"		
        case Type4_KINSTANCETYPE:
    		return "kInstanceType"		
        default:
        	return "kIAMUser"
    }
}

/**
 * Converts Type4Enum Array to its string Array representation
*/
func Type4EnumArrayToValue(type4Enum []Type4Enum) []string {
    convArray := make([]string,len( type4Enum))
    for i:=0; i<len(type4Enum);i++ {
        convArray[i] = Type4EnumToValue(type4Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type4EnumFromValue(value string) Type4Enum {
    switch value {
        case "kIAMUser":
            return Type4_KIAMUSER
        case "kProject":
            return Type4_KPROJECT
        case "kRegion":
            return Type4_KREGION
        case "kAvailabilityZone":
            return Type4_KAVAILABILITYZONE
        case "kVirtualMachine":
            return Type4_KVIRTUALMACHINE
        case "kVPC":
            return Type4_KVPC
        case "kSubnet":
            return Type4_KSUBNET
        case "kNetworkSecurityGroup":
            return Type4_KNETWORKSECURITYGROUP
        case "kInstanceType":
            return Type4_KINSTANCETYPE
        default:
            return Type4_KIAMUSER
    }
}
