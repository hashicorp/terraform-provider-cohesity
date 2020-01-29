// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ExcludeTypeEnum enum
 */
type ExcludeTypeEnum int

/**
 * Value collection for ExcludeTypeEnum enum
 */
const (
    ExcludeType_KVCENTER ExcludeTypeEnum = 1 + iota
    ExcludeType_KFOLDER
    ExcludeType_KDATACENTER
    ExcludeType_KCOMPUTERESOURCE
    ExcludeType_KCLUSTERCOMPUTERESOURCE
    ExcludeType_KRESOURCEPOOL
    ExcludeType_KDATASTORE
    ExcludeType_KHOSTSYSTEM
    ExcludeType_KVIRTUALMACHINE
    ExcludeType_KVIRTUALAPP
    ExcludeType_KSTANDALONEHOST
    ExcludeType_KSTORAGEPOD
    ExcludeType_KNETWORK
    ExcludeType_KDISTRIBUTEDVIRTUALPORTGROUP
    ExcludeType_KTAGCATEGORY
    ExcludeType_KTAG
)

func (r ExcludeTypeEnum) MarshalJSON() ([]byte, error) {
    s := ExcludeTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ExcludeTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ExcludeTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ExcludeTypeEnum to its string representation
 */
func ExcludeTypeEnumToValue(excludeTypeEnum ExcludeTypeEnum) string {
    switch excludeTypeEnum {
        case ExcludeType_KVCENTER:
    		return "kVCenter"
        case ExcludeType_KFOLDER:
    		return "kFolder"
        case ExcludeType_KDATACENTER:
    		return "kDatacenter"
        case ExcludeType_KCOMPUTERESOURCE:
    		return "kComputeResource"
        case ExcludeType_KCLUSTERCOMPUTERESOURCE:
    		return "kClusterComputeResource"
        case ExcludeType_KRESOURCEPOOL:
    		return "kResourcePool"
        case ExcludeType_KDATASTORE:
    		return "kDatastore"
        case ExcludeType_KHOSTSYSTEM:
    		return "kHostSystem"
        case ExcludeType_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case ExcludeType_KVIRTUALAPP:
    		return "kVirtualApp"
        case ExcludeType_KSTANDALONEHOST:
    		return "kStandaloneHost"
        case ExcludeType_KSTORAGEPOD:
    		return "kStoragePod"
        case ExcludeType_KNETWORK:
    		return "kNetwork"
        case ExcludeType_KDISTRIBUTEDVIRTUALPORTGROUP:
    		return "kDistributedVirtualPortgroup"
        case ExcludeType_KTAGCATEGORY:
    		return "kTagCategory"
        case ExcludeType_KTAG:
    		return "kTag"
        default:
        	return "kVCenter"
    }
}

/**
 * Converts ExcludeTypeEnum Array to its string Array representation
*/
func ExcludeTypeEnumArrayToValue(excludeTypeEnum []ExcludeTypeEnum) []string {
    convArray := make([]string,len( excludeTypeEnum))
    for i:=0; i<len(excludeTypeEnum);i++ {
        convArray[i] = ExcludeTypeEnumToValue(excludeTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExcludeTypeEnumFromValue(value string) ExcludeTypeEnum {
    switch value {
        case "kVCenter":
            return ExcludeType_KVCENTER
        case "kFolder":
            return ExcludeType_KFOLDER
        case "kDatacenter":
            return ExcludeType_KDATACENTER
        case "kComputeResource":
            return ExcludeType_KCOMPUTERESOURCE
        case "kClusterComputeResource":
            return ExcludeType_KCLUSTERCOMPUTERESOURCE
        case "kResourcePool":
            return ExcludeType_KRESOURCEPOOL
        case "kDatastore":
            return ExcludeType_KDATASTORE
        case "kHostSystem":
            return ExcludeType_KHOSTSYSTEM
        case "kVirtualMachine":
            return ExcludeType_KVIRTUALMACHINE
        case "kVirtualApp":
            return ExcludeType_KVIRTUALAPP
        case "kStandaloneHost":
            return ExcludeType_KSTANDALONEHOST
        case "kStoragePod":
            return ExcludeType_KSTORAGEPOD
        case "kNetwork":
            return ExcludeType_KNETWORK
        case "kDistributedVirtualPortgroup":
            return ExcludeType_KDISTRIBUTEDVIRTUALPORTGROUP
        case "kTagCategory":
            return ExcludeType_KTAGCATEGORY
        case "kTag":
            return ExcludeType_KTAG
        default:
            return ExcludeType_KVCENTER
    }
}
