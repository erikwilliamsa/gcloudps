package formatters

import (
	"bytes"
	"encoding/json"
)

//Formatter interface for passing types of formatters to workers
type Formatter interface {
	Format(data string) string
}

// JSONFormatter implmentation of format
type JSONFormatter struct{}

func (jf *JSONFormatter) Format(data string) string {
	return FormatJSON(data)
}

// FormatJSON takes a string and returns a formatted JSON string
// if the string passed is a valid JSON string
func FormatJSON(data string) string {
	buffer := new(bytes.Buffer)
	err := json.Indent(buffer, []byte(data), "", "    ")
	if err == nil {
		s := buffer.String()
		return s
	}
	return data
}
