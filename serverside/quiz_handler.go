package main

import (
	"encoding/json"
	"net/http"
	"strings"
)


func handleGenerateConditionalTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	questions := conditionalStatementGenerate(5)
	response, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Error generating conditional tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func handleGenerateDeMorganTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	questions := deMorganGenerate(5)
	response, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Error generating De Morgan tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func handleGenerateTruthTableTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	questions := truthTableGenerate(5)
	response, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Error generating truth table tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func handleAddTask(w http.ResponseWriter, r *http.Request) {
	// Only Post request is allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newQuestion Question
	err := json.NewDecoder(r.Body).Decode(&newQuestion)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	options := strings.Join(newQuestion.Options, ",")
	_, err = db.Exec("INSERT INTO quiz (question, options, correct_answer, difficulty) VALUES (?, ?, ?, ?)",
		newQuestion.Question, options, newQuestion.CorrectAnswer, newQuestion.Difficulty)
	if err != nil {
		http.Error(w, "Failed to add question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete Task
func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/deleteTask/")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM quiz WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	difficulty := strings.TrimPrefix(r.URL.Path, "/getTask/")
	if difficulty != "easy" && difficulty != "medium" && difficulty != "hard" {
		http.Error(w, "Invalid difficulty parameter", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT id, question, options, correct_answer, difficulty FROM quiz WHERE difficulty = ?", difficulty)
	if err != nil {
		http.Error(w, "Failed to retrieve questions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		var options string
		if err := rows.Scan(&q.ID, &q.Question, &options, &q.CorrectAnswer, &q.Difficulty); err != nil {
			http.Error(w, "Failed to scan question", http.StatusInternalServerError)
			return
		}
		q.Options = strings.Split(options, ",")
		questions = append(questions, q)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to iterate over questions", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Failed to marshal questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


func handleAdminPage(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Admin"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	isAdmin, err := isValidAdmin(username, password)
	if err != nil || !isAdmin {
		w.Header().Set("WWW-Authenticate", `Basic realm="Admin"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, "admin.html")
}

func handleApiPage(w http.ResponseWriter, r *http.Request) {
    username, password, ok := r.BasicAuth()
    if !ok {
        w.Header().Set("WWW-Authenticate", `Basic realm="Admin"`)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    isAdmin, err := isValidAdmin(username, password)
    if err != nil || !isAdmin {
        w.Header().Set("WWW-Authenticate", `Basic realm="Admin"`)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    http.ServeFile(w, r, "api.html")
}
