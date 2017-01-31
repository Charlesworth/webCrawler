package jsonPrinter

import (
	"bytes"
	"testing"
)

func TestJSONPrinter(t *testing.T) {
	// give the jsonPrinter a buffer instead of using std out
	buffer := new(bytes.Buffer)
	Writter = buffer

	// test printJSON with a valid URL and Asset slice
	PrintPage("hello", []string{"goodbye"})
	PrintPage("1", []string{"2"})
	End()
	expectedResult := "[{\"url\":\"hello\",\"assets\":[\"goodbye\"]},\n{\"url\":\"1\",\"assets\":[\"2\"]}]\n"
	output := buffer.String()
	if output != expectedResult {
		t.Error("printJSON is output is incorrect, expected:\n", expectedResult, "\nResult:\n", output)
	}

	// give the jsonPrinter a buffer instead of using std out
	buffer = new(bytes.Buffer)
	Writter = buffer

	// Test the reset function
	Reset()
	PrintPage("", []string{})
	End()
	expectedResult = "[{\"url\":\"\",\"assets\":[]}]\n"
	output = buffer.String()
	if output != expectedResult {
		t.Error("Reset did not return jsonPrinter to its oridginal state")
	}

}
