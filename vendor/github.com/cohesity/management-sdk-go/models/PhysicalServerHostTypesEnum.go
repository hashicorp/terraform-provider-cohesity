package models

import(
    "encoding/json"
)

/**
 * Type definition for PhysicalServerHostTypesEnum enum
 */
type PhysicalServerHostTypesEnum int

/**
 * Value collection for PhysicalServerHostTypesEnum enum
 */
const (
    PhysicalServerHostTypes_KLINUX PhysicalServerHostTypesEnum = 1 + iota
    PhysicalServerHostTypes_KWINDOWS
    PhysicalServerHostTypes_KAIX
    PhysicalServerHostTypes_KSOLARIS
)

func (r PhysicalServerHostTypesEnum) MarshalJSON() ([]byte, error) { 
    s := PhysicalServerHostTypesEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *PhysicalServerHostTypesEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  PhysicalServerHostTypesEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts PhysicalServerHostTypesEnum to its string representation
 */
func PhysicalServerHostTypesEnumToValue(physicalServerHostTypesEnum PhysicalServerHostTypesEnum) string {
    switch physicalServerHostTypesEnum {
        case PhysicalServerHostTypes_KLINUX:
    		return "kLinux"		
        case PhysicalServerHostTypes_KWINDOWS:
    		return "kWindows"		
        case PhysicalServerHostTypes_KAIX:
    		return "kAix"		
        case PhysicalServerHostTypes_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts PhysicalServerHostTypesEnum Array to its string Array representation
*/
func PhysicalServerHostTypesEnumArrayToValue(physicalServerHostTypesEnum []PhysicalServerHostTypesEnum) []string {
    convArray := make([]string,len( physicalServerHostTypesEnum))
    for i:=0; i<len(physicalServerHostTypesEnum);i++ {
        convArray[i] = PhysicalServerHostTypesEnumToValue(physicalServerHostTypesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PhysicalServerHostTypesEnumFromValue(value string) PhysicalServerHostTypesEnum {
    switch value {
        case "kLinux":
            return PhysicalServerHostTypes_KLINUX
        case "kWindows":
            return PhysicalServerHostTypes_KWINDOWS
        case "kAix":
            return PhysicalServerHostTypes_KAIX
        case "kSolaris":
            return PhysicalServerHostTypes_KSOLARIS
        default:
            return PhysicalServerHostTypes_KLINUX
    }
}
