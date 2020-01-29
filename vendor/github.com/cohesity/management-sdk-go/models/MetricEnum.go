// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for MetricEnum enum
 */
type MetricEnum int

/**
 * Value collection for MetricEnum enum
 */
const (
    Metric_KSYSTEMUSAGEBYTES MetricEnum = 1 + iota
    Metric_KCHUNKBYTESPHYSICAL
    Metric_KNUMBYTESWRITTEN
    Metric_KNUMBYTESREAD
    Metric_KPEAKREADTHROUGHPUT
    Metric_KPEAKWRITETHROUGHPUT
)

func (r MetricEnum) MarshalJSON() ([]byte, error) {
    s := MetricEnumToValue(r)
    return json.Marshal(s)
}

func (r *MetricEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  MetricEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts MetricEnum to its string representation
 */
func MetricEnumToValue(metricEnum MetricEnum) string {
    switch metricEnum {
        case Metric_KSYSTEMUSAGEBYTES:
    		return "kSystemUsageBytes"
        case Metric_KCHUNKBYTESPHYSICAL:
    		return "kChunkBytesPhysical"
        case Metric_KNUMBYTESWRITTEN:
    		return "kNumBytesWritten"
        case Metric_KNUMBYTESREAD:
    		return "kNumBytesRead"
        case Metric_KPEAKREADTHROUGHPUT:
    		return "kPeakReadThroughput"
        case Metric_KPEAKWRITETHROUGHPUT:
    		return "kPeakWriteThroughput"
        default:
        	return "kSystemUsageBytes"
    }
}

/**
 * Converts MetricEnum Array to its string Array representation
*/
func MetricEnumArrayToValue(metricEnum []MetricEnum) []string {
    convArray := make([]string,len( metricEnum))
    for i:=0; i<len(metricEnum);i++ {
        convArray[i] = MetricEnumToValue(metricEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func MetricEnumFromValue(value string) MetricEnum {
    switch value {
        case "kSystemUsageBytes":
            return Metric_KSYSTEMUSAGEBYTES
        case "kChunkBytesPhysical":
            return Metric_KCHUNKBYTESPHYSICAL
        case "kNumBytesWritten":
            return Metric_KNUMBYTESWRITTEN
        case "kNumBytesRead":
            return Metric_KNUMBYTESREAD
        case "kPeakReadThroughput":
            return Metric_KPEAKREADTHROUGHPUT
        case "kPeakWriteThroughput":
            return Metric_KPEAKWRITETHROUGHPUT
        default:
            return Metric_KSYSTEMUSAGEBYTES
    }
}
