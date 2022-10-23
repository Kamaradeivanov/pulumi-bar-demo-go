package main

import (
	"fmt"

	scaleway "github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cluster, err := scaleway.NewKubernetesCluster(ctx, ctx.Stack(), &scaleway.KubernetesClusterArgs{
			AutoUpgrade: &scaleway.KubernetesClusterAutoUpgradeArgs{
				Enable:                     pulumi.Bool(false),
				MaintenanceWindowDay:       pulumi.String("monday"),
				MaintenanceWindowStartHour: pulumi.Int(02),
			},
			AutoscalerConfig: &scaleway.KubernetesClusterAutoscalerConfigArgs{
				ScaleDownDelayAfterAdd: pulumi.String("10m"),
			},
			Type: pulumi.String("kapsule"),
			Tags: pulumi.StringArray{
				pulumi.String("go"),
			},
			Cni:     pulumi.String("cilium"),
			Version: pulumi.String("1.24.5"),
		})
		if err != nil {
			return fmt.Errorf("error creating kubernetes cluster: %v", err)
		}

		ctx.Export("cluster", cluster.Kubeconfigs)

		return nil
	})
}
