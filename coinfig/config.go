package config

type SMSConfig struct {
	ACCESS_KEY_ID     string `yaml:"ACCESS_KEY_ID"`
	ACCESS_KEY_SECRET string `yaml:"ACCESS_KEY_SECRET"`
}

type CacheConfig struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB int `yaml:"db"`
	Password string `yaml:"password"`
}

type SettingConfig struct {
	SECRET_KEY string `yaml:"SECRET_KEY"`
	DEBUG string `yaml:"DEBUG"`
	DEFAULT_CHARSET string `yaml:"DEFAULT_CHARSET"`
	SMSConfig SMSConfig `yaml:"SMSConfig"`
	CACHES []CacheConfig `yaml:"CACHES"`
}
