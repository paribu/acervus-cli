import {getEventObjectPtr, utilsUUID} from "./external";
import {loadString} from "./";

export function getEventInfo(): string {
	return loadString(getEventObjectPtr());
}

export function getNewUUID(): string {
	return loadString(utilsUUID());
}
