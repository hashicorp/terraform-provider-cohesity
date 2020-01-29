// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CombineMethodEnum enum
 */
type CombineMethodEnum int

/**
 * Value collection for CombineMethodEnum enum
 */
const (
    CombineMethod_LINEAR CombineMethodEnum = 1 + iota
    CombineMethod_STRIPE
    CombineMethod_MIRROR
    CombineMethod_RAID5
    CombineMethod_RAID6
    CombineMethod_ZERO
    CombineMethod_THIN
    CombineMethod_THINPOOL
    CombineMethod_SNAPSHOT
    CombineMethod_CACHE
    CombineMethod_CACHEPOOL
)

func (r CombineMethodEnum) MarshalJSON() ([]byte, error) {
    s := CombineMethodEnumToValue(r)
    return json.Marshal(s)
}

func (r *CombineMethodEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CombineMethodEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CombineMethodEnum to its string representation
 */
func CombineMethodEnumToValue(combineMethodEnum CombineMethodEnum) string {
    switch combineMethodEnum {
        case CombineMethod_LINEAR:
    		return "LINEAR"
        case CombineMethod_STRIPE:
    		return "STRIPE"
        case CombineMethod_MIRROR:
    		return "MIRROR"
        case CombineMethod_RAID5:
    		return "RAID5"
        case CombineMethod_RAID6:
    		return "RAID6"
        case CombineMethod_ZERO:
    		return "ZERO"
        case CombineMethod_THIN:
    		return "THIN"
        case CombineMethod_THINPOOL:
    		return "THINPOOL"
        case CombineMethod_SNAPSHOT:
    		return "SNAPSHOT"
        case CombineMethod_CACHE:
    		return "CACHE"
        case CombineMethod_CACHEPOOL:
    		return "CACHEPOOL"
        default:
        	return "LINEAR"
    }
}

/**
 * Converts CombineMethodEnum Array to its string Array representation
*/
func CombineMethodEnumArrayToValue(combineMethodEnum []CombineMethodEnum) []string {
    convArray := make([]string,len( combineMethodEnum))
    for i:=0; i<len(combineMethodEnum);i++ {
        convArray[i] = CombineMethodEnumToValue(combineMethodEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CombineMethodEnumFromValue(value string) CombineMethodEnum {
    switch value {
        case "LINEAR":
            return CombineMethod_LINEAR
        case "STRIPE":
            return CombineMethod_STRIPE
        case "MIRROR":
            return CombineMethod_MIRROR
        case "RAID5":
            return CombineMethod_RAID5
        case "RAID6":
            return CombineMethod_RAID6
        case "ZERO":
            return CombineMethod_ZERO
        case "THIN":
            return CombineMethod_THIN
        case "THINPOOL":
            return CombineMethod_THINPOOL
        case "SNAPSHOT":
            return CombineMethod_SNAPSHOT
        case "CACHE":
            return CombineMethod_CACHE
        case "CACHEPOOL":
            return CombineMethod_CACHEPOOL
        default:
            return CombineMethod_LINEAR
    }
}
