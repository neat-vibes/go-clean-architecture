package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Log      LogConfig       `json:"log"`
	Server   Server          `json:"server"`
	Context  Context         `json:"context"`
	Database DatabaseOptions `json:"database"`
}

type LogConfig struct {
	Level string `json:"level"`
	Path  string `json:"path"`
}

type Server struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Context struct {
	Timeout int `json:"timeout"`
}

type DatabaseOptions struct {
	Type  string       `json:"type"`
	MySQL *MySQLConfig `json:"mysql,omitempty"`
	// Other database types can be added for configuration
}

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func NewConfig(confFilePath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(confFilePath)
	v.SetConfigType("json")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var conf Config
	err = v.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Config) GetLogLevel() string {
	return c.Log.Level
}

func (c *Config) GetLogPath() string {
	return c.Log.Path
}

func (c *Config) GetDatabaseType() string {
	return c.Database.Type
}

func (c *Config) GetServerAddress() string {
	return c.Server.Address
}

func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func (c *Config) GetContextTimeout() int {
	return c.Context.Timeout
}

func (c *Config) PrintConfig() {
	fmt.Println("Log Level:", c.Log.Level)
	fmt.Println("Log Path:", c.Log.Path)
	fmt.Println("Server Address:", c.Server.Address)
	fmt.Println("Server Port:", c.Server.Port)
	fmt.Println("Context Timeout:", c.Context.Timeout)
	fmt.Println("Database Type:", c.Database.Type)

	if c.Database.MySQL != nil {
		mysqlConfig := c.Database.MySQL
		fmt.Println("MySQL Host:", mysqlConfig.Host)
		fmt.Println("MySQL Port:", mysqlConfig.Port)
		fmt.Println("MySQL User:", mysqlConfig.User)
		fmt.Println("MySQL Password:", mysqlConfig.Password)
		fmt.Println("MySQL Name:", mysqlConfig.DBName)
	}
}
