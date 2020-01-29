// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeProtectionPolicyEnum enum
 */
type TypeProtectionPolicyEnum int

/**
 * Value collection for TypeProtectionPolicyEnum enum
 */
const (
    TypeProtectionPolicy_KREGULAR TypeProtectionPolicyEnum = 1 + iota
    TypeProtectionPolicy_KRPO
)

func (r TypeProtectionPolicyEnum) MarshalJSON() ([]byte, error) {
    s := TypeProtectionPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeProtectionPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeProtectionPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeProtectionPolicyEnum to its string representation
 */
func TypeProtectionPolicyEnumToValue(typeProtectionPolicyEnum TypeProtectionPolicyEnum) string {
    switch typeProtectionPolicyEnum {
        case TypeProtectionPolicy_KREGULAR:
    		return "kRegular"
        case TypeProtectionPolicy_KRPO:
    		return "kRPO"
        default:
        	return "kRegular"
    }
}

/**
 * Converts TypeProtectionPolicyEnum Array to its string Array representation
*/
func TypeProtectionPolicyEnumArrayToValue(typeProtectionPolicyEnum []TypeProtectionPolicyEnum) []string {
    convArray := make([]string,len( typeProtectionPolicyEnum))
    for i:=0; i<len(typeProtectionPolicyEnum);i++ {
        convArray[i] = TypeProtectionPolicyEnumToValue(typeProtectionPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeProtectionPolicyEnumFromValue(value string) TypeProtectionPolicyEnum {
    switch value {
        case "kRegular":
            return TypeProtectionPolicy_KREGULAR
        case "kRPO":
            return TypeProtectionPolicy_KRPO
        default:
            return TypeProtectionPolicy_KREGULAR
    }
}
