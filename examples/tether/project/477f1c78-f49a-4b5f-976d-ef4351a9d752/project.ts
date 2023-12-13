import { generated } from './generated';
import { schema } from './schema';
import { filter } from './schema/filter';

export function handleDeprecateEvent(event: generated.DeprecateEvent): void {
  // Example of creating a new schema object and saving it
  let deprecateEvent = new schema.DeprecateEventSchema();
  deprecateEvent.NewAddress = event.NewAddress;
  deprecateEvent.save();

  // Print the event details
  event.print();
}

export function handleDestroyedBlackFundsEvent(
  event: generated.DestroyedBlackFundsEvent,
): void {
  // Example of querying and updating an existing schema object
  let filters = [new filter.DestroyedBlackFundsEventSchemaFilter()];
  let events = schema.DestroyedBlackFundsEventSchema.query(filters);
  if (events.length > 0) {
    let destroyedEvent = events[0];
    destroyedEvent.BlackListedUser = event.BlackListedUser;
    destroyedEvent.Balance = event.Balance;
    destroyedEvent.update();
  }

  // Print the event details
  event.print();
}

export function handleAddedBlackListEvent(
  event: generated.AddedBlackListEvent,
): void {
  // Example of deleting an existing schema object
  let filters = [new filter.AddedBlackListEventSchemaFilter()];
  let events = schema.AddedBlackListEventSchema.query(filters);
  if (events.length > 0) {
    events[0].delete();
  }

  // Print the event details
  event.print();
}

export function handleApprovalEvent(event: generated.ApprovalEvent): void {
  // Example of printing specific properties of the event
  event.print('owner');
  event.print('spender');
  event.print('value');
}

export function handleTransferEvent(event: generated.TransferEvent): void {
  // Example of simple logging
  console.log(
    `Transfer from ${event.From} to ${event.To} of value ${event.Value}`,
  );
}
