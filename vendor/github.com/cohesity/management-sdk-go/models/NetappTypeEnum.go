// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for NetappTypeEnum enum
 */
type NetappTypeEnum int

/**
 * Value collection for NetappTypeEnum enum
 */
const (
    NetappType_KCLUSTER NetappTypeEnum = 1 + iota
    NetappType_KVSERVER
    NetappType_KVOLUME
)

func (r NetappTypeEnum) MarshalJSON() ([]byte, error) {
    s := NetappTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *NetappTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  NetappTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts NetappTypeEnum to its string representation
 */
func NetappTypeEnumToValue(netappTypeEnum NetappTypeEnum) string {
    switch netappTypeEnum {
        case NetappType_KCLUSTER:
    		return "kCluster"
        case NetappType_KVSERVER:
    		return "kVserver"
        case NetappType_KVOLUME:
    		return "kVolume"
        default:
        	return "kCluster"
    }
}

/**
 * Converts NetappTypeEnum Array to its string Array representation
*/
func NetappTypeEnumArrayToValue(netappTypeEnum []NetappTypeEnum) []string {
    convArray := make([]string,len( netappTypeEnum))
    for i:=0; i<len(netappTypeEnum);i++ {
        convArray[i] = NetappTypeEnumToValue(netappTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NetappTypeEnumFromValue(value string) NetappTypeEnum {
    switch value {
        case "kCluster":
            return NetappType_KCLUSTER
        case "kVserver":
            return NetappType_KVSERVER
        case "kVolume":
            return NetappType_KVOLUME
        default:
            return NetappType_KCLUSTER
    }
}
