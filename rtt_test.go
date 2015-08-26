package main

import (
	"log"
	"os"
	"testing"

	"github.com/dt/thile/gen"
)

var uncompressed gen.HFileService
var compressed gen.HFileService
var maxKey int

func TestMain(m *testing.M) {
	maxKey = 15000000
	var err error
	uncompressed, err = GetTestIntFile("uncompressed", maxKey, false, true)
	if err != nil {
		log.Fatal(err)
	}
	compressed, err = GetTestIntFile("compressed", maxKey, true, true)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestCompressed(t *testing.T) {
	reqs := GetRandomTestReqs("compressed", 10, 5, maxKey)

	for _, req := range reqs {
		if r, err := compressed.GetValuesSingle(req); err != nil {
			t.Fatal("error: ", err)
		} else {
			if len(r.GetValues()) != len(req.GetSortedKeys()) {
				t.Fatal("wrong number of results: ", "\n", req.GetSortedKeys(), "\n", r.GetValues())
			}
		}
	}
}

func BenchmarkUncompressed(b *testing.B) {
	b.StopTimer()
	reqs := GetRandomTestReqs("uncompressed", b.N, 5, maxKey)
	b.StartTimer()

	for _, req := range reqs {
		if _, err := uncompressed.GetValuesSingle(req); err != nil {
			b.Fatal("error: ", err)
		}
	}
}

func BenchmarkCompressed(b *testing.B) {
	b.StopTimer()
	reqs := GetRandomTestReqs("compressed", b.N, 5, maxKey)
	b.StartTimer()

	for _, req := range reqs {
		if _, err := compressed.GetValuesSingle(req); err != nil {
			b.Fatal("error: ", err)
		}
	}
}
