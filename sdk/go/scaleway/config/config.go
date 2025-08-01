// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The Scaleway access key.
func GetAccessKey(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "scaleway:accessKey")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "SCW_ACCESS_KEY"); d != nil {
		value = d.(string)
	}
	return value
}

// The Scaleway API URL to use.
func GetApiUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "scaleway:apiUrl")
}

// The Scaleway organization ID.
func GetOrganizationId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "scaleway:organizationId")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_ORGANIZATION_ID"); d != nil {
		value = d.(string)
	}
	return value
}

// The Scaleway profile to use.
func GetProfile(ctx *pulumi.Context) string {
	return config.Get(ctx, "scaleway:profile")
}

// The Scaleway project ID.
func GetProjectId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "scaleway:projectId")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_PROJECT_ID"); d != nil {
		value = d.(string)
	}
	return value
}

// The region you want to attach the resource to
func GetRegion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "scaleway:region")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_REGION"); d != nil {
		value = d.(string)
	}
	return value
}

// The Scaleway secret Key.
func GetSecretKey(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "scaleway:secretKey")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "SCW_SECRET_KEY"); d != nil {
		value = d.(string)
	}
	return value
}

// The zone you want to attach the resource to
func GetZone(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "scaleway:zone")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_ZONE"); d != nil {
		value = d.(string)
	}
	return value
}
