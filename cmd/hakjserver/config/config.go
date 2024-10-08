package config

import (
	"os"
	"path/filepath"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/viper"
)

const (
	dataDirName    string = "data"
	configFileName string = "hakjserver-config"
	configFileType string = "yaml"
	logFileName    string = "hakjserver.log"

	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
	// ConfigKeyDebugEnabled is the configuration key for debug mode.
	ConfigKeyDebugEnabled string = "debug_enabled"
	// ConfigKeyDefaultDatabase is the configuration key for default database.
	ConfigKeyDefaultDatabase string = "default_db"
	// ConfigKeyLogFileEnabled is the configuration key for enabling log file.
	ConfigKeyLogFileEnabled string = "logfile_enabled"
	// ConfigKeyTlsEnabled is the configuration key for enabling TLS.
	ConfigKeyTLSEnabled string = "tls_enabled"
	// ConfigKeyTLSClientCertAuthEnabled is the configuration key for enabling mTLS.
	ConfigKeyTLSClientCertAuthEnabled string = "tls_client_cert_auth_enabled"
	// ConfigKeyTlsCertPath is the configuration key for TLS certificate file path.
	ConfigKeyTLSCertPath string = "tls_cert_path"
	// ConfigKeyTlsPrivKeyPath is the configuration key for TLS private key file path.
	ConfigKeyTLSPrivKeyPath string = "tls_private_key_path"
	// ConfigKeyTLSCACertPath is the configuration key for TLS CA certificate file path.
	ConfigKeyTLSCACertPath string = "tls_ca_cert_path"
	// ConfigKeyMaxClientConnections is the configuration key for maximum client connections.
	ConfigKeyMaxClientConnections string = "max_client_connections"
	// ConfigKeyLogLevel is the configuration key for log level.
	ConfigKeyLogLevel string = "log_level"
	// VerboseLogsEnabled is the configuration key for enabling verbose logs.
	ConfigKeyVerboseLogsEnabled string = "verbose_logs_enabled"
	// ConfigKeyAuthEnabled is the configuration key for enabling authentication.
	ConfigKeyAuthEnabled string = "auth_enabled"
	// ConfigKeyAuthTokenSecretKey is the configuration key for setting the secret key used to sign JWT tokens.
	ConfigKeyAuthTokenSecretKey string = "auth_token_secret_key"
	// ConfigKeyAuthTokenTTL is the configuration key for setting the JWT token time to live in seconds.
	ConfigKeyAuthTokenTTL string = "auth_token_ttl"
	// ConfigKeyPassword is the configuration key for server password.
	ConfigKeyPassword string = "password"

	// EnvPrefix is the prefix that environment variables use.
	EnvPrefix string = "HAKJ"
	// EnvVarPassword is the environment variable for server password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"

	DefaultLogFileEnabled           bool   = false
	DefaultTLSEnabled               bool   = false
	DefaultTLSClientCertAuthEnabled bool   = false
	DefaultDebugEnabled             bool   = false
	DefaultVerboseLogsEnabled       bool   = false
	DefaultAuthEnabled              bool   = false
	DefaultDatabase                 string = "default"
	DefaultPassword                 string = ""
	DefaultPort                     uint16 = common.ServerDefaultPort
	DefaultLogFilePath              string = ""
	DefaultMaxKeysPerDB             uint32 = common.DbMaxKeyCount
	DefaultMaxHashMapFields         uint32 = common.HashMapMaxFields
	DefaultMaxClientConnections     uint32 = common.DefaultMaxClientConnections
	DefaultTLSCertPath              string = ""
	DefaultTLSPrivKeyPath           string = ""
	DefaultTLSCACertPath            string = ""
	DefaultLogLevel                 string = hakjdb.DefaultLogLevelStr
	DefaultAuthTokenSecretKey       string = ""
	DefaultAuthTokenTTL             uint32 = 900
)

var cfgRegistry *viper.Viper

// ServerConfig holds the server's configuration.
type ServerConfig struct {
	LogFileEnabled           bool
	TLSEnabled               bool
	TLSClientCertAuthEnabled bool
	DebugEnabled             bool
	VerboseLogsEnabled       bool
	AuthEnabled              bool

	// The name of the default database that is created at server startup.
	DefaultDB string
	// File path to the log file if it is enabled.
	// ONLY SERVER CAN CONFIGURE.
	LogFilePath string
	// The maximum number of keys a database can hold.
	// ONLY SERVER CAN CONFIGURE.
	MaxKeysPerDB uint32
	// The maximum number of fields a HashMap can hold.
	// ONLY SERVER CAN CONFIGURE.
	MaxHashMapFields uint32
	// The TCP/IP port the server listens at.
	PortInUse uint16
	// The maximum number of active client connections allowed.
	MaxClientConnections uint32

	TLSCertPath    string
	TLSPrivKeyPath string
	TLSCACertPath  string

	// Secret key used to sign JWT tokens.
	AuthTokenSecretKey string
	// JWT token time to live in seconds.
	AuthTokenTTL uint32
}

