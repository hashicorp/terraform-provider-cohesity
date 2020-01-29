// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for KubernetesTypeEnum enum
 */
type KubernetesTypeEnum int

/**
 * Value collection for KubernetesTypeEnum enum
 */
const (
    KubernetesType_KCLUSTER KubernetesTypeEnum = 1 + iota
    KubernetesType_KNAMESPACE
    KubernetesType_KSERVICE
)

func (r KubernetesTypeEnum) MarshalJSON() ([]byte, error) {
    s := KubernetesTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *KubernetesTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  KubernetesTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts KubernetesTypeEnum to its string representation
 */
func KubernetesTypeEnumToValue(kubernetesTypeEnum KubernetesTypeEnum) string {
    switch kubernetesTypeEnum {
        case KubernetesType_KCLUSTER:
    		return "kCluster"
        case KubernetesType_KNAMESPACE:
    		return "kNamespace"
        case KubernetesType_KSERVICE:
    		return "kService"
        default:
        	return "kCluster"
    }
}

/**
 * Converts KubernetesTypeEnum Array to its string Array representation
*/
func KubernetesTypeEnumArrayToValue(kubernetesTypeEnum []KubernetesTypeEnum) []string {
    convArray := make([]string,len( kubernetesTypeEnum))
    for i:=0; i<len(kubernetesTypeEnum);i++ {
        convArray[i] = KubernetesTypeEnumToValue(kubernetesTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func KubernetesTypeEnumFromValue(value string) KubernetesTypeEnum {
    switch value {
        case "kCluster":
            return KubernetesType_KCLUSTER
        case "kNamespace":
            return KubernetesType_KNAMESPACE
        case "kService":
            return KubernetesType_KSERVICE
        default:
            return KubernetesType_KCLUSTER
    }
}
