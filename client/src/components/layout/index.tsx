import { FC } from 'react';
import { Outlet } from 'react-router';
import { MainNav } from '@/components/layout/mainNav';
import type { NavItem } from '@/types';

const navItems: NavItem[] = [
  {
    label: 'home',
  }
];

const Layout: FC = () => {
  return (
    <>
      <MainNav navItems={navItems} />
      <Outlet />
    </>
  );
};

export default Layout;
