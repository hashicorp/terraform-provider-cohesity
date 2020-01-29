// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PropertyEnum enum
 */
type PropertyEnum int

/**
 * Value collection for PropertyEnum enum
 */
const (
    Property_VIEWBOX PropertyEnum = 1 + iota
    Property_VLAN
    Property_PROTECTIONPOLICY
    Property_ENTITY
    Property_PROTECTIONJOB
    Property_VIEWS
    Property_ACTIVEDIRECTORY
    Property_LDAPPROVIDER
)

func (r PropertyEnum) MarshalJSON() ([]byte, error) {
    s := PropertyEnumToValue(r)
    return json.Marshal(s)
}

func (r *PropertyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PropertyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PropertyEnum to its string representation
 */
func PropertyEnumToValue(propertyEnum PropertyEnum) string {
    switch propertyEnum {
        case Property_VIEWBOX:
    		return "ViewBox"
        case Property_VLAN:
    		return "Vlan"
        case Property_PROTECTIONPOLICY:
    		return "ProtectionPolicy"
        case Property_ENTITY:
    		return "Entity"
        case Property_PROTECTIONJOB:
    		return "ProtectionJob"
        case Property_VIEWS:
    		return "Views"
        case Property_ACTIVEDIRECTORY:
    		return "ActiveDirectory"
        case Property_LDAPPROVIDER:
    		return "LdapProvider"
        default:
        	return "ViewBox"
    }
}

/**
 * Converts PropertyEnum Array to its string Array representation
*/
func PropertyEnumArrayToValue(propertyEnum []PropertyEnum) []string {
    convArray := make([]string,len( propertyEnum))
    for i:=0; i<len(propertyEnum);i++ {
        convArray[i] = PropertyEnumToValue(propertyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PropertyEnumFromValue(value string) PropertyEnum {
    switch value {
        case "ViewBox":
            return Property_VIEWBOX
        case "Vlan":
            return Property_VLAN
        case "ProtectionPolicy":
            return Property_PROTECTIONPOLICY
        case "Entity":
            return Property_ENTITY
        case "ProtectionJob":
            return Property_PROTECTIONJOB
        case "Views":
            return Property_VIEWS
        case "ActiveDirectory":
            return Property_ACTIVEDIRECTORY
        case "LdapProvider":
            return Property_LDAPPROVIDER
        default:
            return Property_VIEWBOX
    }
}
