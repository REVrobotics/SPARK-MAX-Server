const ipcRenderer = window.require("electron").ipcRenderer;

function getParam(parameter) {
  return new Promise((resolve, reject) => {
    ipcRenderer.once("get-param-response", (event, error, response) => {
      if (error) {
        reject(error);
      } else {
        resolve(response);
      }
    });
    ipcRenderer.send("get-param", parameter);
  });
}

function getCanID() {
  return getParam(0);
}

function getInputType() {
  return getParam(1);
}

function getMotorType() {
  return getParam(2);
}

function getCommutationAdvance() {
  return getParam(3);
}

function getSensorType() {
  return getParam(4);
}

function getControlType() {
  return getParam(5);
}

function getIdleMode() {
  return getParam(6);
}

function getDeadband() {
  return getParam(7);
}

export function getAllParameters() {
  return Promise.all([getCanID(), getInputType(), getMotorType(), getCommutationAdvance(), getSensorType(), getControlType(), getIdleMode(), getDeadband()]);
}