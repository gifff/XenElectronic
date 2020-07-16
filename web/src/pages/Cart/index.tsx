import React, { useState, useEffect, useReducer } from 'react';
import client from '../../lib/client';

import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
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
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import { Map as ImmutableMap } from 'immutable';
import { makeStyles } from '@material-ui/core';
import { useCookies } from 'react-cookie';
import { useSnackbar } from 'notistack';
import { useHistory } from 'react-router-dom';

import CartItem from '../../lib/model/CartItem';
import Product from '../../lib/model/Product';
import {
  Action,
  ActionType,
  initialState,
  reducer,
} from '../../lib/reducer/Cart';
import OrderResponse from '../../lib/responses/OrderResponse';

const useStyles = makeStyles((theme) => ({
  table: {
    midWidth: 700,
  },
  checkoutForm: {
    flex: 1,
    padding: theme.spacing(1),
  },
  inputFieldWrapper: {
    paddingLeft: theme.spacing(2),
    paddingRight: theme.spacing(2),
    paddingTop: theme.spacing(1),
    paddingBottom: theme.spacing(1),
  }
}))

const currencyFormatter = new Intl.NumberFormat('id-ID', {
  style: 'currency',
  currency: 'IDR',
})

export default function Cart() {
  const classes = useStyles();
  const history = useHistory();
  const { enqueueSnackbar } = useSnackbar();
  const [cartItems, setCartItems] = useState<CartItem[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isCheckout, setIsCheckout] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);
  const [customerData, setCustomerData] = useState<ImmutableMap<string, string>>(ImmutableMap<string, string>());
  const [state, dispatch] = useReducer(reducer, initialState);
  const [{ cartId },, removeCookie] = useCookies(['cartId']);

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

  const handleCustomerDataChange = (key: string) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setCustomerData(customerData.set(key, event.target.value));
  };

  const handleCheckout = () => {
    setIsCheckout(true);
    const customerName = customerData.get<string>('customer_name', '');
    const customerEmail = customerData.get<string>('customer_email', '');
    const customerAddress = customerData.get<string>('customer_address', '');
    client.createOrder(cartId, customerName, customerEmail, customerAddress)
      .then(response => {
        if (!response.ok) {
          throw new Error('something is broken');
        }

        setCustomerData(ImmutableMap());
        removeCookie('cartId');
        return response.json();
      })
      .then((response: OrderResponse) => {
        history.replace(`/orders/${response.id}`);
      })
      .catch(error => {
        enqueueSnackbar('Unable to create order', {
          autoHideDuration: 3000,
          variant: 'error',
          preventDuplicate: true,
        });
      })
      .finally(() => {
        setIsCheckout(false);
      });
  };

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
      <Grid item xs={12} md={6}>
        <Paper component="form" onSubmit={event => { event.preventDefault(); handleCheckout(); }} className={classes.checkoutForm}>
          <Grid item xs={12} className={classes.inputFieldWrapper}>
            <TextField
              id="customer-name"
              label="Customer Name"
              variant="filled"
              size="small"
              value={customerData.get<string>('customer_name', '')}
              disabled={isCheckout}
              fullWidth
              required
              onChange={handleCustomerDataChange('customer_name')}
            />
          </Grid>
          <Grid item xs={12} className={classes.inputFieldWrapper}>
            <TextField
              id="customer-email"
              label="Customer Email"
              type="email"
              variant="filled"
              size="small"
              value={customerData.get<string>('customer_email', '')}
              disabled={isCheckout}
              fullWidth
              required
              onChange={handleCustomerDataChange('customer_email')}
            />
          </Grid>
          <Grid item xs={12} className={classes.inputFieldWrapper}>
            <TextField
              id="customer-address"
              label="Customer Address"
              variant="filled"
              size="small"
              value={customerData.get<string>('customer_address', '')}
              disabled={isCheckout}
              fullWidth
              required
              onChange={handleCustomerDataChange('customer_address')}
            />
          </Grid>

          <Grid item container xs={12} justify="flex-end" className={classes.inputFieldWrapper}>
            <Button
              variant="contained"
              color="primary"
              size="large"
              startIcon={<LocalMallIcon />}
              disabled={isCheckout}
              type="submit"
            >
              Checkout
            </Button>
          </Grid>
        </Paper>
      </Grid>
    </React.Fragment >
  );

  const errorContent = (
    <Grid item container justify="center" xs={12}>
      <Typography variant="body1" color="error">Sorry, something is broken</Typography>
    </Grid>
  );

  return (
    <Grid container spacing={3} justify="flex-end">
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
