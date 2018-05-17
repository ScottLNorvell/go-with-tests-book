package iteration

import "strings"

func Repeat(char string, count int) (repeated string) {
  for i := 0; i < count; i++ {
    repeated += char
  }
  return repeated
}

func RepeatAnotherWay(s string, i int) string {
  return strings.Repeat(s, i)
}
