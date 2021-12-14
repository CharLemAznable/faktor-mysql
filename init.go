package faktor

import (
    "github.com/desertbit/grumble"
    "github.com/dolthub/go-mysql-server"
    "github.com/dolthub/go-mysql-server/auth"
    "github.com/dolthub/go-mysql-server/memory"
    "github.com/dolthub/go-mysql-server/server"
    "github.com/dolthub/go-mysql-server/sql/information_schema"
    "github.com/sirupsen/logrus"
    "os"
    "strconv"
)

var config = &grumble.Config{
    Name:        "faktor",
    Description: "fake mysql server",

    Flags: func(f *grumble.Flags) {
        f.Int("p", "port", 9900, "mysql server port")
        f.String("u", "username", "auser", "mysql user name")
        f.String("w", "password", "sa", "mysql password")
        f.StringL("log", "faktor.db.log", "fake database log output")
    },
}

var dataSourcePrefix = ""
var dataSourceSuffix = "?charset=utf8&multiStatements=true"
var dbProvider = memory.NewMemoryDBProvider(
    information_schema.NewInformationSchemaDatabase(),
)

var serverStarter = func(_ *grumble.App, flags grumble.FlagMap) error {
    logrus.SetOutput(newLogFile(flags.String("log")))
    username := flags.String("username")
    password := flags.String("password")
    port := strconv.Itoa(flags.Int("port"))

    dataSourcePrefix = username + ":" + password + "@tcp(127.0.0.1:" + port + ")/"
    cfg := server.Config{Protocol: "tcp", Address: "127.0.0.1:" + port,
        Auth: auth.NewNativeSingle(username, password, auth.AllPermissions)}
    engine := sqle.NewDefault(dbProvider)

    s, err := server.NewDefaultServer(cfg, engine)
    if nil == err {
        go func() {
            _ = s.Start()
        }()
    }
    return err
}

func newLogFile(filename string) *os.File {
    f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    return f
}
