package main

import (
	"fmt"
	"time"
)

type DataProvider interface {
	GetData(id string) string
}
type RealDateProvider struct {}

func (rdp *RealDateProvider)GetData(id string) string{
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Real data for Id: %s",id)
}

type CacheDataProvider struct {
	provider DataProvider
	cache map[string]string
}

func(cdp *CacheDataProvider)GetData(id string) string{
	data,ok:=cdp.cache[id]
	if !ok{
		data = cdp.provider.GetData(id)
		cdp.cache[id] = data
	}
	return data
}

func NewCacheDataProvider(provider DataProvider) *CacheDataProvider{
	return &CacheDataProvider{
		provider: provider,
		cache: make(map[string]string),
	}
}

func main(){
	realProvider := &RealDateProvider{}
	cacheProvider := NewCacheDataProvider(realProvider)

	startTime := time.Now()
	data := cacheProvider.GetData("1")
	fmt.Println("Data:", data)
	fmt.Println("TIme spent:",time.Since(startTime))

	startTime = time.Now()
	data = cacheProvider.GetData("1")
	fmt.Println("Data:",data)
	fmt.Println("Time spent:",time.Since(startTime))
}
