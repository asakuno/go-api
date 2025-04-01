package custom_types

import (
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type UUID[T any] uuid.UUID

func (u *UUID[T]) GormDataType() string {
	return "binary(16)"
}

func (u *UUID[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {

	return "binary(16)"
}

func (u *UUID[T]) Scan(value any) (err error) {
	bytes, ok := value.([]byte)
	if !ok {
		str, ok := value.(string)
		if ok {
			bytes = []byte(str)
		} else {
			return errors.New("cannot scan uuid: unsupported type")
		}
	}

	if len(bytes) == 0 {
		*u = UUID[T](uuid.Nil)
		return nil
	}

	parseByte, err := uuid.FromBytes(bytes)
	if err != nil {
		parseStr, errStr := uuid.Parse(string(bytes))
		if errStr != nil {
			return err
		}
		parseByte = parseStr
	}
	*u = UUID[T](parseByte)
	return nil
}

func (u UUID[T]) Value() (driver.Value, error) {
	if uuid.UUID(u) == uuid.Nil {
		return nil, nil
	}
	return uuid.UUID(u).MarshalBinary()
}

func (u UUID[T]) String() string {
	return uuid.UUID(u).String()
}

func NewUUID[T any]() UUID[T] {
	uid, err := uuid.NewV7()
	if err != nil {
		return UUID[T](uuid.New())
	}
	return UUID[T](uid)
}

func ParseUUID[T any](s string) (UUID[T], error) {
	uid, err := uuid.Parse(s)
	return UUID[T](uid), err
}

func MustParseUUID[T any](s string) UUID[T] {
	uid := uuid.Must(uuid.Parse(s))
	return UUID[T](uid)
}
