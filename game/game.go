package game

import (
	"fmt"
	"strings"
)

func ListItems(items []string) string {
    var sb strings.Builder
    sb.WriteString("You can see here ")
    if len(items) < 3 {
        s := fmt.Sprintf("%s and %s.", items[0], items[1])
        sb.WriteString(s)
        return sb.String()
    }
    sb.WriteString(strings.Join(items[:len(items) -1], ", "))
    sb.WriteString(", and ")
    sb.WriteString(items[len(items) -1])
    sb.WriteString(".")
    return sb.String()
}