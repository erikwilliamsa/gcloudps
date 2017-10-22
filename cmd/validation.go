package cmd

import "errors"

// CheckRequiredFlags Checks if required flags are there.
// Will return an error for any invalid flag
func CheckRequiredFlags(flags map[string]string) error {
	if len(ProjectName) <= 0 || len(TopicName) <= 0 {
		return errors.New("project and topic must be set")
	}

	for key, val := range flags {
		if len(val) <= 0 {
			return errors.New(key + " must be set")
		}
	}
	return nil
}
