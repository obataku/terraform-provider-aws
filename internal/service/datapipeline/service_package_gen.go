// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package datapipeline

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	datapipeline_sdkv1 "github.com/aws/aws-sdk-go/service/datapipeline"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourcePipeline,
			TypeName: "aws_datapipeline_pipeline",
		},
		{
			Factory:  DataSourcePipelineDefinition,
			TypeName: "aws_datapipeline_pipeline_definition",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourcePipeline,
			TypeName: "aws_datapipeline_pipeline",
			Name:     "Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourcePipelineDefinition,
			TypeName: "aws_datapipeline_pipeline_definition",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DataPipeline
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*datapipeline_sdkv1.DataPipeline, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	return datapipeline_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
