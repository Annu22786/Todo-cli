package main

import(
  "fmt"
//  "encoding/json"
 // "os"
)
type Storage[T any] struct {
  Filename string
}

// file name declaration func
func declarename[T any](filename string) *Storage[T]{
  return &Storage [T]{
    Filename : filename}
}
// save func

//checker file
func main(){
  fmt.Println("hrllo")
}