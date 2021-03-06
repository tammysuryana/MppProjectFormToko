package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/tammysuryana93/auth"
    _"github.com/tammysuryana93/auth"
    "github.com/tammysuryana93/handler"
    "github.com/tammysuryana93/user"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

)

func main() {
    // connection to databases
    dsn := "root:1993@tcp(127.0.0.1:3306)/mppproject?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    fmt.Println("success to connect databases")

    //  db.AutoMigrate(&user.User{})
    //  db.Set( "gorm:table_options" , "ENGINE=InnoDB" ).AutoMigrate(&user.User{})

    //router := gin.Default()
    //router.GET("/handler")

    userRepository := user.NewRepository(db)
    userService := user.NewService(userRepository)
    authService := auth.NewJwtService()


    userHandler := handler.NewUserHandler(userService, authService)


    router := gin.Default()
    api := router.Group("/api/v1")
    api.POST("/users",userHandler.RegisterUser)
    api.POST("/session",userHandler.Login)
    api.POST("/email_checkers",userHandler.CheckEmaiAvailability)
    api.POST("/avatars",userHandler.UploadAvatar)

   router.Run()
}
