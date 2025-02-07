import { FC } from 'react';
import type { NavigationProps, NavItem } from '@/types';

export const MainNav: FC<NavigationProps> = ({ navItems }) => {
  return (
    <nav>
      <ul>
        {navItems.map((item: NavItem) => {
          return (
            <li key={item.label}>{item.label}</li>
          );
        })}
      </ul>
    </nav>
  );
};
