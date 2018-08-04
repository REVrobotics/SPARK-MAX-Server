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

function saveParam(parameter, value) {
  return new Promise((resolve, reject) => {
    ipcRenderer.once("set-param-response", (event, error, response) => {
      if (error) {
        reject(error);
      } else {
        resolve(response);
      }
    });
    ipcRenderer.send("set-param", parameter, value);
  });
}

function saveConfig(device) {
  return new Promise((resolve, reject) => {
    ipcRenderer.once("save-config-response", (event, error, response) => {
      if (error) {
        reject(error);
      } else {
        resolve(response);
      }
    });
    ipcRenderer.send("save-config", device);
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

// IMPORTANT - The setpoint MUST come in a [-1023, 1024] range!
export function setSetpoint(setpoint) {
  ipcRenderer.send("set-setpoint", setpoint);
}

export function enableHeartbeat(interval) {
  ipcRenderer.send("enable-heartbeat", interval);
}

export function disableHeartbeat() {
  ipcRenderer.once("disable-heartbeat-response", (event, error, response) => {
    console.log(error, response);
  });
  ipcRenderer.send("disable-heartbeat");
}

export function getAllParameters() {
  return Promise.all([getCanID(), getInputType(), getMotorType(), getCommutationAdvance(), getSensorType(), getControlType(), getIdleMode(), getDeadband()]);
}

export function getAllParametersSequential() {
  return new Promise((resolve, reject) => {
    getCanID().then((canID) => {
      getInputType().then((inputType) => {
        getMotorType().then((motorType) => {
          getCommutationAdvance().then((advance) => {
            getSensorType().then((sensorType) => {
              getControlType().then((controlType) => {
                getIdleMode().then((idleMode) => {
                  getDeadband().then((deadband) => {
                    resolve([canID.value, inputType.value, motorType.value, advance.value, sensorType.value, controlType.value, idleMode.value, deadband.value]);
                  }).catch((error) => reject(error));
                }).catch((error) => reject(error));
              }).catch((error) => reject(error));
            }).catch((error) => reject(error));
          }).catch((error) => reject(error));
        }).catch((error) => reject(error));
      }).catch((error) => reject(error));
    }).catch((error) => reject(error));
  });
}

export function saveBasicConfig(device, canID, motorType, idleMode) {
  return new Promise((resolve, reject) => {
    saveParam(0, canID).then((canResponse) => {
      saveParam(2, motorType).then((motorResponse) => {
        saveParam(6, idleMode).then((idleResponse) => {
          saveConfig(device).then((configResponse) => {
            resolve([configResponse, canResponse, motorResponse, idleResponse]);
          }).catch((error) => reject(error));
        }).catch((error) => reject(error));
      }).catch((error) => reject(error));
    }).catch((error) => reject(error));
  });
}