package main

import(
  "fmt"
  "encoding/json"
  "os"
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
func (s *Storage[T]) save(data T)error{
  funcdata,err:=json.MarshalIndent(data," ","")
  if err != nil{
    return nil
  }
  return os.WriteFile((*s).Filename,funcdata,0644)
}

//checker file
func main(){
  fmt.Println("hrllo")
}