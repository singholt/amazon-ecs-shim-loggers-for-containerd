// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"os"

	"github.com/aws/shim-loggers-for-containerd/logger"
	"github.com/aws/shim-loggers-for-containerd/logger/awslogs"

	"github.com/containerd/containerd/runtime/v2/logging"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	initCommonLogOpts()
	initDockerConfigOpts()
	initAWSLogsOpts()
}

func main() {
	pflag.Parse()
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return errors.Wrap(err, "unable to bind command line flags")
	}

	globalArgs, err := getGlobalArgs()
	if err != nil {
		return errors.Wrap(err, "unable to get global arguments")
	}

	// Set UID and/or GID of main goroutine/shim logger process if specified.
	// If you are building with go version includes the following commit, you only need
	// to call this once in main goroutine. Otherwise you need call this function in all
	// goroutines to let this syscall work properly.
	// Commit: https://github.com/golang/go/commit/d1b1145cace8b968307f9311ff611e4bb810710c
	// TODO: remove the above comment once the changes are released: https://go-review.googlesource.com/c/go/+/210639
	/*if err = logger.SetUIDAndGID(globalArgs.UID, globalArgs.GID); err != nil {
		return err
	}*/

	logDriver := globalArgs.LogDriver
	switch logDriver {
	case awslogsDriverName:
		if err := runAWSLogsDriver(globalArgs); err != nil {
			return errors.Wrap(err, "unable to run awslogs driver")
		}
	default:
		return errors.Errorf("unknown log driver: %s", logDriver)
	}

	return nil
}

func runAWSLogsDriver(globalArgs *logger.GlobalArgs) error {
	args, err := getAWSLogsArgs()
	if err != nil {
		return errors.Wrap(err, "unable to get awslogs specified arguments")
	}
	loggerArgs := awslogs.InitLogger(globalArgs, args)
	logging.Run(loggerArgs.RunLogDriver)

	return nil
}
