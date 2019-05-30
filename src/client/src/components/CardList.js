import * as React from "react";
import styled from "@emotion/styled";

import Card from "./Card";

const ListWrapper = styled.div`
  margin: 0 auto;
`;

class CardList extends React.Component {
  render() {
    const { cards } = this.props;

    return (
      <ListWrapper>
        {cards.map(card => (
          <Card card={card} />
        ))}
      </ListWrapper>
    );
  }
}

export default CardList;
