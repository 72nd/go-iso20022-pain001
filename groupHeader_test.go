package pain001

import (
	"regexp"
	"testing"
)

func TestNewGroupHeader(t *testing.T) {
}

func TestGetShortId(t *testing.T) {
	id := getShortId()
	if len(id) != 18 {
		t.Errorf("id \"%s\" has the wrong length %d not 18", id, len(id))
	}
	re := regexp.MustCompile(`^.*-.*-.*$`)
	if !re.MatchString(id) {
		t.Errorf("id \"%s\" has not the format \"^.*-.*-.*$\"", id)
	}
}
