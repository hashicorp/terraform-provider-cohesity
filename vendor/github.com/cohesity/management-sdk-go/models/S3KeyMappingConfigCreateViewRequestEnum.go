// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for S3KeyMappingConfigCreateViewRequestEnum enum
 */
type S3KeyMappingConfigCreateViewRequestEnum int

/**
 * Value collection for S3KeyMappingConfigCreateViewRequestEnum enum
 */
const (
    S3KeyMappingConfigCreateViewRequest_KRANDOM S3KeyMappingConfigCreateViewRequestEnum = 1 + iota
    S3KeyMappingConfigCreateViewRequest_KSHORT
    S3KeyMappingConfigCreateViewRequest_KLONG
    S3KeyMappingConfigCreateViewRequest_KHIERARCHICAL
)

func (r S3KeyMappingConfigCreateViewRequestEnum) MarshalJSON() ([]byte, error) {
    s := S3KeyMappingConfigCreateViewRequestEnumToValue(r)
    return json.Marshal(s)
}

func (r *S3KeyMappingConfigCreateViewRequestEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  S3KeyMappingConfigCreateViewRequestEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts S3KeyMappingConfigCreateViewRequestEnum to its string representation
 */
func S3KeyMappingConfigCreateViewRequestEnumToValue(s3KeyMappingConfigCreateViewRequestEnum S3KeyMappingConfigCreateViewRequestEnum) string {
    switch s3KeyMappingConfigCreateViewRequestEnum {
        case S3KeyMappingConfigCreateViewRequest_KRANDOM:
    		return "kRandom"
        case S3KeyMappingConfigCreateViewRequest_KSHORT:
    		return "kShort"
        case S3KeyMappingConfigCreateViewRequest_KLONG:
    		return "kLong"
        case S3KeyMappingConfigCreateViewRequest_KHIERARCHICAL:
    		return "kHierarchical"
        default:
        	return "kRandom"
    }
}

/**
 * Converts S3KeyMappingConfigCreateViewRequestEnum Array to its string Array representation
*/
func S3KeyMappingConfigCreateViewRequestEnumArrayToValue(s3KeyMappingConfigCreateViewRequestEnum []S3KeyMappingConfigCreateViewRequestEnum) []string {
    convArray := make([]string,len( s3KeyMappingConfigCreateViewRequestEnum))
    for i:=0; i<len(s3KeyMappingConfigCreateViewRequestEnum);i++ {
        convArray[i] = S3KeyMappingConfigCreateViewRequestEnumToValue(s3KeyMappingConfigCreateViewRequestEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func S3KeyMappingConfigCreateViewRequestEnumFromValue(value string) S3KeyMappingConfigCreateViewRequestEnum {
    switch value {
        case "kRandom":
            return S3KeyMappingConfigCreateViewRequest_KRANDOM
        case "kShort":
            return S3KeyMappingConfigCreateViewRequest_KSHORT
        case "kLong":
            return S3KeyMappingConfigCreateViewRequest_KLONG
        case "kHierarchical":
            return S3KeyMappingConfigCreateViewRequest_KHIERARCHICAL
        default:
            return S3KeyMappingConfigCreateViewRequest_KRANDOM
    }
}
