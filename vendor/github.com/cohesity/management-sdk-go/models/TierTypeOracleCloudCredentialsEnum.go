// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TierTypeOracleCloudCredentialsEnum enum
 */
type TierTypeOracleCloudCredentialsEnum int

/**
 * Value collection for TierTypeOracleCloudCredentialsEnum enum
 */
const (
    TierTypeOracleCloudCredentials_KORACLETIERSTANDARD TierTypeOracleCloudCredentialsEnum = 1 + iota
    TierTypeOracleCloudCredentials_KORACLETIERARCHIVE
)

func (r TierTypeOracleCloudCredentialsEnum) MarshalJSON() ([]byte, error) {
    s := TierTypeOracleCloudCredentialsEnumToValue(r)
    return json.Marshal(s)
}

func (r *TierTypeOracleCloudCredentialsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TierTypeOracleCloudCredentialsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TierTypeOracleCloudCredentialsEnum to its string representation
 */
func TierTypeOracleCloudCredentialsEnumToValue(tierTypeOracleCloudCredentialsEnum TierTypeOracleCloudCredentialsEnum) string {
    switch tierTypeOracleCloudCredentialsEnum {
        case TierTypeOracleCloudCredentials_KORACLETIERSTANDARD:
    		return "kOracleTierStandard"
        case TierTypeOracleCloudCredentials_KORACLETIERARCHIVE:
    		return "kOracleTierArchive"
        default:
        	return "kOracleTierStandard"
    }
}

/**
 * Converts TierTypeOracleCloudCredentialsEnum Array to its string Array representation
*/
func TierTypeOracleCloudCredentialsEnumArrayToValue(tierTypeOracleCloudCredentialsEnum []TierTypeOracleCloudCredentialsEnum) []string {
    convArray := make([]string,len( tierTypeOracleCloudCredentialsEnum))
    for i:=0; i<len(tierTypeOracleCloudCredentialsEnum);i++ {
        convArray[i] = TierTypeOracleCloudCredentialsEnumToValue(tierTypeOracleCloudCredentialsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierTypeOracleCloudCredentialsEnumFromValue(value string) TierTypeOracleCloudCredentialsEnum {
    switch value {
        case "kOracleTierStandard":
            return TierTypeOracleCloudCredentials_KORACLETIERSTANDARD
        case "kOracleTierArchive":
            return TierTypeOracleCloudCredentials_KORACLETIERARCHIVE
        default:
            return TierTypeOracleCloudCredentials_KORACLETIERSTANDARD
    }
}
