// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeRestoreTaskEnum enum
 */
type TypeRestoreTaskEnum int

/**
 * Value collection for TypeRestoreTaskEnum enum
 */
const (
    TypeRestoreTask_KRECOVERVMS TypeRestoreTaskEnum = 1 + iota
    TypeRestoreTask_KCLONEVMS
    TypeRestoreTask_KCLONEVIEW
    TypeRestoreTask_KMOUNTVOLUMES
    TypeRestoreTask_KRESTOREFILES
    TypeRestoreTask_KRECOVERAPP
    TypeRestoreTask_KCLONEAPP
    TypeRestoreTask_KRECOVERSANVOLUME
    TypeRestoreTask_KCONVERTANDDEPLOYVMS
    TypeRestoreTask_KMOUNTFILEVOLUME
    TypeRestoreTask_KSYSTEM
    TypeRestoreTask_KRECOVERVOLUMES
    TypeRestoreTask_KDEPLOYVMS
    TypeRestoreTask_KDOWNLOADFILES
    TypeRestoreTask_KRECOVEREMAILS
    TypeRestoreTask_KRECOVERDISKS
)

func (r TypeRestoreTaskEnum) MarshalJSON() ([]byte, error) {
    s := TypeRestoreTaskEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeRestoreTaskEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeRestoreTaskEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeRestoreTaskEnum to its string representation
 */
func TypeRestoreTaskEnumToValue(typeRestoreTaskEnum TypeRestoreTaskEnum) string {
    switch typeRestoreTaskEnum {
        case TypeRestoreTask_KRECOVERVMS:
    		return "kRecoverVMs"
        case TypeRestoreTask_KCLONEVMS:
    		return "kCloneVMs"
        case TypeRestoreTask_KCLONEVIEW:
    		return "kCloneView"
        case TypeRestoreTask_KMOUNTVOLUMES:
    		return "kMountVolumes"
        case TypeRestoreTask_KRESTOREFILES:
    		return "kRestoreFiles"
        case TypeRestoreTask_KRECOVERAPP:
    		return "kRecoverApp"
        case TypeRestoreTask_KCLONEAPP:
    		return "kCloneApp"
        case TypeRestoreTask_KRECOVERSANVOLUME:
    		return "kRecoverSanVolume"
        case TypeRestoreTask_KCONVERTANDDEPLOYVMS:
    		return "kConvertAndDeployVMs"
        case TypeRestoreTask_KMOUNTFILEVOLUME:
    		return "kMountFileVolume"
        case TypeRestoreTask_KSYSTEM:
    		return "kSystem"
        case TypeRestoreTask_KRECOVERVOLUMES:
    		return "kRecoverVolumes"
        case TypeRestoreTask_KDEPLOYVMS:
    		return "kDeployVMs"
        case TypeRestoreTask_KDOWNLOADFILES:
    		return "kDownloadFiles"
        case TypeRestoreTask_KRECOVEREMAILS:
    		return "kRecoverEmails"
        case TypeRestoreTask_KRECOVERDISKS:
    		return "kRecoverDisks"
        default:
        	return "kRecoverVMs"
    }
}

/**
 * Converts TypeRestoreTaskEnum Array to its string Array representation
*/
func TypeRestoreTaskEnumArrayToValue(typeRestoreTaskEnum []TypeRestoreTaskEnum) []string {
    convArray := make([]string,len( typeRestoreTaskEnum))
    for i:=0; i<len(typeRestoreTaskEnum);i++ {
        convArray[i] = TypeRestoreTaskEnumToValue(typeRestoreTaskEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeRestoreTaskEnumFromValue(value string) TypeRestoreTaskEnum {
    switch value {
        case "kRecoverVMs":
            return TypeRestoreTask_KRECOVERVMS
        case "kCloneVMs":
            return TypeRestoreTask_KCLONEVMS
        case "kCloneView":
            return TypeRestoreTask_KCLONEVIEW
        case "kMountVolumes":
            return TypeRestoreTask_KMOUNTVOLUMES
        case "kRestoreFiles":
            return TypeRestoreTask_KRESTOREFILES
        case "kRecoverApp":
            return TypeRestoreTask_KRECOVERAPP
        case "kCloneApp":
            return TypeRestoreTask_KCLONEAPP
        case "kRecoverSanVolume":
            return TypeRestoreTask_KRECOVERSANVOLUME
        case "kConvertAndDeployVMs":
            return TypeRestoreTask_KCONVERTANDDEPLOYVMS
        case "kMountFileVolume":
            return TypeRestoreTask_KMOUNTFILEVOLUME
        case "kSystem":
            return TypeRestoreTask_KSYSTEM
        case "kRecoverVolumes":
            return TypeRestoreTask_KRECOVERVOLUMES
        case "kDeployVMs":
            return TypeRestoreTask_KDEPLOYVMS
        case "kDownloadFiles":
            return TypeRestoreTask_KDOWNLOADFILES
        case "kRecoverEmails":
            return TypeRestoreTask_KRECOVEREMAILS
        case "kRecoverDisks":
            return TypeRestoreTask_KRECOVERDISKS
        default:
            return TypeRestoreTask_KRECOVERVMS
    }
}
