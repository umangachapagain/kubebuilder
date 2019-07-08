/*
Copyright 2018 The Kubernetes Authors.

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
	"fmt"

	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Scaffold a Kubernetes API or webhook.",
		Long:  `Scaffold a Kubernetes API or webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Coming soon.")
		},
	}
	cmd.AddCommand(
		newAPICommand(),
	)

	foundProject, version := getProjectVersion()
	// It add webhook v2 command in the following 2 cases:
	// - There are no PROJECT file found.
	// - version == 2 is found in the PROJECT file.
	if !foundProject || version == "2" {
		cmd.AddCommand(
			newWebhookV2Cmd(),
		)
	}

	return cmd
}