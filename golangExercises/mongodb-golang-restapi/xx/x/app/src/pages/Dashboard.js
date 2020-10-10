import React, { Fragment, useEffect } from "react";
import { NavLink } from "react-router-dom";
import { connect } from "react-redux";
import { getPosts } from "../redux/actions/@posts";
import PropTypes from "prop-types";
import { getCurrentProfile, deleteAccount } from "../redux/actions/@profile";
import { removePost } from "../redux/actions/@posts";
import { loadUser } from "../redux/actions/@auth";

const Dashboard = ({
  getPosts,
  loadUser,
  deleteAccount,
  auth: { isAutheticated },
  profile: { profile, isLoading },
  getCurrentProfile,
  removePost,
  user,
  posts: { posts }
}) => {
  useEffect(() => {
    getCurrentProfile();
  }, [isAutheticated]);

  useEffect(() => {
    getPosts();
  }, []);

  useEffect(() => {
    loadUser();
  }, []);

  const userPostList = user ? posts.filter(post => post.user === user._id) : "";
  const renderPostsList =
    user && posts ? (
      userPostList.map(post => (
        <div className='dash-p-lst' key={post._id}>
          <span>{post.title}</span>
          <button onClick={e => removePost(post._id)}>delete</button>
        </div>
      ))
    ) : (
      <div>loading</div>
    );
  if (!profile) {
    return (
      <Fragment>
        <div className='no-profile'>
          <p>You have not created a profile yet , go ahead and create one</p>
          <NavLink className='dash-navlink' to='/create-profile'>
            create Profile
          </NavLink>
        </div>
      </Fragment>
    );
  }
  return isLoading === true && profile === null ? (
    <div className='loading'>
      <div>loading</div>
      <div></div>
      <div></div>
    </div>
  ) : (
    <div className='dashboard'>
      <div className='prop-list'>
        <div className='prop-val'>
          <div>Name :</div> <div>{profile.universityName}</div>
        </div>
        <div className='prop-val'>
          <div>Role :</div> <div>{profile.role}</div>
        </div>
        <div className='prop-val'>
          <div>Course :</div> <div>{profile.course}</div>
        </div>
        <div className='prop-val'>
          <div>UniversityName :</div> <div>{profile.universityName}</div>
        </div>
        <div className='prop-val'>
          <div>Hobbies :</div> <div>{profile.hobbies.toString()}</div>
        </div>
        <div className='prop-val'>
          <div>facebook :</div>{" "}
          <div>
            <a href={!profile.social ? "" : profile.social.facebook}>
              {!profile.social ? "" : profile.social.facebook}
            </a>
          </div>
        </div>
        <div className='prop-val'>
          <div>instagram :</div>{" "}
          <div>
            <a href={!profile.social ? "" : profile.social.instagram}>
              {!profile.social ? "" : profile.social.instagram}
            </a>
          </div>
        </div>
      </div>
      <div className='dash-btns'>
        <NavLink className='dash-nav' to='/edit-profile'>
          Edit profile
        </NavLink>
        <button onClick={() => deleteAccount()}>Delete Profile</button>
      </div>
      <div className='post-list'>{renderPostsList}</div>
    </div>
  );
};

Dashboard.propTypes = {
  getCurrentProfile: PropTypes.func.isRequired,
  auth: PropTypes.object.isRequired,
  profile: PropTypes.object.isRequired,
  user: PropTypes.object.isRequired,
  posts: PropTypes.object.isRequired,
  getPosts: PropTypes.func.isRequired,
  removePost: PropTypes.func.isRequired,
  deleteAccount: PropTypes.func.isRequired
};

const mapStateToProps = state => {
  return {
    auth: state.auth,
    profile: state.profile,
    user: state.auth.user,
    posts: state.posts
  };
};
export default connect(mapStateToProps, {
  loadUser,
  getCurrentProfile,
  getPosts,
  removePost,
  deleteAccount
})(Dashboard);
