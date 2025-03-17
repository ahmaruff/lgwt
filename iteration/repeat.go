package iteration

import "strings"

func Repeat(char string, count int) string {
	var rep strings.Builder

	for i := 0; i < count; i++ {
		rep.WriteString(char)
	}
	return rep.String()
}
