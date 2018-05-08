package main

import (
  "fmt"
)

func main() {
  var var1 string = "Ha"
  fmt.Println(var1);
  var var2 = "Hyun" // Type can be omitted.
  fmt.Println(var2);
  var3 := "Soo"    // when we use ':' var can be omitted
  fmt.Println(var3);

  var ( // Can declare at once
    site    string = "github.com"
    name    string = "CryptoSalamander"
    job     string = "blockchain Core Developer"
  )

    fmt.Println(site);
    fmt.Println(name);
    fmt.Println(job);
}