package main

import "testing"

func TestSetConfig(t *testing.T) {
	var actual APIConf
	err := SetConfig(&actual)
	if err != nil {
		t.Errorf("read conf file error: %#v", err)
	}
}
