package models

import(
    "encoding/json"
)

/**
 * Type definition for Type7Enum enum
 */
type Type7Enum int

/**
 * Value collection for Type7Enum enum
 */
const (
    Type7_KSCVMMSERVER Type7Enum = 1 + iota
    Type7_KSTANDALONEHOST
    Type7_KSTANDALONECLUSTER
    Type7_KHOSTGROUP
    Type7_KHOST
    Type7_KHOSTCLUSTER
    Type7_KVIRTUALMACHINE
    Type7_KNETWORK
    Type7_KDATASTORE
)

func (r Type7Enum) MarshalJSON() ([]byte, error) { 
    s := Type7EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type7Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type7EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type7Enum to its string representation
 */
func Type7EnumToValue(type7Enum Type7Enum) string {
    switch type7Enum {
        case Type7_KSCVMMSERVER:
    		return "kSCVMMServer"		
        case Type7_KSTANDALONEHOST:
    		return "kStandaloneHost"		
        case Type7_KSTANDALONECLUSTER:
    		return "kStandaloneCluster"		
        case Type7_KHOSTGROUP:
    		return "kHostGroup"		
        case Type7_KHOST:
    		return "kHost"		
        case Type7_KHOSTCLUSTER:
    		return "kHostCluster"		
        case Type7_KVIRTUALMACHINE:
    		return "kVirtualMachine"		
        case Type7_KNETWORK:
    		return "kNetwork"		
        case Type7_KDATASTORE:
    		return "kDatastore"		
        default:
        	return "kSCVMMServer"
    }
}

/**
 * Converts Type7Enum Array to its string Array representation
*/
func Type7EnumArrayToValue(type7Enum []Type7Enum) []string {
    convArray := make([]string,len( type7Enum))
    for i:=0; i<len(type7Enum);i++ {
        convArray[i] = Type7EnumToValue(type7Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type7EnumFromValue(value string) Type7Enum {
    switch value {
        case "kSCVMMServer":
            return Type7_KSCVMMSERVER
        case "kStandaloneHost":
            return Type7_KSTANDALONEHOST
        case "kStandaloneCluster":
            return Type7_KSTANDALONECLUSTER
        case "kHostGroup":
            return Type7_KHOSTGROUP
        case "kHost":
            return Type7_KHOST
        case "kHostCluster":
            return Type7_KHOSTCLUSTER
        case "kVirtualMachine":
            return Type7_KVIRTUALMACHINE
        case "kNetwork":
            return Type7_KNETWORK
        case "kDatastore":
            return Type7_KDATASTORE
        default:
            return Type7_KSCVMMSERVER
    }
}
