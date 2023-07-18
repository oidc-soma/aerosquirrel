import React from 'react';
import { Link, Route } from 'react-router-dom';
import styled from "styled-components";
import { useNavigate } from 'react-router-dom';
import HomeWhite from '../../home-white.png';

 const NavInner = styled.nav`
   position: absolute;
   background-color: #262626;
   width: 88px;
   height: 100vh;
   overflow: hidden;
   overscroll-behavior: contain;
   min-height: 100vh;
   padding-top: 37px;
   z-index: 100;
 `;
 

function Navbar() {
    // Global App Context 정의 예정

  const navigate = useNavigate();


return (
  <NavInner>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      width="24"
      height="24"
      onClick={() => {
        navigate("/dashboard");
      }}
    >
      <path
        d="M12 18V15"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M10.07 2.81997L3.14002 8.36997C2.36002 8.98997 1.86002 10.3 2.03002 11.28L3.36002 19.24C3.60002 20.66 4.96002 21.81 6.40002 21.81H17.6C19.03 21.81 20.4 20.65 20.64 19.24L21.97 11.28C22.13 10.3 21.63 8.98997 20.86 8.36997L13.93 2.82997C12.86 1.96997 11.13 1.96997 10.07 2.81997Z"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
    </svg>
    <br />
    <br />

    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      width="24"
      height="24"
      onClick={() => {
        navigate("/accounts");
      }}
    >
      <path
        d="M5.54003 11.12C0.860029 11.45 0.860029 18.26 5.54003 18.59H7.46007"
        stroke="white"
        stroke-width="1.5"
        stroke-miterlimit="10"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M5.59003 11.12C2.38003 2.19002 15.92 -1.37998 17.47 8.00002C21.8 8.55002 23.55 14.32 20.27 17.19C19.27 18.1 17.98 18.6 16.63 18.59H16.54"
        stroke="white"
        stroke-width="1.5"
        stroke-miterlimit="10"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M17 16.53C17 17.27 16.84 17.97 16.54 18.59C16.46 18.77 16.37 18.94 16.27 19.1C15.41 20.55 13.82 21.53 12 21.53C10.18 21.53 8.58998 20.55 7.72998 19.1C7.62998 18.94 7.54002 18.77 7.46002 18.59C7.16002 17.97 7 17.27 7 16.53C7 13.77 9.24 11.53 12 11.53C14.76 11.53 17 13.77 17 16.53Z"
        stroke="white"
        stroke-width="1.5"
        stroke-miterlimit="10"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M10.4399 16.53L11.4299 17.5201L13.5599 15.55"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
    </svg>
    <br />
    <br />
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      width="24"
      height="24"
      onClick={() => {
        navigate("/inventory");
      }}
    >
      <path
        d="M3.17004 7.43994L12 12.5499L20.77 7.46991"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M12 21.6099V12.5399"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M9.92999 2.48L4.59 5.45003C3.38 6.12003 2.39001 7.80001 2.39001 9.18001V14.83C2.39001 16.21 3.38 17.89 4.59 18.56L9.92999 21.53C11.07 22.16 12.94 22.16 14.08 21.53L19.42 18.56C20.63 17.89 21.62 16.21 21.62 14.83V9.18001C21.62 7.80001 20.63 6.12003 19.42 5.45003L14.08 2.48C12.93 1.84 11.07 1.84 9.92999 2.48Z"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
      <path
        d="M17 13.24V9.58002L7.51001 4.09998"
        stroke="white"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      ></path>
    </svg>
    <br />
    <br />
    <svg
      fill="#FFFFFF"
      height="30px"
      width="30px"
      version="1.1"
      id="Layer_1"
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      viewBox="0 0 460.002 460.002"
      xmlSpace="preserve"
      onClick={() => {
        navigate("/drawer");
      }}
    >
      <g>
        <g>
          <g>
            <path
              d="M367.436,89.7H92.565c-1.58,0-2.865,1.285-2.865,2.865v274.871c0,1.58,1.285,2.865,2.865,2.865H290.29v-65.01
				c0-8.284,6.716-15,15-15h65.011V92.565C370.301,90.986,369.016,89.7,367.436,89.7z"
            />
            <polygon points="320.291,349.088 349.088,320.291 320.291,320.291 			" />
            {/* <path
              d="M427.137,0H32.865C14.743,0,0,14.743,0,32.865v394.272c0,18.122,14.743,32.865,32.865,32.865h394.272
				c18.122,0,32.865-14.743,32.865-32.865V32.865C460.002,14.743,445.259,0,427.137,0z M400.301,305.291
				c0,4.143-1.68,7.894-4.396,10.609L315.9,395.905c-2.715,2.716-6.466,4.396-10.609,4.396H92.565
				c-18.122,0-32.865-14.743-32.865-32.865V92.565C59.7,74.443,74.443,59.7,92.565,59.7h274.871
				c18.122,0,32.865,14.743,32.865,32.865V305.291z"
            /> */}
          </g>
        </g>
      </g>
    </svg>
  </NavInner>

  // <nav className="fixed bottom-0 left-0 top-0 z-20 flex w-[88px] flex-col gap-6 bg-purplin-850 px-5 py-8 dark:bg-black-900 pt-16">
  //     <div className="flex items-center gap-8 text-sm font-semibold text-black-400" rel="noopener noreferrer">
  //        <div>
  //         링크 로고 이미지 삽입예정
  //        </div>
  //     </div>
  // </nav>
);
}

export default Navbar;


