import React from 'react';
import { shallow } from 'enzyme';
import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
configure({ adapter: new Adapter() });

import { Signup } from './Signup';

const props = {
  username: "Alex",
  password: "12345678"
}


describe("Signup Container", () => {
  const signupContainer = shallow(<Signup {...props} />);

  it("renders the right amount of label tags", () => {
    expect(signupContainer.find('label').length).toEqual(2);
  });

  it("renders the username label", () => {
    expect(signupContainer.find("label").at(0).text()).toEqual("Username");
  });

  it("renders the password label", () => {
    expect(signupContainer.find("label").at(1).text()).toEqual("Password");
  });

  it("renders a label component with class `login-username-label`", () => {
    expect(signupContainer.find("label").at(0).hasClass("login-username-label")).toBe(true);
  });

  it("renders a label component with class `login-password-label`", () => {
    expect(signupContainer.find("label").at(1).hasClass("login-password-label")).toBe(true);
  });

  it("renders a Login button", () => {
    expect(signupContainer.find("button").text()).toEqual("Signup");
  });

  it("renders a button component with class `login-submit-btn`", () => {
    expect(signupContainer.find("button").hasClass("login-submit-btn")).toBe(true);
  });
});
