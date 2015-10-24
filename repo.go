package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"
  "log"
	"io/ioutil"
	"net/http"
	"encoding/json"
  "strings"
  "math/rand"
)
type msg struct {
	Pid        bson.ObjectId  `bson:"_id"`
  Id        int       `bson:"id"`
	Name      string    `bson:"name"`
	Address   string    `bson:"address"`
	City      string    `bson:"city"`
	State   	string    `bson:"state"`
	Zip      	string    		`bson:"zip"`
  Coordinate     Coordinate  `bson:"coordinate"`
}

func RepoAddLocation(l Location) msg{
  //uri := os.Getenv("MONGOHQ_URL")
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})

  collection := sess.DB("trip-planner").C("locations")
  pid:=bson.NewObjectId()
  var randno int
  randno=rand.Intn(1000)+rand.Intn(10000)
  doc := msg{Pid: pid,Id:randno,Name: l.Name,Address:l.Address,City:l.City,State:l.State,Zip:l.Zip,Coordinate:l.Coordinate}

  err = collection.Insert(doc)
  if err != nil {
    fmt.Printf("Can't insert document: %v\n", err)
    os.Exit(1)
  }
return doc
}
func RepoShowLocation(l int) Location{
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})

  var updatedmsg Location
  err = sess.DB("trip-planner").C("locations").Find(bson.M{"id": l}).One(&updatedmsg)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }

return updatedmsg
}
func RepoRemoveLocation(l int){
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})


  err = sess.DB("trip-planner").C("locations").Remove(bson.M{"id": l})
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }

}
func RepoUpdateLocation(l int,k Location) Location{
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})
  var updatedmsg Location
  var nupdatedmsg Location

  err = sess.DB("trip-planner").C("locations").Find(bson.M{"id": l}).One(&updatedmsg)
nupdatedmsg.Id=updatedmsg.Id
nupdatedmsg.Name=k.Name
nupdatedmsg.City=k.City
nupdatedmsg.State=k.State
nupdatedmsg.Zip=k.Zip
nupdatedmsg.Coordinate=updatedmsg.Coordinate
nupdatedmsg.Address=updatedmsg.Address


  err = sess.DB("trip-planner").C("locations").Update(updatedmsg,nupdatedmsg)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }
return nupdatedmsg
}
func QueryGMaps(l Location) Coordinate{
  var err error
  sym:=strings.Split(l.Address," ")
  var queryString string
  queryString=sym[0]
  for k:=1;k<len(sym);k++{
queryString=queryString+"+"+sym[k]
}
symmore:=strings.Split(l.City," ")
for k:=1;k<len(symmore);k++{
queryString=queryString+"+"+symmore[k]
}
queryString=queryString+"+"+l.State
url:="http://maps.google.com/maps/api/geocode/json?address="+queryString+"&sensor=false"
fmt.Println(url)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
  	log.Fatal(err)
  }
  req.SetBasicAuth("<token>", "x-oauth-basic")

  client := http.Client{}
  res, err := client.Do(req)
  if err != nil {
  	log.Fatal(err)
  }

  log.Println("StatusCode:", res.StatusCode)
  var dat map[string]interface{}
  // read body
  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
  	log.Fatal(err)
  }

  err = json.Unmarshal(body, &dat);if err != nil {
  panic(err)
  }
//  fmt.Println(dat)

  v:=dat["results"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["location"]

  //v:=dat.results["geometry"].(map[string]interface{})["location"]
  //fmt.Println(v)
  var nCord Coordinate
  nCord.Lat=v.(map[string]interface{})["lat"].(float64)
  nCord.Lng=v.(map[string]interface{})["lng"].(float64)

return nCord
}
