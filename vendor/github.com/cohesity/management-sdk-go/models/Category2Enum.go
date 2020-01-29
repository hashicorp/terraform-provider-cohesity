package models

import(
    "encoding/json"
)

/**
 * Type definition for Category2Enum enum
 */
type Category2Enum int

/**
 * Value collection for Category2Enum enum
 */
const (
    Category2_KDISK Category2Enum = 1 + iota
    Category2_KNODE
    Category2_KCLUSTER
    Category2_KNODEHEALTH
    Category2_KCLUSTERHEALTH
    Category2_KBACKUPRESTORE
    Category2_KENCRYPTION
    Category2_KARCHIVALRESTORE
    Category2_KREMOTEREPLICATION
    Category2_KQUOTA
    Category2_KLICENSE
    Category2_KHELIOSPROACTIVEWELLNESS
    Category2_KHELIOSANALYTICSJOBS
    Category2_KHELIOSSIGNATUREJOBS
    Category2_KSECURITY
)

func (r Category2Enum) MarshalJSON() ([]byte, error) { 
    s := Category2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Category2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Category2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Category2Enum to its string representation
 */
func Category2EnumToValue(category2Enum Category2Enum) string {
    switch category2Enum {
        case Category2_KDISK:
    		return "kDisk"		
        case Category2_KNODE:
    		return "kNode"		
        case Category2_KCLUSTER:
    		return "kCluster"		
        case Category2_KNODEHEALTH:
    		return "kNodeHealth"		
        case Category2_KCLUSTERHEALTH:
    		return "kClusterHealth"		
        case Category2_KBACKUPRESTORE:
    		return "kBackupRestore"		
        case Category2_KENCRYPTION:
    		return "kEncryption"		
        case Category2_KARCHIVALRESTORE:
    		return "kArchivalRestore"		
        case Category2_KREMOTEREPLICATION:
    		return "kRemoteReplication"		
        case Category2_KQUOTA:
    		return "kQuota"		
        case Category2_KLICENSE:
    		return "kLicense"		
        case Category2_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"		
        case Category2_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"		
        case Category2_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"		
        case Category2_KSECURITY:
    		return "kSecurity"		
        default:
        	return "kDisk"
    }
}

/**
 * Converts Category2Enum Array to its string Array representation
*/
func Category2EnumArrayToValue(category2Enum []Category2Enum) []string {
    convArray := make([]string,len( category2Enum))
    for i:=0; i<len(category2Enum);i++ {
        convArray[i] = Category2EnumToValue(category2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Category2EnumFromValue(value string) Category2Enum {
    switch value {
        case "kDisk":
            return Category2_KDISK
        case "kNode":
            return Category2_KNODE
        case "kCluster":
            return Category2_KCLUSTER
        case "kNodeHealth":
            return Category2_KNODEHEALTH
        case "kClusterHealth":
            return Category2_KCLUSTERHEALTH
        case "kBackupRestore":
            return Category2_KBACKUPRESTORE
        case "kEncryption":
            return Category2_KENCRYPTION
        case "kArchivalRestore":
            return Category2_KARCHIVALRESTORE
        case "kRemoteReplication":
            return Category2_KREMOTEREPLICATION
        case "kQuota":
            return Category2_KQUOTA
        case "kLicense":
            return Category2_KLICENSE
        case "kHeliosProActiveWellness":
            return Category2_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return Category2_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return Category2_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return Category2_KSECURITY
        default:
            return Category2_KDISK
    }
}
