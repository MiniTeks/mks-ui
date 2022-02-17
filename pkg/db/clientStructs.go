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

type RClient struct {
	Addr string
	Pass string
	Db   int
}

// Mksresource struct stores the count of the resources.
type Mksresource struct {
	Created   string // resource created
	Failed    string // resource failed
	Completed string // resource completed
	Deleted   string // resource deleted
}

// Application struct stores the respective mks resources
type Application struct {
	Mkstask, Mkspipelinerun, Mkstaskrun Mksresource
}
