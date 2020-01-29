package models

import(
    "encoding/json"
)

/**
 * Type definition for Type32Enum enum
 */
type Type32Enum int

/**
 * Value collection for Type32Enum enum
 */
const (
    Type32_KRECOVERVMS Type32Enum = 1 + iota
    Type32_KCLONEVMS
    Type32_KCLONEVIEW
    Type32_KMOUNTVOLUMES
    Type32_KRESTOREFILES
    Type32_KRECOVERAPP
    Type32_KCLONEAPP
    Type32_KRECOVERSANVOLUME
    Type32_KCONVERTANDDEPLOYVMS
    Type32_KMOUNTFILEVOLUME
    Type32_KSYSTEM
    Type32_KRECOVERVOLUMES
    Type32_KDEPLOYVMS
    Type32_KDOWNLOADFILES
    Type32_KRECOVEREMAILS
    Type32_KRECOVERDISKS
)

func (r Type32Enum) MarshalJSON() ([]byte, error) { 
    s := Type32EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type32Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type32EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type32Enum to its string representation
 */
func Type32EnumToValue(type32Enum Type32Enum) string {
    switch type32Enum {
        case Type32_KRECOVERVMS:
    		return "kRecoverVMs"		
        case Type32_KCLONEVMS:
    		return "kCloneVMs"		
        case Type32_KCLONEVIEW:
    		return "kCloneView"		
        case Type32_KMOUNTVOLUMES:
    		return "kMountVolumes"		
        case Type32_KRESTOREFILES:
    		return "kRestoreFiles"		
        case Type32_KRECOVERAPP:
    		return "kRecoverApp"		
        case Type32_KCLONEAPP:
    		return "kCloneApp"		
        case Type32_KRECOVERSANVOLUME:
    		return "kRecoverSanVolume"		
        case Type32_KCONVERTANDDEPLOYVMS:
    		return "kConvertAndDeployVMs"		
        case Type32_KMOUNTFILEVOLUME:
    		return "kMountFileVolume"		
        case Type32_KSYSTEM:
    		return "kSystem"		
        case Type32_KRECOVERVOLUMES:
    		return "kRecoverVolumes"		
        case Type32_KDEPLOYVMS:
    		return "kDeployVMs"		
        case Type32_KDOWNLOADFILES:
    		return "kDownloadFiles"		
        case Type32_KRECOVEREMAILS:
    		return "kRecoverEmails"		
        case Type32_KRECOVERDISKS:
    		return "kRecoverDisks"		
        default:
        	return "kRecoverVMs"
    }
}

/**
 * Converts Type32Enum Array to its string Array representation
*/
func Type32EnumArrayToValue(type32Enum []Type32Enum) []string {
    convArray := make([]string,len( type32Enum))
    for i:=0; i<len(type32Enum);i++ {
        convArray[i] = Type32EnumToValue(type32Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type32EnumFromValue(value string) Type32Enum {
    switch value {
        case "kRecoverVMs":
            return Type32_KRECOVERVMS
        case "kCloneVMs":
            return Type32_KCLONEVMS
        case "kCloneView":
            return Type32_KCLONEVIEW
        case "kMountVolumes":
            return Type32_KMOUNTVOLUMES
        case "kRestoreFiles":
            return Type32_KRESTOREFILES
        case "kRecoverApp":
            return Type32_KRECOVERAPP
        case "kCloneApp":
            return Type32_KCLONEAPP
        case "kRecoverSanVolume":
            return Type32_KRECOVERSANVOLUME
        case "kConvertAndDeployVMs":
            return Type32_KCONVERTANDDEPLOYVMS
        case "kMountFileVolume":
            return Type32_KMOUNTFILEVOLUME
        case "kSystem":
            return Type32_KSYSTEM
        case "kRecoverVolumes":
            return Type32_KRECOVERVOLUMES
        case "kDeployVMs":
            return Type32_KDEPLOYVMS
        case "kDownloadFiles":
            return Type32_KDOWNLOADFILES
        case "kRecoverEmails":
            return Type32_KRECOVEREMAILS
        case "kRecoverDisks":
            return Type32_KRECOVERDISKS
        default:
            return Type32_KRECOVERVMS
    }
}
