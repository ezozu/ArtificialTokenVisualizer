// cmd/artificialtokenvisualizer/main.go
package main

import (
"flag"
"log"
"os"

"artificialtokenvisualizer/internal/artificialtokenvisualizer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialtokenvisualizer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
