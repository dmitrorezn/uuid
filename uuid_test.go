package uuid

import (
	"slices"
	"testing"

	"github.com/google/uuid"
)

func TestParse(t *testing.T) {
	_uid := uuid.New()
	s := _uid.String()

	uid, err := Parse(s)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	parsed := uid.UUID()

	if !slices.Equal(parsed[:], _uid[:]) {
		t.Fail()
		return
	}
}

func TestFromBytes(t *testing.T) {
	_uid := uuid.New()
	s, err := _uid.MarshalBinary()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	uid, err := FromBytes(s)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	parsed := uid.UUID()

	if !slices.Equal(parsed[:], _uid[:]) {
		t.Fail()
		return
	}
}
