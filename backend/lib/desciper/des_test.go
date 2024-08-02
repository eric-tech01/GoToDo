package desCiper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDES(t *testing.T) {
	var err error
	d := &DesCiper{SecretKey: "12345678", PlainText: "hello taurus"}
	err = d.Encrypt()
	assert.Equal(t, err, nil)

	d2 := &DesCiper{SecretKey: d.SecretKey, CiperText: d.CiperText}
	err = d2.DESDecrypt()
	assert.Equal(t, err, nil)

	assert.Equal(t, d.PlainText, d2.PlainText)
}
