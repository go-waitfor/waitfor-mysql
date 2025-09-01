package mysql_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUse(t *testing.T) {
	config := mysql.Use()
	
	assert.Equal(t, []string{"mysql"}, config.Scheme)
	assert.NotNil(t, config.Factory)
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		url     *url.URL
		wantErr bool
		errType error
	}{
		{
			name:    "valid url",
			url:     &url.URL{Scheme: "mysql", Host: "localhost:3306"},
			wantErr: false,
		},
		{
			name:    "nil url",
			url:     nil,
			wantErr: true,
			errType: waitfor.ErrInvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource, err := mysql.New(tt.url)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resource)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resource)
			}
		})
	}
}

func TestMySQL_Test(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "invalid connection string",
			url:     "mysql://invalid-host:999999/nonexistent",
			wantErr: true,
		},
		{
			name:    "empty host",
			url:     "mysql:///db",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := url.Parse(tt.url)
			require.NoError(t, err)
			
			resource, err := mysql.New(parsedURL)
			require.NoError(t, err)
			
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			
			err = resource.Test(ctx)
			
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMySQL_Test_ContextCancellation(t *testing.T) {
	parsedURL, err := url.Parse("mysql://localhost:3306/test")
	require.NoError(t, err)
	
	resource, err := mysql.New(parsedURL)
	require.NoError(t, err)
	
	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	
	err = resource.Test(ctx)
	assert.Error(t, err)
}

func TestMySQL_Test_URLParsing(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectError bool
	}{
		{
			name:        "valid mysql url with all components",
			url:         "mysql://user:password@localhost:3306/database",
			expectError: true, // Will fail connection but URL parsing should work
		},
		{
			name:        "mysql url without port",
			url:         "mysql://user:password@localhost/database",
			expectError: true, // Will fail connection but URL parsing should work
		},
		{
			name:        "mysql url without database",
			url:         "mysql://user:password@localhost:3306",
			expectError: true, // Will fail connection but URL parsing should work
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := url.Parse(tt.url)
			require.NoError(t, err)
			
			resource, err := mysql.New(parsedURL)
			require.NoError(t, err)
			
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			
			err = resource.Test(ctx)
			
			if tt.expectError {
				assert.Error(t, err) // Connection should fail since no real DB
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIntegrationWithWaitfor(t *testing.T) {
	w := waitfor.New(mysql.Use())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Test with invalid MySQL URL - should fail
	err := w.Test(ctx, []string{"mysql://invalid-host:999999/nonexistent"})
	assert.Error(t, err)
}
