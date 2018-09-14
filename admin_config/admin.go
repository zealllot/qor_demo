package admin_config

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

// Define a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Define another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}

func Admin(db *gorm.DB, mux *http.ServeMux) {
	db.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: db})

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// Mount admin to the mux
	Admin.MountTo("/admin", mux)
}
