package pain001

import (
	"encoding/xml"
	"testing"
)

func compareXmlResult(t *testing.T, data interface{}, expected string) {
	out, err := xml.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	if expected != string(out) {
		t.Errorf("output \"%s\" didn't match expected string \"%s\"", string(out), expected)
	}

}
