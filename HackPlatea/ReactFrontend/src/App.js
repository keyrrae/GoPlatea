import React, { Component } from 'react';
import hack_logo from './hack_logo.svg';
import CodeEditor from './CodeEditor';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={hack_logo} className="App-logo" alt="logo" />
          <h2>Welcome to HackPlatea</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
          <button className="App-button">
              run
          </button>

          <button className="App-button">
              About
          </button>
          <p/>

          <div>
              <CodeEditor/>
          </div>
      </div>
    );
  }
}

export default App;
