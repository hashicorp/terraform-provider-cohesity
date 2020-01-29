// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PrivilegeIdUserParametersEnum enum
 */
type PrivilegeIdUserParametersEnum int

/**
 * Value collection for PrivilegeIdUserParametersEnum enum
 */
const (
    PrivilegeIdUserParameters_KPRINCIPALVIEW PrivilegeIdUserParametersEnum = 1 + iota
    PrivilegeIdUserParameters_KPRINCIPALMODIFY
    PrivilegeIdUserParameters_KAPPLAUNCH
    PrivilegeIdUserParameters_KAPPSMANAGEMENT
    PrivilegeIdUserParameters_KORGANIZATIONVIEW
    PrivilegeIdUserParameters_KORGANIZATIONMODIFY
    PrivilegeIdUserParameters_KORGANIZATIONIMPERSONATE
    PrivilegeIdUserParameters_KCLONEVIEW
    PrivilegeIdUserParameters_KCLONEMODIFY
    PrivilegeIdUserParameters_KCLUSTERVIEW
    PrivilegeIdUserParameters_KCLUSTERMODIFY
    PrivilegeIdUserParameters_KCLUSTERCREATE
    PrivilegeIdUserParameters_KCLUSTERSUPPORT
    PrivilegeIdUserParameters_KCLUSTERUPGRADE
    PrivilegeIdUserParameters_KCLUSTERREMOTEVIEW
    PrivilegeIdUserParameters_KCLUSTERREMOTEMODIFY
    PrivilegeIdUserParameters_KCLUSTEREXTERNALTARGETVIEW
    PrivilegeIdUserParameters_KCLUSTEREXTERNALTARGETMODIFY
    PrivilegeIdUserParameters_KCLUSTERAUDIT
    PrivilegeIdUserParameters_KALERTVIEW
    PrivilegeIdUserParameters_KALERTMODIFY
    PrivilegeIdUserParameters_KVLANVIEW
    PrivilegeIdUserParameters_KVLANMODIFY
    PrivilegeIdUserParameters_KHYBRIDEXTENDERVIEW
    PrivilegeIdUserParameters_KHYBRIDEXTENDERDOWNLOAD
    PrivilegeIdUserParameters_KADLDAPVIEW
    PrivilegeIdUserParameters_KADLDAPMODIFY
    PrivilegeIdUserParameters_KSCHEDULERVIEW
    PrivilegeIdUserParameters_KSCHEDULERMODIFY
    PrivilegeIdUserParameters_KPROTECTIONVIEW
    PrivilegeIdUserParameters_KPROTECTIONMODIFY
    PrivilegeIdUserParameters_KPROTECTIONJOBOPERATE
    PrivilegeIdUserParameters_KPROTECTIONSOURCEMODIFY
    PrivilegeIdUserParameters_KPROTECTIONPOLICYVIEW
    PrivilegeIdUserParameters_KPROTECTIONPOLICYMODIFY
    PrivilegeIdUserParameters_KRESTOREVIEW
    PrivilegeIdUserParameters_KRESTOREMODIFY
    PrivilegeIdUserParameters_KRESTOREDOWNLOAD
    PrivilegeIdUserParameters_KREMOTERESTORE
    PrivilegeIdUserParameters_KSTORAGEVIEW
    PrivilegeIdUserParameters_KSTORAGEMODIFY
    PrivilegeIdUserParameters_KSTORAGEDOMAINVIEW
    PrivilegeIdUserParameters_KSTORAGEDOMAINMODIFY
    PrivilegeIdUserParameters_KANALYTICSVIEW
    PrivilegeIdUserParameters_KANALYTICSMODIFY
    PrivilegeIdUserParameters_KREPORTSVIEW
    PrivilegeIdUserParameters_KMCMMODIFY
    PrivilegeIdUserParameters_KDATASECURITY
    PrivilegeIdUserParameters_KSMBBACKUP
    PrivilegeIdUserParameters_KSMBRESTORE
    PrivilegeIdUserParameters_KSMBTAKEOWNERSHIP
    PrivilegeIdUserParameters_KSMBAUDITING
    PrivilegeIdUserParameters_KMCMUNREGISTER
    PrivilegeIdUserParameters_KMCMUPGRADE
    PrivilegeIdUserParameters_KMCMMODIFYSUPERADMIN
    PrivilegeIdUserParameters_KMCMVIEWSUPERADMIN
    PrivilegeIdUserParameters_KMCMMODIFYCOHESITYADMIN
    PrivilegeIdUserParameters_KMCMVIEWCOHESITYADMIN
    PrivilegeIdUserParameters_KOBJECTSEARCH
    PrivilegeIdUserParameters_KFILEDATALOCKEXPIRYTIMEDECREASE
)

