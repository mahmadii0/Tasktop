# 🗓️ Tasktop — Personal Planning Web App

**Tasktop** is a web-based personal planning system built with **Golang** and **GORM (MySQL)**.  
It allows users to organize, track, and review their **daily, monthly, and yearly goals** in a simple yet powerful dashboard with helpfull reports
Also in the feature we gonna have ai assistant for it🔥
This project is designed to provide both **API-like functionality** for managing plans and a **visual web interface** for user interaction.

---

## 🚀 Features

- MVC architecture
- Create and manage **daily, monthly, and yearly** plans  
- Link daily goals to related **monthly goals**  , Link monthly goals to **annually goals**
- Track completion status (done / pending)  
- User authentication and authorization 
- Private dashboard protected by middleware  
- REST-like routing using **Gorilla Mux**  
- Database integration using **GORM ORM**  --in progress
- Modular and scalable folder structure

---

## 🧱 Project Structure

```

Tasktop/
├── configure/           # Database connection & initialization (GORM + MySQL)
├── controllers/         # Application logic and core functionality
├── middlewares/         # Authentication and access control
├── models/              # Data models (User, Goal, Plan, etc.)
├── routes/              # HTTP routing for main pages and dashboard
├── static/              # CSS, JS, and assets
├── templates/           # HTML templates for frontend rendering
├── utils/               # utils for using in different site sections
├── main.go              # Application entry point
└── go.mod               # Go module definition

````

---

## ⚙️ Requirements

Make sure you have the following installed:

- [Go 1.21+](https://go.dev/dl/)
- [MySQL 8+](https://dev.mysql.com/downloads/)
- [Git](https://git-scm.com/downloads)

---

## 🧩 Environment Configuration

Create a `.env` file in the project root with your database credentials:

```bash
DATABASE_SOURCE = "root:yourPassword@tcp(127.0.0.1:3306)/databaseName?charset=utf8&parseTime=True&loc=Local"
//The full address of templates folder
TEMPLATES_SOURCE = "C:/Users/user/Desktop/Tasktop/templates"

````

---

## 🛠️ Installation & Setup

1. **Clone the repository**

```bash
git clone https://github.com/mahmadii0/Tasktop.git
cd Tasktop
```

2. **Initialize Go modules**

```bash
go mod tidy
```

3. **Run the server**

```bash
go run main.go
```

4. **Access the app**

Open your browser and go to:

```
http://localhost:8080
```

---

## 🗄️ Database

Tasktop uses **GORM** for ORM-based database interaction.
When the app starts, `configure.CreateTables()` automatically migrates all defined models into the database.

Example connection snippet inside `configure/db.go` (conceptually):

```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
    user, pass, host, port, dbname)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

---

## 📚 API & Routing Overview

| Route                                                | Method            | Description                                              |
| ---------------------------------------------------- | ----------------- | -------------------------------------------------------- |
| `/register`                                          | GET               | Register web page                                        |
| `/register`                                          | POST              | Register a new user                                      |
| `/login`                                             | POST              | Authenticate and log in                                  |
| `/logout`                                            | GET               | Log out                                                  |
| `/dashboard`                                         | GET               | User dashboard (requires AuthMiddleware)                 |
| `/{daily/monthly/annually}-goals/{goalId:[0-9]+}`    | GET               | Read Goal(single)                                        |
| `/{daily/monthly/annually}-goals/`                   | GET               | Read Goals(user goals on declear date)                   |
| `/{daily/monthly/annually}-goals/create`             | POST              | Create Goals                                             |
| `/{daily/monthly/annually}-goals/{goalId:[0-9]+}`    | PUT               | Update Goals                                             |
| `/{daily/monthly/annually}/status/{goalId:[0-9]+}`   | GET               | Update status of Goals                                   |
| `/{daily/monthly/annually}-goals/{goalId:[0-9]+}`    | DELETE            | Delete Goals                                             |
| `/notes`                                             | GET               | Get All Notes                                            |
| `/notes/create`                                      | POST              | Create Note                                              |
| `/notes/{noteId:[0-9]}`                              | DELETE            | Delete Note                                              |
| `/report/{daily/monthly/annually}`                   | GET               | get daily/monthly/annually report                        |

> Note: In the code, **“goals”** represent **plans**, meaning your goals are the actual tasks or programs linked to your schedule.

---

## 🔒 Authentication

All `/dashboard/*` routes are protected by the `AuthMiddleware` which ensures only logged-in users can access their planning dashboard.

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -m "Add new feature"`
4. Push the branch: `git push origin feature-name`
5. Create a Pull Request on GitHub

---

## 🧑‍💻 Author

**Mohammad Ahmadi**
📍 GitHub: [@mahmadii0](https://github.com/mahmadii0)

---

## 🪪 License

This project is licensed under the **MIT License** — feel free to use and modify it for your own projects.

---

```

---

Would you like me to add **badges** (for Go version, license, etc.) and a **preview screenshot section** (for when the UI is ready)?  
It’ll make your README look even more professional on GitHub.
```
