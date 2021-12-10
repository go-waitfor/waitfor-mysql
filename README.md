# waitfor-mysql
MySQL resource readiness assertion library

# Quick start

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
		[]string{"mysql://localhost:8080/my-db"},
		waitfor.WithAttempts(5),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
