import React from "react";
import ReactDOM from "react-dom";

import Header from "./components/Header";
import SearchableCardList from "./components/SearchableCardList";

import "./styles.css";

import cards from "./cards.json";

function App() {
  return (
    <div className="App">
      <Header />
      <SearchableCardList cards={cards} />
    </div>
  );
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);
