/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// httpServCmd represents the httpServ command
var httpServCmd = &cobra.Command{
	Use:   "httpServ [port]",
	Short: "serve up a message at a given port",
	Long: `httpServ -m [message] [port]
	`,
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])
		if err != nil || (port <= 0 && port >= 65535) {
			fmt.Println("cannot convert passed arg to a port number in (0, 65535]")
			return
		}

		message, _ := cmd.Flags().GetString("message")

		// set up a server at localhost:port/ returning message
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s\n", message)
		})
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	},
}

func init() {
	rootCmd.AddCommand(httpServCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpServCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpServCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	httpServCmd.Flags().StringP("message", "m", "ok", "set your custom message")
}
