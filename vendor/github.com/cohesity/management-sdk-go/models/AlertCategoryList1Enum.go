package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertCategoryList1Enum enum
 */
type AlertCategoryList1Enum int

/**
 * Value collection for AlertCategoryList1Enum enum
 */
const (
    AlertCategoryList1_KDISK AlertCategoryList1Enum = 1 + iota
    AlertCategoryList1_KNODE
    AlertCategoryList1_KCLUSTER
    AlertCategoryList1_KNODEHEALTH
    AlertCategoryList1_KCLUSTERHEALTH
    AlertCategoryList1_KBACKUPRESTORE
    AlertCategoryList1_KENCRYPTION
    AlertCategoryList1_KARCHIVALRESTORE
    AlertCategoryList1_KREMOTEREPLICATION
    AlertCategoryList1_KQUOTA
)

func (r AlertCategoryList1Enum) MarshalJSON() ([]byte, error) { 
    s := AlertCategoryList1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *AlertCategoryList1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  AlertCategoryList1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts AlertCategoryList1Enum to its string representation
 */
func AlertCategoryList1EnumToValue(alertCategoryList1Enum AlertCategoryList1Enum) string {
    switch alertCategoryList1Enum {
        case AlertCategoryList1_KDISK:
    		return "kDisk"		
        case AlertCategoryList1_KNODE:
    		return "kNode"		
        case AlertCategoryList1_KCLUSTER:
    		return "kCluster"		
        case AlertCategoryList1_KNODEHEALTH:
    		return "kNodeHealth"		
        case AlertCategoryList1_KCLUSTERHEALTH:
    		return "kClusterHealth"		
        case AlertCategoryList1_KBACKUPRESTORE:
    		return "kBackupRestore"		
        case AlertCategoryList1_KENCRYPTION:
    		return "kEncryption"		
        case AlertCategoryList1_KARCHIVALRESTORE:
    		return "kArchivalRestore"		
        case AlertCategoryList1_KREMOTEREPLICATION:
    		return "kRemoteReplication"		
        case AlertCategoryList1_KQUOTA:
    		return "kQuota"		
        default:
        	return "kDisk"
    }
}

/**
 * Converts AlertCategoryList1Enum Array to its string Array representation
*/
func AlertCategoryList1EnumArrayToValue(alertCategoryList1Enum []AlertCategoryList1Enum) []string {
    convArray := make([]string,len( alertCategoryList1Enum))
    for i:=0; i<len(alertCategoryList1Enum);i++ {
        convArray[i] = AlertCategoryList1EnumToValue(alertCategoryList1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertCategoryList1EnumFromValue(value string) AlertCategoryList1Enum {
    switch value {
        case "kDisk":
            return AlertCategoryList1_KDISK
        case "kNode":
            return AlertCategoryList1_KNODE
        case "kCluster":
            return AlertCategoryList1_KCLUSTER
        case "kNodeHealth":
            return AlertCategoryList1_KNODEHEALTH
        case "kClusterHealth":
            return AlertCategoryList1_KCLUSTERHEALTH
        case "kBackupRestore":
            return AlertCategoryList1_KBACKUPRESTORE
        case "kEncryption":
            return AlertCategoryList1_KENCRYPTION
        case "kArchivalRestore":
            return AlertCategoryList1_KARCHIVALRESTORE
        case "kRemoteReplication":
            return AlertCategoryList1_KREMOTEREPLICATION
        case "kQuota":
            return AlertCategoryList1_KQUOTA
        default:
            return AlertCategoryList1_KDISK
    }
}
