package initialize

import (
	"io/ioutil"
	"path"
	"path/filepath"
)

func parsePackage(dir string, pkg *Package) error {

	fsinfo, e := ioutil.ReadDir(dir)

	if e != nil {
		return e
	}

	for _, v := range fsinfo {

		//joined path (of the parent directories and current one)
		jpth := path.Join(dir, v.Name())

		if v.IsDir() {
			//a new package

			var spkg Package

			//because the name is an array (see `Package.go`) we want to get the package names of all the parents
			spkg.Name = v.Name()

			spkg.parent = pkg //set the parent package

			e = parsePackage(jpth, &spkg)

			if e != nil {
				return e
			}

			pkg.ChildPacks = append(pkg.ChildPacks, &spkg)

			continue
		}

		//only append a new class if it's a tusk file
		if filepath.Ext(v.Name()) != ".tusk" {
			continue
		}

		//a new class in the package
		pf, e := parseFile(jpth)

		if e != nil {
			return e
		}

		pkg.Files = append(pkg.Files, pf)
	}

	return nil
}
