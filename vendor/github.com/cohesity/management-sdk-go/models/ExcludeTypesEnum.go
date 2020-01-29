package models

import(
    "encoding/json"
)

/**
 * Type definition for ExcludeTypesEnum enum
 */
type ExcludeTypesEnum int

/**
 * Value collection for ExcludeTypesEnum enum
 */
const (
    ExcludeTypes_KVCENTER ExcludeTypesEnum = 1 + iota
    ExcludeTypes_KFOLDER
    ExcludeTypes_KDATACENTER
    ExcludeTypes_KCOMPUTERESOURCE
    ExcludeTypes_KCLUSTERCOMPUTERESOURCE
    ExcludeTypes_KRESOURCEPOOL
    ExcludeTypes_KDATASTORE
    ExcludeTypes_KHOSTSYSTEM
    ExcludeTypes_KVIRTUALMACHINE
    ExcludeTypes_KVIRTUALAPP
    ExcludeTypes_KSTANDALONEHOST
    ExcludeTypes_KSTORAGEPOD
    ExcludeTypes_KNETWORK
    ExcludeTypes_KDISTRIBUTEDVIRTUALPORTGROUP
    ExcludeTypes_KTAGCATEGORY
    ExcludeTypes_KTAG
)

func (r ExcludeTypesEnum) MarshalJSON() ([]byte, error) { 
    s := ExcludeTypesEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ExcludeTypesEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ExcludeTypesEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ExcludeTypesEnum to its string representation
 */
func ExcludeTypesEnumToValue(excludeTypesEnum ExcludeTypesEnum) string {
    switch excludeTypesEnum {
        case ExcludeTypes_KVCENTER:
    		return "kVCenter"		
        case ExcludeTypes_KFOLDER:
    		return "kFolder"		
        case ExcludeTypes_KDATACENTER:
    		return "kDatacenter"		
        case ExcludeTypes_KCOMPUTERESOURCE:
    		return "kComputeResource"		
        case ExcludeTypes_KCLUSTERCOMPUTERESOURCE:
    		return "kClusterComputeResource"		
        case ExcludeTypes_KRESOURCEPOOL:
    		return "kResourcePool"		
        case ExcludeTypes_KDATASTORE:
    		return "kDatastore"		
        case ExcludeTypes_KHOSTSYSTEM:
    		return "kHostSystem"		
        case ExcludeTypes_KVIRTUALMACHINE:
    		return "kVirtualMachine"		
        case ExcludeTypes_KVIRTUALAPP:
    		return "kVirtualApp"		
        case ExcludeTypes_KSTANDALONEHOST:
    		return "kStandaloneHost"		
        case ExcludeTypes_KSTORAGEPOD:
    		return "kStoragePod"		
        case ExcludeTypes_KNETWORK:
    		return "kNetwork"		
        case ExcludeTypes_KDISTRIBUTEDVIRTUALPORTGROUP:
    		return "kDistributedVirtualPortgroup"		
        case ExcludeTypes_KTAGCATEGORY:
    		return "kTagCategory"		
        case ExcludeTypes_KTAG:
    		return "kTag"		
        default:
        	return "kVCenter"
    }
}

/**
 * Converts ExcludeTypesEnum Array to its string Array representation
*/
func ExcludeTypesEnumArrayToValue(excludeTypesEnum []ExcludeTypesEnum) []string {
    convArray := make([]string,len( excludeTypesEnum))
    for i:=0; i<len(excludeTypesEnum);i++ {
        convArray[i] = ExcludeTypesEnumToValue(excludeTypesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExcludeTypesEnumFromValue(value string) ExcludeTypesEnum {
    switch value {
        case "kVCenter":
            return ExcludeTypes_KVCENTER
        case "kFolder":
            return ExcludeTypes_KFOLDER
        case "kDatacenter":
            return ExcludeTypes_KDATACENTER
        case "kComputeResource":
            return ExcludeTypes_KCOMPUTERESOURCE
        case "kClusterComputeResource":
            return ExcludeTypes_KCLUSTERCOMPUTERESOURCE
        case "kResourcePool":
            return ExcludeTypes_KRESOURCEPOOL
        case "kDatastore":
            return ExcludeTypes_KDATASTORE
        case "kHostSystem":
            return ExcludeTypes_KHOSTSYSTEM
        case "kVirtualMachine":
            return ExcludeTypes_KVIRTUALMACHINE
        case "kVirtualApp":
            return ExcludeTypes_KVIRTUALAPP
        case "kStandaloneHost":
            return ExcludeTypes_KSTANDALONEHOST
        case "kStoragePod":
            return ExcludeTypes_KSTORAGEPOD
        case "kNetwork":
            return ExcludeTypes_KNETWORK
        case "kDistributedVirtualPortgroup":
            return ExcludeTypes_KDISTRIBUTEDVIRTUALPORTGROUP
        case "kTagCategory":
            return ExcludeTypes_KTAGCATEGORY
        case "kTag":
            return ExcludeTypes_KTAG
        default:
            return ExcludeTypes_KVCENTER
    }
}
