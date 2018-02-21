import React from 'react';
import { shallow } from 'enzyme';
import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
configure({ adapter: new Adapter() });

import { SearchContainer } from './SearchContainer';
import {
  propsNBA,
  propsNFL,
  propsNHL,
  propsNBANoResultSearch
} from "../data/fixtures";


describe("SearchContainer Component", () => {
  const nbaSearchContainer = shallow(<SearchContainer {...propsNBA} />);
  const nflSearchContainer = shallow(<SearchContainer {...propsNFL} />);
  const nhlSearchContainer = shallow(<SearchContainer {...propsNHL} />);
  const nbaSearchContainerNoResult = shallow(<SearchContainer {...propsNBANoResultSearch} />);

  describe("tests the Tab Component", () => {
    it("renders tabs components", () => {
      expect(nbaSearchContainer.find("Tab").length).toEqual(3);
    });

    it("renders the NBA Tab", () => {
      expect(nflSearchContainer.find("Tab").at(0).props().children).toEqual("NBA");
    });

    it("tests that the Default Active Tab is NBA when selectedSport === `NBA`", () => {
      expect(nhlSearchContainer.props().children[0].props.children.props.defaultActiveKey).toEqual("NBA");
    });

    it("renders the NFL Tab", () => {
      expect(nflSearchContainer.find("Tab").at(1).props().children).toEqual("NFL");
    });

    it("renders the NHL Tab", () => {
      expect(nhlSearchContainer.find("Tab").at(2).props().children).toEqual("NHL");
    });
  });

  describe("tests the search input field", () => {
    it("tests for `search-input-container` class existence", () => {
      expect(nbaSearchContainer.find(".search-input-container").exists()).toBe(true);
    });

    it("tests for the initial value within the input field", () => {
      expect(nbaSearchContainer.find(".search-input-container").props().children[0].props.value).toEqual(propsNBA.term);
      expect(nflSearchContainer.find(".search-input-container").props().children[0].props.value).toEqual(propsNFL.term);
      expect(nhlSearchContainer.find(".search-input-container").props().children[0].props.value).toEqual(propsNHL.term);
    });
  });

  describe("tests the presence of `<span className='glyphicon glyphicon-search'`>", () => {
    it("renders span with className `glyphicon glyphicon-search`", () => {
      expect(nhlSearchContainer.find(".glyphicon.glyphicon-search").exists()).toBe(true);
    });
  });

  describe("tests the contents of the `searched-teams-container`", () => {
    it("renders the correct amount of NBA teams inside `<div className='team-info'>`", () => {
      expect(nbaSearchContainer.find(".team-info").length).toEqual(propsNBA.teams.length);
    });

    it("renders the correct amount of NFL teams inside `<div className='team-info'>`", () => {
      expect(nflSearchContainer.find(".team-info").length).toEqual(propsNFL.teams.length);
    });

    it("renders the correct amount of NHL teams inside `<div className='team-info'>`", () => {
      expect(nhlSearchContainer.find(".team-info").length).toEqual(propsNHL.teams.length);
    });
  });

  describe("tests for no result search", () => {
    it("renders the `<div className='no-results-container'> div`", () => {
      expect(nbaSearchContainerNoResult.find(".no-results-container").exists()).toBe(true);
    });

    it("renders message inside `div.no-results-text`", () => {
      expect(nbaSearchContainerNoResult.find(".no-results-text").text()).toEqual("No Result");
    });

    it("renders the seahawks meme img", () => {
      expect(nbaSearchContainerNoResult.find(".no-results-img").exists()).toBe(true);
      expect(nbaSearchContainerNoResult.find(".no-results-img").props().src).toEqual("http://i.imgur.com/pFQZf.jpg");
    });
  });

  describe("tests the presence of favorited star", () => {
    expect(nbaSearchContainer.find(".star-favorited").exists()).toBe(true);
    expect(nbaSearchContainer.find(".star-unfavorited").exists()).toBe(true);
  });

  describe("tests the presence of unfavorited star", () => {
    expect(nhlSearchContainer.find(".star-unfavorited").exists()).toBe(true);
    expect(nhlSearchContainer.find(".star-favorited").exists()).toBe(false);
  });
});





