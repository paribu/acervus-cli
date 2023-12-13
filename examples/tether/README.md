# Acervus Project Creation Guide

This guide details the step-by-step process of creating a software development project.

## 1. User Registration and Login

The first step involves user registration and login. If you are registered, you can log in. If not, you must register first.

**Registration Command**:

	    acervus auth register -e <email> -p <password>

This command registers a user with the specified email address and password.

## 2. Adding ABI File and Generating Settings File

In this step, we add the ABI file to the system and create a settings file.

**Required Information for Settings**:

- **Project Name**: This is the name of your project. It's usually chosen to reflect the purpose or theme of the project.
- **Project Description**: This is a brief summary of what your project is about. It might include information about its goals, functionalities, and any unique features.
- **Network**: In the context of blockchain, this refers to the specific blockchain network your project is using.
- **Contract Address**: This is the address of the smart contract on the blockchain that your project is interacting with. Smart contracts are self-executing contracts with the terms of the agreement directly written into code.
- **Contract Name**: This is the name given to the smart contract. It helps in identifying the specific contract within the blockchain.
- **ABI (Application Binary Interface) File Path**: The ABI is essentially how you call functions in a contract and get data back. It's a JSON-formatted file that describes the interfaces of your smart contract, including methods and structures. The ABI file path is the location where this file is stored.
- **Schema File Path**: This could refer to the file path for a database schema related to your project, which defines its database's structure in terms of database tables, the fields in each table, and the relationships between fields and tables.
- **Selected Events**: In blockchain, events are outputs of smart contracts that provide a way of recording the execution of certain processes. Selecting events would involve specifying which events your application is interested in or should respond to.
- **CRUD Events**: CRUD stands for Create, Read, Update, Delete. It's a shorthand for the basic operations that might be performed on the data in your project. If this setting is in CRUD mode, we can create a query, otherwise we can save.
- **Start Block**: In blockchain, each transaction is recorded in a block. The start block would be the block number from where your project begins to consider transactions or events.
- **End Block**: This is the block number where your project stops considering transactions or events. It's useful for limiting the scope of data your project is dealing with.

**Settings File Creation Command**:

      acervus generate settings

After running the create settings command, settings and schema files are created. The settings file contains detailed information about how the project will be structured and which events of the smart contract will be monitored.

## 3. Creating the Project

Now that our ABI file and settings file are ready, we can create our project.

 **Project Creation Command**:

	    acervus projects create

This command will create the project in the specified directory or in the default (project) directory.

>**Project Directory**: The command creates the project in a specified directory or the default (project) directory. This is where all files for the project will be stored. The name of the created directory is your project ID. 
>For example: *477f1c78-f49a-4b5f-976d-ef4351a9d752*

>**Generated  File**: This file usually contains TypeScript code derived from the ABI file. The ABI (Application Binary Interface) defines the functions and structures of a smart contract. The generated.ts file includes functions and classes for interacting with the smart contract on the Ethereum network using these functions and structures.

##### Transfer Event in ABI:
```json
{
    "anonymous":  false,
    "inputs":  [
	    {  "indexed":  true,  "name":  "from",  "type":  "address"  },
	    {  "indexed":  true,  "name":  "to",  "type":  "address"  },
	    {  "indexed":  false,  "name":  "value",  "type":  "uint256"  }
    ],
    "name":  "Transfer",
    "type":  "event"
}
```

##### Transfer Event in Generated File:
```typescript
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
                if(str) { str = convertToString(this.from); }
            } else if (propName == "to") {
                if(str) { str = convertToString(this.to); }
            } else if (propName == "value") {
                if(str) { str = convertToString(this.value); }
            }
        } else {
            str = `{"from":"${this.from}","to":"${this.to}","value":${this.value}}`;
        }
        console.log(str);
    }
}
```


>**Project  File**: This file contains the core configuration of the project and typically includes information like project settings, endpoints, API paths, etc. The project.ts provides essential details on how the project will be executed and how different modules are integrated.

>**Other Files and Folders**: During the creation process, other auxiliary files and folders necessary for the smooth functioning of the project are also created. These may include configuration files, test scripts, and possibly folders containing the frontend and backend code of the project.

## 4. Reviewing and Editing the Project File

We can review the created project file and edit the handle functions in the `project.ts` file as necessary.

 - handleDeprecateEvent: Creates a new schema object and saves it.
 ```typescript
export  function handleDeprecateEvent(event:  generated.DeprecateEvent):  void  {
	// Example of creating a new schema object and saving it
	let  deprecateEvent  =  new  schema.DeprecateEventSchema();
	deprecateEvent.NewAddress  =  event.NewAddress;
	deprecateEvent.save();

	// Print the event details
	event.print();
}
```
   
  - handleDestroyedBlackFundsEvent: Queries and updates an existing
   schema object.

 ```typescript
export  function handleDestroyedBlackFundsEvent( event:  generated.DestroyedBlackFundsEvent,
):  void  {
	// Example of querying and updating an existing schema object
	let  filters  =  [new  filter.DestroyedBlackFundsEventSchemaFilter()];
	let  events  =  schema.DestroyedBlackFundsEventSchema.query(filters);
	if  (events.length  >  0)  {
	let  destroyedEvent  =  events[0];
	destroyedEvent.BlackListedUser  =  event.BlackListedUser;
	destroyedEvent.Balance  =  event.Balance;
	destroyedEvent.update();
	}

	// Print the event details
	event.print();
}
```
   
 - handleAddedBlackListEvent: Deletes an existing schema object.

 ```typescript
export  function handleAddedBlackListEvent( event:  generated.AddedBlackListEvent):  void  {
	// Example of deleting an existing schema object
	let  filters  =  [new  filter.AddedBlackListEventSchemaFilter()];
	let  events  =  schema.AddedBlackListEventSchema.query(filters);
	if  (events.length  >  0)  {
	events[0].delete();
	}

	// Print the event details
	event.print();
}
```
   
 - handleApprovalEvent: Prints specific properties of the event.

 ```typescript
export  function handleApprovalEvent(event:  generated.ApprovalEvent):  void  {
	// Example of printing specific properties of the event
	event.print('owner');
	event.print('spender');
	event.print('value');
}
```
   
 - handleTransferEvent: Logs a message with details from the event.

 ```typescript
export  function handleTransferEvent(event:  generated.TransferEvent):  void  {
	// Example of simple logging
	console.log(
	`Transfer from ${event.From} to ${event.To} of value ${event.Value}`,
	);
}
```

## 5. Testing the Project

Before deploying, we need to check that the system is working correctly with the command.

	    acervus test -i <projectID>

Here, `projectID` is the identifier of the project to be tested.

## 6. Deploying the Project

If the test is successful, we are ready to deploy the project.

 **Deploy Command**:

	    acervus deploy -i <projectID>

Here, `projectID` is the identifier of your project.

--- 

This guide explains how to create and deploy a software development project step by step. By carefully following each step, you can successfully complete your project.

