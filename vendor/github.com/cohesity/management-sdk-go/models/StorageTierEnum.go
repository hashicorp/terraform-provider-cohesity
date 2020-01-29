// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StorageTierEnum enum
 */
type StorageTierEnum int

/**
 * Value collection for StorageTierEnum enum
 */
const (
    StorageTier_KPCIESSD StorageTierEnum = 1 + iota
    StorageTier_KSATASSD
    StorageTier_KSATAHDD
    StorageTier_KCLOUD
)

func (r StorageTierEnum) MarshalJSON() ([]byte, error) {
    s := StorageTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *StorageTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StorageTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StorageTierEnum to its string representation
 */
func StorageTierEnumToValue(storageTierEnum StorageTierEnum) string {
    switch storageTierEnum {
        case StorageTier_KPCIESSD:
    		return "kPCIeSSD"
        case StorageTier_KSATASSD:
    		return "kSATASSD"
        case StorageTier_KSATAHDD:
    		return "kSATAHDD"
        case StorageTier_KCLOUD:
    		return "kCLOUD"
        default:
        	return "kPCIeSSD"
    }
}

/**
 * Converts StorageTierEnum Array to its string Array representation
*/
func StorageTierEnumArrayToValue(storageTierEnum []StorageTierEnum) []string {
    convArray := make([]string,len( storageTierEnum))
    for i:=0; i<len(storageTierEnum);i++ {
        convArray[i] = StorageTierEnumToValue(storageTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StorageTierEnumFromValue(value string) StorageTierEnum {
    switch value {
        case "kPCIeSSD":
            return StorageTier_KPCIESSD
        case "kSATASSD":
            return StorageTier_KSATASSD
        case "kSATAHDD":
            return StorageTier_KSATAHDD
        case "kCLOUD":
            return StorageTier_KCLOUD
        default:
            return StorageTier_KPCIESSD
    }
}
