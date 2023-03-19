// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFunctionResult
	err := ctx.Invoke("scaleway:index/getFunction:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunction.
type LookupFunctionArgs struct {
	FunctionId  *string `pulumi:"functionId"`
	Name        *string `pulumi:"name"`
	NamespaceId string  `pulumi:"namespaceId"`
}

// A collection of values returned by getFunction.
type LookupFunctionResult struct {
	CpuLimit             int               `pulumi:"cpuLimit"`
	Deploy               bool              `pulumi:"deploy"`
	Description          string            `pulumi:"description"`
	DomainName           string            `pulumi:"domainName"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	FunctionId           *string           `pulumi:"functionId"`
	Handler              string            `pulumi:"handler"`
	HttpOption           string            `pulumi:"httpOption"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string            `pulumi:"id"`
	MaxScale                   int               `pulumi:"maxScale"`
	MemoryLimit                int               `pulumi:"memoryLimit"`
	MinScale                   int               `pulumi:"minScale"`
	Name                       *string           `pulumi:"name"`
	NamespaceId                string            `pulumi:"namespaceId"`
	OrganizationId             string            `pulumi:"organizationId"`
	Privacy                    string            `pulumi:"privacy"`
	ProjectId                  string            `pulumi:"projectId"`
	Region                     string            `pulumi:"region"`
	Runtime                    string            `pulumi:"runtime"`
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
	Timeout                    int               `pulumi:"timeout"`
	ZipFile                    string            `pulumi:"zipFile"`
	ZipHash                    string            `pulumi:"zipHash"`
}

func LookupFunctionOutput(ctx *pulumi.Context, args LookupFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionResult, error) {
			args := v.(LookupFunctionArgs)
			r, err := LookupFunction(ctx, &args, opts...)
			var s LookupFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionResultOutput)
}

// A collection of arguments for invoking getFunction.
type LookupFunctionOutputArgs struct {
	FunctionId  pulumi.StringPtrInput `pulumi:"functionId"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	NamespaceId pulumi.StringInput    `pulumi:"namespaceId"`
}

func (LookupFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionArgs)(nil)).Elem()
}

// A collection of values returned by getFunction.
type LookupFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionResult)(nil)).Elem()
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutput() LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutputWithContext(ctx context.Context) LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) CpuLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.CpuLimit }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) Deploy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFunctionResult) bool { return v.Deploy }).(pulumi.BoolOutput)
}

func (o LookupFunctionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o LookupFunctionResultOutput) FunctionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.FunctionId }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionResultOutput) Handler() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Handler }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) HttpOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.HttpOption }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) MaxScale() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.MaxScale }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) MemoryLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.MemoryLimit }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) MinScale() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.MinScale }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) Privacy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Privacy }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Runtime }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]string { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o LookupFunctionResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.Timeout }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) ZipFile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.ZipFile }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) ZipHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.ZipHash }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionResultOutput{})
}
