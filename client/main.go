package main

import (
    "context"
    "fmt"

    "github.com/ttsubo/sample-openapi/client/openapi"
)

func main() {
    cfg := openapi.NewConfiguration()
    c := openapi.NewAPIClient(cfg)

    ctx := context.Background()
    pets, _, err := c.DefaultApi.PetsIdGetExecute(c.DefaultApi.PetsIdGet(ctx, 1))

    fmt.Printf("[{id: %d}, ", *pets.Id)
    fmt.Printf("{name: %s}, ", *pets.Name)
    fmt.Printf("{status: %s}]\n", *pets.Status)
    fmt.Println(err)
}
