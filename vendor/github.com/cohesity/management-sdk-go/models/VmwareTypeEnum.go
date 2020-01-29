// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VmwareTypeEnum enum
 */
type VmwareTypeEnum int

/**
 * Value collection for VmwareTypeEnum enum
 */
const (
    VmwareType_KVCENTER VmwareTypeEnum = 1 + iota
    VmwareType_KFOLDER
    VmwareType_KDATACENTER
    VmwareType_KCOMPUTERESOURCE
    VmwareType_KCLUSTERCOMPUTERESOURCE
    VmwareType_KRESOURCEPOOL
    VmwareType_KDATASTORE
    VmwareType_KHOSTSYSTEM
    VmwareType_KVIRTUALMACHINE
    VmwareType_KVIRTUALAPP
    VmwareType_KSTANDALONEHOST
    VmwareType_KSTORAGEPOD
    VmwareType_KNETWORK
    VmwareType_KDISTRIBUTEDVIRTUALPORTGROUP
    VmwareType_KTAGCATEGORY
    VmwareType_KTAG
    VmwareType_KOPAQUENETWORK
    VmwareType_KVCLOUDDIRECTOR
    VmwareType_KORGANIZATION
    VmwareType_KVIRTUALDATACENTER
    VmwareType_KCATALOG
    VmwareType_KORGMETADATA
    VmwareType_KSTORAGEPOLICY
)

func (r VmwareTypeEnum) MarshalJSON() ([]byte, error) {
    s := VmwareTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *VmwareTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VmwareTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VmwareTypeEnum to its string representation
 */
func VmwareTypeEnumToValue(vmware_TypeEnum VmwareTypeEnum) string {
    switch vmware_TypeEnum {
        case VmwareType_KVCENTER:
    		return "kVCenter"
        case VmwareType_KFOLDER:
    		return "kFolder"
        case VmwareType_KDATACENTER:
    		return "kDatacenter"
        case VmwareType_KCOMPUTERESOURCE:
    		return "kComputeResource"
        case VmwareType_KCLUSTERCOMPUTERESOURCE:
    		return "kClusterComputeResource"
        case VmwareType_KRESOURCEPOOL:
    		return "kResourcePool"
        case VmwareType_KDATASTORE:
    		return "kDatastore"
        case VmwareType_KHOSTSYSTEM:
    		return "kHostSystem"
        case VmwareType_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case VmwareType_KVIRTUALAPP:
    		return "kVirtualApp"
        case VmwareType_KSTANDALONEHOST:
    		return "kStandaloneHost"
        case VmwareType_KSTORAGEPOD:
    		return "kStoragePod"
        case VmwareType_KNETWORK:
    		return "kNetwork"
        case VmwareType_KDISTRIBUTEDVIRTUALPORTGROUP:
    		return "kDistributedVirtualPortgroup"
        case VmwareType_KTAGCATEGORY:
    		return "kTagCategory"
        case VmwareType_KTAG:
    		return "kTag"
        case VmwareType_KOPAQUENETWORK:
    		return "kOpaqueNetwork"
        case VmwareType_KVCLOUDDIRECTOR:
    		return "kVCloudDirector"
        case VmwareType_KORGANIZATION:
    		return "kOrganization"
        case VmwareType_KVIRTUALDATACENTER:
    		return "kVirtualDatacenter"
        case VmwareType_KCATALOG:
    		return "kCatalog"
        case VmwareType_KORGMETADATA:
    		return "kOrgMetadata"
        case VmwareType_KSTORAGEPOLICY:
    		return "kStoragePolicy"
        default:
        	return "kVCenter"
    }
}

/**
 * Converts VmwareTypeEnum Array to its string Array representation
*/
func VmwareTypeEnumArrayToValue(vmware_TypeEnum []VmwareTypeEnum) []string {
    convArray := make([]string,len( vmware_TypeEnum))
    for i:=0; i<len(vmware_TypeEnum);i++ {
        convArray[i] = VmwareTypeEnumToValue(vmware_TypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VmwareTypeEnumFromValue(value string) VmwareTypeEnum {
    switch value {
        case "kVCenter":
            return VmwareType_KVCENTER
        case "kFolder":
            return VmwareType_KFOLDER
        case "kDatacenter":
            return VmwareType_KDATACENTER
        case "kComputeResource":
            return VmwareType_KCOMPUTERESOURCE
        case "kClusterComputeResource":
            return VmwareType_KCLUSTERCOMPUTERESOURCE
        case "kResourcePool":
            return VmwareType_KRESOURCEPOOL
        case "kDatastore":
            return VmwareType_KDATASTORE
        case "kHostSystem":
            return VmwareType_KHOSTSYSTEM
        case "kVirtualMachine":
            return VmwareType_KVIRTUALMACHINE
        case "kVirtualApp":
            return VmwareType_KVIRTUALAPP
        case "kStandaloneHost":
            return VmwareType_KSTANDALONEHOST
        case "kStoragePod":
            return VmwareType_KSTORAGEPOD
        case "kNetwork":
            return VmwareType_KNETWORK
        case "kDistributedVirtualPortgroup":
            return VmwareType_KDISTRIBUTEDVIRTUALPORTGROUP
        case "kTagCategory":
            return VmwareType_KTAGCATEGORY
        case "kTag":
            return VmwareType_KTAG
        case "kOpaqueNetwork":
            return VmwareType_KOPAQUENETWORK
        case "kVCloudDirector":
            return VmwareType_KVCLOUDDIRECTOR
        case "kOrganization":
            return VmwareType_KORGANIZATION
        case "kVirtualDatacenter":
            return VmwareType_KVIRTUALDATACENTER
        case "kCatalog":
            return VmwareType_KCATALOG
        case "kOrgMetadata":
            return VmwareType_KORGMETADATA
        case "kStoragePolicy":
            return VmwareType_KSTORAGEPOLICY
        default:
            return VmwareType_KVCENTER
    }
}
