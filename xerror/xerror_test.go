package xerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	err := Errorf(Normal, "test error")
	// t.Logf("err: %+v", err)
	assert.NotNil(t, err)

	var xerr *XError
	assert.True(t, errors.As(err, &xerr))
	assert.Equal(t, xerr.ErrType, Normal)
	assert.Equal(t, xerr.Err.Error(), "test error")
}

func TestWrap(t *testing.T) {
	err := errors.New("db open error")
	wrappedErr := Wrap(err, DB, "wrapped error")
	// t.Logf("wrappedErr: %+v", wrappedErr)
	assert.NotNil(t, wrappedErr)

	var xerr *XError
	assert.True(t, errors.As(wrappedErr, &xerr))
	assert.Equal(t, xerr.ErrType, DB)
	assert.Equal(t, xerr.Err.Error(), "db open error")
}

func TestWrapf(t *testing.T) {
	err := errors.New("fe test error")
	wrappedErr := Wrapf(err, FE, "wrapped error: %s", "foo")
	// t.Logf("wrappedErr: %+v", wrappedErr)
	assert.NotNil(t, wrappedErr)

	var xerr *XError
	assert.True(t, errors.As(wrappedErr, &xerr))
	assert.Equal(t, xerr.ErrType, FE)
	assert.Equal(t, xerr.Err.Error(), "fe test error")
}
