// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for FileSelectionPolicyEnum enum
 */
type FileSelectionPolicyEnum int

/**
 * Value collection for FileSelectionPolicyEnum enum
 */
const (
    FileSelectionPolicy_KOLDERTHAN FileSelectionPolicyEnum = 1 + iota
    FileSelectionPolicy_KLASTACCESSED
    FileSelectionPolicy_KLASTMODIFIED
)

func (r FileSelectionPolicyEnum) MarshalJSON() ([]byte, error) {
    s := FileSelectionPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *FileSelectionPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  FileSelectionPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts FileSelectionPolicyEnum to its string representation
 */
func FileSelectionPolicyEnumToValue(fileSelectionPolicyEnum FileSelectionPolicyEnum) string {
    switch fileSelectionPolicyEnum {
        case FileSelectionPolicy_KOLDERTHAN:
    		return "kOlderThan"
        case FileSelectionPolicy_KLASTACCESSED:
    		return "kLastAccessed"
        case FileSelectionPolicy_KLASTMODIFIED:
    		return "kLastModified"
        default:
        	return "kOlderThan"
    }
}

/**
 * Converts FileSelectionPolicyEnum Array to its string Array representation
*/
func FileSelectionPolicyEnumArrayToValue(fileSelectionPolicyEnum []FileSelectionPolicyEnum) []string {
    convArray := make([]string,len( fileSelectionPolicyEnum))
    for i:=0; i<len(fileSelectionPolicyEnum);i++ {
        convArray[i] = FileSelectionPolicyEnumToValue(fileSelectionPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func FileSelectionPolicyEnumFromValue(value string) FileSelectionPolicyEnum {
    switch value {
        case "kOlderThan":
            return FileSelectionPolicy_KOLDERTHAN
        case "kLastAccessed":
            return FileSelectionPolicy_KLASTACCESSED
        case "kLastModified":
            return FileSelectionPolicy_KLASTMODIFIED
        default:
            return FileSelectionPolicy_KOLDERTHAN
    }
}
