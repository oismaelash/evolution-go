//go:build integration

package minio_storage

import (
	"os"
	"strconv"
	"testing"
)

// TestNewMinioMediaStorage_CreatesBucket runs against a real MinIO/S3-compatible endpoint.
// Example:
//
//	MINIO_INTEGRATION_TEST=1 \
//	MINIO_ENDPOINT=localhost:9000 \
//	MINIO_ACCESS_KEY=minioadmin \
//	MINIO_SECRET_KEY=minioadmin \
//	MINIO_BUCKET=evolution-test-bucket \
//	MINIO_REGION=us-east-1 \
//	MINIO_USE_SSL=false \
//	go test -tags=integration ./pkg/storage/minio/...
func TestNewMinioMediaStorage_CreatesBucket(t *testing.T) {
	if os.Getenv("MINIO_INTEGRATION_TEST") == "" {
		t.Skip("set MINIO_INTEGRATION_TEST=1 and MinIO env vars to run")
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	access := os.Getenv("MINIO_ACCESS_KEY")
	secret := os.Getenv("MINIO_SECRET_KEY")
	bucket := os.Getenv("MINIO_BUCKET")
	region := os.Getenv("MINIO_REGION")
	if region == "" {
		region = "us-east-1"
	}

	useSSL := false
	if v := os.Getenv("MINIO_USE_SSL"); v != "" {
		var err error
		useSSL, err = strconv.ParseBool(v)
		if err != nil {
			t.Fatalf("MINIO_USE_SSL: %v", err)
		}
	}

	if endpoint == "" || access == "" || secret == "" || bucket == "" {
		t.Fatal("MINIO_ENDPOINT, MINIO_ACCESS_KEY, MINIO_SECRET_KEY, MINIO_BUCKET are required")
	}

	st, err := NewMinioMediaStorage(endpoint, access, secret, bucket, region, useSSL)
	if err != nil {
		t.Fatalf("NewMinioMediaStorage: %v", err)
	}
	if st == nil {
		t.Fatal("expected non-nil storage")
	}
}
