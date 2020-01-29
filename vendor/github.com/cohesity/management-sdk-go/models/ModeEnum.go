// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ModeEnum enum
 */
type ModeEnum int

/**
 * Value collection for ModeEnum enum
 */
const (
    Mode_KFOLDERSUBFOLDERSANDFILES ModeEnum = 1 + iota
    Mode_KFOLDERANDSUBFOLDERS
    Mode_KFOLDERANDFILES
    Mode_KFOLDERONLY
    Mode_KSUBFOLDERSANDFILESONLY
    Mode_KSUBFOLDERSONLY
    Mode_KFILESONLY
)

func (r ModeEnum) MarshalJSON() ([]byte, error) {
    s := ModeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ModeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ModeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ModeEnum to its string representation
 */
func ModeEnumToValue(modeEnum ModeEnum) string {
    switch modeEnum {
        case Mode_KFOLDERSUBFOLDERSANDFILES:
    		return "kFolderSubFoldersAndFiles"
        case Mode_KFOLDERANDSUBFOLDERS:
    		return "kFolderAndSubFolders"
        case Mode_KFOLDERANDFILES:
    		return "kFolderAndFiles"
        case Mode_KFOLDERONLY:
    		return "kFolderOnly"
        case Mode_KSUBFOLDERSANDFILESONLY:
    		return "kSubFoldersAndFilesOnly"
        case Mode_KSUBFOLDERSONLY:
    		return "kSubFoldersOnly"
        case Mode_KFILESONLY:
    		return "kFilesOnly"
        default:
        	return "kFolderSubFoldersAndFiles"
    }
}

/**
 * Converts ModeEnum Array to its string Array representation
*/
func ModeEnumArrayToValue(modeEnum []ModeEnum) []string {
    convArray := make([]string,len( modeEnum))
    for i:=0; i<len(modeEnum);i++ {
        convArray[i] = ModeEnumToValue(modeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ModeEnumFromValue(value string) ModeEnum {
    switch value {
        case "kFolderSubFoldersAndFiles":
            return Mode_KFOLDERSUBFOLDERSANDFILES
        case "kFolderAndSubFolders":
            return Mode_KFOLDERANDSUBFOLDERS
        case "kFolderAndFiles":
            return Mode_KFOLDERANDFILES
        case "kFolderOnly":
            return Mode_KFOLDERONLY
        case "kSubFoldersAndFilesOnly":
            return Mode_KSUBFOLDERSANDFILESONLY
        case "kSubFoldersOnly":
            return Mode_KSUBFOLDERSONLY
        case "kFilesOnly":
            return Mode_KFILESONLY
        default:
            return Mode_KFOLDERSUBFOLDERSANDFILES
    }
}
