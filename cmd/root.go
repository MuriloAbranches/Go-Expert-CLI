package cmd

import (
	"database/sql"
	"os"

	"github.com/MuriloAbranches/Go-Expert-CLI/internal/database"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDB(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

var rootCmd = &cobra.Command{
	Use:   "Go-Expert-CLI",
	Short: "Go Expert CLI",
	Long:  `Go Expert CLI, use this CLI to create and find categories`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
