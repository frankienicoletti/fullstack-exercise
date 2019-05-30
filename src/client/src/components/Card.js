import * as React from "react";
import styled from "@emotion/styled";

const CardWrapper = styled.div`
  border: 0.5px solid #f0eded;
  border-radius: 4px;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.12);

  padding: 8px;
  margin: 8px;

  display: flex;
  flex-direction: row;
`;
const CardIcon = styled.img``;
const Title = styled.div``;

class Card extends React.Component {
  render() {
    const { card } = this.props;

    return (
      <CardWrapper>
        <CardIcon src={card.imgUrl} />
        <Title>{card.name}</Title>
      </CardWrapper>
    );
  }
}

export default Card;
