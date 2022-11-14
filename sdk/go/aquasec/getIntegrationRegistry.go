// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationRegistry(ctx *pulumi.Context, args *LookupIntegrationRegistryArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationRegistryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationRegistryResult
	err := ctx.Invoke("aquasec:index/getIntegrationRegistry:getIntegrationRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntegrationRegistry.
type LookupIntegrationRegistryArgs struct {
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition *string `pulumi:"imageCreationDateCondition"`
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name string `pulumi:"name"`
	// When auto pull image enabled, sets maximum age of auto pulled images
	PullImageAge *string `pulumi:"pullImageAge"`
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
	PullImageCount *int `pulumi:"pullImageCount"`
	// List of scanner names
	ScannerNames []string `pulumi:"scannerNames"`
	// Scanner type
	ScannerType *string `pulumi:"scannerType"`
}

// A collection of values returned by getIntegrationRegistry.
type LookupIntegrationRegistryResult struct {
	// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
	AutoCleanup bool `pulumi:"autoCleanup"`
	// Whether to automatically pull images from the registry on creation and daily
	AutoPull bool `pulumi:"autoPull"`
	// The interval in days to start pulling new images from the registry, Defaults to 1
	AutoPullInterval int `pulumi:"autoPullInterval"`
	// Maximum number of repositories to pull every day, defaults to 100
	AutoPullMax int `pulumi:"autoPullMax"`
	// Whether to automatically pull and rescan images from the registry on creation and daily
	AutoPullRescan bool `pulumi:"autoPullRescan"`
	// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
	AutoPullTime string `pulumi:"autoPullTime"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition string `pulumi:"imageCreationDateCondition"`
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name string `pulumi:"name"`
	// The password for registry authentication
	Password string `pulumi:"password"`
	// List of possible prefixes to image names pulled from the registry
	Prefixes []string `pulumi:"prefixes"`
	// When auto pull image enabled, sets maximum age of auto pulled images
	PullImageAge string `pulumi:"pullImageAge"`
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
	PullImageCount int `pulumi:"pullImageCount"`
	// List of scanner names
	ScannerNames []string `pulumi:"scannerNames"`
	// Scanner type
	ScannerType string `pulumi:"scannerType"`
	// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
	Type string `pulumi:"type"`
	// The URL, address or region of the registry
	Url string `pulumi:"url"`
	// The username for registry authentication.
	Username string `pulumi:"username"`
}

func LookupIntegrationRegistryOutput(ctx *pulumi.Context, args LookupIntegrationRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationRegistryResult, error) {
			args := v.(LookupIntegrationRegistryArgs)
			r, err := LookupIntegrationRegistry(ctx, &args, opts...)
			var s LookupIntegrationRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationRegistryResultOutput)
}

// A collection of arguments for invoking getIntegrationRegistry.
type LookupIntegrationRegistryOutputArgs struct {
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition pulumi.StringPtrInput `pulumi:"imageCreationDateCondition"`
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name pulumi.StringInput `pulumi:"name"`
	// When auto pull image enabled, sets maximum age of auto pulled images
	PullImageAge pulumi.StringPtrInput `pulumi:"pullImageAge"`
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
	PullImageCount pulumi.IntPtrInput `pulumi:"pullImageCount"`
	// List of scanner names
	ScannerNames pulumi.StringArrayInput `pulumi:"scannerNames"`
	// Scanner type
	ScannerType pulumi.StringPtrInput `pulumi:"scannerType"`
}

func (LookupIntegrationRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationRegistryArgs)(nil)).Elem()
}

// A collection of values returned by getIntegrationRegistry.
type LookupIntegrationRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationRegistryResult)(nil)).Elem()
}

func (o LookupIntegrationRegistryResultOutput) ToLookupIntegrationRegistryResultOutput() LookupIntegrationRegistryResultOutput {
	return o
}

func (o LookupIntegrationRegistryResultOutput) ToLookupIntegrationRegistryResultOutputWithContext(ctx context.Context) LookupIntegrationRegistryResultOutput {
	return o
}

// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
func (o LookupIntegrationRegistryResultOutput) AutoCleanup() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) bool { return v.AutoCleanup }).(pulumi.BoolOutput)
}

// Whether to automatically pull images from the registry on creation and daily
func (o LookupIntegrationRegistryResultOutput) AutoPull() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) bool { return v.AutoPull }).(pulumi.BoolOutput)
}

// The interval in days to start pulling new images from the registry, Defaults to 1
func (o LookupIntegrationRegistryResultOutput) AutoPullInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) int { return v.AutoPullInterval }).(pulumi.IntOutput)
}

// Maximum number of repositories to pull every day, defaults to 100
func (o LookupIntegrationRegistryResultOutput) AutoPullMax() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) int { return v.AutoPullMax }).(pulumi.IntOutput)
}

// Whether to automatically pull and rescan images from the registry on creation and daily
func (o LookupIntegrationRegistryResultOutput) AutoPullRescan() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) bool { return v.AutoPullRescan }).(pulumi.BoolOutput)
}

// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
func (o LookupIntegrationRegistryResultOutput) AutoPullTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.AutoPullTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIntegrationRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Additional condition for pulling and rescanning images, Defaults to 'none'
func (o LookupIntegrationRegistryResultOutput) ImageCreationDateCondition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.ImageCreationDateCondition }).(pulumi.StringOutput)
}

// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
func (o LookupIntegrationRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The password for registry authentication
func (o LookupIntegrationRegistryResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.Password }).(pulumi.StringOutput)
}

// List of possible prefixes to image names pulled from the registry
func (o LookupIntegrationRegistryResultOutput) Prefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) []string { return v.Prefixes }).(pulumi.StringArrayOutput)
}

// When auto pull image enabled, sets maximum age of auto pulled images
func (o LookupIntegrationRegistryResultOutput) PullImageAge() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.PullImageAge }).(pulumi.StringOutput)
}

// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
func (o LookupIntegrationRegistryResultOutput) PullImageCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) int { return v.PullImageCount }).(pulumi.IntOutput)
}

// List of scanner names
func (o LookupIntegrationRegistryResultOutput) ScannerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) []string { return v.ScannerNames }).(pulumi.StringArrayOutput)
}

// Scanner type
func (o LookupIntegrationRegistryResultOutput) ScannerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.ScannerType }).(pulumi.StringOutput)
}

// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
func (o LookupIntegrationRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

// The URL, address or region of the registry
func (o LookupIntegrationRegistryResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.Url }).(pulumi.StringOutput)
}

// The username for registry authentication.
func (o LookupIntegrationRegistryResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRegistryResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationRegistryResultOutput{})
}
