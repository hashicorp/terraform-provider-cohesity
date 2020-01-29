// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RecoveryModelEnum enum
 */
type RecoveryModelEnum int

/**
 * Value collection for RecoveryModelEnum enum
 */
const (
    RecoveryModel_KSIMPLERECOVERYMODEL RecoveryModelEnum = 1 + iota
    RecoveryModel_KFULLRECOVERYMODEL
    RecoveryModel_KBULKLOGGEDRECOVERYMODEL
)

func (r RecoveryModelEnum) MarshalJSON() ([]byte, error) {
    s := RecoveryModelEnumToValue(r)
    return json.Marshal(s)
}

func (r *RecoveryModelEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RecoveryModelEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RecoveryModelEnum to its string representation
 */
func RecoveryModelEnumToValue(recoveryModelEnum RecoveryModelEnum) string {
    switch recoveryModelEnum {
        case RecoveryModel_KSIMPLERECOVERYMODEL:
    		return "kSimpleRecoveryModel"
        case RecoveryModel_KFULLRECOVERYMODEL:
    		return "kFullRecoveryModel"
        case RecoveryModel_KBULKLOGGEDRECOVERYMODEL:
    		return "kBulkLoggedRecoveryModel"
        default:
        	return "kSimpleRecoveryModel"
    }
}

/**
 * Converts RecoveryModelEnum Array to its string Array representation
*/
func RecoveryModelEnumArrayToValue(recoveryModelEnum []RecoveryModelEnum) []string {
    convArray := make([]string,len( recoveryModelEnum))
    for i:=0; i<len(recoveryModelEnum);i++ {
        convArray[i] = RecoveryModelEnumToValue(recoveryModelEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RecoveryModelEnumFromValue(value string) RecoveryModelEnum {
    switch value {
        case "kSimpleRecoveryModel":
            return RecoveryModel_KSIMPLERECOVERYMODEL
        case "kFullRecoveryModel":
            return RecoveryModel_KFULLRECOVERYMODEL
        case "kBulkLoggedRecoveryModel":
            return RecoveryModel_KBULKLOGGEDRECOVERYMODEL
        default:
            return RecoveryModel_KSIMPLERECOVERYMODEL
    }
}
