thrift -r --gen go -out ../ kvstepcounter.thrift
rm -Rf ../counter/KVStepCounter/*-remote
thrift -r --gen go -out ../ s2i64kv.thrift
rm -Rf ../common/S2I64KV/*-remote