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
import { makeStyles, CircularProgress } from '@material-ui/core';

import Product from '../../lib/model/Product';
import { useParams } from 'react-router-dom';
import { useCookies } from 'react-cookie';

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

const currencyFormatter = new Intl.NumberFormat('id-ID', {
  style: 'currency',
  currency: 'IDR',
})

export default function ProductList() {
  const classes = useStyles();
  const [products, setProducts] = useState<Product[]>([]);
  const [previousCategoryId, setPreviousCategoryId] = useState<number>(0);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);
  const { categoryId: categoryIdParam } = useParams();
  const categoryId: number = isNaN(categoryIdParam) ? 0 : Number(categoryIdParam);
  const [{ cartId }] = useCookies(['cartId']);

  useEffect(() => {
    if (!isLoading && categoryId > 0 && previousCategoryId !== categoryId) {
      setIsLoading(true);
      setError(null);
      client.listProductByCategory(categoryId)
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
          }
        })
        .catch(error => {
          setProducts([]);
          setError(error)
        })
        .finally(() => {
          setIsLoading(false);
          setPreviousCategoryId(categoryId);
        })
    }
  }, [isLoading, categoryId, previousCategoryId]);

  const addProductToCart = (productId: number) => {
    client.addProductToCart(cartId, productId)
  };

  const loadingContent = (
    <Grid item xs={12}>
      <CircularProgress />
    </Grid>
  );

  const content = products.map(product => (
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
            <Typography variant="body2" color="textPrimary" component="p" gutterBottom>
              {currencyFormatter.format(product.price)}
            </Typography>
            <Typography variant="body1" color="textSecondary" component="p">
              {product.description}
            </Typography>
          </CardContent>
        </CardActionArea>
        <CardActions>
          <Button size="small" color="primary" onClick={() => addProductToCart(product.id)}>
            Add to Cart
          </Button>
        </CardActions>
      </Card>
    </Grid>
  ));

  const invalidCategoryContent = (
    <Grid item xs={12}>
      <Typography variant="body1">Select category first</Typography>
    </Grid>
  );

  const errorContent = (
    <Grid item xs={12}>
      <Typography variant="body1" color="error">Sorry, something is broken</Typography>
    </Grid>
  );

  return (
    <Grid container spacing={3}>
      {
        isLoading ? loadingContent : (
          error !== null ? errorContent : (
            categoryId > 0 ? content : invalidCategoryContent
          )
        )
      }
    </Grid>
  );
}