package auth

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
	"gopkg.in/gomail.v2"
)

func Auth(db *gorm.DB, mux *http.ServeMux) {
	// Migrate AuthIdentity model, AuthIdentity will be used to save auth info, like username/password, oauth token, you could change that.
	db.AutoMigrate(&auth_identity.AuthIdentity{}, &auth_identity.Basic{})

	// Config gomail
	dailer := gomail.NewDialer()
	sender, err := dailer.Dial()
	if err != nil {
		panic(err)
	}

	// Initialize Mailer
	mailer := mailer.New(&mailer.Config{
		Sender: gomailer.New(&gomailer.Config{Sender: sender}),
	})

	// Initialize Auth with configuration
	Auth := clean.New(&auth.Config{
		DB:     db,
		Mailer: mailer,
	})

	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{}))

	//// Allow use Github
	//Auth.RegisterProvider(github.New(&github.Config{
	//	ClientID:     "github client id",
	//	ClientSecret: "github client secret",
	//}))
	//
	//// Allow use Google
	//Auth.RegisterProvider(google.New(&google.Config{
	//	ClientID:     "google client id",
	//	ClientSecret: "google client secret",
	//}))
	//
	//// Allow use Facebook
	//Auth.RegisterProvider(facebook.New(&facebook.Config{
	//	ClientID:     "facebook client id",
	//	ClientSecret: "facebook client secret",
	//}))
	//
	//// Allow use Twitter
	//Auth.RegisterProvider(twitter.New(&twitter.Config{
	//	ClientID:     "twitter client id",
	//	ClientSecret: "twitter client secret",
	//}))

	// Mount Auth to Router
	mux.Handle("/auth/", Auth.NewServeMux())
}
