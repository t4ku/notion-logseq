/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"gopkg.in/go-playground/validator.v9"
)

type Options struct {
	Format string `validate:"oneof=csv"`
}

var (
	o = &Options{}
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page database_id",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var validate = validator.New()
		err := validate.Struct(*o)

		if err != nil {
			var errorText []string
			for _, err := range err.(validator.ValidationErrors) {
				errorText = append(errorText, err.Tag())
			}
			return fmt.Errorf("Parameter error: %s", strings.Join(errorText, "\n"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("page called")
		database_id := args[0]
		//fmt.Println(database_id)

		token := os.Getenv("NOTION_TOKEN")

		url := "https://api.notion.com/v1/databases/" + database_id + "/query"
		payload := strings.NewReader(``)

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("accept", "application/json")
		req.Header.Add("Notion-Version", "2022-06-28")
		req.Header.Add("Authorization", "Bearer "+token+"")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		//fmt.Println(string(body))
		body, _ := io.ReadAll(res.Body)

		results := gjson.Get(string(body), "results")
		results.ForEach(func(key, value gjson.Result) bool {
			page_id := value.Get("id")
			//parent_type := value.Get("parent.type")
			//parent_id := value.Get("parent.id")
			url := value.Get("url")
			//created_time := value.Get("created_time")
			//last_edited_time := value.Get("last_edited_time")
			// TODO: find title property name from properties
			title := value.Get("properties.課題名.title.0.text.content")
			fmt.Println(page_id.String() + "\t" + title.String() + "\t" + url.String())
			return true
		})
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)
	pageCmd.Flags().StringVarP(&o.Format, "format", "f", "csv", "specify output format")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//pageCmd.PersistentFlags().String("database_id", "", "database id where the pages exists")
	//pageCmd.MarkPersistentFlagRequired("database_id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
