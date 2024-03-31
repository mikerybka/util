package util_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mikerybka/util"
)

type Test struct {
	A string
	B int
	C bool
	D []string
	E map[string]string
}

func TestAPI(t *testing.T) {
	defaultt := Test{}
	initial := Test{
		A: "hi",
		B: 12,
		C: true,
		D: []string{"1"},
		E: map[string]string{
			"a": "b",
		},
	}
	api := util.NewAPI(defaultt)

	test(api, t, "GET", "/", nil, util.JSONString(defaultt))
	test(api, t, "PUT", "/", initial, util.JSONString(initial))
	test(api, t, "GET", "/", nil, util.JSONString(initial))

	test(api, t, "GET", "/A", nil, "hi")
	test(api, t, "PUT", "/A", "test", "test")
	test(api, t, "GET", "/A", nil, "test")
	test(api, t, "GET", "/B", nil, 12)
	test(api, t, "PUT", "/B", 22, 22)
	test(api, t, "GET", "/B", nil, 22)
	test(api, t, "GET", "/C", nil, true)
	test(api, t, "PUT", "/C", false, false)
	test(api, t, "GET", "/C", nil, false)

	// test(api, t, "GET", "/D", nil, []string{"1"})
	// test(api, t, "GET", "/D/0", nil, "1")
	// test(api, t, "PUT", "/D", []string{"a", "b"}, []string{"a", "b"})
	// test(api, t, "GET", "/D/0", nil, "a")
	// test(api, t, "GET", "/D/1", nil, "b")
	// test(api, t, "POST", "/D", "c", []string{"a", "b", "c"})
	// test(api, t, "GET", "/D", nil, []string{"a", "b", "c"})
	// test(api, t, "GET", "/D/2", nil, "c")
	// test(api, t, "DELETE", "/D/1", nil, []string{"a", "c"})
	// test(api, t, "GET", "/D/1", nil, "c")

	// test(api, t, "GET", "/E", nil, map[string]string{
	// 	"a": "b",
	// })
	// test(api, t, "PUT", "/E/abc", "123", "123")
	// test(api, t, "GET", "/E", nil, map[string]string{
	// 	"a":   "b",
	// 	"abc": "123",
	// })
	// test(api, t, "GET", "/E/abc", nil, "123")
}

func test(api *util.API[Test], t *testing.T, method string, path string, body any, expected any) {
	w := httptest.NewRecorder()
	b := bytes.NewBuffer(nil)
	if body != nil {
		err := json.NewEncoder(b).Encode(&body)
		if err != nil {
			panic(err)
		}
	}
	r := httptest.NewRequest(method, path, b)

	api.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		panic(err)
	}
	expectedString := string(expectedBytes)
	resBytes, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("invalid response body: %s", err.Error())
	}
	resString := string(resBytes)
	if resString != expectedString {
		t.Errorf("expected %s, got %s", expectedString, resString)
	}
}
