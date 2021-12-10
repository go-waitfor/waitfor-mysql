package mysql_test

import (
	"context"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mysql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUse(t *testing.T) {
	w := waitfor.New(mysql.Use())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := w.Test(ctx, []string{"mongodb://usr:pass@localhost/my-db"})

	assert.Error(t, err)
}
