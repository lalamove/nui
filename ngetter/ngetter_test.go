package ngetter

import "testing"

func TestProvider(t *testing.T) {
	var n = NewProvider(true)

	if true != n.Get().(bool) {
		t.Error("Logger from provider is not the same")
	}

	n.Replace(false)

	if false != n.Get().(bool) {
		t.Error("Logger from provider is not the same as the one we called replace with")
	}
}
