package main

import (
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		registryNamespace, err := scaleway.NewRegistryNamespace(ctx, "test", &scaleway.RegistryNamespaceArgs{
			Name:     pulumi.String("pulumi-scaleway-dirien"),
			IsPublic: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("registryNamespace", registryNamespace.Endpoint)

		return nil
	})
}
