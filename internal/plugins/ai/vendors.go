package ai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

func NewVendorsManager() *VendorsManager {
	return &VendorsManager{
		Vendors:       []Vendor{},
		VendorsByName: map[string]Vendor{},
	}
}

type VendorsManager struct {
	*plugins.PluginBase
	Vendors       []Vendor
	VendorsByName map[string]Vendor
	Models        *VendorsModels
	modelsMu      sync.Mutex
}

// AddVendors registers one or more vendors with the manager.
// Vendors are stored with lowercase keys to enable case-insensitive lookup.
// Membership changes invalidate the model catalog.
func (o *VendorsManager) AddVendors(vendors ...Vendor) {
	o.modelsMu.Lock()
	defer o.modelsMu.Unlock()
	for _, vendor := range vendors {
		name := strings.ToLower(vendor.GetName())
		o.VendorsByName[name] = vendor
		o.Vendors = append(o.Vendors, vendor)
	}
	o.Models = nil
}

func (o *VendorsManager) Clear() {
	o.modelsMu.Lock()
	defer o.modelsMu.Unlock()
	o.VendorsByName = map[string]Vendor{}
	o.Vendors = []Vendor{}
	o.Models = nil
}

func (o *VendorsManager) SetupFillEnvFileContent(envFileContent *bytes.Buffer) {
	for _, vendor := range o.Vendors {
		vendor.SetupFillEnvFileContent(envFileContent)
	}
}

func (o *VendorsManager) GetModels() (ret *VendorsModels, err error) {
	// Many goroutines can call GetModels, for example from the REST server.
	// The lock makes sure that only one call reads the models.
	o.modelsMu.Lock()
	defer o.modelsMu.Unlock()
	if o.Models != nil {
		return o.Models, nil
	}
	// Do not cache a failed discovery. The next call tries the vendors again.
	if ret, err = o.readModels(); err == nil {
		o.Models = ret
	}
	return
}

func (o *VendorsManager) Configure() (err error) {
	for _, vendor := range o.Vendors {
		_ = vendor.Configure()
	}
	return
}

func (o *VendorsManager) HasVendors() bool {
	return len(o.Vendors) > 0
}

// FindByName returns a vendor by name. Lookup is case-insensitive.
// For example, "OpenAI", "openai", and "OPENAI" all match the same vendor.
func (o *VendorsManager) FindByName(name string) Vendor {
	return o.VendorsByName[strings.ToLower(name)]
}

// readModels returns an error if no vendors are configured or if all vendors fail.
// If all vendors fail, it also returns the empty catalog.
func (o *VendorsManager) readModels() (ret *VendorsModels, err error) {
	if len(o.Vendors) == 0 {
		err = errors.New(i18n.T("vendors_no_ai_vendors_configured_read_models"))
		return
	}

	ret = NewVendorsModels()
	failed := 0

	var wg sync.WaitGroup
	resultsChan := make(chan modelResult, len(o.Vendors))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, vendor := range o.Vendors {
		wg.Add(1)
		go o.fetchVendorModels(ctx, &wg, vendor, resultsChan)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		if result.err != nil {
			failed++
			fmt.Println(result.vendorName, result.err)
		} else {
			sort.Slice(result.models, func(i, j int) bool {
				return strings.ToLower(result.models[i]) < strings.ToLower(result.models[j])
			})
			ret.AddGroupItems(result.vendorName, result.models...)
		}
	}
	if failed == len(o.Vendors) {
		err = errors.New(i18n.T("vendors_all_failed_list_models"))
	}
	return
}

func (o *VendorsManager) fetchVendorModels(
	ctx context.Context, wg *sync.WaitGroup, vendor Vendor, resultsChan chan<- modelResult) {

	defer wg.Done()

	models, err := vendor.ListModels(ctx)
	select {
	case <-ctx.Done():
		return
	case resultsChan <- modelResult{vendorName: vendor.GetName(), models: models, err: err}:
	}
}

func (o *VendorsManager) Setup() (ret map[string]Vendor, err error) {
	ret = map[string]Vendor{}
	for _, vendor := range o.Vendors {
		fmt.Println()
		o.setupVendorTo(vendor, ret)
	}
	return
}

func (o *VendorsManager) SetupVendor(vendorName string, configuredVendors map[string]Vendor) (err error) {
	vendor := o.FindByName(vendorName)
	if vendor == nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("vendor_not_found"), vendorName))
		return
	}
	o.setupVendorTo(vendor, configuredVendors)
	return
}

func (o *VendorsManager) setupVendorTo(vendor Vendor, configuredVendors map[string]Vendor) {
	if vendorErr := vendor.Setup(); vendorErr == nil {
		fmt.Printf("%s\n", fmt.Sprintf(i18n.T("plugin_setup_configured"), vendor.GetName()))
		configuredVendors[strings.ToLower(vendor.GetName())] = vendor
	} else {
		delete(configuredVendors, strings.ToLower(vendor.GetName()))
		fmt.Printf("%s", fmt.Sprintf(i18n.T("plugin_setup_skipped"), vendor.GetName()))
	}
}

type modelResult struct {
	vendorName string
	models     []string
	err        error
}
