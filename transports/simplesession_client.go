package transports

import (
	"github.com/apache/thrift/lib/go/thrift"
	ss "github.com/hieubq90/thriftclients/core/auth/SimpleSession"
	thriftpool "github.com/hieubq90/thriftpool"
)

var (
	ssMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (ss.NewTSimpleSessionService_WClient(c)) }),
		thriftpool.DefaultClose)

	ssMapPoolCompact = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (ss.NewTSimpleSessionService_WClient(c)) }),
		thriftpool.DefaultClose)
)

//GetSimpleSessionBinaryClient Get binary client by host:port
func GetSimpleSessionBinaryClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient {
	client, _ := ssMapPool.Get(bsHost, bsPort).Get()
	return client
}

//GetSimpleSessionCompactClient Get compact client by host:port
func GetSimpleSessionCompactClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient {
	client, _ := ssMapPoolCompact.Get(bsHost, bsPort).Get()
	return client
}
