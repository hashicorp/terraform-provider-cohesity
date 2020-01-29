// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeOffice365ProtectionSourceEnum enum
 */
type TypeOffice365ProtectionSourceEnum int

/**
 * Value collection for TypeOffice365ProtectionSourceEnum enum
 */
const (
    TypeOffice365ProtectionSource_KDOMAIN TypeOffice365ProtectionSourceEnum = 1 + iota
    TypeOffice365ProtectionSource_KOUTLOOK
    TypeOffice365ProtectionSource_KMAILBOX
)

func (r TypeOffice365ProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeOffice365ProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeOffice365ProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeOffice365ProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeOffice365ProtectionSourceEnum to its string representation
 */
func TypeOffice365ProtectionSourceEnumToValue(typeOffice365ProtectionSourceEnum TypeOffice365ProtectionSourceEnum) string {
    switch typeOffice365ProtectionSourceEnum {
        case TypeOffice365ProtectionSource_KDOMAIN:
    		return "kDomain"
        case TypeOffice365ProtectionSource_KOUTLOOK:
    		return "kOutlook"
        case TypeOffice365ProtectionSource_KMAILBOX:
    		return "kMailbox"
        default:
        	return "kDomain"
    }
}

/**
 * Converts TypeOffice365ProtectionSourceEnum Array to its string Array representation
*/
func TypeOffice365ProtectionSourceEnumArrayToValue(typeOffice365ProtectionSourceEnum []TypeOffice365ProtectionSourceEnum) []string {
    convArray := make([]string,len( typeOffice365ProtectionSourceEnum))
    for i:=0; i<len(typeOffice365ProtectionSourceEnum);i++ {
        convArray[i] = TypeOffice365ProtectionSourceEnumToValue(typeOffice365ProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeOffice365ProtectionSourceEnumFromValue(value string) TypeOffice365ProtectionSourceEnum {
    switch value {
        case "kDomain":
            return TypeOffice365ProtectionSource_KDOMAIN
        case "kOutlook":
            return TypeOffice365ProtectionSource_KOUTLOOK
        case "kMailbox":
            return TypeOffice365ProtectionSource_KMAILBOX
        default:
            return TypeOffice365ProtectionSource_KDOMAIN
    }
}
