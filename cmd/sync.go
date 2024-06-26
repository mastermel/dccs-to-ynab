/*
Copyright © 2020 Mel Green <mastermel@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/mastermel/dccs-to-ynab/accounts"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Performs a transaction sync from DCCS -> YNAB for each account configured to sync.",
	Run: func(cmd *cobra.Command, args []string) {
		accounts.Sync()
	},
}


// syncLoopCmd represents the sync in loop command
var syncLoopCmd = &cobra.Command{
	Use:   "sync-loop",
	Short: "Performs a transaction sync from DCCS -> YNAB for each account configured to sync on a regular interval.",
	Run: func(cmd *cobra.Command, args []string) {
		accounts.SyncLoop()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(syncLoopCmd)
}
