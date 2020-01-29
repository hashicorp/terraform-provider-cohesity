// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SearchResultFlagEnum enum
 */
type SearchResultFlagEnum int

/**
 * Value collection for SearchResultFlagEnum enum
 */
const (
    SearchResultFlag_KEQUAL SearchResultFlagEnum = 1 + iota
    SearchResultFlag_KNOTEQUAL
    SearchResultFlag_KRESTOREPASSWORDREQUIRED
    SearchResultFlag_KMOVEDONDESTINATION
    SearchResultFlag_KDISABLESUPPORTED
)

func (r SearchResultFlagEnum) MarshalJSON() ([]byte, error) {
    s := SearchResultFlagEnumToValue(r)
    return json.Marshal(s)
}

func (r *SearchResultFlagEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SearchResultFlagEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SearchResultFlagEnum to its string representation
 */
func SearchResultFlagEnumToValue(searchResultFlagEnum SearchResultFlagEnum) string {
    switch searchResultFlagEnum {
        case SearchResultFlag_KEQUAL:
    		return "kEqual"
        case SearchResultFlag_KNOTEQUAL:
    		return "kNotEqual"
        case SearchResultFlag_KRESTOREPASSWORDREQUIRED:
    		return "kRestorePasswordRequired"
        case SearchResultFlag_KMOVEDONDESTINATION:
    		return "kMovedOnDestination"
        case SearchResultFlag_KDISABLESUPPORTED:
    		return "kDisableSupported"
        default:
        	return "kEqual"
    }
}

/**
 * Converts SearchResultFlagEnum Array to its string Array representation
*/
func SearchResultFlagEnumArrayToValue(searchResultFlagEnum []SearchResultFlagEnum) []string {
    convArray := make([]string,len( searchResultFlagEnum))
    for i:=0; i<len(searchResultFlagEnum);i++ {
        convArray[i] = SearchResultFlagEnumToValue(searchResultFlagEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SearchResultFlagEnumFromValue(value string) SearchResultFlagEnum {
    switch value {
        case "kEqual":
            return SearchResultFlag_KEQUAL
        case "kNotEqual":
            return SearchResultFlag_KNOTEQUAL
        case "kRestorePasswordRequired":
            return SearchResultFlag_KRESTOREPASSWORDREQUIRED
        case "kMovedOnDestination":
            return SearchResultFlag_KMOVEDONDESTINATION
        case "kDisableSupported":
            return SearchResultFlag_KDISABLESUPPORTED
        default:
            return SearchResultFlag_KEQUAL
    }
}
