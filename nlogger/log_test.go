package nlogger

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const msg = "I am message"

func TestNewLogger(t *testing.T) {
	var str string
	target := bytes.NewBufferString(str)
	lg := New(target, "")
	assert.NotNil(t, lg)
}

func TestPrintLogInDebug(t *testing.T) {
	var str string

	target := bytes.NewBufferString(str)
	lg := New(target, "")
	lg.Debug(msg)

	assert.Contains(t, target.String(), msg)
}

func TestPrintLogInInfo(t *testing.T) {
	tempFile, _ := ioutil.TempFile(".", "common-logger-output")
	defer os.Remove(tempFile.Name())
	lg := New(tempFile, "")
	lg.Info(msg)

	b, _ := ioutil.ReadFile(tempFile.Name())

	assert.Contains(t, string(b), msg)
}
func TestPrintLogInWarning(t *testing.T) {
	var str string

	target := bytes.NewBufferString(str)
	lg := New(target, "")
	lg.Warn(msg)

	assert.Contains(t, target.String(), msg)
}
func TestPrintLogInError(t *testing.T) {
	var str string

	target := bytes.NewBufferString(str)
	lg := New(target, "")
	lg.Error(msg)

	assert.Contains(t, target.String(), msg)
}

func TestPrintLogInInfoWithPreFix(t *testing.T) {
	var str string
	const prefix = "aaaaaaaaaaaa"
	target := bytes.NewBufferString(str)
	lg := New(target, prefix)
	lg.Error(msg)

	assert.Contains(t, target.String(), msg)
	assert.Contains(t, target.String(), prefix)
}

func TestProvider(t *testing.T) {
	var l = New(ioutil.Discard, "")
	var n = NewProvider(l)

	if l != n.Get() {
		t.Error("Logger from provider is not the same")
	}

	var l2 = New(ioutil.Discard, "")
	n.Replace(l2)

	if l2 != n.Get() {
		t.Error("Logger from provider is not the same as the one we called replace with")
	}
}
