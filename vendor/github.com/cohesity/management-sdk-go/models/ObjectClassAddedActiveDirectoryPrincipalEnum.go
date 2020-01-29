// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassAddedActiveDirectoryPrincipalEnum enum
 */
type ObjectClassAddedActiveDirectoryPrincipalEnum int

/**
 * Value collection for ObjectClassAddedActiveDirectoryPrincipalEnum enum
 */
const (
    ObjectClassAddedActiveDirectoryPrincipal_KUSER ObjectClassAddedActiveDirectoryPrincipalEnum = 1 + iota
    ObjectClassAddedActiveDirectoryPrincipal_KGROUP
    ObjectClassAddedActiveDirectoryPrincipal_KCOMPUTER
    ObjectClassAddedActiveDirectoryPrincipal_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassAddedActiveDirectoryPrincipalEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassAddedActiveDirectoryPrincipalEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassAddedActiveDirectoryPrincipalEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassAddedActiveDirectoryPrincipalEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassAddedActiveDirectoryPrincipalEnum to its string representation
 */
func ObjectClassAddedActiveDirectoryPrincipalEnumToValue(objectClassAddedActiveDirectoryPrincipalEnum ObjectClassAddedActiveDirectoryPrincipalEnum) string {
    switch objectClassAddedActiveDirectoryPrincipalEnum {
        case ObjectClassAddedActiveDirectoryPrincipal_KUSER:
    		return "kUser"
        case ObjectClassAddedActiveDirectoryPrincipal_KGROUP:
    		return "kGroup"
        case ObjectClassAddedActiveDirectoryPrincipal_KCOMPUTER:
    		return "kComputer"
        case ObjectClassAddedActiveDirectoryPrincipal_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassAddedActiveDirectoryPrincipalEnum Array to its string Array representation
*/
func ObjectClassAddedActiveDirectoryPrincipalEnumArrayToValue(objectClassAddedActiveDirectoryPrincipalEnum []ObjectClassAddedActiveDirectoryPrincipalEnum) []string {
    convArray := make([]string,len( objectClassAddedActiveDirectoryPrincipalEnum))
    for i:=0; i<len(objectClassAddedActiveDirectoryPrincipalEnum);i++ {
        convArray[i] = ObjectClassAddedActiveDirectoryPrincipalEnumToValue(objectClassAddedActiveDirectoryPrincipalEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassAddedActiveDirectoryPrincipalEnumFromValue(value string) ObjectClassAddedActiveDirectoryPrincipalEnum {
    switch value {
        case "kUser":
            return ObjectClassAddedActiveDirectoryPrincipal_KUSER
        case "kGroup":
            return ObjectClassAddedActiveDirectoryPrincipal_KGROUP
        case "kComputer":
            return ObjectClassAddedActiveDirectoryPrincipal_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClassAddedActiveDirectoryPrincipal_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClassAddedActiveDirectoryPrincipal_KUSER
    }
}
