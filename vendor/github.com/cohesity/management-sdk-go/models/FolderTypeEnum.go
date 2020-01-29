// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for FolderTypeEnum enum
 */
type FolderTypeEnum int

/**
 * Value collection for FolderTypeEnum enum
 */
const (
    FolderType_KVMFOLDER FolderTypeEnum = 1 + iota
    FolderType_KHOSTFOLDER
    FolderType_KDATASTOREFOLDER
    FolderType_KNETWORKFOLDER
    FolderType_KROOTFOLDER
)

func (r FolderTypeEnum) MarshalJSON() ([]byte, error) {
    s := FolderTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *FolderTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  FolderTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts FolderTypeEnum to its string representation
 */
func FolderTypeEnumToValue(folderTypeEnum FolderTypeEnum) string {
    switch folderTypeEnum {
        case FolderType_KVMFOLDER:
    		return "kVMFolder"
        case FolderType_KHOSTFOLDER:
    		return "kHostFolder"
        case FolderType_KDATASTOREFOLDER:
    		return "kDatastoreFolder"
        case FolderType_KNETWORKFOLDER:
    		return "kNetworkFolder"
        case FolderType_KROOTFOLDER:
    		return "kRootFolder"
        default:
        	return "kVMFolder"
    }
}

/**
 * Converts FolderTypeEnum Array to its string Array representation
*/
func FolderTypeEnumArrayToValue(folderTypeEnum []FolderTypeEnum) []string {
    convArray := make([]string,len( folderTypeEnum))
    for i:=0; i<len(folderTypeEnum);i++ {
        convArray[i] = FolderTypeEnumToValue(folderTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func FolderTypeEnumFromValue(value string) FolderTypeEnum {
    switch value {
        case "kVMFolder":
            return FolderType_KVMFOLDER
        case "kHostFolder":
            return FolderType_KHOSTFOLDER
        case "kDatastoreFolder":
            return FolderType_KDATASTOREFOLDER
        case "kNetworkFolder":
            return FolderType_KNETWORKFOLDER
        case "kRootFolder":
            return FolderType_KROOTFOLDER
        default:
            return FolderType_KVMFOLDER
    }
}
