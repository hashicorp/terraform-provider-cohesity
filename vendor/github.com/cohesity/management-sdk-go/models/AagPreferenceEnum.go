// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AagPreferenceEnum enum
 */
type AagPreferenceEnum int

/**
 * Value collection for AagPreferenceEnum enum
 */
const (
    AagPreference_KPRIMARYREPLICAONLY AagPreferenceEnum = 1 + iota
    AagPreference_KSECONDARYREPLICAONLY
    AagPreference_KPREFERSECONDARYREPLICA
    AagPreference_KANYREPLICA
)

func (r AagPreferenceEnum) MarshalJSON() ([]byte, error) {
    s := AagPreferenceEnumToValue(r)
    return json.Marshal(s)
}

func (r *AagPreferenceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AagPreferenceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AagPreferenceEnum to its string representation
 */
func AagPreferenceEnumToValue(aagPreferenceEnum AagPreferenceEnum) string {
    switch aagPreferenceEnum {
        case AagPreference_KPRIMARYREPLICAONLY:
    		return "kPrimaryReplicaOnly"
        case AagPreference_KSECONDARYREPLICAONLY:
    		return "kSecondaryReplicaOnly"
        case AagPreference_KPREFERSECONDARYREPLICA:
    		return "kPreferSecondaryReplica"
        case AagPreference_KANYREPLICA:
    		return "kAnyReplica"
        default:
        	return "kPrimaryReplicaOnly"
    }
}

/**
 * Converts AagPreferenceEnum Array to its string Array representation
*/
func AagPreferenceEnumArrayToValue(aagPreferenceEnum []AagPreferenceEnum) []string {
    convArray := make([]string,len( aagPreferenceEnum))
    for i:=0; i<len(aagPreferenceEnum);i++ {
        convArray[i] = AagPreferenceEnumToValue(aagPreferenceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AagPreferenceEnumFromValue(value string) AagPreferenceEnum {
    switch value {
        case "kPrimaryReplicaOnly":
            return AagPreference_KPRIMARYREPLICAONLY
        case "kSecondaryReplicaOnly":
            return AagPreference_KSECONDARYREPLICAONLY
        case "kPreferSecondaryReplica":
            return AagPreference_KPREFERSECONDARYREPLICA
        case "kAnyReplica":
            return AagPreference_KANYREPLICA
        default:
            return AagPreference_KPRIMARYREPLICAONLY
    }
}
