package cmd

import (
	"blocky/api"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

func NewListsCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "lists",
		Short: "lists operations",
	}

	c.AddCommand(newRefreshCommand())

	return c
}

func newRefreshCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "refresh",
		Short: "refreshes all lists",
		Run:   refreshList,
	}
}

func refreshList(_ *cobra.Command, _ []string) {
	resp, err := http.Post(apiURL(api.PathListsRefresh), "application/json", nil)
	if err != nil {
		log.Fatal("can't execute", err)

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("NOK: %s %s", resp.Status, string(body))

		return
	}

	log.Info("OK")
}
