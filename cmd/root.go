/*
Copyright 2023 API Testing Authors.

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
package cmd

import (
	ext "github.com/linuxsuren/api-testing/pkg/extension"
	"github.com/linuxsuren/atest-ext-runner-trpc/pkg"
	"github.com/spf13/cobra"
)

func NewRootCommand() (c *cobra.Command) {
	opt := &option{
		Extension: ext.NewExtension("trpc", "runner", 4075),
	}
	c = &cobra.Command{
		Use:   opt.GetFullName(),
		Short: "tRPC runner extension of api-testing",
		RunE:  opt.runE,
	}
	opt.AddFlags(c.Flags())
	return
}

func (o *option) runE(c *cobra.Command, args []string) (err error) {
	remoteServer := pkg.NewRemoteServer()
	err = ext.CreateExtensionRunner(o.Extension, c, remoteServer)
	return
}

type option struct {
	*ext.Extension
}
