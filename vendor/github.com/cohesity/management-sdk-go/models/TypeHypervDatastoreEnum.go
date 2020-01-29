// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeHypervDatastoreEnum enum
 */
type TypeHypervDatastoreEnum int

/**
 * Value collection for TypeHypervDatastoreEnum enum
 */
const (
    TypeHypervDatastore_KFILESHARE TypeHypervDatastoreEnum = 1 + iota
    TypeHypervDatastore_KVOLUME
)

func (r TypeHypervDatastoreEnum) MarshalJSON() ([]byte, error) {
    s := TypeHypervDatastoreEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeHypervDatastoreEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeHypervDatastoreEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeHypervDatastoreEnum to its string representation
 */
func TypeHypervDatastoreEnumToValue(type_hyperv_DatastoreEnum TypeHypervDatastoreEnum) string {
    switch type_hyperv_DatastoreEnum {
        case TypeHypervDatastore_KFILESHARE:
    		return "kFileShare"
        case TypeHypervDatastore_KVOLUME:
    		return "kVolume"
        default:
        	return "kFileShare"
    }
}

/**
 * Converts TypeHypervDatastoreEnum Array to its string Array representation
*/
func TypeHypervDatastoreEnumArrayToValue(type_hyperv_DatastoreEnum []TypeHypervDatastoreEnum) []string {
    convArray := make([]string,len( type_hyperv_DatastoreEnum))
    for i:=0; i<len(type_hyperv_DatastoreEnum);i++ {
        convArray[i] = TypeHypervDatastoreEnumToValue(type_hyperv_DatastoreEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeHypervDatastoreEnumFromValue(value string) TypeHypervDatastoreEnum {
    switch value {
        case "kFileShare":
            return TypeHypervDatastore_KFILESHARE
        case "kVolume":
            return TypeHypervDatastore_KVOLUME
        default:
            return TypeHypervDatastore_KFILESHARE
    }
}
