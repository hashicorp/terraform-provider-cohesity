// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeKvmProtectionSourceEnum enum
 */
type TypeKvmProtectionSourceEnum int

/**
 * Value collection for TypeKvmProtectionSourceEnum enum
 */
const (
    TypeKvmProtectionSource_KOVIRTMANAGER TypeKvmProtectionSourceEnum = 1 + iota
    TypeKvmProtectionSource_KSTANDALONEHOST
    TypeKvmProtectionSource_KDATACENTER
    TypeKvmProtectionSource_KCLUSTER
    TypeKvmProtectionSource_KHOST
    TypeKvmProtectionSource_KVIRTUALMACHINE
    TypeKvmProtectionSource_KNETWORK
    TypeKvmProtectionSource_KSTORAGEDOMAIN
    TypeKvmProtectionSource_KVNICPROFILE
)

func (r TypeKvmProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeKvmProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeKvmProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeKvmProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeKvmProtectionSourceEnum to its string representation
 */
func TypeKvmProtectionSourceEnumToValue(typeKvmProtectionSourceEnum TypeKvmProtectionSourceEnum) string {
    switch typeKvmProtectionSourceEnum {
        case TypeKvmProtectionSource_KOVIRTMANAGER:
    		return "kOVirtManager"
        case TypeKvmProtectionSource_KSTANDALONEHOST:
    		return "kStandaloneHost"
        case TypeKvmProtectionSource_KDATACENTER:
    		return "kDatacenter"
        case TypeKvmProtectionSource_KCLUSTER:
    		return "kCluster"
        case TypeKvmProtectionSource_KHOST:
    		return "kHost"
        case TypeKvmProtectionSource_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case TypeKvmProtectionSource_KNETWORK:
    		return "kNetwork"
        case TypeKvmProtectionSource_KSTORAGEDOMAIN:
    		return "kStorageDomain"
        case TypeKvmProtectionSource_KVNICPROFILE:
    		return "kVNicProfile"
        default:
        	return "kOVirtManager"
    }
}

/**
 * Converts TypeKvmProtectionSourceEnum Array to its string Array representation
*/
func TypeKvmProtectionSourceEnumArrayToValue(typeKvmProtectionSourceEnum []TypeKvmProtectionSourceEnum) []string {
    convArray := make([]string,len( typeKvmProtectionSourceEnum))
    for i:=0; i<len(typeKvmProtectionSourceEnum);i++ {
        convArray[i] = TypeKvmProtectionSourceEnumToValue(typeKvmProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeKvmProtectionSourceEnumFromValue(value string) TypeKvmProtectionSourceEnum {
    switch value {
        case "kOVirtManager":
            return TypeKvmProtectionSource_KOVIRTMANAGER
        case "kStandaloneHost":
            return TypeKvmProtectionSource_KSTANDALONEHOST
        case "kDatacenter":
            return TypeKvmProtectionSource_KDATACENTER
        case "kCluster":
            return TypeKvmProtectionSource_KCLUSTER
        case "kHost":
            return TypeKvmProtectionSource_KHOST
        case "kVirtualMachine":
            return TypeKvmProtectionSource_KVIRTUALMACHINE
        case "kNetwork":
            return TypeKvmProtectionSource_KNETWORK
        case "kStorageDomain":
            return TypeKvmProtectionSource_KSTORAGEDOMAIN
        case "kVNicProfile":
            return TypeKvmProtectionSource_KVNICPROFILE
        default:
            return TypeKvmProtectionSource_KOVIRTMANAGER
    }
}
