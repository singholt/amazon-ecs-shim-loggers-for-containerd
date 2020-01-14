package main

import (
	"github.com/aws/shim-loggers-for-containerd/logger"
	"github.com/aws/shim-loggers-for-containerd/logger/awslogs"

	"github.com/docker/go-units"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	defaultMaxBufferSize = "1m"
	blockingMode         = "blocking"
	nonBlockingMode      = "non-blocking"
)

// getGlobalArgs get arguments that used for any log drivers
func getGlobalArgs() (*logger.GlobalArgs, error) {
	containerID, err := getRequiredValue(containerIDKey)
	if err != nil {
		return nil, err
	}
	containerName, err := getRequiredValue(containerNameKey)
	if err != nil {
		return nil, err
	}
	logDriver, err := getRequiredValue(logDriverTypeKey)
	if err != nil {
		return nil, err
	}
	mode, maxBufferSize, err := getModeAndMaxBufferSize()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get value of flag %s and %s", modeKey, maxBufferSizeKey)
	}

	args := &logger.GlobalArgs{
		ContainerID:   containerID,
		ContainerName: containerName,
		LogDriver:     logDriver,
		Mode:          mode,
		MaxBufferSize: maxBufferSize,
	}

	return args, nil
}

// getAWSLogsArgs gets awslogs specified arguments for awslogs driver
func getAWSLogsArgs() (*awslogs.Args, error) {
	group, err := getRequiredValue(awslogs.GroupKey)
	if err != nil {
		return nil, err
	}
	region, err := getRequiredValue(awslogs.RegionKey)
	if err != nil {
		return nil, err
	}
	stream, err := getRequiredValue(awslogs.StreamKey)
	if err != nil {
		return nil, err
	}
	credentialsEndpoint, err := getRequiredValue(awslogs.CredentialsEndpointKey)
	if err != nil {
		return nil, err
	}

	return &awslogs.Args{
		Group:               group,
		Region:              region,
		Stream:              stream,
		CredentialsEndpoint: credentialsEndpoint,
		CreateGroup:         viper.GetString(awslogs.CreateGroupKey),
		MultilinePattern:    viper.GetString(awslogs.MultilinePatternKey),
		DatetimeFormat:      viper.GetString(awslogs.DatetimeFormatKey),
	}, nil
}

// getRequiredValue parses required arguments or exits if any is missing
func getRequiredValue(flag string) (string, error) {
	isSet := viper.IsSet(flag)
	if !isSet {
		err := errors.Errorf("%s is required", flag)
		return "", err
	}
	val := viper.GetString(flag)

	return val, nil
}

// getModeAndMaxBufferSize gets mode option and max buffer size if in blocking mode
func getModeAndMaxBufferSize() (string, int, error) {
	var (
		mode       string
		maxBufSize int
		err        error
	)

	mode = viper.GetString(modeKey)
	switch mode {
	case "":
		mode = blockingMode
	case blockingMode:
	case nonBlockingMode:
		maxBufSize, err = getMaxBufferSize()
		if err != nil {
			return "", 0, errors.Wrap(err, "unable to get max buffer size")
		}
	default:
		return "", 0, errors.Errorf("unknown mode type: %s", mode)
	}

	return mode, maxBufSize, nil
}

// getMaxBufferSize gets either customer asked buffer size or default size 1m
func getMaxBufferSize() (int, error) {
	var (
		size int64
		err  error
	)
	maxBufferSize := viper.GetString(maxBufferSizeKey)
	if maxBufferSize == "" {
		size, err = units.RAMInBytes(defaultMaxBufferSize)
	} else {
		size, err = units.RAMInBytes(maxBufferSize)
	}

	if err != nil {
		return 0, errors.Wrap(err, "unable to parse buffer size to bytes")
	}

	return int(size), nil
}
