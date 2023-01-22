/*
   Copyright The containerd Authors.

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

package types

// NamespaceCreateCommandOptions specifies options for `nerdctl namespace create`.
type NamespaceCreateCommandOptions struct {
	GOptions GlobalCommandOptions
	// Labels are the namespace labels
	Labels []string
}

// NamespaceUpdateCommandOptions specifies options for `nerdctl namespace update`.
type NamespaceUpdateCommandOptions NamespaceCreateCommandOptions

// NamespaceRemoveCommandOptions specifies options for `nerdctl namespace rm`.
type NamespaceRemoveCommandOptions struct {
	GOptions GlobalCommandOptions
	// CGroup delete the namespace's cgroup
	CGroup bool
}

// NamespaceInspectCommandOptions specifies options for `nerdctl namespace inspect`.
type NamespaceInspectCommandOptions struct {
	GOptions GlobalCommandOptions
	// Format the output using the given Go template, e.g, '{{json .}}'
	Format string
}
