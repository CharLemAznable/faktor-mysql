package faktor

import (
    "github.com/desertbit/grumble"
    "github.com/dolthub/go-mysql-server/sql"
)

var unload = &grumble.Command{
    Name:    "unload",
    Help:    "unload a schema",
    Aliases: []string{"drop"},

    Args: func(a *grumble.Args) {
        a.String("schema", "schema name")
    },

    Run: func(c *grumble.Context) error {
        schema := c.Args.String("schema")
        return dbProvider.DropDatabase(sql.NewEmptyContext(), schema)
    },
}
