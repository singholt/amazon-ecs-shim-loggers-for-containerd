module github.com/aws/shim-loggers-for-containerd

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/aws/aws-sdk-go v1.26.8 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/containerd/containerd v1.4.0-beta.2.0.20200729163537-40b22ef07410
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.7.3-0.20190918143018-ad1b781e44fa
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.4.0
	github.com/fluent/fluent-logger-golang v1.4.0 // indirect
	//github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/mock v1.4.1
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-shellwords v1.0.6 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.6.1
	github.com/tinylib/msgp v1.1.1 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	gotest.tools v2.2.0+incompatible
)

replace github.com/containerd/containerd => /Users/singholt/go/src/github.com/containerd/containerd

go 1.13
