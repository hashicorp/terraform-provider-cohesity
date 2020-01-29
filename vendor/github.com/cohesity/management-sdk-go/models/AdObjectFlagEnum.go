// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AdObjectFlagEnum enum
 */
type AdObjectFlagEnum int

/**
 * Value collection for AdObjectFlagEnum enum
 */
const (
    AdObjectFlag_KEQUAL AdObjectFlagEnum = 1 + iota
    AdObjectFlag_KNOTEQUAL
    AdObjectFlag_KRESTOREPASSWORDREQUIRED
    AdObjectFlag_KMOVEDONDESTINATION
    AdObjectFlag_KDESTINATIONNOTFOUND
    AdObjectFlag_KDISABLESUPPORTED
)

func (r AdObjectFlagEnum) MarshalJSON() ([]byte, error) {
    s := AdObjectFlagEnumToValue(r)
    return json.Marshal(s)
}

func (r *AdObjectFlagEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AdObjectFlagEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AdObjectFlagEnum to its string representation
 */
func AdObjectFlagEnumToValue(adObjectFlagEnum AdObjectFlagEnum) string {
    switch adObjectFlagEnum {
        case AdObjectFlag_KEQUAL:
    		return "kEqual"
        case AdObjectFlag_KNOTEQUAL:
    		return "kNotEqual"
        case AdObjectFlag_KRESTOREPASSWORDREQUIRED:
    		return "kRestorePasswordRequired"
        case AdObjectFlag_KMOVEDONDESTINATION:
    		return "kMovedOnDestination"
        case AdObjectFlag_KDESTINATIONNOTFOUND:
    		return "kDestinationNotFound"
        case AdObjectFlag_KDISABLESUPPORTED:
    		return "kDisableSupported"
        default:
        	return "kEqual"
    }
}

/**
 * Converts AdObjectFlagEnum Array to its string Array representation
*/
func AdObjectFlagEnumArrayToValue(adObjectFlagEnum []AdObjectFlagEnum) []string {
    convArray := make([]string,len( adObjectFlagEnum))
    for i:=0; i<len(adObjectFlagEnum);i++ {
        convArray[i] = AdObjectFlagEnumToValue(adObjectFlagEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AdObjectFlagEnumFromValue(value string) AdObjectFlagEnum {
    switch value {
        case "kEqual":
            return AdObjectFlag_KEQUAL
        case "kNotEqual":
            return AdObjectFlag_KNOTEQUAL
        case "kRestorePasswordRequired":
            return AdObjectFlag_KRESTOREPASSWORDREQUIRED
        case "kMovedOnDestination":
            return AdObjectFlag_KMOVEDONDESTINATION
        case "kDestinationNotFound":
            return AdObjectFlag_KDESTINATIONNOTFOUND
        case "kDisableSupported":
            return AdObjectFlag_KDISABLESUPPORTED
        default:
            return AdObjectFlag_KEQUAL
    }
}
