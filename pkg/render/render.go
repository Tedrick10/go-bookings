// Render: Package
package render

// Go: Libraries
import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Tedrick10/go-bookings/pkg/config"
	"github.com/Tedrick10/go-bookings/pkg/models"
)

var app *config.AppConfig;

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig)  {
	app = a;
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td;
}
/* Functions */
// renderTemplate: Function to render HTML template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData)  {
	// create a template cache
	// tc, err := CreateTemplateCache();
	// if err != nil {
	// 	log.Fatal(err);
	// }
	var tc map[string]*template.Template;

	//get requested template from cache
	if app.UseCache {
		tc = app.TemplateCache;
	} else {
		tc, _ = CreateTemplateCache();
	}
	
	t, ok := tc[tmpl];
	if !ok {
		log.Fatal("Could not get template from tempate cache!");
	}
	
	buf := new(bytes.Buffer);
	td = AddDefaultData(td);
	err := t.Execute(buf, td);
	if err != nil {
		log.Fatal(err);
	}

	// render the template
	_, err = buf.WriteTo(w);
	if err != nil {
		log.Println("Errir writing template to browser", err);
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl");
	// err := parsedTemplate.Execute(w, nil);

	// if err != nil {
	// 	fmt.Println("Error parsing template: ", err);
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// get the template cache from the app config

	myCache := map[string]*template.Template{};

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl");

	// Conditions
	if err != nil {
		return myCache, err;
	}

	// Range through all files ending with *.page.tmpl
	for _, page := range pages {
		// Name stored the name of the specific file
		name := filepath.Base(page);

		ts, err := template.New(name).ParseFiles(page);
		if err != nil {
			return myCache, err;
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl");
		if err != nil {
			return myCache, err;
		}
		
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl");
			if err != nil {
				return myCache, err;
			}
		}

		myCache[name] = ts;
	}

	return myCache, nil;
}