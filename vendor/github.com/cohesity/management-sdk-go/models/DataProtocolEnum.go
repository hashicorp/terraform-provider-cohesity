// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DataProtocolEnum enum
 */
type DataProtocolEnum int

/**
 * Value collection for DataProtocolEnum enum
 */
const (
    DataProtocol_KNFS DataProtocolEnum = 1 + iota
    DataProtocol_KCIFS
    DataProtocol_KISCSI
    DataProtocol_KFC
    DataProtocol_KFCACHE
    DataProtocol_KHTTP
    DataProtocol_KNDMP
    DataProtocol_KMANAGEMENT
)

func (r DataProtocolEnum) MarshalJSON() ([]byte, error) {
    s := DataProtocolEnumToValue(r)
    return json.Marshal(s)
}

func (r *DataProtocolEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DataProtocolEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DataProtocolEnum to its string representation
 */
func DataProtocolEnumToValue(dataProtocolEnum DataProtocolEnum) string {
    switch dataProtocolEnum {
        case DataProtocol_KNFS:
    		return "kNfs"
        case DataProtocol_KCIFS:
    		return "kCifs"
        case DataProtocol_KISCSI:
    		return "kIscsi"
        case DataProtocol_KFC:
    		return "kFc"
        case DataProtocol_KFCACHE:
    		return "kFcache"
        case DataProtocol_KHTTP:
    		return "kHttp"
        case DataProtocol_KNDMP:
    		return "kNdmp"
        case DataProtocol_KMANAGEMENT:
    		return "kManagement"
        default:
        	return "kNfs"
    }
}

/**
 * Converts DataProtocolEnum Array to its string Array representation
*/
func DataProtocolEnumArrayToValue(dataProtocolEnum []DataProtocolEnum) []string {
    convArray := make([]string,len( dataProtocolEnum))
    for i:=0; i<len(dataProtocolEnum);i++ {
        convArray[i] = DataProtocolEnumToValue(dataProtocolEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DataProtocolEnumFromValue(value string) DataProtocolEnum {
    switch value {
        case "kNfs":
            return DataProtocol_KNFS
        case "kCifs":
            return DataProtocol_KCIFS
        case "kIscsi":
            return DataProtocol_KISCSI
        case "kFc":
            return DataProtocol_KFC
        case "kFcache":
            return DataProtocol_KFCACHE
        case "kHttp":
            return DataProtocol_KHTTP
        case "kNdmp":
            return DataProtocol_KNDMP
        case "kManagement":
            return DataProtocol_KMANAGEMENT
        default:
            return DataProtocol_KNFS
    }
}
