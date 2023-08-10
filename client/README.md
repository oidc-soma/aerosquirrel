# Aerosquirrel Front-end Program for Client

이 애플리케이션은 React.js및 Typescript, 일부 Javascript로 작성되었습니다.

사용한 라이브러리는 다음과 같습니다.
* React.js
* react-router-dom
* Recoil
* Tldraw(@tldraw/core, @tldraw/tldraw)
* yorkie-js-sdk
* randomcolor
* react-toastify
* styled-components

## 실행 방법

클라이언트 프로그램을 실행하려면 client 폴더에서 yarn start 명령어를 실행해 주세요.
</br> yarn이 설치되지 않은 경우 npm install을 먼저 실행해 주세요.
### `yarn start`

기본적으로 3000 포트에서 실행되며 웹브라우저 'http://localhost:3000' 페이지에서 확인할 수 있습니다.

### `npx cypress open`

E2E 테스트를 실행하려면 위의 명령어로 Cypress를 실행한 후 브라우저를 이용하여 테스트를 실행해 주세요.
* signup.cy.ts: 회원가입 정상 진행 여부 테스트
* login.cy.ts: 로그인 정상 진행 여부 테스트
* accounttest.cy.ts: Accounts 화면이 정상적으로 표시되는지 여부 테스트
* drawer.cy.ts: Tldraw를 이용한 Drawer 화면 정상 표시 여부 테스트
* inventorytest.cy.ts: 클라우드 인벤토리 목록 화면 정상 표시 여부 테스트

## 각 폴더별 기능

### components
프로그램에서 반복적으로 사용되는 구성요소가 포함된 폴더

* cards: 클라우드 제공자가 포함된 카드 화면 구현(Accounts 화면)
* navbar: Navbar.tsx가 포함된 폴더로 좌측 내비게이션바 컴포넌트 구현
* yorkie-tldraw: Yorkie 및 Tldraw 기능을 활용한 Drawer 기능을 위한 컴포넌트 구현

### hooks

* useMultiplayerState.ts: Yorkie와 Tldraw의 실시간 연동을 위한 컴포넌트로 Yorkie 서버와 요청 응답 및 Tldraw 컴포넌트와의 송/수신 모듈 구현

### Pages

화면에 표시되는 웹페이지의 UI가 구현된 부분입니다.

* Accounts.tsx: 계정 관리 창
* AddAccounts.tsx: 각각의 클라우드 계정의 token값 및 Yorkie의 연동을 위한 token 값을 저장할 수 있는 화면
* AddInstancePrompt.tsx: Tldraw에 인스턴스 도식을 추가하기 위해 현재 인벤토리에 존재하는 인스턴스를 선택하는 화면
* Dashboard.tsx: 현재 클라우드의 비용 및 Metric을 확인할 수 있는 화면
* Drawer.tsx: Tldraw 그리기 화면을 포함하는 화면(component/yorkie-tldraw/Editor.tsx를 import합니다)
* Inventory.tsx: 현재 클라우드 인벤토리에 있는 인스턴스를 목록 Table 형태로 표시합니다.
* Login.tsx, Signup.tsx: 로그인 및 회원가입 관련 화면
* WelcomePage.tsx: 가장 첫 화면으로 서비스 소개 및 홍보 화면

# 프로젝트 기여

프로젝트 기여 및 라이선스 가이드는 Aerosquirrel 전체 프로젝트 가이드에 따릅니다.

[기여 가이드](https://github.com/oidc-soma/aerosquirrel/blob/main/CONTRIBUTING.md)