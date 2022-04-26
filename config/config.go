package config

type (
	Config struct {
		App     App     `yaml:"app"`
		Admin   Admin   `yaml:"admin"`
		Mysql   Mysql   `yaml:"mysql"`
		Account Account `yaml:"account"`
	}

	App struct {
		Name    string `yaml:"name" envconfig:"APP_NAME"`
		Address string `yaml:"port" envconfig:"APP_ADDRESS"`
	}

	Admin struct {
		Username string `yaml:"username" envconfig:"ADMIN_USERNAME"`
		Password string `yaml:"password" envconfig:"ADMIN_PASSWORD"`
	}

	Mysql struct {
		Username string `yaml:"username" envconfig:"MYSQL_USERNAME"`
		Password string `yaml:"password" envconfig:"MYSQL_PASSWORD"`
		DBName   string `yaml:"db_name" envconfig:"MYSQL_DBNAME"`
		Host     string `yaml:"host" envconfig:"MYSQL_HOST"`
		Port     string `yaml:"port" envconfig:"MYSQL_PORT"`
	}

	Account struct {
		MinUsernameLength int `yaml:"min_username_length"`
	}
)

func GetAdmin() Admin {
	return Admin{
		Username: "admin",
		Password: "secret",
	}
}
