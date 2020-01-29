package models

import(
    "encoding/json"
)

/**
 * Type definition for PropertiesEnum enum
 */
type PropertiesEnum int

/**
 * Value collection for PropertiesEnum enum
 */
const (
    Properties_VIEWBOX PropertiesEnum = 1 + iota
    Properties_VLAN
    Properties_PROTECTIONPOLICY
    Properties_ENTITY
    Properties_PROTECTIONJOB
    Properties_VIEWS
    Properties_ACTIVEDIRECTORY
    Properties_LDAPPROVIDER
)

func (r PropertiesEnum) MarshalJSON() ([]byte, error) { 
    s := PropertiesEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *PropertiesEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  PropertiesEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts PropertiesEnum to its string representation
 */
func PropertiesEnumToValue(propertiesEnum PropertiesEnum) string {
    switch propertiesEnum {
        case Properties_VIEWBOX:
    		return "ViewBox"		
        case Properties_VLAN:
    		return "Vlan"		
        case Properties_PROTECTIONPOLICY:
    		return "ProtectionPolicy"		
        case Properties_ENTITY:
    		return "Entity"		
        case Properties_PROTECTIONJOB:
    		return "ProtectionJob"		
        case Properties_VIEWS:
    		return "Views"		
        case Properties_ACTIVEDIRECTORY:
    		return "ActiveDirectory"		
        case Properties_LDAPPROVIDER:
    		return "LdapProvider"		
        default:
        	return "ViewBox"
    }
}

/**
 * Converts PropertiesEnum Array to its string Array representation
*/
func PropertiesEnumArrayToValue(propertiesEnum []PropertiesEnum) []string {
    convArray := make([]string,len( propertiesEnum))
    for i:=0; i<len(propertiesEnum);i++ {
        convArray[i] = PropertiesEnumToValue(propertiesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PropertiesEnumFromValue(value string) PropertiesEnum {
    switch value {
        case "ViewBox":
            return Properties_VIEWBOX
        case "Vlan":
            return Properties_VLAN
        case "ProtectionPolicy":
            return Properties_PROTECTIONPOLICY
        case "Entity":
            return Properties_ENTITY
        case "ProtectionJob":
            return Properties_PROTECTIONJOB
        case "Views":
            return Properties_VIEWS
        case "ActiveDirectory":
            return Properties_ACTIVEDIRECTORY
        case "LdapProvider":
            return Properties_LDAPPROVIDER
        default:
            return Properties_VIEWBOX
    }
}
