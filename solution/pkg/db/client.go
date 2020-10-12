package db

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fdvoracek/go-heroes/pkg/model"
	"github.com/rainycape/memcache"
	"time"
)

type MemcacheClient interface {
	Get(hashedResource []byte, resource string) model.SecurityDefinition
	GetWithChan(hashedResource []byte, resource string, channel chan model.SecurityDefinition)
}

type memcacheClient struct {
	client *memcache.Client
}

func (m *memcacheClient) Get(hashedResource []byte, resource string) model.SecurityDefinition {
	domainKey := fmt.Sprintf(	"@@%s.%x", "DOMAIN_ETL", string(hashedResource))
	//fmt.Printf("About to lookup payload under key %s\n", domainKey)

	item, err := m.client.Get(domainKey)
	if err != nil {
		fmt.Printf("Error occurred while getting key %s due to:\n'%s'\n", domainKey, err)
	}

	var securityDefinition model.SecurityDefinition
	if (item != nil) && (len(item.Value) > 0) {
		err := json.Unmarshal(item.Value, &securityDefinition)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Found %+v\n", securityDefinition)
	} else {
		securityDefinition = model.NewSecurityDefinition(hex.EncodeToString(hashedResource), resource)
		//fmt.Printf("Returning default (CLEAN) result %+v\n", securityDefinition)
	}

	return securityDefinition
}

func (m *memcacheClient) GetWithChan(hashedResource []byte, resource string, channel chan model.SecurityDefinition)  {
	domainKey := fmt.Sprintf(	"@@%s.%x", "DOMAIN_ETL", string(hashedResource))
	//fmt.Printf("About to lookup payload under key %s\n", domainKey)

	item, err := m.client.Get(domainKey)
	if err != nil {
		fmt.Printf("Error occurred while getting key %s due to:\n'%s'\n", domainKey, err)
	}

	var securityDefinition model.SecurityDefinition
	if (item != nil) && (len(item.Value) > 0) {
		err := json.Unmarshal(item.Value, &securityDefinition)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Found %+v\n", securityDefinition)
	} else {
		securityDefinition = model.NewSecurityDefinition(hex.EncodeToString(hashedResource), resource)
		//fmt.Printf("Returning default (CLEAN) result %+v\n", securityDefinition)
	}

	channel <- securityDefinition
}

func NewMemcacheClient(server string, requestTimeout time.Duration, maxIdle int) MemcacheClient {
	mc, err := memcache.New(server)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to memcache")

	mc.SetTimeout(requestTimeout)
	mc.SetMaxIdleConnsPerAddr(maxIdle)
	fmt.Println("Request timeout set to", requestTimeout)
	fmt.Println("Max idle set to", maxIdle)

	clientInstance := &memcacheClient{
		client: mc,
	}
	return clientInstance
}


