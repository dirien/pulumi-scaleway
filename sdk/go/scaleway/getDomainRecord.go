// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a domain record.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.LookupDomainRecord(ctx, &scaleway.LookupDomainRecordArgs{
//				Data:    pulumi.StringRef("1.2.3.4"),
//				DnsZone: pulumi.StringRef("domain.tld"),
//				Name:    pulumi.StringRef("www"),
//				Type:    pulumi.StringRef("A"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupDomainRecord(ctx, &scaleway.LookupDomainRecordArgs{
//				DnsZone:  pulumi.StringRef("domain.tld"),
//				RecordId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDomainRecord(ctx *pulumi.Context, args *LookupDomainRecordArgs, opts ...pulumi.InvokeOption) (*LookupDomainRecordResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainRecordResult
	err := ctx.Invoke("scaleway:index/getDomainRecord:getDomainRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainRecord.
type LookupDomainRecordArgs struct {
	// The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
	// Cannot be used with `recordId`.
	Data *string `pulumi:"data"`
	// The IP address.
	DnsZone *string `pulumi:"dnsZone"`
	// The name of the record (can be an empty string for a root record).
	// Cannot be used with `recordId`.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The record ID.
	// Cannot be used with `name`, `type` and `data`.
	RecordId *string `pulumi:"recordId"`
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	// Cannot be used with `recordId`.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getDomainRecord.
type LookupDomainRecordResult struct {
	Data    *string `pulumi:"data"`
	DnsZone *string `pulumi:"dnsZone"`
	Fqdn    string  `pulumi:"fqdn"`
	// Dynamic record base on user geolocalisation (More information about dynamic records)
	GeoIps []GetDomainRecordGeoIp `pulumi:"geoIps"`
	// Dynamic record base on URL resolve (More information about dynamic records)
	HttpServices []GetDomainRecordHttpService `pulumi:"httpServices"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	KeepEmptyZone bool    `pulumi:"keepEmptyZone"`
	Name          *string `pulumi:"name"`
	// The priority of the record (mostly used with an `MX` record)
	Priority  int     `pulumi:"priority"`
	ProjectId *string `pulumi:"projectId"`
	RecordId  *string `pulumi:"recordId"`
	RootZone  bool    `pulumi:"rootZone"`
	// Time To Live of the record in seconds.
	Ttl  int     `pulumi:"ttl"`
	Type *string `pulumi:"type"`
	// Dynamic record based on the client’s (resolver) subnet (More information about dynamic records)
	Views []GetDomainRecordView `pulumi:"views"`
	// Dynamic record base on IP weights (More information about dynamic records)
	Weighteds []GetDomainRecordWeighted `pulumi:"weighteds"`
}

func LookupDomainRecordOutput(ctx *pulumi.Context, args LookupDomainRecordOutputArgs, opts ...pulumi.InvokeOption) LookupDomainRecordResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainRecordResult, error) {
			args := v.(LookupDomainRecordArgs)
			r, err := LookupDomainRecord(ctx, &args, opts...)
			var s LookupDomainRecordResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainRecordResultOutput)
}

// A collection of arguments for invoking getDomainRecord.
type LookupDomainRecordOutputArgs struct {
	// The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
	// Cannot be used with `recordId`.
	Data pulumi.StringPtrInput `pulumi:"data"`
	// The IP address.
	DnsZone pulumi.StringPtrInput `pulumi:"dnsZone"`
	// The name of the record (can be an empty string for a root record).
	// Cannot be used with `recordId`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The record ID.
	// Cannot be used with `name`, `type` and `data`.
	RecordId pulumi.StringPtrInput `pulumi:"recordId"`
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	// Cannot be used with `recordId`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (LookupDomainRecordOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainRecordArgs)(nil)).Elem()
}

// A collection of values returned by getDomainRecord.
type LookupDomainRecordResultOutput struct{ *pulumi.OutputState }

func (LookupDomainRecordResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainRecordResult)(nil)).Elem()
}

func (o LookupDomainRecordResultOutput) ToLookupDomainRecordResultOutput() LookupDomainRecordResultOutput {
	return o
}

func (o LookupDomainRecordResultOutput) ToLookupDomainRecordResultOutputWithContext(ctx context.Context) LookupDomainRecordResultOutput {
	return o
}

func (o LookupDomainRecordResultOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o LookupDomainRecordResultOutput) DnsZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) *string { return v.DnsZone }).(pulumi.StringPtrOutput)
}

func (o LookupDomainRecordResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Dynamic record base on user geolocalisation (More information about dynamic records)
func (o LookupDomainRecordResultOutput) GeoIps() GetDomainRecordGeoIpArrayOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) []GetDomainRecordGeoIp { return v.GeoIps }).(GetDomainRecordGeoIpArrayOutput)
}

// Dynamic record base on URL resolve (More information about dynamic records)
func (o LookupDomainRecordResultOutput) HttpServices() GetDomainRecordHttpServiceArrayOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) []GetDomainRecordHttpService { return v.HttpServices }).(GetDomainRecordHttpServiceArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDomainRecordResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainRecordResultOutput) KeepEmptyZone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) bool { return v.KeepEmptyZone }).(pulumi.BoolOutput)
}

func (o LookupDomainRecordResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The priority of the record (mostly used with an `MX` record)
func (o LookupDomainRecordResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) int { return v.Priority }).(pulumi.IntOutput)
}

func (o LookupDomainRecordResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupDomainRecordResultOutput) RecordId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) *string { return v.RecordId }).(pulumi.StringPtrOutput)
}

func (o LookupDomainRecordResultOutput) RootZone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) bool { return v.RootZone }).(pulumi.BoolOutput)
}

// Time To Live of the record in seconds.
func (o LookupDomainRecordResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o LookupDomainRecordResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Dynamic record based on the client’s (resolver) subnet (More information about dynamic records)
func (o LookupDomainRecordResultOutput) Views() GetDomainRecordViewArrayOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) []GetDomainRecordView { return v.Views }).(GetDomainRecordViewArrayOutput)
}

// Dynamic record base on IP weights (More information about dynamic records)
func (o LookupDomainRecordResultOutput) Weighteds() GetDomainRecordWeightedArrayOutput {
	return o.ApplyT(func(v LookupDomainRecordResult) []GetDomainRecordWeighted { return v.Weighteds }).(GetDomainRecordWeightedArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainRecordResultOutput{})
}
