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
	"fmt"
)

type Vector []float64

// New allocates a new float64 vector.
func New(numFeatures int) Vector {
	return make([]float64, numFeatures)
}

// Add sums a slice of vectors, and will return an error if the dimensions are incompatible.
func Add(vectors ...Vector) (Vector, error) {
	sum := New(len(vectors[0]))
	for _, v := range vectors {
		if len(v) != len(sum) {
			return nil, fmt.Errorf("cannot add vectors of different dimensions")
		}
		for i, f := range v {
			sum[i] += f
		}
	}

	return sum, nil
}
