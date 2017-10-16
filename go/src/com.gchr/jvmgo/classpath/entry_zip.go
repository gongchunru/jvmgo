package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

//表示 ZIP或者JAR文件形式存在的类路径
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry  {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}


func (self *ZipEntry) readClass(clasName string) ([]byte, Entry, error){
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		 return nil, nil, err
	}

	defer r.Close()

	/*
	for循环的range格式可以对slice进行迭代循环
	 */
	for _, f:= range  r.File{
		if f.Name == clasName {
			rc, err := f.Open()
			if err != nil{
				return nil, nil, err
			}

			defer rc.Close()

			data, err := ioutil.ReadAll(rc)

			if err != nil {
				 return nil, nil, err
			}

			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + clasName)
}

func (self *ZipEntry) String() string {
	return self.absPath
}


