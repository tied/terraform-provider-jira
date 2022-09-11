package jira

import (
	"log"
	"sync"

	jira "github.com/andygrunwald/go-jira"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

type Config struct {
	jiraClient *jira.Client
	jiraLock   sync.Mutex
}

func (c *Config) createAndAuthenticateClient(d *schema.ResourceData) error {
	log.Printf("[INFO] creating jira client using environment variables")
	jiraClient, err := jira.NewClient(nil, d.Get("url").(string))
	if err != nil {
		return errors.Wrap(err, "creating jira client failed")
	}
	jiraClient.Authentication.SetBasicAuth(d.Get("user").(string), d.Get("password").(string))

	c.jiraClient = jiraClient

	return nil
}
