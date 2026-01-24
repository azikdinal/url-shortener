package config

type StoreType string

const (
	InMemoryStoreType StoreType = "IN_MEMORY"
	PostgresStoreType StoreType = "POSTGRES"
)
