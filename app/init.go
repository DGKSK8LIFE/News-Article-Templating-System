package app

import (
	"fmt"
	"io/ioutil"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
	"gopkg.in/yaml.v2"
)

type DB_user struct {
	DB_user string `yaml:"DB_user"`
}

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	DB *sql.DB

	err error

	user DB_user
)

func InitDB() {
	yamlFile, err := ioutil.ReadFile("app/DB_user.yaml")
	if err != nil {
		log.Fatalf("yamlFile err: %s\n", err)
	}
	err = yaml.Unmarshal(yamlFile, &user)
	if err != nil {
		log.Fatalf("Unmarshal err: %s\n", err)
	}
	openString := fmt.Sprintf("%s@tcp(localhost:3306)/articles", user.DB_user)
	DB, err = sql.Open("mysql", openString)
	if err != nil {
		log.Fatalf("DB Error: %s\n", err)
	}
	log.Println("DB Connected")
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}
	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
