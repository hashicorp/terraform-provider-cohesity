// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeCloudDeployTargetDetailsEnum enum
 */
type TypeCloudDeployTargetDetailsEnum int

/**
 * Value collection for TypeCloudDeployTargetDetailsEnum enum
 */
const (
    TypeCloudDeployTargetDetails_KAZURE TypeCloudDeployTargetDetailsEnum = 1 + iota
    TypeCloudDeployTargetDetails_KAWS
    TypeCloudDeployTargetDetails_KGCP
)

func (r TypeCloudDeployTargetDetailsEnum) MarshalJSON() ([]byte, error) {
    s := TypeCloudDeployTargetDetailsEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeCloudDeployTargetDetailsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeCloudDeployTargetDetailsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeCloudDeployTargetDetailsEnum to its string representation
 */
func TypeCloudDeployTargetDetailsEnumToValue(typeCloudDeployTargetDetailsEnum TypeCloudDeployTargetDetailsEnum) string {
    switch typeCloudDeployTargetDetailsEnum {
        case TypeCloudDeployTargetDetails_KAZURE:
    		return "kAzure"
        case TypeCloudDeployTargetDetails_KAWS:
    		return "kAws"
        case TypeCloudDeployTargetDetails_KGCP:
    		return "kGcp"
        default:
        	return "kAzure"
    }
}

/**
 * Converts TypeCloudDeployTargetDetailsEnum Array to its string Array representation
*/
func TypeCloudDeployTargetDetailsEnumArrayToValue(typeCloudDeployTargetDetailsEnum []TypeCloudDeployTargetDetailsEnum) []string {
    convArray := make([]string,len( typeCloudDeployTargetDetailsEnum))
    for i:=0; i<len(typeCloudDeployTargetDetailsEnum);i++ {
        convArray[i] = TypeCloudDeployTargetDetailsEnumToValue(typeCloudDeployTargetDetailsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeCloudDeployTargetDetailsEnumFromValue(value string) TypeCloudDeployTargetDetailsEnum {
    switch value {
        case "kAzure":
            return TypeCloudDeployTargetDetails_KAZURE
        case "kAws":
            return TypeCloudDeployTargetDetails_KAWS
        case "kGcp":
            return TypeCloudDeployTargetDetails_KGCP
        default:
            return TypeCloudDeployTargetDetails_KAZURE
    }
}
