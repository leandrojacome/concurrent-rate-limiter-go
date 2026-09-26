package domain

import "testing"

func TestClientKeyPreservesInvariant(t *testing.T) {
	key, err := NewClientKey(" customer-42 ")
	if err != nil || key != "customer-42" {
		t.Fatalf("key=%q err=%v", key, err)
	}
	if _, err := NewClientKey("x"); err == nil {
		t.Fatal("expected invalid short key")
	}
}
