package models

import(
    "encoding/json"
)

/**
 * Type definition for ExcludeOffice365TypesEnum enum
 */
type ExcludeOffice365TypesEnum int

/**
 * Value collection for ExcludeOffice365TypesEnum enum
 */
const (
    ExcludeOffice365Types_KDOMAIN ExcludeOffice365TypesEnum = 1 + iota
    ExcludeOffice365Types_KOUTLOOK
    ExcludeOffice365Types_KMAILBOX
)

func (r ExcludeOffice365TypesEnum) MarshalJSON() ([]byte, error) { 
    s := ExcludeOffice365TypesEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ExcludeOffice365TypesEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ExcludeOffice365TypesEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ExcludeOffice365TypesEnum to its string representation
 */
func ExcludeOffice365TypesEnumToValue(excludeOffice365TypesEnum ExcludeOffice365TypesEnum) string {
    switch excludeOffice365TypesEnum {
        case ExcludeOffice365Types_KDOMAIN:
    		return "kDomain"		
        case ExcludeOffice365Types_KOUTLOOK:
    		return "kOutlook"		
        case ExcludeOffice365Types_KMAILBOX:
    		return "kMailbox"		
        default:
        	return "kDomain"
    }
}

/**
 * Converts ExcludeOffice365TypesEnum Array to its string Array representation
*/
func ExcludeOffice365TypesEnumArrayToValue(excludeOffice365TypesEnum []ExcludeOffice365TypesEnum) []string {
    convArray := make([]string,len( excludeOffice365TypesEnum))
    for i:=0; i<len(excludeOffice365TypesEnum);i++ {
        convArray[i] = ExcludeOffice365TypesEnumToValue(excludeOffice365TypesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExcludeOffice365TypesEnumFromValue(value string) ExcludeOffice365TypesEnum {
    switch value {
        case "kDomain":
            return ExcludeOffice365Types_KDOMAIN
        case "kOutlook":
            return ExcludeOffice365Types_KOUTLOOK
        case "kMailbox":
            return ExcludeOffice365Types_KMAILBOX
        default:
            return ExcludeOffice365Types_KDOMAIN
    }
}
