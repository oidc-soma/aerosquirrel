import React from "react";
import CloudCompMenu from "./CloudCompMenu";
import "./CloudProvider.css";

function CloudProvider() {
  return (
    <>
      <ul className="list-group list-group-light cloudprovider">
        <CloudCompMenu TypeText="Amazon Web Services" />
        <CloudCompMenu TypeText="Oracle Cloud Infrastructure" />
        <CloudCompMenu TypeText="Yorkie(Draw Panel)" />
      </ul>
    </>
  );
}

export default CloudProvider;
