package utils

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
)

/* Parses an HTML file as a template from scratch (i.e. without Fiber's HTML template package); adapted from [1] */
func Handlerizer(c *fiber.Ctx, s interface{}, f string) error {
	p, _ := template.ParseFiles(f)

	// Assures the content type of the response is HTML
	c.Type("html")

	return p.Execute(c, s) // Imports the "request context" directly as the Writer argument; see [2]
}

/* References:
	* [1]: https://youtube.com/watch?v=Uo9MSE2Gbzs
	* [2]: https://stackoverflow.com/questions/48234198/how-to-use-http-responsewriter-in-fasthttp-for-executing-html-templates-etc
*/
