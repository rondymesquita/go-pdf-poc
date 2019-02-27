 package main

import (
  // "runtime"
  "fmt"
  // "sync"
  "time"
)

func doSomething () {
  time.Sleep(1000 * time.Millisecond)
}

func main() {
    maxGoroutines := 2
    numberOfJobs := 10
    semaphore := make(chan struct{}, maxGoroutines)

    for i := 0; i < numberOfJobs; i++ {
        semaphore <- struct{}{}
        go func(n int) {
          fmt.Println("Processing", n)
            doSomething()
            <-semaphore
        }(i)
    }
}
