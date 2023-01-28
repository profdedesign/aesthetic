package main

import (
	"os"
	"log"
	"fmt"
	"flag"
	"errors"
	"crypto/tls"

	"github.com/gofiber/fiber/v2"

	"server/handlers"
)

func main() {
	// Starts a Fiber instance
	app := fiber.New()

	// Allows to source every static file (or asset) within the web directory
	app.Static("/css", "../web/css")
	app.Static("/fonts", "../web/fonts")
	app.Static("/images", "../web/images")
	app.Static("/js", "../web/js")
	app.Static("/svg", "../web/svg")

	// Renders index.html to the main website
	app.Get("/", handlers.RootPage)

	// Allows to handle 404 errors; adapted from [1]
	app.Use(handlers.Error404Handler)


	// If the application is running on a server, use available certificates to enable HTTPS; adapted from [2] (lifesaver!) & [3]
	flgServer := flag.String("server", "false", "Is the application running inside a production server?")

	if *flgServer == "true" {
		// Allows to define the path of both the key and the certificate chain otherwise
		flgCert := flag.String("chain", "certs/chain.pem", "Path to the (full) certificate chain") // Avoids delivering only the root certificate; see [4]
		flgKey := flag.String("key", "certs/key.pem", "Path to key")
		flag.Parse()

		fmt.Println(*flgCert, *flgKey)

		// Checks if files exist on the defined paths; adapted from [5]
		if _, err := os.Stat(*flgCert); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Certificate doesn't exist at \"" + *flgCert + "\"")
		}

		if _, err := os.Stat(*flgKey); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Key doesn't exist at \"" + *flgKey + "\"")
		}

		// Loads the key-certificate pair to the TLS configuration
		cert, err := tls.LoadX509KeyPair(*flgCert, *flgKey)
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig := &tls.Config{
			Certificates:             []tls.Certificate{cert},
			// ALlows to publicly expose Go safely by keeping TLS up-to-date; see [2]
			CipherSuites:             nil,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS13,
			CurvePreferences: []tls.CurveID{
				// As per Mozilla recommendation [4], explicitly defines ECC cryptography, which allows for smaller key sizes [6]
				tls.CurveP256,
				tls.X25519,
			},
			NextProtos: []string{
				// Due to fasthttp limitation to HTTP v1 only, explicitly defines only to use v1; see [3]
				"http/1.1", "acme-tls/1",
			},
		}

		// Runs the server under the HTTPS port
		ln, err := tls.Listen("tcp", ":443", tlsConfig)
		if err != nil {
			panic(err)
		}

		log.Fatal(app.Listener(ln))
	} else {
		// Otherwise, run as if locally, i.e. under the user port 8000
		log.Fatal(app.Listen(":8000"))
	}
}
/* Sources:
	* [1]: https://github.com/gofiber/recipes/blob/master/404-handler/main.go
	* [2]: https://bruinsslot.jp/post/go-secure-webserver/
	* [3]: https://github.com/gofiber/recipes/blob/master/autocert/main.go (didn't work per se, even under a production server, see post.)
		-> https://github.com/gofiber/recipes/issues/168#issuecomment-1108421521
	* [4]: https://stackoverflow.com/a/40914979
	* [5]: https://stackoverflow.com/a/12518877
	* [6]: https://datatracker.ietf.org/doc/html/rfc4492
 */
