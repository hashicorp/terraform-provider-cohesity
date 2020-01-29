// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertCategoryListEnum enum
 */
type AlertCategoryListEnum int

/**
 * Value collection for AlertCategoryListEnum enum
 */
const (
    AlertCategoryList_KDISK AlertCategoryListEnum = 1 + iota
    AlertCategoryList_KNODE
    AlertCategoryList_KCLUSTER
    AlertCategoryList_KNODEHEALTH
    AlertCategoryList_KCLUSTERHEALTH
    AlertCategoryList_KBACKUPRESTORE
    AlertCategoryList_KENCRYPTION
    AlertCategoryList_KARCHIVALRESTORE
    AlertCategoryList_KREMOTEREPLICATION
    AlertCategoryList_KQUOTA
    AlertCategoryList_KLICENSE
    AlertCategoryList_KHELIOSPROACTIVEWELLNESS
    AlertCategoryList_KHELIOSANALYTICSJOBS
    AlertCategoryList_KHELIOSSIGNATUREJOBS
    AlertCategoryList_KSECURITY
)

func (r AlertCategoryListEnum) MarshalJSON() ([]byte, error) {
    s := AlertCategoryListEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertCategoryListEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertCategoryListEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertCategoryListEnum to its string representation
 */
func AlertCategoryListEnumToValue(alertCategoryListEnum AlertCategoryListEnum) string {
    switch alertCategoryListEnum {
        case AlertCategoryList_KDISK:
    		return "kDisk"
        case AlertCategoryList_KNODE:
    		return "kNode"
        case AlertCategoryList_KCLUSTER:
    		return "kCluster"
        case AlertCategoryList_KNODEHEALTH:
    		return "kNodeHealth"
        case AlertCategoryList_KCLUSTERHEALTH:
    		return "kClusterHealth"
        case AlertCategoryList_KBACKUPRESTORE:
    		return "kBackupRestore"
        case AlertCategoryList_KENCRYPTION:
    		return "kEncryption"
        case AlertCategoryList_KARCHIVALRESTORE:
    		return "kArchivalRestore"
        case AlertCategoryList_KREMOTEREPLICATION:
    		return "kRemoteReplication"
        case AlertCategoryList_KQUOTA:
    		return "kQuota"
        case AlertCategoryList_KLICENSE:
    		return "kLicense"
        case AlertCategoryList_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"
        case AlertCategoryList_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"
        case AlertCategoryList_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"
        case AlertCategoryList_KSECURITY:
    		return "kSecurity"
        default:
        	return "kDisk"
    }
}

/**
 * Converts AlertCategoryListEnum Array to its string Array representation
*/
func AlertCategoryListEnumArrayToValue(alertCategoryListEnum []AlertCategoryListEnum) []string {
    convArray := make([]string,len( alertCategoryListEnum))
    for i:=0; i<len(alertCategoryListEnum);i++ {
        convArray[i] = AlertCategoryListEnumToValue(alertCategoryListEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertCategoryListEnumFromValue(value string) AlertCategoryListEnum {
    switch value {
        case "kDisk":
            return AlertCategoryList_KDISK
        case "kNode":
            return AlertCategoryList_KNODE
        case "kCluster":
            return AlertCategoryList_KCLUSTER
        case "kNodeHealth":
            return AlertCategoryList_KNODEHEALTH
        case "kClusterHealth":
            return AlertCategoryList_KCLUSTERHEALTH
        case "kBackupRestore":
            return AlertCategoryList_KBACKUPRESTORE
        case "kEncryption":
            return AlertCategoryList_KENCRYPTION
        case "kArchivalRestore":
            return AlertCategoryList_KARCHIVALRESTORE
        case "kRemoteReplication":
            return AlertCategoryList_KREMOTEREPLICATION
        case "kQuota":
            return AlertCategoryList_KQUOTA
        case "kLicense":
            return AlertCategoryList_KLICENSE
        case "kHeliosProActiveWellness":
            return AlertCategoryList_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return AlertCategoryList_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return AlertCategoryList_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return AlertCategoryList_KSECURITY
        default:
            return AlertCategoryList_KDISK
    }
}
