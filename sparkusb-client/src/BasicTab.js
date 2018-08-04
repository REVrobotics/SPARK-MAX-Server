import React, {Component} from "react";
import {Alert, Button, FormGroup, NumericInput, Radio, RadioGroup, Switch} from "@blueprintjs/core";
import {Select} from "@blueprintjs/select";
import {MotorTypes, renderMotors} from "./data/MotorTypes";
import {saveBasicConfig} from "./ConfigurationManager";

const ipcRenderer = window.require("electron").ipcRenderer;

const MotorSelect = Select.ofType();

class BasicTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeMotorType:  this.props.parameters[2] === 1 ? MotorTypes[0] : MotorTypes[1],
      canID: this.props.parameters[0] || 0,
      currentLimit: 40,
      isCoastMode: this.props.parameters[6] === 0 || false,
      updateRequested: false,
      savingConfig: false
    };

    /* Form control methods */
    this.selectMotorType = this.selectMotorType.bind(this);
    this.changeCanID = this.changeCanID.bind(this);
    this.changeIdleMode = this.changeIdleMode.bind(this);
    this.changeCurrentLimit = this.changeCurrentLimit.bind(this);

    this.openConfirmModal = this.openConfirmModal.bind(this);
    this.closeConfirmModal = this.closeConfirmModal.bind(this);
    this.updateConfiguration = this.updateConfiguration.bind(this);
  }

  componentDidUpdate(prevProps) {
    if (prevProps.connected !== this.props.connected && this.props.connected === true) {
      console.log("Pre-populating values...");
      this.setState({
        activeMotorType:  this.props.parameters[2] === 1 ? MotorTypes[0] : MotorTypes[1],
        canID: this.props.parameters[0] || 0,
        currentLimit: 40,
        isCoastMode: this.props.parameters[6] === 0 || false,
        updateRequested: false,
        savingConfig: false
      });
    }
  }

  changeCanID(number) {
    this.setState({canID: number});
  }

  selectMotorType(motorType) {
    this.setState({activeMotorType: motorType});
  }

  changeIdleMode() {
    this.setState({isCoastMode: !this.state.isCoastMode});
  }

  changeCurrentLimit(value) {
    this.setState({currentLimit: parseInt(value.currentTarget.value)});
  }

  openConfirmModal() {
    this.setState({updateRequested: true});
  }

  closeConfirmModal() {
    this.setState({updateRequested: false});
  }

  updateConfiguration() {
    this.setState({savingConfig: true});
    const motorType = this.state.activeMotorType.kMotorType;
    const idleMode = this.state.isCoastMode ? 0 : 1;
    const canID = this.state.canID;
    saveBasicConfig(this.props.connectedDevice, canID, motorType, idleMode).then((values) => {
      this.setState({savingConfig: false});
    }).catch((error) => {
      this.setState({savingConfig: false});
      console.log(error);
    });
  }

  render() {
    const {activeMotorType, isCoastMode, currentLimit, canID, updateRequested, savingConfig} = this.state;
    return (
      <div>
        <Alert isOpen={updateRequested} cancelButtonText="Cancel" confirmButtonText="Yes, Update" intent="SUCCESS" onCancel={this.closeConfirmModal} onClose={this.closeConfirmModal} onConfirm={this.updateConfiguration}>
          Are you sure you want to update the configuration of your SPARK controller to a {activeMotorType.name} motor?
        </Alert>
        <div className="form">
          <FormGroup
            label="Select Motor Type"
            labelFor="basic-motor-type"
            className="form-group-half"
          >
            <MotorSelect id="basic-motor-type" filterable={false} items={MotorTypes} itemRenderer={renderMotors} onItemSelect={this.selectMotorType}>
              <Button fill={true} disabled={!this.props.connected} text={activeMotorType.name} rightIcon="double-caret-vertical" />
            </MotorSelect>
          </FormGroup>
          <FormGroup
            label="Can ID"
            labelFor="basic-can-id"
            className="form-group-quarter"
          >
            <NumericInput id="basic-can-id" value={canID} onValueChange={this.changeCanID} min={0} max={24} disabled={!this.props.connected}/>
          </FormGroup>
          <FormGroup
            label="Idle Mode"
            labelFor="basic-idle-mode"
            className="form-group-quarter"
          >
            <Switch checked={isCoastMode} disabled={!this.props.connected} label={isCoastMode ? "Coast" : "Brake"} onChange={this.changeIdleMode} />
          </FormGroup>
        </div>
        <div className="form">
          <FormGroup
            label="Current Limit"
            inline={true}
          >
            <RadioGroup
              inline={true}
              selectedValue={currentLimit}
              onChange={this.changeCurrentLimit}
              disabled={!this.props.connected}
            >
              <Radio label="20A" value={20}/>
              <Radio label="30A" value={30}/>
              <Radio label="40A" value={40}/>
              <Radio label="No Limit" value={-1}/>
            </RadioGroup>
          </FormGroup>
        </div>
        <div className="form">
          <Button className="rev-btn" disabled={!this.props.connected} loading={savingConfig} onClick={this.openConfirmModal}>Update Configuration</Button>
        </div>
      </div>
    );
  }
}

export default BasicTab;