import React, { useState, useEffect } from 'react';
import client from '../../lib/client';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core';

import Product from '../../lib/model/Product';

interface Props {
  categoryId: number
}

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

export default function ProductList(props: Props) {
  const classes = useStyles();
  const [products, setProducts] = useState<Product[]>([]);
  const [categoryId, setCategoryId] = useState<number>(0);

  useEffect(() => {
    if (props.categoryId > 0 && categoryId !== props.categoryId) {
      client.listProductByCategory(props.categoryId)
        .then(response => response.json())
        .then(response => {
          if (Array.isArray(response)) {
            const newProducts: Product[] = response.map(product => new Product(
              product.id,
              product.categoryId,
              product.name,
              product.description,
              product.photo,
              product.price,
            ));
            setProducts(newProducts);
            setCategoryId(props.categoryId);
          }
        })
    }
  });

  return (
    <Grid container spacing={3}>
      {
        categoryId > 0 ?
          products.map(product => (
            <Grid item xs={12} sm={6} md={4}>
              <Card className={classes.root}>
                <CardActionArea>
                  <CardMedia
                    className={classes.media}
                    image={product.photo !== '' ? product.photo : 'https://sabbaytinh.com/media/original/no-image.png'}
                    title={product.name}
                  />
                  <CardContent>
                    <Typography gutterBottom variant="h5" component="h2">
                      {product.name}
                    </Typography>
                    <Typography variant="body2" color="textSecondary" component="p">
                      {product.description}
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
          )) : (
            <Grid item xs={12}>
              <Typography variant='body1'>Select category first</Typography>
            </Grid>
          )
      }
    </Grid>
  );
}
