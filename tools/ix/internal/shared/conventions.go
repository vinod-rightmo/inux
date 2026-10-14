package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
)

// CheckConvention validates every supplied value against the first argument’s regex.
func CheckConvention(args []string, out io.Writer, usage string) error {
	if len(args) < 2 {
		return fmt.Errorf("use: %s", usage)
	}
	pattern, err := regexp.Compile(args[0])
	if err != nil {
		return fmt.Errorf("invalid convention pattern: %w", err)
	}
	// Choose the longest alternative so a|abc can match all of abc.
	pattern.Longest()
	var failed []int
	for i, value := range args[1:] {
		span := pattern.FindStringIndex(value)
		if span == nil || span[0] != 0 || span[1] != len(value) {
			failed = append(failed, i+1)
		}
	}
	if len(failed) > 0 {
		return fmt.Errorf("values at positions %v do not match the supplied convention", failed)
	}
	return json.NewEncoder(out).Encode(map[string]bool{"ok": true})
}
