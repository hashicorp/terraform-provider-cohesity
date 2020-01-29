// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PrivilegeIdUserEnum enum
 */
type PrivilegeIdUserEnum int

/**
 * Value collection for PrivilegeIdUserEnum enum
 */
const (
    PrivilegeIdUser_KPRINCIPALVIEW PrivilegeIdUserEnum = 1 + iota
    PrivilegeIdUser_KPRINCIPALMODIFY
    PrivilegeIdUser_KAPPLAUNCH
    PrivilegeIdUser_KAPPSMANAGEMENT
    PrivilegeIdUser_KORGANIZATIONVIEW
    PrivilegeIdUser_KORGANIZATIONMODIFY
    PrivilegeIdUser_KORGANIZATIONIMPERSONATE
    PrivilegeIdUser_KCLONEVIEW
    PrivilegeIdUser_KCLONEMODIFY
    PrivilegeIdUser_KCLUSTERVIEW
    PrivilegeIdUser_KCLUSTERMODIFY
    PrivilegeIdUser_KCLUSTERCREATE
    PrivilegeIdUser_KCLUSTERSUPPORT
    PrivilegeIdUser_KCLUSTERUPGRADE
    PrivilegeIdUser_KCLUSTERREMOTEVIEW
    PrivilegeIdUser_KCLUSTERREMOTEMODIFY
    PrivilegeIdUser_KCLUSTEREXTERNALTARGETVIEW
    PrivilegeIdUser_KCLUSTEREXTERNALTARGETMODIFY
    PrivilegeIdUser_KCLUSTERAUDIT
    PrivilegeIdUser_KALERTVIEW
    PrivilegeIdUser_KALERTMODIFY
    PrivilegeIdUser_KVLANVIEW
    PrivilegeIdUser_KVLANMODIFY
    PrivilegeIdUser_KHYBRIDEXTENDERVIEW
    PrivilegeIdUser_KHYBRIDEXTENDERDOWNLOAD
    PrivilegeIdUser_KADLDAPVIEW
    PrivilegeIdUser_KADLDAPMODIFY
    PrivilegeIdUser_KSCHEDULERVIEW
    PrivilegeIdUser_KSCHEDULERMODIFY
    PrivilegeIdUser_KPROTECTIONVIEW
    PrivilegeIdUser_KPROTECTIONMODIFY
    PrivilegeIdUser_KPROTECTIONJOBOPERATE
    PrivilegeIdUser_KPROTECTIONSOURCEMODIFY
    PrivilegeIdUser_KPROTECTIONPOLICYVIEW
    PrivilegeIdUser_KPROTECTIONPOLICYMODIFY
    PrivilegeIdUser_KRESTOREVIEW
    PrivilegeIdUser_KRESTOREMODIFY
    PrivilegeIdUser_KRESTOREDOWNLOAD
    PrivilegeIdUser_KREMOTERESTORE
    PrivilegeIdUser_KSTORAGEVIEW
    PrivilegeIdUser_KSTORAGEMODIFY
    PrivilegeIdUser_KSTORAGEDOMAINVIEW
    PrivilegeIdUser_KSTORAGEDOMAINMODIFY
    PrivilegeIdUser_KANALYTICSVIEW
    PrivilegeIdUser_KANALYTICSMODIFY
    PrivilegeIdUser_KREPORTSVIEW
    PrivilegeIdUser_KMCMMODIFY
    PrivilegeIdUser_KDATASECURITY
    PrivilegeIdUser_KSMBBACKUP
    PrivilegeIdUser_KSMBRESTORE
    PrivilegeIdUser_KSMBTAKEOWNERSHIP
    PrivilegeIdUser_KSMBAUDITING
    PrivilegeIdUser_KMCMUNREGISTER
    PrivilegeIdUser_KMCMUPGRADE
    PrivilegeIdUser_KMCMMODIFYSUPERADMIN
    PrivilegeIdUser_KMCMVIEWSUPERADMIN
    PrivilegeIdUser_KMCMMODIFYCOHESITYADMIN
    PrivilegeIdUser_KMCMVIEWCOHESITYADMIN
    PrivilegeIdUser_KOBJECTSEARCH
    PrivilegeIdUser_KFILEDATALOCKEXPIRYTIMEDECREASE
)

func (r PrivilegeIdUserEnum) MarshalJSON() ([]byte, error) {
    s := PrivilegeIdUserEnumToValue(r)
    return json.Marshal(s)
}

