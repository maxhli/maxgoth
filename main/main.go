package main

import (
	"os"
	"net/http"


	"github.com/qor/auth"
	_ "github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/session/manager"
	"github.com/jinzhu/gorm"
)

var (
// Initialize gorm DB
gormDB, _ = gorm.Open("sqlite3", "sample.db")

// Initialize Auth with configuration
Auth = auth.New(&auth.Config{
DB: gormDB,
})
)

func init() {
// Migrate AuthIdentity model, AuthIdentity will be used to save auth info, like username/password, oauth token, you could change that.
// gormDB.AutoMigrate(&auth_identity.AuthIdentity{})

// Register Auth providers
// Allow use username/password
Auth.RegisterProvider(password.New(&password.Config{}))

// Allow use Github
Auth.RegisterProvider(github.New(&github.Config{
ClientID:     "github client id",
ClientSecret: "github client secret",
}))

// Allow use Google
Auth.RegisterProvider(google.New(&google.Config{
ClientID:
	"783416833376-8mjrbd4qob8jujp54nr3huevatvlko5o.apps.googleusercontent.com",
ClientSecret:
	"yDl4aWdM58O0bxqTUMYsVDWw",
}))

// Allow use Facebook
Auth.RegisterProvider(facebook.New(&facebook.Config{
ClientID:     "facebook client id",
ClientSecret: "facebook client secret",
}))

// Allow use Twitter
Auth.RegisterProvider(twitter.New(&twitter.Config{
ClientID:     "twitter client id",
ClientSecret: "twitter client secret",
}))
}

func main() {
mux := http.NewServeMux()

// Mount Auth to Router
mux.Handle("/auth/", Auth.NewServeMux())
port := os.Getenv("PORT")
if port == "" {
	port = ":9000"
} else {
	port = ":" + port
}
http.ListenAndServe(port, manager.SessionManager.Middleware(mux))
}