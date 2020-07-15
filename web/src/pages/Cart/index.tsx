import React, { useState, useEffect, useReducer } from 'react';
import client from '../../lib/client';

import Button from '@material-ui/core/Button';
import DeleteIcon from '@material-ui/icons/Delete';
import Grid from '@material-ui/core/Grid';
import IconButton from '@material-ui/core/IconButton';
import LocalMallIcon from '@material-ui/icons/LocalMall';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';
import { makeStyles, CircularProgress } from '@material-ui/core';
import { useCookies } from 'react-cookie';
import { useSnackbar } from 'notistack';
import { Map as ImmutableMap } from 'immutable';

import CartItem from '../../lib/model/CartItem';
import Product from '../../lib/model/Product';

const useStyles = makeStyles((theme) => ({
  table: {
    midWidth: 700,
  },
  checkoutWrapper: {
    marginBottom: 24,
  },
}))

const currencyFormatter = new Intl.NumberFormat('id-ID', {
  style: 'currency',
  currency: 'IDR',
})

enum ActionType {
  Delete,
  MarkAsDeleted,
  MarkAsFetched,
}

class Action {
  type: ActionType;
  cartItemId: number;
  constructor(type: ActionType, cartItemId?: number) {
    this.type = type;
    this.cartItemId = cartItemId !== undefined ? cartItemId : 0;
  }
}

class StateStore {
  needsFetch: boolean;
  cartItemIdsInDeletion: ImmutableMap<number, boolean>;
  constructor(needsFetch: boolean, cartItemIdsInDeletion: ImmutableMap<number, boolean>) {
    this.needsFetch = needsFetch;
    this.cartItemIdsInDeletion = cartItemIdsInDeletion;
  }
}

const initialState: StateStore = new StateStore(true, ImmutableMap<number, boolean>());
function reducer(state: StateStore, action: Action): StateStore {
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

export default function Cart() {
  const classes = useStyles();
  const { enqueueSnackbar } = useSnackbar();
  const [cartItems, setCartItems] = useState<CartItem[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);
  const [state, dispatch] = useReducer(reducer, initialState);
  const [{ cartId }] = useCookies(['cartId']);

  useEffect(() => {
    if (!isLoading && state.needsFetch) {
      setIsLoading(true);
      setError(null);
      client.listProductsInCart(cartId)
        .then(response => response.json())
        .then(response => {
          if (Array.isArray(response)) {
            setCartItems(response.map(cartItem => new CartItem(
              cartItem.id,
              cartItem.productId,
              new Product(
                cartItem.product.id,
                cartItem.product.categoryId,
                cartItem.product.name,
                cartItem.product.description,
                cartItem.product.photo,
                cartItem.product.price,
              )
            ))
              .sort((a, b) => {
                if (a.product.name < b.product.name) {
                  return -1;
                } else if (a.product.name > b.product.name) {
                  return 1;
                } else {
                  return a.product.id - b.product.id;
                }
              })
            );
          }
        })
        .catch(error => {
          setCartItems([]);
          setError(error)
        })
        .finally(() => {
          setIsLoading(false);
          dispatch(new Action(ActionType.MarkAsFetched));
        })
    }
  }, [isLoading, state.needsFetch, cartId]);

  const removeProductFromCart = (productId: number, cartItemId: number) => {
    dispatch(new Action(ActionType.Delete, cartItemId));
    client.removeProductFromCart(cartId, productId)
      .then(response => {
      })
      .catch(() => {
        enqueueSnackbar('Sorry, something is broken', {
          autoHideDuration: 3000,
          variant: 'error',
          preventDuplicate: true,
        })
      })
      .finally(() => {
        dispatch(new Action(ActionType.MarkAsDeleted, cartItemId));
      })
  }

  const loadingContent = (
    <Grid item container justify="center" xs={12}>
      <CircularProgress />
    </Grid>
  );

  const totalPrice: number = cartItems.reduce((sum, cartItem) => sum + cartItem.product.price, 0)

  const content = (
    <React.Fragment>
      <Grid item xs={12}>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="spanning table">
            <TableHead>
              <TableRow>
                <TableCell size="small">No.</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Price</TableCell>
                <TableCell />
              </TableRow>
            </TableHead>
            <TableBody>
              {cartItems.map((cartItem, index) => (
                <TableRow key={`${cartItem.product.name}_${cartItem.id}`}>
                  <TableCell size="small">{index + 1}</TableCell>
                  <TableCell align="left">{cartItem.product.name}</TableCell>
                  <TableCell align="left">{currencyFormatter.format(cartItem.product.price)}</TableCell>
                  <TableCell align="right" size="small">
                    {
                      state.cartItemIdsInDeletion.has(cartItem.id) ? <CircularProgress color="secondary" /> : (
                        <IconButton onClick={() => removeProductFromCart(cartItem.product.id, cartItem.id)}>
                          <DeleteIcon />
                        </IconButton>
                      )
                    }
                  </TableCell>
                </TableRow>
              ))}
              <TableRow>
                <TableCell />
                <TableCell colSpan={1} align="right">Total</TableCell>
                <TableCell align="left">{currencyFormatter.format(totalPrice)}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </Grid>
      <Grid item container xs={12} justify="flex-end" className={classes.checkoutWrapper}>
        <Button
          variant="contained"
          color="primary"
          size="large"
          startIcon={<LocalMallIcon />}
        >
          Checkout
        </Button>
      </Grid>
    </React.Fragment >
  );

  const errorContent = (
    <Grid item container justify="center" xs={12}>
      <Typography variant="body1" color="error">Sorry, something is broken</Typography>
    </Grid>
  );

  return (
    <Grid container spacing={3}>
      {
        isLoading ? loadingContent : (
          error !== null ? errorContent : (
            content
          )
        )
      }
    </Grid>
  );
}
