package models

import(
    "encoding/json"
)

/**
 * Type definition for State2Enum enum
 */
type State2Enum int

/**
 * Value collection for State2Enum enum
 */
const (
    State2_KINITIALIZING State2Enum = 1 + iota
    State2_KRUNNING
    State2_KPAUSING
    State2_KPAUSED
    State2_KTERMINATING
    State2_KTERMINATED
    State2_KFAILED
)

func (r State2Enum) MarshalJSON() ([]byte, error) { 
    s := State2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *State2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  State2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts State2Enum to its string representation
 */
func State2EnumToValue(state2Enum State2Enum) string {
    switch state2Enum {
        case State2_KINITIALIZING:
    		return "kInitializing"		
        case State2_KRUNNING:
    		return "kRunning"		
        case State2_KPAUSING:
    		return "kPausing"		
        case State2_KPAUSED:
    		return "kPaused"		
        case State2_KTERMINATING:
    		return "kTerminating"		
        case State2_KTERMINATED:
    		return "kTerminated"		
        case State2_KFAILED:
    		return "kFailed"		
        default:
        	return "kInitializing"
    }
}

/**
 * Converts State2Enum Array to its string Array representation
*/
func State2EnumArrayToValue(state2Enum []State2Enum) []string {
    convArray := make([]string,len( state2Enum))
    for i:=0; i<len(state2Enum);i++ {
        convArray[i] = State2EnumToValue(state2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func State2EnumFromValue(value string) State2Enum {
    switch value {
        case "kInitializing":
            return State2_KINITIALIZING
        case "kRunning":
            return State2_KRUNNING
        case "kPausing":
            return State2_KPAUSING
        case "kPaused":
            return State2_KPAUSED
        case "kTerminating":
            return State2_KTERMINATING
        case "kTerminated":
            return State2_KTERMINATED
        case "kFailed":
            return State2_KFAILED
        default:
            return State2_KINITIALIZING
    }
}
