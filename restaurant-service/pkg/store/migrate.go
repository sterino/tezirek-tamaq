package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/pressly/goose/v3"
)

func Migrate(driverName, dataSourceName string) error {
	if !strings.Contains(dataSourceName, "://") {
		return errors.New("store: invalid data source name " + dataSourceName)
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return fmt.Errorf("store: failed to open DB: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect(driverName); err != nil {
		return fmt.Errorf("store: goose set dialect failed: %w", err)
	}

	migrationPath := fmt.Sprintf("migrations/%s", driverName)
	if err := goose.Up(db, migrationPath); err != nil {
		return fmt.Errorf("store: migration failed: %w", err)
	}

	return nil
}
