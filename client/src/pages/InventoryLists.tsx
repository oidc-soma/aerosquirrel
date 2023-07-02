import React from 'react';
import './InventoryLists.css';

function InventoryLists() {

    return (
      <table
        id="InventoryTable"
        className="table-striped table-bordered"
        cellSpacing="0"
        width="100%"
      >
        <thead>
          <tr>
            <th className="th-sm">Cloud</th>
            <th className="th-sm">Service</th>
            <th className="th-sm">Name</th>
            <th className="th-sm">Region</th>
            <th className="th-sm">Account</th>
            <th className="th-sm">Cost</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>AWS</td>
            <td>VPC</td>
            <td>vpc-0b6a762e556d8ee14</td>
            <td>ap-southeast-2</td>
            <td>Sandbox</td>
            <td>$19.94</td>
          </tr>
        </tbody>
        <tfoot></tfoot>
      </table>
    );
}
export default InventoryLists;