package imagick

import (
    "log"
    "strconv"
    "gopkg.in/gographics/imagick.v2/imagick"
    "sync"
)

func main() {

    pdfName := "10pages.pdf"
    // imageName := "test.jpg"

      if err := ConvertPdfToJpg(pdfName); err != nil {
          log.Fatal(err)
      }
}

// ConvertPdfToJpg will take a filename of a pdf file and convert the file into an 
// image which will be saved back to the same location. It will save the image as a 
// high resolution jpg file with minimal compression.
func ConvertPdfToJpg(pdfName string) error {
    res := 300.0
    var quality uint
    quality = 100

    // Setup
    imagick.Initialize()
    log.Println("Initialized")
    defer imagick.Terminate()

    mw := imagick.NewMagickWand()
    log.Println("Created")
    defer mw.Destroy()

    var wg sync.WaitGroup
    var mutex = &sync.Mutex{}

    // Must be *before* ReadImageFile
    // Make sure our image is high quality
    if err := mw.SetResolution(res, res); err != nil {
        return err
    }
    log.Println("Resolution set")

    // Load the image file into imagick
    if err := mw.ReadImage(pdfName); err != nil {
        return err
    }
    log.Println("File readed")

    // Must be *after* ReadImageFile
    // Flatten image and remove alpha channel, to prevent alpha turning black in jpg
    if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
        return err
    }
    log.Println("SetImageAlphaChannel")

    // Set any compression (100 = max quality)
    if err := mw.SetCompressionQuality(quality); err != nil {
        return err
    }
    log.Println("Set Quality")

    numberOfPages := 10
    log.Println("Pages ", numberOfPages)
    wg.Add(numberOfPages)
    defer wg.Wait()

    for i := 0; i < numberOfPages; i++ {
      go func(i int, wg *sync.WaitGroup) {
        defer wg.Done()
        mutex.Lock()
        defer mutex.Unlock()
        log.Println("Page ", i)

        mw.SetIteratorIndex(i)

        if err := mw.SetFormat("jpg"); err != nil {
            log.Fatal(err)
        }
        log.Println("SetFormat")

        imageName := strconv.Itoa(i)
         if err := mw.WriteImage("./output/" + imageName); err != nil {
            log.Fatal(err)
        }
        log.Println("WriteImage")
      }(i, &wg)
    }

    return nil
}
