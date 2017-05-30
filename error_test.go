// Copyright 2017 The casbin Authors. All Rights Reserved.
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

package casbin

import (
	"testing"
	"fmt"
)

func TestPathError(t *testing.T) {
	_, err := NewEnforcerSafe("", "")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		fmt.Print("Test on error: ")
		fmt.Print(err.Error())
	}
}

func TestModelError(t *testing.T) {
	_, err := NewEnforcerSafe("examples/error/error_model.conf", "examples/error/error_policy.csv")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		fmt.Print("Test on error: ")
		fmt.Print(err.Error())
	}
}

func TestPolicyError(t *testing.T) {
	_, err := NewEnforcerSafe("examples/basic_model.conf", "examples/error/error_policy.csv")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		fmt.Print("Test on error: ")
		fmt.Print(err.Error())
	}

	//_, err := e.EnforceSafe("alice", "data1", "read")
	//if err == nil {
	//	t.Errorf("Should be error here.")
	//} else {
	//	fmt.Print("Test on error: ")
	//	fmt.Print(err.Error())
	//}
}