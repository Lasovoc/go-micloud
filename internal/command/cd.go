package command

import (
	"github.com/urfave/cli/v2"
	"go-micloud/internal/file"
	"go-micloud/internal/folder"
	"strings"
)

type Command struct {
	HttpApi file.Api
	Folder  *folder.Folder
}

func (r *Command) Cd() *cli.Command {
	return &cli.Command{
		Name:            "cd",
		Usage:           "改变当前目录，例如：cd movies",
		SkipFlagParsing: true,
		Action: func(context *cli.Context) error {
			var (
				dirName = context.Args().First()
				err     error
			)
			if strings.Trim(dirName, " ") == "/" || strings.Trim(dirName, " ") == "" {
				dirName = "/"
			}
			err = folder.ChangeFolder(r.Folder, dirName)
			if err != nil {
				return err
			}
			files, err := r.HttpApi.GetFolder(r.Folder.Cursor.Id)
			if err != nil {
				return err
			}
			folder.AddFolder(r.Folder, r.Folder.Cursor.Name, files)

			return nil
		},
	}
}
