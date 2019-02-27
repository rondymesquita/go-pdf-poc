package main

import (
  "sync"
  "fmt"
  "image/jpeg"
  "log"
  "os"
  "path/filepath"
  "github.com/gen2brain/go-fitz"
  "reflect"
  "runtime"
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

func generate(doc *fitz.Document, number int) {
  fmt.Println("Generating image from page", number)
  img, err := doc.Image(number)
  if err != nil {
    log.Fatal("Error while extracting image", err)
    panic(err)
  }
  name := fmt.Sprintf("test%03d.jpg", number)
  // fmt.Println("Image Name", name)

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
    runtime.GOMAXPROCS(2)
    doc, err := fitz.New("12MB.pdf")
    fmt.Println(reflect.TypeOf(doc))
    if err != nil {
      panic(err)
    }
    defer doc.Close()

    createOutputFolder()
    var wg sync.WaitGroup
    numPages := doc.NumPage()
    wg.Add(numPages)

    fmt.Println("===> Number of Pages", numPages)
    for number := 0; number < numPages; number++ {
      fmt.Println("=>", number)
      go func(doc *fitz.Document, number int, wg *sync.WaitGroup) {
        fmt.Println("Called")
        defer wg.Done()
        generate(doc, number)
      }(doc, number, &wg)
    }
   wg.Wait()
   fmt.Println("===> Done")
}
