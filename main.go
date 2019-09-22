package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// this handles the main page
func homePageHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "text/html") // set the content header type

	// set the html content [whih normally would be generated automatically] in a conditional way
	htmlContent := `<!DOCTYPE html>
	<html>
		<head>
			<title>Home Page</title>
		</head>
		<body>
			This is a website server by a Go HTTP Server!
		</body>
	</html> `

	io.WriteString(response, htmlContent)
}

// this handles the contact page
func contactPageHandler(response http.ResponseWriter, request *http.Request) {

	// set the content header type
	response.Header().Set("Content-Type", "text/html")

	// page content
	htmlContent := `<!DOCTYPE html>
	<html>
		<head>
			<title>Contact Page</title>
		</head>
		<body>
			To get in touch, please send an email to <a href=\"mailto:damilarelana@gmail.com\">support@gallery.com</a>
		</body>
	</html> `

	io.WriteString(response, htmlContent)
}

// this handles FAQ page
func faqPageHandler(response http.ResponseWriter, request *http.Request) {

	// set the content header type
	response.Header().Set("Content-Type", "text/html")

	// page content
	htmlContent := `<!DOCTYPE html>
	<html>
		<head>
			<title>FAQ Page</title>
		</head>
		<body>
			Frequently Asked Questions .... 
		</body>
	</html> `

	io.WriteString(response, htmlContent)
}

// this handles custome 404 page
func custom404PageHandler(response http.ResponseWriter, request *http.Request) {

	// set the content header type
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusNotFound) // this automatically generates a 404 status code
	// page content
	htmlContent := `<!DOCTYPE html>
	<html>
		<head>
			<title>404 Page</title>
		</head>
		<body>
		 	These are not Gophers you are looking for ... 404 Page!
		</body>
	</html> `

	io.WriteString(response, htmlContent)
}

func main() {
	muxRouter := mux.NewRouter()                                       // instantiate the gorillamux Router
	muxRouter.NotFoundHandler = http.HandlerFunc(custom404PageHandler) // customer 404 Page handler scenario
	muxRouter.HandleFunc("/", homePageHandler)                         // declaring path and the handler function
	muxRouter.HandleFunc("/contact", contactPageHandler)
	muxRouter.HandleFunc("/faq", faqPageHandler)
	http.ListenAndServe(":9009", muxRouter) // set the port where the http server listens and serves. changed `nil` to the instance muxRouter
}
