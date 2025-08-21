📝 Task CLI

A simple command-line task manager written in Go.
Easily add, list, and delete tasks right from your terminal, with tasks saved in a local JSON file for persistence.


## 🔗 Project URL
GitHub Repository: [https://github.com/kuipid01/task-cli](https://github.com/kuipid01/task-cli)

🚀 Features

➕ Add new tasks

📋 List all tasks (with ✅ for done, ❌ for pending)

❌ Delete tasks by ID

💾 Persistent storage in tasks.json

⚡ Fast, lightweight, and dependency-free

📦 Installation
Option 1: Install via go install (Requires Go Installed)
go install github.com/yourusername/task-cli@latest


Now run it globally:

task-cli add "Learn Go"
task-cli list

Option 2: Run from Source
git clone https://github.com/kuipid01/task-cli.git
cd task-cli
go run main.go add "Learn Go"
go run main.go list



List All Tasks
task-cli list


Example output:

[1] ❌ Buy groceries
[2] ❌ Learn Go

Delete a Task
task-cli delete 1


Output:

Task deleted: Buy groceries

📂 Project Structure
task-cli/
│── cmd/             # CLI commands (add, list, delete, etc.)
│── internal/
│   ├── models/      # Task model
│   └── storage/     # JSON file persistence
│── main.go          # Entry point
│── tasks.json       # Tasks are stored here

🛠️ Tech Stack

Go

Cobra
 – CLI framework

JSON file storage for persistence

🤝 Contributing

Pull requests are welcome!
For major changes, please open an issue first to discuss what you’d like to change.

📜 License

MIT
