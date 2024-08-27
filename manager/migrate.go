
package manager

import (
	"fmt"
	"open-btm.com/models"
	"github.com/spf13/cobra"
)

var (
	btmmigrate= &cobra.Command{
		Use:   "migrate",
		Short: "Run Database Migration for found in init migration Models",
		Long:  `Migrate to init database`,
		Run: func(cmd *cobra.Command, args []string) {
			init_migrate()
		},
	}

	btmclean= &cobra.Command{
		Use:   "clean",
		Short: "Drop Database Models for found in init migration Models",
		Long:  `Drop Models found in the models definition`,
		Run: func(cmd *cobra.Command, args []string) {
			clean_database()
		},
	}

)

func init_migrate() {
	models.InitDatabase()
	fmt.Println("Migrated Database Models sucessfully")
}

func clean_database() {
	models.CleanDatabase()
	fmt.Println("Dropped Tables sucessfully")
}


func init() {
	goFrame.AddCommand(btmmigrate)
	goFrame.AddCommand(btmclean)
}

