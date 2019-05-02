namespace cpp counter.KVStepCounter
namespace go counter.KVStepCounter
namespace java counter.KVStepCounter

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TKVCounter{
    1: TKey cid
    2: i64 value=1,
    3: i64 hvalue = 0, //when value exceeds 0, hvalue increase by 1
}

typedef TKVCounter TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TKVCounter data

}

service TDataServiceR{
    TDataResult getData(1: TKey key),
}

service TDataService{
    TDataResult getData(1: TKey key),
    TErrorCode putData(1: TKVCounter data)
}

exception InvalidOperation {
	1: i32 errorCode;
	2: string message;
}

service KVStepCounterService extends TDataService
{
    i32 createGenerator(1: string genName),
    i32 removeGenerator(1: string genName),
    i64 getCurrentValue(1: string genName),
    i64 getValue(1: string genName) ,
    i64 getStepValue(1: string genName, 2: i64 step),
}
