package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func executeCommand(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func doExec() *schema.Resource {
	return &schema.Resource{
		Read: executeCommand,
		Schema: map[string]*schema.Schema{
			"command": {
				Required: true,
				Type:     schema.TypeString,
			},
			"stdout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stderr": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ProvideLocalCommand() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"exec": doExec(),
		},
	}
}
