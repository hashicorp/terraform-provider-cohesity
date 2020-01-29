package models

import(
    "encoding/json"
)

/**
 * Type definition for VaultType1Enum enum
 */
type VaultType1Enum int

/**
 * Value collection for VaultType1Enum enum
 */
const (
    VaultType1_KNEARLINE VaultType1Enum = 1 + iota
    VaultType1_KCOLDLINE
    VaultType1_KGLACIER
    VaultType1_KS3
    VaultType1_KAZURESTANDARD
    VaultType1_KS3COMPATIBLE
    VaultType1_KQSTARTAPE
    VaultType1_KGOOGLESTANDARD
    VaultType1_KGOOGLEDRA
    VaultType1_KAWSGOVCLOUD
    VaultType1_KNAS
    VaultType1_KAZUREGOVCLOUD
)

func (r VaultType1Enum) MarshalJSON() ([]byte, error) { 
    s := VaultType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *VaultType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  VaultType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts VaultType1Enum to its string representation
 */
func VaultType1EnumToValue(vaultType1Enum VaultType1Enum) string {
    switch vaultType1Enum {
        case VaultType1_KNEARLINE:
    		return "kNearline"		
        case VaultType1_KCOLDLINE:
    		return "kColdline"		
        case VaultType1_KGLACIER:
    		return "kGlacier"		
        case VaultType1_KS3:
    		return "kS3"		
        case VaultType1_KAZURESTANDARD:
    		return "kAzureStandard"		
        case VaultType1_KS3COMPATIBLE:
    		return "kS3Compatible"		
        case VaultType1_KQSTARTAPE:
    		return "kQStarTape"		
        case VaultType1_KGOOGLESTANDARD:
    		return "kGoogleStandard"		
        case VaultType1_KGOOGLEDRA:
    		return "kGoogleDRA"		
        case VaultType1_KAWSGOVCLOUD:
    		return "kAWSGovCloud"		
        case VaultType1_KNAS:
    		return "kNAS"		
        case VaultType1_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"		
        default:
        	return "kNearline"
    }
}

/**
 * Converts VaultType1Enum Array to its string Array representation
*/
func VaultType1EnumArrayToValue(vaultType1Enum []VaultType1Enum) []string {
    convArray := make([]string,len( vaultType1Enum))
    for i:=0; i<len(vaultType1Enum);i++ {
        convArray[i] = VaultType1EnumToValue(vaultType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VaultType1EnumFromValue(value string) VaultType1Enum {
    switch value {
        case "kNearline":
            return VaultType1_KNEARLINE
        case "kColdline":
            return VaultType1_KCOLDLINE
        case "kGlacier":
            return VaultType1_KGLACIER
        case "kS3":
            return VaultType1_KS3
        case "kAzureStandard":
            return VaultType1_KAZURESTANDARD
        case "kS3Compatible":
            return VaultType1_KS3COMPATIBLE
        case "kQStarTape":
            return VaultType1_KQSTARTAPE
        case "kGoogleStandard":
            return VaultType1_KGOOGLESTANDARD
        case "kGoogleDRA":
            return VaultType1_KGOOGLEDRA
        case "kAWSGovCloud":
            return VaultType1_KAWSGOVCLOUD
        case "kNAS":
            return VaultType1_KNAS
        case "kAzureGovCloud":
            return VaultType1_KAZUREGOVCLOUD
        default:
            return VaultType1_KNEARLINE
    }
}
