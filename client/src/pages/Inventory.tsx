import React from 'react';
import styled from 'styled-components';
import InventoryLists from './InventoryLists';

const InventoryLabel = styled.h1`
  position: absolute;
  left: 8rem;
  top: 1rem;

  font-size: 20px;
`;

function Inventory() {
    
    return(
        <>
            <InventoryLabel>Inventory</InventoryLabel>
            <InventoryLists />
        </>
    );

}

export default Inventory;