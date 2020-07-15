import React, { useEffect, useState } from 'react';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Container from '@material-ui/core/Container';
import Divider from '@material-ui/core/Divider';
import Grid from '@material-ui/core/Grid';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import ResponsiveDrawer from './components/ResponsiveDrawer';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core';
import { useCookies } from 'react-cookie';

import HomeIcon from '@material-ui/icons/Home';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';

const API_BASE_URL = 'https://xenelectronic.herokuapp.com';

const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  root: {
    maxWidth: 345,
  },
  media: {
    height: 140,
  },
}))

function App() {
  const classes = useStyles();
  const [cookies, setCookie] = useCookies(['cartId']);
  const [categories, setCategories] = useState<string[]>([]);

  useEffect(() => {
    if (cookies.cartId === null || cookies.cartId === undefined || cookies.cartId === '') {
      fetch(`${API_BASE_URL}/carts`, {
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

    if (categories.length === 0) {
      fetch(`${API_BASE_URL}/categories`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/xenelectronic.v1+json'
        }
      })
        .then(response => response.json())
        .then(response => {
          if (Array.isArray(response)) {
            const newCategories = [...categories];
            response.forEach(category => newCategories.push(category.name));
            setCategories(newCategories);
          }
        });
    }
  });

  const drawer = (
    <React.Fragment>
      <List>
        <ListItem button key={"Home"}>
          <ListItemIcon><HomeIcon /></ListItemIcon>
          <ListItemText primary={"Home"} />
        </ListItem>
        <ListItem button key={"Cart"}>
          <ListItemIcon><ShoppingCartIcon /></ListItemIcon>
          <ListItemText primary={"Cart"} />
        </ListItem>
        <Divider />
      </List>
      <List subheader={
        <ListSubheader component="div">
          Categories
        </ListSubheader>
      }>
      {categories.map((text, index) => (
        <ListItem button key={text}>
          <ListItemText primary={text} />
        </ListItem>
      ))}
      </List>
    </React.Fragment>
  );

  return (
    <Container maxWidth="lg">
      <ResponsiveDrawer title="XenElectronic" drawer={drawer}>
        <Grid container spacing={3}>
          {
            [...Array(81)].map(() => (
              <Grid item xs={12} sm={6} md={4}>
                <Card className={classes.root}>
                  <CardActionArea>
                    <CardMedia
                      className={classes.media}
                      image="https://sabbaytinh.com/media/original/no-image.png"
                      title="Product"
                    />
                    <CardContent>
                      <Typography gutterBottom variant="h5" component="h2">
                        Product
                      </Typography>
                      <Typography variant="body2" color="textSecondary" component="p">
                        Product Description
                      </Typography>
                    </CardContent>
                  </CardActionArea>
                  <CardActions>
                    <Button size="small" color="primary">
                      Add to Cart
                    </Button>
                  </CardActions>
                </Card>
              </Grid>
            ))
          }
        </Grid>
      </ResponsiveDrawer>
    </Container>
  );
}

export default App;
