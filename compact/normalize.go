package compact

import "strings"

var r = strings.NewReplacer("_", "", "^", "", "+", "")

// stripAccents removes any accent chars and converts to lower case.
func stripAccents(accented string) string {
	return r.Replace(accented)
}
