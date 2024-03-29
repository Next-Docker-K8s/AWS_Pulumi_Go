package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func CreateVPC(ctx *pulumi.Context, cfg *config.Config) (*ec2.Vpc, error) {

	vpc, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
		CidrBlock: pulumi.String("10.0.0.0/16"),
	})

	vpc_name := cfg.Require("vpc_name")

	ec2.NewTag(
		ctx, "tag", &ec2.TagArgs{
			ResourceId: vpc.ID(),
			Key:        pulumi.String("Name"),
			Value:      pulumi.String(vpc_name),
		})

	return vpc, err
}
