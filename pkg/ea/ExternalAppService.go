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

package ea

import (
	"github.com/devtron-labs/devtron/internal/util"
	"github.com/devtron-labs/devtron/pkg/cluster"
	"go.uber.org/zap"
)

type ExternalAppService interface {
	GetLatestDetailsForHelmApp(clusterId int, namespace string, releaseName string) error
}

type ExternalAppServiceImpl struct {
	logger         *zap.SugaredLogger
	clusterService cluster.ClusterService
	k8sUtil        *util.K8sUtil
}

func NewExternalAppServiceImpl(logger *zap.SugaredLogger, clusterService cluster.ClusterService, k8sUtil *util.K8sUtil) *ExternalAppServiceImpl {
	return &ExternalAppServiceImpl{
		logger:         logger,
		clusterService: clusterService,
		k8sUtil:        k8sUtil,
	}
}

func (impl ExternalAppServiceImpl) GetLatestDetailsForHelmApp(clusterId int, namespace string, releaseName string) error {

	clusterDetail, err := impl.clusterService.FindById(clusterId)
	if err != nil {
		return err
	}

	_, err = impl.clusterService.GetClusterConfig(clusterDetail)
	if err != nil {
		return err
	}

	return nil
}
