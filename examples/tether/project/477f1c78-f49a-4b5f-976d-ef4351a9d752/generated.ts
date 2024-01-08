import { BigInt } from "./lib/bigint";
import { loadString, putString, console, convertToString } from "./utils";
import { EthereumEvent, GetEthereumEvent } from "./model/event";

export namespace generated {
  export class IssueEvent {
    private event: EthereumEvent;
    private amount: BigInt = BigInt.from(0);

    get Amount(): BigInt {
      return this.amount;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let amountParam = this.event.getParameter("amount");
      this.amount = amountParam!.asBigInt();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "amount") {
          if (this.amount) {
            str = convertToString(this.amount);
          }
        }
      } else {
        str = `{"amount":${this.amount}}`;
      }
      console.log(str);
    }
  }

  export class RedeemEvent {
    private event: EthereumEvent;
    private amount: BigInt = BigInt.from(0);

    get Amount(): BigInt {
      return this.amount;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let amountParam = this.event.getParameter("amount");
      this.amount = amountParam!.asBigInt();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "amount") {
          if (this.amount) {
            str = convertToString(this.amount);
          }
        }
      } else {
        str = `{"amount":${this.amount}}`;
      }
      console.log(str);
    }
  }

  export class DeprecateEvent {
    private event: EthereumEvent;
    private newAddress: string = "";

    get NewAddress(): string {
      return this.newAddress;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let newAddressParam = this.event.getParameter("newAddress");
      this.newAddress = newAddressParam!.asString();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "newAddress") {
          if (this.newAddress) {
            str = convertToString(this.newAddress);
          }
        }
      } else {
        str = `{"newAddress":"${this.newAddress}"}`;
      }
      console.log(str);
    }
  }

  export class ParamsEvent {
    private event: EthereumEvent;
    private feeBasisPoints: BigInt = BigInt.from(0);
    private maxFee: BigInt = BigInt.from(0);

    get FeeBasisPoints(): BigInt {
      return this.feeBasisPoints;
    }

    get MaxFee(): BigInt {
      return this.maxFee;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let feeBasisPointsParam = this.event.getParameter("feeBasisPoints");
      this.feeBasisPoints = feeBasisPointsParam!.asBigInt();
      let maxFeeParam = this.event.getParameter("maxFee");
      this.maxFee = maxFeeParam!.asBigInt();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "feeBasisPoints") {
          if (this.feeBasisPoints) {
            str = convertToString(this.feeBasisPoints);
          }
        } else if (propName == "maxFee") {
          if (this.maxFee) {
            str = convertToString(this.maxFee);
          }
        }
      } else {
        str = `{"feeBasisPoints":${this.feeBasisPoints},"maxFee":${this.maxFee}}`;
      }
      console.log(str);
    }
  }

  export class DestroyedBlackFundsEvent {
    private event: EthereumEvent;
    private _blackListedUser: string = "";
    private _balance: BigInt = BigInt.from(0);

    get BlackListedUser(): string {
      return this._blackListedUser;
    }

    get Balance(): BigInt {
      return this._balance;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let _blackListedUserParam = this.event.getParameter("_blackListedUser");
      this._blackListedUser = _blackListedUserParam!.asString();
      let _balanceParam = this.event.getParameter("_balance");
      this._balance = _balanceParam!.asBigInt();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "_blackListedUser") {
          if (this._blackListedUser) {
            str = convertToString(this._blackListedUser);
          }
        } else if (propName == "_balance") {
          if (this._balance) {
            str = convertToString(this._balance);
          }
        }
      } else {
        str = `{"_blackListedUser":"${this._blackListedUser}","_balance":${this._balance}}`;
      }
      console.log(str);
    }
  }

  export class AddedBlackListEvent {
    private event: EthereumEvent;
    private _user: string = "";

    get User(): string {
      return this._user;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let _userParam = this.event.getParameter("_user");
      this._user = _userParam!.asString();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "_user") {
          if (this._user) {
            str = convertToString(this._user);
          }
        }
      } else {
        str = `{"_user":"${this._user}"}`;
      }
      console.log(str);
    }
  }

  export class RemovedBlackListEvent {
    private event: EthereumEvent;
    private _user: string = "";

    get User(): string {
      return this._user;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let _userParam = this.event.getParameter("_user");
      this._user = _userParam!.asString();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "_user") {
          if (this._user) {
            str = convertToString(this._user);
          }
        }
      } else {
        str = `{"_user":"${this._user}"}`;
      }
      console.log(str);
    }
  }

  export class ApprovalEvent {
    private event: EthereumEvent;
    private owner: string = "";
    private spender: string = "";
    private value: BigInt = BigInt.from(0);

    get Owner(): string {
      return this.owner;
    }

    get Spender(): string {
      return this.spender;
    }

    get Value(): BigInt {
      return this.value;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let ownerParam = this.event.getParameter("owner");
      this.owner = ownerParam!.asString();
      let spenderParam = this.event.getParameter("spender");
      this.spender = spenderParam!.asString();
      let valueParam = this.event.getParameter("value");
      this.value = valueParam!.asBigInt();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "owner") {
          if (this.owner) {
            str = convertToString(this.owner);
          }
        } else if (propName == "spender") {
          if (this.spender) {
            str = convertToString(this.spender);
          }
        } else if (propName == "value") {
          if (this.value) {
            str = convertToString(this.value);
          }
        }
      } else {
        str = `{"owner":"${this.owner}","spender":"${this.spender}","value":${this.value}}`;
      }
      console.log(str);
    }
  }

  export class TransferEvent {
    private event: EthereumEvent;
    private from: string = "";
    private to: string = "";
    private value: BigInt = BigInt.from(0);

    get From(): string {
      return this.from;
    }

    get To(): string {
      return this.to;
    }

    get Value(): BigInt {
      return this.value;
    }

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {
      let fromParam = this.event.getParameter("from");
      this.from = fromParam!.asString();
      let toParam = this.event.getParameter("to");
      this.to = toParam!.asString();
      let valueParam = this.event.getParameter("value");
      this.value = valueParam!.asBigInt();
    }

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
        if (propName == "from") {
          if (this.from) {
            str = convertToString(this.from);
          }
        } else if (propName == "to") {
          if (this.to) {
            str = convertToString(this.to);
          }
        } else if (propName == "value") {
          if (this.value) {
            str = convertToString(this.value);
          }
        }
      } else {
        str = `{"from":"${this.from}","to":"${this.to}","value":${this.value}}`;
      }
      console.log(str);
    }
  }

  export class PauseEvent {
    private event: EthereumEvent;

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {}

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
      } else {
        str = `{}`;
      }
      console.log(str);
    }
  }

  export class UnpauseEvent {
    private event: EthereumEvent;

    constructor() {
      this.event = GetEthereumEvent();
      this.init();
    }

    init(): void {}

    print(propName: string = ""): void {
      let str = "";
      if (propName) {
      } else {
        str = `{}`;
      }
      console.log(str);
    }
  }
}
