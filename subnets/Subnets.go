package subnets

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type subnet_availability struct {
	cidr string
	availability_zone string
}

func CreateSubnets(ctx *pulumi.Context, cfg *config.Config, vpc *ec2.Vpc) ([]*ec2.Subnet, []error) {

	// cidr_blocks := [4]string{"10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24"}

	new_cidr_blocks := []subnet_availability {
		{"10.0.1.0/24", "us-east-1a"},
		{"10.0.2.0/24", "us-east-1b"},
		{"10.0.3.0/24", "us-east-1c"},
		{"10.0.4.0/24", "us-east-1d"},
	}

	subnets, errors := subnets(ctx, vpc, new_cidr_blocks[:])

	return subnets, errors

}

func subnets(ctx *pulumi.Context, vpc *ec2.Vpc, cidr_blocks []subnet_availability) ([]*ec2.Subnet, []error) {

	subnets := []*ec2.Subnet{}
	err_array := []error{}

	for i := 0; i < len(cidr_blocks); i++ {

		subnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("Public Subnet %d", i+1), &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String(cidr_blocks[i].cidr),
			AvailabilityZone: pulumi.String(cidr_blocks[i].availability_zone),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(fmt.Sprintf("Public Subnet %d", i+1)),
			},
		})

		subnets = append(subnets, subnet)
		err_array = append(err_array, err)
	}

	return subnets, err_array

}
