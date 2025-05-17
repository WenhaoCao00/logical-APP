# Quizeit

**Quizeit** is a mobile app designed to help students enhance their understanding of formal logic concepts, including truth tables, De Morgan's law, and conditional rules. By offering interactive quizzes on various topics, Quizeit makes learning logic engaging and effective.

## Features

- **User Authentication:** Sign up, log in, and manage your profile.
- **Quiz Categories:** Take quizzes on different logic topics, such as De Morgan's Law and Truth Tables.
- **Timed Quizzes:** Complete quizzes within a time limit.
- **User Statistics:** View your quiz performance, including total quizzes completed, highest scores, and correct answer rates.
- **Admin Dashboard:** Add, delete, and manage quiz questions.

## Project Overview

The project consists of three main parts: the **Frontend** built with Flutter, the **Backend** using GoLang, and a **Task Generator** to dynamically generate quiz questions.

## Getting Started

### Prerequisites

- **Flutter SDK**: Install Flutter from [Flutter website](https://flutter.dev).
- **GoLang**: Install Go from [GoLang website](https://golang.org).
- **SQLite**: Ensure SQLite is installed for backend database operations.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/WenhaoCao00/logical-APP.git
   ```
2. Set up the backend:
   cd serverside
   go run .
3. Set up the frontend:
   cd ../frontend
   flutter pub get
   flutter devices
   flutter run -d <device-id>

Admin Access
To manage the quiz tasks, navigate to the admin interface in your browser at:
http://localhost:7999/admin/

Check all APIs
http://localhost:7999/admin/api/
