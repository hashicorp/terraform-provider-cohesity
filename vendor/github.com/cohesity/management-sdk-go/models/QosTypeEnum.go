// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for QosTypeEnum enum
 */
type QosTypeEnum int

/**
 * Value collection for QosTypeEnum enum
 */
const (
    QosType_KBACKUPHDD QosTypeEnum = 1 + iota
    QosType_KBACKUPSSD
)

func (r QosTypeEnum) MarshalJSON() ([]byte, error) {
    s := QosTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *QosTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  QosTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts QosTypeEnum to its string representation
 */
func QosTypeEnumToValue(qosTypeEnum QosTypeEnum) string {
    switch qosTypeEnum {
        case QosType_KBACKUPHDD:
    		return "kBackupHDD"
        case QosType_KBACKUPSSD:
    		return "kBackupSSD"
        default:
        	return "kBackupHDD"
    }
}

/**
 * Converts QosTypeEnum Array to its string Array representation
*/
func QosTypeEnumArrayToValue(qosTypeEnum []QosTypeEnum) []string {
    convArray := make([]string,len( qosTypeEnum))
    for i:=0; i<len(qosTypeEnum);i++ {
        convArray[i] = QosTypeEnumToValue(qosTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func QosTypeEnumFromValue(value string) QosTypeEnum {
    switch value {
        case "kBackupHDD":
            return QosType_KBACKUPHDD
        case "kBackupSSD":
            return QosType_KBACKUPSSD
        default:
            return QosType_KBACKUPHDD
    }
}
