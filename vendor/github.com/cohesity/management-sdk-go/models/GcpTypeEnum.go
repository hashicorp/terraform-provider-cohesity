// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for GcpTypeEnum enum
 */
type GcpTypeEnum int

/**
 * Value collection for GcpTypeEnum enum
 */
const (
    GcpType_KIAMUSER GcpTypeEnum = 1 + iota
    GcpType_KPROJECT
    GcpType_KREGION
    GcpType_KAVAILABILITYZONE
    GcpType_KVIRTUALMACHINE
    GcpType_KVPC
    GcpType_KSUBNET
    GcpType_KNETWORKSECURITYGROUP
    GcpType_KINSTANCETYPE
    GcpType_KLABEL
    GcpType_KMETADATA
    GcpType_KTAG
    GcpType_KVPCCONNECTOR
)

func (r GcpTypeEnum) MarshalJSON() ([]byte, error) {
    s := GcpTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *GcpTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  GcpTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts GcpTypeEnum to its string representation
 */
func GcpTypeEnumToValue(gcpTypeEnum GcpTypeEnum) string {
    switch gcpTypeEnum {
        case GcpType_KIAMUSER:
    		return "kIAMUser"
        case GcpType_KPROJECT:
    		return "kProject"
        case GcpType_KREGION:
    		return "kRegion"
        case GcpType_KAVAILABILITYZONE:
    		return "kAvailabilityZone"
        case GcpType_KVIRTUALMACHINE:
    		return "kVirtualMachine"
        case GcpType_KVPC:
    		return "kVPC"
        case GcpType_KSUBNET:
    		return "kSubnet"
        case GcpType_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"
        case GcpType_KINSTANCETYPE:
    		return "kInstanceType"
        case GcpType_KLABEL:
    		return "kLabel"
        case GcpType_KMETADATA:
    		return "kMetadata"
        case GcpType_KTAG:
    		return "kTag"
        case GcpType_KVPCCONNECTOR:
    		return "kVPCConnector"
        default:
        	return "kIAMUser"
    }
}

/**
 * Converts GcpTypeEnum Array to its string Array representation
*/
func GcpTypeEnumArrayToValue(gcpTypeEnum []GcpTypeEnum) []string {
    convArray := make([]string,len( gcpTypeEnum))
    for i:=0; i<len(gcpTypeEnum);i++ {
        convArray[i] = GcpTypeEnumToValue(gcpTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func GcpTypeEnumFromValue(value string) GcpTypeEnum {
    switch value {
        case "kIAMUser":
            return GcpType_KIAMUSER
        case "kProject":
            return GcpType_KPROJECT
        case "kRegion":
            return GcpType_KREGION
        case "kAvailabilityZone":
            return GcpType_KAVAILABILITYZONE
        case "kVirtualMachine":
            return GcpType_KVIRTUALMACHINE
        case "kVPC":
            return GcpType_KVPC
        case "kSubnet":
            return GcpType_KSUBNET
        case "kNetworkSecurityGroup":
            return GcpType_KNETWORKSECURITYGROUP
        case "kInstanceType":
            return GcpType_KINSTANCETYPE
        case "kLabel":
            return GcpType_KLABEL
        case "kMetadata":
            return GcpType_KMETADATA
        case "kTag":
            return GcpType_KTAG
        case "kVPCConnector":
            return GcpType_KVPCCONNECTOR
        default:
            return GcpType_KIAMUSER
    }
}
