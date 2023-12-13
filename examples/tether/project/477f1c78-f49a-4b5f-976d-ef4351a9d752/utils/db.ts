import {dbSave, dbUpdate, dbDelete, dbQuery} from "./external";
import {loadString, putString} from "./common";

export function queryFromDb(objectName: string, stringifiedFilters: string): string {
	const resPtr = dbQuery(putString(objectName), putString(stringifiedFilters));
	return loadString(resPtr);
}

export function saveToDb(dataStr: string, objectName: string): void {
	dbSave(putString(dataStr), putString(objectName));
}

export function updateToDb(dataStr: string, objectName: string): void {
	dbUpdate(putString(dataStr), putString(objectName));
}

export function deleteFromDb(id: string, objectName: string): void {
	dbDelete(putString(id), putString(objectName));
}