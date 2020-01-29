// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SqlServerDbStateEnum enum
 */
type SqlServerDbStateEnum int

/**
 * Value collection for SqlServerDbStateEnum enum
 */
const (
    SqlServerDbState_KONLINE SqlServerDbStateEnum = 1 + iota
    SqlServerDbState_KRESTORING
    SqlServerDbState_KRECOVERING
    SqlServerDbState_KRECOVERYPENDING
    SqlServerDbState_KSUSPECT
    SqlServerDbState_KEMERGENCY
    SqlServerDbState_KOFFLINE
    SqlServerDbState_KCOPYING
    SqlServerDbState_KOFFLINESECONDARY
)

func (r SqlServerDbStateEnum) MarshalJSON() ([]byte, error) {
    s := SqlServerDbStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *SqlServerDbStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SqlServerDbStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SqlServerDbStateEnum to its string representation
 */
func SqlServerDbStateEnumToValue(sqlServerDbStateEnum SqlServerDbStateEnum) string {
    switch sqlServerDbStateEnum {
        case SqlServerDbState_KONLINE:
    		return "kOnline"
        case SqlServerDbState_KRESTORING:
    		return "kRestoring"
        case SqlServerDbState_KRECOVERING:
    		return "kRecovering"
        case SqlServerDbState_KRECOVERYPENDING:
    		return "kRecoveryPending"
        case SqlServerDbState_KSUSPECT:
    		return "kSuspect"
        case SqlServerDbState_KEMERGENCY:
    		return "kEmergency"
        case SqlServerDbState_KOFFLINE:
    		return "kOffline"
        case SqlServerDbState_KCOPYING:
    		return "kCopying"
        case SqlServerDbState_KOFFLINESECONDARY:
    		return "kOfflineSecondary"
        default:
        	return "kOnline"
    }
}

/**
 * Converts SqlServerDbStateEnum Array to its string Array representation
*/
func SqlServerDbStateEnumArrayToValue(sqlServerDbStateEnum []SqlServerDbStateEnum) []string {
    convArray := make([]string,len( sqlServerDbStateEnum))
    for i:=0; i<len(sqlServerDbStateEnum);i++ {
        convArray[i] = SqlServerDbStateEnumToValue(sqlServerDbStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SqlServerDbStateEnumFromValue(value string) SqlServerDbStateEnum {
    switch value {
        case "kOnline":
            return SqlServerDbState_KONLINE
        case "kRestoring":
            return SqlServerDbState_KRESTORING
        case "kRecovering":
            return SqlServerDbState_KRECOVERING
        case "kRecoveryPending":
            return SqlServerDbState_KRECOVERYPENDING
        case "kSuspect":
            return SqlServerDbState_KSUSPECT
        case "kEmergency":
            return SqlServerDbState_KEMERGENCY
        case "kOffline":
            return SqlServerDbState_KOFFLINE
        case "kCopying":
            return SqlServerDbState_KCOPYING
        case "kOfflineSecondary":
            return SqlServerDbState_KOFFLINESECONDARY
        default:
            return SqlServerDbState_KONLINE
    }
}
