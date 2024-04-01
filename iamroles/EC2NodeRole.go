package iamroles

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateEC2Role(ctx *pulumi.Context) (*iam.Role, error)  {
	
	assumeRolePolicyDocument := map[string]interface{}{
        "Version": "2012-10-17",
        "Statement": []map[string]interface{}{
            {
                "Effect": "Allow",
                "Action": "sts:AssumeRole",
                "Principal": map[string]interface{}{
                    "Service": "ec2.amazonaws.com",
                },
            },
        },
    }

	

	// policyDocument := map[string]interface{}{
	// 	"Version": "2012-10-17",
	// 	"Statement": []map[string]interface{}{
	// 		{
	// 			"Effect": "Allow",
	// 			"Action": "sts:AssumeRole",
	// 			"Resource": "*",
	// 			"Principal": map[string]interface{}{
	// 				"Service": "ec2.amazonaws.com",
	// 			},
	// 		},
	// 	},
	// }

	// json0, _ := json.Marshal(policyDocument)

	assumeRolePolicyJSON, _ := json.Marshal(assumeRolePolicyDocument)

	ec2NodeRole, err := iam.NewRole(ctx, "ec2_node", &iam.RoleArgs{
		Name: pulumi.String("EKS-Node-Role"),
		AssumeRolePolicy: pulumi.String(string(assumeRolePolicyJSON)),
	})

	iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSWorkerNodePolicy", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
		Role:      ec2NodeRole.Name,
	})

	iam.NewRolePolicyAttachment(ctx, "example-AmazonEKS_CNI_Policy", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"),
		Role:      ec2NodeRole.Name,
	})

	// _, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEC2ContainerRegistryReadOnly", &iam.RolePolicyAttachmentArgs{
	// 	PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
	// 	Role:      example.Name,
	// })

	return ec2NodeRole, err

}