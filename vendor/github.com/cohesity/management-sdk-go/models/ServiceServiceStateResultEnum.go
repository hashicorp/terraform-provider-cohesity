// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ServiceServiceStateResultEnum enum
 */
type ServiceServiceStateResultEnum int

/**
 * Value collection for ServiceServiceStateResultEnum enum
 */
const (
    ServiceServiceStateResult_KAPOLLO ServiceServiceStateResultEnum = 1 + iota
    ServiceServiceStateResult_KBRIDGE
    ServiceServiceStateResult_KGENIE
    ServiceServiceStateResult_KGENIEGOFER
    ServiceServiceStateResult_KMAGNETO
    ServiceServiceStateResult_KIRIS
    ServiceServiceStateResult_KIRISPROXY
    ServiceServiceStateResult_KSCRIBE
    ServiceServiceStateResult_KSTATS
    ServiceServiceStateResult_KYODA
    ServiceServiceStateResult_KALERTS
    ServiceServiceStateResult_KKEYCHAIN
    ServiceServiceStateResult_KLOGWATCHER
    ServiceServiceStateResult_KSTATSCOLLECTER
    ServiceServiceStateResult_KGANDALF
    ServiceServiceStateResult_KNEXUS
    ServiceServiceStateResult_KNEXUSPROXY
    ServiceServiceStateResult_KSTORAGEPROXY
    ServiceServiceStateResult_KTRICORDER
    ServiceServiceStateResult_KRTCLIENT
    ServiceServiceStateResult_KVAULTPROXY
    ServiceServiceStateResult_KSMBPROXY
    ServiceServiceStateResult_KBRIDGEPROXY
    ServiceServiceStateResult_KLIBRARIAN
    ServiceServiceStateResult_KGROOT
    ServiceServiceStateResult_KEAGLEAGENT
    ServiceServiceStateResult_KATHENA
    ServiceServiceStateResult_KBIFROSTBROKER
    ServiceServiceStateResult_KOS
)

func (r ServiceServiceStateResultEnum) MarshalJSON() ([]byte, error) {
    s := ServiceServiceStateResultEnumToValue(r)
    return json.Marshal(s)
}

