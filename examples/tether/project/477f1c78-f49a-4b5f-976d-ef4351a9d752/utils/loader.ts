import { getBlockObjectPtr, getContractCallObjectPtr, getEventObjectPtr, utilsUUID } from "./external";
import { loadString } from "./";

export function getBlockInfo(): string {
	return loadString(getBlockObjectPtr());
}

export function getContractCallInfo(contractCallObjectPtr: i32): string {
	return loadString(getContractCallObjectPtr(contractCallObjectPtr));
}

export function getEventInfo(): string {
	return loadString(getEventObjectPtr());
}

export function getNewUUID(): string {
	return loadString(utilsUUID());
}
