package files

import (
	"testing"
)

func TestOptimizePartSize(t *testing.T) {
	const MiB = 1024 * 1024
	const GiB = 1024 * MiB

	tests := []struct {
		name              string
		contentLength     int64
		explicitPartSize  int64
		wantPartSize      int64
		wantBatchSize     int
	}{
		{
			// 5 MiB / 10 MiB = 1 part; ceil(sqrt(1)) = 1
			name:          "small file 5 MiB returns defaultPartSize batch 1",
			contentLength: 5 * MiB,
			wantPartSize:  defaultPartSize,
			wantBatchSize: 1,
		},
		{
			// 500 MiB / 10 MiB = 50 parts <= 100; ceil(sqrt(50)) = 8
			name:          "500 MiB returns 10 MiB parts batch 8",
			contentLength: 500 * MiB,
			wantPartSize:  10 * MiB,
			wantBatchSize: 8,
		},
		{
			// 5 GiB: 10 MiB = 512 parts (too many), 20 MiB = 256 (too many),
			// 50 MiB = 103 (too many), 100 MiB = 52 <= 100; ceil(sqrt(52)) = 8
			name:          "5 GiB returns 100 MiB parts batch 8",
			contentLength: 5 * GiB,
			wantPartSize:  100 * MiB,
			wantBatchSize: 8,
		},
		{
			// Explicit 20 MiB for 500 MiB file: 500/20 = 25 parts; ceil(sqrt(25)) = 5
			name:             "explicit part size 20 MiB for 500 MiB file batch 5",
			contentLength:    500 * MiB,
			explicitPartSize: 20 * MiB,
			wantPartSize:     20 * MiB,
			wantBatchSize:    5,
		},
		{
			// Unknown size (0): returns defaultPartSize, batch 1
			name:          "unknown size returns defaultPartSize batch 1",
			contentLength: 0,
			wantPartSize:  defaultPartSize,
			wantBatchSize: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPartSize, gotBatchSize := optimizePartSize(tc.contentLength, tc.explicitPartSize)
			if gotPartSize != tc.wantPartSize {
				t.Errorf("optimizePartSize(%d, %d) partSize = %d, want %d",
					tc.contentLength, tc.explicitPartSize, gotPartSize, tc.wantPartSize)
			}
			if gotBatchSize != tc.wantBatchSize {
				t.Errorf("optimizePartSize(%d, %d) batchSize = %d, want %d",
					tc.contentLength, tc.explicitPartSize, gotBatchSize, tc.wantBatchSize)
			}
		})
	}
}
