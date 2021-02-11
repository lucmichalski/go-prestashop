package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/h2non/bimg"
	"github.com/k0kubun/pp"
	"github.com/karrick/godirwalk"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/lucmichalski/go-prestashop/internal/models"
)

var (
	psDir         string
	workDir       string
	dryRun        bool
	db            *gorm.DB
	dbName        string
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbTablePrefix string
	imgTypes      = map[string]string{ // todo. how to deal with other types of images ?
		"p":  "products",
		"c":  "categories",
		"m":  "manufacturers",
		"su": "suppliers",
		"st": "stores",
	}
	imgQuality int
	imgExt     []string
	imgDefExt  = []string{".png", ".jpg", ".jpeg", ".webp"} // jpeg,png, webp,tiff,gif,pdf,svg
)

var ImportCmd = &cobra.Command{
	Use:     "load",
	Aliases: []string{"l"},
	Short:   "load prestashop-shop-creator xml to a mysql database.",
	Long:    "load prestashop-shop-creator xml to a mysql database",
	Run: func(cmd *cobra.Command, args []string) {

		if !dryRun {
			var err error
			// connect to database
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   dbTablePrefix,                   // table name prefix, table for `Lang` would be `eg_lang`
					SingularTable: true,                            // use singular table name, table for `User` would be `user` with this option enabled
					NameReplacer:  strings.NewReplacer("ID", "Id"), // use name replacer to change struct/field name before convert it to db name
					// SkipDefaultTransaction: true,
				},
			})
			if err != nil {
				log.Fatal(err)
			}

		}

		// todo. load image list or parse them by name ?
		// var images []*models.Image
		// db.Find(&images)

		bimgOpts := bimg.Options{
			Lossless: true,
			Quality:  imgQuality,
		}

		for imgType, imgTypeLabel := range imgTypes {

			imgDir := filepath.Join(workDir, "img", imgType)

			// load image types
			var imageTypes []*models.ImageType
			db.Where(imgTypeLabel + " = 1").Find(&imageTypes)

			i := 1
			err := godirwalk.Walk(imgDir, &godirwalk.Options{
				Callback: func(osPathname string, de *godirwalk.Dirent) error {
					if !de.IsDir() {
						// get the extension from the file path

						extension := filepath.Ext(osPathname)
						filename := path.Base(osPathname)
						basename := strings.Replace(filename, extension, "", -1)
						if inSlice(extension, imgExt, false) {
							if options.verbose {
								log.Println("found=", osPathname, "filename=", filename, "extension=", extension, "basename=", basename)
							}

							// Read original image
							buffer, err := bimg.Read(osPathname)
							checkErr("bimg.Read, error", err)

							// Import image buffer
							newImage := bimg.NewImage(buffer)

							// Get the image dimension
							imgDim, err := newImage.Size()

							// Get the image type
							mediaType := newImage.Type()

							switch mediaType {
							case "unknown":
								return nil // errors.New("Unsupported image format")
							}

							bimgOpts.Height = imgDim.Height
							bimgOpts.Width = imgDim.Width

							// Process image quality
							newBytes, err := newImage.Process(bimgOpts)
							checkErr("bimg.Process, error", err)

							// Write final output
							if !dryRun {
								destinationPrepfixPath := filepath.Join(psDir, "img", imgType)
								destinationFinalPath := buildFolderForImage(destinationPrepfixPath, i)
								if err := os.MkdirAll(destinationFinalPath, 0755); err != nil {
									log.Fatal(err)
								}
								destinationFilePath := filepath.Join(destinationFinalPath, fmt.Sprintf("%d%s", i, extension))
								log.Println("destinationFilePath=", destinationFilePath)
								err = bimg.Write(destinationFilePath, newBytes)
								checkErr("bimg.Write, error", err)

								for _, imageType := range imageTypes {

									// Import image buffer
									newImage := bimg.NewImage(buffer)

									// Get the image dimension
									// imgDim, err := newImage.Size()

									// Get the image type
									mediaType := newImage.Type()

									switch mediaType {
									case "unknown":
										return nil // errors.New("Unsupported image format")
									}

									bimgOpts.Height = int(imageType.Height)
									bimgOpts.Width = int(imageType.Width)

									// Process image quality
									newBytes, err := newImage.Process(bimgOpts)
									checkErr("bimg.Process, error", err)

									// Write final output
									destinationPrepfixPath := filepath.Join(psDir, "img", imgType)
									destinationFinalPath := buildFolderForImage(destinationPrepfixPath, i)
									if err := os.MkdirAll(destinationFinalPath, 0755); err != nil {
										log.Fatal(err)
									}
									destinationFilePath := filepath.Join(destinationFinalPath, fmt.Sprintf("%d-%s%s", i, imageType.Name, extension))
									log.Println("destinationFilePath=", destinationFilePath)
									err = bimg.Write(destinationFilePath, newBytes)
									checkErr("bimg.Write, error", err)

								}

								i++
							}
						}
					}
					return nil
				},
				Unsorted: true,
			})
			checkErr("Dir walk, error", err)

		}
	},
}

func init() {
	// todo. manage ssl connections for db
	ImportCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	ImportCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	ImportCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	ImportCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	ImportCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	ImportCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	ImportCmd.Flags().StringVarP(&psDir, "ps-dir", "p", "../../shared/www", "prestashop directory")
	ImportCmd.Flags().IntVarP(&imgQuality, "img-quality", "q", 90, "Image format quality. (works only for JPEG format)")
	ImportCmd.Flags().StringSliceVarP(&imgExt, "img-extension", "e", imgDefExt, "Process only image extensions")
	ImportCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ImportCmd)
}

func buildFolderForImage(basePath string, id int) string {
	idStr := fmt.Sprintf("%d", id)
	idLength := len(idStr)
	for i := 0; i < idLength; i++ {
		basePath = filepath.Join(basePath, string(idStr[i]))
	}
	return basePath
}
