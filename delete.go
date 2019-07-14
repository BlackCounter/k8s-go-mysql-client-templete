package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
  gorm.Model
  Name string
}

func main() {
  db, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/test?charset=utf8&parseTime=True")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  var user User
  db.First(&user, "name = ?", "jotaro")
  if db.Error != nil {
    panic(db.Error)
  }
 
  db.Delete(&user)
  if db.Error != nil {
    panic(db.Error)
  }
}
