package models

import(
    "encoding/json"
)

/**
 * Type definition for QosType1Enum enum
 */
type QosType1Enum int

/**
 * Value collection for QosType1Enum enum
 */
const (
    QosType1_KBACKUPHDD QosType1Enum = 1 + iota
    QosType1_KBACKUPSSD
)

func (r QosType1Enum) MarshalJSON() ([]byte, error) { 
    s := QosType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *QosType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  QosType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts QosType1Enum to its string representation
 */
func QosType1EnumToValue(qosType1Enum QosType1Enum) string {
    switch qosType1Enum {
        case QosType1_KBACKUPHDD:
    		return "kBackupHDD"		
        case QosType1_KBACKUPSSD:
    		return "kBackupSSD"		
        default:
        	return "kBackupHDD"
    }
}

/**
 * Converts QosType1Enum Array to its string Array representation
*/
func QosType1EnumArrayToValue(qosType1Enum []QosType1Enum) []string {
    convArray := make([]string,len( qosType1Enum))
    for i:=0; i<len(qosType1Enum);i++ {
        convArray[i] = QosType1EnumToValue(qosType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func QosType1EnumFromValue(value string) QosType1Enum {
    switch value {
        case "kBackupHDD":
            return QosType1_KBACKUPHDD
        case "kBackupSSD":
            return QosType1_KBACKUPSSD
        default:
            return QosType1_KBACKUPHDD
    }
}
