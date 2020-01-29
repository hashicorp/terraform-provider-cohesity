// Copyright 2019 Cohesity Inc.
package models

import "time"

/*
 * Structure for the custom type ADAttributeRestoreParam
 */
type ADAttributeRestoreParam struct {
	ExcludedPropertyVec *[]string                            `json:"excludedPropertyVec,omitempty" form:"excludedPropertyVec,omitempty"` //Array of LDAP property names to excluded from 'property_vec'. Excluded
	GuidpairVec         []*ADGuidPairADAttributeRestoreParam `json:"guidpairVec,omitempty" form:"guidpairVec,omitempty"`                 //Array of source and destination object guid pairs to restore attributes.
	OptionFlags         *int64                               `json:"optionFlags,omitempty" form:"optionFlags,omitempty"`                 //Attribute restore option flags of type ADAttributeOptionFlags.
	PropertyVec         *[]string                            `json:"propertyVec,omitempty" form:"propertyVec,omitempty"`                 //Array of LDAP property(attribute) names. The name can be standard or
}

/*
 * Structure for the custom type ADObject
 */
type ADObject struct {
	Description       *string                 `json:"description,omitempty" form:"description,omitempty"`             //Specifies the 'description' of an AD object.
	DestinationGuid   *string                 `json:"destinationGuid,omitempty" form:"destinationGuid,omitempty"`     //Specifies the guid of object in the Production AD which is equivalent to
	DisplayName       *string                 `json:"displayName,omitempty" form:"displayName,omitempty"`             //Specifies the display name of the AD object.
	DistinguishedName *string                 `json:"distinguishedName,omitempty" form:"distinguishedName,omitempty"` //Specifies the distinguished name of the AD object.
	ErrorMessage      *string                 `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`           //Specifies the error message while fetching the AD object.
	ObjectClass       *string                 `json:"objectClass,omitempty" form:"objectClass,omitempty"`             //Specifies the class name of an AD Object such as 'user','computer',
	SearchResultFlags *[]SearchResultFlagEnum `json:"searchResultFlags,omitempty" form:"searchResultFlags,omitempty"` //Specifies the SearchResultFlags of the AD object.
	SourceGuid        *string                 `json:"sourceGuid,omitempty" form:"sourceGuid,omitempty"`               //Specifies the guid of the AD object in Snapshot AD.
}

/*
 * Structure for the custom type ADObjectRestoreParam
 */
type ADObjectRestoreParam struct {
	Credentials *Credentials `json:"credentials,omitempty" form:"credentials,omitempty"` //Specifies credentials to access a target source.
	GuidVec     *[]string    `json:"guidVec,omitempty" form:"guidVec,omitempty"`         //Array of AD object guids to restore either from recycle bin or from AD
	OptionFlags *int64       `json:"optionFlags,omitempty" form:"optionFlags,omitempty"` //Restore option flags of type ADObjectRestoreOptionFlags.
	OuPath      *string      `json:"ouPath,omitempty" form:"ouPath,omitempty"`           //Distinguished name(DN) of the target Organization Unit (OU) to restore
}

/*
 * Structure for the custom type ADObjectRestoreStatus
 */
type ADObjectRestoreStatus struct {
	DestGuid          *string                                          `json:"destGuid,omitempty" form:"destGuid,omitempty"`                   //Destination guid string of the AD object that is newly created on
	ObjectFlags       *int64                                           `json:"objectFlags,omitempty" form:"objectFlags,omitempty"`             //Object result flags of type ADObjectFlags.
	PropertyStatusVec []*ADObjectRestoreStatusADAttributeRestoreStatus `json:"propertyStatusVec,omitempty" form:"propertyStatusVec,omitempty"` //AD object attribute(property) restore status vector.
	SourceGuid        *string                                          `json:"sourceGuid,omitempty" form:"sourceGuid,omitempty"`               //Source guid of AD object that was restored. This
	Status            *ErrorProto                                      `json:"status,omitempty" form:"status,omitempty"`                       //TODO: Write general description for this field
	TimetakenMs       *int64                                           `json:"timetakenMs,omitempty" form:"timetakenMs,omitempty"`             //Time taken in milliseconds to restore the individual object or attribute
}

/*
 * Structure for the custom type ADObjectRestoreStatusADAttributeRestoreStatus
 */
type ADObjectRestoreStatusADAttributeRestoreStatus struct {
	AttrstatusVec []*ErrorProto `json:"attrstatusVec,omitempty" form:"attrstatusVec,omitempty"` //Error status. If the 'attrstatus_vec' is empty or contains kNoError,
	LdapName      *string       `json:"ldapName,omitempty" form:"ldapName,omitempty"`           //LDAP name of the attribute.
}

/*
 * Structure for the custom type ADRestoreStatus
 */
type ADRestoreStatus struct {
	ObjectInfo *CompareADObjectsResultADObject `json:"objectInfo,omitempty" form:"objectInfo,omitempty"` //TODO: Write general description for this field
	Status     *ADObjectRestoreStatus          `json:"status,omitempty" form:"status,omitempty"`         //TODO: Write general description for this field
}

/*
 * Structure for the custom type ADUpdateRestoreTaskOptions
 */
type ADUpdateRestoreTaskOptions struct {
	ObjectAttributesParam *ADAttributeRestoreParam `json:"objectAttributesParam,omitempty" form:"objectAttributesParam,omitempty"` //TODO: Write general description for this field
	ObjectParam           *ADObjectRestoreParam    `json:"objectParam,omitempty" form:"objectParam,omitempty"`                     //TODO: Write general description for this field
	Type                  *int64                   `json:"type,omitempty" form:"type,omitempty"`                                   //Specifies the AD restore request type.
}

/*
 * Structure for the custom type AWSSnapshotManagerParams
 */
type AWSSnapshotManagerParams struct {
	AmiCreationFrequency *int64 `json:"amiCreationFrequency,omitempty" form:"amiCreationFrequency,omitempty"` //The frequency of AMI creation. This should be set if the option to create
	CreateAmiForRun      *bool  `json:"createAmiForRun,omitempty" form:"createAmiForRun,omitempty"`           //Whether we need to create an AMI for this run.
	ShouldCreateAmi      *bool  `json:"shouldCreateAmi,omitempty" form:"shouldCreateAmi,omitempty"`           //Whether we need to create an AMI after taking snapshots of the instance.
}

/*
 * Structure for the custom type AagAndDatabases
 */
type AagAndDatabases struct {
	Aag       *ProtectionSource   `json:"aag,omitempty" form:"aag,omitempty"`             //Specifies a generic structure that represents a node
	Databases []*ProtectionSource `json:"databases,omitempty" form:"databases,omitempty"` //Specifies databases found that are members of the AAG.
}

/*
 * Structure for the custom type AccessToken
 */
type AccessToken struct {
	AccessToken *string   `json:"accessToken,omitempty" form:"accessToken,omitempty"` //Generated access token.
	Privileges  *[]string `json:"privileges,omitempty" form:"privileges,omitempty"`   //Privileges for the user.
	TokenType   *string   `json:"tokenType,omitempty" form:"tokenType,omitempty"`     //Access token type.
}

/*
 * Structure for the custom type AccessTokenCredential
 */
type AccessTokenCredential struct {
	Domain   *string `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain the user is logging in to. For a Local user model,
	Password string  `json:"password" form:"password"`                 //Specifies the password of the Cohesity user account.
	Username string  `json:"username" form:"username"`                 //Specifies the login name of the Cohesity user.
}

/*
 * Structure for the custom type AcropolisProtectionSource
 */
type AcropolisProtectionSource struct {
	ClusterUuid *string  `json:"clusterUuid,omitempty" form:"clusterUuid,omitempty"` //Specifies the UUID of the Acropolis cluster instance to which this
	Description *string  `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the Protection Source.
	MountPath   *bool    `json:"mountPath,omitempty" form:"mountPath,omitempty"`     //Specifies whether the the VM is an agent VM. This is applicable to
	Name        *string  `json:"name,omitempty" form:"name,omitempty"`               //Specifies the name of the Acropolis Object.
	Type        TypeEnum `json:"type,omitempty" form:"type,omitempty"`               //Specifies the type of an Acropolis Protection Source Object such as
	Uuid        *string  `json:"uuid,omitempty" form:"uuid,omitempty"`               //Specifies the UUID of the Acropolis Object. This is unique within the
}

/*
 * Structure for the custom type AcropolisRestoreParameters
 */
type AcropolisRestoreParameters struct {
	DisableNetwork     *bool   `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"`         //Specifies whether the network should be left in disabled state.
	NetworkId          *int64  `json:"networkId,omitempty" form:"networkId,omitempty"`                   //Specifies a network configuration to be attached to the cloned or
	PoweredOn          *bool   `json:"poweredOn,omitempty" form:"poweredOn,omitempty"`                   //Specifies the power state of the cloned or recovered objects.
	Prefix             *string `json:"prefix,omitempty" form:"prefix,omitempty"`                         //Specifies a prefix to prepended to the source object name to derive a
	StorageContainerId *int64  `json:"storageContainerId,omitempty" form:"storageContainerId,omitempty"` //A storage container where the VM's files should be restored to. This
	Suffix             *string `json:"suffix,omitempty" form:"suffix,omitempty"`                         //Specifies a suffix to appended to the original source object name
}

/*
 * Structure for the custom type ActivateViewAliasesResult
 */
type ActivateViewAliasesResult struct {
	Aliases []*ViewAliasInfo `json:"aliases,omitempty" form:"aliases,omitempty"` //Aliases created for the view. A view alias allows a directory path inside
}

/*
 * Structure for the custom type ActiveAlertsStats
 */
type ActiveAlertsStats struct {
	NumCriticalAlerts         *int64 `json:"numCriticalAlerts,omitempty" form:"numCriticalAlerts,omitempty"`                 //Specifies the count of active critical Alerts.
	NumHardwareAlerts         *int64 `json:"numHardwareAlerts,omitempty" form:"numHardwareAlerts,omitempty"`                 //Specifies the count of active hardware Alerts.
	NumHardwareCriticalAlerts *int64 `json:"numHardwareCriticalAlerts,omitempty" form:"numHardwareCriticalAlerts,omitempty"` //Specifies the count of active hardware critical Alerts.
	NumHardwareInfoAlerts     *int64 `json:"numHardwareInfoAlerts,omitempty" form:"numHardwareInfoAlerts,omitempty"`         //Specifies the count of active hardware info Alerts.
	NumHardwareWarningAlerts  *int64 `json:"numHardwareWarningAlerts,omitempty" form:"numHardwareWarningAlerts,omitempty"`   //Specifies the count of active hardware warning Alerts.
	NumInfoAlerts             *int64 `json:"numInfoAlerts,omitempty" form:"numInfoAlerts,omitempty"`                         //Specifies the count of active info Alerts.
	NumServiceAlerts          *int64 `json:"numServiceAlerts,omitempty" form:"numServiceAlerts,omitempty"`                   //Specifies the count of active service Alerts.
	NumServiceCriticalAlerts  *int64 `json:"numServiceCriticalAlerts,omitempty" form:"numServiceCriticalAlerts,omitempty"`   //Specifies the count of active service critical Alerts.
	NumServiceInfoAlerts      *int64 `json:"numServiceInfoAlerts,omitempty" form:"numServiceInfoAlerts,omitempty"`           //Specifies the count of active service info Alerts.
	NumServiceWarningAlerts   *int64 `json:"numServiceWarningAlerts,omitempty" form:"numServiceWarningAlerts,omitempty"`     //Specifies the count of active service warning Alerts.
	NumSoftwareAlerts         *int64 `json:"numSoftwareAlerts,omitempty" form:"numSoftwareAlerts,omitempty"`                 //Specifies the count of active software Alerts.
	NumSoftwareCriticalAlerts *int64 `json:"numSoftwareCriticalAlerts,omitempty" form:"numSoftwareCriticalAlerts,omitempty"` //Specifies the count of active software critical Alerts.
	NumSoftwareInfoAlerts     *int64 `json:"numSoftwareInfoAlerts,omitempty" form:"numSoftwareInfoAlerts,omitempty"`         //Specifies the count of active software info Alerts.
	NumSoftwareWarningAlerts  *int64 `json:"numSoftwareWarningAlerts,omitempty" form:"numSoftwareWarningAlerts,omitempty"`   //Specifies the count of active software warning Alerts.
	NumWarningAlerts          *int64 `json:"numWarningAlerts,omitempty" form:"numWarningAlerts,omitempty"`                   //Specifies the count of active warning Alerts.
}

/*
 * Structure for the custom type ActiveDirectoryEntry
 */
type ActiveDirectoryEntry struct {
	DomainName                 *string                      `json:"domainName,omitempty" form:"domainName,omitempty"`                                 //Specifies the fully qualified domain name (FQDN) of an Active Directory.
	FallbackUserIdMappingInfo  *UserIdMapping               `json:"fallbackUserIdMappingInfo,omitempty" form:"fallbackUserIdMappingInfo,omitempty"`   //Specifies how the Unix and Windows users are mapped in an Active Directory.
	IgnoredTrustedDomains      *[]string                    `json:"ignoredTrustedDomains,omitempty" form:"ignoredTrustedDomains,omitempty"`           //Specifies the list of trusted domains that were set by the user to be
	LdapProviderId             *int64                       `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"`                         //Specifies the LDAP provider id which is map to this Active Directory
	MachineAccounts            *[]string                    `json:"machineAccounts,omitempty" form:"machineAccounts,omitempty"`                       //Array of Machine Accounts.
	OuName                     *string                      `json:"ouName,omitempty" form:"ouName,omitempty"`                                         //Specifies an optional Organizational Unit name.
	Password                   *string                      `json:"password,omitempty" form:"password,omitempty"`                                     //Specifies the password for the specified userName.
	PreferredDomainControllers []*PreferredDomainController `json:"preferredDomainControllers,omitempty" form:"preferredDomainControllers,omitempty"` //Specifies Map of Active Directory domain names to its preferred domain
	TenantId                   *string                      `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                     //Specifies the unique id of the tenant.
	TrustedDomains             *[]string                    `json:"trustedDomains,omitempty" form:"trustedDomains,omitempty"`                         //Specifies the trusted domains of the Active Directory domain.
	TrustedDomainsEnabled      *bool                        `json:"trustedDomainsEnabled,omitempty" form:"trustedDomainsEnabled,omitempty"`           //Specifies whether Trusted Domain discovery is disabled.
	UnixRootSid                *string                      `json:"unixRootSid,omitempty" form:"unixRootSid,omitempty"`                               //Specifies the SID of the Active Directory domain user to be mapped to
	UserIdMappingInfo          *UserIdMapping               `json:"userIdMappingInfo,omitempty" form:"userIdMappingInfo,omitempty"`                   //Specifies how the Unix and Windows users are mapped in an Active Directory.
	UserName                   *string                      `json:"userName,omitempty" form:"userName,omitempty"`                                     //Specifies a userName that has administrative privileges in the domain.
	Workgroup                  *string                      `json:"workgroup,omitempty" form:"workgroup,omitempty"`                                   //Specifies an optional Workgroup name.
}

/*
 * Structure for the custom type ActiveDirectoryPrincipal
 */
type ActiveDirectoryPrincipal struct {
	Domain        *string         `json:"domain,omitempty" form:"domain,omitempty"`               //Specifies the domain name of the where the principal' account is
	FullName      *string         `json:"fullName,omitempty" form:"fullName,omitempty"`           //Specifies the full name (first and last names) of the principal.
	ObjectClass   ObjectClassEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"`     //Specifies the object class of the principal (either 'kGroup' or 'kUser').
	PrincipalName *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the principal.
	Sid           *string         `json:"sid,omitempty" form:"sid,omitempty"`                     //Specifies the unique Security id (SID) of the principal.
}

/*
 * Structure for the custom type ActiveDirectoryPrincipalsAddParameters
 */
type ActiveDirectoryPrincipalsAddParameters struct {
	Description   *string                                               `json:"description,omitempty" form:"description,omitempty"`     //Specifies a description about the user or group.
	Domain        *string                                               `json:"domain,omitempty" form:"domain,omitempty"`               //Specifies the domain of the Active Directory where the
	ObjectClass   ObjectClassActiveDirectoryPrincipalsAddParametersEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"`     //Specifies the type of the referenced Active Directory principal.
	PrincipalName *string                                               `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the Active Directory principal,
	Restricted    *bool                                                 `json:"restricted,omitempty" form:"restricted,omitempty"`       //Whether the principal is a restricted principal. A restricted principal
	Roles         *[]string                                             `json:"roles,omitempty" form:"roles,omitempty"`                 //Array of Roles.
}

/*
 * Structure for the custom type AdAttribute
 */
type AdAttribute struct {
	AdAttributeFlags *[]AdAttributeFlagEnum `json:"adAttributeFlags,omitempty" form:"adAttributeFlags,omitempty"` //Specifies the flags related to the attribute of the AD object.
	DestinationValue *AttributeValue        `json:"destinationValue,omitempty" form:"destinationValue,omitempty"` //Represents the information about the values of attribute of the ADObject.
	ErrorMessage     *string                `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`         //Specifies the error message regarding the attribute
	Name             *string                `json:"name,omitempty" form:"name,omitempty"`                         //Specifies the name of the attribute of the AD object.
	SameValue        *AttributeValue        `json:"sameValue,omitempty" form:"sameValue,omitempty"`               //Represents the information about the values of attribute of the ADObject.
	SourceValue      *AttributeValue        `json:"sourceValue,omitempty" form:"sourceValue,omitempty"`           //Represents the information about the values of attribute of the ADObject.
}

/*
 * Structure for the custom type AdDomain
 */
type AdDomain struct {
	DnsRoot       *string           `json:"dnsRoot,omitempty" form:"dnsRoot,omitempty"`             //Specifies DNS root.
	Forest        *string           `json:"forest,omitempty" form:"forest,omitempty"`               //Specifies AD forest name.
	Identity      *AdDomainIdentity `json:"identity,omitempty" form:"identity,omitempty"`           //AD domain identity information.
	NetbiosName   *string           `json:"netbiosName,omitempty" form:"netbiosName,omitempty"`     //Specifies AD NetBIOS name.
	ParentDomain  *string           `json:"parentDomain,omitempty" form:"parentDomain,omitempty"`   //Specifies parent domain name.
	TombstoneDays *int64            `json:"tombstoneDays,omitempty" form:"tombstoneDays,omitempty"` //Specifies tombstone time in days.
}

/*
 * Structure for the custom type AdDomainController
 */
type AdDomainController struct {
	BackupSupported          *bool     `json:"backupSupported,omitempty" form:"backupSupported,omitempty"`                   //Specifies whether backup of this domain controller is supported.
	BackupUnsupportedReasons *[]string `json:"backupUnsupportedReasons,omitempty" form:"backupUnsupportedReasons,omitempty"` //Specifies any reason(s) for domain controller backup not supported.
	Domain                   *AdDomain `json:"domain,omitempty" form:"domain,omitempty"`                                     //Specifies information about an AD Domain.
	HostName                 *string   `json:"hostName,omitempty" form:"hostName,omitempty"`                                 //Specifies FQDN host name of the domain controller.
	IsGlobalCatalog          *bool     `json:"isGlobalCatalog,omitempty" form:"isGlobalCatalog,omitempty"`                   //Specifies whether this domain controller is a global catalog server.
	IsReadOnly               *bool     `json:"isReadOnly,omitempty" form:"isReadOnly,omitempty"`                             //Specifies whether this domain controller is read only.
	UtcOffsetMin             *int64    `json:"utcOffsetMin,omitempty" form:"utcOffsetMin,omitempty"`                         //Specifies UTC time offset of this domain controller in minutes.
}

/*
 * Structure for the custom type AdDomainIdentity
 */
type AdDomainIdentity struct {
	Dn   *string `json:"dn,omitempty" form:"dn,omitempty"`     //Specifies distinguished name of the domain.
	Guid *string `json:"guid,omitempty" form:"guid,omitempty"` //Specifies Unique objectGUID for an AD domain.
	Name *string `json:"name,omitempty" form:"name,omitempty"` //Specifies display name of the domain.
	Sid  *string `json:"sid,omitempty" form:"sid,omitempty"`   //Specifies domain SID.
}

/*
 * Structure for the custom type AdGuidPair
 */
type AdGuidPair struct {
	Destination *string `json:"destination,omitempty" form:"destination,omitempty"` //Specifies the destination guid in production AD object corresponding to
	Source      *string `json:"source,omitempty" form:"source,omitempty"`           //Specifies the source guid string of an AD object in mounted AD snapshot.
}

/*
 * Structure for the custom type AdObjectAttributeParameters
 */
type AdObjectAttributeParameters struct {
	AdGuidPairs             []*AdGuidPair `json:"adGuidPairs,omitempty" form:"adGuidPairs,omitempty"`                         //Specifies the array of source and destination object guid pairs to
	ExcludeLdapProperties   *[]string     `json:"excludeLdapProperties,omitempty" form:"excludeLdapProperties,omitempty"`     //Specifies the array of LDAP property names to excluded from
	LdapProperties          *[]string     `json:"ldapProperties,omitempty" form:"ldapProperties,omitempty"`                   //Specifies the array of LDAP property(attribute) names. The name can be
	MergeMultiValProperties *bool         `json:"mergeMultiValProperties,omitempty" form:"mergeMultiValProperties,omitempty"` //Specifies the Option to merge multi-valued values vs clearing and setting
}

/*
 * Structure for the custom type AdObjectMetaData
 */
type AdObjectMetaData struct {
	Guid           *string `json:"guid,omitempty" form:"guid,omitempty"`                     //Specifies the Guid of the AD object.
	Name           *string `json:"name,omitempty" form:"name,omitempty"`                     //Specifies the name of the AD object.
	SamAccountName *string `json:"samAccountName,omitempty" form:"samAccountName,omitempty"` //Specifies the sam account name of the AD object.
}

/*
 * Structure for the custom type AdObjectRestoreInformation
 */
type AdObjectRestoreInformation struct {
	AttributeRestoreInfo []*AttributeRestoreInformation `json:"attributeRestoreInfo,omitempty" form:"attributeRestoreInfo,omitempty"` //Specifies the list of attributes of the AD object whose restore failed.
	ErrorMessage         *string                        `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`                 //Specifies the error message while restoring the AD Object.
	Name                 *string                        `json:"name,omitempty" form:"name,omitempty"`                                 //Specifies the name of the AD object.
	StartTimeUsecs       *int64                         `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`             //Specifies the start time of the restore of the AD object specified as a
	TimeTakenMsecs       *int64                         `json:"timeTakenMsecs,omitempty" form:"timeTakenMsecs,omitempty"`             //Specifies the time taken for restore of AD Object and its attributes in
}

/*
 * Structure for the custom type AdObjectRestoreParameters
 */
type AdObjectRestoreParameters struct {
	ChangePasswordOnNextLogon *bool     `json:"changePasswordOnNextLogon,omitempty" form:"changePasswordOnNextLogon,omitempty"` //Specifies the option for AD 'user' type of objects to change password when
	LeaveStateDisabled        *bool     `json:"leaveStateDisabled,omitempty" form:"leaveStateDisabled,omitempty"`               //Specifies the option to leave the restored object in disabled state for
	ObjectGuids               *[]string `json:"objectGuids,omitempty" form:"objectGuids,omitempty"`                             //Specifies the array of AD object guids to restore either from recycle bin
	OrganizationUnitPath      *string   `json:"organizationUnitPath,omitempty" form:"organizationUnitPath,omitempty"`           //Specifies the Distinguished name(DN) of the target Organization Unit (OU)
	Password                  *string   `json:"password,omitempty" form:"password,omitempty"`                                   //Specifies the password for restoring user type objects (user,
}

/*
 * Structure for the custom type AdObjectsRestoreStatus
 */
type AdObjectsRestoreStatus struct {
	AdObjectsRestoreInfo []*AdObjectRestoreInformation `json:"adObjectsRestoreInfo,omitempty" form:"adObjectsRestoreInfo,omitempty"` //Specifies the status of all the AD Objects which were requested to
	NumObjectsFailed     *int64                        `json:"numObjectsFailed,omitempty" form:"numObjectsFailed,omitempty"`         //Specifies the number of AD Objects whose restore is in progress.
	NumObjectsSucceeded  *int64                        `json:"numObjectsSucceeded,omitempty" form:"numObjectsSucceeded,omitempty"`   //Specifies the number of AD Objects whose restore is successfull.
}

/*
 * Structure for the custom type AdProtectionSource
 */
type AdProtectionSource struct {
	DomainController *AdDomainController `json:"domainController,omitempty" form:"domainController,omitempty"` //Specifies information about an AD domain controller.
	Name             *string             `json:"name,omitempty" form:"name,omitempty"`                         //Specifies the domain name of the AD entity.
	OwnerId          *int64              `json:"ownerId,omitempty" form:"ownerId,omitempty"`                   //Specifies the entity id of the owner entity.
	Type             *int64              `json:"type,omitempty" form:"type,omitempty"`                         //Specifies the type of the managed object in AD Protection Source.
	Uuid             *string             `json:"uuid,omitempty" form:"uuid,omitempty"`                         //Specifies the UUID for the AD entity.
}

/*
 * Structure for the custom type AdRestoreOptions
 */
type AdRestoreOptions struct {
	ObjectAttributeParameters *AdObjectAttributeParameters `json:"objectAttributeParameters,omitempty" form:"objectAttributeParameters,omitempty"` //AdObjectAttributeParameters are AD attribute recovery parameters for one
	ObjectParameters          *AdObjectRestoreParameters   `json:"objectParameters,omitempty" form:"objectParameters,omitempty"`                   //AdObjectRestoreParameters are the parameters to restore AD objects from
	Type                      TypeAdRestoreOptionsEnum     `json:"type,omitempty" form:"type,omitempty"`                                           //Specifies the AD restore request type.
}

/*
 * Structure for the custom type AdRestoreParameters
 */
type AdRestoreParameters struct {
	Credentials *Credentials `json:"credentials,omitempty" form:"credentials,omitempty"` //Specifies credentials to access a target source.
	Port        *int64       `json:"port,omitempty" form:"port,omitempty"`               //Specifies the port on which the AD domain controller's NTDS database will
}

/*
 * Structure for the custom type AdRootTopologyObject
 */
type AdRootTopologyObject struct {
	ChildObjects      *[]interface{} `json:"childObjects,omitempty" form:"childObjects,omitempty"`           //Specifies the array of children of this object.
	Description       *string        `json:"description,omitempty" form:"description,omitempty"`             //Specifies the 'description' of an object.
	DestGuid          *string        `json:"destGuid,omitempty" form:"destGuid,omitempty"`                   //Specifies the guid of matching 'source_guid' from production AD.
	DisplayName       *string        `json:"displayName,omitempty" form:"displayName,omitempty"`             //Specifies the display name of the object in AD Topology tree.
	DistinguishedName *string        `json:"distinguishedName,omitempty" form:"distinguishedName,omitempty"` //Specifies the distinguished name of the object in AD Topology tree.
	ErrorMessage      *string        `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`           //Specifies the AD error while fetching the ADRootTopologyObject.
	ObjectClass       *string        `json:"objectClass,omitempty" form:"objectClass,omitempty"`             //Specifies the LDAP class name such as 'user','computer',
	SourceGuid        *string        `json:"sourceGuid,omitempty" form:"sourceGuid,omitempty"`               //Specifies the guid string of the object in AD snapshot database.
}

/*
 * Structure for the custom type AddedActiveDirectoryPrincipal
 */
type AddedActiveDirectoryPrincipal struct {
	CreatedTimeMsecs     *int64                                       `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the group or user
	Description          *string                                      `json:"description,omitempty" form:"description,omitempty"`                   //Specifies a description about the user or group.
	Domain               *string                                      `json:"domain,omitempty" form:"domain,omitempty"`                             //Specifies the domain of the Active Directory where the
	LastUpdatedTimeMsecs *int64                                       `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group or user
	ObjectClass          ObjectClassAddedActiveDirectoryPrincipalEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"`                   //Specifies the type of the referenced Active Directory principal.
	PrincipalName        *string                                      `json:"principalName,omitempty" form:"principalName,omitempty"`               //Specifies the name of the Active Directory principal,
	Restricted           *bool                                        `json:"restricted,omitempty" form:"restricted,omitempty"`                     //Whether the principal is a restricted principal. A restricted principal
	Roles                *[]string                                    `json:"roles,omitempty" form:"roles,omitempty"`                               //Array of Roles.
	Sid                  *string                                      `json:"sid,omitempty" form:"sid,omitempty"`                                   //Specifies the unique Security ID (SID) of the Active Directory principal
}

/*
 * Structure for the custom type AddedIdpPrincipal
 */
type AddedIdpPrincipal struct {
	CreatedTimeMsecs     *int64                           `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the group or user
	Domain               *string                          `json:"domain,omitempty" form:"domain,omitempty"`                             //Specifies the name of the Idp where the
	LastUpdatedTimeMsecs *int64                           `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group or user
	ObjectClass          ObjectClassAddedIdpPrincipalEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"`                   //Specifies the type of the referenced Idp principal.
	PrincipalName        *string                          `json:"principalName,omitempty" form:"principalName,omitempty"`               //Specifies the name of the Idp principal,
	Restricted           *bool                            `json:"restricted,omitempty" form:"restricted,omitempty"`                     //Whether the principal is a restricted principal. A restricted principal
	Roles                *[]string                        `json:"roles,omitempty" form:"roles,omitempty"`                               //Array of Roles.
	Sid                  *string                          `json:"sid,omitempty" form:"sid,omitempty"`                                   //Specifies the unique Security ID (SID) of the Idp principal
}

/*
 * Structure for the custom type AdditionalOracleDBParams
 */
type AdditionalOracleDBParams struct {
	AppEntityId      *int64                 `json:"appEntityId,omitempty" form:"appEntityId,omitempty"`           //Database app id.
	DbInfoChannelVec []*OracleDBChannelInfo `json:"dbInfoChannelVec,omitempty" form:"dbInfoChannelVec,omitempty"` //Contains the information for each database and the corresponding hosts
}

/*
 * Structure for the custom type AgentDeploymentStatusResponse
 */
type AgentDeploymentStatusResponse struct {
	CompactVersion       *string                                        `json:"compactVersion,omitempty" form:"compactVersion,omitempty"`             //Specifies the compact version of Cohesity agent. For example, 6.0.1.
	HealthStatus         HealthStatusEnum                               `json:"healthStatus,omitempty" form:"healthStatus,omitempty"`                 //Specifies the health status of the Cohesity agent.
	HostIp               *string                                        `json:"hostIp,omitempty" form:"hostIp,omitempty"`                             //Specifies the IP of the host on which the agent is installed.
	HostOsType           HostOsTypeEnum                                 `json:"hostOsType,omitempty" form:"hostOsType,omitempty"`                     //Specifies the host type on which the agent is installed.
	LastUpgradeStatus    LastUpgradeStatusEnum                          `json:"lastUpgradeStatus,omitempty" form:"lastUpgradeStatus,omitempty"`       //Specifies the status of the last upgrade attempt.
	Upgradability        UpgradabilityAgentDeploymentStatusResponseEnum `json:"upgradability,omitempty" form:"upgradability,omitempty"`               //Specfies the upgradability of the agent running on the server.
	UpgradeStatusMessage *string                                        `json:"upgradeStatusMessage,omitempty" form:"upgradeStatusMessage,omitempty"` //Specifies detailed message about the agent upgrade failure. This field
	Version              *string                                        `json:"version,omitempty" form:"version,omitempty"`                           //Specifies the Cohesity software version of the agent. For example:
}

/*
 * Structure for the custom type AgentInformation
 */
type AgentInformation struct {
	CbmrVersion                     *string                      `json:"cbmrVersion,omitempty" form:"cbmrVersion,omitempty"`                                         //Specifies the version if Cristie BMR product is installed on the host.
	HostType                        HostTypeAgentInformationEnum `json:"hostType,omitempty" form:"hostType,omitempty"`                                               //Specifies the host type where the agent is running. This is only set for
	Id                              *int64                       `json:"id,omitempty" form:"id,omitempty"`                                                           //Specifies the agent's id.
	Name                            *string                      `json:"name,omitempty" form:"name,omitempty"`                                                       //Specifies the agent's name.
	OracleMultiNodeChannelSupported *bool                        `json:"oracleMultiNodeChannelSupported,omitempty" form:"oracleMultiNodeChannelSupported,omitempty"` //Specifies whether oracle multi node multi channel is supported or not.
	RegistrationInfo                *RegisteredSourceInfo        `json:"registrationInfo,omitempty" form:"registrationInfo,omitempty"`                               //Specifies information about a registered Source.
	SourceSideDedupEnabled          *bool                        `json:"sourceSideDedupEnabled,omitempty" form:"sourceSideDedupEnabled,omitempty"`                   //Specifies whether source side dedup is enabled or not.
	Status                          StatusEnum                   `json:"status,omitempty" form:"status,omitempty"`                                                   //Specifies the agent status.
	StatusMessage                   *string                      `json:"statusMessage,omitempty" form:"statusMessage,omitempty"`                                     //Specifies additional details about the agent status.
	Upgradability                   UpgradabilityEnum            `json:"upgradability,omitempty" form:"upgradability,omitempty"`                                     //Specifies the upgradability of the agent running on the physical server.
	UpgradeStatus                   UpgradeStatusEnum            `json:"upgradeStatus,omitempty" form:"upgradeStatus,omitempty"`                                     //Specifies the status of the upgrade of the agent on a physical server.
	UpgradeStatusMessage            *string                      `json:"upgradeStatusMessage,omitempty" form:"upgradeStatusMessage,omitempty"`                       //Specifies detailed message about the agent upgrade failure. This field
	Version                         *string                      `json:"version,omitempty" form:"version,omitempty"`                                                 //Specifies the version of the Agent software.
}

/*
 * Structure for the custom type AggregatedSubtreeInfo
 */
type AggregatedSubtreeInfo struct {
	Environment      EnvironmentAggregatedSubtreeInfoEnum `json:"environment,omitempty" form:"environment,omitempty"`           //Specifies the environment such as 'kSQL' or 'kVMware', where the
	LeavesCount      *int64                               `json:"leavesCount,omitempty" form:"leavesCount,omitempty"`           //Specifies the number of leaf nodes under the subtree of this node.
	TotalLogicalSize *int64                               `json:"totalLogicalSize,omitempty" form:"totalLogicalSize,omitempty"` //Specifies the total logical size of the data under the subtree
}

/*
 * Structure for the custom type Alert
 */
type Alert struct {
	AlertCategory        AlertCategoryEnum       `json:"alertCategory,omitempty" form:"alertCategory,omitempty"`               //Specifies the category of an Alert.
	AlertCode            *string                 `json:"alertCode,omitempty" form:"alertCode,omitempty"`                       //Specifies a unique code that categorizes the Alert,
	AlertDocument        *AlertDocument          `json:"alertDocument,omitempty" form:"alertDocument,omitempty"`               //Specifies documentation about the Alert such as name, description, cause
	AlertState           AlertStateEnum          `json:"alertState,omitempty" form:"alertState,omitempty"`                     //Specifies the current state of the Alert.
	AlertType            *int64                  `json:"alertType,omitempty" form:"alertType,omitempty"`                       //Specifies a 5 digit unique digital id for the Alert Type, such as 00014
	AlertTypeBucket      AlertTypeBucketEnum     `json:"alertTypeBucket,omitempty" form:"alertTypeBucket,omitempty"`           //Specifies the Alert type bucket.
	ClusterId            *int64                  `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //Specifies id of the cluster where the alert was raised.
	ClusterName          *string                 `json:"clusterName,omitempty" form:"clusterName,omitempty"`                   //Specifies name of the cluster where the alert was raised.
	DedupCount           *int64                  `json:"dedupCount,omitempty" form:"dedupCount,omitempty"`                     //Specifies total count of duplicated Alerts even if there are more than
	DedupTimestamps      *[]int64                `json:"dedupTimestamps,omitempty" form:"dedupTimestamps,omitempty"`           //Specifies Unix epoch Timestamps (in microseconds) for the last 25
	EventSource          *string                 `json:"eventSource,omitempty" form:"eventSource,omitempty"`                   //Specifies source where the event occurred.
	FirstTimestampUsecs  *int64                  `json:"firstTimestampUsecs,omitempty" form:"firstTimestampUsecs,omitempty"`   //Specifies Unix epoch Timestamp (in microseconds) of the first
	Id                   *string                 `json:"id,omitempty" form:"id,omitempty"`                                     //Specifies unique id of this Alert.
	LatestTimestampUsecs *int64                  `json:"latestTimestampUsecs,omitempty" form:"latestTimestampUsecs,omitempty"` //Specifies Unix epoch Timestamp (in microseconds) of the most
	PropertyList         []*AlertProperty        `json:"propertyList,omitempty" form:"propertyList,omitempty"`                 //Specifies array of key-value pairs associated with the Alert.
	ResolutionDetails    *AlertResolutionDetails `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"`       //Specifies information about the Alert Resolution such as a summary,
	Severity             SeverityEnum            `json:"severity,omitempty" form:"severity,omitempty"`                         //Specifies the severity level of an Alert.
	SuppressionId        *int64                  `json:"suppressionId,omitempty" form:"suppressionId,omitempty"`               //Specifies unique id generated when the Alert is suppressed by the admin.
	TenantIds            *[]string               `json:"tenantIds,omitempty" form:"tenantIds,omitempty"`                       //Specifies the tenants for which this alert has been raised.
}

/*
 * Structure for the custom type AlertCategoryName
 */
type AlertCategoryName struct {
	Category CategoryEnum `json:"category,omitempty" form:"category,omitempty"` //Specifies alert category.
	Name     *string      `json:"name,omitempty" form:"name,omitempty"`         //Specifies public facing string for alert enums.
}

/*
 * Structure for the custom type AlertDocument
 */
type AlertDocument struct {
	AlertCause       *string `json:"alertCause,omitempty" form:"alertCause,omitempty"`             //Specifies cause of the Alert that is included in the body of the email
	AlertDescription *string `json:"alertDescription,omitempty" form:"alertDescription,omitempty"` //Specifies brief description about the Alert that is used in the subject
	AlertHelpText    *string `json:"alertHelpText,omitempty" form:"alertHelpText,omitempty"`       //Specifies instructions describing how to resolve the Alert that is
	AlertName        *string `json:"alertName,omitempty" form:"alertName,omitempty"`               //Specifies short name that describes the Alert type such as DiskBad,
}

/*
 * Structure for the custom type AlertMetadata
 */
type AlertMetadata struct {
	AlertDocumentList          []*AlertDocument          `json:"alertDocumentList,omitempty" form:"alertDocumentList,omitempty"`                   //Specifies alert documentation one per each language supported.
	AlertTypeBucket            AlertTypeBucketEnum       `json:"alertTypeBucket,omitempty" form:"alertTypeBucket,omitempty"`                       //Specifies the Alert type bucket.
	AlertTypeId                *int64                    `json:"alertTypeId,omitempty" form:"alertTypeId,omitempty"`                               //Specifies unique id for the alert type.
	Category                   CategoryAlertMetadataEnum `json:"category,omitempty" form:"category,omitempty"`                                     //Specifies category of the alert type.
	DedupIntervalSeconds       *int64                    `json:"dedupIntervalSeconds,omitempty" form:"dedupIntervalSeconds,omitempty"`             //Specifies dedup interval in seconds. If the same alert is raised multiple
	DedupUntilResolved         *bool                     `json:"dedupUntilResolved,omitempty" form:"dedupUntilResolved,omitempty"`                 //Specifies if the alerts are to be deduped until the current one (if
	HideAlertFromUser          *bool                     `json:"hideAlertFromUser,omitempty" form:"hideAlertFromUser,omitempty"`                   //Specifies whether to show the alert in the iris UI and CLI.
	IgnoreDuplicateOccurrences *bool                     `json:"ignoreDuplicateOccurrences,omitempty" form:"ignoreDuplicateOccurrences,omitempty"` //Specifies whether to ignore duplicate occurrences completely.
	PrimaryKeyList             *[]string                 `json:"primaryKeyList,omitempty" form:"primaryKeyList,omitempty"`                         //Specifies properties that serve as primary keys.
	PropertyList               *[]string                 `json:"propertyList,omitempty" form:"propertyList,omitempty"`                             //Specifies list of properties that the client is supposed to provide when
	SendSupportNotification    *bool                     `json:"sendSupportNotification,omitempty" form:"sendSupportNotification,omitempty"`       //Specifies whether to send support notification for the alert.
	SnmpNotification           *bool                     `json:"snmpNotification,omitempty" form:"snmpNotification,omitempty"`                     //Specifies whether an SNMP notification is sent when an alert is raised.
	Version                    *int64                    `json:"version,omitempty" form:"version,omitempty"`                                       //Specifies version of the metadata.
}

/*
 * Structure for the custom type AlertProperty
 */
type AlertProperty struct {
	Key   *string `json:"key,omitempty" form:"key,omitempty"`     //Specifies name of the property.
	Value *string `json:"value,omitempty" form:"value,omitempty"` //Specifies value of the property.
}

/*
 * Structure for the custom type AlertResolution
 */
type AlertResolution struct {
	AlertIdList       *[]string               `json:"alertIdList,omitempty" form:"alertIdList,omitempty"`             //Specifies list of Alerts resolved by a Resolution, which are specified by
	ResolutionDetails *AlertResolutionDetails `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies information about the Alert Resolution such as a summary,
	TenantIds         *[]string               `json:"tenantIds,omitempty" form:"tenantIds,omitempty"`                 //Specifies unique tenantIds of the alert contained in this resolution.
}

/*
 * Structure for the custom type AlertResolutionDetails
 */
type AlertResolutionDetails struct {
	ResolutionDetails *string `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies detailed notes about the Resolution.
	ResolutionId      *int64  `json:"resolutionId,omitempty" form:"resolutionId,omitempty"`           //Specifies Unique id assigned by the Cohesity Cluster for this Resolution.
	ResolutionSummary *string `json:"resolutionSummary,omitempty" form:"resolutionSummary,omitempty"` //Specifies short description about the Resolution.
	TimestampUsecs    *int64  `json:"timestampUsecs,omitempty" form:"timestampUsecs,omitempty"`       //Specifies unix epoch timestamp (in microseconds) when the Alerts were
	UserName          *string `json:"userName,omitempty" form:"userName,omitempty"`                   //Specifies name of the Cohesity Cluster user who resolved the Alerts.
}

/*
 * Structure for the custom type AlertResolutionInfo
 */
type AlertResolutionInfo struct {
	ResolutionDetails *string `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies detailed notes about the Resolution.
	ResolutionSummary *string `json:"resolutionSummary,omitempty" form:"resolutionSummary,omitempty"` //Specifies short description about the Resolution.
}

/*
 * Structure for the custom type AlertResolutionRequest
 */
type AlertResolutionRequest struct {
	AlertIdList       *[]string            `json:"alertIdList,omitempty" form:"alertIdList,omitempty"`             //Specifies list of alerts resolved by a Resolution, which are specified by
	ResolutionDetails *AlertResolutionInfo `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Short description and detailed notes about the Resolution.
}

/*
 * Structure for the custom type AlertingConfig
 */
type AlertingConfig struct {
	EmailDeliveryTargets         []*EmailDeliveryTarget `json:"emailDeliveryTargets,omitempty" form:"emailDeliveryTargets,omitempty"`                 //Specifies additional email addresses where alert notifications (configured
	RaiseObjectLevelFailureAlert *bool                  `json:"raiseObjectLevelFailureAlert,omitempty" form:"raiseObjectLevelFailureAlert,omitempty"` //Specifies the boolean to raise per object alert for failures.
}

/*
 * Structure for the custom type AlertingPolicyProto
 */
type AlertingPolicyProto struct {
	DeliveryTargetVec            []*DeliveryRuleProtoDeliveryTarget `json:"deliveryTargetVec,omitempty" form:"deliveryTargetVec,omitempty"`                       //The delivery targets to be alerted.
	Emails                       *[]string                          `json:"emails,omitempty" form:"emails,omitempty"`                                             //The email addresses to send alerts to.
	Policy                       *int64                             `json:"policy,omitempty" form:"policy,omitempty"`                                             //'policy' is declared as int32 because ORing the enums will generate values
	RaiseObjectLevelFailureAlert *bool                              `json:"raiseObjectLevelFailureAlert,omitempty" form:"raiseObjectLevelFailureAlert,omitempty"` //Raise per object alert for failures.
}

/*
 * Structure for the custom type AliasSmbConfig
 */
type AliasSmbConfig struct {
	CachingEnabled              *bool            `json:"cachingEnabled,omitempty" form:"cachingEnabled,omitempty"`                           //Indicate if offline file caching is supported
	DiscoveryEnabled            *bool            `json:"discoveryEnabled,omitempty" form:"discoveryEnabled,omitempty"`                       //Whether the share is discoverable.
	EncryptionEnabled           *bool            `json:"encryptionEnabled,omitempty" form:"encryptionEnabled,omitempty"`                     //Whether SMB encryption is enabled for this share. Encryption is supported
	EncryptionRequired          *bool            `json:"encryptionRequired,omitempty" form:"encryptionRequired,omitempty"`                   //Whether to enforce encryption for all the sessions for this view. When
	IsShareLevelPermissionEmpty *bool            `json:"isShareLevelPermissionEmpty,omitempty" form:"isShareLevelPermissionEmpty,omitempty"` //Indicate if share level permission is cleared by user.
	Permissions                 []*SmbPermission `json:"permissions,omitempty" form:"permissions,omitempty"`                                 //Share level permissions.
}

/*
 * Structure for the custom type AmazonCloudCredentials
 */
type AmazonCloudCredentials struct {
	AccessKeyId      *string          `json:"accessKeyId,omitempty" form:"accessKeyId,omitempty"`           //Specifies the access key for Amazon service account.
	C2sAccessPortal  *C2SAccessPortal `json:"c2sAccessPortal,omitempty" form:"c2sAccessPortal,omitempty"`   //Specifies information required to connect to CAP to get AWS credentials.
	Region           *string          `json:"region,omitempty" form:"region,omitempty"`                     //Specifies the region to use for the Amazon service account.
	SecretAccessKey  *string          `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"`   //Specifies the secret access key for Amazon service account.
	ServiceUrl       *string          `json:"serviceUrl,omitempty" form:"serviceUrl,omitempty"`             //Specifies the URL (Endpoint) for the service such as s3like.notamazon.com.
	SignatureVersion *int64           `json:"signatureVersion,omitempty" form:"signatureVersion,omitempty"` //Specifies the version of the S3 Compliance.
	TierType         TierTypeEnum     `json:"tierType,omitempty" form:"tierType,omitempty"`                 //Specifies the storage class of AWS.
	UseHttps         *bool            `json:"useHttps,omitempty" form:"useHttps,omitempty"`                 //Specifies whether to use http or https to connect to the service.
}

/*
 * Structure for the custom type AntivirusScanConfig
 */
type AntivirusScanConfig struct {
	BlockAccessOnScanFailure *bool                `json:"blockAccessOnScanFailure,omitempty" form:"blockAccessOnScanFailure,omitempty"` //Specifies whether block access to the file when antivirus scan fails.
	IsEnabled                *bool                `json:"isEnabled,omitempty" form:"isEnabled,omitempty"`                               //Specifies whether the antivirus service is enabled or not.
	MaximumScanFileSize      *int64               `json:"maximumScanFileSize,omitempty" form:"maximumScanFileSize,omitempty"`           //Specifies maximum file size that will be sent to antivirus server for
	ScanFilter               *FileExtensionFilter `json:"scanFilter,omitempty" form:"scanFilter,omitempty"`                             //TODO: Write general description for this field
	ScanOnAccess             *bool                `json:"scanOnAccess,omitempty" form:"scanOnAccess,omitempty"`                         //Specifies whether to scan a file when it is opened.
	ScanOnClose              *bool                `json:"scanOnClose,omitempty" form:"scanOnClose,omitempty"`                           //Specifies whether to scan a file when it is closed after modify.
	ScanTimeoutUsecs         *int64               `json:"scanTimeoutUsecs,omitempty" form:"scanTimeoutUsecs,omitempty"`                 //Specifies the maximum amount of time that a scan can take before timing
}

/*
 * Structure for the custom type AntivirusServiceConfig
 */
type AntivirusServiceConfig struct {
	Description *string `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of the Antivirus service. This could be any
	IcapUri     string  `json:"icapUri" form:"icapUri"`                             //Specifies the ICAP uri for this Antivirus service. It is of the form
	Tag         *string `json:"tag,omitempty" form:"tag,omitempty"`                 //Specifies the tag of antivirus service. This is service-specific "cookie"
	TagId       *int64  `json:"tagId,omitempty" form:"tagId,omitempty"`             //Specifies the tag Id of antivirus service.
}

/*
 * Structure for the custom type AntivirusServiceConfigParams
 */
type AntivirusServiceConfigParams struct {
	Description *string `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of the Antivirus service. This could be any
	IcapUri     string  `json:"icapUri" form:"icapUri"`                             //Specifies the ICAP uri for this Antivirus service. It is of the form
}

/*
 * Structure for the custom type AntivirusServiceGroup
 */
type AntivirusServiceGroup struct {
	AntivirusServices []*AntivirusServiceConfig `json:"antivirusServices,omitempty" form:"antivirusServices,omitempty"` //Specifies the Antivirus Services belonging to this antivirus group.
	Description       *string                   `json:"description,omitempty" form:"description,omitempty"`             //Specifies the description of the Antivirus service group.
	Id                int64                     `json:"id" form:"id"`                                                   //Specifies the Id of the Antivirus service group.
	IsEnabled         *bool                     `json:"isEnabled,omitempty" form:"isEnabled,omitempty"`                 //Specifies whether the antivirus service group is enabled or not.
	Name              string                    `json:"name" form:"name"`                                               //Specifies the name of the Antivirus service group.
}

/*
 * Structure for the custom type AntivirusServiceGroupParams
 */
type AntivirusServiceGroupParams struct {
	AntivirusServices []*AntivirusServiceConfigParams `json:"antivirusServices,omitempty" form:"antivirusServices,omitempty"` //Specifies the Antivirus services for this provider.
	Description       *string                         `json:"description,omitempty" form:"description,omitempty"`             //Specifies the description of the Antivirus service group.
	Name              string                          `json:"name" form:"name"`                                               //Specifies the name of the Antivirus service group.
}

/*
 * Structure for the custom type AntivirusServiceGroupStateParams
 */
type AntivirusServiceGroupStateParams struct {
	Enable bool  `json:"enable" form:"enable"` //Specifies the enable flag to enable the Antivirus service group.
	Id     int64 `json:"id" form:"id"`         //Specifies the Id of the Antivirus service group.
}

/*
 * Structure for the custom type AppMetadata
 */
type AppMetadata struct {
	Author           *string `json:"author,omitempty" form:"author,omitempty"`                     //Specifies author of the app.
	CreatedDate      *string `json:"createdDate,omitempty" form:"createdDate,omitempty"`           //Specifies date when the first version of the app was created.
	Description      *string `json:"description,omitempty" form:"description,omitempty"`           //Specifies description about what app does.
	DevVersion       *string `json:"devVersion,omitempty" form:"devVersion,omitempty"`             //Specifies version of the app provided by the developer.
	IconImage        *string `json:"iconImage,omitempty" form:"iconImage,omitempty"`               //Specifies application icon.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty" form:"lastModifiedDate,omitempty"` //Specifies date when the app was last modified.
	Name             *string `json:"name,omitempty" form:"name,omitempty"`                         //Specifies name of the app.
}

/*
 * Structure for the custom type AppOwnerRestoreInfo
 */
type AppOwnerRestoreInfo struct {
	OwnerObject        *RestoreObject       `json:"ownerObject,omitempty" form:"ownerObject,omitempty"`               //TODO: Write general description for this field
	OwnerRestoreParams *RestoreObjectParams `json:"ownerRestoreParams,omitempty" form:"ownerRestoreParams,omitempty"` //TODO: Write general description for this field
	PerformRestore     *bool                `json:"performRestore,omitempty" form:"performRestore,omitempty"`         //If this is set to true, then the owner object needs to be restored. The
}

/*
 * Structure for the custom type AppendHostsParameters
 */
type AppendHostsParameters struct {
	Hosts []*HostEntry `json:"hosts,omitempty" form:"hosts,omitempty"` //Specifies the list of host entries to be added to the Cluster's
}

/*
 * Structure for the custom type ApplicationInfo
 */
type ApplicationInfo struct {
	ApplicationTreeInfo []*ProtectionSourceNode        `json:"applicationTreeInfo,omitempty" form:"applicationTreeInfo,omitempty"` //Application Server and the subtrees below them.
	Environment         EnvironmentApplicationInfoEnum `json:"environment,omitempty" form:"environment,omitempty"`                 //Specifies the environment type of the application such as 'kSQL',
}

/*
 * Structure for the custom type ApplicationParameters
 */
type ApplicationParameters struct {
	TruncateExchangeLog *bool `json:"truncateExchangeLog,omitempty" form:"truncateExchangeLog,omitempty"` //If true, after the Cohesity Cluster successfully captures a Snapshot
}

/*
 * Structure for the custom type ApplicationRestoreObject
 */
type ApplicationRestoreObject struct {
	AdRestoreParameters  *AdRestoreParameters  `json:"adRestoreParameters,omitempty" form:"adRestoreParameters,omitempty"`   //Specifies the parameters specific to Application domain controller.
	ApplicationServerId  *int64                `json:"applicationServerId,omitempty" form:"applicationServerId,omitempty"`   //Specifies the Application Server to restore (for example, kSQL).
	SqlRestoreParameters *SqlRestoreParameters `json:"sqlRestoreParameters,omitempty" form:"sqlRestoreParameters,omitempty"` //Specifies the parameters specific the Application Server instance.
	TargetHostId         *int64                `json:"targetHostId,omitempty" form:"targetHostId,omitempty"`                 //Specifies the target host if the application is to be restored to a
	TargetRootNodeId     *int64                `json:"targetRootNodeId,omitempty" form:"targetRootNodeId,omitempty"`         //Specifies the registered root node, like vCenter, of targetHost.
}

/*
 * Structure for the custom type ApplicationRestoreParameters
 */
type ApplicationRestoreParameters struct {
	ApplicationEnvironment    ApplicationEnvironmentEnum  `json:"applicationEnvironment,omitempty" form:"applicationEnvironment,omitempty"`       //Specifies the Environment of the Application server to restore like
	ApplicationRestoreObjects []*ApplicationRestoreObject `json:"applicationRestoreObjects,omitempty" form:"applicationRestoreObjects,omitempty"` //Specifies the Application Server objects whose data should be restored.
	HostingProtectionSource   *RestoreObjectDetails       `json:"hostingProtectionSource,omitempty" form:"hostingProtectionSource,omitempty"`     //Specifies an object to recover or clone or an object to restore files
}

/*
 * Structure for the custom type ApplicationSpecialParameters
 */
type ApplicationSpecialParameters struct {
	ApplicationEntityIds *[]int64 `json:"applicationEntityIds,omitempty" form:"applicationEntityIds,omitempty"` //Array of Ids of Application Entities like SQL/Oracle instances, and
}

/*
 * Structure for the custom type ApplicationsRestoreTaskRequest
 */
type ApplicationsRestoreTaskRequest struct {
	ApplicationEnvironment    ApplicationEnvironmentApplicationsRestoreTaskRequestEnum `json:"applicationEnvironment" form:"applicationEnvironment"`                           //Specifies the Environment of the Application to restore like 'kSQL', or
	ApplicationRestoreObjects []*ApplicationRestoreObject                              `json:"applicationRestoreObjects,omitempty" form:"applicationRestoreObjects,omitempty"` //Specifies the Application Server objects whose data should be restored
	HostingProtectionSource   RestoreObjectDetails                                     `json:"hostingProtectionSource" form:"hostingProtectionSource"`                         //Specifies an object to recover or clone or an object to restore files
	Name                      string                                                   `json:"name" form:"name"`                                                               //Specifies a name for the new task to be created. This field has to be
	Password                  *string                                                  `json:"password,omitempty" form:"password,omitempty"`                                   //Specifies password of the username to access the target source.
	Username                  *string                                                  `json:"username,omitempty" form:"username,omitempty"`                                   //Specifies username to access the target source.
	VlanParameters            *VlanParameters                                          `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"`                       //Specifies VLAN parameters for the restore operation.
}

/*
 * Structure for the custom type AppsConfig
 */
type AppsConfig struct {
	AllowExternalTraffic  *bool        `json:"allowExternalTraffic,omitempty" form:"allowExternalTraffic,omitempty"`   //Whether to allow pod external traffic.
	AppsMode              AppsModeEnum `json:"appsMode,omitempty" form:"appsMode,omitempty"`                           //Specifies the various modes for running apps.
	AppsSubnet            *Subnet      `json:"appsSubnet,omitempty" form:"appsSubnet,omitempty"`                       //Defines a Subnet (Subnetwork).
	OvercommitMemoryPct   *int64       `json:"overcommitMemoryPct,omitempty" form:"overcommitMemoryPct,omitempty"`     //The system memory to overcommit for apps.
	ReservedCpuMillicores *int64       `json:"reservedCpuMillicores,omitempty" form:"reservedCpuMillicores,omitempty"` //The CPU millicores to reserve for apps.
	ReservedMemoryPct     *int64       `json:"reservedMemoryPct,omitempty" form:"reservedMemoryPct,omitempty"`         //The system memory to reserve for apps.
}

/*
 * Structure for the custom type ArchivalExternalTarget
 */
type ArchivalExternalTarget struct {
	VaultId   *int64        `json:"vaultId,omitempty" form:"vaultId,omitempty"`     //Specifies the id of Archival Vault assigned by the Cohesity Cluster.
	VaultName *string       `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Name of the Archival Vault.
	VaultType VaultTypeEnum `json:"vaultType,omitempty" form:"vaultType,omitempty"` //Specifies the type of the Archival External Target such as 'kCloud',
}

/*
 * Structure for the custom type ArchivalTarget
 */
type ArchivalTarget struct {
	Name    *string `json:"name,omitempty" form:"name,omitempty"`       //The name of the archival target.
	Type    *int64  `json:"type,omitempty" form:"type,omitempty"`       //The type of the archival target.
	VaultId *int64  `json:"vaultId,omitempty" form:"vaultId,omitempty"` //The id of the archival vault.
}

/*
 * Structure for the custom type AttributeRestoreInformation
 */
type AttributeRestoreInformation struct {
	ErrorMessage *[]string `json:"errorMessage,omitempty" form:"errorMessage,omitempty"` //Specifes the error messages corresponding to restore of the attribute.
	Name         *string   `json:"name,omitempty" form:"name,omitempty"`                 //Specifies the name of the attribute of the AD object.
}

/*
 * Structure for the custom type AttributeValue
 */
type AttributeValue struct {
	Flags  *[]FlagEnum `json:"flags,omitempty" form:"flags,omitempty"`   //Specifies the flags related to the attribute values of the AD object.
	Values *[]string   `json:"values,omitempty" form:"values,omitempty"` //Specifies list of values for the attribute.
}

/*
 * Structure for the custom type AuditLogsTile
 */
type AuditLogsTile struct {
	ClusterAuditLogs []*ClusterAuditLog `json:"clusterAuditLogs,omitempty" form:"clusterAuditLogs,omitempty"` //Array of Cluster Audit Logs.
	TotalCount       *int64             `json:"totalCount,omitempty" form:"totalCount,omitempty"`             //Specifies the total number of logs that match the specified
}

/*
 * Structure for the custom type AwsCredentials
 */
type AwsCredentials struct {
	AccessKey          *string     `json:"accessKey,omitempty" form:"accessKey,omitempty"`                   //Specifies Access key of the AWS account.
	AmazonResourceName *string     `json:"amazonResourceName,omitempty" form:"amazonResourceName,omitempty"` //Specifies Amazon Resource Name (owner ID) of the IAM user, act as an
	AwsType            AwsTypeEnum `json:"awsType,omitempty" form:"awsType,omitempty"`                       //Specifies the entity type such as 'kIAMUser' if the environment is kAWS.
	SecretAccessKey    *string     `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"`       //Specifies Secret Access key of the AWS account.
}

/*
 * Structure for the custom type AwsParams
 */
type AwsParams struct {
	InstanceId              *int64     `json:"instanceId,omitempty" form:"instanceId,omitempty"`                           //Specfies id of the AWS instance type in which to deploy the VM.
	NetworkSecurityGroupIds *[]int64   `json:"networkSecurityGroupIds,omitempty" form:"networkSecurityGroupIds,omitempty"` //Specifies ids of the netwrok security groups within above VPC.
	RdsParams               *RdsParams `json:"rdsParams,omitempty" form:"rdsParams,omitempty"`                             //Specifies rds params for the restore operation.
	Region                  *int64     `json:"region,omitempty" form:"region,omitempty"`                                   //Specifies id of the AWS region in which to deploy the VM.
	SubnetId                *int64     `json:"subnetId,omitempty" form:"subnetId,omitempty"`                               //Specifies id of the subnet within above VPC.
	VirtualPrivateCloudId   *int64     `json:"virtualPrivateCloudId,omitempty" form:"virtualPrivateCloudId,omitempty"`     //Specifies id of the Virtual Private Cloud to chose for the instance type.
}

/*
 * Structure for the custom type AwsProtectionSource
 */
type AwsProtectionSource struct {
	AccessKey          *string                     `json:"accessKey,omitempty" form:"accessKey,omitempty"`                   //Specifies Access key of the AWS account.
	AmazonResourceName *string                     `json:"amazonResourceName,omitempty" form:"amazonResourceName,omitempty"` //Specifies Amazon Resource Name (owner ID) of the IAM user, act as an
	AwsType            AwsTypeEnum                 `json:"awsType,omitempty" form:"awsType,omitempty"`                       //Specifies the entity type such as 'kIAMUser' if the environment is kAWS.
	DbEngineId         *string                     `json:"dbEngineId,omitempty" form:"dbEngineId,omitempty"`                 //Specifies DB engine version info of the entity. This is populated only
	HostType           HostTypeEnum                `json:"hostType,omitempty" form:"hostType,omitempty"`                     //Specifies the OS type of the Protection Source of type 'kVirtualMachine'
	IpAddresses        *string                     `json:"ipAddresses,omitempty" form:"ipAddresses,omitempty"`               //Specifies the IP address of the entity of type 'kVirtualMachine'.
	Name               *string                     `json:"name,omitempty" form:"name,omitempty"`                             //Specifies the name of the Object set by the Cloud Provider.
	OwnerId            *string                     `json:"ownerId,omitempty" form:"ownerId,omitempty"`                       //Specifies the owner id of the resource in AWS environment. With type,
	PhysicalSourceId   *int64                      `json:"physicalSourceId,omitempty" form:"physicalSourceId,omitempty"`     //Specifies the Protection Source id of the registered Physical Host.
	RegionId           *string                     `json:"regionId,omitempty" form:"regionId,omitempty"`                     //Specifies the region Id of the entity if the entity is an EC2 instance.
	ResourceId         *string                     `json:"resourceId,omitempty" form:"resourceId,omitempty"`                 //Specifies the unique Id of the resource given by the cloud provider.
	RestoreTaskId      *int64                      `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"`           //Specifies the id of the "convert and deploy" restore task that
	SecretAccessKey    *string                     `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"`       //Specifies Secret Access key of the AWS account.
	TagAttributes      []*TagAttribute             `json:"tagAttributes,omitempty" form:"tagAttributes,omitempty"`           //Specifies the list of AWS tag attributes.
	Type               TypeAwsProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                             //Specifies the type of an AWS Protection Source Object such as
	UserAccountId      *string                     `json:"userAccountId,omitempty" form:"userAccountId,omitempty"`           //Specifies the account id derived from the ARN of the user.
	UserResourceName   *string                     `json:"userResourceName,omitempty" form:"userResourceName,omitempty"`     //Specifies the Amazon Resource Name (ARN) of the user.
}

/*
 * Structure for the custom type AwsSnapshotManagerParameters
 */
type AwsSnapshotManagerParameters struct {
	AmiCreationFrequency *int64 `json:"amiCreationFrequency,omitempty" form:"amiCreationFrequency,omitempty"` //Specifies the frequency of AMI creation. This should be set if the option
	CreateAmi            *bool  `json:"createAmi,omitempty" form:"createAmi,omitempty"`                       //If true, creates an AMI after taking snapshots of the instance. It should
}

/*
 * Structure for the custom type AzureCloudCredentials
 */
type AzureCloudCredentials struct {
	StorageAccessKey   *string                           `json:"storageAccessKey,omitempty" form:"storageAccessKey,omitempty"`     //Specifies the access key to use when accessing a storage tier
	StorageAccountName *string                           `json:"storageAccountName,omitempty" form:"storageAccountName,omitempty"` //Specifies the account name to use when accessing a storage tier
	TierType           TierTypeAzureCloudCredentialsEnum `json:"tierType,omitempty" form:"tierType,omitempty"`                     //Specifies the storage class of Azure.
}

/*
 * Structure for the custom type AzureCredentials
 */
type AzureCredentials struct {
	ApplicationId    *string              `json:"applicationId,omitempty" form:"applicationId,omitempty"`       //Specifies Application Id of the active directory of Azure account.
	ApplicationKey   *string              `json:"applicationKey,omitempty" form:"applicationKey,omitempty"`     //Specifies Application key of the active directory of Azure account.
	AzureType        AzureTypeEnum        `json:"azureType,omitempty" form:"azureType,omitempty"`               //Specifies the entity type such as 'kSubscription' if the environment is
	SubscriptionId   *string              `json:"subscriptionId,omitempty" form:"subscriptionId,omitempty"`     //Specifies Subscription id inside a customer's Azure account. It represents
	SubscriptionType SubscriptionTypeEnum `json:"subscriptionType,omitempty" form:"subscriptionType,omitempty"` //Specifies the subscription type of Azure such as 'kAzureCommercial' or
	TenantId         *string              `json:"tenantId,omitempty" form:"tenantId,omitempty"`                 //Specifies Tenant Id of the active directory of Azure account.
}

/*
 * Structure for the custom type AzureManagedDiskParams
 */
type AzureManagedDiskParams struct {
	DataDisksSkuType *int64 `json:"dataDisksSkuType,omitempty" form:"dataDisksSkuType,omitempty"` //SKU type for data disks.
	OsDiskSkuType    *int64 `json:"osDiskSkuType,omitempty" form:"osDiskSkuType,omitempty"`       //SKU type for OS disk.
}

/*
 * Structure for the custom type AzureParams
 */
type AzureParams struct {
	DataDiskType             DataDiskTypeEnum `json:"dataDiskType,omitempty" form:"dataDiskType,omitempty"`                         //Specifies the disk type used by the data.
	InstanceId               *int64           `json:"instanceId,omitempty" form:"instanceId,omitempty"`                             //Specifies Type of VM (e.g. small, medium, large) when cloning the VM in
	NetworkResourceGroupId   *int64           `json:"networkResourceGroupId,omitempty" form:"networkResourceGroupId,omitempty"`     //Specifies id of the resource group for the selected virtual network.
	OsDiskType               OsDiskTypeEnum   `json:"osDiskType,omitempty" form:"osDiskType,omitempty"`                             //Specifies the disk type used by the OS.
	ResourceGroup            *int64           `json:"resourceGroup,omitempty" form:"resourceGroup,omitempty"`                       //Specifies id of the Azure resource group. Its value is globally unique
	StorageAccount           *int64           `json:"storageAccount,omitempty" form:"storageAccount,omitempty"`                     //Specifies id of the storage account that will contain the storage
	StorageContainer         *int64           `json:"storageContainer,omitempty" form:"storageContainer,omitempty"`                 //Specifies id of the storage container within the above storage account.
	StorageResourceGroupId   *int64           `json:"storageResourceGroupId,omitempty" form:"storageResourceGroupId,omitempty"`     //Specifies id of the resource group for the selected storage account.
	SubnetId                 *int64           `json:"subnetId,omitempty" form:"subnetId,omitempty"`                                 //Specifies Id of the subnet within the above virtual network.
	TempVmResourceGroupId    *int64           `json:"tempVmResourceGroupId,omitempty" form:"tempVmResourceGroupId,omitempty"`       //Specifies the resource group where temporary VM needs to be created.
	TempVmStorageAccountId   *int64           `json:"tempVmStorageAccountId,omitempty" form:"tempVmStorageAccountId,omitempty"`     //Specifies the Storage account where temporary VM needs to be created.
	TempVmStorageContainerId *int64           `json:"tempVmStorageContainerId,omitempty" form:"tempVmStorageContainerId,omitempty"` //Specifies the Storage container where temporary VM needs to be created.
	TempVmSubnetId           *int64           `json:"tempVmSubnetId,omitempty" form:"tempVmSubnetId,omitempty"`                     //Specifies the Subnet where temporary VM needs to be created.
	TempVmVirtualNetworkId   *int64           `json:"tempVmVirtualNetworkId,omitempty" form:"tempVmVirtualNetworkId,omitempty"`     //Specifies the Virtual network where temporary VM needs to be created.
	VirtualNetworkId         *int64           `json:"virtualNetworkId,omitempty" form:"virtualNetworkId,omitempty"`                 //Specifies Id of the Virtual Network.
}

/*
 * Structure for the custom type AzureProtectionSource
 */
type AzureProtectionSource struct {
	ApplicationId    *string                       `json:"applicationId,omitempty" form:"applicationId,omitempty"`       //Specifies Application Id of the active directory of Azure account.
	ApplicationKey   *string                       `json:"applicationKey,omitempty" form:"applicationKey,omitempty"`     //Specifies Application key of the active directory of Azure account.
	AzureType        AzureTypeEnum                 `json:"azureType,omitempty" form:"azureType,omitempty"`               //Specifies the entity type such as 'kSubscription' if the environment is
	HostType         HostTypeEnum                  `json:"hostType,omitempty" form:"hostType,omitempty"`                 //Specifies the OS type of the Protection Source of type 'kVirtualMachine'
	IpAddresses      *[]string                     `json:"ipAddresses,omitempty" form:"ipAddresses,omitempty"`           //Specifies a list of IP addresses for entities of type 'kVirtualMachine'.
	Location         *string                       `json:"location,omitempty" form:"location,omitempty"`                 //Specifies the physical location of the resource group.
	MemoryMbytes     *int64                        `json:"memoryMbytes,omitempty" form:"memoryMbytes,omitempty"`         //Specifies the amount of memory in MegaBytes of the Azure resource of
	Name             *string                       `json:"name,omitempty" form:"name,omitempty"`                         //Specifies the name of the Object set by the Cloud Provider.
	NumCores         *int64                        `json:"numCores,omitempty" form:"numCores,omitempty"`                 //Specifies the number of CPU cores of the Azure resource of
	PhysicalSourceId *int64                        `json:"physicalSourceId,omitempty" form:"physicalSourceId,omitempty"` //Specifies the Protection Source id of the registered Physical Host.
	ResourceId       *string                       `json:"resourceId,omitempty" form:"resourceId,omitempty"`             //Specifies the unique Id of the resource given by the cloud provider.
	RestoreTaskId    *int64                        `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"`       //Specifies the id of the "convert and deploy" restore task that
	SubscriptionId   *string                       `json:"subscriptionId,omitempty" form:"subscriptionId,omitempty"`     //Specifies Subscription id inside a customer's Azure account. It represents
	SubscriptionType SubscriptionTypeEnum          `json:"subscriptionType,omitempty" form:"subscriptionType,omitempty"` //Specifies the subscription type of Azure such as 'kAzureCommercial' or
	TenantId         *string                       `json:"tenantId,omitempty" form:"tenantId,omitempty"`                 //Specifies Tenant Id of the active directory of Azure account.
	Type             TypeAzureProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                         //Specifies the type of an Azure Protection Source Object such as
}

/*
 * Structure for the custom type BackupJobPreOrPostScript
 */
type BackupJobPreOrPostScript struct {
	BackupScript     *ScriptPathAndParams       `json:"backupScript,omitempty" form:"backupScript,omitempty"`         //A message to encapsulate pre or post script associated with a backup job
	FullBackupScript *ScriptPathAndParams       `json:"fullBackupScript,omitempty" form:"fullBackupScript,omitempty"` //A message to encapsulate pre or post script associated with a backup job
	LogBackupScript  *ScriptPathAndParams       `json:"logBackupScript,omitempty" form:"logBackupScript,omitempty"`   //A message to encapsulate pre or post script associated with a backup job
	RemoteHostParams *RemoteHostConnectorParams `json:"remoteHostParams,omitempty" form:"remoteHostParams,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type BackupJobProto
 */
type BackupJobProto struct {
	AbortInExclusionWindow               *bool                               `json:"abortInExclusionWindow,omitempty" form:"abortInExclusionWindow,omitempty"`                             //This field determines whether a backup run should be aborted when it hits
	AlertingPolicy                       *AlertingPolicyProto                `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"`                                             //TODO: Write general description for this field
	BackupQosPrincipal                   *int64                              `json:"backupQosPrincipal,omitempty" form:"backupQosPrincipal,omitempty"`                                     //The backup QoS principal to use for the backup job.
	BackupSourceParams                   []*BackupSourceParams               `json:"backupSourceParams,omitempty" form:"backupSourceParams,omitempty"`                                     //This contains additional backup params that are applicable to sources
	ContinueOnQuiesceFailure             *bool                               `json:"continueOnQuiesceFailure,omitempty" form:"continueOnQuiesceFailure,omitempty"`                         //Whether to continue backing up on quiesce failure.
	CreateRemoteView                     *bool                               `json:"createRemoteView,omitempty" form:"createRemoteView,omitempty"`                                         //If set to false, a remote view will not be created.
	DedupDisabledSourceIdVec             *[]int64                            `json:"dedupDisabledSourceIdVec,omitempty" form:"dedupDisabledSourceIdVec,omitempty"`                         //List of source ids for which source side dedup is disabled from the backup
	DeletionStatus                       *int64                              `json:"deletionStatus,omitempty" form:"deletionStatus,omitempty"`                                             //Determines if the job (and associated backups) should be deleted. Once a
	Description                          *string                             `json:"description,omitempty" form:"description,omitempty"`                                                   //Job description (as entered by the user).
	DrToCloudParams                      *BackupJobProtoDRToCloudParams      `json:"drToCloudParams,omitempty" form:"drToCloudParams,omitempty"`                                           //A Proto needed in case objects backed up by this job need to DR to cloud.
	EhParentSource                       *EntityProto                        `json:"ehParentSource,omitempty" form:"ehParentSource,omitempty"`                                             //Specifies the attributes and the latest statistics about an entity.
	EndTimeUsecs                         *int64                              `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                                 //The time (in usecs) after which no backup for the job should be scheduled.
	EnvBackupParams                      *EnvBackupParams                    `json:"envBackupParams,omitempty" form:"envBackupParams,omitempty"`                                           //Message to capture any additional environment specific backup params at the
	ExcludeSources                       []*BackupJobProtoExcludeSource      `json:"excludeSources,omitempty" form:"excludeSources,omitempty"`                                             //The list of sources to exclude from backups. These can have non-leaf-level
	ExcludeSourcesDEPRECATED             []*EntityProto                      `json:"excludeSources_DEPRECATED,omitempty" form:"excludeSources_DEPRECATED,omitempty"`                       //The list of sources to exclude from backups. These can have non-leaf-level
	ExclusionRanges                      []*BackupJobProtoExclusionTimeRange `json:"exclusionRanges,omitempty" form:"exclusionRanges,omitempty"`                                           //Do not run backups in these time-ranges.
	FullBackupJobPolicy                  *JobPolicyProto                     `json:"fullBackupJobPolicy,omitempty" form:"fullBackupJobPolicy,omitempty"`                                   //A message that specifies the policies to use for a job.
	FullBackupSlaTimeMins                *int64                              `json:"fullBackupSlaTimeMins,omitempty" form:"fullBackupSlaTimeMins,omitempty"`                               //Same as 'sla_time_mins' above, but applies to full backups.
	IndexingPolicy                       *IndexingPolicyProto                `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"`                                             //Proto to encapsulate file level indexing policy for VMs in a backup job.
	IsActive                             *bool                               `json:"isActive,omitempty" form:"isActive,omitempty"`                                                         //Whether the backup job is active or not. Details about what an active job
	IsDeleted                            *bool                               `json:"isDeleted,omitempty" form:"isDeleted,omitempty"`                                                       //Tracks whether the backup job has actually been deleted.
	IsPaused                             *bool                               `json:"isPaused,omitempty" form:"isPaused,omitempty"`                                                         //Whether the backup job is paused. New backup runs are not scheduled for
	IsRpoJob                             *bool                               `json:"isRpoJob,omitempty" form:"isRpoJob,omitempty"`                                                         //Whether the backup job is an RPO policy job. These jobs are hidden from
	JobCreationTimeUsecs                 *int64                              `json:"jobCreationTimeUsecs,omitempty" form:"jobCreationTimeUsecs,omitempty"`                                 //Time when this job was first created.
	JobId                                *int64                              `json:"jobId,omitempty" form:"jobId,omitempty"`                                                               //A unique id for locally created jobs. This should only be used to identify
	JobPolicy                            *JobPolicyProto                     `json:"jobPolicy,omitempty" form:"jobPolicy,omitempty"`                                                       //A message that specifies the policies to use for a job.
	JobUid                               *UniversalIdProto                   `json:"jobUid,omitempty" form:"jobUid,omitempty"`                                                             //TODO: Write general description for this field
	LastModificationTimeUsecs            *int64                              `json:"lastModificationTimeUsecs,omitempty" form:"lastModificationTimeUsecs,omitempty"`                       //Time when this job description was last updated.
	LastPauseModificationTimeUsecs       *int64                              `json:"lastPauseModificationTimeUsecs,omitempty" form:"lastPauseModificationTimeUsecs,omitempty"`             //Time when the job was last paused or unpaused.
	LastPauseReason                      *int64                              `json:"lastPauseReason,omitempty" form:"lastPauseReason,omitempty"`                                           //Last reason for pausing the backup job. Capturing the reason will help in
	LastUpdatedUsername                  *string                             `json:"lastUpdatedUsername,omitempty" form:"lastUpdatedUsername,omitempty"`                                   //The user who modified the job most recently.
	LeverageStorageSnapshots             *bool                               `json:"leverageStorageSnapshots,omitempty" form:"leverageStorageSnapshots,omitempty"`                         //Whether to leverage the storage array based snapshots for this backup
	LeverageStorageSnapshotsForHyperflex *bool                               `json:"leverageStorageSnapshotsForHyperflex,omitempty" form:"leverageStorageSnapshotsForHyperflex,omitempty"` //This is set to true by the user if hyperflex snapshots are requested
	LogBackupJobPolicy                   *JobPolicyProto                     `json:"logBackupJobPolicy,omitempty" form:"logBackupJobPolicy,omitempty"`                                     //A message that specifies the policies to use for a job.
	Name                                 *string                             `json:"name,omitempty" form:"name,omitempty"`                                                                 //The name of this backup job. This must be unique across all jobs.
	NumSnapshotsToKeepOnPrimary          *int64                              `json:"numSnapshotsToKeepOnPrimary,omitempty" form:"numSnapshotsToKeepOnPrimary,omitempty"`                   //Specifies how many recent snapshots of each backed up entity to retain on
	ParentSource                         *EntityProto                        `json:"parentSource,omitempty" form:"parentSource,omitempty"`                                                 //Specifies the attributes and the latest statistics about an entity.
	PerformSourceSideDedup               *bool                               `json:"performSourceSideDedup,omitempty" form:"performSourceSideDedup,omitempty"`                             //Whether or not to perform source side dedup.
	PolicyAppliedTimeMsecs               *int64                              `json:"policyAppliedTimeMsecs,omitempty" form:"policyAppliedTimeMsecs,omitempty"`                             //Epoch time in milliseconds when the policy was last applied to this job.
	PolicyId                             *string                             `json:"policyId,omitempty" form:"policyId,omitempty"`                                                         //Id of the policy being applied to the backup job. It is expected to be of
	PolicyName                           *string                             `json:"policyName,omitempty" form:"policyName,omitempty"`                                                     //The name of the policy referred to by policy_uid. This field can be stale
	PostBackupScript                     *BackupJobPreOrPostScript           `json:"postBackupScript,omitempty" form:"postBackupScript,omitempty"`                                         //A message to encapsulate the pre and post scripts associated with a backup
	PreScript                            *BackupJobPreOrPostScript           `json:"preScript,omitempty" form:"preScript,omitempty"`                                                       //A message to encapsulate the pre and post scripts associated with a backup
	PrimaryJobUid                        *UniversalIdProto                   `json:"primaryJobUid,omitempty" form:"primaryJobUid,omitempty"`                                               //TODO: Write general description for this field
	Priority                             *int64                              `json:"priority,omitempty" form:"priority,omitempty"`                                                         //The priority for the job. This is used at admission time - all admitted
	Quiesce                              *bool                               `json:"quiesce,omitempty" form:"quiesce,omitempty"`                                                           //Whether to take app-consistent snapshots by quiescing apps and the
	RemoteJobUids                        []*UniversalIdProto                 `json:"remoteJobUids,omitempty" form:"remoteJobUids,omitempty"`                                               //The globally unique ids of all remote jobs that are linked to this job
	RemoteViewName                       *string                             `json:"remoteViewName,omitempty" form:"remoteViewName,omitempty"`                                             //A human readable name of the remote view. A remote view is created with
	RequiredFeatureVec                   *[]string                           `json:"requiredFeatureVec,omitempty" form:"requiredFeatureVec,omitempty"`                                     //The features that are strictly required to be supported by the cluster
	SlaTimeMins                          *int64                              `json:"slaTimeMins,omitempty" form:"slaTimeMins,omitempty"`                                                   //If specified, this variable determines the amount of time (after backup
	Sources                              []*BackupJobProtoBackupSource       `json:"sources,omitempty" form:"sources,omitempty"`                                                           //The list of sources that should be backed up. A source in this list could
	StartTime                            *time.Time                          `json:"startTime,omitempty" form:"startTime,omitempty"`                                                       //A message to encapusulate time of a day. Users of this proto will have to
	StubbingPolicy                       *StubbingPolicyProto                `json:"stubbingPolicy,omitempty" form:"stubbingPolicy,omitempty"`                                             //Stubbing jobs do not use protection policies. Instead, schedule and
	TagVec                               *[]string                           `json:"tagVec,omitempty" form:"tagVec,omitempty"`                                                             //Tags associated with the job. User can specify tags/keywords that can
	Timezone                             *string                             `json:"timezone,omitempty" form:"timezone,omitempty"`                                                         //Timezone of the backup job. All time fields (i.e., TimeOfDay) in this
	TruncateLogs                         *bool                               `json:"truncateLogs,omitempty" form:"truncateLogs,omitempty"`                                                 //Whether to truncate logs after a backup run. This is currently only
	Type                                 *int64                              `json:"type,omitempty" form:"type,omitempty"`                                                                 //The type of environment this backup job corresponds to.
	UserInfo                             *UserInformation                    `json:"userInfo,omitempty" form:"userInfo,omitempty"`                                                         //A message to encapsulate information about the user who made the request.
	ViewBoxId                            *int64                              `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                                                       //The view box to which data will be written.
}

/*
 * Structure for the custom type BackupJobProtoBackupSource
 */
type BackupJobProtoBackupSource struct {
	Entities []*EntityProto `json:"entities,omitempty" form:"entities,omitempty"` //Source entities.
}

/*
 * Structure for the custom type BackupJobProtoDRToCloudParams
 */
type BackupJobProtoDRToCloudParams struct {
	NeedToFailOver *bool `json:"needToFailOver,omitempty" form:"needToFailOver,omitempty"` //Whether the objects in this job will be failed over to cloud.
}

/*
 * Structure for the custom type BackupJobProtoExcludeSource
 */
type BackupJobProtoExcludeSource struct {
	Entities []*EntityProto `json:"entities,omitempty" form:"entities,omitempty"` //An intersection of leaf-level entities will be obtained after expanding
}

/*
 * Structure for the custom type BackupJobProtoExclusionTimeRange
 */
type BackupJobProtoExclusionTimeRange struct {
	Day       *int64     `json:"day,omitempty" form:"day,omitempty"`             //If the day is not set, the time range applies to all days.
	EndTime   *time.Time `json:"endTime,omitempty" form:"endTime,omitempty"`     //A message to encapusulate time of a day. Users of this proto will have to
	StartTime *time.Time `json:"startTime,omitempty" form:"startTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProto
 */
type BackupPolicyProto struct {
	ContinuousSchedule      *BackupPolicyProtoContinuousSchedule `json:"continuousSchedule,omitempty" form:"continuousSchedule,omitempty"`           //TODO: Write general description for this field
	DailySchedule           *BackupPolicyProtoDailySchedule      `json:"dailySchedule,omitempty" form:"dailySchedule,omitempty"`                     //The daily schedule encompasses weekly schedules as well. This has been
	MonthlySchedule         *BackupPolicyProtoMonthlySchedule    `json:"monthlySchedule,omitempty" form:"monthlySchedule,omitempty"`                 //TODO: Write general description for this field
	Name                    *string                              `json:"name,omitempty" form:"name,omitempty"`                                       //A backup schedule can have an optional name.
	NumDaysToKeep           *int64                               `json:"numDaysToKeep,omitempty" form:"numDaysToKeep,omitempty"`                     //Specifies how to determine the expiration time for snapshots created by
	NumRetries              *int64                               `json:"numRetries,omitempty" form:"numRetries,omitempty"`                           //The number of retries to perform (for retryable errors) before giving up.
	OneOffSchedule          *BackupPolicyProtoOneOffSchedule     `json:"oneOffSchedule,omitempty" form:"oneOffSchedule,omitempty"`                   //TODO: Write general description for this field
	Periodicity             *int64                               `json:"periodicity,omitempty" form:"periodicity,omitempty"`                         //Determines how often the job should be run.
	RetryDelayMins          *int64                               `json:"retryDelayMins,omitempty" form:"retryDelayMins,omitempty"`                   //The number of minutes to wait before retrying a failed job.
	ScheduleEnd             *BackupPolicyProtoScheduleEnd        `json:"scheduleEnd,omitempty" form:"scheduleEnd,omitempty"`                         //TODO: Write general description for this field
	StartWindowIntervalMins *int64                               `json:"startWindowIntervalMins,omitempty" form:"startWindowIntervalMins,omitempty"` //This field determines the amount of time (in minutes) after which a
	TruncateLogs            *bool                                `json:"truncateLogs,omitempty" form:"truncateLogs,omitempty"`                       //Whether to truncate logs after a backup run. This is currently only
}

/*
 * Structure for the custom type BackupPolicyProtoContinuousSchedule
 */
type BackupPolicyProtoContinuousSchedule struct {
	BackupIntervalMins *int64                                 `json:"backupIntervalMins,omitempty" form:"backupIntervalMins,omitempty"` //If this field is set, backups will be performed periodically every
	ExclusionRanges    []*BackupPolicyProtoExclusionTimeRange `json:"exclusionRanges,omitempty" form:"exclusionRanges,omitempty"`       //Do not start backups in these time-ranges. It's possible for a
}

/*
 * Structure for the custom type BackupPolicyProtoDailySchedule
 */
type BackupPolicyProtoDailySchedule struct {
	Days *[]int64   `json:"days,omitempty" form:"days,omitempty"` //The days of the week backup must be performed. If no days are specified,
	Time *time.Time `json:"time,omitempty" form:"time,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoExclusionTimeRange
 */
type BackupPolicyProtoExclusionTimeRange struct {
	Day       *int64     `json:"day,omitempty" form:"day,omitempty"`             //If the day is not set, the time range applies to all days.
	EndTime   *time.Time `json:"endTime,omitempty" form:"endTime,omitempty"`     //A message to encapusulate time of a day. Users of this proto will have to
	StartTime *time.Time `json:"startTime,omitempty" form:"startTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoMonthlySchedule
 */
type BackupPolicyProtoMonthlySchedule struct {
	Count *int64     `json:"count,omitempty" form:"count,omitempty"` //Count of the day on which to perform the backup (look above for a more
	Day   *int64     `json:"day,omitempty" form:"day,omitempty"`     //The day of the month the backup is to be performed.
	Time  *time.Time `json:"time,omitempty" form:"time,omitempty"`   //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoOneOffSchedule
 */
type BackupPolicyProtoOneOffSchedule struct {
	Time *time.Time `json:"time,omitempty" form:"time,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoScheduleEnd
 */
type BackupPolicyProtoScheduleEnd struct {
	EndAfterNumBackups *int64 `json:"endAfterNumBackups,omitempty" form:"endAfterNumBackups,omitempty"` //The following field has been deprecated.
	EndTimeUsecs       *int64 `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`             //If specified, the backup job will no longer be run after this time.
}

/*
 * Structure for the custom type BackupRun
 */
type BackupRun struct {
	Environment               EnvironmentBackupRunEnum `json:"environment,omitempty" form:"environment,omitempty"`                             //Specifies the environment type that the task is protecting.
	Error                     *string                  `json:"error,omitempty" form:"error,omitempty"`                                         //Specifies if an error occurred (if any) while running this task.
	JobRunId                  *int64                   `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                                   //Specifies the id of the Job Run that ran the backup task and
	Message                   *string                  `json:"message,omitempty" form:"message,omitempty"`                                     //Specifies a message after finishing the task successfully. This field
	MetadataDeleted           *bool                    `json:"metadataDeleted,omitempty" form:"metadataDeleted,omitempty"`                     //Specifies if the metadata and snapshots associated with this Job Run
	Quiesced                  *bool                    `json:"quiesced,omitempty" form:"quiesced,omitempty"`                                   //Specifies if app-consistent snapshot was captured. This field is set to
	RunType                   RunTypeEnum              `json:"runType,omitempty" form:"runType,omitempty"`                                     //Specifies the type of backup such as 'kRegular', 'kFull', 'kLog' or
	SlaViolated               *bool                    `json:"slaViolated,omitempty" form:"slaViolated,omitempty"`                             //Specifies if the SLA was violated for the Job Run. This field is set
	SnapshotsDeleted          *bool                    `json:"snapshotsDeleted,omitempty" form:"snapshotsDeleted,omitempty"`                   //Specifies if backup snapshots associated
	SnapshotsDeletedTimeUsecs *int64                   `json:"snapshotsDeletedTimeUsecs,omitempty" form:"snapshotsDeletedTimeUsecs,omitempty"` //Specifies if backup snapshots associated
	SourceBackupStatus        []*SourceBackupStatus    `json:"sourceBackupStatus,omitempty" form:"sourceBackupStatus,omitempty"`               //Array of Source Object Backup Status.
	Stats                     *ProtectionJobRunStats   `json:"stats,omitempty" form:"stats,omitempty"`                                         //Specifies statistics about a Protection Job Run.
	Status                    StatusBackupRunEnum      `json:"status,omitempty" form:"status,omitempty"`                                       //Specifies the status of Backup task such as 'kRunning', 'kSuccess',
	Warnings                  *[]string                `json:"warnings,omitempty" form:"warnings,omitempty"`                                   //Array of Warnings.
	WormRetentionType         WormRetentionTypeEnum    `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"`                 //Specifies WORM retention type for the snapshot as given by the policy.
}

/*
 * Structure for the custom type BackupScript
 */
type BackupScript struct {
	FullBackupScript        *RemoteScriptPathAndParams `json:"fullBackupScript,omitempty" form:"fullBackupScript,omitempty"`               //Specifies the script that should run for the Full (no CBT) backup schedule
	IncrementalBackupScript *RemoteScriptPathAndParams `json:"incrementalBackupScript,omitempty" form:"incrementalBackupScript,omitempty"` //Specifies the script that should run for the CBT-based backup
	LogBackupScript         *RemoteScriptPathAndParams `json:"logBackupScript,omitempty" form:"logBackupScript,omitempty"`                 //Specifies the script that should run for the Log backup schedule
	RemoteHost              *RemoteHost                `json:"remoteHost,omitempty" form:"remoteHost,omitempty"`                           //Specifies the remote host where the remote scripts are executed.
	Username                *string                    `json:"username,omitempty" form:"username,omitempty"`                               //Specifies the username that will be used to login to the remote host.
}

/*
 * Structure for the custom type BackupSourceParams
 */
type BackupSourceParams struct {
	AppEntityIdVec *[]int64                    `json:"appEntityIdVec,omitempty" form:"appEntityIdVec,omitempty"` //If we are backing up an application (such as SQL), this contains
	OracleParams   *OracleSourceParams         `json:"oracleParams,omitempty" form:"oracleParams,omitempty"`     //Message to capture additional backup/restore params for a Oracle source.
	PhysicalParams *PhysicalBackupSourceParams `json:"physicalParams,omitempty" form:"physicalParams,omitempty"` //Message to capture additional backup params for a Physical type source.
	SkipIndexing   *bool                       `json:"skipIndexing,omitempty" form:"skipIndexing,omitempty"`     //Set to true, if indexing is not required for given source.
	SourceId       *int64                      `json:"sourceId,omitempty" form:"sourceId,omitempty"`             //Source entity id.
	VmwareParams   *VmwareBackupSourceParams   `json:"vmwareParams,omitempty" form:"vmwareParams,omitempty"`     //Message to capture additional backup params for a VMware type source.
}

/*
 * Structure for the custom type BackupSourceStats
 */
type BackupSourceStats struct {
	AdmittedTimeUsecs            *int64 `json:"admittedTimeUsecs,omitempty" form:"admittedTimeUsecs,omitempty"`                       //Specifies the time the task was unqueued from the queue to start running.
	EndTimeUsecs                 *int64 `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                 //Specifies the end time of the Protection Run. The end time
	StartTimeUsecs               *int64 `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                             //Specifies the start time of the Protection Run. The start time
	TimeTakenUsecs               *int64 `json:"timeTakenUsecs,omitempty" form:"timeTakenUsecs,omitempty"`                             //Specifies the actual execution time for the protection run to complete
	TotalBytesReadFromSource     *int64 `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"`         //Specifies the total amount of data read from the source (so far).
	TotalBytesToReadFromSource   *int64 `json:"totalBytesToReadFromSource,omitempty" form:"totalBytesToReadFromSource,omitempty"`     //Specifies the total amount of data expected to be read from the
	TotalLogicalBackupSizeBytes  *int64 `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"`   //Specifies the size of the source object (such as a VM) protected by
	TotalPhysicalBackupSizeBytes *int64 `json:"totalPhysicalBackupSizeBytes,omitempty" form:"totalPhysicalBackupSizeBytes,omitempty"` //Specifies the total amount of physical space used on the Cohesity
	TotalSourceSizeBytes         *int64 `json:"totalSourceSizeBytes,omitempty" form:"totalSourceSizeBytes,omitempty"`                 //Specifies the size of the source object (such as a VM) protected by
}

/*
 * Structure for the custom type BackupTaskInfo
 */
type BackupTaskInfo struct {
	InstanceId     *string `json:"instanceId,omitempty" form:"instanceId,omitempty"`         //Id of that particular instance
	Name           *string `json:"name,omitempty" form:"name,omitempty"`                     //Name of the recovery task.
	StartTimeUsecs *string `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Denotes the start time of the backuptask, needed for deeplinking.
	TaskId         *string `json:"taskId,omitempty" form:"taskId,omitempty"`                 //Id of the recovery task.
}

/*
 * Structure for the custom type BandwidthLimit
 */
type BandwidthLimit struct {
	BandwidthLimitOverrides []*BandwidthLimitOverride `json:"bandwidthLimitOverrides,omitempty" form:"bandwidthLimitOverrides,omitempty"` //Array of Override Bandwidth Limits.
	RateLimitBytesPerSec    *int64                    `json:"rateLimitBytesPerSec,omitempty" form:"rateLimitBytesPerSec,omitempty"`       //Specifies the maximum allowed data transfer rate between the local Cluster
	Timezone                *string                   `json:"timezone,omitempty" form:"timezone,omitempty"`                               //Specifies a time zone for the specified time period.
}

/*
 * Structure for the custom type BandwidthLimitOverride
 */
type BandwidthLimitOverride struct {
	BytesPerSecond *int64       `json:"bytesPerSecond,omitempty" form:"bytesPerSecond,omitempty"` //Specifies the value to override the regular maximum bandwidth rate
	TimePeriods    *TimeOfAWeek `json:"timePeriods,omitempty" form:"timePeriods,omitempty"`       //Specifies a time period by specifying a single daily time period
}

/*
 * Structure for the custom type BasicClusterInfo
 */
type BasicClusterInfo struct {
	AuthenticationType     AuthenticationTypeEnum `json:"authenticationType,omitempty" form:"authenticationType,omitempty"`         //Specifies the authentication scheme for the cluster.
	BannerEnabled          *bool                  `json:"bannerEnabled,omitempty" form:"bannerEnabled,omitempty"`                   //Specifies if banner is enabled on the cluster.
	ClusterSoftwareVersion *string                `json:"clusterSoftwareVersion,omitempty" form:"clusterSoftwareVersion,omitempty"` //Specifies the current release of the Cohesity software running on
	ClusterType            ClusterTypeEnum        `json:"clusterType,omitempty" form:"clusterType,omitempty"`                       //Specifies the type of Cohesity Cluster.
	Domains                *[]string              `json:"domains,omitempty" form:"domains,omitempty"`                               //Array of Domains.
	IdpConfigured          *bool                  `json:"idpConfigured,omitempty" form:"idpConfigured,omitempty"`                   //Specifies Idp is configured for the Cluster.
	IdpTenantExists        *bool                  `json:"idpTenantExists,omitempty" form:"idpTenantExists,omitempty"`               //Specifies Idp is configured for a Tenant.
	LanguageLocale         *string                `json:"languageLocale,omitempty" form:"languageLocale,omitempty"`                 //Specifies the language and locale for the Cohesity Cluster.
	McmMode                *bool                  `json:"mcmMode,omitempty" form:"mcmMode,omitempty"`                               //Specifies whether server is running in mcm-mode. If set to true,
	McmOnPremMode          *bool                  `json:"mcmOnPremMode,omitempty" form:"mcmOnPremMode,omitempty"`                   //Specifies whether server is running in mcm-on-prem-mode. If set to true,
	MultiTenancyEnabled    *bool                  `json:"multiTenancyEnabled,omitempty" form:"multiTenancyEnabled,omitempty"`       //Specifies if multi-tenancy is enabled on the cluster.
	Name                   *string                `json:"name,omitempty" form:"name,omitempty"`                                     //Specifies the name of the Cohesity Cluster.
}

/*
 * Structure for the custom type BasicTaskInfo
 */
type BasicTaskInfo struct {
	Name   *string `json:"name,omitempty" form:"name,omitempty"`     //Name of the recovery task.
	TaskId *string `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
}

/*
 * Structure for the custom type BlackoutPeriod
 */
type BlackoutPeriod struct {
	Day       DayBlackoutPeriodEnum `json:"day,omitempty" form:"day,omitempty"`             //Blackout Day.
	EndTime   *TimeOfDay            `json:"endTime,omitempty" form:"endTime,omitempty"`     //Specifies the end time of the blackout time range.
	StartTime *TimeOfDay            `json:"startTime,omitempty" form:"startTime,omitempty"` //Specifies the start time of the blackout time range.
}

/*
 * Structure for the custom type C2SAccessPortal
 */
type C2SAccessPortal struct {
	Agency                    *string `json:"agency,omitempty" form:"agency,omitempty"`                                       //Name of the agency.
	BaseUrl                   *string `json:"baseUrl,omitempty" form:"baseUrl,omitempty"`                                     //The base url of C2S CAP server.
	ClientCertificatePassword *string `json:"clientCertificatePassword,omitempty" form:"clientCertificatePassword,omitempty"` //Encrypted password of the client private key.
	Mission                   *string `json:"mission,omitempty" form:"mission,omitempty"`                                     //Name of the mission.
	Role                      *string `json:"role,omitempty" form:"role,omitempty"`                                           //Role type.
}

/*
 * Structure for the custom type CancelProtectionJobRunParam
 */
type CancelProtectionJobRunParam struct {
	CopyTaskUid *UniversalId `json:"copyTaskUid,omitempty" form:"copyTaskUid,omitempty"` //Specifies an id for an object that is unique across Cohesity Clusters.
	JobRunId    *int64       `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`       //Run Id of a Protection Job Run that needs to be cancelled. If this Run
}

/*
 * Structure for the custom type CapacityByTier
 */
type CapacityByTier struct {
	StorageTier                  StorageTierEnum `json:"storageTier,omitempty" form:"storageTier,omitempty"`                                   //StorageTier is the type of StorageTier.
	TierMaxPhysicalCapacityBytes *int64          `json:"tierMaxPhysicalCapacityBytes,omitempty" form:"tierMaxPhysicalCapacityBytes,omitempty"` //TierMaxPhysicalCapacityBytes is the maximum physical capacity in bytes of
}

/*
 * Structure for the custom type CentrifyZone
 */
type CentrifyZone struct {
	CentrifySchema    CentrifySchemaEnum `json:"centrifySchema,omitempty" form:"centrifySchema,omitempty"`       //Specifies the schema of this Centrify zone.
	Description       *string            `json:"description,omitempty" form:"description,omitempty"`             //Specifies a description of the Centrify zone.
	DistinguishedName *string            `json:"distinguishedName,omitempty" form:"distinguishedName,omitempty"` //Specifies the distinguished name of the Centrify zone.
}

/*
 * Structure for the custom type CertificateDetails
 */
type CertificateDetails struct {
	CertFileName *string   `json:"certFileName,omitempty" form:"certFileName,omitempty"` //Specifies the filename of the certificate. This is unique to each
	ExpiryDate   *string   `json:"expiryDate,omitempty" form:"expiryDate,omitempty"`     //Specifies the date till when the certificate is valid.
	HostIps      *[]string `json:"hostIps,omitempty" form:"hostIps,omitempty"`           //Each certificate can be deployed to multiple hosts. List of all hosts
}

/*
 * Structure for the custom type ChangeProtectionJobStateParam
 */
type ChangeProtectionJobStateParam struct {
	Pause       *bool  `json:"pause,omitempty" form:"pause,omitempty"`             //If true, the specified Protection Job is paused and no new Runs
	PauseReason *int64 `json:"pauseReason,omitempty" form:"pauseReason,omitempty"` //Specifies the reason of pausing the job so that depending on the pause
}

/*
 * Structure for the custom type ChangeServiceStateParameters
 */
type ChangeServiceStateParameters struct {
	Action   ActionEnum     `json:"action,omitempty" form:"action,omitempty"`     //Specifies the action to take on the specified service.
	Services *[]ServiceEnum `json:"services,omitempty" form:"services,omitempty"` //Specifies the list of services to take the specified action on.
}

/*
 * Structure for the custom type ChangeServiceStateResult
 */
type ChangeServiceStateResult struct {
	Message   *string `json:"message,omitempty" form:"message,omitempty"`     //Specifies a description of the result of the operation.
	StatusUrl *string `json:"statusUrl,omitempty" form:"statusUrl,omitempty"` //Specifies a URL which can be queried to check the status of the
}

/*
 * Structure for the custom type ChassisInfo
 */
type ChassisInfo struct {
	ChassisId   *int64  `json:"chassisId,omitempty" form:"chassisId,omitempty"`     //ChassisId is a unique id assigned to the chassis.
	ChassisName *string `json:"chassisName,omitempty" form:"chassisName,omitempty"` //ChassisName is the name of the chassis. This could be the chassis serial
	Location    *string `json:"location,omitempty" form:"location,omitempty"`       //Location is the location of the chassis within the rack.
	RackId      *int64  `json:"rackId,omitempty" form:"rackId,omitempty"`           //Rack is the rack within which this chassis lives.
}

/*
 * Structure for the custom type CifsShareInfo
 */
type CifsShareInfo struct {
	Acls       *[]string `json:"acls,omitempty" form:"acls,omitempty"`             //Array of Access Control Lists.
	Name       *string   `json:"name,omitempty" form:"name,omitempty"`             //Specifies the name of the CIFS share.
	Path       *string   `json:"path,omitempty" form:"path,omitempty"`             //Specifies the path of this share under the Vserver's root.
	ServerName *string   `json:"serverName,omitempty" form:"serverName,omitempty"` //Specifies the CIFS server name (such as 'NETAPP-01') specified by the
}

/*
 * Structure for the custom type ClearNlmLocksParameters
 */
type ClearNlmLocksParameters struct {
	ClientId *string `json:"clientId,omitempty" form:"clientId,omitempty"` //Specifies the id of the client, related NLM locks needs to be clear.
	FilePath *string `json:"filePath,omitempty" form:"filePath,omitempty"` //Specifies the filepath in the view relative to the root filesystem.
	ViewName *string `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the View in which to search. If a view name is not
}

/*
 * Structure for the custom type CloneAppViewInfoOracle
 */
type CloneAppViewInfoOracle struct {
	MountPathInfoVec *[]string `json:"mountPathInfoVec,omitempty" form:"mountPathInfoVec,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type CloneAppViewInfoProto
 */
type CloneAppViewInfoProto struct {
	OracleAppViewRestoreInfo *CloneAppViewInfoOracle `json:"oracleAppViewRestoreInfo,omitempty" form:"oracleAppViewRestoreInfo,omitempty"` //This message encapsulates backup view Clone operation information of a
}

/*
 * Structure for the custom type CloneAppViewParams
 */
type CloneAppViewParams struct {
	MountPathIdentifier *string `json:"mountPathIdentifier,omitempty" form:"mountPathIdentifier,omitempty"` //Mount path identifier, which identifies the sub-dir where the cohesity
}

/*
 * Structure for the custom type CloneDirectoryParams
 */
type CloneDirectoryParams struct {
	DestinationDirectoryName       *string `json:"destinationDirectoryName,omitempty" form:"destinationDirectoryName,omitempty"`             //Name of the new directory which will contain the clone contents.
	DestinationParentDirectoryPath *string `json:"destinationParentDirectoryPath,omitempty" form:"destinationParentDirectoryPath,omitempty"` //Specifies the path of the destination parent directory. The source dir
	SourceDirectoryPath            *string `json:"sourceDirectoryPath,omitempty" form:"sourceDirectoryPath,omitempty"`                       //Specifies the path of the source directory
}

/*
 * Structure for the custom type CloneRefreshRequest
 */
type CloneRefreshRequest struct {
	CloneTaskId      *int64                  `json:"cloneTaskId,omitempty" form:"cloneTaskId,omitempty"`           //Specifies the ID of the clone task. This is required to determine the
	ContinueOnError  *bool                   `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`   //Specifies if the Restore Task should continue when some operations on some
	Name             string                  `json:"name" form:"name"`                                             //Specifies the name of the Restore Task. This field must be set and
	NewParentId      *int64                  `json:"newParentId,omitempty" form:"newParentId,omitempty"`           //Specify a new registered parent Protection Source. If specified
	Objects          []*RestoreObjectDetails `json:"objects,omitempty" form:"objects,omitempty"`                   //Array of Objects.
	RefreshTimeSecs  *int64                  `json:"refreshTimeSecs,omitempty" form:"refreshTimeSecs,omitempty"`   //Specifies a point in time (unix epoch) to which the database needs to be
	SourceDatabaseId *int64                  `json:"sourceDatabaseId,omitempty" form:"sourceDatabaseId,omitempty"` //Specifies the ID of the source database in the backup job snapshot. This
	VlanParameters   *VlanParameters         `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"`     //Specifies VLAN parameters for the restore operation.
}

/*
 * Structure for the custom type CloneTaskInfo
 */
type CloneTaskInfo struct {
	Name   *string `json:"name,omitempty" form:"name,omitempty"`     //Name of the recovery task.
	TaskId *string `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
}

/*
 * Structure for the custom type CloneTaskRequest
 */
type CloneTaskRequest struct {
	CloneViewParameters  *CloneViewRequest        `json:"cloneViewParameters,omitempty" form:"cloneViewParameters,omitempty"`   //Specifies settings for cloning an existing View.
	ContinueOnError      *bool                    `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`           //Specifies if the Restore Task should continue when some operations on some
	GlacierRetrievalType GlacierRetrievalTypeEnum `json:"glacierRetrievalType,omitempty" form:"glacierRetrievalType,omitempty"` //Specifies the way data needs to be retrieved from the external target.
	HypervParameters     *HypervCloneParameters   `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"`         //Specifies information needed when cloning VMs in HyperV enviroment.
	Name                 string                   `json:"name" form:"name"`                                                     //Specifies the name of the Restore Task. This field must be set and
	NewParentId          *int64                   `json:"newParentId,omitempty" form:"newParentId,omitempty"`                   //Specify a new registered parent Protection Source. If specified
	Objects              []*RestoreObjectDetails  `json:"objects,omitempty" form:"objects,omitempty"`                           //Array of Objects.
	TargetViewName       *string                  `json:"targetViewName,omitempty" form:"targetViewName,omitempty"`             //Specifies the name of the View where the cloned VMs are stored.
	Type                 TypeCloneTaskRequestEnum `json:"type" form:"type"`                                                     //Specifies the type of Restore Task such as 'kCloneVMs' or 'kCloneView'.
	VlanParameters       *VlanParameters          `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"`             //Specifies VLAN parameters for the restore operation.
	VmwareParameters     *VmwareCloneParameters   `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"`         //Specifies the information required for recovering or cloning VmWare VMs.
}

/*
 * Structure for the custom type CloneViewRequest
 */
type CloneViewRequest struct {
	AccessSids                      *[]string                `json:"accessSids,omitempty" form:"accessSids,omitempty"`                                           //Array of Security Identifiers (SIDs)
	AntivirusScanConfig             *AntivirusScanConfig     `json:"antivirusScanConfig,omitempty" form:"antivirusScanConfig,omitempty"`                         //Specifies the antivirus scan config settings for this View.
	CloneViewName                   *string                  `json:"cloneViewName,omitempty" form:"cloneViewName,omitempty"`                                     //Specifies the name of the new View that is cloned from the source View.
	DataLockExpiryUsecs             *int64                   `json:"dataLockExpiryUsecs,omitempty" form:"dataLockExpiryUsecs,omitempty"`                         //DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If
	Description                     *string                  `json:"description,omitempty" form:"description,omitempty"`                                         //Specifies an optional text description about the View.
	EnableFilerAuditLogging         *bool                    `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"`                 //Specifies if Filer Audit Logging is enabled for this view.
	EnableMixedModePermissions      *bool                    `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"`           //If set, mixed mode (NFS and SMB) access is enabled for this view.
	EnableNfsViewDiscovery          *bool                    `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"`                   //If set, it enables discovery of view for NFS.
	EnableOfflineCaching            *bool                    `json:"enableOfflineCaching,omitempty" form:"enableOfflineCaching,omitempty"`                       //Specifies whether to enable offline file caching of the view.
	EnableSmbAccessBasedEnumeration *bool                    `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
	EnableSmbEncryption             *bool                    `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`                         //Specifies the SMB encryption for the View. If set, it enables the SMB
	EnableSmbViewDiscovery          *bool                    `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"`                   //If set, it enables discovery of view for SMB.
	EnforceSmbEncryption            *bool                    `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`                       //Specifies the SMB encryption for all the sessions for the View.
	FileExtensionFilter             *FileExtensionFilter     `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"`                         //TODO: Write general description for this field
	FileLockConfig                  *FileLevelDataLockConfig `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"`                                   //Specifies a config to lock files in a view - to protect from malicious or
	LogicalQuota                    *QuotaPolicy             `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"`                                       //Specifies an optional logical quota limit (in bytes) for the usage allowed
	NfsRootPermissions              *NfsRootPermissions      `json:"nfsRootPermissions,omitempty" form:"nfsRootPermissions,omitempty"`                           //Specifies the config of NFS root permission of a view file system.
	OverrideGlobalWhitelist         *bool                    `json:"overrideGlobalWhitelist,omitempty" form:"overrideGlobalWhitelist,omitempty"`                 //Specifies whether view level client subnet whitelist overrides cluster and
	ProtocolAccess                  ProtocolAccessEnum       `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"`                                   //Specifies the supported Protocols for the View.
	Qos                             *QoS                     `json:"qos,omitempty" form:"qos,omitempty"`                                                         //Specifies the Quality of Service (QoS) Policy for the View.
	SecurityMode                    SecurityModeEnum         `json:"securityMode,omitempty" form:"securityMode,omitempty"`                                       //Specifies the security mode used for this view.
	SharePermissions                []*SmbPermission         `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`                               //Specifies a list of share level permissions.
	SmbPermissionsInfo              *SmbPermissionsInfo      `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"`                           //Specifies information about SMB permissions.
	SourceViewName                  *string                  `json:"sourceViewName,omitempty" form:"sourceViewName,omitempty"`                                   //Specifies the name of the source View that will be cloned.
	StoragePolicyOverride           *StoragePolicyOverride   `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"`                     //Specifies if inline deduplication and compression settings inherited from
	SubnetWhitelist                 []*Subnet                `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`                                 //Array of Subnets.
	TenantId                        *string                  `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                               //Optional tenant id who has access to this View.
}

/*
 * Structure for the custom type CloseSmbFileOpenParameters
 */
type CloseSmbFileOpenParameters struct {
	FilePath *string `json:"filePath,omitempty" form:"filePath,omitempty"` //Specifies the filepath in the view relative to the root filesystem.
	OpenId   *int64  `json:"openId,omitempty" form:"openId,omitempty"`     //Specifies the id of the active open.
	ViewName *string `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the View in which to search. If a view name is not
}

/*
 * Structure for the custom type CloudDeployInfoProto
 */
type CloudDeployInfoProto struct {
	CloudDeployEntityVec          []*CloudDeployInfoProtoCloudDeployEntity `json:"cloudDeployEntityVec,omitempty" form:"cloudDeployEntityVec,omitempty"`                   //Contains the file paths and the information of the entities deployed to
	IsIncremental                 *bool                                    `json:"isIncremental,omitempty" form:"isIncremental,omitempty"`                                 //Whether this Cloud deploy info is for incremental cloudspin.
	RestoreInfo                   *RestoreInfoProto                        `json:"restoreInfo,omitempty" form:"restoreInfo,omitempty"`                                     //Each available extension is listed below along with the location of the
	TargetType                    *int64                                   `json:"targetType,omitempty" form:"targetType,omitempty"`                                       //Specifies the target type for the task. The field is only valid if the
	TotalBytesTransferredToSource *int64                                   `json:"totalBytesTransferredToSource,omitempty" form:"totalBytesTransferredToSource,omitempty"` //Total bytes transferred to source.
	Type                          *int64                                   `json:"type,omitempty" form:"type,omitempty"`                                                   //The type of environment this cloud deploy info pertains to.
}

/*
 * Structure for the custom type CloudDeployInfoProtoCloudDeployEntity
 */
type CloudDeployInfoProtoCloudDeployEntity struct {
	DeployedVmName               *string      `json:"deployedVmName,omitempty" form:"deployedVmName,omitempty"`                             //Optional name that should be used for deployed VM.
	Entity                       *EntityProto `json:"entity,omitempty" form:"entity,omitempty"`                                             //Specifies the attributes and the latest statistics about an entity.
	Error                        *ErrorProto  `json:"error,omitempty" form:"error,omitempty"`                                               //TODO: Write general description for this field
	PreviousRelativeCloneDirPath *string      `json:"previousRelativeCloneDirPath,omitempty" form:"previousRelativeCloneDirPath,omitempty"` //Directory where files of the entity's previous snapshot were cloned to.
	PreviousRelativeClonePaths   *[]string    `json:"previousRelativeClonePaths,omitempty" form:"previousRelativeClonePaths,omitempty"`     //All the paths that the entity's previous snapshot files were cloned to.
	ProgressMonitorTaskPath      *string      `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"`           //Progress monitor task path for this entity which is relative to the root
	PublicStatus                 *int64       `json:"publicStatus,omitempty" form:"publicStatus,omitempty"`                                 //Iris-facing task state. This field is stamped during the export.
	RelativeClonePaths           *[]string    `json:"relativeClonePaths,omitempty" form:"relativeClonePaths,omitempty"`                     //All the paths that the entity's files were cloned to. Each path is
	Status                       *int64       `json:"status,omitempty" form:"status,omitempty"`                                             //The status of the entity.
}

/*
 * Structure for the custom type CloudDeployTarget
 */
type CloudDeployTarget struct {
	DeployVmsToCloudParams *DeployVMsToCloudParams `json:"deployVmsToCloudParams,omitempty" form:"deployVmsToCloudParams,omitempty"` //Contains Cloud specific information needed to identify various resources
	TargetEntity           *EntityProto            `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`                     //Specifies the attributes and the latest statistics about an entity.
	Type                   *int64                  `json:"type,omitempty" form:"type,omitempty"`                                     //The type of the CloudDeploy target.
}

/*
 * Structure for the custom type CloudDeployTargetDetails
 */
type CloudDeployTargetDetails struct {
	AwsParams   *AwsParams                       `json:"awsParams,omitempty" form:"awsParams,omitempty"`     //Specifies various resources when converting and deploying a VM to AWS.
	AzureParams *AzureParams                     `json:"azureParams,omitempty" form:"azureParams,omitempty"` //Specifies various resources when converting and deploying a VM to Azure.
	GcpParams   *GcpParams                       `json:"gcpParams,omitempty" form:"gcpParams,omitempty"`     //Specifies various resources when converting and deploying a VM to GCP.
	Id          *int64                           `json:"id,omitempty" form:"id,omitempty"`                   //Entity corresponding to the cloud deploy target.
	Name        *string                          `json:"name,omitempty" form:"name,omitempty"`               //Specifies the inner object's name or a human-readable string made off the
	Type        TypeCloudDeployTargetDetailsEnum `json:"type,omitempty" form:"type,omitempty"`               //Specifies the type of the CloudDeploy target.
}

/*
 * Structure for the custom type CloudNetworkConfiguration
 */
type CloudNetworkConfiguration struct {
	ClusterGateway    *string   `json:"clusterGateway,omitempty" form:"clusterGateway,omitempty"`       //Specifies the default gateway IP address (or addresses) for the Cluster
	ClusterSubnetMask *string   `json:"clusterSubnetMask,omitempty" form:"clusterSubnetMask,omitempty"` //Specifies the subnet mask (or masks) of the Cluster network.
	DnsServers        *[]string `json:"dnsServers,omitempty" form:"dnsServers,omitempty"`               //Specifies the list of DNS Servers this cluster should be configured with.
	DomainNames       *[]string `json:"domainNames,omitempty" form:"domainNames,omitempty"`             //Specifies the list of domain names this cluster should be configured
	NtpServers        *[]string `json:"ntpServers,omitempty" form:"ntpServers,omitempty"`               //Specifies the list of NTP Servers this cluster should be configured with.
}

/*
 * Structure for the custom type CloudParameters
 */
type CloudParameters struct {
	FailoverToCloud *bool `json:"failoverToCloud,omitempty" form:"failoverToCloud,omitempty"` //Specifies whether the Protection Sources in this Protection Job
}

/*
 * Structure for the custom type Cluster
 */
type Cluster struct {
	AppsSettings                    *AppsConfig                   `json:"appsSettings,omitempty" form:"appsSettings,omitempty"`                                       //TODO: Write general description for this field
	AvailableMetadataSpace          *int64                        `json:"availableMetadataSpace,omitempty" form:"availableMetadataSpace,omitempty"`                   //Information about storage available for metadata
	BannerEnabled                   *bool                         `json:"bannerEnabled,omitempty" form:"bannerEnabled,omitempty"`                                     //Specifies whether UI banner is enabled on the cluster or not. When banner
	BondingMode                     BondingModeEnum               `json:"bondingMode,omitempty" form:"bondingMode,omitempty"`                                         //Specifies the bonding mode to use when bonding NICs to this Cluster.
	ClusterAuditLogConfig           *ClusterAuditLogConfiguration `json:"clusterAuditLogConfig,omitempty" form:"clusterAuditLogConfig,omitempty"`                     //Specifies the settings of the Cluster audit log configuration.
	ClusterSoftwareVersion          *string                       `json:"clusterSoftwareVersion,omitempty" form:"clusterSoftwareVersion,omitempty"`                   //Specifies the current release of the Cohesity software running on
	ClusterType                     ClusterTypeClusterEnum        `json:"clusterType,omitempty" form:"clusterType,omitempty"`                                         //Specifies the type of Cluster such as kPhysical.
	CreatedTimeMsecs                *int64                        `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`                               //Specifies the time when the Cohesity Cluster was created.
	CurrentOpScheduledTimeSecs      *int64                        `json:"currentOpScheduledTimeSecs,omitempty" form:"currentOpScheduledTimeSecs,omitempty"`           //Specifies the time scheduled by the Cohesity Cluster to
	CurrentOperation                CurrentOperationEnum          `json:"currentOperation,omitempty" form:"currentOperation,omitempty"`                               //Specifies the current Cluster-level operation in progress.
	CurrentTimeMsecs                *int64                        `json:"currentTimeMsecs,omitempty" form:"currentTimeMsecs,omitempty"`                               //Specifies the current system time on the Cohesity Cluster.
	DnsServerIps                    *[]string                     `json:"dnsServerIps,omitempty" form:"dnsServerIps,omitempty"`                                       //Array of IP Addresses of DNS Servers.
	DomainNames                     *[]string                     `json:"domainNames,omitempty" form:"domainNames,omitempty"`                                         //Array of Domain Names.
	EnableActiveMonitoring          *bool                         `json:"enableActiveMonitoring,omitempty" form:"enableActiveMonitoring,omitempty"`                   //Specifies if Cohesity can receive monitoring information from the
	EnableUpgradePkgPolling         *bool                         `json:"enableUpgradePkgPolling,omitempty" form:"enableUpgradePkgPolling,omitempty"`                 //If 'true', Cohesity's upgrade server is polled for new releases.
	EncryptionEnabled               *bool                         `json:"encryptionEnabled,omitempty" form:"encryptionEnabled,omitempty"`                             //If 'true', the entire Cohesity Cluster is encrypted including all View
	EncryptionKeyRotationPeriodSecs *int64                        `json:"encryptionKeyRotationPeriodSecs,omitempty" form:"encryptionKeyRotationPeriodSecs,omitempty"` //Specifies the period of time (in seconds) when encryption keys are rotated.
	EulaConfig                      *EulaConfig                   `json:"eulaConfig,omitempty" form:"eulaConfig,omitempty"`                                           //Specifies the End User License Agreement (EULA) acceptance information.
	FilerAuditLogConfig             *FilerAuditLogConfiguration   `json:"filerAuditLogConfig,omitempty" form:"filerAuditLogConfig,omitempty"`                         //Specifies the settings of the filer audit log configuration.
	FipsModeEnabled                 *bool                         `json:"fipsModeEnabled,omitempty" form:"fipsModeEnabled,omitempty"`                                 //Specifies if the Cohesity Cluster should operate in the FIPS mode,
	Gateway                         *string                       `json:"gateway,omitempty" form:"gateway,omitempty"`                                                 //Specifies the gateway IP address.
	GoogleAnalyticsEnabled          *bool                         `json:"googleAnalyticsEnabled,omitempty" form:"googleAnalyticsEnabled,omitempty"`                   //Specifies whether Google Analytics is enabled.
	HardwareInfo                    *ClusterHardwareInfo          `json:"hardwareInfo,omitempty" form:"hardwareInfo,omitempty"`                                       //Specifies a hardware type for motherboard of the Nodes
	Id                              *int64                        `json:"id,omitempty" form:"id,omitempty"`                                                           //Specifies the unique id of Cohesity Cluster.
	IncarnationId                   *int64                        `json:"incarnationId,omitempty" form:"incarnationId,omitempty"`                                     //Specifies the unique incarnation id of the Cohesity Cluster.
	IsDocumentationLocal            *bool                         `json:"isDocumentationLocal,omitempty" form:"isDocumentationLocal,omitempty"`                       //Specifies what version of the documentation is used.
	LanguageLocale                  *string                       `json:"languageLocale,omitempty" form:"languageLocale,omitempty"`                                   //Specifies the language and locale for this Cohesity Cluster.
	LicenseState                    *LicenseState                 `json:"licenseState,omitempty" form:"licenseState,omitempty"`                                       //Specifies the Licensing State information.
	LocalAuthDomainName             *string                       `json:"localAuthDomainName,omitempty" form:"localAuthDomainName,omitempty"`                         //Domain name for SMB local authentication.
	LocalGroupsEnabled              *bool                         `json:"localGroupsEnabled,omitempty" form:"localGroupsEnabled,omitempty"`                           //Specifies whether to enable local groups on cluster. Once it is enabled,
	MetadataFaultToleranceFactor    *int64                        `json:"metadataFaultToleranceFactor,omitempty" form:"metadataFaultToleranceFactor,omitempty"`       //Specifies metadata fault tolerance setting for the cluster. This denotes
	Mtu                             *int64                        `json:"mtu,omitempty" form:"mtu,omitempty"`                                                         //Specifies the Maxium Transmission Unit (MTU) in bytes of
	MultiTenancyEnabled             *bool                         `json:"multiTenancyEnabled,omitempty" form:"multiTenancyEnabled,omitempty"`                         //Specifies if multi tenancy is enabled in the cluster. Authentication &
	Name                            *string                       `json:"name,omitempty" form:"name,omitempty"`                                                       //Specifies the name of the Cohesity Cluster.
	NodeCount                       *int64                        `json:"nodeCount,omitempty" form:"nodeCount,omitempty"`                                             //Specifies the number of Nodes in the Cohesity Cluster.
	NtpSettings                     *NtpSettingsConfig            `json:"ntpSettings,omitempty" form:"ntpSettings,omitempty"`                                         //TODO: Write general description for this field
	ProxyVMSubnet                   *string                       `json:"proxyVMSubnet,omitempty" form:"proxyVMSubnet,omitempty"`                                     //The subnet reserved for ProxyVM
	ReverseTunnelEnabled            *bool                         `json:"reverseTunnelEnabled,omitempty" form:"reverseTunnelEnabled,omitempty"`                       //If 'true', Cohesity's Remote Tunnel is enabled.
	ReverseTunnelEndTimeMsecs       *int64                        `json:"reverseTunnelEndTimeMsecs,omitempty" form:"reverseTunnelEndTimeMsecs,omitempty"`             //ReverseTunnelEndTimeMsecs specifies the end time in milliseconds since
	SchemaInfoList                  []*SchemaInfo                 `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"`                                   //Specifies the time series schema info of the cluster.
	SmbAdDisabled                   *bool                         `json:"smbAdDisabled,omitempty" form:"smbAdDisabled,omitempty"`                                     //Specifies if Active Directory should be disabled for authentication of SMB
	Stats                           *ClusterStats                 `json:"stats,omitempty" form:"stats,omitempty"`                                                     //Specifies statistics about this Cohesity Cluster.
	StigMode                        *bool                         `json:"stigMode,omitempty" form:"stigMode,omitempty"`                                               //Specifies if STIG mode is enabled or not.
	SupportedConfig                 *SupportedConfig              `json:"supportedConfig,omitempty" form:"supportedConfig,omitempty"`                                 //Lists the supported Erasure Coding options for the number of
	SyslogServers                   []*SyslogServer               `json:"syslogServers,omitempty" form:"syslogServers,omitempty"`                                     //Array of Syslog Servers.
	TargetSoftwareVersion           *string                       `json:"targetSoftwareVersion,omitempty" form:"targetSoftwareVersion,omitempty"`                     //Specifies the Cohesity release that this Cluster is being upgraded to
	TenantViewboxSharingEnabled     *bool                         `json:"tenantViewboxSharingEnabled,omitempty" form:"tenantViewboxSharingEnabled,omitempty"`         //In case multi tenancy is enabled, this flag controls whether multiple
	Timezone                        *string                       `json:"timezone,omitempty" form:"timezone,omitempty"`                                               //Specifies the timezone to use for showing time in emails, reports,
	TurboMode                       *bool                         `json:"turboMode,omitempty" form:"turboMode,omitempty"`                                             //Specifies if the cluster is in Turbo mode.
	UsedMetadataSpacePct            *float64                      `json:"usedMetadataSpacePct,omitempty" form:"usedMetadataSpacePct,omitempty"`                       //UsedMetadataSpacePct measures the percentage about storage used for
}

/*
 * Structure for the custom type ClusterAuditLog
 */
type ClusterAuditLog struct {
	Action         *string `json:"action,omitempty" form:"action,omitempty"`                 //Specifies the action that caused the log to be generated.
	Details        *string `json:"details,omitempty" form:"details,omitempty"`               //Specifies more information about the action.
	Domain         *string `json:"domain,omitempty" form:"domain,omitempty"`                 //Specifies the domain of the user who caused the action
	EntityId       *string `json:"entityId,omitempty" form:"entityId,omitempty"`             //Specifies the id of the entity (object) that the action is invoked on.
	EntityName     *string `json:"entityName,omitempty" form:"entityName,omitempty"`         //Specifies the entity (object) name that the action is invoked on.
	EntityType     *string `json:"entityType,omitempty" form:"entityType,omitempty"`         //Specifies the type of the entity (object) that the action is invoked on.
	HumanTimestamp *string `json:"humanTimestamp,omitempty" form:"humanTimestamp,omitempty"` //Specifies the time when the log was generated.
	Impersonation  *bool   `json:"impersonation,omitempty" form:"impersonation,omitempty"`   //Specifies if the log was generated during impersonation.
	NewRecord      *string `json:"newRecord,omitempty" form:"newRecord,omitempty"`           //Specifies the record after the action is invoked.
	OriginalTenant *Tenant `json:"originalTenant,omitempty" form:"originalTenant,omitempty"` //Specifies details about a tenant.
	PreviousRecord *string `json:"previousRecord,omitempty" form:"previousRecord,omitempty"` //Specifies the record before the action is invoked.
	Tenant         *Tenant `json:"tenant,omitempty" form:"tenant,omitempty"`                 //Specifies details about a tenant.
	TimestampUsecs *int64  `json:"timestampUsecs,omitempty" form:"timestampUsecs,omitempty"` //Specifies the time when the log was generated.
	UserName       *string `json:"userName,omitempty" form:"userName,omitempty"`             //Specifies the user who caused the action that generated the log.
}

/*
 * Structure for the custom type ClusterAuditLogConfiguration
 */
type ClusterAuditLogConfiguration struct {
	Enabled             bool  `json:"enabled" form:"enabled"`                         //Specifies if the Cluster audit logging is enabled on the
	RetentionPeriodDays int64 `json:"retentionPeriodDays" form:"retentionPeriodDays"` //Specifies the number of days to keep (retain) the Cluster audit logs.
}

/*
 * Structure for the custom type ClusterAuditLogsSearchResult
 */
type ClusterAuditLogsSearchResult struct {
	ClusterAuditLogs []*ClusterAuditLog `json:"clusterAuditLogs,omitempty" form:"clusterAuditLogs,omitempty"` //Array of Cluster Audit Logs.
	TotalCount       *int64             `json:"totalCount,omitempty" form:"totalCount,omitempty"`             //Specifies the total number of logs that match the specified
}

/*
 * Structure for the custom type ClusterConfigProtoQoSMapping
 */
type ClusterConfigProtoQoSMapping struct {
	PrincipalId *int64                                  `json:"principalId,omitempty" form:"principalId,omitempty"` //Principal id of the QoS principal to which qos_context maps to.
	QosContext  *ClusterConfigProtoQoSMappingQoSContext `json:"qosContext,omitempty" form:"qosContext,omitempty"`   //QoSContext captures the properties that are relevant for QoS in a
}

/*
 * Structure for the custom type ClusterConfigProtoQoSMappingQoSContext
 */
type ClusterConfigProtoQoSMappingQoSContext struct {
	Priority  *int64 `json:"priority,omitempty" form:"priority,omitempty"`   //Priority of a request.
	Type      *int64 `json:"type,omitempty" form:"type,omitempty"`           //TODO: Write general description for this field
	ViewBoxId *int64 `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //View box id of a request.
	ViewId    *int64 `json:"viewId,omitempty" form:"viewId,omitempty"`       //View id of a request.
}

/*
 * Structure for the custom type ClusterConfigProtoSID
 */
type ClusterConfigProtoSID struct {
	IdentifierAuthority *[]int64 `json:"identifierAuthority,omitempty" form:"identifierAuthority,omitempty"` //The authority under which the SID was created. This is always 6 bytes
	RevisionLevel       *int64   `json:"revisionLevel,omitempty" form:"revisionLevel,omitempty"`             //The revision level of the SID.
	SubAuthority        *[]int64 `json:"subAuthority,omitempty" form:"subAuthority,omitempty"`               //List of ids relative to the identifier_authority that uniquely
}

/*
 * Structure for the custom type ClusterConfigProtoStoragePolicyOverride
 */
type ClusterConfigProtoStoragePolicyOverride struct {
	DisableInlineDedupAndCompression *bool `json:"disableInlineDedupAndCompression,omitempty" form:"disableInlineDedupAndCompression,omitempty"` //If this is set to true, we will not do inline dedup and compression even
}

/*
 * Structure for the custom type ClusterConfigProtoSubnet
 */
type ClusterConfigProtoSubnet struct {
	Component     *int64  `json:"component,omitempty" form:"component,omitempty"`         //The component that has claimed this subnet.
	Description   *string `json:"description,omitempty" form:"description,omitempty"`     //Description of the subnet.
	Gateway       *string `json:"gateway,omitempty" form:"gateway,omitempty"`             //Gateway for the subnet.
	Id            *int64  `json:"id,omitempty" form:"id,omitempty"`                       //ID for this subnet.
	Ip            *string `json:"ip,omitempty" form:"ip,omitempty"`                       //ip is subnet IP address given either in v4 or v6. Netmask is
	NetmaskBits   *int64  `json:"netmaskBits,omitempty" form:"netmaskBits,omitempty"`     //TODO: Write general description for this field
	NetmaskIp4    *string `json:"netmaskIp4,omitempty" form:"netmaskIp4,omitempty"`       //TODO: Write general description for this field
	NfsAccess     *int64  `json:"nfsAccess,omitempty" form:"nfsAccess,omitempty"`         //Whether clients from this subnet can mount using NFS protocol.
	NfsRootSquash *bool   `json:"nfsRootSquash,omitempty" form:"nfsRootSquash,omitempty"` //Whether clients from this subnet can mount as root on NFS.
	SmbAccess     *int64  `json:"smbAccess,omitempty" form:"smbAccess,omitempty"`         //Whether clients from this subnet can mount using SMB protocol.
}

/*
 * Structure for the custom type ClusterCreationProgressResult
 */
type ClusterCreationProgressResult struct {
	CompletionPercentage *int64    `json:"completionPercentage,omitempty" form:"completionPercentage,omitempty"` //Specifies an approximate completion percentage for the Cluster creation
	ErrorMessage         *string   `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`                 //Specifies a description of an error if any error was encountered during
	Events               *[]string `json:"events,omitempty" form:"events,omitempty"`                             //Specifies a list of events that took place during Cluster creation.
	InProgress           *bool     `json:"inProgress,omitempty" form:"inProgress,omitempty"`                     //Specifies whether or not the Cluster is still in the process of being
	Message              *string   `json:"message,omitempty" form:"message,omitempty"`                           //Specifies an optional message describing the current state of the
	SecondsRemaining     *int64    `json:"secondsRemaining,omitempty" form:"secondsRemaining,omitempty"`         //Specifies an estimated number of seconds until the Cluster creation
	WarningsFound        *bool     `json:"warningsFound,omitempty" form:"warningsFound,omitempty"`               //Specifies whether or not any warnings were encountered during Cluster
}

/*
 * Structure for the custom type ClusterHardwareInfo
 */
type ClusterHardwareInfo struct {
	HardwareModels  *[]string `json:"hardwareModels,omitempty" form:"hardwareModels,omitempty"`   //TODO: Write general description for this field
	HardwareVendors *[]string `json:"hardwareVendors,omitempty" form:"hardwareVendors,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ClusterIdentifier
 */
type ClusterIdentifier struct {
	ClusterId            *int64 `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //Specifies the cluster id of the cluster.
	ClusterIncarnationId *int64 `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the cluster incarnation id.
}

/*
 * Structure for the custom type ClusterNetworkingEndpoint
 */
type ClusterNetworkingEndpoint struct {
	Fqdn     *string `json:"fqdn,omitempty" form:"fqdn,omitempty"`         //The Fully Qualified Domain Name.
	Ipv4Addr *string `json:"ipv4Addr,omitempty" form:"ipv4Addr,omitempty"` //The IPv4 address.
	Ipv6Addr *string `json:"ipv6Addr,omitempty" form:"ipv6Addr,omitempty"` //The IPv6 address.
}

/*
 * Structure for the custom type ClusterNetworkingResourceInformation
 */
type ClusterNetworkingResourceInformation struct {
	Endpoints []*ClusterNetworkingEndpoint `json:"endpoints,omitempty" form:"endpoints,omitempty"` //The endpoints by which the resource is accessible.
	Type      *string                      `json:"type,omitempty" form:"type,omitempty"`           //The type of the resource.
}

/*
 * Structure for the custom type ClusterPartition
 */
type ClusterPartition struct {
	HostName *string   `json:"hostName,omitempty" form:"hostName,omitempty"` //Specifies that hostname that resolves to one or more Virtual IP
	Id       *int64    `json:"id,omitempty" form:"id,omitempty"`             //Specifies a unique identifier for the Cluster Partition.
	Name     *string   `json:"name,omitempty" form:"name,omitempty"`         //Specifies the name of the Cluster Partition.
	NodeIds  *[]int64  `json:"nodeIds,omitempty" form:"nodeIds,omitempty"`   //Array of Node Ids.
	Vips     *[]string `json:"vips,omitempty" form:"vips,omitempty"`         //Array of VIPs.
	VlanIps  *[]string `json:"vlanIps,omitempty" form:"vlanIps,omitempty"`   //Array of VLAN IPs.
	Vlans    []*Vlan   `json:"vlans,omitempty" form:"vlans,omitempty"`       //Array of VLANs.
}

/*
 * Structure for the custom type ClusterPublicKeys
 */
type ClusterPublicKeys struct {
	SshPublicKey *string `json:"sshPublicKey,omitempty" form:"sshPublicKey,omitempty"` //Specifies the SSH public key used to login to Cluster nodes.
}

/*
 * Structure for the custom type ClusterStats
 */
type ClusterStats struct {
	CloudUsagePerfStats *UsageAndPerformanceStats `json:"cloudUsagePerfStats,omitempty" form:"cloudUsagePerfStats,omitempty"` //Provides usage and performance statistics for the remote data stored on
	DataReductionRatio  *float64                  `json:"dataReductionRatio,omitempty" form:"dataReductionRatio,omitempty"`   //Specifies the ratio of logical bytes (not reduced by
	DataUsageStats      *DataUsageStats           `json:"dataUsageStats,omitempty" form:"dataUsageStats,omitempty"`           //Specifies the data usage metric of the data stored on the Cohesity
	Id                  *int64                    `json:"id,omitempty" form:"id,omitempty"`                                   //Specifies the id of the Cohesity Cluster.
	LocalUsagePerfStats *UsageAndPerformanceStats `json:"localUsagePerfStats,omitempty" form:"localUsagePerfStats,omitempty"` //Provides usage and performance statistics for local data stored directly
	LogicalStats        *LogicalStats             `json:"logicalStats,omitempty" form:"logicalStats,omitempty"`               //Specifies the total logical data size of all the local and
	UsagePerfStats      *UsageAndPerformanceStats `json:"usagePerfStats,omitempty" form:"usagePerfStats,omitempty"`           //Provides usage and performance statistics about the local data
}

/*
 * Structure for the custom type CompareADObjectsResultADAttribute
 */
type CompareADObjectsResultADAttribute struct {
	AttrFlags   *int64                                  `json:"attrFlags,omitempty" form:"attrFlags,omitempty"`     //Object result flags of type ADAttributeFlags.
	DestValue   *CompareADObjectsResultADAttributeValue `json:"destValue,omitempty" form:"destValue,omitempty"`     //TODO: Write general description for this field
	LdapName    *string                                 `json:"ldapName,omitempty" form:"ldapName,omitempty"`       //LDAP attribute name.
	SameValue   *CompareADObjectsResultADAttributeValue `json:"sameValue,omitempty" form:"sameValue,omitempty"`     //TODO: Write general description for this field
	SourceValue *CompareADObjectsResultADAttributeValue `json:"sourceValue,omitempty" form:"sourceValue,omitempty"` //TODO: Write general description for this field
	Status      *ErrorProto                             `json:"status,omitempty" form:"status,omitempty"`           //TODO: Write general description for this field
}

/*
 * Structure for the custom type CompareADObjectsResultADAttributeValue
 */
type CompareADObjectsResultADAttributeValue struct {
	ValueFlags *int64    `json:"valueFlags,omitempty" form:"valueFlags,omitempty"` //Object result flags of type ADAttributeValueFlags.
	ValueVec   *[]string `json:"valueVec,omitempty" form:"valueVec,omitempty"`     //String representation of attribute value. For single valued property,
}

/*
 * Structure for the custom type CompareADObjectsResultADObject
 */
type CompareADObjectsResultADObject struct {
	AttributeVec      []*CompareADObjectsResultADAttribute `json:"attributeVec,omitempty" form:"attributeVec,omitempty"`           //Array of AD attributes of AD object. This will contain distinct
	DestGuid          *string                              `json:"destGuid,omitempty" form:"destGuid,omitempty"`                   //Object guid from dest_server. If empty, compare could not find an AD
	DestPropCount     *int64                               `json:"destPropCount,omitempty" form:"destPropCount,omitempty"`         //Number of attributes in destination object including system properties
	ExcludedPropCount *int64                               `json:"excludedPropCount,omitempty" form:"excludedPropCount,omitempty"` //Number of attributes not compared due to
	MismatchPropCount *int64                               `json:"mismatchPropCount,omitempty" form:"mismatchPropCount,omitempty"` //Number of AD attributes compared based on 'ADCompareOptionFlagsType'
	ObjectFlags       *int64                               `json:"objectFlags,omitempty" form:"objectFlags,omitempty"`             //Object result flags of type ADObjectFlags.
	SourceGuid        *string                              `json:"sourceGuid,omitempty" form:"sourceGuid,omitempty"`               //Object guid from $SourceServer. Guid string with or without '{}' braces.
	SourcePropCount   *int64                               `json:"sourcePropCount,omitempty" form:"sourcePropCount,omitempty"`     //Number of attributes in source object including system properties
	Status            *ErrorProto                          `json:"status,omitempty" form:"status,omitempty"`                       //TODO: Write general description for this field
}

/*
 * Structure for the custom type CompareAdObjectsRequest
 */
type CompareAdObjectsRequest struct {
	RestoreTaskId             int64       `json:"RestoreTaskId" form:"RestoreTaskId"`                                             //Specifies the Restore Task Id corresponding to which we need to compare
	AllowEmptyDestGuids       *bool       `json:"allowEmptyDestGuids,omitempty" form:"allowEmptyDestGuids,omitempty"`             //Specifies the option to get object attributes from Snapshot AD when
	ExcludeSysAttributes      *bool       `json:"excludeSysAttributes,omitempty" form:"excludeSysAttributes,omitempty"`           //Specifies the option to exclude AD system attributes when comparing two AD
	FilterNullValueAttributes *bool       `json:"filterNullValueAttributes,omitempty" form:"filterNullValueAttributes,omitempty"` //Specifies the option to not return attributes where source and destination
	FilterSameValueAttributes *bool       `json:"filterSameValueAttributes,omitempty" form:"filterSameValueAttributes,omitempty"` //Specifies the option to not return attributes where source and
	GuidPairs                 []*GuidPair `json:"guidPairs" form:"guidPairs"`                                                     //Specifies the GuidPair of the AD Objects which we want to compare
	QuickCompare              *bool       `json:"quickCompare,omitempty" form:"quickCompare,omitempty"`                           //Specifies the option to do quick compare of specified guid between Snapshot
}

/*
 * Structure for the custom type ComparedADObject
 */
type ComparedADObject struct {
	AdAttributes      []*AdAttribute      `json:"adAttributes,omitempty" form:"adAttributes,omitempty"`           //Specifies the list of AD attributes for the AD object.
	AdObjectFlags     *[]AdObjectFlagEnum `json:"adObjectFlags,omitempty" form:"adObjectFlags,omitempty"`         //Specifies the flags related to this AD Object.
	DestinationGuid   *string             `json:"destinationGuid,omitempty" form:"destinationGuid,omitempty"`     //Specifies the guid of the object in the Production AD which is equivalent
	ErrorMessage      *string             `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`           //Specifies the error message while fetching the AD object.
	MismatchAttrCount *int64              `json:"mismatchAttrCount,omitempty" form:"mismatchAttrCount,omitempty"` //Specifies the number of attributes of AD Object mismatched on the
	SourceGuid        *string             `json:"sourceGuid,omitempty" form:"sourceGuid,omitempty"`               //Specifies the guid of the AD object in the Snapshot AD.
}

/*
 * Structure for the custom type ConnectorParameters
 */
type ConnectorParameters struct {
	Endpoint    *string                            `json:"endpoint,omitempty" form:"endpoint,omitempty"`       //Specify an IP address or URL of the environment.
	Environment EnvironmentConnectorParametersEnum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment like VMware, SQL, where the
	Id          *int64                             `json:"id,omitempty" form:"id,omitempty"`                   //Specifies a Unique id that is generated when the Source is registered.
	Version     *int64                             `json:"version,omitempty" form:"version,omitempty"`         //Version is updated each time the connector parameters are updated.
}

/*
 * Structure for the custom type ConnectorParams
 */
type ConnectorParams struct {
	AgentEndpoint *string      `json:"agentEndpoint,omitempty" form:"agentEndpoint,omitempty"` //For some of the environments connection to endpoint is done through an
	AgentPort     *int64       `json:"agentPort,omitempty" form:"agentPort,omitempty"`         //Optional agent port to use when connecting to the server.
	Credentials   *Credentials `json:"credentials,omitempty" form:"credentials,omitempty"`     //Specifies credentials to access a target source.
	Endpoint      *string      `json:"endpoint,omitempty" form:"endpoint,omitempty"`           //The endpoint URL of the environment (such as the address of the vCenter
	Entity        *EntityProto `json:"entity,omitempty" form:"entity,omitempty"`               //Specifies the attributes and the latest statistics about an entity.
	HostType      *int64       `json:"hostType,omitempty" form:"hostType,omitempty"`           //The host environment type. This is set for kPhysical type environment.
	Id            *int64       `json:"id,omitempty" form:"id,omitempty"`                       //A unique id associated with this connector params. This is a convenience
	Port          *int64       `json:"port,omitempty" form:"port,omitempty"`                   //Optional port to use when connecting to the server.
	TenantId      *string      `json:"tenantId,omitempty" form:"tenantId,omitempty"`           //The tenant_id for the environment. This is used to remotely access
	Type          *int64       `json:"type,omitempty" form:"type,omitempty"`                   //The type of environment to connect to.
	Version       *int64       `json:"version,omitempty" form:"version,omitempty"`             //A version that is associated with the params. This is updated anytime
}

/*
 * Structure for the custom type Consumer
 */
type Consumer struct {
	Id   *int64           `json:"id,omitempty" form:"id,omitempty"`     //Specifies the id of the consumer.
	Name *string          `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the consumer.
	Type TypeConsumerEnum `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the consumer.
}

/*
 * Structure for the custom type ConsumerStats
 */
type ConsumerStats struct {
	ConsumerType        ConsumerTypeEnum   `json:"consumerType,omitempty" form:"consumerType,omitempty"`               //Specifies the type of the consumer.
	GroupList           []*StatsGroup      `json:"groupList,omitempty" form:"groupList,omitempty"`                     //Specifies a list of groups associated to this consumer.
	Id                  *int64             `json:"id,omitempty" form:"id,omitempty"`                                   //Specifies the id of the consumer.
	Name                *string            `json:"name,omitempty" form:"name,omitempty"`                               //Specifies the name of the consumer.
	QuotaHardLimitBytes *int64             `json:"quotaHardLimitBytes,omitempty" form:"quotaHardLimitBytes,omitempty"` //Specifies the hard limit of logical quota of the consumer. This field
	SchemaInfoList      []*UsageSchemaInfo `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"`           //Specifies a list of schemaInfos of the consumer.
	Stats               *DataUsageStats    `json:"stats,omitempty" form:"stats,omitempty"`                             //Specifies the data usage metric of the data stored on the Cohesity
}

/*
 * Structure for the custom type ContinuousSchedule
 */
type ContinuousSchedule struct {
	BackupIntervalMins *int64 `json:"backupIntervalMins,omitempty" form:"backupIntervalMins,omitempty"` //If specified, this field defines the time interval in minutes when
}

/*
 * Structure for the custom type CopyRun
 */
type CopyRun struct {
	CopySnapshotTasks   []*CopySnapshotTaskStatus `json:"copySnapshotTasks,omitempty" form:"copySnapshotTasks,omitempty"`     //Specifies the status information of each task that copies the snapshot
	Error               *string                   `json:"error,omitempty" form:"error,omitempty"`                             //Specifies if an error occurred (if any) while running this task.
	ExpiryTimeUsecs     *int64                    `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"`         //Specifies expiry time of the copies of the snapshots in this Protection
	HoldForLegalPurpose *bool                     `json:"holdForLegalPurpose,omitempty" form:"holdForLegalPurpose,omitempty"` //Specifies whether legal hold is enabled on this run. It is true if the
	LegalHoldings       []*LegalHoldings          `json:"legalHoldings,omitempty" form:"legalHoldings,omitempty"`             //Specifies the list of Protection Source Ids and the legal hold status.
	RunStartTimeUsecs   *int64                    `json:"runStartTimeUsecs,omitempty" form:"runStartTimeUsecs,omitempty"`     //Specifies start time of the copy run.
	Stats               *CopyRunStats             `json:"stats,omitempty" form:"stats,omitempty"`                             //Stats for one copy task or aggregated stats of a Copy Run in a
	Status              StatusCopyRunEnum         `json:"status,omitempty" form:"status,omitempty"`                           //Specifies the aggregated status of copy tasks such as 'kRunning',
	Target              *SnapshotTargetSettings   `json:"target,omitempty" form:"target,omitempty"`                           //Specifies settings about a target where a copied Snapshot is stored.
	TaskUid             *UniversalId              `json:"taskUid,omitempty" form:"taskUid,omitempty"`                         //Specifies a globally unique id of the copy task.
	UserActionMessage   *string                   `json:"userActionMessage,omitempty" form:"userActionMessage,omitempty"`     //Specifies a message to the user if any manual intervention is needed to
}

/*
 * Structure for the custom type CopyRunStats
 */
type CopyRunStats struct {
	EndTimeUsecs             *int64 `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                         //Specifies the time when this replication ended. If not set, then the
	IsIncremental            *bool  `json:"isIncremental,omitempty" form:"isIncremental,omitempty"`                       //Specifies whether this archival is incremental for archival targets.
	LogicalBytesTransferred  *int64 `json:"logicalBytesTransferred,omitempty" form:"logicalBytesTransferred,omitempty"`   //Specifies the number of logical bytes transferred for this replication
	LogicalSizeBytes         *int64 `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"`                 //Specifies the total amount of logical data to be transferred for this
	LogicalTransferRateBps   *int64 `json:"logicalTransferRateBps,omitempty" form:"logicalTransferRateBps,omitempty"`     //Specifies average logical bytes transfer rate in bytes per second for
	PhysicalBytesTransferred *int64 `json:"physicalBytesTransferred,omitempty" form:"physicalBytesTransferred,omitempty"` //Specifies the number of physical bytes sent over the wire for
	StartTimeUsecs           *int64 `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                     //Specifies the time when this replication was started. If not set, then
}

/*
 * Structure for the custom type CopySnapshotTaskStatus
 */
type CopySnapshotTaskStatus struct {
	Error              *string                          `json:"error,omitempty" form:"error,omitempty"`                           //Specifies if an error occurred (if any) while running this task.
	Source             *ProtectionSource                `json:"source,omitempty" form:"source,omitempty"`                         //Specifies a generic structure that represents a node
	Stats              *CopyRunStats                    `json:"stats,omitempty" form:"stats,omitempty"`                           //Stats for one copy task or aggregated stats of a Copy Run in a
	Status             StatusCopySnapshotTaskStatusEnum `json:"status,omitempty" form:"status,omitempty"`                         //Specifies the status of the source object being protected.
	TaskEndTimeUsecs   *int64                           `json:"taskEndTimeUsecs,omitempty" form:"taskEndTimeUsecs,omitempty"`     //Specifies the end time of the copy task. The end time
	TaskStartTimeUsecs *int64                           `json:"taskStartTimeUsecs,omitempty" form:"taskStartTimeUsecs,omitempty"` //Specifies the start time of the copy task. The start time
}

/*
 * Structure for the custom type CountByTier
 */
type CountByTier struct {
	DiskCount   *int64          `json:"diskCount,omitempty" form:"diskCount,omitempty"`     //DiskCount is the disk number of the storage tier.
	StorageTier StorageTierEnum `json:"storageTier,omitempty" form:"storageTier,omitempty"` //StorageTier is the type of StorageTier.
}

/*
 * Structure for the custom type CreateActiveDirectoryEntryParams
 */
type CreateActiveDirectoryEntryParams struct {
	DomainName                 *string                      `json:"domainName,omitempty" form:"domainName,omitempty"`                                 //Specifies the fully qualified domain name (FQDN) of an Active Directory.
	FallbackUserIdMappingInfo  *UserIdMapping               `json:"fallbackUserIdMappingInfo,omitempty" form:"fallbackUserIdMappingInfo,omitempty"`   //Specifies how the Unix and Windows users are mapped in an Active Directory.
	IgnoredTrustedDomains      *[]string                    `json:"ignoredTrustedDomains,omitempty" form:"ignoredTrustedDomains,omitempty"`           //Specifies the list of trusted domains that were set by the user to be
	LdapProviderId             *int64                       `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"`                         //Specifies the LDAP provider id which is map to this Active Directory
	MachineAccounts            *[]string                    `json:"machineAccounts,omitempty" form:"machineAccounts,omitempty"`                       //Array of Machine Accounts.
	OuName                     *string                      `json:"ouName,omitempty" form:"ouName,omitempty"`                                         //Specifies an optional Organizational Unit name.
	OverwriteExistingAccounts  *bool                        `json:"overwriteExistingAccounts,omitempty" form:"overwriteExistingAccounts,omitempty"`   //Specifies whether the specified machine accounts should overwrite the
	Password                   *string                      `json:"password,omitempty" form:"password,omitempty"`                                     //Specifies the password for the specified userName.
	PreferredDomainControllers []*PreferredDomainController `json:"preferredDomainControllers,omitempty" form:"preferredDomainControllers,omitempty"` //Specifies Map of Active Directory domain names to its preferred domain
	TenantId                   *string                      `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                     //Specifies the unique id of the tenant.
	TrustedDomains             *[]string                    `json:"trustedDomains,omitempty" form:"trustedDomains,omitempty"`                         //Specifies the trusted domains of the Active Directory domain.
	TrustedDomainsEnabled      *bool                        `json:"trustedDomainsEnabled,omitempty" form:"trustedDomainsEnabled,omitempty"`           //Specifies whether Trusted Domain discovery is disabled.
	UnixRootSid                *string                      `json:"unixRootSid,omitempty" form:"unixRootSid,omitempty"`                               //Specifies the SID of the Active Directory domain user to be mapped to
	UserIdMappingInfo          *UserIdMapping               `json:"userIdMappingInfo,omitempty" form:"userIdMappingInfo,omitempty"`                   //Specifies how the Unix and Windows users are mapped in an Active Directory.
	UserName                   *string                      `json:"userName,omitempty" form:"userName,omitempty"`                                     //Specifies a userName that has administrative privileges in the domain.
	Workgroup                  *string                      `json:"workgroup,omitempty" form:"workgroup,omitempty"`                                   //Specifies an optional Workgroup name.
}

/*
 * Structure for the custom type CreateBondParameters
 */
type CreateBondParameters struct {
	BondingMode BondingModeCreateBondParametersEnum `json:"bondingMode,omitempty" form:"bondingMode,omitempty"` //Specifies the bonding mode to use for this bond. If not specified,
	Name        string                              `json:"name" form:"name"`                                   //Specifies a unique name to identify the bond being created.
	Slaves      []string                            `json:"slaves" form:"slaves"`                               //Specifies the names of the slaves of this bond.
}

/*
 * Structure for the custom type CreateBondResult
 */
type CreateBondResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies a message describing the result of the operation.
}

/*
 * Structure for the custom type CreateCloudClusterParameters
 */
type CreateCloudClusterParameters struct {
	ClusterName            string                    `json:"clusterName" form:"clusterName"`                                           //Specifies the name of the new Cluster.
	EncryptionConfig       *EncryptionConfiguration  `json:"encryptionConfig,omitempty" form:"encryptionConfig,omitempty"`             //Specifies the parameters the user wants to use when configuring encryption
	MetadataFaultTolerance *int64                    `json:"metadataFaultTolerance,omitempty" form:"metadataFaultTolerance,omitempty"` //Specifies the metadata fault tolerance.
	NetworkConfig          CloudNetworkConfiguration `json:"networkConfig" form:"networkConfig"`                                       //Specifies all of the parameters needed for network configuration of
	NodeIps                []string                  `json:"nodeIps" form:"nodeIps"`                                                   //Specifies the configuration for the nodes in the new cluster.
}

/*
 * Structure for the custom type CreateClusterResult
 */
type CreateClusterResult struct {
	ClusterId        *int64        `json:"clusterId,omitempty" form:"clusterId,omitempty"`               //Specifies the ID of the new Cluster.
	ClusterName      *string       `json:"clusterName,omitempty" form:"clusterName,omitempty"`           //Specifies the name of the new Cluster.
	ClusterSwVersion *string       `json:"clusterSwVersion,omitempty" form:"clusterSwVersion,omitempty"` //Specifies the software version of the new Cluster.
	HealthyNodes     []*NodeStatus `json:"healthyNodes,omitempty" form:"healthyNodes,omitempty"`         //Specifies the status of the Nodes in the Cluster. All Nodes that
	IncarnationId    *int64        `json:"incarnationId,omitempty" form:"incarnationId,omitempty"`       //Specifies the Incarnation ID of the new Cluster.
	Message          *string       `json:"message,omitempty" form:"message,omitempty"`                   //Specifies an optional message field.
	UnhealthyNodes   []*NodeStatus `json:"unhealthyNodes,omitempty" form:"unhealthyNodes,omitempty"`     //Specifies the status of the Nodes in the Cluster. All Nodes that are
}

/*
 * Structure for the custom type CreateIdpConfigurationRequest
 */
type CreateIdpConfigurationRequest struct {
	AllowLocalAuthentication *bool     `json:"allowLocalAuthentication,omitempty" form:"allowLocalAuthentication,omitempty"` //Specifies whether to allow local authentication. When IdP is configured,
	Certificate              *string   `json:"certificate,omitempty" form:"certificate,omitempty"`                           //Specifies the certificate generated for the app by the IdP service when
	CertificateFilename      *string   `json:"certificateFilename,omitempty" form:"certificateFilename,omitempty"`           //Specifies the filename used to upload the certificate.
	Domain                   *string   `json:"domain,omitempty" form:"domain,omitempty"`                                     //Specifies a unique name for this IdP configuration.
	Enable                   *bool     `json:"enable,omitempty" form:"enable,omitempty"`                                     //Specifies a flag to enable or disable this IdP service. When it is set
	IssuerId                 *string   `json:"issuerId,omitempty" form:"issuerId,omitempty"`                                 //Specifies the IdP provided Issuer ID for the app.
	Name                     *string   `json:"name,omitempty" form:"name,omitempty"`                                         //Specifies the name of the vendor providing IdP service.
	Roles                    *[]string `json:"roles,omitempty" form:"roles,omitempty"`                                       //Specifies a list roles assigned to an IdP user if samlAttributeName is
	SamlAttributeName        *string   `json:"samlAttributeName,omitempty" form:"samlAttributeName,omitempty"`               //Specifies the SAML attribute name that contains a comma separated list
	SignRequest              *bool     `json:"signRequest,omitempty" form:"signRequest,omitempty"`                           //Specifies whether to sign the SAML request or not. When it is set
	SsoUrl                   *string   `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"`                                     //Specifies the SSO URL of the IdP service for the customer. This is the
	TenantId                 *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                 //Specifies the Tenant Id if the IdP is configured for a Tenant. If this is
}

/*
 * Structure for the custom type CreatePhysicalClusterParameters
 */
type CreatePhysicalClusterParameters struct {
	ClusterName            string                       `json:"clusterName" form:"clusterName"`                                           //Specifies the name of the new Cluster.
	EncryptionConfig       *EncryptionConfiguration     `json:"encryptionConfig,omitempty" form:"encryptionConfig,omitempty"`             //Specifies the parameters the user wants to use when configuring encryption
	IpmiConfig             IpmiConfiguration            `json:"ipmiConfig" form:"ipmiConfig"`                                             //Specifies the parameters for configuration of IPMI. This is only needed
	MetadataFaultTolerance *int64                       `json:"metadataFaultTolerance,omitempty" form:"metadataFaultTolerance,omitempty"` //Specifies the metadata fault tolerance.
	NetworkConfig          NetworkConfiguration         `json:"networkConfig" form:"networkConfig"`                                       //Specifies all of the parameters needed for network configuration of
	NodeConfigs            []*PhysicalNodeConfiguration `json:"nodeConfigs" form:"nodeConfigs"`                                           //Specifies the configuration for the nodes in the new cluster.
}

/*
 * Structure for the custom type CreateRemoteVaultRestoreTaskParameters
 */
type CreateRemoteVaultRestoreTaskParameters struct {
	GlacierRetrievalType GlacierRetrievalTypeEnum `json:"glacierRetrievalType,omitempty" form:"glacierRetrievalType,omitempty"` //Specifies the way data needs to be retrieved from the external target.
	RestoreObjects       []*IndexAndSnapshots     `json:"restoreObjects,omitempty" form:"restoreObjects,omitempty"`             //Array of Restore Objects.
	SearchJobUid         UniversalId              `json:"searchJobUid" form:"searchJobUid"`                                     //Specifies the unique id of the remote Vault search Job.
	TaskName             string                   `json:"taskName" form:"taskName"`                                             //Specifies a name of the restore task.
	VaultId              int64                    `json:"vaultId" form:"vaultId"`                                               //Specifies the id of the Vault that contains the index and
}

/*
 * Structure for the custom type CreateRemoteVaultSearchJobParameters
 */
type CreateRemoteVaultSearchJobParameters struct {
	ClusterMatchString *string               `json:"clusterMatchString,omitempty" form:"clusterMatchString,omitempty"` //Filter by specifying a Cluster name prefix string.
	EncryptionKeys     []*VaultEncryptionKey `json:"encryptionKeys,omitempty" form:"encryptionKeys,omitempty"`         //Array of Encryption Keys.
	EndTimeUsecs       *int64                `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`             //Filter by a end time specified as a Unix epoch Timestamp
	JobMatchString     *string               `json:"jobMatchString,omitempty" form:"jobMatchString,omitempty"`         //Filter by specifying a Protection Job name prefix string.
	SearchJobName      string                `json:"searchJobName" form:"searchJobName"`                               //Specifies the search Job name.
	StartTimeUsecs     *int64                `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`         //Filter by a start time specified as a Unix epoch Timestamp
	VaultId            int64                 `json:"vaultId" form:"vaultId"`                                           //Specifies the id of the Vault to search. This id was assigned by the
}

/*
 * Structure for the custom type CreateViewBoxParams
 */
type CreateViewBoxParams struct {
	AdDomainName                    *string        `json:"adDomainName,omitempty" form:"adDomainName,omitempty"`                                       //Specifies an active directory domain that this view box is mapped to.
	ClientSubnetWhiteList           []*Subnet      `json:"clientSubnetWhiteList,omitempty" form:"clientSubnetWhiteList,omitempty"`                     //Array of Subnets.
	CloudDownWaterfallThresholdPct  *int64         `json:"cloudDownWaterfallThresholdPct,omitempty" form:"cloudDownWaterfallThresholdPct,omitempty"`   //Specifies the cloud down water-fall threshold percentage. This indicates
	CloudDownWaterfallThresholdSecs *int64         `json:"cloudDownWaterfallThresholdSecs,omitempty" form:"cloudDownWaterfallThresholdSecs,omitempty"` //Specifies the cloud down water-fall threshold seconds. This indicates
	ClusterPartitionId              int64          `json:"clusterPartitionId" form:"clusterPartitionId"`                                               //Specifies the Cluster Partition id where the Storage Domain (View Box) is
	DefaultUserQuotaPolicy          *QuotaPolicy   `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"`                   //Specifies an optional quota policy/limits that are inherited by all users
	DefaultViewQuotaPolicy          *QuotaPolicy   `json:"defaultViewQuotaPolicy,omitempty" form:"defaultViewQuotaPolicy,omitempty"`                   //Specifies an optional default logical quota limit (in bytes)
	LdapProviderId                  *int64         `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"`                                   //When set, the following provides the LDAP provider the view box is
	Name                            string         `json:"name" form:"name"`                                                                           //Specifies the name of the Storage Domain (View Box).
	PhysicalQuota                   *QuotaPolicy   `json:"physicalQuota,omitempty" form:"physicalQuota,omitempty"`                                     //Specifies an optional quota limit (in bytes) for the physical
	S3BucketsAllowed                *bool          `json:"s3BucketsAllowed,omitempty" form:"s3BucketsAllowed,omitempty"`                               //Specifies whether creation of a S3 bucket is allowed in this
	StoragePolicy                   *StoragePolicy `json:"storagePolicy,omitempty" form:"storagePolicy,omitempty"`                                     //Specifies the storage options applied to a Storage Domain (View Box).
	TenantIdVec                     *[]string      `json:"tenantIdVec,omitempty" form:"tenantIdVec,omitempty"`                                         //Optional ids for the tenants that this view box belongs. This must be
}

/*
 * Structure for the custom type CreateViewRequest
 */
type CreateViewRequest struct {
	AccessSids                      *[]string                               `json:"accessSids,omitempty" form:"accessSids,omitempty"`                                           //Array of Security Identifiers (SIDs)
	AntivirusScanConfig             *AntivirusScanConfig                    `json:"antivirusScanConfig,omitempty" form:"antivirusScanConfig,omitempty"`                         //Specifies the antivirus scan config settings for this View.
	CaseInsensitiveNamesEnabled     *bool                                   `json:"caseInsensitiveNamesEnabled,omitempty" form:"caseInsensitiveNamesEnabled,omitempty"`         //Specifies whether to support case insensitive file/folder names. This
	Description                     *string                                 `json:"description,omitempty" form:"description,omitempty"`                                         //Specifies an optional text description about the View.
	EnableFilerAuditLogging         *bool                                   `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"`                 //Specifies if Filer Audit Logging is enabled for this view.
	EnableMixedModePermissions      *bool                                   `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"`           //If set, mixed mode (NFS and SMB) access is enabled for this view.
	EnableNfsViewDiscovery          *bool                                   `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"`                   //If set, it enables discovery of view for NFS.
	EnableOfflineCaching            *bool                                   `json:"enableOfflineCaching,omitempty" form:"enableOfflineCaching,omitempty"`                       //Specifies whether to enable offline file caching of the view.
	EnableSmbAccessBasedEnumeration *bool                                   `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
	EnableSmbEncryption             *bool                                   `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`                         //Specifies the SMB encryption for the View. If set, it enables the SMB
	EnableSmbViewDiscovery          *bool                                   `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"`                   //If set, it enables discovery of view for SMB.
	EnforceSmbEncryption            *bool                                   `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`                       //Specifies the SMB encryption for all the sessions for the View.
	FileExtensionFilter             *FileExtensionFilter                    `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"`                         //TODO: Write general description for this field
	FileLockConfig                  *FileLevelDataLockConfig                `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"`                                   //Specifies a config to lock files in a view - to protect from malicious or
	LogicalQuota                    *QuotaPolicy                            `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"`                                       //Specifies an optional logical quota limit (in bytes) for the usage allowed
	Name                            string                                  `json:"name" form:"name"`                                                                           //Specifies the name of the new View to create.
	NfsRootPermissions              *NfsRootPermissions                     `json:"nfsRootPermissions,omitempty" form:"nfsRootPermissions,omitempty"`                           //Specifies the config of NFS root permission of a view file system.
	OverrideGlobalWhitelist         *bool                                   `json:"overrideGlobalWhitelist,omitempty" form:"overrideGlobalWhitelist,omitempty"`                 //Specifies whether view level client subnet whitelist overrides cluster and
	ProtocolAccess                  ProtocolAccessEnum                      `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"`                                   //Specifies the supported Protocols for the View.
	Qos                             *QoS                                    `json:"qos,omitempty" form:"qos,omitempty"`                                                         //Specifies the Quality of Service (QoS) Policy for the View.
	S3KeyMappingConfig              S3KeyMappingConfigCreateViewRequestEnum `json:"s3KeyMappingConfig,omitempty" form:"s3KeyMappingConfig,omitempty"`                           //Specifies key mapping config of S3 storage.
	SecurityMode                    SecurityModeEnum                        `json:"securityMode,omitempty" form:"securityMode,omitempty"`                                       //Specifies the security mode used for this view.
	SharePermissions                []*SmbPermission                        `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`                               //Specifies a list of share level permissions.
	SmbPermissionsInfo              *SmbPermissionsInfo                     `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"`                           //Specifies information about SMB permissions.
	StoragePolicyOverride           *StoragePolicyOverride                  `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"`                     //Specifies if inline deduplication and compression settings inherited from
	SubnetWhitelist                 []*Subnet                               `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`                                 //Array of Subnets.
	TenantId                        *string                                 `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                               //Optional tenant id who has access to this View.
	ViewBoxId                       int64                                   `json:"viewBoxId" form:"viewBoxId"`                                                                 //Specifies the id of the Storage Domain (View Box) where the View will be
}

/*
 * Structure for the custom type CreateVirtualClusterParameters
 */
type CreateVirtualClusterParameters struct {
	ClusterName            string                      `json:"clusterName" form:"clusterName"`                                           //Specifies the name of the new Cluster.
	EncryptionConfig       *EncryptionConfiguration    `json:"encryptionConfig,omitempty" form:"encryptionConfig,omitempty"`             //Specifies the parameters the user wants to use when configuring encryption
	MetadataFaultTolerance *int64                      `json:"metadataFaultTolerance,omitempty" form:"metadataFaultTolerance,omitempty"` //Specifies the metadata fault tolerance.
	NetworkConfig          NetworkConfiguration        `json:"networkConfig" form:"networkConfig"`                                       //Specifies all of the parameters needed for network configuration of
	NodeConfigs            []*VirtualNodeConfiguration `json:"nodeConfigs" form:"nodeConfigs"`                                           //Specifies the configuration for the nodes in the new cluster.
}

/*
 * Structure for the custom type CreatedRemoteVaultSearchJobUid
 */
type CreatedRemoteVaultSearchJobUid struct {
	SearchJobUid *UniversalId `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id assigned for the search Job on the Cluster.
}

/*
 * Structure for the custom type Credentials
 */
type Credentials struct {
	Password *string `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
	Username *string `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
}

/*
 * Structure for the custom type CustomUnixIdAttributes
 */
type CustomUnixIdAttributes struct {
	GidAttrName *string `json:"gidAttrName,omitempty" form:"gidAttrName,omitempty"` //Specifies the custom field name in Active Directory user properties to get
	UidAttrName *string `json:"uidAttrName,omitempty" form:"uidAttrName,omitempty"` //Specifies the custom field name in Active Directory user properties to get
}

/*
 * Structure for the custom type DailySchedule
 */
type DailySchedule struct {
	Days *[]DayEnum `json:"days,omitempty" form:"days,omitempty"` //Array of Days.
}

/*
 * Structure for the custom type DataMigrationJobParameters
 */
type DataMigrationJobParameters struct {
	ColdFileWindow      *int64                  `json:"coldFileWindow,omitempty" form:"coldFileWindow,omitempty"`           //Identifies the cold files in the NAS source. Files that haven't been
	FilePathFilter      *FilePathFilter         `json:"filePathFilter,omitempty" form:"filePathFilter,omitempty"`           //Specifies filters to match files and directories on a Server.
	FileSelectionPolicy FileSelectionPolicyEnum `json:"fileSelectionPolicy,omitempty" form:"fileSelectionPolicy,omitempty"` //Specifies policy to select a file to migrate based on its creation, last
	FileSizeBytes       *int64                  `json:"fileSizeBytes,omitempty" form:"fileSizeBytes,omitempty"`             //Gives the size criteria to be used for selecting the files to be migrated
	FileSizePolicy      FileSizePolicyEnum      `json:"fileSizePolicy,omitempty" form:"fileSizePolicy,omitempty"`           //Specifies policy to select a file to migrate based on its size.
	NfsMountPath        *string                 `json:"nfsMountPath,omitempty" form:"nfsMountPath,omitempty"`               //Mount path where the target view must be mounted on all NFS clients for
	TargetViewName      *string                 `json:"targetViewName,omitempty" form:"targetViewName,omitempty"`           //The target view name to which the data will be migrated.
}

/*
 * Structure for the custom type DataMigrationPolicy
 */
type DataMigrationPolicy struct {
	DaysToKeep        *int64                                   `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`               //Specifies how many days to retain Snapshots on the Cohesity Cluster.
	SchedulingPolicy  *SchedulingPolicy                        `json:"schedulingPolicy,omitempty" form:"schedulingPolicy,omitempty"`   //Specifies settings that define a backup schedule for a Protection Job.
	WormRetentionType WormRetentionTypeDataMigrationPolicyEnum `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"` //Specifies WORM retention type for the files. When a WORM retention
}

/*
 * Structure for the custom type DataTransferFromVaultPerTask
 */
type DataTransferFromVaultPerTask struct {
	NumLogicalBytesTransferred  *int64  `json:"numLogicalBytesTransferred,omitempty" form:"numLogicalBytesTransferred,omitempty"`   //Specifies the total number of logical bytes that are transferred from
	NumPhysicalBytesTransferred *int64  `json:"numPhysicalBytesTransferred,omitempty" form:"numPhysicalBytesTransferred,omitempty"` //Specifies the total number of physical bytes that are transferred
	TaskName                    *string `json:"taskName,omitempty" form:"taskName,omitempty"`                                       //Specifies the task name.
	TaskType                    *string `json:"taskType,omitempty" form:"taskType,omitempty"`                                       //Specifies the task type.
}

/*
 * Structure for the custom type DataTransferFromVaultSummary
 */
type DataTransferFromVaultSummary struct {
	DataTransferPerTask                         []*DataTransferFromVaultPerTask `json:"dataTransferPerTask,omitempty" form:"dataTransferPerTask,omitempty"`                                                 //Array of Data Transferred Per Task.
	NumLogicalBytesTransferred                  *int64                          `json:"numLogicalBytesTransferred,omitempty" form:"numLogicalBytesTransferred,omitempty"`                                   //Specifies the total number of logical bytes that have been transferred
	NumPhysicalBytesTransferred                 *int64                          `json:"numPhysicalBytesTransferred,omitempty" form:"numPhysicalBytesTransferred,omitempty"`                                 //Specifies the total number of physical bytes that have been transferred
	NumTasks                                    *int64                          `json:"numTasks,omitempty" form:"numTasks,omitempty"`                                                                       //Specifies the number of recover or clone tasks that have transferred data
	PhysicalDataTransferredBytesDuringTimeRange *[]int64                        `json:"physicalDataTransferredBytesDuringTimeRange,omitempty" form:"physicalDataTransferredBytesDuringTimeRange,omitempty"` //Array of Physical Data Transferred Per Day.
	VaultName                                   *string                         `json:"vaultName,omitempty" form:"vaultName,omitempty"`                                                                     //Specifies the name of the Vault (External Target).
}

/*
 * Structure for the custom type DataUsageStats
 */
type DataUsageStats struct {
	CloudDataWrittenBytes          *int64 `json:"cloudDataWrittenBytes,omitempty" form:"cloudDataWrittenBytes,omitempty"`                   //Specifies the total data written on cloud tiers, as computed by the
	CloudTotalPhysicalUsageBytes   *int64 `json:"cloudTotalPhysicalUsageBytes,omitempty" form:"cloudTotalPhysicalUsageBytes,omitempty"`     //Specifies the total cloud capacity, as computed by the Cohesity Cluster,
	DataInBytes                    *int64 `json:"dataInBytes,omitempty" form:"dataInBytes,omitempty"`                                       //Specifies the data brought into the cluster. This is the usage before data
	DataInBytesAfterDedup          *int64 `json:"dataInBytesAfterDedup,omitempty" form:"dataInBytesAfterDedup,omitempty"`                   //Specifies the the the size of the data has been reduced by change-block
	DataProtectLogicalUsageBytes   *int64 `json:"dataProtectLogicalUsageBytes,omitempty" form:"dataProtectLogicalUsageBytes,omitempty"`     //Specifies the logical data used by Data Protect on Cohesity cluster.
	DataProtectPhysicalUsageBytes  *int64 `json:"dataProtectPhysicalUsageBytes,omitempty" form:"dataProtectPhysicalUsageBytes,omitempty"`   //Specifies the physical data used by Data Protect on Cohesity cluster.
	DataWrittenBytes               *int64 `json:"dataWrittenBytes,omitempty" form:"dataWrittenBytes,omitempty"`                             //Specifies the total data written on local and cloud tiers, as computed
	FileServicesLogicalUsageBytes  *int64 `json:"fileServicesLogicalUsageBytes,omitempty" form:"fileServicesLogicalUsageBytes,omitempty"`   //Specifies the logical data used by File services on Cohesity cluster.
	FileServicesPhysicalUsageBytes *int64 `json:"fileServicesPhysicalUsageBytes,omitempty" form:"fileServicesPhysicalUsageBytes,omitempty"` //Specifies the physical data used by File services on Cohesity cluster.
	LocalDataWrittenBytes          *int64 `json:"localDataWrittenBytes,omitempty" form:"localDataWrittenBytes,omitempty"`                   //Specifies the total data written on local tiers, as computed by the
	LocalTierResiliencyImpactBytes *int64 `json:"localTierResiliencyImpactBytes,omitempty" form:"localTierResiliencyImpactBytes,omitempty"` //Specifies the size of the data has been replicated to other nodes as per
	LocalTotalPhysicalUsageBytes   *int64 `json:"localTotalPhysicalUsageBytes,omitempty" form:"localTotalPhysicalUsageBytes,omitempty"`     //Specifies the total local capacity, as computed by the Cohesity Cluster,
	StorageConsumedBytes           *int64 `json:"storageConsumedBytes,omitempty" form:"storageConsumedBytes,omitempty"`                     //Specifies the total capacity, as computed by the Cohesity Cluster,
	TotalLogicalUsageBytes         *int64 `json:"totalLogicalUsageBytes,omitempty" form:"totalLogicalUsageBytes,omitempty"`                 //Specifies the logical usage as computed by the Cohesity Cluster.
}

/*
 * Structure for the custom type DatastoreInfo
 */
type DatastoreInfo struct {
	Capacity  *int64 `json:"capacity,omitempty" form:"capacity,omitempty"`   //Specifies the capacity of the datastore in bytes.
	FreeSpace *int64 `json:"freeSpace,omitempty" form:"freeSpace,omitempty"` //Specifies the available space on the datastore in bytes.
}

/*
 * Structure for the custom type DbFileInfo
 */
type DbFileInfo struct {
	FileType  FileTypeEnum `json:"fileType,omitempty" form:"fileType,omitempty"`   //Specifies the format type of the file that SQL database stores the data.
	FullPath  *string      `json:"fullPath,omitempty" form:"fullPath,omitempty"`   //Specifies the full path of the database file on the SQL host machine.
	SizeBytes *int64       `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Specifies the last known size of the database file.
}

/*
 * Structure for the custom type DeleteInfectedFileParams
 */
type DeleteInfectedFileParams struct {
	InfectedFileIds []*InfectedFileParam `json:"infectedFileIds,omitempty" form:"infectedFileIds,omitempty"` //Specifies the list of infected file path.
}

/*
 * Structure for the custom type DeleteInfectedFileResponse
 */
type DeleteInfectedFileResponse struct {
	DeleteFailedInfectedFiles    []*InfectedFileId `json:"deleteFailedInfectedFiles,omitempty" form:"deleteFailedInfectedFiles,omitempty"`       //Specifies the failed delete infected files.
	DeleteSucceededInfectedFiles []*InfectedFileId `json:"deleteSucceededInfectedFiles,omitempty" form:"deleteSucceededInfectedFiles,omitempty"` //Specifies the successfully deleted infected files.
}

/*
 * Structure for the custom type DeleteProtectionJobParam
 */
type DeleteProtectionJobParam struct {
	DeleteSnapshots *bool `json:"deleteSnapshots,omitempty" form:"deleteSnapshots,omitempty"` //Specifies if Snapshots generated by the Protection Job should also be
}

/*
 * Structure for the custom type DeleteRouteParam
 */
type DeleteRouteParam struct {
	DestNetwork    *string `json:"destNetwork,omitempty" form:"destNetwork,omitempty"`       //Destination network.
	IfName         *string `json:"ifName,omitempty" form:"ifName,omitempty"`                 //Specifies the network interfaces name to use for communicating with the
	IfaceGroupName *string `json:"ifaceGroupName,omitempty" form:"ifaceGroupName,omitempty"` //Specifies the network interfaces group or vlan interface group to
}

/*
 * Structure for the custom type DeleteViewUsersQuotaParameters
 */
type DeleteViewUsersQuotaParameters struct {
	DeleteAll *bool     `json:"deleteAll,omitempty" form:"deleteAll,omitempty"` //Delete all existing user quota override policies.
	UserIds   []*UserId `json:"userIds,omitempty" form:"userIds,omitempty"`     //The user ids whose policy needs to be deleted.
	ViewName  *string   `json:"viewName,omitempty" form:"viewName,omitempty"`   //View name of input view.
}

/*
 * Structure for the custom type DeliveryRuleProtoDeliveryTarget
 */
type DeliveryRuleProtoDeliveryTarget struct {
	EmailAddress           *string `json:"emailAddress,omitempty" form:"emailAddress,omitempty"`                     //List of email addresses to send notifications.
	ExternalApiCurlOptions *string `json:"externalApiCurlOptions,omitempty" form:"externalApiCurlOptions,omitempty"` //Specifies the curl options used to invoke above rest external api.
	ExternalApiUrl         *string `json:"externalApiUrl,omitempty" form:"externalApiUrl,omitempty"`                 //Specifies the external api to be invoked when an alert matching this
	Locale                 *string `json:"locale,omitempty" form:"locale,omitempty"`                                 //Locale for the delivery target.
}

/*
 * Structure for the custom type DeployCertParameters
 */
type DeployCertParameters struct {
	CertFileName *string   `json:"certFileName,omitempty" form:"certFileName,omitempty"` //Specifies the filename of the certificate.
	Password     *[]string `json:"password,omitempty" form:"password,omitempty"`         //Specifies the passsword of the host to establish SSH connection.
	ServerName   *[]string `json:"serverName,omitempty" form:"serverName,omitempty"`     //Specifies the servername of the host where certificate is to be deployed.
	Target       *[]string `json:"target,omitempty" form:"target,omitempty"`             //Specifies the target location on the host where the certificate is
	UserName     *[]string `json:"userName,omitempty" form:"userName,omitempty"`         //TODO(Sai Akhil S): Accept credentials for the cluster instead of each
	ValidDays    *int64    `json:"validDays,omitempty" form:"validDays,omitempty"`       //Specifies the number of days after which the certificate will expire. The
}

/*
 * Structure for the custom type DeployDBInstancesToRDSParams
 */
type DeployDBInstancesToRDSParams struct {
	AutoMinorVersionUpgrade *bool                                                 `json:"autoMinorVersionUpgrade,omitempty" form:"autoMinorVersionUpgrade,omitempty"` //Whether to enable auto minor version upgrade in the restored DB.
	AvailabilityZone        *EntityProto                                          `json:"availabilityZone,omitempty" form:"availabilityZone,omitempty"`               //Specifies the attributes and the latest statistics about an entity.
	CopyTagsToSnapshots     *bool                                                 `json:"copyTagsToSnapshots,omitempty" form:"copyTagsToSnapshots,omitempty"`         //Whether to enable copying of tags to snapshots of the DB.
	DbInstanceId            *string                                               `json:"dbInstanceId,omitempty" form:"dbInstanceId,omitempty"`                       //The DB instance identifier to use for the restored DB. This field is
	DbOptionGroup           *EntityProto                                          `json:"dbOptionGroup,omitempty" form:"dbOptionGroup,omitempty"`                     //Specifies the attributes and the latest statistics about an entity.
	DbParameterGroup        *EntityProto                                          `json:"dbParameterGroup,omitempty" form:"dbParameterGroup,omitempty"`               //Specifies the attributes and the latest statistics about an entity.
	DbPort                  *int64                                                `json:"dbPort,omitempty" form:"dbPort,omitempty"`                                   //Port to use for the DB in the restored RDS instance.
	IamDbAuthentication     *bool                                                 `json:"iamDbAuthentication,omitempty" form:"iamDbAuthentication,omitempty"`         //Whether to enable IAM authentication for the DB.
	MultiAzDeployment       *bool                                                 `json:"multiAzDeployment,omitempty" form:"multiAzDeployment,omitempty"`             //Whether this is a multi-az deployment or not.
	PointInTimeParams       *DeployDBInstancesToRDSParamsPointInTimeRestoreParams `json:"pointInTimeParams,omitempty" form:"pointInTimeParams,omitempty"`             //Message to capture details of a point in time that the DB needs to be
	PublicAccessibility     *bool                                                 `json:"publicAccessibility,omitempty" form:"publicAccessibility,omitempty"`         //Whether this DB will be publicly accessible or not.
}

/*
 * Structure for the custom type DeployDBInstancesToRDSParamsPointInTimeRestoreParams
 */
type DeployDBInstancesToRDSParamsPointInTimeRestoreParams struct {
	TimestampMsecs *int64 `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Time in milliseconds since epoch that the DB needs to be restored to.
}

/*
 * Structure for the custom type DeployTaskRequest
 */
type DeployTaskRequest struct {
	Name        string                    `json:"name" form:"name"`                                   //Specifies the name of the Deploy Task. This field must be set and must be
	NewParentId *int64                    `json:"newParentId,omitempty" form:"newParentId,omitempty"` //Specifies a new registered parent Protection Source. If specified
	Objects     []*RestoreObjectDetails   `json:"objects,omitempty" form:"objects,omitempty"`         //Array of Objects.
	Target      *CloudDeployTargetDetails `json:"target,omitempty" form:"target,omitempty"`           //Message that specifies the details about CloudDeploy target where backup
}

/*
 * Structure for the custom type DeployVMsToAWSParams
 */
type DeployVMsToAWSParams struct {
	InstanceType          *EntityProto                  `json:"instanceType,omitempty" form:"instanceType,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
	KeyPairName           *EntityProto                  `json:"keyPairName,omitempty" form:"keyPairName,omitempty"`                     //Specifies the attributes and the latest statistics about an entity.
	NetworkSecurityGroups []*EntityProto                `json:"networkSecurityGroups,omitempty" form:"networkSecurityGroups,omitempty"` //Names of the network security groups within the above VPC. At least
	RdsParams             *DeployDBInstancesToRDSParams `json:"rdsParams,omitempty" form:"rdsParams,omitempty"`                         //Contains RDS specfic options that can be supplied while restoring the RDS
	Region                *EntityProto                  `json:"region,omitempty" form:"region,omitempty"`                               //Specifies the attributes and the latest statistics about an entity.
	Subnet                *EntityProto                  `json:"subnet,omitempty" form:"subnet,omitempty"`                               //Specifies the attributes and the latest statistics about an entity.
	Vpc                   *EntityProto                  `json:"vpc,omitempty" form:"vpc,omitempty"`                                     //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DeployVMsToAzureParams
 */
type DeployVMsToAzureParams struct {
	AzureManagedDiskParams *AzureManagedDiskParams `json:"azureManagedDiskParams,omitempty" form:"azureManagedDiskParams,omitempty"` //Contains managed disk parameters needed to deploy to Azure using managed
	ComputeOptions         *EntityProto            `json:"computeOptions,omitempty" form:"computeOptions,omitempty"`                 //Specifies the attributes and the latest statistics about an entity.
	NetworkResourceGroup   *EntityProto            `json:"networkResourceGroup,omitempty" form:"networkResourceGroup,omitempty"`     //Specifies the attributes and the latest statistics about an entity.
	NetworkSecurityGroup   *EntityProto            `json:"networkSecurityGroup,omitempty" form:"networkSecurityGroup,omitempty"`     //Specifies the attributes and the latest statistics about an entity.
	ResourceGroup          *EntityProto            `json:"resourceGroup,omitempty" form:"resourceGroup,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
	StorageAccount         *EntityProto            `json:"storageAccount,omitempty" form:"storageAccount,omitempty"`                 //Specifies the attributes and the latest statistics about an entity.
	StorageContainer       *EntityProto            `json:"storageContainer,omitempty" form:"storageContainer,omitempty"`             //Specifies the attributes and the latest statistics about an entity.
	StorageKey             *EntityProto            `json:"storageKey,omitempty" form:"storageKey,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	StorageResourceGroup   *EntityProto            `json:"storageResourceGroup,omitempty" form:"storageResourceGroup,omitempty"`     //Specifies the attributes and the latest statistics about an entity.
	Subnet                 *EntityProto            `json:"subnet,omitempty" form:"subnet,omitempty"`                                 //Specifies the attributes and the latest statistics about an entity.
	TempVmResourceGroup    *EntityProto            `json:"tempVmResourceGroup,omitempty" form:"tempVmResourceGroup,omitempty"`       //Specifies the attributes and the latest statistics about an entity.
	TempVmStorageAccount   *EntityProto            `json:"tempVmStorageAccount,omitempty" form:"tempVmStorageAccount,omitempty"`     //Specifies the attributes and the latest statistics about an entity.
	TempVmStorageContainer *EntityProto            `json:"tempVmStorageContainer,omitempty" form:"tempVmStorageContainer,omitempty"` //Specifies the attributes and the latest statistics about an entity.
	TempVmSubnet           *EntityProto            `json:"tempVmSubnet,omitempty" form:"tempVmSubnet,omitempty"`                     //Specifies the attributes and the latest statistics about an entity.
	TempVmVirtualNetwork   *EntityProto            `json:"tempVmVirtualNetwork,omitempty" form:"tempVmVirtualNetwork,omitempty"`     //Specifies the attributes and the latest statistics about an entity.
	VirtualNetwork         *EntityProto            `json:"virtualNetwork,omitempty" form:"virtualNetwork,omitempty"`                 //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DeployVMsToCloudParams
 */
type DeployVMsToCloudParams struct {
	DeployVmsToAwsParams          *DeployVMsToAWSParams          `json:"deployVmsToAwsParams,omitempty" form:"deployVmsToAwsParams,omitempty"`                   //Contains AWS specific information needed to identify various resources
	DeployVmsToAzureParams        *DeployVMsToAzureParams        `json:"deployVmsToAzureParams,omitempty" form:"deployVmsToAzureParams,omitempty"`               //Contains Azure specific information needed to identify various resources
	DeployVmsToGcpParams          *DeployVMsToGCPParams          `json:"deployVmsToGcpParams,omitempty" form:"deployVmsToGcpParams,omitempty"`                   //Contains GCP specific information needed to identify various resources
	ReplicateSnapshotsToAwsParams *ReplicateSnapshotsToAWSParams `json:"replicateSnapshotsToAwsParams,omitempty" form:"replicateSnapshotsToAwsParams,omitempty"` //Params required to replicate snapshots to another AWS source. This is
}

/*
 * Structure for the custom type DeployVMsToCloudTaskStateProto
 */
type DeployVMsToCloudTaskStateProto struct {
	DeployVmsToCloudParams *DeployVMsToCloudParams `json:"deployVmsToCloudParams,omitempty" form:"deployVmsToCloudParams,omitempty"` //Contains Cloud specific information needed to identify various resources
}

/*
 * Structure for the custom type DeployVMsToGCPParams
 */
type DeployVMsToGCPParams struct {
	ProjectId *EntityProto `json:"projectId,omitempty" form:"projectId,omitempty"` //Specifies the attributes and the latest statistics about an entity.
	Region    *EntityProto `json:"region,omitempty" form:"region,omitempty"`       //Specifies the attributes and the latest statistics about an entity.
	Subnet    *EntityProto `json:"subnet,omitempty" form:"subnet,omitempty"`       //Specifies the attributes and the latest statistics about an entity.
	Zone      *EntityProto `json:"zone,omitempty" form:"zone,omitempty"`           //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DeployVmsToCloud
 */
type DeployVmsToCloud struct {
	AwsParams   *AwsParams   `json:"awsParams,omitempty" form:"awsParams,omitempty"`     //Specifies various resources when converting and deploying a VM to AWS.
	AzureParams *AzureParams `json:"azureParams,omitempty" form:"azureParams,omitempty"` //Specifies various resources when converting and deploying a VM to Azure.
}

/*
 * Structure for the custom type DestroyCloneAppTaskInfoProto
 */
type DestroyCloneAppTaskInfoProto struct {
	AppEnv                  *int64       `json:"appEnv,omitempty" form:"appEnv,omitempty"`                                   //The application environment.
	Error                   *ErrorProto  `json:"error,omitempty" form:"error,omitempty"`                                     //TODO: Write general description for this field
	Finished                *bool        `json:"finished,omitempty" form:"finished,omitempty"`                               //This will be set to true if the task is complete on the slave.
	TargetEntity            *EntityProto `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`                       //Specifies the attributes and the latest statistics about an entity.
	TargetEntityCredentials *Credentials `json:"targetEntityCredentials,omitempty" form:"targetEntityCredentials,omitempty"` //Specifies credentials to access a target source.
}

/*
 * Structure for the custom type DestroyClonedEntityInfoProto
 */
type DestroyClonedEntityInfoProto struct {
	ClonedEntity             *DestroyClonedEntityInfoProtoClonedEntity `json:"clonedEntity,omitempty" form:"clonedEntity,omitempty"`                         //TODO: Write general description for this field
	ClonedEntityStatus       *int64                                    `json:"clonedEntityStatus,omitempty" form:"clonedEntityStatus,omitempty"`             //TODO: Write general description for this field
	DestroyClonedEntityState *int64                                    `json:"destroyClonedEntityState,omitempty" form:"destroyClonedEntityState,omitempty"` //The state of the destroy/teardown of a cloned entity (i.e, VM).
	Error                    *ErrorProto                               `json:"error,omitempty" form:"error,omitempty"`                                       //TODO: Write general description for this field
	FullViewName             *string                                   `json:"fullViewName,omitempty" form:"fullViewName,omitempty"`                         //The full external view name where cloned objects are placed.
	Type                     *int64                                    `json:"type,omitempty" form:"type,omitempty"`                                         //The type of environment this destroy cloned entity info pertains to.
}

/*
 * Structure for the custom type DestroyClonedEntityInfoProtoClonedEntity
 */
type DestroyClonedEntityInfoProtoClonedEntity struct {
	Entity                 *EntityProto `json:"entity,omitempty" form:"entity,omitempty"`                                 //Specifies the attributes and the latest statistics about an entity.
	RelativeRestorePathVec *[]string    `json:"relativeRestorePathVec,omitempty" form:"relativeRestorePathVec,omitempty"` //Path of all files created by the clone operation. Each path is relative
}

/*
 * Structure for the custom type DestroyClonedTaskStateProto
 */
type DestroyClonedTaskStateProto struct {
	CloneTaskName                *string                           `json:"cloneTaskName,omitempty" form:"cloneTaskName,omitempty"`                               //The name of the clone task.
	DatastoreEntity              *EntityProto                      `json:"datastoreEntity,omitempty" form:"datastoreEntity,omitempty"`                           //Specifies the attributes and the latest statistics about an entity.
	DeployVmsToCloudTaskState    *DeployVMsToCloudTaskStateProto   `json:"deployVmsToCloudTaskState,omitempty" form:"deployVmsToCloudTaskState,omitempty"`       //TODO: Write general description for this field
	DestroyCloneAppTaskInfo      *DestroyCloneAppTaskInfoProto     `json:"destroyCloneAppTaskInfo,omitempty" form:"destroyCloneAppTaskInfo,omitempty"`           //Each available extension is listed below along with the location of the
	DestroyCloneVmTaskInfo       *DestroyClonedVMTaskInfoProto     `json:"destroyCloneVmTaskInfo,omitempty" form:"destroyCloneVmTaskInfo,omitempty"`             //Each available extension is listed below along with the location of the
	DestroyMountVolumesTaskInfo  *DestroyMountVolumesTaskInfoProto `json:"destroyMountVolumesTaskInfo,omitempty" form:"destroyMountVolumesTaskInfo,omitempty"`   //TODO: Write general description for this field
	EndTimeUsecs                 *int64                            `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                 //If the destroy clone task has finished, this field contains the end time
	Error                        *ErrorProto                       `json:"error,omitempty" form:"error,omitempty"`                                               //TODO: Write general description for this field
	FolderEntity                 *EntityProto                      `json:"folderEntity,omitempty" form:"folderEntity,omitempty"`                                 //Specifies the attributes and the latest statistics about an entity.
	FullViewName                 *string                           `json:"fullViewName,omitempty" form:"fullViewName,omitempty"`                                 //The full external view name where cloned objects are placed.
	ParentSourceConnectionParams *ConnectorParams                  `json:"parentSourceConnectionParams,omitempty" form:"parentSourceConnectionParams,omitempty"` //Message that encapsulates the various params required to establish a
	ParentTaskId                 *int64                            `json:"parentTaskId,omitempty" form:"parentTaskId,omitempty"`                                 //The id of the task that triggered the destroy task.
	PerformCloneTaskId           *int64                            `json:"performCloneTaskId,omitempty" form:"performCloneTaskId,omitempty"`                     //The unique id of the task that performed the clone operation.
	RestoreType                  *int64                            `json:"restoreType,omitempty" form:"restoreType,omitempty"`                                   //The type of the restore/clone operation that is being destroyed.
	ScheduledConstituentId       *int64                            `json:"scheduledConstituentId,omitempty" form:"scheduledConstituentId,omitempty"`             //Constituent id (and the gandalf session id) where this task has been
	ScheduledGandalfSessionId    *int64                            `json:"scheduledGandalfSessionId,omitempty" form:"scheduledGandalfSessionId,omitempty"`       //TODO: Write general description for this field
	StartTimeUsecs               *int64                            `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                             //The start time of this destroy clone task.
	Status                       *int64                            `json:"status,omitempty" form:"status,omitempty"`                                             //Status of the destroy clone task.
	TaskId                       *int64                            `json:"taskId,omitempty" form:"taskId,omitempty"`                                             //A globally unique id of this destroy clone task.
	Type                         *int64                            `json:"type,omitempty" form:"type,omitempty"`                                                 //The type of environment that is being operated on.
	User                         *string                           `json:"user,omitempty" form:"user,omitempty"`                                                 //The user who requested this destroy clone task.
	UserInfo                     *UserInformation                  `json:"userInfo,omitempty" form:"userInfo,omitempty"`                                         //A message to encapsulate information about the user who made the request.
	ViewBoxId                    *int64                            `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                                       //The view box id to which 'view_name' belongs to.
	ViewNameDEPRECATED           *string                           `json:"viewName_DEPRECATED,omitempty" form:"viewName_DEPRECATED,omitempty"`                   //The view name as provided by the user for the clone operation.
}

/*
 * Structure for the custom type DestroyClonedVMTaskInfoProto
 */
type DestroyClonedVMTaskInfoProto struct {
	DatastoreNotUnmountedReason *string                         `json:"datastoreNotUnmountedReason,omitempty" form:"datastoreNotUnmountedReason,omitempty"` //If datastore was not unmounted, this field contains the reason for the
	DatastoreUnmounted          *bool                           `json:"datastoreUnmounted,omitempty" form:"datastoreUnmounted,omitempty"`                   //Whether the datastore corresponding to the clone view was unmounted from
	DestroyClonedEntityInfoVec  []*DestroyClonedEntityInfoProto `json:"destroyClonedEntityInfoVec,omitempty" form:"destroyClonedEntityInfoVec,omitempty"`   //Vector of all cloned entities that this destroy task will teardown.
	Type                        *int64                          `json:"type,omitempty" form:"type,omitempty"`                                               //The type of environment this destroy clone task info pertains to.
	ViewDeleted                 *bool                           `json:"viewDeleted,omitempty" form:"viewDeleted,omitempty"`                                 //Whether the clone view was deleted by the destroy task.
}

/*
 * Structure for the custom type DestroyMountVolumesTaskInfoProto
 */
type DestroyMountVolumesTaskInfoProto struct {
	Error                   *ErrorProto            `json:"error,omitempty" form:"error,omitempty"`                                     //TODO: Write general description for this field
	Finished                *bool                  `json:"finished,omitempty" form:"finished,omitempty"`                               //This will be set to true if the task is complete on the slave.
	MountVolumesInfoProto   *MountVolumesInfoProto `json:"mountVolumesInfoProto,omitempty" form:"mountVolumesInfoProto,omitempty"`     //Each available extension is listed below along with the location of the
	SlaveTaskStartTimeUsecs *int64                 `json:"slaveTaskStartTimeUsecs,omitempty" form:"slaveTaskStartTimeUsecs,omitempty"` //This is the timestamp at which the slave task started.
	TargetEntity            *EntityProto           `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`                       //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DeviceNode
 */
type DeviceNode struct {
	IntermediateNode *DeviceTreeDetails  `json:"intermediateNode,omitempty" form:"intermediateNode,omitempty"` //Specifies a logical volume stored as a tree where the leaves are
	LeafNode         *FilePartitionBlock `json:"leafNode,omitempty" form:"leafNode,omitempty"`                 //Defines a leaf node of a device tree. This refers to a logical
}

/*
 * Structure for the custom type DeviceTree
 */
type DeviceTree struct {
	ChildVec     []*DeviceTreeChildDevice `json:"childVec,omitempty" form:"childVec,omitempty"`         //TODO: Write general description for this field
	DeviceLength *int64                   `json:"deviceLength,omitempty" form:"deviceLength,omitempty"` //The length of this device. This should match the length which is
	StripeSize   *int64                   `json:"stripeSize,omitempty" form:"stripeSize,omitempty"`     //In case data is striped, this represents the length of the stripe.
	Type         *int64                   `json:"type,omitempty" form:"type,omitempty"`                 //How to combine the children.
}

/*
 * Structure for the custom type DeviceTreeDetails
 */
type DeviceTreeDetails struct {
	CombineMethod CombineMethodEnum `json:"combineMethod,omitempty" form:"combineMethod,omitempty"` //Specifies how to combine the children of this node.
	DeviceLength  *int64            `json:"deviceLength,omitempty" form:"deviceLength,omitempty"`   //Specifies the length of this device. This number should match the
	DeviceNodes   []*DeviceNode     `json:"deviceNodes,omitempty" form:"deviceNodes,omitempty"`     //Specifies the children of this node in the device tree.
	StripeSize    *int64            `json:"stripeSize,omitempty" form:"stripeSize,omitempty"`       //Specifies the size of the striped data if the data is striped.
}

/*
 * Structure for the custom type DeviceTreeChildDevice
 */
type DeviceTreeChildDevice struct {
	Device         *DeviceTree               `json:"device,omitempty" form:"device,omitempty"`                 //TODO: Write general description for this field
	PartitionSlice *DeviceTreePartitionSlice `json:"partitionSlice,omitempty" form:"partitionSlice,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type DeviceTreePartitionSlice
 */
type DeviceTreePartitionSlice struct {
	DiskFileName    *string `json:"diskFileName,omitempty" form:"diskFileName,omitempty"`       //The disk to use.
	Length          *int64  `json:"length,omitempty" form:"length,omitempty"`                   //The length of data for the LVM volume (for which this device tree is
	LvmDataOffset   *int64  `json:"lvmDataOffset,omitempty" form:"lvmDataOffset,omitempty"`     //Each LVM partition starts with LVM meta data. After the meta data there
	Offset          *int64  `json:"offset,omitempty" form:"offset,omitempty"`                   //This is the offset (in bytes) where data for the LVM volume (for which
	PartitionNumber *int64  `json:"partitionNumber,omitempty" form:"partitionNumber,omitempty"` //The partition to use in the disk above.
}

/*
 * Structure for the custom type DirQuotaConfig
 */
type DirQuotaConfig struct {
	Enabled  *bool   `json:"enabled,omitempty" form:"enabled,omitempty"`   //Specifies whether the directory quota is enabled on the view.
	ViewName *string `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the view.
}

/*
 * Structure for the custom type DirQuotaInfo
 */
type DirQuotaInfo struct {
	Config *DirQuotaConfig   `json:"config,omitempty" form:"config,omitempty"` //Specifies the configuration object of a directory quota.
	Quotas []*DirQuotaPolicy `json:"quotas,omitempty" form:"quotas,omitempty"` //Specifies the list of directory quota policies applied on the view.
}

/*
 * Structure for the custom type DirQuotaPolicy
 */
type DirQuotaPolicy struct {
	UsageBytes *int64       `json:"UsageBytes,omitempty" form:"UsageBytes,omitempty"` //Specifies the current usage (in bytes) by the directory in the view.
	DirPath    *string      `json:"dirPath,omitempty" form:"dirPath,omitempty"`       //Specifies the path of the directory in the view.
	Policy     *QuotaPolicy `json:"policy,omitempty" form:"policy,omitempty"`         //Specifies a quota limit that can be optionally applied to Views and
}

/*
 * Structure for the custom type Disk
 */
type Disk struct {
	DiskBlocks           []*DiskBlock             `json:"diskBlocks,omitempty" form:"diskBlocks,omitempty"`                     //Array of Disk Blocks.
	DiskFormat           DiskFormatEnum           `json:"diskFormat,omitempty" form:"diskFormat,omitempty"`                     //Specifies the format of the virtual disk.
	DiskPartitions       []*DiskPartition         `json:"diskPartitions,omitempty" form:"diskPartitions,omitempty"`             //Array of Partitions.
	PartitionTableFormat PartitionTableFormatEnum `json:"partitionTableFormat,omitempty" form:"partitionTableFormat,omitempty"` //Specifies partition table format on a disk.
	SectorSizeBytes      *int64                   `json:"sectorSizeBytes,omitempty" form:"sectorSizeBytes,omitempty"`           //Specifies the sector size of hard disk. It is used for mapping the disk
	Uuid                 *string                  `json:"uuid,omitempty" form:"uuid,omitempty"`                                 //Specifies the disk uuid.
	VmdkFileName         *string                  `json:"vmdkFileName,omitempty" form:"vmdkFileName,omitempty"`                 //Specifies the disk file name. This is the VMDK name and not the
	VmdkSizeBytes        *int64                   `json:"vmdkSizeBytes,omitempty" form:"vmdkSizeBytes,omitempty"`               //Specifies the disk size in bytes.
}

/*
 * Structure for the custom type DiskBlock
 */
type DiskBlock struct {
	LengthBytes *int64 `json:"lengthBytes,omitempty" form:"lengthBytes,omitempty"` //Specifies the length of the block in bytes.
	OffsetBytes *int64 `json:"offsetBytes,omitempty" form:"offsetBytes,omitempty"` //Specifies the offset of the block (in bytes) from the beginning
}

/*
 * Structure for the custom type DiskPartition
 */
type DiskPartition struct {
	LengthBytes *int64  `json:"lengthBytes,omitempty" form:"lengthBytes,omitempty"` //Specifies the length of the block in bytes.
	Number      *int64  `json:"number,omitempty" form:"number,omitempty"`           //Specifies a unique number of the partition within the linear disk file.
	OffsetBytes *int64  `json:"offsetBytes,omitempty" form:"offsetBytes,omitempty"` //Specifies the offset of the block (in bytes) from the beginning
	TypeUuid    *string `json:"typeUuid,omitempty" form:"typeUuid,omitempty"`       //Specifies the partition type uuid.
	Uuid        *string `json:"uuid,omitempty" form:"uuid,omitempty"`               //Specifies the partition uuid.
}

/*
 * Structure for the custom type DiskUnit
 */
type DiskUnit struct {
	BusNumber      *int64  `json:"busNumber,omitempty" form:"busNumber,omitempty"`           //Specifies the Id of the controller bus that controls the disk.
	ControllerType *string `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
	UnitNumber     *int64  `json:"unitNumber,omitempty" form:"unitNumber,omitempty"`         //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type DomainControllers
 */
type DomainControllers struct {
	DomainControllers *[]string `json:"domainControllers,omitempty" form:"domainControllers,omitempty"` //Domain Controllers of a domain of an Active Directory domain.
}

/*
 * Structure for the custom type DownloadFilesAndFoldersParams
 */
type DownloadFilesAndFoldersParams struct {
	FilesAndFoldersInfo []*FilesAndFoldersInfo `json:"filesAndFoldersInfo,omitempty" form:"filesAndFoldersInfo,omitempty"` //Specifies the absolute paths for list of files and folders to download.
	Name                string                 `json:"name" form:"name"`                                                   //Specifies the name of the Download Task. This field must be set and must
	SourceObjectInfo    *RestoreObjectDetails  `json:"sourceObjectInfo,omitempty" form:"sourceObjectInfo,omitempty"`       //Specifies an object to recover or clone or an object to restore files
}

/*
 * Structure for the custom type DownloadPackageParameters
 */
type DownloadPackageParameters struct {
	Url string `json:"url" form:"url"` //Specifies a URL from which the package can be downloaded to the Cluster.
}

/*
 * Structure for the custom type DownloadPackageResult
 */
type DownloadPackageResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies a message describing the result of the request to download
}

/*
 * Structure for the custom type EditHostsParameters
 */
type EditHostsParameters struct {
	Hosts []*HostEntry `json:"hosts,omitempty" form:"hosts,omitempty"` //Specifies the list of host entries to be edited. Each IP address listed
}

/*
 * Structure for the custom type EmailDeliveryTarget
 */
type EmailDeliveryTarget struct {
	EmailAddress *string `json:"emailAddress,omitempty" form:"emailAddress,omitempty"` //TODO: Write general description for this field
	Locale       *string `json:"locale,omitempty" form:"locale,omitempty"`             //Specifies the language in which the emails sent to the above defined
}

/*
 * Structure for the custom type EmailMetaData
 */
type EmailMetaData struct {
	AllUnderHierarchy     *bool     `json:"allUnderHierarchy,omitempty" form:"allUnderHierarchy,omitempty"`         //AllUnderHierarchy specifies if logs of all the tenants under the hierarchy
	BccRecipientAddresses *[]string `json:"bccRecipientAddresses,omitempty" form:"bccRecipientAddresses,omitempty"` //Specifies the email addresses of the bcc recipients.
	CcRecipientAddresses  *[]string `json:"ccRecipientAddresses,omitempty" form:"ccRecipientAddresses,omitempty"`   //Specifies the email addresses of the cc recipients.
	DomainIds             *[]int64  `json:"domainIds,omitempty" form:"domainIds,omitempty"`                         //Specifies the domain Ids in which mailboxes are registered.
	EmailSubject          *string   `json:"emailSubject,omitempty" form:"emailSubject,omitempty"`                   //Specifies the subject of the email.
	FolderKey             *int64    `json:"folderKey,omitempty" form:"folderKey,omitempty"`                         //Specifes the Parent Folder key.
	FolderName            *string   `json:"folderName,omitempty" form:"folderName,omitempty"`                       //Specifies the parent folder name of the email.
	HasAttachments        *bool     `json:"hasAttachments,omitempty" form:"hasAttachments,omitempty"`               //Specifies whether the emails have any attachments.
	ItemKey               *string   `json:"itemKey,omitempty" form:"itemKey,omitempty"`                             //Specifies the Key(unique within mailbox) for Outlook item such as Email.
	MailboxIds            *[]int64  `json:"mailboxIds,omitempty" form:"mailboxIds,omitempty"`                       //Specifies the mailbox Ids which contains the emails/folders.
	ProtectionJobIds      *[]int64  `json:"protectionJobIds,omitempty" form:"protectionJobIds,omitempty"`           //Specifies the protection job Ids which have backed up mailbox(es)
	ReceivedEndTime       *int64    `json:"receivedEndTime,omitempty" form:"receivedEndTime,omitempty"`             //Specifies the unix end time for querying on email's received time.
	ReceivedStartTime     *int64    `json:"receivedStartTime,omitempty" form:"receivedStartTime,omitempty"`         //Specifies the unix start time for querying on email's received time.
	ReceivedTimeSeconds   *int64    `json:"receivedTimeSeconds,omitempty" form:"receivedTimeSeconds,omitempty"`     //Specifies the unix time when the email was received.
	RecipientAddresses    *[]string `json:"recipientAddresses,omitempty" form:"recipientAddresses,omitempty"`       //Specifies the email addresses of the recipients.
	SenderAddress         *string   `json:"senderAddress,omitempty" form:"senderAddress,omitempty"`                 //Specifies the email address of the sender.
	SentTimeSeconds       *int64    `json:"sentTimeSeconds,omitempty" form:"sentTimeSeconds,omitempty"`             //Specifies the unix time when the email was sent.
	ShowOnlyEmailFolders  *bool     `json:"showOnlyEmailFolders,omitempty" form:"showOnlyEmailFolders,omitempty"`   //Specifies whether the query result should include only Email folders.
	TenantId              *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                           //TenantId specifies the tenant whose action resulted in the audit log.
}

/*
 * Structure for the custom type EncryptionConfiguration
 */
type EncryptionConfiguration struct {
	EnableEncryption *bool  `json:"enableEncryption,omitempty" form:"enableEncryption,omitempty"` //Specifies whether or not to enable encryption. If encryption is enabled,
	EnableFipsMode   *bool  `json:"enableFipsMode,omitempty" form:"enableFipsMode,omitempty"`     //Specifies whether or not to enable FIPS mode. EnableEncryption must be
	RotationPeriod   *int64 `json:"rotationPeriod,omitempty" form:"rotationPeriod,omitempty"`     //Specifies the rotation period for encryption keys in days.
}

/*
 * Structure for the custom type EntityIdentifier
 */
type EntityIdentifier struct {
	EntityId *Value `json:"entityId,omitempty" form:"entityId,omitempty"` //Specifies a data type and data field used to store data.
}

/*
 * Structure for the custom type EntityPermissionInformation
 */
type EntityPermissionInformation struct {
	EntityId *int64       `json:"entityId,omitempty" form:"entityId,omitempty"` //Specifies the entity id.
	Groups   []*GroupInfo `json:"groups,omitempty" form:"groups,omitempty"`     //Specifies groups that have access to entity in case of restricted user.
	Tenant   *TenantInfo  `json:"tenant,omitempty" form:"tenant,omitempty"`     //Specifies struct with basic tenant details.
	Users    []*UserInfo  `json:"users,omitempty" form:"users,omitempty"`       //Specifies users that have access to entity in case of restricted user.
}

/*
 * Structure for the custom type EntityProto
 */
type EntityProto struct {
	AttributeVec    []*KeyValuePair   `json:"attributeVec,omitempty" form:"attributeVec,omitempty"`       //Array of Attributes.
	EntityId        *EntityIdentifier `json:"entityId,omitempty" form:"entityId,omitempty"`               //Specifies a unique identifier for the entity.
	LatestMetricVec []*MetricValue    `json:"latestMetricVec,omitempty" form:"latestMetricVec,omitempty"` //Array of Metric Statistics.
}

/*
 * Structure for the custom type EntitySchemaProto
 */
type EntitySchemaProto struct {
	AttributesDescriptor     *EntitySchemaProtoAttributesDescriptor   `json:"attributesDescriptor,omitempty" form:"attributesDescriptor,omitempty"`         //Specifies a list of attributes about an entity.
	FlushIntervalSecs        *int64                                   `json:"flushIntervalSecs,omitempty" form:"flushIntervalSecs,omitempty"`               //Defines the interval used to flush in memory stats to scribe table.
	IsInternalSchema         *bool                                    `json:"isInternalSchema,omitempty" form:"isInternalSchema,omitempty"`                 //Specifies if this schema should be displayed in Advanced Diagnostics
	LargestFlushIntervalSecs *int64                                   `json:"largestFlushIntervalSecs,omitempty" form:"largestFlushIntervalSecs,omitempty"` //Use can change the flush interval secs via gflag and this store the
	Name                     *string                                  `json:"name,omitempty" form:"name,omitempty"`                                         //Specifies a name that uniquely identifies an entity schema such as
	SchemaDescriptiveName    *string                                  `json:"schemaDescriptiveName,omitempty" form:"schemaDescriptiveName,omitempty"`       //Specifies the name of the Schema as displayed in Advanced Diagnostics
	SchemaHelpText           *string                                  `json:"schemaHelpText,omitempty" form:"schemaHelpText,omitempty"`                     //Specifies an optional informational description about the schema.
	TimeSeriesDescriptorVec  []*EntitySchemaProtoTimeSeriesDescriptor `json:"timeSeriesDescriptorVec,omitempty" form:"timeSeriesDescriptorVec,omitempty"`   //Array of Time Series.
	Version                  *int64                                   `json:"version,omitempty" form:"version,omitempty"`                                   //Specifies the version of the entity schema.
}

/*
 * Structure for the custom type EntitySchemaProtoAttributesDescriptor
 */
type EntitySchemaProtoAttributesDescriptor struct {
	AttributeVec          []*EntitySchemaProtoKeyValueDescriptor `json:"attributeVec,omitempty" form:"attributeVec,omitempty"`                   //Array of Attributes.
	KeyAttributeNameIndex *int64                                 `json:"keyAttributeNameIndex,omitempty" form:"keyAttributeNameIndex,omitempty"` //Specifies the attribute to use as a unique identifier for the entity.
}

/*
 * Structure for the custom type EntitySchemaProtoKeyValueDescriptor
 */
type EntitySchemaProtoKeyValueDescriptor struct {
	KeyName   *string `json:"keyName,omitempty" form:"keyName,omitempty"`     //Specifies the name of a key.
	ValueType *int64  `json:"valueType,omitempty" form:"valueType,omitempty"` //Specifies the type of the value that is associated with the key.
}

/*
 * Structure for the custom type EntitySchemaProtoTimeSeriesDescriptor
 */
type EntitySchemaProtoTimeSeriesDescriptor struct {
	MetricDescriptiveName            *string                                          `json:"metricDescriptiveName,omitempty" form:"metricDescriptiveName,omitempty"`                       //Specifies a descriptive name for the metric that is displayed in the
	MetricName                       *string                                          `json:"metricName,omitempty" form:"metricName,omitempty"`                                             //Specifies the name of the metric such as 'kUnmorphedUsageBytes'.
	MetricUnit                       *EntitySchemaProtoTimeSeriesDescriptorMetricUnit `json:"metricUnit,omitempty" form:"metricUnit,omitempty"`                                             //Specifies the unit of measure for the metric.
	RawMetricPublishIntervalHintSecs *int64                                           `json:"rawMetricPublishIntervalHintSecs,omitempty" form:"rawMetricPublishIntervalHintSecs,omitempty"` //Specifies a suggestion for the interval to collect raw data points.
	TimeToLiveSecs                   *int64                                           `json:"timeToLiveSecs,omitempty" form:"timeToLiveSecs,omitempty"`                                     //Specifies how long the data point will be stored.
	ValueType                        *int64                                           `json:"valueType,omitempty" form:"valueType,omitempty"`                                               //Specifies the value type for this metric.
}

/*
 * Structure for the custom type EntitySchemaProtoTimeSeriesDescriptorMetricUnit
 */
type EntitySchemaProtoTimeSeriesDescriptorMetricUnit struct {
	Type *int64 `json:"type,omitempty" form:"type,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type EnvBackupParams
 */
type EnvBackupParams struct {
	FileStubbingParams    *FileStubbingParams      `json:"fileStubbingParams,omitempty" form:"fileStubbingParams,omitempty"`       //File Stubbing Parameters
	HypervBackupParams    *HypervBackupEnvParams   `json:"hypervBackupParams,omitempty" form:"hypervBackupParams,omitempty"`       //Message to capture any additional backup params for a HyperV environment.
	NasBackupParams       *NasBackupParams         `json:"nasBackupParams,omitempty" form:"nasBackupParams,omitempty"`             //Message to capture any additional backup params for a NAS environment.
	O365BackupParams      *O365BackupEnvParams     `json:"o365BackupParams,omitempty" form:"o365BackupParams,omitempty"`           //TODO: Write general description for this field
	OutlookBackupParams   *OutlookBackupEnvParams  `json:"outlookBackupParams,omitempty" form:"outlookBackupParams,omitempty"`     //Message to capture any additional backup params for an Outlook environment.
	PhysicalBackupParams  *PhysicalBackupEnvParams `json:"physicalBackupParams,omitempty" form:"physicalBackupParams,omitempty"`   //Message to capture any additional backup params for a Physical environment.
	SnapshotManagerParams *SnapshotManagerParams   `json:"snapshotManagerParams,omitempty" form:"snapshotManagerParams,omitempty"` //TODO: Write general description for this field
	SqlBackupJobParams    *SqlBackupJobParams      `json:"sqlBackupJobParams,omitempty" form:"sqlBackupJobParams,omitempty"`       //Message to capture additional backup job params specific to SQL.
	VmwareBackupParams    *VmwareBackupEnvParams   `json:"vmwareBackupParams,omitempty" form:"vmwareBackupParams,omitempty"`       //Message to capture any additional backup params for a VMware environment.
}

/*
 * Structure for the custom type EnvironmentTypeJobParameters
 */
type EnvironmentTypeJobParameters struct {
	AwsSnapshotParameters *AwsSnapshotManagerParameters `json:"awsSnapshotParameters,omitempty" form:"awsSnapshotParameters,omitempty"` //Protection Job parameters applicable to 'kAWSSnapshotManager' Environment
	HypervParameters      *HypervEnvJobParameters       `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"`           //Specifies job parameters applicable for all 'kHyperV' Environment type
	NasParameters         *NasEnvJobParameters          `json:"nasParameters,omitempty" form:"nasParameters,omitempty"`                 //Specifies job parameters applicable for all 'kGenericNas' Environment type
	OutlookParameters     *OutlookEnvJobParameters      `json:"outlookParameters,omitempty" form:"outlookParameters,omitempty"`         //Specifies job parameters applicable for all 'kO365Outlook' Environment type
	PhysicalParameters    *PhysicalEnvJobParameters     `json:"physicalParameters,omitempty" form:"physicalParameters,omitempty"`       //Protection Job parameters applicable to 'kPhysical' Environment type.
	PureParameters        *PureEnvJobParameters         `json:"pureParameters,omitempty" form:"pureParameters,omitempty"`               //Specifies job parameters applicable for all 'kPure' Environment type
	SqlParameters         *SqlEnvJobParameters          `json:"sqlParameters,omitempty" form:"sqlParameters,omitempty"`                 //Specifies job parameters applicable for all 'kSQL' Environment type
	VmwareParameters      *VmwareEnvJobParameters       `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"`           //Specifies job parameters applicable for all 'kVMware' Environment type
}

/*
 * Structure for the custom type ErasureCodingInfo
 */
type ErasureCodingInfo struct {
	Algorithm            AlgorithmEnum `json:"algorithm,omitempty" form:"algorithm,omitempty"`                       //Algorthm used for erasure coding.
	ErasureCodingEnabled *bool         `json:"erasureCodingEnabled,omitempty" form:"erasureCodingEnabled,omitempty"` //Specifies whether Erasure coding is enabled on the Storage Domain
	InlineErasureCoding  *bool         `json:"inlineErasureCoding,omitempty" form:"inlineErasureCoding,omitempty"`   //Specifies if erasure coding should occur inline (as the data is being
	NumCodedStripes      *int64        `json:"numCodedStripes,omitempty" form:"numCodedStripes,omitempty"`           //The number of coded stripes.
	NumDataStripes       *int64        `json:"numDataStripes,omitempty" form:"numDataStripes,omitempty"`             //The number of stripes containing data.
}

/*
 * Structure for the custom type ErrorProto
 */
type ErrorProto struct {
	ErrorMsg *string `json:"errorMsg,omitempty" form:"errorMsg,omitempty"` //An optional detail.
	Type     *int64  `json:"type,omitempty" form:"type,omitempty"`         //Error.
}

/*
 * Structure for the custom type EulaConfig
 */
type EulaConfig struct {
	LicenseKey    string  `json:"licenseKey" form:"licenseKey"`                         //Specifies the license key.
	SignedByUser  *string `json:"signedByUser,omitempty" form:"signedByUser,omitempty"` //Specifies the login account name for the Cohesity user who accepted
	SignedTime    *int64  `json:"signedTime,omitempty" form:"signedTime,omitempty"`     //Specifies the time that the End User License Agreement was accepted.
	SignedVersion int64   `json:"signedVersion" form:"signedVersion"`                   //Specifies the version of the End User License Agreement that was accepted.
}

/*
 * Structure for the custom type ExpandCloudClusterParameters
 */
type ExpandCloudClusterParameters struct {
	NodeIps []string `json:"nodeIps" form:"nodeIps"` //Specifies the list of IPs of the new Nodes.
}

/*
 * Structure for the custom type ExpandPhysicalClusterParameters
 */
type ExpandPhysicalClusterParameters struct {
	NodeConfigs []*PhysicalNodeConfiguration `json:"nodeConfigs" form:"nodeConfigs"`       //Specifies the configuration details of the Nodes in the Cluster.
	Vips        *[]string                    `json:"vips,omitempty" form:"vips,omitempty"` //Specifies the virtual IPs to add to the Cluster.
}

/*
 * Structure for the custom type ExtendedRetentionPolicy
 */
type ExtendedRetentionPolicy struct {
	BackupRunType BackupRunTypeEnum                      `json:"backupRunType,omitempty" form:"backupRunType,omitempty"` //The backup run type to which this extended retention applies to. If this is
	DaysToKeep    *int64                                 `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`       //Specifies the number of days to retain copied Snapshots on the target.
	Multiplier    *int64                                 `json:"multiplier,omitempty" form:"multiplier,omitempty"`       //Specifies a factor to multiply the periodicity by, to determine the copy
	Periodicity   PeriodicityExtendedRetentionPolicyEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"`     //Specifies the frequency that Snapshots should be copied to the
}

/*
 * Structure for the custom type ExternalClientSubnets
 */
type ExternalClientSubnets struct {
	ClientSubnets []*Subnet `json:"clientSubnets,omitempty" form:"clientSubnets,omitempty"` //Specifies the Client Subnets for the cluster.
}

/*
 * Structure for the custom type FileDistributionMetrics
 */
type FileDistributionMetrics struct {
	MetricName *string `json:"metricName,omitempty" form:"metricName,omitempty"` //Specifies the name of the metric.
	Value      *int64  `json:"value,omitempty" form:"value,omitempty"`           //Specifies the value of specified metric name.
}

/*
 * Structure for the custom type FileDistributionStats
 */
type FileDistributionStats struct {
	ClusterId            *int64                     `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //Specifies the cluster Id.
	ClusterIncarnationId *int64                     `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the cluster Incarnation Id.
	EntityId             *int64                     `json:"entityId,omitempty" form:"entityId,omitempty"`                         //Specifies the id of the entity for which file distribution stats are computed.
	EntityName           *string                    `json:"entityName,omitempty" form:"entityName,omitempty"`                     //Specifies the name of the entity for which file distribution stats are computed.
	MetricsList          []*FileDistributionMetrics `json:"metricsList,omitempty" form:"metricsList,omitempty"`                   //Specifies the list of file stats for different file extensions.
}

/*
 * Structure for the custom type FileExtensionFilter
 */
type FileExtensionFilter struct {
	FileExtensionsList *[]string                   `json:"fileExtensionsList,omitempty" form:"fileExtensionsList,omitempty"` //The list of file extensions to apply
	IsEnabled          *bool                       `json:"isEnabled,omitempty" form:"isEnabled,omitempty"`                   //If set, it enables the file extension filter
	Mode               ModeFileExtensionFilterEnum `json:"mode,omitempty" form:"mode,omitempty"`                             //The mode applied to the list of file extensions
}

/*
 * Structure for the custom type FileId
 */
type FileId struct {
	EntityId    *int64 `json:"entityId,omitempty" form:"entityId,omitempty"`       //Specifies the entity id of the file.
	RootInodeId *int64 `json:"rootInodeId,omitempty" form:"rootInodeId,omitempty"` //Specifies the root inode id of the file system that file belongs to.
	ViewId      *int64 `json:"viewId,omitempty" form:"viewId,omitempty"`           //Specifies the id of the View the file belongs to.
}

/*
 * Structure for the custom type FileLevelDataLockConfig
 */
type FileLevelDataLockConfig struct {
	AutoLockAfterDurationIdle         *int64                          `json:"autoLockAfterDurationIdle,omitempty" form:"autoLockAfterDurationIdle,omitempty"`                 //Specifies the duration to lock a file that has not been accessed or
	DefaultFileRetentionDurationMsecs *int64                          `json:"defaultFileRetentionDurationMsecs,omitempty" form:"defaultFileRetentionDurationMsecs,omitempty"` //Specifies a global default retention duration for files in this view, if
	ExpiryTimestampMsecs              *int64                          `json:"expiryTimestampMsecs,omitempty" form:"expiryTimestampMsecs,omitempty"`                           //Specifies a definite timestamp in milliseconds for retaining the file.
	LockingProtocol                   LockingProtocolEnum             `json:"lockingProtocol,omitempty" form:"lockingProtocol,omitempty"`                                     //Specifies the supported mechanisms to explicity lock a file from NFS/SMB
	MaxRetentionDurationMsecs         *int64                          `json:"maxRetentionDurationMsecs,omitempty" form:"maxRetentionDurationMsecs,omitempty"`                 //Specifies a maximum duration in milliseconds for which any file in this
	MinRetentionDurationMsecs         *int64                          `json:"minRetentionDurationMsecs,omitempty" form:"minRetentionDurationMsecs,omitempty"`                 //Specifies a minimum retention duration in milliseconds after a file gets
	Mode                              ModeFileLevelDataLockConfigEnum `json:"mode,omitempty" form:"mode,omitempty"`                                                           //Specifies the mode of file level datalock.
}

/*
 * Structure for the custom type FileLockStatus
 */
type FileLockStatus struct {
	ExpiryTimestampMsecs *int64                 `json:"expiryTimestampMsecs,omitempty" form:"expiryTimestampMsecs,omitempty"` //Specifies a expiry timestamp in milliseconds until the file is locked.
	HoldTimestampMsecs   *int64                 `json:"holdTimestampMsecs,omitempty" form:"holdTimestampMsecs,omitempty"`     //Specifies a override timestamp in milliseconds when an expired file is
	LockTimestampMsecs   *int64                 `json:"lockTimestampMsecs,omitempty" form:"lockTimestampMsecs,omitempty"`     //Specifies the timestamp at which the file was locked.
	Mode                 ModeFileLockStatusEnum `json:"mode,omitempty" form:"mode,omitempty"`                                 //Specifies the mode of the file lock. 'kCompliance', 'kEnterprise'.
	State                *int64                 `json:"state,omitempty" form:"state,omitempty"`                               //Specifies the lock state of the file.
}

/*
 * Structure for the custom type FileNlmLocks
 */
type FileNlmLocks struct {
	FileId   *FileId    `json:"fileId,omitempty" form:"fileId,omitempty"`     //TODO: Write general description for this field
	NlmLocks []*NlmLock `json:"nlmLocks,omitempty" form:"nlmLocks,omitempty"` //Specifies the list of NLM locks in a view.
}

/*
 * Structure for the custom type FilePartitionBlock
 */
type FilePartitionBlock struct {
	DiskFileName *string `json:"diskFileName,omitempty" form:"diskFileName,omitempty"` //Specifies the disk file name where the logical partition is.
	LengthBytes  *int64  `json:"lengthBytes,omitempty" form:"lengthBytes,omitempty"`   //Specifies the length of the block in bytes.
	Number       *int64  `json:"number,omitempty" form:"number,omitempty"`             //Specifies a unique number of the partition within the linear disk file.
	OffsetBytes  *int64  `json:"offsetBytes,omitempty" form:"offsetBytes,omitempty"`   //Specifies the offset of the block (in bytes) from the beginning
}

/*
 * Structure for the custom type FilePathFilter
 */
type FilePathFilter struct {
	ExcludeFilters *[]string `json:"excludeFilters,omitempty" form:"excludeFilters,omitempty"` //Array of Excluded File Path Filters.
	ProtectFilters *[]string `json:"protectFilters,omitempty" form:"protectFilters,omitempty"` //Array of Protected File Path Filters.
}

/*
 * Structure for the custom type FilePathParameters
 */
type FilePathParameters struct {
	BackupFilePath    *string   `json:"backupFilePath,omitempty" form:"backupFilePath,omitempty"`       //Specifies absolute path to a file or a directory in a Physical Server
	ExcludedFilePaths *[]string `json:"excludedFilePaths,omitempty" form:"excludedFilePaths,omitempty"` //Array of Excluded File Paths.
	SkipNestedVolumes *bool     `json:"skipNestedVolumes,omitempty" form:"skipNestedVolumes,omitempty"` //Specifies if any subdirectories under backupFilePath, where local or
}

/*
 * Structure for the custom type FileRestoreInfo
 */
type FileRestoreInfo struct {
	Error            *RequestError     `json:"error,omitempty" form:"error,omitempty"`                       //Details about the Error.
	Filename         *string           `json:"filename,omitempty" form:"filename,omitempty"`                 //Specifies the path of the file/directory.
	FilesystemVolume *FilesystemVolume `json:"filesystemVolume,omitempty" form:"filesystemVolume,omitempty"` //Specifies information about a filesystem volume.
	IsFolder         *bool             `json:"isFolder,omitempty" form:"isFolder,omitempty"`                 //Specifies whether the file path is a folder.
}

/*
 * Structure for the custom type FileSearchResult
 */
type FileSearchResult struct {
	AdObjectMetaData   *AdObjectMetaData        `json:"adObjectMetaData,omitempty" form:"adObjectMetaData,omitempty"`     //Specifies details about the AD objects.
	DocumentType       *string                  `json:"documentType,omitempty" form:"documentType,omitempty"`             //Specifies the inferred document type.
	EmailMetaData      *EmailMetaData           `json:"emailMetaData,omitempty" form:"emailMetaData,omitempty"`           //Specifies details about the emails and the folder containing emails.
	FileVersions       []*FileVersion           `json:"fileVersions,omitempty" form:"fileVersions,omitempty"`             //Array of File Versions.
	Filename           *string                  `json:"filename,omitempty" form:"filename,omitempty"`                     //Specifies the name of the found file or folder.
	IsFolder           *bool                    `json:"isFolder,omitempty" form:"isFolder,omitempty"`                     //Specifies if the found item is a folder.
	JobId              *int64                   `json:"jobId,omitempty" form:"jobId,omitempty"`                           //Specifies the Job id for the Protection Job that is currently
	JobUid             *UniversalId             `json:"jobUid,omitempty" form:"jobUid,omitempty"`                         //Specifies the universal id of the Protection Job that backed up
	ProtectionSource   *ProtectionSource        `json:"protectionSource,omitempty" form:"protectionSource,omitempty"`     //Specifies a generic structure that represents a node
	RegisteredSourceId *int64                   `json:"registeredSourceId,omitempty" form:"registeredSourceId,omitempty"` //Specifies the id of the top-level registered source (such as a
	SourceId           *int64                   `json:"sourceId,omitempty" form:"sourceId,omitempty"`                     //Specifies the source id of the object that contains the file or folder.
	Type               TypeFileSearchResultEnum `json:"type,omitempty" form:"type,omitempty"`                             //Specifies the type of the file document such as KDirectory, kFile, etc.
	ViewBoxId          *int64                   `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                   //Specifies the id of the Domain (View Box) where the source object that
}

/*
 * Structure for the custom type FileSearchResults
 */
type FileSearchResults struct {
	Files      []*FileSearchResult `json:"files,omitempty" form:"files,omitempty"`           //Array of Files and Folders.
	TotalCount *int64              `json:"totalCount,omitempty" form:"totalCount,omitempty"` //Specifies the total number of files and folders that match the filter and
}

/*
 * Structure for the custom type FileSnapshotInformation
 */
type FileSnapshotInformation struct {
	HasArchivalCopy   *bool            `json:"hasArchivalCopy,omitempty" form:"hasArchivalCopy,omitempty"`     //If true, this snapshot is located on an archival target
	HasLocalCopy      *bool            `json:"hasLocalCopy,omitempty" form:"hasLocalCopy,omitempty"`           //If true, this snapshot is located on a local Cohesity Cluster.
	HasRemoteCopy     *bool            `json:"hasRemoteCopy,omitempty" form:"hasRemoteCopy,omitempty"`         //If true, this snapshot is located on a Remote Cohesity Cluster.
	ModifiedTimeUsecs *int64           `json:"modifiedTimeUsecs,omitempty" form:"modifiedTimeUsecs,omitempty"` //Specifies the time when the file or folder was last modified.
	SizeBytes         *int64           `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"`                 //Specifies the size of the file or folder in bytes.
	Snapshot          *SnapshotAttempt `json:"snapshot,omitempty" form:"snapshot,omitempty"`                   //Specifies information about a single snapshot.
}

/*
 * Structure for the custom type FileStubbingParams
 */
type FileStubbingParams struct {
	ColdFileWindow   *int64                `json:"coldFileWindow,omitempty" form:"coldFileWindow,omitempty"`     //Identifies the cold files in the NAS source. Files that haven't been
	FileSelectPolicy *int64                `json:"fileSelectPolicy,omitempty" form:"fileSelectPolicy,omitempty"` //File migrate policy based on file access/modify time and age.
	FileSize         *int64                `json:"fileSize,omitempty" form:"fileSize,omitempty"`                 //Gives the size criteria to be used for selecting the files to be migrated.
	FileSizePolicy   *int64                `json:"fileSizePolicy,omitempty" form:"fileSizePolicy,omitempty"`     //File size policy for selecting files to migrate.
	FilteringPolicy  *FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"`   //Proto to encapsulate the filtering policy for backup objects like files or
	NfsMountPath     *string               `json:"nfsMountPath,omitempty" form:"nfsMountPath,omitempty"`         //Mount path where the Cohesity target view must be mounted on all
	TargetViewName   *string               `json:"targetViewName,omitempty" form:"targetViewName,omitempty"`     //The target view name to which the data will be migrated.
}

/*
 * Structure for the custom type FileVersion
 */
type FileVersion struct {
	ModifiedTimeUsecs *int64             `json:"modifiedTimeUsecs,omitempty" form:"modifiedTimeUsecs,omitempty"` //Specifies the time when the file or folder was last modified.
	SizeBytes         *int64             `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"`                 //Specifies the size of the file or folder (in bytes)
	Snapshots         []*SnapshotAttempt `json:"snapshots,omitempty" form:"snapshots,omitempty"`                 //Array of Snapshots.
}

/*
 * Structure for the custom type FilenamePatternToDirectory
 */
type FilenamePatternToDirectory struct {
	Directory       *string `json:"directory,omitempty" form:"directory,omitempty"`             //Specifies the directory where to keep the files matching the pattern.
	FilenamePattern *string `json:"filenamePattern,omitempty" form:"filenamePattern,omitempty"` //Specifies a pattern to be matched with filenames. This can be a
}

/*
 * Structure for the custom type FilerAuditLogConfiguration
 */
type FilerAuditLogConfiguration struct {
	Enabled             bool  `json:"enabled" form:"enabled"`                         //Specifies if filer audit logging is enabled on the Cohesity Cluster.
	RetentionPeriodDays int64 `json:"retentionPeriodDays" form:"retentionPeriodDays"` //Specifies the number of days to keep (retain) the filer audit logs.
}

/*
 * Structure for the custom type FilesAndFoldersInfo
 */
type FilesAndFoldersInfo struct {
	AbsolutePath *string `json:"absolutePath,omitempty" form:"absolutePath,omitempty"` //AbsolutePath specifies the absolute path of the specified file or folder.
	IsDirectory  *bool   `json:"isDirectory,omitempty" form:"isDirectory,omitempty"`   //IsDirectory specifies if specified object is a directory or not.
}

/*
 * Structure for the custom type FilesToDirectoryMapping
 */
type FilesToDirectoryMapping struct {
	FilePattern     *string `json:"filePattern,omitempty" form:"filePattern,omitempty"`         //Source file name. The file name can be a regex matching source files.
	TargetDirectory *string `json:"targetDirectory,omitempty" form:"targetDirectory,omitempty"` //Target directtory for the source file pattern.
}

/*
 * Structure for the custom type FilesystemVolume
 */
type FilesystemVolume struct {
	Disks             []*Disk               `json:"disks,omitempty" form:"disks,omitempty"`                         //Array of Disks and Partitions.
	DisplayName       *string               `json:"displayName,omitempty" form:"displayName,omitempty"`             //Specifies a description about the filesystem.
	FilesystemType    *string               `json:"filesystemType,omitempty" form:"filesystemType,omitempty"`       //Specifies type of the filesystem on this volume.
	FilesystemUuid    *string               `json:"filesystemUuid,omitempty" form:"filesystemUuid,omitempty"`       //Specifies the uuid of the filesystem.
	IsSupported       *bool                 `json:"isSupported,omitempty" form:"isSupported,omitempty"`             //If true, this is a supported filesystem volume type.
	LogicalVolume     *LogicalVolume        `json:"logicalVolume,omitempty" form:"logicalVolume,omitempty"`         //Specify attributes for a kLMV (Linux) or kLDM (Windows) filesystem.
	LogicalVolumeType LogicalVolumeTypeEnum `json:"logicalVolumeType,omitempty" form:"logicalVolumeType,omitempty"` //Specifies the type of logical volume such as kSimpleVolume, kLVM or kLDM.
	Name              *string               `json:"name,omitempty" form:"name,omitempty"`                           //Specifies the name of the volume such as /C.
	VolumeGuid        *string               `json:"volumeGuid,omitempty" form:"volumeGuid,omitempty"`               //VolumeGuid is the Volume guid.
}

/*
 * Structure for the custom type FilteringPolicyProto
 */
type FilteringPolicyProto struct {
	AllowFilters *[]string `json:"allowFilters,omitempty" form:"allowFilters,omitempty"` //List of filters to allow matched objects for backup.
	DenyFilters  *[]string `json:"denyFilters,omitempty" form:"denyFilters,omitempty"`   //List of filters to deny matched objects for backup.
}

/*
 * Structure for the custom type FixedUnixIdMapping
 */
type FixedUnixIdMapping struct {
	Gid *int64 `json:"gid,omitempty" form:"gid,omitempty"` //Specifies the fixed Unix GID, when mapping type is set to kFixed.
	Uid *int64 `json:"uid,omitempty" form:"uid,omitempty"` //Specifies the fixed Unix UID, when mapping type is set to kFixed.
}

/*
 * Structure for the custom type FlashBladeFileSystem
 */
type FlashBladeFileSystem struct {
	BackupEnabled        *bool              `json:"backupEnabled,omitempty" form:"backupEnabled,omitempty"`               //Specifies whether the .snapshot directory exists on the file system.
	CreatedTimeMsecs     *int64             `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`         //Specifies the time when the filesystem was created in Unix epoch time
	LogicalCapacityBytes *int64             `json:"logicalCapacityBytes,omitempty" form:"logicalCapacityBytes,omitempty"` //Specifies the total capacity in bytes of the file system.
	LogicalUsedBytes     *int64             `json:"logicalUsedBytes,omitempty" form:"logicalUsedBytes,omitempty"`         //Specifies the size of logical data currently represented on the
	NfsInfo              *FlashBladeNfsInfo `json:"nfsInfo,omitempty" form:"nfsInfo,omitempty"`                           //Specifies information specific to NFS protocol exposed by Pure Flash Blade
	PhysicalUsedBytes    *int64             `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"`       //Specifies the size of physical data currently consumed by the file
	Protocols            *[]ProtocolEnum    `json:"protocols,omitempty" form:"protocols,omitempty"`                       //List of Protocols.
	SmbInfo              *FlashBladeSmbInfo `json:"smbInfo,omitempty" form:"smbInfo,omitempty"`                           //Specifies information specific to SMB shares exposed by Pure Flash Blade
	UniqueUsedBytes      *int64             `json:"uniqueUsedBytes,omitempty" form:"uniqueUsedBytes,omitempty"`           //Specifies the size of physical data cconsumed by the file system
}

/*
 * Structure for the custom type FlashBladeNetworkInterface
 */
type FlashBladeNetworkInterface struct {
	IpAddress *string `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the IP address of the Pure Storage FlashBlade Array.
	Name      *string `json:"name,omitempty" form:"name,omitempty"`           //Specifies the name of the network interface.
	Vlan      *int64  `json:"vlan,omitempty" form:"vlan,omitempty"`           //Specifies the id of the VLAN network of the Pure Storage FlashBlade Array.
}

/*
 * Structure for the custom type FlashBladeNfsInfo
 */
type FlashBladeNfsInfo struct {
	ExportRules *string `json:"exportRules,omitempty" form:"exportRules,omitempty"` //Specifies NFS protocol export rules. Rules are in the form host(options).
}

/*
 * Structure for the custom type FlashBladeProtectionSource
 */
type FlashBladeProtectionSource struct {
	FileSystem   *FlashBladeFileSystem              `json:"fileSystem,omitempty" form:"fileSystem,omitempty"`     //Specifies information about a Flash Blade File System in a Storage Array.
	Name         *string                            `json:"name,omitempty" form:"name,omitempty"`                 //Specifies a unique name of the Protection Source.
	StorageArray *FlashBladeStorageArray            `json:"storageArray,omitempty" form:"storageArray,omitempty"` //Specifies information about a Pure Storage FlashBlade Array.
	Type         TypeFlashBladeProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                 //Specifies the type of managed object in a Pure Storage FlashBlade
}

/*
 * Structure for the custom type FlashBladeSmbInfo
 */
type FlashBladeSmbInfo struct {
	AclMode AclModeEnum `json:"aclMode,omitempty" form:"aclMode,omitempty"` //ACL mode for this SMB share.
}

/*
 * Structure for the custom type FlashBladeStorageArray
 */
type FlashBladeStorageArray struct {
	CapacityBytes     *int64                        `json:"capacityBytes,omitempty" form:"capacityBytes,omitempty"`         //Specifies the total capacity in bytes of the Pure Storage FlashBlade
	Id                *string                       `json:"id,omitempty" form:"id,omitempty"`                               //Specifies a unique id of a Pure Storage FlashBlade Array.
	Networks          []*FlashBladeNetworkInterface `json:"networks,omitempty" form:"networks,omitempty"`                   //Specifies the network interfaces of the Pure Storage FlashBlade Array.
	PhysicalUsedBytes *int64                        `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"` //Specifies the space used for physical data in bytes.
	Revision          *string                       `json:"revision,omitempty" form:"revision,omitempty"`                   //Specifies the revision of the Pure Storage FlashBlade software.
	Version           *string                       `json:"version,omitempty" form:"version,omitempty"`                     //Specifies the software version running on the Pure Storage FlashBlade
}

/*
 * Structure for the custom type FreeNodeInformation
 */
type FreeNodeInformation struct {
	ChassisSerial     *string `json:"chassisSerial,omitempty" form:"chassisSerial,omitempty"`         //Specifies the serial number of the Chassis the Node is installed in.
	ConnectedTo       *bool   `json:"connectedTo,omitempty" form:"connectedTo,omitempty"`             //Specifies whether or not this is the Node that is sending the response.
	Id                *int64  `json:"id,omitempty" form:"id,omitempty"`                               //Specifies the ID of the node.
	Ip                *string `json:"ip,omitempty" form:"ip,omitempty"`                               //Specifies the IP address of the Node.
	IpmiIp            *string `json:"ipmiIp,omitempty" form:"ipmiIp,omitempty"`                       //Specifies the IPMI IP of the Node.
	NodeSerial        *string `json:"nodeSerial,omitempty" form:"nodeSerial,omitempty"`               //Specifies the serial number of the Node.
	NodeUiSlot        *string `json:"nodeUiSlot,omitempty" form:"nodeUiSlot,omitempty"`               //Specifies the postition for the UI to display the Node in the Cluster
	NumSlotsInChassis *int64  `json:"numSlotsInChassis,omitempty" form:"numSlotsInChassis,omitempty"` //Specifies the number of Node slots present in the Chassis where this
	SlotNumber        *string `json:"slotNumber,omitempty" form:"slotNumber,omitempty"`               //Specifies the number of the slot the Node is installed in.
	SoftwareVersion   *string `json:"softwareVersion,omitempty" form:"softwareVersion,omitempty"`     //Specifies the version of the software installed on the Node.
}

/*
 * Structure for the custom type FullSnapshotInfo
 */
type FullSnapshotInfo struct {
	RestoreInfo    *RestoreInfo              `json:"restoreInfo,omitempty" form:"restoreInfo,omitempty"`       //Specifies the info regarding a full SQL snapshot.
	SnapshotTarget []*SnapshotTargetSettings `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"` //Specifies the location holding snapshot copies that may be used for
}

/*
 * Structure for the custom type GcpCredentials
 */
type GcpCredentials struct {
	ClientEmailAddress *string     `json:"clientEmailAddress,omitempty" form:"clientEmailAddress,omitempty"` //Specifies Client email address associated with the service account.
	ClientPrivateKey   *string     `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"`     //Specifies Client private associated with the service account.
	GcpType            GcpTypeEnum `json:"gcpType,omitempty" form:"gcpType,omitempty"`                       //Specifies the entity type such as 'kIAMUser' if the environment is kGCP.
	ProjectId          *string     `json:"projectId,omitempty" form:"projectId,omitempty"`                   //Specifies Id of the project associated with Google cloud account.
	VpcNetwork         *string     `json:"vpcNetwork,omitempty" form:"vpcNetwork,omitempty"`                 //Specifies the VPC Network to deploy proxy VMs.
	VpcSubnetwork      *string     `json:"vpcSubnetwork,omitempty" form:"vpcSubnetwork,omitempty"`           //Specifies the subnetwork to deploy proxy VMs.
}

/*
 * Structure for the custom type GcpParams
 */
type GcpParams struct {
	InstanceId            *int64 `json:"instanceId,omitempty" form:"instanceId,omitempty"`                       //Specfies id of the GCP instance type in which to deploy the VM.
	Region                *int64 `json:"region,omitempty" form:"region,omitempty"`                               //Specifies id of the GCP region in which to deploy the VM.
	SubnetId              *int64 `json:"subnetId,omitempty" form:"subnetId,omitempty"`                           //Specifies id of the subnet within above VPC.
	VirtualPrivateCloudId *int64 `json:"virtualPrivateCloudId,omitempty" form:"virtualPrivateCloudId,omitempty"` //Specifies id of the Virtual Private Cloud to chose for the instance type.
}

/*
 * Structure for the custom type GcpProtectionSource
 */
type GcpProtectionSource struct {
	ClientEmailAddress *string                     `json:"clientEmailAddress,omitempty" form:"clientEmailAddress,omitempty"` //Specifies Client email address associated with the service account.
	ClientPrivateKey   *string                     `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"`     //Specifies Client private associated with the service account.
	GcpType            GcpTypeEnum                 `json:"gcpType,omitempty" form:"gcpType,omitempty"`                       //Specifies the entity type such as 'kIAMUser' if the environment is kGCP.
	HostType           HostTypeEnum                `json:"hostType,omitempty" form:"hostType,omitempty"`                     //Specifies the OS type of the Protection Source of type 'kVirtualMachine'
	IpAddressesVM      *string                     `json:"ipAddressesVM,omitempty" form:"ipAddressesVM,omitempty"`           //Specifies the IP address of the entity of type 'kVirtualMachine'.
	Name               *string                     `json:"name,omitempty" form:"name,omitempty"`                             //Specifies the name of the Object set by the Cloud Provider.
	OwnerId            *string                     `json:"ownerId,omitempty" form:"ownerId,omitempty"`                       //Specifies the owner id of the resource in GCP environment. With type,
	PhysicalSourceId   *int64                      `json:"physicalSourceId,omitempty" form:"physicalSourceId,omitempty"`     //Specifies the Protection Source id of the registered Physical Host.
	ProjectId          *string                     `json:"projectId,omitempty" form:"projectId,omitempty"`                   //Specifies the project Id.
	RegionId           *string                     `json:"regionId,omitempty" form:"regionId,omitempty"`                     //Specifies the region Id.
	ResourceId         *string                     `json:"resourceId,omitempty" form:"resourceId,omitempty"`                 //Specifies the unique Id of the resource given by the cloud provider.
	RestoreTaskId      *int64                      `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"`           //Specifies the id of the "convert and deploy" restore task that
	Type               TypeGcpProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                             //Specifies the type of an GCP Protection Source Object such as
	VpcNetwork         *string                     `json:"vpcNetwork,omitempty" form:"vpcNetwork,omitempty"`                 //Specifies the VPC Network to deploy proxy VMs.
	VpcSubnetwork      *string                     `json:"vpcSubnetwork,omitempty" form:"vpcSubnetwork,omitempty"`           //Specifies the subnetwork to deploy proxy VMs.
}

/*
 * Structure for the custom type GdprCopyTask
 */
type GdprCopyTask struct {
	JobId           *int64  `json:"JobId,omitempty" form:"JobId,omitempty"`                     //Specifies the job with which this copy task is tied to.
	CloudTargetType *string `json:"cloudTargetType,omitempty" form:"cloudTargetType,omitempty"` //Specifies the cloud deploy target type. For example 'kAzure','kAWS',
	ExpiryTimeUsecs *int64  `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies the expiry of the copy task.
	TargetId        *int64  `json:"targetId,omitempty" form:"targetId,omitempty"`               //Specifies the id for the target.
	TargetName      *string `json:"targetName,omitempty" form:"targetName,omitempty"`           //Specifies the target of the replication or archival tasks.
	TotalSnapshots  *int64  `json:"totalSnapshots,omitempty" form:"totalSnapshots,omitempty"`   //Specifies the total number of snapshots.
	Type            *string `json:"type,omitempty" form:"type,omitempty"`                       //Specifies details about the Copy Run of a Job Run.
}

/*
 * Structure for the custom type GetAlertTypesParams
 */
type GetAlertTypesParams struct {
	AlertCategoryList *[]AlertCategoryListEnum `json:"alertCategoryList,omitempty" form:"alertCategoryList,omitempty"` //Specifies a list of Alert Categories to filter alert types by.
}

/*
 * Structure for the custom type GetConsumerStatsResult
 */
type GetConsumerStatsResult struct {
	Cookie    *string          `json:"cookie,omitempty" form:"cookie,omitempty"`       //Specifies an opaque string to pass to get the next set of active opens.
	StatsList []*ConsumerStats `json:"statsList,omitempty" form:"statsList,omitempty"` //Specifies a list of consumer stats.
}

/*
 * Structure for the custom type GetRegistrationInfoResponse
 */
type GetRegistrationInfoResponse struct {
	RootNodes  []*ProtectionSourceTreeInfo `json:"rootNodes,omitempty" form:"rootNodes,omitempty"`   //Specifies the registration, protection and permission information of either
	Stats      *ProtectionSummary          `json:"stats,omitempty" form:"stats,omitempty"`           //Specifies the sum of all the stats of protection of Protection Sources
	StatsByEnv []*ProtectionSummaryByEnv   `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"` //Specifies the breakdown of the stats by environment
}

/*
 * Structure for the custom type GetTenantStatsResult
 */
type GetTenantStatsResult struct {
	Cookie    *string        `json:"cookie,omitempty" form:"cookie,omitempty"`       //Specifies an opaque string to pass to get the next set of active opens.
	StatsList []*TenantStats `json:"statsList,omitempty" form:"statsList,omitempty"` //Specifies a list of tenant stats.
}

/*
 * Structure for the custom type GetViewBoxStatsResult
 */
type GetViewBoxStatsResult struct {
	StatsList []*StorageDomainStats `json:"statsList,omitempty" form:"statsList,omitempty"` //Specifies a list of view box stats.
}

/*
 * Structure for the custom type GetViewsByShareNameResult
 */
type GetViewsByShareNameResult struct {
	PaginationCookie *string  `json:"paginationCookie,omitempty" form:"paginationCookie,omitempty"` //If set, i.e. there are more results to display, use this value to get
	SharesList       []*Share `json:"sharesList,omitempty" form:"sharesList,omitempty"`             //Array of Views and Aliases by Share name.
}

/*
 * Structure for the custom type GetViewsResult
 */
type GetViewsResult struct {
	LastResult *bool   `json:"lastResult,omitempty" form:"lastResult,omitempty"` //If false, more Views are available to return. If the number of
	Views      []*View `json:"views,omitempty" form:"views,omitempty"`           //Array of Views.
}

/*
 * Structure for the custom type GoogleAccountInfo
 */
type GoogleAccountInfo struct {
	AccountId *string `json:"accountId,omitempty" form:"accountId,omitempty"` //Specifies the Account Id assigned by Google.
	UserId    *string `json:"userId,omitempty" form:"userId,omitempty"`       //Specifies the User Id assigned by Google.
}

/*
 * Structure for the custom type GoogleCloudCredentials
 */
type GoogleCloudCredentials struct {
	ClientEmailAddress *string                            `json:"clientEmailAddress,omitempty" form:"clientEmailAddress,omitempty"` //Specifies the client email address used to access Google
	ClientPrivateKey   *string                            `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"`     //Specifies the private key used to access Google Cloud Storage that is
	ProjectId          *string                            `json:"projectId,omitempty" form:"projectId,omitempty"`                   //Specifies the project id of an existing Google Cloud project to store
	TierType           TierTypeGoogleCloudCredentialsEnum `json:"tierType,omitempty" form:"tierType,omitempty"`                     //Specifies the storage class of GCP.
}

/*
 * Structure for the custom type GpfsCluster
 */
type GpfsCluster struct {
	CesAddresses  *[]string `json:"cesAddresses,omitempty" form:"cesAddresses,omitempty"`   //Specifies a list of CES(Cluster Export Services) IP addresses of a GPFS
	Id            *int64    `json:"id,omitempty" form:"id,omitempty"`                       //Specifies a globally unique id of a GPFS Cluster.
	PrimaryServer *string   `json:"primaryServer,omitempty" form:"primaryServer,omitempty"` //Specifies a primary server of a GPFS Cluster.
}

/*
 * Structure for the custom type GpfsFileset
 */
type GpfsFileset struct {
	Id        *int64                     `json:"id,omitempty" form:"id,omitempty"`               //Specifies the id of the fileset.
	Name      *string                    `json:"name,omitempty" form:"name,omitempty"`           //Name of the filesystem associated with the fileset
	Path      *string                    `json:"path,omitempty" form:"path,omitempty"`           //Specifies the absolute path of the fileset.
	Protocols *[]ProtocolGpfsFilesetEnum `json:"protocols,omitempty" form:"protocols,omitempty"` //Specifies GPFS supported Protocol information enabled on GPFS File System
}

/*
 * Structure for the custom type GpfsFilesystem
 */
type GpfsFilesystem struct {
	Id   *string `json:"id,omitempty" form:"id,omitempty"`     //Specifies the id of the file system.
	Path *string `json:"path,omitempty" form:"path,omitempty"` //Specifies the path of the file system.
}

/*
 * Structure for the custom type GpfsProtectionSource
 */
type GpfsProtectionSource struct {
	Cluster    *GpfsCluster                 `json:"cluster,omitempty" form:"cluster,omitempty"`       //Specifies information about a GPFS Cluster.
	Fileset    *GpfsFileset                 `json:"fileset,omitempty" form:"fileset,omitempty"`       //Specifies information about a fileset in a GPFS file system.
	Filesystem *GpfsFilesystem              `json:"filesystem,omitempty" form:"filesystem,omitempty"` //Specifies information about filesystem in a GPFS Cluster.
	Name       *string                      `json:"name,omitempty" form:"name,omitempty"`             //Specifies a unique name of the Protection Source.
	Type       TypeGpfsProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`             //Specifies the type of the entity in an GPFS file system
}

/*
 * Structure for the custom type GranularityBucket
 */
type GranularityBucket struct {
	Granularity *int64 `json:"granularity,omitempty" form:"granularity,omitempty"` //The base time period granularity that determines the frequency at which
	Multiplier  *int64 `json:"multiplier,omitempty" form:"multiplier,omitempty"`   //A factor to multiply the granularity by.
}

/*
 * Structure for the custom type Group
 */
type Group struct {
	CreatedTimeMsecs     *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the group was created/added.
	Description          *string         `json:"description,omitempty" form:"description,omitempty"`                   //Specifies a description of the group.
	Domain               *string         `json:"domain,omitempty" form:"domain,omitempty"`                             //Specifies the domain of the group.
	LastUpdatedTimeMsecs *int64          `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group was last modified.
	Name                 *string         `json:"name,omitempty" form:"name,omitempty"`                                 //Specifies the name of the group.
	Restricted           *bool           `json:"restricted,omitempty" form:"restricted,omitempty"`                     //Whether the group is a restricted group. Users belonging to a restricted
	Roles                *[]string       `json:"roles,omitempty" form:"roles,omitempty"`                               //Array of Roles.
	Sid                  *string         `json:"sid,omitempty" form:"sid,omitempty"`                                   //Specifies the unique Security ID (SID) of the group.
	SmbPrincipals        []*SmbPrincipal `json:"smbPrincipals,omitempty" form:"smbPrincipals,omitempty"`               //Specifies the SMB principals. Principals will be added to this group only
	TenantIds            *[]string       `json:"tenantIds,omitempty" form:"tenantIds,omitempty"`                       //Specifies the tenants to which the group belongs to. If not specified,
	Usernames            *[]string       `json:"usernames,omitempty" form:"usernames,omitempty"`                       //Specifies the username of users who are members of the group.
	Users                *[]string       `json:"users,omitempty" form:"users,omitempty"`                               //Specifies the SID of users who are members of the group.
}

/*
 * Structure for the custom type GroupDeleteParameters
 */
type GroupDeleteParameters struct {
	Domain *string   `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain associated with the groups to delete.
	Names  *[]string `json:"names,omitempty" form:"names,omitempty"`   //Array of Groups.
}

/*
 * Structure for the custom type GroupInfo
 */
type GroupInfo struct {
	Domain    *string `json:"domain,omitempty" form:"domain,omitempty"`       //Specifies domain name of the user.
	GroupName *string `json:"groupName,omitempty" form:"groupName,omitempty"` //Specifies group name of the group.
	Sid       *string `json:"sid,omitempty" form:"sid,omitempty"`             //Specifies unique Security ID (SID) of the user.
}

/*
 * Structure for the custom type GroupParameters
 */
type GroupParameters struct {
	Description   *string         `json:"description,omitempty" form:"description,omitempty"`     //Specifies a description of the group.
	Domain        *string         `json:"domain,omitempty" form:"domain,omitempty"`               //Specifies the domain of the group.
	Name          *string         `json:"name,omitempty" form:"name,omitempty"`                   //Specifies the name of the group.
	Restricted    *bool           `json:"restricted,omitempty" form:"restricted,omitempty"`       //Whether the group is a restricted group. Users belonging to a restricted
	Roles         *[]string       `json:"roles,omitempty" form:"roles,omitempty"`                 //Array of Roles.
	SmbPrincipals []*SmbPrincipal `json:"smbPrincipals,omitempty" form:"smbPrincipals,omitempty"` //Specifies the SMB principals. Principals will be added to this group only
	TenantIds     *[]string       `json:"tenantIds,omitempty" form:"tenantIds,omitempty"`         //Specifies the tenants to which the group belongs to. If not specified,
	Users         *[]string       `json:"users,omitempty" form:"users,omitempty"`                 //Specifies the SID of users who are members of the group.
}

/*
 * Structure for the custom type GuidPair
 */
type GuidPair struct {
	DestGuid   *string `json:"destGuid,omitempty" form:"destGuid,omitempty"`     //Specifies the guid of an AD object in the Production AD.
	SourceGuid *string `json:"sourceGuid,omitempty" form:"sourceGuid,omitempty"` //Specifies the guid of an AD object in the Snapshot AD.
}

/*
 * Structure for the custom type HardwareInfo
 */
type HardwareInfo struct {
	ChassisModel          *string `json:"chassisModel,omitempty" form:"chassisModel,omitempty"`                   //TODO: Write general description for this field
	ChassisSerial         *string `json:"chassisSerial,omitempty" form:"chassisSerial,omitempty"`                 //TODO: Write general description for this field
	ChassisType           *string `json:"chassisType,omitempty" form:"chassisType,omitempty"`                     //TODO: Write general description for this field
	CohesityChassisSerial *string `json:"cohesityChassisSerial,omitempty" form:"cohesityChassisSerial,omitempty"` //TODO: Write general description for this field
	CohesityNodeSerial    *string `json:"cohesityNodeSerial,omitempty" form:"cohesityNodeSerial,omitempty"`       //TODO: Write general description for this field
	HbaModel              *string `json:"hbaModel,omitempty" form:"hbaModel,omitempty"`                           //TODO: Write general description for this field
	IpmiLanChannel        *string `json:"ipmiLanChannel,omitempty" form:"ipmiLanChannel,omitempty"`               //TODO: Write general description for this field
	MaxSlots              *string `json:"maxSlots,omitempty" form:"maxSlots,omitempty"`                           //TODO: Write general description for this field
	NodeModel             *string `json:"nodeModel,omitempty" form:"nodeModel,omitempty"`                         //TODO: Write general description for this field
	NodeSerial            *string `json:"nodeSerial,omitempty" form:"nodeSerial,omitempty"`                       //TODO: Write general description for this field
	ProductModel          *string `json:"productModel,omitempty" form:"productModel,omitempty"`                   //TODO: Write general description for this field
	SlotNumber            *string `json:"slotNumber,omitempty" form:"slotNumber,omitempty"`                       //TODO: Write general description for this field
}

/*
 * Structure for the custom type HealthTile
 */
type HealthTile struct {
	CapacityBytes          *int64   `json:"capacityBytes,omitempty" form:"capacityBytes,omitempty"`                   //Raw Cluster Capacity in Bytes. This is not usable capacity and does not
	ClusterCloudUsageBytes *int64   `json:"clusterCloudUsageBytes,omitempty" form:"clusterCloudUsageBytes,omitempty"` //Usage in Bytes on the cloud.
	LastDayAlerts          []*Alert `json:"lastDayAlerts,omitempty" form:"lastDayAlerts,omitempty"`                   //Alerts in last 24 hours.
	LastDayNumCriticals    *int64   `json:"lastDayNumCriticals,omitempty" form:"lastDayNumCriticals,omitempty"`       //Number of Critical Alerts.
	LastDayNumWarnings     *int64   `json:"lastDayNumWarnings,omitempty" form:"lastDayNumWarnings,omitempty"`         //Number of Warning Alerts.
	NumNodes               *int64   `json:"numNodes,omitempty" form:"numNodes,omitempty"`                             //Number of nodes in the cluster.
	NumNodesWithIssues     *int64   `json:"numNodesWithIssues,omitempty" form:"numNodesWithIssues,omitempty"`         //Number of nodes in the cluster that are unhealthy.
	PercentFull            *float64 `json:"percentFull,omitempty" form:"percentFull,omitempty"`                       //Percent the cluster is full.
	RawUsedBytes           *int64   `json:"rawUsedBytes,omitempty" form:"rawUsedBytes,omitempty"`                     //Raw Bytes used in the cluster.
}

/*
 * Structure for the custom type HostEntry
 */
type HostEntry struct {
	DomainNames []string `json:"domainNames" form:"domainNames"` //Specifies the domain names of the host.
	Ip          string   `json:"ip" form:"ip"`                   //Specifies the IP address of the host.
}

/*
 * Structure for the custom type HostResult
 */
type HostResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies a message describing the status of the Hosts request.
}

/*
 * Structure for the custom type HyperFlexProtectionSource
 */
type HyperFlexProtectionSource struct {
	Name           *string                           `json:"name,omitempty" form:"name,omitempty"`                     //Specifies a unique name of the Protection Source
	ProductVersion *string                           `json:"productVersion,omitempty" form:"productVersion,omitempty"` //Specifies the product version of the protection source.
	Type           TypeHyperFlexProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                     //Specifies the type of managed Object in a HyperFlex protection source
	Uuid           *string                           `json:"uuid,omitempty" form:"uuid,omitempty"`                     //Specifies the uuid of the protection source.
}

/*
 * Structure for the custom type HypervBackupEnvParams
 */
type HypervBackupEnvParams struct {
	AllowCrashConsistentSnapshot *bool `json:"allowCrashConsistentSnapshot,omitempty" form:"allowCrashConsistentSnapshot,omitempty"` //Whether to fallback to take a crash-consistent snapshot incase taking
}

/*
 * Structure for the custom type HypervCloneParameters
 */
type HypervCloneParameters struct {
	DisableNetwork *bool   `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
	NetworkId      *int64  `json:"networkId,omitempty" form:"networkId,omitempty"`           //Specifies a network configuration to be attached to the cloned or
	PoweredOn      *bool   `json:"poweredOn,omitempty" form:"poweredOn,omitempty"`           //Specifies the power state of the cloned or recovered objects.
	Prefix         *string `json:"prefix,omitempty" form:"prefix,omitempty"`                 //Specifies a prefix to prepended to the source object name to derive a
	ResourceId     *int64  `json:"resourceId,omitempty" form:"resourceId,omitempty"`         //The resource (HyperV host) to which the restored VM will be attached.
	Suffix         *string `json:"suffix,omitempty" form:"suffix,omitempty"`                 //Specifies a suffix to appended to the original source object name
}

/*
 * Structure for the custom type HypervDatastore
 */
type HypervDatastore struct {
	Capacity    *int64                  `json:"capacity,omitempty" form:"capacity,omitempty"`       //Specifies the capacity of the datastore in bytes.
	FreeSpace   *int64                  `json:"freeSpace,omitempty" form:"freeSpace,omitempty"`     //Specifies the available space on the datastore in bytes.
	MountPoints *[]string               `json:"mountPoints,omitempty" form:"mountPoints,omitempty"` //Specifies the available mount points on the datastore.
	Type        TypeHypervDatastoreEnum `json:"type,omitempty" form:"type,omitempty"`               //Specifies the type of the datastore object like kFileShare or kVolume.
}

/*
 * Structure for the custom type HypervEnvJobParameters
 */
type HypervEnvJobParameters struct {
	FallbackToCrashConsistent *bool `json:"fallbackToCrashConsistent,omitempty" form:"fallbackToCrashConsistent,omitempty"` //If true, takes a crash-consistent snapshot when app-consistent snapshot
}

/*
 * Structure for the custom type HypervProtectionSource
 */
type HypervProtectionSource struct {
	Agents        []*AgentInformation                `json:"agents,omitempty" form:"agents,omitempty"`               //Array of Agents on the Physical Protection Source.
	BackupType    BackupTypeEnum                     `json:"backupType,omitempty" form:"backupType,omitempty"`       //Specifies the type of backup supported by the VM.
	ClusterName   *string                            `json:"clusterName,omitempty" form:"clusterName,omitempty"`     //Specifies the cluster name for 'kHostCluster' objects.
	DatastoreInfo *HypervDatastore                   `json:"datastoreInfo,omitempty" form:"datastoreInfo,omitempty"` //Specifies information about a Datastore Object in HyperV environment.
	Description   *string                            `json:"description,omitempty" form:"description,omitempty"`     //Specifies a description about the Protection Source.
	HostType      HostTypeHypervProtectionSourceEnum `json:"hostType,omitempty" form:"hostType,omitempty"`           //Specifies host OS type for 'kVirtualMachine' objects.
	HypervUuid    *string                            `json:"hypervUuid,omitempty" form:"hypervUuid,omitempty"`       //Specifies the UUID for 'kVirtualMachine' HyperV objects.
	Name          *string                            `json:"name,omitempty" form:"name,omitempty"`                   //Specifies the name of the HyperV Object.
	TagAttributes []*TagAttribute                    `json:"tagAttributes,omitempty" form:"tagAttributes,omitempty"` //Specifies the list of VM Tag attributes associated with this
	Type          TypeHypervProtectionSourceEnum     `json:"type,omitempty" form:"type,omitempty"`                   //Specifies the type of an HyperV Protection Source Object such as
	Uuid          *string                            `json:"uuid,omitempty" form:"uuid,omitempty"`                   //Specifies the UUID of the Object. This is unique within the HyperV
	VmInfo        *HypervVirtualMachine              `json:"vmInfo,omitempty" form:"vmInfo,omitempty"`               //Specifies information about a VirtualMachine Object in HyperV environment.
}

/*
 * Structure for the custom type HypervRestoreParameters
 */
type HypervRestoreParameters struct {
	DatastoreId    *int64  `json:"datastoreId,omitempty" form:"datastoreId,omitempty"`       //A datastore entity where the object's files should be restored to. This
	DisableNetwork *bool   `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
	NetworkId      *int64  `json:"networkId,omitempty" form:"networkId,omitempty"`           //Specifies a network configuration to be attached to the cloned or
	PoweredOn      *bool   `json:"poweredOn,omitempty" form:"poweredOn,omitempty"`           //Specifies the power state of the cloned or recovered objects.
	Prefix         *string `json:"prefix,omitempty" form:"prefix,omitempty"`                 //Specifies a prefix to prepended to the source object name to derive a
	ResourceId     *int64  `json:"resourceId,omitempty" form:"resourceId,omitempty"`         //The resource (HyperV host) to which the restored VM will be attached.
	Suffix         *string `json:"suffix,omitempty" form:"suffix,omitempty"`                 //Specifies a suffix to appended to the original source object name
}

/*
 * Structure for the custom type HypervVirtualMachine
 */
type HypervVirtualMachine struct {
	IsHighlyAvailable *bool              `json:"isHighlyAvailable,omitempty" form:"isHighlyAvailable,omitempty"` //Specifies whether the VM is Highly Availabile or not.
	Version           *string            `json:"version,omitempty" form:"version,omitempty"`                     //Specifies the version of the VM. For example, 8.0, 5.0 etc.
	VmBackupStatus    VmBackupStatusEnum `json:"vmBackupStatus,omitempty" form:"vmBackupStatus,omitempty"`       //Specifies the status of the VM for backup purpose.
	VmBackupType      VmBackupTypeEnum   `json:"vmBackupType,omitempty" form:"vmBackupType,omitempty"`           //Specifies the type of backup supported by the VM.
}

/*
 * Structure for the custom type IcapConnectionStatusResponse
 */
type IcapConnectionStatusResponse struct {
	FailedConnectionStatus    *[]string `json:"failedConnectionStatus,omitempty" form:"failedConnectionStatus,omitempty"`       //Specifies the failed connection status of Icap servers.
	SucceededConnectionStatus *[]string `json:"succeededConnectionStatus,omitempty" form:"succeededConnectionStatus,omitempty"` //Specifies the success connection status of Icap servers.
}

/*
 * Structure for the custom type IdMappingInfo
 */
type IdMappingInfo struct {
	FallbackUserIdMappingInfo *UserIdMapping `json:"fallbackUserIdMappingInfo,omitempty" form:"fallbackUserIdMappingInfo,omitempty"` //Specifies how the Unix and Windows users are mapped in an Active Directory.
	UnixRootSid               *string        `json:"unixRootSid,omitempty" form:"unixRootSid,omitempty"`                             //Specifies the SID of the Active Directory domain user to be mapped to
	UserIdMappingInfo         *UserIdMapping `json:"userIdMappingInfo,omitempty" form:"userIdMappingInfo,omitempty"`                 //Specifies how the Unix and Windows users are mapped in an Active Directory.
}

/*
 * Structure for the custom type IdpPrincipalsAddParameters
 */
type IdpPrincipalsAddParameters struct {
	Domain        *string                                   `json:"domain,omitempty" form:"domain,omitempty"`               //Specifies the name of the Idp where the
	ObjectClass   ObjectClassIdpPrincipalsAddParametersEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"`     //Specifies the type of the referenced Idp principal.
	PrincipalName *string                                   `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the Idp principal,
	Restricted    *bool                                     `json:"restricted,omitempty" form:"restricted,omitempty"`       //Whether the principal is a restricted principal. A restricted principal
	Roles         *[]string                                 `json:"roles,omitempty" form:"roles,omitempty"`                 //Array of Roles.
}

/*
 * Structure for the custom type IdpReachabilityTestResult
 */
type IdpReachabilityTestResult struct {
	Reachable *bool `json:"reachable,omitempty" form:"reachable,omitempty"` //Specifies the flag for Idp reachability.
}

/*
 * Structure for the custom type IdpServiceConfiguration
 */
type IdpServiceConfiguration struct {
	AllowLocalAuthentication *bool     `json:"allowLocalAuthentication,omitempty" form:"allowLocalAuthentication,omitempty"` //Specifies whether to allow local authentication. When IdP is configured,
	Certificate              *string   `json:"certificate,omitempty" form:"certificate,omitempty"`                           //Specifies the certificate generated for the app by the IdP service when
	CertificateFilename      *string   `json:"certificateFilename,omitempty" form:"certificateFilename,omitempty"`           //Specifies the filename used to upload the certificate.
	Domain                   *string   `json:"domain,omitempty" form:"domain,omitempty"`                                     //Specifies a unique name for this IdP configuration.
	Enable                   *bool     `json:"enable,omitempty" form:"enable,omitempty"`                                     //Specifies a flag to enable or disable this IdP service. When it is set
	Id                       *int64    `json:"id,omitempty" form:"id,omitempty"`                                             //Specifies the Id assigned by the Cluster for the IdP service.
	IssuerId                 *string   `json:"issuerId,omitempty" form:"issuerId,omitempty"`                                 //Specifies the IdP provided Issuer ID for the app.
	Name                     *string   `json:"name,omitempty" form:"name,omitempty"`                                         //Specifies the name of the vendor providing IdP service.
	Roles                    *[]string `json:"roles,omitempty" form:"roles,omitempty"`                                       //Specifies a list roles assigned to an IdP user if samlAttributeName is
	SamlAttributeName        *string   `json:"samlAttributeName,omitempty" form:"samlAttributeName,omitempty"`               //Specifies the SAML attribute name that contains a comma separated list
	SignRequest              *bool     `json:"signRequest,omitempty" form:"signRequest,omitempty"`                           //Specifies whether to sign the SAML request or not. When it is set
	SsoUrl                   *string   `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"`                                     //Specifies the SSO URL of the IdP service for the customer. This is the
	TenantId                 *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                 //Specifies the Tenant Id if the IdP is configured for a Tenant. If this is
}

/*
 * Structure for the custom type IdpUserInfo
 */
type IdpUserInfo struct {
	IdpId    *int64  `json:"idpId,omitempty" form:"idpId,omitempty"`       //Specifies the unique Id assigned by the Cluster for the IdP.
	IssuerId *string `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the unique identifier assigned by the vendor for this Cluster.
	UserId   *string `json:"userId,omitempty" form:"userId,omitempty"`     //Specifies the unique identifier assigned by the vendor for the user.
	Vendor   *string `json:"vendor,omitempty" form:"vendor,omitempty"`     //Specifies the vendor providing the IdP service.
}

/*
 * Structure for the custom type IndexAndSnapshots
 */
type IndexAndSnapshots struct {
	ArchiveTaskUid         *UniversalId `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"`                 //Specifies a unique id of the Archive task that originally archived the
	EndTimeUsecs           *int64       `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                     //Specifies the end time as a Unix epoch Timestamp (in microseconds).
	RemoteProtectionJobUid *UniversalId `json:"remoteProtectionJobUid,omitempty" form:"remoteProtectionJobUid,omitempty"` //Specifies a unique id assigned to the original Protection Job
	StartTimeUsecs         *int64       `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                 //Specifies the start time as a Unix epoch Timestamp (in microseconds).
	ViewBoxId              *int64       `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                           //Specifies the id of the local Storage Domain (View Box) where the index
}

/*
 * Structure for the custom type IndexingPolicy
 */
type IndexingPolicy struct {
	AllowPrefixes   *[]string `json:"allowPrefixes,omitempty" form:"allowPrefixes,omitempty"`     //Array of Indexed Directories.
	DenyPrefixes    *[]string `json:"denyPrefixes,omitempty" form:"denyPrefixes,omitempty"`       //Array of Excluded Directories.
	DisableIndexing *bool     `json:"disableIndexing,omitempty" form:"disableIndexing,omitempty"` //Specifies if the files found in an Object (such as a VM) should be
}

/*
 * Structure for the custom type IndexingPolicyProto
 */
type IndexingPolicyProto struct {
	AllowPrefixes   *[]string `json:"allowPrefixes,omitempty" form:"allowPrefixes,omitempty"`     //List of directory prefixes to allow for indexing.
	DenyPrefixes    *[]string `json:"denyPrefixes,omitempty" form:"denyPrefixes,omitempty"`       //List of directory prefixes to filter out.
	DisableIndexing *bool     `json:"disableIndexing,omitempty" form:"disableIndexing,omitempty"` //If this field is set all the files in the VM will be filtered.
}

/*
 * Structure for the custom type InfectedFile
 */
type InfectedFile struct {
	EntityId                    *int64               `json:"entityId,omitempty" form:"entityId,omitempty"`                                       //Specifies the entity id of the infected file.
	FilePath                    *string              `json:"filePath,omitempty" form:"filePath,omitempty"`                                       //Specifies file path of the infected file.
	InfectionDetectionTimestamp *int64               `json:"infectionDetectionTimestamp,omitempty" form:"infectionDetectionTimestamp,omitempty"` //Specifies unix epoch timestamp (in microseconds) at which these threats
	ModifiedTimestampUsecs      *int64               `json:"modifiedTimestampUsecs,omitempty" form:"modifiedTimestampUsecs,omitempty"`           //Specifies unix epoch timestamp (in microseconds) at which this file is
	RemediationState            RemediationStateEnum `json:"remediationState,omitempty" form:"remediationState,omitempty"`                       //Specifies the remediation state of the file.
	RootInodeId                 *int64               `json:"rootInodeId,omitempty" form:"rootInodeId,omitempty"`                                 //Specifies the root inode id of the file system that infected file belongs
	ScanTimestampUsecs          *int64               `json:"scanTimestampUsecs,omitempty" form:"scanTimestampUsecs,omitempty"`                   //Specifies unix epoch timestamp (in microseconds) at which inode was
	ServiceIcapUri              *string              `json:"serviceIcapUri,omitempty" form:"serviceIcapUri,omitempty"`                           //Specifies the instance of an antivirus ICAP server in the cluster config
	ThreatDescriptions          *[]string            `json:"threatDescriptions,omitempty" form:"threatDescriptions,omitempty"`                   //Specifies the list of virus threat descriptions found in the file.
	ViewId                      *int64               `json:"viewId,omitempty" form:"viewId,omitempty"`                                           //Specifies the id of the View the infected file belongs to.
	ViewName                    *string              `json:"viewName,omitempty" form:"viewName,omitempty"`                                       //Specifies the View name corresponding to above view id.
}

/*
 * Structure for the custom type InfectedFileId
 */
type InfectedFileId struct {
	EntityId    *int64 `json:"entityId,omitempty" form:"entityId,omitempty"`       //Specifies the entity id of the infected file.
	RootInodeId *int64 `json:"rootInodeId,omitempty" form:"rootInodeId,omitempty"` //Specifies the root inode id of the file system that infected file belongs
	ViewId      *int64 `json:"viewId,omitempty" form:"viewId,omitempty"`           //Specifies the id of the View the infected file belongs to.
}

/*
 * Structure for the custom type InfectedFileParam
 */
type InfectedFileParam struct {
	EntityId         *int64               `json:"entityId,omitempty" form:"entityId,omitempty"`                 //Specifies the entity id of the infected file.
	RemediationState RemediationStateEnum `json:"remediationState,omitempty" form:"remediationState,omitempty"` //Specifies the remediation state of the file.
	RootInodeId      *int64               `json:"rootInodeId,omitempty" form:"rootInodeId,omitempty"`           //Specifies the root inode id of the file system that infected file belongs
	ViewId           *int64               `json:"viewId,omitempty" form:"viewId,omitempty"`                     //Specifies the id of the View the infected file belongs to.
}

/*
 * Structure for the custom type InfectedFiles
 */
type InfectedFiles struct {
	InfectedFiles    []*InfectedFile `json:"infectedFiles,omitempty" form:"infectedFiles,omitempty"`       //Specifies the infected files.
	PaginationCookie *string         `json:"paginationCookie,omitempty" form:"paginationCookie,omitempty"` //This cookie can be used in the succeeding call to list infected files to
}

/*
 * Structure for the custom type InterfaceGroup
 */
type InterfaceGroup struct {
	Id                  *int64                        `json:"id,omitempty" form:"id,omitempty"`                                   //Interface group Id.
	ModelInterfaceLists []*ProductModelInterfaceTuple `json:"modelInterfaceLists,omitempty" form:"modelInterfaceLists,omitempty"` //Specifies the product model and interface lists.
	Name                *string                       `json:"name,omitempty" form:"name,omitempty"`                               //Specifies the name of the interface group.
}

/*
 * Structure for the custom type IoPreferentialTier
 */
type IoPreferentialTier struct {
	ApolloIOPreferentialTier        *[]ApolloIOPreferentialTierEnum        `json:"apolloIOPreferentialTier,omitempty" form:"apolloIOPreferentialTier,omitempty"`               //Specifies the perferred storage tier used by Apollo as its working directory.
	ApolloWalIOPreferentialTier     *[]ApolloWalIOPreferentialTierEnum     `json:"apolloWalIOPreferentialTier,omitempty" form:"apolloWalIOPreferentialTier,omitempty"`         //Specifies the perferred storage tier used by Apollo as its actions WAL.
	AthenaIOPreferentialTier        *[]AthenaIOPreferentialTierEnum        `json:"athenaIOPreferentialTier,omitempty" form:"athenaIOPreferentialTier,omitempty"`               //Specifies the list of perferred storage tiers used by Athena.
	AthenaSlowerIOPreferentialTier  *[]AthenaSlowerIOPreferentialTierEnum  `json:"athenaSlowerIOPreferentialTier,omitempty" form:"athenaSlowerIOPreferentialTier,omitempty"`   //Specifies the list of perferred storage tiers used by Athena for slower storage.
	DownTierUsagePercentThresholds  *[]int64                               `json:"downTierUsagePercentThresholds,omitempty" form:"downTierUsagePercentThresholds,omitempty"`   //Specifies the usage percentage thresholds for the correponding storage tier.
	GrootIOPreferentialTier         *[]GrootIOPreferentialTierEnum         `json:"grootIOPreferentialTier,omitempty" form:"grootIOPreferentialTier,omitempty"`                 //Specifies the perferred storage tier used by Groot as its working directory.
	HydraDowntierIOPreferentialTier *[]HydraDowntierIOPreferentialTierEnum `json:"hydraDowntierIOPreferentialTier,omitempty" form:"hydraDowntierIOPreferentialTier,omitempty"` //Specifies the list of perferred storage tiers used by Hydra for offloading.
	HydraIOPreferentialTier         *[]HydraIOPreferentialTierEnum         `json:"hydraIOPreferentialTier,omitempty" form:"hydraIOPreferentialTier,omitempty"`                 //Specifies the list of perferred storage tiers used by Hydra.
	LibrarianIOPreferentialTier     *[]LibrarianIOPreferentialTierEnum     `json:"librarianIOPreferentialTier,omitempty" form:"librarianIOPreferentialTier,omitempty"`         //Specifies the list of perferred storage tiers used by librarian.
	RandomIOPreferentialTier        *[]RandomIOPreferentialTierEnum        `json:"randomIOPreferentialTier,omitempty" form:"randomIOPreferentialTier,omitempty"`               //Specifies the order of perferred storage tiers for random IO operations.
	ScribeIOPreferentialTier        *[]ScribeIOPreferentialTierEnum        `json:"scribeIOPreferentialTier,omitempty" form:"scribeIOPreferentialTier,omitempty"`               //Specifies the list of perferred storage tiers used by Scribe.
	SequentialIOPreferentialTier    *[]SequentialIOPreferentialTierEnum    `json:"sequentialIOPreferentialTier,omitempty" form:"sequentialIOPreferentialTier,omitempty"`       //Specifies the perferred storage tier for sequential IO operations.
	YodaIOPreferentialTier          *[]YodaIOPreferentialTierEnum          `json:"yodaIOPreferentialTier,omitempty" form:"yodaIOPreferentialTier,omitempty"`                   //Specifies the list of perferred storage tiers used by Yoda.
}

/*
 * Structure for the custom type IopsTile
 */
type IopsTile struct {
	MaxReadIops      *int64    `json:"maxReadIops,omitempty" form:"maxReadIops,omitempty"`           //Maximum Read IOs per second in last 24 hours.
	MaxWriteIops     *int64    `json:"maxWriteIops,omitempty" form:"maxWriteIops,omitempty"`         //Maximum number of Write IOs per second in last 24 hours.
	ReadIopsSamples  []*Sample `json:"readIopsSamples,omitempty" form:"readIopsSamples,omitempty"`   //Read IOs per second samples taken for the past 24 hours at 10 minutes
	WriteIopsSamples []*Sample `json:"writeIopsSamples,omitempty" form:"writeIopsSamples,omitempty"` //Write IOs per second samples taken for the past 24 hours at 10 minutes
}

/*
 * Structure for the custom type IpmiConfiguration
 */
type IpmiConfiguration struct {
	IpmiGateway    *string `json:"ipmiGateway,omitempty" form:"ipmiGateway,omitempty"`       //Specifies the default Gateway IP Address for the IPMI network.
	IpmiPassword   *string `json:"ipmiPassword,omitempty" form:"ipmiPassword,omitempty"`     //Specifies the IPMI Password.
	IpmiSubnetMask *string `json:"ipmiSubnetMask,omitempty" form:"ipmiSubnetMask,omitempty"` //Specifies the subnet mask for the IPMI network.
	IpmiUsername   *string `json:"ipmiUsername,omitempty" form:"ipmiUsername,omitempty"`     //Specifies the IPMI Username.
}

/*
 * Structure for the custom type IscsiSanPort
 */
type IscsiSanPort struct {
	IpAddress *string `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the IP address of the SAN port.
	Iqn       *string `json:"iqn,omitempty" form:"iqn,omitempty"`             //Specifies the qualified name of the iSCSI port (IQN).
	TcpPort   *int64  `json:"tcpPort,omitempty" form:"tcpPort,omitempty"`     //Specifies the listening port(tcp) of the SAN port.
}

/*
 * Structure for the custom type IsilonAccessZone
 */
type IsilonAccessZone struct {
	Id   *int64  `json:"id,omitempty" form:"id,omitempty"`     //Specifies the id of the access zone.
	Name *string `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the access zone.
	Path *string `json:"path,omitempty" form:"path,omitempty"` //Specifies the path of the access zone in ifs. This should include the
}

/*
 * Structure for the custom type IsilonCluster
 */
type IsilonCluster struct {
	Description *string `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of an Isilon Cluster.
	Guid        *string `json:"guid,omitempty" form:"guid,omitempty"`               //Specifies a globally unique id of an Isilon Cluster.
}

/*
 * Structure for the custom type IsilonMountPoint
 */
type IsilonMountPoint struct {
	AccessZoneName *string                         `json:"accessZoneName,omitempty" form:"accessZoneName,omitempty"` //Specifies the name of access zone.
	NfsMountPoint  *IsilonNfsMountPoint            `json:"nfsMountPoint,omitempty" form:"nfsMountPoint,omitempty"`   //Specifies NFS Mount Point exposed by Isilon Protection Source.
	Path           *string                         `json:"path,omitempty" form:"path,omitempty"`                     //Specifies the path of the access zone in ifs. This should include the
	Protocols      *[]ProtocolIsilonMountPointEnum `json:"protocols,omitempty" form:"protocols,omitempty"`           //List of Protocols on Isilon.
	SmbMountPoints []*IsilonSmbMountPoint          `json:"smbMountPoints,omitempty" form:"smbMountPoints,omitempty"` //Specifies information about an SMB share. This field is set if the
}

/*
 * Structure for custom type LicenceClusterParameters
 */

type LicenceClusterParameters struct {
	SignedVersion *int64
	SignedByUser  string
	SignedTime    *int64
	LicenseKey    string
}

/*
 * Structure for the custom type IsilonNfsMountPoint
 */
type IsilonNfsMountPoint struct {
	AccessZoneName *string `json:"accessZoneName,omitempty" form:"accessZoneName,omitempty"` //Specifies the Access Zone name.
	Description    *string `json:"description,omitempty" form:"description,omitempty"`       //Specifies the description of the NFS mount point.
	Id             *int64  `json:"id,omitempty" form:"id,omitempty"`                         //Specifies the Id of the NFS export.
}

/*
 * Structure for the custom type IsilonProtectionSource
 */
type IsilonProtectionSource struct {
	AccessZone *IsilonAccessZone              `json:"accessZone,omitempty" form:"accessZone,omitempty"` //Specifies information about access zone in an Isilon Cluster.
	Cluster    *IsilonCluster                 `json:"cluster,omitempty" form:"cluster,omitempty"`       //Specifies information about an Isilon Cluster.
	MountPoint *IsilonMountPoint              `json:"mountPoint,omitempty" form:"mountPoint,omitempty"` //Specifies information about a mount point in an Isilon OneFs Cluster.
	Name       *string                        `json:"name,omitempty" form:"name,omitempty"`             //Specifies a unique name of the Protection Source.
	Type       TypeIsilonProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`             //Specifies the type of the entity in an Isilon OneFs file system
}

/*
 * Structure for the custom type IsilonSmbMountPoint
 */
type IsilonSmbMountPoint struct {
	AccessZoneId *int64  `json:"accessZoneId,omitempty" form:"accessZoneId,omitempty"` //Specifies the Access Zone Id.
	Description  *string `json:"description,omitempty" form:"description,omitempty"`   //Specifies the description of the NFS mount point.
	ShareName    *string `json:"shareName,omitempty" form:"shareName,omitempty"`       //Specifies the name of the SMB/CIFS share.
}

/*
 * Structure for the custom type JobPolicyProto
 */
type JobPolicyProto struct {
	BackupPolicy            *BackupPolicyProto           `json:"backupPolicy,omitempty" form:"backupPolicy,omitempty"`                       //If a backup does not get a chance to when it's due (either due to the system
	SnapshotTargetPolicyVec []*SnapshotTargetPolicyProto `json:"snapshotTargetPolicyVec,omitempty" form:"snapshotTargetPolicyVec,omitempty"` //Specifies additional policies that can be used to copy snapshots created
}

/*
 * Structure for the custom type JobRunsTile
 */
type JobRunsTile struct {
	LastDayNumJobErrors        *int64                      `json:"lastDayNumJobErrors,omitempty" form:"lastDayNumJobErrors,omitempty"`               //Number of Error runs in the last 24 hours.
	LastDayNumJobRuns          *int64                      `json:"lastDayNumJobRuns,omitempty" form:"lastDayNumJobRuns,omitempty"`                   //Number of Job Runs in the last 24 hours.
	LastDayNumJobSlaViolations *int64                      `json:"lastDayNumJobSlaViolations,omitempty" form:"lastDayNumJobSlaViolations,omitempty"` //Number of SLA Violations in the last 24 hours.
	NumJobRunning              *int64                      `json:"numJobRunning,omitempty" form:"numJobRunning,omitempty"`                           //Number of Jobs currently running.
	ObjectsProtectedByPolicy   []*ObjectsProtectedByPolicy `json:"objectsProtectedByPolicy,omitempty" form:"objectsProtectedByPolicy,omitempty"`     //Objects Protected By Policy.
}

/*
 * Structure for the custom type KeyValuePair
 */
type KeyValuePair struct {
	Key   *string `json:"key,omitempty" form:"key,omitempty"`     //Specifies the name of the key.
	Value *Value  `json:"value,omitempty" form:"value,omitempty"` //Specifies a data type and data field used to store data.
}

/*
 * Structure for the custom type KmsConfiguration
 */
type KmsConfiguration struct {
	CaCertificate       *string        `json:"caCertificate,omitempty" form:"caCertificate,omitempty"`             //Specifies the CA certificate in PEM format.
	ClientCertificate   *string        `json:"clientCertificate,omitempty" form:"clientCertificate,omitempty"`     //Specifies the client certificate.
	ClientKey           *string        `json:"clientKey,omitempty" form:"clientKey,omitempty"`                     //Specifies the client private key.
	KmipProtocolVersion *string        `json:"kmipProtocolVersion,omitempty" form:"kmipProtocolVersion,omitempty"` //Specifies protocol version used to communicate with the KMS.
	ServerIp            *string        `json:"serverIp,omitempty" form:"serverIp,omitempty"`                       //Specifies the KMS IP address.
	ServerName          *string        `json:"serverName,omitempty" form:"serverName,omitempty"`                   //Specifies the name given to the KMS Server.
	ServerPort          *int64         `json:"serverPort,omitempty" form:"serverPort,omitempty"`                   //Specifies port on which the server is listening.
	ServerType          ServerTypeEnum `json:"serverType,omitempty" form:"serverType,omitempty"`                   //Specifies the type of key mangement system.
}

/*
 * Structure for the custom type KmsConfigurationResponse
 */
type KmsConfigurationResponse struct {
	ClientCertificateExpiryDate *int64  `json:"clientCertificateExpiryDate,omitempty" form:"clientCertificateExpiryDate,omitempty"` //Specifies expiry date of client certificate.
	ConnectionStatus            *bool   `json:"connectionStatus,omitempty" form:"connectionStatus,omitempty"`                       //Specifies if connection to this KMS exists.
	KmipProtocolVersion         *string `json:"kmipProtocolVersion,omitempty" form:"kmipProtocolVersion,omitempty"`                 //Specifies protocol version used to communicate with the KMS.
	ServerIp                    *string `json:"serverIp,omitempty" form:"serverIp,omitempty"`                                       //Specifies the KMS IP address.
	ServerName                  *string `json:"serverName,omitempty" form:"serverName,omitempty"`                                   //Specifies the name given to the KMS Server.
	ServerPort                  *int64  `json:"serverPort,omitempty" form:"serverPort,omitempty"`                                   //Specifies port on which the server is listening.
}

/*
 * Structure for the custom type KubernetesCredentials
 */
type KubernetesCredentials struct {
	ClientPrivateKey *string `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"` //Specifies Client private associated with the service account.
}

/*
 * Structure for the custom type KubernetesProtectionSource
 */
type KubernetesProtectionSource struct {
	Description *string                            `json:"description,omitempty" form:"description,omitempty"` //Specifies an optional description of the object.
	Name        *string                            `json:"name,omitempty" form:"name,omitempty"`               //Specifies a unique name of the Protection Source.
	Type        TypeKubernetesProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`               //Specifies the type of the entity in a Kubernetes environment.
	Uuid        *string                            `json:"uuid,omitempty" form:"uuid,omitempty"`               //Specifies the UUID of the object.
}

/*
 * Structure for the custom type KvmProtectionSource
 */
type KvmProtectionSource struct {
	AgentError   *string                     `json:"agentError,omitempty" form:"agentError,omitempty"`     //Specifies a message when the agent cannot be reached.
	AgentId      *int64                      `json:"agentId,omitempty" form:"agentId,omitempty"`           //Specifies the ID of the Agent with which this KVM entity is associated
	ClusterId    *string                     `json:"clusterId,omitempty" form:"clusterId,omitempty"`       //Specifies the cluster ID for 'kCluster' objects.
	DatacenterId *string                     `json:"datacenterId,omitempty" form:"datacenterId,omitempty"` //Specifies the ID of the 'kDatacenter' objects.
	Description  *string                     `json:"description,omitempty" form:"description,omitempty"`   //Specifies a description about the Protection Source.
	Name         *string                     `json:"name,omitempty" form:"name,omitempty"`                 //Specifies the name of the KVM entity.
	NetworkId    *string                     `json:"networkId,omitempty" form:"networkId,omitempty"`       //Specifies the network ID to which Vnic is attached.
	Type         TypeKvmProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                 //Specifies the type of KVM Protection Source entities such as
	Uuid         *string                     `json:"uuid,omitempty" form:"uuid,omitempty"`                 //Specifies the UUID of the Object. This is unique within the KVM
}

/*
 * Structure for the custom type LastProtectionRunStats
 */
type LastProtectionRunStats struct {
	NumRunsFailed    *int64                         `json:"numRunsFailed,omitempty" form:"numRunsFailed,omitempty"`       //Specifies the number of Protection Jobs for which specified Protection Run failed.
	NumRunsFailedSla *int64                         `json:"numRunsFailedSla,omitempty" form:"numRunsFailedSla,omitempty"` //Specifies the number of Protection Jobs for which specified Protection Run failed SLA.
	NumRunsMetSla    *int64                         `json:"numRunsMetSla,omitempty" form:"numRunsMetSla,omitempty"`       //Specifies the number of Protection Jobs for which specified Protection Run met SLA.
	StatsByEnv       []*LastProtectionRunStatsByEnv `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"`             //Specifies the last Protection Run stats by environment.
}

/*
 * Structure for the custom type LastProtectionRunStatsByEnv
 */
type LastProtectionRunStatsByEnv struct {
	Environment         EnvironmentLastProtectionRunStatsByEnvEnum `json:"environment,omitempty" form:"environment,omitempty"`                 //Specifies the environment.
	NumObjectsFailed    *int64                                     `json:"numObjectsFailed,omitempty" form:"numObjectsFailed,omitempty"`       //Specifies the count of objects that failed last Protection Run.
	NumObjectsFailedSla *int64                                     `json:"numObjectsFailedSla,omitempty" form:"numObjectsFailedSla,omitempty"` //Specifies the count of objects that failed sla in the last Run.
	NumObjectsMetSla    *int64                                     `json:"numObjectsMetSla,omitempty" form:"numObjectsMetSla,omitempty"`       //Specifies the count of objects that met sla in the last Run.
}

/*
 * Structure for the custom type LastProtectionRunSummary
 */
type LastProtectionRunSummary struct {
	NumberOfCancelledProtectionRuns  *int64 `json:"numberOfCancelledProtectionRuns,omitempty" form:"numberOfCancelledProtectionRuns,omitempty"`   //Specifies the number of cancelled Protection Runs the given
	NumberOfFailedProtectionRuns     *int64 `json:"numberOfFailedProtectionRuns,omitempty" form:"numberOfFailedProtectionRuns,omitempty"`         //Specifies the number of failed Protection Runs the given
	NumberOfProtectedSources         *int64 `json:"numberOfProtectedSources,omitempty" form:"numberOfProtectedSources,omitempty"`                 //Specifies the number of Protection Sources protected by the given
	NumberOfRunningProtectionRuns    *int64 `json:"numberOfRunningProtectionRuns,omitempty" form:"numberOfRunningProtectionRuns,omitempty"`       //Specifies the number of running Protection Runs using the current
	NumberOfSlaViolations            *int64 `json:"numberOfSlaViolations,omitempty" form:"numberOfSlaViolations,omitempty"`                       //Specifies the number of SLA violations the given
	NumberOfSuccessfulProtectionRuns *int64 `json:"numberOfSuccessfulProtectionRuns,omitempty" form:"numberOfSuccessfulProtectionRuns,omitempty"` //Specifies the number of successful Protection Runs the given
	TotalLogicalBackupSizeInBytes    *int64 `json:"totalLogicalBackupSizeInBytes,omitempty" form:"totalLogicalBackupSizeInBytes,omitempty"`       //Specifies the aggregated total logical backup performed in all the
}

/*
 * Structure for the custom type LatencyThresholds
 */
type LatencyThresholds struct {
	ActiveTaskMsecs *int64 `json:"activeTaskMsecs,omitempty" form:"activeTaskMsecs,omitempty"` //If the latency of a datastore is above this value, existing backup tasks
	NewTaskMsecs    *int64 `json:"newTaskMsecs,omitempty" form:"newTaskMsecs,omitempty"`       //If the latency of a datastore is above this value, then new backup tasks
}

/*
 * Structure for the custom type LatestProtectionJobRunInfo
 */
type LatestProtectionJobRunInfo struct {
	LatestSnapshotInfo *LatestProtectionRun `json:"latestSnapshotInfo,omitempty" form:"latestSnapshotInfo,omitempty"` //Specifies the information about the latest Protection Run.
	LocationName       *string              `json:"locationName,omitempty" form:"locationName,omitempty"`             //Specifies the name of location that the object is backedup to.
	NumSnapshots       *int64               `json:"numSnapshots,omitempty" form:"numSnapshots,omitempty"`             //Specifies of number of successful snapshots.
}

/*
 * Structure for the custom type LatestProtectionRun
 */
type LatestProtectionRun struct {
	BackupRun           *SourceBackupStatus `json:"backupRun,omitempty" form:"backupRun,omitempty"`                     //Specifies the source object to protect and the current backup status.
	ChangeEventId       *int64              `json:"changeEventId,omitempty" form:"changeEventId,omitempty"`             //Specifies the event id which caused last update on this object.
	CopyRun             *CopyRun            `json:"copyRun,omitempty" form:"copyRun,omitempty"`                         //Specifies details about the Copy Run for a backup run of a Job Run.
	JobRunId            *int64              `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                       //Specifies job run id of the latest successful Protection Job Run.
	ProtectionJobRunUid *RunUid             `json:"protectionJobRunUid,omitempty" form:"protectionJobRunUid,omitempty"` //Specifies the universal id of the latest successful Protection Job Run.
	SnapshotTarget      *string             `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"`           //Specifies the cluster id in case of local or replication snapshots and
	SnapshotTargetType  *int64              `json:"snapshotTargetType,omitempty" form:"snapshotTargetType,omitempty"`   //Specifies the snapshot target type of the latest snapshot.
	TaskStatus          *int64              `json:"taskStatus,omitempty" form:"taskStatus,omitempty"`                   //Specifies the task status of the Protecion Job Run in the final attempt.
	Uuid                *string             `json:"uuid,omitempty" form:"uuid,omitempty"`                               //Specifies the unique id of the Protection Source for which a snapshot is
}

/*
 * Structure for the custom type LdapProvider
 */
type LdapProvider struct {
	AdDomainName            *string      `json:"adDomainName,omitempty" form:"adDomainName,omitempty"`                       //Specifies the domain name of an Active Directory which is mapped to this
	AuthType                AuthTypeEnum `json:"authType,omitempty" form:"authType,omitempty"`                               //Specifies the authentication type used while connecting to LDAP servers.
	BaseDistinguishedName   *string      `json:"baseDistinguishedName,omitempty" form:"baseDistinguishedName,omitempty"`     //Specifies the base distinguished name used as the base for LDAP
	DomainName              *string      `json:"domainName,omitempty" form:"domainName,omitempty"`                           //Specifies the name of the domain name to be used for querying LDAP servers
	Name                    *string      `json:"name,omitempty" form:"name,omitempty"`                                       //Specifies the name of the LDAP provider.
	Port                    *int64       `json:"port,omitempty" form:"port,omitempty"`                                       //Specifies LDAP server port.
	PreferredLdapServerList *[]string    `json:"preferredLdapServerList,omitempty" form:"preferredLdapServerList,omitempty"` //Specifies the preferred LDAP servers. Server names should be either in
	TenantId                *string      `json:"tenantId,omitempty" form:"tenantId,omitempty"`                               //Specifies the unique id of the tenant.
	UseSsl                  *bool        `json:"useSsl,omitempty" form:"useSsl,omitempty"`                                   //Specifies whether to use SSL for LDAP connections.
	UserDistinguishedName   *string      `json:"userDistinguishedName,omitempty" form:"userDistinguishedName,omitempty"`     //Specifies the user distinguished name that is used for LDAP authentication.
	UserPassword            *string      `json:"userPassword,omitempty" form:"userPassword,omitempty"`                       //Specifies the user password that is used for LDAP authentication.
}

/*
 * Structure for the custom type LdapProviderResponse
 */
type LdapProviderResponse struct {
	AdDomainName            *string      `json:"adDomainName,omitempty" form:"adDomainName,omitempty"`                       //Specifies the domain name of an Active Directory which is mapped to this
	AuthType                AuthTypeEnum `json:"authType,omitempty" form:"authType,omitempty"`                               //Specifies the authentication type used while connecting to LDAP servers.
	BaseDistinguishedName   *string      `json:"baseDistinguishedName,omitempty" form:"baseDistinguishedName,omitempty"`     //Specifies the base distinguished name used as the base for LDAP
	DomainName              *string      `json:"domainName,omitempty" form:"domainName,omitempty"`                           //Specifies the name of the domain name to be used for querying LDAP servers
	Id                      *int64       `json:"id,omitempty" form:"id,omitempty"`                                           //Specifies the ID of the LDAP provider.
	Name                    *string      `json:"name,omitempty" form:"name,omitempty"`                                       //Specifies the name of the LDAP provider.
	Port                    *int64       `json:"port,omitempty" form:"port,omitempty"`                                       //Specifies LDAP server port.
	PreferredLdapServerList *[]string    `json:"preferredLdapServerList,omitempty" form:"preferredLdapServerList,omitempty"` //Specifies the preferred LDAP servers. Server names should be either in
	TenantId                *string      `json:"tenantId,omitempty" form:"tenantId,omitempty"`                               //Specifies the unique id of the tenant.
	UseSsl                  *bool        `json:"useSsl,omitempty" form:"useSsl,omitempty"`                                   //Specifies whether to use SSL for LDAP connections.
	UserDistinguishedName   *string      `json:"userDistinguishedName,omitempty" form:"userDistinguishedName,omitempty"`     //Specifies the user distinguished name that is used for LDAP authentication.
	UserPassword            *string      `json:"userPassword,omitempty" form:"userPassword,omitempty"`                       //Specifies the user password that is used for LDAP authentication.
}

/*
 * Structure for the custom type LdapProviderStatus
 */
type LdapProviderStatus struct {
	StatusMessage *string `json:"statusMessage,omitempty" form:"statusMessage,omitempty"` //Specifies the connection status message of an LDAP provider.
}

/*
 * Structure for the custom type LegalHoldings
 */
type LegalHoldings struct {
	HoldForLegalPurpose *bool  `json:"holdForLegalPurpose,omitempty" form:"holdForLegalPurpose,omitempty"` //Specifies whether the source is put on legal hold or not.
	ProtectionSourceId  *int64 `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"`   //Specifies an Protection Source Id in the snapshot.
}

/*
 * Structure for the custom type LicenseState
 */
type LicenseState struct {
	FailedAttempts *int64                `json:"failedAttempts,omitempty" form:"failedAttempts,omitempty"` //Specifies no of failed attempts at claiming the license server
	State          StateLicenseStateEnum `json:"state,omitempty" form:"state,omitempty"`                   //Specifies the current state of licensing workflow.
}

/*
 * Structure for the custom type ListCentrifyZone
 */
type ListCentrifyZone struct {
	CentrifySchema    CentrifySchemaEnum `json:"centrifySchema,omitempty" form:"centrifySchema,omitempty"`       //Specifies the schema of this Centrify zone.
	Description       *string            `json:"description,omitempty" form:"description,omitempty"`             //Specifies a description of the Centrify zone.
	DistinguishedName *string            `json:"distinguishedName,omitempty" form:"distinguishedName,omitempty"` //Specifies the distinguished name of the Centrify zone.
	ZoneName          *string            `json:"zoneName,omitempty" form:"zoneName,omitempty"`                   //Specifies the zone name.
}

/*
 * Structure for the custom type ListNlmLocksResponse
 */
type ListNlmLocksResponse struct {
	Cookie        *string         `json:"cookie,omitempty" form:"cookie,omitempty"`               //Specifies an opaque string to pass to get the next set of NLM locks.
	FilesNlmLocks []*FileNlmLocks `json:"filesNlmLocks,omitempty" form:"filesNlmLocks,omitempty"` //Specifies the list of NLM locks.
}

/*
 * Structure for the custom type LockFileParams
 */
type LockFileParams struct {
	ExpiryTimestampMsecs *int64  `json:"expiryTimestampMsecs,omitempty" form:"expiryTimestampMsecs,omitempty"` //Specifies a definite timestamp in milliseconds for retaining the file, or
	Path                 *string `json:"path,omitempty" form:"path,omitempty"`                                 //Specifies the file path that needs to be locked.
}

/*
 * Structure for the custom type LockRange
 */
type LockRange struct {
	IsExclusive *bool  `json:"isExclusive,omitempty" form:"isExclusive,omitempty"` //TODO: Write general description for this field
	Length      *int64 `json:"length,omitempty" form:"length,omitempty"`           //TODO: Write general description for this field
	Offset      *int64 `json:"offset,omitempty" form:"offset,omitempty"`           //TODO: Write general description for this field
}

/*
 * Structure for the custom type LogicalStats
 */
type LogicalStats struct {
	TotalLogicalUsageBytes *int64 `json:"totalLogicalUsageBytes,omitempty" form:"totalLogicalUsageBytes,omitempty"` //Provides the logical usage as computed by the Cohesity Cluster.
}

/*
 * Structure for the custom type LogicalVolume
 */
type LogicalVolume struct {
	DeviceRootNode *DeviceTreeDetails `json:"deviceRootNode,omitempty" form:"deviceRootNode,omitempty"` //Specifies a logical volume stored as a tree where the leaves are
	GroupName      *string            `json:"groupName,omitempty" form:"groupName,omitempty"`           //Specifies the group name of the logical volume.
	GroupUuid      *string            `json:"groupUuid,omitempty" form:"groupUuid,omitempty"`           //Specifies the group uuid of the logical volume.
	Name           *string            `json:"name,omitempty" form:"name,omitempty"`                     //Specifies the name of the logical volume.
	Uuid           *string            `json:"uuid,omitempty" form:"uuid,omitempty"`                     //Specifies the uuid of the logical volume.
}

/*
 * Structure for the custom type MORef
 */
type MORef struct {
	Item *string `json:"item,omitempty" form:"item,omitempty"` //TODO: Write general description for this field
	Type *string `json:"type,omitempty" form:"type,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type MSExchangeParams
 */
type MSExchangeParams struct {
	PerformLogTruncation *bool `json:"performLogTruncation,omitempty" form:"performLogTruncation,omitempty"` //If this is set to true, then an attempt will be made to truncate
}

/*
 * Structure for the custom type MagnetoInstanceId
 */
type MagnetoInstanceId struct {
	AttemptNum        *int64 `json:"attemptNum,omitempty" form:"attemptNum,omitempty"`               //The attempt of the job instance that took the snapshot.
	JobInstanceId     *int64 `json:"jobInstanceId,omitempty" form:"jobInstanceId,omitempty"`         //Instance of the job that took the snapshot.
	JobStartTimeUsecs *int64 `json:"jobStartTimeUsecs,omitempty" form:"jobStartTimeUsecs,omitempty"` //Start time of the job instance.
}

/*
 * Structure for the custom type MagnetoObjectId
 */
type MagnetoObjectId struct {
	Entity *EntityProto      `json:"entity,omitempty" form:"entity,omitempty"` //Specifies the attributes and the latest statistics about an entity.
	JobId  *int64            `json:"jobId,omitempty" form:"jobId,omitempty"`   //The id of the local job that the object belongs to, which may or may not
	JobUid *UniversalIdProto `json:"jobUid,omitempty" form:"jobUid,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type MapReduceAuxData
 */
type MapReduceAuxData struct {
	PatternVec []*Pattern `json:"patternVec,omitempty" form:"patternVec,omitempty"` //Pattern auxiliary data for a MapReduce.
}

/*
 * Structure for the custom type MetricDataBlock
 */
type MetricDataBlock struct {
	DataPointVec []*MetricDataPoint `json:"dataPointVec,omitempty" form:"dataPointVec,omitempty"` //Array of Data Points.
	MetricName   *string            `json:"metricName,omitempty" form:"metricName,omitempty"`     //Specifies the name of a metric such as 'kDiskAwaitTimeMsecs'.
	Type         *int64             `json:"type,omitempty" form:"type,omitempty"`                 //Specifies the data type of the data points.
}

/*
 * Structure for the custom type MetricDataPoint
 */
type MetricDataPoint struct {
	Data           *ValueData `json:"data,omitempty" form:"data,omitempty"`                     //Specifies the fields to store data of a given type.
	TimestampMsecs *int64     `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Specifies a timestamp when the metric data point was captured.
}

/*
 * Structure for the custom type MetricValue
 */
type MetricValue struct {
	MetricName     *string `json:"metricName,omitempty" form:"metricName,omitempty"`         //Specifies the metric name.
	TimestampMsecs *int64  `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Specifies the creation time of a data point as a Unix epoch Timestamp
	Value          *Value  `json:"value,omitempty" form:"value,omitempty"`                   //Specifies a data type and data field used to store data.
}

/*
 * Structure for the custom type MonthlySchedule
 */
type MonthlySchedule struct {
	Day      DayMonthlyScheduleEnum `json:"day,omitempty" form:"day,omitempty"`           //Specifies the day of the week (such as 'kMonday') to start the Job Run.
	DayCount DayCountEnum           `json:"dayCount,omitempty" form:"dayCount,omitempty"` //Specifies the day count in the month (such as 'kThird') to start
}

/*
 * Structure for the custom type MountVolumeResult
 */
type MountVolumeResult struct {
	Error              *ErrorProto `json:"error,omitempty" form:"error,omitempty"`                           //TODO: Write general description for this field
	FilesystemType     *string     `json:"filesystemType,omitempty" form:"filesystemType,omitempty"`         //Filesystem on this volume.
	MountPoint         *string     `json:"mountPoint,omitempty" form:"mountPoint,omitempty"`                 //This is populated with the mount point where the volume corresponding to
	OriginalVolumeName *string     `json:"originalVolumeName,omitempty" form:"originalVolumeName,omitempty"` //This is the name or mount point of the original volume.
}

/*
 * Structure for the custom type MountVolumeResultDetails
 */
type MountVolumeResultDetails struct {
	MountError *RequestError `json:"mountError,omitempty" form:"mountError,omitempty"` //Specifies the cause of the mount failure if the mounting of a
	MountPoint *string       `json:"mountPoint,omitempty" form:"mountPoint,omitempty"` //Specifies the mount point where the volume is mounted.
	VolumeName *string       `json:"volumeName,omitempty" form:"volumeName,omitempty"` //Specifies the name of the original volume.
}

/*
 * Structure for the custom type MountVolumesHypervParams
 */
type MountVolumesHypervParams struct {
	BringDisksOnline        *bool        `json:"bringDisksOnline,omitempty" form:"bringDisksOnline,omitempty"`               //Optional setting which will determine if the volumes need to be onlined
	TargetEntityCredentials *Credentials `json:"targetEntityCredentials,omitempty" form:"targetEntityCredentials,omitempty"` //Specifies credentials to access a target source.
}

/*
 * Structure for the custom type MountVolumesInfoProto
 */
type MountVolumesInfoProto struct {
	CleanupError              *ErrorProto                    `json:"cleanupError,omitempty" form:"cleanupError,omitempty"`                           //TODO: Write general description for this field
	Error                     *ErrorProto                    `json:"error,omitempty" form:"error,omitempty"`                                         //TODO: Write general description for this field
	Finished                  *bool                          `json:"finished,omitempty" form:"finished,omitempty"`                                   //This will be set to true if the task is complete on the slave.
	MountVolumeResultVec      []*MountVolumeResult           `json:"mountVolumeResultVec,omitempty" form:"mountVolumeResultVec,omitempty"`           //Contains the result information of the mounted volumes.
	RestoreDisksTaskInfoProto *SetupRestoreDiskTaskInfoProto `json:"restoreDisksTaskInfoProto,omitempty" form:"restoreDisksTaskInfoProto,omitempty"` //Each available extension is listed below along with the location of the
	SlaveTaskStartTimeUsecs   *int64                         `json:"slaveTaskStartTimeUsecs,omitempty" form:"slaveTaskStartTimeUsecs,omitempty"`     //This is the timestamp at which the slave task started.
	Type                      *int64                         `json:"type,omitempty" form:"type,omitempty"`                                           //The type of environment this mount volumes info pertains to.
	VmOnlineDisksError        *ErrorProto                    `json:"vmOnlineDisksError,omitempty" form:"vmOnlineDisksError,omitempty"`               //TODO: Write general description for this field
}

/*
 * Structure for the custom type MountVolumesParameters
 */
type MountVolumesParameters struct {
	BringDisksOnline *bool     `json:"bringDisksOnline,omitempty" form:"bringDisksOnline,omitempty"` //Optional setting that determines if the volumes are brought
	Password         *string   `json:"password,omitempty" form:"password,omitempty"`                 //Specifies password of the username to access the target source.
	TargetSourceId   *int64    `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"`     //Specifies the target Protection Source id where the volumes will be
	Username         *string   `json:"username,omitempty" form:"username,omitempty"`                 //Specifies username to access the target source.
	VolumeNames      *[]string `json:"volumeNames,omitempty" form:"volumeNames,omitempty"`           //Array of Volume Names.
}

/*
 * Structure for the custom type MountVolumesParams
 */
type MountVolumesParams struct {
	HypervParams  *MountVolumesHypervParams `json:"hypervParams,omitempty" form:"hypervParams,omitempty"`   //TODO: Write general description for this field
	ReadonlyMount *bool                     `json:"readonlyMount,omitempty" form:"readonlyMount,omitempty"` //Allows the caller to force the Agent to perform a read-only mount. This is
	TargetEntity  *EntityProto              `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`   //Specifies the attributes and the latest statistics about an entity.
	VmwareParams  *MountVolumesVmwareParams `json:"vmwareParams,omitempty" form:"vmwareParams,omitempty"`   //TODO: Write general description for this field
	VolumeNameVec *[]string                 `json:"volumeNameVec,omitempty" form:"volumeNameVec,omitempty"` //Optional names of volumes that need to be mounted. The names here
}

/*
 * Structure for the custom type MountVolumesState
 */
type MountVolumesState struct {
	BringDisksOnline   *bool                       `json:"bringDisksOnline,omitempty" form:"bringDisksOnline,omitempty"`     //Optional setting that determines if the volumes are brought
	MountVolumeResults []*MountVolumeResultDetails `json:"mountVolumeResults,omitempty" form:"mountVolumeResults,omitempty"` //Array of Mount Volume Results.
	OtherError         *RequestError               `json:"otherError,omitempty" form:"otherError,omitempty"`                 //Specifies an error that did not occur during the mount operation.
	TargetSourceId     *int64                      `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"`         //Specifies the target Protection Source Id where the volumes will be
	Username           *string                     `json:"username,omitempty" form:"username,omitempty"`                     //Specifies the username to access the mount target.
}

/*
 * Structure for the custom type MountVolumesTaskStateProto
 */
type MountVolumesTaskStateProto struct {
	FullNasPath *string                `json:"fullNasPath,omitempty" form:"fullNasPath,omitempty"` //Contains the SMB/NFS path of the share we expose to clients. The share
	HostEntity  *EntityProto           `json:"hostEntity,omitempty" form:"hostEntity,omitempty"`   //Specifies the attributes and the latest statistics about an entity.
	MountInfo   *MountVolumesInfoProto `json:"mountInfo,omitempty" form:"mountInfo,omitempty"`     //Each available extension is listed below along with the location of the
	MountParams *MountVolumesParams    `json:"mountParams,omitempty" form:"mountParams,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type MountVolumesVmwareParams
 */
type MountVolumesVmwareParams struct {
	BringDisksOnline        *bool        `json:"bringDisksOnline,omitempty" form:"bringDisksOnline,omitempty"`               //Optional setting which will determine if the volumes need to be onlined
	TargetEntityCredentials *Credentials `json:"targetEntityCredentials,omitempty" form:"targetEntityCredentials,omitempty"` //Specifies credentials to access a target source.
}

/*
 * Structure for the custom type NasBackupParams
 */
type NasBackupParams struct {
	ContinueOnError       *bool                 `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`             //Whether the backup job should continue on errors for snapshot based
	FilteringPolicy       *FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"`             //Proto to encapsulate the filtering policy for backup objects like files or
	MixedModePreference   *int64                `json:"mixedModePreference,omitempty" form:"mixedModePreference,omitempty"`     //If the target entity is a mixed mode volume, which NAS protocol type the
	SnapshotChangeEnabled *bool                 `json:"snapshotChangeEnabled,omitempty" form:"snapshotChangeEnabled,omitempty"` //Whether this backup job should utilize changelist like API when available
}

/*
 * Structure for the custom type NasCredentials
 */
type NasCredentials struct {
	Host      *string       `json:"host,omitempty" form:"host,omitempty"`           //Specifies the hostname or IP address of the NAS server.
	MountPath *string       `json:"mountPath,omitempty" form:"mountPath,omitempty"` //Specifies the mount path to the NAS server.
	Password  *string       `json:"password,omitempty" form:"password,omitempty"`   //If using the CIFS protocol and a Username was specified, specify
	ShareType ShareTypeEnum `json:"shareType,omitempty" form:"shareType,omitempty"` //Specifies the sharing protocol type used to mount the file system.
	Username  *string       `json:"username,omitempty" form:"username,omitempty"`   //If using the CIFS protocol, you can optional specify a username
}

/*
 * Structure for the custom type NasEnvJobParameters
 */
type NasEnvJobParameters struct {
	ContinueOnError                *bool                              `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`                               //Specifies if the backup should continue on with other Protection Sources
	DataMigrationJobParameters     *DataMigrationJobParameters        `json:"dataMigrationJobParameters,omitempty" form:"dataMigrationJobParameters,omitempty"`         //Specifies parameters applicable for data migration jobs in NAS environment.
	EnableFasterIncrementalBackups *bool                              `json:"enableFasterIncrementalBackups,omitempty" form:"enableFasterIncrementalBackups,omitempty"` //Specifies whether this job will enable faster incremental backups using
	FilePathFilters                *FilePathFilter                    `json:"filePathFilters,omitempty" form:"filePathFilters,omitempty"`                               //Specifies filters to match files and directories on a Server.
	NasProtocol                    NasProtocolNasEnvJobParametersEnum `json:"nasProtocol,omitempty" form:"nasProtocol,omitempty"`                                       //Specifies the preferred protocol to use for backup. This does not
}

/*
 * Structure for the custom type NasMountCredentialParams
 */
type NasMountCredentialParams struct {
	Domain         *string         `json:"domain,omitempty" form:"domain,omitempty"`                 //Specifies the domain in which this credential is valid.
	NasProtocol    NasProtocolEnum `json:"nasProtocol,omitempty" form:"nasProtocol,omitempty"`       //Specifies the protocol used by the NAS server.
	NasType        NasTypeEnum     `json:"nasType,omitempty" form:"nasType,omitempty"`               //Specifies the type of a NAS Object such as 'kGroup', or 'kHost'.
	Password       *string         `json:"password,omitempty" form:"password,omitempty"`             //Specifies the password for the username to use for mounting the NAS.
	SkipValidation *bool           `json:"skipValidation,omitempty" form:"skipValidation,omitempty"` //Specifies the flag to disable mount point validation during registration
	Username       *string         `json:"username,omitempty" form:"username,omitempty"`             //Specifies a username to use for mounting the NAS.
}

/*
 * Structure for the custom type NasProtectionSource
 */
type NasProtectionSource struct {
	Description    *string                         `json:"description,omitempty" form:"description,omitempty"`       //Specifies a description about the Protection Source.
	MountPath      *string                         `json:"mountPath,omitempty" form:"mountPath,omitempty"`           //Specifies the mount path of this NAS. For example, for a NFS mount point,
	Name           *string                         `json:"name,omitempty" form:"name,omitempty"`                     //Specifies the name of the NetApp Object.
	Protocol       ProtocolNasProtectionSourceEnum `json:"protocol,omitempty" form:"protocol,omitempty"`             //Specifies the protocol used by the NAS server.
	SkipValidation *bool                           `json:"skipValidation,omitempty" form:"skipValidation,omitempty"` //Specifies whether to skip validation of the given mount point.
	Type           TypeNasProtectionSourceEnum     `json:"type,omitempty" form:"type,omitempty"`                     //Specifies the type of a Protection Source Object in a generic NAS Source
}

/*
 * Structure for the custom type NetappClusterInfo
 */
type NetappClusterInfo struct {
	ContactInfo  *string `json:"contactInfo,omitempty" form:"contactInfo,omitempty"`   //Specifies information about the contact for the NetApp cluster
	Location     *string `json:"location,omitempty" form:"location,omitempty"`         //Specifies where this NetApp cluster is located.
	SerialNumber *string `json:"serialNumber,omitempty" form:"serialNumber,omitempty"` //Specifies the serial number of the NetApp cluster in the
}

/*
 * Structure for the custom type NetappProtectionSource
 */
type NetappProtectionSource struct {
	ClusterInfo *NetappClusterInfo             `json:"clusterInfo,omitempty" form:"clusterInfo,omitempty"` //Specifies information about a NetApp Cluster Protection Source.
	IsTopLevel  *bool                          `json:"isTopLevel,omitempty" form:"isTopLevel,omitempty"`   //Specifies if this Object is a top level Object.
	Name        *string                        `json:"name,omitempty" form:"name,omitempty"`               //Specifies the name of the NetApp Object.
	Type        TypeNetappProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`               //Specifies the type of managed NetApp Object in a NetApp Protection Source
	Uuid        *string                        `json:"uuid,omitempty" form:"uuid,omitempty"`               //Specifies the globally unique ID of this Object assigned by the
	VolumeInfo  *NetappVolumeInfo              `json:"volumeInfo,omitempty" form:"volumeInfo,omitempty"`   //Specifies information about a volume in a NetApp cluster.
	VserverInfo *NetappVserverInfo             `json:"vserverInfo,omitempty" form:"vserverInfo,omitempty"` //Specifies information about a NetApp Vserver in a NetApp Protection Source.
}

/*
 * Structure for the custom type NetappVolumeInfo
 */
type NetappVolumeInfo struct {
	AggregateName     *string                  `json:"aggregateName,omitempty" form:"aggregateName,omitempty"`         //Specifies the containing aggregate name of this volume.
	CapacityBytes     *int64                   `json:"capacityBytes,omitempty" form:"capacityBytes,omitempty"`         //Specifies the total capacity in bytes of this volume.
	CifsShares        []*CifsShareInfo         `json:"cifsShares,omitempty" form:"cifsShares,omitempty"`               //Array of CIFS Shares.
	CreationTimeUsecs *int64                   `json:"creationTimeUsecs,omitempty" form:"creationTimeUsecs,omitempty"` //Specifies the creation time of the volume specified in Unix epoch time
	DataProtocols     *[]DataProtocolEnum      `json:"dataProtocols,omitempty" form:"dataProtocols,omitempty"`         //Array of Data Protocols.
	ExportPolicyName  *string                  `json:"exportPolicyName,omitempty" form:"exportPolicyName,omitempty"`   //Specifies the name of the export policy (which defines the access
	JunctionPath      *string                  `json:"junctionPath,omitempty" form:"junctionPath,omitempty"`           //Specifies the junction path of this volume.
	Name              *string                  `json:"name,omitempty" form:"name,omitempty"`                           //Specifies the name of the NetApp Vserver that this volume belongs to.
	SecurityInfo      *VolumeSecurityInfo      `json:"securityInfo,omitempty" form:"securityInfo,omitempty"`           //Specifies information about NetApp volume security settings.
	State             StateEnum                `json:"state,omitempty" form:"state,omitempty"`                         //Specifies the state of this volume.
	Type              TypeNetappVolumeInfoEnum `json:"type,omitempty" form:"type,omitempty"`                           //Specifies the NetApp type of this volume.
	UsedBytes         *int64                   `json:"usedBytes,omitempty" form:"usedBytes,omitempty"`                 //Specifies the total space (in bytes) used in this volume.
}

/*
 * Structure for the custom type NetappVserverInfo
 */
type NetappVserverInfo struct {
	DataProtocols *[]DataProtocolEnum        `json:"dataProtocols,omitempty" form:"dataProtocols,omitempty"` //Array of Data Protocols.
	Interfaces    []*VserverNetworkInterface `json:"interfaces,omitempty" form:"interfaces,omitempty"`       //Array of Interfaces.
	RootCifsShare *CifsShareInfo             `json:"rootCifsShare,omitempty" form:"rootCifsShare,omitempty"` //Specifies information about a CIFS share of a NetApp volume.
	Type          TypeNetappVserverInfoEnum  `json:"type,omitempty" form:"type,omitempty"`                   //Specifies the type of this Vserver.
}

/*
 * Structure for the custom type NetworkConfiguration
 */
type NetworkConfiguration struct {
	ClusterGateway    *string   `json:"clusterGateway,omitempty" form:"clusterGateway,omitempty"`       //Specifies the default gateway IP address (or addresses) for the Cluster
	ClusterSubnetMask *string   `json:"clusterSubnetMask,omitempty" form:"clusterSubnetMask,omitempty"` //Specifies the subnet mask (or masks) of the Cluster network.
	DnsServers        *[]string `json:"dnsServers,omitempty" form:"dnsServers,omitempty"`               //Specifies the list of DNS Servers this cluster should be configured with.
	DomainNames       *[]string `json:"domainNames,omitempty" form:"domainNames,omitempty"`             //Specifies the list of domain names this cluster should be configured
	NtpServers        *[]string `json:"ntpServers,omitempty" form:"ntpServers,omitempty"`               //Specifies the list of NTP Servers this cluster should be configured with.
	VipHostname       *string   `json:"vipHostname,omitempty" form:"vipHostname,omitempty"`             //Specifies the virtual IP hostname.
	Vips              *[]string `json:"vips,omitempty" form:"vips,omitempty"`                           //Specifies the list of virtual IPs for the new cluster.
}

/*
 * Structure for the custom type NetworkInterface
 */
type NetworkInterface struct {
	BondSlaveSlotTypes *[]string                       `json:"bondSlaveSlotTypes,omitempty" form:"bondSlaveSlotTypes,omitempty"` //Specifies the types of the slots of any slaves if this interface is a
	BondSlaves         *[]string                       `json:"bondSlaves,omitempty" form:"bondSlaves,omitempty"`                 //Specifies the names of any slaves if this interface is a bond.
	BondingMode        BondingModeNetworkInterfaceEnum `json:"bondingMode,omitempty" form:"bondingMode,omitempty"`               //Specifies the bonding mode if this interface is a bond.
	Gateway            *string                         `json:"gateway,omitempty" form:"gateway,omitempty"`                       //Specifies the gateway of the interface.
	Group              *string                         `json:"group,omitempty" form:"group,omitempty"`                           //Specifies the group that this interface belongs to.
	Id                 *int64                          `json:"id,omitempty" form:"id,omitempty"`                                 //Specifies the ID of this network interface.
	IsConnected        *bool                           `json:"isConnected,omitempty" form:"isConnected,omitempty"`               //Specifies whether or not the Interface is connected.
	IsDefaultRoute     *bool                           `json:"isDefaultRoute,omitempty" form:"isDefaultRoute,omitempty"`         //Specifies whether or not to use this interface as the default route.
	IsUp               *bool                           `json:"isUp,omitempty" form:"isUp,omitempty"`                             //Specifies whether or not the interface is currently  up.
	MacAddress         *string                         `json:"macAddress,omitempty" form:"macAddress,omitempty"`                 //Specifies the Mac address of the Interface.
	Mtu                *int64                          `json:"mtu,omitempty" form:"mtu,omitempty"`                               //Specifies the MTU of the interface.
	Name               *string                         `json:"name,omitempty" form:"name,omitempty"`                             //Specifies the name of the interface port.
	Role               RoleNetworkInterfaceEnum        `json:"role,omitempty" form:"role,omitempty"`                             //Specifies the role of this interface.
	Services           *[]ServiceNetworkInterfaceEnum  `json:"services,omitempty" form:"services,omitempty"`                     //Specifies the types of services this interface is used for.
	Speed              *string                         `json:"speed,omitempty" form:"speed,omitempty"`                           //Specifies the speed of the Interface.
	StaticIp           *string                         `json:"staticIp,omitempty" form:"staticIp,omitempty"`                     //Specifies the static IP of the interface.
	Subnet             *string                         `json:"subnet,omitempty" form:"subnet,omitempty"`                         //Specifies the subnet mask of the interface.
	Type               TypeNetworkInterfaceEnum        `json:"type,omitempty" form:"type,omitempty"`                             //Specifies the type of interface.
	VirtualIp          *string                         `json:"virtualIp,omitempty" form:"virtualIp,omitempty"`                   //Specifies the virtual IP of the interface.
}

/*
 * Structure for the custom type NetworkMapping
 */
type NetworkMapping struct {
	DisableNetwork     *bool  `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"`         //Specifies if the network should be disabled. On restore or clone
	PreserveMacAddress *bool  `json:"preserveMacAddress,omitempty" form:"preserveMacAddress,omitempty"` //Specifies if the source mac address should be preserved after restore
	SourceNetworkId    *int64 `json:"sourceNetworkId,omitempty" form:"sourceNetworkId,omitempty"`       //Specifies the id of the source network.
	TargetNetworkId    *int64 `json:"targetNetworkId,omitempty" form:"targetNetworkId,omitempty"`       //Specifies the id of target network.
}

/*
 * Structure for the custom type NetworkMappingProto
 */
type NetworkMappingProto struct {
	DisableNetwork                 *bool        `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"`                                 //This can be set to true to indicate that the attached network should be
	PreserveMacAddressOnNewNetwork *bool        `json:"preserveMacAddressOnNewNetwork,omitempty" form:"preserveMacAddressOnNewNetwork,omitempty"` //VM's MAC address will be preserved on the new network. This value takes
	SourceNetworkEntity            *EntityProto `json:"sourceNetworkEntity,omitempty" form:"sourceNetworkEntity,omitempty"`                       //Specifies the attributes and the latest statistics about an entity.
	TargetNetworkEntity            *EntityProto `json:"targetNetworkEntity,omitempty" form:"targetNetworkEntity,omitempty"`                       //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type NetworkingInformation
 */
type NetworkingInformation struct {
	ResourceVec []*ClusterNetworkingResourceInformation `json:"resourceVec,omitempty" form:"resourceVec,omitempty"` //The list of resources on the the system that are accessible by an
}

/*
 * Structure for the custom type NewS3SecretAccessKey
 */
type NewS3SecretAccessKey struct {
	NewKey *string `json:"newKey,omitempty" form:"newKey,omitempty"` //Specifies the new S3 Secret Access key.
}

/*
 * Structure for the custom type NfsRootPermissions
 */
type NfsRootPermissions struct {
	Gid  *int64 `json:"gid,omitempty" form:"gid,omitempty"`   //Unix GID for the root of the file system.
	Mode *int64 `json:"mode,omitempty" form:"mode,omitempty"` //Unix mode bits for the root of the file system.
	Uid  *int64 `json:"uid,omitempty" form:"uid,omitempty"`   //Unix UID for the root of the file system.
}

/*
 * Structure for the custom type NlmLock
 */
type NlmLock struct {
	ClientId   *string      `json:"clientId,omitempty" form:"clientId,omitempty"`     //Specifies the client ID
	LockRanges []*LockRange `json:"lockRanges,omitempty" form:"lockRanges,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Node
 */
type Node struct {
	CapacityByTier           []*CapacityByTier     `json:"capacityByTier,omitempty" form:"capacityByTier,omitempty"`                     //CapacityByTier describes the capacity of each storage tier.
	ChassisInfo              *ChassisInfo          `json:"chassisInfo,omitempty" form:"chassisInfo,omitempty"`                           //ChassisInfo is the struct for the Chassis.
	ClusterPartitionId       *int64                `json:"clusterPartitionId,omitempty" form:"clusterPartitionId,omitempty"`             //ClusterPartitionId is the Id of the cluster partition to which
	ClusterPartitionName     *string               `json:"clusterPartitionName,omitempty" form:"clusterPartitionName,omitempty"`         //ClusterPartitionName is the name of the cluster to which the Node
	DiskCount                *int64                `json:"diskCount,omitempty" form:"diskCount,omitempty"`                               //DiskCount is the number of disks in a node.
	DiskCountByTier          []*CountByTier        `json:"diskCountByTier,omitempty" form:"diskCountByTier,omitempty"`                   //DiskCountByTier describes the disk number of each storage tier.
	Id                       *int64                `json:"id,omitempty" form:"id,omitempty"`                                             //Id is the Id of the Node.
	Ip                       *string               `json:"ip,omitempty" form:"ip,omitempty"`                                             //Ip is the IP address of the Node.
	IsMarkedForRemoval       *bool                 `json:"isMarkedForRemoval,omitempty" form:"isMarkedForRemoval,omitempty"`             //IsMarkedForRemoval specifies whether the node has been marked for
	MaxPhysicalCapacityBytes *int64                `json:"maxPhysicalCapacityBytes,omitempty" form:"maxPhysicalCapacityBytes,omitempty"` //MaxPhysicalCapacityBytes specifies the maximum physical capacity of the
	NodeHardwareInfo         *NodeHardwareInfo     `json:"nodeHardwareInfo,omitempty" form:"nodeHardwareInfo,omitempty"`                 //NodeHardwareInfo provides the information regarding the hardware.
	NodeIncarnationId        *int64                `json:"nodeIncarnationId,omitempty" form:"nodeIncarnationId,omitempty"`               //NodeIncarnationId is the incarnation id  of this node. The incarnation
	NodeSoftwareVersion      *string               `json:"nodeSoftwareVersion,omitempty" form:"nodeSoftwareVersion,omitempty"`           //NodeSoftwareVersion is the current version of Cohesity software installed
	OfflineMountPathsOfDisks *[]string             `json:"offlineMountPathsOfDisks,omitempty" form:"offlineMountPathsOfDisks,omitempty"` //OfflineMountPathsOfDisks provides the corresponding mount paths for
	RemovalReason            *[]RemovalReasonEnum  `json:"removalReason,omitempty" form:"removalReason,omitempty"`                       //RemovalReason specifies the removal reason of the node.
	RemovalState             RemovalStateEnum      `json:"removalState,omitempty" form:"removalState,omitempty"`                         //RemovalState specifies the removal state of the node.
	SlotNumber               *int64                `json:"slotNumber,omitempty" form:"slotNumber,omitempty"`                             //Slot number occupied by this node within the chassis.
	Stats                    *NodeStats            `json:"stats,omitempty" form:"stats,omitempty"`                                       //NodeStats provides various statistics for the node.
	SystemDisks              []*NodeSystemDiskInfo `json:"systemDisks,omitempty" form:"systemDisks,omitempty"`                           //SystemDisk describes the node system disks.
}

/*
 * Structure for the custom type NodeHardwareInfo
 */
type NodeHardwareInfo struct {
	Cpu             *string `json:"cpu,omitempty" form:"cpu,omitempty"`                         //Cpu provides the information regarding the CPU.
	MemorySizeBytes *int64  `json:"memorySizeBytes,omitempty" form:"memorySizeBytes,omitempty"` //MemorySizeBytes provides the memory size in bytes.
	Network         *string `json:"network,omitempty" form:"network,omitempty"`                 //Network provides the information regarding the network cards.
}

/*
 * Structure for the custom type NodeNetworkInterfaces
 */
type NodeNetworkInterfaces struct {
	ChassisSerial *string             `json:"chassisSerial,omitempty" form:"chassisSerial,omitempty"` //Specifies the serial number of Chassis.
	Id            *int64              `json:"id,omitempty" form:"id,omitempty"`                       //Specifies the ID of the Node.
	Interfaces    []*NetworkInterface `json:"interfaces,omitempty" form:"interfaces,omitempty"`       //Specifies the list of network interfaces present on this Node.
	Message       *string             `json:"message,omitempty" form:"message,omitempty"`             //Specifies an optional message describing the result of the request
	Slot          *int64              `json:"slot,omitempty" form:"slot,omitempty"`                   //Specifies the slot number the Node is located in.
}

/*
 * Structure for the custom type NodeStats
 */
type NodeStats struct {
	Id             *int64                    `json:"id,omitempty" form:"id,omitempty"`                         //Id is the Id of the Node.
	UsagePerfStats *UsageAndPerformanceStats `json:"usagePerfStats,omitempty" form:"usagePerfStats,omitempty"` //Provides usage and performance statistics
}

/*
 * Structure for the custom type NodeStatus
 */
type NodeStatus struct {
	ErrorMessage *string `json:"errorMessage,omitempty" form:"errorMessage,omitempty"` //Specifies an optional message relating to the node status.
	IpmiIp       *string `json:"ipmiIp,omitempty" form:"ipmiIp,omitempty"`             //Specifies the IPMI IP of the node (if physical cluster).
	NodeId       *int64  `json:"nodeId,omitempty" form:"nodeId,omitempty"`             //Specifies the ID of the node.
	NodeIp       *string `json:"nodeIp,omitempty" form:"nodeIp,omitempty"`             //For physical nodes this will specify the IP address of the node.
}

/*
 * Structure for the custom type NodeSystemDiskInfo
 */
type NodeSystemDiskInfo struct {
	DevicePath *string `json:"devicePath,omitempty" form:"devicePath,omitempty"` //DevicePath is the device path of the disk.
	Id         *int64  `json:"id,omitempty" form:"id,omitempty"`                 //Id is the id of the disk.
	Offline    *bool   `json:"offline,omitempty" form:"offline,omitempty"`       //Offline specifies whether a disk is marked offline.
}

/*
 * Structure for the custom type NotificationRule
 */
type NotificationRule struct {
	AlertTypeList          *[]int64                        `json:"alertTypeList,omitempty" form:"alertTypeList,omitempty"`                   //Specifies alert types this rule is applicable to.
	Categories             *[]CategoryNotificationRuleEnum `json:"categories,omitempty" form:"categories,omitempty"`                         //Specifies alert categories this rule is applicable to.
	EmailDeliveryTargets   []*EmailDeliveryTarget          `json:"emailDeliveryTargets,omitempty" form:"emailDeliveryTargets,omitempty"`     //Specifies email addresses to be notified when an alert matching this
	RuleId                 *int64                          `json:"ruleId,omitempty" form:"ruleId,omitempty"`                                 //Specifies id of the alert delivery rule.
	RuleName               *string                         `json:"ruleName,omitempty" form:"ruleName,omitempty"`                             //Specifies name of the alert delivery rule.
	Severities             *[]SeverityNotificationRuleEnum `json:"severities,omitempty" form:"severities,omitempty"`                         //Specifies alert severity types this rule is applicable to.
	TenantId               *string                         `json:"tenantId,omitempty" form:"tenantId,omitempty"`                             //Specifies tenant id this rule is applicable to.
	WebHookDeliveryTargets []*WebHookDeliveryTarget        `json:"webHookDeliveryTargets,omitempty" form:"webHookDeliveryTargets,omitempty"` //Specifies external api urls to be invoked when an alert matching this
}

/*
 * Structure for the custom type Notifications
 */
type Notifications struct {
	Count            *int64              `json:"count,omitempty" form:"count,omitempty"`                       //Notification Count.
	NotificationList []*TaskNotification `json:"notificationList,omitempty" form:"notificationList,omitempty"` //Notification list.
	UnreadCount      *int64              `json:"unreadCount,omitempty" form:"unreadCount,omitempty"`           //Unread Notification Count.
}

/*
 * Structure for the custom type NtpSettingsConfig
 */
type NtpSettingsConfig struct {
	NtpServersInternal *bool `json:"ntpServersInternal,omitempty" form:"ntpServersInternal,omitempty"` //Flag to specify if the NTP servers are on internal network or not.
}

/*
 * Structure for the custom type O365BackupEnvParams
 */
type O365BackupEnvParams struct {
	FilteringPolicy *FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"` //Proto to encapsulate the filtering policy for backup objects like files or
}

/*
 * Structure for the custom type ObjectSearchResults
 */
type ObjectSearchResults struct {
	ObjectSnapshotInfo []*ObjectSnapshotInfo `json:"objectSnapshotInfo,omitempty" form:"objectSnapshotInfo,omitempty"` //Array of Snapshot Objects.
	TotalCount         *int64                `json:"totalCount,omitempty" form:"totalCount,omitempty"`                 //Specifies the total number of backup objects that match the filter and
}

/*
 * Structure for the custom type ObjectSnapshotInfo
 */
type ObjectSnapshotInfo struct {
	ClusterPartitionId *int64             `json:"clusterPartitionId,omitempty" form:"clusterPartitionId,omitempty"` //Specifies the Cohesity Cluster partition id where this object is stored.
	JobId              *int64             `json:"jobId,omitempty" form:"jobId,omitempty"`                           //Specifies the id for the Protection Job that is currently
	JobName            *string            `json:"jobName,omitempty" form:"jobName,omitempty"`                       //Specifies the name of the Protection Job that captured the backup.
	JobUid             *UniversalId       `json:"jobUid,omitempty" form:"jobUid,omitempty"`                         //Specifies the globally unique id of the Protection Job that backed up
	ObjectName         *string            `json:"objectName,omitempty" form:"objectName,omitempty"`                 //Specifies the primary name of the object.
	OsType             *string            `json:"osType,omitempty" form:"osType,omitempty"`                         //Specifies the inferred OS type.
	RegisteredSource   *ProtectionSource  `json:"registeredSource,omitempty" form:"registeredSource,omitempty"`     //Specifies a generic structure that represents a node
	SnapshottedSource  *ProtectionSource  `json:"snapshottedSource,omitempty" form:"snapshottedSource,omitempty"`   //Specifies a generic structure that represents a node
	Versions           []*SnapshotVersion `json:"versions,omitempty" form:"versions,omitempty"`                     //Array of Snapshots.
	ViewBoxId          *int64             `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                   //Specifies the id of the Domain (View Box) where this
	ViewName           *string            `json:"viewName,omitempty" form:"viewName,omitempty"`                     //Specifies the View name where this object is stored.
}

/*
 * Structure for the custom type ObjectSnapshotType
 */
type ObjectSnapshotType struct {
	Msg  *string `json:"msg,omitempty" form:"msg,omitempty"`   //This captures any additional message about the snapshot itself, e.g. if
	Type *int64  `json:"type,omitempty" form:"type,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ObjectsByEnv
 */
type ObjectsByEnv struct {
	EnvType    *string `json:"envType,omitempty" form:"envType,omitempty"`       //Environment Type.
	NumObjects *int64  `json:"numObjects,omitempty" form:"numObjects,omitempty"` //Number of Objects.
}

/*
 * Structure for the custom type ObjectsProtectedByPolicy
 */
type ObjectsProtectedByPolicy struct {
	ObjectsProtected []*ObjectsByEnv `json:"objectsProtected,omitempty" form:"objectsProtected,omitempty"` //Protected Objects.
	PolicyId         *string         `json:"policyId,omitempty" form:"policyId,omitempty"`                 //Id of the policy.
	PolicyName       *string         `json:"policyName,omitempty" form:"policyName,omitempty"`             //Name of the policy.
}

/*
 * Structure for the custom type Office365Credentials
 */
type Office365Credentials struct {
	ClientId     *string `json:"clientId,omitempty" form:"clientId,omitempty"`         //Specifies the application ID that the registration portal
	ClientSecret *string `json:"clientSecret,omitempty" form:"clientSecret,omitempty"` //Specifies the application secret that was created in app registration
	GrantType    *string `json:"grantType,omitempty" form:"grantType,omitempty"`       //Specifies the application grant type. eg: For client credentials flow, set
	Scope        *string `json:"scope,omitempty" form:"scope,omitempty"`               //Specifies a space separated list of scopes/permissions for the user.
}

/*
 * Structure for the custom type Office365ProtectionSource
 */
type Office365ProtectionSource struct {
	Description        *string                           `json:"description,omitempty" form:"description,omitempty"`               //Specifies the description of the Office 365 entity.
	Name               *string                           `json:"name,omitempty" form:"name,omitempty"`                             //Specifies the name of the office 365 entity.
	PrimarySMTPAddress *string                           `json:"primarySMTPAddress,omitempty" form:"primarySMTPAddress,omitempty"` //Specifies the SMTP address for the Outlook source.
	Type               TypeOffice365ProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                             //Specifies the type of the Office 365 entity.
	Uuid               *string                           `json:"uuid,omitempty" form:"uuid,omitempty"`                             //Specifies the UUID of the Office 365 entity.
}

/*
 * Structure for the custom type OracleAppParams
 */
type OracleAppParams struct {
	DatabaseAppId   *int64                       `json:"databaseAppId,omitempty" form:"databaseAppId,omitempty"`     //Specifies the source entity id of the selected app entity.
	NodeChannelList []*OracleDatabaseNodeChannel `json:"nodeChannelList,omitempty" form:"nodeChannelList,omitempty"` //Array of database node channel info.
}

/*
 * Structure for the custom type OracleCloudCredentials
 */
type OracleCloudCredentials struct {
	AccessKeyId     *string                            `json:"accessKeyId,omitempty" form:"accessKeyId,omitempty"`         //Specifies access key to connect to Oracle S3 Compatible vault account.
	Region          *string                            `json:"region,omitempty" form:"region,omitempty"`                   //Specifies the region for Oracle S3 Compatible vault account.
	SecretAccessKey *string                            `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"` //Specifies the secret access key for Oracle S3 Compatible vault account.
	Tenant          *string                            `json:"tenant,omitempty" form:"tenant,omitempty"`                   //Specifies the tenant which is part of the REST endpoints for Oracle S3
	TierType        TierTypeOracleCloudCredentialsEnum `json:"tierType,omitempty" form:"tierType,omitempty"`               //Specifies the storage class of Oracle vault.
}

/*
 * Structure for the custom type OracleDBChannelInfo
 */
type OracleDBChannelInfo struct {
	DbUniqueName *string                        `json:"dbUniqueName,omitempty" form:"dbUniqueName,omitempty"` //The unique name of the database.
	DbUuid       *string                        `json:"dbUuid,omitempty" form:"dbUuid,omitempty"`             //Database id, internal field, is filled by magneto master based on
	HostInfoVec  []*OracleDBChannelInfoHostInfo `json:"hostInfoVec,omitempty" form:"hostInfoVec,omitempty"`   //Vector of Oracle hosts from which we are allowed to take the
	MaxNumHost   *int64                         `json:"maxNumHost,omitempty" form:"maxNumHost,omitempty"`     //Maximum number of hosts from which we are allowed to take backup/restore
	NumChannels  *int64                         `json:"numChannels,omitempty" form:"numChannels,omitempty"`   //The default number of channels to use per host per db. This value is used
}

/*
 * Structure for the custom type OracleDBChannelInfoHostInfo
 */
type OracleDBChannelInfoHostInfo struct {
	Host        *string `json:"host,omitempty" form:"host,omitempty"`               //Host string from which we are allowed to take the backup/restore.
	NumChannels *int64  `json:"numChannels,omitempty" form:"numChannels,omitempty"` //Number of channels we need to create for this host. Default value for
	Portnum     *int64  `json:"portnum,omitempty" form:"portnum,omitempty"`         //port number where database is listening.
}

/*
 * Structure for the custom type OracleDBConfig
 */
type OracleDBConfig struct {
	AuditLogDest         *string                         `json:"auditLogDest,omitempty" form:"auditLogDest,omitempty"`                 //Audit log destination.
	BctFilePath          *string                         `json:"bctFilePath,omitempty" form:"bctFilePath,omitempty"`                   //BCT file path.
	ControlFilePathVec   *[]string                       `json:"controlFilePathVec,omitempty" form:"controlFilePathVec,omitempty"`     //List of paths where the control file needs to be multiplexed.
	DbConfigFilePath     *string                         `json:"dbConfigFilePath,omitempty" form:"dbConfigFilePath,omitempty"`         //Path to the file on oracle host which decides the configuration of
	DiagDest             *string                         `json:"diagDest,omitempty" form:"diagDest,omitempty"`                         //Diag destination.
	EnableArchiveLogMode *bool                           `json:"enableArchiveLogMode,omitempty" form:"enableArchiveLogMode,omitempty"` //If set to false, archive log mode is disabled.
	FraDest              *string                         `json:"fraDest,omitempty" form:"fraDest,omitempty"`                           //FRA path.
	FraSizeMb            *int64                          `json:"fraSizeMb,omitempty" form:"fraSizeMb,omitempty"`                       //FRA Size in MBs.
	NumTempfiles         *int64                          `json:"numTempfiles,omitempty" form:"numTempfiles,omitempty"`                 //How many tempfiles to use for the recovered database.
	RedoLogConf          *OracleDBConfigRedoLogGroupConf `json:"redoLogConf,omitempty" form:"redoLogConf,omitempty"`                   //GROUP1 : {DST1/CH1.log, DST2/CH1.log}
	SgaTargetSize        *string                         `json:"sgaTargetSize,omitempty" form:"sgaTargetSize,omitempty"`               //SGA_TARGET_SIZE size [ Default value same as Source DB ].
	SharedPoolSize       *string                         `json:"sharedPoolSize,omitempty" form:"sharedPoolSize,omitempty"`             //Shared pool size [ Default value same as Source DB ].
}

/*
 * Structure for the custom type OracleDBConfigRedoLogGroupConf
 */
type OracleDBConfigRedoLogGroupConf struct {
	GroupMemberVec *[]string `json:"groupMemberVec,omitempty" form:"groupMemberVec,omitempty"` //List of members of this redo log group.
	MemberPrefix   *string   `json:"memberPrefix,omitempty" form:"memberPrefix,omitempty"`     //Log member name prefix.
	NumGroups      *int64    `json:"numGroups,omitempty" form:"numGroups,omitempty"`           //Number of redo log groups.
	SizeMb         *int64    `json:"sizeMb,omitempty" form:"sizeMb,omitempty"`                 //Size of the member in MB.
}

/*
 * Structure for the custom type OracleDatabaseNode
 */
type OracleDatabaseNode struct {
	ChannelCount *int64  `json:"channelCount,omitempty" form:"channelCount,omitempty"` //Specifies the number of channels user wants for the backup/recovery
	Node         *string `json:"node,omitempty" form:"node,omitempty"`                 //Specifies the ip of the database node.
	Port         *int64  `json:"port,omitempty" form:"port,omitempty"`                 //Specifies the port on which user wants to run the backup/recovery.
}

/*
 * Structure for the custom type OracleDatabaseNodeChannel
 */
type OracleDatabaseNodeChannel struct {
	DatabaseNodeList    []*OracleDatabaseNode `json:"databaseNodeList,omitempty" form:"databaseNodeList,omitempty"`       //Array of nodes of a database.
	DatabaseUniqueName  *string               `json:"databaseUniqueName,omitempty" form:"databaseUniqueName,omitempty"`   //Specifies the unique Name of the database.
	DatabaseUuid        *string               `json:"databaseUuid,omitempty" form:"databaseUuid,omitempty"`               //Specifies the database unique id. This is an internal field and is filled
	DefaultChannelCount *int64                `json:"defaultChannelCount,omitempty" form:"defaultChannelCount,omitempty"` //Specifies the default number of channels to use per node per database.
	MaxNodeCount        *int64                `json:"maxNodeCount,omitempty" form:"maxNodeCount,omitempty"`               //Specifies the maximum number of nodes from which we are allowed to take
}

/*
 * Structure for the custom type OracleHost
 */
type OracleHost struct {
	CpuCount    *int64           `json:"cpuCount,omitempty" form:"cpuCount,omitempty"`       //Specifies the count of CPU available on the host.
	IpAddresses *[]string        `json:"ipAddresses,omitempty" form:"ipAddresses,omitempty"` //Specifies the IP address of the host.
	Ports       *[]int64         `json:"ports,omitempty" form:"ports,omitempty"`             //Specifies ports available for this host.
	Sessions    []*OracleSession `json:"sessions,omitempty" form:"sessions,omitempty"`       //Specifies multiple session configurations available for this host.
}

/*
 * Structure for the custom type OracleProtectionSource
 */
type OracleProtectionSource struct {
	ArchiveLogEnabled *bool                          `json:"archiveLogEnabled,omitempty" form:"archiveLogEnabled,omitempty"` //Specifies whether the database is running in ARCHIVELOG mode. It enables
	BctEnabled        *bool                          `json:"bctEnabled,omitempty" form:"bctEnabled,omitempty"`               //Specifies whether the Block Change Tracking is enabled. BCT improves the
	DbType            DbTypeEnum                     `json:"dbType,omitempty" form:"dbType,omitempty"`                       //Specifies the type of the database in Oracle Protection Source.
	FraSize           *int64                         `json:"fraSize,omitempty" form:"fraSize,omitempty"`                     //Specfies Flash/Fast Recovery area size for the current DB entity.
	Hosts             []*OracleHost                  `json:"hosts,omitempty" form:"hosts,omitempty"`                         //Specifies the list of hosts for the current DB entity.
	Name              *string                        `json:"name,omitempty" form:"name,omitempty"`                           //Specifies the instance name of the Oracle entity.
	OwnerId           *int64                         `json:"ownerId,omitempty" form:"ownerId,omitempty"`                     //Specifies the entity id of the owner entity (such as a VM). This is only
	SgaTargetSize     *string                        `json:"sgaTargetSize,omitempty" form:"sgaTargetSize,omitempty"`         //Specifies System Global Area size for the current DB entity.
	SharedPoolSize    *string                        `json:"sharedPoolSize,omitempty" form:"sharedPoolSize,omitempty"`       //Specifies Shared Pool Size for the current DB entity.
	Size              *int64                         `json:"size,omitempty" form:"size,omitempty"`                           //Specifies database size.
	TempFilesCount    *int64                         `json:"tempFilesCount,omitempty" form:"tempFilesCount,omitempty"`       //Specifies number of temporary files for the current DB entity.
	Type              TypeOracleProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                           //Specifies the type of the managed Object in Oracle Protection Source.
	Uuid              *string                        `json:"uuid,omitempty" form:"uuid,omitempty"`                           //Specifies the UUID for the Oracle entity.
	Version           *string                        `json:"version,omitempty" form:"version,omitempty"`                     //Specifies the Oracle database instance version.
}

/*
 * Structure for the custom type OracleSession
 */
type OracleSession struct {
	Location         *string `json:"location,omitempty" form:"location,omitempty"`                 //Location is the path where Oracle is installed.
	SystemIdentifier *string `json:"systemIdentifier,omitempty" form:"systemIdentifier,omitempty"` //SystemIdentifier is the unique Oracle System Identifier for the DB instance.
}

/*
 * Structure for the custom type OracleSourceParams
 */
type OracleSourceParams struct {
	AdditionalOracleDbParamsVec []*AdditionalOracleDBParams `json:"additionalOracleDbParamsVec,omitempty" form:"additionalOracleDbParamsVec,omitempty"` //Backup channel information for each Oracle database.
}

/*
 * Structure for the custom type OracleSpecialParameters
 */
type OracleSpecialParameters struct {
	AppParamsList        []*OracleAppParams `json:"appParamsList,omitempty" form:"appParamsList,omitempty"`               //Array of application parameters i.e. database parameters for
	ApplicationEntityIds *[]int64           `json:"applicationEntityIds,omitempty" form:"applicationEntityIds,omitempty"` //Array of Ids of Application Entities like Oracle instances, and
}

/*
 * Structure for the custom type OutlookBackupEnvParams
 */
type OutlookBackupEnvParams struct {
	FilteringPolicy *FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"` //Proto to encapsulate the filtering policy for backup objects like files or
}

/*
 * Structure for the custom type OutlookEnvJobParameters
 */
type OutlookEnvJobParameters struct {
	FilePathFilter *FilePathFilter `json:"filePathFilter,omitempty" form:"filePathFilter,omitempty"` //Specifies filters to match files and directories on a Server.
}

/*
 * Structure for the custom type OutlookFolder
 */
type OutlookFolder struct {
	FolderId            *string   `json:"folderId,omitempty" form:"folderId,omitempty"`                       //Specifies the unique ID of the folder.
	FolderKey           *int64    `json:"folderKey,omitempty" form:"folderKey,omitempty"`                     //Specifies the key unique within the mailbox of the folder.
	OutlookItemIdList   *[]string `json:"outlookItemIdList,omitempty" form:"outlookItemIdList,omitempty"`     //Specifies the outlook items within the folder to be restored incase the
	RestoreEntireFolder *bool     `json:"restoreEntireFolder,omitempty" form:"restoreEntireFolder,omitempty"` //Specifies whether the entire folder is to be restored.
}

/*
 * Structure for the custom type OutlookMailbox
 */
type OutlookMailbox struct {
	MailboxObject        *RestoreObjectDetails `json:"mailboxObject,omitempty" form:"mailboxObject,omitempty"`               //Specifies an object to recover or clone or an object to restore files
	OutlookFolderList    []*OutlookFolder      `json:"outlookFolderList,omitempty" form:"outlookFolderList,omitempty"`       //Specifies the list of folders to be restored incase user wishes not to
	RestoreEntireMailbox *bool                 `json:"restoreEntireMailbox,omitempty" form:"restoreEntireMailbox,omitempty"` //Specifies whether the enitre mailbox is to be restored.
}

/*
 * Structure for the custom type OutlookRestoreParameters
 */
type OutlookRestoreParameters struct {
	OutlookMailboxList []*OutlookMailbox `json:"outlookMailboxList,omitempty" form:"outlookMailboxList,omitempty"` //Specifies the list of mailboxes to be restored.
	TargetFolderPath   *string           `json:"targetFolderPath,omitempty" form:"targetFolderPath,omitempty"`     //Specifies the target folder path to restore the mailboxes. This will
	TargetMailbox      *ProtectionSource `json:"targetMailbox,omitempty" form:"targetMailbox,omitempty"`           //Specifies a generic structure that represents a node
}

/*
 * Structure for the custom type OutputSpec
 */
type OutputSpec struct {
	NumReduceShards    *int64  `json:"numReduceShards,omitempty" form:"numReduceShards,omitempty"`       //Number of reduce shards.
	OutputDir          *string `json:"outputDir,omitempty" form:"outputDir,omitempty"`                   //Name of output directory.
	PartitionId        *int64  `json:"partitionId,omitempty" form:"partitionId,omitempty"`               //Partition id where output will go.
	ReduceOutputPrefix *string `json:"reduceOutputPrefix,omitempty" form:"reduceOutputPrefix,omitempty"` //Prefix of the reduce output files. File names will be:
	ViewBoxId          *int64  `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                   //Viewbox id where the output will go.
	ViewName           *string `json:"viewName,omitempty" form:"viewName,omitempty"`                     //Name of the view where output will go. This will be filled up by yoda.
}

/*
 * Structure for the custom type OverwriteViewParam
 */
type OverwriteViewParam struct {
	SourceViewName string `json:"sourceViewName" form:"sourceViewName"` //Specifies the source view name.
	TargetViewName string `json:"targetViewName" form:"targetViewName"` //Specifies the target view name.
}

/*
 * Structure for the custom type PackageDetails
 */
type PackageDetails struct {
	DowntimeRequired *bool    `json:"downtimeRequired,omitempty" form:"downtimeRequired,omitempty"` //Specifies whether or not downtime is required to update to this package.
	InstalledOnNodes *[]int64 `json:"installedOnNodes,omitempty" form:"installedOnNodes,omitempty"` //Specifies the list of IDs of nodes on the cluster where this package is
	PackageName      *string  `json:"packageName,omitempty" form:"packageName,omitempty"`           //Specifies the name of the current package.
	ReleaseDate      *string  `json:"releaseDate,omitempty" form:"releaseDate,omitempty"`           //Specifies the release date of this package.
}

/*
 * Structure for the custom type Pattern
 */
type Pattern struct {
	IsSystemDefined *bool   `json:"isSystemDefined,omitempty" form:"isSystemDefined,omitempty"` //Whether this pattern is system defined.
	Name            *string `json:"name,omitempty" form:"name,omitempty"`                       //Name of the pattern. This is marked optional but is required.
	Type            *int64  `json:"type,omitempty" form:"type,omitempty"`                       //Pattern type.
	Value           *string `json:"value,omitempty" form:"value,omitempty"`                     //Value of the pattern. This is marked optional but is required.
}

/*
 * Structure for the custom type PerformRestoreJobStateProto
 */
type PerformRestoreJobStateProto struct {
	AdmittedTimeUsecs                 *int64                                    `json:"admittedTimeUsecs,omitempty" form:"admittedTimeUsecs,omitempty"`                                 //The time at which the restore job was admitted to run on a Magneto master.
	CancellationRequested             *bool                                     `json:"cancellationRequested,omitempty" form:"cancellationRequested,omitempty"`                         //Whether this restore job has a pending cancellation request.
	ContinueRestoreOnError            *bool                                     `json:"continueRestoreOnError,omitempty" form:"continueRestoreOnError,omitempty"`                       //Whether to continue with the restore operation if restore of any object
	DeployVmsToCloudTaskState         *DeployVMsToCloudTaskStateProto           `json:"deployVmsToCloudTaskState,omitempty" form:"deployVmsToCloudTaskState,omitempty"`                 //TODO: Write general description for this field
	EndTimeUsecs                      *int64                                    `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                           //If the restore job has finished, this field contains the end time for the
	Error                             *ErrorProto                               `json:"error,omitempty" form:"error,omitempty"`                                                         //TODO: Write general description for this field
	Name                              *string                                   `json:"name,omitempty" form:"name,omitempty"`                                                           //The name of the restore job.
	ParentSourceConnectionParams      *ConnectorParams                          `json:"parentSourceConnectionParams,omitempty" form:"parentSourceConnectionParams,omitempty"`           //Message that encapsulates the various params required to establish a
	PowerStateConfig                  *PowerStateConfigProto                    `json:"powerStateConfig,omitempty" form:"powerStateConfig,omitempty"`                                   //TODO: Write general description for this field
	ProgressMonitorTaskPath           *string                                   `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"`                     //Root path of a Pulse task tracking the progress of the restore job.
	RenameRestoredObjectParam         *RenameObjectParamProto                   `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"`                 //Message to specify the prefix/suffix added to rename an object. At least one
	RestoreAcropolisVmsParams         *RestoreAcropolisVMsParams                `json:"restoreAcropolisVmsParams,omitempty" form:"restoreAcropolisVmsParams,omitempty"`                 //TODO: Write general description for this field
	RestoreJobId                      *int64                                    `json:"restoreJobId,omitempty" form:"restoreJobId,omitempty"`                                           //A globally unique id for this restore job.
	RestoreKubernetesNamespacesParams *RestoreKubernetesNamespacesParams        `json:"restoreKubernetesNamespacesParams,omitempty" form:"restoreKubernetesNamespacesParams,omitempty"` //TODO: Write general description for this field
	RestoreKvmVmsParams               *RestoreKVMVMsParams                      `json:"restoreKvmVmsParams,omitempty" form:"restoreKvmVmsParams,omitempty"`                             //TODO: Write general description for this field
	RestoreParentSource               *EntityProto                              `json:"restoreParentSource,omitempty" form:"restoreParentSource,omitempty"`                             //Specifies the attributes and the latest statistics about an entity.
	RestoreTaskVec                    []*PerformRestoreJobStateProtoRestoreTask `json:"restoreTaskVec,omitempty" form:"restoreTaskVec,omitempty"`                                       //Even if the user wanted to restore an entire job from the latest snapshot,
	RestoreVmwareVmParams             *RestoreVmwareVMParams                    `json:"restoreVmwareVmParams,omitempty" form:"restoreVmwareVmParams,omitempty"`                         //TODO: Write general description for this field
	RestoredObjectsNetworkConfig      *RestoredObjectNetworkConfigProto         `json:"restoredObjectsNetworkConfig,omitempty" form:"restoredObjectsNetworkConfig,omitempty"`           //TODO: Write general description for this field
	RestoredToDifferentSource         *bool                                     `json:"restoredToDifferentSource,omitempty" form:"restoredToDifferentSource,omitempty"`                 //Whether restore is being performed to a different parent source.
	StartTimeUsecs                    *int64                                    `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                                       //The start time for this restore job.
	Status                            *int64                                    `json:"status,omitempty" form:"status,omitempty"`                                                       //Status of the restore job.
	Type                              *int64                                    `json:"type,omitempty" form:"type,omitempty"`                                                           //The type of restore being performed.
	User                              *string                                   `json:"user,omitempty" form:"user,omitempty"`                                                           //The user who requested this restore job.
	UserInfo                          *UserInformation                          `json:"userInfo,omitempty" form:"userInfo,omitempty"`                                                   //A message to encapsulate information about the user who made the request.
	ViewBoxId                         *int64                                    `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                                                 //The view box id to which the restore job belongs to.
}

/*
 * Structure for the custom type PerformRestoreJobStateProtoRestoreTask
 */
type PerformRestoreJobStateProtoRestoreTask struct {
	Object                        *RestoreObject `json:"object,omitempty" form:"object,omitempty"`                                               //TODO: Write general description for this field
	ObjectProgressMonitorTaskPath *string        `json:"objectProgressMonitorTaskPath,omitempty" form:"objectProgressMonitorTaskPath,omitempty"` //The relative task path of the progress monitor for the restore of the
	TaskId                        *int64         `json:"taskId,omitempty" form:"taskId,omitempty"`                                               //Id of the task tracking the restore of the above 'object'.
}

/*
 * Structure for the custom type PerformRestoreTaskStateProto
 */
type PerformRestoreTaskStateProto struct {
	Base                                   *RestoreTaskStateBaseProto         `json:"base,omitempty" form:"base,omitempty"`                                                                     //TODO: Write general description for this field
	CanTeardown                            *bool                              `json:"canTeardown,omitempty" form:"canTeardown,omitempty"`                                                       //This is set if the clone operation has created any objects on the primary
	ChildCloneTaskId                       *int64                             `json:"childCloneTaskId,omitempty" form:"childCloneTaskId,omitempty"`                                             //The id of the child clone task triggered by refresh op.
	ChildDestroyTaskId                     *int64                             `json:"childDestroyTaskId,omitempty" form:"childDestroyTaskId,omitempty"`                                         //The following fields are used by clone refresh op. These will be present
	CloneAppViewInfo                       *CloneAppViewInfoProto             `json:"cloneAppViewInfo,omitempty" form:"cloneAppViewInfo,omitempty"`                                             //This message encapsulates the information of Clone operation of backup view
	CloudDeployInfo                        *CloudDeployInfoProto              `json:"cloudDeployInfo,omitempty" form:"cloudDeployInfo,omitempty"`                                               //Each available extension is listed below along with the location of the
	ContinueRestoreOnError                 *bool                              `json:"continueRestoreOnError,omitempty" form:"continueRestoreOnError,omitempty"`                                 //Whether to continue with the restore operation if restore of any object
	CreateView                             *bool                              `json:"createView,omitempty" form:"createView,omitempty"`                                                         //True iff the target view needs to be created.
	DatastoreEntity                        *EntityProto                       `json:"datastoreEntity,omitempty" form:"datastoreEntity,omitempty"`                                               //Specifies the attributes and the latest statistics about an entity.
	DeployVmsToCloudTaskState              *DeployVMsToCloudTaskStateProto    `json:"deployVmsToCloudTaskState,omitempty" form:"deployVmsToCloudTaskState,omitempty"`                           //TODO: Write general description for this field
	FolderEntity                           *EntityProto                       `json:"folderEntity,omitempty" form:"folderEntity,omitempty"`                                                     //Specifies the attributes and the latest statistics about an entity.
	FullViewName                           *string                            `json:"fullViewName,omitempty" form:"fullViewName,omitempty"`                                                     //The full view name (internal or external). This is composed of an optional
	MountVolumesTaskState                  *MountVolumesTaskStateProto        `json:"mountVolumesTaskState,omitempty" form:"mountVolumesTaskState,omitempty"`                                   //TODO: Write general description for this field
	ObjectNameDEPRECATED                   *string                            `json:"objectName_DEPRECATED,omitempty" form:"objectName_DEPRECATED,omitempty"`                                   //An optional name to give to the restored object.
	Objects                                []*RestoreObject                   `json:"objects,omitempty" form:"objects,omitempty"`                                                               //Information on the exact set of objects being restored (along with
	ObjectsProgressMonitorTaskPaths        *[]string                          `json:"objectsProgressMonitorTaskPaths,omitempty" form:"objectsProgressMonitorTaskPaths,omitempty"`               //Vector containing the relative task path of progress monitors of the
	ParentRestoreJobId                     *int64                             `json:"parentRestoreJobId,omitempty" form:"parentRestoreJobId,omitempty"`                                         //If this a child restore task, this field will contain the id of the parent
	ParentRestoreTaskId                    *int64                             `json:"parentRestoreTaskId,omitempty" form:"parentRestoreTaskId,omitempty"`                                       //The id of the parent restore task if this is a restore sub-task.
	PathPrefixDEPRECATED                   *string                            `json:"pathPrefix_DEPRECATED,omitempty" form:"pathPrefix_DEPRECATED,omitempty"`                                   //TODO: Write general description for this field
	PowerStateConfig                       *PowerStateConfigProto             `json:"powerStateConfig,omitempty" form:"powerStateConfig,omitempty"`                                             //TODO: Write general description for this field
	ProgressMonitorTaskPath                *string                            `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"`                               //Root path of a Pulse task tracking the progress of the restore task.
	RecoverDisksTaskState                  *RecoverDisksTaskStateProto        `json:"recoverDisksTaskState,omitempty" form:"recoverDisksTaskState,omitempty"`                                   //TODO: Write general description for this field
	RecoverVolumesTaskState                *RecoverVolumesTaskStateProto      `json:"recoverVolumesTaskState,omitempty" form:"recoverVolumesTaskState,omitempty"`                               //TODO: Write general description for this field
	RelatedRestoreTaskId                   *int64                             `json:"relatedRestoreTaskId,omitempty" form:"relatedRestoreTaskId,omitempty"`                                     //The task id of a related restore task. For example, a SQL restore
	RenameRestoredObjectParam              *RenameObjectParamProto            `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"`                           //Message to specify the prefix/suffix added to rename an object. At least one
	RenameRestoredVappParam                *RenameObjectParamProto            `json:"renameRestoredVappParam,omitempty" form:"renameRestoredVappParam,omitempty"`                               //Message to specify the prefix/suffix added to rename an object. At least one
	ResourcePoolEntity                     *EntityProto                       `json:"resourcePoolEntity,omitempty" form:"resourcePoolEntity,omitempty"`                                         //Specifies the attributes and the latest statistics about an entity.
	RestoreAcropolisVmsParams              *RestoreAcropolisVMsParams         `json:"restoreAcropolisVmsParams,omitempty" form:"restoreAcropolisVmsParams,omitempty"`                           //TODO: Write general description for this field
	RestoreAppTaskState                    *RestoreAppTaskStateProto          `json:"restoreAppTaskState,omitempty" form:"restoreAppTaskState,omitempty"`                                       //TODO: Write general description for this field
	RestoreFilesTaskState                  *RestoreFilesTaskStateProto        `json:"restoreFilesTaskState,omitempty" form:"restoreFilesTaskState,omitempty"`                                   //TODO: Write general description for this field
	RestoreHypervVmParams                  *RestoreHypervVMParams             `json:"restoreHypervVmParams,omitempty" form:"restoreHypervVmParams,omitempty"`                                   //TODO: Write general description for this field
	RestoreInfo                            *RestoreInfoProto                  `json:"restoreInfo,omitempty" form:"restoreInfo,omitempty"`                                                       //Each available extension is listed below along with the location of the
	RestoreKubernetesNamespacesParams      *RestoreKubernetesNamespacesParams `json:"restoreKubernetesNamespacesParams,omitempty" form:"restoreKubernetesNamespacesParams,omitempty"`           //TODO: Write general description for this field
	RestoreKvmVmsParams                    *RestoreKVMVMsParams               `json:"restoreKvmVmsParams,omitempty" form:"restoreKvmVmsParams,omitempty"`                                       //TODO: Write general description for this field
	RestoreOneDriveParams                  *RestoreOneDriveParams             `json:"restoreOneDriveParams,omitempty" form:"restoreOneDriveParams,omitempty"`                                   //TODO: Write general description for this field
	RestoreOutlookParams                   *RestoreOutlookParams              `json:"restoreOutlookParams,omitempty" form:"restoreOutlookParams,omitempty"`                                     //TODO: Write general description for this field
	RestoreParentSource                    *EntityProto                       `json:"restoreParentSource,omitempty" form:"restoreParentSource,omitempty"`                                       //Specifies the attributes and the latest statistics about an entity.
	RestoreSubTaskVec                      *[]int64                           `json:"restoreSubTaskVec,omitempty" form:"restoreSubTaskVec,omitempty"`                                           //Inside Magneto, these are represented as regular restore tasks with their
	RestoreTaskPurged                      *bool                              `json:"restoreTaskPurged,omitempty" form:"restoreTaskPurged,omitempty"`                                           //Whether the restore task is purged. During WAL recovery, purged restore
	RestoreViewDatastoreEntity             *EntityProto                       `json:"restoreViewDatastoreEntity,omitempty" form:"restoreViewDatastoreEntity,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	RestoreVmwareVmParams                  *RestoreVmwareVMParams             `json:"restoreVmwareVmParams,omitempty" form:"restoreVmwareVmParams,omitempty"`                                   //TODO: Write general description for this field
	RestoredObjectsNetworkConfig           *RestoredObjectNetworkConfigProto  `json:"restoredObjectsNetworkConfig,omitempty" form:"restoredObjectsNetworkConfig,omitempty"`                     //TODO: Write general description for this field
	RestoredToDifferentSource              *bool                              `json:"restoredToDifferentSource,omitempty" form:"restoredToDifferentSource,omitempty"`                           //Whether restore is being performed to a different parent source.
	RetrieveArchiveProgressMonitorTaskPath *string                            `json:"retrieveArchiveProgressMonitorTaskPath,omitempty" form:"retrieveArchiveProgressMonitorTaskPath,omitempty"` //The path of the progress monitor for the task that is responsible for
	RetrieveArchiveStubViewName            *string                            `json:"retrieveArchiveStubViewName,omitempty" form:"retrieveArchiveStubViewName,omitempty"`                       //The stub view created by Icebox corresponding to the archive. The stub
	RetrieveArchiveTask                    *RetrieveArchiveTaskStateProto     `json:"retrieveArchiveTask,omitempty" form:"retrieveArchiveTask,omitempty"`                                       //Persistent state of a retrieve of an archive task. Only one of either
	RetrieveArchiveTaskUid                 *UniversalIdProto                  `json:"retrieveArchiveTaskUid,omitempty" form:"retrieveArchiveTaskUid,omitempty"`                                 //TODO: Write general description for this field
	RetrieveArchiveViewName                *string                            `json:"retrieveArchiveViewName,omitempty" form:"retrieveArchiveViewName,omitempty"`                               //The temporary view where the entities that have been retrieved from an
	StubViewRelativeDirName                *string                            `json:"stubViewRelativeDirName,omitempty" form:"stubViewRelativeDirName,omitempty"`                               //Relative directory inside the stub view that corresponds with the archive.
	VaultRestoreParams                     *VaultParamsRestoreParams          `json:"vaultRestoreParams,omitempty" form:"vaultRestoreParams,omitempty"`                                         //TODO: Write general description for this field
	VcdConfig                              *RestoredObjectVCDConfigProto      `json:"vcdConfig,omitempty" form:"vcdConfig,omitempty"`                                                           //TODO: Write general description for this field
	ViewBoxId                              *int64                             `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                                                           //The view box id to which 'view_name' belongs to. In case the restore task
	ViewNameDEPRECATED                     *string                            `json:"viewName_DEPRECATED,omitempty" form:"viewName_DEPRECATED,omitempty"`                                       //The view name as provided by the user for this restore operation.
	ViewParams                             *ViewParams                        `json:"viewParams,omitempty" form:"viewParams,omitempty"`                                                         //TODO(mark): Move this to magneto.proto.
	VolumeInfoVec                          []*VolumeInfo                      `json:"volumeInfoVec,omitempty" form:"volumeInfoVec,omitempty"`                                                   //Information regarding volumes that are required for the restore task. This
}

/*
 * Structure for the custom type PhysicalBackupEnvParams
 */
type PhysicalBackupEnvParams struct {
	EnableIncrementalBackupAfterRestart *bool                 `json:"enableIncrementalBackupAfterRestart,omitempty" form:"enableIncrementalBackupAfterRestart,omitempty"` //If this is set to true, then incremental backup will be performed
	FilteringPolicy                     *FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"`                                         //Proto to encapsulate the filtering policy for backup objects like files or
}

/*
 * Structure for the custom type PhysicalBackupSourceParams
 */
type PhysicalBackupSourceParams struct {
	EnableSystemBackup *bool                     `json:"enableSystemBackup,omitempty" form:"enableSystemBackup,omitempty"` //Allows Magneto to drive a "system" backup using a 3rd-party tool installed
	FileBackupParams   *PhysicalFileBackupParams `json:"fileBackupParams,omitempty" form:"fileBackupParams,omitempty"`     //Message to capture params when backing up files on a Physical source.
	SnapshotParams     *PhysicalSnapshotParams   `json:"snapshotParams,omitempty" form:"snapshotParams,omitempty"`         //This message contains params that controls the snapshot process for a
	SourceAppParams    *SourceAppParams          `json:"sourceAppParams,omitempty" form:"sourceAppParams,omitempty"`       //This message contains params specific to application running on the source
	VolumeGuidVec      *[]string                 `json:"volumeGuidVec,omitempty" form:"volumeGuidVec,omitempty"`           //If this list is non-empty, then only volumes in this will be
}

/*
 * Structure for the custom type PhysicalEnvJobParameters
 */
type PhysicalEnvJobParameters struct {
	FilePathFilters                *FilePathFilter `json:"filePathFilters,omitempty" form:"filePathFilters,omitempty"`                               //Specifies filters to match files and directories on a Server.
	IncrementalSnapshotUponRestart *bool           `json:"incrementalSnapshotUponRestart,omitempty" form:"incrementalSnapshotUponRestart,omitempty"` //If true, performs an incremental backup after server restarts. Otherwise
}

/*
 * Structure for the custom type PhysicalFileBackupParams
 */
type PhysicalFileBackupParams struct {
	BackupPathInfoVec        []*PhysicalFileBackupParamsBackupPathInfo `json:"backupPathInfoVec,omitempty" form:"backupPathInfoVec,omitempty"`               //Specifies the paths to backup on the Physical source.
	SkipNestedVolumesVec     *[]string                                 `json:"skipNestedVolumesVec,omitempty" form:"skipNestedVolumesVec,omitempty"`         //Mount types of nested volumes to be skipped.
	UsesSkipNestedVolumesVec *bool                                     `json:"usesSkipNestedVolumesVec,omitempty" form:"usesSkipNestedVolumesVec,omitempty"` //Specifies whether to use skip_nested_volumes_vec to skip nested mounts.
}

/*
 * Structure for the custom type PhysicalFileBackupParamsBackupPathInfo
 */
type PhysicalFileBackupParamsBackupPathInfo struct {
	ExcludePaths      *[]string `json:"excludePaths,omitempty" form:"excludePaths,omitempty"`           //A list of absolute paths on the Physical source that should not be
	IncludePath       *string   `json:"includePath,omitempty" form:"includePath,omitempty"`             //An absolute path on the Physical source that should be backed up. Any
	SkipNestedVolumes *bool     `json:"skipNestedVolumes,omitempty" form:"skipNestedVolumes,omitempty"` //Whether to skip any nested volumes (both local and network) that are
}

/*
 * Structure for the custom type PhysicalNodeConfiguration
 */
type PhysicalNodeConfiguration struct {
	NodeId     *int64  `json:"nodeId,omitempty" form:"nodeId,omitempty"`         //Specifies the Node ID for this node.
	NodeIp     *string `json:"nodeIp,omitempty" form:"nodeIp,omitempty"`         //Specifies the Node IP address for this node.
	NodeIpmiIp *string `json:"nodeIpmiIp,omitempty" form:"nodeIpmiIp,omitempty"` //Specifies IPMI IP for this node.
}

/*
 * Structure for the custom type PhysicalProtectionSource
 */
type PhysicalProtectionSource struct {
	Agents          []*AgentInformation                  `json:"agents,omitempty" form:"agents,omitempty"`                   //Array of Agents on the Physical Protection Source.
	HostName        *string                              `json:"hostName,omitempty" form:"hostName,omitempty"`               //Specifies the hostname.
	HostType        HostTypePhysicalProtectionSourceEnum `json:"hostType,omitempty" form:"hostType,omitempty"`               //Specifies the environment type for the host.
	Id              *UniversalId                         `json:"id,omitempty" form:"id,omitempty"`                           //Specifies a unique id of a Physical Protection Source.
	MemorySizeBytes *int64                               `json:"memorySizeBytes,omitempty" form:"memorySizeBytes,omitempty"` //Specifies the total memory ont the host in bytes.
	Name            *string                              `json:"name,omitempty" form:"name,omitempty"`                       //Specifies a human readable name of the Protection Source.
	NetworkingInfo  *NetworkingInformation               `json:"networkingInfo,omitempty" form:"networkingInfo,omitempty"`   //Specifies the struct containing information about network addresses
	NumProcessors   *int64                               `json:"numProcessors,omitempty" form:"numProcessors,omitempty"`     //Specifies the number of processors on the host.
	OsName          *string                              `json:"osName,omitempty" form:"osName,omitempty"`                   //Specifies a human readable name of the OS of the Protection Source.
	Type            TypePhysicalProtectionSourceEnum     `json:"type,omitempty" form:"type,omitempty"`                       //Specifies the type of managed Object in a Physical Protection Source.
	Volumes         []*PhysicalVolume                    `json:"volumes,omitempty" form:"volumes,omitempty"`                 //Array of Physical Volumes.
}

/*
 * Structure for the custom type PhysicalSnapshotParams
 */
type PhysicalSnapshotParams struct {
	FetchSnapshotMetadataDisabled *bool     `json:"fetchSnapshotMetadataDisabled,omitempty" form:"fetchSnapshotMetadataDisabled,omitempty"` //Whether fetching and storing of snapshot metadata was disabled.
	NotifyBackupCompleteDisabled  *bool     `json:"notifyBackupCompleteDisabled,omitempty" form:"notifyBackupCompleteDisabled,omitempty"`   //Whether notify backup complete step was disabled.
	VssCopyOnlyBackup             *bool     `json:"vssCopyOnlyBackup,omitempty" form:"vssCopyOnlyBackup,omitempty"`                         //If copy_only_backup option is requrested at the time of the snapshot.
	VssExcludedWriters            *[]string `json:"vssExcludedWriters,omitempty" form:"vssExcludedWriters,omitempty"`                       //List of VSS writers that were excluded.
}

/*
 * Structure for the custom type PhysicalSpecialParameters
 */
type PhysicalSpecialParameters struct {
	ApplicationParameters    *ApplicationParameters         `json:"applicationParameters,omitempty" form:"applicationParameters,omitempty"`       //TODO: Write general description for this field
	EnableSystemBackup       *bool                          `json:"enableSystemBackup,omitempty" form:"enableSystemBackup,omitempty"`             //Specifies whether to allow system backup using 3rd party tools installed
	FilePaths                []*FilePathParameters          `json:"filePaths,omitempty" form:"filePaths,omitempty"`                               //Array of File Paths to Back Up.
	SkipNestedVolumesVec     *[]string                      `json:"skipNestedVolumesVec,omitempty" form:"skipNestedVolumesVec,omitempty"`         //Specifies mounttypes of nested volumes to be skipped.
	UsesSkipNestedVolumesVec *bool                          `json:"usesSkipNestedVolumesVec,omitempty" form:"usesSkipNestedVolumesVec,omitempty"` //Specifies whether to use SkipNestedVolumes vec to skip nested mounts.
	VolumeGuid               *[]string                      `json:"volumeGuid,omitempty" form:"volumeGuid,omitempty"`                             //Array of Mounted Volumes to Back Up.
	WindowsParameters        *WindowsHostSnapshotParameters `json:"windowsParameters,omitempty" form:"windowsParameters,omitempty"`               //Specifies settings that are meaningful only on Windows hosts.
}

/*
 * Structure for the custom type PhysicalVolume
 */
type PhysicalVolume struct {
	DevicePath                    *string   `json:"devicePath,omitempty" form:"devicePath,omitempty"`                                       //Specifies the path to the device that hosts the volume locally.
	Guid                          *string   `json:"guid,omitempty" form:"guid,omitempty"`                                                   //Specifies an id for the Physical Volume.
	IsBootVolume                  *bool     `json:"isBootVolume,omitempty" form:"isBootVolume,omitempty"`                                   //Specifies whether the volume is boot volume.
	IsExtendedAttributesSupported *bool     `json:"isExtendedAttributesSupported,omitempty" form:"isExtendedAttributesSupported,omitempty"` //Specifies whether this volume supports extended attributes (like ACLs)
	IsProtected                   *bool     `json:"isProtected,omitempty" form:"isProtected,omitempty"`                                     //Specifies if a volume is protected by a Job.
	Label                         *string   `json:"label,omitempty" form:"label,omitempty"`                                                 //Specifies a volume label that can be used for displaying additional
	LogicalSizeBytes              *int64    `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"`                           //Specifies the logical size of the volume in bytes that is
	MountPoints                   *[]string `json:"mountPoints,omitempty" form:"mountPoints,omitempty"`                                     //Array of Mount Points.
	MountType                     *string   `json:"mountType,omitempty" form:"mountType,omitempty"`                                         //Specifies mount type of volume e.g. nfs, autofs, ext4 etc.
	NetworkPath                   *string   `json:"networkPath,omitempty" form:"networkPath,omitempty"`                                     //Specifies the full path to connect to the network attached volume.
	UsedSizeBytes                 *int64    `json:"usedSizeBytes,omitempty" form:"usedSizeBytes,omitempty"`                                 //Specifies the size used by the volume in bytes.
}

/*
 * Structure for the custom type PostgresNodeInfo
 */
type PostgresNodeInfo struct {
	DefaultPassword *string `json:"defaultPassword,omitempty" form:"defaultPassword,omitempty"` //Specifies the default password to access the postgres database.
	DefaultUsername *string `json:"defaultUsername,omitempty" form:"defaultUsername,omitempty"` //Specifies the default username to access the postgres database.
	NodeId          *int64  `json:"nodeId,omitempty" form:"nodeId,omitempty"`                   //Specifies the id of the node where postgres database is running.
	NodeIp          *string `json:"nodeIp,omitempty" form:"nodeIp,omitempty"`                   //Specifies the ip of the node where postgres database is running.
	Port            *int64  `json:"port,omitempty" form:"port,omitempty"`                       //Specifies the information where postgres database is running.
}

/*
 * Structure for the custom type PowerStateConfigProto
 */
type PowerStateConfigProto struct {
	PowerOn *bool `json:"powerOn,omitempty" form:"powerOn,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Preferences
 */
type Preferences struct {
	Locale *string `json:"locale,omitempty" form:"locale,omitempty"` //Locale reflects the language settings of the user. Populate using the
}

/*
 * Structure for the custom type PreferredDomainController
 */
type PreferredDomainController struct {
	DomainControllers *[]string `json:"domainControllers,omitempty" form:"domainControllers,omitempty"` //List of Domain controllers DCs in FQDN format that are mapped to an Active
	DomainName        *string   `json:"domainName,omitempty" form:"domainName,omitempty"`               //Specifies the Domain name or the trusted domain of an Active Directory.
}

/*
 * Structure for the custom type Principal
 */
type Principal struct {
	Domain        *string         `json:"domain,omitempty" form:"domain,omitempty"`               //Specifies the domain name of the where the principal' account is
	FullName      *string         `json:"fullName,omitempty" form:"fullName,omitempty"`           //Specifies the full name (first and last names) of the principal.
	ObjectClass   ObjectClassEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"`     //Specifies the object class of the principal (either 'kGroup' or 'kUser').
	PrincipalName *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the principal.
	Sid           *string         `json:"sid,omitempty" form:"sid,omitempty"`                     //Specifies the unique Security id (SID) of the principal.
}

/*
 * Structure for the custom type PrivilegeInfo
 */
type PrivilegeInfo struct {
	PrivilegeId         PrivilegeIdEnum `json:"PrivilegeId,omitempty" form:"PrivilegeId,omitempty"`                 //Specifies unique id for a privilege. This number must be unique when
	Category            *string         `json:"category,omitempty" form:"category,omitempty"`                       //Specifies a category for the privilege such as 'Access Management'.
	Description         *string         `json:"description,omitempty" form:"description,omitempty"`                 //Specifies a description defining what the privilege provides.
	IsAvailableOnHelios *bool           `json:"isAvailableOnHelios,omitempty" form:"isAvailableOnHelios,omitempty"` //Specifies that this privilege is available for Helios operations.
	IsCustomRoleDefault *bool           `json:"isCustomRoleDefault,omitempty" form:"isCustomRoleDefault,omitempty"` //Specifies if this privilege is automatically assigned to custom roles.
	IsSpecial           *bool           `json:"isSpecial,omitempty" form:"isSpecial,omitempty"`                     //Specifies if this privilege is automatically assigned to the default
	IsViewOnly          *bool           `json:"isViewOnly,omitempty" form:"isViewOnly,omitempty"`                   //Specifies if privilege is view-only privilege that cannot make changes.
	Label               *string         `json:"label,omitempty" form:"label,omitempty"`                             //Specifies the label for the privilege as displayed on the Cohesity
	Name                *string         `json:"name,omitempty" form:"name,omitempty"`                               //Specifies the Cluster name for the privilege such as PRINCIPAL_VIEW.
}

/*
 * Structure for the custom type ProductModelInterfaceTuple
 */
type ProductModelInterfaceTuple struct {
	IfaceName        *string `json:"ifaceName,omitempty" form:"ifaceName,omitempty"`               //Specifies the name of the interface.
	ProductModelName *string `json:"productModelName,omitempty" form:"productModelName,omitempty"` //Specifies the product model name.
}

/*
 * Structure for the custom type ProtectObjectParameters
 */
type ProtectObjectParameters struct {
	ProtectionSourceEnvironment ProtectionSourceEnvironmentEnum `json:"protectionSourceEnvironment,omitempty" form:"protectionSourceEnvironment,omitempty"` //Specifies the environment type of the Protection Source object.
	ProtectionSourceIds         []int64                         `json:"protectionSourceIds" form:"protectionSourceIds"`                                     //Specifies the ids of the Protection Sources to protect.
	RpoPolicyId                 string                          `json:"rpoPolicyId" form:"rpoPolicyId"`                                                     //Specifies the Rpo policy id.
}

/*
 * Structure for the custom type ProtectedObject
 */
type ProtectedObject struct {
	JobId                    *UniversalId `json:"jobId,omitempty" form:"jobId,omitempty"`                                       //Specifies an id for an object that is unique across Cohesity Clusters.
	ProtectionFauilureReason *string      `json:"protectionFauilureReason,omitempty" form:"protectionFauilureReason,omitempty"` //If protection fails then specifies why the protection failed on this
	ProtectionSourceId       *int64       `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"`             //Specifies the id of the Protection Source.
}

/*
 * Structure for the custom type ProtectedObjectsByEnv
 */
type ProtectedObjectsByEnv struct {
	EnvType              *string `json:"envType,omitempty" form:"envType,omitempty"`                           //Environment Type.
	ProtectedCount       *int64  `json:"protectedCount,omitempty" form:"protectedCount,omitempty"`             //Number of Protected Objects.
	ProtectedSizeBytes   *int64  `json:"protectedSizeBytes,omitempty" form:"protectedSizeBytes,omitempty"`     //Size of Protected Objects.
	UnprotectedCount     *int64  `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"`         //Number of Unprotected Objects.
	UnprotectedSizeBytes *int64  `json:"unprotectedSizeBytes,omitempty" form:"unprotectedSizeBytes,omitempty"` //Size of Unprotected Objects.
}

/*
 * Structure for the custom type ProtectedObjectsSummary
 */
type ProtectedObjectsSummary struct {
	NumObjectsProtected   *int64                          `json:"numObjectsProtected,omitempty" form:"numObjectsProtected,omitempty"`     //Specifies the total number of protected objects.
	NumObjectsUnprotected *int64                          `json:"numObjectsUnprotected,omitempty" form:"numObjectsUnprotected,omitempty"` //Specifies the total number of unprotected objects.
	ProtectedSizeBytes    *int64                          `json:"protectedSizeBytes,omitempty" form:"protectedSizeBytes,omitempty"`       //Specifies the total size of protected objects in bytes.
	StatsByEnv            []*ProtectedObjectsSummaryByEnv `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"`                       //Specifies the stats of Protected objects by environment.
	UnprotectedSizeBytes  *int64                          `json:"unprotectedSizeBytes,omitempty" form:"unprotectedSizeBytes,omitempty"`   //Specifies the total size of unprotected objects in bytes.
}

/*
 * Structure for the custom type ProtectedObjectsSummaryByEnv
 */
type ProtectedObjectsSummaryByEnv struct {
	Environment           EnvironmentProtectedObjectsSummaryByEnvEnum `json:"environment,omitempty" form:"environment,omitempty"`                     //Specifies the environment.
	NumObjectsProtected   *int64                                      `json:"numObjectsProtected,omitempty" form:"numObjectsProtected,omitempty"`     //Specifies the total number of protected objects.
	NumObjectsUnprotected *int64                                      `json:"numObjectsUnprotected,omitempty" form:"numObjectsUnprotected,omitempty"` //Specifies the total number of unprotected objects.
	ProtectedSizeBytes    *int64                                      `json:"protectedSizeBytes,omitempty" form:"protectedSizeBytes,omitempty"`       //Specifies the total size of protected objects in bytes.
	UnprotectedSizeBytes  *int64                                      `json:"unprotectedSizeBytes,omitempty" form:"unprotectedSizeBytes,omitempty"`   //Specifies the total size of unprotected objects in bytes.
}

/*
 * Structure for the custom type ProtectedObjectsTile
 */
type ProtectedObjectsTile struct {
	ObjectsProtected     []*ProtectedObjectsByEnv `json:"objectsProtected,omitempty" form:"objectsProtected,omitempty"`         //Protected Objects breakdown by object type.
	ProtectedCount       *int64                   `json:"protectedCount,omitempty" form:"protectedCount,omitempty"`             //Number of Protected Objects.
	ProtectedSizeBytes   *int64                   `json:"protectedSizeBytes,omitempty" form:"protectedSizeBytes,omitempty"`     //Size of Protected Objects.
	UnprotectedCount     *int64                   `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"`         //Number of Unprotected Objects.
	UnprotectedSizeBytes *int64                   `json:"unprotectedSizeBytes,omitempty" form:"unprotectedSizeBytes,omitempty"` //Size of Unprotected Objects.
}

/*
 * Structure for the custom type ProtectedSourceSummary
 */
type ProtectedSourceSummary struct {
	BackupRun                  *BackupRun                `json:"backupRun,omitempty" form:"backupRun,omitempty"`                                   //Specifies details about the Backup task for a Job Run.
	CopyRuns                   []*CopyRun                `json:"copyRuns,omitempty" form:"copyRuns,omitempty"`                                     //Specifies details about the Copy tasks of the Job Run.
	IsPaused                   *bool                     `json:"isPaused,omitempty" form:"isPaused,omitempty"`                                     //Specifies the status of the backup job.
	NextProtectionRunTimeUsecs *int64                    `json:"nextProtectionRunTimeUsecs,omitempty" form:"nextProtectionRunTimeUsecs,omitempty"` //Specifies the time at which the next Protection Run is scheduled for the
	ProtectedSourceUid         *UniversalId              `json:"protectedSourceUid,omitempty" form:"protectedSourceUid,omitempty"`                 //Specifies an id for an object that is unique across Cohesity Clusters.
	ProtectionSource           *ProtectionSource         `json:"protectionSource,omitempty" form:"protectionSource,omitempty"`                     //Specifies a generic structure that represents a node
	SourceParameters           []*SourceSpecialParameter `json:"sourceParameters,omitempty" form:"sourceParameters,omitempty"`                     //Specifies additional special settings for a single Protected Source.
}

/*
 * Structure for the custom type ProtectedVmInfo
 */
type ProtectedVmInfo struct {
	ProtectionJobs     []*ProtectionJob    `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"`         //Specifies the list of Protection Jobs that protect the VM.
	ProtectionPolicies []*ProtectionPolicy `json:"protectionPolicies,omitempty" form:"protectionPolicies,omitempty"` //Specifies the list of Policies that are used by the Protection Jobs.
	ProtectionSource   *ProtectionSource   `json:"protectionSource,omitempty" form:"protectionSource,omitempty"`     //Specifies a generic structure that represents a node
	Stats              *ProtectionSummary  `json:"stats,omitempty" form:"stats,omitempty"`                           //Specifies the protection stats of VM.
}

/*
 * Structure for the custom type ProtectionInfo
 */
type ProtectionInfo struct {
	EndTimeUsecs      *int64  `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`           //Specifies the end time for object retention.
	Location          *string `json:"location,omitempty" form:"location,omitempty"`                   //Specifies the location of the object.
	PolicyId          *string `json:"policyId,omitempty" form:"policyId,omitempty"`                   //Specifies the id of the policy.
	ProtectionJobId   *int64  `json:"protectionJobId,omitempty" form:"protectionJobId,omitempty"`     //Specifies the id of the protection job.
	ProtectionJobName *string `json:"protectionJobName,omitempty" form:"protectionJobName,omitempty"` //Specifies the protection job name which protects this object.
	RetentionPeriod   *int64  `json:"retentionPeriod,omitempty" form:"retentionPeriod,omitempty"`     //Specifies the retention period.
	StartTimeUsecs    *int64  `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`       //Specifies the start time for object retention.
	StorageDomain     *string `json:"storageDomain,omitempty" form:"storageDomain,omitempty"`         //Specifies the storage domain name.
	TotalSnapshots    *int64  `json:"totalSnapshots,omitempty" form:"totalSnapshots,omitempty"`       //Specifies the total number of snapshots.
}

/*
 * Structure for the custom type ProtectionJob
 */
type ProtectionJob struct {
	AbortInBlackoutPeriod                *bool                         `json:"abortInBlackoutPeriod,omitempty" form:"abortInBlackoutPeriod,omitempty"`                               //If true, the Cohesity Cluster aborts any currently executing Job Runs
	AlertingConfig                       *AlertingConfig               `json:"alertingConfig,omitempty" form:"alertingConfig,omitempty"`                                             //Specifies optional settings for alerting.
	AlertingPolicy                       *[]AlertingPolicyEnum         `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"`                                             //Array of Job Events.
	CloudParameters                      *CloudParameters              `json:"cloudParameters,omitempty" form:"cloudParameters,omitempty"`                                           //Specifies Cloud parameters that are applicable to all Protection
	ContinueOnQuiesceFailure             *bool                         `json:"continueOnQuiesceFailure,omitempty" form:"continueOnQuiesceFailure,omitempty"`                         //Whether to continue backing up on quiesce failure.
	CreateRemoteView                     *bool                         `json:"createRemoteView,omitempty" form:"createRemoteView,omitempty"`                                         //Specifies whether to create a remote view name to use for view overwrite.
	CreationTimeUsecs                    *int64                        `json:"creationTimeUsecs,omitempty" form:"creationTimeUsecs,omitempty"`                                       //Specifies the time when the Protection Job was created.
	DataMigrationPolicy                  *DataMigrationPolicy          `json:"dataMigrationPolicy,omitempty" form:"dataMigrationPolicy,omitempty"`                                   //Specifies settings for data migration in NAS environment. This also
	DedupDisabledSourceIds               *[]int64                      `json:"dedupDisabledSourceIds,omitempty" form:"dedupDisabledSourceIds,omitempty"`                             //List of source ids for which source side dedup is disabled from the backup
	Description                          *string                       `json:"description,omitempty" form:"description,omitempty"`                                                   //Specifies a text description about the Protection Job.
	EndTimeUsecs                         *int64                        `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                                 //Specifies the epoch time (in microseconds) after which the Protection Job
	Environment                          EnvironmentProtectionJobEnum  `json:"environment,omitempty" form:"environment,omitempty"`                                                   //Specifies the environment type (such as kVMware or kSQL)
	EnvironmentParameters                *EnvironmentTypeJobParameters `json:"environmentParameters,omitempty" form:"environmentParameters,omitempty"`                               //Specifies additional parameters that are common to all Protection
	ExcludeSourceIds                     *[]int64                      `json:"excludeSourceIds,omitempty" form:"excludeSourceIds,omitempty"`                                         //Array of Excluded Source Objects.
	ExcludeVmTagIds                      *[]int64                      `json:"excludeVmTagIds,omitempty" form:"excludeVmTagIds,omitempty"`                                           //Array of Arrays of VM Tag Ids that Specify VMs to Exclude.
	FullProtectionSlaTimeMins            *int64                        `json:"fullProtectionSlaTimeMins,omitempty" form:"fullProtectionSlaTimeMins,omitempty"`                       //If specified, this setting is number of minutes that a Job Run
	FullProtectionStartTime              *TimeOfDay                    `json:"fullProtectionStartTime,omitempty" form:"fullProtectionStartTime,omitempty"`                           //Specifies the time of day to start the Full Protection Schedule.
	Id                                   *int64                        `json:"id,omitempty" form:"id,omitempty"`                                                                     //Specifies an id for the Protection Job.
	IncrementalProtectionSlaTimeMins     *int64                        `json:"incrementalProtectionSlaTimeMins,omitempty" form:"incrementalProtectionSlaTimeMins,omitempty"`         //If specified, this setting is number of minutes that a Job Run
	IncrementalProtectionStartTime       *TimeOfDay                    `json:"incrementalProtectionStartTime,omitempty" form:"incrementalProtectionStartTime,omitempty"`             //Specifies the time of day to start the CBT-based Protection Schedule.
	IndexingPolicy                       *IndexingPolicy               `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"`                                             //Specifies settings for indexing files found in an Object
	IsActive                             *bool                         `json:"isActive,omitempty" form:"isActive,omitempty"`                                                         //Indicates if the current state of the Protection Job is Active
	IsDeleted                            *bool                         `json:"isDeleted,omitempty" form:"isDeleted,omitempty"`                                                       //Equals 'true' if the Protection Job was deleted but some Snapshots
	IsPaused                             *bool                         `json:"isPaused,omitempty" form:"isPaused,omitempty"`                                                         //Indicates if the Protection Job is paused, which means that no new
	LastRun                              *ProtectionRunInstance        `json:"lastRun,omitempty" form:"lastRun,omitempty"`                                                           //Specifies the status of one Job Run.
	LeverageStorageSnapshots             *bool                         `json:"leverageStorageSnapshots,omitempty" form:"leverageStorageSnapshots,omitempty"`                         //Specifies whether to leverage the storage array based snapshots for this
	LeverageStorageSnapshotsForHyperflex *bool                         `json:"leverageStorageSnapshotsForHyperflex,omitempty" form:"leverageStorageSnapshotsForHyperflex,omitempty"` //Specifies whether to leverage Hyperflex as the storage snapshot array
	ModificationTimeUsecs                *int64                        `json:"modificationTimeUsecs,omitempty" form:"modificationTimeUsecs,omitempty"`                               //Specifies the last time this Job was updated.
	ModifiedByUser                       *string                       `json:"modifiedByUser,omitempty" form:"modifiedByUser,omitempty"`                                             //Specifies the last Cohesity user who updated this Job.
	Name                                 string                        `json:"name" form:"name"`                                                                                     //Specifies the name of the Protection Job.
	ParentSourceId                       *int64                        `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"`                                             //Specifies the id of the registered Protection Source that is the
	PerformSourceSideDedup               *bool                         `json:"performSourceSideDedup,omitempty" form:"performSourceSideDedup,omitempty"`                             //Specifies whether source side dedupe should be performed or not.
	PolicyAppliedTimeMsecs               *int64                        `json:"policyAppliedTimeMsecs,omitempty" form:"policyAppliedTimeMsecs,omitempty"`                             //Specifies the epoch time (in milliseconds) when the
	PolicyId                             string                        `json:"policyId" form:"policyId"`                                                                             //Specifies the unique id of the Protection Policy associated with
	PostBackupScript                     *BackupScript                 `json:"postBackupScript,omitempty" form:"postBackupScript,omitempty"`                                         //Specifies the script associated with the backup job. This field must be
	PreBackupScript                      *BackupScript                 `json:"preBackupScript,omitempty" form:"preBackupScript,omitempty"`                                           //Specifies the script associated with the backup job. This field must be
	Priority                             PriorityEnum                  `json:"priority,omitempty" form:"priority,omitempty"`                                                         //Specifies the priority of execution for a Protection Job.
	QosType                              QosTypeEnum                   `json:"qosType,omitempty" form:"qosType,omitempty"`                                                           //Specifies the QoS policy type to use for this Protection Job.
	Quiesce                              *bool                         `json:"quiesce,omitempty" form:"quiesce,omitempty"`                                                           //Indicates if the App-Consistent option is enabled for this Job.
	RemoteScript                         *RemoteJobScript              `json:"remoteScript,omitempty" form:"remoteScript,omitempty"`                                                 //For a Remote Adapter 'kPuppeteer' Job, this field specifies the
	RemoteViewName                       *string                       `json:"remoteViewName,omitempty" form:"remoteViewName,omitempty"`                                             //Specifies the remote view name to use for view overwrite.
	SourceIds                            *[]int64                      `json:"sourceIds,omitempty" form:"sourceIds,omitempty"`                                                       //Array of Protected Source Objects.
	SourceSpecialParameters              []*SourceSpecialParameter     `json:"sourceSpecialParameters,omitempty" form:"sourceSpecialParameters,omitempty"`                           //Array of Special Source Parameters.
	StartTime                            *TimeOfDay                    `json:"startTime,omitempty" form:"startTime,omitempty"`                                                       //Specifies the time of day to start the Protection Schedule.
	SummaryStats                         *ProtectionJobSummaryStats    `json:"summaryStats,omitempty" form:"summaryStats,omitempty"`                                                 //Specifies statistics about a Protection Job.
	TenantId                             *string                       `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                                         //Specifies the unique id of the tenant.
	Timezone                             *string                       `json:"timezone,omitempty" form:"timezone,omitempty"`                                                         //Specifies the timezone to use when calculating time for this
	Uid                                  *UniversalId                  `json:"uid,omitempty" form:"uid,omitempty"`                                                                   //Specifies a global Protection Job id that is unique across Cohesity
	UserSpecifiedTags                    *[]string                     `json:"userSpecifiedTags,omitempty" form:"userSpecifiedTags,omitempty"`                                       //Tags associated with the job. User can specify tags/keywords that can
	ViewBoxId                            int64                         `json:"viewBoxId" form:"viewBoxId"`                                                                           //Specifies the Storage Domain (View Box) id where this Job writes data.
	ViewName                             *string                       `json:"viewName,omitempty" form:"viewName,omitempty"`                                                         //For a Remote Adapter 'kPuppeteer' Job or a 'kView' Job, this field
	VmTagIds                             *[]int64                      `json:"vmTagIds,omitempty" form:"vmTagIds,omitempty"`                                                         //Array of Arrays of VMs Tags Ids that Specify VMs to Protect.
}

/*
 * Structure for the custom type ProtectionJobAuditTrail
 */
type ProtectionJobAuditTrail struct {
	After   *ProtectionJob `json:"after,omitempty" form:"after,omitempty"`     //Provides details about a Protection Job.
	Before  *ProtectionJob `json:"before,omitempty" form:"before,omitempty"`   //Provides details about a Protection Job.
	Changes *[]ChangeEnum  `json:"changes,omitempty" form:"changes,omitempty"` //Specifies the list of changed values in a Protection Job.
}

/*
 * Structure for the custom type ProtectionJobInfo
 */
type ProtectionJobInfo struct {
	JobId   *int64                    `json:"jobId,omitempty" form:"jobId,omitempty"`     //Specifies the id of the Protection Job.
	JobName *string                   `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job.
	Type    TypeProtectionJobInfoEnum `json:"type,omitempty" form:"type,omitempty"`       //Specifies the type of the Protection Job such as kView or kPuppeteer.
}

/*
 * Structure for the custom type ProtectionJobRequestBody
 */
type ProtectionJobRequestBody struct {
	AbortInBlackoutPeriod                *bool                                   `json:"abortInBlackoutPeriod,omitempty" form:"abortInBlackoutPeriod,omitempty"`                               //If true, the Cohesity Cluster aborts any currently executing Job Runs
	AlertingConfig                       *AlertingConfig                         `json:"alertingConfig,omitempty" form:"alertingConfig,omitempty"`                                             //Specifies optional settings for alerting.
	AlertingPolicy                       *[]AlertingPolicyEnum                   `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"`                                             //Array of Job Events.
	CloudParameters                      *CloudParameters                        `json:"cloudParameters,omitempty" form:"cloudParameters,omitempty"`                                           //Specifies Cloud parameters that are applicable to all Protection
	ContinueOnQuiesceFailure             *bool                                   `json:"continueOnQuiesceFailure,omitempty" form:"continueOnQuiesceFailure,omitempty"`                         //Whether to continue backing up on quiesce failure.
	CreateRemoteView                     *bool                                   `json:"createRemoteView,omitempty" form:"createRemoteView,omitempty"`                                         //Specifies whether to create a remote view name to use for view overwrite.
	DataMigrationPolicy                  *DataMigrationPolicy                    `json:"dataMigrationPolicy,omitempty" form:"dataMigrationPolicy,omitempty"`                                   //Specifies settings for data migration in NAS environment. This also
	DedupDisabledSourceIds               *[]int64                                `json:"dedupDisabledSourceIds,omitempty" form:"dedupDisabledSourceIds,omitempty"`                             //List of source ids for which source side dedup is disabled from the backup
	Description                          *string                                 `json:"description,omitempty" form:"description,omitempty"`                                                   //Specifies a text description about the Protection Job.
	EndTimeUsecs                         *int64                                  `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                                 //Specifies the epoch time (in microseconds) after which the Protection Job
	Environment                          EnvironmentProtectionJobRequestBodyEnum `json:"environment,omitempty" form:"environment,omitempty"`                                                   //Specifies the environment type (such as kVMware or kSQL)
	EnvironmentParameters                *EnvironmentTypeJobParameters           `json:"environmentParameters,omitempty" form:"environmentParameters,omitempty"`                               //Specifies additional parameters that are common to all Protection
	ExcludeSourceIds                     *[]int64                                `json:"excludeSourceIds,omitempty" form:"excludeSourceIds,omitempty"`                                         //Array of Excluded Source Objects.
	ExcludeVmTagIds                      *[]int64                                `json:"excludeVmTagIds,omitempty" form:"excludeVmTagIds,omitempty"`                                           //Array of Arrays of VM Tag Ids that Specify VMs to Exclude.
	FullProtectionSlaTimeMins            *int64                                  `json:"fullProtectionSlaTimeMins,omitempty" form:"fullProtectionSlaTimeMins,omitempty"`                       //If specified, this setting is number of minutes that a Job Run
	FullProtectionStartTime              *TimeOfDay                              `json:"fullProtectionStartTime,omitempty" form:"fullProtectionStartTime,omitempty"`                           //Specifies the time of day to start the Full Protection Schedule.
	IncrementalProtectionSlaTimeMins     *int64                                  `json:"incrementalProtectionSlaTimeMins,omitempty" form:"incrementalProtectionSlaTimeMins,omitempty"`         //If specified, this setting is number of minutes that a Job Run
	IncrementalProtectionStartTime       *TimeOfDay                              `json:"incrementalProtectionStartTime,omitempty" form:"incrementalProtectionStartTime,omitempty"`             //Specifies the time of day to start the CBT-based Protection Schedule.
	IndexingPolicy                       *IndexingPolicy                         `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"`                                             //Specifies settings for indexing files found in an Object
	LeverageStorageSnapshots             *bool                                   `json:"leverageStorageSnapshots,omitempty" form:"leverageStorageSnapshots,omitempty"`                         //Specifies whether to leverage the storage array based snapshots for this
	LeverageStorageSnapshotsForHyperflex *bool                                   `json:"leverageStorageSnapshotsForHyperflex,omitempty" form:"leverageStorageSnapshotsForHyperflex,omitempty"` //Specifies whether to leverage Hyperflex as the storage snapshot array
	Name                                 string                                  `json:"name" form:"name"`                                                                                     //Specifies the name of the Protection Job.
	ParentSourceId                       *int64                                  `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"`                                             //Specifies the id of the registered Protection Source that is the
	PerformSourceSideDedup               *bool                                   `json:"performSourceSideDedup,omitempty" form:"performSourceSideDedup,omitempty"`                             //Specifies whether source side dedupe should be performed or not.
	PolicyId                             string                                  `json:"policyId" form:"policyId"`                                                                             //Specifies the unique id of the Protection Policy associated with
	PostBackupScript                     *BackupScript                           `json:"postBackupScript,omitempty" form:"postBackupScript,omitempty"`                                         //Specifies the script associated with the backup job. This field must be
	PreBackupScript                      *BackupScript                           `json:"preBackupScript,omitempty" form:"preBackupScript,omitempty"`                                           //Specifies the script associated with the backup job. This field must be
	Priority                             PriorityEnum                            `json:"priority,omitempty" form:"priority,omitempty"`                                                         //Specifies the priority of execution for a Protection Job.
	QosType                              QosTypeEnum                             `json:"qosType,omitempty" form:"qosType,omitempty"`                                                           //Specifies the QoS policy type to use for this Protection Job.
	Quiesce                              *bool                                   `json:"quiesce,omitempty" form:"quiesce,omitempty"`                                                           //Indicates if the App-Consistent option is enabled for this Job.
	RemoteScript                         *RemoteJobScript                        `json:"remoteScript,omitempty" form:"remoteScript,omitempty"`                                                 //For a Remote Adapter 'kPuppeteer' Job, this field specifies the
	RemoteViewName                       *string                                 `json:"remoteViewName,omitempty" form:"remoteViewName,omitempty"`                                             //Specifies the remote view name to use for view overwrite.
	SourceIds                            *[]int64                                `json:"sourceIds,omitempty" form:"sourceIds,omitempty"`                                                       //Array of Protected Source Objects.
	SourceSpecialParameters              []*SourceSpecialParameter               `json:"sourceSpecialParameters,omitempty" form:"sourceSpecialParameters,omitempty"`                           //Array of Special Source Parameters.
	StartTime                            *TimeOfDay                              `json:"startTime,omitempty" form:"startTime,omitempty"`                                                       //Specifies the time of day to start the Protection Schedule.
	Timezone                             *string                                 `json:"timezone,omitempty" form:"timezone,omitempty"`                                                         //Specifies the timezone to use when calculating time for this
	UserSpecifiedTags                    *[]string                               `json:"userSpecifiedTags,omitempty" form:"userSpecifiedTags,omitempty"`                                       //Tags associated with the job. User can specify tags/keywords that can
	ViewBoxId                            int64                                   `json:"viewBoxId" form:"viewBoxId"`                                                                           //Specifies the Storage Domain (View Box) id where this Job writes data.
	ViewName                             *string                                 `json:"viewName,omitempty" form:"viewName,omitempty"`                                                         //For a Remote Adapter 'kPuppeteer' Job or a 'kView' Job, this field
	VmTagIds                             *[]int64                                `json:"vmTagIds,omitempty" form:"vmTagIds,omitempty"`                                                         //Array of Arrays of VMs Tags Ids that Specify VMs to Protect.
}

/*
 * Structure for the custom type ProtectionJobRunStats
 */
type ProtectionJobRunStats struct {
	AdmittedTimeUsecs            *int64 `json:"admittedTimeUsecs,omitempty" form:"admittedTimeUsecs,omitempty"`                       //Specifies the time the task was unqueued from the queue to start running.
	EndTimeUsecs                 *int64 `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                 //Specifies the end time of the Protection Run. The end time
	NumAppInstances              *int64 `json:"numAppInstances,omitempty" form:"numAppInstances,omitempty"`                           //Specifies the number of application instances backed up by this Run.
	NumCanceledTasks             *int64 `json:"numCanceledTasks,omitempty" form:"numCanceledTasks,omitempty"`                         //Specifies the number of backup tasks that were canceled.
	NumFailedAppObjects          *int64 `json:"numFailedAppObjects,omitempty" form:"numFailedAppObjects,omitempty"`                   //Specifies the number of application objects that were cancelled in this
	NumFailedTasks               *int64 `json:"numFailedTasks,omitempty" form:"numFailedTasks,omitempty"`                             //Specifies the number of backup tasks that failed.
	NumSuccessfulAppObjects      *int64 `json:"numSuccessfulAppObjects,omitempty" form:"numSuccessfulAppObjects,omitempty"`           //Specifies the number of application objects successfully backed up by this
	NumSuccessfulTasks           *int64 `json:"numSuccessfulTasks,omitempty" form:"numSuccessfulTasks,omitempty"`                     //Specifies the number of backup tasks that completed successfully.
	StartTimeUsecs               *int64 `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                             //Specifies the start time of the Protection Run. The start time
	TimeTakenUsecs               *int64 `json:"timeTakenUsecs,omitempty" form:"timeTakenUsecs,omitempty"`                             //Specifies the actual execution time for the protection run to complete
	TotalBytesReadFromSource     *int64 `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"`         //Specifies the total amount of data read from the source (so far).
	TotalBytesToReadFromSource   *int64 `json:"totalBytesToReadFromSource,omitempty" form:"totalBytesToReadFromSource,omitempty"`     //Specifies the total amount of data expected to be read from the
	TotalLogicalBackupSizeBytes  *int64 `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"`   //Specifies the size of the source object (such as a VM) protected by
	TotalPhysicalBackupSizeBytes *int64 `json:"totalPhysicalBackupSizeBytes,omitempty" form:"totalPhysicalBackupSizeBytes,omitempty"` //Specifies the total amount of physical space used on the Cohesity
	TotalSourceSizeBytes         *int64 `json:"totalSourceSizeBytes,omitempty" form:"totalSourceSizeBytes,omitempty"`                 //Specifies the size of the source object (such as a VM) protected by
}

/*
 * Structure for the custom type ProtectionJobSummary
 */
type ProtectionJobSummary struct {
	ClusterId                  *int64  `json:"clusterId,omitempty" form:"clusterId,omitempty"`                                   //Specifies the id of the cluster on which object is protected.
	ClusterIncarnationId       *int64  `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"`             //Specifies the incarnation id of the cluster on which object is protected.
	IsRpoJob                   *bool   `json:"isRpoJob,omitempty" form:"isRpoJob,omitempty"`                                     //Specifies if the Protection Job is created by an RPO policy.
	JobId                      *int64  `json:"jobId,omitempty" form:"jobId,omitempty"`                                           //Specifies the id of the Protection Job.
	JobName                    *string `json:"jobName,omitempty" form:"jobName,omitempty"`                                       //Specifies the name of the Protection Job.
	LastProtectionJobRunStatus *int64  `json:"lastProtectionJobRunStatus,omitempty" form:"lastProtectionJobRunStatus,omitempty"` //Specifies the last job run status.
	PolicyId                   *string `json:"policyId,omitempty" form:"policyId,omitempty"`                                     //Specifies the id of the policy that is used by a Protection Job.
	PolicyName                 *string `json:"policyName,omitempty" form:"policyName,omitempty"`                                 //Specifies the name of the policy that is used by a Protection Job.
}

/*
 * Structure for the custom type ProtectionJobSummaryForPolicies
 */
type ProtectionJobSummaryForPolicies struct {
	BackupRun     *BackupRun     `json:"backupRun,omitempty" form:"backupRun,omitempty"`         //Specifies details about the Backup task for a Job Run.
	CopyRuns      []*CopyRun     `json:"copyRuns,omitempty" form:"copyRuns,omitempty"`           //Specifies details about the Copy tasks of the Job Run.
	ProtectionJob *ProtectionJob `json:"protectionJob,omitempty" form:"protectionJob,omitempty"` //Provides details about a Protection Job.
}

/*
 * Structure for the custom type ProtectionJobSummaryStats
 */
type ProtectionJobSummaryStats struct {
	AverageRunTimeUsecs          *int64 `json:"averageRunTimeUsecs,omitempty" form:"averageRunTimeUsecs,omitempty"`                   //Specifies the average run time of all successful Protection Runs.
	FastestRunTimeUsecs          *int64 `json:"fastestRunTimeUsecs,omitempty" form:"fastestRunTimeUsecs,omitempty"`                   //Specifies the time taken for a fastest successful Protection Run so far.
	NumCanceledRuns              *int64 `json:"numCanceledRuns,omitempty" form:"numCanceledRuns,omitempty"`                           //Specifies the number of runs that were canceled.
	NumFailedRuns                *int64 `json:"numFailedRuns,omitempty" form:"numFailedRuns,omitempty"`                               //Specifies the number of runs that failed to finish.
	NumSlaViolations             *int64 `json:"numSlaViolations,omitempty" form:"numSlaViolations,omitempty"`                         //Specifies the number of runs having SLA violations.
	NumSuccessfulRuns            *int64 `json:"numSuccessfulRuns,omitempty" form:"numSuccessfulRuns,omitempty"`                       //Specifies the number of runs that finished successfully.
	SlowestRunTimeUsecs          *int64 `json:"slowestRunTimeUsecs,omitempty" form:"slowestRunTimeUsecs,omitempty"`                   //Specifies the time taken for a slowest successful Protection Run so far.
	TotalBytesReadFromSource     *int64 `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"`         //Specifies the total amount of data read from the source (so far).
	TotalLogicalBackupSizeBytes  *int64 `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"`   //Specifies the size of the source object (such as a VM) protected by
	TotalPhysicalBackupSizeBytes *int64 `json:"totalPhysicalBackupSizeBytes,omitempty" form:"totalPhysicalBackupSizeBytes,omitempty"` //Specifies the total amount of physical space used on the Cohesity
}

/*
 * Structure for the custom type ProtectionObjectSummary
 */
type ProtectionObjectSummary struct {
	LatestArchivalSnapshotTimeUsecs    *int64                   `json:"latestArchivalSnapshotTimeUsecs,omitempty" form:"latestArchivalSnapshotTimeUsecs,omitempty"`       //Specifies the Unix epoch Timestamp (in microseconds) of the latest
	LatestLocalSnapshotTimeUsecs       *int64                   `json:"latestLocalSnapshotTimeUsecs,omitempty" form:"latestLocalSnapshotTimeUsecs,omitempty"`             //Specifies the Unix epoch Timestamp (in microseconds) of the latest
	LatestReplicationSnapshotTimeUsecs *int64                   `json:"latestReplicationSnapshotTimeUsecs,omitempty" form:"latestReplicationSnapshotTimeUsecs,omitempty"` //Specifies the Unix epoch Timestamp (in microseconds) of the latest
	ParentProtectionSource             *ProtectionSource        `json:"parentProtectionSource,omitempty" form:"parentProtectionSource,omitempty"`                         //Specifies a generic structure that represents a node
	ProtectionJobs                     []*ProtectionRunInstance `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"`                                         //Returns the list of Protection Jobs with summary Information.
	ProtectionSource                   *ProtectionSource        `json:"protectionSource,omitempty" form:"protectionSource,omitempty"`                                     //Specifies the leaf Protection Source Object such as a VM.
	RpoPolicies                        []*ProtectionPolicy      `json:"rpoPolicies,omitempty" form:"rpoPolicies,omitempty"`                                               //Specifies the id of the RPO policy protecting this object.
	TotalArchivalSnapshots             *int64                   `json:"totalArchivalSnapshots,omitempty" form:"totalArchivalSnapshots,omitempty"`                         //Specifies the total number of Archival Snapshots.
	TotalLocalSnapshots                *int64                   `json:"totalLocalSnapshots,omitempty" form:"totalLocalSnapshots,omitempty"`                               //Specifies the total number of Local Snapshots.
	TotalReplicationSnapshots          *int64                   `json:"totalReplicationSnapshots,omitempty" form:"totalReplicationSnapshots,omitempty"`                   //Specifies the total number of Replication Snapshots.
}

/*
 * Structure for the custom type ProtectionPolicy
 */
type ProtectionPolicy struct {
	BlackoutPeriods                 []*BlackoutPeriod                     `json:"blackoutPeriods,omitempty" form:"blackoutPeriods,omitempty"`                                 //Array of Blackout Periods.
	CloudDeployPolicies             []*SnapshotCloudCopyPolicy            `json:"cloudDeployPolicies,omitempty" form:"cloudDeployPolicies,omitempty"`                         //Array of Cloud Deploy Policies.
	DaysToKeep                      *int64                                `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`                                           //Specifies how many days to retain Snapshots on the Cohesity Cluster.
	DaysToKeepLog                   *int64                                `json:"daysToKeepLog,omitempty" form:"daysToKeepLog,omitempty"`                                     //Specifies the number of days to retain log run if Log Schedule exists.
	DaysToKeepSystem                *int64                                `json:"daysToKeepSystem,omitempty" form:"daysToKeepSystem,omitempty"`                               //Specifies the number of days to retain system backups made for bare metal
	Description                     *string                               `json:"description,omitempty" form:"description,omitempty"`                                         //Description of the Protection Policy.
	ExtendedRetentionPolicies       []*ExtendedRetentionPolicy            `json:"extendedRetentionPolicies,omitempty" form:"extendedRetentionPolicies,omitempty"`             //Specifies additional retention policies that should be applied to the
	FullSchedulingPolicy            *SchedulingPolicy                     `json:"fullSchedulingPolicy,omitempty" form:"fullSchedulingPolicy,omitempty"`                       //Specifies the Full (no CBT) backup schedule of a Protection Job and
	Id                              *string                               `json:"id,omitempty" form:"id,omitempty"`                                                           //Specifies a unique Policy id assigned by the Cohesity Cluster.
	IncrementalSchedulingPolicy     *SchedulingPolicy                     `json:"incrementalSchedulingPolicy,omitempty" form:"incrementalSchedulingPolicy,omitempty"`         //Specifies the CBT-based backup schedule of a Protection Job and
	LastModifiedTimeMsecs           *int64                                `json:"lastModifiedTimeMsecs,omitempty" form:"lastModifiedTimeMsecs,omitempty"`                     //Specifies the epoch time (in milliseconds) when the Protection Policy
	LogSchedulingPolicy             *SchedulingPolicy                     `json:"logSchedulingPolicy,omitempty" form:"logSchedulingPolicy,omitempty"`                         //Specifies settings that define a backup schedule for a Protection Job.
	Name                            *string                               `json:"name,omitempty" form:"name,omitempty"`                                                       //Specifies the name of the Protection Policy.
	NumLinkedPolicies               *int64                                `json:"numLinkedPolicies,omitempty" form:"numLinkedPolicies,omitempty"`                             //Species the number of policies linked to a global policy.
	Retries                         *int64                                `json:"retries,omitempty" form:"retries,omitempty"`                                                 //Specifies the number of times to retry capturing Snapshots before
	RetryIntervalMins               *int64                                `json:"retryIntervalMins,omitempty" form:"retryIntervalMins,omitempty"`                             //Specifies the number of minutes before retrying a failed Protection Job.
	RpoPolicySettings               *RpoPolicySettings                    `json:"rpoPolicySettings,omitempty" form:"rpoPolicySettings,omitempty"`                             //Specifies all the additional settings that are applicable only
	SkipIntervalMins                *int64                                `json:"skipIntervalMins,omitempty" form:"skipIntervalMins,omitempty"`                               //Specifies the period of time before skipping the execution of new
	SnapshotArchivalCopyPolicies    []*SnapshotArchivalCopyPolicy         `json:"snapshotArchivalCopyPolicies,omitempty" form:"snapshotArchivalCopyPolicies,omitempty"`       //Array of External Targets.
	SnapshotReplicationCopyPolicies []*SnapshotReplicationCopyPolicy      `json:"snapshotReplicationCopyPolicies,omitempty" form:"snapshotReplicationCopyPolicies,omitempty"` //Array of Remote Clusters.
	SystemSchedulingPolicy          *SchedulingPolicy                     `json:"systemSchedulingPolicy,omitempty" form:"systemSchedulingPolicy,omitempty"`                   //Specifies settings that define a backup schedule for a Protection Job.
	TenantIds                       *[]string                             `json:"tenantIds,omitempty" form:"tenantIds,omitempty"`                                             //Specifies which organizations have been assigned this policy.
	Type                            TypeProtectionPolicyEnum              `json:"type,omitempty" form:"type,omitempty"`                                                       //Specifies the type of the protection policy.
	WormRetentionType               WormRetentionTypeProtectionPolicyEnum `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"`                             //Specifies WORM retention type for the snapshots. When a WORM retention
}

/*
 * Structure for the custom type ProtectionPolicyRequest
 */
type ProtectionPolicyRequest struct {
	BlackoutPeriods                 []*BlackoutPeriod                            `json:"blackoutPeriods,omitempty" form:"blackoutPeriods,omitempty"`                                 //Array of Blackout Periods.
	CloudDeployPolicies             []*SnapshotCloudCopyPolicy                   `json:"cloudDeployPolicies,omitempty" form:"cloudDeployPolicies,omitempty"`                         //Array of Cloud Deploy Policies.
	DaysToKeep                      *int64                                       `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`                                           //Specifies how many days to retain Snapshots on the Cohesity Cluster.
	DaysToKeepLog                   *int64                                       `json:"daysToKeepLog,omitempty" form:"daysToKeepLog,omitempty"`                                     //Specifies the number of days to retain log run if Log Schedule exists.
	DaysToKeepSystem                *int64                                       `json:"daysToKeepSystem,omitempty" form:"daysToKeepSystem,omitempty"`                               //Specifies the number of days to retain system backups made for bare metal
	Description                     *string                                      `json:"description,omitempty" form:"description,omitempty"`                                         //Description of the Protection Policy.
	ExtendedRetentionPolicies       []*ExtendedRetentionPolicy                   `json:"extendedRetentionPolicies,omitempty" form:"extendedRetentionPolicies,omitempty"`             //Specifies additional retention policies that should be applied to the
	FullSchedulingPolicy            *SchedulingPolicy                            `json:"fullSchedulingPolicy,omitempty" form:"fullSchedulingPolicy,omitempty"`                       //Specifies the Full (no CBT) backup schedule of a Protection Job and
	IncrementalSchedulingPolicy     *SchedulingPolicy                            `json:"incrementalSchedulingPolicy,omitempty" form:"incrementalSchedulingPolicy,omitempty"`         //Specifies the CBT-based backup schedule of a Protection Job and
	LogSchedulingPolicy             *SchedulingPolicy                            `json:"logSchedulingPolicy,omitempty" form:"logSchedulingPolicy,omitempty"`                         //Specifies settings that define a backup schedule for a Protection Job.
	Name                            *string                                      `json:"name,omitempty" form:"name,omitempty"`                                                       //Specifies the name of the Protection Policy.
	NumLinkedPolicies               *int64                                       `json:"numLinkedPolicies,omitempty" form:"numLinkedPolicies,omitempty"`                             //Species the number of policies linked to a global policy.
	Retries                         *int64                                       `json:"retries,omitempty" form:"retries,omitempty"`                                                 //Specifies the number of times to retry capturing Snapshots before
	RetryIntervalMins               *int64                                       `json:"retryIntervalMins,omitempty" form:"retryIntervalMins,omitempty"`                             //Specifies the number of minutes before retrying a failed Protection Job.
	RpoPolicySettings               *RpoPolicySettings                           `json:"rpoPolicySettings,omitempty" form:"rpoPolicySettings,omitempty"`                             //Specifies all the additional settings that are applicable only
	SkipIntervalMins                *int64                                       `json:"skipIntervalMins,omitempty" form:"skipIntervalMins,omitempty"`                               //Specifies the period of time before skipping the execution of new
	SnapshotArchivalCopyPolicies    []*SnapshotArchivalCopyPolicy                `json:"snapshotArchivalCopyPolicies,omitempty" form:"snapshotArchivalCopyPolicies,omitempty"`       //Array of External Targets.
	SnapshotReplicationCopyPolicies []*SnapshotReplicationCopyPolicy             `json:"snapshotReplicationCopyPolicies,omitempty" form:"snapshotReplicationCopyPolicies,omitempty"` //Array of Remote Clusters.
	SystemSchedulingPolicy          *SchedulingPolicy                            `json:"systemSchedulingPolicy,omitempty" form:"systemSchedulingPolicy,omitempty"`                   //Specifies settings that define a backup schedule for a Protection Job.
	Type                            TypeProtectionPolicyRequestEnum              `json:"type,omitempty" form:"type,omitempty"`                                                       //Specifies the type of the protection policy.
	WormRetentionType               WormRetentionTypeProtectionPolicyRequestEnum `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"`                             //Specifies WORM retention type for the snapshots. When a WORM retention
}

/*
 * Structure for the custom type ProtectionPolicySummary
 */
type ProtectionPolicySummary struct {
	LastProtectionRunSummary *LastProtectionRunSummary          `json:"lastProtectionRunSummary,omitempty" form:"lastProtectionRunSummary,omitempty"` //LastProtectionRunsSummary is the summary of the last Protection Run for the
	PaginationCookie         *string                            `json:"paginationCookie,omitempty" form:"paginationCookie,omitempty"`                 //If there are more results to display, use this value to get
	ProtectedSourcesSummary  []*ProtectedSourceSummary          `json:"protectedSourcesSummary,omitempty" form:"protectedSourcesSummary,omitempty"`   //Specifies the list of Protection Sources which are protected under the
	ProtectionJobsSummary    []*ProtectionJobSummaryForPolicies `json:"protectionJobsSummary,omitempty" form:"protectionJobsSummary,omitempty"`       //Specifies the list of Protection Jobs associated with the given
	ProtectionPolicy         *ProtectionPolicy                  `json:"protectionPolicy,omitempty" form:"protectionPolicy,omitempty"`                 //TODO: Write general description for this field
	ProtectionRunsSummary    *ProtectionRunsSummary             `json:"protectionRunsSummary,omitempty" form:"protectionRunsSummary,omitempty"`       //ProtectionRunsSummary is the summary of the all the Protection Runs for the
}

/*
 * Structure for the custom type ProtectionRunErrors
 */
type ProtectionRunErrors struct {
	Errors           []*RequestError `json:"errors,omitempty" form:"errors,omitempty"`                     //Specifies the list of errors encountered by a task during a protection
	FileNames        *[]string       `json:"fileNames,omitempty" form:"fileNames,omitempty"`               //Specifies the list of filenames with errors encountered by a task during a
	PaginationCookie *string         `json:"paginationCookie,omitempty" form:"paginationCookie,omitempty"` //Specifies the cookie for next set of results.
}

/*
 * Structure for the custom type ProtectionRunInstance
 */
type ProtectionRunInstance struct {
	BackupRun *BackupRun   `json:"backupRun,omitempty" form:"backupRun,omitempty"` //Specifies details about the Backup task for a Job Run.
	CopyRun   []*CopyRun   `json:"copyRun,omitempty" form:"copyRun,omitempty"`     //Array of Copy Run Tasks.
	JobId     *int64       `json:"jobId,omitempty" form:"jobId,omitempty"`         //Specifies the id of the Protection Job that was run.
	JobName   *string      `json:"jobName,omitempty" form:"jobName,omitempty"`     //Specifies the name of the Protection Job name that was run.
	JobUid    *UniversalId `json:"jobUid,omitempty" form:"jobUid,omitempty"`       //Specifies the globally unique id of the Protection Job that was run.
	ViewBoxId *int64       `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the Storage Domain (View Box) to store the backed up data.
}

/*
 * Structure for the custom type ProtectionRunResponse
 */
type ProtectionRunResponse struct {
	ArchivalRuns    []*LatestProtectionJobRunInfo `json:"archivalRuns,omitempty" form:"archivalRuns,omitempty"`       //Specifies the list of archival job information.
	BackupRuns      []*LatestProtectionJobRunInfo `json:"backupRuns,omitempty" form:"backupRuns,omitempty"`           //Specifies the list of local backup job information.
	ReplicationRuns []*LatestProtectionJobRunInfo `json:"replicationRuns,omitempty" form:"replicationRuns,omitempty"` //Specifies the list of replication job information.
}

/*
 * Structure for the custom type ProtectionRunsStats
 */
type ProtectionRunsStats struct {
	NumArchivalRuns    *int64 `json:"numArchivalRuns,omitempty" form:"numArchivalRuns,omitempty"`       //Specifies the count of archival Runs.
	NumBackupRuns      *int64 `json:"numBackupRuns,omitempty" form:"numBackupRuns,omitempty"`           //Specifies the count of backup Runs.
	NumReplicationRuns *int64 `json:"numReplicationRuns,omitempty" form:"numReplicationRuns,omitempty"` //Specifies the count of replication Runs.
}

/*
 * Structure for the custom type ProtectionRunsSummary
 */
type ProtectionRunsSummary struct {
	NumberOfArchivalRuns              *int64 `json:"numberOfArchivalRuns,omitempty" form:"numberOfArchivalRuns,omitempty"`                           //Specifies the total number of Archival Runs using the current
	NumberOfProtectionRuns            *int64 `json:"numberOfProtectionRuns,omitempty" form:"numberOfProtectionRuns,omitempty"`                       //Specifies the total number of Protection Runs by the given Protection
	NumberOfReplicationRuns           *int64 `json:"numberOfReplicationRuns,omitempty" form:"numberOfReplicationRuns,omitempty"`                     //Specifies the total number of Replication Runs using the current
	NumberOfSuccessfulArchivalRuns    *int64 `json:"numberOfSuccessfulArchivalRuns,omitempty" form:"numberOfSuccessfulArchivalRuns,omitempty"`       //Specifies the number of total successful Archival Runs using the
	NumberOfSuccessfulProtectionRuns  *int64 `json:"numberOfSuccessfulProtectionRuns,omitempty" form:"numberOfSuccessfulProtectionRuns,omitempty"`   //Specifies the number of successful Protection Runs using the current
	NumberOfSuccessfulReplicationRuns *int64 `json:"numberOfSuccessfulReplicationRuns,omitempty" form:"numberOfSuccessfulReplicationRuns,omitempty"` //Specifies the number of total successful Replication Runs using the
}

/*
 * Structure for the custom type ProtectionSource
 */
type ProtectionSource struct {
	AcropolisProtectionSource  *AcropolisProtectionSource  `json:"acropolisProtectionSource,omitempty" form:"acropolisProtectionSource,omitempty"`   //Specifies details about an Acropolis Protection Source
	AdProtectionSource         *AdProtectionSource         `json:"adProtectionSource,omitempty" form:"adProtectionSource,omitempty"`                 //Specifies details about an AD Protection Source
	AwsProtectionSource        *AwsProtectionSource        `json:"awsProtectionSource,omitempty" form:"awsProtectionSource,omitempty"`               //Specifies details about an AWS Protection Source
	AzureProtectionSource      *AzureProtectionSource      `json:"azureProtectionSource,omitempty" form:"azureProtectionSource,omitempty"`           //Specifies details about an Azure Protection Source
	Environment                EnvironmentEnum             `json:"environment,omitempty" form:"environment,omitempty"`                               //Specifies the environment (such as 'kVMware' or 'kSQL') where the
	FlashBladeProtectionSource *FlashBladeProtectionSource `json:"flashBladeProtectionSource,omitempty" form:"flashBladeProtectionSource,omitempty"` //Specifies details about a Pure Storage FlashBlade Protection Source
	GcpProtectionSource        *GcpProtectionSource        `json:"gcpProtectionSource,omitempty" form:"gcpProtectionSource,omitempty"`               //Specifies details about an GCP Protection Source
	GpfsProtectionSource       *GpfsProtectionSource       `json:"gpfsProtectionSource,omitempty" form:"gpfsProtectionSource,omitempty"`             //Specifies details about an GPFS Protection Source
	HyperFlexProtectionSource  *HyperFlexProtectionSource  `json:"hyperFlexProtectionSource,omitempty" form:"hyperFlexProtectionSource,omitempty"`   //Specifies details about a HyperFlex Storage Snapshot source
	HypervProtectionSource     *HypervProtectionSource     `json:"hypervProtectionSource,omitempty" form:"hypervProtectionSource,omitempty"`         //Specifies details about a HyperV Protection Source
	Id                         *int64                      `json:"id,omitempty" form:"id,omitempty"`                                                 //Specifies an id of the Protection Source.
	IsilonProtectionSource     *IsilonProtectionSource     `json:"isilonProtectionSource,omitempty" form:"isilonProtectionSource,omitempty"`         //Specifies details about an Isilon OneFs Protection Source
	KubernetesProtectionSource *KubernetesProtectionSource `json:"kubernetesProtectionSource,omitempty" form:"kubernetesProtectionSource,omitempty"` //Specifies details about a Kubernetes Protection Source
	KvmProtectionSource        *KvmProtectionSource        `json:"kvmProtectionSource,omitempty" form:"kvmProtectionSource,omitempty"`               //Specifies details about a KVM Protection Source
	Name                       *string                     `json:"name,omitempty" form:"name,omitempty"`                                             //Specifies a name of the Protection Source.
	NasProtectionSource        *NasProtectionSource        `json:"nasProtectionSource,omitempty" form:"nasProtectionSource,omitempty"`               //Specifies details about a Generic NAS Protection Source
	NetappProtectionSource     *NetappProtectionSource     `json:"netappProtectionSource,omitempty" form:"netappProtectionSource,omitempty"`         //Specifies details about a NetApp Protection Source
	Office365ProtectionSource  *Office365ProtectionSource  `json:"office365ProtectionSource,omitempty" form:"office365ProtectionSource,omitempty"`   //Specifies details about an Office 365 Protection Source
	OracleProtectionSource     *OracleProtectionSource     `json:"oracleProtectionSource,omitempty" form:"oracleProtectionSource,omitempty"`         //Specifies details about an Oracle Protection Source
	ParentId                   *int64                      `json:"parentId,omitempty" form:"parentId,omitempty"`                                     //Specifies an id of the parent of the Protection Source.
	PhysicalProtectionSource   *PhysicalProtectionSource   `json:"physicalProtectionSource,omitempty" form:"physicalProtectionSource,omitempty"`     //Specifies details about a Physical Protection Source
	PureProtectionSource       *PureProtectionSource       `json:"pureProtectionSource,omitempty" form:"pureProtectionSource,omitempty"`             //Specifies details about a Pure Protection Source
	SqlProtectionSource        *SqlProtectionSource        `json:"sqlProtectionSource,omitempty" form:"sqlProtectionSource,omitempty"`               //Specifies details about a SQL Protection Source
	ViewProtectionSource       *ViewProtectionSource       `json:"viewProtectionSource,omitempty" form:"viewProtectionSource,omitempty"`             //Specifies details about a View Protection Source
	VmwareProtectionSource     *VmwareProtectionSource     `json:"vmWareProtectionSource,omitempty" form:"vmWareProtectionSource,omitempty"`         //Specifies details about a VMware Protection Source
}

/*
 * Structure for the custom type ProtectionSourceNode
 */
type ProtectionSourceNode struct {
	ApplicationNodes          *[]interface{}               `json:"applicationNodes,omitempty" form:"applicationNodes,omitempty"`                   //Array of Child Subtrees.
	EntityPermissionInfo      *EntityPermissionInformation `json:"entityPermissionInfo,omitempty" form:"entityPermissionInfo,omitempty"`           //Specifies the permission information of entities.
	LogicalSize               *int64                       `json:"logicalSize,omitempty" form:"logicalSize,omitempty"`                             //Specifies the logical size of the data in bytes for the Object
	Nodes                     *[]interface{}               `json:"nodes,omitempty" form:"nodes,omitempty"`                                         //Array of Child Nodes.
	ProtectedSourcesSummary   []*AggregatedSubtreeInfo     `json:"protectedSourcesSummary,omitempty" form:"protectedSourcesSummary,omitempty"`     //Array of Protected Objects.
	ProtectionSource          *ProtectionSource            `json:"protectionSource,omitempty" form:"protectionSource,omitempty"`                   //Specifies the Protection Source for the current node.
	RegistrationInfo          *RegisteredSourceInfo        `json:"registrationInfo,omitempty" form:"registrationInfo,omitempty"`                   //Specifies registration information for a root node in a Protection
	UnprotectedSourcesSummary []*AggregatedSubtreeInfo     `json:"unprotectedSourcesSummary,omitempty" form:"unprotectedSourcesSummary,omitempty"` //Array of Unprotected Sources.
}

/*
 * Structure for the custom type ProtectionSourceResponse
 */
type ProtectionSourceResponse struct {
	Jobs                    []*ProtectionJobSummary `json:"jobs,omitempty" form:"jobs,omitempty"`                                       //Specifies the list of Protection Jobs that protect the object.
	LogicalSizeInBytes      *int64                  `json:"logicalSizeInBytes,omitempty" form:"logicalSizeInBytes,omitempty"`           //Specifies the logical size of Protection Source in bytes.
	ParentSource            *ProtectionSource       `json:"parentSource,omitempty" form:"parentSource,omitempty"`                       //Specifies a generic structure that represents a node
	ProtectionSourceUidList []*ProtectionSourceUid  `json:"protectionSourceUidList,omitempty" form:"protectionSourceUidList,omitempty"` //Specifies the list of universal ids of the Protection Source.
	Source                  *ProtectionSource       `json:"source,omitempty" form:"source,omitempty"`                                   //Specifies a generic structure that represents a node
	Uuid                    *string                 `json:"uuid,omitempty" form:"uuid,omitempty"`                                       //Specifies the unique id of the Protection Source.
}

/*
 * Structure for the custom type ProtectionSourceTreeInfo
 */
type ProtectionSourceTreeInfo struct {
	Applications         []*ApplicationInfo           `json:"applications,omitempty" form:"applications,omitempty"`                 //Array of applications hierarchy registered on this node.
	EntityPermissionInfo *EntityPermissionInformation `json:"entityPermissionInfo,omitempty" form:"entityPermissionInfo,omitempty"` //Specifies the permission information of entities.
	LogicalSizeBytes     *int64                       `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"`         //Specifies the logical size of the Protection Source in bytes.
	RegistrationInfo     *RegisteredSourceInfo        `json:"registrationInfo,omitempty" form:"registrationInfo,omitempty"`         //Specifies registration information for a root node in a Protection
	RootNode             *ProtectionSource            `json:"rootNode,omitempty" form:"rootNode,omitempty"`                         //Specifies the Protection Source for the root node of the Protection
	Stats                *ProtectionSummary           `json:"stats,omitempty" form:"stats,omitempty"`                               //Specifies the stats of protection for a Protection Source Tree.
	StatsByEnv           []*ProtectionSummaryByEnv    `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"`                     //Specifies the breakdown of the stats of protection by environment.
}

/*
 * Structure for the custom type ProtectionSourceUid
 */
type ProtectionSourceUid struct {
	ClusterId            *int64 `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //Specifies the id of the cluster on which object is present.
	ClusterIncarnationId *int64 `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the incarnation id of the cluster on which object is present.
	ParentSourceId       *int64 `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"`             //Specifies parent source id of an object.
	SourceId             *int64 `json:"sourceId,omitempty" form:"sourceId,omitempty"`                         //Specifies source id of an object.
}

/*
 * Structure for the custom type ProtectionStats
 */
type ProtectionStats struct {
	NumFailed  *int64 `json:"numFailed,omitempty" form:"numFailed,omitempty"`   //Number of Failed Objects.
	NumObjects *int64 `json:"numObjects,omitempty" form:"numObjects,omitempty"` //Number of Objects.
	SizeBytes  *int64 `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"`   //Size in Bytes.
}

/*
 * Structure for the custom type ProtectionSummary
 */
type ProtectionSummary struct {
	ProtectedCount   *int64 `json:"protectedCount,omitempty" form:"protectedCount,omitempty"`     //Specifies the number of objects that are protected under the given
	ProtectedSize    *int64 `json:"protectedSize,omitempty" form:"protectedSize,omitempty"`       //Specifies the total size of the protected objects under the given entity.
	UnprotectedCount *int64 `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"` //Specifies the number of objects that are not protected under the given
	UnprotectedSize  *int64 `json:"unprotectedSize,omitempty" form:"unprotectedSize,omitempty"`   //Specifies the total size of the unprotected objects under the given
}

/*
 * Structure for the custom type ProtectionSummaryByEnv
 */
type ProtectionSummaryByEnv struct {
	Environment      EnvironmentProtectionSummaryByEnvEnum `json:"environment,omitempty" form:"environment,omitempty"`           //Specifies the type of environment of the source object like kSQL etc.
	ProtectedCount   *int64                                `json:"protectedCount,omitempty" form:"protectedCount,omitempty"`     //Specifies the number of objects that are protected under the given
	ProtectedSize    *int64                                `json:"protectedSize,omitempty" form:"protectedSize,omitempty"`       //Specifies the total size of the protected objects under the given entity.
	UnprotectedCount *int64                                `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"` //Specifies the number of objects that are not protected under the given
	UnprotectedSize  *int64                                `json:"unprotectedSize,omitempty" form:"unprotectedSize,omitempty"`   //Specifies the total size of the unprotected objects under the given
}

/*
 * Structure for the custom type ProtectionTile
 */
type ProtectionTile struct {
	LastDayArchival       *ProtectionStats `json:"lastDayArchival,omitempty" form:"lastDayArchival,omitempty"`             //Protection Statistics.
	LastDayBackup         *ProtectionStats `json:"lastDayBackup,omitempty" form:"lastDayBackup,omitempty"`                 //Protection Statistics.
	LastDayReplicationIn  *ProtectionStats `json:"lastDayReplicationIn,omitempty" form:"lastDayReplicationIn,omitempty"`   //Protection Statistics.
	LastDayReplicationOut *ProtectionStats `json:"lastDayReplicationOut,omitempty" form:"lastDayReplicationOut,omitempty"` //Protection Statistics.
}

/*
 * Structure for the custom type PureEnvJobParameters
 */
type PureEnvJobParameters struct {
	MaxSnapshotsOnPrimary *int64 `json:"maxSnapshotsOnPrimary,omitempty" form:"maxSnapshotsOnPrimary,omitempty"` //Specifies how many recent snapshots of each backed up entity to retain on
}

/*
 * Structure for the custom type PureProtectionSource
 */
type PureProtectionSource struct {
	Name         *string                      `json:"name,omitempty" form:"name,omitempty"`                 //Specifies a unique name of the Protection Source
	StorageArray *PureStorageArray            `json:"storageArray,omitempty" form:"storageArray,omitempty"` //Specifies a Pure Storage Array.
	Type         TypePureProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                 //Specifies the type of managed Object in a pure Protection Source like
	Volume       *PureVolume                  `json:"volume,omitempty" form:"volume,omitempty"`             //Specifies a Pure Volume in a Pure Storage Array.
}

/*
 * Structure for the custom type PureStorageArray
 */
type PureStorageArray struct {
	Id       *string         `json:"id,omitempty" form:"id,omitempty"`             //Specifies a unique id of a Pure Storage Array.
	Ports    []*IscsiSanPort `json:"ports,omitempty" form:"ports,omitempty"`       //Specifies the SAN ports of the Pure Storage Array.
	Revision *string         `json:"revision,omitempty" form:"revision,omitempty"` //Specifies the revision of the Pure Storage Array.
	Version  *string         `json:"version,omitempty" form:"version,omitempty"`   //Specifies the version of the Pure Storage Array.
}

/*
 * Structure for the custom type PureVolume
 */
type PureVolume struct {
	CreatedTime  *string `json:"createdTime,omitempty" form:"createdTime,omitempty"`   //Specifies the created time (e.g., "2015-07-21T17:59:41Z") of the volume.
	ParentVolume *string `json:"parentVolume,omitempty" form:"parentVolume,omitempty"` //Specifies the name of the source volume, if this volume was
	SerialNumber *string `json:"serialNumber,omitempty" form:"serialNumber,omitempty"` //Specifies the serial number of the volume.
	SizeBytes    *int64  `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"`       //Specifies the provisioned size in bytes of the volume.
	UsedBytes    *int64  `json:"usedBytes,omitempty" form:"usedBytes,omitempty"`       //Specifies the total space actually used by the volume.
}

/*
 * Structure for the custom type QStarServerCredentials
 */
type QStarServerCredentials struct {
	Host                *string   `json:"host,omitempty" form:"host,omitempty"`                               //Specifies the IP address or DNS name of the server where QStar
	IntegralVolumeNames *[]string `json:"integralVolumeNames,omitempty" form:"integralVolumeNames,omitempty"` //Array of Integral Volume Names.
	Password            *string   `json:"password,omitempty" form:"password,omitempty"`                       //Specifies the password used to access the QStar host.
	Port                *int64    `json:"port,omitempty" form:"port,omitempty"`                               //Specifies the listening port where QStar WEB API service is running.
	ShareType           *string   `json:"shareType,omitempty" form:"shareType,omitempty"`                     //Specifies the sharing protocol type used by QStar to mount
	UseHttps            *bool     `json:"useHttps,omitempty" form:"useHttps,omitempty"`                       //Specifies whether to use http or https to connect to the service.
	Username            *string   `json:"username,omitempty" form:"username,omitempty"`                       //Specifies the account name used to access the QStar host.
}

/*
 * Structure for the custom type QoS
 */
type QoS struct {
	PrincipalName *string `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the QoS Policy used for the View such as
}

/*
 * Structure for the custom type QoSPolicy
 */
type QoSPolicy struct {
	AlwaysUseSsd        *bool                 `json:"alwaysUseSsd,omitempty" form:"alwaysUseSsd,omitempty"`               //Specifies whether to always write to SSD even if SeqWriteSsdPct is 0.
	Id                  *int64                `json:"id,omitempty" form:"id,omitempty"`                                   //Specifies Id of the QoS Policy.
	MinRequests         *int64                `json:"minRequests,omitempty" form:"minRequests,omitempty"`                 //Specifies minimum number of requests,  corresponding to this Policy,
	Name                *string               `json:"name,omitempty" form:"name,omitempty"`                               //Specifies Name of the Qos Policy.
	Priority            PriorityQoSPolicyEnum `json:"priority,omitempty" form:"priority,omitempty"`                       //Specifies Priority of the Qos Policy.
	RandomWriteHydraPct *int64                `json:"randomWriteHydraPct,omitempty" form:"randomWriteHydraPct,omitempty"` //Specifies percentage of a random write request belonging to this Policy
	RandomWriteSsdPct   *int64                `json:"randomWriteSsdPct,omitempty" form:"randomWriteSsdPct,omitempty"`     //Specifies percentage of a random write request belonging to this Policy
	SeqWriteHydraPct    *int64                `json:"seqWriteHydraPct,omitempty" form:"seqWriteHydraPct,omitempty"`       //Specifies percentage of a sequential write request belonging to this
	SeqWriteSsdPct      *int64                `json:"seqWriteSsdPct,omitempty" form:"seqWriteSsdPct,omitempty"`           //Specifies percentage of a sequential write request belonging to this
	Weight              *int64                `json:"weight,omitempty" form:"weight,omitempty"`                           //Specifies Weight of the QoS Policy used in QoS queue.
	WorkLoadType        *string               `json:"workLoadType,omitempty" form:"workLoadType,omitempty"`               //Specifies Workload type attribute associated with this Policy.
}

/*
 * Structure for the custom type QuotaAndUsageInView
 */
type QuotaAndUsageInView struct {
	Quota      *QuotaPolicy `json:"quota,omitempty" form:"quota,omitempty"`           //Specifies a quota limit that can be optionally applied to Views and
	UsageBytes *int64       `json:"usageBytes,omitempty" form:"usageBytes,omitempty"` //Usage in bytes of this user in this view.
	ViewId     *int64       `json:"viewId,omitempty" form:"viewId,omitempty"`         //The usage and quota policy information of this user for this view.
	ViewName   *string      `json:"viewName,omitempty" form:"viewName,omitempty"`     //View name.
}

/*
 * Structure for the custom type QuotaPolicy
 */
type QuotaPolicy struct {
	AlertLimitBytes          *int64 `json:"alertLimitBytes,omitempty" form:"alertLimitBytes,omitempty"`                   //Specifies if an alert should be triggered when the usage of this
	AlertThresholdPercentage *int64 `json:"alertThresholdPercentage,omitempty" form:"alertThresholdPercentage,omitempty"` //Supported only for user quota policy.
	HardLimitBytes           *int64 `json:"hardLimitBytes,omitempty" form:"hardLimitBytes,omitempty"`                     //Specifies an optional quota limit on the usage allowed for this
}

/*
 * Structure for the custom type RdsParams
 */
type RdsParams struct {
	AvailabilityZoneId            *int64 `json:"availabilityZoneId,omitempty" form:"availabilityZoneId,omitempty"`                       //Entity representing the availability zone to use while restoring the DB.
	DbInstanceId                  string `json:"dbInstanceId" form:"dbInstanceId"`                                                       //The DB instance identifier to use for the restored DB. This field is
	DbOptionGroupId               *int64 `json:"dbOptionGroupId,omitempty" form:"dbOptionGroupId,omitempty"`                             //Entity representing the RDS option group to use while restoring the DB.
	DbParameterGroupId            *int64 `json:"dbParameterGroupId,omitempty" form:"dbParameterGroupId,omitempty"`                       //Entity representing the RDS parameter group to use while restoring the DB.
	DbPort                        *int64 `json:"dbPort,omitempty" form:"dbPort,omitempty"`                                               //Port to use for the DB in the restored RDS instance.
	EnableAutoMinorVersionUpgrade *bool  `json:"enableAutoMinorVersionUpgrade,omitempty" form:"enableAutoMinorVersionUpgrade,omitempty"` //Whether to enable auto minor version upgrade in the restored DB.
	EnableCopyTagsToSnapshots     *bool  `json:"enableCopyTagsToSnapshots,omitempty" form:"enableCopyTagsToSnapshots,omitempty"`         //Whether to enable copying of tags to snapshots of the DB.
	EnableDbAuthentication        *bool  `json:"enableDbAuthentication,omitempty" form:"enableDbAuthentication,omitempty"`               //Whether to enable IAM authentication for the DB.
	EnablePublicAccessibility     *bool  `json:"enablePublicAccessibility,omitempty" form:"enablePublicAccessibility,omitempty"`         //Whether this DB will be publicly accessible or not.
	IsMultiAzDeployment           *bool  `json:"isMultiAzDeployment,omitempty" form:"isMultiAzDeployment,omitempty"`                     //Whether this is a multi-az deployment or not.
}

/*
 * Structure for the custom type RecoverDisksTaskStateProto
 */
type RecoverDisksTaskStateProto struct {
	RecoverVirtualDiskInfo   *RecoverVirtualDiskInfoProto `json:"recoverVirtualDiskInfo,omitempty" form:"recoverVirtualDiskInfo,omitempty"`     //Each available extension is listed below along with the location of the
	RecoverVirtualDiskParams *RecoverVirtualDiskParams    `json:"recoverVirtualDiskParams,omitempty" form:"recoverVirtualDiskParams,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type RecoverTaskRequest
 */
type RecoverTaskRequest struct {
	AcropolisParameters          *AcropolisRestoreParameters   `json:"acropolisParameters,omitempty" form:"acropolisParameters,omitempty"`                   //This field defines the Acropolis specific params for restore tasks of type
	ContinueOnError              *bool                         `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`                           //Specifies if the Restore Task should continue when some operations on some
	DeployVmsToCloud             *DeployVmsToCloud             `json:"deployVmsToCloud,omitempty" form:"deployVmsToCloud,omitempty"`                         //Specifies the details about deploying vms to specific clouds where backup
	GlacierRetrievalType         GlacierRetrievalTypeEnum      `json:"glacierRetrievalType,omitempty" form:"glacierRetrievalType,omitempty"`                 //Specifies the way data needs to be retrieved from the external target.
	HypervParameters             *HypervRestoreParameters      `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"`                         //Specifies information needed when restoring VMs in HyperV enviroment.
	MountParameters              *MountVolumesParameters       `json:"mountParameters,omitempty" form:"mountParameters,omitempty"`                           //Specifies the information required for mounting volumes.
	Name                         string                        `json:"name" form:"name"`                                                                     //Specifies the name of the Restore Task. This field must be set and
	NewParentId                  *int64                        `json:"newParentId,omitempty" form:"newParentId,omitempty"`                                   //Specify a new registered parent Protection Source. If specified
	Objects                      []*RestoreObjectDetails       `json:"objects,omitempty" form:"objects,omitempty"`                                           //Array of Objects.
	OutlookParameters            *OutlookRestoreParameters     `json:"outlookParameters,omitempty" form:"outlookParameters,omitempty"`                       //Specifies information needed for recovering Mailboxes in O365Outlook
	RestoreViewParameters        *UpdateViewParam              `json:"restoreViewParameters,omitempty" form:"restoreViewParameters,omitempty"`               //Specifies the settings that define a View.
	Type                         TypeRecoverTaskRequestEnum    `json:"type" form:"type"`                                                                     //Specifies the type of Restore Task such as 'kRecoverVMs' or
	ViewName                     *string                       `json:"viewName,omitempty" form:"viewName,omitempty"`                                         //Specifie target view into which the objects are to be cloned when doing
	VirtualDiskRestoreParameters *VirtualDiskRestoreParameters `json:"virtualDiskRestoreParameters,omitempty" form:"virtualDiskRestoreParameters,omitempty"` //Specifies the parameters to recover virtual disks of a vm.
	VlanParameters               *VlanParameters               `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"`                             //Specifies VLAN parameters for the restore operation.
	VmwareParameters             *VmwareRestoreParameters      `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"`                         //Specifies the information required for recovering or cloning VmWare VMs.
}

/*
 * Structure for the custom type RecoverVirtualDiskInfoProto
 */
type RecoverVirtualDiskInfoProto struct {
	CleanupError              *ErrorProto                    `json:"cleanupError,omitempty" form:"cleanupError,omitempty"`                           //TODO: Write general description for this field
	DataMigrationError        *ErrorProto                    `json:"dataMigrationError,omitempty" form:"dataMigrationError,omitempty"`               //TODO: Write general description for this field
	Error                     *ErrorProto                    `json:"error,omitempty" form:"error,omitempty"`                                         //TODO: Write general description for this field
	Finished                  *bool                          `json:"finished,omitempty" form:"finished,omitempty"`                                   //This will be set to true if the task is complete on the slave.
	InstantRecoveryFinished   *bool                          `json:"instantRecoveryFinished,omitempty" form:"instantRecoveryFinished,omitempty"`     //This will be set to true once the instant recovery of the virtual disk is
	MigrateTaskMoref          *MORef                         `json:"migrateTaskMoref,omitempty" form:"migrateTaskMoref,omitempty"`                   //TODO: Write general description for this field
	RestoreDisksTaskInfoProto *SetupRestoreDiskTaskInfoProto `json:"restoreDisksTaskInfoProto,omitempty" form:"restoreDisksTaskInfoProto,omitempty"` //Each available extension is listed below along with the location of the
	SlaveTaskStartTimeUsecs   *int64                         `json:"slaveTaskStartTimeUsecs,omitempty" form:"slaveTaskStartTimeUsecs,omitempty"`     //This is the timestamp at which the slave task started.
	TaskState                 *int64                         `json:"taskState,omitempty" form:"taskState,omitempty"`                                 //The state of the task.
	Type                      *int64                         `json:"type,omitempty" form:"type,omitempty"`                                           //The type of environment this recover virtual disk info pertains to.
}

/*
 * Structure for the custom type RecoverVirtualDiskParams
 */
type RecoverVirtualDiskParams struct {
	PowerOffVmBeforeRecovery *bool                                         `json:"powerOffVmBeforeRecovery,omitempty" form:"powerOffVmBeforeRecovery,omitempty"` //Whether to power-off the VM before recovering virtual disks.
	PowerOnVmAfterRecovery   *bool                                         `json:"powerOnVmAfterRecovery,omitempty" form:"powerOnVmAfterRecovery,omitempty"`     //Whether to power-on the VM after recovering virtual disks.
	TargetEntity             *EntityProto                                  `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	VirtualDiskMappings      []*RecoverVirtualDiskParamsVirtualDiskMapping `json:"virtualDiskMappings,omitempty" form:"virtualDiskMappings,omitempty"`           //TODO: Write general description for this field
}

/*
 * Structure for the custom type RecoverVirtualDiskParamsVirtualDiskMapping
 */
type RecoverVirtualDiskParamsVirtualDiskMapping struct {
	DiskToOverwrite *VirtualDiskId `json:"diskToOverwrite,omitempty" form:"diskToOverwrite,omitempty"` //This message defines the proto that can be used to identify the disks in
	SrcDisk         *VirtualDiskId `json:"srcDisk,omitempty" form:"srcDisk,omitempty"`                 //This message defines the proto that can be used to identify the disks in
	TargetLocation  *EntityProto   `json:"targetLocation,omitempty" form:"targetLocation,omitempty"`   //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RecoverVolumesParams
 */
type RecoverVolumesParams struct {
	ForceUnmountVolume *bool                          `json:"forceUnmountVolume,omitempty" form:"forceUnmountVolume,omitempty"` //Whether volume would be dismounted first during LockVolume failure
	MappingVec         []*RecoverVolumesParamsMapping `json:"mappingVec,omitempty" form:"mappingVec,omitempty"`                 //Contains the volume mapping data that defines the restore task.
	TargetEntity       *EntityProto                   `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`             //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RecoverVolumesParamsMapping
 */
type RecoverVolumesParamsMapping struct {
	DstGuid *string `json:"dstGuid,omitempty" form:"dstGuid,omitempty"` //The destination, pertains to the newly rebuilt system.
	SrcGuid *string `json:"srcGuid,omitempty" form:"srcGuid,omitempty"` //The source, pertains to the original backup.
}

/*
 * Structure for the custom type RecoverVolumesTaskStateProto
 */
type RecoverVolumesTaskStateProto struct {
	Params        *RecoverVolumesParams                     `json:"params,omitempty" form:"params,omitempty"`               //TODO(Chinmaya): Rename this to RecoverPhysicalVolumesParams
	TaskResultVec []*RecoverVolumesTaskStateProtoTaskResult `json:"taskResultVec,omitempty" form:"taskResultVec,omitempty"` //Contains high-level per-volume information. This data is here because
}

/*
 * Structure for the custom type RecoverVolumesTaskStateProtoTaskResult
 */
type RecoverVolumesTaskStateProtoTaskResult struct {
	DstGuid                 *string     `json:"dstGuid,omitempty" form:"dstGuid,omitempty"`                                 //Volume GUID for the Target Entity (phy host).
	Error                   *ErrorProto `json:"error,omitempty" form:"error,omitempty"`                                     //TODO: Write general description for this field
	ProgressMonitorTaskPath *string     `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"` //The path relative to the root path of the restore task progress monitor.
}

/*
 * Structure for the custom type RecoveriesTile
 */
type RecoveriesTile struct {
	LastMonthNumRecoveries     *int64                      `json:"lastMonthNumRecoveries,omitempty" form:"lastMonthNumRecoveries,omitempty"`         //Number of Recoveries in the last 30 days.
	LastMonthRecoveriesByType  []*RestoreCountByObjectType `json:"lastMonthRecoveriesByType,omitempty" form:"lastMonthRecoveriesByType,omitempty"`   //Recoveries by Type in the last month.
	LastMonthRecoverySizeBytes *int64                      `json:"lastMonthRecoverySizeBytes,omitempty" form:"lastMonthRecoverySizeBytes,omitempty"` //Bytes recovered in the last 30 days.
	RecoveryNumRunning         *int64                      `json:"recoveryNumRunning,omitempty" form:"recoveryNumRunning,omitempty"`                 //Number of recoveries that are currently running.
}

/*
 * Structure for the custom type RecoveryTaskInfo
 */
type RecoveryTaskInfo struct {
	Name   *string                  `json:"name,omitempty" form:"name,omitempty"`     //Name of the recovery task.
	TaskId *string                  `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
	Type   TypeRecoveryTaskInfoEnum `json:"type,omitempty" form:"type,omitempty"`     //Denotes if the recovery task has an archival target.
}

/*
 * Structure for the custom type RegisterApplicationServersParameters
 */
type RegisterApplicationServersParameters struct {
	Applications       *[]ApplicationEnum `json:"applications,omitempty" form:"applications,omitempty"`             //Specifies the types of applications such as 'kSQL', 'kExchange' running
	HasPersistentAgent *bool              `json:"hasPersistentAgent,omitempty" form:"hasPersistentAgent,omitempty"` //Set this to true if a persistent agent is running on the host. If this is
	Password           *string            `json:"password,omitempty" form:"password,omitempty"`                     //Specifies password of the username to access the target source.
	ProtectionSourceId *int64             `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the Id of the Protection Source that contains one or more
	Username           *string            `json:"username,omitempty" form:"username,omitempty"`                     //Specifies username to access the target source.
}

/*
 * Structure for the custom type RegisterProtectionSourceParameters
 */
type RegisterProtectionSourceParameters struct {
	AgentEndpoint             *string                                           `json:"agentEndpoint,omitempty" form:"agentEndpoint,omitempty"`                         //Specifies the agent endpoint if it is different from the source endpoint.
	AwsCredentials            *AwsCredentials                                   `json:"awsCredentials,omitempty" form:"awsCredentials,omitempty"`                       //Specifies the credentials to authenticate with AWS Cloud Platform.
	AzureCredentials          *AzureCredentials                                 `json:"azureCredentials,omitempty" form:"azureCredentials,omitempty"`                   //Specifies the credentials to authenticate with Azure Cloud Platform.
	Endpoint                  *string                                           `json:"endpoint,omitempty" form:"endpoint,omitempty"`                                   //Specifies the network endpoint of the Protection Source where it is
	Environment               EnvironmentRegisterProtectionSourceParametersEnum `json:"environment,omitempty" form:"environment,omitempty"`                             //Specifies the environment such as 'kPhysical' or 'kVMware' of the
	ForceRegister             *bool                                             `json:"forceRegister,omitempty" form:"forceRegister,omitempty"`                         //ForceRegister is applicable to Physical Environment. By default, the agent
	GcpCredentials            *GcpCredentials                                   `json:"gcpCredentials,omitempty" form:"gcpCredentials,omitempty"`                       //Specifies the credentials to authenticate with Google Cloud Platform.
	HostType                  HostTypeRegisterProtectionSourceParametersEnum    `json:"hostType,omitempty" form:"hostType,omitempty"`                                   //Specifies the optional OS type of the Protection Source (such as kWindows
	HypervType                HypervTypeEnum                                    `json:"hyperVType,omitempty" form:"hyperVType,omitempty"`                               //Specifies the entity type if the environment is kHyperV.
	KubernetesCredentials     *KubernetesCredentials                            `json:"kubernetesCredentials,omitempty" form:"kubernetesCredentials,omitempty"`         //Specifies the credentials to authenticate with a Kubernetes Cluster.
	KubernetesType            KubernetesTypeEnum                                `json:"kubernetesType,omitempty" form:"kubernetesType,omitempty"`                       //Specifies the entity type if the environment is kKubernetes.
	NasMountCredentials       *NasMountCredentialParams                         `json:"nasMountCredentials,omitempty" form:"nasMountCredentials,omitempty"`             //Specifies the server credentials to connect to a NetApp server.
	NetappType                NetappTypeEnum                                    `json:"netappType,omitempty" form:"netappType,omitempty"`                               //Specifies the entity type such as 'kCluster,' if the environment is
	Office365Credentials      *Office365Credentials                             `json:"office365Credentials,omitempty" form:"office365Credentials,omitempty"`           //Specifies the credentials to authenticate with Office365 account.
	Office365Type             Office365TypeEnum                                 `json:"office365Type,omitempty" form:"office365Type,omitempty"`                         //Specifies the entity type such as 'kDomain', 'kOutlook', 'kMailbox', if the
	Password                  *string                                           `json:"password,omitempty" form:"password,omitempty"`                                   //Specifies password of the username to access the target source.
	PhysicalType              PhysicalTypeEnum                                  `json:"physicalType,omitempty" form:"physicalType,omitempty"`                           //Specifies the entity type such as 'kPhysicalHost' if the environment is
	PureType                  PureTypeEnum                                      `json:"pureType,omitempty" form:"pureType,omitempty"`                                   //Specifies the entity type such as 'kStorageArray' if the environment is
	SourceSideDedupEnabled    *bool                                             `json:"sourceSideDedupEnabled,omitempty" form:"sourceSideDedupEnabled,omitempty"`       //This controls whether to use source side dedup on the source or not.
	SslVerification           *SslVerification                                  `json:"sslVerification,omitempty" form:"sslVerification,omitempty"`                     //Specifies information about SSL verification when registering certain
	ThrottlingPolicy          *ThrottlingPolicyParameters                       `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"`                   //Specifies the throttling policy that should be applied to this Source.
	ThrottlingPolicyOverrides []*ThrottlingPolicyOverride                       `json:"throttlingPolicyOverrides,omitempty" form:"throttlingPolicyOverrides,omitempty"` //Array of Throttling Policy Overrides for Datastores.
	Username                  *string                                           `json:"username,omitempty" form:"username,omitempty"`                                   //Specifies username to access the target source.
	VmwareType                VmwareTypeEnum                                    `json:"vmwareType,omitempty" form:"vmwareType,omitempty"`                               //Specifies the entity type such as 'kVCenter' if the environment is
}

/*
 * Structure for the custom type RegisterRemoteCluster
 */
type RegisterRemoteCluster struct {
	AllEndpointsReachable   *bool                  `json:"allEndpointsReachable,omitempty" form:"allEndpointsReachable,omitempty"`     //Specifies whether any endpoint (such as a Node) on the remote Cluster
	BandwidthLimit          *BandwidthLimit        `json:"bandwidthLimit,omitempty" form:"bandwidthLimit,omitempty"`                   //Specifies settings for limiting the data transfer rate between
	ClearInterfaces         *bool                  `json:"clearInterfaces,omitempty" form:"clearInterfaces,omitempty"`                 //TODO: Write general description for this field
	ClearVlanId             *bool                  `json:"clearVlanId,omitempty" form:"clearVlanId,omitempty"`                         //Specifies whether to clear the vlanId field, and thus stop
	ClusterId               *int64                 `json:"clusterId,omitempty" form:"clusterId,omitempty"`                             //Specifies the unique id of the remote Cluster.
	CompressionEnabled      *bool                  `json:"compressionEnabled,omitempty" form:"compressionEnabled,omitempty"`           //Specifies whether to compress the outbound data when
	EncryptionKey           *string                `json:"encryptionKey,omitempty" form:"encryptionKey,omitempty"`                     //Specifies the encryption key used for encrypting the replication data
	IfaceName               *string                `json:"ifaceName,omitempty" form:"ifaceName,omitempty"`                             //Specifies the interface name of the VLAN to use for communicating with
	NetworkInterfaceGroup   *string                `json:"networkInterfaceGroup,omitempty" form:"networkInterfaceGroup,omitempty"`     //Specifies the group name of the network interfaces to
	NetworkInterfaceIds     *[]int64               `json:"networkInterfaceIds,omitempty" form:"networkInterfaceIds,omitempty"`         //Array of Network Interface Ids.
	Password                *string                `json:"password,omitempty" form:"password,omitempty"`                               //Specifies the password for Cohesity user to use when
	PurposeRemoteAccess     *bool                  `json:"purposeRemoteAccess,omitempty" form:"purposeRemoteAccess,omitempty"`         //Whether the remote cluster will be used for remote access for SPOG.
	PurposeReplication      *bool                  `json:"purposeReplication,omitempty" form:"purposeReplication,omitempty"`           //Whether the remote cluster will be used for replication.
	RemoteAccessCredentials *AccessTokenCredential `json:"remoteAccessCredentials,omitempty" form:"remoteAccessCredentials,omitempty"` //Specifies the Cohesity credentials required for generating an access token.
	RemoteIps               *[]string              `json:"remoteIps,omitempty" form:"remoteIps,omitempty"`                             //Array of Remote Node IP Addresses.
	RemoteIrisPorts         *[]int64               `json:"remoteIrisPorts,omitempty" form:"remoteIrisPorts,omitempty"`                 //Array of Ports.
	UserName                *string                `json:"userName,omitempty" form:"userName,omitempty"`                               //Specifies the Cohesity user name used to connect to the
	ValidateOnly            *bool                  `json:"validateOnly,omitempty" form:"validateOnly,omitempty"`                       //Whether to only validate the credentials without saving the information.
	ViewBoxPairInfo         []*ViewBoxPairInfo     `json:"viewBoxPairInfo,omitempty" form:"viewBoxPairInfo,omitempty"`                 //Array of Storage Domain (View Box) Pairs.
	VlanId                  *int64                 `json:"vlanId,omitempty" form:"vlanId,omitempty"`                                   //Specifies the Id of the VLAN to use for communicating with the remote
}

/*
 * Structure for the custom type RegisteredApplicationServer
 */
type RegisteredApplicationServer struct {
	ApplicationServer          *ProtectionSourceNode `json:"applicationServer,omitempty" form:"applicationServer,omitempty"`                   //Specifies the child subtree used to store additional application-level
	RegisteredProtectionSource *ProtectionSource     `json:"registeredProtectionSource,omitempty" form:"registeredProtectionSource,omitempty"` //Specifies the Protection Source like a VM or Physical Server that
}

/*
 * Structure for the custom type RegisteredSourceInfo
 */
type RegisteredSourceInfo struct {
	AccessInfo                 *ConnectorParameters                   `json:"accessInfo,omitempty" form:"accessInfo,omitempty"`                                 //Specifies the parameters required to establish a connection with
	AuthenticationErrorMessage *string                                `json:"authenticationErrorMessage,omitempty" form:"authenticationErrorMessage,omitempty"` //Specifies an authentication error message. This indicates the given
	AuthenticationStatus       AuthenticationStatusEnum               `json:"authenticationStatus,omitempty" form:"authenticationStatus,omitempty"`             //Specifies the status of the authenticating to the Protection Source
	Environments               *[]EnvironmentRegisteredSourceInfoEnum `json:"environments,omitempty" form:"environments,omitempty"`                             //Specifies a list of applications environment that are registered
	IsDbAuthenticated          *bool                                  `json:"isDbAuthenticated,omitempty" form:"isDbAuthenticated,omitempty"`                   //Specifies if application entity dbAuthenticated or not.
	MinimumFreeSpaceGB         *int64                                 `json:"minimumFreeSpaceGB,omitempty" form:"minimumFreeSpaceGB,omitempty"`                 //Specifies the minimum free space in GiB of the space expected to be
	NasMountCredentials        *NasMountCredentialParams              `json:"nasMountCredentials,omitempty" form:"nasMountCredentials,omitempty"`               //Specifies the credentials required to mount directories on the NetApp
	Password                   *string                                `json:"password,omitempty" form:"password,omitempty"`                                     //Specifies password of the username to access the target source.
	RefreshErrorMessage        *string                                `json:"refreshErrorMessage,omitempty" form:"refreshErrorMessage,omitempty"`               //Specifies a message if there was any error encountered during the last
	RefreshTimeUsecs           *int64                                 `json:"refreshTimeUsecs,omitempty" form:"refreshTimeUsecs,omitempty"`                     //Specifies the Unix epoch time (in microseconds) when the Protection
	RegistrationTimeUsecs      *int64                                 `json:"registrationTimeUsecs,omitempty" form:"registrationTimeUsecs,omitempty"`           //Specifies the Unix epoch time (in microseconds) when the Protection
	ThrottlingPolicy           *ThrottlingPolicyParameters            `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"`                     //Specifies the throttling policy for a registered Protection Source.
	ThrottlingPolicyOverrides  []*ThrottlingPolicyOverride            `json:"throttlingPolicyOverrides,omitempty" form:"throttlingPolicyOverrides,omitempty"`   //Array of Throttling Policy Overrides for Datastores.
	UseVmBiosUuid              *bool                                  `json:"useVmBiosUuid,omitempty" form:"useVmBiosUuid,omitempty"`                           //Specifies if registered vCenter is using BIOS UUID to track virtual
	Username                   *string                                `json:"username,omitempty" form:"username,omitempty"`                                     //Specifies username to access the target source.
	WarningMessages            *[]string                              `json:"warningMessages,omitempty" form:"warningMessages,omitempty"`                       //Specifies a list of warnings encountered during registration.
}

/*
 * Structure for the custom type RemoteCluster
 */
type RemoteCluster struct {
	AllEndpointsReachable   *bool                  `json:"allEndpointsReachable,omitempty" form:"allEndpointsReachable,omitempty"`     //Specifies whether any endpoint (such as a Node) on the remote Cluster
	BandwidthLimit          *BandwidthLimit        `json:"bandwidthLimit,omitempty" form:"bandwidthLimit,omitempty"`                   //Specifies settings for limiting the data transfer rate between
	ClearInterfaces         *bool                  `json:"clearInterfaces,omitempty" form:"clearInterfaces,omitempty"`                 //TODO: Write general description for this field
	ClearVlanId             *bool                  `json:"clearVlanId,omitempty" form:"clearVlanId,omitempty"`                         //Specifies whether to clear the vlanId field, and thus stop
	ClusterId               *int64                 `json:"clusterId,omitempty" form:"clusterId,omitempty"`                             //Specifies the unique id of the remote Cluster.
	ClusterIncarnationId    *int64                 `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"`       //Specifies the unique incarnation id of the remote Cluster. This
	CompressionEnabled      *bool                  `json:"compressionEnabled,omitempty" form:"compressionEnabled,omitempty"`           //Specifies whether to compress the outbound data when
	EncryptionKey           *string                `json:"encryptionKey,omitempty" form:"encryptionKey,omitempty"`                     //Specifies the encryption key used for encrypting the replication data
	IfaceName               *string                `json:"ifaceName,omitempty" form:"ifaceName,omitempty"`                             //Specifies the interface name of the VLAN to use for communicating with
	LocalIps                *[]string              `json:"localIps,omitempty" form:"localIps,omitempty"`                               //Array of Local IP Addresses.
	Name                    *string                `json:"name,omitempty" form:"name,omitempty"`                                       //Specifies the name of the remote Cluster.
	NetworkInterfaceGroup   *string                `json:"networkInterfaceGroup,omitempty" form:"networkInterfaceGroup,omitempty"`     //Specifies the group name of the network interfaces to
	NetworkInterfaceIds     *[]int64               `json:"networkInterfaceIds,omitempty" form:"networkInterfaceIds,omitempty"`         //Array of Network Interface Ids.
	PurposeRemoteAccess     *bool                  `json:"purposeRemoteAccess,omitempty" form:"purposeRemoteAccess,omitempty"`         //Whether the remote cluster will be used for remote access for SPOG.
	PurposeReplication      *bool                  `json:"purposeReplication,omitempty" form:"purposeReplication,omitempty"`           //Whether the remote cluster will be used for replication.
	RemoteAccessCredentials *AccessTokenCredential `json:"remoteAccessCredentials,omitempty" form:"remoteAccessCredentials,omitempty"` //Specifies the Cohesity credentials required for generating an access token.
	RemoteIps               *[]string              `json:"remoteIps,omitempty" form:"remoteIps,omitempty"`                             //Array of Remote Node IP Addresses.
	TenantId                *string                `json:"tenantId,omitempty" form:"tenantId,omitempty"`                               //Specifies the tenant Id of the organization that created this remote
	UserName                *string                `json:"userName,omitempty" form:"userName,omitempty"`                               //Specifies the Cohesity user name used to connect to the
	ViewBoxPairInfo         []*ViewBoxPairInfo     `json:"viewBoxPairInfo,omitempty" form:"viewBoxPairInfo,omitempty"`                 //Array of Storage Domain (View Box) Pairs.
	VlanId                  *int64                 `json:"vlanId,omitempty" form:"vlanId,omitempty"`                                   //Specifies the id of the VLAN to use when communicating with the remote
}

/*
 * Structure for the custom type RemoteHost
 */
type RemoteHost struct {
	Address *string            `json:"address,omitempty" form:"address,omitempty"` //Specifies the address (IP, hostname or FQDN) of the remote host
	Type    TypeRemoteHostEnum `json:"type,omitempty" form:"type,omitempty"`       //Specifies the OS type of the remote host that will run the script.
}

/*
 * Structure for the custom type RemoteHostConnectorParams
 */
type RemoteHostConnectorParams struct {
	Credentials *Credentials `json:"credentials,omitempty" form:"credentials,omitempty"` //Specifies credentials to access a target source.
	HostAddress *string      `json:"hostAddress,omitempty" form:"hostAddress,omitempty"` //Address (i.e., IP, hostname or FQDN) of the host to connect to. Magneto
	HostType    *int64       `json:"hostType,omitempty" form:"hostType,omitempty"`       //Type of host to connect to.
}

/*
 * Structure for the custom type RemoteJobScript
 */
type RemoteJobScript struct {
	FullBackupScript        *RemoteScriptPathAndParams `json:"fullBackupScript,omitempty" form:"fullBackupScript,omitempty"`               //Specifies the script that should run for the Full (no CBT) backup schedule
	IncrementalBackupScript *RemoteScriptPathAndParams `json:"incrementalBackupScript,omitempty" form:"incrementalBackupScript,omitempty"` //Specifies the script that should run for the CBT-based backup
	LogBackupScript         *RemoteScriptPathAndParams `json:"logBackupScript,omitempty" form:"logBackupScript,omitempty"`                 //Specifies the script that should run for the Log backup schedule
	RemoteHost              *RemoteHost                `json:"remoteHost,omitempty" form:"remoteHost,omitempty"`                           //Specifies the remote host where the remote scripts are executed.
	Username                *string                    `json:"username,omitempty" form:"username,omitempty"`                               //Specifies the username that will be used to login to the remote host.
}

/*
 * Structure for the custom type RemoteProtectionJobInformation
 */
type RemoteProtectionJobInformation struct {
	ClusterName *string                                       `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the name of the original Cluster that archived the data to the
	Environment EnvironmentRemoteProtectionJobInformationEnum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type (such as kVMware or kSQL)
	JobName     *string                                       `json:"jobName,omitempty" form:"jobName,omitempty"`         //Specifies the name of the Protection Job on the original Cluster.
	JobUid      *UniversalId                                  `json:"jobUid,omitempty" form:"jobUid,omitempty"`           //Specifies the globally unique id of the original Protection Job
}

/*
 * Structure for the custom type RemoteProtectionJobRunInformation
 */
type RemoteProtectionJobRunInformation struct {
	ClusterName       *string                                          `json:"clusterName,omitempty" form:"clusterName,omitempty"`             //Specifies the name of the original Cluster that archived the data to the
	Environment       EnvironmentRemoteProtectionJobRunInformationEnum `json:"environment,omitempty" form:"environment,omitempty"`             //Specifies the environment type (such as kVMware or kSQL)
	JobName           *string                                          `json:"jobName,omitempty" form:"jobName,omitempty"`                     //Specifies the name of the Protection Job on the original Cluster.
	JobUid            *UniversalId                                     `json:"jobUid,omitempty" form:"jobUid,omitempty"`                       //Specifies the globally unique id of the original Protection Job
	ProtectionJobRuns []*RemoteProtectionJobRunInstance                `json:"protectionJobRuns,omitempty" form:"protectionJobRuns,omitempty"` //Array of Protection Job Run Details.
}

/*
 * Structure for the custom type RemoteProtectionJobRunInstance
 */
type RemoteProtectionJobRunInstance struct {
	ArchiveTaskUid    *UniversalId `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"`       //Specifies the globally unique id of the archival task that archived
	ArchiveVersion    *int64       `json:"archiveVersion,omitempty" form:"archiveVersion,omitempty"`       //Specifies the version of the archive.
	ExpiryTimeUsecs   *int64       `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"`     //Specifies the time when the archive expires.
	IndexSizeBytes    *int64       `json:"indexSizeBytes,omitempty" form:"indexSizeBytes,omitempty"`       //Specifies the size of the index for the archive.
	JobRunId          *int64       `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                   //Specifies the instance id of the Job Run task capturing the Snapshot.
	MetadataComplete  *bool        `json:"metadataComplete,omitempty" form:"metadataComplete,omitempty"`   //Specifies whether a full set of metadata is available now.
	SnapshotTimeUsecs *int64       `json:"snapshotTimeUsecs,omitempty" form:"snapshotTimeUsecs,omitempty"` //Specify the time the Snapshot was captured as a Unix epoch Timestamp (in
}

/*
 * Structure for the custom type RemoteRestoreIndexingStatus
 */
type RemoteRestoreIndexingStatus struct {
	EndTimeUsecs               *int64                 `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                             //Specifies the end time of the time range that is being indexed.
	Error                      *string                `json:"error,omitempty" form:"error,omitempty"`                                           //Specifies the error message if the indexing Job/task fails.
	IndexingTaskEndTimeUsecs   *int64                 `json:"indexingTaskEndTimeUsecs,omitempty" form:"indexingTaskEndTimeUsecs,omitempty"`     //Specifies when the indexing task completed. This time is recorded
	IndexingTaskStartTimeUsecs *int64                 `json:"indexingTaskStartTimeUsecs,omitempty" form:"indexingTaskStartTimeUsecs,omitempty"` //Specifies when the indexing task started. This time is recorded
	IndexingTaskStatus         IndexingTaskStatusEnum `json:"indexingTaskStatus,omitempty" form:"indexingTaskStatus,omitempty"`                 //Specifies the status of the indexing Job/task.
	IndexingTaskUid            *UniversalId           `json:"indexingTaskUid,omitempty" form:"indexingTaskUid,omitempty"`                       //Specifies the unique id of the indexing task assigned by this Cluster.
	LatestExpiryTimeUsecs      *int64                 `json:"latestExpiryTimeUsecs,omitempty" form:"latestExpiryTimeUsecs,omitempty"`           //For all the Snapshots retrieved by this Job, specifies the latest time
	ProgressMonitorTask        *string                `json:"progressMonitorTask,omitempty" form:"progressMonitorTask,omitempty"`               //Specifies the path to progress monitor task to track the progress
	StartTimeUsecs             *int64                 `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                         //Specifies the start time of the time range that is being indexed.
}

/*
 * Structure for the custom type RemoteRestoreSnapshotStatus
 */
type RemoteRestoreSnapshotStatus struct {
	ArchiveTaskUid      *UniversalId           `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"`           //Specifies the globally unique id of the archival task that archived
	Error               *string                `json:"error,omitempty" form:"error,omitempty"`                             //Specifies the error message if the indexing task fails.
	ExpiryTimeUsecs     *int64                 `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"`         //Specifies the time when the Snapshot expires on the remote Vault.
	JobRunId            *int64                 `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                       //Specifies the id of the Job Run that originally captured the Snapshot.
	ProgressMonitorTask *string                `json:"progressMonitorTask,omitempty" form:"progressMonitorTask,omitempty"` //Specifies the path to the progress monitor task that tracks the progress
	SnapshotTaskStatus  SnapshotTaskStatusEnum `json:"snapshotTaskStatus,omitempty" form:"snapshotTaskStatus,omitempty"`   //Specifies the status of the indexing task.
	SnapshotTaskUid     *UniversalId           `json:"snapshotTaskUid,omitempty" form:"snapshotTaskUid,omitempty"`         //Specifies the globally unique id of the task capturing the Snapshot.
	SnapshotTimeUsecs   *int64                 `json:"snapshotTimeUsecs,omitempty" form:"snapshotTimeUsecs,omitempty"`     //Specify the time the Snapshot was captured.
}

/*
 * Structure for the custom type RemoteScriptPathAndParams
 */
type RemoteScriptPathAndParams struct {
	ContinueOnError *bool   `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the script needs to continue even if there is an occurence of
	IsActive        *bool   `json:"isActive,omitempty" form:"isActive,omitempty"`               //Specifies if the script is active. If set to false, this script will not
	ScriptParams    *string `json:"scriptParams,omitempty" form:"scriptParams,omitempty"`       //Specifies the parameters and values to pass into the remote script.
	ScriptPath      *string `json:"scriptPath,omitempty" form:"scriptPath,omitempty"`           //Specifies the path to the script on the remote host.
	TimeoutSecs     *int64  `json:"timeoutSecs,omitempty" form:"timeoutSecs,omitempty"`         //Specifies the timeout of the script in seconds. The script will be killed
}

/*
 * Structure for the custom type RemoteVaultRestoreTaskStatus
 */
type RemoteVaultRestoreTaskStatus struct {
	CurrentIndexingStatus          *RemoteRestoreIndexingStatus    `json:"currentIndexingStatus,omitempty" form:"currentIndexingStatus,omitempty"`                   //Specifies the status of an indexing task that builds an index from
	CurrentSnapshotStatus          *RemoteRestoreSnapshotStatus    `json:"currentSnapshotStatus,omitempty" form:"currentSnapshotStatus,omitempty"`                   //Specifies the status of the Snapshot restore task.
	LocalProtectionJobUid          *UniversalId                    `json:"localProtectionJobUid,omitempty" form:"localProtectionJobUid,omitempty"`                   //Specifies the globally unique id of the new inactive Protection Job
	ParentJobUid                   *UniversalId                    `json:"parentJobUid,omitempty" form:"parentJobUid,omitempty"`                                     //Specifies the unique id of the parent Job/task that spawned the
	RemoteProtectionJobInformation *RemoteProtectionJobInformation `json:"remoteProtectionJobInformation,omitempty" form:"remoteProtectionJobInformation,omitempty"` //Specifies details about the original Protection Job and its
	SearchJobUid                   *UniversalId                    `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"`                                     //Specifies the unique id of the search Job that searched the remote Vault.
}

/*
 * Structure for the custom type RemoteVaultSearchJobInformation
 */
type RemoteVaultSearchJobInformation struct {
	ClusterCount    *int64              `json:"clusterCount,omitempty" form:"clusterCount,omitempty"`       //Specifies number of Clusters that have archived to the remote Vault
	EndTimeUsecs    *int64              `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`       //Specifies the end time of the search as a Unix epoch
	Error           *string             `json:"error,omitempty" form:"error,omitempty"`                     //Specifies the error message reported when a search fails.
	JobCount        *int64              `json:"jobCount,omitempty" form:"jobCount,omitempty"`               //Specifies number of Protection Jobs that have archived to the remote Vault
	Name            *string             `json:"name,omitempty" form:"name,omitempty"`                       //Specifies the name of the search Job.
	SearchJobStatus SearchJobStatusEnum `json:"searchJobStatus,omitempty" form:"searchJobStatus,omitempty"` //Specifies the status of the search.
	SearchJobUid    *UniversalId        `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"`       //Specifies the unique id assigned to the search Job by the Cluster.
	StartTimeUsecs  *int64              `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`   //Specifies the start time of the search as a Unix epoch
	VaultId         *int64              `json:"vaultId,omitempty" form:"vaultId,omitempty"`                 //Specifies the id of the remote Vault (External Target) that was searched.
	VaultName       *string             `json:"vaultName,omitempty" form:"vaultName,omitempty"`             //Specifies the name of the remote Vault (External Target) that was searched.
}

/*
 * Structure for the custom type RemoteVaultSearchJobResults
 */
type RemoteVaultSearchJobResults struct {
	ClusterCount       *int64                                         `json:"clusterCount,omitempty" form:"clusterCount,omitempty"`             //Specifies number of Clusters that have archived to the remote Vault
	ClusterMatchString *string                                        `json:"clusterMatchString,omitempty" form:"clusterMatchString,omitempty"` //Specifies the value of the clusterMatchSting if it was set in the
	Cookie             *string                                        `json:"cookie,omitempty" form:"cookie,omitempty"`                         //Specifies an opaque string to pass to the next request to get the
	EndTimeUsecs       *int64                                         `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`             //Specifies the value of endTimeUsecs if it was set in the original
	Error              *string                                        `json:"error,omitempty" form:"error,omitempty"`                           //Specifies the error message if the search fails.
	JobCount           *int64                                         `json:"jobCount,omitempty" form:"jobCount,omitempty"`                     //Specifies number of Protection Jobs that have archived to the remote Vault
	JobMatchString     *string                                        `json:"jobMatchString,omitempty" form:"jobMatchString,omitempty"`         //Specifies the value of the jobMatchSting if it was set in the
	ProtectionJobs     []*RemoteProtectionJobRunInformation           `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"`         //Array of Protection Jobs.
	SearchJobStatus    SearchJobStatusRemoteVaultSearchJobResultsEnum `json:"searchJobStatus,omitempty" form:"searchJobStatus,omitempty"`       //Specifies the status of the search Job.
	SearchJobUid       *UniversalId                                   `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"`             //Specifies the unique id of the search Job assigned by the Cluster.
	StartTimeUsecs     *int64                                         `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`         //Specifies the value of startTimeUsecs if it was set in the original
	VaultId            *int64                                         `json:"vaultId,omitempty" form:"vaultId,omitempty"`                       //Specifies the id of the remote Vault that was searched.
	VaultName          *string                                        `json:"vaultName,omitempty" form:"vaultName,omitempty"`                   //Specifies the name of the remote Vault that was searched.
}

/*
 * Structure for the custom type RenameObjectParamProto
 */
type RenameObjectParamProto struct {
	Prefix *string `json:"prefix,omitempty" form:"prefix,omitempty"` //Prefix to be added to a name.
	Suffix *string `json:"suffix,omitempty" form:"suffix,omitempty"` //Suffix to be added to a name.
}

/*
 * Structure for the custom type RenameViewParam
 */
type RenameViewParam struct {
	NewViewName string `json:"newViewName" form:"newViewName"` //Specifies the new name of the View.
}

/*
 * Structure for the custom type ReplicateSnapshotsToAWSParams
 */
type ReplicateSnapshotsToAWSParams struct {
	Region *EntityProto `json:"region,omitempty" form:"region,omitempty"` //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type ReplicationEncryptionKeyReponse
 */
type ReplicationEncryptionKeyReponse struct {
	EncryptionKey *string `json:"encryptionKey,omitempty" form:"encryptionKey,omitempty"` //Specifies a replication encryption key.
}

/*
 * Structure for the custom type ReplicationTarget
 */
type ReplicationTarget struct {
	ClusterId   *int64  `json:"clusterId,omitempty" form:"clusterId,omitempty"`     //The id of the remote cluster.
	ClusterName *string `json:"clusterName,omitempty" form:"clusterName,omitempty"` //The name of the remote cluster.
}

/*
 * Structure for the custom type ReplicationTargetSettings
 */
type ReplicationTargetSettings struct {
	ClusterId   *int64  `json:"clusterId,omitempty" form:"clusterId,omitempty"`     //Specifies the id of the Remote Cluster.
	ClusterName *string `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the name of the Remote Cluster.
}

/*
 * Structure for the custom type RequestError
 */
type RequestError struct {
	ErrorCode *int64  `json:"errorCode,omitempty" form:"errorCode,omitempty"` //Operation response error code.
	Message   *string `json:"message,omitempty" form:"message,omitempty"`     //Description of the error.
}

/*
 * Structure for the custom type ResetS3SecretKeyParameters
 */
type ResetS3SecretKeyParameters struct {
	Domain   *string `json:"domain,omitempty" form:"domain,omitempty"`     //Specifies the fully qualified domain name (FQDN) of an Active Directory
	TenantId *string `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the tenant for which the the users are to be deleted.
	Username *string `json:"username,omitempty" form:"username,omitempty"` //Specifies the Cohesity user.
}

/*
 * Structure for the custom type RestoreADAppObjectParams
 */
type RestoreADAppObjectParams struct {
	AdRestoreStatusVec []*ADRestoreStatus          `json:"adRestoreStatusVec,omitempty" form:"adRestoreStatusVec,omitempty"` //Status of the AD object/attribute restore operation.
	AdUpdateOptions    *ADUpdateRestoreTaskOptions `json:"adUpdateOptions,omitempty" form:"adUpdateOptions,omitempty"`       //TODO: Write general description for this field
	Credentials        *Credentials                `json:"credentials,omitempty" form:"credentials,omitempty"`               //Specifies credentials to access a target source.
	LdapPort           *int64                      `json:"ldapPort,omitempty" form:"ldapPort,omitempty"`                     //The ldap port on which the AD domain controller's NTDS database will be
	NumFailed          *int64                      `json:"numFailed,omitempty" form:"numFailed,omitempty"`                   //Number of AD objects whose restore failed. Includes both AD object and
	NumRunning         *int64                      `json:"numRunning,omitempty" form:"numRunning,omitempty"`                 //Number of AD objects whose restores are currently running. Includes both
	NumSuccessfull     *int64                      `json:"numSuccessfull,omitempty" form:"numSuccessfull,omitempty"`         //Number of AD objects restored successfully. Includes both AD object and
}

/*
 * Structure for the custom type RestoreAcropolisVMParam
 */
type RestoreAcropolisVMParam struct {
	BaseCbtSnapshotInfoProto *SnapshotInfoProto                        `json:"baseCbtSnapshotInfoProto,omitempty" form:"baseCbtSnapshotInfoProto,omitempty"` //Each available extension is listed below along with the location of the
	NetworkConfig            *RestoreAcropolisVMParamNetworkConfigInfo `json:"networkConfig,omitempty" form:"networkConfig,omitempty"`                       //Proto to define the network configuration to be applied to the restored
}

/*
 * Structure for the custom type RestoreAcropolisVMParamNetworkConfigInfo
 */
type RestoreAcropolisVMParamNetworkConfigInfo struct {
	NetworkStateChange *int64                                             `json:"networkStateChange,omitempty" form:"networkStateChange,omitempty"` //Network state to be applied to the restored VM.
	NicVec             []*RestoreAcropolisVMParamNetworkConfigInfoNICSpec `json:"nicVec,omitempty" form:"nicVec,omitempty"`                         //This field is applicable only if the network_state_change is set to
}

/*
 * Structure for the custom type RestoreAcropolisVMParamNetworkConfigInfoNICSpec
 */
type RestoreAcropolisVMParamNetworkConfigInfoNICSpec struct {
	IpAddress   *string `json:"ipAddress,omitempty" form:"ipAddress,omitempty"`     //IP address to assign to the NIC.
	NetworkUuid *string `json:"networkUuid,omitempty" form:"networkUuid,omitempty"` //The UUID of the network to which the NIC is to be attached.
}

/*
 * Structure for the custom type RestoreAcropolisVMsParams
 */
type RestoreAcropolisVMsParams struct {
	PowerStateConfig             *PowerStateConfigProto            `json:"powerStateConfig,omitempty" form:"powerStateConfig,omitempty"`                         //TODO: Write general description for this field
	RenameRestoredObjectParam    *RenameObjectParamProto           `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"`       //Message to specify the prefix/suffix added to rename an object. At least one
	RestoredObjectsNetworkConfig *RestoredObjectNetworkConfigProto `json:"restoredObjectsNetworkConfig,omitempty" form:"restoredObjectsNetworkConfig,omitempty"` //TODO: Write general description for this field
	StorageContainer             *EntityProto                      `json:"storageContainer,omitempty" form:"storageContainer,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	UuidConfig                   *UuidConfigProto                  `json:"uuidConfig,omitempty" form:"uuidConfig,omitempty"`                                     //TODO: Write general description for this field
}

/*
 * Structure for the custom type RestoreAppObject
 */
type RestoreAppObject struct {
	AppEntity     *EntityProto            `json:"appEntity,omitempty" form:"appEntity,omitempty"`         //Specifies the attributes and the latest statistics about an entity.
	DisplayName   *string                 `json:"displayName,omitempty" form:"displayName,omitempty"`     //The proper display name of this object in the UI, if app_entity is not
	RestoreParams *RestoreAppObjectParams `json:"restoreParams,omitempty" form:"restoreParams,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type RestoreAppObjectParams
 */
type RestoreAppObjectParams struct {
	AdRestoreParams        *RestoreADAppObjectParams     `json:"adRestoreParams,omitempty" form:"adRestoreParams,omitempty"`               //TODO: Write general description for this field
	CloneTaskId            *int64                        `json:"cloneTaskId,omitempty" form:"cloneTaskId,omitempty"`                       //Id of finished clone task which has to be refreshed with different data.
	OracleRestoreParams    *RestoreOracleAppObjectParams `json:"oracleRestoreParams,omitempty" form:"oracleRestoreParams,omitempty"`       //TODO: Write general description for this field
	SqlRestoreParams       *RestoreSqlAppObjectParams    `json:"sqlRestoreParams,omitempty" form:"sqlRestoreParams,omitempty"`             //TODO: Write general description for this field
	TargetHost             *EntityProto                  `json:"targetHost,omitempty" form:"targetHost,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	TargetHostParentSource *EntityProto                  `json:"targetHostParentSource,omitempty" form:"targetHostParentSource,omitempty"` //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RestoreAppParams
 */
type RestoreAppParams struct {
	Credentials         *Credentials         `json:"credentials,omitempty" form:"credentials,omitempty"`                 //Specifies credentials to access a target source.
	OwnerRestoreInfo    *AppOwnerRestoreInfo `json:"ownerRestoreInfo,omitempty" form:"ownerRestoreInfo,omitempty"`       //TODO: Write general description for this field
	RestoreAppObjectVec []*RestoreAppObject  `json:"restoreAppObjectVec,omitempty" form:"restoreAppObjectVec,omitempty"` //The application level objects that needs to be restored. If this vector is
	Type                *int64               `json:"type,omitempty" form:"type,omitempty"`                               //The application environment.
}

/*
 * Structure for the custom type RestoreAppTaskStateProto
 */
type RestoreAppTaskStateProto struct {
	AppRestoreProgressMonitorSubtaskPath *string           `json:"appRestoreProgressMonitorSubtaskPath,omitempty" form:"appRestoreProgressMonitorSubtaskPath,omitempty"` //The Pulse task path to the application restore task sub tree. If the
	LastFinishedLogBackupStartTimeUsecs  *int64            `json:"lastFinishedLogBackupStartTimeUsecs,omitempty" form:"lastFinishedLogBackupStartTimeUsecs,omitempty"`   //The start time of the last finished log backup run. For SQL application,
	RestoreAppParams                     *RestoreAppParams `json:"restoreAppParams,omitempty" form:"restoreAppParams,omitempty"`                                         //This message captures all the necessary arguments specified by the user to
}

/*
 * Structure for the custom type RestoreCountByObjectType
 */
type RestoreCountByObjectType struct {
	ObjectCount *int64  `json:"objectCount,omitempty" form:"objectCount,omitempty"` //Specifies the number of restores of the object type.
	ObjectType  *string `json:"objectType,omitempty" form:"objectType,omitempty"`   //Specifies the type of the restored object.
}

/*
 * Structure for the custom type RestoreEnvStats
 */
type RestoreEnvStats struct {
	Environment EnvironmentRestoreEnvStatsEnum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment.
	ObjectCount *int64                         `json:"objectCount,omitempty" form:"objectCount,omitempty"` //TODO: Write general description for this field
	TotalBytes  *int64                         `json:"totalBytes,omitempty" form:"totalBytes,omitempty"`   //TODO: Write general description for this field
}

/*
 * Structure for the custom type RestoreFileCopyStats
 */
type RestoreFileCopyStats struct {
	EstimationSkipped      *bool  `json:"estimationSkipped,omitempty" form:"estimationSkipped,omitempty"`           //This will be set to true if the estimation step was skipped.
	NumBytesCopied         *int64 `json:"numBytesCopied,omitempty" form:"numBytesCopied,omitempty"`                 //Number of bytes copied so far.
	NumDirectoriesCopied   *int64 `json:"numDirectoriesCopied,omitempty" form:"numDirectoriesCopied,omitempty"`     //Number of directories copied so far.
	NumFilesCopied         *int64 `json:"numFilesCopied,omitempty" form:"numFilesCopied,omitempty"`                 //Number of files copied so far.
	TotalBytesToCopy       *int64 `json:"totalBytesToCopy,omitempty" form:"totalBytesToCopy,omitempty"`             //Total number of bytes to copy.
	TotalDirectoriesToCopy *int64 `json:"totalDirectoriesToCopy,omitempty" form:"totalDirectoriesToCopy,omitempty"` //Total number of directories to copy.
	TotalFilesToCopy       *int64 `json:"totalFilesToCopy,omitempty" form:"totalFilesToCopy,omitempty"`             //Total number of files to copy.
}

/*
 * Structure for the custom type RestoreFileResultInfo
 */
type RestoreFileResultInfo struct {
	CopyStats        *RestoreFileCopyStats `json:"copyStats,omitempty" form:"copyStats,omitempty"`               //This message captures the progress information regarding restore of
	DestinationDir   *string               `json:"destinationDir,omitempty" form:"destinationDir,omitempty"`     //This is set to the destination directory where the file/directory was
	Error            *ErrorProto           `json:"error,omitempty" form:"error,omitempty"`                       //TODO: Write general description for this field
	RestoredFileInfo *RestoredFileInfo     `json:"restoredFileInfo,omitempty" form:"restoredFileInfo,omitempty"` //TODO: Write general description for this field
	Status           *int64                `json:"status,omitempty" form:"status,omitempty"`                     //Status of the restore.
}

/*
 * Structure for the custom type RestoreFilesInfoProto
 */
type RestoreFilesInfoProto struct {
	DownloadFilesPath          *string                  `json:"downloadFilesPath,omitempty" form:"downloadFilesPath,omitempty"`                   //The path that the user should use to download files through the UI. If
	Error                      *ErrorProto              `json:"error,omitempty" form:"error,omitempty"`                                           //TODO: Write general description for this field
	ProxyEntityConnectorParams *ConnectorParams         `json:"proxyEntityConnectorParams,omitempty" form:"proxyEntityConnectorParams,omitempty"` //Message that encapsulates the various params required to establish a
	RestoreFilesResultVec      []*RestoreFileResultInfo `json:"restoreFilesResultVec,omitempty" form:"restoreFilesResultVec,omitempty"`           //Contains the result information of the restored files.
	SlaveTaskStartTimeUsecs    *int64                   `json:"slaveTaskStartTimeUsecs,omitempty" form:"slaveTaskStartTimeUsecs,omitempty"`       //This is the timestamp at which the slave task started.
	TargetType                 *int64                   `json:"targetType,omitempty" form:"targetType,omitempty"`                                 //Specifies the target type for the task. The field is only valid if the
	TeardownError              *ErrorProto              `json:"teardownError,omitempty" form:"teardownError,omitempty"`                           //TODO: Write general description for this field
	Type                       *int64                   `json:"type,omitempty" form:"type,omitempty"`                                             //The type of environment this restore files info pertains to.
}

/*
 * Structure for the custom type RestoreFilesParams
 */
type RestoreFilesParams struct {
	IsFileVolumeRestore      *bool                    `json:"isFileVolumeRestore,omitempty" form:"isFileVolumeRestore,omitempty"`           //Whether this is a file based volume restore.
	IsMountBasedFlr          *bool                    `json:"isMountBasedFlr,omitempty" form:"isMountBasedFlr,omitempty"`                   //Whether this is a mount based file restore operation
	NasProtocolTypeVec       *[]int64                 `json:"nasProtocolTypeVec,omitempty" form:"nasProtocolTypeVec,omitempty"`             //The NAS protocol type(s) of this restore task. Currently we only support a
	ProxyEntity              *EntityProto             `json:"proxyEntity,omitempty" form:"proxyEntity,omitempty"`                           //Specifies the attributes and the latest statistics about an entity.
	ProxyEntityParentSource  *EntityProto             `json:"proxyEntityParentSource,omitempty" form:"proxyEntityParentSource,omitempty"`   //Specifies the attributes and the latest statistics about an entity.
	RestoreFilesPreferences  *RestoreFilesPreferences `json:"restoreFilesPreferences,omitempty" form:"restoreFilesPreferences,omitempty"`   //This message captures preferences from the user while restoring the files
	RestoredFileInfoVec      []*RestoredFileInfo      `json:"restoredFileInfoVec,omitempty" form:"restoredFileInfoVec,omitempty"`           //Information regarding files and directories.
	TargetEntity             *EntityProto             `json:"targetEntity,omitempty" form:"targetEntity,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	TargetEntityCredentials  *Credentials             `json:"targetEntityCredentials,omitempty" form:"targetEntityCredentials,omitempty"`   //Specifies credentials to access a target source.
	TargetEntityParentSource *EntityProto             `json:"targetEntityParentSource,omitempty" form:"targetEntityParentSource,omitempty"` //Specifies the attributes and the latest statistics about an entity.
	TargetHostEntity         *EntityProto             `json:"targetHostEntity,omitempty" form:"targetHostEntity,omitempty"`                 //Specifies the attributes and the latest statistics about an entity.
	TargetHostType           *int64                   `json:"targetHostType,omitempty" form:"targetHostType,omitempty"`                     //The host environment type. This is set in VMware environment to
}

/*
 * Structure for the custom type RestoreFilesPreferences
 */
type RestoreFilesPreferences struct {
	AlternateRestoreBaseDirectory *string `json:"alternateRestoreBaseDirectory,omitempty" form:"alternateRestoreBaseDirectory,omitempty"` //This must be set to a directory path if restore_to_original_paths is
	ContinueOnError               *bool   `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`                             //Whether to continue with the copy in case of encountering an error.
	GenerateSshKeys               *bool   `json:"generateSshKeys,omitempty" form:"generateSshKeys,omitempty"`                             //In case of GCP Linux restores, whether to generate ssh keys to connect to
	OverrideOriginals             *bool   `json:"overrideOriginals,omitempty" form:"overrideOriginals,omitempty"`                         //This is relevant only if restore_to_original_paths is true. If this is
	PreserveAcls                  *bool   `json:"preserveAcls,omitempty" form:"preserveAcls,omitempty"`                                   //Whether to preserve the ACLs of the original file.
	PreserveAttributes            *bool   `json:"preserveAttributes,omitempty" form:"preserveAttributes,omitempty"`                       //Whether to preserve the original attributes.
	PreserveTimestamps            *bool   `json:"preserveTimestamps,omitempty" form:"preserveTimestamps,omitempty"`                       //Whether to preserve the original time stamps.
	RestoreToOriginalPaths        *bool   `json:"restoreToOriginalPaths,omitempty" form:"restoreToOriginalPaths,omitempty"`               //If this is true, then files will be restored to original paths.
	SkipEstimation                *bool   `json:"skipEstimation,omitempty" form:"skipEstimation,omitempty"`                               //Whether to skip the estimation step.
}

/*
 * Structure for the custom type RestoreFilesTaskRequest
 */
type RestoreFilesTaskRequest struct {
	ContinueOnError          *bool                 `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`                   //Specifies if the Restore Task should continue even if the copy operation
	Filenames                *[]string             `json:"filenames,omitempty" form:"filenames,omitempty"`                               //Array of Files or Folders.
	IsFileBasedVolumeRestore *bool                 `json:"isFileBasedVolumeRestore,omitempty" form:"isFileBasedVolumeRestore,omitempty"` //Specifies whether this is a file based volume restore.
	Name                     *string               `json:"name,omitempty" form:"name,omitempty"`                                         //Specifies the name of the Restore Task. This field must be set and
	NewBaseDirectory         *string               `json:"newBaseDirectory,omitempty" form:"newBaseDirectory,omitempty"`                 //Specifies an optional root folder where to recover the selected
	Overwrite                *bool                 `json:"overwrite,omitempty" form:"overwrite,omitempty"`                               //If true, any existing files and folders on the operating system
	Password                 *string               `json:"password,omitempty" form:"password,omitempty"`                                 //Specifies password of the username to access the target source.
	PreserveAttributes       *bool                 `json:"preserveAttributes,omitempty" form:"preserveAttributes,omitempty"`             //If true, the Restore Tasks preserves the original file and
	SourceObjectInfo         *RestoreObjectDetails `json:"sourceObjectInfo,omitempty" form:"sourceObjectInfo,omitempty"`                 //Specifies information about the source object (such as a VM)
	TargetHostType           TargetHostTypeEnum    `json:"targetHostType,omitempty" form:"targetHostType,omitempty"`                     //Specifies the target host types to be restored.
	TargetParentSourceId     *int64                `json:"targetParentSourceId,omitempty" form:"targetParentSourceId,omitempty"`         //Specifies the registered source (such as a vCenter Server)
	TargetSourceId           *int64                `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"`                     //Specifies the id of the target protection source (such as a VM)
	Username                 *string               `json:"username,omitempty" form:"username,omitempty"`                                 //Specifies username to access the target source.
}

/*
 * Structure for the custom type RestoreFilesTaskStateProto
 */
type RestoreFilesTaskStateProto struct {
	RestoreFilesInfo *RestoreFilesInfoProto `json:"restoreFilesInfo,omitempty" form:"restoreFilesInfo,omitempty"` //Each available extension is listed below along with the location of the
	RestoreParams    *RestoreFilesParams    `json:"restoreParams,omitempty" form:"restoreParams,omitempty"`       //This message captures all the necessary arguments specified by the user to
}

/*
 * Structure for the custom type RestoreHypervVMParams
 */
type RestoreHypervVMParams struct {
	CopyRecovery                 *bool                             `json:"copyRecovery,omitempty" form:"copyRecovery,omitempty"`                                 //Whether to perform copy recovery.
	DatastoreEntity              *EntityProto                      `json:"datastoreEntity,omitempty" form:"datastoreEntity,omitempty"`                           //Specifies the attributes and the latest statistics about an entity.
	PowerStateConfig             *PowerStateConfigProto            `json:"powerStateConfig,omitempty" form:"powerStateConfig,omitempty"`                         //TODO: Write general description for this field
	RenameRestoredObjectParam    *RenameObjectParamProto           `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"`       //Message to specify the prefix/suffix added to rename an object. At least one
	ResourceEntity               *EntityProto                      `json:"resourceEntity,omitempty" form:"resourceEntity,omitempty"`                             //Specifies the attributes and the latest statistics about an entity.
	RestoredObjectsNetworkConfig *RestoredObjectNetworkConfigProto `json:"restoredObjectsNetworkConfig,omitempty" form:"restoredObjectsNetworkConfig,omitempty"` //TODO: Write general description for this field
	UuidConfig                   *UuidConfigProto                  `json:"uuidConfig,omitempty" form:"uuidConfig,omitempty"`                                     //TODO: Write general description for this field
}

/*
 * Structure for the custom type RestoreInfo
 */
type RestoreInfo struct {
	ArchivalTarget          *ArchivalExternalTarget   `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`                   //Specifies settings about the Archival External Target (such as Tape or AWS).
	AttemptNumber           *int64                    `json:"attemptNumber,omitempty" form:"attemptNumber,omitempty"`                     //Specifies the attempt number.
	CloudDeployTarget       *CloudDeployTargetDetails `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"`             //Message that specifies the details about CloudDeploy target where backup
	JobRunId                *int64                    `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                               //Specifies the id of the job run.
	JobUid                  *UniversalId              `json:"jobUid,omitempty" form:"jobUid,omitempty"`                                   //Specifies an id for an object that is unique across Cohesity Clusters.
	ParentSource            *ProtectionSource         `json:"parentSource,omitempty" form:"parentSource,omitempty"`                       //Specifies a generic structure that represents a node
	SnapshotRelativeDirPath *string                   `json:"snapshotRelativeDirPath,omitempty" form:"snapshotRelativeDirPath,omitempty"` //Specifies the relative path of the snapshot directory.
	Source                  *ProtectionSource         `json:"source,omitempty" form:"source,omitempty"`                                   //Specifies a generic structure that represents a node
	StartTimeUsecs          *int64                    `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                   //Specifies the start time specified as a Unix epoch Timestamp
	ViewName                *string                   `json:"viewName,omitempty" form:"viewName,omitempty"`                               //Specifies the name of the view.
	VmHadIndependentDisks   *bool                     `json:"vmHadIndependentDisks,omitempty" form:"vmHadIndependentDisks,omitempty"`     //Specifies if the VM had independent disks.
}

/*
 * Structure for the custom type RestoreInfoProto
 */
type RestoreInfoProto struct {
	RestoreEntityVec []*RestoreInfoProtoRestoreEntity `json:"restoreEntityVec,omitempty" form:"restoreEntityVec,omitempty"` //Contains the file paths and the information of the restored entities.
	TargetType       *int64                           `json:"targetType,omitempty" form:"targetType,omitempty"`             //Specifies the target type for the task. The field is only valid if the
	Type             *int64                           `json:"type,omitempty" form:"type,omitempty"`                         //The type of environment this restore info pertains to.
}

/*
 * Structure for the custom type RestoreInfoProtoRestoreEntity
 */
type RestoreInfoProtoRestoreEntity struct {
	Entity                  *EntityProto  `json:"entity,omitempty" form:"entity,omitempty"`                                   //Specifies the attributes and the latest statistics about an entity.
	Error                   *ErrorProto   `json:"error,omitempty" form:"error,omitempty"`                                     //TODO: Write general description for this field
	ProgressMonitorTaskPath *string       `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"` //The path relative to the root path of the restore task progress monitor
	PublicStatus            *int64        `json:"publicStatus,omitempty" form:"publicStatus,omitempty"`                       //Iris-facing task state. This field is stamped during the export.
	RelativeRestorePaths    *[]string     `json:"relativeRestorePaths,omitempty" form:"relativeRestorePaths,omitempty"`       //All the paths that the entity's files were restored to. Each path is
	ResourcePoolEntity      *EntityProto  `json:"resourcePoolEntity,omitempty" form:"resourcePoolEntity,omitempty"`           //Specifies the attributes and the latest statistics about an entity.
	RestoredEntity          *EntityProto  `json:"restoredEntity,omitempty" form:"restoredEntity,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
	Status                  *int64        `json:"status,omitempty" form:"status,omitempty"`                                   //The restore status of the entity.
	Warnings                []*ErrorProto `json:"warnings,omitempty" form:"warnings,omitempty"`                               //Optional warnings if any.
}

/*
 * Structure for the custom type RestoreKVMVMsParams
 */
type RestoreKVMVMsParams struct {
	ClusterEntity                *EntityProto                      `json:"clusterEntity,omitempty" form:"clusterEntity,omitempty"`                               //Specifies the attributes and the latest statistics about an entity.
	DatacenterEntity             *EntityProto                      `json:"datacenterEntity,omitempty" form:"datacenterEntity,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	PowerStateConfig             *PowerStateConfigProto            `json:"powerStateConfig,omitempty" form:"powerStateConfig,omitempty"`                         //TODO: Write general description for this field
	RenameRestoredObjectParam    *RenameObjectParamProto           `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"`       //Message to specify the prefix/suffix added to rename an object. At least one
	RestoredObjectsNetworkConfig *RestoredObjectNetworkConfigProto `json:"restoredObjectsNetworkConfig,omitempty" form:"restoredObjectsNetworkConfig,omitempty"` //TODO: Write general description for this field
	StoragedomainEntity          *EntityProto                      `json:"storagedomainEntity,omitempty" form:"storagedomainEntity,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RestoreKubernetesNamespacesParams
 */
type RestoreKubernetesNamespacesParams struct {
	BackupJobName             *string                 `json:"backupJobName,omitempty" form:"backupJobName,omitempty"`                         //Backup job that needs to be used for recovering the namespace.
	ClusterEntity             *EntityProto            `json:"clusterEntity,omitempty" form:"clusterEntity,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	ManagementNamespace       *string                 `json:"managementNamespace,omitempty" form:"managementNamespace,omitempty"`             //Namespace in which restore job will be created in K8s cluster.
	RenameRestoredObjectParam *RenameObjectParamProto `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"` //Message to specify the prefix/suffix added to rename an object. At least one
}

/*
 * Structure for the custom type RestoreObject
 */
type RestoreObject struct {
	ArchivalTarget          *ArchivalTarget          `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`                   //Message that specifies the details about an archival target (such as cloud
	AttemptNum              *int64                   `json:"attemptNum,omitempty" form:"attemptNum,omitempty"`                           //The attempt number of the job run to restore from.
	CloudDeployTarget       *CloudDeployTarget       `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"`             //Message that specifies the details about CloudDeploy target where backup
	CloudReplicationTarget  *CloudDeployTarget       `json:"cloudReplicationTarget,omitempty" form:"cloudReplicationTarget,omitempty"`   //Message that specifies the details about CloudDeploy target where backup
	Entity                  *EntityProto             `json:"entity,omitempty" form:"entity,omitempty"`                                   //Specifies the attributes and the latest statistics about an entity.
	JobId                   *int64                   `json:"jobId,omitempty" form:"jobId,omitempty"`                                     //The job id from which to restore. This is used while communicating with
	JobInstanceId           *int64                   `json:"jobInstanceId,omitempty" form:"jobInstanceId,omitempty"`                     //Id identifying a specific run to restore from. If this is not specified,
	JobUid                  *UniversalIdProto        `json:"jobUid,omitempty" form:"jobUid,omitempty"`                                   //TODO: Write general description for this field
	ParentSource            *EntityProto             `json:"parentSource,omitempty" form:"parentSource,omitempty"`                       //Specifies the attributes and the latest statistics about an entity.
	RestoreAcropolisVmParam *RestoreAcropolisVMParam `json:"restoreAcropolisVmParam,omitempty" form:"restoreAcropolisVmParam,omitempty"` //TODO: Write general description for this field
	SnapshotRelativeDirPath *string                  `json:"snapshotRelativeDirPath,omitempty" form:"snapshotRelativeDirPath,omitempty"` //The relative path to the directory containing the entity's snapshot.
	StartTimeUsecs          *int64                   `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                   //The start time of the specific job run. Iff 'job_instance_id' is set,
	ViewName                *string                  `json:"viewName,omitempty" form:"viewName,omitempty"`                               //The name of the view where the object's snapshot is located.
	VmHadIndependentDisks   *bool                    `json:"vmHadIndependentDisks,omitempty" form:"vmHadIndependentDisks,omitempty"`     //This is applicable only to VMs and is set to true when the VM being
}

/*
 * Structure for the custom type RestoreObjectDetails
 */
type RestoreObjectDetails struct {
	ArchivalTarget     *ArchivalExternalTarget             `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`         //Specifies settings about the Archival Target (such as Tape or AWS).
	CloudDeployTarget  *CloudDeployTargetDetails           `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"`   //Specifies settings about the Cloud Deploy target.
	Environment        EnvironmentRestoreObjectDetailsEnum `json:"environment,omitempty" form:"environment,omitempty"`               //Specifies the type of the Protection Source.
	JobId              *int64                              `json:"jobId,omitempty" form:"jobId,omitempty"`                           //Protection Job Id.
	JobRunId           *int64                              `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                     //Specifies the id of the Job Run that captured the snapshot.
	JobUid             *UniversalId                        `json:"jobUid,omitempty" form:"jobUid,omitempty"`                         //Specifies the universal id of the Protection Job that backed up
	ProtectionSourceId *int64                              `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the id of the leaf object to recover, clone or recover
	SourceName         *string                             `json:"sourceName,omitempty" form:"sourceName,omitempty"`                 //Specifies the name of the Protection Source.
	StartedTimeUsecs   *int64                              `json:"startedTimeUsecs,omitempty" form:"startedTimeUsecs,omitempty"`     //Specifies the time when the Job Run starts capturing a snapshot.
}

/*
 * Structure for the custom type RestoreObjectParams
 */
type RestoreObjectParams struct {
	Action                       *int64                            `json:"action,omitempty" form:"action,omitempty"`                                             //The action to perform.
	DatastoreEntity              *EntityProto                      `json:"datastoreEntity,omitempty" form:"datastoreEntity,omitempty"`                           //Specifies the attributes and the latest statistics about an entity.
	PowerStateConfig             *PowerStateConfigProto            `json:"powerStateConfig,omitempty" form:"powerStateConfig,omitempty"`                         //TODO: Write general description for this field
	RenameRestoredObjectParam    *RenameObjectParamProto           `json:"renameRestoredObjectParam,omitempty" form:"renameRestoredObjectParam,omitempty"`       //Message to specify the prefix/suffix added to rename an object. At least one
	ResourcePoolEntity           *EntityProto                      `json:"resourcePoolEntity,omitempty" form:"resourcePoolEntity,omitempty"`                     //Specifies the attributes and the latest statistics about an entity.
	RestoreParentSource          *EntityProto                      `json:"restoreParentSource,omitempty" form:"restoreParentSource,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
	RestoredObjectsNetworkConfig *RestoredObjectNetworkConfigProto `json:"restoredObjectsNetworkConfig,omitempty" form:"restoredObjectsNetworkConfig,omitempty"` //TODO: Write general description for this field
	ViewName                     *string                           `json:"viewName,omitempty" form:"viewName,omitempty"`                                         //Target view into which the objects are to be cloned.
}

/*
 * Structure for the custom type RestoreObjectState
 */
type RestoreObjectState struct {
	Error            *RequestError    `json:"error,omitempty" form:"error,omitempty"`                       //Details about the Error.
	ObjectStatus     ObjectStatusEnum `json:"objectStatus,omitempty" form:"objectStatus,omitempty"`         //Specifies the status of an object during a Restore Task.
	ResourcePoolId   *int64           `json:"resourcePoolId,omitempty" form:"resourcePoolId,omitempty"`     //Specifies the id of the Resource Pool that the restored
	RestoredObjectId *int64           `json:"restoredObjectId,omitempty" form:"restoredObjectId,omitempty"` //Specifies the Id of the recovered or cloned object.
	SourceObjectId   *int64           `json:"sourceObjectId,omitempty" form:"sourceObjectId,omitempty"`     //Specifies the Protection Source id of the object to be recovered or
}

/*
 * Structure for the custom type RestoreOneDriveParams
 */
type RestoreOneDriveParams struct {
	DriveOwnerVec     []*RestoreOneDriveParamsDriveOwner `json:"driveOwnerVec,omitempty" form:"driveOwnerVec,omitempty"`         //The list of users/groups whose drives are being restored.
	RestoreToOriginal *bool                              `json:"restoreToOriginal,omitempty" form:"restoreToOriginal,omitempty"` //Whether or not all drive items are restored to original location.
	TargetDriveId     *string                            `json:"targetDriveId,omitempty" form:"targetDriveId,omitempty"`         //The id of the drive in which items will be restored.
	TargetFolderPath  *string                            `json:"targetFolderPath,omitempty" form:"targetFolderPath,omitempty"`   //All drives part of various users listed in drive_owner_vec will be
	TargetUser        *EntityProto                       `json:"targetUser,omitempty" form:"targetUser,omitempty"`               //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RestoreOneDriveParamsDriveItem
 */
type RestoreOneDriveParamsDriveItem struct {
	DriveItemPath *string `json:"driveItemPath,omitempty" form:"driveItemPath,omitempty"` //The path of the drive item relative to root.
	Id            *string `json:"id,omitempty" form:"id,omitempty"`                       //The unique identifier of the item within the Drive.
	IsFileItem    *bool   `json:"isFileItem,omitempty" form:"isFileItem,omitempty"`       //Specify if the item is a file or not.
}

/*
 * Structure for the custom type RestoreOneDriveParamsDriveOwner
 */
type RestoreOneDriveParamsDriveOwner struct {
	DriveVec []*RestoreOneDriveParamsDriveOwnerDrive `json:"driveVec,omitempty" form:"driveVec,omitempty"` //The list of drives that are being restored.
	Object   *RestoreObject                          `json:"object,omitempty" form:"object,omitempty"`     //TODO: Write general description for this field
}

/*
 * Structure for the custom type RestoreOneDriveParamsDriveOwnerDrive
 */
type RestoreOneDriveParamsDriveOwnerDrive struct {
	IsEntireDriveRequired *bool                             `json:"isEntireDriveRequired,omitempty" form:"isEntireDriveRequired,omitempty"` //Specify if the entire drive is to be restored.
	RestoreDriveId        *string                           `json:"restoreDriveId,omitempty" form:"restoreDriveId,omitempty"`               //Id of the drive whose items are being restored.
	RestoreItemVec        []*RestoreOneDriveParamsDriveItem `json:"restoreItemVec,omitempty" form:"restoreItemVec,omitempty"`               //List of drive paths that need to be restored.
}

/*
 * Structure for the custom type RestoreOracleAppObjectParams
 */
type RestoreOracleAppObjectParams struct {
	AlternateLocationParams     *RestoreOracleAppObjectParamsAlternateLocationParams `json:"alternateLocationParams,omitempty" form:"alternateLocationParams,omitempty"`         //For restoring to alternate location this message can not be empty and all
	NoOpenMode                  *bool                                                `json:"noOpenMode,omitempty" form:"noOpenMode,omitempty"`                                   //If set to true, the recovered database will not be opened.
	OracleCloneAppViewParamsVec []*CloneAppViewParams                                `json:"oracleCloneAppViewParamsVec,omitempty" form:"oracleCloneAppViewParamsVec,omitempty"` //Following field contains information related to view expose workflow. Ex
	OracleTargetParams          *OracleSourceParams                                  `json:"oracleTargetParams,omitempty" form:"oracleTargetParams,omitempty"`                   //Message to capture additional backup/restore params for a Oracle source.
	RestoreTimeSecs             *int64                                               `json:"restoreTimeSecs,omitempty" form:"restoreTimeSecs,omitempty"`                         //The time to which the Oracle database needs to be restored. This allows
}

/*
 * Structure for the custom type RestoreOracleAppObjectParamsAlternateLocationParams
 */
type RestoreOracleAppObjectParamsAlternateLocationParams struct {
	BaseDir                 *string         `json:"baseDir,omitempty" form:"baseDir,omitempty"`                                 //Base directory of Oracle at destination.
	DatabaseFileDestination *string         `json:"databaseFileDestination,omitempty" form:"databaseFileDestination,omitempty"` //Location to put the database files(datafiles, logfiles etc.).
	HomeDir                 *string         `json:"homeDir,omitempty" form:"homeDir,omitempty"`                                 //Home directory of Oracle at destination.
	NewDatabaseName         *string         `json:"newDatabaseName,omitempty" form:"newDatabaseName,omitempty"`                 //The name of the Oracle database that we restore to.
	NewSidDeprecated        *string         `json:"newSidDeprecated,omitempty" form:"newSidDeprecated,omitempty"`               //Deprecated field
	OracleDbConfig          *OracleDBConfig `json:"oracleDbConfig,omitempty" form:"oracleDbConfig,omitempty"`                   //This proto captures the oracle database configuration for alternate DB
}

/*
 * Structure for the custom type RestoreOutlookParams
 */
type RestoreOutlookParams struct {
	MailboxVec       []*RestoreOutlookParamsMailbox `json:"mailboxVec,omitempty" form:"mailboxVec,omitempty"`             //In a RestoreJob , user will provide the list of mailboxes
	TargetFolderPath *string                        `json:"targetFolderPath,omitempty" form:"targetFolderPath,omitempty"` //TODO: Write general description for this field
	TargetMailbox    *EntityProto                   `json:"targetMailbox,omitempty" form:"targetMailbox,omitempty"`       //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RestoreOutlookParamsFolder
 */
type RestoreOutlookParamsFolder struct {
	FolderId               *string   `json:"folderId,omitempty" form:"folderId,omitempty"`                             //The Unique ID of the folder.
	FolderKey              *int64    `json:"folderKey,omitempty" form:"folderKey,omitempty"`                           //The Unique key of the folder.
	IsEntireFolderRequired *bool     `json:"isEntireFolderRequired,omitempty" form:"isEntireFolderRequired,omitempty"` //Specify if the entire folder is to be restored.
	ItemIdVec              *[]string `json:"itemIdVec,omitempty" form:"itemIdVec,omitempty"`                           //If is_entire_folder_required is set to false,
}

/*
 * Structure for the custom type RestoreOutlookParamsMailbox
 */
type RestoreOutlookParamsMailbox struct {
	FolderVec               []*RestoreOutlookParamsFolder `json:"folderVec,omitempty" form:"folderVec,omitempty"`                             //If is_entire_mailbox_required is set to false,
	IsEntireMailboxRequired *bool                         `json:"isEntireMailboxRequired,omitempty" form:"isEntireMailboxRequired,omitempty"` //Specify if the entire mailbox is to be restored.
	Object                  *RestoreObject                `json:"object,omitempty" form:"object,omitempty"`                                   //TODO: Write general description for this field
}

/*
 * Structure for the custom type RestorePointsForTimeRange
 */
type RestorePointsForTimeRange struct {
	FullSnapshotInfo []*FullSnapshotInfo  `json:"fullSnapshotInfo,omitempty" form:"fullSnapshotInfo,omitempty"` //Specifies the info related to the recovery object.
	TimeRanges       []*TimeRangeSettings `json:"timeRanges,omitempty" form:"timeRanges,omitempty"`             //Specifies the time ranges of the restore object between full snapshots.
}

/*
 * Structure for the custom type RestorePointsForTimeRangeParam
 */
type RestorePointsForTimeRangeParam struct {
	EndTimeUsecs       *int64                                        `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`             //Specifies the end time specified as a Unix epoch Timestamp
	Environment        EnvironmentRestorePointsForTimeRangeParamEnum `json:"environment,omitempty" form:"environment,omitempty"`               //Specifies the protection source environment type.
	JobUids            []*UniversalId                                `json:"jobUids" form:"jobUids"`                                           //Specifies the jobs for which to get the full snapshot information.
	ProtectionSourceId *int64                                        `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the id of the Protection Source which is to be restored.
	StartTimeUsecs     *int64                                        `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`         //Specifies the start time specified as a Unix epoch Timestamp
}

/*
 * Structure for the custom type RestoreSqlAppObjectParams
 */
type RestoreSqlAppObjectParams struct {
	CaptureTailLogs                 *bool                        `json:"captureTailLogs,omitempty" form:"captureTailLogs,omitempty"`                                 //Set to true if tail logs are to be captured before the restore
	DataFileDestination             *string                      `json:"dataFileDestination,omitempty" form:"dataFileDestination,omitempty"`                         //Which directory to put the database data files. Missing directory will be
	DbRestoreOverwritePolicy        *int64                       `json:"dbRestoreOverwritePolicy,omitempty" form:"dbRestoreOverwritePolicy,omitempty"`               //Policy to overwrite an existing DB during a restore operation.
	InstanceName                    *string                      `json:"instanceName,omitempty" form:"instanceName,omitempty"`                                       //The name of the SQL instance that we restore database to. If target_host
	IsMultiStageRestore             *bool                        `json:"isMultiStageRestore,omitempty" form:"isMultiStageRestore,omitempty"`                         //The following field is set if we are creating a multi-stage SQL restore
	LogFileDestination              *string                      `json:"logFileDestination,omitempty" form:"logFileDestination,omitempty"`                           //Which directory to put the database log files. Missing directory will be
	MultiStageRestoreOptions        *SqlUpdateRestoreTaskOptions `json:"multiStageRestoreOptions,omitempty" form:"multiStageRestoreOptions,omitempty"`               //TODO: Write general description for this field
	NewDatabaseName                 *string                      `json:"newDatabaseName,omitempty" form:"newDatabaseName,omitempty"`                                 //The new name of the database, if it is going to be renamed. app_entity in
	RestoreTimeSecs                 *int64                       `json:"restoreTimeSecs,omitempty" form:"restoreTimeSecs,omitempty"`                                 //The time to which the SQL database needs to be restored. This allows for
	SecondaryDataFileDestination    *string                      `json:"secondaryDataFileDestination,omitempty" form:"secondaryDataFileDestination,omitempty"`       //Which directory to put the secondary data files of the database. Secondary
	SecondaryDataFileDestinationVec []*FilesToDirectoryMapping   `json:"secondaryDataFileDestinationVec,omitempty" form:"secondaryDataFileDestinationVec,omitempty"` //Specify the secondary data files and corresponding direcories of the DB.
	WithNoRecovery                  *bool                        `json:"withNoRecovery,omitempty" form:"withNoRecovery,omitempty"`                                   //Set to true if we want to recover the database in "NO_RECOVERY" mode
}

/*
 * Structure for the custom type RestoreStats
 */
type RestoreStats struct {
	NumClonedObjects    *int64             `json:"numClonedObjects,omitempty" form:"numClonedObjects,omitempty"`       //Specifies the count of cloned objects in the given time frame.
	NumRecoveredObjects *int64             `json:"numRecoveredObjects,omitempty" form:"numRecoveredObjects,omitempty"` //Specifies the count of recovered objects in the given time frame.
	StatsByEnvironment  []*RestoreEnvStats `json:"statsByEnvironment,omitempty" form:"statsByEnvironment,omitempty"`   //Specifies the stats of recovery jobs aggregated by the environment type.
}

/*
 * Structure for the custom type RestoreTask
 */
type RestoreTask struct {
	AcropolisParameters     *AcropolisRestoreParameters   `json:"acropolisParameters,omitempty" form:"acropolisParameters,omitempty"`         //This field defines the Acropolis specific params for restore tasks of type
	ApplicationParameters   *ApplicationRestoreParameters `json:"applicationParameters,omitempty" form:"applicationParameters,omitempty"`     //Specifies the information regarding the application restore parameters.
	ArchiveTaskUid          *UniversalId                  `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"`                   //Specifies the uid of the Restore Task that retrieves objects from
	CloneViewParameters     *UpdateViewParam              `json:"cloneViewParameters,omitempty" form:"cloneViewParameters,omitempty"`         //Specifies the View settings used when cloning a View.
	ContinueOnError         *bool                         `json:"continueOnError,omitempty" form:"continueOnError,omitempty"`                 //Specifies if the Restore Task should continue when some operations on some
	DatastoreId             *int64                        `json:"datastoreId,omitempty" form:"datastoreId,omitempty"`                         //Specifies the datastore where the object's files are recovered to.
	DeployVmsToCloud        *DeployVmsToCloud             `json:"deployVmsToCloud,omitempty" form:"deployVmsToCloud,omitempty"`               //Specifies the details about deploying vms to specific clouds where backup
	EndTimeUsecs            *int64                        `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                       //Specifies the end time of the Restore Task as a Unix epoch
	Error                   *RequestError                 `json:"error,omitempty" form:"error,omitempty"`                                     //Specifies the error reported by the Restore Task (if any) after the
	FullViewName            *string                       `json:"fullViewName,omitempty" form:"fullViewName,omitempty"`                       //Specifies the full name of a View.
	HypervParameters        *HypervRestoreParameters      `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"`               //Specifies information needed when restoring VMs in HyperV enviroment.
	Id                      *int64                        `json:"id,omitempty" form:"id,omitempty"`                                           //Specifies the id of the Restore Task assigned by
	MountVolumesState       *MountVolumesState            `json:"mountVolumesState,omitempty" form:"mountVolumesState,omitempty"`             //Specifies the states of mounting all the volumes onto a mount target
	Name                    string                        `json:"name" form:"name"`                                                           //Specifies the name of the Restore Task. This field must be set and
	NewParentId             *int64                        `json:"newParentId,omitempty" form:"newParentId,omitempty"`                         //Specify a new registered parent Protection Source. If specified
	Objects                 []*RestoreObjectDetails       `json:"objects,omitempty" form:"objects,omitempty"`                                 //Array of Objects.
	OutlookParameters       *OutlookRestoreParameters     `json:"outlookParameters,omitempty" form:"outlookParameters,omitempty"`             //Specifies information needed for recovering Mailboxes in O365Outlook
	RestoreObjectState      []*RestoreObjectState         `json:"restoreObjectState,omitempty" form:"restoreObjectState,omitempty"`           //Array of Object States.
	StartTimeUsecs          *int64                        `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                   //Specifies the start time for the Restore Task as a Unix epoch
	Status                  StatusRestoreTaskEnum         `json:"status,omitempty" form:"status,omitempty"`                                   //Specifies the overall status of the Restore Task.
	TargetViewCreated       *bool                         `json:"targetViewCreated,omitempty" form:"targetViewCreated,omitempty"`             //Is true if a new View was created by a 'kCloneVMs' Restore Task.
	Type                    TypeRestoreTaskEnum           `json:"type,omitempty" form:"type,omitempty"`                                       //Specifies the type of Restore Task.
	Username                *string                       `json:"username,omitempty" form:"username,omitempty"`                               //Specifies the Cohesity user who requested this Restore Task.
	ViewBoxId               *int64                        `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                             //Specifies the id of the Domain (View Box) where the View is stored.
	VirtualDiskRestoreState *VirtualDiskRecoverTaskState  `json:"virtualDiskRestoreState,omitempty" form:"virtualDiskRestoreState,omitempty"` //Specifies the complete information about a recover virtual disk task state.
	VlanParameters          *VlanParameters               `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"`                   //Specifies VLAN parameters for the restore operation.
	VmwareParameters        *VmwareRestoreParameters      `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"`               //Specifies the information required for recovering or cloning VmWare VMs.
}

/*
 * Structure for the custom type RestoreTaskStateBaseProto
 */
type RestoreTaskStateBaseProto struct {
	CancellationRequested        *bool              `json:"cancellationRequested,omitempty" form:"cancellationRequested,omitempty"`               //Whether this task has a pending cancellation request.
	EndTimeUsecs                 *int64             `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                                 //If the restore task has finished, this field contains the end time for the
	Error                        *ErrorProto        `json:"error,omitempty" form:"error,omitempty"`                                               //TODO: Write general description for this field
	Name                         *string            `json:"name,omitempty" form:"name,omitempty"`                                                 //The name of the restore task.
	ParentSourceConnectionParams *ConnectorParams   `json:"parentSourceConnectionParams,omitempty" form:"parentSourceConnectionParams,omitempty"` //Message that encapsulates the various params required to establish a
	PublicStatus                 *int64             `json:"publicStatus,omitempty" form:"publicStatus,omitempty"`                                 //Iris-facing task state. This field is stamped during the export.
	RefreshStatus                *int64             `json:"refreshStatus,omitempty" form:"refreshStatus,omitempty"`                               //Status of the refresh task.
	RestoreVlanParams            *RestoreVlanParams `json:"restoreVlanParams,omitempty" form:"restoreVlanParams,omitempty"`                       //TODO: Write general description for this field
	ScheduledConstituentId       *int64             `json:"scheduledConstituentId,omitempty" form:"scheduledConstituentId,omitempty"`             //Constituent id (and the gandalf session id) where this task has been
	ScheduledGandalfSessionId    *int64             `json:"scheduledGandalfSessionId,omitempty" form:"scheduledGandalfSessionId,omitempty"`       //TODO: Write general description for this field
	StartTimeUsecs               *int64             `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                             //The start time for this restore task.
	Status                       *int64             `json:"status,omitempty" form:"status,omitempty"`                                             //Status of the restore task.
	TaskId                       *int64             `json:"taskId,omitempty" form:"taskId,omitempty"`                                             //A globally unique id for this task.
	TotalLogicalSizeBytes        *int64             `json:"totalLogicalSizeBytes,omitempty" form:"totalLogicalSizeBytes,omitempty"`               //Logical size of this restore task. This is the amount of data that needs
	TotalPhysicalSizeBytes       *int64             `json:"totalPhysicalSizeBytes,omitempty" form:"totalPhysicalSizeBytes,omitempty"`             //Physical size of this restore task. This is the amount of data that was
	Type                         *int64             `json:"type,omitempty" form:"type,omitempty"`                                                 //The type of restore being performed.
	User                         *string            `json:"user,omitempty" form:"user,omitempty"`                                                 //The user who requested this restore task.
	UserInfo                     *UserInformation   `json:"userInfo,omitempty" form:"userInfo,omitempty"`                                         //A message to encapsulate information about the user who made the request.
	UserMessages                 *[]string          `json:"userMessages,omitempty" form:"userMessages,omitempty"`                                 //Messages displayed to the user for this task (if any).
	Warnings                     []*ErrorProto      `json:"warnings,omitempty" form:"warnings,omitempty"`                                         //The warnings encountered by this task (if any) during its execution.
}

/*
 * Structure for the custom type RestoreTaskWrapper
 */
type RestoreTaskWrapper struct {
	RestoreTask *RestoreWrapperProto `json:"restoreTask,omitempty" form:"restoreTask,omitempty"` //If this message is a checkpoint record in WAL-logging or if this message is
}

/*
 * Structure for the custom type RestoreVmwareVMParams
 */
type RestoreVmwareVMParams struct {
	CopyRecovery            *bool        `json:"copyRecovery,omitempty" form:"copyRecovery,omitempty"`                       //Whether to perform copy recovery instead of instant recovery.
	DatastoreEntity         *EntityProto `json:"datastoreEntity,omitempty" form:"datastoreEntity,omitempty"`                 //Specifies the attributes and the latest statistics about an entity.
	PreserveTagsDuringClone *bool        `json:"preserveTagsDuringClone,omitempty" form:"preserveTagsDuringClone,omitempty"` //Whether to preserve tags for the clone op.
	ResourcePoolEntity      *EntityProto `json:"resourcePoolEntity,omitempty" form:"resourcePoolEntity,omitempty"`           //Specifies the attributes and the latest statistics about an entity.
	TargetDatastoreFolder   *EntityProto `json:"targetDatastoreFolder,omitempty" form:"targetDatastoreFolder,omitempty"`     //Specifies the attributes and the latest statistics about an entity.
	TargetVmFolder          *EntityProto `json:"targetVmFolder,omitempty" form:"targetVmFolder,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RestoreVlanParams
 */
type RestoreVlanParams struct {
	DisableVlan   *bool   `json:"disableVlan,omitempty" form:"disableVlan,omitempty"`     //If this is set to true, then even if VLANs are configured on the system,
	InterfaceName *string `json:"interfaceName,omitempty" form:"interfaceName,omitempty"` //Interface group to use for restore. If this is not specified, primary
	VlanId        *int64  `json:"vlanId,omitempty" form:"vlanId,omitempty"`               //If this is set, then the Cohesity host name or the IP address associated
}

/*
 * Structure for the custom type RestoreWrapperProto
 */
type RestoreWrapperProto struct {
	DestroyClonedTaskStateVec     []*DestroyClonedTaskStateProto  `json:"destroyClonedTaskStateVec,omitempty" form:"destroyClonedTaskStateVec,omitempty"`         //For a restore task of type 'Clone', this field contains the info of the
	OwnerRestoreWrapperProto      *RestoreWrapperProto            `json:"ownerRestoreWrapperProto,omitempty" form:"ownerRestoreWrapperProto,omitempty"`           //TODO: Write general description for this field
	PerformRefreshTaskStateVec    []*PerformRestoreTaskStateProto `json:"performRefreshTaskStateVec,omitempty" form:"performRefreshTaskStateVec,omitempty"`       //Contains information of the refresh tasks for a clone
	PerformRestoreJobState        *PerformRestoreJobStateProto    `json:"performRestoreJobState,omitempty" form:"performRestoreJobState,omitempty"`               //Proto to define the persistent information of a wrapper restore job that
	PerformRestoreTaskState       *PerformRestoreTaskStateProto   `json:"performRestoreTaskState,omitempty" form:"performRestoreTaskState,omitempty"`             //TODO: Write general description for this field
	RestoreSubTaskWrapperProtoVec *[]interface{}                  `json:"restoreSubTaskWrapperProtoVec,omitempty" form:"restoreSubTaskWrapperProtoVec,omitempty"` //If this restore has sub tasks, the following field will get populated
}

/*
 * Structure for the custom type RestoredFileInfo
 */
type RestoredFileInfo struct {
	AbsolutePath      *string `json:"absolutePath,omitempty" form:"absolutePath,omitempty"`           //Full path of the file being restored: the actual file path without the
	AttachedDiskId    *int64  `json:"attachedDiskId,omitempty" form:"attachedDiskId,omitempty"`       //Disk information of where the source file is currently located.
	DiskPartitionId   *int64  `json:"diskPartitionId,omitempty" form:"diskPartitionId,omitempty"`     //Disk partition to which the file belongs to.
	IsDirectory       *bool   `json:"isDirectory,omitempty" form:"isDirectory,omitempty"`             //Whether the path points to a directory.
	IsNonSimpleLdmVol *bool   `json:"isNonSimpleLdmVol,omitempty" form:"isNonSimpleLdmVol,omitempty"` //This will be set to true for recovery workflows for non-simple volumes
	RestoreMountPoint *string `json:"restoreMountPoint,omitempty" form:"restoreMountPoint,omitempty"` //Mount point of the volume on which the file to be restored is located.
	SizeBytes         *int64  `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"`                 //Size of the file in bytes. Required in FLR in GCP using Cloud Functions.
	VirtualDiskFile   *string `json:"virtualDiskFile,omitempty" form:"virtualDiskFile,omitempty"`     //Virtual disk file to which this file belongs to.
	VolumeId          *string `json:"volumeId,omitempty" form:"volumeId,omitempty"`                   //Id of the volume.
	VolumePath        *string `json:"volumePath,omitempty" form:"volumePath,omitempty"`               //Original volume name (or drive letter). This is used while performing the
}

/*
 * Structure for the custom type RestoredObjectNetworkConfigProto
 */
type RestoredObjectNetworkConfigProto struct {
	DetachNetwork                  *bool                  `json:"detachNetwork,omitempty" form:"detachNetwork,omitempty"`                                   //If this is set to true, then the network will be detached from the
	DisableNetwork                 *bool                  `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"`                                 //This can be set to true to indicate that the attached network should be
	Mappings                       []*NetworkMappingProto `json:"mappings,omitempty" form:"mappings,omitempty"`                                             //The network mappings to be applied to the target object.
	NetworkEntity                  *EntityProto           `json:"networkEntity,omitempty" form:"networkEntity,omitempty"`                                   //Specifies the attributes and the latest statistics about an entity.
	PreserveMacAddressOnNewNetwork *bool                  `json:"preserveMacAddressOnNewNetwork,omitempty" form:"preserveMacAddressOnNewNetwork,omitempty"` //If this is true and we are attaching to a new network entity, then the
	VnicEntity                     *EntityProto           `json:"vnicEntity,omitempty" form:"vnicEntity,omitempty"`                                         //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RestoredObjectVCDConfigProto
 */
type RestoredObjectVCDConfigProto struct {
	IsVapp                 *bool            `json:"isVapp,omitempty" form:"isVapp,omitempty"`                                 //Whether the restored object is a VApp.
	VappEntity             *EntityProto     `json:"vappEntity,omitempty" form:"vappEntity,omitempty"`                         //Specifies the attributes and the latest statistics about an entity.
	VcenterConnectorParams *ConnectorParams `json:"vcenterConnectorParams,omitempty" form:"vcenterConnectorParams,omitempty"` //Message that encapsulates the various params required to establish a
	VdcEntity              *EntityProto     `json:"vdcEntity,omitempty" form:"vdcEntity,omitempty"`                           //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type RetentionPolicyProto
 */
type RetentionPolicyProto struct {
	NumDaysToKeep *int64              `json:"numDaysToKeep,omitempty" form:"numDaysToKeep,omitempty"` //The number of days to keep the snapshots for a backup run.
	WormRetention *WormRetentionProto `json:"wormRetention,omitempty" form:"wormRetention,omitempty"` //Message that specifies the WORM attributes. WORM attributes can be
}

/*
 * Structure for the custom type RetrieveArchiveInfo
 */
type RetrieveArchiveInfo struct {
	AvgLogicalTransferRateBps *int64                                `json:"avgLogicalTransferRateBps,omitempty" form:"avgLogicalTransferRateBps,omitempty"` //Average logical bytes transfer rate in bytes per second as seen by Icebox.
	BytesTransferred          *int64                                `json:"bytesTransferred,omitempty" form:"bytesTransferred,omitempty"`                   //Number of physical bytes transferred for this retrieval task so far.
	EndTimeUsecs              *int64                                `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                           //Time when this retrieval task ended at Icebox side. If not set, then the
	Error                     *ErrorProto                           `json:"error,omitempty" form:"error,omitempty"`                                         //TODO: Write general description for this field
	LogicalBytesTransferred   *int64                                `json:"logicalBytesTransferred,omitempty" form:"logicalBytesTransferred,omitempty"`     //Number of logical bytes transferred so far.
	LogicalSizeBytes          *int64                                `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"`                   //Total logical size of the retrieval task.
	ProgressMonitorTaskPath   *string                               `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"`     //The root path of the progress monitor for this task.
	RetrievedEntityVec        []*RetrieveArchiveInfoRetrievedEntity `json:"retrievedEntityVec,omitempty" form:"retrievedEntityVec,omitempty"`               //Contains info about all retrieved entities.
	StartTimeUsecs            *int64                                `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                       //Time when this retrieval task was started by Icebox. If not set, then
	StubViewName              *string                               `json:"stubViewName,omitempty" form:"stubViewName,omitempty"`                           //The stub view that Icebox created. Stub view can be used for selectively
	StubViewRelativeDirName   *string                               `json:"stubViewRelativeDirName,omitempty" form:"stubViewRelativeDirName,omitempty"`     //Relative directory inside the stub view that corresponds with the archive.
	TargetViewName            *string                               `json:"targetViewName,omitempty" form:"targetViewName,omitempty"`                       //The name of the target view where Icebox has retrieved and staged the
	UserActionRequiredMsg     *string                               `json:"userActionRequiredMsg,omitempty" form:"userActionRequiredMsg,omitempty"`         //Message to display in the UI if any manual intervention is needed to make
}

/*
 * Structure for the custom type RetrieveArchiveInfoRetrievedEntity
 */
type RetrieveArchiveInfoRetrievedEntity struct {
	BytesTransferred        *int64       `json:"bytesTransferred,omitempty" form:"bytesTransferred,omitempty"`               //Number of physical bytes transferred over wire for this entity.
	EndTimestampUsecs       *int64       `json:"endTimestampUsecs,omitempty" form:"endTimestampUsecs,omitempty"`             //Time in microseconds when retrieve of this entity finished or failed.
	Entity                  *EntityProto `json:"entity,omitempty" form:"entity,omitempty"`                                   //Specifies the attributes and the latest statistics about an entity.
	Error                   *ErrorProto  `json:"error,omitempty" form:"error,omitempty"`                                     //TODO: Write general description for this field
	LogicalBytesTransferred *int64       `json:"logicalBytesTransferred,omitempty" form:"logicalBytesTransferred,omitempty"` //Number of logical bytes transferred so far.
	LogicalSizeBytes        *int64       `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"`               //Total logical size of this entity.
	ProgressMonitorTaskPath *string      `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"` //The path relative to the root path of the retrieval task progress
	RelativeSnapshotDir     *string      `json:"relativeSnapshotDir,omitempty" form:"relativeSnapshotDir,omitempty"`         //The path relative to the root of the file system where the snapshot of
	StartTimestampUsecs     *int64       `json:"startTimestampUsecs,omitempty" form:"startTimestampUsecs,omitempty"`         //Time in microseconds when retrieve of this entity started.
	Status                  *int64       `json:"status,omitempty" form:"status,omitempty"`                                   //The retrieval status of this entity.
}

/*
 * Structure for the custom type RetrieveArchiveTaskStateProto
 */
type RetrieveArchiveTaskStateProto struct {
	ArchivalTarget          *ArchivalTarget                                 `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`                   //Message that specifies the details about an archival target (such as cloud
	ArchiveTaskUid          *UniversalIdProto                               `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"`                   //TODO: Write general description for this field
	BackupRunStartTimeUsecs *int64                                          `json:"backupRunStartTimeUsecs,omitempty" form:"backupRunStartTimeUsecs,omitempty"` //The start time of the backup run whose corresponding archive is being
	CancellationRequested   *bool                                           `json:"cancellationRequested,omitempty" form:"cancellationRequested,omitempty"`     //Whether this retrieval task has a pending cancellation request.
	DownloadFilesInfo       *RetrieveArchiveTaskStateProtoDownloadFilesInfo `json:"downloadFilesInfo,omitempty" form:"downloadFilesInfo,omitempty"`             //Information required for Icebox when downloading files from an archived
	EndTimeUsecs            *int64                                          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`                       //If the retrieval task has finished, this field contains the end time for
	EntityVec               []*EntityProto                                  `json:"entityVec,omitempty" form:"entityVec,omitempty"`                             //Information on the exact set of objects to retrieve from archive. Even if
	Error                   *ErrorProto                                     `json:"error,omitempty" form:"error,omitempty"`                                     //TODO: Write general description for this field
	FullViewNameDEPRECATED  *string                                         `json:"fullViewName_DEPRECATED,omitempty" form:"fullViewName_DEPRECATED,omitempty"` //The full view name (external). This is composed of a Cohesity specific
	JobUid                  *UniversalIdProto                               `json:"jobUid,omitempty" form:"jobUid,omitempty"`                                   //TODO: Write general description for this field
	Name                    *string                                         `json:"name,omitempty" form:"name,omitempty"`                                       //The name of the retrieval task.
	ProgressMonitorTaskPath *string                                         `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"` //The path of the progress monitor for this task.
	RestoreTaskId           *int64                                          `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"`                     //For retrieve tasks created after the 2.8 release, this will contain the id
	RetrievalInfo           *RetrieveArchiveInfo                            `json:"retrievalInfo,omitempty" form:"retrievalInfo,omitempty"`                     //Proto to describe information about the retrieval of an archive task as
	StartTimeUsecs          *int64                                          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"`                   //The start time for this retrieval task.
	Status                  *int64                                          `json:"status,omitempty" form:"status,omitempty"`                                   //The status of this task.
	TaskUid                 *UniversalIdProto                               `json:"taskUid,omitempty" form:"taskUid,omitempty"`                                 //TODO: Write general description for this field
	User                    *string                                         `json:"user,omitempty" form:"user,omitempty"`                                       //The user who requested this retrieval task.
	VaultRestoreParams      *VaultParamsRestoreParams                       `json:"vaultRestoreParams,omitempty" form:"vaultRestoreParams,omitempty"`           //TODO: Write general description for this field
	ViewBoxId               *int64                                          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                             //The view box id to which 'view_name' belongs to.
	ViewNameDEPRECATED      *string                                         `json:"viewName_DEPRECATED,omitempty" form:"viewName_DEPRECATED,omitempty"`         //The view name as provided by the user for this retrieval task. Retrieved
}

/*
 * Structure for the custom type RetrieveArchiveTaskStateProtoDownloadFilesInfo
 */
type RetrieveArchiveTaskStateProtoDownloadFilesInfo struct {
	FilePath              *string            `json:"filePath,omitempty" form:"filePath,omitempty"`                           //The file to download from the archive.
	MagnetoInstanceId     *MagnetoInstanceId `json:"magnetoInstanceId,omitempty" form:"magnetoInstanceId,omitempty"`         //TODO: Write general description for this field
	ObjectId              *MagnetoObjectId   `json:"objectId,omitempty" form:"objectId,omitempty"`                           //TODO(apurv): This message type should be moved to the Yoda namespace.
	SkipRestoreInStubView *bool              `json:"skipRestoreInStubView,omitempty" form:"skipRestoreInStubView,omitempty"` //Ask Icebox to only create the stub view and skip restoring files in
	TargetDirPath         *string            `json:"targetDirPath,omitempty" form:"targetDirPath,omitempty"`                 //Path to the directory under which the downloaded files will be placed.
	TargetViewName        *string            `json:"targetViewName,omitempty" form:"targetViewName,omitempty"`               //Target view name where the downloaded files will be placed.
}

/*
 * Structure for the custom type Role
 */
type Role struct {
	CreatedTimeMsecs     *int64    `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the role was created.
	Description          *string   `json:"description,omitempty" form:"description,omitempty"`                   //Specifies a description about the role.
	IsCustomRole         *bool     `json:"isCustomRole,omitempty" form:"isCustomRole,omitempty"`                 //Specifies if the role is a user-defined custom role.
	Label                *string   `json:"label,omitempty" form:"label,omitempty"`                               //Specifies the label for the role as displayed on the Cohesity
	LastUpdatedTimeMsecs *int64    `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the role was last modified.
	Name                 *string   `json:"name,omitempty" form:"name,omitempty"`                                 //Specifies the internal Cluster name for the role such as COHESITY_VIEWER.
	Privileges           *[]string `json:"privileges,omitempty" form:"privileges,omitempty"`                     //Array of Privileges.
	TenantId             *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                         //Specifies unique id of the tenant owning the role.
	TenantIds            *[]string `json:"tenantIds,omitempty" form:"tenantIds,omitempty"`                       //Specifies id of tenants using this role.
}

/*
 * Structure for the custom type RoleCreateParameters
 */
type RoleCreateParameters struct {
	Description *string   `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the role.
	Name        *string   `json:"name,omitempty" form:"name,omitempty"`               //Specifies the name of the custom role.
	Privileges  *[]string `json:"privileges,omitempty" form:"privileges,omitempty"`   //Array of Privileges.
}

/*
 * Structure for the custom type RoleDeleteParameters
 */
type RoleDeleteParameters struct {
	Names []string `json:"names" form:"names"` //Array of Role Names.
}

/*
 * Structure for the custom type RoleUpdateParameters
 */
type RoleUpdateParameters struct {
	Description *string   `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the role.
	Privileges  *[]string `json:"privileges,omitempty" form:"privileges,omitempty"`   //Array of Privileges.
}

/*
 * Structure for the custom type Route
 */
type Route struct {
	Description    *string `json:"description,omitempty" form:"description,omitempty"`       //Specifies a description of the Static Route.
	DestNetwork    *string `json:"destNetwork,omitempty" form:"destNetwork,omitempty"`       //Destination network.
	IfName         *string `json:"ifName,omitempty" form:"ifName,omitempty"`                 //Specifies the network interfaces name to use for communicating with the
	IfaceGroupName *string `json:"ifaceGroupName,omitempty" form:"ifaceGroupName,omitempty"` //Specifies the network interfaces group or interface group with vlan id to
	NextHop        *string `json:"nextHop,omitempty" form:"nextHop,omitempty"`               //Specifies the next hop to the destination network.
}

/*
 * Structure for the custom type RpoPolicySettings
 */
type RpoPolicySettings struct {
	AlertingConfig           *AlertingConfig               `json:"alertingConfig,omitempty" form:"alertingConfig,omitempty"`                     //Specifies optional settings for alerting.
	AlertingPolicy           *[]AlertingPolicyEnum         `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"`                     //Array of Job Events.
	EnvironmentTypeJobParams *EnvironmentTypeJobParameters `json:"environmentTypeJobParams,omitempty" form:"environmentTypeJobParams,omitempty"` //Specifies additional parameters that are common to all Protection
	IndexingPolicy           *IndexingPolicy               `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"`                     //Specifies settings for indexing files found in an Object
	QosType                  QosTypeRpoPolicySettingsEnum  `json:"qosType,omitempty" form:"qosType,omitempty"`                                   //Specifies the QoS policy type to use.
	StorageDomainId          *int64                        `json:"storageDomainId,omitempty" form:"storageDomainId,omitempty"`                   //Specifies the Storage Domain to which data will be written.
}

/*
 * Structure for the custom type RpoSchedule
 */
type RpoSchedule struct {
	IntervalUnit IntervalUnitEnum `json:"intervalUnit,omitempty" form:"intervalUnit,omitempty"` //Specifies an RPO policy interval unit which will be used along with the
	Multiplier   *int64           `json:"multiplier,omitempty" form:"multiplier,omitempty"`     //Specifies the multiplier value to be used with the  RPO interval unit
}

/*
 * Structure for the custom type RunDiagnosticsMessage
 */
type RunDiagnosticsMessage struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies the status message returned after initiating a run
}

/*
 * Structure for the custom type RunJobSnapshotTarget
 */
type RunJobSnapshotTarget struct {
	ArchivalTarget      *ArchivalExternalTarget      `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`           //Specifies settings about the Archival External Target (such as Tape or AWS).
	DaysToKeep          *int64                       `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`                   //Specifies the number of days to retain copied Snapshots on the target.
	HoldForLegalPurpose *bool                        `json:"holdForLegalPurpose,omitempty" form:"holdForLegalPurpose,omitempty"` //Specifies optionally whether to retain the snapshot for legal purpose.
	ReplicationTarget   *ReplicationTargetSettings   `json:"replicationTarget,omitempty" form:"replicationTarget,omitempty"`     //Specifies settings about the Remote Cohesity Cluster where Snapshots
	Type                TypeRunJobSnapshotTargetEnum `json:"type,omitempty" form:"type,omitempty"`                               //Specifies the type of a Snapshot target such as 'kLocal', 'kRemote' or
}

/*
 * Structure for the custom type RunNowParameters
 */
type RunNowParameters struct {
	DatabaseIds *[]int64 `json:"databaseIds,omitempty" form:"databaseIds,omitempty"` //Specifies the ids of the DB's to perform run now on.
	SourceId    *int64   `json:"sourceId,omitempty" form:"sourceId,omitempty"`       //Specifies the source id of the Databases to perform the Run Now
}

/*
 * Structure for the custom type RunProtectionJobParam
 */
type RunProtectionJobParam struct {
	CopyRunTargets   []*RunJobSnapshotTarget          `json:"copyRunTargets,omitempty" form:"copyRunTargets,omitempty"`     //Optional parameter to be set if you want specific replication or archival
	RunNowParameters []*RunNowParameters              `json:"runNowParameters,omitempty" form:"runNowParameters,omitempty"` //Optional parameters of a Run Now operation.
	RunType          RunTypeRunProtectionJobParamEnum `json:"runType,omitempty" form:"runType,omitempty"`                   //Specifies the type of backup. If not specified, 'kRegular' is assumed.
	SourceIds        *[]int64                         `json:"sourceIds,omitempty" form:"sourceIds,omitempty"`               //Optional parameter if you want to back up only a subset of sources that
}

/*
 * Structure for the custom type RunUid
 */
type RunUid struct {
	JobUid         *UniversalId `json:"jobUid,omitempty" form:"jobUid,omitempty"`                 //Specifies an id for an object that is unique across Cohesity Clusters.
	StartTimeUsecs *int64       `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the Protection Job Run.
}

/*
 * Structure for the custom type SQLServerInstanceVersion
 */
type SQLServerInstanceVersion struct {
	Build         *int64  `json:"build,omitempty" form:"build,omitempty"`                 //Specfies the build.
	MajorVersion  *int64  `json:"majorVersion,omitempty" form:"majorVersion,omitempty"`   //Specfies the major version.
	MinorVersion  *int64  `json:"minorVersion,omitempty" form:"minorVersion,omitempty"`   //Specfies the minor version.
	Revision      *int64  `json:"revision,omitempty" form:"revision,omitempty"`           //Specfies the revision.
	VersionString *string `json:"versionString,omitempty" form:"versionString,omitempty"` //Specfies the version string.
}

/*
 * Structure for the custom type SalesforceAccountInfo
 */
type SalesforceAccountInfo struct {
	AccountId *string `json:"accountId,omitempty" form:"accountId,omitempty"` //Specifies the Account Id assigned by Salesforce.
	UserId    *string `json:"userId,omitempty" form:"userId,omitempty"`       //Specifies the User Id assigned by Salesforce.
}

/*
 * Structure for the custom type Sample
 */
type Sample struct {
	FloatValue     *float64 `json:"floatValue,omitempty" form:"floatValue,omitempty"`         //Specifies the value of the data sample if the type is float64.
	TimestampMsecs *int64   `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Specifies the timestamp when the data sample occured.
	Value          *int64   `json:"value,omitempty" form:"value,omitempty"`                   //Specifies the value of the data sample if the type is int64.
}

/*
 * Structure for the custom type SchedulerProtoSchedulerJobSchedule
 */
type SchedulerProtoSchedulerJobSchedule struct {
	Day      *int64  `json:"day,omitempty" form:"day,omitempty"`           //The day of the week when schedule should be executed (0-6).
	Hour     *int64  `json:"hour,omitempty" form:"hour,omitempty"`         //The hour of the day when schedule should be executed (0-23).
	Timezone *string `json:"timezone,omitempty" form:"timezone,omitempty"` //The timezone for the execution of the schedule.
}

/*
 * Structure for the custom type SchedulingPolicy
 */
type SchedulingPolicy struct {
	ContinuousSchedule *ContinuousSchedule `json:"continuousSchedule,omitempty" form:"continuousSchedule,omitempty"` //Specifies the time interval between two Job Runs of a continuous backup
	DailySchedule      *DailySchedule      `json:"dailySchedule,omitempty" form:"dailySchedule,omitempty"`           //Specifies a daily or weekly backup schedule.
	MonthlySchedule    *MonthlySchedule    `json:"monthlySchedule,omitempty" form:"monthlySchedule,omitempty"`       //Specifies a monthly backup schedule.
	Periodicity        PeriodicityEnum     `json:"periodicity,omitempty" form:"periodicity,omitempty"`               //Specifies how often to start new Job Runs of a Protection Job.
	RpoSchedule        *RpoSchedule        `json:"rpoSchedule,omitempty" form:"rpoSchedule,omitempty"`               //Specifies an RPO backup schedule.
}

/*
 * Structure for the custom type SchedulingPolicyProto
 */
type SchedulingPolicyProto struct {
	ContinuousSchedule *SchedulingPolicyProtoContinuousSchedule `json:"continuousSchedule,omitempty" form:"continuousSchedule,omitempty"` //TODO: Write general description for this field
	DailySchedule      *SchedulingPolicyProtoDailySchedule      `json:"dailySchedule,omitempty" form:"dailySchedule,omitempty"`           //The daily schedule encompasses weekly schedules as well. This has been
	MonthlySchedule    *SchedulingPolicyProtoMonthlySchedule    `json:"monthlySchedule,omitempty" form:"monthlySchedule,omitempty"`       //TODO: Write general description for this field
	Periodicity        *int64                                   `json:"periodicity,omitempty" form:"periodicity,omitempty"`               //Determines how often the job should be run.
	RpoSchedule        *SchedulingPolicyProtoRPOSchedule        `json:"rpoSchedule,omitempty" form:"rpoSchedule,omitempty"`               //TODO: Write general description for this field
}

/*
 * Structure for the custom type SchedulingPolicyProtoContinuousSchedule
 */
type SchedulingPolicyProtoContinuousSchedule struct {
	BackupIntervalMins *int64 `json:"backupIntervalMins,omitempty" form:"backupIntervalMins,omitempty"` //If this field is set, backups will be performed periodically every
}

/*
 * Structure for the custom type SchedulingPolicyProtoDailySchedule
 */
type SchedulingPolicyProtoDailySchedule struct {
	Days *[]int64 `json:"days,omitempty" form:"days,omitempty"` //The days of the week backup must be performed. If no days are specified,
}

/*
 * Structure for the custom type SchedulingPolicyProtoMonthlySchedule
 */
type SchedulingPolicyProtoMonthlySchedule struct {
	Count *int64 `json:"count,omitempty" form:"count,omitempty"` //Count of the day on which to perform the backup (look above for a more
	Day   *int64 `json:"day,omitempty" form:"day,omitempty"`     //The day of the month the backup is to be performed.
}

/*
 * Structure for the custom type SchedulingPolicyProtoRPOSchedule
 */
type SchedulingPolicyProtoRPOSchedule struct {
	RpoIntervalMins *int64 `json:"rpoIntervalMins,omitempty" form:"rpoIntervalMins,omitempty"` //If this field is set, then at any point, a recovery point should be
}

/*
 * Structure for the custom type SchemaInfo
 */
type SchemaInfo struct {
	EntityId   *string `json:"entityId,omitempty" form:"entityId,omitempty"`     //Specifies the id of the entity represented as a string.
	Key        *string `json:"key,omitempty" form:"key,omitempty"`               //Specifies the key which is public facing name for metric name.
	MetricName *string `json:"metricName,omitempty" form:"metricName,omitempty"` //Specifies the Apollo schema metric name.
	SchemaName *string `json:"schemaName,omitempty" form:"schemaName,omitempty"` //Specifies the name of entity schema such as 'ApolloViewBoxStats'.
}

/*
 * Structure for the custom type ScriptExecutionStatus
 */
type ScriptExecutionStatus struct {
	Error     *ErrorProto `json:"error,omitempty" form:"error,omitempty"`         //TODO: Write general description for this field
	Executing *bool       `json:"executing,omitempty" form:"executing,omitempty"` //Indicates if a script is executing. This is particularly useful when
	ExitCode  *int64      `json:"exitCode,omitempty" form:"exitCode,omitempty"`   //Exit code of the script.
	State     *int64      `json:"state,omitempty" form:"state,omitempty"`         //Execution state of the script.
}

/*
 * Structure for the custom type ScriptPathAndParams
 */
type ScriptPathAndParams struct {
	ContinueOnError *bool   `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Applicable only for pre backup scripts. If this flag is set to true, then
	IsActive        *bool   `json:"isActive,omitempty" form:"isActive,omitempty"`               //Indicates if the script is active. If 'is_active' is set to false, this
	ScriptParams    *string `json:"scriptParams,omitempty" form:"scriptParams,omitempty"`       //Custom parameters that users want to pass to the script. For example,
	ScriptPath      *string `json:"scriptPath,omitempty" form:"scriptPath,omitempty"`           //For backup jobs of type 'kPuppeteer', 'script_path' is full path of
	TimeoutSecs     *int64  `json:"timeoutSecs,omitempty" form:"timeoutSecs,omitempty"`         //Timeout of the script. The script will be killed if it exceeds this value.
}

/*
 * Structure for the custom type SendSnmpTestTrapResult
 */
type SendSnmpTestTrapResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ServiceStateResult
 */
type ServiceStateResult struct {
	Service ServiceServiceStateResultEnum `json:"service,omitempty" form:"service,omitempty"` //Specifies the name of the service.
	State   StateServiceStateResultEnum   `json:"state,omitempty" form:"state,omitempty"`     //Specifies the state of the service.
}

/*
 * Structure for the custom type SetupRestoreDiskTaskInfoProto
 */
type SetupRestoreDiskTaskInfoProto struct {
	Entity                      *EntityProto `json:"entity,omitempty" form:"entity,omitempty"`                                           //Specifies the attributes and the latest statistics about an entity.
	ProgressMonitorRootTaskPath *string      `json:"progressMonitorRootTaskPath,omitempty" form:"progressMonitorRootTaskPath,omitempty"` //The path to the progress monitor root task if any.
	RootEntity                  *EntityProto `json:"rootEntity,omitempty" form:"rootEntity,omitempty"`                                   //Specifies the attributes and the latest statistics about an entity.
	SourceViewName              *string      `json:"sourceViewName,omitempty" form:"sourceViewName,omitempty"`                           //The source view which contains the backups for the 'entity'.
	TaskId                      *int64       `json:"taskId,omitempty" form:"taskId,omitempty"`                                           //The id of the associated task.
	ViewBoxId                   *int64       `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                                     //The view box id containing the backups for the 'entity'.
	ViewName                    *string      `json:"viewName,omitempty" form:"viewName,omitempty"`                                       //Destination view into which the files will be cloned.
}

/*
 * Structure for the custom type Share
 */
type Share struct {
	AllSmbMountPaths       *[]string        `json:"allSmbMountPaths,omitempty" form:"allSmbMountPaths,omitempty"`             //Array of SMB Paths.
	EnableSmbEncryption    *bool            `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`       //Specifies the SMB encryption for the View Alias. If set, it enables the
	EnableSmbViewDiscovery *bool            `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view alias for SMB.
	EnforceSmbEncryption   *bool            `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`     //Specifies the SMB encryption for all the sessions for the View Alias.
	NfsMountPath           *string          `json:"nfsMountPath,omitempty" form:"nfsMountPath,omitempty"`                     //Specifies the path for mounting this Share as an NFS share.
	Path                   *string          `json:"path,omitempty" form:"path,omitempty"`                                     //Specifies the path information for this share.
	S3AccessPath           *string          `json:"s3AccessPath,omitempty" form:"s3AccessPath,omitempty"`                     //Specifies the path to access this View as an S3 share.
	ShareName              *string          `json:"shareName,omitempty" form:"shareName,omitempty"`                           //The name of the share.
	SharePermissions       []*SmbPermission `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`             //Specifies a list of share level permissions.
	SmbMountPath           *string          `json:"smbMountPath,omitempty" form:"smbMountPath,omitempty"`                     //Specifies the main path for mounting this Share as an SMB share.
	SubnetWhitelist        []*Subnet        `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`               //Specifies a list of Subnets with IP addresses that have permissions to
	TenantId               *string          `json:"tenantId,omitempty" form:"tenantId,omitempty"`                             //Specifies the unique id of the tenant.
	ViewName               *string          `json:"viewName,omitempty" form:"viewName,omitempty"`                             //Specifies the view name this share belongs to.
}

/*
 * Structure for the custom type ShowSystemLedInfoParameters
 */
type ShowSystemLedInfoParameters struct {
	NodeIp *string `json:"nodeIp,omitempty" form:"nodeIp,omitempty"` //Optional field.
}

/*
 * Structure for the custom type ShowSystemLedInfoResult
 */
type ShowSystemLedInfoResult struct {
	LedInfo *string `json:"ledInfo,omitempty" form:"ledInfo,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type SmbActiveFileOpensResponse
 */
type SmbActiveFileOpensResponse struct {
	ActiveFilePaths []*SmbActiveFilePath `json:"activeFilePaths,omitempty" form:"activeFilePaths,omitempty"` //Specifies the active opens for an SMB file in a view.
	Cookie          *string              `json:"cookie,omitempty" form:"cookie,omitempty"`                   //Specifies an opaque string to pass to get the next set of active opens.
}

/*
 * Structure for the custom type SmbActiveFilePath
 */
type SmbActiveFilePath struct {
	ActiveSessions []*SmbActiveSession `json:"activeSessions,omitempty" form:"activeSessions,omitempty"` //Specifies the sessions where the file is open.
	FilePath       *string             `json:"filePath,omitempty" form:"filePath,omitempty"`             //Specifies the filepath in the view.
	ViewId         *int64              `json:"viewId,omitempty" form:"viewId,omitempty"`                 //Specifies the id of the View assigned by the Cohesity Cluster.
	ViewName       *string             `json:"viewName,omitempty" form:"viewName,omitempty"`             //Specifies the name of the View.
}

/*
 * Structure for the custom type SmbActiveOpen
 */
type SmbActiveOpen struct {
	AccessInfoList  *[]AccessInfoListEnum `json:"accessInfoList,omitempty" form:"accessInfoList,omitempty"`   //Specifies the access information.
	OpenId          *int64                `json:"openId,omitempty" form:"openId,omitempty"`                   //Specifies the id of the active open.
	OthersCanDelete *bool                 `json:"othersCanDelete,omitempty" form:"othersCanDelete,omitempty"` //Specifies whether others are allowed to delete.
	OthersCanRead   *bool                 `json:"othersCanRead,omitempty" form:"othersCanRead,omitempty"`     //Specifies whether others are allowed to read.
	OthersCanWrite  *bool                 `json:"othersCanWrite,omitempty" form:"othersCanWrite,omitempty"`   //Specifies whether others are allowed to write.
}

/*
 * Structure for the custom type SmbActiveSession
 */
type SmbActiveSession struct {
	ActiveOpens []*SmbActiveOpen `json:"activeOpens,omitempty" form:"activeOpens,omitempty"` //Specifies the list of active opens of the file in this session.
	ClientIp    *string          `json:"clientIp,omitempty" form:"clientIp,omitempty"`       //Specifies the IP address from which the file is still open.
	Domain      *string          `json:"domain,omitempty" form:"domain,omitempty"`           //Specifies the domain of the user.
	ServerIp    *string          `json:"serverIp,omitempty" form:"serverIp,omitempty"`       //Specifies the IP address of the server where the file exists.
	SessionId   *int64           `json:"sessionId,omitempty" form:"sessionId,omitempty"`     //Specifies the id of the session.
	Username    *string          `json:"username,omitempty" form:"username,omitempty"`       //Specifies the username who keeps the file open.
}

/*
 * Structure for the custom type SmbConnection
 */
type SmbConnection struct {
	ClientIp   *string   `json:"clientIp,omitempty" form:"clientIp,omitempty"`     //Specifies the Client IP address of the connection.
	DomainName *string   `json:"domainName,omitempty" form:"domainName,omitempty"` //Domain name of the corresponding user.
	NodeIp     *string   `json:"nodeIp,omitempty" form:"nodeIp,omitempty"`         //Specifies a Node IP address where the connection request is received.
	Path       *string   `json:"path,omitempty" form:"path,omitempty"`             //Mount path.
	ServerIp   *string   `json:"serverIp,omitempty" form:"serverIp,omitempty"`     //Specifies the Server IP address of the connection.
	SessionId  *int64    `json:"sessionId,omitempty" form:"sessionId,omitempty"`   //Session id.
	Sids       *[]string `json:"sids,omitempty" form:"sids,omitempty"`             //List of SIDs in the SMB session token.
	UserName   *string   `json:"userName,omitempty" form:"userName,omitempty"`     //User name used to login for this session.
	ViewId     *int64    `json:"viewId,omitempty" form:"viewId,omitempty"`         //Specifies the id of the view.
	ViewName   *string   `json:"viewName,omitempty" form:"viewName,omitempty"`     //Specifies the name of the view.
}

/*
 * Structure for the custom type SmbPermission
 */
type SmbPermission struct {
	Access            AccessEnum            `json:"access,omitempty" form:"access,omitempty"`                       //Specifies the read/write access to the SMB share.
	Mode              ModeEnum              `json:"mode,omitempty" form:"mode,omitempty"`                           //Specifies how the permission should be applied to folders and/or files.
	Sid               *string               `json:"sid,omitempty" form:"sid,omitempty"`                             //Specifies the security identifier (SID) of the principal.
	SpecialAccessMask *int64                `json:"specialAccessMask,omitempty" form:"specialAccessMask,omitempty"` //Specifies custom access permissions.
	SpecialType       *int64                `json:"specialType,omitempty" form:"specialType,omitempty"`             //Specifies a custom type.
	Type              TypeSmbPermissionEnum `json:"type,omitempty" form:"type,omitempty"`                           //Specifies the type of permission.
}

/*
 * Structure for the custom type SmbPermissionsInfo
 */
type SmbPermissionsInfo struct {
	OwnerSid    *string          `json:"ownerSid,omitempty" form:"ownerSid,omitempty"`       //Specifies the security identifier (SID) of the owner of the SMB share.
	Permissions []*SmbPermission `json:"permissions,omitempty" form:"permissions,omitempty"` //Array of SMB Permissions.
}

/*
 * Structure for the custom type SmbPrincipal
 */
type SmbPrincipal struct {
	Domain *string `json:"domain,omitempty" form:"domain,omitempty"` //Specifies domain name of the principal.
	Name   *string `json:"name,omitempty" form:"name,omitempty"`     //Specifies name of the SMB principal which may be a group or user.
	Sid    *string `json:"sid,omitempty" form:"sid,omitempty"`       //Specifies unique Security ID (SID) of the principal that look similar to
	Type   *string `json:"type,omitempty" form:"type,omitempty"`     //Specifies the type. This can be a user or a group.
}

/*
 * Structure for the custom type SnapshotArchivalCopyPolicy
 */
type SnapshotArchivalCopyPolicy struct {
	CopyPartial *bool                                     `json:"copyPartial,omitempty" form:"copyPartial,omitempty"` //Specifies if Snapshots are copied from the first completely successful
	DaysToKeep  *int64                                    `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`   //Specifies the number of days to retain copied Snapshots on the target.
	Multiplier  *int64                                    `json:"multiplier,omitempty" form:"multiplier,omitempty"`   //Specifies a factor to multiply the periodicity by, to determine the copy
	Periodicity PeriodicitySnapshotArchivalCopyPolicyEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
	Target      *ArchivalExternalTarget                   `json:"target,omitempty" form:"target,omitempty"`           //Specifies the archival target to copy the Snapshots to.
}

/*
 * Structure for the custom type SnapshotAttempt
 */
type SnapshotAttempt struct {
	AttemptNumber    *int64 `json:"attemptNumber,omitempty" form:"attemptNumber,omitempty"`       //Specifies the number of the attempts made by the Job Run
	JobRunId         *int64 `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                 //Specifies the id of the Job Run that captured the snapshot.
	StartedTimeUsecs *int64 `json:"startedTimeUsecs,omitempty" form:"startedTimeUsecs,omitempty"` //Specifies the time when the Job Run starts capturing a snapshot.
}

/*
 * Structure for the custom type SnapshotCloudCopyPolicy
 */
type SnapshotCloudCopyPolicy struct {
	CopyPartial *bool                                  `json:"copyPartial,omitempty" form:"copyPartial,omitempty"` //Specifies if Snapshots are copied from the first completely successful
	DaysToKeep  *int64                                 `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`   //Specifies the number of days to retain copied Snapshots on the target.
	Multiplier  *int64                                 `json:"multiplier,omitempty" form:"multiplier,omitempty"`   //Specifies a factor to multiply the periodicity by, to determine the copy
	Periodicity PeriodicitySnapshotCloudCopyPolicyEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
	Target      *CloudDeployTargetDetails              `json:"target,omitempty" form:"target,omitempty"`           //Message that specifies the details about CloudDeploy target where backup
}

/*
 * Structure for the custom type SnapshotCopyTask
 */
type SnapshotCopyTask struct {
	CopyStatus      *string                 `json:"copyStatus,omitempty" form:"copyStatus,omitempty"`           //Specifies the status of the copy task.
	ExpiryTimeUsecs *int64                  `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies when the Snapshot expires on the target.
	Message         *string                 `json:"message,omitempty" form:"message,omitempty"`                 //Specifies warning or error information when the copy task is not
	SnapshotTarget  *SnapshotTargetSettings `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"`   //Specifies settings about a target where a copied Snapshot is stored.
}

/*
 * Structure for the custom type SnapshotInfo
 */
type SnapshotInfo struct {
	Environment               EnvironmentSnapshotInfoEnum `json:"environment,omitempty" form:"environment,omitempty"`                             //Specifies the environment type (such as kVMware or kSQL) that
	RelativeSnapshotDirectory *string                     `json:"relativeSnapshotDirectory,omitempty" form:"relativeSnapshotDirectory,omitempty"` //Specifies the relative directory path from root path where the snapshot
	RootPath                  *string                     `json:"rootPath,omitempty" form:"rootPath,omitempty"`                                   //Specifies the root path where the snapshot is stored, using the following
	ViewName                  *string                     `json:"viewName,omitempty" form:"viewName,omitempty"`                                   //Specifies the name of the View that is cloned.
}

/*
 * Structure for the custom type SnapshotInfoProto
 */
type SnapshotInfoProto struct {
	ErrorRocksdbName              *string                        `json:"errorRocksdbName,omitempty" form:"errorRocksdbName,omitempty"`                           //The name of the rocksdb directory for errors seen during this backup,
	FileWalkDone                  *bool                          `json:"fileWalkDone,omitempty" form:"fileWalkDone,omitempty"`                                   //This field is only applicable for NAS and file backup jobs. It indicates
	NumAppInstances               *int64                         `json:"numAppInstances,omitempty" form:"numAppInstances,omitempty"`                             //Number of application instances backed up by this task. For example, if
	NumAppObjects                 *int64                         `json:"numAppObjects,omitempty" form:"numAppObjects,omitempty"`                                 //Number of application objects in total backed up by this task. For
	PostBackupScriptStatus        *ScriptExecutionStatus         `json:"postBackupScriptStatus,omitempty" form:"postBackupScriptStatus,omitempty"`               //TODO: Write general description for this field
	PreBackupScriptStatus         *ScriptExecutionStatus         `json:"preBackupScriptStatus,omitempty" form:"preBackupScriptStatus,omitempty"`                 //TODO: Write general description for this field
	RelativeSnapshotDir           *string                        `json:"relativeSnapshotDir,omitempty" form:"relativeSnapshotDir,omitempty"`                     //This is the path relative to 'root_path' under which the snapshot lives.
	RootPath                      *string                        `json:"rootPath,omitempty" form:"rootPath,omitempty"`                                           //The root path under which the snapshot is stored. This is of the form
	ScribeTableColumn             *string                        `json:"scribeTableColumn,omitempty" form:"scribeTableColumn,omitempty"`                         //If this backup task stores any auxiliary state in Scribe table, this field
	ScribeTableRow                *string                        `json:"scribeTableRow,omitempty" form:"scribeTableRow,omitempty"`                               //If this backup task stores any auxiliary state in Scribe table, this field
	SlaveTaskStartTimeUsecs       *int64                         `json:"slaveTaskStartTimeUsecs,omitempty" form:"slaveTaskStartTimeUsecs,omitempty"`             //This is the timestamp at which the slave task started.
	SnapshotType                  *ObjectSnapshotType            `json:"snapshotType,omitempty" form:"snapshotType,omitempty"`                                   //TODO: Write general description for this field
	StorageSnapshotProvider       *StorageSnapshotProviderParams `json:"storageSnapshotProvider,omitempty" form:"storageSnapshotProvider,omitempty"`             //TODO: Write general description for this field
	TargetType                    *int64                         `json:"targetType,omitempty" form:"targetType,omitempty"`                                       //Specifies the target type for the task. The field is only valid if the
	TotalBytesReadFromSource      *int64                         `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"`           //Contains the information regarding number of bytes that are read from the
	TotalBytesToReadFromSource    *int64                         `json:"totalBytesToReadFromSource,omitempty" form:"totalBytesToReadFromSource,omitempty"`       //Contains the total number of bytes that will be read from the source
	TotalChangedEntityCount       *int64                         `json:"totalChangedEntityCount,omitempty" form:"totalChangedEntityCount,omitempty"`             //The total number of file and directory entities that have changed since
	TotalEntityCount              *int64                         `json:"totalEntityCount,omitempty" form:"totalEntityCount,omitempty"`                           //The total number of file and directory entities visited in this
	TotalLogicalBackupSizeBytes   *int64                         `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"`     //Logical size of the source whose snapshot is being taken. This is the
	TotalPrimaryPhysicalSizeBytes *int64                         `json:"totalPrimaryPhysicalSizeBytes,omitempty" form:"totalPrimaryPhysicalSizeBytes,omitempty"` //Contains the information regarding number of bytes that the source (such
	Type                          *int64                         `json:"type,omitempty" form:"type,omitempty"`                                                   //The type of environment this snapshot info pertains to.
	ViewCaseInsensitivityAltered  *bool                          `json:"viewCaseInsensitivityAltered,omitempty" form:"viewCaseInsensitivityAltered,omitempty"`   //Whether during the backup, the backup view's case insensitivity property
	ViewName                      *string                        `json:"viewName,omitempty" form:"viewName,omitempty"`                                           //The view name under which the snapshot was created.
	ViewNameToGc                  *string                        `json:"viewNameToGc,omitempty" form:"viewNameToGc,omitempty"`                                   //The view name under which the snapshot of the migrated data was created.
	Warnings                      []*ErrorProto                  `json:"warnings,omitempty" form:"warnings,omitempty"`                                           //Warnings if any. These warnings will be propogated to the UI by master.
}

/*
 * Structure for the custom type SnapshotManagerParams
 */
type SnapshotManagerParams struct {
	AwsSnapshotManagerParams *AWSSnapshotManagerParams `json:"awsSnapshotManagerParams,omitempty" form:"awsSnapshotManagerParams,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type SnapshotReplicationCopyPolicy
 */
type SnapshotReplicationCopyPolicy struct {
	CloudTarget *CloudDeployTargetDetails                    `json:"cloudTarget,omitempty" form:"cloudTarget,omitempty"` //Message that specifies the details about CloudDeploy target where backup
	CopyPartial *bool                                        `json:"copyPartial,omitempty" form:"copyPartial,omitempty"` //Specifies if Snapshots are copied from the first completely successful
	DaysToKeep  *int64                                       `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"`   //Specifies the number of days to retain copied Snapshots on the target.
	Multiplier  *int64                                       `json:"multiplier,omitempty" form:"multiplier,omitempty"`   //Specifies a factor to multiply the periodicity by, to determine the copy
	Periodicity PeriodicitySnapshotReplicationCopyPolicyEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
	Target      *ReplicationTargetSettings                   `json:"target,omitempty" form:"target,omitempty"`           //Specifies the replication target to copy the Snapshots to.
}

/*
 * Structure for the custom type SnapshotTarget
 */
type SnapshotTarget struct {
	ArchivalTarget    *ArchivalTarget    `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`       //Message that specifies the details about an archival target (such as cloud
	CloudDeployTarget *CloudDeployTarget `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"` //Message that specifies the details about CloudDeploy target where backup
	ReplicationTarget *ReplicationTarget `json:"replicationTarget,omitempty" form:"replicationTarget,omitempty"` //Message that specifies the details about a remote cluster where backup
	Type              *int64             `json:"type,omitempty" form:"type,omitempty"`                           //The type of snapshot target this proto represents.
}

/*
 * Structure for the custom type SnapshotTargetPolicyProto
 */
type SnapshotTargetPolicyProto struct {
	CopyPartiallySuccessfulRun *bool                 `json:"copyPartiallySuccessfulRun,omitempty" form:"copyPartiallySuccessfulRun,omitempty"` //If this is false, then only snapshots from the first completely successful
	GranularityBucket          *GranularityBucket    `json:"granularityBucket,omitempty" form:"granularityBucket,omitempty"`                   //Message that specifies the frequency granularity at which to copy the
	NumDaysToKeep              *int64                `json:"numDaysToKeep,omitempty" form:"numDaysToKeep,omitempty"`                           //Specifies how to determine the expiration time for snapshots copied due to
	RetentionPolicy            *RetentionPolicyProto `json:"retentionPolicy,omitempty" form:"retentionPolicy,omitempty"`                       //Message that specifies the retention policy for backup snapshots.
	SnapshotTarget             *SnapshotTarget       `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"`                         //Message that specifies details about a target (such as a replication or
}

/*
 * Structure for the custom type SnapshotTargetSettings
 */
type SnapshotTargetSettings struct {
	ArchivalTarget    *ArchivalExternalTarget        `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"`       //Specifies settings about the Archival External Target (such as Tape or AWS).
	ReplicationTarget *ReplicationTargetSettings     `json:"replicationTarget,omitempty" form:"replicationTarget,omitempty"` //Specifies settings about the Remote Cohesity Cluster where Snapshots
	Type              TypeSnapshotTargetSettingsEnum `json:"type,omitempty" form:"type,omitempty"`                           //Specifies the type of a Snapshot target such as 'kLocal', 'kRemote' or
}

/*
 * Structure for the custom type SnapshotVersion
 */
type SnapshotVersion struct {
	AttemptNumber            *int64  `json:"attemptNumber,omitempty" form:"attemptNumber,omitempty"`                       //Specifies the number of the attempts made by the Job Run
	DeltaSizeBytes           *int64  `json:"deltaSizeBytes,omitempty" form:"deltaSizeBytes,omitempty"`                     //Specifies the size of the data captured from the source object.
	IsAppConsistent          *bool   `json:"isAppConsistent,omitempty" form:"isAppConsistent,omitempty"`                   //Specifies if an app-consistent snapshot was captured. For example,
	IsFullBackup             *bool   `json:"isFullBackup,omitempty" form:"isFullBackup,omitempty"`                         //Specifies if the snapshot is a full backup. For example, all blocks
	JobRunId                 *int64  `json:"jobRunId,omitempty" form:"jobRunId,omitempty"`                                 //Specifies the id of the Job Run that captured the snapshot.
	LocalMountPath           *string `json:"localMountPath,omitempty" form:"localMountPath,omitempty"`                     //Specifies the local path relative to the View, without the
	LogicalSizeBytes         *int64  `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"`                 //Specifies the size of the snapshot if the data
	PhysicalSizeBytes        *int64  `json:"physicalSizeBytes,omitempty" form:"physicalSizeBytes,omitempty"`               //Specifies the amount of data actually used on the disk to store this
	PrimaryPhysicalSizeBytes *int64  `json:"primaryPhysicalSizeBytes,omitempty" form:"primaryPhysicalSizeBytes,omitempty"` //Specifies the total amount of disk space used to store this
	StartedTimeUsecs         *int64  `json:"startedTimeUsecs,omitempty" form:"startedTimeUsecs,omitempty"`                 //Specifies the time when the Job Run starts capturing a snapshot.
}

/*
 * Structure for the custom type SourceAppParams
 */
type SourceAppParams struct {
	IsVssCopyOnly    *bool             `json:"isVssCopyOnly,omitempty" form:"isVssCopyOnly,omitempty"`       //If the backup is a VSS full backup with the copy-only option specified.
	MsExchangeParams *MSExchangeParams `json:"msExchangeParams,omitempty" form:"msExchangeParams,omitempty"` //All the params specific to MS exchange application.
}

/*
 * Structure for the custom type SourceBackupStatus
 */
type SourceBackupStatus struct {
	CurrentSnapshotInfo     *SnapshotInfo                `json:"currentSnapshotInfo,omitempty" form:"currentSnapshotInfo,omitempty"`         //Specifies details about the snapshot task created to backup or copy one
	Error                   *string                      `json:"error,omitempty" form:"error,omitempty"`                                     //Specifies if an error occurred (if any) while running this task.
	IsFullBackup            *bool                        `json:"isFullBackup,omitempty" form:"isFullBackup,omitempty"`                       //Specifies whether this is a 'kFull' or 'kRegular' backup of the Run.
	NumRestarts             *int64                       `json:"numRestarts,omitempty" form:"numRestarts,omitempty"`                         //Specifies the number of times the the task was restarted because of the
	ParentSourceId          *int64                       `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"`                   //Specifies the id of the registered Protection Source that is the
	ProgressMonitorTaskPath *string                      `json:"progressMonitorTaskPath,omitempty" form:"progressMonitorTaskPath,omitempty"` //Specifies the yoda progress monitor task path which is used to get pulse
	Quiesced                *bool                        `json:"quiesced,omitempty" form:"quiesced,omitempty"`                               //Specifies if app-consistent snapshot was captured. This field is set to
	SlaViolated             *bool                        `json:"slaViolated,omitempty" form:"slaViolated,omitempty"`                         //Specifies if the SLA was violated for the Job Run. This field is set
	Source                  *ProtectionSource            `json:"source,omitempty" form:"source,omitempty"`                                   //Specifies a generic structure that represents a node
	Stats                   *BackupSourceStats           `json:"stats,omitempty" form:"stats,omitempty"`                                     //Specifies statistics about a Backup task in a Protection Job Run.
	Status                  StatusSourceBackupStatusEnum `json:"status,omitempty" form:"status,omitempty"`                                   //Specifies the status of the source object being protected.
	Warnings                *[]string                    `json:"warnings,omitempty" form:"warnings,omitempty"`                               //Array of Warnings.
}

/*
 * Structure for the custom type SourceForPrincipalParam
 */
type SourceForPrincipalParam struct {
	ProtectionSourceIds *[]int64  `json:"protectionSourceIds,omitempty" form:"protectionSourceIds,omitempty"` //Array of Protection Source Ids.
	Sid                 *string   `json:"sid,omitempty" form:"sid,omitempty"`                                 //Specifies the SID of the principal to grant access permissions to.
	ViewNames           *[]string `json:"viewNames,omitempty" form:"viewNames,omitempty"`                     //Array of View names.
}

/*
 * Structure for the custom type SourceSpecialParameter
 */
type SourceSpecialParameter struct {
	AdSpecialParameters       *ApplicationSpecialParameters `json:"adSpecialParameters,omitempty" form:"adSpecialParameters,omitempty"`             //Specifies additional special settings applicable for a Protection Source
	OracleSpecialParameters   *OracleSpecialParameters      `json:"oracleSpecialParameters,omitempty" form:"oracleSpecialParameters,omitempty"`     //Specifies special settings applicable for 'kOracle' environment.
	PhysicalSpecialParameters *PhysicalSpecialParameters    `json:"physicalSpecialParameters,omitempty" form:"physicalSpecialParameters,omitempty"` //Specifies additional special settings applicable for a Protection Source
	SkipIndexing              *bool                         `json:"skipIndexing,omitempty" form:"skipIndexing,omitempty"`                           //Specifies not to index the objects in the Protection Source when
	SourceId                  *int64                        `json:"sourceId,omitempty" form:"sourceId,omitempty"`                                   //Specifies the object id of the Protection Source that these
	SqlSpecialParameters      *ApplicationSpecialParameters `json:"sqlSpecialParameters,omitempty" form:"sqlSpecialParameters,omitempty"`           //Specifies additional special settings applicable for a Protection Source
	TruncateExchangeLog       *bool                         `json:"truncateExchangeLog,omitempty" form:"truncateExchangeLog,omitempty"`             //If true, after the Cohesity Cluster successfully captures a Snapshot
	VmCredentials             *Credentials                  `json:"vmCredentials,omitempty" form:"vmCredentials,omitempty"`                         //Specifies the administrator credentials to log in to the
	VmwareSpecialParameters   *VmwareSpecialParameters      `json:"vmwareSpecialParameters,omitempty" form:"vmwareSpecialParameters,omitempty"`     //Specifies additional special settings applicable for a Protection Source
}

/*
 * Structure for the custom type SourcesForSid
 */
type SourcesForSid struct {
	ProtectionSources []*ProtectionSource `json:"protectionSources,omitempty" form:"protectionSources,omitempty"` //Array of Protection Sources.
	Sid               *string             `json:"sid,omitempty" form:"sid,omitempty"`                             //Specifies the security identifier (SID) of the principal.
	Views             []*View             `json:"views,omitempty" form:"views,omitempty"`                         //Array of View Names.
}

/*
 * Structure for the custom type SqlAagHostAndDatabases
 */
type SqlAagHostAndDatabases struct {
	AagDatabases    []*AagAndDatabases    `json:"aagDatabases,omitempty" form:"aagDatabases,omitempty"`       //Specifies a list of AAGs and database members in each AAG.
	ApplicationNode *ProtectionSourceNode `json:"applicationNode,omitempty" form:"applicationNode,omitempty"` //Many different node types are supported such as
	Databases       []*ProtectionSource   `json:"databases,omitempty" form:"databases,omitempty"`             //Specifies all database entities found on the server. Database may or
	ErrorMessage    *string               `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`       //Specifies an error message when the host is not registered as an SQL
	UnknownHostName *string               `json:"unknownHostName,omitempty" form:"unknownHostName,omitempty"` //Specifies the name of the host that is not registered as an SQL server
}

/*
 * Structure for the custom type SqlBackupJobParams
 */
type SqlBackupJobParams struct {
	AagBackupPreferenceType        *int64 `json:"aagBackupPreferenceType,omitempty" form:"aagBackupPreferenceType,omitempty"`               //Preference type for backing up databases that are part of an AAG.
	BackupDatabaseVolumesOnly      *bool  `json:"backupDatabaseVolumesOnly,omitempty" form:"backupDatabaseVolumesOnly,omitempty"`           //If set to true, only the volumes associated with databases should be
	BackupSystemDbs                *bool  `json:"backupSystemDbs,omitempty" form:"backupSystemDbs,omitempty"`                               //Set to true if system databases should be backed up.
	FullBackupType                 *int64 `json:"fullBackupType,omitempty" form:"fullBackupType,omitempty"`                                 //The type of SQL full backup to be used for this job.
	IsCopyOnlyFull                 *bool  `json:"isCopyOnlyFull,omitempty" form:"isCopyOnlyFull,omitempty"`                                 //Whether full backups should be copy-only.
	UseAagPreferencesFromSqlServer *bool  `json:"useAagPreferencesFromSqlServer,omitempty" form:"useAagPreferencesFromSqlServer,omitempty"` //Set to true if we should use AAG preferences specified at the SQL server
	UserDbPreferenceType           *int64 `json:"userDbPreferenceType,omitempty" form:"userDbPreferenceType,omitempty"`                     //Preference type for backing up user databases on the host.
}

/*
 * Structure for the custom type SqlEnvJobParameters
 */
type SqlEnvJobParameters struct {
	AagPreference              AagPreferenceEnum                 `json:"aagPreference,omitempty" form:"aagPreference,omitempty"`                           //Specifies the preference for backing up databases that are part of an AAG.
	AagPreferenceFromSqlServer *bool                             `json:"aagPreferenceFromSqlServer,omitempty" form:"aagPreferenceFromSqlServer,omitempty"` //If true, AAG preferences are taken from the SQL server host. If this is
	BackupSystemDatabases      *bool                             `json:"backupSystemDatabases,omitempty" form:"backupSystemDatabases,omitempty"`           //If true, system databases are backed up. If this is set to false,
	BackupType                 BackupTypeSqlEnvJobParametersEnum `json:"backupType,omitempty" form:"backupType,omitempty"`                                 //Specifies the type of the 'kFull' backup job. Specifies whether it is
	BackupVolumesOnly          *bool                             `json:"backupVolumesOnly,omitempty" form:"backupVolumesOnly,omitempty"`                   //If set to true, only the volumes associated with databases should be
	IsCopyOnlyFull             *bool                             `json:"isCopyOnlyFull,omitempty" form:"isCopyOnlyFull,omitempty"`                         //If true, the backup is a full backup with the copy-only option specified.
	UserDatabasePreference     UserDatabasePreferenceEnum        `json:"userDatabasePreference,omitempty" form:"userDatabasePreference,omitempty"`         //Specifies the preference for backing up user databases on the host.
}

/*
 * Structure for the custom type SqlProtectionSource
 */
type SqlProtectionSource struct {
	IsAvailableForVssBackup  *bool                       `json:"IsAvailableForVssBackup,omitempty" form:"IsAvailableForVssBackup,omitempty"`   //Specifies whether the database is marked as available for backup according
	CreatedTimestamp         *string                     `json:"createdTimestamp,omitempty" form:"createdTimestamp,omitempty"`                 //Specifies the time when the database was created. It is displayed in the
	DatabaseName             *string                     `json:"databaseName,omitempty" form:"databaseName,omitempty"`                         //Specifies the database name of the SQL Protection Source, if the type
	DbAagEntityId            *int64                      `json:"dbAagEntityId,omitempty" form:"dbAagEntityId,omitempty"`                       //Specifies the AAG entity id if the database is part of an AAG.
	DbAagName                *string                     `json:"dbAagName,omitempty" form:"dbAagName,omitempty"`                               //Specifies the name of the AAG if the database is part of an AAG.
	DbCompatibilityLevel     *int64                      `json:"dbCompatibilityLevel,omitempty" form:"dbCompatibilityLevel,omitempty"`         //Specifies the versions of SQL server that the database is compatible
	DbFileGroups             *[]string                   `json:"dbFileGroups,omitempty" form:"dbFileGroups,omitempty"`                         //Specifies the information about the set of file groups for this db on
	DbFiles                  []*DbFileInfo               `json:"dbFiles,omitempty" form:"dbFiles,omitempty"`                                   //Specifies the last known information about the set of database files
	DbOwnerUsername          *string                     `json:"dbOwnerUsername,omitempty" form:"dbOwnerUsername,omitempty"`                   //Specifies the name of the database owner.
	Id                       *SqlSourceId                `json:"id,omitempty" form:"id,omitempty"`                                             //Specifies a unique id for a SQL Protection Source.
	Name                     *string                     `json:"name,omitempty" form:"name,omitempty"`                                         //Specifies the instance name of the SQL Protection Source
	OwnerId                  *int64                      `json:"ownerId,omitempty" form:"ownerId,omitempty"`                                   //Specifies the id of the container VM for the SQL Protection Source.
	RecoveryModel            RecoveryModelEnum           `json:"recoveryModel,omitempty" form:"recoveryModel,omitempty"`                       //Specifies the Recovery Model for the database in SQL environment.
	SqlServerDbState         SqlServerDbStateEnum        `json:"sqlServerDbState,omitempty" form:"sqlServerDbState,omitempty"`                 //The state of the database as returned by SQL Server.
	SqlServerInstanceVersion *SQLServerInstanceVersion   `json:"sqlServerInstanceVersion,omitempty" form:"sqlServerInstanceVersion,omitempty"` //Specifies the Server Instance Version.
	Type                     TypeSqlProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"`                                         //Specifies the type of the managed Object in a SQL Protection Source.
}

/*
 * Structure for the custom type SqlRestoreParameters
 */
type SqlRestoreParameters struct {
	CaptureTailLogs                       *bool                         `json:"captureTailLogs,omitempty" form:"captureTailLogs,omitempty"`                                             //Set this to true if tail logs are to be captured before the restore
	KeepOffline                           *bool                         `json:"keepOffline,omitempty" form:"keepOffline,omitempty"`                                                     //Set this to true if we want to restore the database and do not want to
	NewDatabaseName                       *string                       `json:"newDatabaseName,omitempty" form:"newDatabaseName,omitempty"`                                             //Specifies optionally a new name for the restored database.
	NewInstanceName                       *string                       `json:"newInstanceName,omitempty" form:"newInstanceName,omitempty"`                                             //Specifies an instance name of the SQL Server that should be restored.
	RestoreTimeSecs                       *int64                        `json:"restoreTimeSecs,omitempty" form:"restoreTimeSecs,omitempty"`                                             //Specifies the time in the past to which the SQL database needs to be
	TargetDataFilesDirectory              *string                       `json:"targetDataFilesDirectory,omitempty" form:"targetDataFilesDirectory,omitempty"`                           //Specifies the directory where to put the database data files.
	TargetLogFilesDirectory               *string                       `json:"targetLogFilesDirectory,omitempty" form:"targetLogFilesDirectory,omitempty"`                             //Specifies the directory where to put the database log files. Missing
	TargetSecondaryDataFilesDirectoryList []*FilenamePatternToDirectory `json:"targetSecondaryDataFilesDirectoryList,omitempty" form:"targetSecondaryDataFilesDirectoryList,omitempty"` //Specifies the secondary data filename pattern and corresponding
}

/*
 * Structure for the custom type SqlSourceId
 */
type SqlSourceId struct {
	CreatedDateMsecs *int64   `json:"createdDateMsecs,omitempty" form:"createdDateMsecs,omitempty"` //Specifies a unique identifier generated from the date the database is
	DatabaseId       *int64   `json:"databaseId,omitempty" form:"databaseId,omitempty"`             //Specifies a unique id of the database but only for the life of the
	InstanceId       *[]int64 `json:"instanceId,omitempty" form:"instanceId,omitempty"`             //Array of bytes that stores the SQL Server Instance id.
}

/*
 * Structure for the custom type SqlUpdateRestoreTaskOptions
 */
type SqlUpdateRestoreTaskOptions struct {
	MultiStageRestoreAction *int64 `json:"multiStageRestoreAction,omitempty" form:"multiStageRestoreAction,omitempty"` //This field is set if we are performing an action on a multi-stage SQL
}

/*
 * Structure for the custom type SslCertificateConfig
 */
type SslCertificateConfig struct {
	Certificate         *string `json:"certificate,omitempty" form:"certificate,omitempty"`                 //Certificate is a SSL certificate used by Iris HTTPS webserver.
	LastUpdateTimeMsecs *int64  `json:"lastUpdateTimeMsecs,omitempty" form:"lastUpdateTimeMsecs,omitempty"` //LastUpdateTimeMsecs is a time in milliseconds at which certificate was
	PrivateKey          *string `json:"privateKey,omitempty" form:"privateKey,omitempty"`                   //PrivateKey is a matching private key of the above certificate.
}

/*
 * Structure for the custom type SslVerification
 */
type SslVerification struct {
	CaCertificate *string `json:"caCertificate,omitempty" form:"caCertificate,omitempty"` //Contains the contents of CA cert/cert chain.
	IsEnabled     *bool   `json:"isEnabled,omitempty" form:"isEnabled,omitempty"`         //Whether SSL verification should be performed.
}

/*
 * Structure for the custom type StaticRoute
 */
type StaticRoute struct {
	Description           *string  `json:"description,omitempty" form:"description,omitempty"`                     //Specifies a description of the Static Route.
	IsUpdate              *bool    `json:"isUpdate,omitempty" form:"isUpdate,omitempty"`                           //Specifies if the route is currently being updated on the Cohesity Cluster.
	NetworkInterfaceGroup *string  `json:"networkInterfaceGroup,omitempty" form:"networkInterfaceGroup,omitempty"` //Specifies the group name of the network interfaces to
	NetworkInterfaceIds   *[]int64 `json:"networkInterfaceIds,omitempty" form:"networkInterfaceIds,omitempty"`     //Array of Network Interface Ids.
	Subnet                *Subnet  `json:"subnet,omitempty" form:"subnet,omitempty"`                               //Specifies the destination subnet of the Static Route.
	VlanId                *int64   `json:"vlanId,omitempty" form:"vlanId,omitempty"`                               //Specifies the ID of the VLAN to use for communication with the
}

/*
 * Structure for the custom type StatsGroup
 */
type StatsGroup struct {
	Consumer    *Consumer `json:"consumer,omitempty" form:"consumer,omitempty"`       //Consumer is the storage consumer of a group.
	EntityId    *string   `json:"entityId,omitempty" form:"entityId,omitempty"`       //Specifies the entity id of the group.
	Id          *int64    `json:"id,omitempty" form:"id,omitempty"`                   //Specifies the id of the group.
	Name        *string   `json:"name,omitempty" form:"name,omitempty"`               //Specifies the name of the group.
	TenantId    *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`       //Specifies the id of the organization (tenant) with respect to this group.
	TenantName  *string   `json:"tenantName,omitempty" form:"tenantName,omitempty"`   //Specifies the name of the organization (tenant) with respect to this
	ViewBoxId   *int64    `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`     //Specifies the id of the view box (storage domain) with respect to this
	ViewBoxName *string   `json:"viewBoxName,omitempty" form:"viewBoxName,omitempty"` //Specifies the name of the view box (storage domain) with respect to this
}

/*
 * Structure for the custom type StopRemoteVaultSearchJobParameters
 */
type StopRemoteVaultSearchJobParameters struct {
	SearchJobUid *UniversalId `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id of the Remote Vault search job in progress.
}

/*
 * Structure for the custom type StorageDomainStats
 */
type StorageDomainStats struct {
	CloudSpillVaultId   *int64             `json:"cloudSpillVaultId,omitempty" form:"cloudSpillVaultId,omitempty"`     //Specifies the cloud spill vault id of the view box (storage domain).
	GroupList           []*StatsGroup      `json:"groupList,omitempty" form:"groupList,omitempty"`                     //Specifies a list of groups associated to this view box (storage domain).
	Id                  *int64             `json:"id,omitempty" form:"id,omitempty"`                                   //Specifies the id of the view box (storage domain).
	Name                *string            `json:"name,omitempty" form:"name,omitempty"`                               //Specifies the name of the view box (storage domain).
	QuotaHardLimitBytes *int64             `json:"quotaHardLimitBytes,omitempty" form:"quotaHardLimitBytes,omitempty"` //Specifies the hard limit of physical quota of the view box
	SchemaInfoList      []*UsageSchemaInfo `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"`           //Specifies a list of schemaInfos of the view box (storage domain).
	Stats               *DataUsageStats    `json:"stats,omitempty" form:"stats,omitempty"`                             //Specifies the data usage metric of the data stored on the Cohesity
}

/*
 * Structure for the custom type StorageEfficiencyTile
 */
type StorageEfficiencyTile struct {
	DataInBytes               *int64    `json:"dataInBytes,omitempty" form:"dataInBytes,omitempty"`                             //Specifies the size of data brought into the cluster. This is the usage
	DataInBytesSamples        []*Sample `json:"dataInBytesSamples,omitempty" form:"dataInBytesSamples,omitempty"`               //Specifies the samples taken for Data brought into the cluster in bytes
	DataInDedupedBytes        *int64    `json:"dataInDedupedBytes,omitempty" form:"dataInDedupedBytes,omitempty"`               //Specifies the size of data after compression and or dedupe operations
	DataInDedupedBytesSamples []*Sample `json:"dataInDedupedBytesSamples,omitempty" form:"dataInDedupedBytesSamples,omitempty"` //Specifies the samples taken for morphed data in bytes in ascending order
	DedupeRatio               *float64  `json:"dedupeRatio,omitempty" form:"dedupeRatio,omitempty"`                             //Specifies the current dedupe ratio on the cluster. It is the ratio of
	DedupeRatioSamples        []*Sample `json:"dedupeRatioSamples,omitempty" form:"dedupeRatioSamples,omitempty"`               //Specifies the samples for data reduction ratio in ascending order of time.
	DurationDays              *int64    `json:"durationDays,omitempty" form:"durationDays,omitempty"`                           //Specifies the duration in days in which the samples were taken.
	IntervalSeconds           *int64    `json:"intervalSeconds,omitempty" form:"intervalSeconds,omitempty"`                     //Specifies the interval between the samples in seconds.
	LogicalUsedBytes          *int64    `json:"logicalUsedBytes,omitempty" form:"logicalUsedBytes,omitempty"`                   //Specifies the size of logical data currently represented on the cluster.
	LogicalUsedBytesSamples   []*Sample `json:"logicalUsedBytesSamples,omitempty" form:"logicalUsedBytesSamples,omitempty"`     //Specifies the samples taken for logical data represented in bytes in
	PhysicalUsedBytes         *int64    `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"`                 //Specifies the size of physical data currently consumed on the cluster.
	PhysicalUsedBytesSamples  []*Sample `json:"physicalUsedBytesSamples,omitempty" form:"physicalUsedBytesSamples,omitempty"`   //Specifies the samples taken for physical data consumed in bytes in
	StorageReductionRatio     *float64  `json:"storageReductionRatio,omitempty" form:"storageReductionRatio,omitempty"`         //Specifies the current storage reduction ratio on the cluster.
	StorageReductionSamples   []*Sample `json:"storageReductionSamples,omitempty" form:"storageReductionSamples,omitempty"`     //Specifies the samples for storage reduction ratio in ascending order of
}

/*
 * Structure for the custom type StoragePolicy
 */
type StoragePolicy struct {
	AppMarkerDetection           *bool                 `json:"appMarkerDetection,omitempty" form:"appMarkerDetection,omitempty"`                     //Specifies Whether to support app marker detection. When this is set to
	CloudSpillVaultId            *int64                `json:"cloudSpillVaultId,omitempty" form:"cloudSpillVaultId,omitempty"`                       //Specifies the vault id assigned for an external Storage
	CompressionPolicy            CompressionPolicyEnum `json:"compressionPolicy,omitempty" form:"compressionPolicy,omitempty"`                       //Specifies the compression setting to be applied to a Storage Domain
	DeduplicateCompressDelaySecs *int64                `json:"deduplicateCompressDelaySecs,omitempty" form:"deduplicateCompressDelaySecs,omitempty"` //Specifies the time in seconds when deduplication and compression
	DeduplicationEnabled         *bool                 `json:"deduplicationEnabled,omitempty" form:"deduplicationEnabled,omitempty"`                 //Specifies if deduplication is enabled for the Storage Domain (View Box).
	EncryptionPolicy             EncryptionPolicyEnum  `json:"encryptionPolicy,omitempty" form:"encryptionPolicy,omitempty"`                         //Specifies the encryption setting for the Storage Domain (View Box).
	ErasureCodingInfo            *ErasureCodingInfo    `json:"erasureCodingInfo,omitempty" form:"erasureCodingInfo,omitempty"`                       //Specifies information for erasure coding.
	InlineCompress               *bool                 `json:"inlineCompress,omitempty" form:"inlineCompress,omitempty"`                             //Specifies if compression should occur inline (as the data is being
	InlineDeduplicate            *bool                 `json:"inlineDeduplicate,omitempty" form:"inlineDeduplicate,omitempty"`                       //Specifies if deduplication should occur inline (as the data is being
	NumFailuresTolerated         *int64                `json:"numFailuresTolerated,omitempty" form:"numFailuresTolerated,omitempty"`                 //Number of disk failures to tolerate. This is an optional field. Default value
	NumNodeFailuresTolerated     *int64                `json:"numNodeFailuresTolerated,omitempty" form:"numNodeFailuresTolerated,omitempty"`         //Number of node failures to tolerate. If NumNodeFailuresTolerated is set to
}

/*
 * Structure for the custom type StoragePolicyOverride
 */
type StoragePolicyOverride struct {
	DisableInlineDedupAndCompression *bool `json:"disableInlineDedupAndCompression,omitempty" form:"disableInlineDedupAndCompression,omitempty"` //If false, the inline deduplication and compression settings inherited
}

/*
 * Structure for the custom type StorageSnapshotProviderParams
 */
type StorageSnapshotProviderParams struct {
	ConnectorParams *ConnectorParams `json:"connectorParams,omitempty" form:"connectorParams,omitempty"` //Message that encapsulates the various params required to establish a
	Entity          *EntityProto     `json:"entity,omitempty" form:"entity,omitempty"`                   //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type StorageStats
 */
type StorageStats struct {
	DataProtectionLogicalUsageBytes  *int64 `json:"dataProtectionLogicalUsageBytes,omitempty" form:"dataProtectionLogicalUsageBytes,omitempty"`   //Specifies the logical size of protected objects in bytes.
	DataProtectionPhysicalUsageBytes *int64 `json:"dataProtectionPhysicalUsageBytes,omitempty" form:"dataProtectionPhysicalUsageBytes,omitempty"` //Specifies the physical size of protected objects in bytes.
	FileServicesLogicalUsageBytes    *int64 `json:"fileServicesLogicalUsageBytes,omitempty" form:"fileServicesLogicalUsageBytes,omitempty"`       //Specifies the logical size consumed by file services in bytes.
	FileServicesPhysicalUsageBytes   *int64 `json:"fileServicesPhysicalUsageBytes,omitempty" form:"fileServicesPhysicalUsageBytes,omitempty"`     //Specifies the physical size consumed by file services in bytes.
	LocalAvailableBytes              *int64 `json:"localAvailableBytes,omitempty" form:"localAvailableBytes,omitempty"`                           //Specifies the local storage currently available on the cluster in bytes.
	LocalUsageBytes                  *int64 `json:"localUsageBytes,omitempty" form:"localUsageBytes,omitempty"`                                   //Specifies the local storage currently in use on the cluster in bytes.
	TotalCapacityBytes               *int64 `json:"totalCapacityBytes,omitempty" form:"totalCapacityBytes,omitempty"`                             //Specifies the total capacity of the cluster in bytes.
}

/*
 * Structure for the custom type StubbingPolicyProto
 */
type StubbingPolicyProto struct {
	RetentionPolicy  *RetentionPolicyProto  `json:"retentionPolicy,omitempty" form:"retentionPolicy,omitempty"`   //Message that specifies the retention policy for backup snapshots.
	SchedulingPolicy *SchedulingPolicyProto `json:"schedulingPolicy,omitempty" form:"schedulingPolicy,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Subnet
 */
type Subnet struct {
	Component     *string       `json:"component,omitempty" form:"component,omitempty"`         //Component that has reserved the subnet.
	Description   *string       `json:"description,omitempty" form:"description,omitempty"`     //Description of the subnet.
	Id            *int64        `json:"id,omitempty" form:"id,omitempty"`                       //ID of the subnet.
	Ip            *string       `json:"ip,omitempty" form:"ip,omitempty"`                       //Specifies either an IPv6 address or an IPv4 address.
	NetmaskBits   *int64        `json:"netmaskBits,omitempty" form:"netmaskBits,omitempty"`     //Specifies the netmask using bits.
	NetmaskIp4    *string       `json:"netmaskIp4,omitempty" form:"netmaskIp4,omitempty"`       //Specifies the netmask using an IP4 address.
	NfsAccess     NfsAccessEnum `json:"nfsAccess,omitempty" form:"nfsAccess,omitempty"`         //Specifies whether clients from this subnet can mount using NFS protocol.
	NfsRootSquash *bool         `json:"nfsRootSquash,omitempty" form:"nfsRootSquash,omitempty"` //Specifies whether clients from this subnet can mount as root on NFS.
	SmbAccess     SmbAccessEnum `json:"smbAccess,omitempty" form:"smbAccess,omitempty"`         //Specifies whether clients from this subnet can mount using SMB protocol.
}

/*
 * Structure for the custom type SupportedConfig
 */
type SupportedConfig struct {
	MinNodesAllowed        *int64    `json:"minNodesAllowed,omitempty" form:"minNodesAllowed,omitempty"`               //Specifies the minimum number of Nodes supported for this Cluster type.
	SupportedErasureCoding *[]string `json:"supportedErasureCoding,omitempty" form:"supportedErasureCoding,omitempty"` //Array of Supported Erasure Coding Options.
}

/*
 * Structure for the custom type SyslogServer
 */
type SyslogServer struct {
	Address                  string                   `json:"address" form:"address"`                                                       //Specifies the IP address or hostname of the syslog server.
	IsClusterAuditingEnabled *bool                    `json:"isClusterAuditingEnabled,omitempty" form:"isClusterAuditingEnabled,omitempty"` //Specifies if Cluster audit logs should be sent to this syslog server.
	IsFilerAuditingEnabled   *bool                    `json:"isFilerAuditingEnabled,omitempty" form:"isFilerAuditingEnabled,omitempty"`     //Specifies if filer audit logs should be sent to this syslog server.
	Name                     *string                  `json:"name,omitempty" form:"name,omitempty"`                                         //Specifies a unique name for the syslog server on the Cluster.
	Port                     int64                    `json:"port" form:"port"`                                                             //Specifies the port where the syslog server listens.
	Protocol                 ProtocolSyslogServerEnum `json:"protocol" form:"protocol"`                                                     //Specifies the protocol used to send the logs.
}

/*
 * Structure for the custom type TagAttribute
 */
type TagAttribute struct {
	Id   *int64  `json:"id,omitempty" form:"id,omitempty"`     //Specifies the Coheisty id of the VM tag.
	Name *string `json:"name,omitempty" form:"name,omitempty"` //Specifies the VMware name of the VM tag.
	Uuid *string `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the VMware Universally Unique Identifier (UUID) of the
}

/*
 * Structure for the custom type TapeMediaInformation
 */
type TapeMediaInformation struct {
	Barcode  *string `json:"barcode,omitempty" form:"barcode,omitempty"`   //Specifies a unique identifier for the media.
	Location *string `json:"location,omitempty" form:"location,omitempty"` //Specifies the location of the offline media as recorded by the
	Online   *bool   `json:"online,omitempty" form:"online,omitempty"`     //Specifies a flag that indicates if the media is online or offline.
}

/*
 * Structure for the custom type Task
 */
type Task struct {
	Attributes               []*TaskAttribute `json:"attributes,omitempty" form:"attributes,omitempty"`                             //The latest attributes reported for this task.
	EndTimeSeconds           *int64           `json:"endTimeSeconds,omitempty" form:"endTimeSeconds,omitempty"`                     //Specifies the end time of the task.
	ErrorMessage             *string          `json:"errorMessage,omitempty" form:"errorMessage,omitempty"`                         //Specifies an optional error message for this task.
	Events                   []*TaskEvent     `json:"events,omitempty" form:"events,omitempty"`                                     //Specifies the events reported for this task.
	ExpectedEndTimeSeconds   *int64           `json:"expectedEndTimeSeconds,omitempty" form:"expectedEndTimeSeconds,omitempty"`     //Specifies the estimated end time of the task.
	ExpectedSecondsRemaining *int64           `json:"expectedSecondsRemaining,omitempty" form:"expectedSecondsRemaining,omitempty"` //Specifies the expected remaining time for this task in seconds.
	ExpectedTotalWorkCount   *int64           `json:"expectedTotalWorkCount,omitempty" form:"expectedTotalWorkCount,omitempty"`     //The expected raw count of the total work remaining. This is the highest
	LastUpdateTimeSeconds    *int64           `json:"lastUpdateTimeSeconds,omitempty" form:"lastUpdateTimeSeconds,omitempty"`       //Specifies the timestamp when the last progress was reported.
	PercentFinished          *float64         `json:"percentFinished,omitempty" form:"percentFinished,omitempty"`                   //Specifies the reported progress on the task.
	StartTimeSeconds         *int64           `json:"startTimeSeconds,omitempty" form:"startTimeSeconds,omitempty"`                 //Specifies the start time of the task.
	Status                   StatusTaskEnum   `json:"status,omitempty" form:"status,omitempty"`                                     //Specifies the status of the task being queried.
	SubTasks                 *[]interface{}   `json:"subTasks,omitempty" form:"subTasks,omitempty"`                                 //Specifies a list of subtasks belonging to this task.
	TaskPath                 *string          `json:"taskPath,omitempty" form:"taskPath,omitempty"`                                 //Specifes the path of this task.
}

/*
 * Structure for the custom type TaskAttribute
 */
type TaskAttribute struct {
	Name      *string       `json:"name,omitempty" form:"name,omitempty"`           //Specifies the name of this Task Attribute.
	Value     *string       `json:"value,omitempty" form:"value,omitempty"`         //Specifies the value of this Task Attribute.
	ValueType ValueTypeEnum `json:"valueType,omitempty" form:"valueType,omitempty"` //Specifies the type of the value contained here. All values are returned as
}

/*
 * Structure for the custom type TaskEvent
 */
type TaskEvent struct {
	EventMessage       *string  `json:"eventMessage,omitempty" form:"eventMessage,omitempty"`             //Specifies the message associated with an event.
	PercentFinished    *float64 `json:"percentFinished,omitempty" form:"percentFinished,omitempty"`       //Specifies the completion percentage of the task attached to this event.
	RemainingWorkCount *int64   `json:"remainingWorkCount,omitempty" form:"remainingWorkCount,omitempty"` //Specifies the amount of work remaining for the task attached to this event.
	TimestampSeconds   *int64   `json:"timestampSeconds,omitempty" form:"timestampSeconds,omitempty"`     //Specifies the Unix timestamp that the event occurred.
}

/*
 * Structure for the custom type TaskNotification
 */
type TaskNotification struct {
	BackupTask        *BackupTaskInfo            `json:"backupTask,omitempty" form:"backupTask,omitempty"`               //TODO: Write general description for this field
	CloneTask         *CloneTaskInfo             `json:"cloneTask,omitempty" form:"cloneTask,omitempty"`                 //Parameters for a clone op.
	CreatedTimeSecs   *int64                     `json:"createdTimeSecs,omitempty" form:"createdTimeSecs,omitempty"`     //Timestamp at which the notification was created.
	Description       *string                    `json:"description,omitempty" form:"description,omitempty"`             //Description holds the actual notification text generated for
	Dismissed         *bool                      `json:"dismissed,omitempty" form:"dismissed,omitempty"`                 //Dismissed keeps track of whether a notification has been seen
	DismissedTimeSecs *int64                     `json:"dismissedTimeSecs,omitempty" form:"dismissedTimeSecs,omitempty"` //Timestamp at which user dismissed this notification event.
	FieldMessageTask  *BasicTaskInfo             `json:"fieldMessageTask,omitempty" form:"fieldMessageTask,omitempty"`   //TODO: Write general description for this field
	Id                *string                    `json:"id,omitempty" form:"id,omitempty"`                               //id identifies a user notification event uniquely.
	RecoveryTask      *RecoveryTaskInfo          `json:"recoveryTask,omitempty" form:"recoveryTask,omitempty"`           //Parameters for a recovery op.
	Status            StatusTaskNotificationEnum `json:"status,omitempty" form:"status,omitempty"`                       //Status of the task.
	TaskType          TaskTypeEnum               `json:"taskType,omitempty" form:"taskType,omitempty"`                   //Task type denotes which type of task this notification is for.
	Visited           *bool                      `json:"visited,omitempty" form:"visited,omitempty"`                     //Visited keeps track of whether a notification has been seen or not.
	VisitedTimeSecs   *int64                     `json:"visitedTimeSecs,omitempty" form:"visitedTimeSecs,omitempty"`     //Timestamp at which user visited this notification event.
}

/*
 * Structure for the custom type Tenant
 */
type Tenant struct {
	ActiveDirectories      []*ActiveDirectoryEntry `json:"activeDirectories,omitempty" form:"activeDirectories,omitempty"`           //Specifies the active directories this tenant is associated to.
	BifrostEnabled         *bool                   `json:"bifrostEnabled,omitempty" form:"bifrostEnabled,omitempty"`                 //Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
	CreatedTimeMsecs       *int64                  `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`             //Specifies the epoch time in milliseconds when the tenant account
	Deleted                *bool                   `json:"deleted,omitempty" form:"deleted,omitempty"`                               //Specifies if the Tenant is deleted.
	DeletedTimeMsecs       *int64                  `json:"deletedTimeMsecs,omitempty" form:"deletedTimeMsecs,omitempty"`             //Specifies the timestamp at which the tenant was deleted.
	DeletionFinished       *bool                   `json:"deletionFinished,omitempty" form:"deletionFinished,omitempty"`             //Specifies if the object collection is complete for the tenant.
	DeletionInfoVec        []*TenantDeletionInfo   `json:"deletionInfoVec,omitempty" form:"deletionInfoVec,omitempty"`               //Specifies the current deletion state of object categories.
	Description            *string                 `json:"description,omitempty" form:"description,omitempty"`                       //Specifies the description of this tenant.
	EntityIds              *[]int64                `json:"entityIds,omitempty" form:"entityIds,omitempty"`                           //Specifies the EntityIds this tenant is associated to.
	LastUpdatedTimeMsecs   *int64                  `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"`     //Specifies the epoch time in milliseconds when the tenant account was last
	LdapProviders          []*LdapProviderResponse `json:"ldapProviders,omitempty" form:"ldapProviders,omitempty"`                   //Specifies the ldap providers this tenant is associated to.
	Name                   *string                 `json:"name,omitempty" form:"name,omitempty"`                                     //Specifies the name of the tenant.
	OrgSuffix              *string                 `json:"orgSuffix,omitempty" form:"orgSuffix,omitempty"`                           //Specifies the organization suffix needed to construct tenant id. Tenant id
	ParentTenantId         *string                 `json:"parentTenantId,omitempty" form:"parentTenantId,omitempty"`                 //Specifies the parent tenant of this tenant if available.
	PolicyIds              *[]string               `json:"policyIds,omitempty" form:"policyIds,omitempty"`                           //Specifies the PolicyIds this tenant is associated to.
	ProtectionJobs         []*BackupJobProto       `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"`                 //Specifies the ProtectionJobs this tenant is associated to.
	SubscribeToAlertEmails *bool                   `json:"subscribeToAlertEmails,omitempty" form:"subscribeToAlertEmails,omitempty"` //Service provider can optionally unsubscribe from the Tenant Alert Emails.
	TenantId               *string                 `json:"tenantId,omitempty" form:"tenantId,omitempty"`                             //Specifies the unique id of the tenant.
	ViewBoxIds             *[]int64                `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"`                         //Specifies the ViewBoxIds this tenant is associated to.
	Views                  []*View                 `json:"views,omitempty" form:"views,omitempty"`                                   //Specifies the Views this tenant is associated to.
	VlanIfaceNames         *[]string               `json:"vlanIfaceNames,omitempty" form:"vlanIfaceNames,omitempty"`                 //Specifies the VlanIfaceNames this tenant is associated to,
}

/*
 * Structure for the custom type TenantActiveDirectoryUpdate
 */
type TenantActiveDirectoryUpdate struct {
	ActiveDirectoryDomains *[]string `json:"activeDirectoryDomains,omitempty" form:"activeDirectoryDomains,omitempty"` //Specifies the ActiveDirectoryDomain vec for respective tenant.
	TenantId               *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                             //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantActiveDirectoryUpdateParameters
 */
type TenantActiveDirectoryUpdateParameters struct {
	ActiveDirectoryDomains *[]string `json:"activeDirectoryDomains,omitempty" form:"activeDirectoryDomains,omitempty"` //Specifies the ActiveDirectoryDomain vec for respective tenant.
	TenantId               *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                             //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantCreateParameters
 */
type TenantCreateParameters struct {
	BifrostEnabled         *bool   `json:"bifrostEnabled,omitempty" form:"bifrostEnabled,omitempty"`                 //Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
	Description            *string `json:"description,omitempty" form:"description,omitempty"`                       //Specifies the description of this tenant.
	Name                   *string `json:"name,omitempty" form:"name,omitempty"`                                     //Specifies the name of the tenant.
	OrgSuffix              *string `json:"orgSuffix,omitempty" form:"orgSuffix,omitempty"`                           //Specifies the organization suffix needed to construct tenant id. Tenant id
	ParentTenantId         *string `json:"parentTenantId,omitempty" form:"parentTenantId,omitempty"`                 //Specifies the parent tenant of this tenant if available.
	SubscribeToAlertEmails *bool   `json:"subscribeToAlertEmails,omitempty" form:"subscribeToAlertEmails,omitempty"` //Service provider can optionally unsubscribe from the Tenant Alert Emails.
}

/*
 * Structure for the custom type TenantDeletionInfo
 */
type TenantDeletionInfo struct {
	Category            *int64  `json:"category,omitempty" form:"category,omitempty"`                       //Specifies the category of objects whose deletion state is being captured.
	FinishedAtTimeMsecs *int64  `json:"finishedAtTimeMsecs,omitempty" form:"finishedAtTimeMsecs,omitempty"` //Specifies the time when the process finished.
	ProcessedAtNode     *string `json:"processedAtNode,omitempty" form:"processedAtNode,omitempty"`         //Specifies the node ip where the process ran. Typically this would be
	RetryCount          *int64  `json:"retryCount,omitempty" form:"retryCount,omitempty"`                   //Specifies the number of times this task has been retried.
	StartedAtTimeMsecs  *int64  `json:"startedAtTimeMsecs,omitempty" form:"startedAtTimeMsecs,omitempty"`   //Specifies the time when the process started.
	State               *int64  `json:"state,omitempty" form:"state,omitempty"`                             //Specifies the deletion completion state of the object category.
}

/*
 * Structure for the custom type TenantEntityUpdate
 */
type TenantEntityUpdate struct {
	EntityIds *[]int64 `json:"entityIds,omitempty" form:"entityIds,omitempty"` //Specifies the EntityIds for respective tenant.
	TenantId  *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantEntityUpdateParameters
 */
type TenantEntityUpdateParameters struct {
	EntityIds *[]int64 `json:"entityIds,omitempty" form:"entityIds,omitempty"` //Specifies the EntityIds for respective tenant.
	TenantId  *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantInfo
 */
type TenantInfo struct {
	Name     *string `json:"name,omitempty" form:"name,omitempty"`         //Specifies name of the tenant.
	TenantId *string `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantLdapProviderUpdate
 */
type TenantLdapProviderUpdate struct {
	LdapProviderIds *[]int64 `json:"ldapProviderIds,omitempty" form:"ldapProviderIds,omitempty"` //Specifies the ids of ldap providers for respective tenant.
	TenantId        *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`               //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantLdapProviderUpdateParameters
 */
type TenantLdapProviderUpdateParameters struct {
	LdapProviderIds *[]int64 `json:"ldapProviderIds,omitempty" form:"ldapProviderIds,omitempty"` //Specifies the ids of ldap providers for respective tenant.
	TenantId        *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`               //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionJobUpdate
 */
type TenantProtectionJobUpdate struct {
	ProtectionJobIds *[]int64 `json:"protectionJobIds,omitempty" form:"protectionJobIds,omitempty"` //Specifies the ProtectionJobIds vec for respective tenant.
	TenantId         *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`                 //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionJobUpdateParameters
 */
type TenantProtectionJobUpdateParameters struct {
	ProtectionJobIds *[]int64 `json:"protectionJobIds,omitempty" form:"protectionJobIds,omitempty"` //Specifies the ProtectionJobIds vec for respective tenant.
	TenantId         *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`                 //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionPolicyUpdate
 */
type TenantProtectionPolicyUpdate struct {
	PolicyIds *[]string `json:"policyIds,omitempty" form:"policyIds,omitempty"` //Specifies the PolicyIds for respective tenant.
	TenantId  *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionPolicyUpdateParameters
 */
type TenantProtectionPolicyUpdateParameters struct {
	PolicyIds *[]string `json:"policyIds,omitempty" form:"policyIds,omitempty"` //Specifies the PolicyIds for respective tenant.
	TenantId  *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProxy
 */
type TenantProxy struct {
	IpAddress *string `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the ip address of the proxy.
	TenantId  *string `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantStats
 */
type TenantStats struct {
	GroupList      []*StatsGroup      `json:"groupList,omitempty" form:"groupList,omitempty"`           //Specifies a list of groups associated to this tenant (organization).
	Id             *string            `json:"id,omitempty" form:"id,omitempty"`                         //Specifies the id of the tenant (organization).
	Name           *string            `json:"name,omitempty" form:"name,omitempty"`                     //Specifies the name of the tenant (organization).
	SchemaInfoList []*UsageSchemaInfo `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"` //Specifies a list of schemaInfos of the tenant (organization).
	Stats          *DataUsageStats    `json:"stats,omitempty" form:"stats,omitempty"`                   //Specifies the data usage metric of the data stored on the Cohesity
}

/*
 * Structure for the custom type TenantUpdate
 */
type TenantUpdate struct {
	BifrostEnabled         *bool   `json:"bifrostEnabled,omitempty" form:"bifrostEnabled,omitempty"`                 //Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
	Description            *string `json:"description,omitempty" form:"description,omitempty"`                       //Specifies the description of this tenant.
	Name                   *string `json:"name,omitempty" form:"name,omitempty"`                                     //Specifies the name of the tenant.
	SubscribeToAlertEmails *bool   `json:"subscribeToAlertEmails,omitempty" form:"subscribeToAlertEmails,omitempty"` //Service provider can optionally unsubscribe from the Tenant Alert Emails.
	TenantId               *string `json:"tenantId,omitempty" form:"tenantId,omitempty"`                             //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantUserUpdateParameters
 */
type TenantUserUpdateParameters struct {
	Sids     *[]string `json:"sids,omitempty" form:"sids,omitempty"`         //Specifies the array of Sid of the users.
	TenantId *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantViewBoxUpdate
 */
type TenantViewBoxUpdate struct {
	TenantId   *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`     //Specifies the unique id of the tenant.
	ViewBoxIds *[]int64 `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //Specifies the ViewBoxIds for respective tenant.
}

/*
 * Structure for the custom type TenantViewBoxUpdateParameters
 */
type TenantViewBoxUpdateParameters struct {
	TenantId   *string  `json:"tenantId,omitempty" form:"tenantId,omitempty"`     //Specifies the unique id of the tenant.
	ViewBoxIds *[]int64 `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //Specifies the ViewBoxIds for respective tenant.
}

/*
 * Structure for the custom type TenantViewUpdate
 */
type TenantViewUpdate struct {
	TenantId  *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
	ViewNames *[]string `json:"viewNames,omitempty" form:"viewNames,omitempty"` //Specifies the PolicyIds for respective tenant.
}

/*
 * Structure for the custom type TenantViewUpdateParameters
 */
type TenantViewUpdateParameters struct {
	TenantId  *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`   //Specifies the unique id of the tenant.
	ViewNames *[]string `json:"viewNames,omitempty" form:"viewNames,omitempty"` //Specifies the PolicyIds for respective tenant.
}

/*
 * Structure for the custom type TenantVlanUpdate
 */
type TenantVlanUpdate struct {
	TenantId       *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`             //Specifies the unique id of the tenant.
	VlanIfaceNames *[]string `json:"vlanIfaceNames,omitempty" form:"vlanIfaceNames,omitempty"` //Specifies the VlanIfaceNames for respective tenant,
}

/*
 * Structure for the custom type TenantVlanUpdateParameters
 */
type TenantVlanUpdateParameters struct {
	TenantId       *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`             //Specifies the unique id of the tenant.
	VlanIfaceNames *[]string `json:"vlanIfaceNames,omitempty" form:"vlanIfaceNames,omitempty"` //Specifies the VlanIfaceNames for respective tenant,
}

/*
 * Structure for the custom type TestIdpReachability
 */
type TestIdpReachability struct {
	IssuerId *string `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the IdP provided Issuer ID for the app.
	SsoUrl   *string `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"`     //Specifies the SSO URL of the IdP service for the customer. This is the
}

/*
 * Structure for the custom type ThrottlingPolicyOverride
 */
type ThrottlingPolicyOverride struct {
	DatastoreId      *int64                      `json:"datastoreId,omitempty" form:"datastoreId,omitempty"`           //Specifies the Protection Source id of the Datastore.
	DatastoreName    *string                     `json:"datastoreName,omitempty" form:"datastoreName,omitempty"`       //Specifies the display name of the Datastore.
	ThrottlingPolicy *ThrottlingPolicyParameters `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"` //Specifies the throttling policy for a registered Protection Source.
}

/*
 * Structure for the custom type ThrottlingPolicyParameters
 */
type ThrottlingPolicyParameters struct {
	EnforceMaxStreams    *bool              `json:"enforceMaxStreams,omitempty" form:"enforceMaxStreams,omitempty"`       //Specifies whether datastore streams are configured for all datastores
	IsEnabled            *bool              `json:"isEnabled,omitempty" form:"isEnabled,omitempty"`                       //Indicates whether read operations to the datastores, which are
	LatencyThresholds    *LatencyThresholds `json:"latencyThresholds,omitempty" form:"latencyThresholds,omitempty"`       //Specifies latency thresholds that trigger throttling for all datastores
	MaxConcurrentStreams *int64             `json:"maxConcurrentStreams,omitempty" form:"maxConcurrentStreams,omitempty"` //Specifies the limit on the number of streams Cohesity cluster will make
}

/*
 * Structure for the custom type ThroughputTile
 */
type ThroughputTile struct {
	MaxReadThroughput      *int64    `json:"maxReadThroughput,omitempty" form:"maxReadThroughput,omitempty"`           //Maxium Read throughput in last 24 hours.
	MaxWriteThroughput     *int64    `json:"maxWriteThroughput,omitempty" form:"maxWriteThroughput,omitempty"`         //Maximum Write throughput in last 24 hours.
	ReadThroughputSamples  []*Sample `json:"readThroughputSamples,omitempty" form:"readThroughputSamples,omitempty"`   //Read throughput samples taken for the past 24 hours at 10 minutes
	WriteThroughputSamples []*Sample `json:"writeThroughputSamples,omitempty" form:"writeThroughputSamples,omitempty"` //Write throughput samples taken for the past 24 hours at 10 minutes
}

/*
 * Structure for the custom type Time
 */
type Time struct {
	Hour   *int64 `json:"hour,omitempty" form:"hour,omitempty"`     //The hour when backup should be performed (0 - 23).
	Minute *int64 `json:"minute,omitempty" form:"minute,omitempty"` //The minute when backup should be performed (0 - 59).
}

/*
 * Structure for the custom type TimeOfAWeek
 */
type TimeOfAWeek struct {
	Days      *[]DayEnum `json:"days,omitempty" form:"days,omitempty"`           //Array of Week Days.
	EndTime   *TimeOfDay `json:"endTime,omitempty" form:"endTime,omitempty"`     //Specifies the end time for the daily time period.
	StartTime *TimeOfDay `json:"startTime,omitempty" form:"startTime,omitempty"` //Specifies the start time for the daily time period.
}

/*
 * Structure for the custom type TimeOfDay
 */
type TimeOfDay struct {
	Hour   *int64 `json:"hour,omitempty" form:"hour,omitempty"`     //Specifies an (0-23) hour in a day.
	Minute *int64 `json:"minute,omitempty" form:"minute,omitempty"` //Specifies a (0-59) minute in an hour.
}

/*
 * Structure for the custom type TimeRangeSettings
 */
type TimeRangeSettings struct {
	EndTimeUsecs   *int64       `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"`     //Specifies the end time specified as a Unix epoch Timestamp
	JobUid         *UniversalId `json:"jobUid,omitempty" form:"jobUid,omitempty"`                 //Specifies an id for an object that is unique across Cohesity Clusters.
	StartTimeUsecs *int64       `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time specified as a Unix epoch Timestamp
}

/*
 * Structure for the custom type TimeSeriesSchemaResponse
 */
type TimeSeriesSchemaResponse struct {
	SchemaInfoList []*SchemaInfo `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"` //Specifies the list of the schema info for an entity.
}

/*
 * Structure for the custom type TrendingData
 */
type TrendingData struct {
	Cancelled           *int64  `json:"cancelled,omitempty" form:"cancelled,omitempty"`                     //Specifies number of cancelled runs.
	Failed              *int64  `json:"failed,omitempty" form:"failed,omitempty"`                           //Specifies number of failed runs.
	Running             *int64  `json:"running,omitempty" form:"running,omitempty"`                         //Specifies number of in-progress runs.
	Successful          *int64  `json:"successful,omitempty" form:"successful,omitempty"`                   //Specifies number of successful runs.
	Total               *int64  `json:"total,omitempty" form:"total,omitempty"`                             //Specifies total number of runs.
	TrendName           *string `json:"trendName,omitempty" form:"trendName,omitempty"`                     //Specifies trend name. This is start of the day/week/month.
	TrendStartTimeUsecs *int64  `json:"trendStartTimeUsecs,omitempty" form:"trendStartTimeUsecs,omitempty"` //Specifies start of the day/week/month in micro seconds
}

/*
 * Structure for the custom type UnRegisterApplicationServersParameters
 */
type UnRegisterApplicationServersParameters struct {
	Applications *[]ApplicationEnum `json:"applications,omitempty" form:"applications,omitempty"` //Specifies the types of applications such as 'kSQL', 'kExchange' running
}

/*
 * Structure for the custom type UniversalId
 */
type UniversalId struct {
	ClusterId            *int64 `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //Specifies the Cohesity Cluster id where the object was created.
	ClusterIncarnationId *int64 `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies an id for the Cohesity Cluster that is generated when
	Id                   *int64 `json:"id,omitempty" form:"id,omitempty"`                                     //Specifies a unique id assigned to an object (such as a Job)
}

/*
 * Structure for the custom type UniversalIdProto
 */
type UniversalIdProto struct {
	ClusterId            *int64 `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //The id of the cluster at which the object was created.
	ClusterIncarnationId *int64 `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //The incarnation id of the above cluster.
	ObjectId             *int64 `json:"objectId,omitempty" form:"objectId,omitempty"`                         //The object id - this is unique within the above cluster.
}

/*
 * Structure for the custom type UnprotectObjectParams
 */
type UnprotectObjectParams struct {
	DeleteSnapshots    *bool  `json:"deleteSnapshots,omitempty" form:"deleteSnapshots,omitempty"` //Specifies whether to delete the snapshots of the Protection Object.
	ProtectionSourceId int64  `json:"protectionSourceId" form:"protectionSourceId"`               //Specifies the id of the Protection Source to be unprotected.
	RpoPolicyId        string `json:"rpoPolicyId" form:"rpoPolicyId"`                             //Specifies the id of the Rpo Policy from which to unprotect the object.
}

/*
 * Structure for the custom type UpdateAntivirusServiceGroupParams
 */
type UpdateAntivirusServiceGroupParams struct {
	AntivirusServices []*AntivirusServiceConfigParams `json:"antivirusServices,omitempty" form:"antivirusServices,omitempty"` //Specifies the Antivirus services for this provider.
	Description       *string                         `json:"description,omitempty" form:"description,omitempty"`             //Specifies the description of the Antivirus service group.
	Id                int64                           `json:"id" form:"id"`                                                   //Specifies the Id of the Antivirus service group.
	IsEnabled         *bool                           `json:"isEnabled,omitempty" form:"isEnabled,omitempty"`                 //Specifies whether the antivirus service group is enabled or not.
	Name              string                          `json:"name" form:"name"`                                               //Specifies the name of the Antivirus service group.
}

/*
 * Structure for the custom type UpdateApplicationServerParameters
 */
type UpdateApplicationServerParameters struct {
	Applications       *[]ApplicationEnum `json:"applications,omitempty" form:"applications,omitempty"`             //Specifies the types of applications such as 'kSQL', 'kExchange' running
	HasPersistentAgent *bool              `json:"hasPersistentAgent,omitempty" form:"hasPersistentAgent,omitempty"` //Set this to true if a persistent agent is running on the host. If this is
	Password           *string            `json:"password,omitempty" form:"password,omitempty"`                     //Specifies password of the username to access the target source.
	ProtectionSourceId *int64             `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the Id of the Protection Source that contains one or more
	Username           *string            `json:"username,omitempty" form:"username,omitempty"`                     //Specifies username to access the target source.
}

/*
 * Structure for the custom type UpdateBondParameters
 */
type UpdateBondParameters struct {
	BondingMode    BondingModeUpdateBondParametersEnum `json:"bondingMode" form:"bondingMode"`                           //Specifies the new bonding mode.
	LacpRate       *string                             `json:"lacpRate,omitempty" form:"lacpRate,omitempty"`             //Specifies the LACP rate. If not specified,
	Name           string                              `json:"name" form:"name"`                                         //Specifies the name of the bond being updated.
	XmitHashPolicy *string                             `json:"xmitHashPolicy,omitempty" form:"xmitHashPolicy,omitempty"` //Specifies the xmit hash policy. If not specified,
}

/*
 * Structure for the custom type UpdateBondResult
 */
type UpdateBondResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies a message describing the result of the operation.
}

/*
 * Structure for the custom type UpdateClusterParams
 */
type UpdateClusterParams struct {
	AppsSettings                    *AppsConfig                   `json:"appsSettings,omitempty" form:"appsSettings,omitempty"`                                       //TODO: Write general description for this field
	BannerEnabled                   *bool                         `json:"bannerEnabled,omitempty" form:"bannerEnabled,omitempty"`                                     //Specifies whether UI banner is enabled on the cluster or not. When banner
	BondingMode                     BondingModeEnum               `json:"bondingMode,omitempty" form:"bondingMode,omitempty"`                                         //Specifies the bonding mode to use when bonding NICs to this Cluster.
	ClusterAuditLogConfig           *ClusterAuditLogConfiguration `json:"clusterAuditLogConfig,omitempty" form:"clusterAuditLogConfig,omitempty"`                     //Specifies the settings of the Cluster audit log configuration.
	DnsServerIps                    *[]string                     `json:"dnsServerIps,omitempty" form:"dnsServerIps,omitempty"`                                       //Array of IP Addresses of DNS Servers.
	DomainNames                     *[]string                     `json:"domainNames,omitempty" form:"domainNames,omitempty"`                                         //Array of Domain Names.
	EnableActiveMonitoring          *bool                         `json:"enableActiveMonitoring,omitempty" form:"enableActiveMonitoring,omitempty"`                   //Specifies if Cohesity can receive monitoring information from the
	EnableUpgradePkgPolling         *bool                         `json:"enableUpgradePkgPolling,omitempty" form:"enableUpgradePkgPolling,omitempty"`                 //If 'true', Cohesity's upgrade server is polled for new releases.
	EncryptionKeyRotationPeriodSecs *int64                        `json:"encryptionKeyRotationPeriodSecs,omitempty" form:"encryptionKeyRotationPeriodSecs,omitempty"` //Specifies the period of time (in seconds) when encryption keys are rotated.
	FilerAuditLogConfig             *FilerAuditLogConfiguration   `json:"filerAuditLogConfig,omitempty" form:"filerAuditLogConfig,omitempty"`                         //Specifies the settings of the filer audit log configuration.
	Gateway                         *string                       `json:"gateway,omitempty" form:"gateway,omitempty"`                                                 //Specifies the gateway IP address.
	GoogleAnalyticsEnabled          *bool                         `json:"googleAnalyticsEnabled,omitempty" form:"googleAnalyticsEnabled,omitempty"`                   //Specifies whether Google Analytics is enabled.
	IsDocumentationLocal            *bool                         `json:"isDocumentationLocal,omitempty" form:"isDocumentationLocal,omitempty"`                       //Specifies what version of the documentation is used.
	LanguageLocale                  *string                       `json:"languageLocale,omitempty" form:"languageLocale,omitempty"`                                   //Specifies the language and locale for this Cohesity Cluster.
	LocalAuthDomainName             *string                       `json:"localAuthDomainName,omitempty" form:"localAuthDomainName,omitempty"`                         //Domain name for SMB local authentication.
	LocalGroupsEnabled              *bool                         `json:"localGroupsEnabled,omitempty" form:"localGroupsEnabled,omitempty"`                           //Specifies whether to enable local groups on cluster. Once it is enabled,
	MetadataFaultToleranceFactor    *int64                        `json:"metadataFaultToleranceFactor,omitempty" form:"metadataFaultToleranceFactor,omitempty"`       //Specifies metadata fault tolerance setting for the cluster. This denotes
	Mtu                             *int64                        `json:"mtu,omitempty" form:"mtu,omitempty"`                                                         //Specifies the Maxium Transmission Unit (MTU) in bytes of
	MultiTenancyEnabled             *bool                         `json:"multiTenancyEnabled,omitempty" form:"multiTenancyEnabled,omitempty"`                         //Specifies if multi tenancy is enabled in the cluster. Authentication &
	Name                            *string                       `json:"name,omitempty" form:"name,omitempty"`                                                       //Specifies the name of the Cohesity Cluster.
	NtpSettings                     *NtpSettingsConfig            `json:"ntpSettings,omitempty" form:"ntpSettings,omitempty"`                                         //TODO: Write general description for this field
	ReverseTunnelEnabled            *bool                         `json:"reverseTunnelEnabled,omitempty" form:"reverseTunnelEnabled,omitempty"`                       //If 'true', Cohesity's Remote Tunnel is enabled.
	ReverseTunnelEndTimeMsecs       *int64                        `json:"reverseTunnelEndTimeMsecs,omitempty" form:"reverseTunnelEndTimeMsecs,omitempty"`             //ReverseTunnelEndTimeMsecs specifies the end time in milliseconds since
	SmbAdDisabled                   *bool                         `json:"smbAdDisabled,omitempty" form:"smbAdDisabled,omitempty"`                                     //Specifies if Active Directory should be disabled for authentication of SMB
	StigMode                        *bool                         `json:"stigMode,omitempty" form:"stigMode,omitempty"`                                               //Specifies if STIG mode is enabled or not.
	SyslogServers                   []*SyslogServer               `json:"syslogServers,omitempty" form:"syslogServers,omitempty"`                                     //Array of Syslog Servers.
	TenantViewboxSharingEnabled     *bool                         `json:"tenantViewboxSharingEnabled,omitempty" form:"tenantViewboxSharingEnabled,omitempty"`         //In case multi tenancy is enabled, this flag controls whether multiple
	Timezone                        *string                       `json:"timezone,omitempty" form:"timezone,omitempty"`                                               //Specifies the timezone to use for showing time in emails, reports,
	TurboMode                       *bool                         `json:"turboMode,omitempty" form:"turboMode,omitempty"`                                             //Specifies if the cluster is in Turbo mode.
}

/*
 * Structure for the custom type UpdateDirQuotaArgs
 */
type UpdateDirQuotaArgs struct {
	Quota    *DirQuotaPolicy `json:"quota,omitempty" form:"quota,omitempty"`       //Specifies a policy configuration for the directory quota. A policy is the
	ViewName *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the view.
}

/*
 * Structure for the custom type UpdateEulaConfig
 */
type UpdateEulaConfig struct {
	SignedVersion *int64 `json:"signedVersion,omitempty" form:"signedVersion,omitempty"` //Specifies the version of the End User License Agreement that was accepted.
}

/*
 * Structure for the custom type UpdateIdpConfigurationRequest
 */
type UpdateIdpConfigurationRequest struct {
	AllowLocalAuthentication *bool     `json:"allowLocalAuthentication,omitempty" form:"allowLocalAuthentication,omitempty"` //Specifies whether to allow local authentication. When IdP is configured,
	Certificate              *string   `json:"certificate,omitempty" form:"certificate,omitempty"`                           //Specifies the certificate generated for the app by the IdP service when
	CertificateFilename      *string   `json:"certificateFilename,omitempty" form:"certificateFilename,omitempty"`           //Specifies the filename used to upload the certificate.
	Enable                   *bool     `json:"enable,omitempty" form:"enable,omitempty"`                                     //Specifies a flag to enable or disable this IdP service. When it is set
	IssuerId                 *string   `json:"issuerId,omitempty" form:"issuerId,omitempty"`                                 //Specifies the IdP provided Issuer ID for the app.
	Roles                    *[]string `json:"roles,omitempty" form:"roles,omitempty"`                                       //Specifies a list roles assigned to an IdP user if samlAttributeName is
	SamlAttributeName        *string   `json:"samlAttributeName,omitempty" form:"samlAttributeName,omitempty"`               //Specifies the SAML attribute name that contains a comma separated list
	SignRequest              *bool     `json:"signRequest,omitempty" form:"signRequest,omitempty"`                           //Specifies whether to sign the SAML request or not. When it is set
	SsoUrl                   *string   `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"`                                     //Specifies the SSO URL of the IdP service for the customer. This is the
}

/*
 * Structure for the custom type UpdateIgnoredTrustedDomainsParams
 */
type UpdateIgnoredTrustedDomainsParams struct {
	IgnoredTrustedDomains *[]string `json:"ignoredTrustedDomains,omitempty" form:"ignoredTrustedDomains,omitempty"` //Specifies the list of trusted domains that were set by the user to be
}

/*
 * Structure for the custom type UpdateInfectedFileParams
 */
type UpdateInfectedFileParams struct {
	InfectedFileIds  []*InfectedFileParam                         `json:"infectedFileIds,omitempty" form:"infectedFileIds,omitempty"`   //Specifies the list of infected file identifiers.
	RemediationState RemediationStateUpdateInfectedFileParamsEnum `json:"remediationState,omitempty" form:"remediationState,omitempty"` //Specifies the remediation state of the file. Not setting any value to
}

/*
 * Structure for the custom type UpdateInfectedFileResponse
 */
type UpdateInfectedFileResponse struct {
	UpdateFailedInfectedFiles    []*InfectedFileId `json:"updateFailedInfectedFiles,omitempty" form:"updateFailedInfectedFiles,omitempty"`       //Specifies the failed update infected files.
	UpdateSucceededInfectedFiles []*InfectedFileId `json:"updateSucceededInfectedFiles,omitempty" form:"updateSucceededInfectedFiles,omitempty"` //Specifies the successfully updated infected files.
}

/*
 * Structure for the custom type UpdateLdapProviderParam
 */
type UpdateLdapProviderParam struct {
	AdDomainName            *string      `json:"adDomainName,omitempty" form:"adDomainName,omitempty"`                       //Specifies the domain name of an Active Directory which is mapped to this
	AuthType                AuthTypeEnum `json:"authType,omitempty" form:"authType,omitempty"`                               //Specifies the authentication type used while connecting to LDAP servers.
	BaseDistinguishedName   *string      `json:"baseDistinguishedName,omitempty" form:"baseDistinguishedName,omitempty"`     //Specifies the base distinguished name used as the base for LDAP
	DomainName              *string      `json:"domainName,omitempty" form:"domainName,omitempty"`                           //Specifies the name of the domain name to be used for querying LDAP servers
	Id                      *int64       `json:"id,omitempty" form:"id,omitempty"`                                           //Specifies the ID of the LDAP provider.
	Name                    *string      `json:"name,omitempty" form:"name,omitempty"`                                       //Specifies the name of the LDAP provider.
	Port                    *int64       `json:"port,omitempty" form:"port,omitempty"`                                       //Specifies LDAP server port.
	PreferredLdapServerList *[]string    `json:"preferredLdapServerList,omitempty" form:"preferredLdapServerList,omitempty"` //Specifies the preferred LDAP servers. Server names should be either in
	TenantId                *string      `json:"tenantId,omitempty" form:"tenantId,omitempty"`                               //Specifies the unique id of the tenant.
	UseSsl                  *bool        `json:"useSsl,omitempty" form:"useSsl,omitempty"`                                   //Specifies whether to use SSL for LDAP connections.
	UserDistinguishedName   *string      `json:"userDistinguishedName,omitempty" form:"userDistinguishedName,omitempty"`     //Specifies the user distinguished name that is used for LDAP authentication.
	UserPassword            *string      `json:"userPassword,omitempty" form:"userPassword,omitempty"`                       //Specifies the user password that is used for LDAP authentication.
}

/*
 * Structure for the custom type UpdateLdapProviderParams
 */
type UpdateLdapProviderParams struct {
	LdapProviderId *int64 `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"` //Specifies the LDAP provider id which is mapped to an Active Directory
}

/*
 * Structure for the custom type UpdateLinuxPasswordReqParams
 */
type UpdateLinuxPasswordReqParams struct {
	ClusterId     *int64    `json:"clusterId,omitempty" form:"clusterId,omitempty"` //If cluster ID is specified, then the password is updated for all the nodes
	LinuxPassword string    `json:"linuxPassword" form:"linuxPassword"`             //Specifies the new linux password.
	LinuxUsername string    `json:"linuxUsername" form:"linuxUsername"`             //Specifies the linux username for which the password will be updated.
	NodeIps       *[]string `json:"nodeIps,omitempty" form:"nodeIps,omitempty"`     //Specifies the node IP address on which the linux password will be updated.
}

/*
 * Structure for the custom type UpdateLinuxPasswordResult
 */
type UpdateLinuxPasswordResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type UpdateMachineAccountsParams
 */
type UpdateMachineAccountsParams struct {
	MachineAccounts           *[]string `json:"machineAccounts,omitempty" form:"machineAccounts,omitempty"`                     //Array of Machine Accounts.
	OverwriteExistingAccounts *bool     `json:"overwriteExistingAccounts,omitempty" form:"overwriteExistingAccounts,omitempty"` //Specifies whether the specified machine accounts should overwrite the
	Password                  *string   `json:"password,omitempty" form:"password,omitempty"`                                   //Specifies the password for the specified userName.
	UserName                  *string   `json:"userName,omitempty" form:"userName,omitempty"`                                   //Specifies a userName that has administrative privileges in the domain.
}

/*
 * Structure for the custom type UpdateProtectionJobRun
 */
type UpdateProtectionJobRun struct {
	CopyRunTargets    []*RunJobSnapshotTarget `json:"copyRunTargets,omitempty" form:"copyRunTargets,omitempty"`       //Specifies the retention for archival, replication or extended local
	ExpiryTimeUsecs   *int64                  `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"`     //Specifies a new expiration time as a Unix epoch Timestamp
	JobUid            *UniversalId            `json:"jobUid,omitempty" form:"jobUid,omitempty"`                       //Specifies a unique universal id for the Job.
	RunStartTimeUsecs *int64                  `json:"runStartTimeUsecs,omitempty" form:"runStartTimeUsecs,omitempty"` //Specifies the start time of the Job Run to update. The start time
	SourceIds         *[]int64                `json:"sourceIds,omitempty" form:"sourceIds,omitempty"`                 //Ids of the Protection Sources. If this is specified, retention time will
}

/*
 * Structure for the custom type UpdateProtectionJobRunsParam
 */
type UpdateProtectionJobRunsParam struct {
	JobRuns []*UpdateProtectionJobRun `json:"jobRuns,omitempty" form:"jobRuns,omitempty"` //Array of Job Runs.
}

/*
 * Structure for the custom type UpdateProtectionJobsState
 */
type UpdateProtectionJobsState struct {
	FailedJobIds     *[]int64 `json:"failedJobIds,omitempty" form:"failedJobIds,omitempty"`         //Specifies a list of Protection Job ids for which updation of state failed.
	SuccessfulJobIds *[]int64 `json:"successfulJobIds,omitempty" form:"successfulJobIds,omitempty"` //Specifies a list of Protection Job ids for which updation of state is
}

/*
 * Structure for the custom type UpdateProtectionJobsStateParams
 */
type UpdateProtectionJobsStateParams struct {
	Action ActionUpdateProtectionJobsStateParamsEnum `json:"action,omitempty" form:"action,omitempty"` //Specifies the action to be performed on all the specfied Protection Jobs.
	JobIds *[]int64                                  `json:"jobIds,omitempty" form:"jobIds,omitempty"` //Specifies a list of Protection Job ids for which the state should change.
}

/*
 * Structure for the custom type UpdateProtectionObjectParameters
 */
type UpdateProtectionObjectParameters struct {
	PauseBackup        *bool                     `json:"pauseBackup,omitempty" form:"pauseBackup,omitempty"`           //Specifies if the protection for the Protection Object is to be paused.
	ProtectedSourceUid UniversalId               `json:"protectedSourceUid" form:"protectedSourceUid"`                 //Specifies an id for an object that is unique across Cohesity Clusters.
	RpoPolicyId        *string                   `json:"rpoPolicyId,omitempty" form:"rpoPolicyId,omitempty"`           //Specifies the unique id of the new RPO policy to assign to the object.
	SourceParameters   []*SourceSpecialParameter `json:"sourceParameters,omitempty" form:"sourceParameters,omitempty"` //Specifies the additional special settings for a Protected Source.
}

/*
 * Structure for the custom type UpdateProtectionSourceParameters
 */
type UpdateProtectionSourceParameters struct {
	AgentEndpoint             *string                                      `json:"agentEndpoint,omitempty" form:"agentEndpoint,omitempty"`                         //Specifies the agent endpoint if it is different from the source endpoint.
	AwsCredentials            *AwsCredentials                              `json:"awsCredentials,omitempty" form:"awsCredentials,omitempty"`                       //Specifies the credentials to authenticate with AWS Cloud Platform.
	AzureCredentials          *AzureCredentials                            `json:"azureCredentials,omitempty" form:"azureCredentials,omitempty"`                   //Specifies the credentials to authenticate with Azure Cloud Platform.
	Endpoint                  *string                                      `json:"endpoint,omitempty" form:"endpoint,omitempty"`                                   //Specifies the network endpoint of the Protection Source where it is
	ForceRegister             *bool                                        `json:"forceRegister,omitempty" form:"forceRegister,omitempty"`                         //ForceRegister is applicable to Physical Environment. By default, the agent
	GcpCredentials            *GcpCredentials                              `json:"gcpCredentials,omitempty" form:"gcpCredentials,omitempty"`                       //Specifies the credentials to authenticate with Google Cloud Platform.
	HostType                  HostTypeUpdateProtectionSourceParametersEnum `json:"hostType,omitempty" form:"hostType,omitempty"`                                   //Specifies the optional OS type of the Protection Source (such as kWindows
	KubernetesCredentials     *KubernetesCredentials                       `json:"kubernetesCredentials,omitempty" form:"kubernetesCredentials,omitempty"`         //Specifies the credentials to authenticate with a Kubernetes Cluster.
	MinimumFreeSpaceGB        *int64                                       `json:"minimumFreeSpaceGB,omitempty" form:"minimumFreeSpaceGB,omitempty"`               //Specifies the minimum space in GB after which backup jobs will be canceled
	NasMountCredentials       *NasMountCredentialParams                    `json:"nasMountCredentials,omitempty" form:"nasMountCredentials,omitempty"`             //Specifies the server credentials to connect to a NetApp server.
	Office365Credentials      *Office365Credentials                        `json:"office365Credentials,omitempty" form:"office365Credentials,omitempty"`           //Specifies the credentials to authenticate with Office365 account.
	Password                  *string                                      `json:"password,omitempty" form:"password,omitempty"`                                   //Specifies password of the username to access the target source.
	SourceSideDedupEnabled    *bool                                        `json:"sourceSideDedupEnabled,omitempty" form:"sourceSideDedupEnabled,omitempty"`       //This controls whether to use source side dedup on the source or not.
	SslVerification           *SslVerification                             `json:"sslVerification,omitempty" form:"sslVerification,omitempty"`                     //Specifies information about SSL verification when registering certain
	ThrottlingPolicy          *ThrottlingPolicyParameters                  `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"`                   //Specifies the throttling policy that should be applied to this Source.
	ThrottlingPolicyOverrides []*ThrottlingPolicyOverride                  `json:"throttlingPolicyOverrides,omitempty" form:"throttlingPolicyOverrides,omitempty"` //Array of Throttling Policy Overrides for Datastores.
	Username                  *string                                      `json:"username,omitempty" form:"username,omitempty"`                                   //Specifies username to access the target source.
}

/*
 * Structure for the custom type UpdateResolutionParams
 */
type UpdateResolutionParams struct {
	AlertIdList *[]string `json:"alertIdList,omitempty" form:"alertIdList,omitempty"` //Specifies the Alerts to resolve, which are specified by Alert Ids.
}

/*
 * Structure for the custom type UpdateRestoreTaskParams
 */
type UpdateRestoreTaskParams struct {
	AdOptions     *AdRestoreOptions `json:"adOptions,omitempty" form:"adOptions,omitempty"`         //AdRestoreOptions are the AD specific options for the restore task being
	RestoreTaskId *int64            `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"` //Specifies the ID of the existing Restore Task to update.
	SqlOptions    SqlOptionsEnum    `json:"sqlOptions,omitempty" form:"sqlOptions,omitempty"`       //Specifies the sql options to update the Restore Task with.
}

/*
 * Structure for the custom type UpdateSourcesForPrincipalsParams
 */
type UpdateSourcesForPrincipalsParams struct {
	SourcesForPrincipals []*SourceForPrincipalParam `json:"sourcesForPrincipals,omitempty" form:"sourcesForPrincipals,omitempty"` //Array of Principals, Sources and Views.
}

/*
 * Structure for the custom type UpdateUserQuotaSettingsForView
 */
type UpdateUserQuotaSettingsForView struct {
	DefaultUserQuotaPolicy          *QuotaPolicy `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"`                   //Specifies a quota limit that can be optionally applied to Views and
	EnableUserQuota                 *bool        `json:"enableUserQuota,omitempty" form:"enableUserQuota,omitempty"`                                 //If set, it enables/disables the user quota overrides for a view.
	InheritDefaultPolicyFromViewbox *bool        `json:"inheritDefaultPolicyFromViewbox,omitempty" form:"inheritDefaultPolicyFromViewbox,omitempty"` //If set to true, the default_policy in view metadata will be cleared and
	ViewName                        *string      `json:"viewName,omitempty" form:"viewName,omitempty"`                                               //View name of input view.
}

/*
 * Structure for the custom type UpdateViewAliasParam
 */
type UpdateViewAliasParam struct {
	AliasName              *string          `json:"aliasName,omitempty" form:"aliasName,omitempty"`                           //Name of the alias to be updated.
	EnableSmbEncryption    *bool            `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`       //Specifies the SMB encryption for the View Alias. If set, it enables the
	EnableSmbViewDiscovery *bool            `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view alias for SMB.
	EnforceSmbEncryption   *bool            `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`     //Specifies the SMB encryption for all the sessions for the View Alias.
	SharePermissions       []*SmbPermission `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`             //Specifies a list of share level permissions.
	SubnetWhitelist        []*Subnet        `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`               //Specifies a list of Subnets with IP addresses that have permissions to
}

/*
 * Structure for the custom type UpdateViewParam
 */
type UpdateViewParam struct {
	AccessSids                      *[]string                `json:"accessSids,omitempty" form:"accessSids,omitempty"`                                           //Array of Security Identifiers (SIDs)
	AntivirusScanConfig             *AntivirusScanConfig     `json:"antivirusScanConfig,omitempty" form:"antivirusScanConfig,omitempty"`                         //Specifies the antivirus scan config settings for this View.
	Description                     *string                  `json:"description,omitempty" form:"description,omitempty"`                                         //Specifies an optional text description about the View.
	EnableFilerAuditLogging         *bool                    `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"`                 //Specifies if Filer Audit Logging is enabled for this view.
	EnableMixedModePermissions      *bool                    `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"`           //If set, mixed mode (NFS and SMB) access is enabled for this view.
	EnableNfsViewDiscovery          *bool                    `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"`                   //If set, it enables discovery of view for NFS.
	EnableOfflineCaching            *bool                    `json:"enableOfflineCaching,omitempty" form:"enableOfflineCaching,omitempty"`                       //Specifies whether to enable offline file caching of the view.
	EnableSmbAccessBasedEnumeration *bool                    `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
	EnableSmbEncryption             *bool                    `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`                         //Specifies the SMB encryption for the View. If set, it enables the SMB
	EnableSmbViewDiscovery          *bool                    `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"`                   //If set, it enables discovery of view for SMB.
	EnforceSmbEncryption            *bool                    `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`                       //Specifies the SMB encryption for all the sessions for the View.
	FileExtensionFilter             *FileExtensionFilter     `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"`                         //TODO: Write general description for this field
	FileLockConfig                  *FileLevelDataLockConfig `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"`                                   //Specifies a config to lock files in a view - to protect from malicious or
	LogicalQuota                    *QuotaPolicy             `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"`                                       //Specifies an optional logical quota limit (in bytes) for the usage allowed
	NfsRootPermissions              *NfsRootPermissions      `json:"nfsRootPermissions,omitempty" form:"nfsRootPermissions,omitempty"`                           //Specifies the config of NFS root permission of a view file system.
	OverrideGlobalWhitelist         *bool                    `json:"overrideGlobalWhitelist,omitempty" form:"overrideGlobalWhitelist,omitempty"`                 //Specifies whether view level client subnet whitelist overrides cluster and
	ProtocolAccess                  ProtocolAccessEnum       `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"`                                   //Specifies the supported Protocols for the View.
	Qos                             *QoS                     `json:"qos,omitempty" form:"qos,omitempty"`                                                         //Specifies the Quality of Service (QoS) Policy for the View.
	SecurityMode                    SecurityModeEnum         `json:"securityMode,omitempty" form:"securityMode,omitempty"`                                       //Specifies the security mode used for this view.
	SharePermissions                []*SmbPermission         `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`                               //Specifies a list of share level permissions.
	SmbPermissionsInfo              *SmbPermissionsInfo      `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"`                           //Specifies information about SMB permissions.
	StoragePolicyOverride           *StoragePolicyOverride   `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"`                     //Specifies if inline deduplication and compression settings inherited from
	SubnetWhitelist                 []*Subnet                `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`                                 //Array of Subnets.
	TenantId                        *string                  `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                               //Optional tenant id who has access to this View.
}

/*
 * Structure for the custom type UpgradeClusterParameters
 */
type UpgradeClusterParameters struct {
	TargetSwVersion string `json:"targetSwVersion" form:"targetSwVersion"` //Specifies the target software version. If specified, all Nodes on the
}

/*
 * Structure for the custom type UpgradeClusterResult
 */
type UpgradeClusterResult struct {
	Message   *string `json:"message,omitempty" form:"message,omitempty"`     //Specifies a message describing the result of the request.
	StatusUrl *string `json:"statusUrl,omitempty" form:"statusUrl,omitempty"` //Specifies the URL that can be queried to get the status of the operation
}

/*
 * Structure for the custom type UpgradeNodeParameters
 */
type UpgradeNodeParameters struct {
	NodeIds             *[]int64 `json:"nodeIds,omitempty" form:"nodeIds,omitempty"`                         //Specifies a list of IDs of additional nodes to be upgraded. These must
	TargetSwVersion     *string  `json:"targetSwVersion,omitempty" form:"targetSwVersion,omitempty"`         //Specifies the target software version. The node that the request is sent
	UpgradeAllFreeNodes *bool    `json:"upgradeAllFreeNodes,omitempty" form:"upgradeAllFreeNodes,omitempty"` //Specifies whether or not to attempt to upgrade all free nodes which
	UpgradeSelf         *bool    `json:"upgradeSelf,omitempty" form:"upgradeSelf,omitempty"`                 //Specifies that the node that the request is being sent to should be
}

/*
 * Structure for the custom type UpgradeNodeResult
 */
type UpgradeNodeResult struct {
	Message   *string `json:"message,omitempty" form:"message,omitempty"`     //Specifies a message describing the result of the request.
	StatusUrl *string `json:"statusUrl,omitempty" form:"statusUrl,omitempty"` //Specifies a URL that can be queried to get the status of the operation
}

/*
 * Structure for the custom type UpgradePhysicalAgentsMessage
 */
type UpgradePhysicalAgentsMessage struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies the status message returned after initiating an upgrade request.
}

/*
 * Structure for the custom type UpgradePhysicalServerAgents
 */
type UpgradePhysicalServerAgents struct {
	AgentIds []int64 `json:"agentIds" form:"agentIds"` //Array of Agent Ids.
}

/*
 * Structure for the custom type UploadPackageResult
 */
type UploadPackageResult struct {
	Message *string `json:"message,omitempty" form:"message,omitempty"` //Specifies a message describing the result of the request to upload
}

/*
 * Structure for the custom type UsageAndPerformanceStats
 */
type UsageAndPerformanceStats struct {
	DataInBytes                    *int64   `json:"dataInBytes,omitempty" form:"dataInBytes,omitempty"`                                       //Data brought into the cluster. This is the usage before data reduction if
	DataInBytesAfterReduction      *int64   `json:"dataInBytesAfterReduction,omitempty" form:"dataInBytesAfterReduction,omitempty"`           //Morphed Usage before data is replicated to other nodes as per RF or
	MinUsablePhysicalCapacityBytes *int64   `json:"minUsablePhysicalCapacityBytes,omitempty" form:"minUsablePhysicalCapacityBytes,omitempty"` //Specifies the minimum usable capacity available
	NumBytesRead                   *int64   `json:"numBytesRead,omitempty" form:"numBytesRead,omitempty"`                                     //Provides the total number of bytes read in the last 30 seconds.
	NumBytesWritten                *int64   `json:"numBytesWritten,omitempty" form:"numBytesWritten,omitempty"`                               //Provides the total number of bytes written in the last 30 second.
	PhysicalCapacityBytes          *int64   `json:"physicalCapacityBytes,omitempty" form:"physicalCapacityBytes,omitempty"`                   //Provides the total physical capacity in bytes as computed
	ReadIos                        *int64   `json:"readIos,omitempty" form:"readIos,omitempty"`                                               //Provides the number of Read IOs that occurred in the last 30 seconds.
	ReadLatencyMsecs               *float64 `json:"readLatencyMsecs,omitempty" form:"readLatencyMsecs,omitempty"`                             //Provides the Read latency in milliseconds for the Read IOs that occurred
	SystemCapacityBytes            *int64   `json:"systemCapacityBytes,omitempty" form:"systemCapacityBytes,omitempty"`                       //Provides the total available capacity as computed by
	SystemUsageBytes               *int64   `json:"systemUsageBytes,omitempty" form:"systemUsageBytes,omitempty"`                             //Provides the usage of bytes, as computed by the Linux 'statfs' command,
	TotalPhysicalRawUsageBytes     *int64   `json:"totalPhysicalRawUsageBytes,omitempty" form:"totalPhysicalRawUsageBytes,omitempty"`         //Provides the usage of bytes, as computed by the Cohesity Cluster,
	TotalPhysicalUsageBytes        *int64   `json:"totalPhysicalUsageBytes,omitempty" form:"totalPhysicalUsageBytes,omitempty"`               //Provides the total capacity, as computed by the Cohesity Cluster,
	WriteIos                       *int64   `json:"writeIos,omitempty" form:"writeIos,omitempty"`                                             //Provides the number of Write IOs that occurred in the last 30 seconds.
	WriteLatencyMsecs              *float64 `json:"writeLatencyMsecs,omitempty" form:"writeLatencyMsecs,omitempty"`                           //Provides the Write latency in milliseconds for the Write IOs that occurred
}

/*
 * Structure for the custom type UsageSchemaInfo
 */
type UsageSchemaInfo struct {
	SchemaInfoList []*SchemaInfo `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"` //Specifies the list of the schema info for an entity.
}

/*
 * Structure for the custom type User
 */
type User struct {
	AdditionalGroupNames *[]string                  `json:"additionalGroupNames,omitempty" form:"additionalGroupNames,omitempty"` //Array of Additional Groups.
	AuthenticationType   AuthenticationTypeUserEnum `json:"authenticationType,omitempty" form:"authenticationType,omitempty"`     //Specifies the authentication type of the user.
	ClusterIdentifiers   []*ClusterIdentifier       `json:"clusterIdentifiers,omitempty" form:"clusterIdentifiers,omitempty"`     //Specifies the list of clusters this user has access to. If this is not
	CreatedTimeMsecs     *int64                     `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the user account
	Description          *string                    `json:"description,omitempty" form:"description,omitempty"`                   //Specifies a description about the user.
	Domain               *string                    `json:"domain,omitempty" form:"domain,omitempty"`                             //Specifies the fully qualified domain name (FQDN) of an Active Directory
	EffectiveTimeMsecs   *int64                     `json:"effectiveTimeMsecs,omitempty" form:"effectiveTimeMsecs,omitempty"`     //Specifies the epoch time in milliseconds when the user becomes
	EmailAddress         *string                    `json:"emailAddress,omitempty" form:"emailAddress,omitempty"`                 //Specifies the email address of the user.
	ExpiredTimeMsecs     *int64                     `json:"expiredTimeMsecs,omitempty" form:"expiredTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the user becomes
	GoogleAccount        *GoogleAccountInfo         `json:"googleAccount,omitempty" form:"googleAccount,omitempty"`               //Google Account Information of a Helios BaaS user.
	IdpUserInfo          *IdpUserInfo               `json:"idpUserInfo,omitempty" form:"idpUserInfo,omitempty"`                   //Specifies an IdP User's information logged in using an IdP.
	LastUpdatedTimeMsecs *int64                     `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user account was last
	Password             *string                    `json:"password,omitempty" form:"password,omitempty"`                         //Specifies the password of this user.
	Preferences          *Preferences               `json:"preferences,omitempty" form:"preferences,omitempty"`                   //TODO: Write general description for this field
	PrimaryGroupName     *string                    `json:"primaryGroupName,omitempty" form:"primaryGroupName,omitempty"`         //Specifies the name of the primary group of this User.
	PrivilegeIds         *[]PrivilegeIdUserEnum     `json:"privilegeIds,omitempty" form:"privilegeIds,omitempty"`                 //Array of Privileges.
	Restricted           *bool                      `json:"restricted,omitempty" form:"restricted,omitempty"`                     //Whether the user is a restricted user. A restricted user can only view
	Roles                *[]string                  `json:"roles,omitempty" form:"roles,omitempty"`                               //Array of Roles.
	S3AccessKeyId        *string                    `json:"s3AccessKeyId,omitempty" form:"s3AccessKeyId,omitempty"`               //Specifies the S3 Account Access Key ID.
	S3AccountId          *string                    `json:"s3AccountId,omitempty" form:"s3AccountId,omitempty"`                   //Specifies the S3 Account Canonical User ID.
	S3SecretKey          *string                    `json:"s3SecretKey,omitempty" form:"s3SecretKey,omitempty"`                   //Specifies the S3 Account Secret Key.
	SalesforceAccount    *SalesforceAccountInfo     `json:"salesforceAccount,omitempty" form:"salesforceAccount,omitempty"`       //Salesforce Account Information of a Helios user.
	Sid                  *string                    `json:"sid,omitempty" form:"sid,omitempty"`                                   //Specifies the unique Security ID (SID) of the user.
	TenantId             *string                    `json:"tenantId,omitempty" form:"tenantId,omitempty"`                         //Specifies the effective Tenant ID of the user.
	Username             *string                    `json:"username,omitempty" form:"username,omitempty"`                         //Specifies the login name of the user.
}

/*
 * Structure for the custom type UserDeleteParameters
 */
type UserDeleteParameters struct {
	Domain   *string   `json:"domain,omitempty" form:"domain,omitempty"`     //Specifies the domain associated with the users to delete.
	TenantId *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the tenant for which the the users are to be deleted.
	Users    *[]string `json:"users,omitempty" form:"users,omitempty"`       //Array of Users.
}

/*
 * Structure for the custom type UserId
 */
type UserId struct {
	Sid     *string `json:"sid,omitempty" form:"sid,omitempty"`         //If interested in a user via smb_client, include SID.
	UnixUid *int64  `json:"unixUid,omitempty" form:"unixUid,omitempty"` //If interested in a user via unix-identifier, include UnixUid.
}

/*
 * Structure for the custom type UserIdMapping
 */
type UserIdMapping struct {
	CentrifyZoneMapping     *CentrifyZone           `json:"centrifyZoneMapping,omitempty" form:"centrifyZoneMapping,omitempty"`         //Specifies the properties associated to a Centrify zone of an Active
	CustomAttributesMapping *CustomUnixIdAttributes `json:"customAttributesMapping,omitempty" form:"customAttributesMapping,omitempty"` //Specifies the custom attributes when mapping type is set to
	FixedMapping            *FixedUnixIdMapping     `json:"fixedMapping,omitempty" form:"fixedMapping,omitempty"`                       //Specifies the fields when mapping type is set to 'kFixed'. It maps all
	Type                    TypeUserIdMappingEnum   `json:"type,omitempty" form:"type,omitempty"`                                       //Specifies the mapping type used.
}

/*
 * Structure for the custom type UserInfo
 */
type UserInfo struct {
	Domain   *string `json:"domain,omitempty" form:"domain,omitempty"`     //Specifies domain name of the user.
	Sid      *string `json:"sid,omitempty" form:"sid,omitempty"`           //Specifies unique Security ID (SID) of the user.
	UserName *string `json:"userName,omitempty" form:"userName,omitempty"` //Specifies user name of the user.
}

/*
 * Structure for the custom type UserInformation
 */
type UserInformation struct {
	IncludeSubtenantObjects *bool                    `json:"includeSubtenantObjects,omitempty" form:"includeSubtenantObjects,omitempty"` //Whether objects owned by subtenants should be returned. This would
	PulseAttributeVec       []*KeyValuePair          `json:"pulseAttributeVec,omitempty" form:"pulseAttributeVec,omitempty"`             //Specifies the KeyValuePair that client (eg. Iris) wants to persist along
	SidVec                  []*ClusterConfigProtoSID `json:"sidVec,omitempty" form:"sidVec,omitempty"`                                   //If specified, only the objects associated with these SIDs should be
	TenantIdVec             *[]string                `json:"tenantIdVec,omitempty" form:"tenantIdVec,omitempty"`                         //If specified, only the objects associated with this tenant should be
}

/*
 * Structure for the custom type UserParameters
 */
type UserParameters struct {
	AdditionalGroupNames *[]string                        `json:"additionalGroupNames,omitempty" form:"additionalGroupNames,omitempty"` //Array of Additional Groups.
	ClusterIdentifiers   []*ClusterIdentifier             `json:"clusterIdentifiers,omitempty" form:"clusterIdentifiers,omitempty"`     //Specifies the list of clusters this user has access to. If this is not
	Description          *string                          `json:"description,omitempty" form:"description,omitempty"`                   //Specifies a description about the user.
	Domain               *string                          `json:"domain,omitempty" form:"domain,omitempty"`                             //Specifies the fully qualified domain name (FQDN) of an Active Directory
	EffectiveTimeMsecs   *int64                           `json:"effectiveTimeMsecs,omitempty" form:"effectiveTimeMsecs,omitempty"`     //Specifies the epoch time in milliseconds when the user becomes
	EmailAddress         *string                          `json:"emailAddress,omitempty" form:"emailAddress,omitempty"`                 //Specifies the email address of the user.
	ExpiredTimeMsecs     *int64                           `json:"expiredTimeMsecs,omitempty" form:"expiredTimeMsecs,omitempty"`         //Specifies the epoch time in milliseconds when the user becomes
	Password             *string                          `json:"password,omitempty" form:"password,omitempty"`                         //Specifies the password of this user.
	PrimaryGroupName     *string                          `json:"primaryGroupName,omitempty" form:"primaryGroupName,omitempty"`         //Specifies the name of the primary group of this User.
	PrivilegeIds         *[]PrivilegeIdUserParametersEnum `json:"privilegeIds,omitempty" form:"privilegeIds,omitempty"`                 //Array of Privileges.
	Restricted           *bool                            `json:"restricted,omitempty" form:"restricted,omitempty"`                     //Whether the user is a restricted user. A restricted user can only view
	Roles                *[]string                        `json:"roles,omitempty" form:"roles,omitempty"`                               //Array of Roles.
	Username             *string                          `json:"username,omitempty" form:"username,omitempty"`                         //Specifies the login name of the user.
}

/*
 * Structure for the custom type UserQuota
 */
type UserQuota struct {
	QuotaPolicy *QuotaPolicy `json:"quotaPolicy,omitempty" form:"quotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
	Sid         *string      `json:"sid,omitempty" form:"sid,omitempty"`                 //If interested in a user via smb_client, include SID.
	UnixUid     *int64       `json:"unixUid,omitempty" form:"unixUid,omitempty"`         //If interested in a user via unix-identifier, include UnixUid.
}

/*
 * Structure for the custom type UserQuotaAndUsage
 */
type UserQuotaAndUsage struct {
	QuotaPolicy *QuotaPolicy `json:"quotaPolicy,omitempty" form:"quotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
	Sid         *string      `json:"sid,omitempty" form:"sid,omitempty"`                 //If interested in a user via smb_client, include SID.
	UnixUid     *int64       `json:"unixUid,omitempty" form:"unixUid,omitempty"`         //If interested in a user via unix-identifier, include UnixUid.
	UsageBytes  *int64       `json:"usageBytes,omitempty" form:"usageBytes,omitempty"`   //Current logical usage of user id in the input view.
}

/*
 * Structure for the custom type UserQuotaSettings
 */
type UserQuotaSettings struct {
	DefaultUserQuotaPolicy *QuotaPolicy `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
	EnableUserQuota        *bool        `json:"enableUserQuota,omitempty" form:"enableUserQuota,omitempty"`               //If set, it enables/disables the user quota overrides for a view.
}

/*
 * Structure for the custom type UserQuotaSummaryForUser
 */
type UserQuotaSummaryForUser struct {
	NumViewsAboveAlertThreshold *int64 `json:"numViewsAboveAlertThreshold,omitempty" form:"numViewsAboveAlertThreshold,omitempty"` //Number of views in which user has exceeded alert threshold limit.
	NumViewsAboveHardLimit      *int64 `json:"numViewsAboveHardLimit,omitempty" form:"numViewsAboveHardLimit,omitempty"`           //Number of views in which the user has exceeded hard limit.
	TotalNumViews               *int64 `json:"totalNumViews,omitempty" form:"totalNumViews,omitempty"`                             //Total number of views in which the user has a quota policy specified
}

/*
 * Structure for the custom type UserQuotaSummaryForView
 */
type UserQuotaSummaryForView struct {
	DefaultUserQuotaPolicy      *QuotaPolicy `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"`           //Specifies a quota limit that can be optionally applied to Views and
	NumUsersAboveAlertThreshold *int64       `json:"numUsersAboveAlertThreshold,omitempty" form:"numUsersAboveAlertThreshold,omitempty"` //Number of users who has exceeded their specified alert limit.
	NumUsersAboveHardLimit      *int64       `json:"numUsersAboveHardLimit,omitempty" form:"numUsersAboveHardLimit,omitempty"`           //Number of users who has exceeded their specified quota hard limit.
	TotalNumUsers               *int64       `json:"totalNumUsers,omitempty" form:"totalNumUsers,omitempty"`                             //Total number of users who has either a user quota policy override
}

/*
 * Structure for the custom type UuidConfigProto
 */
type UuidConfigProto struct {
	PreserveUuid *bool `json:"preserveUuid,omitempty" form:"preserveUuid,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type VcloudDirectorInfo
 */
type VcloudDirectorInfo struct {
	Endpoint *string `json:"endpoint,omitempty" form:"endpoint,omitempty"` //vCenter endpoint.
	Name     *string `json:"name,omitempty" form:"name,omitempty"`         //vCenter name.
}

/*
 * Structure for the custom type VmwareBackupEnvParams
 */
type VmwareBackupEnvParams struct {
	AllowCrashConsistentSnapshot *bool                       `json:"allowCrashConsistentSnapshot,omitempty" form:"allowCrashConsistentSnapshot,omitempty"` //Whether to fallback to take a crash-consistent snapshot incase taking
	AllowVmsWithPhysicalRdmDisks *bool                       `json:"allowVmsWithPhysicalRdmDisks,omitempty" form:"allowVmsWithPhysicalRdmDisks,omitempty"` //Physical RDM disks cannot be backed up using VADP. By default the backups
	VmwareDiskExclusionInfo      []*VmwareDiskExclusionProto `json:"vmwareDiskExclusionInfo,omitempty" form:"vmwareDiskExclusionInfo,omitempty"`           //List of Virtual Disk(s) to be excluded from the backup job. These disks
}

/*
 * Structure for the custom type VmwareBackupSourceParams
 */
type VmwareBackupSourceParams struct {
	SourceAppParams         *SourceAppParams            `json:"sourceAppParams,omitempty" form:"sourceAppParams,omitempty"`                 //This message contains params specific to application running on the source
	VmCredentials           *Credentials                `json:"vmCredentials,omitempty" form:"vmCredentials,omitempty"`                     //Specifies credentials to access a target source.
	VmwareDiskExclusionInfo []*VmwareDiskExclusionProto `json:"vmwareDiskExclusionInfo,omitempty" form:"vmwareDiskExclusionInfo,omitempty"` //List of Virtual Disk(s) to be excluded from the backup job for the source.
}

/*
 * Structure for the custom type VmwareDiskExclusionProto
 */
type VmwareDiskExclusionProto struct {
	ControllerBusNumber *int64  `json:"controllerBusNumber,omitempty" form:"controllerBusNumber,omitempty"` //Controller's bus-id controlling the virtual disk in question.
	ControllerType      *string `json:"controllerType,omitempty" form:"controllerType,omitempty"`           //Controller's type (SCSI, IDE etc).
	UnitNumber          *int64  `json:"unitNumber,omitempty" form:"unitNumber,omitempty"`                   //Disk unit number to identify the virtual disk within a controller.
}

/*
 * Structure for the custom type VmwareObjectId
 */
type VmwareObjectId struct {
	MorItem *string `json:"morItem,omitempty" form:"morItem,omitempty"` //Specifies the Managed Object Reference Item.
	MorType *string `json:"morType,omitempty" form:"morType,omitempty"` //Specifies the Managed Object Reference Type.
	Uuid    *string `json:"uuid,omitempty" form:"uuid,omitempty"`       //Specifies a Universally Unique Identifier (UUID) of a VMware Object.
}

/*
 * Structure for the custom type VmwareProtectionSource
 */
type VmwareProtectionSource struct {
	AgentId            *int64                             `json:"agentId,omitempty" form:"agentId,omitempty"`                       //Specifies the id of the persistent agent.
	Agents             []*AgentInformation                `json:"agents,omitempty" form:"agents,omitempty"`                         //Specifies the list of agent information on the Virtual Machine.
	ConnectionState    ConnectionStateEnum                `json:"connectionState,omitempty" form:"connectionState,omitempty"`       //Specifies the connection state of the Object and are only valid for
	DatastoreInfo      *DatastoreInfo                     `json:"datastoreInfo,omitempty" form:"datastoreInfo,omitempty"`           //TODO: Write general description for this field
	FolderType         FolderTypeEnum                     `json:"folderType,omitempty" form:"folderType,omitempty"`                 //Specifies the folder type for the 'kFolder' Object.
	HasPersistentAgent *bool                              `json:"hasPersistentAgent,omitempty" form:"hasPersistentAgent,omitempty"` //Set to true if a persistent agent is running on the Virtual Machine.
	HostType           HostTypeVmwareProtectionSourceEnum `json:"hostType,omitempty" form:"hostType,omitempty"`                     //Specifies the host type for the 'kVirtualMachine' Object.
	Id                 *VmwareObjectId                    `json:"id,omitempty" form:"id,omitempty"`                                 //Specifies a unique Protection Source id across Cohesity Clusters.
	IsVmTemplate       *bool                              `json:"isVmTemplate,omitempty" form:"isVmTemplate,omitempty"`             //IsTemplate specifies if the VM is a template or not.
	Name               *string                            `json:"name,omitempty" form:"name,omitempty"`                             //Specifies a human readable name of the Protection Source.
	TagAttributes      []*TagAttribute                    `json:"tagAttributes,omitempty" form:"tagAttributes,omitempty"`           //Specifies the optional list of VM Tag attributes associated with this
	ToolsRunningStatus ToolsRunningStatusEnum             `json:"toolsRunningStatus,omitempty" form:"toolsRunningStatus,omitempty"` //Specifies the status of VMware Tools for the guest OS on the VM.
	Type               TypeVmwareProtectionSourceEnum     `json:"type,omitempty" form:"type,omitempty"`                             //Specifies the type of managed Object in a VMware Protection Source.
	VcloudDirectorInfo []*VcloudDirectorInfo              `json:"vCloudDirectorInfo,omitempty" form:"vCloudDirectorInfo,omitempty"` //Specifies an array of vCenters to be registered
	VirtualDisks       []*VirtualDiskInfo                 `json:"virtualDisks,omitempty" form:"virtualDisks,omitempty"`             //Specifies an array of virtual disks that are part of the Virtual Machine.
}

/*
 * Structure for the custom type Value
 */
type Value struct {
	Data *ValueData `json:"data,omitempty" form:"data,omitempty"` //Specifies the fields to store data of a given type.
	Type *int64     `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of value.
}

/*
 * Structure for the custom type ValueData
 */
type ValueData struct {
	BytesValue  *[]int64 `json:"bytesValue,omitempty" form:"bytesValue,omitempty"`   //Specifies the field to store an array of bytes if the current
	DoubleValue *float64 `json:"doubleValue,omitempty" form:"doubleValue,omitempty"` //Specifies the field to store data if the current data type is a
	Int64Value  *int64   `json:"int64Value,omitempty" form:"int64Value,omitempty"`   //Specifies the field to store data if the current data type is a
	StringValue *string  `json:"stringValue,omitempty" form:"stringValue,omitempty"` //Specifies the field to store data if the current data type is a
}

/*
 * Structure for the custom type Vault
 */
type Vault struct {
	CaTrustedCertificate           *string                    `json:"caTrustedCertificate,omitempty" form:"caTrustedCertificate,omitempty"`                     //Specifies the CA (certificate authority) trusted certificate.
	ClientCertificate              *string                    `json:"clientCertificate,omitempty" form:"clientCertificate,omitempty"`                           //Specifies the client CA  certificate. This certificate is in pem format.
	ClientPrivateKey               *string                    `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"`                             //Specifies the client private key. This certificate is in pem format.
	CompressionPolicy              CompressionPolicyVaultEnum `json:"compressionPolicy,omitempty" form:"compressionPolicy,omitempty"`                           //Specifies whether to send data to the Vault in a
	Config                         *VaultConfig               `json:"config,omitempty" form:"config,omitempty"`                                                 //Specifies the settings required to connect to a specific Vault type.
	CustomerManagingEncryptionKeys *bool                      `json:"customerManagingEncryptionKeys,omitempty" form:"customerManagingEncryptionKeys,omitempty"` //Specifies whether to manage the encryption key manually or let the
	DedupEnabled                   *bool                      `json:"dedupEnabled,omitempty" form:"dedupEnabled,omitempty"`                                     //Specifies whether to deduplicate data before sending it to the Vault.
	Description                    *string                    `json:"description,omitempty" form:"description,omitempty"`                                       //Specifies a description about the Vault.
	DesiredWalLocation             DesiredWalLocationEnum     `json:"desiredWalLocation,omitempty" form:"desiredWalLocation,omitempty"`                         //Desired location for write ahead logs(wal).
	EncryptionKeyFileDownloaded    *bool                      `json:"encryptionKeyFileDownloaded,omitempty" form:"encryptionKeyFileDownloaded,omitempty"`       //Specifies if the encryption key file has been downloaded using the
	EncryptionPolicy               EncryptionPolicyVaultEnum  `json:"encryptionPolicy,omitempty" form:"encryptionPolicy,omitempty"`                             //Specifies whether to send and store data in an encrypted format.
	ExternalTargetType             ExternalTargetTypeEnum     `json:"externalTargetType,omitempty" form:"externalTargetType,omitempty"`                         //Specifies the type of Vault.
	FullArchiveIntervalDays        *int64                     `json:"fullArchiveIntervalDays,omitempty" form:"fullArchiveIntervalDays,omitempty"`               //Specifies the number days between full archives to the Vault.
	Id                             *int64                     `json:"id,omitempty" form:"id,omitempty"`                                                         //Specifies an id that identifies the Vault.
	IncrementalArchivesEnabled     *bool                      `json:"incrementalArchivesEnabled,omitempty" form:"incrementalArchivesEnabled,omitempty"`         //Specifies whether to perform incremental archival when sending data
	KeyFileDownloadTimeUsecs       *int64                     `json:"keyFileDownloadTimeUsecs,omitempty" form:"keyFileDownloadTimeUsecs,omitempty"`             //Specifies the time (in microseconds) when the encryption key file was
	KeyFileDownloadUser            *string                    `json:"keyFileDownloadUser,omitempty" form:"keyFileDownloadUser,omitempty"`                       //Specifies the user who downloaded the encryption key from the
	Name                           *string                    `json:"name,omitempty" form:"name,omitempty"`                                                     //Specifies the name of the Vault.
	Type                           TypeVaultEnum              `json:"type,omitempty" form:"type,omitempty"`                                                     //Specifies the type of Vault.
	UsageType                      UsageTypeEnum              `json:"usageType,omitempty" form:"usageType,omitempty"`                                           //Specifies the usage type of the Vault.
	VaultBandwidthLimits           *VaultBandwidthLimits      `json:"vaultBandwidthLimits,omitempty" form:"vaultBandwidthLimits,omitempty"`                     //VaultBandwidthLimits represents the network bandwidth limits
}

/*
 * Structure for the custom type VaultBandwidthLimits
 */
type VaultBandwidthLimits struct {
	Download *BandwidthLimit `json:"download,omitempty" form:"download,omitempty"` //Specifies settings for limiting the data transfer rate between
	Upload   *BandwidthLimit `json:"upload,omitempty" form:"upload,omitempty"`     //Specifies settings for limiting the data transfer rate between
}

/*
 * Structure for the custom type VaultConfig
 */
type VaultConfig struct {
	Amazon     *AmazonCloudCredentials `json:"amazon,omitempty" form:"amazon,omitempty"`         //Specifies the cloud credentials to connect to a Amazon
	Azure      *AzureCloudCredentials  `json:"azure,omitempty" form:"azure,omitempty"`           //Specifies the cloud credentials to connect to a Microsoft
	BucketName *string                 `json:"bucketName,omitempty" form:"bucketName,omitempty"` //Specifies the name of a storage location of the Vault,
	Google     *GoogleCloudCredentials `json:"google,omitempty" form:"google,omitempty"`         //Specifies the cloud credentials to connect to a Google service account.
	Nas        *NasCredentials         `json:"nas,omitempty" form:"nas,omitempty"`               //Specifies the server credentials to connect to a NetApp server.
	Oracle     *OracleCloudCredentials `json:"oracle,omitempty" form:"oracle,omitempty"`         //Specifies the Oracle Cloud Credentials to connect to an Oracle S3 Compatible
	Qstar      *QStarServerCredentials `json:"qstar,omitempty" form:"qstar,omitempty"`           //Specifies the server credentials to connect to a QStar service
}

/*
 * Structure for the custom type VaultEncryptionKey
 */
type VaultEncryptionKey struct {
	ClusterName       *string      `json:"clusterName,omitempty" form:"clusterName,omitempty"`             //Specifies the name of the source Cohesity Cluster
	EncryptionKeyData *string      `json:"encryptionKeyData,omitempty" form:"encryptionKeyData,omitempty"` //Specifies the encryption key data corresponding to the specified keyUid.
	KeyUid            *UniversalId `json:"keyUid,omitempty" form:"keyUid,omitempty"`                       //Specifies the universal id of the Data Encryption Key.
	VaultId           *int64       `json:"vaultId,omitempty" form:"vaultId,omitempty"`                     //Specifies the id of the Vault whose data is encrypted by
	VaultName         *string      `json:"vaultName,omitempty" form:"vaultName,omitempty"`                 //Specifies the name of the Vault whose data is encrypted by this key.
}

/*
 * Structure for the custom type VaultParamsRestoreParams
 */
type VaultParamsRestoreParams struct {
	Glacier *VaultParamsRestoreParamsGlacier `json:"glacier,omitempty" form:"glacier,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type VaultParamsRestoreParamsGlacier
 */
type VaultParamsRestoreParamsGlacier struct {
	RetrievalType *int64 `json:"retrievalType,omitempty" form:"retrievalType,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type VaultProviderStatsByEnv
 */
type VaultProviderStatsByEnv struct {
	Count       *int64                                 `json:"count,omitempty" form:"count,omitempty"`             //Specifies the count of the objects of the specified environment.
	Environment EnvironmentVaultProviderStatsByEnvEnum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type.
	Size        *int64                                 `json:"size,omitempty" form:"size,omitempty"`               //Specifies the size of the entities of the specified environment.
}

/*
 * Structure for the custom type VaultProviderStatsInfo
 */
type VaultProviderStatsInfo struct {
	ChangeRate           *int64                              `json:"changeRate,omitempty" form:"changeRate,omitempty"`                     //Specifies the relative change of size of entities on the vault.
	ClusterId            *int64                              `json:"clusterId,omitempty" form:"clusterId,omitempty"`                       //Specifies the cluster id.
	ClusterIncarnationId *int64                              `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the cluster incarnation id.
	ClusterName          *string                             `json:"clusterName,omitempty" form:"clusterName,omitempty"`                   //Specifies the cluster name.
	ReadBandwidth        *int64                              `json:"readBandwidth,omitempty" form:"readBandwidth,omitempty"`               //Specifies the average read bandwidth over the last 24 hours.
	StatsByEnv           []*VaultProviderStatsByEnv          `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"`                     //Specifies the stats by environments.
	VaultGroup           VaultGroupEnum                      `json:"vaultGroup,omitempty" form:"vaultGroup,omitempty"`                     //Specifies the cloud vendor type.
	VaultId              *int64                              `json:"vaultId,omitempty" form:"vaultId,omitempty"`                           //Specifies the Vault id.
	VaultType            VaultTypeVaultProviderStatsInfoEnum `json:"vaultType,omitempty" form:"vaultType,omitempty"`                       //Specifies the External Target type.
	Vaultname            *string                             `json:"vaultname,omitempty" form:"vaultname,omitempty"`                       //Specifies the Vault name.
	WriteBandwidth       *int64                              `json:"writeBandwidth,omitempty" form:"writeBandwidth,omitempty"`             //Specifies the average write bandwidth over the last 24 hours.
}

/*
 * Structure for the custom type VaultRunInfo
 */
type VaultRunInfo struct {
	Count     *int64 `json:"count,omitempty" form:"count,omitempty"`         //Specifies the count of runs that ended in the specified state between the start time passed in and the current timestamp.
	Timestamp *int64 `json:"timestamp,omitempty" form:"timestamp,omitempty"` //Specifies the Unix timestamp at which the run entered the specified state.
}

/*
 * Structure for the custom type VaultRunStatsSummary
 */
type VaultRunStatsSummary struct {
	FailureTimeSeries []*VaultRunInfo `json:"failureTimeSeries,omitempty" form:"failureTimeSeries,omitempty"` //Specifies the time series for the failed runs that ended in the given time frame.
	NumFailedRuns     *int64          `json:"numFailedRuns,omitempty" form:"numFailedRuns,omitempty"`         //Specifies the number of runs that ended in failure during the given time frame.
	NumInProgressRuns *int64          `json:"numInProgressRuns,omitempty" form:"numInProgressRuns,omitempty"` //Specifies the number of runs that were currently in progress at the time that the API call was made.
	NumQueuedRuns     *int64          `json:"numQueuedRuns,omitempty" form:"numQueuedRuns,omitempty"`         //Specifies the number of runs that were currently queued at the time that the API call was made.
	NumSuccessfulRuns *int64          `json:"numSuccessfulRuns,omitempty" form:"numSuccessfulRuns,omitempty"` //Specifies the number of runs that ended in success during the given time frame.
	SuccessTimeSeries []*VaultRunInfo `json:"successTimeSeries,omitempty" form:"successTimeSeries,omitempty"` //Specifies the time series for the successful runs that ended in the given time frame.
}

/*
 * Structure for the custom type VaultStats
 */
type VaultStats struct {
	AwsUsageBytes    *int64            `json:"awsUsageBytes,omitempty" form:"awsUsageBytes,omitempty"`       //Specifies the usage on AWS vaults.
	AzureUsageBytes  *int64            `json:"azureUsageBytes,omitempty" form:"azureUsageBytes,omitempty"`   //Specifies the usage on Azure vaults.
	GcpUsageBytes    *int64            `json:"gcpUsageBytes,omitempty" form:"gcpUsageBytes,omitempty"`       //Specifies the usage on GCP vaults.
	NasUsageBytes    *int64            `json:"nasUsageBytes,omitempty" form:"nasUsageBytes,omitempty"`       //Specifies the usage on NAS vaults.
	OracleUsageBytes *int64            `json:"oracleUsageBytes,omitempty" form:"oracleUsageBytes,omitempty"` //Specifies the usage on Oracle vaults.
	QstarUsageBytes  *int64            `json:"qstarUsageBytes,omitempty" form:"qstarUsageBytes,omitempty"`   //Specifies the usage on QStar Tape vaults.
	S3cUsageBytes    *int64            `json:"s3cUsageBytes,omitempty" form:"s3cUsageBytes,omitempty"`       //Specifies the usage on S3 Compatible vaults.
	VaultStatsList   []*VaultStatsInfo `json:"vaultStatsList,omitempty" form:"vaultStatsList,omitempty"`     //Specifies the stats of all vaults on the cluster.
}

/*
 * Structure for the custom type VaultStatsInfo
 */
type VaultStatsInfo struct {
	Id         *int64                 `json:"id,omitempty" form:"id,omitempty"`                 //Specifies the Vault Id.
	Name       *string                `json:"name,omitempty" form:"name,omitempty"`             //Specifies the Vault name.
	Type       TypeVaultStatsInfoEnum `json:"type,omitempty" form:"type,omitempty"`             //Specifies the Vault type.
	UsageBytes *int64                 `json:"usageBytes,omitempty" form:"usageBytes,omitempty"` //Specifies the bytes used by the Vault.
}

/*
 * Structure for the custom type View
 */
type View struct {
	AccessSids                      *[]string                `json:"accessSids,omitempty" form:"accessSids,omitempty"`                                           //Array of Security Identifiers (SIDs)
	Aliases                         []*ViewAliasInfo         `json:"aliases,omitempty" form:"aliases,omitempty"`                                                 //Aliases created for the view. A view alias allows a directory path inside
	AllSmbMountPaths                *[]string                `json:"allSmbMountPaths,omitempty" form:"allSmbMountPaths,omitempty"`                               //Array of SMB Paths.
	AntivirusScanConfig             *AntivirusScanConfig     `json:"antivirusScanConfig,omitempty" form:"antivirusScanConfig,omitempty"`                         //Specifies the antivirus scan config settings for this View.
	BasicMountPath                  *string                  `json:"basicMountPath,omitempty" form:"basicMountPath,omitempty"`                                   //Specifies the NFS mount path of the View (without the hostname
	CaseInsensitiveNamesEnabled     *bool                    `json:"caseInsensitiveNamesEnabled,omitempty" form:"caseInsensitiveNamesEnabled,omitempty"`         //Specifies whether to support case insensitive file/folder names. This
	CreateTimeMsecs                 *int64                   `json:"createTimeMsecs,omitempty" form:"createTimeMsecs,omitempty"`                                 //Specifies the time that the View was created in milliseconds.
	DataLockExpiryUsecs             *int64                   `json:"dataLockExpiryUsecs,omitempty" form:"dataLockExpiryUsecs,omitempty"`                         //DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If
	Description                     *string                  `json:"description,omitempty" form:"description,omitempty"`                                         //Specifies an optional text description about the View.
	EnableFilerAuditLogging         *bool                    `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"`                 //Specifies if Filer Audit Logging is enabled for this view.
	EnableMixedModePermissions      *bool                    `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"`           //If set, mixed mode (NFS and SMB) access is enabled for this view.
	EnableNfsViewDiscovery          *bool                    `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"`                   //If set, it enables discovery of view for NFS.
	EnableOfflineCaching            *bool                    `json:"enableOfflineCaching,omitempty" form:"enableOfflineCaching,omitempty"`                       //Specifies whether to enable offline file caching of the view.
	EnableSmbAccessBasedEnumeration *bool                    `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
	EnableSmbEncryption             *bool                    `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`                         //Specifies the SMB encryption for the View. If set, it enables the SMB
	EnableSmbViewDiscovery          *bool                    `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"`                   //If set, it enables discovery of view for SMB.
	EnforceSmbEncryption            *bool                    `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`                       //Specifies the SMB encryption for all the sessions for the View.
	FileExtensionFilter             *FileExtensionFilter     `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"`                         //TODO: Write general description for this field
	FileLockConfig                  *FileLevelDataLockConfig `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"`                                   //Specifies a config to lock files in a view - to protect from malicious or
	IsTargetForMigratedData         *bool                    `json:"isTargetForMigratedData,omitempty" form:"isTargetForMigratedData,omitempty"`                 //Specifies if a view contains migrated data.
	LogicalQuota                    *QuotaPolicy             `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"`                                       //Specifies an optional logical quota limit (in bytes) for the usage allowed
	LogicalUsageBytes               *int64                   `json:"logicalUsageBytes,omitempty" form:"logicalUsageBytes,omitempty"`                             //LogicalUsageBytes is the logical usage in bytes for the view.
	Name                            *string                  `json:"name,omitempty" form:"name,omitempty"`                                                       //Specifies the name of the View.
	NfsMountPath                    *string                  `json:"nfsMountPath,omitempty" form:"nfsMountPath,omitempty"`                                       //Specifies the path for mounting this View as an NFS share.
	NfsRootPermissions              *NfsRootPermissions      `json:"nfsRootPermissions,omitempty" form:"nfsRootPermissions,omitempty"`                           //Specifies the config of NFS root permission of a view file system.
	OverrideGlobalWhitelist         *bool                    `json:"overrideGlobalWhitelist,omitempty" form:"overrideGlobalWhitelist,omitempty"`                 //Specifies whether view level client subnet whitelist overrides cluster and
	ProtocolAccess                  ProtocolAccessEnum       `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"`                                   //Specifies the supported Protocols for the View.
	Qos                             *QoS                     `json:"qos,omitempty" form:"qos,omitempty"`                                                         //Specifies the Quality of Service (QoS) Policy for the View.
	S3AccessPath                    *string                  `json:"s3AccessPath,omitempty" form:"s3AccessPath,omitempty"`                                       //Specifies the path to access this View as an S3 share.
	S3KeyMappingConfig              S3KeyMappingConfigEnum   `json:"s3KeyMappingConfig,omitempty" form:"s3KeyMappingConfig,omitempty"`                           //Specifies the S3 key mapping config of the view. This parameter can only
	SecurityMode                    SecurityModeEnum         `json:"securityMode,omitempty" form:"securityMode,omitempty"`                                       //Specifies the security mode used for this view.
	SharePermissions                []*SmbPermission         `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`                               //Specifies a list of share level permissions.
	SmbMountPath                    *string                  `json:"smbMountPath,omitempty" form:"smbMountPath,omitempty"`                                       //Specifies the main path for mounting this View as an SMB share.
	SmbPermissionsInfo              *SmbPermissionsInfo      `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"`                           //Specifies information about SMB permissions.
	Stats                           *ViewStats               `json:"stats,omitempty" form:"stats,omitempty"`                                                     //Provides statistics about the View.
	StoragePolicyOverride           *StoragePolicyOverride   `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"`                     //Specifies if inline deduplication and compression settings inherited from
	SubnetWhitelist                 []*Subnet                `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`                                 //Array of Subnets.
	TenantId                        *string                  `json:"tenantId,omitempty" form:"tenantId,omitempty"`                                               //Optional tenant id who has access to this View.
	ViewBoxId                       *int64                   `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"`                                             //Specifies the id of the Storage Domain (View Box) where the View is stored.
	ViewBoxName                     *string                  `json:"viewBoxName,omitempty" form:"viewBoxName,omitempty"`                                         //Specifies the name of the Storage Domain (View Box) where the View is stored.
	ViewId                          *int64                   `json:"viewId,omitempty" form:"viewId,omitempty"`                                                   //Specifies an id of the View assigned by the Cohesity Cluster.
	ViewProtection                  *ViewProtection          `json:"viewProtection,omitempty" form:"viewProtection,omitempty"`                                   //Specifies information about the Protection Jobs that are protecting the
}

/*
 * Structure for the custom type ViewAlias
 */
type ViewAlias struct {
	AliasName              *string          `json:"aliasName,omitempty" form:"aliasName,omitempty"`                           //Alias name.
	EnableSmbEncryption    *bool            `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"`       //Specifies the SMB encryption for the View Alias. If set, it enables the
	EnableSmbViewDiscovery *bool            `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view alias for SMB.
	EnforceSmbEncryption   *bool            `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"`     //Specifies the SMB encryption for all the sessions for the View Alias.
	SharePermissions       []*SmbPermission `json:"sharePermissions,omitempty" form:"sharePermissions,omitempty"`             //Specifies a list of share level permissions.
	SubnetWhitelist        []*Subnet        `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"`               //Specifies a list of Subnets with IP addresses that have permissions to
	ViewName               *string          `json:"viewName,omitempty" form:"viewName,omitempty"`                             //View name.
	ViewPath               *string          `json:"viewPath,omitempty" form:"viewPath,omitempty"`                             //View path for the alias.
}

/*
 * Structure for the custom type ViewAliasInfo
 */
type ViewAliasInfo struct {
	AliasName             *string                     `json:"aliasName,omitempty" form:"aliasName,omitempty"`                         //Alias name.
	ClientSubnetWhitelist []*ClusterConfigProtoSubnet `json:"clientSubnetWhitelist,omitempty" form:"clientSubnetWhitelist,omitempty"` //List of external client subnet IPs that are allowed to access the share.
	SmbConfig             *AliasSmbConfig             `json:"smbConfig,omitempty" form:"smbConfig,omitempty"`                         //Message defining SMB config for IRIS. SMB config contains SMB encryption
	ViewPath              *string                     `json:"viewPath,omitempty" form:"viewPath,omitempty"`                           //View path for the alias.
}

/*
 * Structure for the custom type ViewBox
 */
type ViewBox struct {
	AdDomainName                    *string                 `json:"adDomainName,omitempty" form:"adDomainName,omitempty"`                                       //Specifies an active directory domain that this view box is mapped to.
	ClientSubnetWhiteList           []*Subnet               `json:"clientSubnetWhiteList,omitempty" form:"clientSubnetWhiteList,omitempty"`                     //Array of Subnets.
	CloudDownWaterfallThresholdPct  *int64                  `json:"cloudDownWaterfallThresholdPct,omitempty" form:"cloudDownWaterfallThresholdPct,omitempty"`   //Specifies the cloud down water-fall threshold percentage. This indicates
	CloudDownWaterfallThresholdSecs *int64                  `json:"cloudDownWaterfallThresholdSecs,omitempty" form:"cloudDownWaterfallThresholdSecs,omitempty"` //Specifies the cloud down water-fall threshold seconds. This indicates
	ClusterPartitionId              int64                   `json:"clusterPartitionId" form:"clusterPartitionId"`                                               //Specifies the Cluster Partition id where the Storage Domain (View Box) is
	ClusterPartitionName            *string                 `json:"clusterPartitionName,omitempty" form:"clusterPartitionName,omitempty"`                       //Specifies the Cohesity Cluster name where the Storage Domain (View Box) is
	DefaultUserQuotaPolicy          *QuotaPolicy            `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"`                   //Specifies an optional quota policy/limits that are inherited by all users
	DefaultViewQuotaPolicy          *QuotaPolicy            `json:"defaultViewQuotaPolicy,omitempty" form:"defaultViewQuotaPolicy,omitempty"`                   //Specifies an optional default logical quota limit (in bytes)
	Id                              *int64                  `json:"id,omitempty" form:"id,omitempty"`                                                           //Specifies the Id of the Storage Domain (View Box).
	LdapProviderId                  *int64                  `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"`                                   //When set, the following provides the LDAP provider the view box is
	Name                            string                  `json:"name" form:"name"`                                                                           //Specifies the name of the Storage Domain (View Box).
	PhysicalQuota                   *QuotaPolicy            `json:"physicalQuota,omitempty" form:"physicalQuota,omitempty"`                                     //Specifies an optional quota limit (in bytes) for the physical
	RemovalState                    RemovalStateViewBoxEnum `json:"removalState,omitempty" form:"removalState,omitempty"`                                       //Specifies the current removal state of the Storage Domain (View Box).
	S3BucketsAllowed                *bool                   `json:"s3BucketsAllowed,omitempty" form:"s3BucketsAllowed,omitempty"`                               //Specifies whether creation of a S3 bucket is allowed in this
	SchemaInfoList                  []*SchemaInfo           `json:"schemaInfoList,omitempty" form:"schemaInfoList,omitempty"`                                   //Specifies the time series schema info of the view box.
	Stats                           *ViewBoxStats           `json:"stats,omitempty" form:"stats,omitempty"`                                                     //Provides statistics about the Storage Domain (View Box).
	StoragePolicy                   *StoragePolicy          `json:"storagePolicy,omitempty" form:"storagePolicy,omitempty"`                                     //Specifies the storage options applied to a Storage Domain (View Box).
	TenantIdVec                     *[]string               `json:"tenantIdVec,omitempty" form:"tenantIdVec,omitempty"`                                         //Optional ids for the tenants that this view box belongs. This must be
	TreatFileSyncAsDataSync         *bool                   `json:"treatFileSyncAsDataSync,omitempty" form:"treatFileSyncAsDataSync,omitempty"`                 //If 'true', when the Cohesity Cluster is writing to a file, the
}

/*
 * Structure for the custom type ViewBoxPairInfo
 */
type ViewBoxPairInfo struct {
	LocalViewBoxId    *int64  `json:"localViewBoxId,omitempty" form:"localViewBoxId,omitempty"`       //Specifies the id of the Storage Domain (View Box) on the local Cluster.
	LocalViewBoxName  *string `json:"localViewBoxName,omitempty" form:"localViewBoxName,omitempty"`   //Specifies the name of the Storage Domain (View Box) on the local Cluster.
	RemoteViewBoxId   *int64  `json:"remoteViewBoxId,omitempty" form:"remoteViewBoxId,omitempty"`     //Specifies the id of the Storage Domain (View Box) on the remote Cluster.
	RemoteViewBoxName *string `json:"remoteViewBoxName,omitempty" form:"remoteViewBoxName,omitempty"` //Specifies the name of the Storage Domain (View Box) on the remote Cluster.
}

/*
 * Structure for the custom type ViewBoxStats
 */
type ViewBoxStats struct {
	CloudUsagePerfStats *UsageAndPerformanceStats `json:"cloudUsagePerfStats,omitempty" form:"cloudUsagePerfStats,omitempty"` //Provides usage and performance statistics
	DataUsageStats      *DataUsageStats           `json:"dataUsageStats,omitempty" form:"dataUsageStats,omitempty"`           //Specifies the data usage metric of the data stored on the Cohesity
	Id                  *int64                    `json:"id,omitempty" form:"id,omitempty"`                                   //Specifies the id of the Storage Domain (View Box).
	LocalUsagePerfStats *UsageAndPerformanceStats `json:"localUsagePerfStats,omitempty" form:"localUsagePerfStats,omitempty"` //Provides usage and performance statistics
	LogicalStats        *LogicalStats             `json:"logicalStats,omitempty" form:"logicalStats,omitempty"`               //Provides logical statistics for logical entities such as Clusters
	UsagePerfStats      *UsageAndPerformanceStats `json:"usagePerfStats,omitempty" form:"usagePerfStats,omitempty"`           //Provides usage and performance statistics
}

/*
 * Structure for the custom type ViewIdMappingProtoProtocolAccessInfo
 */
type ViewIdMappingProtoProtocolAccessInfo struct {
	IscsiAccess *int64 `json:"iscsiAccess,omitempty" form:"iscsiAccess,omitempty"` //Access control for iSCSI protocol for this view.
	NfsAccess   *int64 `json:"nfsAccess,omitempty" form:"nfsAccess,omitempty"`     //Access control for NFS protocol for this view.
	S3Access    *int64 `json:"s3Access,omitempty" form:"s3Access,omitempty"`       //Access control for S3 protocol for this view.
	SmbAccess   *int64 `json:"smbAccess,omitempty" form:"smbAccess,omitempty"`     //Access control for SMB protocol for this view.
}

/*
 * Structure for the custom type ViewParams
 */
type ViewParams struct {
	ClientSubnetWhitelistVec []*ClusterConfigProtoSubnet              `json:"clientSubnetWhitelistVec,omitempty" form:"clientSubnetWhitelistVec,omitempty"` //List of external client subnets from where requests will be received for
	DisableNfsAccess         *bool                                    `json:"disableNfsAccess,omitempty" form:"disableNfsAccess,omitempty"`                 //Whether to disable NFS access in the new view.
	ProtocolAccessInfo       *ViewIdMappingProtoProtocolAccessInfo    `json:"protocolAccessInfo,omitempty" form:"protocolAccessInfo,omitempty"`             //TODO: Write general description for this field
	QosMappingVec            []*ClusterConfigProtoQoSMapping          `json:"qosMappingVec,omitempty" form:"qosMappingVec,omitempty"`                       //The qos mappings (if any) for the new view.
	StoragePolicyOverride    *ClusterConfigProtoStoragePolicyOverride `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"`       //TODO: Write general description for this field
	ViewDescription          *string                                  `json:"viewDescription,omitempty" form:"viewDescription,omitempty"`                   //The description to be applied to the new view.
	WormLockExpiryUsecs      *int64                                   `json:"wormLockExpiryUsecs,omitempty" form:"wormLockExpiryUsecs,omitempty"`           //This value 'worm_lock_expiry_usecs' if specified will be set on the cloned
}

/*
 * Structure for the custom type ViewProtection
 */
type ViewProtection struct {
	Inactive        *bool                `json:"inactive,omitempty" form:"inactive,omitempty"`               //Specifies if this View is an inactive View that was created on this
	MagnetoEntityId *int64               `json:"magnetoEntityId,omitempty" form:"magnetoEntityId,omitempty"` //Specifies the id of the Protection Source that is using this View.
	ProtectionJobs  []*ProtectionJobInfo `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"`   //Array of Protection Jobs.
}

/*
 * Structure for the custom type ViewProtectionSource
 */
type ViewProtectionSource struct {
	Id   *UniversalId                 `json:"id,omitempty" form:"id,omitempty"`     //Specifies a unique id of a Protection Source for a View.
	Name *string                      `json:"name,omitempty" form:"name,omitempty"` //Specifies a human readable name of the Protection Source of a View.
	Type TypeViewProtectionSourceEnum `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed Object in a View Protection Source
}

/*
 * Structure for the custom type ViewProtocolStats
 */
type ViewProtocolStats struct {
	Protocols *[]ProtocolViewProtocolStatsEnum `json:"protocols,omitempty" form:"protocols,omitempty"` //Specifies the protocols supported on these Views.
	SizeBytes *int64                           `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Specifies the size of all the Views in bytes which are using the specified protocol.
	ViewCount *int64                           `json:"viewCount,omitempty" form:"viewCount,omitempty"` //Specifies the number of Views which are using the specified protocol.
}

/*
 * Structure for the custom type ViewStatInfo
 */
type ViewStatInfo struct {
	ClusterId             *int64                      `json:"clusterId,omitempty" form:"clusterId,omitempty"`                         //Specifies the cluster Id.
	ClusterIncarnationId  *int64                      `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"`   //Specifies the cluster Incarnation Id.
	DataReadBytes         *int64                      `json:"dataReadBytes,omitempty" form:"dataReadBytes,omitempty"`                 //Specifies the data read in bytes.
	DataWrittenBytes      *int64                      `json:"dataWrittenBytes,omitempty" form:"dataWrittenBytes,omitempty"`           //Specifies the data written in bytes.
	LogicalUsedBytes      *int64                      `json:"logicalUsedBytes,omitempty" form:"logicalUsedBytes,omitempty"`           //Specifies the logical size used in bytes.
	PeakReadThroughput    *int64                      `json:"peakReadThroughput,omitempty" form:"peakReadThroughput,omitempty"`       //Specifies the peak data read in bytes per second in the last day.
	PeakWriteThroughput   *int64                      `json:"peakWriteThroughput,omitempty" form:"peakWriteThroughput,omitempty"`     //Specifies the peak data written in bytes per second in the last day.
	PhysicalUsedBytes     *int64                      `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"`         //Specifies the physical size used in bytes.
	Protocols             *[]ProtocolViewStatInfoEnum `json:"protocols,omitempty" form:"protocols,omitempty"`                         //Specifies the protocols of this view.
	StorageReductionRatio *float64                    `json:"storageReductionRatio,omitempty" form:"storageReductionRatio,omitempty"` //Specifies the storage reduction ratio.
	ViewId                *int64                      `json:"viewId,omitempty" form:"viewId,omitempty"`                               //Specifies the view Id.
	ViewName              *string                     `json:"viewName,omitempty" form:"viewName,omitempty"`                           //Specifies the view name.
}

/*
 * Structure for the custom type ViewStats
 */
type ViewStats struct {
	DataUsageStats *DataUsageStats `json:"dataUsageStats,omitempty" form:"dataUsageStats,omitempty"` //Specifies the data usage metric of the data stored on the Cohesity
	Id             *int64          `json:"id,omitempty" form:"id,omitempty"`                         //Specifies the id of the View.
}

/*
 * Structure for the custom type ViewStatsSnapshot
 */
type ViewStatsSnapshot struct {
	Timestamp     *int64          `json:"timestamp,omitempty" form:"timestamp,omitempty"`         //Specifies the unix time in milliseconds when these values were generated
	ViewStatsList []*ViewStatInfo `json:"viewStatsList,omitempty" form:"viewStatsList,omitempty"` //Specifies the list of Views and their statistics at the given timestamp.
}

/*
 * Structure for the custom type ViewUserQuotaParameters
 */
type ViewUserQuotaParameters struct {
	UserQuotaPolicy *UserQuota `json:"userQuotaPolicy,omitempty" form:"userQuotaPolicy,omitempty"` //Specifies the quota policy applied to a user.
	ViewName        *string    `json:"viewName,omitempty" form:"viewName,omitempty"`               //View name of input view.
}

/*
 * Structure for the custom type ViewUserQuotas
 */
type ViewUserQuotas struct {
	Cookie                  *string                  `json:"cookie,omitempty" form:"cookie,omitempty"`                                   //This cookie can be used in the succeeding call to list user quotas and
	QuotaAndUsageInAllViews []*QuotaAndUsageInView   `json:"quotaAndUsageInAllViews,omitempty" form:"quotaAndUsageInAllViews,omitempty"` //The quota and usage information for a user in all his views.
	SummaryForUser          *UserQuotaSummaryForUser `json:"summaryForUser,omitempty" form:"summaryForUser,omitempty"`                   //Speifies the summary of quota information for a particular user.
	SummaryForView          *UserQuotaSummaryForView `json:"summaryForView,omitempty" form:"summaryForView,omitempty"`                   //Specifies the user quota summary information/result for a view.
	UserQuotaSettings       *UserQuotaSettings       `json:"userQuotaSettings,omitempty" form:"userQuotaSettings,omitempty"`             //Specifies the quota settings parameters for a particular user.
	UsersQuotaAndUsage      []*UserQuotaAndUsage     `json:"usersQuotaAndUsage,omitempty" form:"usersQuotaAndUsage,omitempty"`           //The list of user quota policies/overrides and usages.
}

/*
 * Structure for the custom type VirtualDiskId
 */
type VirtualDiskId struct {
	ControllerBusNumber *int64  `json:"controllerBusNumber,omitempty" form:"controllerBusNumber,omitempty"` //Controller's bus-id controlling the virtual disk in question.
	ControllerType      *string `json:"controllerType,omitempty" form:"controllerType,omitempty"`           //Controller's type (SCSI, IDE etc).
	DiskId              *string `json:"diskId,omitempty" form:"diskId,omitempty"`                           //Original disk id. This is sufficient to identify the disk information, but
	UnitNumber          *int64  `json:"unitNumber,omitempty" form:"unitNumber,omitempty"`                   //Disk unit number to identify the virtual disk within a controller.
}

/*
 * Structure for the custom type VirtualDiskIdInformation
 */
type VirtualDiskIdInformation struct {
	BusNumber      *int64  `json:"busNumber,omitempty" form:"busNumber,omitempty"`           //Specifies the Id of the controller bus that controls the disk.
	ControllerType *string `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
	DiskId         *string `json:"diskId,omitempty" form:"diskId,omitempty"`                 //Specfies the uuid of the virtual disk.
	UnitNumber     *int64  `json:"unitNumber,omitempty" form:"unitNumber,omitempty"`         //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type VirtualDiskInfo
 */
type VirtualDiskInfo struct {
	BusNumber      *int64  `json:"busNumber,omitempty" form:"busNumber,omitempty"`           //Specifies the Id of the controller bus that controls the disk.
	ControllerType *string `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
	Filename       *string `json:"filename,omitempty" form:"filename,omitempty"`             //Specifies the host file name used as the virtual disk.
	UnitNumber     *int64  `json:"unitNumber,omitempty" form:"unitNumber,omitempty"`         //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type VirtualDiskInformation
 */
type VirtualDiskInformation struct {
	BusNumber       *int64            `json:"busNumber,omitempty" form:"busNumber,omitempty"`             //Specifies the Id of the controller bus that controls the disk.
	ControllerType  *string           `json:"controllerType,omitempty" form:"controllerType,omitempty"`   //Specifies the controller type like SCSI, or IDE etc.
	DiskId          *string           `json:"diskId,omitempty" form:"diskId,omitempty"`                   //Specifies original disk id. This is sufficient to identify the disk
	DiskLocation    *ProtectionSource `json:"diskLocation,omitempty" form:"diskLocation,omitempty"`       //Specifies a generic structure that represents a node
	DiskSizeInBytes *int64            `json:"diskSizeInBytes,omitempty" form:"diskSizeInBytes,omitempty"` //Specifies size of the virtual disk in bytes.
	FilePath        *string           `json:"filePath,omitempty" form:"filePath,omitempty"`               //Specifies the original file path if applicable.
	MountPoints     *[]string         `json:"mountPoints,omitempty" form:"mountPoints,omitempty"`         //Specifies the list of mount points.
	UnitNumber      *int64            `json:"unitNumber,omitempty" form:"unitNumber,omitempty"`           //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type VirtualDiskMapping
 */
type VirtualDiskMapping struct {
	DiskToOverwrite  *VirtualDiskIdInformation `json:"diskToOverwrite,omitempty" form:"diskToOverwrite,omitempty"`   //Specifies information about virtual disk which includes disk uuid,
	SourceDisk       *VirtualDiskIdInformation `json:"sourceDisk,omitempty" form:"sourceDisk,omitempty"`             //Specifies information about virtual disk which includes disk uuid,
	TargetLocationId *int64                    `json:"targetLocationId,omitempty" form:"targetLocationId,omitempty"` //Specifies the target location information, for e.g. a datastore in
}

/*
 * Structure for the custom type VirtualDiskMappingResponse
 */
type VirtualDiskMappingResponse struct {
	DiskToOverwrite *VirtualDiskIdInformation `json:"diskToOverwrite,omitempty" form:"diskToOverwrite,omitempty"` //Specifies information about virtual disk which includes disk uuid,
	SourceDisk      *VirtualDiskIdInformation `json:"sourceDisk,omitempty" form:"sourceDisk,omitempty"`           //Specifies information about virtual disk which includes disk uuid,
	TargetLocation  *ProtectionSource         `json:"targetLocation,omitempty" form:"targetLocation,omitempty"`   //Specifies a generic structure that represents a node
}

/*
 * Structure for the custom type VirtualDiskRecoverTaskState
 */
type VirtualDiskRecoverTaskState struct {
	Error                      *RequestError               `json:"error,omitempty" form:"error,omitempty"`                                           //Details about the Error.
	IsInstantRecoveryFinished  *bool                       `json:"isInstantRecoveryFinished,omitempty" form:"isInstantRecoveryFinished,omitempty"`   //Specifies if instant recovery of the virtual disk is complete.
	TaskState                  TaskStateEnum               `json:"taskState,omitempty" form:"taskState,omitempty"`                                   //Specifies the current state of the restore virtual disks task.
	VirtualDiskRestoreResponse *VirtualDiskRestoreResponse `json:"virtualDiskRestoreResponse,omitempty" form:"virtualDiskRestoreResponse,omitempty"` //Specifies the parameters to recover virtual disks of a vm with full
}

/*
 * Structure for the custom type VirtualDiskRestoreParameters
 */
type VirtualDiskRestoreParameters struct {
	PowerOffVmBeforeRecovery *bool                 `json:"powerOffVmBeforeRecovery,omitempty" form:"powerOffVmBeforeRecovery,omitempty"` //Specifies whether to power off the VM before recovering virtual disks.
	PowerOnVmAfterRecovery   *bool                 `json:"powerOnVmAfterRecovery,omitempty" form:"powerOnVmAfterRecovery,omitempty"`     //Specifies whether to power on the VM after recovering virtual disks.
	TargetSourceId           *int64                `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"`                     //Specifies the target entity to which the disks should be attached.
	VirtualDiskMappings      []*VirtualDiskMapping `json:"virtualDiskMappings,omitempty" form:"virtualDiskMappings,omitempty"`           //Specifies the list of virtual disks mappings.
}

/*
 * Structure for the custom type VirtualDiskRestoreResponse
 */
type VirtualDiskRestoreResponse struct {
	PowerOffVmBeforeRecovery *bool                         `json:"powerOffVmBeforeRecovery,omitempty" form:"powerOffVmBeforeRecovery,omitempty"` //Specifies whether to power off the VM before recovering virtual disks.
	PowerOnVmAfterRecovery   *bool                         `json:"powerOnVmAfterRecovery,omitempty" form:"powerOnVmAfterRecovery,omitempty"`     //Specifies whether to power on the VM after recovering virtual disks.
	TargetSource             *ProtectionSource             `json:"targetSource,omitempty" form:"targetSource,omitempty"`                         //Specifies a generic structure that represents a node
	VirtualDiskMappings      []*VirtualDiskMappingResponse `json:"virtualDiskMappings,omitempty" form:"virtualDiskMappings,omitempty"`           //Specifies the list of virtual disks mappings.
}

/*
 * Structure for the custom type VirtualNodeConfiguration
 */
type VirtualNodeConfiguration struct {
	NodeId *int64  `json:"nodeId,omitempty" form:"nodeId,omitempty"` //Specifies the Node ID for this node.
	NodeIp *string `json:"nodeIp,omitempty" form:"nodeIp,omitempty"` //Specifies the Node IP address for this node.
}

/*
 * Structure for the custom type Vlan
 */
type Vlan struct {
	AddToClusterPartition *bool     `json:"addToClusterPartition,omitempty" form:"addToClusterPartition,omitempty"` //Specifies whether to add the VLAN IPs to the cluster partition
	AllTenantAccess       *bool     `json:"allTenantAccess,omitempty" form:"allTenantAccess,omitempty"`             //Specifies if this VLAN can be used by all tenants without explicit
	Description           *string   `json:"description,omitempty" form:"description,omitempty"`                     //Specifies a description of the VLAN.
	Gateway               *string   `json:"gateway,omitempty" form:"gateway,omitempty"`                             //Specifies the Gateway of the VLAN.
	Hostname              *string   `json:"hostname,omitempty" form:"hostname,omitempty"`                           //Specifies the hostname of the VLAN.
	Id                    *int64    `json:"id,omitempty" form:"id,omitempty"`                                       //Specifies the id of the VLAN.
	IfaceGroupName        *string   `json:"ifaceGroupName,omitempty" form:"ifaceGroupName,omitempty"`               //Specifies the interface group name of the VLAN. It is in the format of
	InterfaceName         *string   `json:"interfaceName,omitempty" form:"interfaceName,omitempty"`                 //Specifies the interface name of the VLAN.
	Ips                   *[]string `json:"ips,omitempty" form:"ips,omitempty"`                                     //Array of IPs.
	Subnet                *Subnet   `json:"subnet,omitempty" form:"subnet,omitempty"`                               //Specifies the subnet of the VLAN.
	TenantId              *string   `json:"tenantId,omitempty" form:"tenantId,omitempty"`                           //Optional tenant id that this vlan belongs to.
	VlanName              *string   `json:"vlanName,omitempty" form:"vlanName,omitempty"`                           //Specifies the VLAN name of the vlanId.
}

/*
 * Structure for the custom type VlanParameters
 */
type VlanParameters struct {
	DisableVlan   *bool   `json:"disableVlan,omitempty" form:"disableVlan,omitempty"`     //Specifies whether to use the VIPs even when VLANs are configured on the
	InterfaceName *string `json:"interfaceName,omitempty" form:"interfaceName,omitempty"` //Specifies the physical interface group name to use for mounting
	Vlan          *int64  `json:"vlan,omitempty" form:"vlan,omitempty"`                   //Specifies the VLAN to use for mounting Cohesity's view on the remote
}

/*
 * Structure for the custom type VmVolumesInformation
 */
type VmVolumesInformation struct {
	FilesystemVolumes []*FilesystemVolume `json:"filesystemVolumes,omitempty" form:"filesystemVolumes,omitempty"` //Array of Filesystem Volumes.
}

/*
 * Structure for the custom type VmwareCloneParameters
 */
type VmwareCloneParameters struct {
	DatastoreFolderId *int64            `json:"datastoreFolderId,omitempty" form:"datastoreFolderId,omitempty"` //Specifies the folder where the restore datastore should be created.
	DetachNetwork     *bool             `json:"detachNetwork,omitempty" form:"detachNetwork,omitempty"`         //Specifies whether the network should be detached from the
	DisableNetwork    *bool             `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"`       //Specifies whether the network should be left in disabled state.
	NetworkId         *int64            `json:"networkId,omitempty" form:"networkId,omitempty"`                 //Specifies a network configuration to be attached to the cloned or
	NetworkMappings   []*NetworkMapping `json:"networkMappings,omitempty" form:"networkMappings,omitempty"`     //Specifies the parameters for mapping the source and target
	PoweredOn         *bool             `json:"poweredOn,omitempty" form:"poweredOn,omitempty"`                 //Specifies the power state of the cloned or recovered objects.
	Prefix            *string           `json:"prefix,omitempty" form:"prefix,omitempty"`                       //Specifies a prefix to prepended to the source object name to derive a
	ResourcePoolId    *int64            `json:"resourcePoolId,omitempty" form:"resourcePoolId,omitempty"`       //Specifies the resource pool where the cloned or recovered objects are
	Suffix            *string           `json:"suffix,omitempty" form:"suffix,omitempty"`                       //Specifies a suffix to appended to the original source object name
	VmFolderId        *int64            `json:"vmFolderId,omitempty" form:"vmFolderId,omitempty"`               //Specifies a folder where the VMs should be restored. This is applicable
}

/*
 * Structure for the custom type VmwareEnvJobParameters
 */
type VmwareEnvJobParameters struct {
	ExcludedDisks             []*DiskUnit `json:"excludedDisks,omitempty" form:"excludedDisks,omitempty"`                         //Specifies the list of Disks to be excluded from backing up. These disks
	FallbackToCrashConsistent *bool       `json:"fallbackToCrashConsistent,omitempty" form:"fallbackToCrashConsistent,omitempty"` //If true, takes a crash-consistent snapshot when app-consistent snapshot
	SkipPhysicalRdmDisks      *bool       `json:"skipPhysicalRdmDisks,omitempty" form:"skipPhysicalRdmDisks,omitempty"`           //If true, skip physical RDM disks when backing up VMs. Otherwise, backup
}

/*
 * Structure for the custom type VmwareRestoreParameters
 */
type VmwareRestoreParameters struct {
	DatastoreFolderId *int64            `json:"datastoreFolderId,omitempty" form:"datastoreFolderId,omitempty"` //Specifies the folder where the restore datastore should be created.
	DatastoreId       *int64            `json:"datastoreId,omitempty" form:"datastoreId,omitempty"`             //Specifies the datastore where the object's files should be
	DetachNetwork     *bool             `json:"detachNetwork,omitempty" form:"detachNetwork,omitempty"`         //Specifies whether the network should be detached from the
	DisableNetwork    *bool             `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"`       //Specifies whether the network should be left in disabled state.
	NetworkId         *int64            `json:"networkId,omitempty" form:"networkId,omitempty"`                 //Specifies a network configuration to be attached to the cloned or
	NetworkMappings   []*NetworkMapping `json:"networkMappings,omitempty" form:"networkMappings,omitempty"`     //Specifies the parameters for mapping the source and target
	PoweredOn         *bool             `json:"poweredOn,omitempty" form:"poweredOn,omitempty"`                 //Specifies the power state of the cloned or recovered objects.
	Prefix            *string           `json:"prefix,omitempty" form:"prefix,omitempty"`                       //Specifies a prefix to prepended to the source object name to derive a
	ResourcePoolId    *int64            `json:"resourcePoolId,omitempty" form:"resourcePoolId,omitempty"`       //Specifies the resource pool where the cloned or recovered objects are
	Suffix            *string           `json:"suffix,omitempty" form:"suffix,omitempty"`                       //Specifies a suffix to appended to the original source object name
	VmFolderId        *int64            `json:"vmFolderId,omitempty" form:"vmFolderId,omitempty"`               //Specifies a folder where the VMs should be restored. This is applicable
}

/*
 * Structure for the custom type VmwareSpecialParameters
 */
type VmwareSpecialParameters struct {
	ApplicationParameters *ApplicationParameters `json:"applicationParameters,omitempty" form:"applicationParameters,omitempty"` //TODO: Write general description for this field
	ExcludedDisks         []*DiskUnit            `json:"excludedDisks,omitempty" form:"excludedDisks,omitempty"`                 //Specifies the list of Disks to be excluded from backing up. These disks
	VmCredentials         *Credentials           `json:"vmCredentials,omitempty" form:"vmCredentials,omitempty"`                 //Specifies the administrator credentials to log in to the
}

/*
 * Structure for the custom type VolumeInfo
 */
type VolumeInfo struct {
	DiskVec        []*VolumeInfoDiskInfo        `json:"diskVec,omitempty" form:"diskVec,omitempty"`               //Information about all the disks and partitions needed to mount this
	DisplayName    *string                      `json:"displayName,omitempty" form:"displayName,omitempty"`       //Display name.
	FilesystemType *string                      `json:"filesystemType,omitempty" form:"filesystemType,omitempty"` //Filesystem on this volume.
	FsUuid         *string                      `json:"fsUuid,omitempty" form:"fsUuid,omitempty"`                 //Filesystem uuid.
	IsBootable     *bool                        `json:"isBootable,omitempty" form:"isBootable,omitempty"`         //Is this volume bootable?
	IsDedup        *bool                        `json:"isDedup,omitempty" form:"isDedup,omitempty"`               //Is this a dedup volume?
	IsSupported    *bool                        `json:"isSupported,omitempty" form:"isSupported,omitempty"`       //Is this a supported Volume (filesystem)?
	LvInfo         *VolumeInfoLogicalVolumeInfo `json:"lvInfo,omitempty" form:"lvInfo,omitempty"`                 //This is extra attribute which uniquely identifies a logical volume in LVM
	VolumeGuid     *string                      `json:"volumeGuid,omitempty" form:"volumeGuid,omitempty"`         //The guid of the volume represented by this virtual disk.
	VolumeType     *int64                       `json:"volumeType,omitempty" form:"volumeType,omitempty"`         //Whether this volume is simple, lvm or ldm.
}

/*
 * Structure for the custom type VolumeInfoDiskInfo
 */
type VolumeInfoDiskInfo struct {
	DiskFileName     *string                            `json:"diskFileName,omitempty" form:"diskFileName,omitempty"`         //Disk name. This is the vmdk names, and not the flat file name.
	DiskFormat       *int64                             `json:"diskFormat,omitempty" form:"diskFormat,omitempty"`             //Disk format type of this file.
	DiskUuid         *string                            `json:"diskUuid,omitempty" form:"diskUuid,omitempty"`                 //Disk uuid.
	PartitionType    *int64                             `json:"partitionType,omitempty" form:"partitionType,omitempty"`       //Disk partition type.
	PartitionVec     []*VolumeInfoDiskInfoPartitionInfo `json:"partitionVec,omitempty" form:"partitionVec,omitempty"`         //Information about all the partitions in this disk.
	PhysicalRangeVec []*VolumeInfoDiskInfoPhysicalRange `json:"physicalRangeVec,omitempty" form:"physicalRangeVec,omitempty"` //This disk is formed by following physical ranges.
	SectorSize       *int64                             `json:"sectorSize,omitempty" form:"sectorSize,omitempty"`             //Sector size of disk.
	VmdkSize         *int64                             `json:"vmdkSize,omitempty" form:"vmdkSize,omitempty"`                 //Disk size in bytes.
}

/*
 * Structure for the custom type VolumeInfoDiskInfoPartitionInfo
 */
type VolumeInfoDiskInfoPartitionInfo struct {
	Length            *int64  `json:"length,omitempty" form:"length,omitempty"`                       //Length of partition in bytes.
	PartitionNumber   *int64  `json:"partitionNumber,omitempty" form:"partitionNumber,omitempty"`     //Partition number.
	PartitionTypeUuid *string `json:"partitionTypeUuid,omitempty" form:"partitionTypeUuid,omitempty"` //Partition type uuid.
	PartitionUuid     *string `json:"partitionUuid,omitempty" form:"partitionUuid,omitempty"`         //Partition uuid.
	StartOffset       *int64  `json:"startOffset,omitempty" form:"startOffset,omitempty"`             //Start offset of partition in bytes.
}

/*
 * Structure for the custom type VolumeInfoDiskInfoPhysicalRange
 */
type VolumeInfoDiskInfoPhysicalRange struct {
	Length *int64 `json:"length,omitempty" form:"length,omitempty"` //Length of this range in bytes.
	Offset *int64 `json:"offset,omitempty" form:"offset,omitempty"` //Offset of this range in disk file from beginning of file.
}

/*
 * Structure for the custom type VolumeInfoLogicalVolumeInfo
 */
type VolumeInfoLogicalVolumeInfo struct {
	DeviceTree        *DeviceTree `json:"deviceTree,omitempty" form:"deviceTree,omitempty"`               //TODO: Write general description for this field
	LogicalVolumeName *string     `json:"logicalVolumeName,omitempty" form:"logicalVolumeName,omitempty"` //Logical volume name.
	LogicalVolumeUuid *string     `json:"logicalVolumeUuid,omitempty" form:"logicalVolumeUuid,omitempty"` //Logical volume uuid.
	VolumeGroupName   *string     `json:"volumeGroupName,omitempty" form:"volumeGroupName,omitempty"`     //Volume group name.
	VolumeGroupUuid   *string     `json:"volumeGroupUuid,omitempty" form:"volumeGroupUuid,omitempty"`     //Volume group uuid.
}

/*
 * Structure for the custom type VolumeSecurityInfo
 */
type VolumeSecurityInfo struct {
	GroupId     *int64    `json:"groupId,omitempty" form:"groupId,omitempty"`         //Specifies the Unix group ID for this volume. 0 indicates the root id.
	Permissions *string   `json:"permissions,omitempty" form:"permissions,omitempty"` //Specifies the Unix permission bits in octal string format.
	Style       StyleEnum `json:"style,omitempty" form:"style,omitempty"`             //Specifies the security style associated with this volume.
	UserId      *int64    `json:"userId,omitempty" form:"userId,omitempty"`           //Specifies the Unix user id for this volume. 0 indicates the root id.
}

/*
 * Structure for the custom type VserverNetworkInterface
 */
type VserverNetworkInterface struct {
	DataProtocols *[]DataProtocolEnum `json:"dataProtocols,omitempty" form:"dataProtocols,omitempty"` //Array of Data Protocols.
	IpAddress     *string             `json:"ipAddress,omitempty" form:"ipAddress,omitempty"`         //Specifies the IP address of this interface.
	Name          *string             `json:"name,omitempty" form:"name,omitempty"`                   //Specifies the name of this interface.
}

/*
 * Structure for the custom type WebHookDeliveryTarget
 */
type WebHookDeliveryTarget struct {
	CurlOptions    *string `json:"curlOptions,omitempty" form:"curlOptions,omitempty"`       //Specifies curl options used to invoke external api url defined above.
	ExternalApiUrl *string `json:"externalApiUrl,omitempty" form:"externalApiUrl,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type WindowsHostSnapshotParameters
 */
type WindowsHostSnapshotParameters struct {
	CopyOnlyBackup      *bool     `json:"copyOnlyBackup,omitempty" form:"copyOnlyBackup,omitempty"`           //Specifies whether to backup regardless of the state of each file's
	DisableMetadata     *bool     `json:"disableMetadata,omitempty" form:"disableMetadata,omitempty"`         //Specifies whether to disable fetching and storing of some metadata
	DisableNotification *bool     `json:"disableNotification,omitempty" form:"disableNotification,omitempty"` //Specifies whether to disable some notification steps when taking
	ExcludedVssWriters  *[]string `json:"excludedVssWriters,omitempty" form:"excludedVssWriters,omitempty"`   //Specifies a list of Windows VSS writers that are excluded from backups.
}

/*
 * Structure for the custom type WormRetentionProto
 */
type WormRetentionProto struct {
	PolicyType *int64 `json:"policyType,omitempty" form:"policyType,omitempty"` //The type of WORM policy set on this run. This field is irrelevant
}

/*
 * Structure for the custom type ADGuidPairADAttributeRestoreParam
 */
type ADGuidPairADAttributeRestoreParam struct {
	Destination *string `json:"destination,omitempty" form:"destination,omitempty"` //Destination guid in production AD object corresponding to source. If
	Source      *string `json:"source,omitempty" form:"source,omitempty"`           //Source guid string of an AD object in mounted AD snapshot. This cannot be
}
