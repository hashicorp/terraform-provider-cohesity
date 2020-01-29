// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeProtectionPolicyRequestEnum enum
 */
type TypeProtectionPolicyRequestEnum int

/**
 * Value collection for TypeProtectionPolicyRequestEnum enum
 */
const (
    TypeProtectionPolicyRequest_KREGULAR TypeProtectionPolicyRequestEnum = 1 + iota
    TypeProtectionPolicyRequest_KRPO
)

func (r TypeProtectionPolicyRequestEnum) MarshalJSON() ([]byte, error) {
    s := TypeProtectionPolicyRequestEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeProtectionPolicyRequestEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeProtectionPolicyRequestEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeProtectionPolicyRequestEnum to its string representation
 */
func TypeProtectionPolicyRequestEnumToValue(typeProtectionPolicyRequestEnum TypeProtectionPolicyRequestEnum) string {
    switch typeProtectionPolicyRequestEnum {
        case TypeProtectionPolicyRequest_KREGULAR:
    		return "kRegular"
        case TypeProtectionPolicyRequest_KRPO:
    		return "kRPO"
        default:
        	return "kRegular"
    }
}

/**
 * Converts TypeProtectionPolicyRequestEnum Array to its string Array representation
*/
func TypeProtectionPolicyRequestEnumArrayToValue(typeProtectionPolicyRequestEnum []TypeProtectionPolicyRequestEnum) []string {
    convArray := make([]string,len( typeProtectionPolicyRequestEnum))
    for i:=0; i<len(typeProtectionPolicyRequestEnum);i++ {
        convArray[i] = TypeProtectionPolicyRequestEnumToValue(typeProtectionPolicyRequestEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeProtectionPolicyRequestEnumFromValue(value string) TypeProtectionPolicyRequestEnum {
    switch value {
        case "kRegular":
            return TypeProtectionPolicyRequest_KREGULAR
        case "kRPO":
            return TypeProtectionPolicyRequest_KRPO
        default:
            return TypeProtectionPolicyRequest_KREGULAR
    }
}
