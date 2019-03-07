package pdf

import (
  "testing"
  "log"
  "io/ioutil"
  "flag"
)

var (
  mode = flag.String("mode", "", "Set parallel or sequence mode for image extraction")
  maxGoroutines = flag.Int("maxGoroutines", 1, "Set max of goroutines to run in parallel mode")
  file = flag.String("file", "", "PDF file to be used")
)


func BenchmarkGenerate(b *testing.B) {
  log.SetOutput(ioutil.Discard)
  for n := 0; n < b.N; n++ {
    Generate(*mode, *file, *maxGoroutines)
  }
}
