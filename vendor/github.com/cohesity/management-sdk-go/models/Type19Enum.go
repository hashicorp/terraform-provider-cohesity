package models

import(
    "encoding/json"
)

/**
 * Type definition for Type19Enum enum
 */
type Type19Enum int

/**
 * Value collection for Type19Enum enum
 */
const (
    Type19_KVCENTER Type19Enum = 1 + iota
    Type19_KFOLDER
    Type19_KDATACENTER
    Type19_KCOMPUTERESOURCE
    Type19_KCLUSTERCOMPUTERESOURCE
    Type19_KRESOURCEPOOL
    Type19_KDATASTORE
    Type19_KHOSTSYSTEM
    Type19_KVIRTUALMACHINE
    Type19_KVIRTUALAPP
    Type19_KSTANDALONEHOST
    Type19_KSTORAGEPOD
    Type19_KNETWORK
    Type19_KDISTRIBUTEDVIRTUALPORTGROUP
    Type19_KTAGCATEGORY
    Type19_KTAG
    Type19_KOPAQUENETWORK
    Type19_KVCLOUDDIRECTOR
    Type19_KORGANIZATION
    Type19_KVIRTUALDATACENTER
    Type19_KCATALOG
    Type19_KORGMETADATA
    Type19_KSTORAGEPOLICY
)

func (r Type19Enum) MarshalJSON() ([]byte, error) { 
    s := Type19EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type19Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type19EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type19Enum to its string representation
 */
func Type19EnumToValue(type19Enum Type19Enum) string {
    switch type19Enum {
        case Type19_KVCENTER:
    		return "kVCenter"		
        case Type19_KFOLDER:
    		return "kFolder"		
        case Type19_KDATACENTER:
    		return "kDatacenter"		
        case Type19_KCOMPUTERESOURCE:
    		return "kComputeResource"		
        case Type19_KCLUSTERCOMPUTERESOURCE:
    		return "kClusterComputeResource"		
        case Type19_KRESOURCEPOOL:
    		return "kResourcePool"		
        case Type19_KDATASTORE:
    		return "kDatastore"		
        case Type19_KHOSTSYSTEM:
    		return "kHostSystem"		
        case Type19_KVIRTUALMACHINE:
    		return "kVirtualMachine"		
        case Type19_KVIRTUALAPP:
    		return "kVirtualApp"		
        case Type19_KSTANDALONEHOST:
    		return "kStandaloneHost"		
        case Type19_KSTORAGEPOD:
    		return "kStoragePod"		
        case Type19_KNETWORK:
    		return "kNetwork"		
        case Type19_KDISTRIBUTEDVIRTUALPORTGROUP:
    		return "kDistributedVirtualPortgroup"		
        case Type19_KTAGCATEGORY:
    		return "kTagCategory"		
        case Type19_KTAG:
    		return "kTag"		
        case Type19_KOPAQUENETWORK:
    		return "kOpaqueNetwork"		
        case Type19_KVCLOUDDIRECTOR:
    		return "kVCloudDirector"		
        case Type19_KORGANIZATION:
    		return "kOrganization"		
        case Type19_KVIRTUALDATACENTER:
    		return "kVirtualDatacenter"		
        case Type19_KCATALOG:
    		return "kCatalog"		
        case Type19_KORGMETADATA:
    		return "kOrgMetadata"		
        case Type19_KSTORAGEPOLICY:
    		return "kStoragePolicy"		
        default:
        	return "kVCenter"
    }
}

/**
 * Converts Type19Enum Array to its string Array representation
*/
func Type19EnumArrayToValue(type19Enum []Type19Enum) []string {
    convArray := make([]string,len( type19Enum))
    for i:=0; i<len(type19Enum);i++ {
        convArray[i] = Type19EnumToValue(type19Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type19EnumFromValue(value string) Type19Enum {
    switch value {
        case "kVCenter":
            return Type19_KVCENTER
        case "kFolder":
            return Type19_KFOLDER
        case "kDatacenter":
            return Type19_KDATACENTER
        case "kComputeResource":
            return Type19_KCOMPUTERESOURCE
        case "kClusterComputeResource":
            return Type19_KCLUSTERCOMPUTERESOURCE
        case "kResourcePool":
            return Type19_KRESOURCEPOOL
        case "kDatastore":
            return Type19_KDATASTORE
        case "kHostSystem":
            return Type19_KHOSTSYSTEM
        case "kVirtualMachine":
            return Type19_KVIRTUALMACHINE
        case "kVirtualApp":
            return Type19_KVIRTUALAPP
        case "kStandaloneHost":
            return Type19_KSTANDALONEHOST
        case "kStoragePod":
            return Type19_KSTORAGEPOD
        case "kNetwork":
            return Type19_KNETWORK
        case "kDistributedVirtualPortgroup":
            return Type19_KDISTRIBUTEDVIRTUALPORTGROUP
        case "kTagCategory":
            return Type19_KTAGCATEGORY
        case "kTag":
            return Type19_KTAG
        case "kOpaqueNetwork":
            return Type19_KOPAQUENETWORK
        case "kVCloudDirector":
            return Type19_KVCLOUDDIRECTOR
        case "kOrganization":
            return Type19_KORGANIZATION
        case "kVirtualDatacenter":
            return Type19_KVIRTUALDATACENTER
        case "kCatalog":
            return Type19_KCATALOG
        case "kOrgMetadata":
            return Type19_KORGMETADATA
        case "kStoragePolicy":
            return Type19_KSTORAGEPOLICY
        default:
            return Type19_KVCENTER
    }
}
