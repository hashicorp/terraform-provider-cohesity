// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for LibrarianIOPreferentialTierEnum enum
 */
type LibrarianIOPreferentialTierEnum int

/**
 * Value collection for LibrarianIOPreferentialTierEnum enum
 */
const (
    LibrarianIOPreferentialTier_KPCIESSD LibrarianIOPreferentialTierEnum = 1 + iota
    LibrarianIOPreferentialTier_KSATASSD
    LibrarianIOPreferentialTier_KSATAHDD
    LibrarianIOPreferentialTier_KCLOUD
)

func (r LibrarianIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := LibrarianIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *LibrarianIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  LibrarianIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts LibrarianIOPreferentialTierEnum to its string representation
 */
func LibrarianIOPreferentialTierEnumToValue(librarianIOPreferentialTierEnum LibrarianIOPreferentialTierEnum) string {
    switch librarianIOPreferentialTierEnum {
        case LibrarianIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case LibrarianIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case LibrarianIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case LibrarianIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts LibrarianIOPreferentialTierEnum Array to its string Array representation
*/
func LibrarianIOPreferentialTierEnumArrayToValue(librarianIOPreferentialTierEnum []LibrarianIOPreferentialTierEnum) []string {
    convArray := make([]string,len( librarianIOPreferentialTierEnum))
    for i:=0; i<len(librarianIOPreferentialTierEnum);i++ {
        convArray[i] = LibrarianIOPreferentialTierEnumToValue(librarianIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func LibrarianIOPreferentialTierEnumFromValue(value string) LibrarianIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return LibrarianIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return LibrarianIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return LibrarianIOPreferentialTier_KSATAHDD
        case "kCloud":
            return LibrarianIOPreferentialTier_KCLOUD
        default:
            return LibrarianIOPreferentialTier_KPCIESSD
    }
}
