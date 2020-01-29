// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CategoryEnum enum
 */
type CategoryEnum int

/**
 * Value collection for CategoryEnum enum
 */
const (
    Category_KDISK CategoryEnum = 1 + iota
    Category_KNODE
    Category_KCLUSTER
    Category_KNODEHEALTH
    Category_KCLUSTERHEALTH
    Category_KBACKUPRESTORE
    Category_KENCRYPTION
    Category_KARCHIVALRESTORE
    Category_KREMOTEREPLICATION
    Category_KQUOTA
    Category_KLICENSE
    Category_KHELIOSPROACTIVEWELLNESS
    Category_KHELIOSANALYTICSJOBS
    Category_KHELIOSSIGNATUREJOBS
    Category_KSECURITY
)

func (r CategoryEnum) MarshalJSON() ([]byte, error) {
    s := CategoryEnumToValue(r)
    return json.Marshal(s)
}

func (r *CategoryEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CategoryEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CategoryEnum to its string representation
 */
func CategoryEnumToValue(categoryEnum CategoryEnum) string {
    switch categoryEnum {
        case Category_KDISK:
    		return "kDisk"
        case Category_KNODE:
    		return "kNode"
        case Category_KCLUSTER:
    		return "kCluster"
        case Category_KNODEHEALTH:
    		return "kNodeHealth"
        case Category_KCLUSTERHEALTH:
    		return "kClusterHealth"
        case Category_KBACKUPRESTORE:
    		return "kBackupRestore"
        case Category_KENCRYPTION:
    		return "kEncryption"
        case Category_KARCHIVALRESTORE:
    		return "kArchivalRestore"
        case Category_KREMOTEREPLICATION:
    		return "kRemoteReplication"
        case Category_KQUOTA:
    		return "kQuota"
        case Category_KLICENSE:
    		return "kLicense"
        case Category_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"
        case Category_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"
        case Category_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"
        case Category_KSECURITY:
    		return "kSecurity"
        default:
        	return "kDisk"
    }
}

/**
 * Converts CategoryEnum Array to its string Array representation
*/
func CategoryEnumArrayToValue(categoryEnum []CategoryEnum) []string {
    convArray := make([]string,len( categoryEnum))
    for i:=0; i<len(categoryEnum);i++ {
        convArray[i] = CategoryEnumToValue(categoryEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CategoryEnumFromValue(value string) CategoryEnum {
    switch value {
        case "kDisk":
            return Category_KDISK
        case "kNode":
            return Category_KNODE
        case "kCluster":
            return Category_KCLUSTER
        case "kNodeHealth":
            return Category_KNODEHEALTH
        case "kClusterHealth":
            return Category_KCLUSTERHEALTH
        case "kBackupRestore":
            return Category_KBACKUPRESTORE
        case "kEncryption":
            return Category_KENCRYPTION
        case "kArchivalRestore":
            return Category_KARCHIVALRESTORE
        case "kRemoteReplication":
            return Category_KREMOTEREPLICATION
        case "kQuota":
            return Category_KQUOTA
        case "kLicense":
            return Category_KLICENSE
        case "kHeliosProActiveWellness":
            return Category_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return Category_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return Category_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return Category_KSECURITY
        default:
            return Category_KDISK
    }
}
