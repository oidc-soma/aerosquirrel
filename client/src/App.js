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


function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" title="Dashboard - Aerosqirrel" element={<Accounts />}></Route>
        <Route path="/accounts/*" element={<Accounts />}></Route>
        <Route path="/inventory/*" element={<Inventory />}></Route>
        <Route path="/dashboard/*" element={<Dashboard />}></Route>
        <Route path="/settings/*" element={<Settings />}></Route>
        <Route path="/signup/*" element={<SignUp />}></Route>
        <Route path="/drawer/*" element={<Drawer />}></Route>
      </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
