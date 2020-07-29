package main

import "fmt"
import "github.com/gofiber/fiber"
import "os"
import "log"
import "net/http"
import "github.com/PuerkitoBio/goquery"
import "github.com/gofiber/template/html"
import "github.com/gofiber/fiber/middleware"

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ShowBuscadorForm(c *fiber.Ctx) {
	if err := c.Render("buscador",
		fiber.Map{}); err != nil {
		c.Status(500).Send(err.Error())
	}
}

func PostBuscadorForm(c *fiber.Ctx) {
	query := c.FormValue("coso")
	c.Send(query)
	fmt.Println("Printing query:" + query)
}

func main() {

	engine := html.New(".", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Use(middleware.Logger(os.Stdout))
	app.Use(middleware.Logger("UTC"))
	app.Use(middleware.Logger("15:04:05"))
	app.Use(middleware.Logger("${time} ${method} ${path}"))
	app.Get("/buscador", ShowBuscadorForm)
	app.Post("/buscador", PostBuscadorForm)

	query := "tarjeta de video"
	mercadoLibreUrl := "https://listado.mercadolibre.com.mx/"
	input := "#D[A:" + query + "]"
	busqueda := mercadoLibreUrl + query + input
	response, error := http.Get(busqueda)
	fmt.Println(busqueda)
	errorHandling(error)
	fmt.Println(response)
	fmt.Println(response.Body)
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			fmt.Println(imgSrc)
		}
	})
	app.Listen(3000)
}
