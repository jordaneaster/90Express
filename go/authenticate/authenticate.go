package authenticate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"   // Import JWT library
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// TokenExpirationTime represents the expiration time for JWT tokens (e.g., 24 hours)
const TokenExpirationTime = time.Hour * 24

type Profile struct {
	UserID   int
	Username string
	Email    string
	FullName string
	Company  string
}

type AuthDB interface {
	Close()
	InsertUser(username, password string) error
	AuthenticateUserDB(username, password string) (bool, error)
}

type authDB struct {
	db *sql.DB
}

func NewAuthDB() (*authDB, error) {
	// Ensure the directory exists
	err := os.MkdirAll("go/auth", 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// Create or open the SQLite file
	dbFile := "go/auth/users.db"
	log.Printf("Creating or opening database file: %s\n", dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create users table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password TEXT,
        profile_id INTEGER
    );`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create users table: %v", err)
	}

	// Create profiles table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS profiles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        email TEXT,
        full_name TEXT,
        company TEXT
    );`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create profiles table: %v", err)
	}

	return &authDB{db: db}, nil
}

func CheckUserProfileExists(username string) bool {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "go/auth/users.db")
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return false
	}
	defer db.Close()

	// Query the database to check if the user profile exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM profiles WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		log.Printf("Error checking user profile: %v\n", err)
		return false
	}

	return exists
}

// createUserProfile inserts a new user profile into the database.
func CreateUserProfile(profile Profile) error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "go/auth/users.db")
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return err
	}
	defer db.Close()

	// Prepare the SQL statement for inserting a new profile
	stmt, err := db.Prepare("INSERT INTO profiles (username, email, full_name, company) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v\n", err)
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to insert the new profile
	_, err = stmt.Exec(profile.Username, profile.Email, profile.FullName, profile.Company)
	if err != nil {
		log.Printf("Error inserting new profile: %v\n", err)
		return err
	}

	log.Printf("New profile created for user: %s\n", profile.Username)
	return nil
}

// GenerateToken generates a JWT token for the given username
func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(TokenExpirationTime).Unix(),
	})
	signedToken, err := token.SignedString([]byte("secret")) // Use a secret key to sign the token
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

type AuthenticationResponse struct {
	Authenticated bool   `json:"authenticated"`
	Token         string `json:"token,omitempty"`
	Username      string `json:"username,omitempty"` // Add Username field
}

// AuthenticateUserDB authenticates the user and returns a token upon successful authentication
func (adb *authDB) AuthenticateUserDB(username, password string) (AuthenticationResponse, error) {
	var response AuthenticationResponse

	var count int
	err := adb.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND password = ?", username, password).Scan(&count)
	if err != nil {
		return response, err
	}

	if count > 0 {
		// User exists, generate token and include it in the response
		token, err := GenerateToken(username)
		if err != nil {
			return response, err
		}
		response.Authenticated = true
		response.Token = token
	} else {
		// User does not exist
		response.Authenticated = false
	}

	return response, nil
}
