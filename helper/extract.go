package helper

import (
	"strings"
)

func ExtractDuplicateField(err string) string {
	items := strings.Split(err,"_")
	return items[1]
}
