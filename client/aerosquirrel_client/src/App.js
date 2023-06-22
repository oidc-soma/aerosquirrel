import logo from './logo.svg';
import React from 'react';
import Navbar from './components/navbar/Navbar';
import './App.css';
import Accounts from './pages/Accounts';

function App() {
  return (
    <div className="App">
      <Navbar />
      <Accounts />
    </div>
  );
}

export default App;
