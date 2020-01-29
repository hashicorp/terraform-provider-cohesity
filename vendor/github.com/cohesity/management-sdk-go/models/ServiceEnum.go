// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ServiceEnum enum
 */
type ServiceEnum int

/**
 * Value collection for ServiceEnum enum
 */
const (
    Service_KAPOLLO ServiceEnum = 1 + iota
    Service_KBRIDGE
    Service_KGENIE
    Service_KGENIEGOFER
    Service_KMAGNETO
    Service_KIRIS
    Service_KIRISPROXY
    Service_KSCRIBE
    Service_KSTATS
    Service_KYODA
    Service_KALERTS
    Service_KKEYCHAIN
    Service_KLOGWATCHER
    Service_KSTATSCOLLECTER
    Service_KGANDALF
    Service_KNEXUS
    Service_KNEXUSPROXY
    Service_KSTORAGEPROXY
    Service_KTRICORDER
    Service_KRTCLIENT
    Service_KVAULTPROXY
    Service_KSMBPROXY
    Service_KBRIDGEPROXY
    Service_KLIBRARIAN
    Service_KGROOT
    Service_KEAGLEAGENT
    Service_KATHENA
    Service_KBIFROSTBROKER
    Service_KOS
)

func (r ServiceEnum) MarshalJSON() ([]byte, error) {
    s := ServiceEnumToValue(r)
    return json.Marshal(s)
}

func (r *ServiceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ServiceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ServiceEnum to its string representation
 */
func ServiceEnumToValue(serviceEnum ServiceEnum) string {
    switch serviceEnum {
        case Service_KAPOLLO:
    		return "kApollo"
        case Service_KBRIDGE:
    		return "kBridge"
        case Service_KGENIE:
    		return "kGenie"
        case Service_KGENIEGOFER:
    		return "kGenieGofer"
        case Service_KMAGNETO:
    		return "kMagneto"
        case Service_KIRIS:
    		return "kIris"
        case Service_KIRISPROXY:
    		return "kIrisProxy"
        case Service_KSCRIBE:
    		return "kScribe"
        case Service_KSTATS:
    		return "kStats"
        case Service_KYODA:
    		return "kYoda"
        case Service_KALERTS:
    		return "kAlerts"
        case Service_KKEYCHAIN:
    		return "kKeychain"
        case Service_KLOGWATCHER:
    		return "kLogWatcher"
        case Service_KSTATSCOLLECTER:
    		return "kStatsCollecter"
        case Service_KGANDALF:
    		return "kGandalf"
        case Service_KNEXUS:
    		return "kNexus"
        case Service_KNEXUSPROXY:
    		return "kNexusProxy"
        case Service_KSTORAGEPROXY:
    		return "kStorageProxy"
        case Service_KTRICORDER:
    		return "kTricorder"
        case Service_KRTCLIENT:
    		return "kRtClient"
        case Service_KVAULTPROXY:
    		return "kVaultProxy"
        case Service_KSMBPROXY:
    		return "kSmbProxy"
        case Service_KBRIDGEPROXY:
    		return "kBridgeProxy"
        case Service_KLIBRARIAN:
    		return "kLibrarian"
        case Service_KGROOT:
    		return "kGroot"
        case Service_KEAGLEAGENT:
    		return "kEagleAgent"
        case Service_KATHENA:
    		return "kAthena"
        case Service_KBIFROSTBROKER:
    		return "kBifrostBroker"
        case Service_KOS:
    		return "kOs"
        default:
        	return "kApollo"
    }
}

/**
 * Converts ServiceEnum Array to its string Array representation
*/
func ServiceEnumArrayToValue(serviceEnum []ServiceEnum) []string {
    convArray := make([]string,len( serviceEnum))
    for i:=0; i<len(serviceEnum);i++ {
        convArray[i] = ServiceEnumToValue(serviceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ServiceEnumFromValue(value string) ServiceEnum {
    switch value {
        case "kApollo":
            return Service_KAPOLLO
        case "kBridge":
            return Service_KBRIDGE
        case "kGenie":
            return Service_KGENIE
        case "kGenieGofer":
            return Service_KGENIEGOFER
        case "kMagneto":
            return Service_KMAGNETO
        case "kIris":
            return Service_KIRIS
        case "kIrisProxy":
            return Service_KIRISPROXY
        case "kScribe":
            return Service_KSCRIBE
        case "kStats":
            return Service_KSTATS
        case "kYoda":
            return Service_KYODA
        case "kAlerts":
            return Service_KALERTS
        case "kKeychain":
            return Service_KKEYCHAIN
        case "kLogWatcher":
            return Service_KLOGWATCHER
        case "kStatsCollecter":
            return Service_KSTATSCOLLECTER
        case "kGandalf":
            return Service_KGANDALF
        case "kNexus":
            return Service_KNEXUS
        case "kNexusProxy":
            return Service_KNEXUSPROXY
        case "kStorageProxy":
            return Service_KSTORAGEPROXY
        case "kTricorder":
            return Service_KTRICORDER
        case "kRtClient":
            return Service_KRTCLIENT
        case "kVaultProxy":
            return Service_KVAULTPROXY
        case "kSmbProxy":
            return Service_KSMBPROXY
        case "kBridgeProxy":
            return Service_KBRIDGEPROXY
        case "kLibrarian":
            return Service_KLIBRARIAN
        case "kGroot":
            return Service_KGROOT
        case "kEagleAgent":
            return Service_KEAGLEAGENT
        case "kAthena":
            return Service_KATHENA
        case "kBifrostBroker":
            return Service_KBIFROSTBROKER
        case "kOs":
            return Service_KOS
        default:
            return Service_KAPOLLO
    }
}
