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
          <h3 style={{marginLeft: 20}}>Welcome to HackPlatea</h3>
        </div>
        <p className="App-intro">
          Enter your Hacklang code, and click run to execute the program.
        </p>
          <div>
              <CodeEditor/>
          </div>
      </div>
    );
  }
}

export default App;
