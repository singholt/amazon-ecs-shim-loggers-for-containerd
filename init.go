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
	"github.com/aws/shim-loggers-for-containerd/logger/awslogs"
	"github.com/spf13/pflag"
)

const (
	// Container options
	containerIDKey   = "container-id"
	containerNameKey = "container-name"

	// Mode and buffer size options
	modeKey          = "mode"
	maxBufferSizeKey = "max-buffer-size"

	// LogDriver options
	logDriverTypeKey  = "log-driver"
	awslogsDriverName = "awslogs"
	fluentdDriverName = "fluentd"
	splunkDriverName  = "splunk"

	// Verbose mode option
	verboseKey = "verbose"

	// UID/GID option
	uidKey = "uid"
	gidKey = "gid"

	// cleanup time option
	cleanupTimeKey = "cleanup-time"

	// docker config options
	ContainerImageIDKey   = "container-image-id"
	ContainerImageNameKey = "container-image-name"
	ContainerEnvKey       = "container-env"
	ContainerLabelsKey    = "container-labels"
)

// initCommonLogOpts initialize common options that get used by any log drivers
func initCommonLogOpts() {
	// container info
	pflag.String(containerIDKey, "", "Id of the container")
	pflag.String(containerNameKey, "", "Name of the container")

	// log driver options
	pflag.String(logDriverTypeKey, "", "`awslogs`, `fluentd` or `splunk`")

	// mode options
	pflag.String(modeKey, "", "Whether the writer is blocked or not blocked")
	pflag.String(maxBufferSizeKey, "", "The size of intermediate buffer for non-blocking mode")

	// verbose mode option
	pflag.Bool(verboseKey, false, "If set, then more logs will be printed for debugging")

	// set uid/gid option
	pflag.Int(uidKey, -1, "Customized uid for all the goroutines in shim logger process")
	pflag.Int(gidKey, -1, "Customized gid for all the goroutines in shim logger process")

	// cleanup time option
	pflag.String(cleanupTimeKey, "5s", "Cleanup time after pipes are closed, default to 5 seconds")
}

// initDockerConfigOpts initialize the docker configuration variables for the container
func initDockerConfigOpts() {
	pflag.String(ContainerImageIDKey, "", "Image id of the container")
	pflag.String(ContainerImageNameKey, "", "Image name of the container")
	pflag.String(ContainerEnvKey, "", "Environment variables of the container")
	pflag.String(ContainerLabelsKey, "", "Labels of the container")
}

// initAWSLogsOpts initialize awslogs driver specified options
func initAWSLogsOpts() {
	pflag.String(awslogs.GroupKey, "", "The CloudWatch log group to use")
	pflag.String(awslogs.RegionKey, "", "The CloudWatch region to use")
	pflag.String(awslogs.StreamKey, "", "The CloudWatch log stream to use")
	pflag.String(awslogs.CreateGroupKey, "", "Is this a new group that needs to be created?")
	pflag.String(awslogs.CredentialsEndpointKey, "", "The endpoint for iam credentials")
	pflag.String(awslogs.MultilinePatternKey, "", "Support multiline pattern for debug")
	pflag.String(awslogs.DatetimeFormatKey, "", "Multiline pattern in strftime format")
}