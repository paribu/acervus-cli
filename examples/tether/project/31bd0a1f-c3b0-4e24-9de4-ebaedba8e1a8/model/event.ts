import { JSON } from "../lib/json";
import { BigInt } from "../lib/bigint";
import { getEventInfo } from "../utils/loader";
import { console } from "../utils";

type ParamName = string;
export class EthereumEventParameter {
	constructor(
		public name: string,
		public value: string,
		public type: string
	) {}

	private getValuesAsArr(value: string, key: string = "values"): JSON.Arr {
		const valuesAsJSON: JSON.Obj = <JSON.Obj>JSON.parse(value);
		const valuesAsArr = valuesAsJSON.getArr(key);

		return valuesAsArr != null ? valuesAsArr : new JSON.Arr();
	}

	asBigInt(): BigInt {
		return BigInt.fromString(this.value);
	}

	asInt8(): i8 {
		return i8.parse(this.value);
	}

	asInt16(): i16 {
		return i16.parse(this.value);
	}

	asInt32(): i32 {
		return i32.parse(this.value);
	}

	asInt64(): i64 {
		return i64.parse(this.value);
	}

	asUInt8(): u8 {
		return u8.parse(this.value);
	}

	asUInt16(): u16 {
		return u16.parse(this.value);
	}

	asUInt32(): u32 {
		return i32.parse(this.value);
	}

	asUInt64(): u64 {
		return u64.parse(this.value);
	}

	asString(): string {
		return this.value;
	}

	asBool(): bool {
		return this.value == "true";
	}

	asInt8Array(): Array<i8> {
		const arrayOfInt: i8[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = i8.parse(valuesArr[i].toString());
			arrayOfInt.push(val);
		}

		return arrayOfInt;
	}

	asInt16Array(): Array<i16> {
		const arrayOfInt: i16[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = i16.parse(valuesArr[i].toString());
			arrayOfInt.push(val);
		}

		return arrayOfInt;
	}

	asInt32Array(): Array<i32> {
		const arrayOfInt: i32[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = i32.parse(valuesArr[i].toString());
			arrayOfInt.push(val);
		}

		return arrayOfInt;
	}

	asInt64Array(): Array<i64> {
		const arrayOfInt: i64[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = i64.parse(valuesArr[i].toString());
			arrayOfInt.push(val);
		}

		return arrayOfInt;
	}

	asUInt8Array(): Array<u8> {
		const arrayOfUInt: u8[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = u8.parse(valuesArr[i].toString());
			arrayOfUInt.push(val);
		}

		return arrayOfUInt;
	}

	asUInt16Array(): Array<u16> {
		const arrayOfUInt: u16[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = u16.parse(valuesArr[i].toString());
			arrayOfUInt.push(val);
		}

		return arrayOfUInt;
	}

	asUInt32Array(): Array<u32> {
		const arrayOfUInt: u32[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = u32.parse(valuesArr[i].toString());
			arrayOfUInt.push(val);
		}

		return arrayOfUInt;
	}

	asUInt64Array(): Array<u64> {
		const arrayOfUInt: u64[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = u64.parse(valuesArr[i].toString());
			arrayOfUInt.push(val);
		}

		return arrayOfUInt;
	}

	asStringArray(): Array<string> {
		const arrayOfString: string[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = (<JSON.Str>valuesArr[i]).valueOf();
			arrayOfString.push(val);
		}

		return arrayOfString;
	}

	asBigIntArray(): Array<BigInt> {
		const arrayOfBigInt: BigInt[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = BigInt.fromString(valuesArr[i].toString());
			arrayOfBigInt.push(val);
		}

		return arrayOfBigInt;
	}

	asBoolArray(): Array<bool> {
		const arrayOfBool: bool[] = [];
		const valuesAsArr = this.getValuesAsArr(this.value);
		const valuesArr = valuesAsArr.valueOf();
		for (let i = 0; i < valuesArr.length; i++) {
			let val = (<JSON.Bool>valuesArr[i]).valueOf();
			arrayOfBool.push(val);
		}

		return arrayOfBool;
	}

	asInt8Matrix(): Array<Array<i8>> {
		const matrixOfInt: i8[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfInt: i8[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = i8.parse(valuesArr[j].toString());
				arrayOfInt.push(val);
			}

			matrixOfInt.push(arrayOfInt);
		}

		return matrixOfInt;
	}

	asInt16Matrix(): Array<Array<i16>> {
		const matrixOfInt: i16[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfInt: i16[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = i16.parse(valuesArr[j].toString());
				arrayOfInt.push(val);
			}

			matrixOfInt.push(arrayOfInt);
		}

		return matrixOfInt;
	}

	asInt32Matrix(): Array<Array<i32>> {
		const matrixOfInt: i32[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfInt: i32[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = i32.parse(valuesArr[j].toString());
				arrayOfInt.push(val);
			}

			matrixOfInt.push(arrayOfInt);
		}

		return matrixOfInt;
	}

	asInt64Matrix(): Array<Array<i64>> {
		const matrixOfInt: i64[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfInt: i64[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = i64.parse(valuesArr[j].toString());
				arrayOfInt.push(val);
			}

			matrixOfInt.push(arrayOfInt);
		}

		return matrixOfInt;
	}

	asUInt8Matrix(): Array<Array<u8>> {
		const matrixOfUInt: u8[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfUInt: u8[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = u8.parse(valuesArr[j].toString());
				arrayOfUInt.push(val);
			}

			matrixOfUInt.push(arrayOfUInt);
		}

		return matrixOfUInt;
	}

