// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for FlagEnum enum
 */
type FlagEnum int

/**
 * Value collection for FlagEnum enum
 */
const (
    Flag_KERROR FlagEnum = 1 + iota
    Flag_KTRUNCATED
    Flag_KCSV
)

func (r FlagEnum) MarshalJSON() ([]byte, error) {
    s := FlagEnumToValue(r)
    return json.Marshal(s)
}

func (r *FlagEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  FlagEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts FlagEnum to its string representation
 */
func FlagEnumToValue(flagEnum FlagEnum) string {
    switch flagEnum {
        case Flag_KERROR:
    		return "kError"
        case Flag_KTRUNCATED:
    		return "kTruncated"
        case Flag_KCSV:
    		return "kCSV"
        default:
        	return "kError"
    }
}

/**
 * Converts FlagEnum Array to its string Array representation
*/
func FlagEnumArrayToValue(flagEnum []FlagEnum) []string {
    convArray := make([]string,len( flagEnum))
    for i:=0; i<len(flagEnum);i++ {
        convArray[i] = FlagEnumToValue(flagEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func FlagEnumFromValue(value string) FlagEnum {
    switch value {
        case "kError":
            return Flag_KERROR
        case "kTruncated":
            return Flag_KTRUNCATED
        case "kCSV":
            return Flag_KCSV
        default:
            return Flag_KERROR
    }
}
