import React from "react";
import {MenuItem} from "@blueprintjs/core";

export const MotorTypes = [
  {motorID: 1, name: "REV Brushless", kMotorType: 1},
  {motorID: 2, name: "CIM", kMotorType: 0},
  {motorID: 3, name: "Mini CIM", kMotorType: 0},
  {motorID: 4, name: "775 Pro/Redline", kMotorType: 0},
  {motorID: 5, name: "Bag", kMotorType: 0},
  {motorID: 6, name: "REV HD Hex", kMotorType: 0}
];

export const renderMotors = (motorType, itemProps) => {
  return (
    <MenuItem
      active={itemProps.modifiers.active}
      key={motorType.motorID}
      onClick={itemProps.handleClick}
      text={motorType.name}
    />
  );
};