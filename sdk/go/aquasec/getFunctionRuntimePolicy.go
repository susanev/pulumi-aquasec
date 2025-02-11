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
//			functionRuntimePolicy, err := aquasec.LookupFunctionRuntimePolicy(ctx, &GetFunctionRuntimePolicyArgs{
//				Name: "FunctionRuntimePolicyName",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("functionRuntimePolicyDetails", functionRuntimePolicy)
//			return nil
//		})
//	}
//
// ```
func LookupFunctionRuntimePolicy(ctx *pulumi.Context, args *LookupFunctionRuntimePolicyArgs, opts ...pulumi.InvokeOption) (*LookupFunctionRuntimePolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFunctionRuntimePolicyResult
	err := ctx.Invoke("aquasec:index/getFunctionRuntimePolicy:getFunctionRuntimePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunctionRuntimePolicy.
type LookupFunctionRuntimePolicyArgs struct {
	// Name of the function runtime policy
	Name string `pulumi:"name"`
}

// A collection of values returned by getFunctionRuntimePolicy.
type LookupFunctionRuntimePolicyResult struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author string `pulumi:"author"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables bool `pulumi:"blockMaliciousExecutables"`
	// List of processes that will be allowed
	BlockMaliciousExecutablesAllowedProcesses []string `pulumi:"blockMaliciousExecutablesAllowedProcesses"`
	// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
	BlockRunningExecutablesInTmpFolder bool `pulumi:"blockRunningExecutablesInTmpFolder"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables []string `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description string `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce bool `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey string `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns []string `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey string `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName string `pulumi:"honeypotServerlessAppName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the function runtime policy
	Name string `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []GetFunctionRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
}

func LookupFunctionRuntimePolicyOutput(ctx *pulumi.Context, args LookupFunctionRuntimePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionRuntimePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionRuntimePolicyResult, error) {
			args := v.(LookupFunctionRuntimePolicyArgs)
			r, err := LookupFunctionRuntimePolicy(ctx, &args, opts...)
			var s LookupFunctionRuntimePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionRuntimePolicyResultOutput)
}

// A collection of arguments for invoking getFunctionRuntimePolicy.
type LookupFunctionRuntimePolicyOutputArgs struct {
	// Name of the function runtime policy
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupFunctionRuntimePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionRuntimePolicyArgs)(nil)).Elem()
}

// A collection of values returned by getFunctionRuntimePolicy.
type LookupFunctionRuntimePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionRuntimePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionRuntimePolicyResult)(nil)).Elem()
}

func (o LookupFunctionRuntimePolicyResultOutput) ToLookupFunctionRuntimePolicyResultOutput() LookupFunctionRuntimePolicyResultOutput {
	return o
}

func (o LookupFunctionRuntimePolicyResultOutput) ToLookupFunctionRuntimePolicyResultOutputWithContext(ctx context.Context) LookupFunctionRuntimePolicyResultOutput {
	return o
}

// Indicates the application scope of the service.
func (o LookupFunctionRuntimePolicyResultOutput) ApplicationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) []string { return v.ApplicationScopes }).(pulumi.StringArrayOutput)
}

// Username of the account that created the service.
func (o LookupFunctionRuntimePolicyResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.Author }).(pulumi.StringOutput)
}

// If true, prevent creation of malicious executables in functions during their runtime post invocation.
func (o LookupFunctionRuntimePolicyResultOutput) BlockMaliciousExecutables() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) bool { return v.BlockMaliciousExecutables }).(pulumi.BoolOutput)
}

// List of processes that will be allowed
func (o LookupFunctionRuntimePolicyResultOutput) BlockMaliciousExecutablesAllowedProcesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) []string { return v.BlockMaliciousExecutablesAllowedProcesses }).(pulumi.StringArrayOutput)
}

// If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
func (o LookupFunctionRuntimePolicyResultOutput) BlockRunningExecutablesInTmpFolder() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) bool { return v.BlockRunningExecutablesInTmpFolder }).(pulumi.BoolOutput)
}

// List of executables that are prevented from running in containers.
func (o LookupFunctionRuntimePolicyResultOutput) BlockedExecutables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) []string { return v.BlockedExecutables }).(pulumi.StringArrayOutput)
}

// The description of the function runtime policy
func (o LookupFunctionRuntimePolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Indicates if the runtime policy is enabled or not.
func (o LookupFunctionRuntimePolicyResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Indicates that policy should effect container execution (not just for audit).
func (o LookupFunctionRuntimePolicyResultOutput) Enforce() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) bool { return v.Enforce }).(pulumi.BoolOutput)
}

// Honeypot User ID (Access Key)
func (o LookupFunctionRuntimePolicyResultOutput) HoneypotAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.HoneypotAccessKey }).(pulumi.StringOutput)
}

// List of options to apply the honeypot on (Environment Vairable, Layer, File)
func (o LookupFunctionRuntimePolicyResultOutput) HoneypotApplyOns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) []string { return v.HoneypotApplyOns }).(pulumi.StringArrayOutput)
}

// Honeypot User Password (Secret Key)
func (o LookupFunctionRuntimePolicyResultOutput) HoneypotSecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.HoneypotSecretKey }).(pulumi.StringOutput)
}

// Serverless application name
func (o LookupFunctionRuntimePolicyResultOutput) HoneypotServerlessAppName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.HoneypotServerlessAppName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFunctionRuntimePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the function runtime policy
func (o LookupFunctionRuntimePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Logical expression of how to compute the dependency of the scope variables.
func (o LookupFunctionRuntimePolicyResultOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) string { return v.ScopeExpression }).(pulumi.StringOutput)
}

// List of scope attributes.
func (o LookupFunctionRuntimePolicyResultOutput) ScopeVariables() GetFunctionRuntimePolicyScopeVariableArrayOutput {
	return o.ApplyT(func(v LookupFunctionRuntimePolicyResult) []GetFunctionRuntimePolicyScopeVariable {
		return v.ScopeVariables
	}).(GetFunctionRuntimePolicyScopeVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionRuntimePolicyResultOutput{})
}
