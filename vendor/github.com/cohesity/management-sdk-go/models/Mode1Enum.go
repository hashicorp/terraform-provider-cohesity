package models

import(
    "encoding/json"
)

/**
 * Type definition for Mode1Enum enum
 */
type Mode1Enum int

/**
 * Value collection for Mode1Enum enum
 */
const (
    Mode1_KFOLDERSUBFOLDERSANDFILES Mode1Enum = 1 + iota
    Mode1_KFOLDERANDSUBFOLDERS
    Mode1_KFOLDERANDFILES
    Mode1_KFOLDERONLY
    Mode1_KSUBFOLDERSANDFILESONLY
    Mode1_KSUBFOLDERSONLY
    Mode1_KFILESONLY
)

func (r Mode1Enum) MarshalJSON() ([]byte, error) { 
    s := Mode1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Mode1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Mode1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Mode1Enum to its string representation
 */
func Mode1EnumToValue(mode1Enum Mode1Enum) string {
    switch mode1Enum {
        case Mode1_KFOLDERSUBFOLDERSANDFILES:
    		return "kFolderSubFoldersAndFiles"		
        case Mode1_KFOLDERANDSUBFOLDERS:
    		return "kFolderAndSubFolders"		
        case Mode1_KFOLDERANDFILES:
    		return "kFolderAndFiles"		
        case Mode1_KFOLDERONLY:
    		return "kFolderOnly"		
        case Mode1_KSUBFOLDERSANDFILESONLY:
    		return "kSubFoldersAndFilesOnly"		
        case Mode1_KSUBFOLDERSONLY:
    		return "kSubFoldersOnly"		
        case Mode1_KFILESONLY:
    		return "kFilesOnly"		
        default:
        	return "kFolderSubFoldersAndFiles"
    }
}

/**
 * Converts Mode1Enum Array to its string Array representation
*/
func Mode1EnumArrayToValue(mode1Enum []Mode1Enum) []string {
    convArray := make([]string,len( mode1Enum))
    for i:=0; i<len(mode1Enum);i++ {
        convArray[i] = Mode1EnumToValue(mode1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Mode1EnumFromValue(value string) Mode1Enum {
    switch value {
        case "kFolderSubFoldersAndFiles":
            return Mode1_KFOLDERSUBFOLDERSANDFILES
        case "kFolderAndSubFolders":
            return Mode1_KFOLDERANDSUBFOLDERS
        case "kFolderAndFiles":
            return Mode1_KFOLDERANDFILES
        case "kFolderOnly":
            return Mode1_KFOLDERONLY
        case "kSubFoldersAndFilesOnly":
            return Mode1_KSUBFOLDERSANDFILESONLY
        case "kSubFoldersOnly":
            return Mode1_KSUBFOLDERSONLY
        case "kFilesOnly":
            return Mode1_KFILESONLY
        default:
            return Mode1_KFOLDERSUBFOLDERSANDFILES
    }
}
