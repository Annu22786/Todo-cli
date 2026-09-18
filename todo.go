package main

import (
  "fmt"
  "time"
  "errors"
  "os"
  "github.com/aquasecurity/table"
  "strconv"
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
  t := *todos
  if err:=t.validateindex(index);err!=nil{
    return err
  }
  t = append(t[:index],t[index + 1 :]...)
  return nil
}
//edit title with index
func (todos *Todos) ed(index int,title string)error{
  t := *todos
  if err:=t.validateindex(index);err!=nil{
    return err
  }
  t[index].Title=title
  return nil
  }
//func for toggle(trye/false) by index number
func (todos *Todos) tog(index int) error{
  t := *todos
  if err:=t.validateindex(index);err!=nil{
    return err
  }
  check := t[index].completed
  if check == false{
    t[index].completed = true
    //completed time logic
    com := time.Now() 
    t[index].completedat = &com
  }else {
    t[index].completed = false
    t[index].completedat = nil
}
  
  return nil
}

//checker

// func for printing all todos in cli

func (todos *Todos) print(){
  t:= *todos
  table := table.New(os.Stdout)
  table.SetHeaders("#","Title","Completed","Createdat","Completedat")
  table.SetRowLines(false)
  for index,value := range t {
    completed := ""
    completedat := ""
    if value.completed == true {
      completed = "☆"
      if value.completedat != nil {
        completedat = value.completedat.Format(time.RFC1123)
      }
    }
    table.AddRow(strconv.Itoa(index),value.Title,completed,value.createdat.Format(time.RFC1123),completedat)
 }
table.Render()
}

//checker
/*
func main(){
  fmt.Println("done")
}

*/