package cv

import (
	"context"
	"fmt"
	
	"github.com/joshsummers4/golang_cv/libs/utils/database"
	"github.com/joshsummers4/golang_cv/libs/utils/logger"
)

 const createContactTableSQL = `
CREATE TABLE IF NOT EXISTS contact (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT,
	message TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func init() {
	// Initialize the contact database and create the table if it doesn't exist
	db, err := database.Open("./contact.db")
	if err != nil {
		logger.Error(context.Background(), "failed to open contact database", err, []string{"server"}, nil)
		return
	}
	defer db.Close()

	_, err = db.Exec(createContactTableSQL)
	if err != nil {
		logger.Error(context.Background(), "failed to create contact table", err, []string{"server"}, nil)
	}
}

func AddContact(ctx context.Context, name, email, message string) error {
	db, err := database.Open("./contact.db")
	if err != nil {
		logger.Error(ctx, "failed to open database", err, []string{"server"}, nil)
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, `
		INSERT INTO contact (name, email, message)
		VALUES ($1, $2, $3)
	`, name, email, message)

	if err != nil {
		logger.Error(ctx, "failed to add contact", err, []string{"server"}, map[string]any{"name": name, "email": email})
		return fmt.Errorf("failed to add contact: %w", err)
	}

	logger.Info(ctx, "contact form submitted", []string{"server"}, map[string]any{"name": name, "email": email})
	return nil
}
