// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	appstream_sdkv1 "github.com/aws/aws-sdk-go/service/appstream"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDirectoryConfig,
			TypeName: "aws_appstream_directory_config",
		},
		{
			Factory:  ResourceFleet,
			TypeName: "aws_appstream_fleet",
			Name:     "Fleet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceFleetStackAssociation,
			TypeName: "aws_appstream_fleet_stack_association",
		},
		{
			Factory:  ResourceImageBuilder,
			TypeName: "aws_appstream_image_builder",
			Name:     "Image Builder",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceStack,
			TypeName: "aws_appstream_stack",
			Name:     "Stack",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_appstream_user",
		},
		{
			Factory:  ResourceUserStackAssociation,
			TypeName: "aws_appstream_user_stack_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppStream
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*appstream_sdkv1.AppStream, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	return appstream_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
