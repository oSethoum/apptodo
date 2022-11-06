package entgqlplus

import "os"

type (
	extensionOption = func(*extension)
	database        = string
)

const (
	SQLite     database = "sqlite"
	MySQL      database = "mysql"
	PostgreSQL database = "postgres"
)

// WithMutation(b bool) enables entgqlplus to generate the Mutations.
// Default value is WithMutation(false).
func WithMutation(b bool) extensionOption {
	return func(e *extension) {
		e.config.Mutation = b
	}
}

// WithSubscription(b bool) enables entgqlplus to generate the Subscriptions.
// Works only if WithMutation(true) is enabled.
// Default value is WithSubscription(false).
func WithSubscription(b bool) extensionOption {
	return func(e *extension) {
		if e.config.Mutation {
			e.config.Subscription = b
		}
	}
}

// WithEchoServer(b bool) enables entgqlplus to generate the server, routes and the handlers.
// Default value is WithEchoServer(false).
func WithEchoServer(b bool) extensionOption {
	return func(e *extension) {
		if b && !e.config.Echo {
			e.config.Echo = b
		}
	}
}

// WithJWTAuth(b bool) enables entgqlplus to generate the login route and the Protected middleware
// Works only if WithEcho(true) is enabled.
// Default value is WithJWTAuth(false).
func WithJWTAuth(b bool) extensionOption {
	return func(e *extension) {
		if e.config.Echo {
			e.config.JWT = b
		}
	}
}

// WithDatabase(b Database) enables entgqlplus to generate the necessary code to connect to the database and migration.
// Default value is WithDatabase(entgql.SQLite).
func WithDatabase(d database, dbconfig ...string) extensionOption {
	return func(e *extension) {
		e.config.Database = d
		if d == SQLite {
			if len(dbconfig) == 1 {
				e.config.DBConfig = dbconfig
			} else {
				e.config.DBConfig = []string{"db"}
			}
		} else if d == MySQL || d == PostgreSQL {
			if len(dbconfig) == 3 {
				e.config.DBConfig = dbconfig
			} else {
				e.config.DBConfig = []string{"user", "pass", "db"}
			}
		}
	}
}

// WithConfigPath(p string) enables entgqlplus locate the gqlgen.yml config file.
// Default value is With WithConfigPath("../gqlgen.yml").
func WithConfigPath(p string) extensionOption {
	return func(e *extension) {
		_, err := os.Stat(p)
		catch(err)
		e.config.GqlGenPath = p
	}
}

// WithFileUpload(b bool) adds upload mutation.
// this only works if WithMutation(true) is enabled.
// Default is WithFileUpload(false).
func WithFileUpload(b bool) extensionOption {
	return func(e *extension) {
		if e.config.Mutation {
			e.config.FileUpload = b
		}
	}
}

// WithPrivacy(b bool) adds upload mutation.
// Default is WithPrivacy(false).
func WithPrivacy(b bool) extensionOption {
	return func(e *extension) {
		if e.config.Mutation {
			e.config.Privacy = b
		}
	}
}
