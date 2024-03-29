// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Database DocumentDB Users.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			dbPassword, err := random.NewRandomPassword(ctx, "dbPassword", &random.RandomPasswordArgs{
//				Length:  pulumi.Int(16),
//				Special: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDocumentDBUser(ctx, "dbAdmin", &scaleway.DocumentDBUserArgs{
//				InstanceId: pulumi.String("11111111-1111-1111-1111-111111111111"),
//				Password:   dbPassword.Result,
//				IsAdmin:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/documentDBUser:DocumentDBUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
//
// ```
type DocumentDBUser struct {
	pulumi.CustomResourceState

	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Grant admin permissions to the Database User.
	IsAdmin pulumi.BoolPtrOutput `pulumi:"isAdmin"`
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name pulumi.StringOutput `pulumi:"name"`
	// Database User password.
	Password pulumi.StringOutput `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDocumentDBUser registers a new resource with the given unique name, arguments, and options.
func NewDocumentDBUser(ctx *pulumi.Context,
	name string, args *DocumentDBUserArgs, opts ...pulumi.ResourceOption) (*DocumentDBUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentDBUser
	err := ctx.RegisterResource("scaleway:index/documentDBUser:DocumentDBUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentDBUser gets an existing DocumentDBUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentDBUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentDBUserState, opts ...pulumi.ResourceOption) (*DocumentDBUser, error) {
	var resource DocumentDBUser
	err := ctx.ReadResource("scaleway:index/documentDBUser:DocumentDBUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentDBUser resources.
type documentDBUserState struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId *string `pulumi:"instanceId"`
	// Grant admin permissions to the Database User.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name *string `pulumi:"name"`
	// Database User password.
	Password *string `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region *string `pulumi:"region"`
}

type DocumentDBUserState struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId pulumi.StringPtrInput
	// Grant admin permissions to the Database User.
	IsAdmin pulumi.BoolPtrInput
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name pulumi.StringPtrInput
	// Database User password.
	Password pulumi.StringPtrInput
	// The Scaleway region this resource resides in.
	Region pulumi.StringPtrInput
}

func (DocumentDBUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentDBUserState)(nil)).Elem()
}

