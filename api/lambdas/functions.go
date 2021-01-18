/*
   Copyright 2020 Docker Compose CLI authors

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

package lambdas

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/mitchellh/mapstructure"
)

type Function struct {
	Name     string `yaml:",omitempty" json:"name,omitempty"`
	Lang     string `yaml:",omitempty" json:"lang,omitempty"`
	Src      string `yaml:",omitempty" json:"src,omitempty"`
	Handler      string `yaml:",omitempty" json:"handler,omitempty"`
	Consume  []string `yaml:",omitempty" json:"consume,omitempty"`
	Produce  []string `yaml:",omitempty" json:"produce,omitempty"`
}

type Functions map[string]Function

// TODO to be moved to compose-go loader.Load once compose-spec has approved `functions` design
func LoadFunctions(project *types.Project) error {
	x, ok := project.Extensions["x-functions"]
	if !ok {
		return nil
	}

	functions := Functions{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &functions,
	})
	if err != nil {
		return err
	}
	err = decoder.Decode(x)
	if err != nil {
		return err
	}
	project.Extensions["x-functions"] = functions
	return nil
}

// GetQueues returns project's queues.
// TODO replace use of this with direct access `project.Queues` once compose-spec approvec `queues` concept
func GetFunctions(project *types.Project) Functions {
	x, ok := project.Extensions["x-functions"]
	if !ok {
		return nil
	}
	return x.(Functions)
}
