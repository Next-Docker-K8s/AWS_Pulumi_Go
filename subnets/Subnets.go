package subnets

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func CreateSubnets(ctx *pulumi.Context, cfg *config.Config, vpc *ec2.Vpc) ([]*ec2.Subnet, []error) {

	cidr_blocks := [4]string{"10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24"}

	cidr_values := cfg.Require("subnets")

	for i := 0; i < len(cidr_values); i++ {
		fmt.Printf("Here are the values %s \n", cidr_values[i])
	}

	fmt.Println(cidr_values)

	subnets, errors := subnets(ctx, vpc, cidr_blocks[:])

	return subnets, errors

}

func subnets(ctx *pulumi.Context, vpc *ec2.Vpc, cidr_blocks []string) ([]*ec2.Subnet, []error) {

	subnets := []*ec2.Subnet{}
	err_array := []error{}

	for i := 0; i < len(cidr_blocks); i++ {

		subnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("Public Subnet %d", i+1), &ec2.SubnetArgs{
			VpcId:     vpc.ID(),
			CidrBlock: pulumi.String(cidr_blocks[i]),
			Tags: pulumi.StringMap{
				"Name": pulumi.String(fmt.Sprintf("Public Subnet %d", i+1)),
			},
		})

		subnets = append(subnets, subnet)
		err_array = append(err_array, err)
	}

	return subnets, err_array

}
