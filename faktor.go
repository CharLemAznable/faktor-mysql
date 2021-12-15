package faktor

import (
    "github.com/desertbit/grumble"
)

var app = grumble.New(config)

func Run() {
    logo := func(a *grumble.App) {
        _, _ = a.Println("########    ###    ##    ## ########  #######  ########  ")
        _, _ = a.Println("##         ## ##   ##   ##     ##    ##     ## ##     ## ")
        _, _ = a.Println("##        ##   ##  ##  ##      ##    ##     ## ##     ## ")
        _, _ = a.Println("######   ##     ## #####       ##    ##     ## ########  ")
        _, _ = a.Println("##       ######### ##  ##      ##    ##     ## ##   ##   ")
        _, _ = a.Println("##       ##     ## ##   ##     ##    ##     ## ##    ##  ")
        _, _ = a.Println("##       ##     ## ##    ##    ##     #######  ##     ## ")
        _, _ = a.Println()
    }
    app.SetPrintASCIILogo(logo)
    app.SetPrintHelp(func(a *grumble.App, shell bool) {
        logo(a)
        _, _ = a.Println(a.Config().Description)
        _, _ = a.Println()
        _, _ = a.Println("Usage:")
        _, _ = a.Printf("  %s [flags]\n", a.Config().Name)
        _, _ = a.Println()
        _, _ = a.Println("Commands:")
        _, _ = a.Println("  load, create  load a schema [ with init sql script ]")
        _, _ = a.Println("  unload, drop  unload a schema")
        _, _ = a.Println("  help          use 'help [command]' for command help")
        _, _ = a.Println()
        _, _ = a.Println("Flags:")
        _, _ = a.Println("  -p, --port     int       mysql server port (default: 9900)")
        _, _ = a.Println("  -u, --username string    mysql user name (default: auser)")
        _, _ = a.Println("  -w, --password string    mysql password (default: sa)")
        _, _ = a.Println("      --log      string    fake database log output (default: faktor.db.log)")
        _, _ = a.Println("      --nocolor            disable color output")
        _, _ = a.Println("  -h, --help               display help")
        _, _ = a.Println()
    })

    app.OnInit(serverStarter)
    app.AddCommand(load)
    app.AddCommand(unload)

    grumble.Main(app)
}
