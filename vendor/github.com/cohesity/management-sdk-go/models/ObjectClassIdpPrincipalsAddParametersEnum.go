// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClassIdpPrincipalsAddParametersEnum enum
 */
type ObjectClassIdpPrincipalsAddParametersEnum int

/**
 * Value collection for ObjectClassIdpPrincipalsAddParametersEnum enum
 */
const (
    ObjectClassIdpPrincipalsAddParameters_KUSER ObjectClassIdpPrincipalsAddParametersEnum = 1 + iota
    ObjectClassIdpPrincipalsAddParameters_KGROUP
    ObjectClassIdpPrincipalsAddParameters_KCOMPUTER
    ObjectClassIdpPrincipalsAddParameters_KWELLKNOWNPRINCIPAL
)

func (r ObjectClassIdpPrincipalsAddParametersEnum) MarshalJSON() ([]byte, error) {
    s := ObjectClassIdpPrincipalsAddParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectClassIdpPrincipalsAddParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectClassIdpPrincipalsAddParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectClassIdpPrincipalsAddParametersEnum to its string representation
 */
func ObjectClassIdpPrincipalsAddParametersEnumToValue(objectClassIdpPrincipalsAddParametersEnum ObjectClassIdpPrincipalsAddParametersEnum) string {
    switch objectClassIdpPrincipalsAddParametersEnum {
        case ObjectClassIdpPrincipalsAddParameters_KUSER:
    		return "kUser"
        case ObjectClassIdpPrincipalsAddParameters_KGROUP:
    		return "kGroup"
        case ObjectClassIdpPrincipalsAddParameters_KCOMPUTER:
    		return "kComputer"
        case ObjectClassIdpPrincipalsAddParameters_KWELLKNOWNPRINCIPAL:
    		return "kWellKnownPrincipal"
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClassIdpPrincipalsAddParametersEnum Array to its string Array representation
*/
func ObjectClassIdpPrincipalsAddParametersEnumArrayToValue(objectClassIdpPrincipalsAddParametersEnum []ObjectClassIdpPrincipalsAddParametersEnum) []string {
    convArray := make([]string,len( objectClassIdpPrincipalsAddParametersEnum))
    for i:=0; i<len(objectClassIdpPrincipalsAddParametersEnum);i++ {
        convArray[i] = ObjectClassIdpPrincipalsAddParametersEnumToValue(objectClassIdpPrincipalsAddParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClassIdpPrincipalsAddParametersEnumFromValue(value string) ObjectClassIdpPrincipalsAddParametersEnum {
    switch value {
        case "kUser":
            return ObjectClassIdpPrincipalsAddParameters_KUSER
        case "kGroup":
            return ObjectClassIdpPrincipalsAddParameters_KGROUP
        case "kComputer":
            return ObjectClassIdpPrincipalsAddParameters_KCOMPUTER
        case "kWellKnownPrincipal":
            return ObjectClassIdpPrincipalsAddParameters_KWELLKNOWNPRINCIPAL
        default:
            return ObjectClassIdpPrincipalsAddParameters_KUSER
    }
}
