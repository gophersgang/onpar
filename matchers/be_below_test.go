package matchers_test

import (
	"reflect"
	"testing"

	"github.com/apoydence/onpar/matchers"
)

func TestBelow(t *testing.T) {
	t.Parallel()

	m := matchers.BeBelow(101)

	v, err := m.Match(99.0)
	if err != nil {
		t.Error("expected err to be nil")
	}

	if !reflect.DeepEqual(v, 99.0) {
		t.Errorf("expected %v to equal %v", v, 99.0)
	}

	_, err = m.Match(103.0)
	if err == nil {
		t.Error("expected err to not be nil")
	}

	_, err = m.Match("invalid")
	if err == nil {
		t.Error("expected err to not be nil")
	}
}
