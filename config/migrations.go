package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	migrate "github.com/rubenv/sql-migrate"

	_ "github.com/lib/pq"
)

func RunMigrations() {
	// Bangun DSN dari variabel lingkungan
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || dbname == "" {
		log.Printf("DB_HOST: %s, DB_PORT: %s, DB_USER: %s, DB_NAME: %s", host, port, user, dbname)
		log.Fatal("Missing required database environment variables")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	// Gunakan jalur absolut untuk direktori migrasi
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	migrationDir := filepath.Join(execDir, "migrations")

	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		log.Fatalf("Migration directory '%s' does not exist", migrationDir)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: migrationDir,
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migration failed: %v. Check your migration files and database connection.", err)
	}

	log.Printf("Successfully applied %d migrations\n", n)
}
