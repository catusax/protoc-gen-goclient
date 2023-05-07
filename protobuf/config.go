package protobuf

import (
	"flag"
)

var Hostname string
var SvcName string
var Namespace string
var Port string
var Options string

var Flags = new(flag.FlagSet)

func init() {
	Flags.StringVar(&Hostname, "hostname", "localhost", "host of your service")
	Flags.StringVar(&SvcName, "svc", "", "k8s service name")
	Flags.StringVar(&Namespace, "namespace", "default", "k8s namespace of your service")
	Flags.StringVar(&Port, "port", "8080", "grpc port of your service")
	Flags.StringVar(&Options, "grpc-option", "", "grpc options package where your can define your middlewares")
}
