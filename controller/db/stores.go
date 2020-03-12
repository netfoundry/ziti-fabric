/*
	Copyright 2020 NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package db

import (
	"github.com/netfoundry/ziti-foundation/storage/boltz"
)

type Stores struct {
	Endpoint EndpointStore
	Router   RouterStore
	Service  ServiceStore
}

type stores struct {
	endpoint *endpointStoreImpl
	router   *routerStoreImpl
	service  *serviceStoreImpl
}

func InitStores(db boltz.Db) (*Stores, error) {
	internalStores := &stores{}

	internalStores.endpoint = newEndpointStore(internalStores)
	internalStores.router = newRouterStore(internalStores)
	internalStores.service = newServiceStore(internalStores)

	stores := &Stores{
		Endpoint: internalStores.endpoint,
		Router:   internalStores.router,
		Service:  internalStores.service,
	}

	internalStores.endpoint.initializeLinked()
	internalStores.router.initializeLinked()
	internalStores.service.initializeLinked()

	mm := boltz.NewMigratorManager(db)
	if err := mm.Migrate("fabric", internalStores.migrate); err != nil {
		return nil, err
	}

	return stores, nil
}
