package namegen // import "github.com/docker/docker/pkg/namesgenerator"

import (
	"strings"
	"testing"
)

func Test_GetName_Format(t *testing.T) {
	name := GetName(0)
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func Test_GetNameWithSeed_Format(t *testing.T) {
	name := GetNameWithSeed(nil, 0)
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func Test_GetName_Retries(t *testing.T) {
	name := GetName(1)
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
	if !strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name doesn't contain a number")
	}
}

func Test_GetNameWithSeed_Retries(t *testing.T) {
	name := GetNameWithSeed(nil, 1)
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
	if !strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name doesn't contain a number")
	}
}

func Benchmark_GetName(b *testing.B) {
	b.ReportAllocs()
	var out string
	for n := 0; n < b.N; n++ {
		out = GetName(5)
	}
	_ = out
	// b.Log("Last result:", out)
}

func Benchmark_GetNameWithSeed(b *testing.B) {
	b.ReportAllocs()
	var out string
	for n := 0; n < b.N; n++ {
		out = GetNameWithSeed(nil, 5)
	}
	_ = out
	// b.Log("Last result:", out)
}
