import { Map as ImmutableMap } from 'immutable';

export enum ActionType {
  Delete,
  MarkAsDeleted,
  MarkAsFetched,
}

export class Action {
  type: ActionType;
  cartItemId: number;
  constructor(type: ActionType, cartItemId?: number) {
    this.type = type;
    this.cartItemId = cartItemId !== undefined ? cartItemId : 0;
  }
}

export class StateStore {
  needsFetch: boolean;
  cartItemIdsInDeletion: ImmutableMap<number, boolean>;
  // customerData: ImmutableMap<number, boolean>;

  constructor(needsFetch: boolean, cartItemIdsInDeletion: ImmutableMap<number, boolean>) {
    this.needsFetch = needsFetch;
    this.cartItemIdsInDeletion = cartItemIdsInDeletion;
  }

  withNeedsFetch(needsFetch: boolean): StateStore {
    this.needsFetch = needsFetch;
    return this;
  }

  withCartItemIdsInDeletion(cartItemIdsInDeletion: ImmutableMap<number, boolean>): StateStore {
    this.cartItemIdsInDeletion = cartItemIdsInDeletion;
    return this;
  }

  static fromPrevious(state: StateStore): StateStore {
    const newState = new StateStore(
      state.needsFetch,
      state.cartItemIdsInDeletion,
    );

    return newState
  }
}

export const initialState: StateStore = new StateStore(true, ImmutableMap<number, boolean>());
export function reducer(state: StateStore, action: Action): StateStore {
  switch (action.type) {
    case ActionType.Delete:
      return new StateStore(state.needsFetch, state.cartItemIdsInDeletion.set(action.cartItemId, false));
    case ActionType.MarkAsDeleted:
      const newCartItemIdsInDeletion = state.cartItemIdsInDeletion.set(action.cartItemId, true);
      const isAllDeletionFinished = newCartItemIdsInDeletion.reduce((isAllFinished, isFinished) => isAllFinished && isFinished, true);
      return new StateStore(isAllDeletionFinished, isAllDeletionFinished ? ImmutableMap<number, boolean>() : newCartItemIdsInDeletion);
    case ActionType.MarkAsFetched:
      return new StateStore(false, state.cartItemIdsInDeletion);
    default:
      throw new Error();
  }
}
