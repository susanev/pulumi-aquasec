// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationRegistry struct {
	pulumi.CustomResourceState

	// The username of the user who created or last modified the registry
	Author pulumi.StringOutput `pulumi:"author"`
	// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
	AutoCleanup pulumi.BoolPtrOutput `pulumi:"autoCleanup"`
	// Whether to automatically pull images from the registry on creation and daily
	AutoPull pulumi.BoolPtrOutput `pulumi:"autoPull"`
	// The interval in days to start pulling new images from the registry, Defaults to 1
	AutoPullInterval pulumi.IntPtrOutput `pulumi:"autoPullInterval"`
	// Maximum number of repositories to pull every day, defaults to 100
	AutoPullMax pulumi.IntPtrOutput `pulumi:"autoPullMax"`
	// Whether to automatically pull and rescan images from the registry on creation and daily
	AutoPullRescan pulumi.BoolPtrOutput `pulumi:"autoPullRescan"`
	// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
	AutoPullTime pulumi.StringPtrOutput `pulumi:"autoPullTime"`
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition pulumi.StringOutput `pulumi:"imageCreationDateCondition"`
	// The last time the registry was modified in UNIX time
	LastUpdated pulumi.StringOutput `pulumi:"lastUpdated"`
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name    pulumi.StringOutput                  `pulumi:"name"`
	Options IntegrationRegistryOptionArrayOutput `pulumi:"options"`
	// The password for registry authentication
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// List of possible prefixes to image names pulled from the registry
	Prefixes pulumi.StringArrayOutput `pulumi:"prefixes"`
	// When auto pull image enabled, sets maximum age of auto pulled images (for example for 5 Days the value should be: 5D), Requires `imageCreationDateCondition = "imageAge"`
	PullImageAge pulumi.StringOutput `pulumi:"pullImageAge"`
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository (based on image creation date) Requires `imageCreationDateCondition = "imageCount"`
	PullImageCount pulumi.IntOutput `pulumi:"pullImageCount"`
	// List of scanner names
	ScannerNames pulumi.StringArrayOutput `pulumi:"scannerNames"`
	// The Scanner type
	ScannerType pulumi.StringOutput `pulumi:"scannerType"`
	// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
	Type pulumi.StringOutput `pulumi:"type"`
	// The URL, address or region of the registry
	Url pulumi.StringOutput `pulumi:"url"`
	// The username for registry authentication.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewIntegrationRegistry registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRegistry(ctx *pulumi.Context,
	name string, args *IntegrationRegistryArgs, opts ...pulumi.ResourceOption) (*IntegrationRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IntegrationRegistry
	err := ctx.RegisterResource("aquasec:index/integrationRegistry:IntegrationRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRegistry gets an existing IntegrationRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRegistryState, opts ...pulumi.ResourceOption) (*IntegrationRegistry, error) {
	var resource IntegrationRegistry
	err := ctx.ReadResource("aquasec:index/integrationRegistry:IntegrationRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRegistry resources.
type integrationRegistryState struct {
	// The username of the user who created or last modified the registry
	Author *string `pulumi:"author"`
	// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
	AutoCleanup *bool `pulumi:"autoCleanup"`
	// Whether to automatically pull images from the registry on creation and daily
	AutoPull *bool `pulumi:"autoPull"`
	// The interval in days to start pulling new images from the registry, Defaults to 1
	AutoPullInterval *int `pulumi:"autoPullInterval"`
	// Maximum number of repositories to pull every day, defaults to 100
	AutoPullMax *int `pulumi:"autoPullMax"`
	// Whether to automatically pull and rescan images from the registry on creation and daily
	AutoPullRescan *bool `pulumi:"autoPullRescan"`
	// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
	AutoPullTime *string `pulumi:"autoPullTime"`
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition *string `pulumi:"imageCreationDateCondition"`
	// The last time the registry was modified in UNIX time
	LastUpdated *string `pulumi:"lastUpdated"`
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name    *string                     `pulumi:"name"`
	Options []IntegrationRegistryOption `pulumi:"options"`
	// The password for registry authentication
	Password *string `pulumi:"password"`
	// List of possible prefixes to image names pulled from the registry
	Prefixes []string `pulumi:"prefixes"`
	// When auto pull image enabled, sets maximum age of auto pulled images (for example for 5 Days the value should be: 5D), Requires `imageCreationDateCondition = "imageAge"`
	PullImageAge *string `pulumi:"pullImageAge"`
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository (based on image creation date) Requires `imageCreationDateCondition = "imageCount"`
	PullImageCount *int `pulumi:"pullImageCount"`
	// List of scanner names
	ScannerNames []string `pulumi:"scannerNames"`
	// The Scanner type
	ScannerType *string `pulumi:"scannerType"`
	// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
	Type *string `pulumi:"type"`
	// The URL, address or region of the registry
	Url *string `pulumi:"url"`
	// The username for registry authentication.
	Username *string `pulumi:"username"`
}

type IntegrationRegistryState struct {
	// The username of the user who created or last modified the registry
	Author pulumi.StringPtrInput
	// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
	AutoCleanup pulumi.BoolPtrInput
	// Whether to automatically pull images from the registry on creation and daily
	AutoPull pulumi.BoolPtrInput
	// The interval in days to start pulling new images from the registry, Defaults to 1
	AutoPullInterval pulumi.IntPtrInput
	// Maximum number of repositories to pull every day, defaults to 100
	AutoPullMax pulumi.IntPtrInput
	// Whether to automatically pull and rescan images from the registry on creation and daily
	AutoPullRescan pulumi.BoolPtrInput
	// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
	AutoPullTime pulumi.StringPtrInput
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition pulumi.StringPtrInput
	// The last time the registry was modified in UNIX time
	LastUpdated pulumi.StringPtrInput
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name    pulumi.StringPtrInput
	Options IntegrationRegistryOptionArrayInput
	// The password for registry authentication
	Password pulumi.StringPtrInput
	// List of possible prefixes to image names pulled from the registry
	Prefixes pulumi.StringArrayInput
	// When auto pull image enabled, sets maximum age of auto pulled images (for example for 5 Days the value should be: 5D), Requires `imageCreationDateCondition = "imageAge"`
	PullImageAge pulumi.StringPtrInput
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository (based on image creation date) Requires `imageCreationDateCondition = "imageCount"`
	PullImageCount pulumi.IntPtrInput
	// List of scanner names
	ScannerNames pulumi.StringArrayInput
	// The Scanner type
	ScannerType pulumi.StringPtrInput
	// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
	Type pulumi.StringPtrInput
	// The URL, address or region of the registry
	Url pulumi.StringPtrInput
	// The username for registry authentication.
	Username pulumi.StringPtrInput
}

func (IntegrationRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRegistryState)(nil)).Elem()
}

type integrationRegistryArgs struct {
	// The username of the user who created or last modified the registry
	Author *string `pulumi:"author"`
	// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
	AutoCleanup *bool `pulumi:"autoCleanup"`
	// Whether to automatically pull images from the registry on creation and daily
	AutoPull *bool `pulumi:"autoPull"`
	// The interval in days to start pulling new images from the registry, Defaults to 1
	AutoPullInterval *int `pulumi:"autoPullInterval"`
	// Maximum number of repositories to pull every day, defaults to 100
	AutoPullMax *int `pulumi:"autoPullMax"`
	// Whether to automatically pull and rescan images from the registry on creation and daily
	AutoPullRescan *bool `pulumi:"autoPullRescan"`
	// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
	AutoPullTime *string `pulumi:"autoPullTime"`
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition *string `pulumi:"imageCreationDateCondition"`
	// The last time the registry was modified in UNIX time
	LastUpdated *string `pulumi:"lastUpdated"`
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name    *string                     `pulumi:"name"`
	Options []IntegrationRegistryOption `pulumi:"options"`
	// The password for registry authentication
	Password *string `pulumi:"password"`
	// List of possible prefixes to image names pulled from the registry
	Prefixes []string `pulumi:"prefixes"`
	// When auto pull image enabled, sets maximum age of auto pulled images (for example for 5 Days the value should be: 5D), Requires `imageCreationDateCondition = "imageAge"`
	PullImageAge *string `pulumi:"pullImageAge"`
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository (based on image creation date) Requires `imageCreationDateCondition = "imageCount"`
	PullImageCount *int `pulumi:"pullImageCount"`
	// List of scanner names
	ScannerNames []string `pulumi:"scannerNames"`
	// The Scanner type
	ScannerType *string `pulumi:"scannerType"`
	// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
	Type string `pulumi:"type"`
	// The URL, address or region of the registry
	Url *string `pulumi:"url"`
	// The username for registry authentication.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a IntegrationRegistry resource.
type IntegrationRegistryArgs struct {
	// The username of the user who created or last modified the registry
	Author pulumi.StringPtrInput
	// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
	AutoCleanup pulumi.BoolPtrInput
	// Whether to automatically pull images from the registry on creation and daily
	AutoPull pulumi.BoolPtrInput
	// The interval in days to start pulling new images from the registry, Defaults to 1
	AutoPullInterval pulumi.IntPtrInput
	// Maximum number of repositories to pull every day, defaults to 100
	AutoPullMax pulumi.IntPtrInput
	// Whether to automatically pull and rescan images from the registry on creation and daily
	AutoPullRescan pulumi.BoolPtrInput
	// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
	AutoPullTime pulumi.StringPtrInput
	// Additional condition for pulling and rescanning images, Defaults to 'none'
	ImageCreationDateCondition pulumi.StringPtrInput
	// The last time the registry was modified in UNIX time
	LastUpdated pulumi.StringPtrInput
	// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
	Name    pulumi.StringPtrInput
	Options IntegrationRegistryOptionArrayInput
	// The password for registry authentication
	Password pulumi.StringPtrInput
	// List of possible prefixes to image names pulled from the registry
	Prefixes pulumi.StringArrayInput
	// When auto pull image enabled, sets maximum age of auto pulled images (for example for 5 Days the value should be: 5D), Requires `imageCreationDateCondition = "imageAge"`
	PullImageAge pulumi.StringPtrInput
	// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository (based on image creation date) Requires `imageCreationDateCondition = "imageCount"`
	PullImageCount pulumi.IntPtrInput
	// List of scanner names
	ScannerNames pulumi.StringArrayInput
	// The Scanner type
	ScannerType pulumi.StringPtrInput
	// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
	Type pulumi.StringInput
	// The URL, address or region of the registry
	Url pulumi.StringPtrInput
	// The username for registry authentication.
	Username pulumi.StringPtrInput
}

func (IntegrationRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRegistryArgs)(nil)).Elem()
}

