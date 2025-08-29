package ai

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

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
}

func (o *VendorsManager) AddVendors(vendors ...Vendor) {
	for _, vendor := range vendors {
		o.VendorsByName[vendor.GetName()] = vendor
		o.Vendors = append(o.Vendors, vendor)
	}
}

func (o *VendorsManager) Clear(vendors ...Vendor) {
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
	if o.Models == nil {
		err = o.readModels()
	}
	ret = o.Models
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

func (o *VendorsManager) FindByName(name string) Vendor {
	return o.VendorsByName[name]
}

func (o *VendorsManager) readModels() (err error) {
	if len(o.Vendors) == 0 {

		err = fmt.Errorf("no AI vendors configured to read models from. Please configure at least one AI vendor")
		return
	}

	o.Models = NewVendorsModels()

	var wg sync.WaitGroup
	resultsChan := make(chan modelResult, len(o.Vendors))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, vendor := range o.Vendors {
		wg.Add(1)
		go o.fetchVendorModels(ctx, &wg, vendor, resultsChan)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect results
	for result := range resultsChan {
		if result.err != nil {
			fmt.Println(result.vendorName, result.err)
		} else {
			sort.Slice(result.models, func(i, j int) bool {
				return strings.ToLower(result.models[i]) < strings.ToLower(result.models[j])
			})
			o.Models.AddGroupItems(result.vendorName, result.models...)
		}
	}
	return
}

func (o *VendorsManager) fetchVendorModels(
	ctx context.Context, wg *sync.WaitGroup, vendor Vendor, resultsChan chan<- modelResult) {

	defer wg.Done()

	models, err := vendor.ListModels()
	select {
	case <-ctx.Done():
		// Context canceled, don't send the result
		return
	case resultsChan <- modelResult{vendorName: vendor.GetName(), models: models, err: err}:
		// Result sent
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
		err = fmt.Errorf("vendor %s not found", vendorName)
		return
	}
	o.setupVendorTo(vendor, configuredVendors)
	return
}

func (o *VendorsManager) setupVendorTo(vendor Vendor, configuredVendors map[string]Vendor) {
	if vendorErr := vendor.Setup(); vendorErr == nil {
		fmt.Printf("[%v] configured\n", vendor.GetName())
		configuredVendors[vendor.GetName()] = vendor
	} else {
		delete(configuredVendors, vendor.GetName())
		fmt.Printf("[%v] skipped\n", vendor.GetName())
	}
}

type modelResult struct {
	vendorName string
	models     []string
	err        error
}
