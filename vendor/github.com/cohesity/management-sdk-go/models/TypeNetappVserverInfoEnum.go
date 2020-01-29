// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeNetappVserverInfoEnum enum
 */
type TypeNetappVserverInfoEnum int

/**
 * Value collection for TypeNetappVserverInfoEnum enum
 */
const (
    TypeNetappVserverInfo_KDATA TypeNetappVserverInfoEnum = 1 + iota
    TypeNetappVserverInfo_KADMIN
    TypeNetappVserverInfo_KSYSTEM
    TypeNetappVserverInfo_KNODE
    TypeNetappVserverInfo_KUNKNOWN
)

func (r TypeNetappVserverInfoEnum) MarshalJSON() ([]byte, error) {
    s := TypeNetappVserverInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeNetappVserverInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeNetappVserverInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeNetappVserverInfoEnum to its string representation
 */
func TypeNetappVserverInfoEnumToValue(typeNetappVserverInfoEnum TypeNetappVserverInfoEnum) string {
    switch typeNetappVserverInfoEnum {
        case TypeNetappVserverInfo_KDATA:
    		return "kData"
        case TypeNetappVserverInfo_KADMIN:
    		return "kAdmin"
        case TypeNetappVserverInfo_KSYSTEM:
    		return "kSystem"
        case TypeNetappVserverInfo_KNODE:
    		return "kNode"
        case TypeNetappVserverInfo_KUNKNOWN:
    		return "kUnknown"
        default:
        	return "kData"
    }
}

/**
 * Converts TypeNetappVserverInfoEnum Array to its string Array representation
*/
func TypeNetappVserverInfoEnumArrayToValue(typeNetappVserverInfoEnum []TypeNetappVserverInfoEnum) []string {
    convArray := make([]string,len( typeNetappVserverInfoEnum))
    for i:=0; i<len(typeNetappVserverInfoEnum);i++ {
        convArray[i] = TypeNetappVserverInfoEnumToValue(typeNetappVserverInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeNetappVserverInfoEnumFromValue(value string) TypeNetappVserverInfoEnum {
    switch value {
        case "kData":
            return TypeNetappVserverInfo_KDATA
        case "kAdmin":
            return TypeNetappVserverInfo_KADMIN
        case "kSystem":
            return TypeNetappVserverInfo_KSYSTEM
        case "kNode":
            return TypeNetappVserverInfo_KNODE
        case "kUnknown":
            return TypeNetappVserverInfo_KUNKNOWN
        default:
            return TypeNetappVserverInfo_KDATA
    }
}
