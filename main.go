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

  "github.com/qorpress/qorpress-auth-example/pkg/config"
  "github.com/qorpress/qorpress-auth-example/pkg/models"
)

var (
  // Initialize gorm DB
  gormDB, _ = gorm.Open("sqlite3", "qorpress.db")

  // Initialize Auth with configuration
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

  gormDB.LogMode(config.Config.DB.Debug)
  
  gormDB.AutoMigrate(
    &models.User{},
    &auth_identity.AuthIdentity{},
  )


  // Register Auth providers
  // Allow use username/password
  Auth.RegisterProvider(password.New(&password.Config{}))

  // Allow use Github
  Auth.RegisterProvider(github.New(&github.Config{
    ClientID:     config.Config.Auth.Github.ClientID,
    ClientSecret: config.Config.Auth.Github.ClientSecret,
  }))

  // Allow use Google
  Auth.RegisterProvider(google.New(&google.Config{
    ClientID:     config.Config.Auth.Google.ClientID,
    ClientSecret: config.Config.Auth.Google.ClientID,
    AllowedDomains: []string{}, // Accept all domains, instead you can pass a whitelist of acceptable domains
  }))

  // Allow use Facebook
  Auth.RegisterProvider(facebook.New(&facebook.Config{
    ClientID:     config.Config.Auth.Facebook.ClientID,
    ClientSecret: config.Config.Auth.Facebook.ClientID,
  }))

  // Allow use Twitter
  Auth.RegisterProvider(twitter.New(&twitter.Config{
    ClientID:     config.Config.Auth.Twitter.ClientID,
    ClientSecret: config.Config.Auth.Twitter.ClientID,
  }))

}

func main() {
  mux := http.NewServeMux()

  // Mount Auth to Router
  mux.Handle("/auth/", Auth.NewServeMux())

  router := gin.Default()

  if !config.Config.App.Debug {
    gin.SetMode(gin.ReleaseMode)
  }

  router.Any("/*resources", gin.WrapH(mux))

  router.Run(fmt.Sprintf("%s:%d", "", config.Config.App.Port))

  // http.ListenAndServe(":9000", manager.SessionManager.Middleware(RedirectBack.Middleware(mux)))
  // http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
}
