package casbinx

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	pgxadapter "github.com/pckhoi/casbin-pgx-adapter"
	"os"
)

const (
	GRPC = "gRPC"
)

func NewCasbinX(modelPolicy string) *casbin.Enforcer {
	// Initialize a PGX adapter and use it in a Casbin enforcer:
	// The adapter will use the Postgres database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := pgxadapter.NewAdapter(
		os.Getenv("PQ_CONNECTION"),
		pgxadapter.WithDatabase(os.Getenv("PQ_DATABASE")),
		pgxadapter.WithTableName("rbacs"),
		pgxadapter.WithSkipTableCreate(),
	)

	// Load policy model from text
	m, err := model.NewModelFromString(modelPolicy)
	if err != nil {
		fmt.Println("Failed to load policy from text: ", err)
	}

	// The adapter will use the table named "casbin_rule" by default.
	// If it doesn't exist, the adapter will create it automatically.
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		fmt.Println("Failed to new enforcer: ", err)
	}

	// Enable logs
	e.EnableLog(true)

	// Load the policy from the Database
	if err = e.LoadPolicy(); err != nil {
		fmt.Println("Failed to load policy: ", err)
	}

	return e
}
