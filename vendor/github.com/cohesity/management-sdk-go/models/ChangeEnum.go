// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ChangeEnum enum
 */
type ChangeEnum int

/**
 * Value collection for ChangeEnum enum
 */
const (
    Change_KPROTECTIONJOBNAME ChangeEnum = 1 + iota
    Change_KPROTECTIONJOBDESCRIPTION
    Change_KPROTECTIONJOBSOURCES
    Change_KPROTECTIONJOBSCHEDULE
    Change_KPROTECTIONJOBFULLSCHEDULE
    Change_KPROTECTIONJOBRETRYSETTINGS
    Change_KPROTECTIONJOBRETENTIONPOLICY
    Change_KPROTECTIONJOBINDEXINGPOLICY
    Change_KPROTECTIONJOBALERTINGPOLICY
    Change_KPROTECTIONJOBPRIORITY
    Change_KPROTECTIONJOBQUIESCE
    Change_KPROTECTIONJOBSLA
    Change_KPROTECTIONJOBPOLICYID
    Change_KPROTECTIONJOBTIMEZONE
    Change_KPROTECTIONJOBFUTURERUNSPAUSED
    Change_KPROTECTIONJOBFUTURERUNSRESUMED
    Change_KSNAPSHOTTARGETPOLICY
    Change_KPROTECTIONJOBBLACKOUTWINDOW
    Change_KPROTECTIONJOBQOS
    Change_KPROTECTIONJOBINVALIDFIELD
)

func (r ChangeEnum) MarshalJSON() ([]byte, error) {
    s := ChangeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ChangeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ChangeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ChangeEnum to its string representation
 */
func ChangeEnumToValue(changeEnum ChangeEnum) string {
    switch changeEnum {
        case Change_KPROTECTIONJOBNAME:
    		return "kProtectionJobName"
        case Change_KPROTECTIONJOBDESCRIPTION:
    		return "kProtectionJobDescription"
        case Change_KPROTECTIONJOBSOURCES:
    		return "kProtectionJobSources"
        case Change_KPROTECTIONJOBSCHEDULE:
    		return "kProtectionJobSchedule"
        case Change_KPROTECTIONJOBFULLSCHEDULE:
    		return "kProtectionJobFullSchedule"
        case Change_KPROTECTIONJOBRETRYSETTINGS:
    		return "kProtectionJobRetrySettings"
        case Change_KPROTECTIONJOBRETENTIONPOLICY:
    		return "kProtectionJobRetentionPolicy"
        case Change_KPROTECTIONJOBINDEXINGPOLICY:
    		return "kProtectionJobIndexingPolicy"
        case Change_KPROTECTIONJOBALERTINGPOLICY:
    		return "kProtectionJobAlertingPolicy"
        case Change_KPROTECTIONJOBPRIORITY:
    		return "kProtectionJobPriority"
        case Change_KPROTECTIONJOBQUIESCE:
    		return "kProtectionJobQuiesce"
        case Change_KPROTECTIONJOBSLA:
    		return "kProtectionJobSla"
        case Change_KPROTECTIONJOBPOLICYID:
    		return "kProtectionJobPolicyId"
        case Change_KPROTECTIONJOBTIMEZONE:
    		return "kProtectionJobTimezone"
        case Change_KPROTECTIONJOBFUTURERUNSPAUSED:
    		return "kProtectionJobFutureRunsPaused"
        case Change_KPROTECTIONJOBFUTURERUNSRESUMED:
    		return "kProtectionJobFutureRunsResumed"
        case Change_KSNAPSHOTTARGETPOLICY:
    		return "kSnapshotTargetPolicy"
        case Change_KPROTECTIONJOBBLACKOUTWINDOW:
    		return "kProtectionJobBlackoutWindow"
        case Change_KPROTECTIONJOBQOS:
    		return "kProtectionJobQOS"
        case Change_KPROTECTIONJOBINVALIDFIELD:
    		return "kProtectionJobInvalidField"
        default:
        	return "kProtectionJobName"
    }
}

/**
 * Converts ChangeEnum Array to its string Array representation
*/
func ChangeEnumArrayToValue(changeEnum []ChangeEnum) []string {
    convArray := make([]string,len( changeEnum))
    for i:=0; i<len(changeEnum);i++ {
        convArray[i] = ChangeEnumToValue(changeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ChangeEnumFromValue(value string) ChangeEnum {
    switch value {
        case "kProtectionJobName":
            return Change_KPROTECTIONJOBNAME
        case "kProtectionJobDescription":
            return Change_KPROTECTIONJOBDESCRIPTION
        case "kProtectionJobSources":
            return Change_KPROTECTIONJOBSOURCES
        case "kProtectionJobSchedule":
            return Change_KPROTECTIONJOBSCHEDULE
        case "kProtectionJobFullSchedule":
            return Change_KPROTECTIONJOBFULLSCHEDULE
        case "kProtectionJobRetrySettings":
            return Change_KPROTECTIONJOBRETRYSETTINGS
        case "kProtectionJobRetentionPolicy":
            return Change_KPROTECTIONJOBRETENTIONPOLICY
        case "kProtectionJobIndexingPolicy":
            return Change_KPROTECTIONJOBINDEXINGPOLICY
        case "kProtectionJobAlertingPolicy":
            return Change_KPROTECTIONJOBALERTINGPOLICY
        case "kProtectionJobPriority":
            return Change_KPROTECTIONJOBPRIORITY
        case "kProtectionJobQuiesce":
            return Change_KPROTECTIONJOBQUIESCE
        case "kProtectionJobSla":
            return Change_KPROTECTIONJOBSLA
        case "kProtectionJobPolicyId":
            return Change_KPROTECTIONJOBPOLICYID
        case "kProtectionJobTimezone":
            return Change_KPROTECTIONJOBTIMEZONE
        case "kProtectionJobFutureRunsPaused":
            return Change_KPROTECTIONJOBFUTURERUNSPAUSED
        case "kProtectionJobFutureRunsResumed":
            return Change_KPROTECTIONJOBFUTURERUNSRESUMED
        case "kSnapshotTargetPolicy":
            return Change_KSNAPSHOTTARGETPOLICY
        case "kProtectionJobBlackoutWindow":
            return Change_KPROTECTIONJOBBLACKOUTWINDOW
        case "kProtectionJobQOS":
            return Change_KPROTECTIONJOBQOS
        case "kProtectionJobInvalidField":
            return Change_KPROTECTIONJOBINVALIDFIELD
        default:
            return Change_KPROTECTIONJOBNAME
    }
}
