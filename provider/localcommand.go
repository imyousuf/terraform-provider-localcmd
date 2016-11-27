package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"bytes"
	"os/exec"
	"log"
)

const command_attr_key string = "command"
const stdout_attr_key string = "stdout"
const stderr_attr_key string = "stderr"

func executeCommand(data *schema.ResourceData, meta interface{}) error {
	passed_cmd, _ := data.Get(command_attr_key).(string)
	cmd := exec.Command("bash", "-c", passed_cmd)
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err := cmd.Run()
	log.Println("Done running")
	if err != nil {
		log.Printf("Error: %s", errOut.String())
		log.Fatal(err)
	}
	log.Printf("StdOut %s", out.String())
	log.Printf("StdErr %s", errOut.String())
	data.Set(stdout_attr_key, out.String())
	data.Set(stderr_attr_key, errOut.String())
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
				Sensitive: false,
			},
			stderr_attr_key: {
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: false,
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
