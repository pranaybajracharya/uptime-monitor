# Uptime Monitor

Uptime Monitor is a simple Go application that periodically checks the availability of specified websites and logs their status to a database.

## Features

- Monitors multiple URLs concurrently
- Checks website availability every X seconds
- Logs uptime status (up/down) and HTTP status codes to a database
- Easy to configure and extend

## Prerequisites

Before running this application, make sure you have the following installed:

- Go (latest version recommended)
- SQLite

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/PranayBajracharya/uptime-monitor.git
   cd uptime-monitor
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Set up your database and update the connection details in `db/database.go`.

4. Run the application:

   ```
   go run main.go
   ```

## Configuration

### Modify URLs

Edit the `urls` slice in `main.go` to include the websites you want to monitor:

```
var urls = []string{
	"https://www.yourwebsite.com",
	"https://www.yourwebsite.com/path",
}
```

### Modify Check Interval

Modify the check interval in `main.go`:

```
var checkInterval = 60 * time.Second
```
