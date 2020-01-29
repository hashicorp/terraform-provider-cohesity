// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TierTypeAzureCloudCredentialsEnum enum
 */
type TierTypeAzureCloudCredentialsEnum int

/**
 * Value collection for TierTypeAzureCloudCredentialsEnum enum
 */
const (
    TierTypeAzureCloudCredentials_KAZURETIERHOT TierTypeAzureCloudCredentialsEnum = 1 + iota
    TierTypeAzureCloudCredentials_KAZURETIERCOOL
    TierTypeAzureCloudCredentials_KAZURETIERARCHIVE
)

func (r TierTypeAzureCloudCredentialsEnum) MarshalJSON() ([]byte, error) {
    s := TierTypeAzureCloudCredentialsEnumToValue(r)
    return json.Marshal(s)
}

func (r *TierTypeAzureCloudCredentialsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TierTypeAzureCloudCredentialsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TierTypeAzureCloudCredentialsEnum to its string representation
 */
func TierTypeAzureCloudCredentialsEnumToValue(tierTypeAzureCloudCredentialsEnum TierTypeAzureCloudCredentialsEnum) string {
    switch tierTypeAzureCloudCredentialsEnum {
        case TierTypeAzureCloudCredentials_KAZURETIERHOT:
    		return "kAzureTierHot"
        case TierTypeAzureCloudCredentials_KAZURETIERCOOL:
    		return "kAzureTierCool"
        case TierTypeAzureCloudCredentials_KAZURETIERARCHIVE:
    		return "kAzureTierArchive"
        default:
        	return "kAzureTierHot"
    }
}

/**
 * Converts TierTypeAzureCloudCredentialsEnum Array to its string Array representation
*/
func TierTypeAzureCloudCredentialsEnumArrayToValue(tierTypeAzureCloudCredentialsEnum []TierTypeAzureCloudCredentialsEnum) []string {
    convArray := make([]string,len( tierTypeAzureCloudCredentialsEnum))
    for i:=0; i<len(tierTypeAzureCloudCredentialsEnum);i++ {
        convArray[i] = TierTypeAzureCloudCredentialsEnumToValue(tierTypeAzureCloudCredentialsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierTypeAzureCloudCredentialsEnumFromValue(value string) TierTypeAzureCloudCredentialsEnum {
    switch value {
        case "kAzureTierHot":
            return TierTypeAzureCloudCredentials_KAZURETIERHOT
        case "kAzureTierCool":
            return TierTypeAzureCloudCredentials_KAZURETIERCOOL
        case "kAzureTierArchive":
            return TierTypeAzureCloudCredentials_KAZURETIERARCHIVE
        default:
            return TierTypeAzureCloudCredentials_KAZURETIERHOT
    }
}
