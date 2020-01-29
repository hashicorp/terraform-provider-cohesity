// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassActiveDirectoryPrincipalsAddParametersEnum enum
 */
type ObjectClassActiveDirectoryPrincipalsAddParametersEnum int

/**
 * Value collection for ObjectClassActiveDirectoryPrincipalsAddParametersEnum enum
 */
const (
    ObjectClassActiveDirectoryPrincipalsAddParameters_KUSER ObjectClassActiveDirectoryPrincipalsAddParametersEnum = 1 + iota
    ObjectClassActiveDirectoryPrincipalsAddParameters_KGROUP
    ObjectClassActiveDirectoryPrincipalsAddParameters_KCOMPUTER
    ObjectClassActiveDirectoryPrincipalsAddParameters_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassActiveDirectoryPrincipalsAddParametersEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassActiveDirectoryPrincipalsAddParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassActiveDirectoryPrincipalsAddParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassActiveDirectoryPrincipalsAddParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassActiveDirectoryPrincipalsAddParametersEnum to its string representation
 */
func ObjectClassActiveDirectoryPrincipalsAddParametersEnumToValue(objectClassActiveDirectoryPrincipalsAddParametersEnum ObjectClassActiveDirectoryPrincipalsAddParametersEnum) string {
    switch objectClassActiveDirectoryPrincipalsAddParametersEnum {
        case ObjectClassActiveDirectoryPrincipalsAddParameters_KUSER:
    		return "kUser"
        case ObjectClassActiveDirectoryPrincipalsAddParameters_KGROUP:
    		return "kGroup"
        case ObjectClassActiveDirectoryPrincipalsAddParameters_KCOMPUTER:
    		return "kComputer"
        case ObjectClassActiveDirectoryPrincipalsAddParameters_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassActiveDirectoryPrincipalsAddParametersEnum Array to its string Array representation
*/
func ObjectClassActiveDirectoryPrincipalsAddParametersEnumArrayToValue(objectClassActiveDirectoryPrincipalsAddParametersEnum []ObjectClassActiveDirectoryPrincipalsAddParametersEnum) []string {
    convArray := make([]string,len( objectClassActiveDirectoryPrincipalsAddParametersEnum))
    for i:=0; i<len(objectClassActiveDirectoryPrincipalsAddParametersEnum);i++ {
        convArray[i] = ObjectClassActiveDirectoryPrincipalsAddParametersEnumToValue(objectClassActiveDirectoryPrincipalsAddParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassActiveDirectoryPrincipalsAddParametersEnumFromValue(value string) ObjectClassActiveDirectoryPrincipalsAddParametersEnum {
    switch value {
        case "kUser":
            return ObjectClassActiveDirectoryPrincipalsAddParameters_KUSER
        case "kGroup":
            return ObjectClassActiveDirectoryPrincipalsAddParameters_KGROUP
        case "kComputer":
            return ObjectClassActiveDirectoryPrincipalsAddParameters_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClassActiveDirectoryPrincipalsAddParameters_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClassActiveDirectoryPrincipalsAddParameters_KUSER
    }
}
