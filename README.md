# Acervus CLI

Welcome to Acervus CLI, your dedicated tool for interacting with the Acervus Cloud with ease.

## What does Acervus do?

Indexing and tracking blockchains can be complex. Acervus simplifies this.

Rather than relying on general-purpose indexers or APIs, you can write a few lines of code, upload it to Acervus, and access blockchain data tailored to your specifications. We've already indexed all the data from certain blockchains, enabling you to filter, transform, and query them to meet your specific requirements.

## Installation

### macOS

#### Using Homebrew

You should tap the `paribu/acervus` repository to install `acervus`.

```
brew install paribu/acervus/acervus
```

Or `brew tap paribu/acervus` and then `brew install acervus`.

This tool is constantly updated, so don't forget to run following command to get latest updates:

```
brew upgrade acervus
```

You can visit the tap at https://github.com/paribu/homebrew-acervus

### Linux

You can use Homebrew to install ( see macOS instructions above ). Other installation methods will soon be available.

### Windows

Currently, only building from source is supported. Other installation methods will soon be available.

## How it works? An example:

Assuming you're keen on tracking an NFT collection, let's say you're curious to find out which NFT is the most popular based on the number of times it's been sold or changed owners. Take the [MutantApe YachtClub](https://etherscan.io/token/0x60e4d786628fea6478f785a6d7e704777c86a7c6) NFT collection as an example.

First, we need to obtain the contract's ABI to identify the events emitted by that contract. This can be done by visiting [this link](https://etherscan.io/token/0x60e4d786628fea6478f785a6d7e704777c86a7c6#code). Since we're focused on tracking ownership changes, our main interest lies in the `Transfer` event, which is inherited from the ERC721 standard.

We don't need the entire ABI for our purpose, just the part relevant to our tracking. Therefore, we can create a trimmed version of the ABI, named `ABI.json`, which will include only the necessary information. It would look something like this:


```json
[
     {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "tokenId",
                "type": "uint256"
            }
        ],
        "name": "Transfer",
        "type": "event"
    }
]
```

Alongside the ABI, we require a structured way to store our data. This calls for defining our GraphQL schema. To do this, we'll craft a `graphql.schema` file. It's structured capture and organize the data we're interested in. Here's an example of what our `graphql.schema` file might look like:


```graphql
type TransferEventSchema @entity {
	tokenId: BigInt
    transferCount: Int
}
```

In our `entity` within the GraphQL schema, we're going to track specific details: the `tokenId` of each NFT and the number of times it has been transferred (`transferCount`).

Once we've set up the schema, the next step involves generating some helper code. This can be efficiently done using our CLI tool. With these helper files in place (refer to the imported files for details), we're all set to dive into the main part of our project - writing the indexer code. This will be done in the `project.ts` file.


```typescript
import { generated } from "./generated";
import { schema } from "./schema";
import { filter } from "./schema/filter";
import { console } from "./utils";

// This function is invoked for each transfer event during the indexing process.
// It processes events sequentially from startBlock to endBlock as defined in settings.
export function handleTransferEvent(event: generated.TransferEvent): void {
    // "event" object corresponds to ABI item
    // "schema" and "filter" are generated from our graphql entity

    // Extracting tokenID from the event.
    const tokenID = event.TokenId;

    // Retrieving sender and receiver information.
    const sender = event.From;
    const receiver = event.to;

    // Optional: Implement logic to filter out irrelevant events.
    // Example: Skipping the event if the sender and receiver are the same.
    if (sender == receiver) {
        return; // Skip indexing
    }

    // Now we need to get our previous records for this tokenId from DB if it exists.
    // Filter name is derived from our graphql entity.
    const filter = new filter.TransferEventSchemaFilter();
    filter.tokenId = tokenID; // Specifying the query parameter.

    // Executing the query.
    const queryResult = schema.TransferEventSchema.query([filter]);

    // Handling a new NFT (tokenID not previously indexed).
    if (queryResult.length == 0) {
        const transferEvent = new schema.TransferEventSchema();
        console.log("Found new NFT with ID" + tokenId); // Logging for new NFT discovery.
        transferEvent.TokenId = tokenId; // Setting the token ID.
        // Set transfer count to 1 since it is the first time we encountered this tokenId.
        transferEvent.TransferCount = 1; 
        // Now we can save our object
        transferEvent.save(); // Stored in the database now.
    } else { 
        // If the tokenId (our NFT) is indexed before, let's update it
        const result = queryResult[0]; // Assuming a single entity per tokenID.
        result.transferEvent = result.transferEvent + 1; // Increment the transfer count.
        transferEvent.update(); // Updating the existing record in the database.
    }
}
```

Once you upload your code to Acervus, the platform takes over by sequentially sending events from the specified `startBlock` to the `endBlock`, and your code will process each of these events.

But, you might wonder, how does Acervus determine which function to invoke for each specific event? This is where the `setting.yml` file comes into play. You need to provide this file to guide Acervus in matching functions to events. The structure of the `setting.yml` file is straightforward and will look something like this:


```yaml
project: MyMutantApeIndexer
description: It will count how many times an NFT transferred
schema: ./schema.graphql
sources:
- track: ethereum/contract
  name: MutantApeYachtClub
  network: Ethereum Mainnet
  source:
    address: 0x60E4d786628Fea6478F785A6d7e704777c86a7c6
    abi: ./abi.json
  code:
    handlers:
    - type: ethereum/event
      function: handleTransferEvent # Your function in project.ts
      name: Transfer(address,address,uint256) # Event from ABI
      startBlock: 10000000 # Index this event starting from 1 millionth block.
      endBlock: 0 # This zero means there is no "end block" so it will continue indexing forever (as new blocks are generated)
```

After a period of waiting (keep in mind, indexing does take some time, but it's usually not too lengthy), your database will be populated with the data structured according to the schema entity you previously defined. Once the data is in place, you can query it - along with the logs - to view the results.

To do this, simply run the following command in your terminal:

```bash
$ acervus query data
```

When you execute the command, you can expect to receive a result similar to the following:

```json
{
    "results":[
        {
            "$id": "some-object-id",
            "tokenId": 1523,
            "transferCount": 51
        },
        {
            "$id": "some-other-object-id",
            "tokenId": 1024,
            "transferCount": 13
        }
        ...
    ]
}
```

Note: `$id` field is autogenerated by Acervus, and it's used to identify the object in the database.

## Conclusion

Acervus introduces a more adaptable approach to indexing blockchain data, setting it apart from traditional APIs. The beauty of Acervus lies in its simplicity - there's no need for you to manage servers or set up databases. All you need to do is upload your code, and Acervus takes care of delivering the data you need, tailored to your requirements.

At present, Acervus supports only `Events` and is limited to the `Ethereum Mainnet`. However, we have plans in the pipeline to extend our capabilities to include other blockchains and their respective data.

For more detailed information on how to use the CLI and to get a deeper understanding of its functionalities, please refer to the `docs` folder. Additionally, for a more comprehensive example of what you can achieve with Acervus, take a look at the `examples` folder.
