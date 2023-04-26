package factory

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-go/config"
)

const (
	dbConfigFileName         = "config.toml"
	defaultType              = "LvlDBSerial"
	defaultBatchDelaySeconds = 2
	defaultMaxBatchSize      = 100
	defaultMaxOpenFiles      = 10
)

type dbConfigHandler struct {
	dbType              string
	batchDelaySeconds   int
	maxBatchSize        int
	maxOpenFiles        int
	shardIDProviderType string
	numShards           int32
}

// NewDBConfigHandler will create a new db config handler instance
func NewDBConfigHandler(config config.DBConfig) *dbConfigHandler {
	return &dbConfigHandler{
		dbType:              config.Type,
		batchDelaySeconds:   config.BatchDelaySeconds,
		maxBatchSize:        config.MaxBatchSize,
		maxOpenFiles:        config.MaxOpenFiles,
		shardIDProviderType: config.ShardIDProviderType,
		numShards:           config.NumShards,
	}
}

// GetDBConfig will get the db config based on path
func (dh *dbConfigHandler) GetDBConfig(path string) (*config.DBConfig, error) {
	dbConfigFromFile := &config.DBConfig{}
	err := core.LoadTomlFile(dbConfigFromFile, getPersisterConfigFilePath(path))
	if err == nil {
		log.Debug("getDBConfig: loaded db config from toml config file", "path", dbConfigFromFile)
		return dbConfigFromFile, nil
	}

	empty := checkIfDirIsEmpty(path)
	if !empty {
		dbConfig := &config.DBConfig{
			Type:              defaultType,
			BatchDelaySeconds: defaultBatchDelaySeconds,
			MaxBatchSize:      defaultMaxBatchSize,
			MaxOpenFiles:      defaultMaxOpenFiles,
		}

		log.Debug("getDBConfig: loaded default db config")
		return dbConfig, nil
	}

	dbConfig := &config.DBConfig{
		Type:                dh.dbType,
		BatchDelaySeconds:   dh.batchDelaySeconds,
		MaxBatchSize:        dh.maxBatchSize,
		MaxOpenFiles:        dh.maxOpenFiles,
		ShardIDProviderType: dh.shardIDProviderType,
		NumShards:           dh.numShards,
	}

	log.Debug("getDBConfig: loaded db config from main config file")
	return dbConfig, nil
}

// SaveDBConfigToFilePath will save the provided db config to specified path
func (dh *dbConfigHandler) SaveDBConfigToFilePath(path string, dbConfig *config.DBConfig) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			// in memory db, no files available
			log.Debug("createPersisterConfigFile: provided path not available, config file will not be created")
			return nil
		}

		return err
	}

	configFilePath := getPersisterConfigFilePath(path)
	f, err := core.OpenFile(configFilePath)
	defer func() {
		_ = f.Close()
	}()
	if err == nil {
		// config file already exists, no need to save config
		return nil
	}

	err = core.SaveTomlFile(dbConfig, configFilePath)
	if err != nil {
		return err
	}

	return nil
}

func getPersisterConfigFilePath(path string) string {
	return filepath.Join(
		path,
		dbConfigFileName,
	)
}

func checkIfDirIsEmpty(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Debug("getDBConfig: failed to check if dir is empty", "path", path, "error", err.Error())
		return true
	}

	if len(files) == 0 {
		return true
	}

	return false
}

// IsInterfaceNil returns true if there is no value under the interface
func (dh *dbConfigHandler) IsInterfaceNil() bool {
	return dh == nil
}
