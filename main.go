package main

import "fmt"
import "github.com/gofiber/fiber"
import "os"
import "log"
import "net/http"
import "github.com/PuerkitoBio/goquery"
import "github.com/gofiber/template/html"

func errorHandling(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	engine := html.New(".", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Get("/home", func(c *fiber.Ctx) {
		_ = c.Render("buscador", fiber.Map{})
		if c.Fasthttp.Request.Header.Method == "POST" {
			coso := c.FormValue("coso")
		}
	})

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

	// Create a goquery document from the HTTP response
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
