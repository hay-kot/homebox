export enum EventTypes {
  // ClearStores event is used to inform the stores that _all_ the data they are using
  // is now out of date and they should refresh - This is used when the user makes large
  // changes to the data such as bulk actions or importing a CSV file
  ClearStores,
}

export type EventFn = () => void;

export interface IEventBus {
  on(event: EventTypes, fn: EventFn, key: string): void;
  off(event: EventTypes, key: string): void;
  emit(event: EventTypes): void;
}

class EventBus implements IEventBus {
  private listeners: Record<EventTypes, Record<string, EventFn>> = {
    [EventTypes.ClearStores]: {},
  };

  on(event: EventTypes, fn: EventFn, key: string): void {
    this.listeners[event][key] = fn;
  }

  off(event: EventTypes, key: string): void {
    delete this.listeners[event][key];
  }

  emit(event: EventTypes): void {
    Object.values(this.listeners[event]).forEach(fn => fn());
  }
}

const bus = new EventBus();

export function useEventBus(): IEventBus {
  return bus;
}
