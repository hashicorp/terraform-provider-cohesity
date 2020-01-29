// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AgentTypeEnum enum
 */
type AgentTypeEnum int

/**
 * Value collection for AgentTypeEnum enum
 */
const (
    AgentType_KCPP AgentTypeEnum = 1 + iota
    AgentType_KJAVA
    AgentType_KGO
)

func (r AgentTypeEnum) MarshalJSON() ([]byte, error) {
    s := AgentTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AgentTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AgentTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AgentTypeEnum to its string representation
 */
func AgentTypeEnumToValue(agentTypeEnum AgentTypeEnum) string {
    switch agentTypeEnum {
        case AgentType_KCPP:
    		return "kCpp"
        case AgentType_KJAVA:
    		return "kJava"
        case AgentType_KGO:
    		return "kGo"
        default:
        	return "kCpp"
    }
}

/**
 * Converts AgentTypeEnum Array to its string Array representation
*/
func AgentTypeEnumArrayToValue(agentTypeEnum []AgentTypeEnum) []string {
    convArray := make([]string,len( agentTypeEnum))
    for i:=0; i<len(agentTypeEnum);i++ {
        convArray[i] = AgentTypeEnumToValue(agentTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AgentTypeEnumFromValue(value string) AgentTypeEnum {
    switch value {
        case "kCpp":
            return AgentType_KCPP
        case "kJava":
            return AgentType_KJAVA
        case "kGo":
            return AgentType_KGO
        default:
            return AgentType_KCPP
    }
}
