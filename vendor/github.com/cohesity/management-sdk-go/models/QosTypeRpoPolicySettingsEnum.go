// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for QosTypeRpoPolicySettingsEnum enum
 */
type QosTypeRpoPolicySettingsEnum int

/**
 * Value collection for QosTypeRpoPolicySettingsEnum enum
 */
const (
    QosTypeRpoPolicySettings_KBACKUPHDD QosTypeRpoPolicySettingsEnum = 1 + iota
    QosTypeRpoPolicySettings_KBACKUPSSD
)

func (r QosTypeRpoPolicySettingsEnum) MarshalJSON() ([]byte, error) {
    s := QosTypeRpoPolicySettingsEnumToValue(r)
    return json.Marshal(s)
}

func (r *QosTypeRpoPolicySettingsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  QosTypeRpoPolicySettingsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts QosTypeRpoPolicySettingsEnum to its string representation
 */
func QosTypeRpoPolicySettingsEnumToValue(qosTypeRpoPolicySettingsEnum QosTypeRpoPolicySettingsEnum) string {
    switch qosTypeRpoPolicySettingsEnum {
        case QosTypeRpoPolicySettings_KBACKUPHDD:
    		return "kBackupHDD"
        case QosTypeRpoPolicySettings_KBACKUPSSD:
    		return "kBackupSSD"
        default:
        	return "kBackupHDD"
    }
}

/**
 * Converts QosTypeRpoPolicySettingsEnum Array to its string Array representation
*/
func QosTypeRpoPolicySettingsEnumArrayToValue(qosTypeRpoPolicySettingsEnum []QosTypeRpoPolicySettingsEnum) []string {
    convArray := make([]string,len( qosTypeRpoPolicySettingsEnum))
    for i:=0; i<len(qosTypeRpoPolicySettingsEnum);i++ {
        convArray[i] = QosTypeRpoPolicySettingsEnumToValue(qosTypeRpoPolicySettingsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func QosTypeRpoPolicySettingsEnumFromValue(value string) QosTypeRpoPolicySettingsEnum {
    switch value {
        case "kBackupHDD":
            return QosTypeRpoPolicySettings_KBACKUPHDD
        case "kBackupSSD":
            return QosTypeRpoPolicySettings_KBACKUPSSD
        default:
            return QosTypeRpoPolicySettings_KBACKUPHDD
    }
}
