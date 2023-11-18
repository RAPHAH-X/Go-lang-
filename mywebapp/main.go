package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a route and a handler function
	http.HandleFunc("/", homeHandler)

	// Start the web server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve a simple HTML page
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>My Go Web App</title>
    </head>
		<h1>Am creating my first website</h1>
    <body>
        <h1>Hello, this is my Go web app!</h1>
    </body>
		<h2> please add some imagies for verification</h2>
    </html>
    `

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
