package main

//云厂商接入案例
import (
	"io"
	"sync"
)


type Interface interface {
	Instances() (Instances,bool)
	Zones() (Zones, bool)
}

type Factory func(config io.Reader) (Interface, error)

var providersMutex sync.Mutex
var providers = make(map[string]Factory)

func RegisterCloudProvider(name string, cloud Factory){
	providersMutex.Lock()
	defer providersMutex.Unlock()
	providers[name] = cloud
}

func GetCloudProvider(name string,config io.Reader) (Interface, error){
	providersMutex.Lock()
	defer providersMutex.Unlock()
	f, found:=providers[name]
	if !found {
		return nil,nil
	}
	return f(config)
}

func main(){
	RegisterCloudProvider("gce",newGCECloud)
}


type GCECloud struct {
	projectID string
	zone string
}

func newGCECloud(config io.Reader) (Interface,error){
	return &GCECloud{},nil
}

func (gce *GCECloud) Instances()(Instances,bool){
	return gce,true
}

func (gce *GCECloud) Zones() (Zones, bool){
	return gce,true
}
