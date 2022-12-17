// Handlers: Package
package handlers

// Go: Libraries
import (
	"net/http"

	"github.com/Tedrick10/go-bookings/pkg/config"
	"github.com/Tedrick10/go-bookings/pkg/models"
	"github.com/Tedrick10/go-bookings/pkg/render"
)

// Repo: the repository used by the handlers
var Repo *Repository;

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repostiory
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers set the repository for the handlers
func NewHandler(r *Repository)  {
	Repo = r;
}

/* HTTP Functions */
// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	remoteIP := r.RemoteAddr;
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP);
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{});
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request)  {
	// Perfrom some logic
	stringMap := make(map[string]string);
	stringMap["test"] = "Hello, Again!";
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip");
	stringMap["remote_ip"] = remoteIP;

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	});
}
