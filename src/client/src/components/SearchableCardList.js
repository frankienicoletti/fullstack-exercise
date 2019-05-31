import * as React from "react";

import CardList from "./CardList";
import SearchInput from "./SearchInput";

import API from '../utils/API';

class SearchableCardList extends React.PureComponent {
  state = {
    searchInput: "",
    cards: [],
    apiError: false
  };

  async getPeople(searchInput) {
    try {
      const response = await API.get('/people', {
        params: {
          q: searchInput
        }
      });
      this.setState({
        cards: response.data,
        apiError: false
      });
    } catch (e) {
      this.setState({
        apiError: true
      });
    }
  }

  async componentDidMount() {
    await this.getPeople(this.state.searchInput);
  }

  async componentDidUpdate(prevProps, prevState) {
    if (this.state.searchInput !== prevState.searchInput) {
      await this.getPeople(this.state.searchInput);
    }
  }

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
    return (
      <div>
        <SearchInput
          placeholder="Search here..."
          value={this.state.searchInput}
          onChange={this.onInputChange}
        />
        {
          this.state.apiError
            ? <div>Failed to connect to backend</div>
            : <CardList cards={this.state.cards} />
        }
      </div>
    );
  }
}

export default SearchableCardList;
