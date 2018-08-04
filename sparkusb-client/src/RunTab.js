import React, {Component} from "react";
import {Button, FormGroup, Slider, Radio, RadioGroup} from "@blueprintjs/core";
import ReactEcharts from "echarts-for-react";
import {disableHeartbeat, enableHeartbeat, setSetpoint} from "./ConfigurationManager";

const ipcRenderer = window.require("electron").ipcRenderer;

class RunTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      option: {
        xAxis: {
          type: 'category',
          name: "Time"
        },
        yAxis: {
          type: 'value',
          name: "Percent"
        },
        series: [{
          data: [],
          type: 'line'
        }]
      },
      mode: "Percent",
      output: 0,
      running: false
    };

    this.changeMode = this.changeMode.bind(this);
    this.changeOutput = this.changeOutput.bind(this);

    this.run = this.run.bind(this);
    this.stop = this.stop.bind(this);

    this.updateGraph = this.updateGraph.bind(this);

    ipcRenderer.on("enable-heartbeat-response", this.updateGraph);
  }

  componentWillUnmount() {
    if (this.state.running) {
      this.stop();
    }
    ipcRenderer.removeListener("enable-heartbeat-response", this.updateGraph);
  }

  componentDidUpdate(prevProps) {
    if (prevProps.connected !== this.props.connected && this.props.connected === false) {
      if (this.state.running) {
        this.stop();
      }
    }
  }

  updateGraph(event, error, response) {
    // this.state.option.series.data.push(0.1);
    this.forceUpdate();
  }

  run() {
    console.log("Starting heartbeat");
    this.setState({running: true});
    enableHeartbeat(20);
  }

  stop() {
    console.log("Stopping heartbeat");
    this.setState({running: false});
    disableHeartbeat();
  }

  changeMode(value) {
    this.setState({
      option: {
        xAxis: {
          type: 'category',
          name: "Time"
        },
        yAxis: {
          type: 'value',
          name: value.currentTarget.value
        },
        series: [{
          data: [],
          type: 'line'
        }]
      },
      mode: value.currentTarget.value
    });
  }

  changeOutput(value) {
    if (Math.abs(this.state.output - value) > 0.1) {
      if (this.state.output - value < 0) {
        value = this.state.output + 0.1;
      } else {
        value = this.state.output - 0.1;
      }
    }
    this.setState({output: value});
    if (value < 0) {
      setSetpoint(value * 1024);
    } else {
      setSetpoint(value * 1023);
    }
  }

  render() {
    const {option, mode, output, running} = this.state;
    return (
      <div>
        <ReactEcharts
          option={option}
          notMerge={true}
          lazyUpdate={true}
          className="echart-container"
        />
        <div className="form">
          <FormGroup
            labelFor="run-mode"
            className="form-group-quarter inline"
          >
            <RadioGroup
              label="Mode"
              selectedValue={mode}
              onChange={this.changeMode}
              disabled={!this.props.connected}
            >
              <Radio label="Percent" value={"Percent"}/>
              <Radio label="Velocity" value={"Velocity"}/>
              <Radio label="Position" value={"Position"}/>
            </RadioGroup>
          </FormGroup>
          <FormGroup className="form-group-quarter"/>
          <FormGroup
            label="Motor Output"
            className="form-group-half"
          >
            <Slider initialValue={output} disabled={!this.props.connected} value={output} min={-1.0} max={1.0} stepSize={0.01} onChange={this.changeOutput} />
          </FormGroup>
          <FormGroup
            className="form-group-quarter"
          >
            <Button className="rev-btn" disabled={!this.props.connected} onClick={running ? this.stop : this.run}>{running ? "Stop" : "Run"}</Button>
          </FormGroup>
        </div>
      </div>
    );
  }
}

export default RunTab;