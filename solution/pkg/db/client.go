package db

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fdvoracek/go-heroes/solution/pkg/model"
	"github.com/rainycape/memcache"
	"time"
)

type MemcacheClient interface {
	Get(hashedResource []byte, resource string) model.SecurityDefinition
}

type memcacheClient struct {
	client *memcache.Client
}

func (m *memcacheClient) Get(hashedResource []byte, resource string) model.SecurityDefinition {
	domainKey := fmt.Sprintf(	"@@%s.%x", "DOMAIN_ETL", string(hashedResource))
	//fmt.Printf("About to lookup payload under key %s\n", domainKey)

	item, err := m.client.Get(domainKey)
	if err != nil {
		//fmt.Printf("Error occurred while getting key %s due to:\n'%s'\n", domainKey, err)
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

func NewMemcacheClient(server string, requestTimeout time.Duration) MemcacheClient {
	mc, err := memcache.New(server)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to memcache")

	mc.SetTimeout(requestTimeout)
	fmt.Println("Request timeout set to", requestTimeout)

	clientInstance := &memcacheClient{
		client: mc,
	}
	return clientInstance
}


