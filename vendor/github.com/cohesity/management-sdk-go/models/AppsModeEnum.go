// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AppsModeEnum enum
 */
type AppsModeEnum int

/**
 * Value collection for AppsModeEnum enum
 */
const (
    AppsMode_KDISABLED AppsModeEnum = 1 + iota
    AppsMode_KBAREMETAL
    AppsMode_KVMONLY
)

func (r AppsModeEnum) MarshalJSON() ([]byte, error) {
    s := AppsModeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AppsModeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AppsModeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AppsModeEnum to its string representation
 */
func AppsModeEnumToValue(appsModeEnum AppsModeEnum) string {
    switch appsModeEnum {
        case AppsMode_KDISABLED:
    		return "kDisabled"
        case AppsMode_KBAREMETAL:
    		return "kBareMetal"
        case AppsMode_KVMONLY:
    		return "kVmOnly"
        default:
        	return "kDisabled"
    }
}

/**
 * Converts AppsModeEnum Array to its string Array representation
*/
func AppsModeEnumArrayToValue(appsModeEnum []AppsModeEnum) []string {
    convArray := make([]string,len( appsModeEnum))
    for i:=0; i<len(appsModeEnum);i++ {
        convArray[i] = AppsModeEnumToValue(appsModeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AppsModeEnumFromValue(value string) AppsModeEnum {
    switch value {
        case "kDisabled":
            return AppsMode_KDISABLED
        case "kBareMetal":
            return AppsMode_KBAREMETAL
        case "kVmOnly":
            return AppsMode_KVMONLY
        default:
            return AppsMode_KDISABLED
    }
}
