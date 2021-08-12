import React, { Component } from 'react';
import { Router, browserHistory, Route, Link } from 'react-router';
import logo from './logo.svg';
import './App.css';
import Login from './Login.js'
import LinkElement from './LinkElement.js'
import LinksArray from './LinksArray.js'
import portal from './portal-pixel-art.gif'
import { ReactTinyLink } from 'react-tiny-link'

const Page = ({ title }) => (
    <div className="App">
      <div className="App-header">
        <img src={portal} className="App-logo" alt="logo" />
        <br></br>
        <h2>{title}</h2>
      </div>
      <p className="App-intro">
      </p>
      <p>
        <Link to="/">LinkPortal</Link>
      </p>
      <LinksArray/>
      
    </div>
);

const Home = (props) => (
  <Page title="LinkPortal"/>
);

const About = (props) => (
  <Page title="About"/>
);

const Settings = (props) => (
  <Page title="Settings"/>
);

const login = (props) => (
  <Page title="login"/>
);

class App extends Component {
  render() {
    return (
      <div>
      <Router history={browserHistory}>
        <Route path="/" component={Home}/>
        <Route path="/login"><Login/></Route>
        <Route path="/about" component={About}/>
        <Route path="/settings" component={Settings}/>
      </Router>
      </div>
    );
  }
}

export default App;