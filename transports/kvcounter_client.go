package transports

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/hieubq90/thriftclients/counter/KVStepCounter"
	thriftpool "github.com/hieubq90/thriftpool"
)

var (
	kvCounterMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (KVStepCounter.NewKVStepCounterServiceClient(c)) }),
		thriftpool.DefaultClose)

	kvCounterMapPoolCompact = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (KVStepCounter.NewKVStepCounterServiceClient(c)) }),
		thriftpool.DefaultClose)
)

// func init() {
// 	fmt.Println("init thrift kvcounter client ")
// }

//GetKVCounterBinaryClient client by host:port
func GetKVCounterBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := kvCounterMapPool.Get(aHost, aPort).Get()
	return client
}

//GetKVCounterCompactClient get compact client by host:port
func GetKVCounterCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := kvCounterMapPoolCompact.Get(aHost, aPort).Get()
	return client
}
