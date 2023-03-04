package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)


var DB *sql.DB

// ConnectDB tries to connect DB and on succcesful it returns
// DB connection string and nil error, otherwise return empty DB and the corresponding error.
func ConnectDB() (error) {
	// database variables
	// usually we should get them from env like os.Getenv("variableName")
	var (
		host			= os.Getenv("DB_HOST")
		portStr		= os.Getenv("DB_PORT")
		user     	= os.Getenv("DB_USER")
		password 	= os.Getenv("DB_PASSWORD")
		dbname   	= os.Getenv("DB_NAME")
	)

	port, err := strconv.ParseInt(portStr, 10, 0);
	if err != nil {
		log.Fatal("Invalid DB port");
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname);
	db, err := sql.Open("postgres", connString);

	if err != nil {
		log.Printf("failed to connect to database: %v", err);
		return err;
	}

	DB = db;
	return nil;
}

func RunMigrations() (error) {
	driver, err := postgres.WithInstance(DB, &postgres.Config{});
	if err != nil {
		log.Print("migration: unable to connect to existing DB instance");
		log.Fatal(err);

		return err;
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://database/migrations",	"postgres", driver);
	if err != nil {
		log.Print("migration: Unable to get migrations");
		log.Fatal(err)

		return err;
	}

	return upMigrations(migrations);
}

func upMigrations(migrations *migrate.Migrate) (error) {
	log.Print("migration: Starting UP migrations...");
	err := migrations.Up();
	if (err != nil && err != migrate.ErrNoChange) {

		if (err == migrate.ErrNoChange) {
			log.Print("migration: No new migrations found.");
		} else {
			log.Print("migration: Unable to run UP migration");
			log.Fatal(err);
		}
	}

	log.Print("migration: UP migrations completed");

	return nil;
}
