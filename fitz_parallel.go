package main

import (
  "fmt"
  "image/jpeg"
  "log"
  "os"
  "path/filepath"
  "github.com/gen2brain/go-fitz"
  "time"
  "sync"
)
var tmpDir = "./output"

func createOutputFolder() {
  if _, err := os.Stat(tmpDir); !os.IsNotExist(err) {
    fmt.Println("Deleting existing folder")
    os.RemoveAll(tmpDir)
  }
  fmt.Println("Creating folder")
  os.Mkdir(tmpDir, os.ModePerm)
}

func doSomething (n int) {
  fmt.Println("Processing", n)
  time.Sleep(1000 * time.Millisecond)
}

func generate(doc *fitz.Document, number int) {
  fmt.Println("Generating image from page", number)
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

  err = jpeg.Encode(f, img, &jpeg.Options{1})
  if err != nil {
    log.Fatal("Error while saving image")
    panic(err)
  }
  fmt.Println("Write finished", number)
}

func main() {
  maxGoroutines := 4
  semaphore := make(chan struct{}, maxGoroutines)

  doc, err := fitz.New("12MB.pdf")
  // doc, err := fitz.New("10page.pdf")
  if err != nil {
    panic(err)
  }
  defer doc.Close()

  createOutputFolder()
  numPages := doc.NumPage()
  var wg sync.WaitGroup
  wg.Add(numPages)

  fmt.Println("===> Number of Pages", numPages)
  for number := 0; number < numPages; number++ {
    semaphore <- struct{}{}
    go func(doc *fitz.Document, number int, wg *sync.WaitGroup) {
      generate(doc, number)
      // doSomething(number)
      <-semaphore
      wg.Done()
    }(doc, number, &wg)
  }

  close(semaphore)
  wg.Wait()
  fmt.Println("===> Done")
}
