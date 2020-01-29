// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassSearchActiveDirectoryPrincipalsEnum enum
 */
type ObjectClassSearchActiveDirectoryPrincipalsEnum int

/**
 * Value collection for ObjectClassSearchActiveDirectoryPrincipalsEnum enum
 */
const (
    ObjectClassSearchActiveDirectoryPrincipals_KUSER ObjectClassSearchActiveDirectoryPrincipalsEnum = 1 + iota
    ObjectClassSearchActiveDirectoryPrincipals_KGROUP
    ObjectClassSearchActiveDirectoryPrincipals_KCOMPUTER
    ObjectClassSearchActiveDirectoryPrincipals_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassSearchActiveDirectoryPrincipalsEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassSearchActiveDirectoryPrincipalsEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassSearchActiveDirectoryPrincipalsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassSearchActiveDirectoryPrincipalsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassSearchActiveDirectoryPrincipalsEnum to its string representation
 */
func ObjectClassSearchActiveDirectoryPrincipalsEnumToValue(objectClassSearchActiveDirectoryPrincipalsEnum ObjectClassSearchActiveDirectoryPrincipalsEnum) string {
    switch objectClassSearchActiveDirectoryPrincipalsEnum {
        case ObjectClassSearchActiveDirectoryPrincipals_KUSER:
    		return "kUser"
        case ObjectClassSearchActiveDirectoryPrincipals_KGROUP:
    		return "kGroup"
        case ObjectClassSearchActiveDirectoryPrincipals_KCOMPUTER:
    		return "kComputer"
        case ObjectClassSearchActiveDirectoryPrincipals_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassSearchActiveDirectoryPrincipalsEnum Array to its string Array representation
*/
func ObjectClassSearchActiveDirectoryPrincipalsEnumArrayToValue(objectClassSearchActiveDirectoryPrincipalsEnum []ObjectClassSearchActiveDirectoryPrincipalsEnum) []string {
    convArray := make([]string,len( objectClassSearchActiveDirectoryPrincipalsEnum))
    for i:=0; i<len(objectClassSearchActiveDirectoryPrincipalsEnum);i++ {
        convArray[i] = ObjectClassSearchActiveDirectoryPrincipalsEnumToValue(objectClassSearchActiveDirectoryPrincipalsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassSearchActiveDirectoryPrincipalsEnumFromValue(value string) ObjectClassSearchActiveDirectoryPrincipalsEnum {
    switch value {
        case "kUser":
            return ObjectClassSearchActiveDirectoryPrincipals_KUSER
        case "kGroup":
            return ObjectClassSearchActiveDirectoryPrincipals_KGROUP
        case "kComputer":
            return ObjectClassSearchActiveDirectoryPrincipals_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClassSearchActiveDirectoryPrincipals_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClassSearchActiveDirectoryPrincipals_KUSER
    }
}
