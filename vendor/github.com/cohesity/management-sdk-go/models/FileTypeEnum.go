// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for FileTypeEnum enum
 */
type FileTypeEnum int

/**
 * Value collection for FileTypeEnum enum
 */
const (
    FileType_KROWS FileTypeEnum = 1 + iota
    FileType_KLOG
    FileType_KFILESTREAM
    FileType_KNOTSUPPORTEDTYPE
    FileType_KFULLTEXT
)

func (r FileTypeEnum) MarshalJSON() ([]byte, error) {
    s := FileTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *FileTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  FileTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts FileTypeEnum to its string representation
 */
func FileTypeEnumToValue(fileTypeEnum FileTypeEnum) string {
    switch fileTypeEnum {
        case FileType_KROWS:
    		return "kRows"
        case FileType_KLOG:
    		return "kLog"
        case FileType_KFILESTREAM:
    		return "kFileStream"
        case FileType_KNOTSUPPORTEDTYPE:
    		return "kNotSupportedType"
        case FileType_KFULLTEXT:
    		return "kFullText"
        default:
        	return "kRows"
    }
}

/**
 * Converts FileTypeEnum Array to its string Array representation
*/
func FileTypeEnumArrayToValue(fileTypeEnum []FileTypeEnum) []string {
    convArray := make([]string,len( fileTypeEnum))
    for i:=0; i<len(fileTypeEnum);i++ {
        convArray[i] = FileTypeEnumToValue(fileTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func FileTypeEnumFromValue(value string) FileTypeEnum {
    switch value {
        case "kRows":
            return FileType_KROWS
        case "kLog":
            return FileType_KLOG
        case "kFileStream":
            return FileType_KFILESTREAM
        case "kNotSupportedType":
            return FileType_KNOTSUPPORTEDTYPE
        case "kFullText":
            return FileType_KFULLTEXT
        default:
            return FileType_KROWS
    }
}
