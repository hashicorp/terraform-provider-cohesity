// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PkgTypeEnum enum
 */
type PkgTypeEnum int

/**
 * Value collection for PkgTypeEnum enum
 */
const (
    PkgType_KSCRIPT PkgTypeEnum = 1 + iota
    PkgType_KRPM
    PkgType_KSUSERPM
    PkgType_KDEB
)

func (r PkgTypeEnum) MarshalJSON() ([]byte, error) {
    s := PkgTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *PkgTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PkgTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PkgTypeEnum to its string representation
 */
func PkgTypeEnumToValue(pkgTypeEnum PkgTypeEnum) string {
    switch pkgTypeEnum {
        case PkgType_KSCRIPT:
    		return "kScript"
        case PkgType_KRPM:
    		return "kRPM"
        case PkgType_KSUSERPM:
    		return "kSuseRPM"
        case PkgType_KDEB:
    		return "kDEB"
        default:
        	return "kScript"
    }
}

/**
 * Converts PkgTypeEnum Array to its string Array representation
*/
func PkgTypeEnumArrayToValue(pkgTypeEnum []PkgTypeEnum) []string {
    convArray := make([]string,len( pkgTypeEnum))
    for i:=0; i<len(pkgTypeEnum);i++ {
        convArray[i] = PkgTypeEnumToValue(pkgTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PkgTypeEnumFromValue(value string) PkgTypeEnum {
    switch value {
        case "kScript":
            return PkgType_KSCRIPT
        case "kRPM":
            return PkgType_KRPM
        case "kSuseRPM":
            return PkgType_KSUSERPM
        case "kDEB":
            return PkgType_KDEB
        default:
            return PkgType_KSCRIPT
    }
}
