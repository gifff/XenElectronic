import React, { ReactNode } from 'react';

import ListItem from '@material-ui/core/ListItem';
import {
  Link as RouterLink,
  LinkProps as RouterLinkProps,
} from 'react-router-dom';

interface ListItemLinkProps {
  children?: ReactNode;
  to: string;
}

export default function ListItemLink(props: ListItemLinkProps) {
  const { children, to } = props;

  const renderLink = React.useMemo(
    () => React.forwardRef<any, Omit<RouterLinkProps, 'to'>>((itemProps, ref) => (
      <RouterLink to={to} ref={ref} {...itemProps} />
    )),
    [to],
  );

  return (
    <li>
      <ListItem button component={renderLink}>
        {children}
      </ListItem>
    </li>
  );
}
