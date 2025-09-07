package orderedmap

import "testing"

var tests = []struct {
	name     string
	key      string
	value    string
	expected bool
}{
	{
		name:     "foo",
		key:      "foo",
		value:    "fooval",
		expected: true,
	},
	{
		name:     "bar",
		key:      "bar",
		value:    "barval",
		expected: true,
	},
	{
		name:     "fooagain",
		key:      "foo",
		value:    "fooval2",
		expected: false,
	},
	{
		name:     "baz",
		key:      "baz",
		value:    "bazval",
		expected: true,
	},
}

func TestOrderedMapSet(t *testing.T) {
	om := NewOrderedMap[string, string]()
	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ok := om.Set(tst.key, tst.value)
			if ok != tst.expected {
				t.Errorf("expected %v, got %v", tst.expected, ok)
			}
		})
	}

	for k, v := range om.KVPairs() {
		t.Logf("%s => %s", k, v)
	}
}

func TestOrderedMapInsert(t *testing.T) {
	om := NewOrderedMap[string, string]()
	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ok := om.Insert(tst.key, tst.value)
			if ok != tst.expected {
				t.Errorf("expected %v, got %v", tst.expected, ok)
			}
		})
	}
	for k, v := range om.KVPairs() {
		t.Logf("%s => %s", k, v)
	}
}

func TestMapUpdate(t *testing.T) {
	om := NewOrderedMap[string, string]()
	// seed the values initially since update will fail if not exists
	for _, tst := range tests {
		om.Set(tst.key, tst.value)
	}
	if om.Update("thisisnotset", "someval") {
		t.Errorf("expected false")
	}
	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if !om.Update(tst.key, tst.value) {
				t.Errorf("expected true")
			}
		})
	}

	for k, v := range om.KVPairs() {
		t.Logf("%s => %s", k, v)
	}
}

func TestMapDelete(t *testing.T) {
	om := NewOrderedMap[string, string]()
	for _, tst := range tests {
		om.Set(tst.key, tst.value)
	}
	if om.Delete("thisisnotset") {
		t.Errorf("expected false")
	}
	if !om.Delete("foo") {
		t.Errorf("expected true")
	}
	for k, v := range om.KVPairs() {
		t.Logf("%s => %s", k, v)
	}
	t.Logf("%#v", om.ord)
}
