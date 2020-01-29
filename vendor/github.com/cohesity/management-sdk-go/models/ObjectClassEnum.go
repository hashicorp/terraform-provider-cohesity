// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassEnum enum
 */
type ObjectClassEnum int

/**
 * Value collection for ObjectClassEnum enum
 */
const (
    ObjectClass_KUSER ObjectClassEnum = 1 + iota
    ObjectClass_KGROUP
    ObjectClass_KCOMPUTER
    ObjectClass_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassEnum to its string representation
 */
func ObjectClassEnumToValue(objectClassEnum ObjectClassEnum) string {
    switch objectClassEnum {
        case ObjectClass_KUSER:
    		return "kUser"
        case ObjectClass_KGROUP:
    		return "kGroup"
        case ObjectClass_KCOMPUTER:
    		return "kComputer"
        case ObjectClass_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassEnum Array to its string Array representation
*/
func ObjectClassEnumArrayToValue(objectClassEnum []ObjectClassEnum) []string {
    convArray := make([]string,len( objectClassEnum))
    for i:=0; i<len(objectClassEnum);i++ {
        convArray[i] = ObjectClassEnumToValue(objectClassEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassEnumFromValue(value string) ObjectClassEnum {
    switch value {
        case "kUser":
            return ObjectClass_KUSER
        case "kGroup":
            return ObjectClass_KGROUP
        case "kComputer":
            return ObjectClass_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClass_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClass_KUSER
    }
}
