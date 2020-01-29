package models

import(
    "encoding/json"
)

/**
 * Type definition for LastRunTypeEnum enum
 */
type LastRunTypeEnum int

/**
 * Value collection for LastRunTypeEnum enum
 */
const (
    LastRunType_KREGULAR LastRunTypeEnum = 1 + iota
    LastRunType_KFULL
    LastRunType_KLOG
    LastRunType_KSYSTEM
)

func (r LastRunTypeEnum) MarshalJSON() ([]byte, error) { 
    s := LastRunTypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *LastRunTypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  LastRunTypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts LastRunTypeEnum to its string representation
 */
func LastRunTypeEnumToValue(lastRunTypeEnum LastRunTypeEnum) string {
    switch lastRunTypeEnum {
        case LastRunType_KREGULAR:
    		return "kRegular"		
        case LastRunType_KFULL:
    		return "kFull"		
        case LastRunType_KLOG:
    		return "kLog"		
        case LastRunType_KSYSTEM:
    		return "kSystem"		
        default:
        	return "kRegular"
    }
}

/**
 * Converts LastRunTypeEnum Array to its string Array representation
*/
func LastRunTypeEnumArrayToValue(lastRunTypeEnum []LastRunTypeEnum) []string {
    convArray := make([]string,len( lastRunTypeEnum))
    for i:=0; i<len(lastRunTypeEnum);i++ {
        convArray[i] = LastRunTypeEnumToValue(lastRunTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func LastRunTypeEnumFromValue(value string) LastRunTypeEnum {
    switch value {
        case "kRegular":
            return LastRunType_KREGULAR
        case "kFull":
            return LastRunType_KFULL
        case "kLog":
            return LastRunType_KLOG
        case "kSystem":
            return LastRunType_KSYSTEM
        default:
            return LastRunType_KREGULAR
    }
}
