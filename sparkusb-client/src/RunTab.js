import React, {Component} from "react";
import {FormGroup, Slider, Radio, RadioGroup} from "@blueprintjs/core";
import ReactEcharts from "echarts-for-react";

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
          name: "Voltage"
        },
        series: [{
          data: [0.1, 0.25, 0.33, 0.5, 0.75, 0.8, 1.0],
          type: 'line'
        }]
      },
      mode: "Percent",
      output: 0,
    };

    this.changeMode = this.changeMode.bind(this);
    this.changeOutput = this.changeOutput.bind(this);
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
          data: [0.1, 0.25, 0.33, 0.5, 0.75, 0.8, 1.0],
          type: 'line'
        }]
      },
      mode: value.currentTarget.value
    });
  }

  changeOutput(value) {
    this.setState({output: value});
  }

  render() {
    const {option, mode, output} = this.state;
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
            <Slider disabled={!this.props.connected} initialValue={output} value={output} min={0} max={1.0} stepSize={0.01} onChange={this.changeOutput} />
          </FormGroup>
        </div>
      </div>
    );
  }
}

export default RunTab;