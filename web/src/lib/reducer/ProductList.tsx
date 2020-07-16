import {
  Map as ImmutableMap,
  Stack as ImmutableStack,
} from 'immutable';

export enum ActionType {
  AddToCart,
  MarkAsAdded,
}

export class Action {
  type: ActionType;
  productId: number;
  constructor(type: ActionType, productId: number = 0) {
    this.type = type;
    this.productId = productId;
  }
}

export class StateStore {
  inFlightAddToCartRequests: ImmutableMap<number, ImmutableStack<boolean>>;
  constructor(inFlightAddToCartRequests: ImmutableMap<number, ImmutableStack<boolean>>) {
    this.inFlightAddToCartRequests = inFlightAddToCartRequests;
  }
}

export const initialState: StateStore = new StateStore(ImmutableMap<number, ImmutableStack<boolean>>());
export function reducer(state: StateStore, action: Action): StateStore {
  switch (action.type) {
    case ActionType.AddToCart:
      {
        const currentStack = state.inFlightAddToCartRequests.get(action.productId);
        const newStack = currentStack !== undefined ? currentStack.push(true) : ImmutableStack<boolean>().push(true);
        return new StateStore(state.inFlightAddToCartRequests.set(action.productId, newStack));
      }
    case ActionType.MarkAsAdded:
      {
        const currentStack = state.inFlightAddToCartRequests.get(action.productId);
        if (currentStack !== undefined && currentStack.size > 1) {
          const newStack = currentStack.pop();
          return new StateStore(state.inFlightAddToCartRequests.set(action.productId, newStack));
        } else {
          return new StateStore(state.inFlightAddToCartRequests.remove(action.productId));
        }
      }
    default:
      throw new Error();
  }
}
