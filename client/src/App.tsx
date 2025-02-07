import { Suspense, lazy } from 'react';
import { Routes, Route } from 'react-router';
import Layout from '@/components/layout';
import './App.css'

// PAGES
const Home = lazy(() => import('@/pages/home'));

function App() {

  return (
    <Suspense>
      <Routes>
        <Route path='/' element={<Layout />}>
          <Route index element={<Home />} />
        </Route>
      </Routes>
    </Suspense>
  )
}

export default App
