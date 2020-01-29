// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for S3KeyMappingConfigEnum enum
 */
type S3KeyMappingConfigEnum int

/**
 * Value collection for S3KeyMappingConfigEnum enum
 */
const (
    S3KeyMappingConfig_KRANDOM S3KeyMappingConfigEnum = 1 + iota
    S3KeyMappingConfig_KSHORT
    S3KeyMappingConfig_KLONG
    S3KeyMappingConfig_KHIERARCHICAL
)

func (r S3KeyMappingConfigEnum) MarshalJSON() ([]byte, error) {
    s := S3KeyMappingConfigEnumToValue(r)
    return json.Marshal(s)
}

func (r *S3KeyMappingConfigEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  S3KeyMappingConfigEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts S3KeyMappingConfigEnum to its string representation
 */
func S3KeyMappingConfigEnumToValue(s3KeyMappingConfigEnum S3KeyMappingConfigEnum) string {
    switch s3KeyMappingConfigEnum {
        case S3KeyMappingConfig_KRANDOM:
    		return "kRandom"
        case S3KeyMappingConfig_KSHORT:
    		return "kShort"
        case S3KeyMappingConfig_KLONG:
    		return "kLong"
        case S3KeyMappingConfig_KHIERARCHICAL:
    		return "kHierarchical"
        default:
        	return "kRandom"
    }
}

/**
 * Converts S3KeyMappingConfigEnum Array to its string Array representation
*/
func S3KeyMappingConfigEnumArrayToValue(s3KeyMappingConfigEnum []S3KeyMappingConfigEnum) []string {
    convArray := make([]string,len( s3KeyMappingConfigEnum))
    for i:=0; i<len(s3KeyMappingConfigEnum);i++ {
        convArray[i] = S3KeyMappingConfigEnumToValue(s3KeyMappingConfigEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func S3KeyMappingConfigEnumFromValue(value string) S3KeyMappingConfigEnum {
    switch value {
        case "kRandom":
            return S3KeyMappingConfig_KRANDOM
        case "kShort":
            return S3KeyMappingConfig_KSHORT
        case "kLong":
            return S3KeyMappingConfig_KLONG
        case "kHierarchical":
            return S3KeyMappingConfig_KHIERARCHICAL
        default:
            return S3KeyMappingConfig_KRANDOM
    }
}
