import React, { Component } from 'react';

const ipcRenderer = window.require("electron").ipcRenderer;

class App extends Component {
  constructor() {
    super();
    this.state = {
      connected: false,
      canID: -1
    };
    this.retryConnection = this.retryConnection.bind(this);
    this.getCanID = this.getCanID.bind(this);
    this.listConnections = this.listConnections.bind(this);
  }

  componentDidMount() {
    this.retryConnection();
  }

  retryConnection() {
    ipcRenderer.once("test-response", (event, error, response) => {
        if (error) {
          if (error.details === "Access is denied.") {
              this.setState({connected: true});
          } else {
              this.setState({connected: false});
          }
        } else {
            this.setState({connected: true});
            this.getCanID();
        }
    });
    ipcRenderer.send("test");
  }

  getCanID() {
    ipcRenderer.once("get-can-id-response", (event, error, response) => {
      console.log(error, response);
    });
    ipcRenderer.send("get-can-id", 0);
  }

  listConnections() {
    ipcRenderer.once("list-devices-response", (event, error, response) => {
      console.log(error);
      console.log(response);
    });
    ipcRenderer.send("list-devices");
  }

  render() {
    const {connected, canID} = this.state;
    return (
      <div>
        <div>
          <span>REVBLDC Connection Status: {connected ? "CONNECTED" : "NOT CONNECTED"}</span>
        </div>
        <div>
          <button onClick={this.retryConnection}>Retry Connection</button>
        </div>
        <div>
          <button onClick={this.listConnections}>List Devices</button>
        </div>
        <div>
          <span>Motor Controller Parameters</span>
          <div>
            <span>CanID: {canID}</span>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
