// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import "github.com/go-redis/redis/v8"

// GetValue returns the value of a given key from the redis database.
func GetValue(rClient *redis.Client, key string) string {
	val, err := rClient.Get(ctx, key).Result()
	if err != nil {
		return "0"
	}
	return val
}

// GetValues returns the keys in the Application struct using the GetValue function.
func GetValues(rClient *redis.Client) (*Application, error) {
	dataPoint := []string{"CREATED", "DELETED", "ACTIVE", "FAILED", "COMPLETED"}
	resources := []string{"MKSTASK", "MKSTASKRUN", "MKSPIPELINERUN"}
	ans := [3]Mksresource{}
	for j, r := range resources {
		resArr := [5]string{}
		for i, d := range dataPoint {
			val := GetValue(rClient, r+d)
			resArr[i] = val
		}
		ans[j] = Mksresource{
			Created:   resArr[0],
			Deleted:   resArr[1],
			Active:    resArr[2],
			Failed:    resArr[3],
			Completed: resArr[4],
		}
	}
	res := &Application{
		Mkstask:        ans[0],
		Mkstaskrun:     ans[1],
		Mkspipelinerun: ans[2],
	}
	return res, nil
}
