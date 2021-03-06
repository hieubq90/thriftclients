package transports

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/hieubq90/thriftclients/bigset/Generic"
	"github.com/hieubq90/thriftpool"
)

var (
	bsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (Generic.NewTStringBigSetKVServiceClient(c)) }),
		thriftpool.DefaultClose)

	ibsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (Generic.NewTIBSDataServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetBsGenericClient client by host:port
func GetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := bsGenericMapPool.Get(aHost, aPort).Get()
	return client
}

//GetIBsGenericClient client by host:port
func GetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := ibsGenericMapPool.Get(aHost, aPort).Get()
	return client
}
