# waitfor-mysql

[![Build Status](https://github.com/go-waitfor/waitfor-mysql/workflows/Build/badge.svg)](https://github.com/go-waitfor/waitfor-mysql/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-waitfor/waitfor-mysql)](https://goreportcard.com/report/github.com/go-waitfor/waitfor-mysql)
[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

A MySQL plugin for the [waitfor](https://github.com/go-waitfor/waitfor) library that provides MySQL database readiness assertion capabilities. This library allows you to wait for MySQL databases to become available before proceeding with your application startup or tests.

## Features

- **Simple Integration**: Easy-to-use plugin for the waitfor framework
- **Flexible URL Support**: Supports various MySQL URL formats
- **Configurable Timeouts**: Customizable retry attempts and timeouts
- **Context Awareness**: Proper context handling for cancellation and timeouts
- **Error Handling**: Detailed error reporting for connection issues

## Installation

```bash
go get github.com/go-waitfor/waitfor-mysql
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mysql"
	"os"
)

func main() {
	runner := waitfor.New(mysql.Use())

	err := runner.Test(
		context.Background(),
		[]string{"mysql://user:password@localhost:3306/mydb"},
		waitfor.WithAttempts(5),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

## Usage Examples

### Basic Usage

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mysql"
)

func main() {
	// Create a new waitfor runner with MySQL support
	runner := waitfor.New(mysql.Use())

	// Test MySQL readiness
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := runner.Test(ctx, []string{"mysql://root:password@localhost:3306/myapp"})
	if err != nil {
		log.Fatalf("MySQL is not ready: %v", err)
	}

	log.Println("MySQL is ready!")
}
```

### Advanced Configuration

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mysql"
)

func main() {
	runner := waitfor.New(mysql.Use())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	err := runner.Test(
		ctx,
		[]string{
			"mysql://user1:pass1@db1.example.com:3306/app1",
			"mysql://user2:pass2@db2.example.com:3306/app2",
		},
		waitfor.WithAttempts(10),
		waitfor.WithInterval(2*time.Second),
	)

	if err != nil {
		log.Fatalf("One or more MySQL databases are not ready: %v", err)
	}

	log.Println("All MySQL databases are ready!")
}
```

## MySQL URL Format

The library supports standard MySQL connection URLs with the following format:

```
mysql://[username[:password]@][host[:port]][/database][?parameters]
```

### Examples

- `mysql://localhost:3306` - Local MySQL on port 3306, no authentication
- `mysql://user:password@localhost/mydb` - With authentication and database
- `mysql://user:password@mysql.example.com:3306/production` - Remote MySQL with full URL
- `mysql://root@localhost:3307` - Custom port with username only

### Supported URL Components

- **username**: MySQL username (optional)
- **password**: MySQL password (optional)
- **host**: MySQL server hostname or IP address (default: localhost)
- **port**: MySQL server port (default: 3306)  
- **database**: Database name to connect to (optional)
- **parameters**: Additional MySQL driver parameters (optional)

## Configuration Options

The library supports all waitfor framework configuration options:

- `waitfor.WithAttempts(n)`: Set maximum number of connection attempts
- `waitfor.WithInterval(duration)`: Set interval between attempts
- `waitfor.WithTimeout(duration)`: Set timeout for individual attempts

## Error Handling

The library provides detailed error information for troubleshooting:

```go
import (
    "context"
    "log"
    "strings"
)

err := runner.Test(ctx, []string{"mysql://invalid-host:3306/db"})
if err != nil {
    // err will contain details about the connection failure
    log.Printf("MySQL connection failed: %v", err)
    
    // You can check for specific error types or patterns
    if strings.Contains(err.Error(), "connection refused") {
        log.Println("MySQL server is not running or not accessible")
    }
}
```

## Requirements

- Go 1.23 or later
- MySQL server (any version supported by [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql))

## Dependencies

This library depends on:
- [github.com/go-waitfor/waitfor](https://github.com/go-waitfor/waitfor) - The core waitfor framework
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## Related Projects

- [waitfor](https://github.com/go-waitfor/waitfor) - The main waitfor framework
- [waitfor-postgresql](https://github.com/go-waitfor/waitfor-postgresql) - PostgreSQL plugin
- [waitfor-redis](https://github.com/go-waitfor/waitfor-redis) - Redis plugin
