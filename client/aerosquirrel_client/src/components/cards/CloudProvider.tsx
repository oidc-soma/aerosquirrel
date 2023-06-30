import React from 'react';
import CloudCompMenu from './CloudCompMenu';

function CloudProvider() {


    return (
      <>
        <ul className="list-group list-group-light">
         <CloudCompMenu TypeText="Amazon Web Services" />
         <CloudCompMenu TypeText="Oracle Cloud Infrastructure" />
         <CloudCompMenu TypeText="On-Premise" />
        </ul>
      </>
    );
}

export default CloudProvider;