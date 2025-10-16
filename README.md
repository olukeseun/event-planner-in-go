# 🗓️ Event Planner App

An **Event Planner Web API** built with **Go (Gin Framework)** that allows users to create, manage, and join events.  
The system supports **user authentication**, **role-based access**, and **event scheduling**, making it ideal for individuals, teams, or organizations planning meetups, conferences, or social gatherings.

---

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database:** PostgreSQL (or MySQL)
- **ORM:** GORM
- **Authentication:** JWT (JSON Web Tokens)
- **Configuration:** GoDotEnv
- **API Documentation:** Swagger (optional)
- **Containerization:** Docker (optional)

---

## 📦 Features

- 👤 **User Authentication** (Sign up, Login, Logout)
- 🎟️ **Create, Update, Delete Events**
- 🗓️ **View Upcoming Events**
- 🤝 **Join or Leave Events**
- 🧑‍💼 **Role-based Access Control** (Organizer / Attendee)
- 📩 **Event Invitations**
- 🔔 **Email or Push Notifications (optional)**

---

## ⚙️ Project Structure

event-planner-app/
│
├── cmd/
│ └── main.go # Entry point of the application
│
├── config/
│ └── config.go # Loads environment variables and DB config
│
├── controllers/
│ ├── auth_controller.go # Handles login/signup logic
│ ├── event_controller.go # Handles CRUD for events
│ └── user_controller.go # Handles user profiles
│
├── middleware/
│ ├── auth_middleware.go # JWT authentication middleware
│ └── role_middleware.go # Role-based access control
│
├── models/
│ ├── user.go # User model
│ ├── event.go # Event model
│ └── participant.go # Event participant model
│
├── routes/
│ └── routes.go # All application routes
│
├── services/
│ ├── auth_service.go # Business logic for authentication
│ ├── event_service.go # Business logic for events
│ └── jwt_service.go # Token generation/validation
│
├── utils/
│ └── response.go # Standardized API response helper
│
├── .env # Environment variables
├── go.mod # Go module definition
├── go.sum # Dependency lock file
└── README.md # Project documentation


---

## ⚙️ Setup Instructions

### 1️⃣ Clone Repository

```bash
git clone https://github.com/olukeseun/event-planner-app.git
cd event-planner-app

