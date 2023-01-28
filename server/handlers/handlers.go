package handlers

import (
	"html/template"

	"github.com/gofiber/fiber/v2"

	"server/utils"
)

/* Renders the main website */
type Map struct {
	A template.HTMLAttr // Allows to safely apply the attribute, avoiding issues; see [1] & [2]
}

func RootPage(c *fiber.Ctx) error {
	m := Map{
		A: `target="__blank"`, // Allows to redirect a page in a new window/tab [3]
	}

	return utils.Handlerizer(c, m, "../web/html/index.html")
}

/* Renders 4xx error websites */
type Map4xx struct {
	Error   string
	Apology string
	Enquiry string
	Button  string
	HREF    string
}

func Error404Handler(c *fiber.Ctx) error {
	m := Map4xx{
		Error:   "HTTP Error 404",
		Apology: "Shoot. Seems I can't find whatever you're looking for.",
		Enquiry: "However, you can get back to the main page:",
		Button:  "Go Back",
		HREF:    "/",
	}
	c.Status(404)
	return utils.Handlerizer(c, m, "../web/html/4xx.html")
}

/* References:
	* [1]: https://stackoverflow.com/a/14796642
	* [2]: https://stackoverflow.com/a/58276921
	* [3]: https://www.w3schools.com/html/html_links.asp
*/
