package ormerrors

import (
	"github.com/cosmos/cosmos-sdk/errors"
	"google.golang.org/grpc/codes"
)

var codespace = "orm"

// IsNotFound returns true if the error indicates that the record was not found.
func IsNotFound(err error) bool {
	return errors.IsOf(err, NotFound)
}

var (
	UniqueKeyViolation = errors.RegisterWithGRPCCode(codespace, 24, codes.FailedPrecondition, "unique key violation")
	NotFound           = errors.RegisterWithGRPCCode(codespace, 29, codes.NotFound, "not found")
	AlreadyExists      = errors.RegisterWithGRPCCode(codespace, 31, codes.AlreadyExists, "already exists")

	InvalidTableId                = errors.New(codespace, 1, "invalid or missing table or single id, need a non-zero value")
	MissingPrimaryKey             = errors.New(codespace, 2, "table is missing primary key")
	FieldNotFound                 = errors.New(codespace, 5, "field not found")
	InvalidAutoIncrementKey       = errors.New(codespace, 6, "an auto-increment primary key must specify a single uint64 field")
	InvalidIndexId                = errors.New(codespace, 7, "invalid or missing index id, need a value >= 0 and < 32768")
	DuplicateIndexId              = errors.New(codespace, 8, "duplicate index id")
	PrimaryKeyConstraintViolation = errors.New(codespace, 9, "object with primary key already exists")
	PrimaryKeyInvalidOnUpdate     = errors.New(codespace, 11, "can't update object with missing or invalid primary key")
	AutoIncrementKeyAlreadySet    = errors.New(codespace, 12, "can't create with auto-increment primary key already set")
	UnexpectedDecodePrefix        = errors.New(codespace, 14, "unexpected prefix while trying to decode an entry")
	BytesFieldTooLong             = errors.New(codespace, 15, "bytes field is longer than 255 bytes")
	BadDecodeEntry                = errors.New(codespace, 17, "bad decode entry")
	IndexOutOfBounds              = errors.New(codespace, 18, "index out of bounds")
	UnsupportedKeyField           = errors.New(codespace, 20, "unsupported key field")
	UnexpectedError               = errors.New(codespace, 21, "unexpected error")
	InvalidRangeIterationKeys     = errors.New(codespace, 22, "invalid range iteration keys")
	JSONImportError               = errors.New(codespace, 23, "json import error")
	InvalidTableDefinition        = errors.New(codespace, 25, "invalid table definition")
	InvalidFileDescriptorID       = errors.New(codespace, 26, "invalid file descriptor ID")
	TableNotFound                 = errors.New(codespace, 27, "table not found")
	JSONValidationError           = errors.New(codespace, 28, "invalid JSON")
	ReadOnly                      = errors.New(codespace, 30, "database is read-only")
)
