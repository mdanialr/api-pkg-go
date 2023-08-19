package conf

import "github.com/spf13/viper"

// InitConfigYml init config that use yaml file. Optionally use given string as
// the config file name to look for.
func InitConfigYml(configName ...string) (*viper.Viper, error) {
	name := "app"
	if len(configName) > 0 {
		name = configName[0]
	}
	return initConfig("yaml", name)
}

// initConfig return new viper config and error if any. Use given
// string as the config file name to look for.
func initConfig(t string, configName string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(t)
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	v.WatchConfig()

	return v, nil
}