func (r PrivilegeIdUserParametersEnum) MarshalJSON() ([]byte, error) {
    s := PrivilegeIdUserParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *PrivilegeIdUserParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PrivilegeIdUserParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PrivilegeIdUserParametersEnum to its string representation
 */
func PrivilegeIdUserParametersEnumToValue(privilegeIdUserParametersEnum PrivilegeIdUserParametersEnum) string {
    switch privilegeIdUserParametersEnum {
        case PrivilegeIdUserParameters_KPRINCIPALVIEW:
    		return "kPrincipalView"
        case PrivilegeIdUserParameters_KPRINCIPALMODIFY:
    		return "kPrincipalModify"
        case PrivilegeIdUserParameters_KAPPLAUNCH:
    		return "kAppLaunch"
        case PrivilegeIdUserParameters_KAPPSMANAGEMENT:
    		return "kAppsManagement"
        case PrivilegeIdUserParameters_KORGANIZATIONVIEW:
    		return "kOrganizationView"
        case PrivilegeIdUserParameters_KORGANIZATIONMODIFY:
    		return "kOrganizationModify"
        case PrivilegeIdUserParameters_KORGANIZATIONIMPERSONATE:
    		return "kOrganizationImpersonate"
        case PrivilegeIdUserParameters_KCLONEVIEW:
    		return "kCloneView"
        case PrivilegeIdUserParameters_KCLONEMODIFY:
    		return "kCloneModify"
        case PrivilegeIdUserParameters_KCLUSTERVIEW:
    		return "kClusterView"
        case PrivilegeIdUserParameters_KCLUSTERMODIFY:
    		return "kClusterModify"
        case PrivilegeIdUserParameters_KCLUSTERCREATE:
    		return "kClusterCreate"
        case PrivilegeIdUserParameters_KCLUSTERSUPPORT:
    		return "kClusterSupport"
        case PrivilegeIdUserParameters_KCLUSTERUPGRADE:
    		return "kClusterUpgrade"
        case PrivilegeIdUserParameters_KCLUSTERREMOTEVIEW:
    		return "kClusterRemoteView"
        case PrivilegeIdUserParameters_KCLUSTERREMOTEMODIFY:
    		return "kClusterRemoteModify"
        case PrivilegeIdUserParameters_KCLUSTEREXTERNALTARGETVIEW:
    		return "kClusterExternalTargetView"
        case PrivilegeIdUserParameters_KCLUSTEREXTERNALTARGETMODIFY:
    		return "kClusterExternalTargetModify"
        case PrivilegeIdUserParameters_KCLUSTERAUDIT:
    		return "kClusterAudit"
        case PrivilegeIdUserParameters_KALERTVIEW:
    		return "kAlertView"
        case PrivilegeIdUserParameters_KALERTMODIFY:
    		return "kAlertModify"
        case PrivilegeIdUserParameters_KVLANVIEW:
    		return "kVlanView"
        case PrivilegeIdUserParameters_KVLANMODIFY:
    		return "kVlanModify"
        case PrivilegeIdUserParameters_KHYBRIDEXTENDERVIEW:
    		return "kHybridExtenderView"
        case PrivilegeIdUserParameters_KHYBRIDEXTENDERDOWNLOAD:
    		return "kHybridExtenderDownload"
        case PrivilegeIdUserParameters_KADLDAPVIEW:
    		return "kAdLdapView"
        case PrivilegeIdUserParameters_KADLDAPMODIFY:
    		return "kAdLdapModify"
        case PrivilegeIdUserParameters_KSCHEDULERVIEW:
    		return "kSchedulerView"
        case PrivilegeIdUserParameters_KSCHEDULERMODIFY:
    		return "kSchedulerModify"
        case PrivilegeIdUserParameters_KPROTECTIONVIEW:
    		return "kProtectionView"
        case PrivilegeIdUserParameters_KPROTECTIONMODIFY:
    		return "kProtectionModify"
        case PrivilegeIdUserParameters_KPROTECTIONJOBOPERATE:
    		return "kProtectionJobOperate"
        case PrivilegeIdUserParameters_KPROTECTIONSOURCEMODIFY:
    		return "kProtectionSourceModify"
        case PrivilegeIdUserParameters_KPROTECTIONPOLICYVIEW:
    		return "kProtectionPolicyView"
        case PrivilegeIdUserParameters_KPROTECTIONPOLICYMODIFY:
    		return "kProtectionPolicyModify"
        case PrivilegeIdUserParameters_KRESTOREVIEW:
    		return "kRestoreView"
        case PrivilegeIdUserParameters_KRESTOREMODIFY:
    		return "kRestoreModify"
        case PrivilegeIdUserParameters_KRESTOREDOWNLOAD:
    		return "kRestoreDownload"
        case PrivilegeIdUserParameters_KREMOTERESTORE:
    		return "kRemoteRestore"
        case PrivilegeIdUserParameters_KSTORAGEVIEW:
    		return "kStorageView"
        case PrivilegeIdUserParameters_KSTORAGEMODIFY:
    		return "kStorageModify"
        case PrivilegeIdUserParameters_KSTORAGEDOMAINVIEW:
    		return "kStorageDomainView"
        case PrivilegeIdUserParameters_KSTORAGEDOMAINMODIFY:
    		return "kStorageDomainModify"
        case PrivilegeIdUserParameters_KANALYTICSVIEW:
    		return "kAnalyticsView"
        case PrivilegeIdUserParameters_KANALYTICSMODIFY:
    		return "kAnalyticsModify"
        case PrivilegeIdUserParameters_KREPORTSVIEW:
    		return "kReportsView"
        case PrivilegeIdUserParameters_KMCMMODIFY:
    		return "kMcmModify"
        case PrivilegeIdUserParameters_KDATASECURITY:
    		return "kDataSecurity"
        case PrivilegeIdUserParameters_KSMBBACKUP:
    		return "kSmbBackup"
        case PrivilegeIdUserParameters_KSMBRESTORE:
    		return "kSmbRestore"
        case PrivilegeIdUserParameters_KSMBTAKEOWNERSHIP:
    		return "kSmbTakeOwnership"
        case PrivilegeIdUserParameters_KSMBAUDITING:
    		return "kSmbAuditing"
        case PrivilegeIdUserParameters_KMCMUNREGISTER:
    		return "kMcmUnregister"
        case PrivilegeIdUserParameters_KMCMUPGRADE:
    		return "kMcmUpgrade"
        case PrivilegeIdUserParameters_KMCMMODIFYSUPERADMIN:
    		return "kMcmModifySuperAdmin"
        case PrivilegeIdUserParameters_KMCMVIEWSUPERADMIN:
    		return "kMcmViewSuperAdmin"
        case PrivilegeIdUserParameters_KMCMMODIFYCOHESITYADMIN:
    		return "kMcmModifyCohesityAdmin"
        case PrivilegeIdUserParameters_KMCMVIEWCOHESITYADMIN:
    		return "kMcmViewCohesityAdmin"
        case PrivilegeIdUserParameters_KOBJECTSEARCH:
    		return "kObjectSearch"
        case PrivilegeIdUserParameters_KFILEDATALOCKEXPIRYTIMEDECREASE:
    		return "kFileDatalockExpiryTimeDecrease"
        default:
        	return "kPrincipalView"
    }
}

/**
 * Converts PrivilegeIdUserParametersEnum Array to its string Array representation
*/
func PrivilegeIdUserParametersEnumArrayToValue(privilegeIdUserParametersEnum []PrivilegeIdUserParametersEnum) []string {
    convArray := make([]string,len( privilegeIdUserParametersEnum))
    for i:=0; i<len(privilegeIdUserParametersEnum);i++ {
        convArray[i] = PrivilegeIdUserParametersEnumToValue(privilegeIdUserParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PrivilegeIdUserParametersEnumFromValue(value string) PrivilegeIdUserParametersEnum {
    switch value {
        case "kPrincipalView":
            return PrivilegeIdUserParameters_KPRINCIPALVIEW
        case "kPrincipalModify":
            return PrivilegeIdUserParameters_KPRINCIPALMODIFY
        case "kAppLaunch":
            return PrivilegeIdUserParameters_KAPPLAUNCH
        case "kAppsManagement":
            return PrivilegeIdUserParameters_KAPPSMANAGEMENT
        case "kOrganizationView":
            return PrivilegeIdUserParameters_KORGANIZATIONVIEW
        case "kOrganizationModify":
            return PrivilegeIdUserParameters_KORGANIZATIONMODIFY
        case "kOrganizationImpersonate":
            return PrivilegeIdUserParameters_KORGANIZATIONIMPERSONATE
        case "kCloneView":
            return PrivilegeIdUserParameters_KCLONEVIEW
        case "kCloneModify":
            return PrivilegeIdUserParameters_KCLONEMODIFY
        case "kClusterView":
            return PrivilegeIdUserParameters_KCLUSTERVIEW
        case "kClusterModify":
            return PrivilegeIdUserParameters_KCLUSTERMODIFY
        case "kClusterCreate":
            return PrivilegeIdUserParameters_KCLUSTERCREATE
        case "kClusterSupport":
            return PrivilegeIdUserParameters_KCLUSTERSUPPORT
        case "kClusterUpgrade":
            return PrivilegeIdUserParameters_KCLUSTERUPGRADE
        case "kClusterRemoteView":
            return PrivilegeIdUserParameters_KCLUSTERREMOTEVIEW
        case "kClusterRemoteModify":
            return PrivilegeIdUserParameters_KCLUSTERREMOTEMODIFY
        case "kClusterExternalTargetView":
            return PrivilegeIdUserParameters_KCLUSTEREXTERNALTARGETVIEW
        case "kClusterExternalTargetModify":
            return PrivilegeIdUserParameters_KCLUSTEREXTERNALTARGETMODIFY
        case "kClusterAudit":
            return PrivilegeIdUserParameters_KCLUSTERAUDIT
        case "kAlertView":
            return PrivilegeIdUserParameters_KALERTVIEW
        case "kAlertModify":
            return PrivilegeIdUserParameters_KALERTMODIFY
        case "kVlanView":
            return PrivilegeIdUserParameters_KVLANVIEW
        case "kVlanModify":
            return PrivilegeIdUserParameters_KVLANMODIFY
        case "kHybridExtenderView":
            return PrivilegeIdUserParameters_KHYBRIDEXTENDERVIEW
        case "kHybridExtenderDownload":
            return PrivilegeIdUserParameters_KHYBRIDEXTENDERDOWNLOAD
        case "kAdLdapView":
            return PrivilegeIdUserParameters_KADLDAPVIEW
        case "kAdLdapModify":
            return PrivilegeIdUserParameters_KADLDAPMODIFY
        case "kSchedulerView":
            return PrivilegeIdUserParameters_KSCHEDULERVIEW
        case "kSchedulerModify":
            return PrivilegeIdUserParameters_KSCHEDULERMODIFY
        case "kProtectionView":
            return PrivilegeIdUserParameters_KPROTECTIONVIEW
        case "kProtectionModify":
            return PrivilegeIdUserParameters_KPROTECTIONMODIFY
        case "kProtectionJobOperate":
            return PrivilegeIdUserParameters_KPROTECTIONJOBOPERATE
        case "kProtectionSourceModify":
            return PrivilegeIdUserParameters_KPROTECTIONSOURCEMODIFY
        case "kProtectionPolicyView":
            return PrivilegeIdUserParameters_KPROTECTIONPOLICYVIEW
        case "kProtectionPolicyModify":
            return PrivilegeIdUserParameters_KPROTECTIONPOLICYMODIFY
        case "kRestoreView":
            return PrivilegeIdUserParameters_KRESTOREVIEW
        case "kRestoreModify":
            return PrivilegeIdUserParameters_KRESTOREMODIFY
        case "kRestoreDownload":
            return PrivilegeIdUserParameters_KRESTOREDOWNLOAD
        case "kRemoteRestore":
            return PrivilegeIdUserParameters_KREMOTERESTORE
        case "kStorageView":
            return PrivilegeIdUserParameters_KSTORAGEVIEW
        case "kStorageModify":
            return PrivilegeIdUserParameters_KSTORAGEMODIFY
        case "kStorageDomainView":
            return PrivilegeIdUserParameters_KSTORAGEDOMAINVIEW
        case "kStorageDomainModify":
            return PrivilegeIdUserParameters_KSTORAGEDOMAINMODIFY
        case "kAnalyticsView":
            return PrivilegeIdUserParameters_KANALYTICSVIEW
        case "kAnalyticsModify":
            return PrivilegeIdUserParameters_KANALYTICSMODIFY
        case "kReportsView":
            return PrivilegeIdUserParameters_KREPORTSVIEW
        case "kMcmModify":
            return PrivilegeIdUserParameters_KMCMMODIFY
        case "kDataSecurity":
            return PrivilegeIdUserParameters_KDATASECURITY
        case "kSmbBackup":
            return PrivilegeIdUserParameters_KSMBBACKUP
        case "kSmbRestore":
            return PrivilegeIdUserParameters_KSMBRESTORE
        case "kSmbTakeOwnership":
            return PrivilegeIdUserParameters_KSMBTAKEOWNERSHIP
        case "kSmbAuditing":
            return PrivilegeIdUserParameters_KSMBAUDITING
        case "kMcmUnregister":
            return PrivilegeIdUserParameters_KMCMUNREGISTER
        case "kMcmUpgrade":
            return PrivilegeIdUserParameters_KMCMUPGRADE
        case "kMcmModifySuperAdmin":
            return PrivilegeIdUserParameters_KMCMMODIFYSUPERADMIN
        case "kMcmViewSuperAdmin":
            return PrivilegeIdUserParameters_KMCMVIEWSUPERADMIN
        case "kMcmModifyCohesityAdmin":
            return PrivilegeIdUserParameters_KMCMMODIFYCOHESITYADMIN
        case "kMcmViewCohesityAdmin":
            return PrivilegeIdUserParameters_KMCMVIEWCOHESITYADMIN
        case "kObjectSearch":
            return PrivilegeIdUserParameters_KOBJECTSEARCH
        case "kFileDatalockExpiryTimeDecrease":
            return PrivilegeIdUserParameters_KFILEDATALOCKEXPIRYTIMEDECREASE
        default:
            return PrivilegeIdUserParameters_KPRINCIPALVIEW
    }
}
