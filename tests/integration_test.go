// Copyright 2025 The casbin Authors. All Rights Reserved.
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
// See the License for the specific language governing permissions
// and limitations under the License.

package tests

import (
	"testing"

	"github.com/cici0602/pml-to-selinux/compiler"
	"github.com/cici0602/pml-to-selinux/models"
)

func TestBasicPolicyGeneration(t *testing.T) {
	decoded := &models.DecodedPML{
		Policies: []models.DecodedPolicy{
			{
				Policy: models.Policy{
					Subject: "webapp",
					Object:  "/var/www/html",
					Action:  "read",
					Effect:  "allow",
				},
			},
		},
	}

	gen := compiler.NewGenerator(decoded, "webapp")
	policy, err := gen.Generate()

	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if policy.ModuleName != "webapp" {
		t.Errorf("Expected module name 'webapp', got '%s'", policy.ModuleName)
	}

	if len(policy.Types) == 0 {
		t.Error("Expected types to be generated")
	}

	if len(policy.Rules) == 0 {
		t.Error("Expected rules to be generated")
	}
}
