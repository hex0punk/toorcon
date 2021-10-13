package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hex0punk/toorcon/cachicamo/helpers"
	"github.com/hex0punk/toorcon/cachicamo/visitors"
	"log"
	"mime/multipart"
	"net/http"
	"runtime"
	"time"
)

type PhrasePayload struct {
	Phrase   string `json:"phrase"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/upload", uploadFile)
		api.POST("/testPhrase", testPhrase)
		api.GET("/logUsage", logUsage)
		api.GET("/addVisitor", addVisitorCount)
		api.GET("/subtractVisitor", subtractVisitorCount)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8081")
}

func testPhrase(c *gin.Context) {
	var p PhrasePayload
	if err := c.ShouldBindJSON(&p); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := helpers.ParseSecretPass(p.Phrase)
	if err != nil {
		c.JSON(403, gin.H{
			"msg": "Invalid phrase",
		})
		c.Abort()
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("YOU GOT IT!"))
}

func uploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")

	ch := make(chan error)
	go func(ctx *gin.Context, fileCpy *multipart.FileHeader) {
		err := saveFile(ctx, fileCpy)
		ch <- err
		saveFile(ctx, fileCpy)
	}(c, file)

	select {
	case err := <-ch:
		if err != nil {
			fmt.Sprintf("'%s' uploaded!", file.Filename)
		}
		c.JSON(500, gin.H{
			"msg": "Unable to save file",
		})
		c.Abort()
	case <-time.After(time.Second * 3):
		c.JSON(500, gin.H{
			"msg": "TIMEOUT!",
		})
		c.Abort()
	}
	fmt.Printf("GOROUTINES: %d\n", runtime.NumGoroutine())
	helpers.PrintMemUsage()
}

func addVisitorCount(c *gin.Context){
	v := visitors.New()
	v.Add()
	c.String(http.StatusOK, fmt.Sprintf("Visitor Count: %d\n", v.GetCount()))
}

func subtractVisitorCount(c *gin.Context){
	v := visitors.New()
	err := v.Subtract()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("Visitor Count: %d\n", v.GetCount()))
}

func logUsage(c *gin.Context) {
	helpers.PrintMemUsage()
	c.String(http.StatusOK, fmt.Sprintf("logged"))
}

func saveFile(c *gin.Context, file *multipart.FileHeader) error {
	log.Println(file.Filename)

	time.Sleep(time.Second * 10)

	err := c.SaveUploadedFile(file, "./Uploads/" + file.Filename)

	if err != nil {
		return err
	}

	return nil
}
