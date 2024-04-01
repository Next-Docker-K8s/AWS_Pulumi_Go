package iamroles

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func EKSRole(ctx *pulumi.Context) (*iam.Role, error)  {

	fmt.Print("Inside EKS Role Creation")

	assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect: pulumi.StringRef("Allow"),
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type: "Service",
						Identifiers: []string{
							"eks.amazonaws.com",
						},
					},
				},
				Actions: []string{
					"sts:AssumeRole",
				},
			},
		},
	}, nil)

	EKSRole, err := iam.NewRole(ctx, "EKS_Role", &iam.RoleArgs{
		Name:             pulumi.String("EKS_Role_By_Pulumi"),
		AssumeRolePolicy: pulumi.String(assumeRole.Json),
	})

	_, err = iam.NewRolePolicyAttachment(ctx, "AmazonEKSClusterPolicy_Pulumi", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
		Role:      EKSRole.Name,
	})

	return EKSRole, err

}