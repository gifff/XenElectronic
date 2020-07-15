import React, { useEffect } from 'react';

import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import ResponsiveDrawer from './components/ResponsiveDrawer';
import { makeStyles } from '@material-ui/core';
import { useCookies } from 'react-cookie';

const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
}))

function App() {
  const classes = useStyles();
  const [cookies, setCookie] = useCookies(['cartId']);

  useEffect(() => {
    if (cookies.cartId === null || cookies.cartId === undefined || cookies.cartId === '') {
      fetch('https://xenelectronic.herokuapp.com/carts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/xenelectronic.v1+json'
        }
      })
        .then(response => response.json())
        .then(response => {
          setCookie('cartId', response['cart_id'], { path: '/' });
        });
    }
  });

  return (
    <Container maxWidth="lg">
      <ResponsiveDrawer title="XenElectronic">
        <Grid container spacing={3}>
          {
            [...Array(81)].map(() => (
              <Grid item xs={12} sm={6} md={4}>
                <Paper className={classes.paper}>Product</Paper>
              </Grid>
            ))
          }
        </Grid>
      </ResponsiveDrawer>
    </Container>
  );
}

export default App;
