import React from 'react';
import { shallow } from 'enzyme';
import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
configure({ adapter: new Adapter() });

import { Login } from './Login';

const props = {
  username: "Alex",
  password: "12345678"
}


describe("Login Container", () => {
  const loginContainer = shallow(<Login {...props} />);

  it("renders the right amount of label tags", () => {
    expect(loginContainer.find('label').length).toEqual(2);
  });

  it("renders the username label", () => {
    expect(loginContainer.find("label").at(0).text()).toEqual("Username");
  });

  it("renders the password label", () => {
    expect(loginContainer.find("label").at(1).text()).toEqual("Password");
  });

  it("renders a label component with class `login-username-label`", () => {
    expect(loginContainer.find("label").at(0).hasClass("login-username-label")).toBe(true);
  });

  it("renders a label component with class `login-password-label`", () => {
    expect(loginContainer.find("label").at(1).hasClass("login-password-label")).toBe(true);
  });

  it("renders a Login button", () => {
    expect(loginContainer.find("button").text()).toEqual("Login");
  });

  it("renders a button component with class `login-submit-btn`", () => {
    expect(loginContainer.find("button").hasClass("login-submit-btn")).toBe(true);
  });
});
