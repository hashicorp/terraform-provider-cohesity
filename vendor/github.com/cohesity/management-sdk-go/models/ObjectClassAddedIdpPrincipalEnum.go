// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassAddedIdpPrincipalEnum enum
 */
type ObjectClassAddedIdpPrincipalEnum int

/**
 * Value collection for ObjectClassAddedIdpPrincipalEnum enum
 */
const (
    ObjectClassAddedIdpPrincipal_KUSER ObjectClassAddedIdpPrincipalEnum = 1 + iota
    ObjectClassAddedIdpPrincipal_KGROUP
    ObjectClassAddedIdpPrincipal_KCOMPUTER
    ObjectClassAddedIdpPrincipal_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassAddedIdpPrincipalEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassAddedIdpPrincipalEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassAddedIdpPrincipalEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassAddedIdpPrincipalEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassAddedIdpPrincipalEnum to its string representation
 */
func ObjectClassAddedIdpPrincipalEnumToValue(objectClassAddedIdpPrincipalEnum ObjectClassAddedIdpPrincipalEnum) string {
    switch objectClassAddedIdpPrincipalEnum {
        case ObjectClassAddedIdpPrincipal_KUSER:
    		return "kUser"
        case ObjectClassAddedIdpPrincipal_KGROUP:
    		return "kGroup"
        case ObjectClassAddedIdpPrincipal_KCOMPUTER:
    		return "kComputer"
        case ObjectClassAddedIdpPrincipal_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassAddedIdpPrincipalEnum Array to its string Array representation
*/
func ObjectClassAddedIdpPrincipalEnumArrayToValue(objectClassAddedIdpPrincipalEnum []ObjectClassAddedIdpPrincipalEnum) []string {
    convArray := make([]string,len( objectClassAddedIdpPrincipalEnum))
    for i:=0; i<len(objectClassAddedIdpPrincipalEnum);i++ {
        convArray[i] = ObjectClassAddedIdpPrincipalEnumToValue(objectClassAddedIdpPrincipalEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassAddedIdpPrincipalEnumFromValue(value string) ObjectClassAddedIdpPrincipalEnum {
    switch value {
        case "kUser":
            return ObjectClassAddedIdpPrincipal_KUSER
        case "kGroup":
            return ObjectClassAddedIdpPrincipal_KGROUP
        case "kComputer":
            return ObjectClassAddedIdpPrincipal_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClassAddedIdpPrincipal_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClassAddedIdpPrincipal_KUSER
    }
}
