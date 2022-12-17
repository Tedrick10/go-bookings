// Main: Package
package main

// Go: Libraries
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Tedrick10/go-bookings/pkg/config"
	"github.com/Tedrick10/go-bookings/pkg/handlers"
	"github.com/Tedrick10/go-bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

/* Variables */
const portNumber string = ":8080"; // The port number to be used as localhost.

// Declarations
var appConfig config.AppConfig;
var session *scs.SessionManager;

/* Index Function */
// Main is the index function of the application
func main()  {
	

	// Change this to true when in production
	appConfig.InProduction = false;

	session = scs.New();
	session.Lifetime = 24 * time.Hour;
	session.Cookie.Persist = true;
	session.Cookie.SameSite = http.SameSiteLaxMode;
	session.Cookie.Secure = appConfig.InProduction;
	appConfig.Session = session;

	tc, err := render.CreateTemplateCache();
	if err != nil {
		log.Fatal("Cannot create template cache");
	}

	appConfig.TemplateCache = tc;
	appConfig.UseCache = false;
	repo := handlers.NewRepo(&appConfig);
	handlers.NewHandler(repo);
	render.NewTemplates(&appConfig);

	// Actions
	// http.HandleFunc("/", handlers.Repo.Home);
	// http.HandleFunc("/about", handlers.Repo.About);

	// Printings
	fmt.Println(fmt.Sprintf("Starting application on port %s.", portNumber));

	// Listener
	// _ = http.ListenAndServe(portNumber, nil);
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe();
	log.Fatal(err);
}