type IntegrationRegistryInput interface {
	pulumi.Input

	ToIntegrationRegistryOutput() IntegrationRegistryOutput
	ToIntegrationRegistryOutputWithContext(ctx context.Context) IntegrationRegistryOutput
}

func (*IntegrationRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRegistry)(nil)).Elem()
}

func (i *IntegrationRegistry) ToIntegrationRegistryOutput() IntegrationRegistryOutput {
	return i.ToIntegrationRegistryOutputWithContext(context.Background())
}

func (i *IntegrationRegistry) ToIntegrationRegistryOutputWithContext(ctx context.Context) IntegrationRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRegistryOutput)
}

// IntegrationRegistryArrayInput is an input type that accepts IntegrationRegistryArray and IntegrationRegistryArrayOutput values.
// You can construct a concrete instance of `IntegrationRegistryArrayInput` via:
//
//	IntegrationRegistryArray{ IntegrationRegistryArgs{...} }
type IntegrationRegistryArrayInput interface {
	pulumi.Input

	ToIntegrationRegistryArrayOutput() IntegrationRegistryArrayOutput
	ToIntegrationRegistryArrayOutputWithContext(context.Context) IntegrationRegistryArrayOutput
}

type IntegrationRegistryArray []IntegrationRegistryInput

func (IntegrationRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationRegistry)(nil)).Elem()
}

