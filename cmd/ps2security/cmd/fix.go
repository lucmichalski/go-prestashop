package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"

	// "github.com/k0kubun/pp"
	"github.com/karrick/godirwalk"
	"github.com/spf13/cobra"
)

var (
	workDir         string
	project         string
	author          string
	company         string
	copyright       string
	email           string
	sorted          bool
	noChmod         bool
	noIndex         bool
	dryRun          bool
	skipDirectories = []string{"vendor", "cache"}
)

const (
	psCheckFilePath = `config/defines.inc.php`
	psIndexTemplate = `<?php

/**
 * {{ .project }}
 *
 * Copyright {{ .year }} {{ .company }}
 *
 * @author   {{ .author }}
 * @email    {{ .email }}
 * @copyright Copyright (c) {{ .year }} {{ .company }}. All rights reserved
 *
 */

header("Pragma: no-cache");

header("Cache-Control: post-check=0, pre-check=0", false);
header("Cache-Control: no-store, no-cache, must-revalidate");

header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Expires: Mon, 26 Jul 1997 05:00:00 GMT");

header("Location: ../");

exit;`
)

var FixCmd = &cobra.Command{
	Use:     "fix",
	Aliases: []string{"f"},
	Short:   "Fix missing index.php files and permissions on a prestashop worktree.",
	Long:    "Fix missing index.php files and permissions on a prestashop worktree.",
	Run: func(cmd *cobra.Command, args []string) {

		// check if it is a prestashop install, does config/defines.inc.php exists
		psDefineInc := filepath.Join(workDir, psCheckFilePath)
		if _, err := os.Stat(psDefineInc); os.IsNotExist(err) {
			log.Fatalf("Sounds like it is not a prestashop workdir. Could not find file '%s'\n", psDefineInc)
		}

		indexPHP := bytes.NewBufferString("")
		indexTemplate, _ := template.New("").Parse(psIndexTemplate)
		indexTemplate.Execute(indexPHP, map[string]string{
			"project": fmt.Sprintf("%s", project),
			"company": fmt.Sprintf("%s", company),
			"author":  fmt.Sprintf("%s", author),
			"email":   fmt.Sprintf("%s", email),
			"year":    time.Now().Format("Y"),
		})

		// err :=
		godirwalk.Walk(workDir, &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				if !de.IsDir() {

					err := os.Chmod(osPathname, 0644)
					if err != nil {
						return err
					} else {
						log.Infof("Changed file '%s' permissions to 0644\n", osPathname)
					}

				} else {
					err := os.Chmod(osPathname, 0755)
					if err != nil {
						return err
					} else {
						log.Infof("Changed directory '%s' permissions to 0755\n", osPathname)
					}

					if inSliceContains(osPathname, skipDirectories, false) {
						return nil
					}
					indexFile := filepath.Join(osPathname, "index.php")
					if _, err := os.Stat(indexFile); os.IsNotExist(err) {
						log.Warnf("Missing index.php at '%s'\n", osPathname)
						if !dryRun {
							err := ioutil.WriteFile(indexFile, indexPHP.Bytes(), 0644)
							return err
						}
					} else {
						log.Infof("Success 'index.php' exists at '%s'\n", osPathname)
					}
				}
				return nil
			},
			Unsorted: !sorted,
		})
		// checkErr("walk errors", err)

	},
}

func init() {
	FixCmd.Flags().StringVarP(&workDir, "workdir", "w", "../../shared/www", "workdir to explore and fix.")
	FixCmd.Flags().StringVarP(&project, "project", "p", "Prestashop++", "Project name.")
	FixCmd.Flags().StringVarP(&company, "company", "c", "Evolutive Group", "Company name.")
	FixCmd.Flags().StringVarP(&author, "author", "a", "Luc Michalski", "Author name.")
	FixCmd.Flags().StringVarP(&email, "email", "t", "technique@evolutive-business.com", "Contact email address.")
	FixCmd.Flags().BoolVarP(&noChmod, "no-chmod", "m", false, "Do not add fix files and directories permissions.")
	FixCmd.Flags().BoolVarP(&noIndex, "no-index", "i", false, "Do not add missing index.php.")
	FixCmd.Flags().BoolVarP(&sorted, "sorted", "s", false, "walk directories in sorted way (slower).")
	FixCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "dry run over wortree and list potential changes.")
	RootCmd.AddCommand(FixCmd)
}
