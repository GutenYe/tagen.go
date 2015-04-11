package strings2

import (
	"regexp"
	"bytes"
  "strings"
)

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(str string) string {
  chunks := camelingRegex.FindAll([]byte(str), -1)
  for i, v := range chunks {
    if i > 0 {
      chunks[i] = bytes.Title(v)
    }
  }
  return string(bytes.Join(chunks, nil))
}

func ClassCase(str string) string {
  return strings.Title(CamelCase(str))
}