type documentDBUserArgs struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId string `pulumi:"instanceId"`
	// Grant admin permissions to the Database User.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name *string `pulumi:"name"`
	// Database User password.
	Password string `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DocumentDBUser resource.
type DocumentDBUserArgs struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId pulumi.StringInput
	// Grant admin permissions to the Database User.
	IsAdmin pulumi.BoolPtrInput
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name pulumi.StringPtrInput
	// Database User password.
	Password pulumi.StringInput
	// The Scaleway region this resource resides in.
	Region pulumi.StringPtrInput
}

func (DocumentDBUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentDBUserArgs)(nil)).Elem()
}

type DocumentDBUserInput interface {
	pulumi.Input

	ToDocumentDBUserOutput() DocumentDBUserOutput
	ToDocumentDBUserOutputWithContext(ctx context.Context) DocumentDBUserOutput
}

func (*DocumentDBUser) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentDBUser)(nil)).Elem()
}

func (i *DocumentDBUser) ToDocumentDBUserOutput() DocumentDBUserOutput {
	return i.ToDocumentDBUserOutputWithContext(context.Background())
}

func (i *DocumentDBUser) ToDocumentDBUserOutputWithContext(ctx context.Context) DocumentDBUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBUserOutput)
}

// DocumentDBUserArrayInput is an input type that accepts DocumentDBUserArray and DocumentDBUserArrayOutput values.
// You can construct a concrete instance of `DocumentDBUserArrayInput` via:
//
//	DocumentDBUserArray{ DocumentDBUserArgs{...} }
type DocumentDBUserArrayInput interface {
	pulumi.Input

	ToDocumentDBUserArrayOutput() DocumentDBUserArrayOutput
	ToDocumentDBUserArrayOutputWithContext(context.Context) DocumentDBUserArrayOutput
}

type DocumentDBUserArray []DocumentDBUserInput

func (DocumentDBUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentDBUser)(nil)).Elem()
}

func (i DocumentDBUserArray) ToDocumentDBUserArrayOutput() DocumentDBUserArrayOutput {
	return i.ToDocumentDBUserArrayOutputWithContext(context.Background())
}

func (i DocumentDBUserArray) ToDocumentDBUserArrayOutputWithContext(ctx context.Context) DocumentDBUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBUserArrayOutput)
}

// DocumentDBUserMapInput is an input type that accepts DocumentDBUserMap and DocumentDBUserMapOutput values.
// You can construct a concrete instance of `DocumentDBUserMapInput` via:
//
//	DocumentDBUserMap{ "key": DocumentDBUserArgs{...} }
type DocumentDBUserMapInput interface {
	pulumi.Input

	ToDocumentDBUserMapOutput() DocumentDBUserMapOutput
	ToDocumentDBUserMapOutputWithContext(context.Context) DocumentDBUserMapOutput
}

type DocumentDBUserMap map[string]DocumentDBUserInput

func (DocumentDBUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentDBUser)(nil)).Elem()
}

func (i DocumentDBUserMap) ToDocumentDBUserMapOutput() DocumentDBUserMapOutput {
	return i.ToDocumentDBUserMapOutputWithContext(context.Background())
}

func (i DocumentDBUserMap) ToDocumentDBUserMapOutputWithContext(ctx context.Context) DocumentDBUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBUserMapOutput)
}

type DocumentDBUserOutput struct{ *pulumi.OutputState }

func (DocumentDBUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentDBUser)(nil)).Elem()
}

func (o DocumentDBUserOutput) ToDocumentDBUserOutput() DocumentDBUserOutput {
	return o
}

func (o DocumentDBUserOutput) ToDocumentDBUserOutputWithContext(ctx context.Context) DocumentDBUserOutput {
	return o
}

// UUID of the documentDB instance.
//
// > **Important:** Updates to `instanceId` will recreate the Database User.
func (o DocumentDBUserOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBUser) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Grant admin permissions to the Database User.
func (o DocumentDBUserOutput) IsAdmin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DocumentDBUser) pulumi.BoolPtrOutput { return v.IsAdmin }).(pulumi.BoolPtrOutput)
}

// Database Username.
//
// > **Important:** Updates to `name` will recreate the Database User.
func (o DocumentDBUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Database User password.
func (o DocumentDBUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The Scaleway region this resource resides in.
func (o DocumentDBUserOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBUser) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type DocumentDBUserArrayOutput struct{ *pulumi.OutputState }

func (DocumentDBUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentDBUser)(nil)).Elem()
}

func (o DocumentDBUserArrayOutput) ToDocumentDBUserArrayOutput() DocumentDBUserArrayOutput {
	return o
}

func (o DocumentDBUserArrayOutput) ToDocumentDBUserArrayOutputWithContext(ctx context.Context) DocumentDBUserArrayOutput {
	return o
}

func (o DocumentDBUserArrayOutput) Index(i pulumi.IntInput) DocumentDBUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentDBUser {
		return vs[0].([]*DocumentDBUser)[vs[1].(int)]
	}).(DocumentDBUserOutput)
}

type DocumentDBUserMapOutput struct{ *pulumi.OutputState }

func (DocumentDBUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentDBUser)(nil)).Elem()
}

func (o DocumentDBUserMapOutput) ToDocumentDBUserMapOutput() DocumentDBUserMapOutput {
	return o
}

func (o DocumentDBUserMapOutput) ToDocumentDBUserMapOutputWithContext(ctx context.Context) DocumentDBUserMapOutput {
	return o
}

func (o DocumentDBUserMapOutput) MapIndex(k pulumi.StringInput) DocumentDBUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentDBUser {
		return vs[0].(map[string]*DocumentDBUser)[vs[1].(string)]
	}).(DocumentDBUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBUserInput)(nil)).Elem(), &DocumentDBUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBUserArrayInput)(nil)).Elem(), DocumentDBUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBUserMapInput)(nil)).Elem(), DocumentDBUserMap{})
	pulumi.RegisterOutputType(DocumentDBUserOutput{})
	pulumi.RegisterOutputType(DocumentDBUserArrayOutput{})
	pulumi.RegisterOutputType(DocumentDBUserMapOutput{})
}
