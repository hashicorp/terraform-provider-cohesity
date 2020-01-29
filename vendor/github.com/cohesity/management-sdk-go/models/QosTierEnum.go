package models

import(
    "encoding/json"
)

/**
 * Type definition for QosTierEnum enum
 */
type QosTierEnum int

/**
 * Value collection for QosTierEnum enum
 */
const (
    QosTier_KLOW QosTierEnum = 1 + iota
    QosTier_KMEDIUM
    QosTier_KHIGH
)

func (r QosTierEnum) MarshalJSON() ([]byte, error) { 
    s := QosTierEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *QosTierEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  QosTierEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts QosTierEnum to its string representation
 */
func QosTierEnumToValue(qosTierEnum QosTierEnum) string {
    switch qosTierEnum {
        case QosTier_KLOW:
    		return "kLow"		
        case QosTier_KMEDIUM:
    		return "kMedium"		
        case QosTier_KHIGH:
    		return "kHigh"		
        default:
        	return "kLow"
    }
}

/**
 * Converts QosTierEnum Array to its string Array representation
*/
func QosTierEnumArrayToValue(qosTierEnum []QosTierEnum) []string {
    convArray := make([]string,len( qosTierEnum))
    for i:=0; i<len(qosTierEnum);i++ {
        convArray[i] = QosTierEnumToValue(qosTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func QosTierEnumFromValue(value string) QosTierEnum {
    switch value {
        case "kLow":
            return QosTier_KLOW
        case "kMedium":
            return QosTier_KMEDIUM
        case "kHigh":
            return QosTier_KHIGH
        default:
            return QosTier_KLOW
    }
}
