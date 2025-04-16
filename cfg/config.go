package cfg

type Config struct {
	Host              string
	Port              string
	ChunkSize         int
	HttpClientTimeout int
	StoragePath       string
	Limits            map[string]int
	SQLiteDBPath      string
}

func LoadNewDefaultConfig() *Config {
	var cfg = &Config{
		Host:              "127.0.0.1",
		Port:              "50051",
		ChunkSize:         4096,
		HttpClientTimeout: 30,
		StoragePath:       "./storage",
		Limits: map[string]int{
			"upload":   10,
			"download": 10,
			"list":     10,
		},
		SQLiteDBPath: "./storage/sqlite/metadata.db",
	}

	return cfg
}
