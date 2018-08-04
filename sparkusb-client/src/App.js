import React, { Component } from 'react';
import {Tab, Tabs} from "@blueprintjs/core";
import BasicTab from "./BasicTab";
import AdvancedTab from "./AdvancedTab";
import ConnectionStatusBar from "./ConnectionStatusBar";
import RunTab from "./RunTab";
import PIDTunerTab from "./PIDTunerTab";
import FirmwareTab from "./FirmwareTab";
import {getAllParametersSequential} from "./ConfigurationManager";
import HelpTab from "./HelpTab";

const ipcRenderer = window.require("electron").ipcRenderer;

class App extends Component {
  constructor() {
    super();
    this.state = {
      connected: false,
      connecting: false,
      connectedDevice: "",
      connectionStatus: "CONNECTION FAILED",
      requestParam: 0,
      response: "",
      appLogs: [],
      parameters: []
    };
    this.retryConnection = this.retryConnection.bind(this);
    this.disconnect = this.disconnect.bind(this);
    this.pushAppLog = this.pushAppLog.bind(this);

    ipcRenderer.on("disconnection", (event, device) => {
      if (this.state.connected) {
        console.log("USB unplugged. Forcing disconnection...");
        this.disconnect();
      }
    });
  }

  pushAppLog(log) {
    this.state.appLogs.push(log);
    this.setState({appLogs: this.state.appLogs});
  }

  componentDidMount() {
    this.retryConnection();
  }

  disconnect() {
    this.setState({connecting: true, connected: false, connectionStatus: "DISCONNECTING..."});
    ipcRenderer.once("disconnect-response", (event, error, response) => {
      console.log(error, response);
      this.setState({connecting: false, connected: false, connectionStatus: "READY TO CONNECT"});
    });
    ipcRenderer.send("disconnect", this.state.connectedDevice);
  }

  retryConnection() {
    this.setState({connecting: true, connectionStatus: "CONNECTING...", connectedDevice: "", parameters: []});
    this.listDevices().then((deviceList) => {
      if (deviceList.length > 0) {
        ipcRenderer.once("connect-response", (event, error, response) => {
          if (error) {
            this.pushAppLog("Error while connecting to device. " + error.details);
            if (error.details === "Access is denied.") {
              getAllParametersSequential().then((values) => {
                console.log(values);
                this.pushAppLog("Successfully pulled device parameters.");
                this.setState({connected: true, connecting: false, connectionStatus: "CONNECTED", connectedDevice: deviceList[0], parameters: values});
              }).catch((error) => {
                this.pushAppLog("Error while getting device parameters. " + error.details);
                this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
              });
            } else {
              this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
            }
          } else {
            getAllParametersSequential().then((values) => {
              console.log(values);
              this.pushAppLog("Successfully pulled device parameters.");
              this.setState({connected: true, connecting: false, connectionStatus: "CONNECTED", connectedDevice: deviceList[0], parameters: values});
            }).catch((error) => {
              this.pushAppLog("Error while getting device parameters. " + error.details);
              this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
            });
          }
        });
        ipcRenderer.send("connect", deviceList[0]);
      } else {
        this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
      }
    }).catch((error) => {
      this.setState({connected: false, connecting: false, connectionStatus: "CONNECTION FAILED"});
      this.pushAppLog("Error while connecting to device. " + error.details);
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
    const {connected, connecting, connectionStatus, connectedDevice, appLogs, parameters} = this.state;
    return (
      <div id="main-container">
        <ConnectionStatusBar connected={connected} connecting={connecting} connectionStatus={connectionStatus} onConnect={connected ? this.disconnect : this.retryConnection} onLog={this.pushAppLog} />
        <Tabs id="main-tabs" defaultSelectedTabId="main-tab-basic" renderActiveTabPanelOnly={true}>
          <Tab id="main-tab-basic" title="Basic" panel={<BasicTab connectedDevice={connectedDevice} connected={connected} parameters={parameters} />} />
          <Tab id="main-tab-advanced" title="Advanced" panel={<AdvancedTab connected={connected} />} />
          <Tab id="main-tab-run" title="Run" panel={<RunTab connected={connected} />} />
          <Tab id="main-tab-pid" title="PID Tuner" panel={<PIDTunerTab connected={connected} />} />
          <Tab id="main-tab-network" title="Network" panel={<span>Network</span>} />
          <Tab id="main-tab-firmware" title="Firmware" panel={<FirmwareTab connected={connected} />} />
          <Tab id="main-tab-help" title="Help" panel={<HelpTab logs={appLogs}/>} />
        </Tabs>
      </div>
    );
  }
}

export default App;
