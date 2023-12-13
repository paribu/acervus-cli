import { BigInt } from "../lib/bigint";
import {
	queryFromDb,
	saveToDb,
	updateToDb,
	deleteFromDb,
	console,
	convertToString,
	generateUUID
} from "../utils";

import { JSON } from "../lib/json";

import { filter } from "./filter";
export namespace schema {
export class IssueEventSchema extends filter.IssueEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.IssueEventSchemaFilter[]): IssueEventSchema[] {
		const stringifiedRes = queryFromDb(
			"IssueEventSchema",
			IssueEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: IssueEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new IssueEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const amount = jsonObj.getNum("amount");
			if (amount != null) res[i].Amount = BigInt.from(amount.valueOf());
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "IssueEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "IssueEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "IssueEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "amount") {
				if(str) { str = convertToString(this.amount); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.IssueEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetAmount) jsonObj.set("amount", filter.Amount);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("amount", convertToString(this.amount));

		return jsonObj.toString();
	}
}

export class RedeemEventSchema extends filter.RedeemEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.RedeemEventSchemaFilter[]): RedeemEventSchema[] {
		const stringifiedRes = queryFromDb(
			"RedeemEventSchema",
			RedeemEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: RedeemEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new RedeemEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const amount = jsonObj.getNum("amount");
			if (amount != null) res[i].Amount = BigInt.from(amount.valueOf());
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "RedeemEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "RedeemEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "RedeemEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "amount") {
				if(str) { str = convertToString(this.amount); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.RedeemEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetAmount) jsonObj.set("amount", filter.Amount);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("amount", convertToString(this.amount));

		return jsonObj.toString();
	}
}

export class DeprecateEventSchema extends filter.DeprecateEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.DeprecateEventSchemaFilter[]): DeprecateEventSchema[] {
		const stringifiedRes = queryFromDb(
			"DeprecateEventSchema",
			DeprecateEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: DeprecateEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new DeprecateEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const newAddress = jsonObj.getString("newAddress");
			if (newAddress != null) res[i].NewAddress = newAddress.valueOf();
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "DeprecateEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "DeprecateEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "DeprecateEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "newAddress") {
				if(str) { str = convertToString(this.newAddress); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.DeprecateEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetNewAddress) jsonObj.set("newAddress", filter.NewAddress);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("newAddress", this.newAddress);

		return jsonObj.toString();
	}
}

export class ParamsEventSchema extends filter.ParamsEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.ParamsEventSchemaFilter[]): ParamsEventSchema[] {
		const stringifiedRes = queryFromDb(
			"ParamsEventSchema",
			ParamsEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: ParamsEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new ParamsEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const feeBasisPoints = jsonObj.getNum("feeBasisPoints");
			if (feeBasisPoints != null) res[i].FeeBasisPoints = BigInt.from(feeBasisPoints.valueOf());

			const maxFee = jsonObj.getNum("maxFee");
			if (maxFee != null) res[i].MaxFee = BigInt.from(maxFee.valueOf());
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "ParamsEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "ParamsEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "ParamsEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "feeBasisPoints") {
				if(str) { str = convertToString(this.feeBasisPoints); }
			} else if (propName == "maxFee") {
				if(str) { str = convertToString(this.maxFee); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.ParamsEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetFeeBasisPoints) jsonObj.set("feeBasisPoints", filter.FeeBasisPoints);

			if (filter.IsSetMaxFee) jsonObj.set("maxFee", filter.MaxFee);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("feeBasisPoints", convertToString(this.feeBasisPoints));
		jsonObj.set("maxFee", convertToString(this.maxFee));

		return jsonObj.toString();
	}
}

export class DestroyedBlackFundsEventSchema extends filter.DestroyedBlackFundsEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.DestroyedBlackFundsEventSchemaFilter[]): DestroyedBlackFundsEventSchema[] {
		const stringifiedRes = queryFromDb(
			"DestroyedBlackFundsEventSchema",
			DestroyedBlackFundsEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: DestroyedBlackFundsEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new DestroyedBlackFundsEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const _blackListedUser = jsonObj.getString("_blackListedUser");
			if (_blackListedUser != null) res[i].BlackListedUser = _blackListedUser.valueOf();

			const _balance = jsonObj.getNum("_balance");
			if (_balance != null) res[i].Balance = BigInt.from(_balance.valueOf());
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "DestroyedBlackFundsEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "DestroyedBlackFundsEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "DestroyedBlackFundsEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "_blackListedUser") {
				if(str) { str = convertToString(this._blackListedUser); }
			} else if (propName == "_balance") {
				if(str) { str = convertToString(this._balance); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.DestroyedBlackFundsEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetBlackListedUser) jsonObj.set("_blackListedUser", filter.BlackListedUser);

			if (filter.IsSetBalance) jsonObj.set("_balance", filter.Balance);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("_blackListedUser", this._blackListedUser);
		jsonObj.set("_balance", convertToString(this._balance));

		return jsonObj.toString();
	}
}

export class AddedBlackListEventSchema extends filter.AddedBlackListEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.AddedBlackListEventSchemaFilter[]): AddedBlackListEventSchema[] {
		const stringifiedRes = queryFromDb(
			"AddedBlackListEventSchema",
			AddedBlackListEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: AddedBlackListEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new AddedBlackListEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const _user = jsonObj.getString("_user");
			if (_user != null) res[i].User = _user.valueOf();
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "AddedBlackListEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "AddedBlackListEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "AddedBlackListEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "_user") {
				if(str) { str = convertToString(this._user); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.AddedBlackListEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetUser) jsonObj.set("_user", filter.User);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("_user", this._user);

		return jsonObj.toString();
	}
}

