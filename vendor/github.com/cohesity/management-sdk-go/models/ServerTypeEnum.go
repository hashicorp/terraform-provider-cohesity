// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ServerTypeEnum enum
 */
type ServerTypeEnum int

/**
 * Value collection for ServerTypeEnum enum
 */
const (
    ServerType_KINTERNALKMS ServerTypeEnum = 1 + iota
    ServerType_KAWSKMS
    ServerType_KCRYPTSOFTKMS
)

func (r ServerTypeEnum) MarshalJSON() ([]byte, error) {
    s := ServerTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ServerTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ServerTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ServerTypeEnum to its string representation
 */
func ServerTypeEnumToValue(serverTypeEnum ServerTypeEnum) string {
    switch serverTypeEnum {
        case ServerType_KINTERNALKMS:
    		return "kInternalKms"
        case ServerType_KAWSKMS:
    		return "kAwsKms"
        case ServerType_KCRYPTSOFTKMS:
    		return "kCryptsoftKms"
        default:
        	return "kInternalKms"
    }
}

/**
 * Converts ServerTypeEnum Array to its string Array representation
*/
func ServerTypeEnumArrayToValue(serverTypeEnum []ServerTypeEnum) []string {
    convArray := make([]string,len( serverTypeEnum))
    for i:=0; i<len(serverTypeEnum);i++ {
        convArray[i] = ServerTypeEnumToValue(serverTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ServerTypeEnumFromValue(value string) ServerTypeEnum {
    switch value {
        case "kInternalKms":
            return ServerType_KINTERNALKMS
        case "kAwsKms":
            return ServerType_KAWSKMS
        case "kCryptsoftKms":
            return ServerType_KCRYPTSOFTKMS
        default:
            return ServerType_KINTERNALKMS
    }
}
