import React, { Fragment, useState, useEffect } from "react";
import { Redirect } from "react-router-dom";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { login } from "../redux/actions/@auth";

const Login = ({ isAuthenticated, login, message }) => {
  const [formData, setFormData] = useState({
    email: "",
    password: ""
  });
  const [warn, setWarn] = useState("");
  const [spinner, setSpinner] = useState("");

  const { email, password } = formData;

  const onChange = e => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const onSubmit = e => {
    e.preventDefault();
    if (email === "" || password === "")
      return [
        setWarn("email or password cannot be empty"),
        setSpinner(false),
        setTimeout(() => {
          setWarn("");
        }, 3000)
      ];
    else if (true) {
      return [
        setWarn(""),
        setSpinner(true),
        login(formData),
        setTimeout(() => {
          setSpinner(false);
        }, 3000)
      ];
    }
  };

  //redirect is already logged in
  // if (isAuthenticated) {
  //   return <Redirect to='/dashboard' />;
  // }
  return (
    <Fragment>
      <form className='login' onSubmit={e => onSubmit(e)}>
        <input
          name='email'
          type='text'
          value={email}
          onChange={e => onChange(e)}
          placeholder='Enter Email'
        />
        <input
          name='password'
          type='password'
          value={password}
          onChange={e => onChange(e)}
          placeholder='******'
        />
        <button
          className={spinner ? "spinner" : ""}
          type='submit'
          value='Login'
        >
          Login
        </button>
        <div className='warn'>{warn ? warn : "" || message ? message : ""}</div>{" "}
      </form>
    </Fragment>
  );
};

Login.propTypes = {
  isAuthenticated: PropTypes.bool.isRequired,
  message: PropTypes.object.isRequired,
  login: PropTypes.func.isRequired
};

const mapStateToProps = state => {
  return {
    isAuthenticated: state.auth.isAuthenticated,
    message: state.auth.message ? state.auth.message.error : ""
  };
};
export default connect(mapStateToProps, { login })(Login);
