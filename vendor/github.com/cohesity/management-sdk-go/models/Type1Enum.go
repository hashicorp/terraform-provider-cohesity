package models

import(
    "encoding/json"
)

/**
 * Type definition for Type1Enum enum
 */
type Type1Enum int

/**
 * Value collection for Type1Enum enum
 */
const (
    Type1_KIAMUSER Type1Enum = 1 + iota
    Type1_KREGION
    Type1_KAVAILABILITYZONE
    Type1_KEC2INSTANCE
    Type1_KVPC
    Type1_KSUBNET
    Type1_KNETWORKSECURITYGROUP
    Type1_KINSTANCETYPE
    Type1_KKEYPAIR
)

func (r Type1Enum) MarshalJSON() ([]byte, error) { 
    s := Type1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type1Enum to its string representation
 */
func Type1EnumToValue(type1Enum Type1Enum) string {
    switch type1Enum {
        case Type1_KIAMUSER:
    		return "kIAMUser"		
        case Type1_KREGION:
    		return "kRegion"		
        case Type1_KAVAILABILITYZONE:
    		return "kAvailabilityZone"		
        case Type1_KEC2INSTANCE:
    		return "kEC2Instance"		
        case Type1_KVPC:
    		return "kVPC"		
        case Type1_KSUBNET:
    		return "kSubnet"		
        case Type1_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"		
        case Type1_KINSTANCETYPE:
    		return "kInstanceType"		
        case Type1_KKEYPAIR:
    		return "kKeyPair"		
        default:
        	return "kIAMUser"
    }
}

/**
 * Converts Type1Enum Array to its string Array representation
*/
func Type1EnumArrayToValue(type1Enum []Type1Enum) []string {
    convArray := make([]string,len( type1Enum))
    for i:=0; i<len(type1Enum);i++ {
        convArray[i] = Type1EnumToValue(type1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type1EnumFromValue(value string) Type1Enum {
    switch value {
        case "kIAMUser":
            return Type1_KIAMUSER
        case "kRegion":
            return Type1_KREGION
        case "kAvailabilityZone":
            return Type1_KAVAILABILITYZONE
        case "kEC2Instance":
            return Type1_KEC2INSTANCE
        case "kVPC":
            return Type1_KVPC
        case "kSubnet":
            return Type1_KSUBNET
        case "kNetworkSecurityGroup":
            return Type1_KNETWORKSECURITYGROUP
        case "kInstanceType":
            return Type1_KINSTANCETYPE
        case "kKeyPair":
            return Type1_KKEYPAIR
        default:
            return Type1_KIAMUSER
    }
}
