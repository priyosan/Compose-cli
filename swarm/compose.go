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

package swarm

import (
	"context"
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose-cli/api/compose"
	"io"
)

func (b *swarmAPIService) Up(ctx context.Context, project *types.Project, detach bool) error {
	// TODO adapt code from https://github.com/docker/cli/blob/master/cli/command/stack/swarm/deploy_composefile.go#L19
	return nil
}

func (b *swarmAPIService) Down(ctx context.Context, projectName string) error {
	panic("implement me")
}

func (b *swarmAPIService) Logs(ctx context.Context, projectName string, w io.Writer) error {
	panic("implement me")
}

func (b *swarmAPIService) Ps(ctx context.Context, projectName string) ([]compose.ServiceStatus, error) {
	panic("implement me")
}

func (b *swarmAPIService) List(ctx context.Context, projectName string) ([]compose.Stack, error) {
	panic("implement me")
}

func (b *swarmAPIService) Convert(ctx context.Context, project *types.Project, format string) ([]byte, error) {
	panic("implement me")
}

