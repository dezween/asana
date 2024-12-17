# Asana Data Extractor

**Asana Data Extractor** is a microservice written in **Golang** that connects to the Asana API, retrieves users and projects, and saves their data into individual **JSON files**.

---

## 📦 **Features**

- Fetch data about **users** and **projects** from the Asana API.
- Save each user's and project's data into a separate JSON file.
- Periodically extract data at two intervals:
  - **Every 30 seconds** (fast extraction).
  - **Every 5 minutes** (standard extraction).
- Scalable using **Worker Pool** for handling multiple records.
- Error handling and graceful shutdown using **context.Context**.

---

## 🚀 **Installation and Setup**

### **1. Clone the Repository**

```bash
git clone https://github.com/yourusername/yourrepository.git
cd yourrepository

## 🚀 Install Dependencies

Ensure **Go 1.18+** is installed, then run:

```bash
go mod tidy

## 🚀 **Configuration**

Set your Asana API access token and other configurations in the config.yml file (or config/config.go).

Example config.yml:

```bash
asana_token: "Bearer YOUR_ACCESS_TOKEN"
output_folder: "./output"
requests_per_sec: 5
worker_count: 5
fetch_interval_fast: 30s
fetch_interval_slow: 5m

## ▶️ **How to Run the Project**

```bash
go run cmd/main.go

By default, the program will:

- Extract user and project data from the Asana API.
- Save each user and project as a separate JSON file into the ./output folder.

## 🧪 Running Tests

To execute the tests, use the following command:

```bash
go test ./internal/controllers -v

Test Details:
- Validate successful retrieval of users and projects.

## 📂 **Project Structure**

```bash
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/                 # Application configuration
│   ├── controllers/            # Logic for fetching data from Asana
│   ├── transport/              # HTTP client
│   ├── application/            # Core logic, worker pool, and saving data
│   └── logger/                 # Logging utilities
├── output/                     # Directory for saving JSON files
├── scripts/                    # Scripts and configurations
│   └── config.yml              # Configuration file
├── go.mod                      # Project dependencies
├── go.sum                      # Checksums for dependencies
└── README.md                   # Documentation

## 📝 **Example Output Files**
The program saves individual JSON files for each user and project with unique names:

Users:

```bash
output/user_123456789_20240702_131234.json
output/user_987654321_20240702_131235.json
Projects:

```bash
output/project_555555555_20240702_131236.json
output/project_666666666_20240702_131237.json

## ⏹ **Stopping the Application**
The program runs by default for 15 minutes and then shuts down gracefully.
To stop the program manually, press:

```bash
Ctrl + C

## 🛠 **Requirements**

Go version 1.18 or higher.
A valid Asana API access token.
