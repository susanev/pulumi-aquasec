// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-aquasec/sdk/go/aquasec"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aquasec.NewFunctionRuntimePolicy(ctx, "functionRuntimePolicy", &aquasec.FunctionRuntimePolicyArgs{
//				ApplicationScopes: pulumi.StringArray{
//					pulumi.String("Global"),
//				},
//				BlockMaliciousExecutables: pulumi.Bool(true),
//				BlockMaliciousExecutablesAllowedProcesses: pulumi.StringArray{
//					pulumi.String("proc1"),
//					pulumi.String("proc2"),
//				},
//				BlockRunningExecutablesInTmpFolder: pulumi.Bool(true),
//				BlockedExecutables: pulumi.StringArray{
//					pulumi.String("exe1"),
//					pulumi.String("exe2"),
//				},
//				Description: pulumi.String("function_runtime_policy"),
//				Enabled:     pulumi.Bool(true),
//				Enforce:     pulumi.Bool(false),
//				ScopeVariables: FunctionRuntimePolicyScopeVariableArray{
//					&FunctionRuntimePolicyScopeVariableArgs{
//						Attribute: pulumi.String("kubernetes.cluster"),
//						Value:     pulumi.String("default"),
//					},
//					&FunctionRuntimePolicyScopeVariableArgs{
//						Attribute: pulumi.String("kubernetes.label"),
//						Name:      pulumi.String("app"),
//						Value:     pulumi.String("aqua"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FunctionRuntimePolicy struct {
	pulumi.CustomResourceState

	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayOutput `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author pulumi.StringOutput `pulumi:"author"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables pulumi.BoolPtrOutput `pulumi:"blockMaliciousExecutables"`
	// List of processes that will be allowed
	BlockMaliciousExecutablesAllowedProcesses pulumi.StringArrayOutput `pulumi:"blockMaliciousExecutablesAllowedProcesses"`
	// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
	BlockRunningExecutablesInTmpFolder pulumi.BoolPtrOutput `pulumi:"blockRunningExecutablesInTmpFolder"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables pulumi.StringArrayOutput `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey pulumi.StringPtrOutput `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns pulumi.StringArrayOutput `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey pulumi.StringPtrOutput `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName pulumi.StringPtrOutput `pulumi:"honeypotServerlessAppName"`
	// Name of the function runtime policy
	Name pulumi.StringOutput `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringOutput `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables FunctionRuntimePolicyScopeVariableArrayOutput `pulumi:"scopeVariables"`
}

// NewFunctionRuntimePolicy registers a new resource with the given unique name, arguments, and options.
func NewFunctionRuntimePolicy(ctx *pulumi.Context,
	name string, args *FunctionRuntimePolicyArgs, opts ...pulumi.ResourceOption) (*FunctionRuntimePolicy, error) {
	if args == nil {
		args = &FunctionRuntimePolicyArgs{}
	}

	if args.HoneypotSecretKey != nil {
		args.HoneypotSecretKey = pulumi.ToSecret(args.HoneypotSecretKey).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"honeypotSecretKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource FunctionRuntimePolicy
	err := ctx.RegisterResource("aquasec:index/functionRuntimePolicy:FunctionRuntimePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionRuntimePolicy gets an existing FunctionRuntimePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionRuntimePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionRuntimePolicyState, opts ...pulumi.ResourceOption) (*FunctionRuntimePolicy, error) {
	var resource FunctionRuntimePolicy
	err := ctx.ReadResource("aquasec:index/functionRuntimePolicy:FunctionRuntimePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionRuntimePolicy resources.
type functionRuntimePolicyState struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author *string `pulumi:"author"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables *bool `pulumi:"blockMaliciousExecutables"`
	// List of processes that will be allowed
	BlockMaliciousExecutablesAllowedProcesses []string `pulumi:"blockMaliciousExecutablesAllowedProcesses"`
	// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
	BlockRunningExecutablesInTmpFolder *bool `pulumi:"blockRunningExecutablesInTmpFolder"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables []string `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description *string `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce *bool `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey *string `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns []string `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey *string `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName *string `pulumi:"honeypotServerlessAppName"`
	// Name of the function runtime policy
	Name *string `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []FunctionRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
}

