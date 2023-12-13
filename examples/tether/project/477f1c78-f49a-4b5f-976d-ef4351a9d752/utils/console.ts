import { logPrint, logPrintError } from "./external";
import { putString } from "./common"

export namespace console {
  export function log(str: string): void {
    logPrint(putString(str));
  }

  export function error(str: string): void {
    logPrintError(putString(str));
  }
}   