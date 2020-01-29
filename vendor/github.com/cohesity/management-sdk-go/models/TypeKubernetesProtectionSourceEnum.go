// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeKubernetesProtectionSourceEnum enum
 */
type TypeKubernetesProtectionSourceEnum int

/**
 * Value collection for TypeKubernetesProtectionSourceEnum enum
 */
const (
    TypeKubernetesProtectionSource_KCLUSTER TypeKubernetesProtectionSourceEnum = 1 + iota
    TypeKubernetesProtectionSource_KNAMESPACE
    TypeKubernetesProtectionSource_KSERVICE
)

func (r TypeKubernetesProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeKubernetesProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeKubernetesProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeKubernetesProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeKubernetesProtectionSourceEnum to its string representation
 */
func TypeKubernetesProtectionSourceEnumToValue(typeKubernetesProtectionSourceEnum TypeKubernetesProtectionSourceEnum) string {
    switch typeKubernetesProtectionSourceEnum {
        case TypeKubernetesProtectionSource_KCLUSTER:
    		return "kCluster"
        case TypeKubernetesProtectionSource_KNAMESPACE:
    		return "kNamespace"
        case TypeKubernetesProtectionSource_KSERVICE:
    		return "kService"
        default:
        	return "kCluster"
    }
}

/**
 * Converts TypeKubernetesProtectionSourceEnum Array to its string Array representation
*/
func TypeKubernetesProtectionSourceEnumArrayToValue(typeKubernetesProtectionSourceEnum []TypeKubernetesProtectionSourceEnum) []string {
    convArray := make([]string,len( typeKubernetesProtectionSourceEnum))
    for i:=0; i<len(typeKubernetesProtectionSourceEnum);i++ {
        convArray[i] = TypeKubernetesProtectionSourceEnumToValue(typeKubernetesProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeKubernetesProtectionSourceEnumFromValue(value string) TypeKubernetesProtectionSourceEnum {
    switch value {
        case "kCluster":
            return TypeKubernetesProtectionSource_KCLUSTER
        case "kNamespace":
            return TypeKubernetesProtectionSource_KNAMESPACE
        case "kService":
            return TypeKubernetesProtectionSource_KSERVICE
        default:
            return TypeKubernetesProtectionSource_KCLUSTER
    }
}
