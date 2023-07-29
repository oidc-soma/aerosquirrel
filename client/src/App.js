import logo from './logo.svg';
import React from 'react';
import Navbar from './components/navbar/Navbar';
import './App.css';
import Accounts from './pages/Accounts';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
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

function App() {

  return (
    <div className="App">
      <RecoilRoot>
      <BrowserRouter>
      
      <Routes>
        <Route path="/" title="Dashboard - Aerosqirrel" element={<><Navbar /><Accounts /></>}></Route>
        <Route path="/accounts/*" element={<><Navbar /><Accounts /></>}></Route>
        <Route path="/inventory/*" element={<><Navbar /><Inventory /></>}></Route>
        <Route path="/dashboard/*" element={<><Navbar /><Dashboard /></>}></Route>
        <Route path="/settings/*" element={<><Navbar /><Settings /></>}></Route>
        <Route path="/signup/*" element={<><SignUp /></>}></Route>
        <Route path="/login/*" element={<><Login /></>}></Route>
        <Route path="/drawer/*" element={<><Navbar /><Drawer /></>}></Route>
        <Route path="/addinst/*" element={<><Navbar /><AddInstancePrompt /></>}></Route>
      </Routes>
      </BrowserRouter>
      </RecoilRoot>
    </div>
  );
}

export default App;
