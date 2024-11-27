package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	scaleway "github.com/scaleway/terraform-provider-scaleway/v2/internal/provider"
)

func NewProvider() *schema.Provider {
	p := scaleway.Provider(scaleway.DefaultConfig())
	return p()
}
