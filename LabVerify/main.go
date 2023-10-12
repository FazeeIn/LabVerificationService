package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	/*err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}*/

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.LoadHTMLGlob("src/**/*")
	//r.Static("/static", "./static")
	//r.StaticFS("/src/*", http.Dir("src"))
	a := controller.NewApp()
	a.Routes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
