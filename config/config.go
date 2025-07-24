package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql   Mssql   `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql   Pgsql   `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle  Oracle  `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Sqlite  Sqlite  `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Timer   Timer   `mapstructure:"timer" json:"timer" yaml:"timer"`
}
