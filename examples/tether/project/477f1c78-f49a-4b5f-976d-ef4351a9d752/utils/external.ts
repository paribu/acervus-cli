// @ts-nocheck

@external("block", "getBlockInfo")
export declare function getBlockObjectPtr(): i32;

@external("event", "getEvent")
export declare function getEventObjectPtr(): i32;

@external("contractcall", "getContractCallInfo")
export declare function getContractCallObjectPtr(contractCallObjectPtr: i32): i32;

@external("log", "log")
export declare function logPrint(strPtr: i32): void;

@external("log", "error")
export declare function logPrintError(strPtr: i32): void;

@external("db", "query")
export declare function dbQuery(objectNamePtr: i32, filters: i32): i32;

@external("db", "save")
export declare function dbSave(dataStrPtr: i32, objectNamePtr: i32): void;

@external("db", "update")
export declare function dbUpdate(dataStrPtr: i32, objectNamePtr: i32): void;

@external("db", "delete")
export declare function dbDelete(idPtr: i32, objectNamePtr: i32): void;

@external("uuid", "generate")
export declare function utilsUUID(): i32;