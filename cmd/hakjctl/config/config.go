package config

import (
	"time"

	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName     string = "hakjctl-config"
	configFileType     string = "yaml"
	CacheDirSubDirName string = ".hakjctl"
	TokenCacheFileName string = "._hakjctl_token_cache___"

	// ConfigKeyHost is the configuration key for host.
	ConfigKeyHost string = "host"
	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
	// ConfigKeyDatabase is the configuration key for default database.
	ConfigKeyDatabase string = "default_db"
	// ConfigKeyTlsEnabled is the configuration key for enabling TLS.
	ConfigKeyTlsEnabled string = "tls_enabled"
	// ConfigKeyTlsCertPath is the configuration key for TLS certificate path.
	ConfigKeyTlsCertPath string = "tls_cert_path"
	// ConfigKeyCommandTimeout is the configuration key for setting command timeout.
	ConfigKeyCommandTimeout string = "command_timeout"

	// EnvPrefix is the prefix for environment variables.
	EnvPrefix string = "HAKJCTL"
	// EnvVarPassword is the environment variable for password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"

	// DefaultDatabase is the name of the default database to use.
	DefaultDatabase string = "default"
	// DefaultCommandTimeout is the default command timeout in seconds.
	DefaultCommandTimeout uint32 = 10
)

// InitConfig initializes and loads configurations.
func InitConfig() {
	configDirPath, err := common.GetExecParentDirPath()
	cobra.CheckErr(err)

	viper.AddConfigPath(configDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SetDefault(ConfigKeyHost, common.ServerDefaultHost)
	viper.SetDefault(ConfigKeyPort, common.ServerDefaultPort)
	viper.SetDefault(ConfigKeyDatabase, DefaultDatabase)
	viper.SetDefault(ConfigKeyTlsEnabled, false)
	viper.SetDefault(ConfigKeyTlsCertPath, "")
	viper.SetDefault(ConfigKeyCommandTimeout, DefaultCommandTimeout)

	viper.SafeWriteConfig()
	cobra.CheckErr(viper.ReadInConfig())
}

// GetCmdTimeout gets the configured command timeout.
// Command timeout is the maximum number of seconds to wait before a request is cancelled.
func GetCmdTimeout() time.Duration {
	return time.Duration(viper.GetUint32(ConfigKeyCommandTimeout)) * time.Second
}