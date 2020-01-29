// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for OsDiskTypeEnum enum
 */
type OsDiskTypeEnum int

/**
 * Value collection for OsDiskTypeEnum enum
 */
const (
    OsDiskType_KPREMIUMSSD OsDiskTypeEnum = 1 + iota
    OsDiskType_KSTANDARDSSD
    OsDiskType_KSTANDARDHDD
)

func (r OsDiskTypeEnum) MarshalJSON() ([]byte, error) {
    s := OsDiskTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *OsDiskTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  OsDiskTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts OsDiskTypeEnum to its string representation
 */
func OsDiskTypeEnumToValue(osDiskTypeEnum OsDiskTypeEnum) string {
    switch osDiskTypeEnum {
        case OsDiskType_KPREMIUMSSD:
    		return "kPremiumSSD"
        case OsDiskType_KSTANDARDSSD:
    		return "kStandardSSD"
        case OsDiskType_KSTANDARDHDD:
    		return "kStandardHDD"
        default:
        	return "kPremiumSSD"
    }
}

/**
 * Converts OsDiskTypeEnum Array to its string Array representation
*/
func OsDiskTypeEnumArrayToValue(osDiskTypeEnum []OsDiskTypeEnum) []string {
    convArray := make([]string,len( osDiskTypeEnum))
    for i:=0; i<len(osDiskTypeEnum);i++ {
        convArray[i] = OsDiskTypeEnumToValue(osDiskTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func OsDiskTypeEnumFromValue(value string) OsDiskTypeEnum {
    switch value {
        case "kPremiumSSD":
            return OsDiskType_KPREMIUMSSD
        case "kStandardSSD":
            return OsDiskType_KSTANDARDSSD
        case "kStandardHDD":
            return OsDiskType_KSTANDARDHDD
        default:
            return OsDiskType_KPREMIUMSSD
    }
}
