package main

import (
	"next_kubernetes/subnets"
	vpc "next_kubernetes/vpc"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")

		// Create an AWS resource (S3 Bucket)
		// bucket, err := s3.NewBucket(ctx, "new-bucket-from-pulumi", nil)
		// if err != nil {
		// 	return err
		// }

		vpc, err := vpc.CreateVPC(ctx, cfg)

		if err != nil {
			return err
		}

		subnets.CreateSubnets(ctx, cfg, vpc)

		// Export the name of the bucket
		ctx.Export("vpc name", vpc.ID())
		return nil
	})
}
