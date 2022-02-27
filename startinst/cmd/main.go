package main
import (

  //"golang.org/x/net/context"
  //"google.golang.org/api/compute/v1"
  //"golang.org/x/oauth2/google"
  "gleb.ca/startinstance"
  "fmt"
  "strconv"
  "log"
)
func main() {
  projects := [...]string{
    "gleb-sandbox",
    //"asset-inventory-tester",
  }
  //filters := [...]string{
  //  "status = RUNNING",
  //  "name != my-uninteresting-instance-one",
  //  "name != my-uninteresting-instance-two",
  //}
  //instances, err := listinstances.ListInstances(project) if err != nil { log.Fatalln(err) }
  instances := [...]string{
    "mssqlinst01std-02-gleb-sandbox-01",
  }

  for _, project := range projects {
    for _, instance := range instances {
      Instance, err := startinstance.StartInstance(project,instances)
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Println(Instance.Name)
    }
  }
}
