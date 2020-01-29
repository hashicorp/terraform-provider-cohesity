// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PrivilegeIdEnum enum
 */
type PrivilegeIdEnum int

/**
 * Value collection for PrivilegeIdEnum enum
 */
const (
    PrivilegeId_KPRINCIPALVIEW PrivilegeIdEnum = 1 + iota
    PrivilegeId_KPRINCIPALMODIFY
    PrivilegeId_KAPPLAUNCH
    PrivilegeId_KAPPSMANAGEMENT
    PrivilegeId_KORGANIZATIONVIEW
    PrivilegeId_KORGANIZATIONMODIFY
    PrivilegeId_KORGANIZATIONIMPERSONATE
    PrivilegeId_KCLONEVIEW
    PrivilegeId_KCLONEMODIFY
    PrivilegeId_KCLUSTERVIEW
    PrivilegeId_KCLUSTERMODIFY
    PrivilegeId_KCLUSTERCREATE
    PrivilegeId_KCLUSTERSUPPORT
    PrivilegeId_KCLUSTERUPGRADE
    PrivilegeId_KCLUSTERREMOTEVIEW
    PrivilegeId_KCLUSTERREMOTEMODIFY
    PrivilegeId_KCLUSTEREXTERNALTARGETVIEW
    PrivilegeId_KCLUSTEREXTERNALTARGETMODIFY
    PrivilegeId_KCLUSTERAUDIT
    PrivilegeId_KALERTVIEW
    PrivilegeId_KALERTMODIFY
    PrivilegeId_KVLANVIEW
    PrivilegeId_KVLANMODIFY
    PrivilegeId_KHYBRIDEXTENDERVIEW
    PrivilegeId_KHYBRIDEXTENDERDOWNLOAD
    PrivilegeId_KADLDAPVIEW
    PrivilegeId_KADLDAPMODIFY
    PrivilegeId_KSCHEDULERVIEW
    PrivilegeId_KSCHEDULERMODIFY
    PrivilegeId_KPROTECTIONVIEW
    PrivilegeId_KPROTECTIONMODIFY
    PrivilegeId_KPROTECTIONJOBOPERATE
    PrivilegeId_KPROTECTIONSOURCEMODIFY
    PrivilegeId_KPROTECTIONPOLICYVIEW
    PrivilegeId_KPROTECTIONPOLICYMODIFY
    PrivilegeId_KRESTOREVIEW
    PrivilegeId_KRESTOREMODIFY
    PrivilegeId_KRESTOREDOWNLOAD
    PrivilegeId_KREMOTERESTORE
    PrivilegeId_KSTORAGEVIEW
    PrivilegeId_KSTORAGEMODIFY
    PrivilegeId_KSTORAGEDOMAINVIEW
    PrivilegeId_KSTORAGEDOMAINMODIFY
    PrivilegeId_KANALYTICSVIEW
    PrivilegeId_KANALYTICSMODIFY
    PrivilegeId_KREPORTSVIEW
    PrivilegeId_KMCMMODIFY
    PrivilegeId_KDATASECURITY
    PrivilegeId_KSMBBACKUP
    PrivilegeId_KSMBRESTORE
    PrivilegeId_KSMBTAKEOWNERSHIP
    PrivilegeId_KSMBAUDITING
    PrivilegeId_KMCMUNREGISTER
    PrivilegeId_KMCMUPGRADE
    PrivilegeId_KMCMMODIFYSUPERADMIN
    PrivilegeId_KMCMVIEWSUPERADMIN
    PrivilegeId_KMCMMODIFYCOHESITYADMIN
    PrivilegeId_KMCMVIEWCOHESITYADMIN
    PrivilegeId_KOBJECTSEARCH
    PrivilegeId_KFILEDATALOCKEXPIRYTIMEDECREASE
)

func (r PrivilegeIdEnum) MarshalJSON() ([]byte, error) {
    s := PrivilegeIdEnumToValue(r)
    return json.Marshal(s)
}

