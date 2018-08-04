import React, {Component} from "react";
import {Button, FormGroup, NumericInput} from "@blueprintjs/core";

class PIDTunerTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      profile: 0,
      p: 0,
      i: 0,
      d: 0,
      f: 0
    };
    this.changeProfile = this.changeProfile.bind(this);
    this.changeP = this.changeP.bind(this);
    this.changeI = this.changeI.bind(this);
    this.changeD = this.changeD.bind(this);
    this.changeF = this.changeF.bind(this);
    this.updatePID = this.updatePID.bind(this);
  }

  changeProfile(value) {
    this.setState({profile: value});
  }

  changeP(value) {
    this.setState({p: value});
  }

  changeI(value) {
    this.setState({i: value});
  }

  changeD(value) {
    this.setState({d: value});
  }

  changeF(value) {
    this.setState({f: value});
  }

  updatePID() {
    console.log("PLEASE HELP ME");
  }

  render() {
    const {profile, p, i, d, f} = this.state;
    return (
      <div>
        <div className="form">
          <FormGroup
            label="PID Profile"
            className="form-group-fifth"
          >
            <NumericInput id="pid-profile" disabled={!this.props.connected} value={profile} onValueChange={this.changeProfile} min={0} max={3}/>
          </FormGroup>
          <FormGroup
            label="P"
            className="form-group-fifth"
          >
            <NumericInput id="pid-profile" disabled={!this.props.connected} value={p} onValueChange={this.changeP} min={0} max={3}/>
          </FormGroup>
          <FormGroup
            label="I"
            className="form-group-fifth"
          >
            <NumericInput id="pid-profile" disabled={!this.props.connected} value={i} onValueChange={this.changeI} min={0} max={3}/>
          </FormGroup>
          <FormGroup
            label="D"
            className="form-group-fifth"
          >
            <NumericInput id="pid-profile" disabled={!this.props.connected} value={d} onValueChange={this.changeD} min={0} max={3}/>
          </FormGroup>
          <FormGroup
            label="F"
            className="form-group-fifth"
          >
            <NumericInput id="pid-profile" disabled={!this.props.connected} value={f} onValueChange={this.changeF} min={0} max={3}/>
          </FormGroup>
        </div>
        <div className="form">
          <Button className="rev-btn" disabled={!this.props.connected} onClick={this.updatePID}>Update PIDF Configuration</Button>
        </div>
      </div>
    );
  }
}

export default PIDTunerTab;