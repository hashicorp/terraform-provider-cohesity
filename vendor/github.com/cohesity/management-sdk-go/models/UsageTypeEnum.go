// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for UsageTypeEnum enum
 */
type UsageTypeEnum int

/**
 * Value collection for UsageTypeEnum enum
 */
const (
    UsageType_KARCHIVAL UsageTypeEnum = 1 + iota
    UsageType_KCLOUDSPILL
)

func (r UsageTypeEnum) MarshalJSON() ([]byte, error) {
    s := UsageTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *UsageTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  UsageTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts UsageTypeEnum to its string representation
 */
func UsageTypeEnumToValue(usageTypeEnum UsageTypeEnum) string {
    switch usageTypeEnum {
        case UsageType_KARCHIVAL:
    		return "kArchival"
        case UsageType_KCLOUDSPILL:
    		return "kCloudSpill"
        default:
        	return "kArchival"
    }
}

/**
 * Converts UsageTypeEnum Array to its string Array representation
*/
func UsageTypeEnumArrayToValue(usageTypeEnum []UsageTypeEnum) []string {
    convArray := make([]string,len( usageTypeEnum))
    for i:=0; i<len(usageTypeEnum);i++ {
        convArray[i] = UsageTypeEnumToValue(usageTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func UsageTypeEnumFromValue(value string) UsageTypeEnum {
    switch value {
        case "kArchival":
            return UsageType_KARCHIVAL
        case "kCloudSpill":
            return UsageType_KCLOUDSPILL
        default:
            return UsageType_KARCHIVAL
    }
}
