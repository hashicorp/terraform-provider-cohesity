package models

import(
    "encoding/json"
)

/**
 * Type definition for State1Enum enum
 */
type State1Enum int

/**
 * Value collection for State1Enum enum
 */
const (
    State1_KINITIALIZING State1Enum = 1 + iota
    State1_KRUNNING
    State1_KPAUSING
    State1_KPAUSED
    State1_KTERMINATING
    State1_KTERMINATED
    State1_KFAILED
)

func (r State1Enum) MarshalJSON() ([]byte, error) { 
    s := State1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *State1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  State1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts State1Enum to its string representation
 */
func State1EnumToValue(state1Enum State1Enum) string {
    switch state1Enum {
        case State1_KINITIALIZING:
    		return "kInitializing"		
        case State1_KRUNNING:
    		return "kRunning"		
        case State1_KPAUSING:
    		return "kPausing"		
        case State1_KPAUSED:
    		return "kPaused"		
        case State1_KTERMINATING:
    		return "kTerminating"		
        case State1_KTERMINATED:
    		return "kTerminated"		
        case State1_KFAILED:
    		return "kFailed"		
        default:
        	return "kInitializing"
    }
}

/**
 * Converts State1Enum Array to its string Array representation
*/
func State1EnumArrayToValue(state1Enum []State1Enum) []string {
    convArray := make([]string,len( state1Enum))
    for i:=0; i<len(state1Enum);i++ {
        convArray[i] = State1EnumToValue(state1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func State1EnumFromValue(value string) State1Enum {
    switch value {
        case "kInitializing":
            return State1_KINITIALIZING
        case "kRunning":
            return State1_KRUNNING
        case "kPausing":
            return State1_KPAUSING
        case "kPaused":
            return State1_KPAUSED
        case "kTerminating":
            return State1_KTERMINATING
        case "kTerminated":
            return State1_KTERMINATED
        case "kFailed":
            return State1_KFAILED
        default:
            return State1_KINITIALIZING
    }
}
