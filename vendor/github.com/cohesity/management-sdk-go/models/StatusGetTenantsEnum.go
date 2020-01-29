// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusGetTenantsEnum enum
 */
type StatusGetTenantsEnum int

/**
 * Value collection for StatusGetTenantsEnum enum
 */
const (
    StatusGetTenants_ACTIVE StatusGetTenantsEnum = 1 + iota
    StatusGetTenants_DEACTIVATED
    StatusGetTenants_DELETED
)

func (r StatusGetTenantsEnum) MarshalJSON() ([]byte, error) {
    s := StatusGetTenantsEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusGetTenantsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusGetTenantsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusGetTenantsEnum to its string representation
 */
func StatusGetTenantsEnumToValue(statusGetTenantsEnum StatusGetTenantsEnum) string {
    switch statusGetTenantsEnum {
        case StatusGetTenants_ACTIVE:
    		return "Active"
        case StatusGetTenants_DEACTIVATED:
    		return "Deactivated"
        case StatusGetTenants_DELETED:
    		return "Deleted"
        default:
        	return "Active"
    }
}

/**
 * Converts StatusGetTenantsEnum Array to its string Array representation
*/
func StatusGetTenantsEnumArrayToValue(statusGetTenantsEnum []StatusGetTenantsEnum) []string {
    convArray := make([]string,len( statusGetTenantsEnum))
    for i:=0; i<len(statusGetTenantsEnum);i++ {
        convArray[i] = StatusGetTenantsEnumToValue(statusGetTenantsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusGetTenantsEnumFromValue(value string) StatusGetTenantsEnum {
    switch value {
        case "Active":
            return StatusGetTenants_ACTIVE
        case "Deactivated":
            return StatusGetTenants_DEACTIVATED
        case "Deleted":
            return StatusGetTenants_DELETED
        default:
            return StatusGetTenants_ACTIVE
    }
}
