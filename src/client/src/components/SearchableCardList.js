import * as React from "react";

import CardList from "./CardList";
import SearchInput from "./SearchInput";

class SearchableCardList extends React.PureComponent {
  state = {
    searchInput: ""
  };

  onInputChange = event => {
    this.setState({ searchInput: event.target.value });
  };

  filterCards = (cards, searchInput) => {
    return searchInput.length === 0
      ? cards
      : cards.filter(c => {
          return (
            c.name.toLowerCase().includes(searchInput) ||
            c.colors.map(c => c.toLowerCase()).includes(searchInput)
          );
        });
  };

  render() {
    const { cards } = this.props;

    return (
      <div>
        <SearchInput
          placeholder="Search here..."
          value={this.state.searchInput}
          onChange={this.onInputChange}
        />
        <CardList cards={this.filterCards(cards, this.state.searchInput)} />
      </div>
    );
  }
}

export default SearchableCardList;
