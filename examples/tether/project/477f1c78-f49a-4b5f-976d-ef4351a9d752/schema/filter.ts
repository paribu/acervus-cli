import { BigInt } from "../lib/bigint";
export namespace filter {
  export class IssueEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected amount: BigInt = BigInt.from(0);
    protected isSetAmount: bool = false;
    set Amount(value: BigInt) {
      this.isSetAmount = true;
      this.amount = value;
    }
    get Amount(): BigInt {
      return this.amount;
    }

    get IsSetAmount(): bool {
      return this.isSetAmount;
    }
  }

  export class RedeemEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected amount: BigInt = BigInt.from(0);
    protected isSetAmount: bool = false;
    set Amount(value: BigInt) {
      this.isSetAmount = true;
      this.amount = value;
    }
    get Amount(): BigInt {
      return this.amount;
    }

    get IsSetAmount(): bool {
      return this.isSetAmount;
    }
  }

  export class DeprecateEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected newAddress: string = "";
    protected isSetNewAddress: bool = false;
    set NewAddress(value: string) {
      this.isSetNewAddress = true;
      this.newAddress = value;
    }
    get NewAddress(): string {
      return this.newAddress;
    }

    get IsSetNewAddress(): bool {
      return this.isSetNewAddress;
    }
  }

  export class ParamsEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected feeBasisPoints: BigInt = BigInt.from(0);
    protected isSetFeeBasisPoints: bool = false;
    set FeeBasisPoints(value: BigInt) {
      this.isSetFeeBasisPoints = true;
      this.feeBasisPoints = value;
    }
    get FeeBasisPoints(): BigInt {
      return this.feeBasisPoints;
    }

    get IsSetFeeBasisPoints(): bool {
      return this.isSetFeeBasisPoints;
    }

    protected maxFee: BigInt = BigInt.from(0);
    protected isSetMaxFee: bool = false;
    set MaxFee(value: BigInt) {
      this.isSetMaxFee = true;
      this.maxFee = value;
    }
    get MaxFee(): BigInt {
      return this.maxFee;
    }

    get IsSetMaxFee(): bool {
      return this.isSetMaxFee;
    }
  }

  export class DestroyedBlackFundsEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected _blackListedUser: string = "";
    protected isSetBlackListedUser: bool = false;
    set BlackListedUser(value: string) {
      this.isSetBlackListedUser = true;
      this._blackListedUser = value;
    }
    get BlackListedUser(): string {
      return this._blackListedUser;
    }

    get IsSetBlackListedUser(): bool {
      return this.isSetBlackListedUser;
    }

    protected _balance: BigInt = BigInt.from(0);
    protected isSetBalance: bool = false;
    set Balance(value: BigInt) {
      this.isSetBalance = true;
      this._balance = value;
    }
    get Balance(): BigInt {
      return this._balance;
    }

    get IsSetBalance(): bool {
      return this.isSetBalance;
    }
  }

  export class AddedBlackListEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected _user: string = "";
    protected isSetUser: bool = false;
    set User(value: string) {
      this.isSetUser = true;
      this._user = value;
    }
    get User(): string {
      return this._user;
    }

    get IsSetUser(): bool {
      return this.isSetUser;
    }
  }

  export class RemovedBlackListEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected _user: string = "";
    protected isSetUser: bool = false;
    set User(value: string) {
      this.isSetUser = true;
      this._user = value;
    }
    get User(): string {
      return this._user;
    }

    get IsSetUser(): bool {
      return this.isSetUser;
    }
  }

  export class ApprovalEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected owner: string = "";
    protected isSetOwner: bool = false;
    set Owner(value: string) {
      this.isSetOwner = true;
      this.owner = value;
    }
    get Owner(): string {
      return this.owner;
    }

    get IsSetOwner(): bool {
      return this.isSetOwner;
    }

    protected spender: string = "";
    protected isSetSpender: bool = false;
    set Spender(value: string) {
      this.isSetSpender = true;
      this.spender = value;
    }
    get Spender(): string {
      return this.spender;
    }

    get IsSetSpender(): bool {
      return this.isSetSpender;
    }

    protected value: BigInt = BigInt.from(0);
    protected isSetValue: bool = false;
    set Value(value: BigInt) {
      this.isSetValue = true;
      this.value = value;
    }
    get Value(): BigInt {
      return this.value;
    }

    get IsSetValue(): bool {
      return this.isSetValue;
    }
  }

  export class TransferEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }

    protected from: string = "";
    protected isSetFrom: bool = false;
    set From(value: string) {
      this.isSetFrom = true;
      this.from = value;
    }
    get From(): string {
      return this.from;
    }

    get IsSetFrom(): bool {
      return this.isSetFrom;
    }

    protected to: string = "";
    protected isSetTo: bool = false;
    set To(value: string) {
      this.isSetTo = true;
      this.to = value;
    }
    get To(): string {
      return this.to;
    }

    get IsSetTo(): bool {
      return this.isSetTo;
    }

    protected value: BigInt = BigInt.from(0);
    protected isSetValue: bool = false;
    set Value(value: BigInt) {
      this.isSetValue = true;
      this.value = value;
    }
    get Value(): BigInt {
      return this.value;
    }

    get IsSetValue(): bool {
      return this.isSetValue;
    }
  }

  export class PauseEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }
  }

  export class UnpauseEventSchemaFilter {
    protected $id: string = "";
    protected isSet$Id: bool = false;
    set $Id(value: string) {
      this.isSet$Id = true;
      this.$id = value;
    }
    get $Id(): string {
      return this.$id;
    }

    get IsSet$Id(): bool {
      return this.isSet$Id;
    }

    protected id: string = "";
    protected isSetId: bool = false;
    set Id(value: string) {
      this.isSetId = true;
      this.id = value;
    }
    get Id(): string {
      return this.id;
    }

    get IsSetId(): bool {
      return this.isSetId;
    }
  }
}
