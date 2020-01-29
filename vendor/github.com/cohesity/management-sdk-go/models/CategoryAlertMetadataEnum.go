// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CategoryAlertMetadataEnum enum
 */
type CategoryAlertMetadataEnum int

/**
 * Value collection for CategoryAlertMetadataEnum enum
 */
const (
    CategoryAlertMetadata_KDISK CategoryAlertMetadataEnum = 1 + iota
    CategoryAlertMetadata_KNODE
    CategoryAlertMetadata_KCLUSTER
    CategoryAlertMetadata_KNODEHEALTH
    CategoryAlertMetadata_KCLUSTERHEALTH
    CategoryAlertMetadata_KBACKUPRESTORE
    CategoryAlertMetadata_KENCRYPTION
    CategoryAlertMetadata_KARCHIVALRESTORE
    CategoryAlertMetadata_KREMOTEREPLICATION
    CategoryAlertMetadata_KQUOTA
    CategoryAlertMetadata_KLICENSE
    CategoryAlertMetadata_KHELIOSPROACTIVEWELLNESS
    CategoryAlertMetadata_KHELIOSANALYTICSJOBS
    CategoryAlertMetadata_KHELIOSSIGNATUREJOBS
    CategoryAlertMetadata_KSECURITY
)

func (r CategoryAlertMetadataEnum) MarshalJSON() ([]byte, error) {
    s := CategoryAlertMetadataEnumToValue(r)
    return json.Marshal(s)
}

func (r *CategoryAlertMetadataEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CategoryAlertMetadataEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CategoryAlertMetadataEnum to its string representation
 */
func CategoryAlertMetadataEnumToValue(categoryAlertMetadataEnum CategoryAlertMetadataEnum) string {
    switch categoryAlertMetadataEnum {
        case CategoryAlertMetadata_KDISK:
    		return "kDisk"
        case CategoryAlertMetadata_KNODE:
    		return "kNode"
        case CategoryAlertMetadata_KCLUSTER:
    		return "kCluster"
        case CategoryAlertMetadata_KNODEHEALTH:
    		return "kNodeHealth"
        case CategoryAlertMetadata_KCLUSTERHEALTH:
    		return "kClusterHealth"
        case CategoryAlertMetadata_KBACKUPRESTORE:
    		return "kBackupRestore"
        case CategoryAlertMetadata_KENCRYPTION:
    		return "kEncryption"
        case CategoryAlertMetadata_KARCHIVALRESTORE:
    		return "kArchivalRestore"
        case CategoryAlertMetadata_KREMOTEREPLICATION:
    		return "kRemoteReplication"
        case CategoryAlertMetadata_KQUOTA:
    		return "kQuota"
        case CategoryAlertMetadata_KLICENSE:
    		return "kLicense"
        case CategoryAlertMetadata_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"
        case CategoryAlertMetadata_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"
        case CategoryAlertMetadata_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"
        case CategoryAlertMetadata_KSECURITY:
    		return "kSecurity"
        default:
        	return "kDisk"
    }
}

/**
 * Converts CategoryAlertMetadataEnum Array to its string Array representation
*/
func CategoryAlertMetadataEnumArrayToValue(categoryAlertMetadataEnum []CategoryAlertMetadataEnum) []string {
    convArray := make([]string,len( categoryAlertMetadataEnum))
    for i:=0; i<len(categoryAlertMetadataEnum);i++ {
        convArray[i] = CategoryAlertMetadataEnumToValue(categoryAlertMetadataEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CategoryAlertMetadataEnumFromValue(value string) CategoryAlertMetadataEnum {
    switch value {
        case "kDisk":
            return CategoryAlertMetadata_KDISK
        case "kNode":
            return CategoryAlertMetadata_KNODE
        case "kCluster":
            return CategoryAlertMetadata_KCLUSTER
        case "kNodeHealth":
            return CategoryAlertMetadata_KNODEHEALTH
        case "kClusterHealth":
            return CategoryAlertMetadata_KCLUSTERHEALTH
        case "kBackupRestore":
            return CategoryAlertMetadata_KBACKUPRESTORE
        case "kEncryption":
            return CategoryAlertMetadata_KENCRYPTION
        case "kArchivalRestore":
            return CategoryAlertMetadata_KARCHIVALRESTORE
        case "kRemoteReplication":
            return CategoryAlertMetadata_KREMOTEREPLICATION
        case "kQuota":
            return CategoryAlertMetadata_KQUOTA
        case "kLicense":
            return CategoryAlertMetadata_KLICENSE
        case "kHeliosProActiveWellness":
            return CategoryAlertMetadata_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return CategoryAlertMetadata_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return CategoryAlertMetadata_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return CategoryAlertMetadata_KSECURITY
        default:
            return CategoryAlertMetadata_KDISK
    }
}
