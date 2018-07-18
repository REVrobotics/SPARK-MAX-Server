import React from "react";
import {MenuItem} from "@blueprintjs/core";

export const MotorTypes = [
  {motorID: 1, name: "REV Brushless"},
  {motorID: 2, name: "CIM"},
  {motorID: 3, name: "Mini CIM"},
  {motorID: 4, name: "775 Pro/Redline"},
  {motorID: 5, name: "Bag"},
  {motorID: 6, name: "REV HD Hex"}
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