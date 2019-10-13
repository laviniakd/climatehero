import React, { Component } from 'react';
import logo from '../logo.svg';
import '../App.css';

export class Main extends Component {
    displayName = Main.name

    render() {
        return (
            <div className="main">
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