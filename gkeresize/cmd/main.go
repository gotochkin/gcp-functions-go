package main
import (

  //"golang.org/x/net/context"
  //"google.golang.org/api/compute/v1"
  //"golang.org/x/oauth2/google"
  "gleb.ca/gkeresize"
  "fmt"
  //"strconv"
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
  //instances := [...]string{
  //  "mssqlinst01std-02-gleb-sandbox-01",
  //}

  for _, project := range projects {
    listClusters, err := resizegke.ListClusters(project) 
    if err != nil {
      log.Fatalln(err)
    }
    for i, cluster := range listClusters {
      //fmt.Printf("Number: %s instance name: %s\n",strconv.Itoa(i),instance.Name)
      fmt.Println("Number:",strconv.Itoa(i),
                  //"instance name:",instance.Name,
                  //"State:", instance.State,
                  //"Settings.ActivationPolicy:", instance.Settings.ActivationPolicy,
                  //"Settings.SettingsVersion:", instance.Settings.SettingsVersion,
                )
      //fmt.Printf("%#v\n",instance.BackupConfiguration)
    }
  }
}