/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

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

package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"k8c.io/kubermatic/v2/pkg/install/crdmigration"
	kubermaticmaster "k8c.io/kubermatic/v2/pkg/install/stack/kubermatic-master"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	shutdownKubeContextFlag = cli.StringFlag{
		Name:   "kube-context",
		Usage:  "Context to use from the given kubeconfig",
		EnvVar: "KUBE_CONTEXT",
	}
	safetyFlag = cli.StringFlag{
		Name:  "stop-the-world",
		Usage: "Safety flag, must be set to 'yes' to continue",
	}
)

func ShutdownCommand(logger *logrus.Logger) cli.Command {
	return cli.Command{
		Name:   "shutdown",
		Usage:  "[CRD migration] Scales all KKP controllers on all clusters to 0, in preparation for the CRD migration",
		Action: ShutdownAction(logger),
		Flags: []cli.Flag{
			shutdownKubeContextFlag,
			safetyFlag,
		},
	}
}

func ShutdownAction(logger *logrus.Logger) cli.ActionFunc {
	return handleErrors(logger, setupLogger(logger, func(ctx *cli.Context) error {
		if strings.ToLower(ctx.String(safetyFlag.Name)) != "yes" {
			return errors.New("to prevent accidental shutdowns, the --stop-the-world flag must be set to 'yes'")
		}

		appContext := context.Background()
		namespace := kubermaticmaster.KubermaticOperatorNamespace

		// get kube client to master cluster
		kubeContext := ctx.String(shutdownKubeContextFlag.Name)

		logger.Info("Creating Kubernetes client to the master cluster…")

		kubeClient, err := getKubeClient(appContext, logger, kubeContext)
		if err != nil {
			return fmt.Errorf("failed to create Kubernetes client: %w", err)
		}

		// retrieve legacy KubermaticConfiguration (note: this is NOT defaulted, because
		// the defaulting code is only working for the new API group)
		config, err := loadLegacyKubermaticConfiguration(appContext, kubeClient, namespace)
		if err != nil {
			return fmt.Errorf("failed to retrieve KubermaticConfiguration: %w", err)
		}

		logger.Info("Retrieving Seeds…")

		allSeeds, err := getLegacySeeds(appContext, kubeClient, namespace)
		if err != nil {
			return fmt.Errorf("failed to list Seeds: %w", err)
		}

		logger.Infof("Found %d Seeds.", len(allSeeds))

		// build kube client for each seed cluster
		seedClients := map[string]ctrlruntimeclient.Client{}

		logger.Info("Creating Kubernetes client for each Seed…")

		for _, seed := range allSeeds {
			seedClient, err := getSeedClient(appContext, kubeClient, seed)
			if err != nil {
				return fmt.Errorf("failed to create Kubernetes client for Seed %q: %w", seed.Name, err)
			}

			seedClients[seed.Name] = seedClient
		}

		// assemble migration options
		opt := crdmigration.Options{
			KubermaticNamespace:     namespace,
			KubermaticConfiguration: config,
			MasterClient:            kubeClient,
			Seeds:                   allSeeds,
			SeedClients:             seedClients,
		}

		// here we go
		if err := crdmigration.ShutdownControllers(appContext, logger, &opt); err != nil {
			return fmt.Errorf("operation failed: %w", err)
		}

		logger.Info("All controllers have been scaled down to 0 replicas and webhooks have been removed now.")
		logger.Info("Please run the `migrate-crds` command now to migrate your resources.")

		return nil
	}))
}
