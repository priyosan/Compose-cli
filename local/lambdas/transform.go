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
	"fmt"
	"io/ioutil"

	"github.com/docker/compose-cli/api/lambdas"

	"github.com/compose-spec/compose-go/types"
)

func TransformForLambdas(project *types.Project) error {
	functions := lambdas.GetFunctions(project)

	if len(functions) == 0 {
		return nil
	}

	for name, f := range functions {
		var dockerfile string
		switch f.Lang {
		case "javascript":
			dockerfile = `
FROM amazon/aws-lambda-nodejs:12
COPY . ./
RUN npm install
`

		default:
			return fmt.Errorf("unsupported function language %s", f.Lang)
		}

		// FIXME would be nice one can inline a Dockerfile, see https://github.com/docker/compose/issues/7095
		// FIXME can't use os.Tempdir, see https://github.com/docker/cli/issues/2249
		file, err := ioutil.TempFile(project.WorkingDir, "Dockerfile")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(file.Name(), []byte(dockerfile), 0700)
		if err != nil {
			return err
		}
		defer file.Close()

		project.Services = append(project.Services, types.ServiceConfig{
			Name: name +"_function",
			Build: &types.BuildConfig{
				Context: f.Src,
				Dockerfile: file.Name(),
			},
			Command: types.ShellCommand{f.Handler},
			PullPolicy: types.PullPolicyBuild,
			Ports: []types.ServicePortConfig{
				{
					Target: 8080,
				},
			},
			Labels: map[string]string{
				"com.docker.compose.function": name,
			},
		})
	}
	return nil
}
