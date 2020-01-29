// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RunTypeGetVaultProviderStatsEnum enum
 */
type RunTypeGetVaultProviderStatsEnum int

/**
 * Value collection for RunTypeGetVaultProviderStatsEnum enum
 */
const (
    RunTypeGetVaultProviderStats_KARCHIVED RunTypeGetVaultProviderStatsEnum = 1 + iota
    RunTypeGetVaultProviderStats_KRESTORED
)

func (r RunTypeGetVaultProviderStatsEnum) MarshalJSON() ([]byte, error) {
    s := RunTypeGetVaultProviderStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *RunTypeGetVaultProviderStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RunTypeGetVaultProviderStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RunTypeGetVaultProviderStatsEnum to its string representation
 */
func RunTypeGetVaultProviderStatsEnumToValue(runTypeGetVaultProviderStatsEnum RunTypeGetVaultProviderStatsEnum) string {
    switch runTypeGetVaultProviderStatsEnum {
        case RunTypeGetVaultProviderStats_KARCHIVED:
    		return "kArchived"
        case RunTypeGetVaultProviderStats_KRESTORED:
    		return "kRestored"
        default:
        	return "kArchived"
    }
}

/**
 * Converts RunTypeGetVaultProviderStatsEnum Array to its string Array representation
*/
func RunTypeGetVaultProviderStatsEnumArrayToValue(runTypeGetVaultProviderStatsEnum []RunTypeGetVaultProviderStatsEnum) []string {
    convArray := make([]string,len( runTypeGetVaultProviderStatsEnum))
    for i:=0; i<len(runTypeGetVaultProviderStatsEnum);i++ {
        convArray[i] = RunTypeGetVaultProviderStatsEnumToValue(runTypeGetVaultProviderStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunTypeGetVaultProviderStatsEnumFromValue(value string) RunTypeGetVaultProviderStatsEnum {
    switch value {
        case "kArchived":
            return RunTypeGetVaultProviderStats_KARCHIVED
        case "kRestored":
            return RunTypeGetVaultProviderStats_KRESTORED
        default:
            return RunTypeGetVaultProviderStats_KARCHIVED
    }
}
