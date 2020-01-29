// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PureTypeEnum enum
 */
type PureTypeEnum int

/**
 * Value collection for PureTypeEnum enum
 */
const (
    PureType_KSTORAGEARRAY PureTypeEnum = 1 + iota
    PureType_KVOLUME
)

func (r PureTypeEnum) MarshalJSON() ([]byte, error) {
    s := PureTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *PureTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PureTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PureTypeEnum to its string representation
 */
func PureTypeEnumToValue(pureTypeEnum PureTypeEnum) string {
    switch pureTypeEnum {
        case PureType_KSTORAGEARRAY:
    		return "kStorageArray"
        case PureType_KVOLUME:
    		return "kVolume"
        default:
        	return "kStorageArray"
    }
}

/**
 * Converts PureTypeEnum Array to its string Array representation
*/
func PureTypeEnumArrayToValue(pureTypeEnum []PureTypeEnum) []string {
    convArray := make([]string,len( pureTypeEnum))
    for i:=0; i<len(pureTypeEnum);i++ {
        convArray[i] = PureTypeEnumToValue(pureTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PureTypeEnumFromValue(value string) PureTypeEnum {
    switch value {
        case "kStorageArray":
            return PureType_KSTORAGEARRAY
        case "kVolume":
            return PureType_KVOLUME
        default:
            return PureType_KSTORAGEARRAY
    }
}
