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

//Car Model
type Car struct {
	ID    int64  `json:"id"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Make  string `json:"make"`
}

//Person Model
type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Cars      []Car  `json:"cars"`
}

func main() {

	var CAR_API_BASE_URL = "https://cars-rest-api.herokuapp.com/api"

	app := iris.Default()
	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html").Reload(true))
	app.StaticWeb("/", "./static")

	// Render Home View
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Title", "Home Page")
		ctx.ViewData("Name", "oscar")
		ctx.Gzip(true)

		ctx.View("home.html")
	})

	// REST API Endpoints, group routes by endpoint
	api := app.Party("/api")

	// showing how you can also do an inline function vs a handler.
	api.Get("/help", func(ctx iris.Context) {
		ctx.Writef("GET / -- fetch all people\n")
		ctx.Writef("GET /$id -- fetch a person by id\n")
		ctx.Writef("POST / -- create new person\n")
		ctx.Writef("PUT /$id -- update an existing person\n")
		ctx.Writef("DELETE /$id -- delete an existing person\n")
	})

	// http://localhost:8080/api/person - get all people
	api.Get("/person", func(ctx iris.Context) {

		response, err := http.Get(fmt.Sprintf("%s/person", CAR_API_BASE_URL))

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// create an array of Person Objects
		// each person also owns multiple cars
		// so we iterate through those
		var people []Person

		json.Unmarshal(responseData, &people)
		fmt.Printf("%#v", people)
		fmt.Println(len(people))

		for i := 0; i < len(people); i++ {
			fmt.Printf("%v owns %v  car(s)\n\n", people[i].FirstName, len(people[i].Cars))

			// iterate through owner cars
			for c := range people[i].Cars {
				fmt.Printf("Make: '%v' Model: '%s'\n", people[i].Cars[c].Make, people[i].Cars[c].Model)
			}
		}

		ctx.StatusCode(200)
		ctx.JSON(people)
	})

	// http://localhost:8080/api/person/42 - get a specific person by Id
	api.Get("/person/{id:string}", func(ctx iris.Context) {
		println(ctx.Path())

		// the person id being passed as parameter
		id := ctx.Params().GetString("id")
		ctx.Writef("get user by id: %s", id)

		response, err := http.Get(fmt.Sprintf("%s/person/%s", CAR_API_BASE_URL, id))

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var person Person

		json.Unmarshal(responseData, &person)
		//fmt.Printf("%#v", person)

		ctx.StatusCode(200)
		ctx.JSON(person)

	})

	// http://localhost:8080/person/cars - get all cars person owns.
	api.Get("/person/{id:int}/cars", func(ctx iris.Context) {})

	app.Configure(iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
	app.Run(iris.Addr(":8080"))

}