func (r *ServiceServiceStateResultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ServiceServiceStateResultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ServiceServiceStateResultEnum to its string representation
 */
func ServiceServiceStateResultEnumToValue(serviceServiceStateResultEnum ServiceServiceStateResultEnum) string {
    switch serviceServiceStateResultEnum {
        case ServiceServiceStateResult_KAPOLLO:
    		return "kApollo"
        case ServiceServiceStateResult_KBRIDGE:
    		return "kBridge"
        case ServiceServiceStateResult_KGENIE:
    		return "kGenie"
        case ServiceServiceStateResult_KGENIEGOFER:
    		return "kGenieGofer"
        case ServiceServiceStateResult_KMAGNETO:
    		return "kMagneto"
        case ServiceServiceStateResult_KIRIS:
    		return "kIris"
        case ServiceServiceStateResult_KIRISPROXY:
    		return "kIrisProxy"
        case ServiceServiceStateResult_KSCRIBE:
    		return "kScribe"
        case ServiceServiceStateResult_KSTATS:
    		return "kStats"
        case ServiceServiceStateResult_KYODA:
    		return "kYoda"
        case ServiceServiceStateResult_KALERTS:
    		return "kAlerts"
        case ServiceServiceStateResult_KKEYCHAIN:
    		return "kKeychain"
        case ServiceServiceStateResult_KLOGWATCHER:
    		return "kLogWatcher"
        case ServiceServiceStateResult_KSTATSCOLLECTER:
    		return "kStatsCollecter"
        case ServiceServiceStateResult_KGANDALF:
    		return "kGandalf"
        case ServiceServiceStateResult_KNEXUS:
    		return "kNexus"
        case ServiceServiceStateResult_KNEXUSPROXY:
    		return "kNexusProxy"
        case ServiceServiceStateResult_KSTORAGEPROXY:
    		return "kStorageProxy"
        case ServiceServiceStateResult_KTRICORDER:
    		return "kTricorder"
        case ServiceServiceStateResult_KRTCLIENT:
    		return "kRtClient"
        case ServiceServiceStateResult_KVAULTPROXY:
    		return "kVaultProxy"
        case ServiceServiceStateResult_KSMBPROXY:
    		return "kSmbProxy"
        case ServiceServiceStateResult_KBRIDGEPROXY:
    		return "kBridgeProxy"
        case ServiceServiceStateResult_KLIBRARIAN:
    		return "kLibrarian"
        case ServiceServiceStateResult_KGROOT:
    		return "kGroot"
        case ServiceServiceStateResult_KEAGLEAGENT:
    		return "kEagleAgent"
        case ServiceServiceStateResult_KATHENA:
    		return "kAthena"
        case ServiceServiceStateResult_KBIFROSTBROKER:
    		return "kBifrostBroker"
        case ServiceServiceStateResult_KOS:
    		return "kOs"
        default:
        	return "kApollo"
    }
}

/**
 * Converts ServiceServiceStateResultEnum Array to its string Array representation
*/
func ServiceServiceStateResultEnumArrayToValue(serviceServiceStateResultEnum []ServiceServiceStateResultEnum) []string {
    convArray := make([]string,len( serviceServiceStateResultEnum))
    for i:=0; i<len(serviceServiceStateResultEnum);i++ {
        convArray[i] = ServiceServiceStateResultEnumToValue(serviceServiceStateResultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ServiceServiceStateResultEnumFromValue(value string) ServiceServiceStateResultEnum {
    switch value {
        case "kApollo":
            return ServiceServiceStateResult_KAPOLLO
        case "kBridge":
            return ServiceServiceStateResult_KBRIDGE
        case "kGenie":
            return ServiceServiceStateResult_KGENIE
        case "kGenieGofer":
            return ServiceServiceStateResult_KGENIEGOFER
        case "kMagneto":
            return ServiceServiceStateResult_KMAGNETO
        case "kIris":
            return ServiceServiceStateResult_KIRIS
        case "kIrisProxy":
            return ServiceServiceStateResult_KIRISPROXY
        case "kScribe":
            return ServiceServiceStateResult_KSCRIBE
        case "kStats":
            return ServiceServiceStateResult_KSTATS
        case "kYoda":
            return ServiceServiceStateResult_KYODA
        case "kAlerts":
            return ServiceServiceStateResult_KALERTS
        case "kKeychain":
            return ServiceServiceStateResult_KKEYCHAIN
        case "kLogWatcher":
            return ServiceServiceStateResult_KLOGWATCHER
        case "kStatsCollecter":
            return ServiceServiceStateResult_KSTATSCOLLECTER
        case "kGandalf":
            return ServiceServiceStateResult_KGANDALF
        case "kNexus":
            return ServiceServiceStateResult_KNEXUS
        case "kNexusProxy":
            return ServiceServiceStateResult_KNEXUSPROXY
        case "kStorageProxy":
            return ServiceServiceStateResult_KSTORAGEPROXY
        case "kTricorder":
            return ServiceServiceStateResult_KTRICORDER
        case "kRtClient":
            return ServiceServiceStateResult_KRTCLIENT
        case "kVaultProxy":
            return ServiceServiceStateResult_KVAULTPROXY
        case "kSmbProxy":
            return ServiceServiceStateResult_KSMBPROXY
        case "kBridgeProxy":
            return ServiceServiceStateResult_KBRIDGEPROXY
        case "kLibrarian":
            return ServiceServiceStateResult_KLIBRARIAN
        case "kGroot":
            return ServiceServiceStateResult_KGROOT
        case "kEagleAgent":
            return ServiceServiceStateResult_KEAGLEAGENT
        case "kAthena":
            return ServiceServiceStateResult_KATHENA
        case "kBifrostBroker":
            return ServiceServiceStateResult_KBIFROSTBROKER
        case "kOs":
            return ServiceServiceStateResult_KOS
        default:
            return ServiceServiceStateResult_KAPOLLO
    }
}
