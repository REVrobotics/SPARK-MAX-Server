import React, {Component} from "react";
import {Button} from "@blueprintjs/core";

class ConnectionStatusBar extends Component {
  constructor(props) {
    super(props);
    this.attemptConnection = this.attemptConnection.bind(this);
  }

  attemptConnection() {
    if (typeof this.props.onConnect === "function") {
        this.props.onConnect();
    }
  }

  render() {
    return (
      <div id="status-bar">
        <span id="status-bar-status">Motor Controller Connection Status: {this.props.connectionStatus}</span>
        <span id="status-bar-button"><Button fill={true} disabled={this.props.connected} loading={this.props.connecting} onClick={this.attemptConnection}>Connect</Button></span>
      </div>
    );
  }
}

export default ConnectionStatusBar;