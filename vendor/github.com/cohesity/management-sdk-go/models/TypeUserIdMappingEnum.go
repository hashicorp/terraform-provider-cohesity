// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeUserIdMappingEnum enum
 */
type TypeUserIdMappingEnum int

/**
 * Value collection for TypeUserIdMappingEnum enum
 */
const (
    TypeUserIdMapping_KRID TypeUserIdMappingEnum = 1 + iota
    TypeUserIdMapping_KRFC2307
    TypeUserIdMapping_KSFU30
    TypeUserIdMapping_KCENTRIFY
    TypeUserIdMapping_KFIXED
    TypeUserIdMapping_KCUSTOMATTRIBUTES
    TypeUserIdMapping_KLDAPPROVIDER
)

func (r TypeUserIdMappingEnum) MarshalJSON() ([]byte, error) {
    s := TypeUserIdMappingEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeUserIdMappingEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeUserIdMappingEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeUserIdMappingEnum to its string representation
 */
func TypeUserIdMappingEnumToValue(typeUserIdMappingEnum TypeUserIdMappingEnum) string {
    switch typeUserIdMappingEnum {
        case TypeUserIdMapping_KRID:
    		return "kRid"
        case TypeUserIdMapping_KRFC2307:
    		return "kRfc2307"
        case TypeUserIdMapping_KSFU30:
    		return "kSfu30"
        case TypeUserIdMapping_KCENTRIFY:
    		return "kCentrify"
        case TypeUserIdMapping_KFIXED:
    		return "kFixed"
        case TypeUserIdMapping_KCUSTOMATTRIBUTES:
    		return "kCustomAttributes"
        case TypeUserIdMapping_KLDAPPROVIDER:
    		return "kLdapProvider"
        default:
        	return "kRid"
    }
}

/**
 * Converts TypeUserIdMappingEnum Array to its string Array representation
*/
func TypeUserIdMappingEnumArrayToValue(typeUserIdMappingEnum []TypeUserIdMappingEnum) []string {
    convArray := make([]string,len( typeUserIdMappingEnum))
    for i:=0; i<len(typeUserIdMappingEnum);i++ {
        convArray[i] = TypeUserIdMappingEnumToValue(typeUserIdMappingEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeUserIdMappingEnumFromValue(value string) TypeUserIdMappingEnum {
    switch value {
        case "kRid":
            return TypeUserIdMapping_KRID
        case "kRfc2307":
            return TypeUserIdMapping_KRFC2307
        case "kSfu30":
            return TypeUserIdMapping_KSFU30
        case "kCentrify":
            return TypeUserIdMapping_KCENTRIFY
        case "kFixed":
            return TypeUserIdMapping_KFIXED
        case "kCustomAttributes":
            return TypeUserIdMapping_KCUSTOMATTRIBUTES
        case "kLdapProvider":
            return TypeUserIdMapping_KLDAPPROVIDER
        default:
            return TypeUserIdMapping_KRID
    }
}
