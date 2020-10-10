import React, { useState } from "react";
import PropTypes from "prop-types";
import Login from "./Login";
import Register from "./Register";

const Auth = props => {
  const [n, setN] = useState(true);
  return (
    <div className='auth'>
      <div className='tog-auth'>
        <span onClick={() => setN(true)}>Login</span>
        <span onClick={() => setN(false)}>Register</span>
      </div>

      {n ? <Login /> : <Register />}
    </div>
  );
};

Auth.propTypes = {};

export default Auth;