type FunctionRuntimePolicyState struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// Username of the account that created the service.
	Author pulumi.StringPtrInput
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables pulumi.BoolPtrInput
	// List of processes that will be allowed
	BlockMaliciousExecutablesAllowedProcesses pulumi.StringArrayInput
	// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
	BlockRunningExecutablesInTmpFolder pulumi.BoolPtrInput
	// List of executables that are prevented from running in containers.
	BlockedExecutables pulumi.StringArrayInput
	// The description of the function runtime policy
	Description pulumi.StringPtrInput
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrInput
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrInput
	// Honeypot User ID (Access Key)
	HoneypotAccessKey pulumi.StringPtrInput
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns pulumi.StringArrayInput
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey pulumi.StringPtrInput
	// Serverless application name
	HoneypotServerlessAppName pulumi.StringPtrInput
	// Name of the function runtime policy
	Name pulumi.StringPtrInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables FunctionRuntimePolicyScopeVariableArrayInput
}

func (FunctionRuntimePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionRuntimePolicyState)(nil)).Elem()
}

type functionRuntimePolicyArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables *bool `pulumi:"blockMaliciousExecutables"`
	// List of processes that will be allowed
	BlockMaliciousExecutablesAllowedProcesses []string `pulumi:"blockMaliciousExecutablesAllowedProcesses"`
	// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
	BlockRunningExecutablesInTmpFolder *bool `pulumi:"blockRunningExecutablesInTmpFolder"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables []string `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description *string `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce *bool `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey *string `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns []string `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey *string `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName *string `pulumi:"honeypotServerlessAppName"`
	// Name of the function runtime policy
	Name *string `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []FunctionRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
}

// The set of arguments for constructing a FunctionRuntimePolicy resource.
type FunctionRuntimePolicyArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables pulumi.BoolPtrInput
	// List of processes that will be allowed
	BlockMaliciousExecutablesAllowedProcesses pulumi.StringArrayInput
	// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
	BlockRunningExecutablesInTmpFolder pulumi.BoolPtrInput
	// List of executables that are prevented from running in containers.
	BlockedExecutables pulumi.StringArrayInput
	// The description of the function runtime policy
	Description pulumi.StringPtrInput
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrInput
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrInput
	// Honeypot User ID (Access Key)
	HoneypotAccessKey pulumi.StringPtrInput
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns pulumi.StringArrayInput
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey pulumi.StringPtrInput
	// Serverless application name
	HoneypotServerlessAppName pulumi.StringPtrInput
	// Name of the function runtime policy
	Name pulumi.StringPtrInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables FunctionRuntimePolicyScopeVariableArrayInput
}

func (FunctionRuntimePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionRuntimePolicyArgs)(nil)).Elem()
}

type FunctionRuntimePolicyInput interface {
	pulumi.Input

	ToFunctionRuntimePolicyOutput() FunctionRuntimePolicyOutput
	ToFunctionRuntimePolicyOutputWithContext(ctx context.Context) FunctionRuntimePolicyOutput
}

func (*FunctionRuntimePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionRuntimePolicy)(nil)).Elem()
}

func (i *FunctionRuntimePolicy) ToFunctionRuntimePolicyOutput() FunctionRuntimePolicyOutput {
	return i.ToFunctionRuntimePolicyOutputWithContext(context.Background())
}

func (i *FunctionRuntimePolicy) ToFunctionRuntimePolicyOutputWithContext(ctx context.Context) FunctionRuntimePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionRuntimePolicyOutput)
}

