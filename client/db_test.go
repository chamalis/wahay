package client

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "gopkg.in/check.v1"
)

func (s *clientSuite) Test_write_writesDBContent(c *C) {
	tempDir, err := ioutil.TempDir("", "test")
	c.Assert(err, IsNil)
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "testfile.txt")

	file, err := os.Create(tempFile)
	c.Assert(err, IsNil)

	defer file.Close()
	defer os.Remove(tempFile)

	content := []byte("example content")

	db := &dbData{
		filename: tempFile,
		content:  content,
	}

	err = db.write()
	c.Assert(err, IsNil)

	readContent, err := ioutil.ReadFile(tempFile)

	c.Assert(err, IsNil)
	c.Assert(readContent, DeepEquals, content)
}

func (s *clientSuite) Test_exists_returnsTrueIfStringExistsInContent(c *C) {
	db := &dbData{
		content: []byte("this is true"),
	}

	exampleString := "true"

	res := db.exists(exampleString)

	c.Assert(res, Equals, true)
}

func (s *clientSuite) Test_exists_returnsFalseIfStringDoesntExistInContent(c *C) {
	db := &dbData{
		content: []byte("this is false"),
	}

	exampleString := "true"

	res := db.exists(exampleString)

	c.Assert(res, Equals, false)
}

func (s *clientSuite) Test_replaceString_findsAndReplacesContentInDB(c *C) {
	db := &dbData{
		content: []byte("this is true"),
	}

	find := "true"
	replace := "false"
	expectedContent := []byte("this is false")

	db.replaceString(find, replace)

	c.Assert(db.content, DeepEquals, expectedContent)
}
