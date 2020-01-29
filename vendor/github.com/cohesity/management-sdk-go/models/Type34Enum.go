package models

import(
    "encoding/json"
)

/**
 * Type definition for Type34Enum enum
 */
type Type34Enum int

/**
 * Value collection for Type34Enum enum
 */
const (
    Type34_KNEARLINE Type34Enum = 1 + iota
    Type34_KCOLDLINE
    Type34_KGLACIER
    Type34_KS3
    Type34_KAZURESTANDARD
    Type34_KS3COMPATIBLE
    Type34_KQSTARTAPE
    Type34_KGOOGLESTANDARD
    Type34_KGOOGLEDRA
    Type34_KAWSGOVCLOUD
    Type34_KNAS
    Type34_KAZUREGOVCLOUD
)

func (r Type34Enum) MarshalJSON() ([]byte, error) { 
    s := Type34EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type34Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type34EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type34Enum to its string representation
 */
func Type34EnumToValue(type34Enum Type34Enum) string {
    switch type34Enum {
        case Type34_KNEARLINE:
    		return "kNearline"		
        case Type34_KCOLDLINE:
    		return "kColdline"		
        case Type34_KGLACIER:
    		return "kGlacier"		
        case Type34_KS3:
    		return "kS3"		
        case Type34_KAZURESTANDARD:
    		return "kAzureStandard"		
        case Type34_KS3COMPATIBLE:
    		return "kS3Compatible"		
        case Type34_KQSTARTAPE:
    		return "kQStarTape"		
        case Type34_KGOOGLESTANDARD:
    		return "kGoogleStandard"		
        case Type34_KGOOGLEDRA:
    		return "kGoogleDRA"		
        case Type34_KAWSGOVCLOUD:
    		return "kAWSGovCloud"		
        case Type34_KNAS:
    		return "kNAS"		
        case Type34_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"		
        default:
        	return "kNearline"
    }
}

/**
 * Converts Type34Enum Array to its string Array representation
*/
func Type34EnumArrayToValue(type34Enum []Type34Enum) []string {
    convArray := make([]string,len( type34Enum))
    for i:=0; i<len(type34Enum);i++ {
        convArray[i] = Type34EnumToValue(type34Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type34EnumFromValue(value string) Type34Enum {
    switch value {
        case "kNearline":
            return Type34_KNEARLINE
        case "kColdline":
            return Type34_KCOLDLINE
        case "kGlacier":
            return Type34_KGLACIER
        case "kS3":
            return Type34_KS3
        case "kAzureStandard":
            return Type34_KAZURESTANDARD
        case "kS3Compatible":
            return Type34_KS3COMPATIBLE
        case "kQStarTape":
            return Type34_KQSTARTAPE
        case "kGoogleStandard":
            return Type34_KGOOGLESTANDARD
        case "kGoogleDRA":
            return Type34_KGOOGLEDRA
        case "kAWSGovCloud":
            return Type34_KAWSGOVCLOUD
        case "kNAS":
            return Type34_KNAS
        case "kAzureGovCloud":
            return Type34_KAZUREGOVCLOUD
        default:
            return Type34_KNEARLINE
    }
}
