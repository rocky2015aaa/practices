package main

import "testing"

func TestPrintConvertedUri(t *testing.T) {
	encodingMap := map[string]string{
		"%20": " ",
		"%21": "!",
		"%24": "$",
		"%25": "%",
		"%28": "(",
		"%29": ")",
		"%2a": "*",
	}

	inputStrings := []string{"http://%21%24", "ftp://%28%29"}

	newInputStrings := ConvertUri(inputStrings, encodingMap)
	if newInputStrings[0] != "http://!$" {
		t.Error("expected", true, "got", false)
	}
	if newInputStrings[1] != "ftp://()" {
		t.Error("expected", true, "got", false)
	}
}
