/*
Copyright 2015 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/resource"
)

func TestResourceHelpers(t *testing.T) {
	cpuLimit := resource.MustParse("10")
	memoryLimit := resource.MustParse("10G")
	resourceSpec := ResourceRequirementSpec{
		Limits: ResourceList{
			"cpu":             cpuLimit,
			"memory":          memoryLimit,
			"kube.io/storage": memoryLimit,
		},
	}
	if res := resourceSpec.Limits.Cpu(); *res != cpuLimit {
		t.Errorf("expected cpulimit %d, got %d", cpuLimit, res)
	}
	if res := resourceSpec.Limits.Memory(); *res != memoryLimit {
		t.Errorf("expected memorylimit %d, got %d", memoryLimit, res)
	}
	resourceSpec = ResourceRequirementSpec{
		Limits: ResourceList{
			"memory":          memoryLimit,
			"kube.io/storage": memoryLimit,
		},
	}
	if res := resourceSpec.Limits.Cpu(); res.Value() != 0 {
		t.Errorf("expected cpulimit %d, got %d", 0, res)
	}
	if res := resourceSpec.Limits.Memory(); *res != memoryLimit {
		t.Errorf("expected memorylimit %d, got %d", memoryLimit, res)
	}
}
