package pdf

import (
  "log"
  "image/jpeg"
  "fmt"
  "os"
  "path/filepath"
  "github.com/gen2brain/go-fitz"
  "time"
  "sync"
)
var tmpDir = "./output"
const FILENAME = "../10pages.pdf"

func createOutputFolder() {
  if _, err := os.Stat(tmpDir); !os.IsNotExist(err) {
    log.Println("Deleting existing folder")
    os.RemoveAll(tmpDir)
  }
  log.Println("Creating folder")
  os.Mkdir(tmpDir, os.ModePerm)
}

func fakeGenerate (n int) {
  log.Println("Processing", n)
  time.Sleep(1000 * time.Millisecond)
}

func extractImages(doc *fitz.Document, number int) {
  log.Println("Generating image from page", number)
  img, err := doc.ImageDPI(number, 72)
  if err != nil {
    log.Fatal("Error while extracting image", err)
    panic(err)
  }
  name := fmt.Sprintf("test%03d.jpg", number)

  f, err := os.Create(filepath.Join(tmpDir, name))
  defer f.Close()

  if err != nil {
    log.Fatal("Error while creating file")
    panic(err)
  }

  err = jpeg.Encode(f, img, &jpeg.Options{80})
  if err != nil {
    log.Fatal("Error while saving image")
    panic(err)
  }
  log.Println("Write finished", number)
}

func generateInSequence(file string) {
  doc, err := fitz.New("../" + file)
  if err != nil {
    panic(err)
  }
  defer doc.Close()

  createOutputFolder()
  numPages := doc.NumPage()
  log.Println("===> Number of Pages", numPages)
  for number := 0; number < numPages; number++ {
    extractImages(doc, number)
  }

  log.Println("===> Done")
}

func generateInParallel(file string, maxGoroutines int) {
  semaphore := make(chan struct{}, maxGoroutines)

  doc, err := fitz.New("../" + file)
  if err != nil {
    panic(err)
  }
  defer doc.Close()

  createOutputFolder()
  numPages := doc.NumPage()
  var wg sync.WaitGroup
  wg.Add(numPages)

  log.Println("===> Number of Pages", numPages)
  for number := 0; number < numPages; number++ {
    semaphore <- struct{}{}
    go func(doc *fitz.Document, number int, wg *sync.WaitGroup) {
      extractImages(doc, number)
      <-semaphore
      wg.Done()
    }(doc, number, &wg)
  }

  close(semaphore)
  wg.Wait()
  log.Println("===> Done")
}

func Generate(mode, file string, maxGoroutines int) {
  if mode == "sequence" {
    // fmt.Println("Running in Sequence with file", file)
    generateInSequence(file)
  } else {
    // fmt.Println("Running in Parallel with file", file)
    generateInParallel(file, maxGoroutines)
  }
}

func main() {
  Generate("sequence", "10pages.pdf", 1)
}
