// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertCategoryListGetAlertsEnum enum
 */
type AlertCategoryListGetAlertsEnum int

/**
 * Value collection for AlertCategoryListGetAlertsEnum enum
 */
const (
    AlertCategoryListGetAlerts_KDISK AlertCategoryListGetAlertsEnum = 1 + iota
    AlertCategoryListGetAlerts_KNODE
    AlertCategoryListGetAlerts_KCLUSTER
    AlertCategoryListGetAlerts_KNODEHEALTH
    AlertCategoryListGetAlerts_KCLUSTERHEALTH
    AlertCategoryListGetAlerts_KBACKUPRESTORE
    AlertCategoryListGetAlerts_KENCRYPTION
    AlertCategoryListGetAlerts_KARCHIVALRESTORE
    AlertCategoryListGetAlerts_KREMOTEREPLICATION
    AlertCategoryListGetAlerts_KQUOTA
)

func (r AlertCategoryListGetAlertsEnum) MarshalJSON() ([]byte, error) {
    s := AlertCategoryListGetAlertsEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertCategoryListGetAlertsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertCategoryListGetAlertsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertCategoryListGetAlertsEnum to its string representation
 */
func AlertCategoryListGetAlertsEnumToValue(alertCategoryListGetAlertsEnum AlertCategoryListGetAlertsEnum) string {
    switch alertCategoryListGetAlertsEnum {
        case AlertCategoryListGetAlerts_KDISK:
    		return "kDisk"
        case AlertCategoryListGetAlerts_KNODE:
    		return "kNode"
        case AlertCategoryListGetAlerts_KCLUSTER:
    		return "kCluster"
        case AlertCategoryListGetAlerts_KNODEHEALTH:
    		return "kNodeHealth"
        case AlertCategoryListGetAlerts_KCLUSTERHEALTH:
    		return "kClusterHealth"
        case AlertCategoryListGetAlerts_KBACKUPRESTORE:
    		return "kBackupRestore"
        case AlertCategoryListGetAlerts_KENCRYPTION:
    		return "kEncryption"
        case AlertCategoryListGetAlerts_KARCHIVALRESTORE:
    		return "kArchivalRestore"
        case AlertCategoryListGetAlerts_KREMOTEREPLICATION:
    		return "kRemoteReplication"
        case AlertCategoryListGetAlerts_KQUOTA:
    		return "kQuota"
        default:
        	return "kDisk"
    }
}

/**
 * Converts AlertCategoryListGetAlertsEnum Array to its string Array representation
*/
func AlertCategoryListGetAlertsEnumArrayToValue(alertCategoryListGetAlertsEnum []AlertCategoryListGetAlertsEnum) []string {
    convArray := make([]string,len( alertCategoryListGetAlertsEnum))
    for i:=0; i<len(alertCategoryListGetAlertsEnum);i++ {
        convArray[i] = AlertCategoryListGetAlertsEnumToValue(alertCategoryListGetAlertsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertCategoryListGetAlertsEnumFromValue(value string) AlertCategoryListGetAlertsEnum {
    switch value {
        case "kDisk":
            return AlertCategoryListGetAlerts_KDISK
        case "kNode":
            return AlertCategoryListGetAlerts_KNODE
        case "kCluster":
            return AlertCategoryListGetAlerts_KCLUSTER
        case "kNodeHealth":
            return AlertCategoryListGetAlerts_KNODEHEALTH
        case "kClusterHealth":
            return AlertCategoryListGetAlerts_KCLUSTERHEALTH
        case "kBackupRestore":
            return AlertCategoryListGetAlerts_KBACKUPRESTORE
        case "kEncryption":
            return AlertCategoryListGetAlerts_KENCRYPTION
        case "kArchivalRestore":
            return AlertCategoryListGetAlerts_KARCHIVALRESTORE
        case "kRemoteReplication":
            return AlertCategoryListGetAlerts_KREMOTEREPLICATION
        case "kQuota":
            return AlertCategoryListGetAlerts_KQUOTA
        default:
            return AlertCategoryListGetAlerts_KDISK
    }
}
