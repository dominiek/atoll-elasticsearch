
package main

import (
  "os"
  "log"
  "fmt"
  "github.com/codegangsta/cli"
)

func fatalError(err error) {
  log.Fatalf("Error: %v", err)
}

func main() {
  app := cli.NewApp()
  app.Name = "atoll-elasticsearch"
  app.Usage = "Elasticsearch monitoring plugin for Atoll"
  app.Version = "0.1.1"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "host",
      Value: "localhost",
      Usage: "ES host",
    },
    cli.IntFlag{
      Name: "port",
      Value: 9200,
      Usage: "ES port",
    },
  }
  app.Action = func(c *cli.Context) {
    elasticsearch := Elasticsearch{c.String("host"), uint16(c.Int("port"))};
    data, err := elasticsearch.Monitor();
    if err != nil {
      fatalError(err);
    } else {
      fmt.Printf("%s", data);
    }
  }

  app.Run(os.Args)
}
