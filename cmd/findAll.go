package cmd

import (
	"fmt"

	"github.com/MuriloAbranches/Go-Expert-CLI/internal/database"
	"github.com/spf13/cobra"
)

func newFindAllCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "findAll",
		Short: "Find all categories",
		Long:  "FInd all catgories from database",
		RunE:  runFindAll(categoryDB),
	}
}

func runFindAll(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		categories, err := categoryDB.FindAll()
		if err != nil {
			return err
		}

		for _, category := range categories {
			fmt.Printf("%s|%s|%s\n", category.ID, category.Name, category.Description)
		}

		return nil
	}
}

func init() {
	findAllCmd := newFindAllCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(findAllCmd)
}