	asUInt16Matrix(): Array<Array<u16>> {
		const matrixOfUInt: u16[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfUInt: u16[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = u16.parse(valuesArr[j].toString());
				arrayOfUInt.push(val);
			}

			matrixOfUInt.push(arrayOfUInt);
		}

		return matrixOfUInt;
	}

	asUInt32Matrix(): Array<Array<u32>> {
		const matrixOfUInt: u32[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfUInt: u32[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = u32.parse(valuesArr[j].toString());
				arrayOfUInt.push(val);
			}

			matrixOfUInt.push(arrayOfUInt);
		}

		return matrixOfUInt;
	}

	asUInt64Matrix(): Array<Array<u64>> {
		const matrixOfUInt: u64[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfUInt: u64[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = u64.parse(valuesArr[j].toString());
				arrayOfUInt.push(val);
			}

			matrixOfUInt.push(arrayOfUInt);
		}

		return matrixOfUInt;
	}

	asStringMatrix(): Array<Array<string>> {
		const matrixOfString: string[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfString: string[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = (<JSON.Str>valuesArr[j]).valueOf();
				arrayOfString.push(val);
			}

			matrixOfString.push(arrayOfString);
		}

		return matrixOfString;
	}

	asBigIntMatrix(): Array<Array<BigInt>> {
		const matrixOfBigInt: BigInt[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfBigInt: BigInt[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = BigInt.fromString(valuesArr[j].toString());
				arrayOfBigInt.push(val);
			}

			matrixOfBigInt.push(arrayOfBigInt);
		}

		return matrixOfBigInt;
	}

	asBoolMatrix(): Array<Array<bool>> {
		const matrixOfBool: bool[][] = [];
		const matrixValuesAsArr = this.getValuesAsArr(this.value);
		const matrixValuesArr = matrixValuesAsArr.valueOf();
		for (let i = 0; i < matrixValuesArr.length; i++) {
			const arrayOfBool: bool[] = [];
			const valuesArr = (<JSON.Arr>matrixValuesArr[i]).valueOf();
			for (let j = 0; j < valuesArr.length; j++) {
				let val = (<JSON.Bool>valuesArr[j]).valueOf();
				arrayOfBool.push(val);
			}

			matrixOfBool.push(arrayOfBool);
		}

		return matrixOfBool;
	}
}

export class EthereumEvent {
	public eventID: u32;
	public address: string;
	public blockNumber: u32;
	public blockHash: string;
	public transactionHash: string;
	public logIndex: u32;
	public parameters: Map<ParamName, EthereumEventParameter>;

	constructor(
		eventID: u32,
		address: string,
		blockNumber: u32,
		blockHash: string,
		transactionHash: string,
		logIndex: u32,
		parameters: Map<ParamName, EthereumEventParameter>
	) {
		this.eventID = eventID;
		this.address = address;
		this.blockNumber = blockNumber;
		this.blockHash = blockHash;
		this.transactionHash = transactionHash;
		this.logIndex = logIndex;
		this.parameters = parameters;
	}

	getParameter(name: string): EthereumEventParameter | null {
		return this.parameters.get(name) || null;
	}
}

export function GetEthereumEvent(): EthereumEvent {
	const eventAsString = getEventInfo();
	const eventAsJSON: JSON.Obj = <JSON.Obj>JSON.parse(eventAsString);

	if (eventAsJSON.has("error_message")) {
		const jsonErrorMessage = eventAsJSON.getString("error_message");
		const errorMessage = jsonErrorMessage !== null ? jsonErrorMessage.valueOf() : "";

		console.log(errorMessage);
		throw new Error(errorMessage);
	}

	const jsonEventID = eventAsJSON.getInteger("id");
	const eventID: u32 = jsonEventID != null ? <u32>jsonEventID.valueOf() : 0;

	const jsonAddress = eventAsJSON.getString("address");
	const address = jsonAddress != null ? jsonAddress.valueOf() : "";

	const jsonBlockNumber = eventAsJSON.getInteger("block_number");
	const blockNumber: u32 = jsonBlockNumber != null ? <u32>jsonBlockNumber.valueOf() : 0;

	const jsonBlockHash = eventAsJSON.getString("block_hash");
	const blockHash = jsonBlockHash != null ? jsonBlockHash.valueOf() : "";

	const jsonTransactionHash = eventAsJSON.getString("transaction_hash");
	const transactionHash = jsonTransactionHash != null ? jsonTransactionHash.valueOf() : "";

	const jsonLogIndex = eventAsJSON.getInteger("log_index");
	const logIndex: u32 = jsonLogIndex != null ? <u32>jsonLogIndex.valueOf() : 0;

	const parametersMap: Map<ParamName, EthereumEventParameter> = new Map<
		ParamName,
		EthereumEventParameter
	>();
	const parametersAsArr = eventAsJSON.getArr("parameters");

	if (parametersAsArr != null) {
		const parametersArr = parametersAsArr.valueOf();
		for (let i = 0; i < parametersArr.length; i++) {
			const parameterAsJSON: JSON.Obj = <JSON.Obj>parametersArr[i];

			const jsonParameterName = parameterAsJSON.getString("name");
			const parameterName = jsonParameterName != null ? jsonParameterName.valueOf() : "";

			const jsonParameterValue = parameterAsJSON.getString("value");
			const parameterValue = jsonParameterValue != null ? jsonParameterValue.valueOf() : "";

			const jsonParameterType = parameterAsJSON.getString("type");
			const parameterType = jsonParameterType != null ? jsonParameterType.valueOf() : "";

			parametersMap.set(
				parameterName,
				new EthereumEventParameter(parameterName, parameterValue, parameterType)
			);
		}
	}

	const event = new EthereumEvent(
		eventID,
		address,
		blockNumber,
		blockHash,
		transactionHash,
		logIndex,
		parametersMap
	);

	return event;
}
