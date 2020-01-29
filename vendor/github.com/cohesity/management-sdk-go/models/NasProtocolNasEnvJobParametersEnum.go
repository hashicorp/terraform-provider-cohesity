// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for NasProtocolNasEnvJobParametersEnum enum
 */
type NasProtocolNasEnvJobParametersEnum int

/**
 * Value collection for NasProtocolNasEnvJobParametersEnum enum
 */
const (
    NasProtocolNasEnvJobParameters_KNFS3 NasProtocolNasEnvJobParametersEnum = 1 + iota
    NasProtocolNasEnvJobParameters_KCIFS1
)

func (r NasProtocolNasEnvJobParametersEnum) MarshalJSON() ([]byte, error) {
    s := NasProtocolNasEnvJobParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *NasProtocolNasEnvJobParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  NasProtocolNasEnvJobParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts NasProtocolNasEnvJobParametersEnum to its string representation
 */
func NasProtocolNasEnvJobParametersEnumToValue(nasProtocolNasEnvJobParametersEnum NasProtocolNasEnvJobParametersEnum) string {
    switch nasProtocolNasEnvJobParametersEnum {
        case NasProtocolNasEnvJobParameters_KNFS3:
    		return "kNfs3"
        case NasProtocolNasEnvJobParameters_KCIFS1:
    		return "kCifs1"
        default:
        	return "kNfs3"
    }
}

/**
 * Converts NasProtocolNasEnvJobParametersEnum Array to its string Array representation
*/
func NasProtocolNasEnvJobParametersEnumArrayToValue(nasProtocolNasEnvJobParametersEnum []NasProtocolNasEnvJobParametersEnum) []string {
    convArray := make([]string,len( nasProtocolNasEnvJobParametersEnum))
    for i:=0; i<len(nasProtocolNasEnvJobParametersEnum);i++ {
        convArray[i] = NasProtocolNasEnvJobParametersEnumToValue(nasProtocolNasEnvJobParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NasProtocolNasEnvJobParametersEnumFromValue(value string) NasProtocolNasEnvJobParametersEnum {
    switch value {
        case "kNfs3":
            return NasProtocolNasEnvJobParameters_KNFS3
        case "kCifs1":
            return NasProtocolNasEnvJobParameters_KCIFS1
        default:
            return NasProtocolNasEnvJobParameters_KNFS3
    }
}
