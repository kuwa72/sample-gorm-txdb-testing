package testutil

import (
	"fmt"
	"kuwa72/sample-gorm-txdb-testing/usecase"
	"os"
	"path"
	"path/filepath"

	"github.com/DATA-DOG/go-txdb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB(db *gorm.DB) {
	db.AutoMigrate(usecase.User{})
}

func findModuleRoot(dir string) (string, error) {
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir, nil
		}
		d := filepath.Dir(dir)
		if d == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = d
	}
}

func NewTestDB(name string) (*gorm.DB, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	rootdir, err := findModuleRoot(wd)
	if err != nil {
		return nil, err
	}
	dsn := path.Join(rootdir, "test.db")
	txdb.Register(name, "sqlite3", dsn)
	dialector := sqlite.New(sqlite.Config{
		DriverName: name,
		DSN:        dsn,
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	initDB(db)
	return db, nil
}
