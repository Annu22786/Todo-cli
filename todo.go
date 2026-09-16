package main

import (
  "fmt"
  "time"
  "errors"
)
type Todo struct{
  Title string
  completed bool
  createdat time.Time
  completedat *time.Time  
}

//because we want to update and delete and we will use slice 
type Todos []Todo

// func for adding new task 

func (todos *Todos) add(title string){
  todo :=Todo{
    Title : title,
    completed : false,
    createdat : time.Now(),
    completedat : nil,
  }
  *todos = append (*todos,todo)
}
// func to check whether index is correct or not
func (todos *Todos) validateindex(index int) error{
  if index<0 || index>=len(*todos){
    err:= errors.New("invalid input")
    fmt.Println(err)
    return err
  }
  return nil
}
// delete task from todo
func(todos *Todos) del(index int)error{
  if err:=todos.validateindex(index);err!=nil{
    return err
  }
  *todos = append((*todos)[:index],(*todos)[index + 1 :]...)
  return nil
}
func main(){
  fmt.Println("done")
}