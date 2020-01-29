// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeFileSearchResultEnum enum
 */
type TypeFileSearchResultEnum int

/**
 * Value collection for TypeFileSearchResultEnum enum
 */
const (
    TypeFileSearchResult_KDIRECTORY TypeFileSearchResultEnum = 1 + iota
    TypeFileSearchResult_KFILE
    TypeFileSearchResult_KEMAIL
    TypeFileSearchResult_KSYMLINK
)

func (r TypeFileSearchResultEnum) MarshalJSON() ([]byte, error) {
    s := TypeFileSearchResultEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeFileSearchResultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeFileSearchResultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeFileSearchResultEnum to its string representation
 */
func TypeFileSearchResultEnumToValue(typeFileSearchResultEnum TypeFileSearchResultEnum) string {
    switch typeFileSearchResultEnum {
        case TypeFileSearchResult_KDIRECTORY:
    		return "kDirectory"
        case TypeFileSearchResult_KFILE:
    		return "kFile"
        case TypeFileSearchResult_KEMAIL:
    		return "kEmail"
        case TypeFileSearchResult_KSYMLINK:
    		return "kSymlink"
        default:
        	return "kDirectory"
    }
}

/**
 * Converts TypeFileSearchResultEnum Array to its string Array representation
*/
func TypeFileSearchResultEnumArrayToValue(typeFileSearchResultEnum []TypeFileSearchResultEnum) []string {
    convArray := make([]string,len( typeFileSearchResultEnum))
    for i:=0; i<len(typeFileSearchResultEnum);i++ {
        convArray[i] = TypeFileSearchResultEnumToValue(typeFileSearchResultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeFileSearchResultEnumFromValue(value string) TypeFileSearchResultEnum {
    switch value {
        case "kDirectory":
            return TypeFileSearchResult_KDIRECTORY
        case "kFile":
            return TypeFileSearchResult_KFILE
        case "kEmail":
            return TypeFileSearchResult_KEMAIL
        case "kSymlink":
            return TypeFileSearchResult_KSYMLINK
        default:
            return TypeFileSearchResult_KDIRECTORY
    }
}
