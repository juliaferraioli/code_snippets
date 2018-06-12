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
	"fmt"
)

type Matrix [][]float64

// New allocates a new 2-dimensional, float64 matrix.
func New(rows, cols int) Matrix {
	mat := make([][]float64, rows)
	for col := range mat {
		mat[col] = make([]float64, cols)
	}
	return mat
}

// Add sums a slice of matrices, and will return an error if the dimensions are incompatible.
func Add(matrices ...Matrix) (Matrix, error) {
	sum := New(len(matrices[0]), len(matrices[0][0]))
	for _, m := range matrices {
		if len(m) != len(sum) {
			return nil, fmt.Errorf("cannot add matrices of different dimensions")
		}
		for i, row := range m {
			if len(row) != len(sum[i]) {
				return nil, fmt.Errorf("cannot add matrices of different dimensions")
			}
			for j, col := range row {
				sum[i][j] = sum[i][j] + col
			}
		}
	}

	return sum, nil
}
