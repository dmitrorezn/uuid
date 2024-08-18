package uuid

import (
	"database/sql/driver"
	"unsafe"

	"github.com/google/uuid"
)

var (
	NameSpaceDNS  = Create(uuid.NameSpaceDNS)
	NameSpaceURL  = Create(uuid.NameSpaceURL)
	NameSpaceOID  = Create(uuid.NameSpaceOID)
	NameSpaceX500 = Create(uuid.NameSpaceX500)
	Max           = Create(uuid.Max)

	Nil = Create(uuid.Nil)
)

func New() UUID {
	return Create(
		uuid.New(),
	)
}

func Create(id uuid.UUID) UUID {
	return UUID{
		Uuid: id[:],
	}
}

func NewZero() UUID {
	return UUID{
		Uuid: make([]byte, uuidLen),
	}
}

func Parse(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}

	return Create(id), nil
}

func ParseBytes(b []byte) (UUID, error) {
	id, err := uuid.ParseBytes(b)
	if err != nil {
		return UUID{}, err
	}

	return Create(id), nil
}

func MustParse(s string) UUID {
	return Create(uuid.MustParse(s))
}

func FromBytes(b []byte) (UUID, error) {
	var uid = NewZero()
	if err := uid.UnmarshalBinary(b); err != nil {
		return uid, err
	}

	return uid, nil
}

func Validate(s string) error {
	return uuid.Validate(s)
}

const uuidLen = len(uuid.UUID{})

func (x *UUID) UUID() uuid.UUID {
	return *(*uuid.UUID)(unsafe.Slice(unsafe.SliceData(x.Uuid), uuidLen))
}

func (x *UUID) String() string {
	return x.UUID().String()
}

func (x *UUID) URN() string {
	return x.UUID().URN()
}

func (x *UUID) Variant() uuid.Variant {
	return x.UUID().Variant()
}

func (x *UUID) Version() uuid.Version {
	return x.UUID().Version()
}

func (x *UUID) Domain() uuid.Domain {
	return x.UUID().Domain()
}

func (x *UUID) ID() uint32 {
	return x.UUID().ID()
}

func (x *UUID) Time() uuid.Time {
	return x.UUID().Time()
}

func (x *UUID) ClockSequence() int {
	return x.UUID().ClockSequence()
}

func (x *UUID) MarshalText() ([]byte, error) {
	uid := x.UUID()

	return uid.MarshalText()
}

func (x *UUID) UnmarshalText(data []byte) error {
	uid := x.UUID()

	return uid.UnmarshalText(data)
}

func (x *UUID) MarshalBinary() ([]byte, error) {
	uid := x.UUID()

	return uid.MarshalBinary()
}

func (x *UUID) UnmarshalBinary(data []byte) error {
	uid := x.UUID()

	return uid.UnmarshalBinary(data)
}

func (x *UUID) Scan(src interface{}) error {
	uid := x.UUID()

	return uid.Scan(src)
}

func (x *UUID) Value() (driver.Value, error) {
	return x.UUID().Value()
}

func (x *UUID) NodeID() []byte {
	return x.UUID().NodeID()
}
