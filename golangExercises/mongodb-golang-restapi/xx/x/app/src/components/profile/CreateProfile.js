import React, { Fragment, useState } from "react";
import { withRouter } from "react-router-dom";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { createProfile } from "../../redux/actions/@profile";

const CreateProfile = ({ createProfile, history }) => {
  const [profileData, setProfileData] = useState({
    role: "",
    course: "",
    universityName: "",
    hobbies: "",
    facebook: "",
    instagram: ""
  });
  let {
    role,
    course,
    universityName,
    hobbies,
    facebook,
    instagram
  } = profileData;

  const onChange = e => {
    setProfileData({ ...profileData, [e.target.name]: e.target.value });
  };
  const onSubmit = e => {
    e.preventDefault();
    createProfile(profileData, history);
    console.log("from create profile", profileData);
  };

  return (
    <div className='create-prof-wrapper'>
      <form className='create-profile' onSubmit={e => onSubmit(e)}>
        <Fragment>
          <select name='role' value={role} onChange={e => onChange(e)}>
            <option value='0'>Select Role</option>
            <option value='Student'>Student</option>
            <option value='Reacher'>Teacher</option>
          </select>

          <select name='course' value={course} onChange={e => onChange(e)}>
            <option value='0'>choose course</option>
            <option value='B.A'>B.A</option>
            <option value='BSC'>BSC</option>
          </select>
        </Fragment>
        <input
          type='text'
          name='universityName'
          value={universityName}
          onChange={e => onChange(e)}
          placeholder='Enter university name'
        />
        <input
          type='text'
          name='hobbies'
          value={hobbies}
          onChange={e => onChange(e)}
          placeholder='Do you have hobbies'
        />
        <div>
          <h4>Social network</h4>
          <input
            type='text'
            name='facebook'
            value={facebook}
            onChange={e => onChange(e)}
            placeholder='facebook link'
          />
          <input
            type='text'
            name='instagram'
            value={instagram}
            onChange={e => onChange(e)}
            placeholder='instagram link'
          />
        </div>
        <button type='submit'>submit</button>
      </form>
    </div>
  );
};

CreateProfile.propTypes = {};

const mapStateToProps = state => {
  console.log("this is state", state.profile);
  return { profile: state.profile };
};
export default connect(mapStateToProps, { createProfile })(
  withRouter(CreateProfile)
);
