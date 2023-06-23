import logo from './logo.svg';
import React from 'react';
import Navbar from './components/navbar/Navbar';
import './App.css';
import Accounts from './pages/Accounts';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Accounts />}></Route>
        <Route path="/accounts/*" element={<Accounts />}></Route>
      </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
