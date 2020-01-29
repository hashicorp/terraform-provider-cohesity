// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeRemoteHostEnum enum
 */
type TypeRemoteHostEnum int

/**
 * Value collection for TypeRemoteHostEnum enum
 */
const (
    TypeRemoteHost_KLINUX TypeRemoteHostEnum = 1 + iota
    TypeRemoteHost_KWINDOWS
    TypeRemoteHost_KAIX
    TypeRemoteHost_KSOLARIS
    TypeRemoteHost_KSAPHANA
    TypeRemoteHost_KOTHER
)

func (r TypeRemoteHostEnum) MarshalJSON() ([]byte, error) {
    s := TypeRemoteHostEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeRemoteHostEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeRemoteHostEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeRemoteHostEnum to its string representation
 */
func TypeRemoteHostEnumToValue(typeRemoteHostEnum TypeRemoteHostEnum) string {
    switch typeRemoteHostEnum {
        case TypeRemoteHost_KLINUX:
    		return "kLinux"
        case TypeRemoteHost_KWINDOWS:
    		return "kWindows"
        case TypeRemoteHost_KAIX:
    		return "kAix"
        case TypeRemoteHost_KSOLARIS:
    		return "kSolaris"
        case TypeRemoteHost_KSAPHANA:
    		return "kSapHana"
        case TypeRemoteHost_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts TypeRemoteHostEnum Array to its string Array representation
*/
func TypeRemoteHostEnumArrayToValue(typeRemoteHostEnum []TypeRemoteHostEnum) []string {
    convArray := make([]string,len( typeRemoteHostEnum))
    for i:=0; i<len(typeRemoteHostEnum);i++ {
        convArray[i] = TypeRemoteHostEnumToValue(typeRemoteHostEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeRemoteHostEnumFromValue(value string) TypeRemoteHostEnum {
    switch value {
        case "kLinux":
            return TypeRemoteHost_KLINUX
        case "kWindows":
            return TypeRemoteHost_KWINDOWS
        case "kAix":
            return TypeRemoteHost_KAIX
        case "kSolaris":
            return TypeRemoteHost_KSOLARIS
        case "kSapHana":
            return TypeRemoteHost_KSAPHANA
        case "kOther":
            return TypeRemoteHost_KOTHER
        default:
            return TypeRemoteHost_KLINUX
    }
}
