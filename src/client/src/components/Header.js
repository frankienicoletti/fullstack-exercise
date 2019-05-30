import * as React from "react";
import styled from "@emotion/styled";

const HeaderWrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  margin: 24px 0;
`;

const HeaderGroup = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
`;

const Heading = styled.div`
  font-size: 22px;
  font-weight: 500;
`;

const Logo = styled.img`
  margin-right: 8px;
`;
Logo.defaultProps = {
  src: "https://assets.usejournal.com/accounts/logo.svg"
};

class Header extends React.Component {
  render() {
    return (
      <HeaderWrapper>
        <HeaderGroup>
          <Logo width={32} height={32} />
          <Heading>Journal</Heading>
        </HeaderGroup>
      </HeaderWrapper>
    );
  }
}

export default Header;
