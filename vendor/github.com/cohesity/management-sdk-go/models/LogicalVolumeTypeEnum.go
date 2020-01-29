// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for LogicalVolumeTypeEnum enum
 */
type LogicalVolumeTypeEnum int

/**
 * Value collection for LogicalVolumeTypeEnum enum
 */
const (
    LogicalVolumeType_KSIMPLEVOLUME LogicalVolumeTypeEnum = 1 + iota
    LogicalVolumeType_KLVM
    LogicalVolumeType_KLDM
)

func (r LogicalVolumeTypeEnum) MarshalJSON() ([]byte, error) {
    s := LogicalVolumeTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *LogicalVolumeTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  LogicalVolumeTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts LogicalVolumeTypeEnum to its string representation
 */
func LogicalVolumeTypeEnumToValue(logicalVolumeTypeEnum LogicalVolumeTypeEnum) string {
    switch logicalVolumeTypeEnum {
        case LogicalVolumeType_KSIMPLEVOLUME:
    		return "kSimpleVolume"
        case LogicalVolumeType_KLVM:
    		return "kLVM"
        case LogicalVolumeType_KLDM:
    		return "kLDM"
        default:
        	return "kSimpleVolume"
    }
}

/**
 * Converts LogicalVolumeTypeEnum Array to its string Array representation
*/
func LogicalVolumeTypeEnumArrayToValue(logicalVolumeTypeEnum []LogicalVolumeTypeEnum) []string {
    convArray := make([]string,len( logicalVolumeTypeEnum))
    for i:=0; i<len(logicalVolumeTypeEnum);i++ {
        convArray[i] = LogicalVolumeTypeEnumToValue(logicalVolumeTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func LogicalVolumeTypeEnumFromValue(value string) LogicalVolumeTypeEnum {
    switch value {
        case "kSimpleVolume":
            return LogicalVolumeType_KSIMPLEVOLUME
        case "kLVM":
            return LogicalVolumeType_KLVM
        case "kLDM":
            return LogicalVolumeType_KLDM
        default:
            return LogicalVolumeType_KSIMPLEVOLUME
    }
}
