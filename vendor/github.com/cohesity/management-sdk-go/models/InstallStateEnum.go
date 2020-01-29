package models

import(
    "encoding/json"
)

/**
 * Type definition for InstallStateEnum enum
 */
type InstallStateEnum int

/**
 * Value collection for InstallStateEnum enum
 */
const (
    InstallState_KNOTINSTALLED InstallStateEnum = 1 + iota
    InstallState_KINSTALLINPROGRESS
    InstallState_KINSTALLED
    InstallState_KINSTALLFAILED
    InstallState_KUNINSTALLINPROGRESS
    InstallState_KUNINSTALLFAILED
)

func (r InstallStateEnum) MarshalJSON() ([]byte, error) { 
    s := InstallStateEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *InstallStateEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  InstallStateEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts InstallStateEnum to its string representation
 */
func InstallStateEnumToValue(installStateEnum InstallStateEnum) string {
    switch installStateEnum {
        case InstallState_KNOTINSTALLED:
    		return "kNotInstalled"		
        case InstallState_KINSTALLINPROGRESS:
    		return "kInstallInProgress"		
        case InstallState_KINSTALLED:
    		return "kInstalled"		
        case InstallState_KINSTALLFAILED:
    		return "kInstallFailed"		
        case InstallState_KUNINSTALLINPROGRESS:
    		return "kUninstallInProgress"		
        case InstallState_KUNINSTALLFAILED:
    		return "kUninstallFailed"		
        default:
        	return "kNotInstalled"
    }
}

/**
 * Converts InstallStateEnum Array to its string Array representation
*/
func InstallStateEnumArrayToValue(installStateEnum []InstallStateEnum) []string {
    convArray := make([]string,len( installStateEnum))
    for i:=0; i<len(installStateEnum);i++ {
        convArray[i] = InstallStateEnumToValue(installStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func InstallStateEnumFromValue(value string) InstallStateEnum {
    switch value {
        case "kNotInstalled":
            return InstallState_KNOTINSTALLED
        case "kInstallInProgress":
            return InstallState_KINSTALLINPROGRESS
        case "kInstalled":
            return InstallState_KINSTALLED
        case "kInstallFailed":
            return InstallState_KINSTALLFAILED
        case "kUninstallInProgress":
            return InstallState_KUNINSTALLINPROGRESS
        case "kUninstallFailed":
            return InstallState_KUNINSTALLFAILED
        default:
            return InstallState_KNOTINSTALLED
    }
}
