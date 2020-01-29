// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassSearchPrincipalsEnum enum
 */
type ObjectClassSearchPrincipalsEnum int

/**
 * Value collection for ObjectClassSearchPrincipalsEnum enum
 */
const (
    ObjectClassSearchPrincipals_KUSER ObjectClassSearchPrincipalsEnum = 1 + iota
    ObjectClassSearchPrincipals_KGROUP
    ObjectClassSearchPrincipals_KCOMPUTER
    ObjectClassSearchPrincipals_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassSearchPrincipalsEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassSearchPrincipalsEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassSearchPrincipalsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassSearchPrincipalsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassSearchPrincipalsEnum to its string representation
 */
func ObjectClassSearchPrincipalsEnumToValue(objectClassSearchPrincipalsEnum ObjectClassSearchPrincipalsEnum) string {
    switch objectClassSearchPrincipalsEnum {
        case ObjectClassSearchPrincipals_KUSER:
    		return "kUser"
        case ObjectClassSearchPrincipals_KGROUP:
    		return "kGroup"
        case ObjectClassSearchPrincipals_KCOMPUTER:
    		return "kComputer"
        case ObjectClassSearchPrincipals_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassSearchPrincipalsEnum Array to its string Array representation
*/
func ObjectClassSearchPrincipalsEnumArrayToValue(objectClassSearchPrincipalsEnum []ObjectClassSearchPrincipalsEnum) []string {
    convArray := make([]string,len( objectClassSearchPrincipalsEnum))
    for i:=0; i<len(objectClassSearchPrincipalsEnum);i++ {
        convArray[i] = ObjectClassSearchPrincipalsEnumToValue(objectClassSearchPrincipalsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassSearchPrincipalsEnumFromValue(value string) ObjectClassSearchPrincipalsEnum {
    switch value {
        case "kUser":
            return ObjectClassSearchPrincipals_KUSER
        case "kGroup":
            return ObjectClassSearchPrincipals_KGROUP
        case "kComputer":
            return ObjectClassSearchPrincipals_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClassSearchPrincipals_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClassSearchPrincipals_KUSER
    }
}
