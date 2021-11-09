/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package restHandler

import (
	"github.com/devtron-labs/devtron/pkg/ea"
	"go.uber.org/zap"
	"net/http"
)

type ExternalAppRestHandler interface {
	GetLatestDetailsForHelmApp(w http.ResponseWriter, r *http.Request)
}

type ExternalAppRestHandlerImpl struct {
	logger *zap.SugaredLogger
	externalAppService ea.ExternalAppService
}

func NewExternalAppRestHandlerImpl(logger *zap.SugaredLogger, externalAppService ea.ExternalAppService) *ExternalAppRestHandlerImpl {
	return &ExternalAppRestHandlerImpl{
		logger: logger,
		externalAppService: externalAppService,
	}
}

func (handler ExternalAppRestHandlerImpl) GetLatestDetailsForHelmApp(w http.ResponseWriter, r *http.Request) {

	clusterId := 1
	namespace := "manish"
	releaseName := "manish-chart-release"

	handler.externalAppService.GetLatestDetailsForHelmApp(clusterId, namespace, releaseName)
}
