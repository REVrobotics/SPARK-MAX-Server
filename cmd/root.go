// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	spark0mq "github.com/willtoth/USB-BLDC-TOOL/spark0mq"
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

var cfgFile string

// Device is the COM port
var Device string

// Persist mode keeps connection alive while application is running
var Persist bool

// Remote mode sets up a TCP/IP server to stream the command line
var Remote bool

// port for grpc server
var port uint

// default port
var defaultPort uint = 8001

// Version of application
const (
	AppVersion = "0.1.1"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: AppVersion,
	Use:     "revbldc",
	Short:   "configure and control REV CAN BLDC controller",
	Long: `Use this tool to configure and command the REV CAN BLDC
controller over USB, either by CLI interface or GUI:

revbldc tool provides bindings to talk to the REV motor
controller and can be called via command line or
externally. It can update firmware, set and get parameters
and save/load configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Remote == true {
			server, err := spark0mq.Spark0mqStart(int(port))
			if err != nil {
				log.Fatalf("Failed to start server %v", err)
			}
			//Press enter to kill server
			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')
			server.Stop()
		} else if Persist == true {
			fmt.Println("Interactive feature not yet implmeneted")
		} else {
			cmd.Usage()
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Remote == false && Persist == false {
			err := sparkusb.Connect(Device)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if Remote == false && Persist == false {
			sparkusb.Disconnect()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.USB-BLDC-TOOL.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Device, "device", "d", "", "Set the device COM port")
	rootCmd.PersistentFlags().BoolVarP(&Persist, "interactive", "i", false, "Keep connection alive between commands")
	rootCmd.PersistentFlags().BoolVarP(&Remote, "remote", "r", false, "Run a TCP/IP server to stream commands")
	rootCmd.PersistentFlags().UintVarP(&port, "port", "p", defaultPort, "Set port for 0mq server")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".USB-BLDC-TOOL" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".USB-BLDC-TOOL")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
