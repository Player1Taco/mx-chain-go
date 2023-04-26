package factory_test

import (
	"testing"

	"github.com/multiversx/mx-chain-go/storage"
	"github.com/multiversx/mx-chain-go/storage/factory"
	"github.com/stretchr/testify/require"
)

func TestNewPersisterFactory(t *testing.T) {
	t.Parallel()

	dbConfigHandler := factory.NewDBConfigHandler(createDefaultDBConfig())
	pf, err := factory.NewPersisterFactory(dbConfigHandler)
	require.NotNil(t, pf)
	require.Nil(t, err)
}

func TestPersisterFactory_Create(t *testing.T) {
	t.Parallel()

	t.Run("invalid file path, should fail", func(t *testing.T) {
		t.Parallel()

		dbConfigHandler := factory.NewDBConfigHandler(createDefaultDBConfig())
		pf, _ := factory.NewPersisterFactory(dbConfigHandler)

		p, err := pf.Create("")
		require.Nil(t, p)
		require.Equal(t, storage.ErrInvalidFilePath, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		dbConfigHandler := factory.NewDBConfigHandler(createDefaultDBConfig())
		pf, _ := factory.NewPersisterFactory(dbConfigHandler)

		dir := t.TempDir()

		p, err := pf.Create(dir)
		require.NotNil(t, p)
		require.Nil(t, err)
	})
}

// func TestPersisterFactory_GetDBConfig(t *testing.T) {
// 	t.Parallel()

// 	t.Run("create persister factory twice, should load from config file", func(t *testing.T) {
// 		t.Parallel()

// 		expectedDBConfig := createDefaultDBConfig()

// 		dbConfigHandler := factory.NewDBConfigHandler(createDefaultDBConfig())
// 		pf, _ := factory.NewPersisterFactory(dbConfigHandler)

// 		dirPath := t.TempDir()

// 		_, err := pf.Create(dirPath)
// 		require.Nil(t, err)

// 		newDBConfig := config.DBConfig{
// 			Type:              "type1",
// 			BatchDelaySeconds: 1,
// 			MaxBatchSize:      2,
// 			MaxOpenFiles:      3,
// 			NumShards:         4,
// 		}

// 		dbConfigHandler = factory.NewDBConfigHandler(newDBConfig)
// 		pf, _ = factory.NewPersisterFactory(dbConfigHandler)

// 		conf, err := pf.GetDBConfig(dirPath)
// 		require.Nil(t, err)
// 		require.Equal(t, &expectedDBConfig, conf)
// 	})
// }
