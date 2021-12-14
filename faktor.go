package faktor

import (
    "github.com/desertbit/grumble"
)

var app = grumble.New(config)

func Run() {
    app.SetPrintASCIILogo(func(a *grumble.App) {
        _, _ = a.Println("########    ###    ##    ## ########  #######  ########  ")
        _, _ = a.Println("##         ## ##   ##   ##     ##    ##     ## ##     ## ")
        _, _ = a.Println("##        ##   ##  ##  ##      ##    ##     ## ##     ## ")
        _, _ = a.Println("######   ##     ## #####       ##    ##     ## ########  ")
        _, _ = a.Println("##       ######### ##  ##      ##    ##     ## ##   ##   ")
        _, _ = a.Println("##       ##     ## ##   ##     ##    ##     ## ##    ##  ")
        _, _ = a.Println("##       ##     ## ##    ##    ##     #######  ##     ## ")
        _, _ = a.Println()
    })

    app.OnInit(serverStarter)
    app.AddCommand(load)
    app.AddCommand(unload)

    grumble.Main(app)
}