func (i IntegrationRegistryArray) ToIntegrationRegistryArrayOutput() IntegrationRegistryArrayOutput {
	return i.ToIntegrationRegistryArrayOutputWithContext(context.Background())
}

func (i IntegrationRegistryArray) ToIntegrationRegistryArrayOutputWithContext(ctx context.Context) IntegrationRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRegistryArrayOutput)
}

// IntegrationRegistryMapInput is an input type that accepts IntegrationRegistryMap and IntegrationRegistryMapOutput values.
// You can construct a concrete instance of `IntegrationRegistryMapInput` via:
//
//	IntegrationRegistryMap{ "key": IntegrationRegistryArgs{...} }
type IntegrationRegistryMapInput interface {
	pulumi.Input

	ToIntegrationRegistryMapOutput() IntegrationRegistryMapOutput
	ToIntegrationRegistryMapOutputWithContext(context.Context) IntegrationRegistryMapOutput
}

type IntegrationRegistryMap map[string]IntegrationRegistryInput

func (IntegrationRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationRegistry)(nil)).Elem()
}

func (i IntegrationRegistryMap) ToIntegrationRegistryMapOutput() IntegrationRegistryMapOutput {
	return i.ToIntegrationRegistryMapOutputWithContext(context.Background())
}

func (i IntegrationRegistryMap) ToIntegrationRegistryMapOutputWithContext(ctx context.Context) IntegrationRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRegistryMapOutput)
}

type IntegrationRegistryOutput struct{ *pulumi.OutputState }

func (IntegrationRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRegistry)(nil)).Elem()
}

func (o IntegrationRegistryOutput) ToIntegrationRegistryOutput() IntegrationRegistryOutput {
	return o
}

func (o IntegrationRegistryOutput) ToIntegrationRegistryOutputWithContext(ctx context.Context) IntegrationRegistryOutput {
	return o
}

// The username of the user who created or last modified the registry
func (o IntegrationRegistryOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

// Automatically clean up images and repositories which are no longer present in the registry from Aqua console
func (o IntegrationRegistryOutput) AutoCleanup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.BoolPtrOutput { return v.AutoCleanup }).(pulumi.BoolPtrOutput)
}

// Whether to automatically pull images from the registry on creation and daily
func (o IntegrationRegistryOutput) AutoPull() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.BoolPtrOutput { return v.AutoPull }).(pulumi.BoolPtrOutput)
}

