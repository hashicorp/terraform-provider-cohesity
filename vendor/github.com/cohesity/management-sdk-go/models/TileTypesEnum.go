package models

import(
    "encoding/json"
)

/**
 * Type definition for TileTypesEnum enum
 */
type TileTypesEnum int

/**
 * Value collection for TileTypesEnum enum
 */
const (
    TileTypes_KHEALTH TileTypesEnum = 1 + iota
    TileTypes_KJOBRUNS
    TileTypes_KRECOVERIES
    TileTypes_KPROTECTEDOBJECTS
    TileTypes_KPROTECTION
    TileTypes_KAUDITLOGS
    TileTypes_KIOPS
    TileTypes_KTHROUGHPUT
    TileTypes_KSTORAGEEFFICIENCY
)

func (r TileTypesEnum) MarshalJSON() ([]byte, error) { 
    s := TileTypesEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TileTypesEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TileTypesEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TileTypesEnum to its string representation
 */
func TileTypesEnumToValue(tileTypesEnum TileTypesEnum) string {
    switch tileTypesEnum {
        case TileTypes_KHEALTH:
    		return "kHealth"		
        case TileTypes_KJOBRUNS:
    		return "kJobRuns"		
        case TileTypes_KRECOVERIES:
    		return "kRecoveries"		
        case TileTypes_KPROTECTEDOBJECTS:
    		return "kProtectedObjects"		
        case TileTypes_KPROTECTION:
    		return "kProtection"		
        case TileTypes_KAUDITLOGS:
    		return "kAuditLogs"		
        case TileTypes_KIOPS:
    		return "kIops"		
        case TileTypes_KTHROUGHPUT:
    		return "kThroughput"		
        case TileTypes_KSTORAGEEFFICIENCY:
    		return "kStorageEfficiency"		
        default:
        	return "kHealth"
    }
}

/**
 * Converts TileTypesEnum Array to its string Array representation
*/
func TileTypesEnumArrayToValue(tileTypesEnum []TileTypesEnum) []string {
    convArray := make([]string,len( tileTypesEnum))
    for i:=0; i<len(tileTypesEnum);i++ {
        convArray[i] = TileTypesEnumToValue(tileTypesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TileTypesEnumFromValue(value string) TileTypesEnum {
    switch value {
        case "kHealth":
            return TileTypes_KHEALTH
        case "kJobRuns":
            return TileTypes_KJOBRUNS
        case "kRecoveries":
            return TileTypes_KRECOVERIES
        case "kProtectedObjects":
            return TileTypes_KPROTECTEDOBJECTS
        case "kProtection":
            return TileTypes_KPROTECTION
        case "kAuditLogs":
            return TileTypes_KAUDITLOGS
        case "kIops":
            return TileTypes_KIOPS
        case "kThroughput":
            return TileTypes_KTHROUGHPUT
        case "kStorageEfficiency":
            return TileTypes_KSTORAGEEFFICIENCY
        default:
            return TileTypes_KHEALTH
    }
}
