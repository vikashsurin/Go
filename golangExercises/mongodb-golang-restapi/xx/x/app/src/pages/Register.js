import React, { Fragment, useState, useEffect } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { register } from "../redux/actions/@auth";
import PropTypes from "prop-types";

const Register = ({ isAuthenticated, register, message }) => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
    password2: ""
  });
  const { name, email, password, password2 } = formData;

  const onChange = e => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const [warn, setWarn] = useState("");
  const [spinner, setSpinner] = useState(false);

  const onSubmit = async e => {
    e.preventDefault();
    if (name === "" || email === "")
      return [
        setWarn("name or email field cannot be empty"),
        setTimeout(() => {
          setWarn("");
        }, 3000)
      ];
    else if (formData.password !== formData.password2)
      return [
        setWarn("password must be same"),
        setTimeout(() => {
          setWarn("");
        }, 3000)
      ];
    else if (formData.password.length < 6 || formData.password2.length < 6)
      return [
        setWarn("password atleast 6 cahracters long"),
        setTimeout(() => {
          setWarn("");
        }, 3000)
      ];
    else
      return [
        setSpinner(true),
        register(formData),
        setTimeout(() => {
          setSpinner(false);
        }, 3000)
      ];
  };

  //redirect is already logged in
  if (isAuthenticated) {
    return <Redirect to='/dashboard' />;
  }

  return (
    <Fragment>
      <form className='register ' onSubmit={e => onSubmit(e)}>
        <input
          name='name'
          type='text'
          value={name}
          onChange={e => onChange(e)}
          placeholder='Enter Name'
        />
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
          placeholder='   ******'
        />
        <input
          name='password2'
          type='password'
          value={password2}
          onChange={e => onChange(e)}
          placeholder='  ******'
        />
        <button
          className={spinner ? "spinner" : ""}
          type='submit'
          value='Register'
        >
          Register
        </button>
        <div className='warn'>{warn ? warn : "" || message ? message : ""}</div>
      </form>
    </Fragment>
  );
};

Register.propTypes = {
  isAuthenticated: PropTypes.bool.isRequired,
  message: PropTypes.object.isRequired,
  register: PropTypes.func.isRequired
};

const mapStateToProps = state => {
  return {
    isAuthenticated: state.auth.isAuthenticated,
    message: state.auth.message ? state.auth.message.error : ""
  };
};
export default connect(mapStateToProps, { register })(Register);
