package models

import(
    "encoding/json"
)

/**
 * Type definition for BackupType1Enum enum
 */
type BackupType1Enum int

/**
 * Value collection for BackupType1Enum enum
 */
const (
    BackupType1_KSQLVSSVOLUME BackupType1Enum = 1 + iota
    BackupType1_KSQLVSSFILE
)

func (r BackupType1Enum) MarshalJSON() ([]byte, error) { 
    s := BackupType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *BackupType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  BackupType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts BackupType1Enum to its string representation
 */
func BackupType1EnumToValue(backupType1Enum BackupType1Enum) string {
    switch backupType1Enum {
        case BackupType1_KSQLVSSVOLUME:
    		return "kSqlVSSVolume"		
        case BackupType1_KSQLVSSFILE:
    		return "kSqlVSSFile"		
        default:
        	return "kSqlVSSVolume"
    }
}

/**
 * Converts BackupType1Enum Array to its string Array representation
*/
func BackupType1EnumArrayToValue(backupType1Enum []BackupType1Enum) []string {
    convArray := make([]string,len( backupType1Enum))
    for i:=0; i<len(backupType1Enum);i++ {
        convArray[i] = BackupType1EnumToValue(backupType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BackupType1EnumFromValue(value string) BackupType1Enum {
    switch value {
        case "kSqlVSSVolume":
            return BackupType1_KSQLVSSVOLUME
        case "kSqlVSSFile":
            return BackupType1_KSQLVSSFILE
        default:
            return BackupType1_KSQLVSSVOLUME
    }
}
