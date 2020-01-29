// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeEnum enum
 */
type TypeEnum int

/**
 * Value collection for TypeEnum enum
 */
const (
    Type_KPRISMCENTRAL TypeEnum = 1 + iota
    Type_KSTANDALONECLUSTER
    Type_KCLUSTER
    Type_KHOST
    Type_KVIRTUALMACHINE
    Type_KNETWORK
    Type_KSTORAGECONTAINER
)

func (r TypeEnum) MarshalJSON() ([]byte, error) {
    s := TypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeEnum to its string representation
 */
func TypeEnumToValue(typeEnum TypeEnum) string {
    switch typeEnum {
        case Type_KPRISMCENTRAL:
    		return "kPrismCentral"
        case Type_KSTANDALONECLUSTER:
    		return "kStandaloneCluster"
        case Type_KCLUSTER:
    		return "kCluster"
        case Type_KHOST:
    		return "kHost"
        case Type_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case Type_KNETWORK:
    		return "kNetwork"
        case Type_KSTORAGECONTAINER:
    		return "kStorageContainer"
        default:
        	return "kPrismCentral"
    }
}

/**
 * Converts TypeEnum Array to its string Array representation
*/
func TypeEnumArrayToValue(typeEnum []TypeEnum) []string {
    convArray := make([]string,len( typeEnum))
    for i:=0; i<len(typeEnum);i++ {
        convArray[i] = TypeEnumToValue(typeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeEnumFromValue(value string) TypeEnum {
    switch value {
        case "kPrismCentral":
            return Type_KPRISMCENTRAL
        case "kStandaloneCluster":
            return Type_KSTANDALONECLUSTER
        case "kCluster":
            return Type_KCLUSTER
        case "kHost":
            return Type_KHOST
        case "kVirtualMachine":
            return Type_KVIRTUALMACHINE
        case "kNetwork":
            return Type_KNETWORK
        case "kStorageContainer":
            return Type_KSTORAGECONTAINER
        default:
            return Type_KPRISMCENTRAL
    }
}
