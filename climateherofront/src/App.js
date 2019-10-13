import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { Main } from './components/Main';
import { CreateAccount } from './components/CreateAccount';


export default class App extends Component {
  render() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Welcome to ClimateHero, Emma! 
        </p>
        <img src={logo} className="App-logo" alt="logo" />
        
        
      </header>
    </div>
  );
  }
}