// Reload reloads configurations that can be changed at runtime.
func (cfg *ServerConfig) Reload(lg hakjdb.Logger) {
	// Config that can be changed at runtime:
	// - log file enabled
	// - verbose logs
	// - log level
	// - auth enabled
	// - password
	// - auth token secret key
	// - auth token ttl
	// - max client connections

	lg.Info("Reloading configurations ...")
	dataDirPath := getDataDirPath(lg)
	initConfigFile(dataDirPath)
	cfgRegistry.SafeWriteConfig()
	readConfigFile(lg)

	cfg.LogFileEnabled = cfgRegistry.GetBool(ConfigKeyLogFileEnabled)
	cfg.VerboseLogsEnabled = cfgRegistry.GetBool(ConfigKeyVerboseLogsEnabled)
	cfg.AuthEnabled = cfgRegistry.GetBool(ConfigKeyAuthEnabled)
	cfg.AuthTokenSecretKey = cfgRegistry.GetString(ConfigKeyAuthTokenSecretKey)
	cfg.AuthTokenTTL = cfgRegistry.GetUint32(ConfigKeyAuthTokenTTL)
	cfg.MaxClientConnections = cfgRegistry.GetUint32(ConfigKeyMaxClientConnections)
}

func InitCfgRegistry() {
	if cfgRegistry == nil {
		cfgRegistry = viper.New()
	}
}

// LoadConfig loads server configurations.
func LoadConfig(lg hakjdb.Logger) ServerConfig {
	lg.Info("Loading configurations ...")
	dataDirPath := getDataDirPath(lg)
	initConfigFile(dataDirPath)
	setConfigDefaults()

	cfgRegistry.SetEnvPrefix(EnvPrefix)
	cfgRegistry.AutomaticEnv()
	cfgRegistry.SafeWriteConfig()
	readConfigFile(lg)

	logLevel, logLevelStr, ok := hakjdb.GetLogLevelFromStr(GetLogLevelStr())
	if !ok {
		lg.Warning("Invalid log level configured. Default log level will be used")
	}
	lg.Infof("Using log level %s", logLevelStr)
	lg.SetLogLevel(logLevel)

	if cfgRegistry.GetBool(ConfigKeyVerboseLogsEnabled) {
		lg.Info("Verbose logs are enabled")
	}

	return ServerConfig{
		LogFileEnabled:           cfgRegistry.GetBool(ConfigKeyLogFileEnabled),
		TLSEnabled:               cfgRegistry.GetBool(ConfigKeyTLSEnabled),
		TLSClientCertAuthEnabled: cfgRegistry.GetBool(ConfigKeyTLSClientCertAuthEnabled),
		DebugEnabled:             cfgRegistry.GetBool(ConfigKeyDebugEnabled),
		VerboseLogsEnabled:       cfgRegistry.GetBool(ConfigKeyVerboseLogsEnabled),
		AuthEnabled:              cfgRegistry.GetBool(ConfigKeyAuthEnabled),
		DefaultDB:                cfgRegistry.GetString(ConfigKeyDefaultDatabase),
		LogFilePath:              filepath.Join(dataDirPath, logFileName),
		MaxKeysPerDB:             DefaultMaxKeysPerDB,
		MaxHashMapFields:         DefaultMaxHashMapFields,
		PortInUse:                cfgRegistry.GetUint16(ConfigKeyPort),
		MaxClientConnections:     cfgRegistry.GetUint32(ConfigKeyMaxClientConnections),
		TLSCertPath:              cfgRegistry.GetString(ConfigKeyTLSCertPath),
		TLSPrivKeyPath:           cfgRegistry.GetString(ConfigKeyTLSPrivKeyPath),
		TLSCACertPath:            cfgRegistry.GetString(ConfigKeyTLSCACertPath),
		AuthTokenSecretKey:       cfgRegistry.GetString(ConfigKeyAuthTokenSecretKey),
		AuthTokenTTL:             cfgRegistry.GetUint32(ConfigKeyAuthTokenTTL),
	}
}

