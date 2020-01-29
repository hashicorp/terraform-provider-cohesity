// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ServiceNetworkInterfaceEnum enum
 */
type ServiceNetworkInterfaceEnum int

/**
 * Value collection for ServiceNetworkInterfaceEnum enum
 */
const (
    ServiceNetworkInterface_KREPLICATIONSERVICE ServiceNetworkInterfaceEnum = 1 + iota
    ServiceNetworkInterface_KREMOTETUNNELSERVICE
    ServiceNetworkInterface_KCLUSTERDATASERVICE
    ServiceNetworkInterface_KAVAHIDISCOVERSERVICE
)

func (r ServiceNetworkInterfaceEnum) MarshalJSON() ([]byte, error) {
    s := ServiceNetworkInterfaceEnumToValue(r)
    return json.Marshal(s)
}

func (r *ServiceNetworkInterfaceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ServiceNetworkInterfaceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ServiceNetworkInterfaceEnum to its string representation
 */
func ServiceNetworkInterfaceEnumToValue(serviceNetworkInterfaceEnum ServiceNetworkInterfaceEnum) string {
    switch serviceNetworkInterfaceEnum {
        case ServiceNetworkInterface_KREPLICATIONSERVICE:
    		return "kReplicationService"
        case ServiceNetworkInterface_KREMOTETUNNELSERVICE:
    		return "kRemoteTunnelService"
        case ServiceNetworkInterface_KCLUSTERDATASERVICE:
    		return "kClusterDataService"
        case ServiceNetworkInterface_KAVAHIDISCOVERSERVICE:
    		return "kAvahiDiscoverService"
        default:
        	return "kReplicationService"
    }
}

/**
 * Converts ServiceNetworkInterfaceEnum Array to its string Array representation
*/
func ServiceNetworkInterfaceEnumArrayToValue(serviceNetworkInterfaceEnum []ServiceNetworkInterfaceEnum) []string {
    convArray := make([]string,len( serviceNetworkInterfaceEnum))
    for i:=0; i<len(serviceNetworkInterfaceEnum);i++ {
        convArray[i] = ServiceNetworkInterfaceEnumToValue(serviceNetworkInterfaceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ServiceNetworkInterfaceEnumFromValue(value string) ServiceNetworkInterfaceEnum {
    switch value {
        case "kReplicationService":
            return ServiceNetworkInterface_KREPLICATIONSERVICE
        case "kRemoteTunnelService":
            return ServiceNetworkInterface_KREMOTETUNNELSERVICE
        case "kClusterDataService":
            return ServiceNetworkInterface_KCLUSTERDATASERVICE
        case "kAvahiDiscoverService":
            return ServiceNetworkInterface_KAVAHIDISCOVERSERVICE
        default:
            return ServiceNetworkInterface_KREPLICATIONSERVICE
    }
}
