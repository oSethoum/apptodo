package entgqlplus

import (
	"embed"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

//go:embed templates
var assets embed.FS

type extension struct {
	entc.DefaultExtension
	hooks  []gen.Hook
	config *config
}

func (e *extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		files := []file{}
		schemaDir := path.Join(path.Dir(e.config.GqlGenPath), path.Dir(e.config.GqlGen.Schema[0]))
		resolverDir := path.Join(path.Dir(e.config.GqlGenPath), e.config.GqlGen.Resolver.Dir)
		dbDir := path.Join(path.Dir(e.config.GqlGenPath), "db")
		authDir := path.Join(path.Dir(e.config.GqlGenPath), "auth")

		data := &templateData{
			Config: e.config,
		}

		data.parse(g)

		files = append(files,
			file{
				Path:   path.Join(resolverDir, "schema.resolvers.go"),
				Buffer: parseTemplate("schema.resolvers.go.tmpl", data),
			},
			file{
				Path:   path.Join(resolverDir, "resolver.go"),
				Buffer: parseTemplate("resolver.go.tmpl", data),
			},

			file{
				Path:   path.Join(schemaDir, "types.graphqls"),
				Buffer: parseTemplate("types.graphqls.go.tmpl", nil),
			},
		)

		if e.config.Mutation {
			for i := range data.Nodes {
				data.Node = data.Nodes[i]
				files = append(files,
					file{
						Path:   path.Join(resolverDir, snake(data.Node.Name)+".resolvers.go"),
						Buffer: parseTemplate("node.resolvers.go.tmpl", data),
					},
					file{
						Path:   path.Join(schemaDir, snake(data.Node.Name)+".graphqls"),
						Buffer: parseTemplate("node.graphqls.go.tmpl", data),
					},
				)
			}
		}

		if e.config.Echo {
			files = append(files,
				file{
					Path:   path.Join(path.Dir(e.config.GqlGenPath), "routes/routes.go"),
					Buffer: parseTemplate("routes.go.tmpl", data),
				},
				file{
					Path:   path.Join(path.Dir(e.config.GqlGenPath), "handlers/handlers.go"),
					Buffer: parseTemplate("handlers.go.tmpl", data),
				},
				file{
					Path:   path.Join(path.Dir(e.config.GqlGenPath), "server.go"),
					Buffer: parseTemplate("server.go.tmpl", data),
				},
			)
		}

		if len(e.config.Database) > 0 {
			files = append(files,
				file{
					Path:   path.Join(dbDir, "db.go"),
					Buffer: parseTemplate("db.go.tmpl", data),
				},
			)
		}

		if e.config.FileUpload {
			files = append(files,
				file{
					Path:   path.Join(resolverDir, "upload.resolvers.go"),
					Buffer: parseTemplate("upload.resolvers.go.tmpl", data),
				},
				file{
					Path:   path.Join(schemaDir, "upload.graphqls"),
					Buffer: parseTemplate("upload.graphqls.go.tmpl", nil),
				},
			)
		}

		if e.config.Subscription && data.HasSubscription {
			files = append(files,
				file{
					Path:   path.Join(resolverDir, "notifiers.go"),
					Buffer: parseTemplate("notifiers.go.tmpl", data),
				},
				file{
					Path:   path.Join(resolverDir, "types.go"),
					Buffer: parseTemplate("types.go.tmpl", data),
				},
			)
		}

		if e.config.JWT {
			files = append(files,
				file{
					Path:   path.Join(authDir, "login.go"),
					Buffer: parseTemplate("auth.login.go.tmpl", data),
				},
				file{
					Path:   path.Join(authDir, "middleware.go"),
					Buffer: parseTemplate("auth.middleware.go.tmpl", data),
				},
				file{
					Path:   path.Join(authDir, "types.go"),
					Buffer: parseTemplate("auth.types.go.tmpl", data),
				},
			)
		}

		if e.config.Privacy {
			files = append(files, file{
				Path:   path.Join(authDir, "privacy.go"),
				Buffer: parseTemplate("auth.privacy.go.tmpl", data),
			})

			for _, n := range g.Nodes {
				fpath := path.Join(path.Dir(e.config.GqlGenPath), "ent/schema/", lower(n.Name)+".go")
				f, err := os.OpenFile(fpath, os.O_APPEND, 0666)
				if err != nil {
					log.Fatalln(fpath, err)
				}
				data.Node = node{
					Name: n.Name,
				}
				buff, _ := io.ReadAll(f)
				if !strings.Contains(string(buff), "Policy") {
					f.WriteString(parseTemplate("node.privacy.go.tmpl", data))
				}
				f.Close()
			}
		}

		cleanFiles(resolverDir, schemaDir)
		writeFiles(files)

		return next.Generate(g)
	})
}

func (e *extension) Hooks() []gen.Hook {
	return e.hooks
}

func NewExtension(opts ...extensionOption) *extension {
	ex := &extension{
		// Default Config
		config: &config{
			FileUpload:   false,
			Subscription: false,
			Mutation:     false,
			Database:     SQLite,
			Echo:         false,
			GqlGenPath:   "../gqlgen.yaml",
			JWT:          false,
		},
		hooks: []gen.Hook{},
	}

	for i := range opts {
		opts[i](ex)
	}

	gen.Funcs["camel"] = func(s string) string { return camel(snake(s)) }
	gen.Funcs["lower"] = strings.ToLower

	ex.config.GqlGen = readGqlGen(ex.config.GqlGenPath)
	ex.hooks = append(ex.hooks, ex.generate)
	return ex
}