// FunctionRuntimePolicyArrayInput is an input type that accepts FunctionRuntimePolicyArray and FunctionRuntimePolicyArrayOutput values.
// You can construct a concrete instance of `FunctionRuntimePolicyArrayInput` via:
//
//	FunctionRuntimePolicyArray{ FunctionRuntimePolicyArgs{...} }
type FunctionRuntimePolicyArrayInput interface {
	pulumi.Input

	ToFunctionRuntimePolicyArrayOutput() FunctionRuntimePolicyArrayOutput
	ToFunctionRuntimePolicyArrayOutputWithContext(context.Context) FunctionRuntimePolicyArrayOutput
}

type FunctionRuntimePolicyArray []FunctionRuntimePolicyInput

func (FunctionRuntimePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionRuntimePolicy)(nil)).Elem()
}

func (i FunctionRuntimePolicyArray) ToFunctionRuntimePolicyArrayOutput() FunctionRuntimePolicyArrayOutput {
	return i.ToFunctionRuntimePolicyArrayOutputWithContext(context.Background())
}

func (i FunctionRuntimePolicyArray) ToFunctionRuntimePolicyArrayOutputWithContext(ctx context.Context) FunctionRuntimePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionRuntimePolicyArrayOutput)
}

// FunctionRuntimePolicyMapInput is an input type that accepts FunctionRuntimePolicyMap and FunctionRuntimePolicyMapOutput values.
// You can construct a concrete instance of `FunctionRuntimePolicyMapInput` via:
//
//	FunctionRuntimePolicyMap{ "key": FunctionRuntimePolicyArgs{...} }
type FunctionRuntimePolicyMapInput interface {
	pulumi.Input

	ToFunctionRuntimePolicyMapOutput() FunctionRuntimePolicyMapOutput
	ToFunctionRuntimePolicyMapOutputWithContext(context.Context) FunctionRuntimePolicyMapOutput
}

type FunctionRuntimePolicyMap map[string]FunctionRuntimePolicyInput

func (FunctionRuntimePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionRuntimePolicy)(nil)).Elem()
}

func (i FunctionRuntimePolicyMap) ToFunctionRuntimePolicyMapOutput() FunctionRuntimePolicyMapOutput {
	return i.ToFunctionRuntimePolicyMapOutputWithContext(context.Background())
}

func (i FunctionRuntimePolicyMap) ToFunctionRuntimePolicyMapOutputWithContext(ctx context.Context) FunctionRuntimePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionRuntimePolicyMapOutput)
}

type FunctionRuntimePolicyOutput struct{ *pulumi.OutputState }

func (FunctionRuntimePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionRuntimePolicy)(nil)).Elem()
}

func (o FunctionRuntimePolicyOutput) ToFunctionRuntimePolicyOutput() FunctionRuntimePolicyOutput {
	return o
}

func (o FunctionRuntimePolicyOutput) ToFunctionRuntimePolicyOutputWithContext(ctx context.Context) FunctionRuntimePolicyOutput {
	return o
}

// Indicates the application scope of the service.
func (o FunctionRuntimePolicyOutput) ApplicationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringArrayOutput { return v.ApplicationScopes }).(pulumi.StringArrayOutput)
}

// Username of the account that created the service.
func (o FunctionRuntimePolicyOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

// If true, prevent creation of malicious executables in functions during their runtime post invocation.
func (o FunctionRuntimePolicyOutput) BlockMaliciousExecutables() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.BoolPtrOutput { return v.BlockMaliciousExecutables }).(pulumi.BoolPtrOutput)
}

// List of processes that will be allowed
func (o FunctionRuntimePolicyOutput) BlockMaliciousExecutablesAllowedProcesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringArrayOutput {
		return v.BlockMaliciousExecutablesAllowedProcesses
	}).(pulumi.StringArrayOutput)
}

// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
func (o FunctionRuntimePolicyOutput) BlockRunningExecutablesInTmpFolder() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.BoolPtrOutput { return v.BlockRunningExecutablesInTmpFolder }).(pulumi.BoolPtrOutput)
}

// List of executables that are prevented from running in containers.
func (o FunctionRuntimePolicyOutput) BlockedExecutables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringArrayOutput { return v.BlockedExecutables }).(pulumi.StringArrayOutput)
}