func readConfigFile(lg hakjdb.Logger) {
	if err := cfgRegistry.ReadInConfig(); err != nil {
		lg.Errorf("Failed to read configuration file: %v", err)
	}
}

func getDataDirPath(lg hakjdb.Logger) string {
	parentDir, err := common.GetExecParentDirPath()
	if err != nil {
		lg.Fatalf("Failed to get the executable's parent directory: %v", err)
	}
	dataDirPath, err := common.GetDirPath(parentDir, dataDirName)
	if err != nil {
		lg.Fatalf("Failed to get the data directory: %v", err)
	}
	return dataDirPath
}

func initConfigFile(dataDirPath string) {
	cfgRegistry.AddConfigPath(dataDirPath)
	cfgRegistry.SetConfigType(configFileType)
	cfgRegistry.SetConfigName(configFileName)
}

func setConfigDefaults() {
	cfgRegistry.SetDefault(ConfigKeyDebugEnabled, DefaultDebugEnabled)
	cfgRegistry.SetDefault(ConfigKeyLogFileEnabled, DefaultLogFileEnabled)
	cfgRegistry.SetDefault(ConfigKeyTLSEnabled, DefaultTLSEnabled)
	cfgRegistry.SetDefault(ConfigKeyTLSClientCertAuthEnabled, DefaultTLSClientCertAuthEnabled)
	cfgRegistry.SetDefault(ConfigKeyVerboseLogsEnabled, DefaultVerboseLogsEnabled)
	cfgRegistry.SetDefault(ConfigKeyAuthEnabled, DefaultAuthEnabled)
	cfgRegistry.SetDefault(ConfigKeyPassword, DefaultPassword)
	cfgRegistry.SetDefault(ConfigKeyPort, DefaultPort)
	cfgRegistry.SetDefault(ConfigKeyDefaultDatabase, DefaultDatabase)
	cfgRegistry.SetDefault(ConfigKeyTLSCertPath, DefaultTLSCertPath)
	cfgRegistry.SetDefault(ConfigKeyTLSPrivKeyPath, DefaultTLSPrivKeyPath)
	cfgRegistry.SetDefault(ConfigKeyTLSCACertPath, DefaultTLSCACertPath)
	cfgRegistry.SetDefault(ConfigKeyMaxClientConnections, DefaultMaxClientConnections)
	cfgRegistry.SetDefault(ConfigKeyLogLevel, DefaultLogLevel)
	cfgRegistry.SetDefault(ConfigKeyAuthTokenSecretKey, DefaultAuthTokenSecretKey)
	cfgRegistry.SetDefault(ConfigKeyAuthTokenTTL, DefaultAuthTokenTTL)
}

// DefaultConfig returns the default configurations.
func DefaultConfig() ServerConfig {
	return ServerConfig{
		LogFileEnabled:           DefaultLogFileEnabled,
		TLSEnabled:               DefaultTLSEnabled,
		TLSClientCertAuthEnabled: DefaultTLSClientCertAuthEnabled,
		DebugEnabled:             DefaultDebugEnabled,
		VerboseLogsEnabled:       DefaultVerboseLogsEnabled,
		AuthEnabled:              DefaultAuthEnabled,
		DefaultDB:                DefaultDatabase,
		LogFilePath:              DefaultLogFilePath,
		MaxKeysPerDB:             DefaultMaxKeysPerDB,
		MaxHashMapFields:         DefaultMaxHashMapFields,
		PortInUse:                DefaultPort,
		MaxClientConnections:     DefaultMaxClientConnections,
		TLSCertPath:              DefaultTLSCertPath,
		TLSPrivKeyPath:           DefaultTLSPrivKeyPath,
		TLSCACertPath:            DefaultTLSCACertPath,
		AuthTokenSecretKey:       DefaultAuthTokenSecretKey,
		AuthTokenTTL:             DefaultAuthTokenTTL,
	}
}

// GetPassword returns the server password.
// The returned bool is true if it is set and false if not.
// Clears the password after reading it so it won't live in memory improving security.
func GetPassword() (string, bool) {
	password := cfgRegistry.GetString(ConfigKeyPassword)
	clearPassword()
	return password, password != ""
}

func clearPassword() {
	cfgRegistry.Set(ConfigKeyPassword, "")
}

func getEnvVar(envVar string) (string, bool) {
	return os.LookupEnv(envVar)
}

func GetLogLevelStr() string {
	return cfgRegistry.GetString(ConfigKeyLogLevel)
}

func GetCfgRegistry() *viper.Viper {
	return cfgRegistry
}
