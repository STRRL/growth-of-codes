package command

import (
	"context"
	"github.com/STRRL/growth-of-codes/pkg/persistent/sink"
	"github.com/STRRL/growth-of-codes/pkg/persistent/source"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewLoadCommand() (*cobra.Command, error) {
	options := &LoadOptions{}
	cmd := &cobra.Command{
		Use:   "load",
		Short: "load data to mysql backend",
		RunE: func(cmd *cobra.Command, args []string) error {
			return load(context.TODO(), args[0], options)
		},
	}
	cmd.Flags().StringVar(&options.MySQLDSN, "mysql-dsn", "", "mysql dsn")
	cmd.Flags().StringVar(&options.FromCSV, "from-csv", "", "from csv file")
	return cmd, nil
}

type LoadOptions struct {
	MySQLDSN string
	FromCSV  string
}

func load(ctx context.Context, projectName string, options *LoadOptions) error {
	csvSource := source.NewCSVSource(projectName, options.FromCSV)
	db, err := gorm.Open(mysql.Open(options.MySQLDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	gormSink := sink.NewGORMRepository(db)

	all, err := csvSource.ListAll()
	if err != nil {
		return err
	}
	return gormSink.SaveMulti(ctx, all)
}
