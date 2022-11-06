package entgqlplus

import (
	"bytes"
	"text/template"

	"entgo.io/ent/entc/gen"
)

func parseTemplate(filename string, data *templateData) string {
	buffer, err := assets.ReadFile("templates/" + filename)
	catch(err)
	t, err := template.New(filename).Funcs(gen.Funcs).Parse(string(buffer))
	catch(err)
	out := bytes.Buffer{}
	err = t.Execute(&out, data)
	catch(err)
	return out.String()
}
