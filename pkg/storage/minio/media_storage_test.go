package minio_storage

import (
	"context"
	"errors"
	"testing"

	"github.com/minio/minio-go/v7"
)

type fakeMinioBucket struct {
	exists      bool
	existsErr   error
	makeErr     error
	makeCalls   int
	lastBucket  string
	lastRegion  string
	existsCalls int
}

func (f *fakeMinioBucket) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	_ = ctx
	f.existsCalls++
	f.lastBucket = bucketName
	return f.exists, f.existsErr
}

func (f *fakeMinioBucket) MakeBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	_ = ctx
	f.makeCalls++
	f.lastBucket = bucketName
	f.lastRegion = opts.Region
	return f.makeErr
}

func TestEnsureBucketExists_AlreadyThere(t *testing.T) {
	f := &fakeMinioBucket{exists: true}
	err := ensureBucketExists(context.Background(), f, "my-bucket", "us-east-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.makeCalls != 0 {
		t.Fatalf("MakeBucket should not run when bucket exists, calls=%d", f.makeCalls)
	}
	if f.existsCalls != 1 {
		t.Fatalf("expected 1 BucketExists call, got %d", f.existsCalls)
	}
}

func TestEnsureBucketExists_CreatesWhenMissing(t *testing.T) {
	f := &fakeMinioBucket{exists: false}
	err := ensureBucketExists(context.Background(), f, "new-bucket", "eu-west-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.makeCalls != 1 {
		t.Fatalf("expected 1 MakeBucket call, got %d", f.makeCalls)
	}
	if f.lastBucket != "new-bucket" || f.lastRegion != "eu-west-1" {
		t.Fatalf("unexpected MakeBucket args: bucket=%q region=%q", f.lastBucket, f.lastRegion)
	}
}

func TestEnsureBucketExists_BucketExistsError(t *testing.T) {
	f := &fakeMinioBucket{existsErr: errors.New("network down")}
	err := ensureBucketExists(context.Background(), f, "b", "r")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEnsureBucketExists_MakeBucketError(t *testing.T) {
	f := &fakeMinioBucket{exists: false, makeErr: errors.New("access denied")}
	err := ensureBucketExists(context.Background(), f, "b", "r")
	if err == nil {
		t.Fatal("expected error")
	}
}
