import React, { Component } from 'react';
import {Tab, Tabs} from "@blueprintjs/core";
import BasicTab from "./BasicTab";
import AdvancedTab from "./AdvancedTab";
import ConnectionStatusBar from "./ConnectionStatusBar";

const ipcRenderer = window.require("electron").ipcRenderer;

class App extends Component {
  constructor() {
    super();
    this.state = {
      connected: false,
      connecting: false,
      connectionStatus: "CONNECTION FAILED",
      requestParam: 0,
      response: ""
    };
    this.retryConnection = this.retryConnection.bind(this);
  }

  componentDidMount() {
    this.retryConnection();
  }

  retryConnection() {
    this.setState({connecting: true});
    this.listDevices().then((deviceList) => {
      if (deviceList.length > 0) {
        ipcRenderer.once("connect-response", (event, error, response) => {
          setTimeout(() => {
            if (error) {
              if (error.details === "Access is denied.") {
                this.setState({connected: true, connecting: false, connectionStatus: "CONNECTED"});
              } else {
                this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
              }
            } else {
              this.setState({connected: true, connecting: false, connectionStatus: "CONNECTED"});
            }
          }, 500);
        });
        ipcRenderer.send("connect", deviceList[0]);
      } else {
        this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
      }
    }).catch((error) => {
      this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
      console.log(error);
    });
  }

  listDevices() {
    return new Promise((resolve, reject) => {
      ipcRenderer.once("list-devices-response", (event, error, response) => {
        if (error) {
          reject(error);
        } else {
          resolve(response.deviceList);
        }
      });
      ipcRenderer.send("list-devices");
    });
  }

  render() {
    const {connected, connecting, connectionStatus, canID} = this.state;
    return (
      <div id="main-container">
        <ConnectionStatusBar connected={connected} connecting={connecting} connectionStatus={connectionStatus} onConnect={this.retryConnection} />
        <Tabs id="main-tabs" defaultSelectedTabId="main-tab-basic">
          <Tab id="main-tab-basic" title="Basic" panel={<BasicTab connected={connected} />} />
          <Tab id="main-tab-advanced" title="Advanced" panel={<AdvancedTab/>} />
        </Tabs>
      </div>
    );
  }
}

export default App;
