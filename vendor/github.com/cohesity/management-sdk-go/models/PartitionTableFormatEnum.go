// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PartitionTableFormatEnum enum
 */
type PartitionTableFormatEnum int

/**
 * Value collection for PartitionTableFormatEnum enum
 */
const (
    PartitionTableFormat_KNOPARTITION PartitionTableFormatEnum = 1 + iota
    PartitionTableFormat_KMBRPARTITION
    PartitionTableFormat_KGPTPARTITION
    PartitionTableFormat_KSGIPARTITION
    PartitionTableFormat_KSUNPARTITION
)

func (r PartitionTableFormatEnum) MarshalJSON() ([]byte, error) {
    s := PartitionTableFormatEnumToValue(r)
    return json.Marshal(s)
}

func (r *PartitionTableFormatEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PartitionTableFormatEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PartitionTableFormatEnum to its string representation
 */
func PartitionTableFormatEnumToValue(partitionTableFormatEnum PartitionTableFormatEnum) string {
    switch partitionTableFormatEnum {
        case PartitionTableFormat_KNOPARTITION:
    		return "kNoPartition"
        case PartitionTableFormat_KMBRPARTITION:
    		return "kMBRPartition"
        case PartitionTableFormat_KGPTPARTITION:
    		return "kGPTPartition"
        case PartitionTableFormat_KSGIPARTITION:
    		return "kSGIPartition"
        case PartitionTableFormat_KSUNPARTITION:
    		return "kSUNPartition"
        default:
        	return "kNoPartition"
    }
}

/**
 * Converts PartitionTableFormatEnum Array to its string Array representation
*/
func PartitionTableFormatEnumArrayToValue(partitionTableFormatEnum []PartitionTableFormatEnum) []string {
    convArray := make([]string,len( partitionTableFormatEnum))
    for i:=0; i<len(partitionTableFormatEnum);i++ {
        convArray[i] = PartitionTableFormatEnumToValue(partitionTableFormatEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PartitionTableFormatEnumFromValue(value string) PartitionTableFormatEnum {
    switch value {
        case "kNoPartition":
            return PartitionTableFormat_KNOPARTITION
        case "kMBRPartition":
            return PartitionTableFormat_KMBRPARTITION
        case "kGPTPartition":
            return PartitionTableFormat_KGPTPARTITION
        case "kSGIPartition":
            return PartitionTableFormat_KSGIPARTITION
        case "kSUNPartition":
            return PartitionTableFormat_KSUNPARTITION
        default:
            return PartitionTableFormat_KNOPARTITION
    }
}
