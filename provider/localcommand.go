package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"bytes"
	"os/exec"
)

const command_attr_key string = "command"
const stdout_attr_key string = "stdout"
const stderr_attr_key string = "stderr"

func executeCommand(data *schema.ResourceData, meta interface{}) error {
	cmd := exec.Command("bash", "-c", data.Get(command_attr_key).(string))
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	cmd.Run()
	data.Set(stdout_attr_key, out.String())
	data.Set(stderr_attr_key, out.String())
	return nil
}

func doExec() *schema.Resource {
	return &schema.Resource{
		Read: executeCommand,
		Schema: map[string]*schema.Schema{
			command_attr_key: {
				Required: true,
				Type:     schema.TypeString,
			},
			stdout_attr_key: {
				Type:     schema.TypeString,
				Computed: true,
			},
			stderr_attr_key: {
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
			"localcmd_exec": doExec(),
		},
	}
}
