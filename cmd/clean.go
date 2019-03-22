// Copyright © 2019 John Carlo Salter <juuncaerlum@gmail.com>
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
	"fmt"

	"github.com/spf13/cobra"
)

var cleanAll bool
var cleanPkgdbs bool
var pkgdbs string

// var cleanCache bool
// var srcCaches string
// var binaryCaches string

// cleanCmd represents the clean command
var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean up cached information",
	Long:  "clean up cached source files and binaries, as well as the saved package database.",
	RunE:  clean,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("clean called")
	//	},
}

func init() {
	CleanCmd.Flags().BoolVar(&cleanAll, "all", false, "clean all cached items")
	CleanCmd.Flags().BoolVar(&cleanPkgdbs, "pkgdbs", false, "Remove cached package databases.")
	CleanCmd.Flags().StringVar(&pkgdbs, "dbs", "ALL", "Package databases to remove.")
	CleanCmd.Flags().BoolVar(&cleanCache, "cache", false, "Remove cache sources and/or binaries")
	// cleanCmd.Flags().StringVar(&srcCaches, "src", "ALL", "Clean src caches in clean --cache")
	// cleanCmd.Flags().StringVar(&binaryCaches, "binary", "ALL", "Clean binary caches in clean --cache")

	RootCmd.AddCommand(CleanCmd)
}

func clean(cmd *cobra.Command, args []string) error {

	if !cleanAll && !cleanPkgdbs && !cleanCache {
		fmt.Println("No clean options passed -- not cleaning.")
	}
	if cleanAll {
		fmt.Println("Cleaning all.")
	} else {
		if cleanPkgdbs {
			if pkgdbs == "ALL" {
				fmt.Println("Cleaning all pkgdbs")
			} else {
				fmt.Println(fmt.Sprintf("Cleaning specific package databases: %s", pkgdbs))
			}
		}
		if cleanCache {
			cleanCacheFolders()
		}
	}
	fmt.Println("Donezo.")
	return nil
}

/*
func cleanCacheFolders() {
	cachePath := userCache(cfg.Cache)

	if srcCaches == "ALL" {
		fmt.Println("Cleaning all src caches.")
		_ = deleteCacheSubfolders(nil, "src", cachePath)
	} else {
		fmt.Println(fmt.Sprintf("Cleaning specific src caches: %s", srcCaches))
		srcRepos := strings.Split(srcCaches, ",")
		_ = deleteCacheSubfolders(srcRepos, "src", cachePath)
	}

	if binaryCaches == "ALL" {
		fmt.Println("Cleaning all binary caches.")
		deleteCacheSubfolders(nil, "binary", cachePath)
	} else {
		fmt.Println(fmt.Sprintf("Cleaning specific binary caches: %s", binaryCaches))
		binaryRepos := strings.Split(binaryCaches, ",")
		_ = deleteCacheSubfolders(binaryRepos, "binary", cachePath)
	}
}

func deleteAllCacheSubfolders(cacheDirectory string) {
	deleteCacheSubfolders(nil, "src", cacheDirectory)
	deleteCacheSubfolders(nil, "binary", cacheDirectory)
}

func deleteCacheSubfolders(repos []string, subfolder string, cacheDirectory string) error {
	cacheDirFsObject, err := fs.Open(cacheDirectory)
	if err != nil {
		return err
	}

	repoFolders, _ := cacheDirFsObject.Readdir(0)

	if repos == nil || len(repos) == 0 {
		for _, repoFolder := range repoFolders {
			subfolderPath := filepath.Join(cacheDirectory, repoFolder.Name(), subfolder)
			fs.RemoveAll(subfolderPath)
		}
	} else {
		for _, repoToClear := range repos {
			for _, repoFolder := range repoFolders {
				if repoToClear == repoFolder.Name() {
					subfolderPath := filepath.Join(cacheDirectory, repoFolder.Name(), subfolder)
					fs.Remove(subfolderPath)
				}
			}
		}
	}
	return nil
}
*/
