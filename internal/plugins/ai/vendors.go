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
	mu            sync.Mutex
	modelsVersion uint64
	modelsLoading *modelsDiscovery
	Vendors       []Vendor
	VendorsByName map[string]Vendor
	Models        *VendorsModels
}

type modelsDiscovery struct {
	done chan struct{}
	once sync.Once
}

func (d *modelsDiscovery) notify() {
	d.once.Do(func() { close(d.done) })
}

// Called with mu held. Obsolete discovery can finish for its original caller,
// but readers and new membership must not wait on a removed provider.
func (o *VendorsManager) invalidateModels() {
	o.Models = nil
	o.modelsVersion++
	if loading := o.modelsLoading; loading != nil {
		o.modelsLoading = nil
		loading.notify()
	}
}

// AddVendors registers one or more vendors with the manager.
// Vendors are stored with lowercase keys to enable case-insensitive lookup.
// Membership changes invalidate the model catalog under the manager lock.
func (o *VendorsManager) AddVendors(vendors ...Vendor) {
	if len(vendors) == 0 {
		return
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	for _, vendor := range vendors {
		name := strings.ToLower(vendor.GetName())
		o.VendorsByName[name] = vendor
		o.Vendors = append(o.Vendors, vendor)
	}
	o.invalidateModels()
}

func (o *VendorsManager) Clear() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.VendorsByName = map[string]Vendor{}
	o.Vendors = []Vendor{}
	o.invalidateModels()
}

func (o *VendorsManager) vendorsSnapshot() []Vendor {
	o.mu.Lock()
	defer o.mu.Unlock()
	return append([]Vendor(nil), o.Vendors...)
}

func (o *VendorsManager) SetupFillEnvFileContent(envFileContent *bytes.Buffer) {
	for _, vendor := range o.vendorsSnapshot() {
		vendor.SetupFillEnvFileContent(envFileContent)
	}
}

func (o *VendorsManager) GetModels() (ret *VendorsModels, err error) {
	for {
		o.mu.Lock()
		if o.Models != nil {
			ret = o.Models
			o.mu.Unlock()
			return
		}
		if loading := o.modelsLoading; loading != nil {
			o.mu.Unlock()
			<-loading.done
			continue
		}
		vendors := append([]Vendor(nil), o.Vendors...)
		version := o.modelsVersion
		loading := &modelsDiscovery{done: make(chan struct{})}
		o.modelsLoading = loading
		o.mu.Unlock()

		// Fetch without the manager lock: provider callbacks and registrations
		// must not be blocked by an in-flight network request.
		ret, err = o.readModels(vendors)
		o.mu.Lock()
		if version == o.modelsVersion && err == nil {
			o.Models = ret
		}
		if o.modelsLoading == loading {
			o.modelsLoading = nil
		}
		loading.notify()
		o.mu.Unlock()
		return
	}
}

func (o *VendorsManager) Configure() (err error) {
	for _, vendor := range o.vendorsSnapshot() {
		_ = vendor.Configure()
	}
	return
}

func (o *VendorsManager) HasVendors() bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	return len(o.Vendors) > 0
}

// FindByName returns a vendor by name. Lookup is case-insensitive.
// For example, "OpenAI", "openai", and "OPENAI" all match the same vendor.
func (o *VendorsManager) FindByName(name string) Vendor {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.VendorsByName[strings.ToLower(name)]
}

func (o *VendorsManager) readModels(vendors []Vendor) (ret *VendorsModels, err error) {
	if len(vendors) == 0 {
		err = errors.New(i18n.T("vendors_no_ai_vendors_configured_read_models"))
		return
	}

	ret = NewVendorsModels()

	var wg sync.WaitGroup
	resultsChan := make(chan modelResult, len(vendors))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, vendor := range vendors {
		wg.Add(1)
		go o.fetchVendorModels(ctx, &wg, vendor, resultsChan)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		if result.err != nil {
			fmt.Println(result.vendorName, result.err)
		} else {
			models := append([]string(nil), result.models...)
			sort.Slice(models, func(i, j int) bool {
				return strings.ToLower(models[i]) < strings.ToLower(models[j])
			})
			ret.AddGroupItems(result.vendorName, models...)
		}
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
	for _, vendor := range o.vendorsSnapshot() {
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
