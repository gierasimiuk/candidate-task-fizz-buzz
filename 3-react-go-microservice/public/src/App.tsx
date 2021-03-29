import React, { Component } from 'react';
import './App.css';

export default class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1>Welcome to the FizzBuzz calculator</h1>
          <p>Enter the start and end numbers for the calculation and then press Go!</p>
          <form action="http://127.0.0.1:8000/fizzbuzz" method="post">
            <label>Start: </label>
            <input type="text" name="start" />
            <label>End: </label>
            <input type="text" name="end" />
            <input type="submit" value="Go!" id="go-button" />
          </form>
        </header>
      </div>
    );
  }
}

// export default App;
