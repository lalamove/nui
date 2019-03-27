package nlogger

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	var buf = new(bytes.Buffer)
	var l = New(buf, "")
	var n = NewProvider(l)

	if l != n.Get() {
		t.Error("Logger from provider is not the same")
	}

	type testStructured struct {
		basicStructured
	}
	// test the panic handler
	n.Replace(&testStructured{
		basicStructured{&defaultLogger{log.New(ioutil.Discard, "", log.LstdFlags)}},
	})
	var msg = buf.String()
	var expMsg = "your new logger is not the same concrete type of logger as your old logger, " +
		"we will continue using the old logger oldLoggerType=*nlogger.basicStructured newLoggerType=*nlogger.testStructured"

	if strings.Compare(msg[20:], expMsg) == 0 {
		t.Error("Did not reach the panic handler when trying to replace logger with another concrete type ")
	}

	if l != n.Get() {
		t.Error("Logger from provider was replaced with that of another concrete type")
	}

	var l2 = New(ioutil.Discard, "")
	n.Replace(l2)

	if l2 != n.Get() {
		t.Error("Logger from provider is not the same as the one we called replace with")
	}
}