// The description of the function runtime policy
func (o FunctionRuntimePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates if the runtime policy is enabled or not.
func (o FunctionRuntimePolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Indicates that policy should effect container execution (not just for audit).
func (o FunctionRuntimePolicyOutput) Enforce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.BoolPtrOutput { return v.Enforce }).(pulumi.BoolPtrOutput)
}

// Honeypot User ID (Access Key)
func (o FunctionRuntimePolicyOutput) HoneypotAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringPtrOutput { return v.HoneypotAccessKey }).(pulumi.StringPtrOutput)
}

// List of options to apply the honeypot on (Environment Vairable, Layer, File)
func (o FunctionRuntimePolicyOutput) HoneypotApplyOns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringArrayOutput { return v.HoneypotApplyOns }).(pulumi.StringArrayOutput)
}

// Honeypot User Password (Secret Key)
func (o FunctionRuntimePolicyOutput) HoneypotSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringPtrOutput { return v.HoneypotSecretKey }).(pulumi.StringPtrOutput)
}

// Serverless application name
func (o FunctionRuntimePolicyOutput) HoneypotServerlessAppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringPtrOutput { return v.HoneypotServerlessAppName }).(pulumi.StringPtrOutput)
}

// Name of the function runtime policy
func (o FunctionRuntimePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Logical expression of how to compute the dependency of the scope variables.
func (o FunctionRuntimePolicyOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) pulumi.StringOutput { return v.ScopeExpression }).(pulumi.StringOutput)
}

// List of scope attributes.
func (o FunctionRuntimePolicyOutput) ScopeVariables() FunctionRuntimePolicyScopeVariableArrayOutput {
	return o.ApplyT(func(v *FunctionRuntimePolicy) FunctionRuntimePolicyScopeVariableArrayOutput { return v.ScopeVariables }).(FunctionRuntimePolicyScopeVariableArrayOutput)
}

type FunctionRuntimePolicyArrayOutput struct{ *pulumi.OutputState }

func (FunctionRuntimePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionRuntimePolicy)(nil)).Elem()
}

func (o FunctionRuntimePolicyArrayOutput) ToFunctionRuntimePolicyArrayOutput() FunctionRuntimePolicyArrayOutput {
	return o
}

func (o FunctionRuntimePolicyArrayOutput) ToFunctionRuntimePolicyArrayOutputWithContext(ctx context.Context) FunctionRuntimePolicyArrayOutput {
	return o
}

func (o FunctionRuntimePolicyArrayOutput) Index(i pulumi.IntInput) FunctionRuntimePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionRuntimePolicy {
		return vs[0].([]*FunctionRuntimePolicy)[vs[1].(int)]
	}).(FunctionRuntimePolicyOutput)
}

type FunctionRuntimePolicyMapOutput struct{ *pulumi.OutputState }

func (FunctionRuntimePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionRuntimePolicy)(nil)).Elem()
}

func (o FunctionRuntimePolicyMapOutput) ToFunctionRuntimePolicyMapOutput() FunctionRuntimePolicyMapOutput {
	return o
}

func (o FunctionRuntimePolicyMapOutput) ToFunctionRuntimePolicyMapOutputWithContext(ctx context.Context) FunctionRuntimePolicyMapOutput {
	return o
}

func (o FunctionRuntimePolicyMapOutput) MapIndex(k pulumi.StringInput) FunctionRuntimePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionRuntimePolicy {
		return vs[0].(map[string]*FunctionRuntimePolicy)[vs[1].(string)]
	}).(FunctionRuntimePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionRuntimePolicyInput)(nil)).Elem(), &FunctionRuntimePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionRuntimePolicyArrayInput)(nil)).Elem(), FunctionRuntimePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionRuntimePolicyMapInput)(nil)).Elem(), FunctionRuntimePolicyMap{})
	pulumi.RegisterOutputType(FunctionRuntimePolicyOutput{})
	pulumi.RegisterOutputType(FunctionRuntimePolicyArrayOutput{})
	pulumi.RegisterOutputType(FunctionRuntimePolicyMapOutput{})
}
