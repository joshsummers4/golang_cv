package cv

import (
	"context"
	"fmt"
	
	"github.com/joshsummers4/golang_cv/libs/utils/database"
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
		fmt.Printf("Init contact: failed to open database: %v\n", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(createContactTableSQL)
	if err != nil {
		fmt.Printf("Init contact: failed to create contact table: %v\n", err)
		return
	}
}

func AddContact(ctx context.Context, name, email, message string) error {
	db, err := database.Open("./contact.db")
	if err != nil {
		fmt.Printf("contact add contact: failed to open database: %v\n", err)
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, `
		INSERT INTO contact (name, email, message)
		VALUES ($1, $2, $3)
	`, name, email, message)

	if err != nil {
		fmt.Printf("contact add contact: failed to add contact: %v\n", err)
		return fmt.Errorf("failed to add contact: %w", err)
	}

	fmt.Printf("contact add contact: contact form submitted: %v\n", map[string]any{"name": name, "email": email})
	return nil
}
