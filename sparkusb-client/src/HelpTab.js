import React, {Component} from "react";
import {Button, Dialog} from "@blueprintjs/core";

class HelpTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      viewingLogs: false
    };
    this.viewLogs = this.viewLogs.bind(this);
    this.unviewLogs = this.unviewLogs.bind(this);
  }

  viewLogs() {
    this.setState({viewingLogs: true});
  }

  unviewLogs() {
    this.setState({viewingLogs: false});
  }

  render() {
    const {viewingLogs} = this.state;
    return (
      <div>
        <Dialog
          isOpen={viewingLogs}
          onClose={this.unviewLogs}
        >
          <div className="bp3-dialog-header">
            <h4 className="bp3-heading">Application Logs</h4>
          </div>
          <div className="bp3-dialog-body">
            {this.props.logs.length === 0 && <span><i>There are currently no application logs.</i></span>}
            {this.props.logs.map(log => {
              return <p key={log}>{log}</p>
            })}
          </div>
        </Dialog>
        <div id="help-troubleshoot">
          <h2>Troubleshooting</h2>
          <ol>
            <li>Try restarting the program.</li>
            <li>After restarting the program, unplug the usb from the computer and plug it in again.</li>
            <li>Contact <a href="mailto:greg@revrobotics.com">greg@revrobotics.com</a></li>
          </ol>
        </div>
        <div className="form">
          <Button className="rev-btn" onClick={this.viewLogs}>View Application Logs</Button>
        </div>
      </div>
    );
  }
}

export default HelpTab;