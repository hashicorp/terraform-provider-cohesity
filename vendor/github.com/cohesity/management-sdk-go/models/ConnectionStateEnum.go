// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ConnectionStateEnum enum
 */
type ConnectionStateEnum int

/**
 * Value collection for ConnectionStateEnum enum
 */
const (
    ConnectionState_KCONNECTED ConnectionStateEnum = 1 + iota
    ConnectionState_KDISCONNECTED
    ConnectionState_KINACCESSIBLE
    ConnectionState_KINVALID
    ConnectionState_KORPHANED
    ConnectionState_KNOTRESPONDING
)

func (r ConnectionStateEnum) MarshalJSON() ([]byte, error) {
    s := ConnectionStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *ConnectionStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ConnectionStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ConnectionStateEnum to its string representation
 */
func ConnectionStateEnumToValue(connectionStateEnum ConnectionStateEnum) string {
    switch connectionStateEnum {
        case ConnectionState_KCONNECTED:
    		return "kConnected"
        case ConnectionState_KDISCONNECTED:
    		return "kDisconnected"
        case ConnectionState_KINACCESSIBLE:
    		return "kInaccessible"
        case ConnectionState_KINVALID:
    		return "kInvalid"
        case ConnectionState_KORPHANED:
    		return "kOrphaned"
        case ConnectionState_KNOTRESPONDING:
    		return "kNotResponding"
        default:
        	return "kConnected"
    }
}

/**
 * Converts ConnectionStateEnum Array to its string Array representation
*/
func ConnectionStateEnumArrayToValue(connectionStateEnum []ConnectionStateEnum) []string {
    convArray := make([]string,len( connectionStateEnum))
    for i:=0; i<len(connectionStateEnum);i++ {
        convArray[i] = ConnectionStateEnumToValue(connectionStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ConnectionStateEnumFromValue(value string) ConnectionStateEnum {
    switch value {
        case "kConnected":
            return ConnectionState_KCONNECTED
        case "kDisconnected":
            return ConnectionState_KDISCONNECTED
        case "kInaccessible":
            return ConnectionState_KINACCESSIBLE
        case "kInvalid":
            return ConnectionState_KINVALID
        case "kOrphaned":
            return ConnectionState_KORPHANED
        case "kNotResponding":
            return ConnectionState_KNOTRESPONDING
        default:
            return ConnectionState_KCONNECTED
    }
}
