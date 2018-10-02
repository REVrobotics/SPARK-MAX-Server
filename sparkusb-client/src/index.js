import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';

window.eval = global.eval = () => {
    throw new Error("This app does not support window.eval().")
};

const ipcRenderer = window.require("electron").ipcRenderer;

const headless = false;

if (headless) {
  ReactDOM.render(<App />, document.getElementById('root'));
} else {
  ipcRenderer.once("start-server-success", () => {
    ReactDOM.render(<App />, document.getElementById('root'));
  });
  ipcRenderer.once("start-server-error", (event, error) => {
    console.log(error);
  });
  ipcRenderer.send("start-server");
}