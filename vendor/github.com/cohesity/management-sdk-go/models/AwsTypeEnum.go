// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AwsTypeEnum enum
 */
type AwsTypeEnum int

/**
 * Value collection for AwsTypeEnum enum
 */
const (
    AwsType_KIAMUSER AwsTypeEnum = 1 + iota
    AwsType_KREGION
    AwsType_KAVAILABILITYZONE
    AwsType_KEC2INSTANCE
    AwsType_KVPC
    AwsType_KSUBNET
    AwsType_KNETWORKSECURITYGROUP
    AwsType_KINSTANCETYPE
    AwsType_KKEYPAIR
    AwsType_KTAG
    AwsType_KRDSOPTIONGROUP
    AwsType_KRDSPARAMETERGROUP
    AwsType_KRDSINSTANCE
    AwsType_KRDSSUBNET
    AwsType_KRDSTAG
)

func (r AwsTypeEnum) MarshalJSON() ([]byte, error) {
    s := AwsTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AwsTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AwsTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AwsTypeEnum to its string representation
 */
func AwsTypeEnumToValue(awsTypeEnum AwsTypeEnum) string {
    switch awsTypeEnum {
        case AwsType_KIAMUSER:
    		return "kIAMUser"
        case AwsType_KREGION:
    		return "kRegion"
        case AwsType_KAVAILABILITYZONE:
    		return "kAvailabilityZone"
        case AwsType_KEC2INSTANCE:
    		return "kEC2Instance"
        case AwsType_KVPC:
    		return "kVPC"
        case AwsType_KSUBNET:
    		return "kSubnet"
        case AwsType_KNETWORKSECURITYGROUP:
    		return "kNetworkSecurityGroup"
        case AwsType_KINSTANCETYPE:
    		return "kInstanceType"
        case AwsType_KKEYPAIR:
    		return "kKeyPair"
        case AwsType_KTAG:
    		return "kTag"
        case AwsType_KRDSOPTIONGROUP:
    		return "kRDSOptionGroup"
        case AwsType_KRDSPARAMETERGROUP:
    		return "kRDSParameterGroup"
        case AwsType_KRDSINSTANCE:
    		return "kRDSInstance"
        case AwsType_KRDSSUBNET:
    		return "kRDSSubnet"
        case AwsType_KRDSTAG:
    		return "kRDSTag"
        default:
        	return "kIAMUser"
    }
}

/**
 * Converts AwsTypeEnum Array to its string Array representation
*/
func AwsTypeEnumArrayToValue(awsTypeEnum []AwsTypeEnum) []string {
    convArray := make([]string,len( awsTypeEnum))
    for i:=0; i<len(awsTypeEnum);i++ {
        convArray[i] = AwsTypeEnumToValue(awsTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AwsTypeEnumFromValue(value string) AwsTypeEnum {
    switch value {
        case "kIAMUser":
            return AwsType_KIAMUSER
        case "kRegion":
            return AwsType_KREGION
        case "kAvailabilityZone":
            return AwsType_KAVAILABILITYZONE
        case "kEC2Instance":
            return AwsType_KEC2INSTANCE
        case "kVPC":
            return AwsType_KVPC
        case "kSubnet":
            return AwsType_KSUBNET
        case "kNetworkSecurityGroup":
            return AwsType_KNETWORKSECURITYGROUP
        case "kInstanceType":
            return AwsType_KINSTANCETYPE
        case "kKeyPair":
            return AwsType_KKEYPAIR
        case "kTag":
            return AwsType_KTAG
        case "kRDSOptionGroup":
            return AwsType_KRDSOPTIONGROUP
        case "kRDSParameterGroup":
            return AwsType_KRDSPARAMETERGROUP
        case "kRDSInstance":
            return AwsType_KRDSINSTANCE
        case "kRDSSubnet":
            return AwsType_KRDSSUBNET
        case "kRDSTag":
            return AwsType_KRDSTAG
        default:
            return AwsType_KIAMUSER
    }
}
