/* Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package matrix

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		name               string
		rows, cols         int
		wantRows, wantCols int
	}{
		{
			name:     "Zero matrix",
			rows:     0,
			cols:     0,
			wantRows: 0,
			wantCols: 0,
		},
		{
			name:     "2x2 matrix",
			rows:     2,
			cols:     2,
			wantRows: 2,
			wantCols: 2,
		},
		{
			name:     "8x3 matrix",
			rows:     8,
			cols:     3,
			wantRows: 8,
			wantCols: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.rows, tt.cols)
			if len(m) != tt.wantRows {
				t.Errorf("got %v rows, wanted %v rows", len(m), tt.wantRows)
			}
			if len(m) != 0 && len(m[0]) != tt.wantCols {
				t.Errorf("got %v columns, wanted %v columns", len(m[0]), tt.wantRows)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		name     string
		matrices []Matrix
		want     Matrix
		wantErr  bool
	}{
		{
			name: "Zero matrices should add to zero",
			matrices: []Matrix{
				Matrix{{0, 0}, {0, 0}},
				Matrix{{0, 0}, {0, 0}},
			},
			want:    Matrix{{0, 0}, {0, 0}},
			wantErr: false,
		},
		{
			name: "One matrix should add to itself",
			matrices: []Matrix{
				Matrix{{2, 3}, {4, 5}},
			},
			want:    Matrix{{2, 3}, {4, 5}},
			wantErr: false,
		},
		{
			name: "2x2 matrices should add properly",
			matrices: []Matrix{
				Matrix{{2, 6}, {1, 9}},
				Matrix{{4, 7}, {0, 3}},
			},
			want:    Matrix{{6, 13}, {1, 12}},
			wantErr: false,
		},
		{
			name: "Matrices with different numbers of columns should error",
			matrices: []Matrix{
				Matrix{{2, 6}, {1, 9}},
				Matrix{{4, 7, 5}, {0, 3, 0}},
			},
			wantErr: true,
		},
		{
			name: "Matrices with different numbers of rows should error",
			matrices: []Matrix{
				Matrix{{2, 6}, {1, 9}},
				Matrix{{4, 7}, {0, 3}, {1, 1}},
			},
			wantErr: true,
		},
		{
			name: "Matrices with float values should add properly",
			matrices: []Matrix{
				Matrix{{4.69, 8.74, -7.46}, {1.73, -4.76, -6.73}, {-5.27, 5.30, -4.43}},
				Matrix{{-1.39, -3.28, -9.67}, {-2.16, -8.18, -9.16}, {2.44, -5.99, -1.13}},
				Matrix{{-3.2, -0.4, -1.82}, {-1.12, 5.4, 5.14}, {.3, 3.67, -1.66}},
			},
			want:    Matrix{{.1, 5.06, -18.95}, {-1.55, -7.54, -10.75}, {-2.53, 2.98, -7.22}},
			wantErr: false,
		},
	}
	tolerance := .00001
	opt := cmp.Comparer(func(x, y float64) bool {
		diff := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		if math.IsNaN(diff / mean) {
			return true
		}
		return diff/mean < tolerance
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Add(tt.matrices...)

			if err != nil && !tt.wantErr {
				t.Fatalf("got error %v, wanted no error", err)
			}

			if !cmp.Equal(m, tt.want, opt) {
				t.Fatalf("got %v, wanted %v", m, tt.want)
			}
		})
	}
}
