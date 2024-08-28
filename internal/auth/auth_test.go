package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"ApiKey YWxhZGRpbjpvcGVuc2VzYW1l"},
	}
	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	want := "YWxhZGRpbjpvcGVuc2VzYW1l"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
