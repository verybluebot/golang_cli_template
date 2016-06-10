package main

import(
    "fmt"
    "flag"
    "header"
    "errors"
    log "logger"
)


func showHelp() {
    fmt.Println(`

Usage: CLI Template [OPTIONS]

Options:
    -s, --string     Prints string input.       (default='some string').
    -i, --int        Prints int.                 (defualt=0).
    -e, --error      Prints a costum made error.
    -w, --warning    Prints the warning you entered.

    -h, --help       prints this help info.


    `)
}


func setFlag(flag *flag.FlagSet) {
    flag.Usage = func() {
        showHelp()
    }
}

func main() {
    var sStr string
    flag.StringVar(&sStr, "s", "", "")
    flag.StringVar(&sStr, "string", "", "")

    var sInt int
    flag.IntVar(&sInt, "i", 0, "")
    flag.IntVar(&sInt, "int", 0, "")

    var sErr bool
    flag.BoolVar(&sErr, "e", false, "")
    flag.BoolVar(&sErr, "error", false, "")

    var sWarn string
    flag.StringVar(&sWarn, "w", "", "")
    flag.StringVar(&sWarn, "warning", "", "")

    var sHelp bool
    flag.BoolVar(&sHelp, "h", false, "")
    flag.BoolVar(&sHelp, "help", false, "")

    setFlag(flag.CommandLine)

    flag.Parse()

    if sHelp {
        showHelp()
        return
    }

    header.GetHeader()

    log.InitLogger()

    if sStr != "" {
        log.Printf("This is user string input: %s", sStr)
    }

    if sInt != 0 {
        log.Printf("This is user integer input: %d", sInt)
    }

    if sErr {
        log.Errorf("Oh got some error: %s", errors.New("Something borke!!!"))
    }

    if sWarn != "" {
        log.Warningf("Your warning: %s", sWarn)
    }
}