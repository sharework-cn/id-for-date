package idp

import (
	"fmt"
	"os"

	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd *cobra.Command = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "An identifier probe that operates on a table without a time-based index",
		Long:  `If you've got a table without any time-based index, searching for info based on a specific time can be a real headache. However, this tool is incredibly helpful as it allows you to effortlessly locate the ID associated with a particular time, eliminating the need for an exhaustive search through the entire table.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is the yaml file in $HOME/.config/idp)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
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

		// Search config in home directory with name ".yaml"
		viper.AddConfigPath(home + ".config/idp")
		viper.SetConfigName(".yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func RootCommand(name string) *cobra.Command {
	return rootCmd
}
