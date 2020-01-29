// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TierTypeGoogleCloudCredentialsEnum enum
 */
type TierTypeGoogleCloudCredentialsEnum int

/**
 * Value collection for TierTypeGoogleCloudCredentialsEnum enum
 */
const (
    TierTypeGoogleCloudCredentials_KGOOGLESTANDARD TierTypeGoogleCloudCredentialsEnum = 1 + iota
    TierTypeGoogleCloudCredentials_KGOOGLENEARLINE
    TierTypeGoogleCloudCredentials_KGOOGLECOLDLINE
    TierTypeGoogleCloudCredentials_KGOOGLEREGIONAL
    TierTypeGoogleCloudCredentials_KGOOGLEMULTIREGIONAL
)

func (r TierTypeGoogleCloudCredentialsEnum) MarshalJSON() ([]byte, error) {
    s := TierTypeGoogleCloudCredentialsEnumToValue(r)
    return json.Marshal(s)
}

func (r *TierTypeGoogleCloudCredentialsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TierTypeGoogleCloudCredentialsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TierTypeGoogleCloudCredentialsEnum to its string representation
 */
func TierTypeGoogleCloudCredentialsEnumToValue(tierTypeGoogleCloudCredentialsEnum TierTypeGoogleCloudCredentialsEnum) string {
    switch tierTypeGoogleCloudCredentialsEnum {
        case TierTypeGoogleCloudCredentials_KGOOGLESTANDARD:
    		return "kGoogleStandard"
        case TierTypeGoogleCloudCredentials_KGOOGLENEARLINE:
    		return "kGoogleNearline"
        case TierTypeGoogleCloudCredentials_KGOOGLECOLDLINE:
    		return "kGoogleColdline"
        case TierTypeGoogleCloudCredentials_KGOOGLEREGIONAL:
    		return "kGoogleRegional"
        case TierTypeGoogleCloudCredentials_KGOOGLEMULTIREGIONAL:
    		return "kGoogleMultiRegional"
        default:
        	return "kGoogleStandard"
    }
}

/**
 * Converts TierTypeGoogleCloudCredentialsEnum Array to its string Array representation
*/
func TierTypeGoogleCloudCredentialsEnumArrayToValue(tierTypeGoogleCloudCredentialsEnum []TierTypeGoogleCloudCredentialsEnum) []string {
    convArray := make([]string,len( tierTypeGoogleCloudCredentialsEnum))
    for i:=0; i<len(tierTypeGoogleCloudCredentialsEnum);i++ {
        convArray[i] = TierTypeGoogleCloudCredentialsEnumToValue(tierTypeGoogleCloudCredentialsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierTypeGoogleCloudCredentialsEnumFromValue(value string) TierTypeGoogleCloudCredentialsEnum {
    switch value {
        case "kGoogleStandard":
            return TierTypeGoogleCloudCredentials_KGOOGLESTANDARD
        case "kGoogleNearline":
            return TierTypeGoogleCloudCredentials_KGOOGLENEARLINE
        case "kGoogleColdline":
            return TierTypeGoogleCloudCredentials_KGOOGLECOLDLINE
        case "kGoogleRegional":
            return TierTypeGoogleCloudCredentials_KGOOGLEREGIONAL
        case "kGoogleMultiRegional":
            return TierTypeGoogleCloudCredentials_KGOOGLEMULTIREGIONAL
        default:
            return TierTypeGoogleCloudCredentials_KGOOGLESTANDARD
    }
}
