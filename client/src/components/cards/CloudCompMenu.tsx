import React from 'react';
import AWSIcon from '../../aws-icon.png';
import OracleIcon from '../../oracle-icon.png';
import OnPremiseIcon from '../../on-premise.png';

type CloudCompProps = {
 TypeText: string;
};


function CloudCompMenu({TypeText}:CloudCompProps) {
    if(TypeText==="Amazon Web Services")
    {
        return (
          <li className="list-group-item d-flex justify-content-between align-items-center">
            <div className="d-flex align-items-center">
                <img src={AWSIcon} width={40}></img>
              <div className="ms-3">
                <p className="fw-bold mb-1">{TypeText}</p>
              </div>
            </div>
            {/* <span className="badge rounded-pill badge-success">Continue</span> */}
          </li>
        );
    }
    else if(TypeText==="Oracle Cloud Infrastructure")
    { return (
      <li className="list-group-item d-flex justify-content-between align-items-center">
        <div className="d-flex align-items-center">
          <img src={OracleIcon} width={40}></img>
          <div className="ms-3">
            <p className="fw-bold mb-1">{TypeText}</p>
          </div>
        </div>
        {/* <span className="badge rounded-pill badge-success">Continue</span> */}
      </li>
    );
        
    }
    else
    {
         return (
           <li className="list-group-item d-flex justify-content-between align-items-center">
             <div className="d-flex align-items-center">
               <img src={OnPremiseIcon} width={40}></img>
               <div className="ms-3">
                 <p className="fw-bold mb-1">{TypeText}</p>
               </div>
             </div>
             {/* <span className="badge rounded-pill badge-success">Continue</span> */}
           </li>
         );
    }

    
}

export default CloudCompMenu;
