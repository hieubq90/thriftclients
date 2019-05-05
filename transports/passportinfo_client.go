package transports

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/hieubq90/thriftclients/core/auth/Passport"
	"github.com/hieubq90/thriftpool"
)

var (
	passportBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (Passport.NewTPassportServiceClient(c)) }),
		thriftpool.DefaultClose)

	passportCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (Passport.NewTPassportServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetPassportBinaryClient client by host:port
func GetPassportBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := passportBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetPassportCompactClient get compact client by host:port
func GetPassportCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := passportCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
