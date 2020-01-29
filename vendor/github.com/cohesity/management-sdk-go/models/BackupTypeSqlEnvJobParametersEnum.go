// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BackupTypeSqlEnvJobParametersEnum enum
 */
type BackupTypeSqlEnvJobParametersEnum int

/**
 * Value collection for BackupTypeSqlEnvJobParametersEnum enum
 */
const (
    BackupTypeSqlEnvJobParameters_KSQLVSSVOLUME BackupTypeSqlEnvJobParametersEnum = 1 + iota
    BackupTypeSqlEnvJobParameters_KSQLVSSFILE
)

func (r BackupTypeSqlEnvJobParametersEnum) MarshalJSON() ([]byte, error) {
    s := BackupTypeSqlEnvJobParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *BackupTypeSqlEnvJobParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BackupTypeSqlEnvJobParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BackupTypeSqlEnvJobParametersEnum to its string representation
 */
func BackupTypeSqlEnvJobParametersEnumToValue(backupTypeSqlEnvJobParametersEnum BackupTypeSqlEnvJobParametersEnum) string {
    switch backupTypeSqlEnvJobParametersEnum {
        case BackupTypeSqlEnvJobParameters_KSQLVSSVOLUME:
    		return "kSqlVSSVolume"
        case BackupTypeSqlEnvJobParameters_KSQLVSSFILE:
    		return "kSqlVSSFile"
        default:
        	return "kSqlVSSVolume"
    }
}

/**
 * Converts BackupTypeSqlEnvJobParametersEnum Array to its string Array representation
*/
func BackupTypeSqlEnvJobParametersEnumArrayToValue(backupTypeSqlEnvJobParametersEnum []BackupTypeSqlEnvJobParametersEnum) []string {
    convArray := make([]string,len( backupTypeSqlEnvJobParametersEnum))
    for i:=0; i<len(backupTypeSqlEnvJobParametersEnum);i++ {
        convArray[i] = BackupTypeSqlEnvJobParametersEnumToValue(backupTypeSqlEnvJobParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BackupTypeSqlEnvJobParametersEnumFromValue(value string) BackupTypeSqlEnvJobParametersEnum {
    switch value {
        case "kSqlVSSVolume":
            return BackupTypeSqlEnvJobParameters_KSQLVSSVOLUME
        case "kSqlVSSFile":
            return BackupTypeSqlEnvJobParameters_KSQLVSSFILE
        default:
            return BackupTypeSqlEnvJobParameters_KSQLVSSVOLUME
    }
}
