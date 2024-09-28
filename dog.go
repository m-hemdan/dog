package dog

import (
	"strings"
)

func WhenGrownUps(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: lowercase " + strings.ToLower(s)
}
