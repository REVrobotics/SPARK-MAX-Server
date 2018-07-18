import React, {Component} from "react";
import {Button} from "@blueprintjs/core";

const ipcRenderer = window.require("electron").ipcRenderer;

class FirmwareTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      outputText: ["Please connect to a device to see it's firmware.", "Connected. Loading firmware..."],
      firmwareVersion: "0.0.1"
    };
    this.openFirmwareDialog = this.openFirmwareDialog.bind(this);
  }

  openFirmwareDialog() {
    ipcRenderer.once("request-firmware-response", (event, filePath) => {
      this.state.outputText.push("Loading firmware from " + filePath + "...");
      this.setState({outputText: this.state.outputText});
    });
    ipcRenderer.send("request-firmware");
  }

  render() {
    const {outputText, firmwareVersion} = this.state;
    return (
      <div>
        <div id="firmware-console">
          {outputText.map(line => {
            return <p>{line}</p>;
          })}
        </div>
        <div id="firmware-bar">
          <span>Current Firmware: {firmwareVersion}</span>
          <span><Button className="rev-btn" disabled={!this.props.connected} onClick={this.openFirmwareDialog}>Load Firmware</Button></span>
        </div>
      </div>
    );
  }
}

export default FirmwareTab;