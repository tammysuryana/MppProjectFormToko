package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
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

 //   db.AutoMigrate(&user.User{})
  //  db.Set( "gorm:table_options" , "ENGINE=InnoDB" ).AutoMigrate(&user.User{})

    //router := gin.Default()
    //router.GET("/handler")

    userRepository := user.NewRepository(db)
    userService := user.NewService(userRepository)
  //  userService.SaveAvatar(1, "images/1-profile.png")


    userHandler := handler.NewUserHandler(userService)


    // TEST DARI SERVICE.GO
    //input := user.LoginInput{
    //    Email: "kjhkjh@domain.com",
    //    Password: "$2a$04$StqRxVQGKxpVZpJo7iSLu.KUoJ299fIJkDtWUB3UDre78rS78ZIuq",
    //
    //}
    //user, err := userService.Login(input)
    //if err != nil{
    //    fmt.Println("terjadi kesalahan ")
    //    fmt.Println(err.Error())
    //}
    //fmt.Println(user.Email)
    //fmt.Println(user.Name)


         //   CODE INPUT HARDCORE
    // userByEmail, err := userRepository.FindById("kjhkjh@domain.com")
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
    //if (userByEmail.ID == 0){
    //    fmt.Println("teu aya user na eung ")
    //}else{
    //    fmt.Println("hallo ",userByEmail.Name)
    //}



    router := gin.Default()
    api := router.Group("/api/v1")
    api.POST("/users",userHandler.RegisterUser)
    api.POST("/session",userHandler.Login)
    api.POST("/email_checkers",userHandler.CheckEmaiAvailability)
    api.POST("/avatars",userHandler.UploadAvatar)

   //handler.NewUserHandler(userService)
 //  userHandler := handler.NewUserHandler(userService)
   //api := router.Group("/api/v1")
  // api.POST("/users", userHandler.RegisterUser)


   router.Run()
}
