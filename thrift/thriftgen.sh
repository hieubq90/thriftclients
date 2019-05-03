thrift -r --gen go -out ../ kvstepcounter.thrift
rm -Rf ../counter/KVStepCounter/*-remote
thrift -r --gen go -out ../ s2i64kv.thrift
rm -Rf ../common/S2I64KV/*-remote
thrift -r --gen go -out ../ bigsetlistint.thrift
rm -Rf ../bigset/ListI64/*-remote
thrift -r --gen go -out ../ bigsetgenericdata.thrift
rm -Rf ../bigset/Generic/*-remote