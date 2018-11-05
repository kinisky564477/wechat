package core_test

import (
	"testing"

	"github.com/zjxpcyc/wechat/core"
)

func TestParse(t *testing.T) {
	xmlStr := "<xml><a><![CDATA[b]]></a><c>d</c></xml>"
	expected := map[string]string{
		"a": "b",
		"c": "d",
	}

	xml := &core.XMLParse{}
	res, err := xml.Parse(xmlStr)
	if err != nil {
		t.Fatalf("Parse xml fail, %s", err.Error())
	}

	found := true
	for k, v := range res {
		if !found {
			continue
		}

		val := expected[k]
		if val != v {
			found = false
		}
	}

	if !found {
		t.Fatalf("Parse xml fail")
	}
}
