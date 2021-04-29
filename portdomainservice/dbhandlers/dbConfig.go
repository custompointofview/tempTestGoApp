package dbhandlers

type DBConfig struct {
	database   string
	collection string
}

func NewDBConfig(database string, collection string) *DBConfig {
	return &DBConfig{database: database, collection: collection}
}
