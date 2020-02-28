package main

import (
  "net/http"
  "fmt"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "github.com/qorpress/redirect_back"
  "github.com/qorpress/auth"
  "github.com/qorpress/auth/auth_identity"
  "github.com/qorpress/auth/providers/github"
  "github.com/qorpress/auth/providers/google"
  "github.com/qorpress/auth/providers/password"
  "github.com/qorpress/auth/providers/facebook"
  "github.com/qorpress/auth/providers/twitter"
  "github.com/qorpress/session/manager"
  _ "github.com/mattn/go-sqlite3"
  "github.com/qorpress/auth_themes/clean"

  "github.com/qorpress/qorpress-auth-example/config"
  "github.com/qorpress/qorpress-auth-example/models"
)

var (
  // Initialize gorm DB
  gormDB, _ = gorm.Open("sqlite3", "qorpress.db")
  // Initialize Auth with configuration
  /*
  Auth = auth.New(&auth.Config{
    DB: gormDB,
    Redirector: auth.Redirector{RedirectBack},
  })
  */

  Auth = clean.New(&auth.Config{
    DB:         gormDB,
    // NO NEED TO CONFIG RENDER, AS IT'S CONFIGED IN CLEAN THEME
    // Render:     render.New(&render.Config{AssetFileSystem: bindatafs.AssetFS.NameSpace("auth")}),
    Mailer:     config.Mailer,
    UserModel:  models.User{},
    Redirector: auth.Redirector{RedirectBack},
  })
)

var RedirectBack = redirect_back.New(&redirect_back.Config{
  SessionManager:  manager.SessionManager,
  IgnoredPrefixes: []string{"/auth"},
})

func init() {
  // Migrate AuthIdentity model, AuthIdentity will be used to save auth info, like username/password, oauth token, you could change that.
  // gormDB.AutoMigrate(&auth_identity.AuthIdentity{})

  gormDB.LogMode(true)
  
  gormDB.AutoMigrate(
    &models.User{},
    &auth_identity.AuthIdentity{},
  )


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
    ClientID:     "google client id",
    ClientSecret: "google client secret",
    AllowedDomains: []string{}, // Accept all domains, instead you can pass a whitelist of acceptable domains
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

  router := gin.Default()
  router.Any("/*resources", gin.WrapH(mux))

  router.Run(fmt.Sprintf("%s:%s", "127.0.0.1", "9000"))

  // http.ListenAndServe(":9000", manager.SessionManager.Middleware(RedirectBack.Middleware(mux)))
  // http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
}
