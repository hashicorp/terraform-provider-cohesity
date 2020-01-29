// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CurrentOperationEnum enum
 */
type CurrentOperationEnum int

/**
 * Value collection for CurrentOperationEnum enum
 */
const (
    CurrentOperation_KREMOVENODE CurrentOperationEnum = 1 + iota
    CurrentOperation_KUPGRADE
    CurrentOperation_KNONE
    CurrentOperation_KDESTROY
    CurrentOperation_KCLEAN
    CurrentOperation_KRESTARTSERVICES
)

func (r CurrentOperationEnum) MarshalJSON() ([]byte, error) {
    s := CurrentOperationEnumToValue(r)
    return json.Marshal(s)
}

func (r *CurrentOperationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CurrentOperationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CurrentOperationEnum to its string representation
 */
func CurrentOperationEnumToValue(currentOperationEnum CurrentOperationEnum) string {
    switch currentOperationEnum {
        case CurrentOperation_KREMOVENODE:
    		return "kRemoveNode"
        case CurrentOperation_KUPGRADE:
    		return "kUpgrade"
        case CurrentOperation_KNONE:
    		return "kNone"
        case CurrentOperation_KDESTROY:
    		return "kDestroy"
        case CurrentOperation_KCLEAN:
    		return "kClean"
        case CurrentOperation_KRESTARTSERVICES:
    		return "kRestartServices"
        default:
        	return "kRemoveNode"
    }
}

/**
 * Converts CurrentOperationEnum Array to its string Array representation
*/
func CurrentOperationEnumArrayToValue(currentOperationEnum []CurrentOperationEnum) []string {
    convArray := make([]string,len( currentOperationEnum))
    for i:=0; i<len(currentOperationEnum);i++ {
        convArray[i] = CurrentOperationEnumToValue(currentOperationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CurrentOperationEnumFromValue(value string) CurrentOperationEnum {
    switch value {
        case "kRemoveNode":
            return CurrentOperation_KREMOVENODE
        case "kUpgrade":
            return CurrentOperation_KUPGRADE
        case "kNone":
            return CurrentOperation_KNONE
        case "kDestroy":
            return CurrentOperation_KDESTROY
        case "kClean":
            return CurrentOperation_KCLEAN
        case "kRestartServices":
            return CurrentOperation_KRESTARTSERVICES
        default:
            return CurrentOperation_KREMOVENODE
    }
}