export class RemovedBlackListEventSchema extends filter.RemovedBlackListEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.RemovedBlackListEventSchemaFilter[]): RemovedBlackListEventSchema[] {
		const stringifiedRes = queryFromDb(
			"RemovedBlackListEventSchema",
			RemovedBlackListEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: RemovedBlackListEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new RemovedBlackListEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const _user = jsonObj.getString("_user");
			if (_user != null) res[i].User = _user.valueOf();
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "RemovedBlackListEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "RemovedBlackListEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "RemovedBlackListEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "_user") {
				if(str) { str = convertToString(this._user); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.RemovedBlackListEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetUser) jsonObj.set("_user", filter.User);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("_user", this._user);

		return jsonObj.toString();
	}
}

export class ApprovalEventSchema extends filter.ApprovalEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.ApprovalEventSchemaFilter[]): ApprovalEventSchema[] {
		const stringifiedRes = queryFromDb(
			"ApprovalEventSchema",
			ApprovalEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: ApprovalEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new ApprovalEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const owner = jsonObj.getString("owner");
			if (owner != null) res[i].Owner = owner.valueOf();

			const spender = jsonObj.getString("spender");
			if (spender != null) res[i].Spender = spender.valueOf();

			const value = jsonObj.getNum("value");
			if (value != null) res[i].Value = BigInt.from(value.valueOf());
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "ApprovalEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "ApprovalEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "ApprovalEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "owner") {
				if(str) { str = convertToString(this.owner); }
			} else if (propName == "spender") {
				if(str) { str = convertToString(this.spender); }
			} else if (propName == "value") {
				if(str) { str = convertToString(this.value); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.ApprovalEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetOwner) jsonObj.set("owner", filter.Owner);

			if (filter.IsSetSpender) jsonObj.set("spender", filter.Spender);

			if (filter.IsSetValue) jsonObj.set("value", filter.Value);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("owner", this.owner);
		jsonObj.set("spender", this.spender);
		jsonObj.set("value", convertToString(this.value));

		return jsonObj.toString();
	}
}

export class TransferEventSchema extends filter.TransferEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.TransferEventSchemaFilter[]): TransferEventSchema[] {
		const stringifiedRes = queryFromDb(
			"TransferEventSchema",
			TransferEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: TransferEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new TransferEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();

			const from = jsonObj.getString("from");
			if (from != null) res[i].From = from.valueOf();

			const to = jsonObj.getString("to");
			if (to != null) res[i].To = to.valueOf();

			const value = jsonObj.getNum("value");
			if (value != null) res[i].Value = BigInt.from(value.valueOf());
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "TransferEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "TransferEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "TransferEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			} else if (propName == "from") {
				if(str) { str = convertToString(this.from); }
			} else if (propName == "to") {
				if(str) { str = convertToString(this.to); }
			} else if (propName == "value") {
				if(str) { str = convertToString(this.value); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.TransferEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			if (filter.IsSetFrom) jsonObj.set("from", filter.From);

			if (filter.IsSetTo) jsonObj.set("to", filter.To);

			if (filter.IsSetValue) jsonObj.set("value", filter.Value);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);
		jsonObj.set("from", this.from);
		jsonObj.set("to", this.to);
		jsonObj.set("value", convertToString(this.value));

		return jsonObj.toString();
	}
}

export class PauseEventSchema extends filter.PauseEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.PauseEventSchemaFilter[]): PauseEventSchema[] {
		const stringifiedRes = queryFromDb(
			"PauseEventSchema",
			PauseEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: PauseEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new PauseEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "PauseEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "PauseEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "PauseEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.PauseEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);

		return jsonObj.toString();
	}
}

export class UnpauseEventSchema extends filter.UnpauseEventSchemaFilter {
	constructor() {
		super();
		this.id = generateUUID();
	}

	static query(filters: filter.UnpauseEventSchemaFilter[]): UnpauseEventSchema[] {
		const stringifiedRes = queryFromDb(
			"UnpauseEventSchema",
			UnpauseEventSchema.stringifyFilters(filters)
		)

		const dataList = (
			<JSON.Arr>JSON.parse(stringifiedRes)
		).valueOf();

		const res: UnpauseEventSchema[] = [];
		for (let i = 0; i < dataList.length; i++) {
			const jsonObj = <JSON.Obj>JSON.parse(dataList[i].toString());
			res[i] = new UnpauseEventSchema();

			const id = jsonObj.getString("id");
			if (id != null) res[i].Id = id.valueOf();
		}

		return res;
	}

	save(): void {
		saveToDb(this.toString(), "UnpauseEventSchema");
	}

	update(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		updateToDb(this.toString(), "UnpauseEventSchema");
	}

	delete(): void {
		if (!this.id) {
			throw new Error("Id is required");
		}

		deleteFromDb(this.id.toString(), "UnpauseEventSchema");
	}

	print(propName: string = ""): void {
		let str = "";
		if (propName) {
			if (propName == "id") {
				if(str) { str = convertToString(this.id); }
			}
		} else {
			str = this.toString();
		}
		console.log(str);
	}

	static stringifyFilters(filters: filter.UnpauseEventSchemaFilter[]): string {
		const jsonArr = new JSON.Arr();
		for (let i = 0; i < filters.length; i++) {
			const filter = filters[i];
			const jsonObj = new JSON.Obj();

			if (filter.IsSetId) jsonObj.set("id", filter.Id);

			jsonArr.push(jsonObj)
		}

		return jsonArr.stringify();
	}

	toString(): string {
		const jsonObj = new JSON.Obj();
		jsonObj.set("id", this.id);

		return jsonObj.toString();
	}
}

}
