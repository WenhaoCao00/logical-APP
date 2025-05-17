package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./quizeit.db")
	if err != nil {
		log.Fatal(err)
	}

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		is_admin BOOLEAN DEFAULT 0,
		avatarUrl TEXT,
		websiteUrls TEXT,
		bio TEXT,
		description TEXT,
		totalQuizNumber INTEGER DEFAULT 0,
		quizScores TEXT
	);`

	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}

	createQuizTable := `CREATE TABLE IF NOT EXISTS quiz (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		question TEXT NOT NULL,
		options TEXT NOT NULL,
		correct_answer TEXT NOT NULL,
		difficulty TEXT NOT NULL
	);`

	_, err = db.Exec(createQuizTable)
	if err != nil {
		log.Fatal(err)
	}

	
}

func insertDefaultAdmin() {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE is_admin = 1").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		adminUsername := "admin"
		adminEmail := "admin@example.com"
		adminPassword := "admin"

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Failed to hash admin password:", err)
		}

		query := `INSERT INTO users (uuid, name, username, email, password, is_admin) VALUES (?, ?, ?, ?, ?, 1)`
		_, err = db.Exec(query, "admin-uuid", "Administrator", adminUsername, adminEmail, string(hashedPassword))
		if err != nil {
			log.Fatal("Failed to insert default admin:", err)
		}

		fmt.Println("Default admin account created with username: admin and password: admin123")
	} else {
		fmt.Println("Admin account already exists.")
	}
}
// func insertDefaultTester() {
// 	var count int
// 	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'asdfgh'").Scan(&count)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if count == 0 {
// 		testUsername := "asdfgh"
// 		testEmail := "asdfgh"
// 		testPassword := "asdfgh"

// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)
// 		if err != nil {
// 			log.Fatal("Failed to hash tester password:", err)
// 		}

// 		query := `INSERT INTO users (uuid, name, username, email, password, is_admin) VALUES (?, ?, ?, ?, ?, 0)`
// 		_, err = db.Exec(query, "tester-uuid", "Tester", testUsername, testEmail, string(hashedPassword))
// 		if err != nil {
// 			log.Fatal("Failed to insert default tester:", err)
// 		}

// 		fmt.Println("Default tester account created with username: asdfgh and password: asdfgh")
// 	} else {
// 		fmt.Println("Tester account already exists.")
// 	}
// }


// check if the valid admin
func isValidAdmin(username, password string) (bool, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ? AND is_admin = 1", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil 
		}
		return false, err
	}

	// check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil 
	}

	return true, nil
}

func registerUser(user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (uuid, name, username, email, password, avatarUrl, websiteUrls, bio, description, totalQuizNumber, quizScores) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, user.UUID, user.Name, user.Username, user.Email, string(hashedPassword), user.AvatarUrl, strings.Join(user.WebsiteUrls, ","), user.Bio, user.Description, user.TotalQuizNumber, formatQuizScores(user.QuizScores))
	return err
}

func authenticateUser(email, password string) (User, error) {
	var user User
	var websiteUrls, quizScores string

	err := db.QueryRow("SELECT uuid, name, username, email, password, avatarUrl, websiteUrls, bio, description, totalQuizNumber, quizScores FROM users WHERE email = ?", email).Scan(&user.UUID, &user.Name, &user.Username, &user.Email, &user.Password, &user.AvatarUrl, &websiteUrls, &user.Bio, &user.Description, &user.TotalQuizNumber, &quizScores)
	if err != nil {
		return User{}, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, fmt.Errorf("incorrect password")
	}

	user.WebsiteUrls = strings.Split(websiteUrls, ",")
	user.QuizScores = parseQuizScores(quizScores)

	user.Password = ""

	return user, nil
}

func getUserFromDB(username string) (User, error) {
	var user User
	var avatarUrl, bio, description, websiteUrls, quizScores sql.NullString
	var totalQuizNumber sql.NullInt64

	query := `SELECT uuid, name, username, email, avatarUrl, websiteUrls, bio, description, totalQuizNumber, quizScores FROM users WHERE username = ?`
	row := db.QueryRow(query, username)

	// Scanning nullable fields using sql.NullString and sql.NullInt64
	err := row.Scan(&user.UUID, &user.Name, &user.Username, &user.Email, &avatarUrl, &websiteUrls, &bio, &description, &totalQuizNumber, &quizScores)
	if err != nil {
		return User{}, err
	}

	// Convert nullable fields to standard Go types
	user.AvatarUrl = getStringFromNullString(avatarUrl)
	user.WebsiteUrls = getStringArrayFromNullString(websiteUrls)
	user.Bio = getStringFromNullString(bio)
	user.Description = getStringFromNullString(description)
	user.TotalQuizNumber = getIntFromNullInt64(totalQuizNumber)
	user.QuizScores = getQuizScoresFromNullString(quizScores)

	return user, nil
}

func saveUserToDB(user User) error {
	websiteUrls := strings.Join(user.WebsiteUrls, ",")
	quizScores := formatQuizScores(user.QuizScores)

	query := `UPDATE users 
	          SET name = ?, avatarUrl = ?, websiteUrls = ?, bio = ?, description = ?, totalQuizNumber = ?, quizScores = ? 
	          WHERE username = ?`

	_, err := db.Exec(query, user.Name, user.AvatarUrl, websiteUrls, user.Bio, user.Description, user.TotalQuizNumber, quizScores, user.Username)
	return err
}


func insertFixedQuestions() {
	//to avoid inserting questions multiple times
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM quiz").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}


	if count > 0 {
		fmt.Println("Quiz table already contains questions. Skipping insertion.")
		return
	}

	questions := generateFixedQuestions()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO quiz (question, options, correct_answer, difficulty) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, difficulty := range []string{"easy", "medium", "hard"} {
		for _, q := range questions[difficulty] {
			options := strings.Join(q.Options, ",")
			_, err = stmt.Exec(q.Question, options, q.CorrectAnswer, q.Difficulty)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}




//helper function for user
// Convert sql.NullString to string
func getStringFromNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// Convert sql.NullString to []string (splitting comma-separated values)
func getStringArrayFromNullString(ns sql.NullString) []string {
	if ns.Valid {
		return strings.Split(ns.String, ",")
	}
	return []string{}
}

// Convert sql.NullInt64 to int
func getIntFromNullInt64(ni sql.NullInt64) int {
	if ni.Valid {
		return int(ni.Int64)
	}
	return 0
}

// Convert sql.NullString to map[string]int (for quiz scores)
func getQuizScoresFromNullString(ns sql.NullString) map[string]int {
	if ns.Valid {
		return parseQuizScores(ns.String)
	}
	return make(map[string]int)
}

// Convert quiz scores from string to map[string]int
func parseQuizScores(quizScores string) map[string]int {
	result := make(map[string]int)
	entries := strings.Split(quizScores, ",")
	for _, entry := range entries {
		if entry == "" {
			continue
		}
		parts := strings.Split(entry, ":")
		score, _ := strconv.Atoi(parts[1])
		result[parts[0]] = score
	}
	return result
}

// Convert map[string]int to string (formatting quiz scores as a string)
func formatQuizScores(quizScores map[string]int) string {
	var entries []string
	for k, v := range quizScores {
		entries = append(entries, fmt.Sprintf("%s:%d", k, v))
	}
	return strings.Join(entries, ",")
}