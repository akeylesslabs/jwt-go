package jwt

import "testing"

func Test_mapClaims_string_aud(t *testing.T){
	mapClaims := MapClaims{
		"aud": "foo",
	}
	got := mapClaims.VerifyAudience("foo", true)
	if got != true {
		t.Fatalf("Failed to verify claims, wanted: true got %v",  got)
	}

	got = mapClaims.VerifyAudience("bar", true)
	if got != false {
		t.Fatalf("Failed to verify claims, wanted: false got %v",  got)
	}
}

func Test_mapClaims_string_aud_not_required(t *testing.T){
	mapClaims := MapClaims{
		"aud": "",
	}
	got := mapClaims.VerifyAudience("foo", true)
	if got != false {
		t.Fatalf("Failed to verify claims, wanted: false got %v",  got)
	}

	got = mapClaims.VerifyAudience("foo", false)
	if got != true {
		t.Fatalf("Failed to verify claims, wanted: true got %v",  got)
	}
}

func Test_mapClaims_list_aud_failed(t *testing.T){
	mapClaims := MapClaims{
		"aud": []string{"foo"},
	}
	got := mapClaims.VerifyAudience("foo", true)
	if got != false {
		t.Fatalf("Failed to verify claims, wanted: false got %v", got)
	}

	got = mapClaims.VerifyAudience("bar", true)
	if got != false {
		t.Fatalf("Failed to verify claims, wanted: false got %v", got)
	}

	got = mapClaims.VerifyAudience("bar", false)
	if got != true {
		t.Fatalf("Failed to verify claims, wanted: true got %v", got)
	}
}