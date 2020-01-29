// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CentrifySchemaEnum enum
 */
type CentrifySchemaEnum int

/**
 * Value collection for CentrifySchemaEnum enum
 */
const (
    CentrifySchema_KCENTRIFYDYNAMICSCHEMA_1_0 CentrifySchemaEnum = 1 + iota
    CentrifySchema_KCENTRIFYDYNAMICSCHEMA_2_0
    CentrifySchema_KCENTRIFYSFU_3_0
    CentrifySchema_KCENTRIFYSFU_4_0
    CentrifySchema_KCENTRIFYCDCRFC2307
    CentrifySchema_KCENTRIFYDYNAMICSCHEMA_3_0
    CentrifySchema_KCENTRIFYCDCRFC2307_2
    CentrifySchema_KCENTRIFYDYNAMICSCHEMA_5_0
    CentrifySchema_KCENTRIFYCDCRFC2307_3
    CentrifySchema_KCENTRIFYSFU_3_0_V5
)

func (r CentrifySchemaEnum) MarshalJSON() ([]byte, error) {
    s := CentrifySchemaEnumToValue(r)
    return json.Marshal(s)
}

func (r *CentrifySchemaEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CentrifySchemaEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CentrifySchemaEnum to its string representation
 */
func CentrifySchemaEnumToValue(centrifySchemaEnum CentrifySchemaEnum) string {
    switch centrifySchemaEnum {
        case CentrifySchema_KCENTRIFYDYNAMICSCHEMA_1_0:
    		return "kCentrifyDynamicSchema_1_0"
        case CentrifySchema_KCENTRIFYDYNAMICSCHEMA_2_0:
    		return "kCentrifyDynamicSchema_2_0"
        case CentrifySchema_KCENTRIFYSFU_3_0:
    		return "kCentrifySfu_3_0"
        case CentrifySchema_KCENTRIFYSFU_4_0:
    		return "kCentrifySfu_4_0"
        case CentrifySchema_KCENTRIFYCDCRFC2307:
    		return "kCentrifyCdcRfc2307"
        case CentrifySchema_KCENTRIFYDYNAMICSCHEMA_3_0:
    		return "kCentrifyDynamicSchema_3_0"
        case CentrifySchema_KCENTRIFYCDCRFC2307_2:
    		return "kCentrifyCdcRfc2307_2"
        case CentrifySchema_KCENTRIFYDYNAMICSCHEMA_5_0:
    		return "kCentrifyDynamicSchema_5_0"
        case CentrifySchema_KCENTRIFYCDCRFC2307_3:
    		return "kCentrifyCdcRfc2307_3"
        case CentrifySchema_KCENTRIFYSFU_3_0_V5:
    		return "kCentrifySfu_3_0_V5"
        default:
        	return "kCentrifyDynamicSchema_1_0"
    }
}

/**
 * Converts CentrifySchemaEnum Array to its string Array representation
*/
func CentrifySchemaEnumArrayToValue(centrifySchemaEnum []CentrifySchemaEnum) []string {
    convArray := make([]string,len( centrifySchemaEnum))
    for i:=0; i<len(centrifySchemaEnum);i++ {
        convArray[i] = CentrifySchemaEnumToValue(centrifySchemaEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CentrifySchemaEnumFromValue(value string) CentrifySchemaEnum {
    switch value {
        case "kCentrifyDynamicSchema_1_0":
            return CentrifySchema_KCENTRIFYDYNAMICSCHEMA_1_0
        case "kCentrifyDynamicSchema_2_0":
            return CentrifySchema_KCENTRIFYDYNAMICSCHEMA_2_0
        case "kCentrifySfu_3_0":
            return CentrifySchema_KCENTRIFYSFU_3_0
        case "kCentrifySfu_4_0":
            return CentrifySchema_KCENTRIFYSFU_4_0
        case "kCentrifyCdcRfc2307":
            return CentrifySchema_KCENTRIFYCDCRFC2307
        case "kCentrifyDynamicSchema_3_0":
            return CentrifySchema_KCENTRIFYDYNAMICSCHEMA_3_0
        case "kCentrifyCdcRfc2307_2":
            return CentrifySchema_KCENTRIFYCDCRFC2307_2
        case "kCentrifyDynamicSchema_5_0":
            return CentrifySchema_KCENTRIFYDYNAMICSCHEMA_5_0
        case "kCentrifyCdcRfc2307_3":
            return CentrifySchema_KCENTRIFYCDCRFC2307_3
        case "kCentrifySfu_3_0_V5":
            return CentrifySchema_KCENTRIFYSFU_3_0_V5
        default:
            return CentrifySchema_KCENTRIFYDYNAMICSCHEMA_1_0
    }
}
