package models

import(
    "encoding/json"
)

/**
 * Type definition for Type9Enum enum
 */
type Type9Enum int

/**
 * Value collection for Type9Enum enum
 */
const (
    Type9_KOVIRTMANAGER Type9Enum = 1 + iota
    Type9_KSTANDALONEHOST
    Type9_KDATACENTER
    Type9_KCLUSTER
    Type9_KHOST
    Type9_KVIRTUALMACHINE
    Type9_KNETWORK
    Type9_KSTORAGEDOMAIN
    Type9_KVNICPROFILE
)

func (r Type9Enum) MarshalJSON() ([]byte, error) { 
    s := Type9EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type9Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type9EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type9Enum to its string representation
 */
func Type9EnumToValue(type9Enum Type9Enum) string {
    switch type9Enum {
        case Type9_KOVIRTMANAGER:
    		return "kOVirtManager"		
        case Type9_KSTANDALONEHOST:
    		return "kStandaloneHost"		
        case Type9_KDATACENTER:
    		return "kDatacenter"		
        case Type9_KCLUSTER:
    		return "kCluster"		
        case Type9_KHOST:
    		return "kHost"		
        case Type9_KVIRTUALMACHINE:
    		return "kVirtualMachine"		
        case Type9_KNETWORK:
    		return "kNetwork"		
        case Type9_KSTORAGEDOMAIN:
    		return "kStorageDomain"		
        case Type9_KVNICPROFILE:
    		return "kVNicProfile"		
        default:
        	return "kOVirtManager"
    }
}

/**
 * Converts Type9Enum Array to its string Array representation
*/
func Type9EnumArrayToValue(type9Enum []Type9Enum) []string {
    convArray := make([]string,len( type9Enum))
    for i:=0; i<len(type9Enum);i++ {
        convArray[i] = Type9EnumToValue(type9Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type9EnumFromValue(value string) Type9Enum {
    switch value {
        case "kOVirtManager":
            return Type9_KOVIRTMANAGER
        case "kStandaloneHost":
            return Type9_KSTANDALONEHOST
        case "kDatacenter":
            return Type9_KDATACENTER
        case "kCluster":
            return Type9_KCLUSTER
        case "kHost":
            return Type9_KHOST
        case "kVirtualMachine":
            return Type9_KVIRTUALMACHINE
        case "kNetwork":
            return Type9_KNETWORK
        case "kStorageDomain":
            return Type9_KSTORAGEDOMAIN
        case "kVNicProfile":
            return Type9_KVNICPROFILE
        default:
            return Type9_KOVIRTMANAGER
    }
}
