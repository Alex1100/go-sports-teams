import React,
{
  Component
} from 'react';

import {
  signupUser
} from '../actions/auth';

import {
  inputUsername,
  inputPassword
} from '../actions/login';

import {
  connect
} from 'react-redux';


class Signup extends Component {
  constructor(props) {
    super(props);

    this.handleUsername = this.handleUsername.bind(this);
    this.handlePassword = this.handlePassword.bind(this);
    this.SignupUser = this.SignupUser.bind(this);
  }


  handleUsername(e) {
    this.props.dispatch(inputUsername(e.target.value));
  }

  handlePassword(e) {
    this.props.dispatch(inputPassword(e.target.value));
  }

  SignupUser(){
    const {
      username,
      password,
      history,
      dispatch
    } = this.props;

    if (password.length > 5) {
      dispatch(
        signupUser(
          {
            username,
            password
          },
          history
        )
      );
    } else {
      console.log('password not long enough');
    }
  }

  render() {
    const {
      username,
      password
    } = this.props;

    return (
      <div>
        <div className="login-container">
          <form>
            <label className="login-username-label">
              Username
            </label>
            <input
              className="login-username-input"
              type='text'
              name='username'
              onChange={(e) => {
                e.preventDefault();
                this.handleUsername(e);
              }}
              value={username}
            />
            <br/>
            <label className="login-password-label">
              Password
            </label>
            <input
              className="login-password-input"
              type='password'
              name='password'
              onChange={(e) => {
                e.preventDefault();
                this.handlePassword(e);
              }}
              value={password}
            />
            <br/>
            <button
              className="login-submit-btn"
              onClick={(e) => {
                e.preventDefault();
                this.SignupUser();
              }}>
              Signup
            </button>
          </form>
        </div>
      </div>
    );
  }
}


const mapStateToProps = (state) => {
  const {
    signup
  } = state;

  const {
    username,
    password
  } = signup;

  return {
    username,
    password,
  };
};

export default connect(mapStateToProps)(Signup);
