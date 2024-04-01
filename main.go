package main

import (
	"next_kubernetes/eks"
	"next_kubernetes/iamroles"
	"next_kubernetes/subnets"
	"next_kubernetes/vpc"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")

		vpc, err := vpc.CreateVPC(ctx, cfg)

		if err != nil {
			return err
		}

		subnets, errors := subnets.CreateSubnets(ctx, cfg, vpc)

		if errors[0] != nil{
			return errors[0]
		}

		eksRole, err := iamroles.EKSRole(ctx)

		if err != nil{
			return err
		}

		ec2NodeRole, err := iamroles.CreateEC2Role(ctx)

		if err != nil {
			return err
		}

		eks_cluster, err := eks.CreateEKS(ctx, vpc, subnets, eksRole, ec2NodeRole)

		if err != nil {
			return err
		}

		ctx.Export("vpc name", vpc.ID())
		ctx.Export("First Subnet", subnets[0].ID())
		ctx.Export("EKS Name", eks_cluster.Name)

		
		return nil

	})
}
