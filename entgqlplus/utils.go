package entgqlplus

import (
	"log"
	"os"
	"path"
	"strings"

	"entgo.io/ent/entc/gen"
	"gopkg.in/yaml.v3"
)

var (
	lower = strings.ToLower
	camel = gen.Funcs["camel"].(func(string) string)
	snake = gen.Funcs["snake"].(func(string) string)
)

func writeFiles(files []file) {
	for i := range files {
		writeFile(files[i])
	}
}

func writeFile(f file) {
	err := os.MkdirAll(path.Dir(f.Path), 0666)
	catch(err)
	os.WriteFile(f.Path, []byte(f.Buffer), 0666)
}

func catch(err error) {
	if err != nil {
		log.Fatalln("entgqlplus:", err)
	}
}

func readGqlGen(fpath string) gqlGen {
	buffer, err := os.ReadFile(fpath)
	catch(err)
	out := gqlGen{}
	err = yaml.Unmarshal(buffer, &out)
	catch(err)
	out.Exec.Dir = path.Dir(out.Exec.FileName)
	out.Model.Dir = path.Dir(out.Model.FileName)
	return out
}

func inArray[T string | int | uint](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func cleanFiles(resolverDir, schemaDir string) {
	os.RemoveAll(resolverDir)
	os.RemoveAll(schemaDir)
}