func (r *PrivilegeIdEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PrivilegeIdEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PrivilegeIdEnum to its string representation
 */
func PrivilegeIdEnumToValue(privilegeIdEnum PrivilegeIdEnum) string {
    switch privilegeIdEnum {
        case PrivilegeId_KPRINCIPALVIEW:
    		return "kPrincipalView"
        case PrivilegeId_KPRINCIPALMODIFY:
    		return "kPrincipalModify"
        case PrivilegeId_KAPPLAUNCH:
    		return "kAppLaunch"
        case PrivilegeId_KAPPSMANAGEMENT:
    		return "kAppsManagement"
        case PrivilegeId_KORGANIZATIONVIEW:
    		return "kOrganizationView"
        case PrivilegeId_KORGANIZATIONMODIFY:
    		return "kOrganizationModify"
        case PrivilegeId_KORGANIZATIONIMPERSONATE:
    		return "kOrganizationImpersonate"
        case PrivilegeId_KCLONEVIEW:
    		return "kCloneView"
        case PrivilegeId_KCLONEMODIFY:
    		return "kCloneModify"
        case PrivilegeId_KCLUSTERVIEW:
    		return "kClusterView"
        case PrivilegeId_KCLUSTERMODIFY:
    		return "kClusterModify"
        case PrivilegeId_KCLUSTERCREATE:
    		return "kClusterCreate"
        case PrivilegeId_KCLUSTERSUPPORT:
    		return "kClusterSupport"
        case PrivilegeId_KCLUSTERUPGRADE:
    		return "kClusterUpgrade"
        case PrivilegeId_KCLUSTERREMOTEVIEW:
    		return "kClusterRemoteView"
        case PrivilegeId_KCLUSTERREMOTEMODIFY:
    		return "kClusterRemoteModify"
        case PrivilegeId_KCLUSTEREXTERNALTARGETVIEW:
    		return "kClusterExternalTargetView"
        case PrivilegeId_KCLUSTEREXTERNALTARGETMODIFY:
    		return "kClusterExternalTargetModify"
        case PrivilegeId_KCLUSTERAUDIT:
    		return "kClusterAudit"
        case PrivilegeId_KALERTVIEW:
    		return "kAlertView"
        case PrivilegeId_KALERTMODIFY:
    		return "kAlertModify"
        case PrivilegeId_KVLANVIEW:
    		return "kVlanView"
        case PrivilegeId_KVLANMODIFY:
    		return "kVlanModify"
        case PrivilegeId_KHYBRIDEXTENDERVIEW:
    		return "kHybridExtenderView"
        case PrivilegeId_KHYBRIDEXTENDERDOWNLOAD:
    		return "kHybridExtenderDownload"
        case PrivilegeId_KADLDAPVIEW:
    		return "kAdLdapView"
        case PrivilegeId_KADLDAPMODIFY:
    		return "kAdLdapModify"
        case PrivilegeId_KSCHEDULERVIEW:
    		return "kSchedulerView"
        case PrivilegeId_KSCHEDULERMODIFY:
    		return "kSchedulerModify"
        case PrivilegeId_KPROTECTIONVIEW:
    		return "kProtectionView"
        case PrivilegeId_KPROTECTIONMODIFY:
    		return "kProtectionModify"
        case PrivilegeId_KPROTECTIONJOBOPERATE:
    		return "kProtectionJobOperate"
        case PrivilegeId_KPROTECTIONSOURCEMODIFY:
    		return "kProtectionSourceModify"
        case PrivilegeId_KPROTECTIONPOLICYVIEW:
    		return "kProtectionPolicyView"
        case PrivilegeId_KPROTECTIONPOLICYMODIFY:
    		return "kProtectionPolicyModify"
        case PrivilegeId_KRESTOREVIEW:
    		return "kRestoreView"
        case PrivilegeId_KRESTOREMODIFY:
    		return "kRestoreModify"
        case PrivilegeId_KRESTOREDOWNLOAD:
    		return "kRestoreDownload"
        case PrivilegeId_KREMOTERESTORE:
    		return "kRemoteRestore"
        case PrivilegeId_KSTORAGEVIEW:
    		return "kStorageView"
        case PrivilegeId_KSTORAGEMODIFY:
    		return "kStorageModify"
        case PrivilegeId_KSTORAGEDOMAINVIEW:
    		return "kStorageDomainView"
        case PrivilegeId_KSTORAGEDOMAINMODIFY:
    		return "kStorageDomainModify"
        case PrivilegeId_KANALYTICSVIEW:
    		return "kAnalyticsView"
        case PrivilegeId_KANALYTICSMODIFY:
    		return "kAnalyticsModify"
        case PrivilegeId_KREPORTSVIEW:
    		return "kReportsView"
        case PrivilegeId_KMCMMODIFY:
    		return "kMcmModify"
        case PrivilegeId_KDATASECURITY:
    		return "kDataSecurity"
        case PrivilegeId_KSMBBACKUP:
    		return "kSmbBackup"
        case PrivilegeId_KSMBRESTORE:
    		return "kSmbRestore"
        case PrivilegeId_KSMBTAKEOWNERSHIP:
    		return "kSmbTakeOwnership"
        case PrivilegeId_KSMBAUDITING:
    		return "kSmbAuditing"
        case PrivilegeId_KMCMUNREGISTER:
    		return "kMcmUnregister"
        case PrivilegeId_KMCMUPGRADE:
    		return "kMcmUpgrade"
        case PrivilegeId_KMCMMODIFYSUPERADMIN:
    		return "kMcmModifySuperAdmin"
        case PrivilegeId_KMCMVIEWSUPERADMIN:
    		return "kMcmViewSuperAdmin"
        case PrivilegeId_KMCMMODIFYCOHESITYADMIN:
    		return "kMcmModifyCohesityAdmin"
        case PrivilegeId_KMCMVIEWCOHESITYADMIN:
    		return "kMcmViewCohesityAdmin"
        case PrivilegeId_KOBJECTSEARCH:
    		return "kObjectSearch"
        case PrivilegeId_KFILEDATALOCKEXPIRYTIMEDECREASE:
    		return "kFileDatalockExpiryTimeDecrease"
        default:
        	return "kPrincipalView"
    }
}

/**
 * Converts PrivilegeIdEnum Array to its string Array representation
*/
func PrivilegeIdEnumArrayToValue(privilegeIdEnum []PrivilegeIdEnum) []string {
    convArray := make([]string,len( privilegeIdEnum))
    for i:=0; i<len(privilegeIdEnum);i++ {
        convArray[i] = PrivilegeIdEnumToValue(privilegeIdEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PrivilegeIdEnumFromValue(value string) PrivilegeIdEnum {
    switch value {
        case "kPrincipalView":
            return PrivilegeId_KPRINCIPALVIEW
        case "kPrincipalModify":
            return PrivilegeId_KPRINCIPALMODIFY
        case "kAppLaunch":
            return PrivilegeId_KAPPLAUNCH
        case "kAppsManagement":
            return PrivilegeId_KAPPSMANAGEMENT
        case "kOrganizationView":
            return PrivilegeId_KORGANIZATIONVIEW
        case "kOrganizationModify":
            return PrivilegeId_KORGANIZATIONMODIFY
        case "kOrganizationImpersonate":
            return PrivilegeId_KORGANIZATIONIMPERSONATE
        case "kCloneView":
            return PrivilegeId_KCLONEVIEW
        case "kCloneModify":
            return PrivilegeId_KCLONEMODIFY
        case "kClusterView":
            return PrivilegeId_KCLUSTERVIEW
        case "kClusterModify":
            return PrivilegeId_KCLUSTERMODIFY
        case "kClusterCreate":
            return PrivilegeId_KCLUSTERCREATE
        case "kClusterSupport":
            return PrivilegeId_KCLUSTERSUPPORT
        case "kClusterUpgrade":
            return PrivilegeId_KCLUSTERUPGRADE
        case "kClusterRemoteView":
            return PrivilegeId_KCLUSTERREMOTEVIEW
        case "kClusterRemoteModify":
            return PrivilegeId_KCLUSTERREMOTEMODIFY
        case "kClusterExternalTargetView":
            return PrivilegeId_KCLUSTEREXTERNALTARGETVIEW
        case "kClusterExternalTargetModify":
            return PrivilegeId_KCLUSTEREXTERNALTARGETMODIFY
        case "kClusterAudit":
            return PrivilegeId_KCLUSTERAUDIT
        case "kAlertView":
            return PrivilegeId_KALERTVIEW
        case "kAlertModify":
            return PrivilegeId_KALERTMODIFY
        case "kVlanView":
            return PrivilegeId_KVLANVIEW
        case "kVlanModify":
            return PrivilegeId_KVLANMODIFY
        case "kHybridExtenderView":
            return PrivilegeId_KHYBRIDEXTENDERVIEW
        case "kHybridExtenderDownload":
            return PrivilegeId_KHYBRIDEXTENDERDOWNLOAD
        case "kAdLdapView":
            return PrivilegeId_KADLDAPVIEW
        case "kAdLdapModify":
            return PrivilegeId_KADLDAPMODIFY
        case "kSchedulerView":
            return PrivilegeId_KSCHEDULERVIEW
        case "kSchedulerModify":
            return PrivilegeId_KSCHEDULERMODIFY
        case "kProtectionView":
            return PrivilegeId_KPROTECTIONVIEW
        case "kProtectionModify":
            return PrivilegeId_KPROTECTIONMODIFY
        case "kProtectionJobOperate":
            return PrivilegeId_KPROTECTIONJOBOPERATE
        case "kProtectionSourceModify":
            return PrivilegeId_KPROTECTIONSOURCEMODIFY
        case "kProtectionPolicyView":
            return PrivilegeId_KPROTECTIONPOLICYVIEW
        case "kProtectionPolicyModify":
            return PrivilegeId_KPROTECTIONPOLICYMODIFY
        case "kRestoreView":
            return PrivilegeId_KRESTOREVIEW
        case "kRestoreModify":
            return PrivilegeId_KRESTOREMODIFY
        case "kRestoreDownload":
            return PrivilegeId_KRESTOREDOWNLOAD
        case "kRemoteRestore":
            return PrivilegeId_KREMOTERESTORE
        case "kStorageView":
            return PrivilegeId_KSTORAGEVIEW
        case "kStorageModify":
            return PrivilegeId_KSTORAGEMODIFY
        case "kStorageDomainView":
            return PrivilegeId_KSTORAGEDOMAINVIEW
        case "kStorageDomainModify":
            return PrivilegeId_KSTORAGEDOMAINMODIFY
        case "kAnalyticsView":
            return PrivilegeId_KANALYTICSVIEW
        case "kAnalyticsModify":
            return PrivilegeId_KANALYTICSMODIFY
        case "kReportsView":
            return PrivilegeId_KREPORTSVIEW
        case "kMcmModify":
            return PrivilegeId_KMCMMODIFY
        case "kDataSecurity":
            return PrivilegeId_KDATASECURITY
        case "kSmbBackup":
            return PrivilegeId_KSMBBACKUP
        case "kSmbRestore":
            return PrivilegeId_KSMBRESTORE
        case "kSmbTakeOwnership":
            return PrivilegeId_KSMBTAKEOWNERSHIP
        case "kSmbAuditing":
            return PrivilegeId_KSMBAUDITING
        case "kMcmUnregister":
            return PrivilegeId_KMCMUNREGISTER
        case "kMcmUpgrade":
            return PrivilegeId_KMCMUPGRADE
        case "kMcmModifySuperAdmin":
            return PrivilegeId_KMCMMODIFYSUPERADMIN
        case "kMcmViewSuperAdmin":
            return PrivilegeId_KMCMVIEWSUPERADMIN
        case "kMcmModifyCohesityAdmin":
            return PrivilegeId_KMCMMODIFYCOHESITYADMIN
        case "kMcmViewCohesityAdmin":
            return PrivilegeId_KMCMVIEWCOHESITYADMIN
        case "kObjectSearch":
            return PrivilegeId_KOBJECTSEARCH
        case "kFileDatalockExpiryTimeDecrease":
            return PrivilegeId_KFILEDATALOCKEXPIRYTIMEDECREASE
        default:
            return PrivilegeId_KPRINCIPALVIEW
    }
}
