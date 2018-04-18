package config

var GlobCfg = Config{}

type Config struct {
	TablePrefix string `toml:"table_prefix"`
}
