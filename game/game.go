package game

import "strings"

func ListItems(items []string) string {
    var sb strings.Builder
    sb.WriteString("You can see here ")
    sb.WriteString(strings.Join(items[:len(items) -1], ", "))
    sb.WriteString(", and ")
    sb.WriteString(items[len(items) -1])
    sb.WriteString(".")
    return sb.String()
}