import React, { useState, useEffect } from 'react';
import client from '../../lib/client';

import CircularProgress from '@material-ui/core/CircularProgress';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core';
import { useParams } from 'react-router-dom';

import CartItem from '../../lib/model/CartItem';
import Product from '../../lib/model/Product';
import Order from '../../lib/model/Order';
import OrderResponse from '../../lib/responses/OrderResponse';

const useStyles = makeStyles((theme) => ({
  table: {
    midWidth: 700,
  },
}))

const currencyFormatter = new Intl.NumberFormat('id-ID', {
  style: 'currency',
  currency: 'IDR',
})

export default function ViewOrder() {
  const classes = useStyles();
  const [order, setOrder] = useState<Order>(new Order());
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isFetched, setIsFetched] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);
  const { orderId } = useParams();

  useEffect(() => {
    if (!isLoading && !isFetched) {
      setIsLoading(true);
      setError(null);
      client.viewOrder(orderId)
        .then(response => {
          if (!response.ok) {
            throw new Error('Error occured');
          }

          return response.json();
        })
        .then((response: OrderResponse) => {
          let cartItems: CartItem[] = new Array<CartItem>();
          if (Array.isArray(response.cart_items)) {
            cartItems = response.cart_items.map(cartItem => new CartItem(
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
            ));

            cartItems.sort((a, b) => {
              if (a.product.name < b.product.name) {
                return -1;
              } else if (a.product.name > b.product.name) {
                return 1;
              } else {
                return a.product.id - b.product.id;
              }
            });
          }

          setOrder(new Order(
            response.id,
            response.customer_name,
            response.customer_email,
            response.customer_address,
            response.payment_method,
            response.payment_account_number,
            response.payment_amount,
            cartItems,
          ))
        })
        .catch(error => {
          setError(error)
        })
        .finally(() => {
          setIsLoading(false);
          setIsFetched(true);
        })
    }
  }, [isLoading, isFetched, orderId]);

  const loadingContent = (
    <Grid item container justify="center" xs={12}>
      <CircularProgress />
    </Grid>
  );

  const totalPrice: number = order.cartItems.reduce((sum, cartItem) => sum + cartItem.product.price, 0)

  const content = (
    <React.Fragment>
      <Grid item xs={12}>
        <TableContainer component={Paper}>
          <Table className={classes.table} size="small">
            <TableBody>
              <TableRow>
                <TableCell padding="none" align="right">Order ID :</TableCell>
                <TableCell align="left">{order.id}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell padding="none" align="right">Customer Name :</TableCell>
                <TableCell align="left">{order.customerName}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell padding="none" align="right">Customer Email :</TableCell>
                <TableCell align="left">{order.customerEmail}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell padding="none" align="right">Customer Address :</TableCell>
                <TableCell align="left">{order.customerAddress}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell padding="none" align="right">Payment Method :</TableCell>
                <TableCell align="left">{order.paymentMethod === '' ? '-' : order.paymentMethod}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell padding="none" align="right">Payment Account Number :</TableCell>
                <TableCell align="left">{order.paymentAccountNumber === '' ? '-' : order.paymentAccountNumber}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell padding="none" align="right">Payment Amount :</TableCell>
                <TableCell align="left">{currencyFormatter.format(order.paymentAmount)}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </Grid>
      <Grid item xs={12}>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="spanning table">
            <TableHead>
              <TableRow>
                <TableCell size="small">No.</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Price</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {order.cartItems.map((cartItem, index) => (
                <TableRow key={`${cartItem.product.name}_${cartItem.id}`}>
                  <TableCell size="small">{index + 1}</TableCell>
                  <TableCell align="left">{cartItem.product.name}</TableCell>
                  <TableCell align="left">{currencyFormatter.format(cartItem.product.price)}</TableCell>
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
