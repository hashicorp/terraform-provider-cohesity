// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeNetappVolumeInfoEnum enum
 */
type TypeNetappVolumeInfoEnum int

/**
 * Value collection for TypeNetappVolumeInfoEnum enum
 */
const (
    TypeNetappVolumeInfo_KREADWRITE TypeNetappVolumeInfoEnum = 1 + iota
    TypeNetappVolumeInfo_KLOADSHARING
    TypeNetappVolumeInfo_KDATAPROTECTION
    TypeNetappVolumeInfo_KDATACACHE
    TypeNetappVolumeInfo_KTMP
    TypeNetappVolumeInfo_KUNKNOWNTYPE
)

func (r TypeNetappVolumeInfoEnum) MarshalJSON() ([]byte, error) {
    s := TypeNetappVolumeInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeNetappVolumeInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeNetappVolumeInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeNetappVolumeInfoEnum to its string representation
 */
func TypeNetappVolumeInfoEnumToValue(typeNetappVolumeInfoEnum TypeNetappVolumeInfoEnum) string {
    switch typeNetappVolumeInfoEnum {
        case TypeNetappVolumeInfo_KREADWRITE:
    		return "kReadWrite"
        case TypeNetappVolumeInfo_KLOADSHARING:
    		return "kLoadSharing"
        case TypeNetappVolumeInfo_KDATAPROTECTION:
    		return "kDataProtection"
        case TypeNetappVolumeInfo_KDATACACHE:
    		return "kDataCache"
        case TypeNetappVolumeInfo_KTMP:
    		return "kTmp"
        case TypeNetappVolumeInfo_KUNKNOWNTYPE:
    		return "kUnknownType"
        default:
        	return "kReadWrite"
    }
}

/**
 * Converts TypeNetappVolumeInfoEnum Array to its string Array representation
*/
func TypeNetappVolumeInfoEnumArrayToValue(typeNetappVolumeInfoEnum []TypeNetappVolumeInfoEnum) []string {
    convArray := make([]string,len( typeNetappVolumeInfoEnum))
    for i:=0; i<len(typeNetappVolumeInfoEnum);i++ {
        convArray[i] = TypeNetappVolumeInfoEnumToValue(typeNetappVolumeInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeNetappVolumeInfoEnumFromValue(value string) TypeNetappVolumeInfoEnum {
    switch value {
        case "kReadWrite":
            return TypeNetappVolumeInfo_KREADWRITE
        case "kLoadSharing":
            return TypeNetappVolumeInfo_KLOADSHARING
        case "kDataProtection":
            return TypeNetappVolumeInfo_KDATAPROTECTION
        case "kDataCache":
            return TypeNetappVolumeInfo_KDATACACHE
        case "kTmp":
            return TypeNetappVolumeInfo_KTMP
        case "kUnknownType":
            return TypeNetappVolumeInfo_KUNKNOWNTYPE
        default:
            return TypeNetappVolumeInfo_KREADWRITE
    }
}
