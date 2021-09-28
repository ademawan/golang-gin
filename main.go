package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()

	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/hello", hello)
	// r.GET("/hello/:name/:age", helloWithParams)
	r.GET("/hello/:name", helloWithParams)
	r.GET("/hello/:name/:age", helloWithParams)

	r.GET("/list", listBook)
	r.GET("/login", login)

	r.Run()
}

type Login struct {
	Username string `binding:"required" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func login(c *gin.Context) {

	// username := c.PostForm("username")

	login := Login{}
	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
	}

	if login.Username == "root" && login.Password == "123" {
		c.JSON(200, gin.H{"code": 200, "message": "ok"})
	}

	// fmt.Println(login)
}

type Book struct {
	Name string
	Page int
}

func listBook(c *gin.Context) {
	var books []Book
	book1 := Book{Name: "bobo", Page: 10}
	book2 := Book{Name: "story", Page: 20}
	books = append(books, book1)
	books = append(books, book2)

	c.JSON(200, gin.H{"code": 200, "message": "Ok", "data": books})
	// page := c.DefaultQuery("page", "1")
	// order := c.DefaultQuery("order", "ASC")
	// fmt.Println("page " + page + " order " + order)

	// fmt.Println(books[0])
}

func hello(c *gin.Context) {
	//http://localhost/hello?firstname=ademawan
	Fname := c.Query("firstname")
	fmt.Println(Fname)

	c.String(201, "hello")
}
func helloWithParams(c *gin.Context) {
	name := c.Param("name")

	age := c.Param("age")
	fmt.Println(age, name)
	c.String(201, "hello "+name+age)
}
