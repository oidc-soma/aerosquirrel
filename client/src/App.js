import logo from './logo.svg';
import React from 'react';
import Navbar from './components/navbar/Navbar';
import './App.css';
import Accounts from './pages/Accounts';
import { BrowserRouter, Route, Routes, useLocation } from 'react-router-dom';
import Inventory from './pages/Inventory';
import Dashboard from './pages/Dashboard';
import Settings from './pages/Settings';
import styled from 'styled-components';
import AddAccount from './pages/AddAccount';
import SignUp from './pages/SignUp';
import Drawer from './pages/Drawer';
import Login from './pages/Login';
import AddInstancePrompt from './pages/AddInstancePrompt';
import { RecoilRoot, atom, selector, useRecoilState, useRecoilValue } from 'recoil';
import { InventoryAtom } from './atoms';
import { TransitionGroup, CSSTransition } from 'react-transition-group';
import AnimatedCursor from 'react-animated-cursor';
import WelcomePage from './pages/WelcomePage';
import {ToastContainer} from 'react-toastify';

const AnimatedRoutes = () => {
  const location = useLocation();

  return (
    <>
    <Navbar />
    <TransitionGroup className="transition-group">
      <CSSTransition key={location.key} classNames="fade" timeout={300}>
        <Routes location={location}>
          <Route path="/" title="Dashboard - Aerosqirrel" element={<><WelcomePage /></>}></Route>
          <Route path="/accounts/*" element={<Accounts />}></Route>
          <Route path="/inventory/*" element={<Inventory />}></Route>
          <Route path="/dashboard/*" element={<Dashboard />}></Route>
          <Route path="/settings/*" element={<Settings />}></Route>
          <Route path="/signup/*" element={<SignUp />}></Route>
          <Route path="/login/*" element={<Login />}></Route>
          <Route path="/drawer/*" element={<Drawer />}></Route>
          <Route path="/addinst/*" element={<AddInstancePrompt />}></Route>
          <Route path="/welcome/*" element={<WelcomePage />}></Route>
        </Routes>
      </CSSTransition>
    </TransitionGroup>
    </>
  );
};

function App() {
  return (
    <div className="App">
      <ToastContainer
        position="top-right"
        autoClose={3000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="light"
      />
      <RecoilRoot>
        <BrowserRouter>
          <AnimatedRoutes />
        </BrowserRouter>
      </RecoilRoot>
       <AnimatedCursor
      innerSize={10}
      outerSize={10}
      color='0, 0, 0'
      outerAlpha={0.2}
      innerScale={0.5}
      outerScale={5}
    />
    </div>
  );
}

export default App;
