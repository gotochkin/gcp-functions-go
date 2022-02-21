package main
import (

  //"golang.org/x/net/context"
  //"google.golang.org/api/compute/v1"
  //"golang.org/x/oauth2/google"
  "gleb.ca/listinstances"
  "fmt"
  "strconv"
  //"strings"
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

  for _, project := range projects {
    listInstances, err := listinstances.ListInstances(project) 
    if err != nil {
      log.Fatalln(err)
    }
    for i, instance := range listInstances {
      //m := strings(instance)
      fmt.Printf("Number: %s: instance name: %s",strconv.Itoa(i),instance.Name)
    }
  }
}
