import React from 'react';

import { shallow } from 'enzyme';

import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';

configure({ adapter: new Adapter() });


import UserNav from './UserNav';
const props = {
  user: {
    username: "Alex"
  }
};


describe("UserNav", () => {
  const userNav = shallow(<UserNav {...props} />);


  it("renders username", () => {
    expect(userNav.find('h3').text()).toEqual(props.user.username + " ");
  });
});