// The interval in days to start pulling new images from the registry, Defaults to 1
func (o IntegrationRegistryOutput) AutoPullInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.IntPtrOutput { return v.AutoPullInterval }).(pulumi.IntPtrOutput)
}

// Maximum number of repositories to pull every day, defaults to 100
func (o IntegrationRegistryOutput) AutoPullMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.IntPtrOutput { return v.AutoPullMax }).(pulumi.IntPtrOutput)
}

// Whether to automatically pull and rescan images from the registry on creation and daily
func (o IntegrationRegistryOutput) AutoPullRescan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.BoolPtrOutput { return v.AutoPullRescan }).(pulumi.BoolPtrOutput)
}

// The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
func (o IntegrationRegistryOutput) AutoPullTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringPtrOutput { return v.AutoPullTime }).(pulumi.StringPtrOutput)
}

// Additional condition for pulling and rescanning images, Defaults to 'none'
func (o IntegrationRegistryOutput) ImageCreationDateCondition() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.ImageCreationDateCondition }).(pulumi.StringOutput)
}

// The last time the registry was modified in UNIX time
func (o IntegrationRegistryOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.LastUpdated }).(pulumi.StringOutput)
}

// The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
func (o IntegrationRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationRegistryOutput) Options() IntegrationRegistryOptionArrayOutput {
	return o.ApplyT(func(v *IntegrationRegistry) IntegrationRegistryOptionArrayOutput { return v.Options }).(IntegrationRegistryOptionArrayOutput)
}

// The password for registry authentication
func (o IntegrationRegistryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// List of possible prefixes to image names pulled from the registry
func (o IntegrationRegistryOutput) Prefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringArrayOutput { return v.Prefixes }).(pulumi.StringArrayOutput)
}

// When auto pull image enabled, sets maximum age of auto pulled images (for example for 5 Days the value should be: 5D), Requires `imageCreationDateCondition = "imageAge"`
func (o IntegrationRegistryOutput) PullImageAge() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.PullImageAge }).(pulumi.StringOutput)
}

// When auto pull image enabled, sets maximum age of auto pulled images tags from each repository (based on image creation date) Requires `imageCreationDateCondition = "imageCount"`
func (o IntegrationRegistryOutput) PullImageCount() pulumi.IntOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.IntOutput { return v.PullImageCount }).(pulumi.IntOutput)
}

// List of scanner names
func (o IntegrationRegistryOutput) ScannerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringArrayOutput { return v.ScannerNames }).(pulumi.StringArrayOutput)
}

// The Scanner type
func (o IntegrationRegistryOutput) ScannerType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.ScannerType }).(pulumi.StringOutput)
}

// Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
func (o IntegrationRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The URL, address or region of the registry
func (o IntegrationRegistryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The username for registry authentication.
func (o IntegrationRegistryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationRegistry) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type IntegrationRegistryArrayOutput struct{ *pulumi.OutputState }

func (IntegrationRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationRegistry)(nil)).Elem()
}

func (o IntegrationRegistryArrayOutput) ToIntegrationRegistryArrayOutput() IntegrationRegistryArrayOutput {
	return o
}

func (o IntegrationRegistryArrayOutput) ToIntegrationRegistryArrayOutputWithContext(ctx context.Context) IntegrationRegistryArrayOutput {
	return o
}

func (o IntegrationRegistryArrayOutput) Index(i pulumi.IntInput) IntegrationRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationRegistry {
		return vs[0].([]*IntegrationRegistry)[vs[1].(int)]
	}).(IntegrationRegistryOutput)
}

type IntegrationRegistryMapOutput struct{ *pulumi.OutputState }

func (IntegrationRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationRegistry)(nil)).Elem()
}

func (o IntegrationRegistryMapOutput) ToIntegrationRegistryMapOutput() IntegrationRegistryMapOutput {
	return o
}

func (o IntegrationRegistryMapOutput) ToIntegrationRegistryMapOutputWithContext(ctx context.Context) IntegrationRegistryMapOutput {
	return o
}

func (o IntegrationRegistryMapOutput) MapIndex(k pulumi.StringInput) IntegrationRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationRegistry {
		return vs[0].(map[string]*IntegrationRegistry)[vs[1].(string)]
	}).(IntegrationRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRegistryInput)(nil)).Elem(), &IntegrationRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRegistryArrayInput)(nil)).Elem(), IntegrationRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRegistryMapInput)(nil)).Elem(), IntegrationRegistryMap{})
	pulumi.RegisterOutputType(IntegrationRegistryOutput{})
	pulumi.RegisterOutputType(IntegrationRegistryArrayOutput{})
	pulumi.RegisterOutputType(IntegrationRegistryMapOutput{})
}
