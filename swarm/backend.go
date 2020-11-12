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
	"github.com/docker/compose-cli/api/compose"
	"github.com/docker/compose-cli/api/containers"
	"github.com/docker/compose-cli/api/resources"
	"github.com/docker/compose-cli/api/secrets"
	"github.com/docker/compose-cli/api/volumes"
	"github.com/docker/compose-cli/backend"
	apicontext "github.com/docker/compose-cli/context"
	"github.com/docker/compose-cli/context/cloud"
	"github.com/docker/compose-cli/context/store"
)

const backendType = store.SwarmContextType

func init() {
	backend.Register(backendType, backendType, service, cloud.NotImplementedCloudService)
}

func service(ctx context.Context) (backend.Service, error) {
	contextStore := store.ContextStore(ctx)
	currentContext := apicontext.CurrentContext(ctx)
	var context store.Endpoint

	if err := contextStore.GetEndpoint(currentContext, &context); err != nil {
		return nil, err
	}

	return &swarmAPIService{context}, nil
}

type swarmAPIService struct {
	ctx    store.Endpoint
}

func (b *swarmAPIService) ContainerService() containers.Service {
	return nil
}

func (b *swarmAPIService) ComposeService() compose.Service {
	return b
}

func (b *swarmAPIService) SecretsService() secrets.Service {
	return nil
}

func (b *swarmAPIService) VolumeService() volumes.Service {
	return nil
}

func (b *swarmAPIService) ResourceService() resources.Service {
	return nil
}



