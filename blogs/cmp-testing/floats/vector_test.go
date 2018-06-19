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

package vector

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		name    string
		vectors []Vector
		want    Vector
		wantErr bool
	}{
		{
			name: "Zero vectors should add to zero",
			vectors: []Vector{
				Vector{0, 0, 0, 0},
				Vector{0, 0, 0, 0},
			},
			want:    Vector{0, 0, 0, 0},
			wantErr: false,
		},
		{
			name: "One vector should add to itself",
			vectors: []Vector{
				Vector{2, 3, 4, 5},
			},
			want:    Vector{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "Two vectors should add properly",
			vectors: []Vector{
				Vector{2, 6, 1, 9},
				Vector{4, 7, 0, 3},
			},
			want:    Vector{6, 13, 1, 12},
			wantErr: false,
		},
		{
			name: "Vectors with different numbers of features should error",
			vectors: []Vector{
				Vector{2, 6, 1, 9},
				Vector{4, 7, 5, 0, 3, 0},
			},
			wantErr: true,
		},
		{
			name: "Vectors with float values should add properly",
			vectors: []Vector{
				Vector{4.69, 8.74, -7.46, 1.73, -4.76, -6.73, -5.27, 5.30, -4.43},
				Vector{-1.39, -3.28, -9.67, -2.16, -8.18, -9.16, 2.44, -5.99, -1.13},
				Vector{-3.2, -0.4, -1.82, -1.12, 5.4, 5.14, .3, 3.67, -1.66},
			},
			want:    Vector{.1, 5.06, -18.95, -1.55, -7.54, -10.75, -2.53, 2.98, -7.22},
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
			m, err := Add(tt.vectors...)

			if err != nil && !tt.wantErr {
				t.Fatalf("got error %v, wanted no error", err)
			}

			if !cmp.Equal(m, tt.want, opt) {
				t.Fatalf("got %v, wanted %v", m, tt.want)
			}
		})
	}
}
