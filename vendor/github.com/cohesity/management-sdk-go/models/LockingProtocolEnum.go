// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for LockingProtocolEnum enum
 */
type LockingProtocolEnum int

/**
 * Value collection for LockingProtocolEnum enum
 */
const (
    LockingProtocol_KSETREADONLY LockingProtocolEnum = 1 + iota
    LockingProtocol_KSETATIME
)

func (r LockingProtocolEnum) MarshalJSON() ([]byte, error) {
    s := LockingProtocolEnumToValue(r)
    return json.Marshal(s)
}

func (r *LockingProtocolEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  LockingProtocolEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts LockingProtocolEnum to its string representation
 */
func LockingProtocolEnumToValue(lockingProtocolEnum LockingProtocolEnum) string {
    switch lockingProtocolEnum {
        case LockingProtocol_KSETREADONLY:
    		return "kSetReadOnly"
        case LockingProtocol_KSETATIME:
    		return "kSetAtime"
        default:
        	return "kSetReadOnly"
    }
}

/**
 * Converts LockingProtocolEnum Array to its string Array representation
*/
func LockingProtocolEnumArrayToValue(lockingProtocolEnum []LockingProtocolEnum) []string {
    convArray := make([]string,len( lockingProtocolEnum))
    for i:=0; i<len(lockingProtocolEnum);i++ {
        convArray[i] = LockingProtocolEnumToValue(lockingProtocolEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func LockingProtocolEnumFromValue(value string) LockingProtocolEnum {
    switch value {
        case "kSetReadOnly":
            return LockingProtocol_KSETREADONLY
        case "kSetAtime":
            return LockingProtocol_KSETATIME
        default:
            return LockingProtocol_KSETREADONLY
    }
}
