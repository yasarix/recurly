package recurly

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestTypeHREFUnmarshal(t *testing.T) {
	type h struct {
		XMLName xml.Name   `xml:"foo"`
		Account hrefString `xml:"account"`
		Invoice hrefInt    `xml:"invoice"`
	}

	expected := h{
		XMLName: xml.Name{Local: "foo"},
		Account: "100abc",
		Invoice: 1108,
	}

	str := bytes.NewBufferString(`<foo><account href="https://your-subdomain.recurly.com/v2/accounts/100abc"/>
    <invoice href="https://your-subdomain.recurly.com/v2/invoices/1108"/></foo>`)

	var given h
	if err := xml.NewDecoder(str).Decode(&given); err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if !reflect.DeepEqual(expected, given) {
		t.Fatalf("unexpected result: %v", given)
	}
}
