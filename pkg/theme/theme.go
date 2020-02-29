package auththeme

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/qorpress/auth"
	"github.com/qorpress/i18n"
	"github.com/qorpress/i18n/backends/yaml"
	"github.com/qorpress/qor"
	"github.com/qorpress/qor/utils"
	"github.com/qorpress/render"
)

func New(config *auth.Config) *auth.Auth {
	if config == nil {
		config = &auth.Config{}
	}
	config.ViewPaths = append(config.ViewPaths, filepath.Join(utils.AppRoot, "app/views"))

	if config.DB == nil {
		fmt.Print("Please configure *gorm.DB for Auth theme clean")
	}

	if config.Render == nil {
		yamlBackend := yaml.New()
		I18n := i18n.New(yamlBackend)

		filePath := filepath.Join(utils.AppRoot, "app/locales/en-US.yml")

		if content, err := ioutil.ReadFile(filePath); err == nil {
			translations, err := yamlBackend.LoadYAMLContent(content)
			if err != nil {
				fmt.Println(err.Error())
			}
			for _, translation := range translations {
				I18n.AddTranslation(translation)
			}
		} else if err != nil {
			fmt.Println(err.Error())
		}

		config.Render = render.New(&render.Config{
			FuncMapMaker: func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
				return template.FuncMap{
					"t": func(key string, args ...interface{}) template.HTML {
						return I18n.T(utils.GetLocale(&qor.Context{Request: req}), key, args...)
					},
				}
			},
		})
	}

	Auth := auth.New(config)

	// Auth.RegisterProvider(password.New(&password.Config{
	// 	Confirmable: true,
	// 	RegisterHandler: func(context *auth.Context) (*claims.Claims, error) {
	// 		context.Request.ParseForm()

	// 		if context.Request.Form.Get("confirm_password") != context.Request.Form.Get("password") {
	// 			return nil, ErrPasswordConfirmationNotMatch
	// 		}

	// 		return password.DefaultRegisterHandler(context)
	// 	},
	// }))

	if Auth.Config.DB != nil {
		// Migrate Auth Identity model
		Auth.Config.DB.AutoMigrate(Auth.Config.AuthIdentityModel)
	}
	return Auth
}
