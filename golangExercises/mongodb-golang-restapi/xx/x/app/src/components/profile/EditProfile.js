import React, { Fragment, useState, useEffect } from "react";
import { withRouter } from "react-router-dom";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { updateProfile } from "../../redux/actions/@profile";
import { getCurrentProfile } from "../../redux/actions/@profile";

const EditProfile = ({
  getCurrentProfile,
  updateProfile,
  profile: { profile, isLoading },
  history
}) => {
  const [profileData, setProfileData] = useState({
    role: "",
    course: "",
    universityName: "",
    hobbies: "",
    facebook: "",
    instagram: ""
  });
  const {
    role,
    course,
    universityName,
    hobbies,
    facebook,
    instagram
  } = profileData;

  useEffect(() => {
    console.log("hello there");
    getCurrentProfile();
    setProfileData({
      role: !profile.role ? "" : profile.role,
      course: !profile.course ? "" : profile.course,
      universityName: !profile.universityName ? "" : profile.universityName,
      hobbies: !profile.hobbies ? "" : profile.hobbies.toString(),
      facebook: !profile.social ? "" : profile.social.facebook,
      instagram: !profile.social ? "" : profile.social.instagram
    });
  }, [isLoading, getCurrentProfile]);

  const onChange = e => {
    console.log("changed");
    setProfileData({ ...profileData, [e.target.name]: e.target.value });
  };
  const onSubmit = async e => {
    e.preventDefault();
    updateProfile(profileData, history);
    console.log("from update profile", profileData);
  };

  return (
    <div className='edit-profile-wrapper'>
      <form className='edit-profile' onSubmit={e => onSubmit(e)}>
        <select name='role' value={role} onChange={e => onChange(e)}>
          <option value='0'>Select Role</option>
          <option value='Student'>Student</option>
          <option value='Teacher'>Teacher</option>
        </select>

        <select name='course' value={course} onChange={e => onChange(e)}>
          <option value='0'>Choose Course</option>
          <option value='B.A'>B.A</option>
          <option value='BSC'>BSC</option>
        </select>
        <input
          className='inp-uname'
          type='text'
          name='universityName'
          value={universityName}
          onChange={e => onChange(e)}
          placeholder='Enter university name'
        />
        <input
          className='inp-hob'
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

EditProfile.propTypes = {
  getCurrentProfile: PropTypes.func.isRequired,
  updateProfile: PropTypes.func.isRequired,
  profile: PropTypes.object.isRequired
};

const mapStateToProps = state => {
  console.log("this is state", state.profile);
  return { profile: state.profile };
};
export default connect(mapStateToProps, { updateProfile, getCurrentProfile })(
  withRouter(EditProfile)
);
