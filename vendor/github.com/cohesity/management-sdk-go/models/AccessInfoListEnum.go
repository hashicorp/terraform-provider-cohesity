// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AccessInfoListEnum enum
 */
type AccessInfoListEnum int

/**
 * Value collection for AccessInfoListEnum enum
 */
const (
    AccessInfoList_KFILEREADDATA AccessInfoListEnum = 1 + iota
    AccessInfoList_KFILEWRITEDATA
    AccessInfoList_KFILEAPPENDDATA
    AccessInfoList_KFILEREADEA
    AccessInfoList_KFILEWRITEEA
    AccessInfoList_KFILEEXECUTE
    AccessInfoList_KFILEDELETECHILD
    AccessInfoList_KFILEREADATTRIBUTES
    AccessInfoList_KFILEWRITEATTRIBUTES
    AccessInfoList_KDELETE
    AccessInfoList_KREADCONTROL
    AccessInfoList_KWRITEDAC
    AccessInfoList_KWRITEOWNER
    AccessInfoList_KSYNCHRONIZE
    AccessInfoList_KACCESSSYSTEMSECURITY
    AccessInfoList_KMAXIMUMALLOWED
    AccessInfoList_KGENERICALL
    AccessInfoList_KGENERICEXECUTE
    AccessInfoList_KGENERICWRITE
    AccessInfoList_KGENERICREAD
)

func (r AccessInfoListEnum) MarshalJSON() ([]byte, error) {
    s := AccessInfoListEnumToValue(r)
    return json.Marshal(s)
}

func (r *AccessInfoListEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AccessInfoListEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AccessInfoListEnum to its string representation
 */
func AccessInfoListEnumToValue(accessInfoListEnum AccessInfoListEnum) string {
    switch accessInfoListEnum {
        case AccessInfoList_KFILEREADDATA:
    		return "kFileReadData"
        case AccessInfoList_KFILEWRITEDATA:
    		return "kFileWriteData"
        case AccessInfoList_KFILEAPPENDDATA:
    		return "kFileAppendData"
        case AccessInfoList_KFILEREADEA:
    		return "kFileReadEa"
        case AccessInfoList_KFILEWRITEEA:
    		return "kFileWriteEa"
        case AccessInfoList_KFILEEXECUTE:
    		return "kFileExecute"
        case AccessInfoList_KFILEDELETECHILD:
    		return "kFileDeleteChild"
        case AccessInfoList_KFILEREADATTRIBUTES:
    		return "kFileReadAttributes"
        case AccessInfoList_KFILEWRITEATTRIBUTES:
    		return "kFileWriteAttributes"
        case AccessInfoList_KDELETE:
    		return "kDelete"
        case AccessInfoList_KREADCONTROL:
    		return "kReadControl"
        case AccessInfoList_KWRITEDAC:
    		return "kWriteDac"
        case AccessInfoList_KWRITEOWNER:
    		return "kWriteOwner"
        case AccessInfoList_KSYNCHRONIZE:
    		return "kSynchronize"
        case AccessInfoList_KACCESSSYSTEMSECURITY:
    		return "kAccessSystemSecurity"
        case AccessInfoList_KMAXIMUMALLOWED:
    		return "kMaximumAllowed"
        case AccessInfoList_KGENERICALL:
    		return "kGenericAll"
        case AccessInfoList_KGENERICEXECUTE:
    		return "kGenericExecute"
        case AccessInfoList_KGENERICWRITE:
    		return "kGenericWrite"
        case AccessInfoList_KGENERICREAD:
    		return "kGenericRead"
        default:
        	return "kFileReadData"
    }
}

/**
 * Converts AccessInfoListEnum Array to its string Array representation
*/
func AccessInfoListEnumArrayToValue(accessInfoListEnum []AccessInfoListEnum) []string {
    convArray := make([]string,len( accessInfoListEnum))
    for i:=0; i<len(accessInfoListEnum);i++ {
        convArray[i] = AccessInfoListEnumToValue(accessInfoListEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AccessInfoListEnumFromValue(value string) AccessInfoListEnum {
    switch value {
        case "kFileReadData":
            return AccessInfoList_KFILEREADDATA
        case "kFileWriteData":
            return AccessInfoList_KFILEWRITEDATA
        case "kFileAppendData":
            return AccessInfoList_KFILEAPPENDDATA
        case "kFileReadEa":
            return AccessInfoList_KFILEREADEA
        case "kFileWriteEa":
            return AccessInfoList_KFILEWRITEEA
        case "kFileExecute":
            return AccessInfoList_KFILEEXECUTE
        case "kFileDeleteChild":
            return AccessInfoList_KFILEDELETECHILD
        case "kFileReadAttributes":
            return AccessInfoList_KFILEREADATTRIBUTES
        case "kFileWriteAttributes":
            return AccessInfoList_KFILEWRITEATTRIBUTES
        case "kDelete":
            return AccessInfoList_KDELETE
        case "kReadControl":
            return AccessInfoList_KREADCONTROL
        case "kWriteDac":
            return AccessInfoList_KWRITEDAC
        case "kWriteOwner":
            return AccessInfoList_KWRITEOWNER
        case "kSynchronize":
            return AccessInfoList_KSYNCHRONIZE
        case "kAccessSystemSecurity":
            return AccessInfoList_KACCESSSYSTEMSECURITY
        case "kMaximumAllowed":
            return AccessInfoList_KMAXIMUMALLOWED
        case "kGenericAll":
            return AccessInfoList_KGENERICALL
        case "kGenericExecute":
            return AccessInfoList_KGENERICEXECUTE
        case "kGenericWrite":
            return AccessInfoList_KGENERICWRITE
        case "kGenericRead":
            return AccessInfoList_KGENERICREAD
        default:
            return AccessInfoList_KFILEREADDATA
    }
}
