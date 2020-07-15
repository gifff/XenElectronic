import React, { ReactNode, MouseEventHandler } from 'react';

import ListItem from '@material-ui/core/ListItem';
import {
  Link as RouterLink,
  LinkProps as RouterLinkProps,
} from 'react-router-dom';

interface ListItemLinkProps {
  children?: ReactNode;
  selected?: boolean;
  to: string;
  onClick?: MouseEventHandler<HTMLAnchorElement>
}

export default function ListItemLink(props: ListItemLinkProps) {
  const { children, selected, to, onClick } = props;

  const renderLink = React.useMemo(
    () => React.forwardRef<any, Omit<RouterLinkProps, 'to'>>((itemProps, ref) => (
      <RouterLink to={to} ref={ref} {...itemProps} />
    )),
    [to],
  );

  return (
    <li>
      <ListItem button component={renderLink} selected={selected} onClick={onClick} >
        {children}
      </ListItem>
    </li>
  );
}
