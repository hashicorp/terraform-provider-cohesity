// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeVmwareProtectionSourceEnum enum
 */
type TypeVmwareProtectionSourceEnum int

/**
 * Value collection for TypeVmwareProtectionSourceEnum enum
 */
const (
    TypeVmwareProtectionSource_KVCENTER TypeVmwareProtectionSourceEnum = 1 + iota
    TypeVmwareProtectionSource_KFOLDER
    TypeVmwareProtectionSource_KDATACENTER
    TypeVmwareProtectionSource_KCOMPUTERESOURCE
    TypeVmwareProtectionSource_KCLUSTERCOMPUTERESOURCE
    TypeVmwareProtectionSource_KRESOURCEPOOL
    TypeVmwareProtectionSource_KDATASTORE
    TypeVmwareProtectionSource_KHOSTSYSTEM
    TypeVmwareProtectionSource_KVIRTUALMACHINE
    TypeVmwareProtectionSource_KVIRTUALAPP
    TypeVmwareProtectionSource_KSTANDALONEHOST
    TypeVmwareProtectionSource_KSTORAGEPOD
    TypeVmwareProtectionSource_KNETWORK
    TypeVmwareProtectionSource_KDISTRIBUTEDVIRTUALPORTGROUP
    TypeVmwareProtectionSource_KTAGCATEGORY
    TypeVmwareProtectionSource_KTAG
    TypeVmwareProtectionSource_KOPAQUENETWORK
    TypeVmwareProtectionSource_KVCLOUDDIRECTOR
    TypeVmwareProtectionSource_KORGANIZATION
    TypeVmwareProtectionSource_KVIRTUALDATACENTER
    TypeVmwareProtectionSource_KCATALOG
    TypeVmwareProtectionSource_KORGMETADATA
    TypeVmwareProtectionSource_KSTORAGEPOLICY
)

func (r TypeVmwareProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeVmwareProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeVmwareProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeVmwareProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeVmwareProtectionSourceEnum to its string representation
 */
func TypeVmwareProtectionSourceEnumToValue(type_vmware_ProtectionSourceEnum TypeVmwareProtectionSourceEnum) string {
    switch type_vmware_ProtectionSourceEnum {
        case TypeVmwareProtectionSource_KVCENTER:
    		return "kVCenter"
        case TypeVmwareProtectionSource_KFOLDER:
    		return "kFolder"
        case TypeVmwareProtectionSource_KDATACENTER:
    		return "kDatacenter"
        case TypeVmwareProtectionSource_KCOMPUTERESOURCE:
    		return "kComputeResource"
        case TypeVmwareProtectionSource_KCLUSTERCOMPUTERESOURCE:
    		return "kClusterComputeResource"
        case TypeVmwareProtectionSource_KRESOURCEPOOL:
    		return "kResourcePool"
        case TypeVmwareProtectionSource_KDATASTORE:
    		return "kDatastore"
        case TypeVmwareProtectionSource_KHOSTSYSTEM:
    		return "kHostSystem"
        case TypeVmwareProtectionSource_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case TypeVmwareProtectionSource_KVIRTUALAPP:
    		return "kVirtualApp"
        case TypeVmwareProtectionSource_KSTANDALONEHOST:
    		return "kStandaloneHost"
        case TypeVmwareProtectionSource_KSTORAGEPOD:
    		return "kStoragePod"
        case TypeVmwareProtectionSource_KNETWORK:
    		return "kNetwork"
        case TypeVmwareProtectionSource_KDISTRIBUTEDVIRTUALPORTGROUP:
    		return "kDistributedVirtualPortgroup"
        case TypeVmwareProtectionSource_KTAGCATEGORY:
    		return "kTagCategory"
        case TypeVmwareProtectionSource_KTAG:
    		return "kTag"
        case TypeVmwareProtectionSource_KOPAQUENETWORK:
    		return "kOpaqueNetwork"
        case TypeVmwareProtectionSource_KVCLOUDDIRECTOR:
    		return "kVCloudDirector"
        case TypeVmwareProtectionSource_KORGANIZATION:
    		return "kOrganization"
        case TypeVmwareProtectionSource_KVIRTUALDATACENTER:
    		return "kVirtualDatacenter"
        case TypeVmwareProtectionSource_KCATALOG:
    		return "kCatalog"
        case TypeVmwareProtectionSource_KORGMETADATA:
    		return "kOrgMetadata"
        case TypeVmwareProtectionSource_KSTORAGEPOLICY:
    		return "kStoragePolicy"
        default:
        	return "kVCenter"
    }
}

/**
 * Converts TypeVmwareProtectionSourceEnum Array to its string Array representation
*/
func TypeVmwareProtectionSourceEnumArrayToValue(type_vmware_ProtectionSourceEnum []TypeVmwareProtectionSourceEnum) []string {
    convArray := make([]string,len( type_vmware_ProtectionSourceEnum))
    for i:=0; i<len(type_vmware_ProtectionSourceEnum);i++ {
        convArray[i] = TypeVmwareProtectionSourceEnumToValue(type_vmware_ProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeVmwareProtectionSourceEnumFromValue(value string) TypeVmwareProtectionSourceEnum {
    switch value {
        case "kVCenter":
            return TypeVmwareProtectionSource_KVCENTER
        case "kFolder":
            return TypeVmwareProtectionSource_KFOLDER
        case "kDatacenter":
            return TypeVmwareProtectionSource_KDATACENTER
        case "kComputeResource":
            return TypeVmwareProtectionSource_KCOMPUTERESOURCE
        case "kClusterComputeResource":
            return TypeVmwareProtectionSource_KCLUSTERCOMPUTERESOURCE
        case "kResourcePool":
            return TypeVmwareProtectionSource_KRESOURCEPOOL
        case "kDatastore":
            return TypeVmwareProtectionSource_KDATASTORE
        case "kHostSystem":
            return TypeVmwareProtectionSource_KHOSTSYSTEM
        case "kVirtualMachine":
            return TypeVmwareProtectionSource_KVIRTUALMACHINE
        case "kVirtualApp":
            return TypeVmwareProtectionSource_KVIRTUALAPP
        case "kStandaloneHost":
            return TypeVmwareProtectionSource_KSTANDALONEHOST
        case "kStoragePod":
            return TypeVmwareProtectionSource_KSTORAGEPOD
        case "kNetwork":
            return TypeVmwareProtectionSource_KNETWORK
        case "kDistributedVirtualPortgroup":
            return TypeVmwareProtectionSource_KDISTRIBUTEDVIRTUALPORTGROUP
        case "kTagCategory":
            return TypeVmwareProtectionSource_KTAGCATEGORY
        case "kTag":
            return TypeVmwareProtectionSource_KTAG
        case "kOpaqueNetwork":
            return TypeVmwareProtectionSource_KOPAQUENETWORK
        case "kVCloudDirector":
            return TypeVmwareProtectionSource_KVCLOUDDIRECTOR
        case "kOrganization":
            return TypeVmwareProtectionSource_KORGANIZATION
        case "kVirtualDatacenter":
            return TypeVmwareProtectionSource_KVIRTUALDATACENTER
        case "kCatalog":
            return TypeVmwareProtectionSource_KCATALOG
        case "kOrgMetadata":
            return TypeVmwareProtectionSource_KORGMETADATA
        case "kStoragePolicy":
            return TypeVmwareProtectionSource_KSTORAGEPOLICY
        default:
            return TypeVmwareProtectionSource_KVCENTER
    }
}
