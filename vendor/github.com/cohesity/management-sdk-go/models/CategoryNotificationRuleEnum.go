// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CategoryNotificationRuleEnum enum
 */
type CategoryNotificationRuleEnum int

/**
 * Value collection for CategoryNotificationRuleEnum enum
 */
const (
    CategoryNotificationRule_KDISK CategoryNotificationRuleEnum = 1 + iota
    CategoryNotificationRule_KNODE
    CategoryNotificationRule_KCLUSTER
    CategoryNotificationRule_KNODEHEALTH
    CategoryNotificationRule_KCLUSTERHEALTH
    CategoryNotificationRule_KBACKUPRESTORE
    CategoryNotificationRule_KENCRYPTION
    CategoryNotificationRule_KARCHIVALRESTORE
    CategoryNotificationRule_KREMOTEREPLICATION
    CategoryNotificationRule_KQUOTA
    CategoryNotificationRule_KLICENSE
    CategoryNotificationRule_KHELIOSPROACTIVEWELLNESS
    CategoryNotificationRule_KHELIOSANALYTICSJOBS
    CategoryNotificationRule_KHELIOSSIGNATUREJOBS
    CategoryNotificationRule_KSECURITY
)

func (r CategoryNotificationRuleEnum) MarshalJSON() ([]byte, error) {
    s := CategoryNotificationRuleEnumToValue(r)
    return json.Marshal(s)
}

func (r *CategoryNotificationRuleEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CategoryNotificationRuleEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CategoryNotificationRuleEnum to its string representation
 */
func CategoryNotificationRuleEnumToValue(categoryNotificationRuleEnum CategoryNotificationRuleEnum) string {
    switch categoryNotificationRuleEnum {
        case CategoryNotificationRule_KDISK:
    		return "kDisk"
        case CategoryNotificationRule_KNODE:
    		return "kNode"
        case CategoryNotificationRule_KCLUSTER:
    		return "kCluster"
        case CategoryNotificationRule_KNODEHEALTH:
    		return "kNodeHealth"
        case CategoryNotificationRule_KCLUSTERHEALTH:
    		return "kClusterHealth"
        case CategoryNotificationRule_KBACKUPRESTORE:
    		return "kBackupRestore"
        case CategoryNotificationRule_KENCRYPTION:
    		return "kEncryption"
        case CategoryNotificationRule_KARCHIVALRESTORE:
    		return "kArchivalRestore"
        case CategoryNotificationRule_KREMOTEREPLICATION:
    		return "kRemoteReplication"
        case CategoryNotificationRule_KQUOTA:
    		return "kQuota"
        case CategoryNotificationRule_KLICENSE:
    		return "kLicense"
        case CategoryNotificationRule_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"
        case CategoryNotificationRule_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"
        case CategoryNotificationRule_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"
        case CategoryNotificationRule_KSECURITY:
    		return "kSecurity"
        default:
        	return "kDisk"
    }
}

/**
 * Converts CategoryNotificationRuleEnum Array to its string Array representation
*/
func CategoryNotificationRuleEnumArrayToValue(categoryNotificationRuleEnum []CategoryNotificationRuleEnum) []string {
    convArray := make([]string,len( categoryNotificationRuleEnum))
    for i:=0; i<len(categoryNotificationRuleEnum);i++ {
        convArray[i] = CategoryNotificationRuleEnumToValue(categoryNotificationRuleEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CategoryNotificationRuleEnumFromValue(value string) CategoryNotificationRuleEnum {
    switch value {
        case "kDisk":
            return CategoryNotificationRule_KDISK
        case "kNode":
            return CategoryNotificationRule_KNODE
        case "kCluster":
            return CategoryNotificationRule_KCLUSTER
        case "kNodeHealth":
            return CategoryNotificationRule_KNODEHEALTH
        case "kClusterHealth":
            return CategoryNotificationRule_KCLUSTERHEALTH
        case "kBackupRestore":
            return CategoryNotificationRule_KBACKUPRESTORE
        case "kEncryption":
            return CategoryNotificationRule_KENCRYPTION
        case "kArchivalRestore":
            return CategoryNotificationRule_KARCHIVALRESTORE
        case "kRemoteReplication":
            return CategoryNotificationRule_KREMOTEREPLICATION
        case "kQuota":
            return CategoryNotificationRule_KQUOTA
        case "kLicense":
            return CategoryNotificationRule_KLICENSE
        case "kHeliosProActiveWellness":
            return CategoryNotificationRule_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return CategoryNotificationRule_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return CategoryNotificationRule_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return CategoryNotificationRule_KSECURITY
        default:
            return CategoryNotificationRule_KDISK
    }
}
