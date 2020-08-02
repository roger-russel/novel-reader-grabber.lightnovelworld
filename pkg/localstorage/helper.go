package localstorage

import "os"

//ToDO check if is a directory or if it was a file
func getNovelPath(source string, novelSlug string) (path string, err error) {

	path = fullPathRootFolder + "/" + source + "/" + novelSlug

	err = os.MkdirAll(path, 0750)

	if os.IsExist(err) {
		return path, nil
	}

	return path, err

}
