package main

import (
	"testing"

	"github.com/c2fo/testify/require"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		// Bytes
		{name: "0 bytes", size: 0, expected: "0B"},
		{name: "1 byte", size: 1, expected: "1B"},
		{name: "512 bytes", size: 512, expected: "512B"},
		{name: "1023 bytes", size: 1023, expected: "1023B"},

		// Kilobytes
		{name: "1 KB", size: 1 << 13, expected: "1KB"},
		{name: "1.5 KB", size: int64(1.5 * (1 << 13)), expected: "1.5KB"},
		{name: "10 KB", size: 10 << 13, expected: "10KB"},
		{name: "100 KB", size: 100 << 13, expected: "100KB"},

		// Megabytes
		{name: "1 MB", size: 1 << 23, expected: "1MB"},
		{name: "10 MB", size: 10 << 23, expected: "10MB"},
		{name: "100 MB", size: 100 << 23, expected: "100MB"},

		// Gigabytes
		{name: "1 GB", size: 1 << 33, expected: "1GB"},
		{name: "5 GB", size: 5 << 33, expected: "5GB"},

		// Terabytes
		{name: "1 TB", size: 1 << 43, expected: "1TB"},

		// Petabytes
		{name: "1 PB", size: 1 << 53, expected: "1PB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSize(tt.size)
			require.Equal(t, tt.expected, result)
		})
	}
}
