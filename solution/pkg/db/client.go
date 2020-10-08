package db

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"github.com/CoufalJa/go-workshop/pkg/model"
	"github.com/rainycape/memcache"
	"time"
)

type MemcacheClient interface {
	Get(hashedResource []byte, resource string, resultChain chan model.SecurityDefinition)
}

type memcacheClient struct {
	client *memcache.Client
}

func (m *memcacheClient) Get(hashedResource []byte, resource string, resultChain chan model.SecurityDefinition) {
	domainKey := fmt.Sprintf(	"@@%s.%x", "DOMAIN", string(hashedResource))
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

	resultChain <- securityDefinition
}

func NewMemcacheClient(server string, requestTimeout time.Duration) MemcacheClient {
	mc, err := memcache.New(server)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to memcache")

	mc.SetTimeout(requestTimeout)
	fmt.Println("Request timeout set to ", requestTimeout)

	clientInstance := &memcacheClient{
		client: mc,
	}
	return clientInstance
}


