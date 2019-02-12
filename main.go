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
	Year  int    `json:"year"`
	Make  string `json:"genre"`
}

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	email     string `json:"email"`
}

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html").Reload(true))
	app.StaticWeb("/", "./static")

	// Database connection
	// session, err := mgo.Dial("localhost")
	// if nil != err {
	// 	panic(err)
	// }
	// defer session.Close()
	// session.SetMode(mgo.Monotonic, true)

	// // Database name and collection name
	// // car-db is database name car is colletion name
	// c := session.DB("car-db").C("car")
	// c.Insert(&Car{"Audi", "Luxury car"})

	// Render home.html
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Title", "Home Page")
		ctx.ViewData("Name", "iris") // {{.Name}} will render: iris
		ctx.Gzip(true)

		ctx.View("home.html")
	})

	// REST API Endpoints
	app.Handle("GET", "/api/person", func(ctx iris.Context) {

		response, err := http.Get("https://cars-rest-api.herokuapp.com/api/person")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject []Person
		json.Unmarshal(responseData, &responseObject)

		fmt.Println(responseObject)
		fmt.Println(len(responseObject))

		for i := 0; i < len(responseObject); i++ {
			fmt.Println(responseObject[i].FirstName)
		}

		// person.FirstName = params["id"]
		// people = append(people, person)
		// json.NewEncoder(w).Encode(people)
		/* 		response, err := http.Get("https://cars-rest-api.herokuapp.com/api/person")
		   		if err != nil {
		   			fmt.Print(err.Error())
		   			os.Exit(1)
		   		}

		   		responseData, err := ioutil.ReadAll(response.Body)
		   		if err != nil {
		   			log.Fatal(err)
		   		}

		   		var responseObject Person
		   		json.Unmarshal(responseData, &responseObject)

		   		fmt.Println(responseObject.FirstName) */

	})

	// Get /car from mongodb
	// app.Get("/car", func(ctx iris.Context) {
	// 	result := Car{}
	// 	err = c.Find(bson.M{"name": "Audi"}).One(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ctx.JSON(result)
	// })

	app.Configure(iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
	app.Run(iris.Addr(":8080"))

}
