package config

import (
	"os"
	"strconv"

	"github.com/qorpress/mailer/gomailer"
	"github.com/go-gomail/gomail"
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
	"github.com/qorpress/auth/providers/facebook"
	"github.com/qorpress/auth/providers/github"
	"github.com/qorpress/auth/providers/google"
	"github.com/qorpress/auth/providers/twitter"
	"github.com/qorpress/location"
	"github.com/qorpress/mailer"
	// "github.com/qorpress/mailer/logger"
	"github.com/qorpress/media/oss"
	"github.com/qorpress/oss/s3"
	"github.com/qorpress/redirect_back"
	"github.com/qorpress/session/manager"
	"github.com/unrolled/render"
	// "github.com/qorpress/i18n"
	// "github.com/qorpress/i18n/backends/yaml"
)

/*
- Fixed：Unknown column 'basics.provider' in 'where clause')
  - https://github.com/qorpress/auth/pull/20/files
*/ 

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type AuthConfig struct {
	Github   github.Config
	Google   google.Config
	Facebook facebook.Config
	Twitter  twitter.Config
}

type ApiKeyConfig struct {
	GoogleAPIKey string `env:"GoogleAPIKey"`
	BaiduAPIKey  string `env:"BaiduAPIKey"`
	Twitter      TwitterApiConfig
}

type TwitterApiConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type BucketConfig struct {
	S3 struct {
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
		Region          string `env:"AWS_Region"`
		S3Bucket        string `env:"AWS_Bucket"`
	}
}

var Config = struct {
	HTTPS bool `default:"false" env:"HTTPS"`
	Port  uint `default:"4000" env:"PORT"`
	DB    struct {
		Name     string `env:"DBName" default:"gopress"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
	Bucket BucketConfig
	ApiKey ApiKeyConfig
	Auth   AuthConfig
	SMTP   SMTPConfig
	// Schema []generator.Model
}{}

var (
	Root         = os.Getenv("GOPATH") + "/src/github.com/qorpress/qorpress-auth-example"
	Mailer       *mailer.Mailer
	Render       = render.New()
	RedirectBack = redirect_back.New(&redirect_back.Config{
		SessionManager:  manager.SessionManager,
		IgnoredPrefixes: []string{"/auth"},
	})
)

func init() {
	if err := configor.Load(&Config, ".config/gopress.yml", ".config/gopress.yaml"); err != nil {
		panic(err)
	}

	pp.Println(Config)

	location.GoogleAPIKey = Config.ApiKey.GoogleAPIKey
	location.BaiduAPIKey = Config.ApiKey.BaiduAPIKey

	if Config.Bucket.S3.AccessKeyID != "" {
		oss.Storage = s3.New(&s3.Config{
			AccessID:  Config.Bucket.S3.AccessKeyID,
			AccessKey: Config.Bucket.S3.SecretAccessKey,
			Region:    Config.Bucket.S3.Region,
			Bucket:    Config.Bucket.S3.S3Bucket,
		})
	}

	portSmtp, err := strconv.Atoi(Config.SMTP.Port)
	if err != nil {
		panic(err)
	}

	dialer := gomail.NewDialer(Config.SMTP.Host, portSmtp, Config.SMTP.User, Config.SMTP.Password)
	sender, err := dialer.Dial()
	if err != nil {
		panic(err)
	}
	pp.Println(sender)

	Mailer = mailer.New(&mailer.Config{
		Sender: gomailer.New(&gomailer.Config{Sender: sender}),
	})
	//Mailer = mailer.New(&mailer.Config{
	//	Sender: logger.New(&logger.Config{}),
	//})
}
