import React, {Component} from "react";
import {MotorTypes, renderMotors} from "./data/MotorTypes";
import {Alert, Button, FormGroup, NumericInput, Slider, Switch} from "@blueprintjs/core";
import {Select} from "@blueprintjs/select/lib/esm/index";

const MotorSelect = Select.ofType();

class AdvancedTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeMotorType: MotorTypes[0],
      canID: -1,
      currentLimitEnabled: false,
      currentLimit: 40.0,
      isCoastMode: false,
      deadband: 0.0,
      currentProfile: 1,
      currentSetpoint: 0,
      velocityProfile: 1,
      velocitySetpoint: 0,
      positionProfile: 0,
      positionSetpoint: 0,
      slaveMode: false,
      masterID: -1,
      outputRampLimitEnabled: false,
      outputRampLimit: 0.0,
      inputRampLimitEnabled: false,
      inputRampLimit: 0.0,
      updateRequested: false
    };

    /* Form control methods */
    this.changeMotorType = this.changeMotorType.bind(this);
    this.changeCanID = this.changeCanID.bind(this);
    this.changeCurrentLimitEnabled = this.changeCurrentLimitEnabled.bind(this);
    this.changeIdleMode = this.changeIdleMode.bind(this);
    this.changeCurrentLimit = this.changeCurrentLimit.bind(this);
    this.changeDeadband = this.changeDeadband.bind(this);
    this.changeSlaveMode = this.changeSlaveMode.bind(this);
    this.changeMasterID = this.changeMasterID.bind(this);
    this.changeOutputLimitEnabled = this.changeOutputLimitEnabled.bind(this);
    this.changeOutputLimitRate = this.changeOutputLimitRate.bind(this);
    this.changeInputLimitEnabled = this.changeInputLimitEnabled.bind(this);
    this.changeInputLimitRate = this.changeInputLimitRate.bind(this);

    /* Form control for profile control methods */
    this.changeCurrentProfile = this.changeCurrentProfile.bind(this);
    this.changeCurrentSetpoint = this.changeCurrentSetpoint.bind(this);
    this.changePositionProfile = this.changePositionProfile.bind(this);
    this.changePositionSetpoint = this.changePositionSetpoint.bind(this);
    this.changeVelocityProfile = this.changeVelocityProfile.bind(this);
    this.changeVelocitySetpoint = this.changeVelocitySetpoint.bind(this);

    this.openConfirmModal = this.openConfirmModal.bind(this);
    this.closeConfirmModal = this.closeConfirmModal.bind(this);
    this.updateConfiguration = this.updateConfiguration.bind(this);
  }

  changeCanID(number) {
    this.setState({canID: number});
  }

  changeMotorType(motorType) {
    this.setState({activeMotorType: motorType});
  }

  changeCurrentLimitEnabled() {
    this.setState({currentLimitEnabled: !this.state.currentLimitEnabled});
  }

  changeIdleMode() {
    this.setState({isCoastMode: !this.state.isCoastMode});
  }

  changeCurrentLimit(value) {
    this.setState({currentLimit: value});
  }

  changeDeadband(value) {
    this.setState({deadband: value});
  }

  openConfirmModal() {
    this.setState({updateRequested: true});
  }

  closeConfirmModal() {
    this.setState({updateRequested: false});
  }

  changeCurrentProfile(value) {
    this.setState({currentProfile: value});
  }

  changeCurrentSetpoint(value) {
    this.setState({currentSetpoint: value});
  }

  changeVelocityProfile(value) {
    this.setState({velocityProfile: value});
  }

  changeVelocitySetpoint(value) {
    this.setState({velocitySetpoint: value});
  }

  changePositionProfile(value) {
    this.setState({positionProfile: value});
  }

  changePositionSetpoint(value) {
    this.setState({positionSetpoint: value});
  }

  changeSlaveMode() {
    this.setState({slaveMode: !this.state.slaveMode});
  }

  changeMasterID(value) {
    this.setState({masterID: value});
  }

  changeOutputLimitEnabled() {
    this.setState({outputRampLimitEnabled: !this.state.outputRampLimitEnabled});
  }

  changeOutputLimitRate(value) {
    this.setState({outputRampLimit: value});
  }

  changeInputLimitEnabled() {
    this.setState({inputRampLimitEnabled: !this.state.inputRampLimitEnabled});
  }

  changeInputLimitRate(value) {
    this.setState({inputRateLimit: value});
  }

  updateConfiguration() {
    console.log("Configuration updated!");
    // TODO - !!!!
  }

  render() {
    const {
      activeMotorType, isCoastMode, currentLimitEnabled, currentLimit, canID, deadband, updateRequested,
      currentProfile, currentSetpoint, positionProfile, positionSetpoint, velocityProfile, velocitySetpoint,
      outputRampLimitEnabled, outputRampLimit, inputRampLimitEnabled, inputRampLimit,
      slaveMode, masterID
    } = this.state;
    return (
      <div>
        <Alert isOpen={updateRequested} cancelButtonText="Cancel" confirmButtonText="Yes, Update" intent="SUCCESS" onCancel={this.closeConfirmModal} onClose={this.closeConfirmModal} onConfirm={this.updateConfiguration}>
          Are you sure you want to update the configuration of your SPARK controller to a {activeMotorType.name} motor?
        </Alert>
        <div className="form">
          <FormGroup
            label="Select Motor Type"
            labelFor="advanced-motor-type"
            className="form-group-two-fifths"
          >
            <MotorSelect id="advanced-motor-type" filterable={false} items={MotorTypes} itemRenderer={renderMotors} onItemSelect={this.changeMotorType}>
              <Button fill={true} disabled={!this.props.connected} text={activeMotorType.name} rightIcon="double-caret-vertical" />
            </MotorSelect>
          </FormGroup>
          <FormGroup
            label="Current Limit"
            labelFor="advanced-has-limit"
            className="form-group-fifth"
          >
            <Switch checked={currentLimitEnabled} disabled={!this.props.connected} label={currentLimitEnabled ? "On" : "No Limit"} onChange={this.changeCurrentLimitEnabled} />
          </FormGroup>
          <FormGroup
            label="Manual Limit"
            labelFor="advanced-current-limit"
            className="form-group-fifth"
          >
            <NumericInput id="advanced-current-limit" value={currentLimit} disabled={!currentLimitEnabled} onValueChange={this.changeCurrentLimit} stepSize={0.5} min={0} max={100}/>
          </FormGroup>
          <FormGroup
            label="Can ID"
            labelFor="advanced-can-id"
            className="form-group-fifth"
          >
            <NumericInput id="advanced-can-id" disabled={!this.props.connected} value={canID} onValueChange={this.changeCanID} min={0} max={24}/>
          </FormGroup>
        </div>
        <div className="form">
          <FormGroup
            label="Idle Mode"
            labelFor="advanced-idle-mode"
            className="form-group-quarter"
          >
            <Switch checked={isCoastMode} disabled={!this.props.connected} label={isCoastMode ? "Coast" : "Brake"} onChange={this.changeIdleMode} />
          </FormGroup>
          <FormGroup
            label="Motor Deadband"
            labelFor="advanced-deadband"
            className="form-group-three-quarters"
          >
            <Slider initialValue={deadband} disabled={!this.props.connected} value={deadband} min={0} max={1.0} stepSize={0.01} onChange={this.changeDeadband} />
          </FormGroup>
        </div>
        <div className="form">
          <FormGroup
            label="Current Profile And Setpoint"
            labelFor="advanced-current-profile"
            className="form-group-half inline"
          >
            <NumericInput id="advanced-current-profile" disabled={!this.props.connected} value={currentProfile} onValueChange={this.changeCurrentProfile} min={1} max={4}/>
            <NumericInput id="advanced-current-setpoint" disabled={!this.props.connected} value={currentSetpoint} onValueChange={this.changeCurrentSetpoint} min={1} max={1024}/>
          </FormGroup>
          <FormGroup
            label="Velocity Profile And Setpoint"
            labelFor="advanced-velocity-profile"
            className="form-group-half inline"
          >
            <NumericInput id="advanced-velocity-profile" disabled={!this.props.connected} value={velocityProfile} onValueChange={this.changeVelocityProfile} min={1} max={4}/>
            <NumericInput id="advanced-velocity-setpoint" disabled={!this.props.connected} value={velocitySetpoint} onValueChange={this.changeVelocitySetpoint} min={1} max={1024}/>
          </FormGroup>
        </div>
        <div className="form">
          <FormGroup
            label="Position Profile And Setpoint"
            labelFor="advanced-position-profile"
            className="form-group-half inline"
          >
            <NumericInput id="advanced-position-profile" disabled={!this.props.connected} value={positionProfile} onValueChange={this.changePositionProfile} min={1} max={4}/>
            <NumericInput id="advanced-position-setpoint" disabled={!this.props.connected} value={positionSetpoint} onValueChange={this.changePositionSetpoint} min={1} max={1024}/>
          </FormGroup>
          <FormGroup
            label="Slave Mode"
            labelFor="advanced-is-slave"
            className="form-group-quarter"
          >
            <Switch checked={slaveMode} disabled={!this.props.connected} label={slaveMode ? "Enabled" : "Disabled"} onChange={this.changeSlaveMode} />
          </FormGroup>
          <FormGroup
            label="Master ID"
            labelFor="advanced-master-id"
            className="form-group-quarter"
          >
            <NumericInput id="advanced-master-id" value={masterID} disabled={!slaveMode} onValueChange={this.changeMasterID} min={0} max={24}/>
          </FormGroup>
        </div>
        <div className="form">
          <FormGroup
            label="Output Ramp Limit"
            labelFor="advanced-output-limit"
            className="form-group-quarter"
          >
            <Switch checked={outputRampLimitEnabled} disabled={!this.props.connected} label={outputRampLimitEnabled ? "Enabled" : "Disabled"} onChange={this.changeOutputLimitEnabled} />
          </FormGroup>
          <FormGroup
            label="Ramp Rate Limit"
            labelFor="advanced-output-rate"
            className="form-group-quarter"
          >
            <NumericInput id="advanced-output-rate" value={outputRampLimit} disabled={!outputRampLimitEnabled} onValueChange={this.changeOutputLimitRate} min={0} max={1024}/>
          </FormGroup>
          <FormGroup
            label="Input Ramp Limit"
            labelFor="advanced-input-limit"
            className="form-group-quarter"
          >
            <Switch checked={inputRampLimitEnabled} disabled={!this.props.connected} label={inputRampLimitEnabled ? "Enabled" : "Disabled"} onChange={this.changeInputLimitEnabled} />
          </FormGroup>
          <FormGroup
            label="Ramp Rate Limit"
            labelFor="advanced-input-rate"
            className="form-group-quarter"
          >
            <NumericInput id="advanced-input-rate" value={inputRampLimit} disabled={!inputRampLimitEnabled} onValueChange={this.changeInputLimitRate} min={0} max={1024}/>
          </FormGroup>
        </div>
        <div className="form">
          <Button className="rev-btn" disabled={!this.props.connected} onClick={this.openConfirmModal}>Update Configuration</Button>
        </div>
      </div>
    );
  }
}

export default AdvancedTab;