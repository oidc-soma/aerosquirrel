import React from "react";
import { Link, Route } from "react-router-dom";
import styled from "styled-components";

const AccountsLabel = styled.h1`
position: absolute;
left: 8rem;
top: 1rem;

    font-size: 20px;
`



function Accounts() {
  // Global App Context 정의 예정

  return (
    <AccountsLabel>Accounts</AccountsLabel>
    // <nav className="fixed bottom-0 left-0 top-0 z-20 flex w-[88px] flex-col gap-6 bg-purplin-850 px-5 py-8 dark:bg-black-900 pt-16">
    //     <div className="flex items-center gap-8 text-sm font-semibold text-black-400" rel="noopener noreferrer">
    //        <div>
    //         링크 로고 이미지 삽입예정
    //        </div>
    //     </div>
    // </nav>
  );
}

export default Accounts;
