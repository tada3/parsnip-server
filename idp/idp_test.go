package idp_test

import (
	"fmt"
	"testing"

	"git-dev.linecorp.com/tadafumi-yoshihara/idpgo/idp"
)

func Test_getSession(t *testing.T) {
	baseURL := "https://biz-idp-internal-api.line-apps-beta.com/"
	token := "lMG4qbDt8reNLtkOL9cK2H9UyJz4P6ZEZwDocYfFgaU"
	sessionID := "hoge"

	c, err := idp.New(baseURL, token, 1000)
	if err != nil {
		t.Fatal(err)
	}
	sr, err := c.GetSession(sessionID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("sr = %+v\n", sr)
}

func Test_generateLogoutURL(t *testing.T) {
	baseURL := "https://biz-idp-internal-api.line-apps-beta.com/"
	token := "lMG4qbDt8reNLtkOL9cK2H9UyJz4P6ZEZwDocYfFgaU"
	sessionID := "hoge"

	c, err := idp.New(baseURL, token, 1000)
	if err != nil {
		t.Fatal(err)
	}
	sr, err := c.GetSession(sessionID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("sr = %+v\n", sr)
}
