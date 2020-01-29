// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DataDiskTypeEnum enum
 */
type DataDiskTypeEnum int

/**
 * Value collection for DataDiskTypeEnum enum
 */
const (
    DataDiskType_KPREMIUMSSD DataDiskTypeEnum = 1 + iota
    DataDiskType_KSTANDARDSSD
    DataDiskType_KSTANDARDHDD
)

func (r DataDiskTypeEnum) MarshalJSON() ([]byte, error) {
    s := DataDiskTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *DataDiskTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DataDiskTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DataDiskTypeEnum to its string representation
 */
func DataDiskTypeEnumToValue(dataDiskTypeEnum DataDiskTypeEnum) string {
    switch dataDiskTypeEnum {
        case DataDiskType_KPREMIUMSSD:
    		return "kPremiumSSD"
        case DataDiskType_KSTANDARDSSD:
    		return "kStandardSSD"
        case DataDiskType_KSTANDARDHDD:
    		return "kStandardHDD"
        default:
        	return "kPremiumSSD"
    }
}

/**
 * Converts DataDiskTypeEnum Array to its string Array representation
*/
func DataDiskTypeEnumArrayToValue(dataDiskTypeEnum []DataDiskTypeEnum) []string {
    convArray := make([]string,len( dataDiskTypeEnum))
    for i:=0; i<len(dataDiskTypeEnum);i++ {
        convArray[i] = DataDiskTypeEnumToValue(dataDiskTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DataDiskTypeEnumFromValue(value string) DataDiskTypeEnum {
    switch value {
        case "kPremiumSSD":
            return DataDiskType_KPREMIUMSSD
        case "kStandardSSD":
            return DataDiskType_KSTANDARDSSD
        case "kStandardHDD":
            return DataDiskType_KSTANDARDHDD
        default:
            return DataDiskType_KPREMIUMSSD
    }
}
