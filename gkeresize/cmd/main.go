package main
import (

  //"golang.org/x/net/context"
  //"google.golang.org/api/compute/v1"
  //"golang.org/x/oauth2/google"
  "gleb.ca/gkeresize"
  "fmt"
  "strconv"
  "log"
)
func main() {
  projects := [...]string{
    "gleb-sandbox",
    "asset-inventory-tester",
  }
  clusternames := [...]string{
    "us-sandbox-cluster-2",
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
    parent := fmt.Sprintf("projects/%s/locations/-", project)
    listClusters, err := gkeresize.ListClusters(parent) 
    if err != nil {
      log.Fatalln(err)
    }
    //fmt.Println(parent)
    //fmt.Printf("%#v\n",listClusters)
    //fmt.Println("Name:",listClusters.Clusters,)
    for i, cluster := range listClusters.Clusters {
      for _, clusters := range clusternames {
        if cluster.Name == clusters {
          fmt.Println("Number:",strconv.Itoa(i),
                      "Project:", project,
                      "cluster name:",cluster.Name,
                      "Location:", cluster.Location,
                      "Node count:", cluster.CurrentNodeCount,
        )
          for _, nodepool := range cluster.NodePools {
            fmt.Println("Node Pool Name:",nodepool.Name)
            parentsize := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", project,cluster.Location,cluster.Name,nodepool.Name)
            fmt.Println(parentsize)
            resizeClusters, err := gkeresize.ResizeClusters(parentsize,0)
            if err != nil {
              log.Fatalln(err)
            }
            fmt.Println(resizeClusters.Name)
          }
        }
      }
    }
  }
}