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

package context

import (
	"fmt"
	"github.com/docker/compose-cli/context/store"
	"github.com/spf13/cobra"
)

func init() {
	extraCommands = append(extraCommands, createSwarmCommand)
	extraHelp = append(extraHelp, `
Create Docker Swarm context:
$ docker context create swarm CONTEXT [flags]
(see docker context create swarm --help)
`)
}

func createSwarmCommand() *cobra.Command {
	var opts descriptionCreateOpts
	cmd := &cobra.Command{
		Use:   "swarm CONTEXT [flags]",
		Short: "Create a context for Docker Swarm",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			s := store.ContextStore(cmd.Context())
			name := args[0]
			result := s.Create(
				name, store.SwarmContextType, opts.description, store.Endpoint{},
			)
			fmt.Printf("Successfully created %s context %q\n", store.SwarmContextType, name)
			return result
		},
	}
	return cmd
}

