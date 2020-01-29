// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectStatusEnum enum
 */
type ObjectStatusEnum int

/**
 * Value collection for ObjectStatusEnum enum
 */
const (
    ObjectStatus_KFILESCLONED ObjectStatusEnum = 1 + iota
    ObjectStatus_KFETCHEDENTITYINFO
    ObjectStatus_KVMCREATED
    ObjectStatus_KRELOCATIONSTARTED
    ObjectStatus_KFINISHED
    ObjectStatus_KABORTED
    ObjectStatus_KDATACOPYSTARTED
    ObjectStatus_KINPROGRESS
)

func (r ObjectStatusEnum) MarshalJSON() ([]byte, error) {
    s := ObjectStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *ObjectStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ObjectStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ObjectStatusEnum to its string representation
 */
func ObjectStatusEnumToValue(objectStatusEnum ObjectStatusEnum) string {
    switch objectStatusEnum {
        case ObjectStatus_KFILESCLONED:
    		return "kFilesCloned"
        case ObjectStatus_KFETCHEDENTITYINFO:
    		return "kFetchedEntityInfo"
        case ObjectStatus_KVMCREATED:
    		return "kVMCreated"
        case ObjectStatus_KRELOCATIONSTARTED:
    		return "kRelocationStarted"
        case ObjectStatus_KFINISHED:
    		return "kFinished"
        case ObjectStatus_KABORTED:
    		return "kAborted"
        case ObjectStatus_KDATACOPYSTARTED:
    		return "kDataCopyStarted"
        case ObjectStatus_KINPROGRESS:
    		return "kInProgress"
        default:
        	return "kFilesCloned"
    }
}

/**
 * Converts ObjectStatusEnum Array to its string Array representation
*/
func ObjectStatusEnumArrayToValue(objectStatusEnum []ObjectStatusEnum) []string {
    convArray := make([]string,len( objectStatusEnum))
    for i:=0; i<len(objectStatusEnum);i++ {
        convArray[i] = ObjectStatusEnumToValue(objectStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectStatusEnumFromValue(value string) ObjectStatusEnum {
    switch value {
        case "kFilesCloned":
            return ObjectStatus_KFILESCLONED
        case "kFetchedEntityInfo":
            return ObjectStatus_KFETCHEDENTITYINFO
        case "kVMCreated":
            return ObjectStatus_KVMCREATED
        case "kRelocationStarted":
            return ObjectStatus_KRELOCATIONSTARTED
        case "kFinished":
            return ObjectStatus_KFINISHED
        case "kAborted":
            return ObjectStatus_KABORTED
        case "kDataCopyStarted":
            return ObjectStatus_KDATACOPYSTARTED
        case "kInProgress":
            return ObjectStatus_KINPROGRESS
        default:
            return ObjectStatus_KFILESCLONED
    }
}
