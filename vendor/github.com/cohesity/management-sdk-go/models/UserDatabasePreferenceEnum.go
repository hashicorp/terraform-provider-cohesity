// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for UserDatabasePreferenceEnum enum
 */
type UserDatabasePreferenceEnum int

/**
 * Value collection for UserDatabasePreferenceEnum enum
 */
const (
    UserDatabasePreference_KBACKUPALLDATABASES UserDatabasePreferenceEnum = 1 + iota
    UserDatabasePreference_KBACKUPALLEXCEPTAAGDATABASES
    UserDatabasePreference_KBACKUPONLYAAGDATABASES
)

func (r UserDatabasePreferenceEnum) MarshalJSON() ([]byte, error) {
    s := UserDatabasePreferenceEnumToValue(r)
    return json.Marshal(s)
}

func (r *UserDatabasePreferenceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  UserDatabasePreferenceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts UserDatabasePreferenceEnum to its string representation
 */
func UserDatabasePreferenceEnumToValue(userDatabasePreferenceEnum UserDatabasePreferenceEnum) string {
    switch userDatabasePreferenceEnum {
        case UserDatabasePreference_KBACKUPALLDATABASES:
    		return "kBackupAllDatabases"
        case UserDatabasePreference_KBACKUPALLEXCEPTAAGDATABASES:
    		return "kBackupAllExceptAAGDatabases"
        case UserDatabasePreference_KBACKUPONLYAAGDATABASES:
    		return "kBackupOnlyAAGDatabases"
        default:
        	return "kBackupAllDatabases"
    }
}

/**
 * Converts UserDatabasePreferenceEnum Array to its string Array representation
*/
func UserDatabasePreferenceEnumArrayToValue(userDatabasePreferenceEnum []UserDatabasePreferenceEnum) []string {
    convArray := make([]string,len( userDatabasePreferenceEnum))
    for i:=0; i<len(userDatabasePreferenceEnum);i++ {
        convArray[i] = UserDatabasePreferenceEnumToValue(userDatabasePreferenceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func UserDatabasePreferenceEnumFromValue(value string) UserDatabasePreferenceEnum {
    switch value {
        case "kBackupAllDatabases":
            return UserDatabasePreference_KBACKUPALLDATABASES
        case "kBackupAllExceptAAGDatabases":
            return UserDatabasePreference_KBACKUPALLEXCEPTAAGDATABASES
        case "kBackupOnlyAAGDatabases":
            return UserDatabasePreference_KBACKUPONLYAAGDATABASES
        default:
            return UserDatabasePreference_KBACKUPALLDATABASES
    }
}
