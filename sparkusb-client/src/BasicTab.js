import React, {Component} from "react";
import {Button, FormGroup, NumericInput} from "@blueprintjs/core";

const ipcRenderer = window.require("electron").ipcRenderer;

class BasicTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      canID: 0
    };

    this.changeCanID = this.changeCanID.bind(this);
    this.setCanID = this.setCanID.bind(this);
  }

  componentDidUpdate(prevProps) {
    if (prevProps.connected !== this.props.connected && this.props.connected === true) {
      this.getCanID();
    }
  }

  getCanID() {
    ipcRenderer.once("get-param-response", (event, error, response) => {
      console.log("Response!");
      console.log(error, response);
      if (!error) {
        this.setState({canID: response.value});
      }
    });
    ipcRenderer.send("get-param", 0);
  }

  setCanID() {
    console.log(this.state.canID);
    ipcRenderer.once("set-param-response", (event, error, response) => {
      console.log(error, response);
    });
    ipcRenderer.send("set-param", 0, this.state.canID);
  }

  changeCanID(number) {
    this.setState({canID: number});
  }

  render() {
    const {canID} = this.state;
    return (
      <div className="form">
        <FormGroup
          label="Set Can ID"
          labelFor="basic-can-id"
        >
          <NumericInput id="basic-can-id" value={canID} onValueChange={this.changeCanID} min={0} max={24} disabled={!this.props.connected}/>
        </FormGroup>
        <FormGroup
          label="Get Can ID"
          labelFor="basic-can-id-get"
        >
          <Button id="basic-can-id-get" disabled={!this.props.connected} onClick={this.getCanID}>Get CAN ID</Button>
        </FormGroup>
        {/*<span><Button onClick={this.setCanID} disabled={!this.props.connected}>Set CAN ID</Button></span>*/}
      </div>
    );
  }
}

export default BasicTab;