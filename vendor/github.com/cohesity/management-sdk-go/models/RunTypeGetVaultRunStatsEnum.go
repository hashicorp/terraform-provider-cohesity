// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RunTypeGetVaultRunStatsEnum enum
 */
type RunTypeGetVaultRunStatsEnum int

/**
 * Value collection for RunTypeGetVaultRunStatsEnum enum
 */
const (
    RunTypeGetVaultRunStats_KARCHIVE RunTypeGetVaultRunStatsEnum = 1 + iota
    RunTypeGetVaultRunStats_KRESTORE
    RunTypeGetVaultRunStats_KCLOUDSPIN
)

func (r RunTypeGetVaultRunStatsEnum) MarshalJSON() ([]byte, error) {
    s := RunTypeGetVaultRunStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *RunTypeGetVaultRunStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RunTypeGetVaultRunStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RunTypeGetVaultRunStatsEnum to its string representation
 */
func RunTypeGetVaultRunStatsEnumToValue(runTypeGetVaultRunStatsEnum RunTypeGetVaultRunStatsEnum) string {
    switch runTypeGetVaultRunStatsEnum {
        case RunTypeGetVaultRunStats_KARCHIVE:
    		return "kArchive"
        case RunTypeGetVaultRunStats_KRESTORE:
    		return "kRestore"
        case RunTypeGetVaultRunStats_KCLOUDSPIN:
    		return "kCloudSpin"
        default:
        	return "kArchive"
    }
}

/**
 * Converts RunTypeGetVaultRunStatsEnum Array to its string Array representation
*/
func RunTypeGetVaultRunStatsEnumArrayToValue(runTypeGetVaultRunStatsEnum []RunTypeGetVaultRunStatsEnum) []string {
    convArray := make([]string,len( runTypeGetVaultRunStatsEnum))
    for i:=0; i<len(runTypeGetVaultRunStatsEnum);i++ {
        convArray[i] = RunTypeGetVaultRunStatsEnumToValue(runTypeGetVaultRunStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunTypeGetVaultRunStatsEnumFromValue(value string) RunTypeGetVaultRunStatsEnum {
    switch value {
        case "kArchive":
            return RunTypeGetVaultRunStats_KARCHIVE
        case "kRestore":
            return RunTypeGetVaultRunStats_KRESTORE
        case "kCloudSpin":
            return RunTypeGetVaultRunStats_KCLOUDSPIN
        default:
            return RunTypeGetVaultRunStats_KARCHIVE
    }
}
