package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kataras/iris"
)

type Car struct {
	ID    int64  `json:"id"`
	Model string `json:"model"`
	Year  int64  `json:"year"`
	Make  string `json:"make"`
}

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Cars      []Car  `json:"cars"`
}

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html").Reload(true))
	app.StaticWeb("/", "./static")

	// Render home.html
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Title", "Home Page")
		ctx.ViewData("Name", "oscar")
		ctx.Gzip(true)

		ctx.View("home.html")
	})

	// REST API Endpoints
	app.Handle("GET", "/api/person", func(ctx iris.Context) {

		// call the Cars API
		response, err := http.Get("https://cars-rest-api.herokuapp.com/api/person")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var people []Person

		json.Unmarshal(responseData, &people)
		fmt.Printf("%#v", people)
		fmt.Println(len(people))

		for i := 0; i < len(people); i++ {
			fmt.Printf("%v owns %v  car(s)\n\n", people[i].FirstName, len(people[i].Cars))

			// iterate throug owner cars
			for c := range people[i].Cars {
				fmt.Printf("Make: '%v' Model: '%s'\n", people[i].Cars[c].Make, people[i].Cars[c].Model)
			}
		}

	})

	app.Configure(iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
	app.Run(iris.Addr(":8080"))

}
