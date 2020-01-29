// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertCategoryEnum enum
 */
type AlertCategoryEnum int

/**
 * Value collection for AlertCategoryEnum enum
 */
const (
    AlertCategory_KDISK AlertCategoryEnum = 1 + iota
    AlertCategory_KNODE
    AlertCategory_KCLUSTER
    AlertCategory_KNODEHEALTH
    AlertCategory_KCLUSTERHEALTH
    AlertCategory_KBACKUPRESTORE
    AlertCategory_KENCRYPTION
    AlertCategory_KARCHIVALRESTORE
    AlertCategory_KREMOTEREPLICATION
    AlertCategory_KQUOTA
    AlertCategory_KLICENSE
    AlertCategory_KHELIOSPROACTIVEWELLNESS
    AlertCategory_KHELIOSANALYTICSJOBS
    AlertCategory_KHELIOSSIGNATUREJOBS
    AlertCategory_KSECURITY
)

func (r AlertCategoryEnum) MarshalJSON() ([]byte, error) {
    s := AlertCategoryEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertCategoryEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertCategoryEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertCategoryEnum to its string representation
 */
func AlertCategoryEnumToValue(alertCategoryEnum AlertCategoryEnum) string {
    switch alertCategoryEnum {
        case AlertCategory_KDISK:
    		return "kDisk"
        case AlertCategory_KNODE:
    		return "kNode"
        case AlertCategory_KCLUSTER:
    		return "kCluster"
        case AlertCategory_KNODEHEALTH:
    		return "kNodeHealth"
        case AlertCategory_KCLUSTERHEALTH:
    		return "kClusterHealth"
        case AlertCategory_KBACKUPRESTORE:
    		return "kBackupRestore"
        case AlertCategory_KENCRYPTION:
    		return "kEncryption"
        case AlertCategory_KARCHIVALRESTORE:
    		return "kArchivalRestore"
        case AlertCategory_KREMOTEREPLICATION:
    		return "kRemoteReplication"
        case AlertCategory_KQUOTA:
    		return "kQuota"
        case AlertCategory_KLICENSE:
    		return "kLicense"
        case AlertCategory_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"
        case AlertCategory_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"
        case AlertCategory_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"
        case AlertCategory_KSECURITY:
    		return "kSecurity"
        default:
        	return "kDisk"
    }
}

/**
 * Converts AlertCategoryEnum Array to its string Array representation
*/
func AlertCategoryEnumArrayToValue(alertCategoryEnum []AlertCategoryEnum) []string {
    convArray := make([]string,len( alertCategoryEnum))
    for i:=0; i<len(alertCategoryEnum);i++ {
        convArray[i] = AlertCategoryEnumToValue(alertCategoryEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertCategoryEnumFromValue(value string) AlertCategoryEnum {
    switch value {
        case "kDisk":
            return AlertCategory_KDISK
        case "kNode":
            return AlertCategory_KNODE
        case "kCluster":
            return AlertCategory_KCLUSTER
        case "kNodeHealth":
            return AlertCategory_KNODEHEALTH
        case "kClusterHealth":
            return AlertCategory_KCLUSTERHEALTH
        case "kBackupRestore":
            return AlertCategory_KBACKUPRESTORE
        case "kEncryption":
            return AlertCategory_KENCRYPTION
        case "kArchivalRestore":
            return AlertCategory_KARCHIVALRESTORE
        case "kRemoteReplication":
            return AlertCategory_KREMOTEREPLICATION
        case "kQuota":
            return AlertCategory_KQUOTA
        case "kLicense":
            return AlertCategory_KLICENSE
        case "kHeliosProActiveWellness":
            return AlertCategory_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return AlertCategory_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return AlertCategory_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return AlertCategory_KSECURITY
        default:
            return AlertCategory_KDISK
    }
}
