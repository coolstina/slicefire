// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slicefire

// InStringSlice check element is in string slice
func InStringSlice(haystack []string, needle string) (exists bool) {
	for _, v := range haystack {
		if v == needle {
			exists = true
		}
	}
	return
}

// InSlice check element is in slice
func InSlice(haystack []interface{}, needle interface{}) (exists bool) {
	for _, v := range haystack {
		if v == needle {
			exists = true
		}
	}
	return
}
