import React from 'react';

import { mount } from 'enzyme';

import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';

configure({ adapter: new Adapter() });


import LandingPage from './LandingPage';
const props = {};

describe("LandingPage", () => {
  const landingPage = mount(<LandingPage {...props} />);

  it("renders `Log in to keep track of your favorite teams` in h1", () => {
    expect(landingPage.find('h1').text()).toEqual("Log in to keep track of your favorite teams");
  });
});
