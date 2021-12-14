package faktor

import (
    "github.com/CharLemAznable/sqlx"
    "github.com/desertbit/grumble"
    "github.com/dolthub/go-mysql-server/sql"
    "io/ioutil"

    _ "github.com/go-sql-driver/mysql"
)

var load = &grumble.Command{
    Name:    "load",
    Help:    "load a schema [ with init sql script ]",
    Aliases: []string{"create"},

    Args: func(a *grumble.Args) {
        a.String("schema", "schema name")
        a.String("script", "schema init script", grumble.Default(""))
    },

    Run: func(c *grumble.Context) error {
        schema := c.Args.String("schema")
        if err := dbProvider.CreateDatabase(
            sql.NewEmptyContext(), schema); nil != err {
            return err
        }

        script := c.Args.String("script")
        if "" == script {
            return nil
        }
        source, err := ioutil.ReadFile(script)
        if nil != err {
            return err
        }

        db, err := sqlx.Open("mysql",
            dataSourcePrefix+schema+dataSourceSuffix)
        if nil != err {
            return err
        }
        defer func() { _ = db.Close() }()

        if err = db.Ping(); err != nil {
            return err
        }
        _, err = db.ExecX(string(source))
        return err
    },
}
