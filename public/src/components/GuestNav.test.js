import React from 'react';
import { shallow } from 'enzyme';
import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
configure({ adapter: new Adapter() });

import GuestNav from './GuestNav';
const props = {};

describe("GuestNav", () => {
  const gNav = shallow(<GuestNav {...props} />);

  it("renders the Action-Network text in h1 tag", () => {
    expect(gNav.find('Link h1').text()).toEqual("Action-Network");
  });
});
