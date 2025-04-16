package pkg

import "strings"

type StringFormatter struct{}

func (s *StringFormatter) UpperCase(str string) string {
	return strings.ToUpper(str)
}

func (s *StringFormatter) LowerCase(str string) string {
	return strings.ToLower(str)
}

func NewStringFormatter() *StringFormatter {
	return &StringFormatter{}
}
