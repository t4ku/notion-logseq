/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// databaseCmd represents the list command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "List Notion Databases",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("database called")

		token := os.Getenv("NOTION_TOKEN")

		url := "https://api.notion.com/v1/search"
		payload := strings.NewReader("{\"page_size\":100, \"filter\": { \"value\": \"database\", \"property\":\"object\" }}")

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("accept", "application/json")
		req.Header.Add("Notion-Version", "2022-06-28")
		req.Header.Add("Authorization", "Bearer "+token+"")
		req.Header.Add("content-type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		result := gjson.Get(string(body), "results")

		result.ForEach(func(key, value gjson.Result) bool {
			// println(value.String())
			database_id := value.Get("id")
			title := value.Get("title.0.text.content")

			fmt.Println(database_id.String() + "\t" + title.String())
			// fmt.Println(key)
			// database := value[key]
			return true
		})
		// basic print
		//fmt.Println(res)
		//fmt.Println(string(body))

		// pretty print
		// var buf bytes.Buffer
		// err := json.Indent(&buf, []byte(body), "", "  ")
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