func (r *PrivilegeIdUserEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PrivilegeIdUserEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PrivilegeIdUserEnum to its string representation
 */
func PrivilegeIdUserEnumToValue(privilegeIdUserEnum PrivilegeIdUserEnum) string {
    switch privilegeIdUserEnum {
        case PrivilegeIdUser_KPRINCIPALVIEW:
    		return "kPrincipalView"
        case PrivilegeIdUser_KPRINCIPALMODIFY:
    		return "kPrincipalModify"
        case PrivilegeIdUser_KAPPLAUNCH:
    		return "kAppLaunch"
        case PrivilegeIdUser_KAPPSMANAGEMENT:
    		return "kAppsManagement"
        case PrivilegeIdUser_KORGANIZATIONVIEW:
    		return "kOrganizationView"
        case PrivilegeIdUser_KORGANIZATIONMODIFY:
    		return "kOrganizationModify"
        case PrivilegeIdUser_KORGANIZATIONIMPERSONATE:
    		return "kOrganizationImpersonate"
        case PrivilegeIdUser_KCLONEVIEW:
    		return "kCloneView"
        case PrivilegeIdUser_KCLONEMODIFY:
    		return "kCloneModify"
        case PrivilegeIdUser_KCLUSTERVIEW:
    		return "kClusterView"
        case PrivilegeIdUser_KCLUSTERMODIFY:
    		return "kClusterModify"
        case PrivilegeIdUser_KCLUSTERCREATE:
    		return "kClusterCreate"
        case PrivilegeIdUser_KCLUSTERSUPPORT:
    		return "kClusterSupport"
        case PrivilegeIdUser_KCLUSTERUPGRADE:
    		return "kClusterUpgrade"
        case PrivilegeIdUser_KCLUSTERREMOTEVIEW:
    		return "kClusterRemoteView"
        case PrivilegeIdUser_KCLUSTERREMOTEMODIFY:
    		return "kClusterRemoteModify"
        case PrivilegeIdUser_KCLUSTEREXTERNALTARGETVIEW:
    		return "kClusterExternalTargetView"
        case PrivilegeIdUser_KCLUSTEREXTERNALTARGETMODIFY:
    		return "kClusterExternalTargetModify"
        case PrivilegeIdUser_KCLUSTERAUDIT:
    		return "kClusterAudit"
        case PrivilegeIdUser_KALERTVIEW:
    		return "kAlertView"
        case PrivilegeIdUser_KALERTMODIFY:
    		return "kAlertModify"
        case PrivilegeIdUser_KVLANVIEW:
    		return "kVlanView"
        case PrivilegeIdUser_KVLANMODIFY:
    		return "kVlanModify"
        case PrivilegeIdUser_KHYBRIDEXTENDERVIEW:
    		return "kHybridExtenderView"
        case PrivilegeIdUser_KHYBRIDEXTENDERDOWNLOAD:
    		return "kHybridExtenderDownload"
        case PrivilegeIdUser_KADLDAPVIEW:
    		return "kAdLdapView"
        case PrivilegeIdUser_KADLDAPMODIFY:
    		return "kAdLdapModify"
        case PrivilegeIdUser_KSCHEDULERVIEW:
    		return "kSchedulerView"
        case PrivilegeIdUser_KSCHEDULERMODIFY:
    		return "kSchedulerModify"
        case PrivilegeIdUser_KPROTECTIONVIEW:
    		return "kProtectionView"
        case PrivilegeIdUser_KPROTECTIONMODIFY:
    		return "kProtectionModify"
        case PrivilegeIdUser_KPROTECTIONJOBOPERATE:
    		return "kProtectionJobOperate"
        case PrivilegeIdUser_KPROTECTIONSOURCEMODIFY:
    		return "kProtectionSourceModify"
        case PrivilegeIdUser_KPROTECTIONPOLICYVIEW:
    		return "kProtectionPolicyView"
        case PrivilegeIdUser_KPROTECTIONPOLICYMODIFY:
    		return "kProtectionPolicyModify"
        case PrivilegeIdUser_KRESTOREVIEW:
    		return "kRestoreView"
        case PrivilegeIdUser_KRESTOREMODIFY:
    		return "kRestoreModify"
        case PrivilegeIdUser_KRESTOREDOWNLOAD:
    		return "kRestoreDownload"
        case PrivilegeIdUser_KREMOTERESTORE:
    		return "kRemoteRestore"
        case PrivilegeIdUser_KSTORAGEVIEW:
    		return "kStorageView"
        case PrivilegeIdUser_KSTORAGEMODIFY:
    		return "kStorageModify"
        case PrivilegeIdUser_KSTORAGEDOMAINVIEW:
    		return "kStorageDomainView"
        case PrivilegeIdUser_KSTORAGEDOMAINMODIFY:
    		return "kStorageDomainModify"
        case PrivilegeIdUser_KANALYTICSVIEW:
    		return "kAnalyticsView"
        case PrivilegeIdUser_KANALYTICSMODIFY:
    		return "kAnalyticsModify"
        case PrivilegeIdUser_KREPORTSVIEW:
    		return "kReportsView"
        case PrivilegeIdUser_KMCMMODIFY:
    		return "kMcmModify"
        case PrivilegeIdUser_KDATASECURITY:
    		return "kDataSecurity"
        case PrivilegeIdUser_KSMBBACKUP:
    		return "kSmbBackup"
        case PrivilegeIdUser_KSMBRESTORE:
    		return "kSmbRestore"
        case PrivilegeIdUser_KSMBTAKEOWNERSHIP:
    		return "kSmbTakeOwnership"
        case PrivilegeIdUser_KSMBAUDITING:
    		return "kSmbAuditing"
        case PrivilegeIdUser_KMCMUNREGISTER:
    		return "kMcmUnregister"
        case PrivilegeIdUser_KMCMUPGRADE:
    		return "kMcmUpgrade"
        case PrivilegeIdUser_KMCMMODIFYSUPERADMIN:
    		return "kMcmModifySuperAdmin"
        case PrivilegeIdUser_KMCMVIEWSUPERADMIN:
    		return "kMcmViewSuperAdmin"
        case PrivilegeIdUser_KMCMMODIFYCOHESITYADMIN:
    		return "kMcmModifyCohesityAdmin"
        case PrivilegeIdUser_KMCMVIEWCOHESITYADMIN:
    		return "kMcmViewCohesityAdmin"
        case PrivilegeIdUser_KOBJECTSEARCH:
    		return "kObjectSearch"
        case PrivilegeIdUser_KFILEDATALOCKEXPIRYTIMEDECREASE:
    		return "kFileDatalockExpiryTimeDecrease"
        default:
        	return "kPrincipalView"
    }
}

/**
 * Converts PrivilegeIdUserEnum Array to its string Array representation
*/
func PrivilegeIdUserEnumArrayToValue(privilegeIdUserEnum []PrivilegeIdUserEnum) []string {
    convArray := make([]string,len( privilegeIdUserEnum))
    for i:=0; i<len(privilegeIdUserEnum);i++ {
        convArray[i] = PrivilegeIdUserEnumToValue(privilegeIdUserEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PrivilegeIdUserEnumFromValue(value string) PrivilegeIdUserEnum {
    switch value {
        case "kPrincipalView":
            return PrivilegeIdUser_KPRINCIPALVIEW
        case "kPrincipalModify":
            return PrivilegeIdUser_KPRINCIPALMODIFY
        case "kAppLaunch":
            return PrivilegeIdUser_KAPPLAUNCH
        case "kAppsManagement":
            return PrivilegeIdUser_KAPPSMANAGEMENT
        case "kOrganizationView":
            return PrivilegeIdUser_KORGANIZATIONVIEW
        case "kOrganizationModify":
            return PrivilegeIdUser_KORGANIZATIONMODIFY
        case "kOrganizationImpersonate":
            return PrivilegeIdUser_KORGANIZATIONIMPERSONATE
        case "kCloneView":
            return PrivilegeIdUser_KCLONEVIEW
        case "kCloneModify":
            return PrivilegeIdUser_KCLONEMODIFY
        case "kClusterView":
            return PrivilegeIdUser_KCLUSTERVIEW
        case "kClusterModify":
            return PrivilegeIdUser_KCLUSTERMODIFY
        case "kClusterCreate":
            return PrivilegeIdUser_KCLUSTERCREATE
        case "kClusterSupport":
            return PrivilegeIdUser_KCLUSTERSUPPORT
        case "kClusterUpgrade":
            return PrivilegeIdUser_KCLUSTERUPGRADE
        case "kClusterRemoteView":
            return PrivilegeIdUser_KCLUSTERREMOTEVIEW
        case "kClusterRemoteModify":
            return PrivilegeIdUser_KCLUSTERREMOTEMODIFY
        case "kClusterExternalTargetView":
            return PrivilegeIdUser_KCLUSTEREXTERNALTARGETVIEW
        case "kClusterExternalTargetModify":
            return PrivilegeIdUser_KCLUSTEREXTERNALTARGETMODIFY
        case "kClusterAudit":
            return PrivilegeIdUser_KCLUSTERAUDIT
        case "kAlertView":
            return PrivilegeIdUser_KALERTVIEW
        case "kAlertModify":
            return PrivilegeIdUser_KALERTMODIFY
        case "kVlanView":
            return PrivilegeIdUser_KVLANVIEW
        case "kVlanModify":
            return PrivilegeIdUser_KVLANMODIFY
        case "kHybridExtenderView":
            return PrivilegeIdUser_KHYBRIDEXTENDERVIEW
        case "kHybridExtenderDownload":
            return PrivilegeIdUser_KHYBRIDEXTENDERDOWNLOAD
        case "kAdLdapView":
            return PrivilegeIdUser_KADLDAPVIEW
        case "kAdLdapModify":
            return PrivilegeIdUser_KADLDAPMODIFY
        case "kSchedulerView":
            return PrivilegeIdUser_KSCHEDULERVIEW
        case "kSchedulerModify":
            return PrivilegeIdUser_KSCHEDULERMODIFY
        case "kProtectionView":
            return PrivilegeIdUser_KPROTECTIONVIEW
        case "kProtectionModify":
            return PrivilegeIdUser_KPROTECTIONMODIFY
        case "kProtectionJobOperate":
            return PrivilegeIdUser_KPROTECTIONJOBOPERATE
        case "kProtectionSourceModify":
            return PrivilegeIdUser_KPROTECTIONSOURCEMODIFY
        case "kProtectionPolicyView":
            return PrivilegeIdUser_KPROTECTIONPOLICYVIEW
        case "kProtectionPolicyModify":
            return PrivilegeIdUser_KPROTECTIONPOLICYMODIFY
        case "kRestoreView":
            return PrivilegeIdUser_KRESTOREVIEW
        case "kRestoreModify":
            return PrivilegeIdUser_KRESTOREMODIFY
        case "kRestoreDownload":
            return PrivilegeIdUser_KRESTOREDOWNLOAD
        case "kRemoteRestore":
            return PrivilegeIdUser_KREMOTERESTORE
        case "kStorageView":
            return PrivilegeIdUser_KSTORAGEVIEW
        case "kStorageModify":
            return PrivilegeIdUser_KSTORAGEMODIFY
        case "kStorageDomainView":
            return PrivilegeIdUser_KSTORAGEDOMAINVIEW
        case "kStorageDomainModify":
            return PrivilegeIdUser_KSTORAGEDOMAINMODIFY
        case "kAnalyticsView":
            return PrivilegeIdUser_KANALYTICSVIEW
        case "kAnalyticsModify":
            return PrivilegeIdUser_KANALYTICSMODIFY
        case "kReportsView":
            return PrivilegeIdUser_KREPORTSVIEW
        case "kMcmModify":
            return PrivilegeIdUser_KMCMMODIFY
        case "kDataSecurity":
            return PrivilegeIdUser_KDATASECURITY
        case "kSmbBackup":
            return PrivilegeIdUser_KSMBBACKUP
        case "kSmbRestore":
            return PrivilegeIdUser_KSMBRESTORE
        case "kSmbTakeOwnership":
            return PrivilegeIdUser_KSMBTAKEOWNERSHIP
        case "kSmbAuditing":
            return PrivilegeIdUser_KSMBAUDITING
        case "kMcmUnregister":
            return PrivilegeIdUser_KMCMUNREGISTER
        case "kMcmUpgrade":
            return PrivilegeIdUser_KMCMUPGRADE
        case "kMcmModifySuperAdmin":
            return PrivilegeIdUser_KMCMMODIFYSUPERADMIN
        case "kMcmViewSuperAdmin":
            return PrivilegeIdUser_KMCMVIEWSUPERADMIN
        case "kMcmModifyCohesityAdmin":
            return PrivilegeIdUser_KMCMMODIFYCOHESITYADMIN
        case "kMcmViewCohesityAdmin":
            return PrivilegeIdUser_KMCMVIEWCOHESITYADMIN
        case "kObjectSearch":
            return PrivilegeIdUser_KOBJECTSEARCH
        case "kFileDatalockExpiryTimeDecrease":
            return PrivilegeIdUser_KFILEDATALOCKEXPIRYTIMEDECREASE
        default:
            return PrivilegeIdUser_KPRINCIPALVIEW
    }
}
