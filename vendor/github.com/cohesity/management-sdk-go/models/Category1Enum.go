package models

import(
    "encoding/json"
)

/**
 * Type definition for Category1Enum enum
 */
type Category1Enum int

/**
 * Value collection for Category1Enum enum
 */
const (
    Category1_KDISK Category1Enum = 1 + iota
    Category1_KNODE
    Category1_KCLUSTER
    Category1_KNODEHEALTH
    Category1_KCLUSTERHEALTH
    Category1_KBACKUPRESTORE
    Category1_KENCRYPTION
    Category1_KARCHIVALRESTORE
    Category1_KREMOTEREPLICATION
    Category1_KQUOTA
    Category1_KLICENSE
    Category1_KHELIOSPROACTIVEWELLNESS
    Category1_KHELIOSANALYTICSJOBS
    Category1_KHELIOSSIGNATUREJOBS
    Category1_KSECURITY
)

func (r Category1Enum) MarshalJSON() ([]byte, error) { 
    s := Category1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Category1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Category1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Category1Enum to its string representation
 */
func Category1EnumToValue(category1Enum Category1Enum) string {
    switch category1Enum {
        case Category1_KDISK:
    		return "kDisk"		
        case Category1_KNODE:
    		return "kNode"		
        case Category1_KCLUSTER:
    		return "kCluster"		
        case Category1_KNODEHEALTH:
    		return "kNodeHealth"		
        case Category1_KCLUSTERHEALTH:
    		return "kClusterHealth"		
        case Category1_KBACKUPRESTORE:
    		return "kBackupRestore"		
        case Category1_KENCRYPTION:
    		return "kEncryption"		
        case Category1_KARCHIVALRESTORE:
    		return "kArchivalRestore"		
        case Category1_KREMOTEREPLICATION:
    		return "kRemoteReplication"		
        case Category1_KQUOTA:
    		return "kQuota"		
        case Category1_KLICENSE:
    		return "kLicense"		
        case Category1_KHELIOSPROACTIVEWELLNESS:
    		return "kHeliosProActiveWellness"		
        case Category1_KHELIOSANALYTICSJOBS:
    		return "kHeliosAnalyticsJobs"		
        case Category1_KHELIOSSIGNATUREJOBS:
    		return "kHeliosSignatureJobs"		
        case Category1_KSECURITY:
    		return "kSecurity"		
        default:
        	return "kDisk"
    }
}

/**
 * Converts Category1Enum Array to its string Array representation
*/
func Category1EnumArrayToValue(category1Enum []Category1Enum) []string {
    convArray := make([]string,len( category1Enum))
    for i:=0; i<len(category1Enum);i++ {
        convArray[i] = Category1EnumToValue(category1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Category1EnumFromValue(value string) Category1Enum {
    switch value {
        case "kDisk":
            return Category1_KDISK
        case "kNode":
            return Category1_KNODE
        case "kCluster":
            return Category1_KCLUSTER
        case "kNodeHealth":
            return Category1_KNODEHEALTH
        case "kClusterHealth":
            return Category1_KCLUSTERHEALTH
        case "kBackupRestore":
            return Category1_KBACKUPRESTORE
        case "kEncryption":
            return Category1_KENCRYPTION
        case "kArchivalRestore":
            return Category1_KARCHIVALRESTORE
        case "kRemoteReplication":
            return Category1_KREMOTEREPLICATION
        case "kQuota":
            return Category1_KQUOTA
        case "kLicense":
            return Category1_KLICENSE
        case "kHeliosProActiveWellness":
            return Category1_KHELIOSPROACTIVEWELLNESS
        case "kHeliosAnalyticsJobs":
            return Category1_KHELIOSANALYTICSJOBS
        case "kHeliosSignatureJobs":
            return Category1_KHELIOSSIGNATUREJOBS
        case "kSecurity":
            return Category1_KSECURITY
        default:
            return Category1_KDISK
    }
}
