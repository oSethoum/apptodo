package entgqlplus

import (
	"log"
	"strings"

	"entgo.io/ent/entc/gen"
)

type (
	templateData struct {
		Package         string
		Nodes           []node
		Node            node
		AuthNode        node
		HasSubscription bool
		Config          *config
	}

	config struct {
		Database     database
		DBConfig     []string
		Echo         bool
		JWT          bool
		Mutation     bool
		Privacy      bool
		FileUpload   bool
		Subscription bool
		GqlGenPath   string
		GqlGen       gqlGen
	}

	node struct {
		Name         string
		Subscription bool
		Auth         bool
	}

	file struct {
		Path   string
		Buffer string
	}

	gqlGen struct {
		Schema []string `yaml:"schema"`
		Exec   struct {
			FileName string `yaml:"filename"`
			Dir      string
			Package  string `yaml:"package"`
		} `yaml:"exec"`
		Model struct {
			FileName string `yaml:"filename"`
			Dir      string
			Package  string `yaml:"package"`
		} `yaml:"model"`
		Resolver struct {
			Dir     string `yaml:"dir"`
			Package string `yaml:"package"`
		} `yaml:"resolver"`
	}
)

func (t *templateData) parse(g *gen.Graph) {
	t.Package = strings.ReplaceAll(g.Package, "/ent", "")
	asCount := 0
	for i := range g.Nodes {
		a := &annotation{}
		a.decode(g.Nodes[i].Annotations[schemaAnnotationKey])
		n := node{Name: g.Nodes[i].Name}
		if a.SchemaOptions != nil {
			n.Subscription = inArray(a.SchemaOptions, Subscription)
			if n.Subscription {
				t.HasSubscription = true
			}
			as := inArray(a.SchemaOptions, AuthSchema)
			if as {
				asCount++
				if asCount > 1 {
					log.Fatalln("entgqlplus: can't use more than one AuthSchema")
				}
				t.AuthNode = n
			}
		}
		t.Nodes = append(t.Nodes, n)
	}

	if t.Config.JWT {
		if asCount == 0 {
			log.Fatalln("entgqlplus: No auth schema found, use entgqlplus.AuthSchema()")
		}
	}
}
