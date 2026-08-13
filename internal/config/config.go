package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	App           AppConfig
	HTTP          HTTPConfig
	Database      DatabaseConfig
	Redis         RedisConfig
	Fabric        FabricConfig
	Security      SecurityConfig
	Logging       LoggingConfig
	Observability ObservabilityConfig
}

type AppConfig struct {
	Name        string
	Environment string
	Version     string
}

type HTTPConfig struct {
	Host              string
	Port              int
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ShutdownTimeout   time.Duration
	MaxBodyBytes      int64
}

type DatabaseConfig struct {
	Host                  string
	Port                  int
	Name                  string
	User                  string
	Password              string
	SSLMode               string
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type FabricConfig struct {
	Network            string
	MSPID              string
	GatewayAddress     string
	CertificatePath    string
	PrivateKeyPath     string
	TLSCertificatePath string
}

type SecurityConfig struct {
	EnableTLS   bool
	TLSCertPath string
	TLSKeyPath  string
	TrustProxy  bool
}

type LoggingConfig struct {
	Level  string
	Format string
}

type ObservabilityConfig struct {
	Enabled      bool
	ServiceName  string
	OTLPEndpoint string
}

func Load() (Config, error) {
	cfg := Config{}

	// ------------------------------------------------------------
	// Application
	// ------------------------------------------------------------

	cfg.App = AppConfig{
		Name:        getEnv("APP_NAME", "national-kyc-system"),
		Environment: getEnv("APP_ENV", "development"),
		Version:     getEnv("APP_VERSION", "0.1.0"),
	}

	// ------------------------------------------------------------
	// HTTP
	// ------------------------------------------------------------

	httpPort, err := getEnvInt("HTTP_PORT", 8080)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_PORT: %w",
			err,
		)
	}

	httpMaxBodyBytes, err := getEnvInt64(
		"HTTP_MAX_BODY_BYTES",
		1048576,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_MAX_BODY_BYTES: %w",
			err,
		)
	}

	readTimeout, err := getEnvDuration(
		"HTTP_READ_TIMEOUT",
		15*time.Second,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_READ_TIMEOUT: %w",
			err,
		)
	}

	readHeaderTimeout, err := getEnvDuration(
		"HTTP_READ_HEADER_TIMEOUT",
		5*time.Second,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_READ_HEADER_TIMEOUT: %w",
			err,
		)
	}

	writeTimeout, err := getEnvDuration(
		"HTTP_WRITE_TIMEOUT",
		30*time.Second,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_WRITE_TIMEOUT: %w",
			err,
		)
	}

	idleTimeout, err := getEnvDuration(
		"HTTP_IDLE_TIMEOUT",
		60*time.Second,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_IDLE_TIMEOUT: %w",
			err,
		)
	}

	shutdownTimeout, err := getEnvDuration(
		"HTTP_SHUTDOWN_TIMEOUT",
		10*time.Second,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_SHUTDOWN_TIMEOUT: %w",
			err,
		)
	}

	cfg.HTTP = HTTPConfig{
		Host:              getEnv("HTTP_HOST", "0.0.0.0"),
		Port:              httpPort,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ShutdownTimeout:   shutdownTimeout,
		MaxBodyBytes:      httpMaxBodyBytes,
	}

	// ------------------------------------------------------------
	// Database
	// ------------------------------------------------------------

	databasePort, err := getEnvInt(
		"DATABASE_PORT",
		5432,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid DATABASE_PORT: %w",
			err,
		)
	}

	maxOpenConnections, err := getEnvInt(
		"DATABASE_MAX_OPEN_CONNECTIONS",
		25,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid DATABASE_MAX_OPEN_CONNECTIONS: %w",
			err,
		)
	}

	maxIdleConnections, err := getEnvInt(
		"DATABASE_MAX_IDLE_CONNECTIONS",
		10,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid DATABASE_MAX_IDLE_CONNECTIONS: %w",
			err,
		)
	}

	connectionMaxLifetime, err := getEnvDuration(
		"DATABASE_CONNECTION_MAX_LIFETIME",
		time.Hour,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid DATABASE_CONNECTION_MAX_LIFETIME: %w",
			err,
		)
	}

	cfg.Database = DatabaseConfig{
		Host: getEnv(
			"DATABASE_HOST",
			"localhost",
		),
		Port: databasePort,
		Name: getEnv(
			"DATABASE_NAME",
			"national_kyc",
		),
		User: getEnv(
			"DATABASE_USER",
			"kyc_app",
		),
		Password: os.Getenv(
			"DATABASE_PASSWORD",
		),
		SSLMode: getEnv(
			"DATABASE_SSL_MODE",
			"disable",
		),
		MaxOpenConnections:    maxOpenConnections,
		MaxIdleConnections:    maxIdleConnections,
		ConnectionMaxLifetime: connectionMaxLifetime,
	}

	// ------------------------------------------------------------
	// Redis
	// ------------------------------------------------------------

	redisPort, err := getEnvInt(
		"REDIS_PORT",
		6379,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid REDIS_PORT: %w",
			err,
		)
	}

	redisDB, err := getEnvInt(
		"REDIS_DB",
		0,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid REDIS_DB: %w",
			err,
		)
	}

	cfg.Redis = RedisConfig{
		Host: getEnv(
			"REDIS_HOST",
			"localhost",
		),
		Port: redisPort,
		Password: os.Getenv(
			"REDIS_PASSWORD",
		),
		DB: redisDB,
	}

	// ------------------------------------------------------------
	// Security
	// ------------------------------------------------------------

	securityTLS, err := getEnvBool(
		"SECURITY_ENABLE_TLS",
		false,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid SECURITY_ENABLE_TLS: %w",
			err,
		)
	}

	trustProxy, err := getEnvBool(
		"SECURITY_TRUST_PROXY",
		false,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid SECURITY_TRUST_PROXY: %w",
			err,
		)
	}

	cfg.Security = SecurityConfig{
		EnableTLS: securityTLS,
		TLSCertPath: os.Getenv(
			"SECURITY_TLS_CERTIFICATE_PATH",
		),
		TLSKeyPath: os.Getenv(
			"SECURITY_TLS_PRIVATE_KEY_PATH",
		),
		TrustProxy: trustProxy,
	}

	// ------------------------------------------------------------
	// Hyperledger Fabric
	// ------------------------------------------------------------

	cfg.Fabric = FabricConfig{
		Network: getEnv(
			"FABRIC_NETWORK",
			"kycchannel",
		),
		MSPID: getEnv(
			"FABRIC_MSP_ID",
			"NKAOrgMSP",
		),
		GatewayAddress: os.Getenv(
			"FABRIC_GATEWAY_ADDRESS",
		),
		CertificatePath: os.Getenv(
			"FABRIC_CERTIFICATE_PATH",
		),
		PrivateKeyPath: os.Getenv(
			"FABRIC_PRIVATE_KEY_PATH",
		),
		TLSCertificatePath: os.Getenv(
			"FABRIC_TLS_CERTIFICATE_PATH",
		),
	}

	// ------------------------------------------------------------
	// Logging
	// ------------------------------------------------------------

	cfg.Logging = LoggingConfig{
		Level: getEnv(
			"LOG_LEVEL",
			"info",
		),
		Format: getEnv(
			"LOG_FORMAT",
			"json",
		),
	}

	// ------------------------------------------------------------
	// Observability
	// ------------------------------------------------------------

	otelEnabled, err := getEnvBool(
		"OTEL_ENABLED",
		false,
	)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid OTEL_ENABLED: %w",
			err,
		)
	}

	cfg.Observability = ObservabilityConfig{
		Enabled: otelEnabled,
		ServiceName: getEnv(
			"OTEL_SERVICE_NAME",
			"national-kyc-system",
		),
		OTLPEndpoint: os.Getenv(
			"OTEL_EXPORTER_OTLP_ENDPOINT",
		),
	}

	// ------------------------------------------------------------
	// Validate
	// ------------------------------------------------------------

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) Validate() error {

	// ------------------------------------------------------------
	// Application validation
	// ------------------------------------------------------------

	if c.App.Name == "" {
		return fmt.Errorf(
			"APP_NAME cannot be empty",
		)
	}

	switch c.App.Environment {
	case "development":
	case "test":
	case "staging":
	case "production":
	default:
		return fmt.Errorf(
			"APP_ENV must be one of development, test, staging, production",
		)
	}

	if c.App.Version == "" {
		return fmt.Errorf(
			"APP_VERSION cannot be empty",
		)
	}

	// ------------------------------------------------------------
	// HTTP validation
	// ------------------------------------------------------------

	if c.HTTP.Host == "" {
		return fmt.Errorf(
			"HTTP_HOST cannot be empty",
		)
	}

	if c.HTTP.Port < 1 || c.HTTP.Port > 65535 {
		return fmt.Errorf(
			"HTTP_PORT must be between 1 and 65535",
		)
	}

	if c.HTTP.ReadTimeout <= 0 {
		return fmt.Errorf(
			"HTTP_READ_TIMEOUT must be greater than zero",
		)
	}

	if c.HTTP.ReadHeaderTimeout <= 0 {
		return fmt.Errorf(
			"HTTP_READ_HEADER_TIMEOUT must be greater than zero",
		)
	}

	if c.HTTP.WriteTimeout <= 0 {
		return fmt.Errorf(
			"HTTP_WRITE_TIMEOUT must be greater than zero",
		)
	}

	if c.HTTP.IdleTimeout <= 0 {
		return fmt.Errorf(
			"HTTP_IDLE_TIMEOUT must be greater than zero",
		)
	}

	if c.HTTP.ShutdownTimeout <= 0 {
		return fmt.Errorf(
			"HTTP_SHUTDOWN_TIMEOUT must be greater than zero",
		)
	}

	if c.HTTP.MaxBodyBytes <= 0 {
		return fmt.Errorf(
			"HTTP_MAX_BODY_BYTES must be greater than zero",
		)
	}

	const maxGatewayBodySize int64 = 10 * 1024 * 1024

	if c.HTTP.MaxBodyBytes > maxGatewayBodySize {
		return fmt.Errorf(
			"HTTP_MAX_BODY_BYTES cannot exceed 10 MiB",
		)
	}

	// ------------------------------------------------------------
	// Database validation
	// ------------------------------------------------------------

	if c.Database.Host == "" {
		return fmt.Errorf(
			"DATABASE_HOST cannot be empty",
		)
	}

	if c.Database.Port < 1 || c.Database.Port > 65535 {
		return fmt.Errorf(
			"DATABASE_PORT must be between 1 and 65535",
		)
	}

	if c.Database.Name == "" {
		return fmt.Errorf(
			"DATABASE_NAME cannot be empty",
		)
	}

	if c.Database.User == "" {
		return fmt.Errorf(
			"DATABASE_USER cannot be empty",
		)
	}

	if c.Database.MaxOpenConnections <= 0 {
		return fmt.Errorf(
			"DATABASE_MAX_OPEN_CONNECTIONS must be greater than zero",
		)
	}

	if c.Database.MaxIdleConnections < 0 {
		return fmt.Errorf(
			"DATABASE_MAX_IDLE_CONNECTIONS cannot be negative",
		)
	}

	if c.Database.MaxIdleConnections >
		c.Database.MaxOpenConnections {
		return fmt.Errorf(
			"DATABASE_MAX_IDLE_CONNECTIONS cannot exceed DATABASE_MAX_OPEN_CONNECTIONS",
		)
	}

	if c.Database.ConnectionMaxLifetime <= 0 {
		return fmt.Errorf(
			"DATABASE_CONNECTION_MAX_LIFETIME must be greater than zero",
		)
	}

	switch c.Database.SSLMode {
	case "disable":
	case "require":
	case "verify-ca":
	case "verify-full":
	default:
		return fmt.Errorf(
			"DATABASE_SSL_MODE is invalid",
		)
	}

	// ------------------------------------------------------------
	// Redis validation
	// ------------------------------------------------------------

	if c.Redis.Host == "" {
		return fmt.Errorf(
			"REDIS_HOST cannot be empty",
		)
	}

	if c.Redis.Port < 1 || c.Redis.Port > 65535 {
		return fmt.Errorf(
			"REDIS_PORT must be between 1 and 65535",
		)
	}

	if c.Redis.DB < 0 {
		return fmt.Errorf(
			"REDIS_DB cannot be negative",
		)
	}

	// ------------------------------------------------------------
	// Security validation
	// ------------------------------------------------------------

	if c.Security.EnableTLS {

		if c.Security.TLSCertPath == "" {
			return fmt.Errorf(
				"SECURITY_TLS_CERTIFICATE_PATH is required when TLS is enabled",
			)
		}

		if c.Security.TLSKeyPath == "" {
			return fmt.Errorf(
				"SECURITY_TLS_PRIVATE_KEY_PATH is required when TLS is enabled",
			)
		}
	}

	// ------------------------------------------------------------
	// Fabric validation
	// ------------------------------------------------------------

	if c.Fabric.Network == "" {
		return fmt.Errorf(
			"FABRIC_NETWORK cannot be empty",
		)
	}

	if c.Fabric.MSPID == "" {
		return fmt.Errorf(
			"FABRIC_MSP_ID cannot be empty",
		)
	}

	// ------------------------------------------------------------
	// Logging validation
	// ------------------------------------------------------------

	switch c.Logging.Level {
	case "debug":
	case "info":
	case "warn":
	case "error":
	default:
		return fmt.Errorf(
			"LOG_LEVEL must be one of debug, info, warn, error",
		)
	}

	switch c.Logging.Format {
	case "json":
	case "text":
	default:
		return fmt.Errorf(
			"LOG_FORMAT must be either json or text",
		)
	}

	// ------------------------------------------------------------
	// Observability validation
	// ------------------------------------------------------------

	if c.Observability.ServiceName == "" {
		return fmt.Errorf(
			"OTEL_SERVICE_NAME cannot be empty",
		)
	}

	if c.Observability.Enabled &&
		c.Observability.OTLPEndpoint == "" {
		return fmt.Errorf(
			"OTEL_EXPORTER_OTLP_ENDPOINT is required when observability is enabled",
		)
	}

	// ------------------------------------------------------------
	// Production security policy
	// ------------------------------------------------------------

	if c.App.Environment == "production" {

		if !c.Security.EnableTLS {
			return fmt.Errorf(
				"TLS must be enabled in production",
			)
		}

		if c.Database.SSLMode == "disable" {
			return fmt.Errorf(
				"DATABASE_SSL_MODE=disable is not allowed in production",
			)
		}

		if c.Database.Password == "" {
			return fmt.Errorf(
				"DATABASE_PASSWORD is required in production",
			)
		}

		if c.Redis.Password == "" {
			return fmt.Errorf(
				"REDIS_PASSWORD is required in production",
			)
		}

		if c.Fabric.GatewayAddress == "" {
			return fmt.Errorf(
				"FABRIC_GATEWAY_ADDRESS is required in production",
			)
		}

		if c.Fabric.CertificatePath == "" {
			return fmt.Errorf(
				"FABRIC_CERTIFICATE_PATH is required in production",
			)
		}

		if c.Fabric.PrivateKeyPath == "" {
			return fmt.Errorf(
				"FABRIC_PRIVATE_KEY_PATH is required in production",
			)
		}

		if c.Fabric.TLSCertificatePath == "" {
			return fmt.Errorf(
				"FABRIC_TLS_CERTIFICATE_PATH is required in production",
			)
		}
	}

	return nil
}

func getEnv(
	key string,
	fallback string,
) string {
	value, exists := os.LookupEnv(key)

	if !exists || value == "" {
		return fallback
	}

	return value
}

func getEnvInt(
	key string,
	fallback int,
) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	return strconv.Atoi(value)
}

func getEnvInt64(
	key string,
	fallback int64,
) (int64, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	return strconv.ParseInt(
		value,
		10,
		64,
	)
}

func getEnvBool(
	key string,
	fallback bool,
) (bool, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	return strconv.ParseBool(value)
}

func getEnvDuration(
	key string,
	fallback time.Duration,
) (time.Duration, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	return time.ParseDuration(value)
}
