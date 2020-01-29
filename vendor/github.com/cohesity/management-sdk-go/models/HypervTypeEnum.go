// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HypervTypeEnum enum
 */
type HypervTypeEnum int

/**
 * Value collection for HypervTypeEnum enum
 */
const (
    HypervType_KSCVMMSERVER HypervTypeEnum = 1 + iota
    HypervType_KSTANDALONEHOST
    HypervType_KSTANDALONECLUSTER
    HypervType_KHOSTGROUP
    HypervType_KHYPERVHOST
    HypervType_KHOSTCLUSTER
    HypervType_KVIRTUALMACHINE
    HypervType_KNETWORK
    HypervType_KDATASTORE
    HypervType_KTAG
    HypervType_KCUSTOMPROPERTY
)

func (r HypervTypeEnum) MarshalJSON() ([]byte, error) {
    s := HypervTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *HypervTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HypervTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HypervTypeEnum to its string representation
 */
func HypervTypeEnumToValue(hyperv_TypeEnum HypervTypeEnum) string {
    switch hyperv_TypeEnum {
        case HypervType_KSCVMMSERVER:
    		return "kSCVMMServer"
        case HypervType_KSTANDALONEHOST:
    		return "kStandaloneHost"
        case HypervType_KSTANDALONECLUSTER:
    		return "kStandaloneCluster"
        case HypervType_KHOSTGROUP:
    		return "kHostGroup"
        case HypervType_KHYPERVHOST:
    		return "kHypervHost"
        case HypervType_KHOSTCLUSTER:
    		return "kHostCluster"
        case HypervType_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case HypervType_KNETWORK:
    		return "kNetwork"
        case HypervType_KDATASTORE:
    		return "kDatastore"
        case HypervType_KTAG:
    		return "kTag"
        case HypervType_KCUSTOMPROPERTY:
    		return "kCustomProperty"
        default:
        	return "kSCVMMServer"
    }
}

/**
 * Converts HypervTypeEnum Array to its string Array representation
*/
func HypervTypeEnumArrayToValue(hyperv_TypeEnum []HypervTypeEnum) []string {
    convArray := make([]string,len( hyperv_TypeEnum))
    for i:=0; i<len(hyperv_TypeEnum);i++ {
        convArray[i] = HypervTypeEnumToValue(hyperv_TypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HypervTypeEnumFromValue(value string) HypervTypeEnum {
    switch value {
        case "kSCVMMServer":
            return HypervType_KSCVMMSERVER
        case "kStandaloneHost":
            return HypervType_KSTANDALONEHOST
        case "kStandaloneCluster":
            return HypervType_KSTANDALONECLUSTER
        case "kHostGroup":
            return HypervType_KHOSTGROUP
        case "kHypervHost":
            return HypervType_KHYPERVHOST
        case "kHostCluster":
            return HypervType_KHOSTCLUSTER
        case "kVirtualMachine":
            return HypervType_KVIRTUALMACHINE
        case "kNetwork":
            return HypervType_KNETWORK
        case "kDatastore":
            return HypervType_KDATASTORE
        case "kTag":
            return HypervType_KTAG
        case "kCustomProperty":
            return HypervType_KCUSTOMPROPERTY
        default:
            return HypervType_KSCVMMSERVER
    }
}
