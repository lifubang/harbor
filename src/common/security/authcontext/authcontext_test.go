// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authcontext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSysAdmin(t *testing.T) {
	// nil roles
	ctx := &AuthContext{}
	assert.False(t, ctx.IsSysAdmin())

	// has no admin role
	ctx = &AuthContext{
		Roles: []string{projectAdminRole, developerRole, guestRole},
	}
	assert.False(t, ctx.IsSysAdmin())

	// has admin role
	ctx = &AuthContext{
		Roles: []string{sysAdminRole},
	}
	assert.True(t, ctx.IsSysAdmin())
}

func TestGetProjectRoles(t *testing.T) {
	ctx := &AuthContext{
		Projects: []*project{
			&project{
				Name:  "project",
				Roles: []string{projectAdminRole, developerRole, guestRole},
			},
		},
	}

	// test with name
	roles := ctx.GetProjectRoles("project")
	assert.Equal(t, 3, len(roles))

	// TODO add test case with ID
}

func TestGetMyProjects(t *testing.T) {
	ctx := &AuthContext{
		Projects: []*project{
			&project{
				Name:  "project1",
				Roles: []string{projectAdminRole},
			},
			&project{
				Name:  "project2",
				Roles: []string{developerRole},
			},
		},
	}

	projects := ctx.GetMyProjects()
	assert.Equal(t, 2, len(projects))
}
