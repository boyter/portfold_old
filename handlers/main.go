package handlers

import (
	"boyter/portfold/data/mysql"
	"html/template"
	"log"
	"net/http"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type Application struct {
	ErrorLog     *log.Logger
	InfoLog      *log.Logger
	ProjectModel *mysql.ProjectModel
}

func (app *Application) Routes() *http.ServeMux {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	// It is good practice to create a new one to avoid the default global one
	// being polluted by imports
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(app.Home))
	mux.Handle("/help/", http.HandlerFunc(app.Help))
	mux.Handle("/health-check/", http.HandlerFunc(app.HealthCheck))

	// Setup to serve files from the supplied directory
	fileServer := http.FileServer(http.Dir("./assets/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

// Define a home handler function which writes a byte slice containing
// "Hello from Portfold" as the response body.
func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// TODO should cache these templates and better yet build them into the application

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./assets/ui/html/home.page.tmpl",
		"./assets/ui/html/base.layout.tmpl",
		"./assets/ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *Application) Help(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Help"))
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("health check"))
}
