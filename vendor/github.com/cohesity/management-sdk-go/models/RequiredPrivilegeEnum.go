package models

import(
    "encoding/json"
)

/**
 * Type definition for RequiredPrivilegeEnum enum
 */
type RequiredPrivilegeEnum int

/**
 * Value collection for RequiredPrivilegeEnum enum
 */
const (
    RequiredPrivilege_KREADACCESS RequiredPrivilegeEnum = 1 + iota
    RequiredPrivilege_KREADWRITEACCESS
    RequiredPrivilege_KMANAGEMENTACCESS
)

func (r RequiredPrivilegeEnum) MarshalJSON() ([]byte, error) { 
    s := RequiredPrivilegeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *RequiredPrivilegeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  RequiredPrivilegeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts RequiredPrivilegeEnum to its string representation
 */
func RequiredPrivilegeEnumToValue(requiredPrivilegeEnum RequiredPrivilegeEnum) string {
    switch requiredPrivilegeEnum {
        case RequiredPrivilege_KREADACCESS:
    		return "kReadAccess"		
        case RequiredPrivilege_KREADWRITEACCESS:
    		return "kReadWriteAccess"		
        case RequiredPrivilege_KMANAGEMENTACCESS:
    		return "kManagementAccess"		
        default:
        	return "kReadAccess"
    }
}

/**
 * Converts RequiredPrivilegeEnum Array to its string Array representation
*/
func RequiredPrivilegeEnumArrayToValue(requiredPrivilegeEnum []RequiredPrivilegeEnum) []string {
    convArray := make([]string,len( requiredPrivilegeEnum))
    for i:=0; i<len(requiredPrivilegeEnum);i++ {
        convArray[i] = RequiredPrivilegeEnumToValue(requiredPrivilegeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RequiredPrivilegeEnumFromValue(value string) RequiredPrivilegeEnum {
    switch value {
        case "kReadAccess":
            return RequiredPrivilege_KREADACCESS
        case "kReadWriteAccess":
            return RequiredPrivilege_KREADWRITEACCESS
        case "kManagementAccess":
            return RequiredPrivilege_KMANAGEMENTACCESS
        default:
            return RequiredPrivilege_KREADACCESS
    }
}
