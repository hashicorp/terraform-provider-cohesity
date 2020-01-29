package models

import(
    "encoding/json"
)

/**
 * Type definition for Type31Enum enum
 */
type Type31Enum int

/**
 * Value collection for Type31Enum enum
 */
const (
    Type31_KRECOVERVMS Type31Enum = 1 + iota
    Type31_KCLONEVMS
    Type31_KCLONEVIEW
    Type31_KMOUNTVOLUMES
    Type31_KRESTOREFILES
    Type31_KRECOVERAPP
    Type31_KCLONEAPP
    Type31_KRECOVERSANVOLUME
    Type31_KCONVERTANDDEPLOYVMS
    Type31_KMOUNTFILEVOLUME
    Type31_KSYSTEM
    Type31_KRECOVERVOLUMES
    Type31_KDEPLOYVMS
    Type31_KDOWNLOADFILES
    Type31_KRECOVEREMAILS
    Type31_KRECOVERDISKS
)

func (r Type31Enum) MarshalJSON() ([]byte, error) { 
    s := Type31EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type31Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type31EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type31Enum to its string representation
 */
func Type31EnumToValue(type31Enum Type31Enum) string {
    switch type31Enum {
        case Type31_KRECOVERVMS:
    		return "kRecoverVMs"		
        case Type31_KCLONEVMS:
    		return "kCloneVMs"		
        case Type31_KCLONEVIEW:
    		return "kCloneView"		
        case Type31_KMOUNTVOLUMES:
    		return "kMountVolumes"		
        case Type31_KRESTOREFILES:
    		return "kRestoreFiles"		
        case Type31_KRECOVERAPP:
    		return "kRecoverApp"		
        case Type31_KCLONEAPP:
    		return "kCloneApp"		
        case Type31_KRECOVERSANVOLUME:
    		return "kRecoverSanVolume"		
        case Type31_KCONVERTANDDEPLOYVMS:
    		return "kConvertAndDeployVMs"		
        case Type31_KMOUNTFILEVOLUME:
    		return "kMountFileVolume"		
        case Type31_KSYSTEM:
    		return "kSystem"		
        case Type31_KRECOVERVOLUMES:
    		return "kRecoverVolumes"		
        case Type31_KDEPLOYVMS:
    		return "kDeployVMs"		
        case Type31_KDOWNLOADFILES:
    		return "kDownloadFiles"		
        case Type31_KRECOVEREMAILS:
    		return "kRecoverEmails"		
        case Type31_KRECOVERDISKS:
    		return "kRecoverDisks"		
        default:
        	return "kRecoverVMs"
    }
}

/**
 * Converts Type31Enum Array to its string Array representation
*/
func Type31EnumArrayToValue(type31Enum []Type31Enum) []string {
    convArray := make([]string,len( type31Enum))
    for i:=0; i<len(type31Enum);i++ {
        convArray[i] = Type31EnumToValue(type31Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type31EnumFromValue(value string) Type31Enum {
    switch value {
        case "kRecoverVMs":
            return Type31_KRECOVERVMS
        case "kCloneVMs":
            return Type31_KCLONEVMS
        case "kCloneView":
            return Type31_KCLONEVIEW
        case "kMountVolumes":
            return Type31_KMOUNTVOLUMES
        case "kRestoreFiles":
            return Type31_KRESTOREFILES
        case "kRecoverApp":
            return Type31_KRECOVERAPP
        case "kCloneApp":
            return Type31_KCLONEAPP
        case "kRecoverSanVolume":
            return Type31_KRECOVERSANVOLUME
        case "kConvertAndDeployVMs":
            return Type31_KCONVERTANDDEPLOYVMS
        case "kMountFileVolume":
            return Type31_KMOUNTFILEVOLUME
        case "kSystem":
            return Type31_KSYSTEM
        case "kRecoverVolumes":
            return Type31_KRECOVERVOLUMES
        case "kDeployVMs":
            return Type31_KDEPLOYVMS
        case "kDownloadFiles":
            return Type31_KDOWNLOADFILES
        case "kRecoverEmails":
            return Type31_KRECOVEREMAILS
        case "kRecoverDisks":
            return Type31_KRECOVERDISKS
        default:
            return Type31_KRECOVERVMS
    }
}
